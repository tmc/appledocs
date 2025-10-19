// Hello World using generated bindings (darwinkit-style)
//
// This example demonstrates how close we can get to darwinkit's API
// using only our generated bindings, without modifying generated code.
package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

func init() {
	// Load AppKit framework
	_, err := purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// Foundation types (would normally be in foundation package)
type Point struct {
	X, Y float64
}

type Size struct {
	Width, Height float64
}

type Rect struct {
	Origin Point
	Size   Size
}

var (
	clickCount   int
	counterLabel appkit.TextField
)

func createButtonHandler() objc.ID {
	className := "ButtonHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		buttonClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			clickCount++
			fmt.Printf("Button clicked! Count: %d\n", clickCount)
			counterLabel.SetStringValue(fmt.Sprintf("Button clicks: %d", clickCount))
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{}, []objc.MethodDef{
			{Cmd: objc.RegisterName("buttonClicked:"), Fn: buttonClicked},
		})
	}
	handler := objc.ID(class).Send(objc.RegisterName("alloc"))
	handler = handler.Send(objc.RegisterName("init"))
	return handler
}

func main() {
	fmt.Println("=== Hello World (Generated Bindings - Darwinkit Style) ===\n")

	// Create application (darwinkit: app := appkit.Application_SharedApplication())
	app := appkit.SharedApplication()
	app.SetActivationPolicy(0) // ApplicationActivationPolicyRegular

	// Create window (darwinkit: appkit.NewWindowWithContentRectStyleMaskBackingDefer)
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		unsafe.Pointer(&Rect{
			Origin: Point{X: 100, Y: 100},
			Size:   Size{Width: 400, Height: 300},
		}),
		1|2, // TitledWindowMask | ClosableWindowMask
		2,   // BackingStoreBuffered
		false,
	)

	// Set window properties (darwinkit: window.SetTitle("..."))
	window.SetTitle("Hello from Generated Bindings!")

	// Get content view (darwinkit: window.ContentView())
	// Note: window.ContentView() currently returns unsafe.Pointer, not View
	contentViewPtr := window.ID.Send(objc.RegisterName("contentView"))
	contentView := appkit.ViewFrom(unsafe.Pointer(contentViewPtr))

	// Create label (darwinkit: label := appkit.NewTextField())
	label := appkit.TextFieldClass.New()
	label.SetStringValue("This uses generated bindings!")
	label.ID.Send(objc.RegisterName("setFrameOrigin:"), Point{X: 50, Y: 200})
	label.ID.Send(objc.RegisterName("setFrameSize:"), Size{Width: 300, Height: 50})
	label.SetEditable(false)
	label.SetBordered(false)
	label.SetBackgroundColor(nil)

	// Add label to content view (darwinkit: contentView.AddSubview(label))
	contentView.AddSubview(unsafe.Pointer(label.ID))

	// Create counter label
	counterLabel = appkit.TextFieldClass.New()
	counterLabel.SetStringValue("Button clicks: 0")
	counterLabel.ID.Send(objc.RegisterName("setFrameOrigin:"), Point{X: 50, Y: 80})
	counterLabel.ID.Send(objc.RegisterName("setFrameSize:"), Size{Width: 300, Height: 30})
	counterLabel.SetEditable(false)
	counterLabel.SetBordered(false)
	counterLabel.SetBackgroundColor(nil)
	counterLabel.SetAlignment(2) // TextAlignmentCenter

	contentView.AddSubview(unsafe.Pointer(counterLabel.ID))

	// Create button (darwinkit: button := appkit.NewButtonWithTitle("Click Me!"))
	button := appkit.ButtonClass.New()
	button.SetTitle("Click Me!")
	button.ID.Send(objc.RegisterName("setFrameOrigin:"), Point{X: 150, Y: 130})
	button.ID.Send(objc.RegisterName("setFrameSize:"), Size{Width: 100, Height: 40})
	button.SetButtonType(0) // ButtonTypeMomentaryLight
	button.SetBezelStyle(1) // BezelStyleRounded

	// Set button action (darwinkit: button.SetTarget/SetAction)
	buttonHandler := createButtonHandler()
	button.SetTarget(buttonHandler)
	button.SetAction(objc.RegisterName("buttonClicked:"))

	contentView.AddSubview(unsafe.Pointer(button.ID))

	// Show window (darwinkit: window.MakeKeyAndOrderFront(nil))
	window.MakeKeyAndOrderFront(0)

	// Activate app (darwinkit: app.ActivateIgnoringOtherApps(true))
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ Window created using generated bindings")
	fmt.Println("   API style closely matches darwinkit!")
	fmt.Println("   Click the button to see the counter increment!")
	fmt.Println("   Close window or press Cmd+Q to quit\n")

	// Run application (darwinkit: app.Run())
	app.Run()
}
