#!/bin/bash

# PhaseLimiter GUI Run Script
# Ensures the latest version is used and handles common issues

echo "🎵 PhaseLimiter GUI - Vibe-coded Audio Mastering"
echo "=================================================="

# Kill any existing processes
echo "🔄 Stopping any existing processes..."
pkill -f phaselimiter-gui 2>/dev/null || true

# Clean temporary files
echo "🧹 Cleaning temporary files..."
rm -rf /tmp/phaselimiter 2>/dev/null || true

# Ensure scripts are up to date
echo "📝 Ensuring latest scripts..."
cp phaselimiter/bin/phase_limiter PhaseLimiter.app/Contents/MacOS/phaselimiter/bin/phase_limiter

# Build if needed
if [ ! -f "./phaselimiter-gui" ]; then
    echo "🔨 Building application..."
    ./build.sh
fi

# Check dependencies
echo "🔍 Checking dependencies..."
if ! command -v ffmpeg &> /dev/null; then
    echo "⚠️  Warning: FFmpeg not found. Audio processing may fail."
fi

if ! command -v sox &> /dev/null; then
    echo "⚠️  Warning: SoX not found. Audio processing may fail."
fi

# Run the application
echo "🚀 Starting PhaseLimiter GUI..."
echo ""
echo "💡 Tips:"
echo "  - Drag & drop audio files onto the window"
echo "  - Use standard formats (MP3, M4A, WAV) for best results"
echo "  - If processing fails, try restarting the app"
echo ""

./phaselimiter-gui 