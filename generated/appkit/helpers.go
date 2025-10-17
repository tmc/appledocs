// Package-level convenience helpers for AppKit.
//
// This file contains hand-written helpers that complement the generated bindings.
// It is not overwritten during code generation.
package appkit

import (
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// RunApp is a convenience function that handles the boilerplate of setting up and running a macOS application.
// It:
// - Locks the calling goroutine to the OS thread (required for AppKit)
// - Gets the shared NSApplication instance
// - Sets the activation policy to Regular (shows in Dock)
// - Calls the setup function to allow the caller to configure the app and create windows
// - Finishes launching the application
// - Activates the application
// - Runs the main event loop
//
// The setup function receives the Application instance and should create and configure
// the application's windows and other UI elements.
//
// Example:
//
//	func main() {
//	    appkit.RunApp(func(app appkit.Application) {
//	        window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(...)
//	        window.MakeKeyAndOrderFront(0)
//	    })
//	}
func RunApp(setup func(app Application)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Get NSApplication shared instance
	appClass := objc.GetClass("NSApplication")
	app := ApplicationFrom(unsafe.Pointer(
		objc.ID(appClass).Send(objc.RegisterName("sharedApplication")),
	))

	// Set activation policy to regular (makes it a proper app with dock icon)
	// NSApplicationActivationPolicyRegular = 0
	app.ID.Send(objc.RegisterName("setActivationPolicy:"), 0)

	// Call setup function to let caller configure the app and create windows
	setup(app)

	// Finish launching and activate app
	app.ID.Send(objc.RegisterName("finishLaunching"))
	app.ID.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	// Run event loop (this blocks until the app quits)
	app.ID.Send(objc.RegisterName("run"))
}
