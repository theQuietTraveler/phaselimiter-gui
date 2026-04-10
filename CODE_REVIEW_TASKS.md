# Kodegjennomgang: foreslåtte oppgaver

## 1) Skrivefeil (typo) – oppgave
**Problem funnet:** I GUI-teksten står det `Same algorithm with bakuage.com/aimastering.com`, som er språklig feil/uklar formulering på engelsk.

**Oppgave:**
- Endre teksten til en tydelig setning, f.eks. `Uses the same algorithm as bakuage.com / aimastering.com`.
- Samtidig gå gjennom resten av faste GUI-strenger for små skrivefeil og inkonsekvent språkbruk (store/små bokstaver, bindestrek osv.).

**Akseptansekriterier:**
- Alle synlige GUI-strenger er grammatisk korrekte og konsistente.
- Endringen er verifisert ved manuell kjøring av appen.

## 2) Bug-fiks – oppgave
**Problem funnet:** Ved dra-og-slipp brukes `url.Parse(line)` og deretter `fileUrl.Path` direkte som input-fil. Dette kan gi feil for filnavn med URL-kodede tegn (f.eks. mellomrom som `%20`) og potensielt feil sti-håndtering på enkelte plattformer.

**Oppgave:**
- Parse URI robust og dekod stien korrekt (`url.PathUnescape` eller tilsvarende trygg løsning).
- Håndter ugyldige/ikke-fil-URI-er eksplisitt med feilmelding i UI i stedet for stille ignorering.
- Legg til defensiv validering før jobben legges i kø.

**Akseptansekriterier:**
- Filer med mellomrom, `%`, `#` og ikke-ASCII-tegn prosesseres korrekt via drag-and-drop.
- Ugyldige URI-er vises som forståelig feilstatus.

## 3) Kommentar-/dokumentasjonsavvik – oppgave
**Problem funnet:** README beskriver en avansert «professional 8-stage mastering pipeline» med SoX-basert kjede, men denne Go-koden kaller i praksis `phase_limiter`-binæren med argumenter og viser ikke den beskrevne kjeden. Dokumentasjonen og implementasjonen er derfor ute av synk.

**Oppgave:**
- Oppdater README til å reflektere faktisk runtime-arkitektur i dette repoet.
- Skill tydelig mellom hva GUI gjør (parameterinnsamling + prosesskjøring) og hva ekstern `phase_limiter` gjør.
- Fjern/omskriv påstander som ikke kan spores til kode i repoet.

**Akseptansekriterier:**
- README stemmer med observerbar kodeflyt i `main.go`/`mastering.go`.
- Teknisk beskrivelse er verifiserbar uten antakelser om ekstern, udokumentert logikk.

## 4) Testforbedring – oppgave
**Problem funnet:** Prosjektet mangler målrettede enhetstester for kritisk parsing- og statuslogikk.

**Oppgave:**
- Introduser testbar funksjon for parsing av progresjonslinjer (regex + float-konvertering) fra prosessoutput.
- Skriv tabelltester for gyldige/ugyldige linjer, grenseverdier og locale-uavhengig parsing.
- Legg til test for statusoverganger (`waiting -> processing -> succeeded/failed`) med mock/stub av kommandoeksekvering.

**Akseptansekriterier:**
- `go test ./...` kjører og dekker parsing + sentrale statusoverganger.
- Feil i progresjonsformat fanges av testene før produksjon.
