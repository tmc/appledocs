package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	localObjc "github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
)

var (
	e2e          = flag.Bool("e2e", false, "run end-to-end tests")
	clickCount   int
	counterLabel appkit.TextField
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETest()
		return
	}

	// Create application using the class method
	app := appkit.ApplicationFrom(appkit.ApplicationClass.SharedApplication())
	// NSApplicationActivationPolicyRegular = 0
	app.SetActivationPolicy(unsafe.Pointer(uintptr(0)))

	// Create window with proper constructor
	contentRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 400, Height: 300},
	}

	// NSWindowStyleMaskTitled = 1, NSWindowStyleMaskClosable = 2
	styleMask := appkit.WindowStyleMask(1 | 2)
	// NSBackingStoreBuffered = 2
	backingStoreType := appkit.BackingStoreType(2)

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		contentRect,
		styleMask,
		backingStoreType,
		false,
	)

	// Set window title using the generated helper
	titleStr := localObjc.String("Hello from Generated Bindings!")
	window.SetTitle(unsafe.Pointer(titleStr))

	// Create counter label
	labelRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 200},
		Size:   coregraphics.CGSize{Width: 200, Height: 40},
	}
	counterLabel = appkit.NewTextField()
	counterLabel.ID.Send(objc.RegisterName("setFrame:"), labelRect)

	// Set label properties using generated helper
	clicksStr := localObjc.String("Clicks: 0")
	counterLabel.SetStringValue(unsafe.Pointer(clicksStr))
	counterLabel.SetEditable(false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setDrawsBackground:"), false)
	// NSTextAlignmentCenter = 2
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2)

	// Add label to window's content view
	contentView := window.ContentView()
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button
	buttonRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 150, Y: 120},
		Size:   coregraphics.CGSize{Width: 100, Height: 40},
	}
	button := appkit.NewButton()
	button.ID.Send(objc.RegisterName("setFrame:"), buttonRect)

	// Set button title - Button doesn't have SetTitle in generated code, use objc.Send
	buttonTitle := localObjc.String("Click Me!")
	button.ID.Send(objc.RegisterName("setTitle:"), buttonTitle)

	button.SetTarget(localObjc.ID(buttonHandler))
	button.SetAction(objc.RegisterName("buttonClicked:"))

	// Add button to window's content view
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(button.ID))

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(objc.ID(0))

	// Activate ignoring other apps
	app.ActivateIgnoringOtherApps(true)

	// Run the application event loop
	app.Run()
}

// runE2ETest runs automated end-to-end test with small delays for visibility
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Clickme Generated Bindings) ===")

	// Create application
	app := appkit.ApplicationFrom(appkit.ApplicationClass.SharedApplication())
	// NSApplicationActivationPolicyAccessory = 1 (No dock icon in tests)
	app.SetActivationPolicy(unsafe.Pointer(uintptr(1)))
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	contentRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 400, Height: 300},
	}

	styleMask := appkit.WindowStyleMask(1 | 2)
	backingStoreType := appkit.BackingStoreType(2)

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		contentRect,
		styleMask,
		backingStoreType,
		false,
	)

	titleStr := localObjc.String("E2E Test Window")
	window.SetTitle(unsafe.Pointer(titleStr))
	fmt.Println("✓ Created window with title")
	time.Sleep(100 * time.Millisecond)

	// Get content view
	contentView := window.ContentView()
	fmt.Println("✓ Got content view")
	time.Sleep(100 * time.Millisecond)

	// Create counter label
	labelRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 200},
		Size:   coregraphics.CGSize{Width: 200, Height: 40},
	}
	counterLabel = appkit.NewTextField()
	counterLabel.ID.Send(objc.RegisterName("setFrame:"), labelRect)

	clicksStr := localObjc.String("Clicks: 0")
	counterLabel.SetStringValue(unsafe.Pointer(clicksStr))
	counterLabel.SetEditable(false)
	counterLabel.ID.Send(objc.RegisterName("setBordered:"), false)
	counterLabel.ID.Send(objc.RegisterName("setDrawsBackground:"), false)
	counterLabel.ID.Send(objc.RegisterName("setAlignment:"), 2)
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(counterLabel.ID))
	fmt.Println("✓ Created counter label")
	time.Sleep(100 * time.Millisecond)

	// Create button handler
	buttonHandler := createButtonHandler()
	fmt.Println("✓ Created button handler")
	time.Sleep(100 * time.Millisecond)

	// Create button
	buttonRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 150, Y: 120},
		Size:   coregraphics.CGSize{Width: 100, Height: 40},
	}
	button := appkit.NewButton()
	button.ID.Send(objc.RegisterName("setFrame:"), buttonRect)

	buttonTitle := localObjc.String("Click Me!")
	button.ID.Send(objc.RegisterName("setTitle:"), buttonTitle)
	button.SetTarget(localObjc.ID(buttonHandler))
	button.SetAction(objc.RegisterName("buttonClicked:"))
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(button.ID))
	fmt.Println("✓ Created and configured button")
	time.Sleep(100 * time.Millisecond)

	// Verify button is valid
	if button.ID == 0 {
		fmt.Println("✗ FAIL: Button not created")
		os.Exit(1)
	}
	fmt.Println("✓ Button validation passed")
	time.Sleep(100 * time.Millisecond)

	// Verify label string values
	labelValue := counterLabel.StringValue()
	if labelValue == nil {
		fmt.Println("✗ FAIL: Label value not set")
		os.Exit(1)
	}
	fmt.Println("✓ Label values verified")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(objc.ID(0))
	fmt.Println("✓ Window displayed")
	time.Sleep(200 * time.Millisecond)

	// Close window
	window.Close()
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	os.Exit(0)
}

// createButtonHandler creates an NSObject subclass that handles button clicks
func createButtonHandler() objc.ID {
	className := "ButtonHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			clickCount++
			fmt.Printf("Button clicked! Count: %d\n", clickCount)

			labelText := localObjc.String(fmt.Sprintf("Clicks: %d", clickCount))
			counterLabel.SetStringValue(unsafe.Pointer(labelText))
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{},
			[]objc.MethodDef{{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked}})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	return handler.Send(objc.RegisterName("init"))
}
