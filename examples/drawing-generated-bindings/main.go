// CoreGraphics Drawing using only generated bindings
//
// This advanced example demonstrates:
// - Drawing shapes using CoreGraphics
// - Creating NSImage and drawing into it
// - Using only generated bindings (no darwinkit)
// - Combining AppKit, Foundation, and CoreGraphics
package main

import (
	"fmt"
	"strings"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

// Foundation/AppKit types
type (
	NSPoint struct{ X, Y float64 }
	NSSize  struct{ Width, Height float64 }
	NSRect  struct {
		Origin NSPoint
		Size   NSSize
	}
)

// CoreGraphics types
type (
	CGFloat           float64
	CGContextRef      unsafe.Pointer
	CGPoint           struct{ X, Y CGFloat }
	CGSize            struct{ Width, Height CGFloat }
	CGRect            struct{ Origin CGPoint; Size CGSize }
	CGPathDrawingMode int32
)

const (
	kCGPathFill       CGPathDrawingMode = 0
	kCGPathStroke     CGPathDrawingMode = 2
	kCGPathFillStroke CGPathDrawingMode = 3
)

// CoreGraphics function bindings
var (
	cgLib                      uintptr
	CGContextSaveGState        func(c CGContextRef)
	CGContextRestoreGState     func(c CGContextRef)
	CGContextSetRGBFillColor   func(c CGContextRef, r, g, b, a CGFloat)
	CGContextSetRGBStrokeColor func(c CGContextRef, r, g, b, a CGFloat)
	CGContextSetLineWidth      func(c CGContextRef, width CGFloat)
	CGContextFillRect          func(c CGContextRef, rect CGRect)
	CGContextStrokeRect        func(c CGContextRef, rect CGRect)
	CGContextBeginPath         func(c CGContextRef)
	CGContextClosePath         func(c CGContextRef)
	CGContextAddArc            func(c CGContextRef, x, y, radius, startAngle, endAngle CGFloat, clockwise int32)
	CGContextDrawPath          func(c CGContextRef, mode CGPathDrawingMode)
	CGContextMoveToPoint       func(c CGContextRef, x, y CGFloat)
	CGContextAddLineToPoint    func(c CGContextRef, x, y CGFloat)
)

func init() {
	runtime.LockOSThread()

	// Load frameworks
	var err error
	_, err = purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	cgLib, err = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// Register CoreGraphics functions
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

func createAppDelegate() objc.ID {
	className := "DrawingAppDelegate"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		windowShouldClose := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
			appClass := objc.GetClass("NSApplication")
			app := objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))
			app.Send(objc.RegisterName("terminate:"), 0)
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
	image = image.Send(objc.RegisterName("initWithSize:"), NSSize{Width: width, Height: height})

	// Lock focus to draw
	image.Send(objc.RegisterName("lockFocus"))

	// Get current graphics context
	ctxClass := objc.GetClass("NSGraphicsContext")
	nsCtx := objc.ID(ctxClass).Send(objc.RegisterName("currentContext"))
	cgCtxID := nsCtx.Send(objc.RegisterName("CGContext"))
	cgCtx := CGContextRef(unsafe.Pointer(uintptr(cgCtxID)))

	// Draw shapes
	drawShapes(cgCtx, CGSize{Width: CGFloat(width), Height: CGFloat(height)})

	// Unlock focus
	image.Send(objc.RegisterName("unlockFocus"))

	return image
}

