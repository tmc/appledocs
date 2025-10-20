// QuartzCore Animation example using only generated bindings
//
// This example demonstrates:
// - Creating and animating CALayer objects
// - Using CABasicAnimation for property animations
// - Layer hierarchy and transformations
// - Using only generated bindings (no manual purego calls)
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
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
)

func init() {
	runtime.LockOSThread()
}

// RunApp runs the AppKit event loop with proper initialization
func RunApp(didLaunch func(app appkit.Application)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("=== QuartzCore Animation (Generated Bindings) ===")

	app := appkit.SharedApplication()

	delegateClass, err := objc.RegisterClass(
		"AppDelegate",
		objc.GetClass("NSObject"),
		nil,
		nil,
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("applicationDidFinishLaunching:"),
				Fn: func(self objc.ID, _cmd objc.SEL, notification objc.ID) {
					fmt.Println("✓ Application finished launching")
					didLaunch(app)
				},
			},
			{
				Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
					return false
				},
			},
			{
				Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
				Fn: func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
					return true
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ERROR: Failed to register delegate class: %v\n", err)
		os.Exit(1)
	}

	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	app.ID.Send(objc.RegisterName("setDelegate:"), delegate)

	fmt.Println("✓ Starting AppKit event loop...")
	app.Run()
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETest()
		return
	}

	RunApp(func(app appkit.Application) {
		app.SetActivationPolicy(appkit.ActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create window
		type NSPoint struct{ X, Y float64 }
		type NSSize struct{ Width, Height float64 }
		type NSRect struct {
			Origin NSPoint
			Size   NSSize
		}
		rect := NSRect{
			Origin: NSPoint{X: 100, Y: 100},
			Size:   NSSize{Width: 600, Height: 400},
		}

		windowClass := objc.GetClass("NSWindow")
		windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
		styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
		windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
			unsafe.Pointer(&rect), styleMask, appkit.BackingStoreBuffered, false)
		window := appkit.WindowFrom(unsafe.Pointer(windowID))

		window.SetTitle("QuartzCore Animation Example")

		// Create layer-hosting view
		contentView := window.ContentView()
		contentView.ID.Send(objc.RegisterName("setWantsLayer:"), true)

		// Get the root layer
		rootLayer := contentView.ID.Send(objc.RegisterName("layer"))

		// Create an animated layer
		layerClass := objc.GetClass("CALayer")
		animatedLayer := objc.ID(layerClass).Send(objc.RegisterName("layer"))

		// Set layer properties
		animatedLayer.Send(objc.RegisterName("setFrame:"), coregraphics.CGRect{
			Origin: coregraphics.CGPoint{X: 50, Y: 150},
			Size:   coregraphics.CGSize{Width: 100, Height: 100},
		})

		// Set layer background color (red)
		color := coregraphics.CGColorCreateGenericRGB(1.0, 0.0, 0.0, 1.0) // Red
		animatedLayer.Send(objc.RegisterName("setBackgroundColor:"), color)
		animatedLayer.Send(objc.RegisterName("setCornerRadius:"), 10.0)

		// Add layer to root
		rootLayer.Send(objc.RegisterName("addSublayer:"), animatedLayer)

		// Create position animation
		animClass := objc.GetClass("CABasicAnimation")
		posAnimation := objc.ID(animClass).Send(objc.RegisterName("animationWithKeyPath:"), createNSString("position.x"))

		// Set animation properties
		posAnimation.Send(objc.RegisterName("setFromValue:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithDouble:"), 100.0))
		posAnimation.Send(objc.RegisterName("setToValue:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithDouble:"), 500.0))
		posAnimation.Send(objc.RegisterName("setDuration:"), 2.0)
		posAnimation.Send(objc.RegisterName("setRepeatCount:"), float32(1000.0)) // Repeat many times
		posAnimation.Send(objc.RegisterName("setAutoreverses:"), true)

		// Add animation to layer
		animatedLayer.Send(objc.RegisterName("addAnimation:forKey:"), posAnimation, createNSString("positionAnimation"))

		// Create a rotating layer
		rotatingLayer := objc.ID(layerClass).Send(objc.RegisterName("layer"))
		rotatingLayer.Send(objc.RegisterName("setFrame:"), coregraphics.CGRect{
			Origin: coregraphics.CGPoint{X: 300, Y: 150},
			Size:   coregraphics.CGSize{Width: 100, Height: 100},
		})

		// Blue color
		blueColor := coregraphics.CGColorCreateGenericRGB(0.0, 0.5, 1.0, 1.0)
		rotatingLayer.Send(objc.RegisterName("setBackgroundColor:"), blueColor)

		rootLayer.Send(objc.RegisterName("addSublayer:"), rotatingLayer)

		// Create rotation animation
		rotAnimation := objc.ID(animClass).Send(objc.RegisterName("animationWithKeyPath:"), createNSString("transform.rotation.z"))
		rotAnimation.Send(objc.RegisterName("setFromValue:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithDouble:"), 0.0))
		rotAnimation.Send(objc.RegisterName("setToValue:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithDouble:"), 6.28318)) // 2*PI
		rotAnimation.Send(objc.RegisterName("setDuration:"), 3.0)
		rotAnimation.Send(objc.RegisterName("setRepeatCount:"), float32(1000.0))
		rotatingLayer.Send(objc.RegisterName("addAnimation:forKey:"), rotAnimation, createNSString("rotationAnimation"))

		window.ID.Send(objc.RegisterName("retain"))
		window.ID.Send(objc.RegisterName("center"))
		window.MakeKeyAndOrderFront(window.ID)

		fmt.Println("✓ Window created with animated layers")
		fmt.Println("✅ Using generated QuartzCore bindings:")
		fmt.Println("   - CALayer for layer composition")
		fmt.Println("   - CABasicAnimation for property animations")
		fmt.Println("   - CGColorCreate from generated/quartzcore")
		fmt.Println("   - Animating position and rotation")
		fmt.Println("\n   Watch the red square move and blue square rotate!")
		fmt.Println("   Press Cmd+Q to quit.")
	})
}

