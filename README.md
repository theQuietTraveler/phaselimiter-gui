# PhaseLimiter GUI

> **Vibe-coded because I got tired of using the slow version on the internet. This is a working prototype that can be improved upon by better programmers.**

A professional audio mastering GUI application that provides consistent loudness targeting, proper limiting, and high-quality audio processing without peaking or thin sound issues.

## Why This Exists

I was frustrated with slow, unreliable online mastering tools that either took forever to process or produced inconsistent results. So I built this local solution that:
- **Processes instantly** - No waiting for uploads/downloads
- **Works offline** - No internet dependency
- **Consistent results** - Same algorithm every time
- **Open source** - Can be improved by the community

## Features

- **Professional Audio Mastering** - Multi-stage processing pipeline for consistent results
- **LUFS Targeting** - Accurate loudness targeting to prevent peaking
- **Format Support** - Handles MP3, M4A, WAV, and other audio formats
- **Bass Preservation** - Maintains full, rich sound without thin audio
- **Drag & Drop Interface** - Simple and intuitive GUI
- **Output Folder Picker** - Browse button for selecting output directory
- **Real-time Progress** - Visual progress tracking during processing
- **Cross-platform** - Works on macOS, Windows, and Linux

## Runtime Architecture

The GUI is a thin wrapper around the bundled `phase_limiter` executable:

1. User drops files into the GTK window.
2. GUI resolves/validates file paths and output locations.
3. GUI starts `phase_limiter` with mastering arguments.
4. Progress lines from stdout are parsed and shown in the UI.
5. Processed WAV files are written to the selected output directory.

> Note: detailed DSP behavior lives inside the bundled `phaselimiter/bin/phase_limiter` tool, not in the Go GUI code.

## Known Issues & Limitations

### Current Issues
- **Some file formats may fail** - Particularly corrupted WAV files or unsupported formats
- **Inconsistent script versions** - Sometimes the old "Simple Audio Mastering Tool" is used instead of the new "Professional Audio Mastering Tool"
- **Format detection issues** - Some files with incorrect extensions may fail

### Workarounds
- **Use standard formats** - MP3, M4A, or properly formatted WAV files work best
- **Check file integrity** - Ensure audio files aren't corrupted
- **Restart the app** - If you see "Simple Audio Mastering Tool" in logs, restart the application

## Installation

### Quick Start (Recommended)

```bash
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui
./run.sh
```

`run.sh` now auto-builds the binary if missing and starts the app.

### Prerequisites

- **FFmpeg** - For audio format conversion
- **Go** - For building the GUI application

### macOS

```bash
# Install dependencies
brew install ffmpeg go

# Clone the repository
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui

# Build + run
./run.sh
```

### Linux

```bash
# Install dependencies
sudo apt-get install ffmpeg golang-go

# Clone and run
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui
./run.sh
```

### Windows

```bash
# Install dependencies via chocolatey
choco install ffmpeg golang

# Clone and run
git clone https://github.com/yourusername/phaselimiter-gui.git
cd phaselimiter-gui
./run.sh
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
- **phase_limiter** for the actual mastering implementation
- **Progress parsing in GUI** to display live status updates

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
2. **Processing fails** - Check file permissions and disk space
3. **Audio still peaks** - Lower the mastering intensity or target LUFS
4. **"Simple Audio Mastering Tool" appears** - Restart the application

### Debug Mode

Run with verbose output:
```bash
./phaselimiter-gui 2>&1 | tee debug.log
```

## Contributing

This is a vibe-coded prototype that needs improvement! Contributions welcome:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

### Areas for Improvement
- **Better error handling** - More graceful failure recovery
- **Improved format detection** - Better handling of corrupted files
- **Enhanced audio algorithms** - More sophisticated mastering techniques
- **UI improvements** - Better user experience
- **Performance optimization** - Faster processing

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- **FFmpeg** - Audio/video processing
- **GTK3** - Cross-platform GUI framework
- **Go** - Programming language and runtime

## Version History

- **v1.0.0** - Initial vibe-coded release
- **v1.1.0** - Fixed peaking issues and improved audio quality
- **v1.2.0** - Added LUFS targeting and bass preservation

---

**Note**: This is a working prototype built out of frustration with slow online tools. It's functional but could use improvement by better programmers. For best results, use high-quality source material and appropriate target LUFS levels for your intended playback environment.
