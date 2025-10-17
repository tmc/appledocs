package appkit_test

import (
	"fmt"

	"github.com/tmc/appledocs/generated/appkit"
)

// ExampleSharedApplication demonstrates getting the NSApplication singleton.
func ExampleSharedApplication() {
	app := appkit.SharedApplication()
	fmt.Printf("Got application instance: %T\n", app)

	// Output:
	// Got application instance: appkit.Application
}

// ExampleApplication_SetActivationPolicy demonstrates setting activation policy.
func ExampleApplication_SetActivationPolicy() {
	app := appkit.SharedApplication()

	// Set to regular app (shows in Dock)
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	fmt.Println("Set activation policy to Regular")
	fmt.Println("Available policies:")
	fmt.Println("- ActivationPolicyRegular: Normal app with Dock icon")
	fmt.Println("- ActivationPolicyAccessory: No Dock icon")
	fmt.Println("- ActivationPolicyProhibited: Cannot activate")

	// Output:
	// Set activation policy to Regular
	// Available policies:
	// - ActivationPolicyRegular: Normal app with Dock icon
	// - ActivationPolicyAccessory: No Dock icon
	// - ActivationPolicyProhibited: Cannot activate
}

// ExampleWindow_SetTitle demonstrates setting a window's title.
func ExampleWindow_SetTitle() {
	window := appkit.NewWindow()

	window.SetTitle("My Application Window")

	fmt.Println("Window title set successfully")

	// Output:
	// Window title set successfully
}

// ExampleWindow_ContentView demonstrates getting a window's content view.
func ExampleWindow_ContentView() {
	window := appkit.NewWindow()

	contentView := window.ContentView()

	fmt.Printf("Got content view: %T\n", contentView)

	// Output:
	// Got content view: appkit.View
}

// ExampleRunApp demonstrates how to use RunApp to create a simple macOS application.
// This example shows the minimal code needed to create and display a window.
//
// Note: This example will block when run interactively. In automated tests, it's
// typically skipped or run with a timeout.
func ExampleRunApp() {
	// This example demonstrates the RunApp API but doesn't actually run
	// to avoid blocking tests. In a real application, this would be in main():
	//
	// func main() {
	//     appkit.RunApp(func(app appkit.Application) {
	//         window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(...)
	//         window.MakeKeyAndOrderFront(0)
	//     })
	// }

	fmt.Println("RunApp provides a simple way to create macOS applications")
	fmt.Println("It handles all the boilerplate:")
	fmt.Println("- Thread locking")
	fmt.Println("- Getting shared application")
	fmt.Println("- Setting activation policy")
	fmt.Println("- Finishing launch")
	fmt.Println("- Running the event loop")

	// Output:
	// RunApp provides a simple way to create macOS applications
	// It handles all the boilerplate:
	// - Thread locking
	// - Getting shared application
	// - Setting activation policy
	// - Finishing launch
	// - Running the event loop
}
