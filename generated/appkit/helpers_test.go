package appkit_test

import (
	"fmt"
)

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
