// main.go - Launch AppKit UI from Go using Swift bindings

package main

import (
	"fmt"
	"log"

	"github.com/ebitengine/purego"
)

// SwiftUI wraps the Swift AppKit UI functions
type SwiftUI struct {
	initApp      func()
	createWindow func()
	runApp       func()
	savePNG      func(path *byte, outSuccess *bool)
	quit         func()
}

func NewSwiftUI() (*SwiftUI, error) {
	lib, err := purego.Dlopen("./libCoreGraphicsUI.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return nil, fmt.Errorf("failed to load library: %w", err)
	}

	ui := &SwiftUI{}
	purego.RegisterLibFunc(&ui.initApp, lib, "ui_init_app")
	purego.RegisterLibFunc(&ui.createWindow, lib, "ui_create_window")
	purego.RegisterLibFunc(&ui.runApp, lib, "ui_run_app")
	purego.RegisterLibFunc(&ui.savePNG, lib, "ui_save_png")
	purego.RegisterLibFunc(&ui.quit, lib, "ui_quit")

	return ui, nil
}

func (ui *SwiftUI) InitApp() {
	fmt.Println("Go: Initializing Swift UI...")
	ui.initApp()
}

func (ui *SwiftUI) CreateWindow() {
	fmt.Println("Go: Creating window...")
	ui.createWindow()
}

func (ui *SwiftUI) RunApp() {
	fmt.Println("Go: Starting event loop (this will block)...")
	ui.runApp()
	fmt.Println("Go: Event loop ended")
}

func (ui *SwiftUI) SavePNG(path string) error {
	pathBytes := append([]byte(path), 0)
	var success bool

	ui.savePNG(&pathBytes[0], &success)

	if !success {
		return fmt.Errorf("failed to save PNG")
	}
	return nil
}

func main() {
	fmt.Println("=== CoreGraphics Drawing with Swift UI ===")
	fmt.Println("Go → Swift → AppKit → CoreGraphics")
	fmt.Println()

	// Load Swift UI library
	ui, err := NewSwiftUI()
	if err != nil {
		log.Fatalf("Failed to load Swift UI: %v", err)
	}

	// Initialize AppKit application
	ui.InitApp()

	// Create and show window
	ui.CreateWindow()

	// Optional: Save screenshot before showing
	// (Note: window must be visible for this to work)

	fmt.Println()
	fmt.Println("Window should now be visible!")
	fmt.Println("Close the window to exit.")
	fmt.Println()

	// Run event loop (blocks until window closes)
	ui.RunApp()

	fmt.Println("Application closed.")
}
