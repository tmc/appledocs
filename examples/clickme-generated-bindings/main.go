package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
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

	// Create application
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 400, 300,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("Hello from Generated Bindings!")

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	window.ContentView().AddSubviewTyped(counterLabel)

	// Create button handler
	buttonHandler := createButtonHandler()

	// Create button
	button := appkit.NewButtonWithFrame(150, 120, 100, 40)
	button.SetTitleString("Click Me!")
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))

	// Add button to window's content view
	window.ContentView().AddSubviewTyped(button)

	// Make window key and bring to front
	window.MakeKeyAndOrderFront(0)

	// Activate ignoring other apps
	app.ActivateIgnoringOtherApps(true)

	// Run the application event loop
	app.Run()
}

// runE2ETest runs automated end-to-end test with small delays for visibility
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Clickme Generated Bindings) ===")

	// Create application
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory) // No dock icon in tests
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 400, 300,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window with title")
	time.Sleep(100 * time.Millisecond)

	// Get content view
	contentView := window.ContentView()
	fmt.Println("✓ Got content view")
	time.Sleep(100 * time.Millisecond)

	// Create counter label
	counterLabel = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetDrawsBackground(false)
	counterLabel.SetAlignment(appkit.TextAlignmentCenter)
	contentView.AddSubviewTyped(counterLabel)
	fmt.Println("✓ Created counter label")
	time.Sleep(100 * time.Millisecond)

	// Create button handler
	buttonHandler := createButtonHandler()
	fmt.Println("✓ Created button handler")
	time.Sleep(100 * time.Millisecond)

	// Create button
	button := appkit.NewButtonWithFrame(150, 120, 100, 40)
	button.SetTitleString("Click Me!")
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))
	contentView.AddSubviewTyped(button)
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
	labelValue := counterLabel.ID.Send(objc.RegisterName("stringValue"))
	if labelValue == 0 {
		fmt.Println("✗ FAIL: Label value not set")
		os.Exit(1)
	}
	fmt.Println("✓ Label values verified")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(0)
	fmt.Println("✓ Window displayed")
	time.Sleep(200 * time.Millisecond)

	// Close window
	window.ID.Send(objc.RegisterName("close"))
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
			counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{},
			[]objc.MethodDef{{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked}})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	return handler.Send(objc.RegisterName("init"))
}
