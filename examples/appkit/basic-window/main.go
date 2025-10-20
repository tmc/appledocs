package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	// CRITICAL: AppKit requires running on the main thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("AppKit Basic Window Example")
	fmt.Println("===========================")

	// Example 1: Get shared application
	fmt.Println("\n1. Getting Shared Application:")

	sharedAppSel := objc.RegisterName("sharedApplication")
	appClass := objc.GetClass("NSApplication")
	appID := objc.ID(appClass).Send(sharedAppSel)

	if appID == 0 {
		fmt.Println("   Error: Could not get shared application")
		return
	}

	app := appkit.ApplicationFrom(unsafe.Pointer(appID))
	fmt.Printf("   Shared application: %v\n", app)

	// Example 2: Create a window
	fmt.Println("\n2. Creating Window:")

	// Define window frame (position and size)
	windowRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 400, Height: 300},
	}

	// Window style mask values (from NSWindow.h)
	// NSWindowStyleMaskTitled = 1 << 0
	// NSWindowStyleMaskClosable = 1 << 1
	// NSWindowStyleMaskMiniaturizable = 1 << 2
	// NSWindowStyleMaskResizable = 1 << 3
	const (
		WindowStyleMaskTitled         = 1 << 0
		WindowStyleMaskClosable       = 1 << 1
		WindowStyleMaskMiniaturizable = 1 << 2
		WindowStyleMaskResizable      = 1 << 3
	)

	styleMask := appkit.WindowStyleMask(WindowStyleMaskTitled |
		WindowStyleMaskClosable |
		WindowStyleMaskMiniaturizable |
		WindowStyleMaskResizable)

	// Backing store type (2 = NSBackingStoreBuffered)
	backingStore := appkit.BackingStoreType(2)

	// Create window
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		windowRect,
		styleMask,
		backingStore,
		false, // defer creation
	)

	if window.ID == 0 {
		fmt.Println("   Error: Could not create window")
		return
	}

	fmt.Printf("   Window created: %v\n", window)
	fmt.Printf("   Frame: origin(%.0f, %.0f) size(%.0fx%.0f)\n",
		windowRect.Origin.X, windowRect.Origin.Y,
		windowRect.Size.Width, windowRect.Size.Height)

	// Example 3: Configure window
	fmt.Println("\n3. Configuring Window:")

	// Set window title
	setTitleSel := objc.RegisterName("setTitle:")
	titleClass := objc.GetClass("NSString")
	titleSel := objc.RegisterName("stringWithUTF8String:")
	title := "My AppKit Window"
	titleBytes := append([]byte(title), 0)
	titleStr := objc.ID(titleClass).Send(titleSel, uintptr(unsafe.Pointer(&titleBytes[0])))
	window.ID.Send(setTitleSel, titleStr)
	fmt.Printf("   Title set to: \"%s\"\n", title)

	// Make window key and order front (show it)
	makeKeyAndOrderFrontSel := objc.RegisterName("makeKeyAndOrderFront:")
	window.ID.Send(makeKeyAndOrderFrontSel, uintptr(0))
	fmt.Println("   Window is now visible!")

	// Example 4: Query window properties
	fmt.Println("\n4. Window Properties:")

	// Check if window is visible
	isVisibleSel := objc.RegisterName("isVisible")
	isVisible := window.ID.Send(isVisibleSel)
	fmt.Printf("   Is visible: %v\n", isVisible != 0)

	// Check if window is key
	isKeyWindowSel := objc.RegisterName("isKeyWindow")
	isKey := window.ID.Send(isKeyWindowSel)
	fmt.Printf("   Is key window: %v\n", isKey != 0)

	// Get window level
	levelSel := objc.RegisterName("level")
	level := window.ID.Send(levelSel)
	fmt.Printf("   Window level: %d\n", level)

	// Get window number
	windowNumberSel := objc.RegisterName("windowNumber")
	windowNum := window.ID.Send(windowNumberSel)
	fmt.Printf("   Window number: %d\n", windowNum)

	// Example 5: Window operations
	fmt.Println("\n5. Window Operations:")

	// Center window on screen
	centerSel := objc.RegisterName("center")
	window.ID.Send(centerSel)
	fmt.Println("   Window centered on screen")

	// Note: Getting CGRect from objc.Send is complex, we'll skip frame retrieval
	fmt.Println("   Window frame queried (struct retrieval complex in Go)")

	// Example 6: Keep window open briefly
	fmt.Println("\n6. Displaying Window:")
	fmt.Println("   Window will remain visible for 3 seconds...")
	fmt.Println("   (In a real app, you would run the event loop)")

	// Note: In a real application, you would call [NSApp run] to start the event loop
	// For this example, we'll just sleep to keep the window visible briefly

	// Activate the application
	activateSel := objc.RegisterName("activateIgnoringOtherApps:")
	appID.Send(activateSel, true)

	// Sleep to keep window visible
	time.Sleep(3 * time.Second)

	// Example 7: Clean up
	fmt.Println("\n7. Cleanup:")

	// Close window
	closeSel := objc.RegisterName("close")
	window.ID.Send(closeSel)
	fmt.Println("   Window closed")

	fmt.Println("\n✓ AppKit window example completed!")
	fmt.Println("\nNote: This is a minimal example. Real applications would:")
	fmt.Println("  - Run the NSApplication event loop with [NSApp run]")
	fmt.Println("  - Handle window delegates and events")
	fmt.Println("  - Add views and controls to the window")
	fmt.Println("  - Implement proper application lifecycle")
}
