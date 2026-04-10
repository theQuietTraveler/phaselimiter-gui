#!/usr/bin/env bash
set -euo pipefail

echo "🔨 Building PhaseLimiter GUI"

if ! command -v go >/dev/null 2>&1; then
  echo "❌ Go is not installed."
  exit 1
fi

if ! command -v ffmpeg >/dev/null 2>&1; then
  echo "⚠️  FFmpeg not found in PATH. Processing may fail."
fi

rm -f phaselimiter-gui phaselimiter-gui.exe

if [[ "${OSTYPE:-}" == "msys" || "${OSTYPE:-}" == "win32" ]]; then
  go build -o phaselimiter-gui.exe .
  chmod +x phaselimiter-gui.exe || true
  echo "✅ Built phaselimiter-gui.exe"
else
  go build -o phaselimiter-gui .
  chmod +x phaselimiter-gui
  echo "✅ Built phaselimiter-gui"
fi
