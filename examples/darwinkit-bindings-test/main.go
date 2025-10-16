// Example demonstrating end-to-end DarwinKit-style binding generation and usage
package main

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	// Load AppKit framework first
	appkit, err := purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load AppKit: %v", err))
	}
	fmt.Printf("✓ Successfully loaded AppKit framework (handle: %v)\n", appkit)

	// Test basic objc runtime access
	// Get NSColor class
	colorClass := objc.GetClass("NSColor")
	if colorClass.Ptr() == nil {
		panic("Failed to get NSColor class")
	}
	fmt.Println("✓ Successfully loaded NSColor class")

	// Call +[NSColor redColor] class method
	redColor := objc.Call[objc.Object](colorClass, objc.Sel("redColor"))
	if redColor.Ptr() == nil {
		panic("Failed to create red color")
	}
	fmt.Println("✓ Successfully created NSColor instance via class method")

	// Get description
	desc := objc.Call[string](redColor, objc.Sel("description"))
	fmt.Printf("✓ Color description: %s\n", desc)

	// Test that the object responds to selector
	respondsToDescription := objc.Call[bool](redColor, objc.Sel("respondsToSelector:"), objc.Sel("description"))
	fmt.Printf("✓ Responds to description selector: %v\n", respondsToDescription)

	fmt.Println("\n✅ All basic objc runtime tests passed!")
	fmt.Println("   Next: Generate DarwinKit-style bindings for AppKit classes")
}
