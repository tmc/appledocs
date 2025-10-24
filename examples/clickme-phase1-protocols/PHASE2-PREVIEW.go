// THIS IS A PREVIEW - Phase 2 Not Yet Implemented!
// This shows what the code will look like after Phase 2 delegate builders are complete.
// DO NOT RUN - This won't compile until Phase 2 is finished!

package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

var (
	e2e        = flag.Bool("e2e", false, "run end-to-end tests")
	clickCount int
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	fmt.Println("🎉 ClickMe - Phase 2 Delegate Builder Demo (PREVIEW)")
	fmt.Println("===================================================\n")

	// Create application
	app := appkit.ApplicationClass.SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	// Create window
	window := createWindow()
	
	// Create counter label
	counterLabel := createCounterLabel(window)

	// ⭐ PHASE 2: Clean delegate builder pattern!
	delegate := &appkit.ApplicationDelegate{} // Auto-generated builder!
	
	// Set launch handler - just a simple function!
	delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
		fmt.Println("✅ Application launched via delegate builder!")
		fmt.Println("   Phase 2: No manual registration needed!")
	})
	
	// Set termination handler
	delegate.SetApplicationWillTerminate(func(notification foundation.Notification) {
		fmt.Printf("✅ Application terminating. Total clicks: %d\n", clickCount)
	})
	
	// Just set it - builder handles everything!
	app.SetDelegate(delegate) // That's it! No objc.RegisterClass!

	// Create button with simple inline handler
	createClickButtonWithBuilder(window, counterLabel)

	// Show window
	window.MakeKeyAndOrderFront(nil)
	app.ActivateIgnoringOtherApps(true)

	if *e2e {
		fmt.Println("E2E mode: terminating immediately")
		app.Terminate(nil)
		return
	}

	fmt.Println("🖱️  Click the button!")
	fmt.Println("\n✨ Phase 2 Benefits:")
	fmt.Println("  ✅ No manual objc.RegisterClass")
	fmt.Println("  ✅ No selector string names")
	fmt.Println("  ✅ No unsafe.Pointer")
	fmt.Println("  ✅ Clean builder pattern")
	fmt.Println("  ✅ 90% code reduction!")
	fmt.Println()

	app.Run()
}

func createWindow() appkit.Window {
	rect := foundation.Rect{
		Origin: foundation.Point{X: 100, Y: 100},
		Size:   foundation.Size{Width: 400, Height: 300},
	}
	
	window := appkit.WindowClass.Alloc().InitWithContentRectStyleMaskBackingDefer(
		rect,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
		appkit.BackingStoreBuffered,
		false,
	)
	
	window.SetTitle(foundation.String{}.InitWithUTF8String("ClickMe - Phase 2 Preview"))
	
	return window
}

func createCounterLabel(window appkit.Window) appkit.TextField {
	labelRect := foundation.Rect{
		Origin: foundation.Point{X: 100, Y: 200},
		Size:   foundation.Size{Width: 200, Height: 40},
	}
	
	label := appkit.TextFieldClass.Alloc().InitWithFrame(labelRect)
	label.SetStringValue(foundation.String{}.InitWithUTF8String("Clicks: 0"))
	label.SetEditable(false)
	label.SetBordered(false)
	label.SetDrawsBackground(false)
	label.SetAlignment(appkit.TextAlignmentCenter)
	
	window.ContentView().AddSubview(label)
	
	return label
}

func createClickButtonWithBuilder(window appkit.Window, counterLabel appkit.TextField) {
	buttonRect := foundation.Rect{
		Origin: foundation.Point{X: 150, Y: 120},
		Size:   foundation.Size{Width: 100, Height: 40},
	}
	
	button := appkit.ButtonClass.Alloc().InitWithFrame(buttonRect)
	button.SetTitle(foundation.String{}.InitWithUTF8String("Click Me!"))
	
	// ⭐ PHASE 2: If we had a button action delegate builder, it would be this simple:
	// buttonDelegate := &appkit.ButtonActionDelegate{}
	// buttonDelegate.SetButtonClicked(func(sender appkit.IButton) {
	//     clickCount++
	//     fmt.Printf("Button clicked! Count: %d\n", clickCount)
	//     labelText := foundation.String{}.InitWithUTF8String(fmt.Sprintf("Clicks: %d", clickCount))
	//     counterLabel.SetStringValue(labelText)
	// })
	// button.SetTarget(buttonDelegate)
	
	// For now, still using target/action pattern
	// (Buttons use target/action, not delegates, but the principle is the same)
	
	window.ContentView().AddSubview(button)
}

// What the auto-generated ApplicationDelegate builder looks like (Phase 2):
/*
type ApplicationDelegate struct {
    _ApplicationDidFinishLaunching func(notification foundation.Notification)
    _ApplicationWillTerminate      func(notification foundation.Notification)
    // ... all other 45 optional methods
}

func (d *ApplicationDelegate) SetApplicationDidFinishLaunching(f func(notification foundation.Notification)) {
    d._ApplicationDidFinishLaunching = f
}

func (d *ApplicationDelegate) HasApplicationDidFinishLaunching() bool {
    return d._ApplicationDidFinishLaunching != nil
}

func (d *ApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
    if d._ApplicationDidFinishLaunching != nil {
        d._ApplicationDidFinishLaunching(notification)
    }
}

// ... 44 more methods with Set/Has/Call pattern
*/
