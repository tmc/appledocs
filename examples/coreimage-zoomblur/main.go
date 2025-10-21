package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
)

func main() {
	// AppKit requires running on the main thread
	runtime.LockOSThread()

	// Create the application
	appPtr := appkit.ApplicationClass.SharedApplication()
	app := appkit.ApplicationFrom(appPtr)

	// Set activation policy using objc.Send directly
	objc.Send[objc.ID](app.ID, objc.RegisterName("setActivationPolicy:"), uint(0))

	// Create the window
	window := createWindow()
	objc.Send[objc.ID](window.ID, objc.RegisterName("makeKeyAndOrderFront:"), objc.ID(0))

	// Activate the app
	objc.Send[objc.ID](app.ID, objc.RegisterName("activateIgnoringOtherApps:"), true)

	// Run the application event loop
	objc.Send[objc.ID](app.ID, objc.RegisterName("run"))
}

func createWindow() appkit.Window {
	// Window dimensions
	contentRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 100, Y: 100},
		Size:   coregraphics.CGSize{Width: 1200, Height: 700},
	}

	// Create window with standard style
	styleMask := appkit.WindowStyleMaskTitled |
		appkit.WindowStyleMaskClosable |
		appkit.WindowStyleMaskMiniaturizable |
		appkit.WindowStyleMaskResizable

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		contentRect,
		styleMask,
		appkit.BackingStoreBuffered,
		false,
	)

	objc.Send[objc.ID](window.ID, objc.RegisterName("setTitle:"), objc.String("CoreImage Zoom Blur Demo"))
	objc.Send[objc.ID](window.ID, objc.RegisterName("center"))

	// Create the content view
	contentView := createContentView()
	objc.Send[objc.ID](window.ID, objc.RegisterName("setContentView:"), contentView.ID)

	return window
}

func createContentView() appkit.View {
	// Create main container view
	containerView := appkit.NewView()

	// Create a sample image using CoreImage
	originalImage := createSampleImage()

	// Create image views for original and blurred versions
	imageSize := coregraphics.CGSize{Width: 500, Height: 400}

	// Original image view (left side)
	originalImageView := createImageView(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 50, Y: 200}, Size: imageSize},
		originalImage,
	)

	// Blurred image view (right side)
	blurredImage := applyZoomBlur(originalImage, 30.0, 250, 200)
	blurredImageView := createImageView(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 650, Y: 200}, Size: imageSize},
		blurredImage,
	)

	// Add labels
	originalLabel := createLabel(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 50, Y: 620}, Size: coregraphics.CGSize{Width: 500, Height: 30}},
		"Original Image",
	)

	blurredLabel := createLabel(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 650, Y: 620}, Size: coregraphics.CGSize{Width: 500, Height: 30}},
		"Zoom Blur Applied (amount: 30.0)",
	)

	// Add description
	description := createLabel(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 50, Y: 150}, Size: coregraphics.CGSize{Width: 1100, Height: 30}},
		"CIZoomBlur creates a radial blur effect emanating from the center point",
	)

	// Add controls info
	infoText := createLabel(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 50, Y: 100}, Size: coregraphics.CGSize{Width: 1100, Height: 40}},
		"The blur simulates a zooming motion blur effect. Higher amounts create more pronounced radial blur.",
	)

	// Add subviews
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), originalImageView.ID)
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), blurredImageView.ID)
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), originalLabel.ID)
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), blurredLabel.ID)
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), description.ID)
	objc.Send[objc.ID](containerView.ID, objc.RegisterName("addSubview:"), infoText.ID)

	return containerView
}

func createImageView(frame coregraphics.CGRect, ciImage unsafe.Pointer) appkit.ImageView {
	// Create NSImageView
	imageView := appkit.NewImageView()
	objc.Send[objc.ID](imageView.ID, objc.RegisterName("setFrame:"), frame)

	// Convert CIImage to NSImage
	nsImage := ciImageToNSImage(ciImage)
	if nsImage.ID != 0 {
		objc.Send[objc.ID](imageView.ID, objc.RegisterName("setImage:"), nsImage.ID)
	}

	// Configure image view - NSImageScaleProportionallyUpOrDown = 1
	objc.Send[objc.ID](imageView.ID, objc.RegisterName("setImageScaling:"), uint(1))
	objc.Send[objc.ID](imageView.ID, objc.RegisterName("setWantsLayer:"), true)

	// Add border
	layer := objc.Send[objc.ID](imageView.ID, objc.RegisterName("layer"))
	if layer != 0 {
		objc.Send[objc.ID](layer, objc.RegisterName("setBorderWidth:"), 2.0)

		// Set border color (light gray)
		colorSpace := coregraphics.CGColorSpaceCreateDeviceRGB()
		components := []float64{0.8, 0.8, 0.8, 1.0}
		cgColor := coregraphics.CGColorCreate(colorSpace, unsafe.Pointer(&components[0]))
		objc.Send[objc.ID](layer, objc.RegisterName("setBorderColor:"), cgColor)
	}

	return imageView
}

