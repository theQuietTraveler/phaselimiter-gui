#!/usr/bin/env bash
set -euo pipefail

echo "🎵 PhaseLimiter GUI"
echo "==================="

if ! command -v go >/dev/null 2>&1; then
  echo "❌ Go is not installed. Install Go and re-run ./run.sh"
  exit 1
fi

if ! command -v ffmpeg >/dev/null 2>&1; then
  echo "⚠️  FFmpeg not found in PATH. Processing may fail."
fi

if [[ ! -x ./phaselimiter-gui ]]; then
  echo "🔨 Binary not found, building..."
  ./build.sh
fi

echo "🚀 Starting app..."
exec ./phaselimiter-gui
