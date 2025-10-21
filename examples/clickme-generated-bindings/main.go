package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	localobjc "github.com/tmc/appledocs/generated/objc"
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
	nsAppClass := localobjc.GetClass("NSApplication")
	sharedApp := localobjc.ID(nsAppClass).Send(localobjc.Sel("sharedApplication"))
	app := appkit.ApplicationFrom(unsafe.Pointer(sharedApp))

	// Set activation policy to regular (appears in dock)
	// NSApplicationActivationPolicyRegular = 0
	sharedApp.Send(localobjc.Sel("setActivationPolicy:"), 0)

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
	window.SetTitle("Hello from Generated Bindings!")

	// Create counter label with frame (like purego example)
	labelRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 200},
		Size:   coregraphics.CGSize{Width: 200, Height: 40},
	}
	// Use low-level objc calls like purego does
	nsTextFieldClass := localobjc.GetClass("NSTextField")
	counterLabelID := localobjc.ID(nsTextFieldClass).Send(localobjc.Sel("alloc"))
	counterLabel = appkit.TextFieldFrom(unsafe.Pointer(counterLabelID))
	counterLabel.ID = counterLabel.ID.Send(localobjc.Sel("initWithFrame:"), labelRect)

	// Set label properties using generated helper
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBezeled(false)
	counterLabel.SetDrawsBackground(false)
	// NSTextAlignmentCenter = 2
	appkit.ControlFrom(unsafe.Pointer(counterLabel.ID)).SetAlignment(unsafe.Pointer(uintptr(2)))

	// Add label to window's content view
	contentView := window.ContentView()
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button with frame (like purego example)
	buttonRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 150, Y: 120},
		Size:   coregraphics.CGSize{Width: 100, Height: 40},
	}
	// Use low-level objc calls like purego does
	nsButtonClass := localobjc.GetClass("NSButton")
	buttonID := localobjc.ID(nsButtonClass).Send(localobjc.Sel("alloc"))
	button := appkit.ButtonFrom(unsafe.Pointer(buttonID))
	button.ID = button.ID.Send(localobjc.Sel("initWithFrame:"), buttonRect)

	// Set button title (Button doesn't have SetTitle, use low-level call)
	button.ID.Send(localobjc.Sel("setTitle:"), localobjc.String("Click Me!"))

	button.SetTarget(localobjc.ID(buttonHandler))
	button.SetAction(localobjc.Sel("buttonClicked:"))

	// Add button to window's content view
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(button.ID))

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(localobjc.ID(0))

	// Activate ignoring other apps
	app.ActivateIgnoringOtherApps(true)

	// Run the application event loop (call via Send like purego does)
	app.ID.Send(localobjc.Sel("run"))
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

	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window with title")
	time.Sleep(100 * time.Millisecond)

	// Get content view
	contentView := window.ContentView()
	fmt.Println("✓ Got content view")
	time.Sleep(100 * time.Millisecond)

	// Create counter label with frame (like purego example)
	labelRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 200},
		Size:   coregraphics.CGSize{Width: 200, Height: 40},
	}
	// Use low-level objc calls like purego does
	nsTextFieldClass := localobjc.GetClass("NSTextField")
	counterLabelID := localobjc.ID(nsTextFieldClass).Send(localobjc.Sel("alloc"))
	counterLabel = appkit.TextFieldFrom(unsafe.Pointer(counterLabelID))
	counterLabel.ID = counterLabel.ID.Send(localobjc.Sel("initWithFrame:"), labelRect)

	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBezeled(false)
	counterLabel.SetDrawsBackground(false)
	// NSTextAlignmentCenter = 2
	appkit.ControlFrom(unsafe.Pointer(counterLabel.ID)).SetAlignment(unsafe.Pointer(uintptr(2)))
	appkit.ViewFrom(contentView).AddSubview(unsafe.Pointer(counterLabel.ID))
	fmt.Println("✓ Created counter label")
	time.Sleep(100 * time.Millisecond)

	// Create button handler
	buttonHandler := createButtonHandler()
	fmt.Println("✓ Created button handler")
	time.Sleep(100 * time.Millisecond)

	// Create button with frame (like purego example)
	buttonRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 150, Y: 120},
		Size:   coregraphics.CGSize{Width: 100, Height: 40},
	}
	// Use low-level objc calls like purego does
	nsButtonClass := localobjc.GetClass("NSButton")
	buttonID := localobjc.ID(nsButtonClass).Send(localobjc.Sel("alloc"))
	button := appkit.ButtonFrom(unsafe.Pointer(buttonID))
	button.ID = button.ID.Send(localobjc.Sel("initWithFrame:"), buttonRect)

	// Set button title (Button doesn't have SetTitle, use low-level call)
	button.ID.Send(localobjc.Sel("setTitle:"), localobjc.String("Click Me!"))
	button.SetTarget(localobjc.ID(buttonHandler))
	button.SetAction(localobjc.Sel("buttonClicked:"))
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
	if labelValue == "" {
		fmt.Println("✗ FAIL: Label value not set")
		os.Exit(1)
	}
	fmt.Println("✓ Label values verified")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(localobjc.ID(0))
	fmt.Println("✓ Window displayed")
	time.Sleep(200 * time.Millisecond)

	// Close window
	window.Close()
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	os.Exit(0)
}

// createButtonHandler creates an NSObject subclass that handles button clicks
func createButtonHandler() localobjc.ID {
	className := "ButtonHandler"
	class := localobjc.GetClass(className)
	if class == 0 {
		superClass := localobjc.GetClass("NSObject")
		buttonClicked := func(self localobjc.ID, _cmd localobjc.SEL, sender localobjc.ID) {
			clickCount++
			fmt.Printf("Button clicked! Count: %d\n", clickCount)

			counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{},
			[]objc.MethodDef{{Cmd: localobjc.Sel("buttonClicked:"), Fn: buttonClicked}})
	}
	handler := localobjc.ID(class).Send(localobjc.Sel("alloc"))
	return handler.Send(localobjc.Sel("init"))
}
