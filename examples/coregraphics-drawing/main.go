// CoreGraphics Drawing Example
//
// This example demonstrates actual CoreGraphics drawing operations.
// It creates an NSImage, draws into it using CGContext, and displays it.
//
// Key concepts demonstrated:
// - Loading CoreGraphics framework
// - Binding CG functions with purego
// - Drawing into an NSImage using CGContext
// - Drawing rectangles, circles, and paths
// - Using CGContext fill and stroke operations
package main

import (
	"fmt"
	"strings"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// CoreGraphics types
type (
	CGFloat           float64
	CGContextRef      unsafe.Pointer
	CGColorSpaceRef   unsafe.Pointer
	CGPoint           struct{ X, Y CGFloat }
	CGSize            struct{ Width, Height CGFloat }
	CGRect            struct{ Origin CGPoint; Size CGSize }
	CGPathDrawingMode int32
)

// CGPathDrawingMode constants
const (
	kCGPathFill       CGPathDrawingMode = 0
	kCGPathStroke     CGPathDrawingMode = 2
	kCGPathFillStroke CGPathDrawingMode = 3
)

// CoreGraphics function bindings
var (
	cgLib uintptr

	// Context state
	CGContextSaveGState    func(c CGContextRef)
	CGContextRestoreGState func(c CGContextRef)

	// Color operations
	CGContextSetRGBFillColor   func(c CGContextRef, red, green, blue, alpha CGFloat)
	CGContextSetRGBStrokeColor func(c CGContextRef, red, green, blue, alpha CGFloat)
	CGContextSetLineWidth      func(c CGContextRef, width CGFloat)

	// Drawing operations
	CGContextFillRect       func(c CGContextRef, rect CGRect)
	CGContextStrokeRect     func(c CGContextRef, rect CGRect)
	CGContextBeginPath      func(c CGContextRef)
	CGContextClosePath      func(c CGContextRef)
	CGContextAddArc         func(c CGContextRef, x, y, radius, startAngle, endAngle CGFloat, clockwise int32)
	CGContextDrawPath       func(c CGContextRef, mode CGPathDrawingMode)
	CGContextMoveToPoint    func(c CGContextRef, x, y CGFloat)
	CGContextAddLineToPoint func(c CGContextRef, x, y CGFloat)
)

func init() {
	runtime.LockOSThread()

	// Load CoreGraphics framework
	var err error
	cgLib, err = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load CoreGraphics: %v", err))
	}

	// Register CG functions
	purego.RegisterLibFunc(&CGContextSaveGState, cgLib, "CGContextSaveGState")
	purego.RegisterLibFunc(&CGContextRestoreGState, cgLib, "CGContextRestoreGState")
	purego.RegisterLibFunc(&CGContextSetRGBFillColor, cgLib, "CGContextSetRGBFillColor")
	purego.RegisterLibFunc(&CGContextSetRGBStrokeColor, cgLib, "CGContextSetRGBStrokeColor")
	purego.RegisterLibFunc(&CGContextSetLineWidth, cgLib, "CGContextSetLineWidth")
	purego.RegisterLibFunc(&CGContextFillRect, cgLib, "CGContextFillRect")
	purego.RegisterLibFunc(&CGContextStrokeRect, cgLib, "CGContextStrokeRect")
	purego.RegisterLibFunc(&CGContextBeginPath, cgLib, "CGContextBeginPath")
	purego.RegisterLibFunc(&CGContextClosePath, cgLib, "CGContextClosePath")
	purego.RegisterLibFunc(&CGContextAddArc, cgLib, "CGContextAddArc")
	purego.RegisterLibFunc(&CGContextDrawPath, cgLib, "CGContextDrawPath")
	purego.RegisterLibFunc(&CGContextMoveToPoint, cgLib, "CGContextMoveToPoint")
	purego.RegisterLibFunc(&CGContextAddLineToPoint, cgLib, "CGContextAddLineToPoint")
}

func main() {
	fmt.Println("=== CoreGraphics Drawing Example ===\n")

	// Create application
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	// Create window
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		foundation.Rect{
			Origin: foundation.Point{X: 100, Y: 100},
			Size:   foundation.Size{Width: 800, Height: 600},
		},
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable,
		appkit.BackingStoreBuffered,
		false,
	)
	window.SetTitle("CoreGraphics Drawing Demo")

	// Create an image and draw into it
	image := createDrawnImage()

	// Create image view to display the drawn image
	imageView := appkit.NewImageView()
	imageView.SetImage(image)
	imageView.SetFrame(foundation.Rect{
		Origin: foundation.Point{X: 0, Y: 0},
		Size:   foundation.Size{Width: 800, Height: 600},
	})
	imageView.SetImageScaling(appkit.ImageScaleProportionallyUpOrDown)

	window.SetContentView(imageView)

	// Show window
	window.MakeKeyAndOrderFront(nil)
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ CoreGraphics drawing example running")
	fmt.Println("   Drawing shapes using CGContext")
	fmt.Println("   - Blue filled rectangle")
	fmt.Println("   - Red circle outline")
	fmt.Println("   - Green triangle")
	fmt.Println("   - And more shapes!")
	fmt.Println("   Press Cmd+Q to quit\n")

	// Run application
	app.Run()
}