func drawShapes(ctx CGContextRef, size CGSize) {
	CGContextSaveGState(ctx)

	// White background
	CGContextSetRGBFillColor(ctx, 1.0, 1.0, 1.0, 1.0)
	CGContextFillRect(ctx, CGRect{Origin: CGPoint{X: 0, Y: 0}, Size: size})

	// Blue rectangle
	CGContextSetRGBFillColor(ctx, 0.2, 0.4, 0.8, 1.0)
	CGContextFillRect(ctx, CGRect{Origin: CGPoint{X: 50, Y: 450}, Size: CGSize{Width: 200, Height: 100}})

	// Red circle
	CGContextSetRGBStrokeColor(ctx, 0.9, 0.2, 0.2, 1.0)
	CGContextSetLineWidth(ctx, 3.0)
	CGContextBeginPath(ctx)
	CGContextAddArc(ctx, 400, 500, 60, 0, 6.28318, 0)
	CGContextDrawPath(ctx, kCGPathStroke)

	// Green triangle
	CGContextSetRGBFillColor(ctx, 0.2, 0.8, 0.3, 1.0)
	CGContextBeginPath(ctx)
	CGContextMoveToPoint(ctx, 600, 450)
	CGContextAddLineToPoint(ctx, 700, 450)
	CGContextAddLineToPoint(ctx, 650, 550)
	CGContextClosePath(ctx)
	CGContextDrawPath(ctx, kCGPathFill)

	// Purple rectangle outline
	CGContextSetRGBStrokeColor(ctx, 0.7, 0.2, 0.8, 1.0)
	CGContextSetLineWidth(ctx, 4.0)
	CGContextStrokeRect(ctx, CGRect{Origin: CGPoint{X: 50, Y: 300}, Size: CGSize{Width: 300, Height: 100}})

	// Orange circle
	CGContextSetRGBFillColor(ctx, 1.0, 0.6, 0.0, 1.0)
	CGContextBeginPath(ctx)
	CGContextAddArc(ctx, 500, 350, 40, 0, 6.28318, 0)
	CGContextDrawPath(ctx, kCGPathFill)

	// Cyan path with fill and stroke
	CGContextSetRGBFillColor(ctx, 0.3, 0.8, 0.9, 0.5)
	CGContextSetRGBStrokeColor(ctx, 0.0, 0.5, 0.6, 1.0)
	CGContextSetLineWidth(ctx, 2.0)
	CGContextBeginPath(ctx)
	CGContextMoveToPoint(ctx, 100, 200)
	CGContextAddLineToPoint(ctx, 200, 250)
	CGContextAddLineToPoint(ctx, 150, 150)
	CGContextAddLineToPoint(ctx, 250, 180)
	CGContextClosePath(ctx)
	CGContextDrawPath(ctx, kCGPathFillStroke)

	// Rainbow bars
	colors := []struct{ r, g, b CGFloat }{
		{1.0, 0.0, 0.0}, {1.0, 0.5, 0.0}, {1.0, 1.0, 0.0},
		{0.0, 1.0, 0.0}, {0.0, 0.0, 1.0}, {0.5, 0.0, 0.5},
	}
	for i, c := range colors {
		CGContextSetRGBFillColor(ctx, c.r, c.g, c.b, 0.8)
		CGContextFillRect(ctx, CGRect{
			Origin: CGPoint{X: CGFloat(350 + i*50), Y: 150},
			Size:   CGSize{Width: 40, Height: 150},
		})
	}

	CGContextRestoreGState(ctx)
}

func main() {
	fmt.Println("=== CoreGraphics Drawing (Generated Bindings) ===\n")

	// Create app
	appClass := objc.GetClass("NSApplication")
	app := objc.ID(appClass).Send(objc.RegisterName("sharedApplication"))
	app.Send(objc.RegisterName("setActivationPolicy:"), 0)

	// Create window
	windowClass := objc.GetClass("NSWindow")
	window := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
	window = window.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		NSRect{Origin: NSPoint{X: 100, Y: 100}, Size: NSSize{Width: 800, Height: 600}},
		1|2|8, 2, false)

	// Set title
	titleStr := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"), "CoreGraphics Drawing (Generated Bindings)")
	window.Send(objc.RegisterName("setTitle:"), titleStr)

	// Set delegate
	delegate := createAppDelegate()
	window.Send(objc.RegisterName("setDelegate:"), delegate)

	// Create and draw image
	image := createDrawnImage(800, 600)

	// Create image view
	imageViewClass := objc.GetClass("NSImageView")
	imageView := objc.ID(imageViewClass).Send(objc.RegisterName("alloc"))
	imageView = imageView.Send(objc.RegisterName("initWithFrame:"),
		NSRect{Origin: NSPoint{X: 0, Y: 0}, Size: NSSize{Width: 800, Height: 600}})
	imageView.Send(objc.RegisterName("setImage:"), image)
	imageView.Send(objc.RegisterName("setImageScaling:"), 1) // NSImageScaleProportionallyUpOrDown

	// Set as content view
	window.Send(objc.RegisterName("setContentView:"), imageView)

	// Show window
	window.Send(objc.RegisterName("makeKeyAndOrderFront:"), 0)
	app.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)

	fmt.Println("✅ Drawing example using only generated bindings")
	fmt.Println("   - CoreGraphics for drawing")
	fmt.Println("   - AppKit for windowing")
	fmt.Println("   - Foundation types")
	fmt.Println("   - No darwinkit dependency!")
	fmt.Println("   Press Cmd+Q or close window to quit\n")

	app.Send(objc.RegisterName("run"))
}
