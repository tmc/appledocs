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

	// Load required frameworks
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

	_, err = purego.Dlopen("/System/Library/Frameworks/WebKit.framework/WebKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Printf("Failed to load WebKit: %v\n", err)
		return
	}

	// Get NSApplication shared application
	nsAppClass := objc.GetClass("NSApplication")
	sharedApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	// Set activation policy to regular (appears in dock)
	// NSApplicationActivationPolicyRegular = 0
	sharedApp.Send(objc.RegisterName("setActivationPolicy:"), 0)

	// Activate ignoring other apps
	sharedApp.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	// Create URL for loading
	nsURLClass := objc.GetClass("NSURL")
	urlString := createNSString("http://progrium.com")
	url := objc.ID(nsURLClass).Send(objc.RegisterName("URLWithString:"), urlString)

	// Create URL request
	nsURLRequestClass := objc.GetClass("NSURLRequest")
	request := objc.ID(nsURLRequestClass).Send(objc.RegisterName("requestWithURL:"), url)

	// Create frame rect (1440x900)
	frame := NSRect{
		Origin: NSPoint{X: 0, Y: 0},
		Size:   NSSize{Width: 1440, Height: 900},
	}

	// Create WebView configuration
	wkWebViewConfigClass := objc.GetClass("WKWebViewConfiguration")
	config := objc.ID(wkWebViewConfigClass).Send(objc.RegisterName("alloc"))
	config = config.Send(objc.RegisterName("init"))

	// Create WebView
	wkWebViewClass := objc.GetClass("WKWebView")
	webView := objc.ID(wkWebViewClass).Send(objc.RegisterName("alloc"))
	webView = webView.Send(objc.RegisterName("initWithFrame:configuration:"), frame, config)

	// Load request in WebView
	webView.Send(objc.RegisterName("loadRequest:"), request)

	// Create window
	// NSWindowStyleMaskTitled = 1, NSWindowStyleMaskClosable = 2
	styleMask := uintptr(1 | 2) // Titled | Closable
	// NSBackingStoreBuffered = 2
	backingStoreType := uintptr(2)

	nsWindowClass := objc.GetClass("NSWindow")
	window := objc.ID(nsWindowClass).Send(objc.RegisterName("alloc"))
	windowFrame := NSRect{
		Origin: NSPoint{X: 0, Y: 0},
		Size:   NSSize{Width: 1440, Height: 900},
	}
	window = window.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		windowFrame, styleMask, backingStoreType, false)

	// Retain window to prevent deallocation
	window.Send(objc.RegisterName("retain"))

	// Set WebView as content view
	window.Send(objc.RegisterName("setContentView:"), webView)

	// Center window on screen
	window.Send(objc.RegisterName("center"))

	// Make window key and bring to front
	window.Send(objc.RegisterName("makeKeyAndOrderFront:"), window)

	// Create and set app delegate for handling window close
	delegateClass, err := createAppDelegateClass()
	if err != nil {
		fmt.Printf("Failed to create delegate class: %v\n", err)
		return
	}
	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc"))
	delegate = delegate.Send(objc.RegisterName("init"))
	sharedApp.Send(objc.RegisterName("setDelegate:"), delegate)

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

// createAppDelegateClass creates a custom NSApplicationDelegate class
// that handles applicationShouldTerminateAfterLastWindowClosed
func createAppDelegateClass() (objc.Class, error) {
	className := "AppDelegate"

	// Method that returns true for applicationShouldTerminateAfterLastWindowClosed
	shouldTerminate := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
		return true
	}

	// Register the class with objc.RegisterClass
	class, err := objc.RegisterClass(
		className,
		objc.GetClass("NSObject"),
		nil, // no protocols
		nil, // no fields
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
				Fn:  shouldTerminate,
			},
		},
	)

	return class, err
}