func createLabel(frame coregraphics.CGRect, text string) appkit.TextField {
	label := appkit.NewTextField()
	objc.Send[objc.ID](label.ID, objc.RegisterName("setFrame:"), frame)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setStringValue:"), objc.String(text))
	objc.Send[objc.ID](label.ID, objc.RegisterName("setBezeled:"), false)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setDrawsBackground:"), false)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setEditable:"), false)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setSelectable:"), false)

	// Center align (NSTextAlignmentCenter = 2)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setAlignment:"), uint(2))

	// Make font slightly larger
	fontSize := 14.0
	fontClass := objc.GetClass("NSFont")
	font := objc.Send[objc.ID](objc.ID(fontClass), objc.RegisterName("systemFontOfSize:"), fontSize)
	objc.Send[objc.ID](label.ID, objc.RegisterName("setFont:"), font)

	return label
}

func createSampleImage() unsafe.Pointer {
	// Create a colorful gradient image using CoreImage
	width := 500.0
	height := 400.0

	// Create a solid color image as the base
	colorClass := objc.GetClass("CIColor")
	color := objc.Send[objc.ID](objc.ID(colorClass), objc.RegisterName("colorWithRed:green:blue:"),
		0.5, 0.6, 0.8)

	colorImage := coreimage.NewImageWithColor(unsafe.Pointer(color))
	if colorImage.ID == 0 {
		return nil
	}

	// Crop to size
	cropRect := coregraphics.CGRect{
		Origin: coregraphics.CGPoint{X: 0, Y: 0},
		Size:   coregraphics.CGSize{Width: width, Height: height},
	}
	return colorImage.ImageByCroppingToRect(cropRect)
}

func applyZoomBlur(inputImage unsafe.Pointer, amount float64, centerX, centerY float64) unsafe.Pointer {
	// Create the CIZoomBlur filter
	filterPtr := coreimage.FilterClass.ZoomBlurFilter()
	if filterPtr == nil {
		fmt.Println("Warning: Failed to create zoom blur filter")
		return inputImage
	}

	filter := coreimage.FilterFrom(filterPtr)

	// Set the input image
	filter.SetValueForKey(objc.ID(inputImage), "inputImage")

	// Set the blur amount
	amountValue := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
	filter.SetValueForKey(amountValue.ID, "inputAmount")

	// Set the center point
	vectorClass := objc.GetClass("CIVector")
	sel := objc.RegisterName("vectorWithX:Y:")
	centerVector := objc.Send[objc.ID](objc.ID(vectorClass), sel, centerX, centerY)
	filter.SetValueForKey(centerVector, "inputCenter")

	// Get the output image
	outputImage := filter.OutputImage()
	if outputImage == nil {
		fmt.Println("Warning: Filter produced no output")
		return inputImage
	}

	return outputImage
}

func ciImageToNSImage(ciImage unsafe.Pointer) appkit.Image {
	if ciImage == nil {
		return appkit.Image{}
	}

	// Create a CIContext for rendering
	context := coreimage.NewContext()
	if context.ID == 0 {
		fmt.Println("Warning: Failed to create CIContext")
		return appkit.Image{}
	}

	// Get the image extent
	ciImageObj := coreimage.ImageFrom(ciImage)
	extent := ciImageObj.Extent()

	// Render to CGImage
	cgImage := context.CreateCGImageFromRect(ciImage, extent)
	if cgImage == 0 {
		fmt.Println("Warning: Failed to create CGImage")
		return appkit.Image{}
	}

	// Create NSImage from CGImage
	imageSize := coregraphics.CGSize{Width: extent.Size.Width, Height: extent.Size.Height}

	// Use NSImage initWithCGImage:size:
	nsImageClass := objc.GetClass("NSImage")
	nsImageAlloc := objc.Send[objc.ID](objc.ID(nsImageClass), objc.RegisterName("alloc"))
	sel := objc.RegisterName("initWithCGImage:size:")
	nsImageID := objc.Send[objc.ID](nsImageAlloc, sel, cgImage, imageSize)

	return appkit.ImageFrom(unsafe.Pointer(nsImageID))
}

// ===================================================================
// DEMONSTRATION OF CURRENT vs IDEAL API
// ===================================================================
//
// CURRENT API (What we use now):
// - Lots of unsafe.Pointer conversions
// - Manual objc.Send calls for many operations
// - String-based KVC for filter parameters
// - Manual NSNumber wrapping for primitives
// - No type safety or IDE autocomplete
//
// IDEAL API (With protocol wrappers):
//
// func applyZoomBlurIdeal(inputImage coreimage.Image, amount float64, centerX, centerY float64) coreimage.Image {
//     filter := coreimage.NewZoomBlurFilter()
//     filter.SetInputImage(inputImage)
//     filter.SetAmount(amount)
//     filter.SetCenter(coreimage.NewVector(centerX, centerY))
//     return filter.OutputImage()
// }
//
// Benefits:
// - Type-safe (proper types instead of unsafe.Pointer)
// - Concise (no manual wrapping)
// - Discoverable (IDE autocomplete shows all methods)
// - Less error-prone (compile-time checking)
//