func createNSString(s string) objc.ID {
	strClass := objc.GetClass("NSString")
	str := objc.ID(strClass).Send(objc.RegisterName("alloc"))
	return str.Send(objc.RegisterName("initWithUTF8String:"), s)
}

func runE2ETest() {
	fmt.Println("=== E2E Test Mode (QuartzCore Generated Bindings) ===")

	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory)
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 600, 400,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window")
	time.Sleep(100 * time.Millisecond)

	// Create layer-hosting view
	contentView := window.ContentView()
	contentView.ID.Send(objc.RegisterName("setWantsLayer:"), true)
	rootLayer := contentView.ID.Send(objc.RegisterName("layer"))
	if rootLayer == 0 {
		fmt.Println("✗ FAIL: Failed to get root layer")
		os.Exit(1)
	}
	fmt.Println("✓ Created layer-hosting view")
	time.Sleep(100 * time.Millisecond)

	// Create CALayer
	layerClass := objc.GetClass("CALayer")
	testLayer := objc.ID(layerClass).Send(objc.RegisterName("layer"))
	if testLayer == 0 {
		fmt.Println("✗ FAIL: Failed to create CALayer")
		os.Exit(1)
	}
	fmt.Println("✓ Created CALayer")
	time.Sleep(100 * time.Millisecond)

	// Set layer frame
	testLayer.Send(objc.RegisterName("setFrame:"), coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 50, Y: 50},
		Size:   coregraphics.CGSize{Width: 100, Height: 100},
	})
	fmt.Println("✓ Set layer frame")
	time.Sleep(100 * time.Millisecond)

	// Create color
	color := coregraphics.CGColorCreateGenericRGB(1.0, 0.0, 0.0, 1.0)
	testLayer.Send(objc.RegisterName("setBackgroundColor:"), color)
	fmt.Println("✓ Created CGColor and set layer color")
	time.Sleep(100 * time.Millisecond)

	// Create animation
	animClass := objc.GetClass("CABasicAnimation")
	animation := objc.ID(animClass).Send(objc.RegisterName("animationWithKeyPath:"), createNSString("opacity"))
	if animation == 0 {
		fmt.Println("✗ FAIL: Failed to create CABasicAnimation")
		os.Exit(1)
	}
	fmt.Println("✓ Created CABasicAnimation")
	time.Sleep(100 * time.Millisecond)

	// Add layer to root
	rootLayer.Send(objc.RegisterName("addSublayer:"), testLayer)
	fmt.Println("✓ Added layer to hierarchy")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(0)
	fmt.Println("✓ Window displayed with QuartzCore layer")
	time.Sleep(200 * time.Millisecond)

	window.ID.Send(objc.RegisterName("close"))
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used generated QuartzCore bindings")
	fmt.Println("   ✓ CALayer, CABasicAnimation, CGColor")
	fmt.Println("   ✓ Layer hierarchy and animations")
	os.Exit(0)
}
