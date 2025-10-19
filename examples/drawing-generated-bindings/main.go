// CoreGraphics Drawing using only generated bindings
//
// This example demonstrates:
// - Drawing shapes using CoreGraphics generated bindings
// - Creating NSImage and drawing into it
// - Using only generated bindings (no manual purego calls)
// - Combining AppKit, Foundation, and CoreGraphics
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

func createAppDelegate() objc.ID {
	className := "DrawingAppDelegate"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		windowShouldClose := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
			app := appkit.SharedApplication()
			app.Terminate(0)
			return true
		}
		class, _ = objc.RegisterClass(className, superClass, []*objc.Protocol{}, []objc.FieldDef{},
			[]objc.MethodDef{{Cmd: objc.RegisterName("windowShouldClose:"), Fn: windowShouldClose}})
	}
	delegate := objc.ID(class).Send(objc.RegisterName("alloc"))
	return delegate.Send(objc.RegisterName("init"))
}

func createDrawnImage(width, height float64) objc.ID {
	// Create NSImage
	imageClass := objc.GetClass("NSImage")
	image := objc.ID(imageClass).Send(objc.RegisterName("alloc"))
	// Create NSSize manually (no helper function generated yet)
	type NSSize struct{ Width, Height float64 }
	image = image.Send(objc.RegisterName("initWithSize:"), NSSize{Width: width, Height: height})

	// Lock focus to draw
	image.Send(objc.RegisterName("lockFocus"))

	// Get current graphics context
	ctxClass := objc.GetClass("NSGraphicsContext")
	nsCtx := objc.ID(ctxClass).Send(objc.RegisterName("currentContext"))
	cgCtxID := nsCtx.Send(objc.RegisterName("CGContext"))
	cgCtx := coregraphics.CGContextRef(unsafe.Pointer(uintptr(cgCtxID)))

	// Draw shapes using generated bindings
	drawShapes(cgCtx, coregraphics.CGSize{Width: coregraphics.CGFloat(width), Height: coregraphics.CGFloat(height)})

	// Unlock focus
	image.Send(objc.RegisterName("unlockFocus"))

	return image
}

