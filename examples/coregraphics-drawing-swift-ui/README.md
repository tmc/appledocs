# CoreGraphics Drawing with Swift UI

Complete end-to-end example: **Go → Swift → AppKit → CoreGraphics**

## Overview

This example demonstrates the full stack:
1. **Go** launches the application using purego
2. **Swift** provides AppKit UI layer with `@_cdecl` exports
3. **AppKit** creates native macOS window
4. **CoreGraphics** renders the drawing

## Architecture

```
┌─────────────────────────────────────────┐
│            Go Application               │
│         (main.go + purego)              │
└──────────────┬──────────────────────────┘
               │ dlsym/dlopen
               ▼
┌─────────────────────────────────────────┐
│         Swift UI Layer                  │
│      (AppKitUI.swift + @_cdecl)         │
└──────────────┬──────────────────────────┘
               │ imports
               ▼
┌─────────────────────────────────────────┐
│           AppKit Framework              │
│    (NSApplication, NSWindow, NSView)    │
└──────────────┬──────────────────────────┘
               │ uses
               ▼
┌─────────────────────────────────────────┐
│       CoreGraphics Framework            │
│         (CGContext, CGPath)             │
└─────────────────────────────────────────┘
```

## What It Does

- Creates a native macOS window (400×400)
- Displays CoreGraphics drawing:
  - Blue rectangle
  - Red circle with black stroke
  - Green triangle
- Runs full AppKit event loop
- Closes cleanly when window is closed

## Building

```bash
make          # Build Swift library
make run      # Build and run
```

## Running

```bash
go run main.go
```

You should see:
1. Terminal output showing initialization
2. A macOS window appear with the drawing
3. Window behaves like a native macOS app
4. Close window to exit

## Key Features

### No cgo
Uses purego for all Swift interop - no C compiler needed!

### Native macOS App
Full AppKit integration:
- Native window decorations
- Proper app activation
- Event loop integration
- Window management

### Clean Architecture
Clear separation of concerns:
- Go: Application logic and launching
- Swift: UI layer and AppKit bridging
- AppKit: Native macOS windowing
- CoreGraphics: Rendering

## API

The Swift library exports:
- `ui_init_app()` - Initialize NSApplication
- `ui_create_window()` - Create and show window
- `ui_run_app()` - Run event loop (blocks)
- `ui_save_png(path)` - Save window contents
- `ui_quit()` - Quit application

## Comparison

This demonstrates the **complete workflow** for Go→Swift→macOS apps:

| Layer | Technology | Purpose |
|-------|-----------|---------|
| Application | Go | Main logic, lifecycle |
| Interop | purego | Call Swift without cgo |
| UI Framework | Swift + AppKit | Native macOS UI |
| Graphics | CoreGraphics | 2D rendering |

## Why This Matters

This proves you can build **native macOS apps in Go** using Swift for the UI layer, with:
- ✅ No cgo overhead
- ✅ Native performance
- ✅ Full AppKit capabilities
- ✅ Clean Swift code
- ✅ Type-safe Go bindings

Perfect template for Go apps that need native macOS UI!
