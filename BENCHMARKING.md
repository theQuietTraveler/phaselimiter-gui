# Benchmarking and Improvement Notes

## Current Fit-for-Purpose Assessment

For what this app does (single-file mastering orchestration via GUI), the current design is generally fit for purpose:

- The GUI does lightweight work (collecting parameters, launching process, rendering progress).
- Heavy CPU work is delegated to the external `phase_limiter` binary.
- Jobs are processed serially by `MasteringRunner`, which simplifies state handling.

## Stability Changes Implemented (April 10, 2026)

The following stability improvements are now implemented in code:

- **Graceful cancellation path**: job execution now uses `exec.CommandContext(...)` so active mastering processes are cancelled when the app shuts down.
- **Safe runner shutdown**: `MasteringRunner` uses context cancellation + `sync.Once` and closes `MasteringUpdate` on exit to avoid goroutine leaks.
- **Scanner error handling**: process stdout scanner errors are surfaced as failed job status instead of being ignored.
- **Defensive enqueue behavior**: new jobs are ignored after termination rather than blocking.
- **Output directory validation**: drag-and-drop now validates/creates output directory up front and reports an explicit failure row on errors.
- **Error visibility in UI**: added a `message` column in the table so parsing/runtime failures are visible to users.

## How to Benchmark

### 1) Microbenchmarks for parsing hot paths

```bash
go test ./internal/parsing -bench . -benchmem
```

This benchmarks:
- dropped URI parsing/decoding (`ParseDroppedFilePath`)
- progress output parsing (`ExtractProgression`)

### 2) End-to-end throughput benchmark (practical)

Use a folder of representative audio files and measure total wall-clock time:

```bash
time ./phaselimiter-gui
```

Then drag the same corpus and compare:
- total runtime
- CPU utilization
- memory peaks
- failure rate

### 3) Profiling ideas

- Add optional timing logs around process start/stop and per-file completion.
- Track average and p95 processing time per file in logs.

## Possible Improvements

1. **Parallel job execution (opt-in)**
   - Process multiple files concurrently with bounded worker pool.
   - Gate with CPU-core-aware default and user override.

2. **Better status/error visibility in UI**
   - Add a dedicated error/message column so parse/processing errors are obvious.

3. **I/O optimizations**
   - Skip work if output already exists and source mtime/hash unchanged.

4. **Backpressure and cancellation**
   - Add cancel button per queued item and global stop-all.

5. **Stable benchmark dataset**
   - Maintain a fixed corpus + expected outputs for reproducible performance checks.
