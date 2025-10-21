package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
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
	app := appkit.ApplicationClass.SharedApplication()
	appkit.ApplicationFrom(app).SetActivationPolicy(appkit.ActivationPolicyRegular)

	// Create window with proper constructor
	contentRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 400, Height: 300},
	}
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		contentRect,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
		appkit.BackingStoreBuffered,
		false,
	)
	window.SetTitle(objc.String("Hello from Generated Bindings!"))

	// Create counter label
	labelRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 200},
		Size:   coregraphics.CGSize{Width: 200, Height: 40},
	}
	counterLabel = appkit.TextFieldClass.TextFieldWithFrame(labelRect)
	counterLabel.SetStringValue(objc.String("Clicks: 0"))
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)

	// Add label to window's content view
	contentView := window.ContentView()
	appkit.ViewFrom(contentView).AddSubview(counterLabel.ID)

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button
	buttonRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 150, Y: 120},
		Size:   coregraphics.CGSize{Width: 100, Height: 40},
	}
	button := appkit.ButtonClass.ButtonWithFrame(buttonRect)
	button.SetTitle(objc.String("Click Me!"))
	button.SetTarget(buttonHandler)
	button.SetAction(objc.Sel("buttonClicked:"))

	// Add button to window's content view
	appkit.ViewFrom(contentView).AddSubview(button.ID)

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(nil)

	// Activate ignoring other apps
	appkit.ApplicationFrom(app).ActivateIgnoringOtherApps(true)

	// Run the application event loop
	appkit.ApplicationFrom(app).Run()
}

// runE2ETest runs automated end-to-end test with small delays for visibility
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Clickme Generated Bindings) ===")

	// Create application
	app := appkit.ApplicationClass.SharedApplication()
	appkit.ApplicationFrom(app).SetActivationPolicy(appkit.ActivationPolicyAccessory) // No dock icon in tests
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	contentRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 400, Height: 300},
	}
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		contentRect,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
		appkit.BackingStoreBuffered,
		false,
	)
	window.SetTitle(objc.String("E2E Test Window"))
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
	counterLabel = appkit.TextFieldClass.TextFieldWithFrame(labelRect)
	counterLabel.SetStringValue(objc.String("Clicks: 0"))
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	appkit.ViewFrom(contentView).AddSubview(counterLabel.ID)
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
	button := appkit.ButtonClass.ButtonWithFrame(buttonRect)
	button.SetTitle(objc.String("Click Me!"))
	button.SetTarget(buttonHandler)
	button.SetAction(objc.Sel("buttonClicked:"))
	appkit.ViewFrom(contentView).AddSubview(button.ID)
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
	window.MakeKeyAndOrderFront(nil)
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
			counterLabel.SetStringValue(objc.String(fmt.Sprintf("Clicks: %d", clickCount)))
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{},
			[]objc.MethodDef{{Cmd: objc.Sel("buttonClicked:"), Fn: buttonClicked}})
	}
	handler := objc.ID(class).Send(objc.Sel("alloc"))
	return handler.Send(objc.Sel("init"))
}
