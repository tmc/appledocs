// ClickMe with Phase 1 Protocol Interfaces
// This demonstrates type-safe application delegate using generated protocol interfaces
package main

import (
	"flag"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
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

// MyAppDelegate implements PApplicationDelegate with Phase 1 type safety
type MyAppDelegate struct {
	counterLabel appkit.TextField
	clickCount   int
}

// ApplicationDidFinishLaunching - called when app finishes launching
func (d *MyAppDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	fmt.Println("✅ Application launched via type-safe protocol!")
	fmt.Println("   Using Phase 1 PApplicationDelegate interface")
}

func (d *MyAppDelegate) HasApplicationDidFinishLaunching() bool {
	return true
}

// ApplicationWillTerminate - called before app terminates
func (d *MyAppDelegate) ApplicationWillTerminate(notification foundation.Notification) {
	fmt.Printf("✅ Application terminating. Total clicks: %d\n", d.clickCount)
}

func (d *MyAppDelegate) HasApplicationWillTerminate() bool {
	return true
}

// Compile-time verification we implement the protocol!
var _ appkit.PApplicationDelegate = (*MyAppDelegate)(nil)

func main() {
	flag.Parse()

	fmt.Println("🎉 ClickMe - Phase 1 Protocol Interface Demo")
	fmt.Println("==========================================\n")

	// Create application
	app := appkit.ApplicationClass.SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	// Create our delegate (Phase 1: still needs manual registration)
	delegate := &MyAppDelegate{}
	
	fmt.Println("📋 Delegate Features (Phase 1):")
	fmt.Println("  ✅ Implements PApplicationDelegate interface")
	fmt.Println("  ✅ Type-safe method signatures")
	fmt.Println("  ✅ Compile-time protocol verification")
	fmt.Println("  ⏳ Manual registration still required (Phase 2 will fix)")
	fmt.Println()

	// Register delegate class (Phase 2 will automate this!)
	delegateClass := registerAppDelegate(delegate)
	delegateInstance := objc.ID(delegateClass).
		Send(objc.RegisterName("alloc")).
		Send(objc.RegisterName("init"))
	
	// Set as app delegate
	app.SetDelegate(appkit.IObject(delegateInstance))

	// Create window
	window := createWindow()
	
	// Create counter label
	delegate.counterLabel = createCounterLabel(window)
	
	// Create button with click handler
	createClickButton(window, delegate)

	// Show window
	window.MakeKeyAndOrderFront(nil)
	app.ActivateIgnoringOtherApps(true)

	// E2E mode
	if *e2e {
		fmt.Println("E2E mode: terminating immediately")
		app.Terminate(nil)
		return
	}

	fmt.Println("🖱️  Click the button to see the counter increment!")
	fmt.Println("   Delegate methods will be called on launch and terminate\n")

	// Run event loop
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
	
	window.SetTitle(foundation.String{}.InitWithUTF8String("ClickMe - Phase 1 Protocols"))
	
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

func createClickButton(window appkit.Window, delegate *MyAppDelegate) {
	buttonRect := foundation.Rect{
		Origin: foundation.Point{X: 150, Y: 120},
		Size:   foundation.Size{Width: 100, Height: 40},
	}
	
	button := appkit.ButtonClass.Alloc().InitWithFrame(buttonRect)
	button.SetTitle(foundation.String{}.InitWithUTF8String("Click Me!"))
	
	// Create button handler
	handlerClass := registerButtonHandler(delegate)
	handler := objc.ID(handlerClass).
		Send(objc.RegisterName("alloc")).
		Send(objc.RegisterName("init"))
	
	button.SetTarget(appkit.IObject(handler))
	button.SetAction(objc.RegisterName("buttonClicked:"))
	
	window.ContentView().AddSubview(button)
}

// registerAppDelegate registers our delegate class (Phase 2 will automate this!)
func registerAppDelegate(delegate *MyAppDelegate) objc.Class {
	className := "MyAppDelegate"
	class := objc.GetClass(className)
	if class != 0 {
		return class
	}

	didFinishLaunching := func(self objc.ID, _cmd objc.SEL, notification objc.ID) {
		delegate.ApplicationDidFinishLaunching(foundation.Notification{ID: notification})
	}

	willTerminate := func(self objc.ID, _cmd objc.SEL, notification objc.ID) {
		delegate.ApplicationWillTerminate(foundation.Notification{ID: notification})
	}

	class, _ = objc.RegisterClass(
		className,
		objc.GetClass("NSObject"),
		[]*objc.Protocol{appkit.ApplicationDelegateProtocol},
		nil,
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("applicationDidFinishLaunching:"),
				Fn:  didFinishLaunching,
			},
			{
				Cmd: objc.RegisterName("applicationWillTerminate:"),
				Fn:  willTerminate,
			},
		},
	)

	return class
}

// registerButtonHandler registers button click handler
func registerButtonHandler(delegate *MyAppDelegate) objc.Class {
	className := "ButtonHandler"
	class := objc.GetClass(className)
	if class != 0 {
		return class
	}

	buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
		delegate.clickCount++
		fmt.Printf("Button clicked! Count: %d\n", delegate.clickCount)
		
		labelText := foundation.String{}.InitWithUTF8String(
			fmt.Sprintf("Clicks: %d", delegate.clickCount),
		)
		delegate.counterLabel.SetStringValue(labelText)
	}

	class, _ = objc.RegisterClass(
		className,
		objc.GetClass("NSObject"),
		nil,
		nil,
		[]objc.MethodDef{
			{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
		},
	)

	return class
}
