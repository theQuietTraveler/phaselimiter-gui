#!/usr/bin/env python3
"""
Create a simple audio mastering icon for the PhaseLimiter app
"""

from PIL import Image, ImageDraw, ImageFont
import os

def create_icon():
    # Create a 512x512 icon (standard macOS app icon size)
    size = 512
    img = Image.new('RGBA', (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)
    
    # Background circle
    center = size // 2
    radius = size // 3
    draw.ellipse([center - radius, center - radius, center + radius, center + radius], 
                 fill=(45, 85, 255, 255), outline=(30, 60, 200, 255), width=8)
    
    # Audio waveform lines
    waveform_points = []
    for i in range(0, 360, 20):
        x = center + int(radius * 0.7 * (i - 180) / 180)
        y = center + int(radius * 0.3 * (i % 60 - 30) / 30)
        waveform_points.append((x, y))
    
    # Draw waveform
    for i in range(len(waveform_points) - 1):
        draw.line([waveform_points[i], waveform_points[i + 1]], 
                 fill=(255, 255, 255, 255), width=6)
    
    # Add "PL" text
    try:
        # Try to use a system font
        font = ImageFont.truetype("/System/Library/Fonts/Arial.ttf", 120)
    except:
        # Fallback to default font
        font = ImageFont.load_default()
    
    text = "PL"
    bbox = draw.textbbox((0, 0), text, font=font)
    text_width = bbox[2] - bbox[0]
    text_height = bbox[3] - bbox[1]
    
    text_x = center - text_width // 2
    text_y = center - text_height // 2 + 20
    
    # Draw text with shadow
    draw.text((text_x + 3, text_y + 3), text, fill=(0, 0, 0, 100), font=font)
    draw.text((text_x, text_y), text, fill=(255, 255, 255, 255), font=font)
    
    return img

def create_icon_set():
    """Create different icon sizes for macOS"""
    base_icon = create_icon()
    
    # Create icons directory
    os.makedirs("icons", exist_ok=True)
    
    # Standard macOS icon sizes
    sizes = [16, 32, 64, 128, 256, 512]
    
    for size in sizes:
        resized = base_icon.resize((size, size), Image.Resampling.LANCZOS)
        resized.save(f"icons/icon_{size}x{size}.png")
    
    # Create icns file structure
    os.makedirs("PhaseLimiter.iconset", exist_ok=True)
    
    # Copy to iconset with proper naming
    for size in sizes:
        resized = base_icon.resize((size, size), Image.Resampling.LANCZOS)
        resized.save(f"PhaseLimiter.iconset/icon_{size}x{size}.png")
        # Also create @2x versions for retina displays
        if size <= 256:
            resized_2x = base_icon.resize((size * 2, size * 2), Image.Resampling.LANCZOS)
            resized_2x.save(f"PhaseLimiter.iconset/icon_{size}x{size}@2x.png")
    
    print("Icons created successfully!")

if __name__ == "__main__":
    create_icon_set() 