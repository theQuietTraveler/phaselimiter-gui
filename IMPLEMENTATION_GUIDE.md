# Implementasjonsguide: stabilitetsforbedringene i PhaseLimiter GUI

Denne guiden viser hvordan du implementerer (og verifiserer) forbedringene som ble gjort for robust parsing, trygg shutdown og bedre feilsynlighet i UI.

## 1) Robust parsing i egen modul

**Mål:** Flytt parsing ut av GUI-kode slik at den kan testes/benchmarkes uten GTK-avhengigheter.

### Steg
1. Opprett `internal/parsing/parsing.go`.
2. Implementer:
   - `ParseDroppedFilePath(line string, goos string) (string, error)`
   - `ExtractProgression(line string) (float64, bool)`
3. I `main.go`: bruk `parsing.ParseDroppedFilePath(...)` i drag-and-drop-løkken.
4. I `mastering.go`: bruk `parsing.ExtractProgression(...)` når stdout leses.

## 2) Bedre feilsynlighet i UI

**Mål:** Brukeren skal se *hvorfor* en jobb feiler.

### Steg
1. Utvid tabellmodell med ekstra kolonne `message`.
2. Oppdater `updateListItem(...)` til å sette `Message`.
3. Når parsing feiler eller output-katalog er ugyldig:
   - lag en `Mastering`-rad med `Status=failed`
   - sett `Message` med konkret feiltekst.

## 3) Stabil runner og trygg avslutning

**Mål:** Unngå heng/prosesslekkasje ved app-lukk.

### Steg
1. Endre `MasteringRunner` til å bruke:
   - `context.Context`
   - `context.CancelFunc`
   - `sync.Once` for idempotent `Terminate()`.
2. Kjør prosessen med `exec.CommandContext(ctx, ...)`.
3. I `Run()`:
   - lytt på `ctx.Done()`
   - `defer close(MasteringUpdate)`.
4. I UI-goroutine:
   - bruk `for m := range masteringRunner.MasteringUpdate`.
5. I `execute(...)`:
   - håndter `scanner.Err()` og rapporter dette som `failed`.

## 4) Defensiv output-katalog-håndtering

**Mål:** Feile tidlig og tydelig ved dårlig output-path.

### Steg
1. Les output-katalog fra inputfelt.
2. `strings.TrimSpace(...)`.
3. Hvis tom: legg inn failed-rad med melding.
4. Kjør `os.MkdirAll(outputDir, 0o755)`:
   - ved feil: failed-rad med melding.

## 5) Tester og benchmark

### Enhetstester
- Legg tester i `internal/parsing/parsing_test.go`:
  - URL-escaped paths
  - unicode paths
  - unsupported scheme
  - invalid URI
  - progression parsing.

### Benchmark
- Legg benchmark i `internal/parsing/parsing_benchmark_test.go`.
- Kjør:

```bash
go test ./internal/parsing
go test ./internal/parsing -bench . -benchmem
```

## 6) Praktisk implementeringsrekkefølge (anbefalt)

1. Parsing-modul + tester
2. Integrasjon i `main.go` og `mastering.go`
3. Runner-shutdown/cancellation
4. UI message-kolonne
5. Output-dir validering
6. Benchmark + dokumentasjon

## 7) Vanlige fallgruver

- **Kopi av `sync.Once`**: bruk pointer receiver på `MasteringRunner`-metoder.
- **Goroutine som aldri stopper**: range over kanal som lukkes i `Run()`.
- **Uforklarlige brukerfeil**: ikke dropp input stille; skriv `Message`.
- **Misvisende testpåstander**: rapporter miljøbegrensninger (GTK/glib) eksplisitt.