func createDrawnImage() appkit.Image {
	width := 800.0
	height := 600.0

	// Create NSImage
	size := foundation.Size{Width: width, Height: height}
	img := appkit.NewImageWithSize(size)

	// Lock focus to draw into the image
	objc.ID(img.Ptr()).Send(objc.RegisterName("lockFocus"))

	// Get the current NSGraphicsContext and extract CGContext
	nsCtx := appkit.GraphicsContext_CurrentContext()
	cgContextID := objc.ID(nsCtx.Ptr()).Send(objc.RegisterName("CGContext"))
	cgContext := CGContextRef(unsafe.Pointer(uintptr(cgContextID)))

	// Perform CoreGraphics drawing
	drawShapes(cgContext, CGSize{Width: CGFloat(width), Height: CGFloat(height)})

	// Unlock focus
	objc.ID(img.Ptr()).Send(objc.RegisterName("unlockFocus"))

	return img
}

func drawShapes(cgContext CGContextRef, size CGSize) {
	// Save graphics state
	CGContextSaveGState(cgContext)

	// Draw white background
	CGContextSetRGBFillColor(cgContext, 1.0, 1.0, 1.0, 1.0)
	CGContextFillRect(cgContext, CGRect{
		Origin: CGPoint{X: 0, Y: 0},
		Size:   size,
	})

	// Draw blue filled rectangle
	CGContextSetRGBFillColor(cgContext, 0.2, 0.4, 0.8, 1.0)
	CGContextFillRect(cgContext, CGRect{
		Origin: CGPoint{X: 50, Y: 450},
		Size:   CGSize{Width: 200, Height: 100},
	})

	// Draw red stroked circle
	CGContextSetRGBStrokeColor(cgContext, 0.9, 0.2, 0.2, 1.0)
	CGContextSetLineWidth(cgContext, 3.0)
	CGContextBeginPath(cgContext)
	CGContextAddArc(cgContext, 400, 500, 60, 0, 6.28318, 0) // 2π radians = full circle
	CGContextDrawPath(cgContext, kCGPathStroke)

	// Draw green filled triangle
	CGContextSetRGBFillColor(cgContext, 0.2, 0.8, 0.3, 1.0)
	CGContextBeginPath(cgContext)
	CGContextMoveToPoint(cgContext, 600, 450)
	CGContextAddLineToPoint(cgContext, 700, 450)
	CGContextAddLineToPoint(cgContext, 650, 550)
	CGContextClosePath(cgContext)
	CGContextDrawPath(cgContext, kCGPathFill)

	// Draw purple stroked rectangle
	CGContextSetRGBStrokeColor(cgContext, 0.7, 0.2, 0.8, 1.0)
	CGContextSetLineWidth(cgContext, 4.0)
	CGContextStrokeRect(cgContext, CGRect{
		Origin: CGPoint{X: 50, Y: 300},
		Size:   CGSize{Width: 300, Height: 100},
	})

	// Draw orange filled circle
	CGContextSetRGBFillColor(cgContext, 1.0, 0.6, 0.0, 1.0)
	CGContextBeginPath(cgContext)
	CGContextAddArc(cgContext, 500, 350, 40, 0, 6.28318, 0)
	CGContextDrawPath(cgContext, kCGPathFill)

	// Draw cyan stroked and filled path
	CGContextSetRGBFillColor(cgContext, 0.3, 0.8, 0.9, 0.5)
	CGContextSetRGBStrokeColor(cgContext, 0.0, 0.5, 0.6, 1.0)
	CGContextSetLineWidth(cgContext, 2.0)
	CGContextBeginPath(cgContext)
	CGContextMoveToPoint(cgContext, 100, 200)
	CGContextAddLineToPoint(cgContext, 200, 250)
	CGContextAddLineToPoint(cgContext, 150, 150)
	CGContextAddLineToPoint(cgContext, 250, 180)
	CGContextClosePath(cgContext)
	CGContextDrawPath(cgContext, kCGPathFillStroke)

	// Draw rainbow colored rectangles
	colors := []struct{ r, g, b CGFloat }{
		{1.0, 0.0, 0.0}, // Red
		{1.0, 0.5, 0.0}, // Orange
		{1.0, 1.0, 0.0}, // Yellow
		{0.0, 1.0, 0.0}, // Green
		{0.0, 0.0, 1.0}, // Blue
		{0.5, 0.0, 0.5}, // Purple
	}
	for i, color := range colors {
		CGContextSetRGBFillColor(cgContext, color.r, color.g, color.b, 0.8)
		CGContextFillRect(cgContext, CGRect{
			Origin: CGPoint{X: CGFloat(350 + i*50), Y: 150},
			Size:   CGSize{Width: 40, Height: 150},
		})
	}

	// Draw info box with border
	CGContextSetRGBFillColor(cgContext, 0.95, 0.95, 0.95, 1.0)
	CGContextFillRect(cgContext, CGRect{
		Origin: CGPoint{X: 50, Y: 50},
		Size:   CGSize{Width: 700, Height: 60},
	})
	CGContextSetRGBStrokeColor(cgContext, 0.3, 0.3, 0.3, 1.0)
	CGContextSetLineWidth(cgContext, 1.0)
	CGContextStrokeRect(cgContext, CGRect{
		Origin: CGPoint{X: 50, Y: 50},
		Size:   CGSize{Width: 700, Height: 60},
	})

	// Restore graphics state
	CGContextRestoreGState(cgContext)
}
