package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

var (
	clickCountSimple   int
	counterLabelSimple appkit.TextField
)

func init() {
	runtime.LockOSThread()
}

// Simple version matching darwinkit's clickme - NO RunApp, NO delegate
func mainSimple() {
	runtime.LockOSThread()
	flag.Parse()

	fmt.Println("=== Click Me (Simple - No Delegate) ===")

	// Create application
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	fmt.Println("DEBUG: Creating window...")

	// Create window using generated constructor (BEFORE app.Run())
	styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
	window := appkit.NewWindowWithFrame(100, 100, 400, 300, styleMask)
	fmt.Printf("DEBUG: Created window, ID: %v\n", window.ID)

	window.SetTitle("Click Me (Simple)")
	fmt.Println("DEBUG: Set window title")

	// Create counter label
	counterLabelSimple = appkit.NewTextFieldWithFrame(100, 200, 200, 40)
	counterLabelSimple.SetStringValue("Clicks: 0")
	counterLabelSimple.SetEditable(false)
	counterLabelSimple.SetBordered(false)
	counterLabelSimple.SetDrawsBackground(false)
	counterLabelSimple.SetAlignment(appkit.TextAlignmentCenter)
	window.ContentView().AddSubviewTyped(counterLabelSimple)
	fmt.Println("DEBUG: Created and added counter label")

	// Create button handler
	buttonHandler := createSimpleButtonHandler()

	// Create button
	button := appkit.NewButtonWithFrame(150, 120, 100, 40)
	button.SetTitleString("Click Me!")
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))
	window.ContentView().AddSubviewTyped(button)
	fmt.Println("DEBUG: Created and added button")

	// Show window (BEFORE app.Run())
	window.MakeKeyAndOrderFront(window.ID)
	fmt.Println("DEBUG: Made window key and ordered front")

	// Activate app
	app.ActivateIgnoringOtherApps(true)
	fmt.Println("DEBUG: Activated app")

	// Check window state
	isVisible := window.ID.Send(objc.RegisterName("isVisible"))
	canBecomeKey := window.ID.Send(objc.RegisterName("canBecomeKeyWindow"))
	isKeyWindow := window.ID.Send(objc.RegisterName("isKeyWindow"))
	fmt.Printf("DEBUG: Window state - isVisible: %v, canBecomeKey: %v, isKeyWindow: %v\n",
		isVisible != 0, canBecomeKey != 0, isKeyWindow != 0)

	fmt.Println("✓ Window created and displayed")
	fmt.Println("Click the button! Press Cmd+Q to quit.")

	// Run application (window already created)
	app.Run()
}

// createSimpleButtonHandler creates an NSObject subclass that handles button clicks
func createSimpleButtonHandler() objc.ID {
	className := "SimpleButtonHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			clickCountSimple++
			fmt.Printf("Button clicked! Count: %d\n", clickCountSimple)
			counterLabelSimple.SetStringValue(fmt.Sprintf("Clicks: %d", clickCountSimple))
		}
		class, _ = objc.RegisterClass(className, superClass, nil, nil, []objc.MethodDef{
			{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
		})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	return handler.Send(objc.RegisterName("init"))
}