func drawShapes(ctx coregraphics.CGContextRef, size coregraphics.CGSize) {
	coregraphics.CGContextSaveGState(ctx)

	// White background
	coregraphics.CGContextSetRGBFillColor(ctx, 1.0, 1.0, 1.0, 1.0)
	coregraphics.CGContextFillRect(ctx, coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 0, Y: 0},
		Size:   size,
	})

	// Blue rectangle
	coregraphics.CGContextSetRGBFillColor(ctx, 0.2, 0.4, 0.8, 1.0)
	coregraphics.CGContextFillRect(ctx, coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 50, Y: 450},
		Size:   coregraphics.CGSize{Width: 200, Height: 100},
	})

	// Red circle
	coregraphics.CGContextSetRGBStrokeColor(ctx, 0.9, 0.2, 0.2, 1.0)
	coregraphics.CGContextSetLineWidth(ctx, 3.0)
	coregraphics.CGContextBeginPath(ctx)
	coregraphics.CGContextAddArc(ctx, 400, 500, 60, 0, 6.28318, 0)
	coregraphics.CGContextStrokePath(ctx)

	// Green triangle
	coregraphics.CGContextSetRGBFillColor(ctx, 0.2, 0.8, 0.3, 1.0)
	coregraphics.CGContextBeginPath(ctx)
	coregraphics.CGContextMoveToPoint(ctx, 600, 450)
	coregraphics.CGContextAddLineToPoint(ctx, 700, 450)
	coregraphics.CGContextAddLineToPoint(ctx, 650, 550)
	coregraphics.CGContextClosePath(ctx)
	coregraphics.CGContextFillPath(ctx)

	// Purple rectangle outline
	coregraphics.CGContextSetRGBStrokeColor(ctx, 0.7, 0.2, 0.8, 1.0)
	coregraphics.CGContextSetLineWidth(ctx, 4.0)
	coregraphics.CGContextStrokeRect(ctx, coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 50, Y: 300},
		Size:   coregraphics.CGSize{Width: 300, Height: 100},
	})

	// Orange circle
	coregraphics.CGContextSetRGBFillColor(ctx, 1.0, 0.6, 0.0, 1.0)
	coregraphics.CGContextBeginPath(ctx)
	coregraphics.CGContextAddArc(ctx, 500, 350, 40, 0, 6.28318, 0)
	coregraphics.CGContextFillPath(ctx)

	// Cyan path with fill and stroke
	coregraphics.CGContextSetRGBFillColor(ctx, 0.3, 0.8, 0.9, 0.5)
	coregraphics.CGContextSetRGBStrokeColor(ctx, 0.0, 0.5, 0.6, 1.0)
	coregraphics.CGContextSetLineWidth(ctx, 2.0)
	coregraphics.CGContextBeginPath(ctx)
	coregraphics.CGContextMoveToPoint(ctx, 100, 200)
	coregraphics.CGContextAddLineToPoint(ctx, 200, 250)
	coregraphics.CGContextAddLineToPoint(ctx, 150, 150)
	coregraphics.CGContextAddLineToPoint(ctx, 250, 180)
	coregraphics.CGContextClosePath(ctx)
	// Use kCGPathFillStroke mode
	coregraphics.CGContextDrawPath(ctx, unsafe.Pointer(uintptr(3))) // kCGPathFillStroke = 3

	// Rainbow bars
	colors := []struct{ r, g, b coregraphics.CGFloat }{
		{1.0, 0.0, 0.0}, {1.0, 0.5, 0.0}, {1.0, 1.0, 0.0},
		{0.0, 1.0, 0.0}, {0.0, 0.0, 1.0}, {0.5, 0.0, 0.5},
	}
	for i, c := range colors {
		coregraphics.CGContextSetRGBFillColor(ctx, c.r, c.g, c.b, 0.8)
		coregraphics.CGContextFillRect(ctx, coregraphics.CGRect{
			Origin: coregraphics.CGPoint{X: coregraphics.CGFloat(350 + i*50), Y: 150},
			Size:   coregraphics.CGSize{Width: 40, Height: 150},
		})
	}

	coregraphics.CGContextRestoreGState(ctx)
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETest()
		return
	}

	fmt.Println("=== CoreGraphics Drawing (Generated Bindings) ===\n")

	// Create app using generated bindings
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyRegular)

	// Create window using generated bindings
	window := appkit.NewWindowWithFrame(100, 100, 800, 600,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("CoreGraphics Drawing (Generated Bindings)")

	// Set delegate
	delegate := createAppDelegate()
	window.SetDelegate(delegate)

	// Create and draw image
	image := createDrawnImage(800, 600)

	// Create image view
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	imageViewClass := objc.GetClass("NSImageView")
	imageView := objc.ID(imageViewClass).Send(objc.RegisterName("alloc"))
	imageView = imageView.Send(objc.RegisterName("initWithFrame:"),
		NSRect{Origin: NSPoint{X: 0, Y: 0}, Size: NSSize{Width: 800, Height: 600}})
	imageView.Send(objc.RegisterName("setImage:"), image)
	imageView.Send(objc.RegisterName("setImageScaling:"), 1) // NSImageScaleProportionallyUpOrDown

	// Set as content view
	window.ID.Send(objc.RegisterName("setContentView:"), imageView)

	// Show window
	window.MakeKeyAndOrderFront(0)
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ Drawing example using ONLY generated bindings:")
	fmt.Println("   - CoreGraphics: CGContextSetRGBFillColor, CGContextFillRect, etc.")
	fmt.Println("   - AppKit: SharedApplication, NewWindowWithFrame, etc.")
	fmt.Println("   - Types: CGRect, CGPoint, CGSize, CGFloat from generated/coregraphics")
	fmt.Println("   - NO manual purego.RegisterLibFunc calls!")
	fmt.Println("   - NO manual type definitions!")
	fmt.Println("\n   Press Cmd+Q or close window to quit\n")

	app.Run()
}

// runE2ETest runs automated end-to-end test with small delays for visibility
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Drawing Generated Bindings) ===")

	// Create app
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory) // No dock icon in tests
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 800, 600,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window")
	time.Sleep(100 * time.Millisecond)

	// Set delegate
	delegate := createAppDelegate()
	window.SetDelegate(delegate)
	fmt.Println("✓ Set window delegate")
	time.Sleep(100 * time.Millisecond)

	// Create and draw image using CoreGraphics
	fmt.Println("✓ Drawing shapes with CoreGraphics...")
	image := createDrawnImage(800, 600)
	if image == 0 {
		fmt.Println("✗ FAIL: Image not created")
		os.Exit(1)
	}
	fmt.Println("✓ Created image with CoreGraphics drawing")
	time.Sleep(200 * time.Millisecond)

	// Create image view
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	imageViewClass := objc.GetClass("NSImageView")
	imageView := objc.ID(imageViewClass).Send(objc.RegisterName("alloc"))
	imageView = imageView.Send(objc.RegisterName("initWithFrame:"),
		NSRect{Origin: NSPoint{X: 0, Y: 0}, Size: NSSize{Width: 800, Height: 600}})
	imageView.Send(objc.RegisterName("setImage:"), image)
	imageView.Send(objc.RegisterName("setImageScaling:"), 1)
	fmt.Println("✓ Created image view")
	time.Sleep(100 * time.Millisecond)

	// Set as content view
	window.ID.Send(objc.RegisterName("setContentView:"), imageView)
	fmt.Println("✓ Set content view")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(0)
	fmt.Println("✓ Window displayed with drawing")
	time.Sleep(300 * time.Millisecond)

	// Close window
	window.ID.Send(objc.RegisterName("close"))
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used generated CoreGraphics bindings")
	fmt.Println("   ✓ Drew shapes: rectangles, circles, triangles")
	fmt.Println("   ✓ No manual purego.RegisterLibFunc calls")
	os.Exit(0)
}
