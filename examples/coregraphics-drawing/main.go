// CoreGraphics Drawing Example
//
// This example demonstrates actual CoreGraphics drawing operations.
// It creates a custom NSView that uses CGContext to draw shapes, paths, and colors.
//
// Key concepts demonstrated:
// - Loading CoreGraphics framework
// - Binding CG functions with purego
// - Creating a custom NSView with drawRect:
// - Drawing rectangles, circles, and paths
// - Using CGContext fill and stroke operations
package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// CoreGraphics types from generated bindings
type (
	CGFloat           float64
	CGContextRef      unsafe.Pointer
	CGColorSpaceRef   unsafe.Pointer
	CGColorRef        unsafe.Pointer
	CGPoint           struct{ X, Y CGFloat }
	CGSize            struct{ Width, Height CGFloat }
	CGRect            struct{ Origin CGPoint; Size CGSize }
	CGPathRef         unsafe.Pointer
	CGPathDrawingMode int32
)

// CGPathDrawingMode constants
const (
	kCGPathFill        CGPathDrawingMode = 0
	kCGPathStroke      CGPathDrawingMode = 2
	kCGPathFillStroke  CGPathDrawingMode = 3
)

// CoreGraphics function bindings
var (
	cgLib uintptr

	// Context state
	CGContextSaveGState    func(c CGContextRef)
	CGContextRestoreGState func(c CGContextRef)

	// Color operations
	CGColorSpaceCreateDeviceRGB func() CGColorSpaceRef
	CGColorSpaceRelease         func(space CGColorSpaceRef)
	CGContextSetRGBFillColor    func(c CGContextRef, red, green, blue, alpha CGFloat)
	CGContextSetRGBStrokeColor  func(c CGContextRef, red, green, blue, alpha CGFloat)
	CGContextSetLineWidth       func(c CGContextRef, width CGFloat)

	// Drawing operations
	CGContextFillRect   func(c CGContextRef, rect CGRect)
	CGContextStrokeRect func(c CGContextRef, rect CGRect)
	CGContextBeginPath  func(c CGContextRef)
	CGContextClosePath  func(c CGContextRef)
	CGContextAddArc     func(c CGContextRef, x, y, radius, startAngle, endAngle CGFloat, clockwise int32)
	CGContextAddRect    func(c CGContextRef, rect CGRect)
	CGContextDrawPath   func(c CGContextRef, mode CGPathDrawingMode)
	CGContextMoveToPoint func(c CGContextRef, x, y CGFloat)
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
	purego.RegisterLibFunc(&CGColorSpaceCreateDeviceRGB, cgLib, "CGColorSpaceCreateDeviceRGB")
	purego.RegisterLibFunc(&CGColorSpaceRelease, cgLib, "CGColorSpaceRelease")
	purego.RegisterLibFunc(&CGContextSetRGBFillColor, cgLib, "CGContextSetRGBFillColor")
	purego.RegisterLibFunc(&CGContextSetRGBStrokeColor, cgLib, "CGContextSetRGBStrokeColor")
	purego.RegisterLibFunc(&CGContextSetLineWidth, cgLib, "CGContextSetLineWidth")
	purego.RegisterLibFunc(&CGContextFillRect, cgLib, "CGContextFillRect")
	purego.RegisterLibFunc(&CGContextStrokeRect, cgLib, "CGContextStrokeRect")
	purego.RegisterLibFunc(&CGContextBeginPath, cgLib, "CGContextBeginPath")
	purego.RegisterLibFunc(&CGContextClosePath, cgLib, "CGContextClosePath")
	purego.RegisterLibFunc(&CGContextAddArc, cgLib, "CGContextAddArc")
	purego.RegisterLibFunc(&CGContextAddRect, cgLib, "CGContextAddRect")
	purego.RegisterLibFunc(&CGContextDrawPath, cgLib, "CGContextDrawPath")
	purego.RegisterLibFunc(&CGContextMoveToPoint, cgLib, "CGContextMoveToPoint")
	purego.RegisterLibFunc(&CGContextAddLineToPoint, cgLib, "CGContextAddLineToPoint")
}

// CustomView is our drawing view
type CustomView struct {
	appkit.View
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

	// Create custom view with drawing implementation
	customView := createCustomDrawingView()
	window.SetContentView(customView)

	// Show window
	window.MakeKeyAndOrderFront(nil)
	app.ActivateIgnoringOtherApps(true)

	fmt.Println("✅ CoreGraphics drawing example running")
	fmt.Println("   Drawing shapes using CGContext")
	fmt.Println("   - Blue filled rectangle")
	fmt.Println("   - Red circle outline")
	fmt.Println("   - Green triangle")
	fmt.Println("   Press Cmd+Q to quit\n")

	// Run application
	app.Run()
}

func createCustomDrawingView() appkit.View {
	// Create NSView subclass that implements drawRect:
	var drawRectCallback func(self objc.ID, rect foundation.Rect)
	drawRectCallback = func(self objc.ID, rect foundation.Rect) {
		// Get current graphics context
		ctx := appkit.GraphicsContext_CurrentContext()
		if ctx.Ptr() == nil {
			return
		}

		// Get CGContext from NSGraphicsContext
		cgContextPtr := objc.ID(ctx.Ptr()).Send(objc.RegisterName("CGContext"))
		if cgContextPtr == 0 {
			return
		}
		cgContext := CGContextRef(unsafe.Pointer(uintptr(cgContextPtr)))

		// Save graphics state
		CGContextSaveGState(cgContext)

		// Draw white background
		CGContextSetRGBFillColor(cgContext, 1.0, 1.0, 1.0, 1.0)
		CGContextFillRect(cgContext, CGRect{
			Origin: CGPoint{X: 0, Y: 0},
			Size:   CGSize{Width: CGFloat(rect.Size.Width), Height: CGFloat(rect.Size.Height)},
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
		// AddArc: center (x,y), radius, startAngle, endAngle, clockwise
		CGContextAddArc(cgContext, 400, 500, 60, 0, 6.28318, 0) // Full circle (2π radians)
		CGContextDrawPath(cgContext, kCGPathStroke)

		// Draw green filled triangle
		CGContextSetRGBFillColor(cgContext, 0.2, 0.8, 0.3, 1.0)
		CGContextBeginPath(cgContext)
		CGContextMoveToPoint(cgContext, 600, 450)
		CGContextAddLineToPoint(cgContext, 700, 450)
		CGContextAddLineToPoint(cgContext, 650, 550)
		CGContextClosePath(cgContext)
		CGContextDrawPath(cgContext, kCGPathFill)

		// Draw purple stroked rounded rectangle
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

		// Draw multiple colored rectangles
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

		// Draw text label (simplified - just drawing background for text area)
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

	// Register NSView subclass
	className := "CustomDrawingView"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSView")
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("drawRect:"),
					Fn:  drawRectCallback,
				},
			},
		)
	}

	// Create instance
	view := objc.ID(class).Send(objc.RegisterName("alloc"))
	view = view.Send(objc.RegisterName("init"))

	// Wrap in darwinkit View
	return appkit.ViewFrom(unsafe.Pointer(view))
}
