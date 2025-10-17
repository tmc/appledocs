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

// Application lifecycle convenience methods

// SharedApplication returns the singleton NSApplication instance.
// This is equivalent to [NSApplication sharedApplication].
func SharedApplication() Application {
	appClass := objc.GetClass("NSApplication")
	return ApplicationFrom(unsafe.Pointer(
		objc.ID(appClass).Send(objc.RegisterName("sharedApplication")),
	))
}

// ActivationPolicy values for SetActivationPolicy.
const (
	ActivationPolicyRegular    = 0 // Regular app with Dock icon and menu bar
	ActivationPolicyAccessory  = 1 // Accessory app without Dock icon
	ActivationPolicyProhibited = 2 // Cannot activate or become frontmost
)

// SetActivationPolicy sets the application's activation policy.
// Use ActivationPolicyRegular for normal applications.
func (a Application) SetActivationPolicy(policy int) bool {
	return objc.Send[bool](a.ID, objc.RegisterName("setActivationPolicy:"), policy)
}

// ActivateIgnoringOtherApps activates the application and brings it to the front.
// If ignoreOtherApps is true, the application becomes active regardless of other apps.
func (a Application) ActivateIgnoringOtherApps(ignoreOtherApps bool) {
	objc.Send[objc.ID](a.ID, objc.RegisterName("activateIgnoringOtherApps:"), ignoreOtherApps)
}

// Run starts the main event loop. This blocks until the application terminates.
func (a Application) Run() {
	objc.Send[objc.ID](a.ID, objc.RegisterName("run"))
}

// Terminate terminates the application, closing all windows.
// Sender is typically the object that initiated the termination (can be nil/0).
func (a Application) Terminate(sender objc.ID) {
	objc.Send[objc.ID](a.ID, objc.RegisterName("terminate:"), sender)
}

// FinishLaunching completes the application launch process.
// This should be called before Run() if you're managing the app lifecycle manually.
func (a Application) FinishLaunching() {
	objc.Send[objc.ID](a.ID, objc.RegisterName("finishLaunching"))
}

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

	app := SharedApplication()
	app.SetActivationPolicy(ActivationPolicyRegular)

	// Call setup function to let caller configure the app and create windows
	setup(app)

	// Finish launching and activate app
	app.FinishLaunching()
	app.ActivateIgnoringOtherApps(true)

	// Run event loop (this blocks until the app quits)
	app.Run()
}

// Window convenience methods

// SetTitle sets the window's title string.
func (w Window) SetTitle(title string) {
	strClass := objc.GetClass("NSString")
	nsStr := objc.ID(strClass).Send(objc.RegisterName("stringWithUTF8String:"), title)
	objc.Send[objc.ID](w.ID, objc.RegisterName("setTitle:"), nsStr)
}

// Title returns the window's title string.
func (w Window) Title() string {
	nsStr := objc.Send[objc.ID](w.ID, objc.RegisterName("title"))
	if nsStr == 0 {
		return ""
	}
	cStr := objc.Send[*byte](nsStr, objc.RegisterName("UTF8String"))
	if cStr == nil {
		return ""
	}
	length := 0
	for ptr := cStr; *ptr != 0; ptr = (*byte)(unsafe.Add(unsafe.Pointer(ptr), 1)) {
		length++
	}
	return string(unsafe.Slice(cStr, length))
}

// SetDelegate sets the window's delegate.
func (w Window) SetDelegate(delegate objc.ID) {
	objc.Send[objc.ID](w.ID, objc.RegisterName("setDelegate:"), delegate)
}

// ContentView returns the window's content view.
func (w Window) ContentView() View {
	view := objc.Send[objc.ID](w.ID, objc.RegisterName("contentView"))
	return ViewFrom(unsafe.Pointer(view))
}

// SetContentView sets the window's content view.
func (w Window) SetContentView(view View) {
	objc.Send[objc.ID](w.ID, objc.RegisterName("setContentView:"), view.ID)
}
