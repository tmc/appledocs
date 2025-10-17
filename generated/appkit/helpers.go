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

// View hierarchy convenience methods

// AddSubviewTyped adds a subview to the view hierarchy.
// This is a type-safe wrapper around AddSubview that accepts a View instead of unsafe.Pointer.
func (v View) AddSubviewTyped(subview View) {
	v.AddSubview(unsafe.Pointer(subview.ID))
}

// SetFrameRect sets the view's frame rectangle.
func (v View) SetFrameRect(x, y, width, height float64) {
	// Create NSRect structure
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	frame := NSRect{
		Origin: NSPoint{X: x, Y: y},
		Size:   NSSize{Width: width, Height: height},
	}
	objc.Send[objc.ID](v.ID, objc.RegisterName("setFrame:"), frame)
}

// Frame returns the view's frame rectangle as (x, y, width, height).
func (v View) Frame() (x, y, width, height float64) {
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	frame := objc.Send[NSRect](v.ID, objc.RegisterName("frame"))
	return frame.Origin.X, frame.Origin.Y, frame.Size.Width, frame.Size.Height
}

// Button convenience methods

// ButtonType constants for SetButtonType.
const (
	ButtonTypeMomentaryLight    = 0
	ButtonTypePushOnPushOff     = 1
	ButtonTypeToggle            = 2
	ButtonTypeSwitch            = 3
	ButtonTypeRadio             = 4
	ButtonTypeMomentaryChange   = 5
	ButtonTypeOnOff             = 6
	ButtonTypeMomentaryPushIn   = 7
	ButtonTypeAccelerator       = 8
	ButtonTypeMultiLevelAccelerator = 9
)

// BezelStyle constants for SetBezelStyle.
const (
	BezelStyleRounded           = 1
	BezelStyleRegularSquare     = 2
	BezelStyleDisclosure        = 5
	BezelStyleShadowlessSquare  = 6
	BezelStyleCircular          = 7
	BezelStyleTexturedSquare    = 8
	BezelStyleHelpButton        = 9
	BezelStyleSmallSquare       = 10
	BezelStyleTexturedRounded   = 11
	BezelStyleRoundRect         = 12
	BezelStyleRecessed          = 13
	BezelStyleRoundedDisclosure = 14
	BezelStyleInline            = 15
)

// SetButtonType sets the button's type.
func (b Button) SetButtonType(buttonType int) {
	objc.Send[objc.ID](b.ID, objc.RegisterName("setButtonType:"), buttonType)
}

// SetBezelStyle sets the button's bezel style.
func (b Button) SetBezelStyle(bezelStyle int) {
	objc.Send[objc.ID](b.ID, objc.RegisterName("setBezelStyle:"), bezelStyle)
}

// SetTarget sets the button's action target.
func (b Button) SetTarget(target objc.ID) {
	objc.Send[objc.ID](b.ID, objc.RegisterName("setTarget:"), target)
}

// SetAction sets the button's action selector.
func (b Button) SetAction(action objc.SEL) {
	objc.Send[objc.ID](b.ID, objc.RegisterName("setAction:"), action)
}

// SetTitleString sets the button's title from a Go string.
func (b Button) SetTitleString(title string) {
	strClass := objc.GetClass("NSString")
	nsStr := objc.ID(strClass).Send(objc.RegisterName("stringWithUTF8String:"), title)
	objc.Send[objc.ID](b.ID, objc.RegisterName("setTitle:"), nsStr)
}

// Control convenience methods (inherited by Button, TextField, etc.)

// TextAlignment constants for SetAlignment.
const (
	TextAlignmentLeft      = 0
	TextAlignmentCenter    = 1
	TextAlignmentRight     = 2
	TextAlignmentJustified = 3
	TextAlignmentNatural   = 4
)

// SetStringValue sets the control's string value.
func (c Control) SetStringValue(value string) {
	strClass := objc.GetClass("NSString")
	nsStr := objc.ID(strClass).Send(objc.RegisterName("stringWithUTF8String:"), value)
	objc.Send[objc.ID](c.ID, objc.RegisterName("setStringValue:"), nsStr)
}

// StringValue returns the control's string value as a Go string.
func (c Control) StringValue() string {
	nsStr := objc.Send[objc.ID](c.ID, objc.RegisterName("stringValue"))
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

// SetEditable sets whether the control is editable.
func (c Control) SetEditable(editable bool) {
	objc.Send[objc.ID](c.ID, objc.RegisterName("setEditable:"), editable)
}

// SetBordered sets whether the control has a border.
func (c Control) SetBordered(bordered bool) {
	objc.Send[objc.ID](c.ID, objc.RegisterName("setBordered:"), bordered)
}

// SetAlignment sets the text alignment.
func (c Control) SetAlignment(alignment int) {
	objc.Send[objc.ID](c.ID, objc.RegisterName("setAlignment:"), alignment)
}

// TextField convenience methods

// SetDrawsBackground sets whether the text field draws its background.
func (t TextField) SetDrawsBackground(draws bool) {
	objc.Send[objc.ID](t.ID, objc.RegisterName("setDrawsBackground:"), draws)
}

// SetBackgroundColor sets the text field's background color.
func (t TextField) SetBackgroundColor(color Color) {
	objc.Send[objc.ID](t.ID, objc.RegisterName("setBackgroundColor:"), color.ID)
}

// BackgroundColor returns the text field's background color.
func (t TextField) BackgroundColor() Color {
	colorID := objc.Send[objc.ID](t.ID, objc.RegisterName("backgroundColor"))
	return ColorFrom(unsafe.Pointer(colorID))
}
