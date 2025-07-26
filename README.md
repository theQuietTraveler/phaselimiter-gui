# PhaseLimiter GUI

A professional audio mastering GUI application that provides consistent loudness targeting, proper limiting, and high-quality audio processing without peaking or thin sound issues.

## Features

- **Professional Audio Mastering** - Multi-stage processing pipeline for consistent results
- **LUFS Targeting** - Accurate loudness targeting to prevent peaking
- **Format Support** - Handles MP3, M4A, WAV, and other audio formats
- **Bass Preservation** - Maintains full, rich sound without thin audio
- **Drag & Drop Interface** - Simple and intuitive GUI
- **Real-time Progress** - Visual progress tracking during processing
- **Cross-platform** - Works on macOS, Windows, and Linux

## Audio Processing Pipeline

The application uses a professional 8-stage mastering pipeline:

1. **Format Conversion** - Convert any audio format to compatible WAV
2. **Audio Analysis** - Analyze audio characteristics and statistics
3. **Normalization** - Normalize to -0.5dB to prevent clipping
4. **Compression** - Gentle compression to control dynamics
5. **Limiting** - Hard limiting to prevent peaks
6. **EQ** - Subtle EQ to preserve bass and body
7. **LUFS Targeting** - Final gain adjustment for target loudness
8. **Final Conversion** - Output to WAV format

## Installation

### Prerequisites

- **FFmpeg** - For audio format conversion
- **SoX** - For audio processing and effects
- **Go** - For building the GUI application

### macOS

```bash
# Install dependencies
brew install ffmpeg sox go

# Clone the repository
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui

# Build the application
go build -o phaselimiter-gui main.go mastering.go cmd_hide_window.go

# Run the application
./phaselimiter-gui
```

### Linux

```bash
# Install dependencies
sudo apt-get install ffmpeg sox golang-go

# Clone and build (same as macOS)
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui
go build -o phaselimiter-gui main.go mastering.go cmd_hide_window.go
./phaselimiter-gui
```

### Windows

```bash
# Install dependencies via chocolatey
choco install ffmpeg sox golang

# Clone and build (same as other platforms)
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui
go build -o phaselimiter-gui.exe main.go mastering.go cmd_hide_window.go
phaselimiter-gui.exe
```

## Usage

1. **Launch the application** - Run the built executable
2. **Set parameters**:
   - **Output directory** - Where processed files will be saved
   - **Target loudness** - LUFS target (default: -9 dB)
   - **Mastering intensity** - Processing strength (0.0-1.0)
   - **Preserve bass** - Check to maintain bass frequencies
3. **Drag & drop** audio files onto the application window
4. **Monitor progress** - Watch real-time processing status
5. **Find results** - Processed files in your output directory

## Configuration

### Target LUFS Settings

- **-12 dB and below**: Very conservative processing, minimal compression
- **-9 dB**: Standard mastering level, moderate compression
- **-6 dB and above**: Aggressive processing, heavy compression

### Mastering Intensity

- **0.0**: Minimal processing, preserves original dynamics
- **0.5**: Moderate processing, balanced approach
- **1.0**: Full processing, maximum loudness and consistency

## Technical Details

### Audio Processing

The application uses a sophisticated audio processing chain:

- **FFmpeg** for format conversion and compatibility
- **SoX** for audio effects and processing
- **Multi-stage limiting** to prevent peaking
- **Adaptive compression** based on target LUFS
- **Conservative EQ** to preserve audio quality

### File Structure

```
phaselimiter-gui/
├── main.go                 # Main GUI application
├── mastering.go            # Audio processing logic
├── cmd_hide_window.go      # Platform-specific window handling
├── phaselimiter/
│   └── bin/
│       └── phase_limiter   # Python audio processing script
├── PhaseLimiter.app/       # macOS app bundle
└── icons/                  # Application icons
```

## Troubleshooting

### Common Issues

1. **"FFmpeg not found"** - Install FFmpeg via package manager
2. **"SoX not found"** - Install SoX via package manager
3. **Processing fails** - Check file permissions and disk space
4. **Audio still peaks** - Lower the mastering intensity or target LUFS

### Debug Mode

Run with verbose output:
```bash
./phaselimiter-gui 2>&1 | tee debug.log
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- **FFmpeg** - Audio/video processing
- **SoX** - Audio effects and processing
- **GTK3** - Cross-platform GUI framework
- **Go** - Programming language and runtime

## Version History

- **v1.0.0** - Initial release with professional mastering pipeline
- **v1.1.0** - Fixed peaking issues and improved audio quality
- **v1.2.0** - Added LUFS targeting and bass preservation

---

**Note**: This application is designed for professional audio mastering. For best results, use high-quality source material and appropriate target LUFS levels for your intended playback environment. 