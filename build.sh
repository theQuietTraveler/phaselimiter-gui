#!/bin/bash

# PhaseLimiter GUI Build Script
# This script builds the application for different platforms

set -e

echo "🎵 PhaseLimiter GUI Build Script"
echo "=================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed. Please install Go first."
    exit 1
fi

# Check if required dependencies are installed
echo "🔍 Checking dependencies..."

if ! command -v ffmpeg &> /dev/null; then
    echo "⚠️  Warning: FFmpeg not found. Audio processing may fail."
fi

if ! command -v sox &> /dev/null; then
    echo "⚠️  Warning: SoX not found. Audio processing may fail."
fi

# Build for current platform
echo "🔨 Building PhaseLimiter GUI..."

# Remove old build
rm -f phaselimiter-gui phaselimiter-gui.exe

# Build the application
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    # Windows
    go build -o phaselimiter-gui.exe main.go mastering.go cmd_hide_window.go
    echo "✅ Built phaselimiter-gui.exe for Windows"
else
    # macOS and Linux
    go build -o phaselimiter-gui main.go mastering.go cmd_hide_window.go
    echo "✅ Built phaselimiter-gui for $(uname -s)"
fi

# Make executable
chmod +x phaselimiter-gui*

echo ""
echo "🎉 Build completed successfully!"
echo ""
echo "To run the application:"
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    echo "  ./phaselimiter-gui.exe"
else
    echo "  ./phaselimiter-gui"
fi
echo ""
echo "Make sure FFmpeg and SoX are installed for full functionality." 