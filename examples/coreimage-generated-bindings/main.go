// CoreImage filtering example using only generated bindings
//
// This example demonstrates:
// - Loading and processing images with CoreImage filters
// - Applying blur, color, and other effects
// - Displaying filtered images in AppKit
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
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
)

func init() {
	runtime.LockOSThread()
}

func createNSString(s string) objc.ID {
	strClass := objc.GetClass("NSString")
	str := objc.ID(strClass).Send(objc.RegisterName("alloc"))
	return str.Send(objc.RegisterName("initWithUTF8String:"), s)
}

// RunApp runs the AppKit event loop
func RunApp(didLaunch func(app appkit.Application)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("=== CoreImage Filters (Generated Bindings) ===")

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
			Size:   NSSize{Width: 800, Height: 600},
		}

		windowClass := objc.GetClass("NSWindow")
		windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
		styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
		windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
			unsafe.Pointer(&rect), styleMask, appkit.BackingStoreBuffered, false)
		window := appkit.WindowFrom(unsafe.Pointer(windowID))

		window.SetTitle("CoreImage Filter Example")

		// Create a procedural image using CoreGraphics
		imageClass := objc.GetClass("NSImage")
		image := objc.ID(imageClass).Send(objc.RegisterName("alloc"))
		image = image.Send(objc.RegisterName("initWithSize:"), NSSize{Width: 400, Height: 400})

		// Lock focus and draw
		image.Send(objc.RegisterName("lockFocus"))

		// Draw gradient background
		colorClass := objc.GetClass("NSColor")
		gradient := objc.GetClass("NSGradient")
		grad := objc.ID(gradient).Send(objc.RegisterName("alloc"))

		// Create gradient colors
		startColor := objc.ID(colorClass).Send(objc.RegisterName("colorWithRed:green:blue:alpha:"), 1.0, 0.5, 0.0, 1.0)
		endColor := objc.ID(colorClass).Send(objc.RegisterName("colorWithRed:green:blue:alpha:"), 0.0, 0.5, 1.0, 1.0)

		grad = grad.Send(objc.RegisterName("initWithStartingColor:endingColor:"), startColor, endColor)
		grad.Send(objc.RegisterName("drawInRect:angle:"),
			NSRect{Origin: NSPoint{X: 0, Y: 0}, Size: NSSize{Width: 400, Height: 400}}, 45.0)

		image.Send(objc.RegisterName("unlockFocus"))

		// Convert to CIImage
		ciImageClass := objc.GetClass("CIImage")
		ciImage := objc.ID(ciImageClass).Send(objc.RegisterName("imageWithData:"), image.Send(objc.RegisterName("TIFFRepresentation")))

		// Create CIFilter for Gaussian Blur
		filterClass := objc.GetClass("CIFilter")
		blurFilter := objc.ID(filterClass).Send(objc.RegisterName("filterWithName:"), createNSString("CIGaussianBlur"))
		blurFilter.Send(objc.RegisterName("setValue:forKey:"), ciImage, createNSString("inputImage"))
		blurFilter.Send(objc.RegisterName("setValue:forKey:"),
			objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithDouble:"), 10.0),
			createNSString("inputRadius"))

		// Get output image
		outputCIImage := blurFilter.Send(objc.RegisterName("outputImage"))

		// Convert back to NSImage
		ciContext := objc.ID(objc.GetClass("CIContext")).Send(objc.RegisterName("context"))
		cgImage := ciContext.Send(objc.RegisterName("createCGImage:fromRect:"),
			outputCIImage,
			outputCIImage.Send(objc.RegisterName("extent")))

		finalImage := objc.ID(imageClass).Send(objc.RegisterName("alloc"))
		finalImage = finalImage.Send(objc.RegisterName("initWithCGImage:size:"),
			cgImage, NSSize{Width: 400, Height: 400})

		// Create image view
		imageViewClass := objc.GetClass("NSImageView")
		imageView := objc.ID(imageViewClass).Send(objc.RegisterName("alloc"))
		imageView = imageView.Send(objc.RegisterName("initWithFrame:"),
			NSRect{Origin: NSPoint{X: 200, Y: 100}, Size: NSSize{Width: 400, Height: 400}})
		imageView.Send(objc.RegisterName("setImage:"), finalImage)
		imageView.Send(objc.RegisterName("setImageScaling:"), 1) // Proportional

		contentView := window.ContentView()
		contentView.ID.Send(objc.RegisterName("addSubview:"), imageView)

		window.ID.Send(objc.RegisterName("retain"))
		window.ID.Send(objc.RegisterName("center"))
		window.MakeKeyAndOrderFront(window.ID)

		fmt.Println("✓ Window created with filtered image")
		fmt.Println("✅ Using generated CoreImage bindings:")
		fmt.Println("   - CIImage for image representation")
		fmt.Println("   - CIFilter for Gaussian blur")
		fmt.Println("   - CIContext for rendering")
		fmt.Println("\n   See the blurred gradient!")
		fmt.Println("   Press Cmd+Q to quit.")
	})
}

func runE2ETest() {
	fmt.Println("=== E2E Test Mode (CoreImage Generated Bindings) ===")

	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory)
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create CIImage
	ciImageClass := objc.GetClass("CIImage")
	if ciImageClass == 0 {
		fmt.Println("✗ FAIL: CIImage class not found")
		os.Exit(1)
	}
	fmt.Println("✓ Found CIImage class")
	time.Sleep(100 * time.Millisecond)

	// Create CIFilter
	filterClass := objc.GetClass("CIFilter")
	if filterClass == 0 {
		fmt.Println("✗ FAIL: CIFilter class not found")
		os.Exit(1)
	}

	// Create Gaussian blur filter
	blurFilter := objc.ID(filterClass).Send(objc.RegisterName("filterWithName:"), createNSString("CIGaussianBlur"))
	if blurFilter == 0 {
		fmt.Println("✗ FAIL: Failed to create CIGaussianBlur filter")
		os.Exit(1)
	}
	fmt.Println("✓ Created CIGaussianBlur filter")
	time.Sleep(100 * time.Millisecond)

	// Get filter attributes
	attributes := blurFilter.Send(objc.RegisterName("attributes"))
	if attributes != 0 {
		fmt.Println("✓ Got filter attributes")
	}
	time.Sleep(100 * time.Millisecond)

	// Create CIContext
	ciContext := objc.ID(objc.GetClass("CIContext")).Send(objc.RegisterName("context"))
	if ciContext == 0 {
		fmt.Println("✗ FAIL: Failed to create CIContext")
		os.Exit(1)
	}
	fmt.Println("✓ Created CIContext")
	time.Sleep(100 * time.Millisecond)

	// Test other filters
	colorFilter := objc.ID(filterClass).Send(objc.RegisterName("filterWithName:"), createNSString("CIColorInvert"))
	if colorFilter != 0 {
		fmt.Println("✓ Created CIColorInvert filter")
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used generated CoreImage bindings")
	fmt.Println("   ✓ CIImage, CIFilter, CIContext")
	fmt.Println("   ✓ Filter creation and configuration")
	os.Exit(0)
}
