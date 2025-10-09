package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

// NSPoint represents a point in a 2D coordinate system
type NSPoint struct {
	X, Y float64
}

// NSSize represents a width and height
type NSSize struct {
	Width, Height float64
}

// NSRect represents a rectangle
type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	// Load Foundation and AppKit frameworks
	_, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Printf("Failed to load Foundation: %v\n", err)
		return
	}

	_, err = purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Printf("Failed to load AppKit: %v\n", err)
		return
	}

	// Get NSApplication shared application
	nsAppClass := objc.GetClass("NSApplication")
	sharedApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	// Set activation policy to regular (appears in dock)
	// NSApplicationActivationPolicyRegular = 0
	sharedApp.Send(objc.RegisterName("setActivationPolicy:"), 0)

	// Create window
	// NSWindowStyleMaskTitled = 1, NSWindowStyleMaskClosable = 2
	styleMask := uintptr(1 | 2)
	// NSBackingStoreBuffered = 2
	backingStoreType := uintptr(2)

	windowFrame := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 400, Height: 300},
	}

	nsWindowClass := objc.GetClass("NSWindow")
	window := objc.ID(nsWindowClass).Send(objc.RegisterName("alloc"))
	window = window.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		windowFrame, styleMask, backingStoreType, false)

	// Set window title
	title := createNSString("Hello from Purego!")
	window.Send(objc.RegisterName("setTitle:"), title)

	// Create button
	nsButtonClass := objc.GetClass("NSButton")
	button := objc.ID(nsButtonClass).Send(objc.RegisterName("alloc"))

	buttonFrame := NSRect{
		Origin: NSPoint{X: 150, Y: 120},
		Size:   NSSize{Width: 100, Height: 40},
	}
	button = button.Send(objc.RegisterName("initWithFrame:"), buttonFrame)

	// Set button title
	buttonTitle := createNSString("Click Me!")
	button.Send(objc.RegisterName("setTitle:"), buttonTitle)

	// Add button to window's content view
	contentView := window.Send(objc.RegisterName("contentView"))
	objc.ID(contentView).Send(objc.RegisterName("addSubview:"), button)

	// Make window key and bring to front
	window.Send(objc.RegisterName("makeKeyAndOrderFront:"), 0)

	// Activate ignoring other apps
	sharedApp.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	// In E2E mode, terminate immediately
	if *e2e {
		fmt.Println("E2E mode: terminating immediately")
		sharedApp.Send(objc.RegisterName("terminate:"), 0)
		return
	}

	// Run the application event loop
	sharedApp.Send(objc.RegisterName("run"))
}

// createNSString creates an NSString from a Go string
func createNSString(s string) objc.ID {
	nsStringClass := objc.GetClass("NSString")
	str := objc.ID(nsStringClass).Send(objc.RegisterName("alloc"))
	return str.Send(objc.RegisterName("initWithUTF8String:"), objc.RegisterName(s))
}
