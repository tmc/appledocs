// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Filter] class.
var (
	filterClass     _FilterClass
	filterClassOnce sync.Once
)

func getFilterClass() _FilterClass {
	filterClassOnce.Do(func() {
		filterClass = _FilterClass{objc.GetClass("CIFilter")}
	})
	return filterClass
}

type _FilterClass struct {
	class objc.Class
}

// An interface definition for the [Filter] class.
type IFilter interface {
	objectivec.IObject
	ApplyArgumentsOptions(k unsafe.Pointer, args unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer
	Apply(k unsafe.Pointer) unsafe.Pointer
	SetDefaults()
	ViewForUIConfigurationExcludedKeys(inUIConfiguration unsafe.Pointer, inKeys unsafe.Pointer) unsafe.Pointer
}

// An image processor that produces an image by manipulating one or more input images or by generating new image data.
//
// The class produces a object as output. Typically, a filter takes one or more images as input. Some filters, however, generate an image based on other types of input parameters. The par swift.class` object are set and retrieved through the use of key-value pairs. You use the object in conjunction with other Core Image classes, such as , , and , to take advantage of the built-in Core Image filters when processing images, creating filter generators, or writing custom filters. objects are mutable, and thus cannot be shared safely among threads. Each thread must create its own objects, but you can pass a filter’s immutable input and output objects between threads. To get a quick overview of how to set up and use Core Image filters, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
type Filter struct {
	objectivec.Object
}

// FilterFrom constructs a [Filter] from an unsafe.Pointer.
//
// An image processor that produces an image by manipulating one or more input images or by generating new image data.
func FilterFrom(ptr unsafe.Pointer) Filter {
	return Filter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterClass) Alloc() Filter {
	rv := objc.Send[Filter](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterClass) New() Filter {
	rv := objc.Send[Filter](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Filter) Init() Filter {
	rv := objc.Send[Filter](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Filter) Autorelease() Filter {
	rv := objc.Send[Filter](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilter creates a new Filter instance.
func NewFilter() Filter {
	return getFilterClass().New()
}


// Creates a filter that allows the processing of RAW images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(imageURL:options:)
func NewFilterWithImageURLOptions(url unsafe.Pointer, options unsafe.Pointer) Filter {
	rv := objc.Send[Filter](objc.ID(getFilterClass().class), objc.Sel("filterWithImageURL:options:"), url, options)
	return rv
}

// Creates a object for a specific kind of filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:)
func NewFilterWithName(name string) Filter {
	rv := objc.Send[Filter](objc.ID(getFilterClass().class), objc.Sel("filterWithName:"), objc.String(name))
	return rv
}

// Creates a object for a specific kind of filter and initializes the input values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:withInputParameters:)
func NewFilterWithNameWithInputParameters(name string, params unsafe.Pointer) Filter {
	rv := objc.Send[Filter](objc.ID(getFilterClass().class), objc.Sel("filterWithName:withInputParameters:"), objc.String(name), params)
	return rv
}

// Creates a filter from a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(CVPixelBuffer:properties:options:)
func NewFilterWithCVPixelBufferPropertiesOptions(pixelBuffer unsafe.Pointer, properties unsafe.Pointer, options unsafe.Pointer) Filter {
	rv := objc.Send[Filter](objc.ID(getFilterClass().class), objc.Sel("filterWithCVPixelBuffer:properties:options:"), pixelBuffer, properties, options)
	return rv
}

// Creates a filter that allows the processing of RAW images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(imageData:options:)
func NewFilterWithImageDataOptions(data unsafe.Pointer, options unsafe.Pointer) Filter {
	rv := objc.Send[Filter](objc.ID(getFilterClass().class), objc.Sel("filterWithImageData:options:"), data, options)
	return rv
}


// Transitions by folding and crossfading an image to reveal the target image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/accordionFoldTransition()
func (fc _FilterClass) AccordionFoldTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("accordionFoldTransitionFilter"))
	return rv
}

// Blends colors from two images by addition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/additionCompositing()
func (fc _FilterClass) AdditionCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("additionCompositingFilter"))
	return rv
}

// Performs a transform on the image and extends the image edges to infinity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/affineClamp()
func (fc _FilterClass) AffineClampFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("affineClampFilter"))
	return rv
}

// Performs a transform on the image and tiles the result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/affineTile()
func (fc _FilterClass) AffineTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("affineTileFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAlphaWeightedHistogram()
func (fc _FilterClass) AreaAlphaWeightedHistogramFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaAlphaWeightedHistogramFilter"))
	return rv
}

// Returns a 1 x 1 pixel image that contains the average color for the region of interest.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAverage()
func (fc _FilterClass) AreaAverageFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaAverageFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAverageMaximumRed()
func (fc _FilterClass) AreaAverageMaximumRedFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaAverageMaximumRedFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaBoundsRed()
func (fc _FilterClass) AreaBoundsRedFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaBoundsRedFilter"))
	return rv
}

// Returns a histogram of a specified area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaHistogram()
func (fc _FilterClass) AreaHistogramFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaHistogramFilter"))
	return rv
}

// Returns a logarithmic histogram of a specified area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaLogarithmicHistogram()
func (fc _FilterClass) AreaLogarithmicHistogramFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaLogarithmicHistogramFilter"))
	return rv
}

// Calculates the maximum color components of a specified area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMaximum()
func (fc _FilterClass) AreaMaximumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMaximumFilter"))
	return rv
}

// Finds the pixel with the highest alpha value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMaximumAlpha()
func (fc _FilterClass) AreaMaximumAlphaFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMaximumAlphaFilter"))
	return rv
}

// Calculates minimum and maximum color components for a specified area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinMax()
func (fc _FilterClass) AreaMinMaxFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMinMaxFilter"))
	return rv
}

// Calculates the minimum and maximum red component value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinMaxRed()
func (fc _FilterClass) AreaMinMaxRedFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMinMaxRedFilter"))
	return rv
}

// Calculates the minimum color component values for a specified area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinimum()
func (fc _FilterClass) AreaMinimumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMinimumFilter"))
	return rv
}

// Calculates the pixel within a specified area that has the smallest alpha value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinimumAlpha()
func (fc _FilterClass) AreaMinimumAlphaFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("areaMinimumAlphaFilter"))
	return rv
}

// Generates an attributed-text image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/attributedTextImageGenerator()
func (fc _FilterClass) AttributedTextImageGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("attributedTextImageGeneratorFilter"))
	return rv
}

// Generates a low-density barcode.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/aztecCodeGenerator()
func (fc _FilterClass) AztecCodeGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("aztecCodeGeneratorFilter"))
	return rv
}

// Generates a barcode as an image from the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/barcodeGenerator()
func (fc _FilterClass) BarcodeGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("barcodeGeneratorFilter"))
	return rv
}

// Transitions between two images by removing rectangular portions of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/barsSwipeTransition()
func (fc _FilterClass) BarsSwipeTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("barsSwipeTransitionFilter"))
	return rv
}

// Produces a high-quality scaled version of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bicubicScaleTransform()
func (fc _FilterClass) BicubicScaleTransformFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("bicubicScaleTransformFilter"))
	return rv
}

// Blends two images by using an alpha mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithAlphaMask()
func (fc _FilterClass) BlendWithAlphaMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blendWithAlphaMaskFilter"))
	return rv
}

// Blends two images by using a blue mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithBlueMask()
func (fc _FilterClass) BlendWithBlueMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blendWithBlueMaskFilter"))
	return rv
}

// Blends two images by using a mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithMask()
func (fc _FilterClass) BlendWithMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blendWithMaskFilter"))
	return rv
}

// Blends two images by using a red mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithRedMask()
func (fc _FilterClass) BlendWithRedMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blendWithRedMaskFilter"))
	return rv
}

// Adjusts an image’s colors by applying a blur effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bloom()
func (fc _FilterClass) BloomFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("bloomFilter"))
	return rv
}

// Generates a blurred rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blurredRectangleGenerator()
func (fc _FilterClass) BlurredRectangleGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blurredRectangleGeneratorFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blurredRoundedRectangleGenerator()
func (fc _FilterClass) BlurredRoundedRectangleGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blurredRoundedRectangleGeneratorFilter"))
	return rv
}

// Applies a bokeh effect to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bokehBlur()
func (fc _FilterClass) BokehBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("bokehBlurFilter"))
	return rv
}

// Applies a square-shaped blur to an area of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/boxBlur()
func (fc _FilterClass) BoxBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("boxBlurFilter"))
	return rv
}

// Distorts an image with a concave or convex bump.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bumpDistortion()
func (fc _FilterClass) BumpDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("bumpDistortionFilter"))
	return rv
}

// Linearly distorts an image with a concave or convex bump.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bumpDistortionLinear()
func (fc _FilterClass) BumpDistortionLinearFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("bumpDistortionLinearFilter"))
	return rv
}

// Applies the Canny edge-detection algorithm to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/cannyEdgeDetector()
func (fc _FilterClass) CannyEdgeDetectorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("cannyEdgeDetectorFilter"))
	return rv
}

// Generates a checkerboard image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/checkerboardGenerator()
func (fc _FilterClass) CheckerboardGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("checkerboardGeneratorFilter"))
	return rv
}

// Distorts an image with radiating circles to the periphery of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circleSplashDistortion()
func (fc _FilterClass) CircleSplashDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("circleSplashDistortionFilter"))
	return rv
}

// Adds a circular overlay to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circularScreen()
func (fc _FilterClass) CircularScreenFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("circularScreenFilter"))
	return rv
}

// Distorts an image by increasing the distance of the center of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circularWrap()
func (fc _FilterClass) CircularWrapFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("circularWrapFilter"))
	return rv
}

// Adds a series of colorful dots to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/cmykHalftone()
func (fc _FilterClass) CMYKHalftone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("CMYKHalftone"))
	return rv
}

// Generates a high-density, linear barcode.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/code128BarcodeGenerator()
func (fc _FilterClass) Code128BarcodeGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("code128BarcodeGeneratorFilter"))
	return rv
}

// Calculates the absolute difference between each color component in the input images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorAbsoluteDifference()
func (fc _FilterClass) ColorAbsoluteDifferenceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorAbsoluteDifferenceFilter"))
	return rv
}

// Blends color from two images using the luminance values from the background image and the hue and saturation values from the input image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorBlendMode()
func (fc _FilterClass) ColorBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorBlendModeFilter"))
	return rv
}

// Blends color from two images while darkening the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorBurnBlendMode()
func (fc _FilterClass) ColorBurnBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorBurnBlendModeFilter"))
	return rv
}

// Alters the colors in an image based on color components.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorClamp()
func (fc _FilterClass) ColorClampFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorClampFilter"))
	return rv
}

// Alters the brightness, contrast, and saturation of an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorControls()
func (fc _FilterClass) ColorControlsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorControlsFilter"))
	return rv
}

// Adjusts an image’s color by applying polynomial cross-products.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCrossPolynomial()
func (fc _FilterClass) ColorCrossPolynomialFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorCrossPolynomialFilter"))
	return rv
}

// Adjusts an image’s pixels using a three-dimensional color table.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCube()
func (fc _FilterClass) ColorCubeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorCubeFilter"))
	return rv
}

// Adjusts an image’s pixels using a three-dimensional color table in specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCubeWithColorSpace()
func (fc _FilterClass) ColorCubeWithColorSpaceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorCubeWithColorSpaceFilter"))
	return rv
}

// Alters an image’s pixels using a three-dimensional color tables and a mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCubesMixedWithMask()
func (fc _FilterClass) ColorCubesMixedWithMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorCubesMixedWithMaskFilter"))
	return rv
}

// Adjusts an image’s color curves.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCurves()
func (fc _FilterClass) ColorCurvesFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorCurvesFilter"))
	return rv
}

// Blends color from two images using dodging.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorDodgeBlendMode()
func (fc _FilterClass) ColorDodgeBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorDodgeBlendModeFilter"))
	return rv
}

// Inverts an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorInvert()
func (fc _FilterClass) ColorInvertFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorInvertFilter"))
	return rv
}

// Performs a transformation of the input image colors to colors from a gradient image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMap()
func (fc _FilterClass) ColorMapFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorMapFilter"))
	return rv
}

// Alters the colors in an image based on vectors provided.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMatrix()
func (fc _FilterClass) ColorMatrixFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorMatrixFilter"))
	return rv
}

// Adjusts an image’s colors to shades of a single color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMonochrome()
func (fc _FilterClass) ColorMonochromeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorMonochromeFilter"))
	return rv
}

// Alters an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorPolynomial()
func (fc _FilterClass) ColorPolynomialFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorPolynomialFilter"))
	return rv
}

// Flattens an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorPosterize()
func (fc _FilterClass) ColorPosterizeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorPosterizeFilter"))
	return rv
}

// Compares the red, green, and blue components of the input image to a threshold and sets them to 1 or 0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorThreshold()
func (fc _FilterClass) ColorThresholdFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorThresholdFilter"))
	return rv
}

// Compares the red, green, and blue components of the input image against a threshold calculated using Otsu’s algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorThresholdOtsu()
func (fc _FilterClass) ColorThresholdOtsuFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("colorThresholdOtsuFilter"))
	return rv
}

// Calculates the average color for a specified column of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/columnAverage()
func (fc _FilterClass) ColumnAverageFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("columnAverageFilter"))
	return rv
}

// Creates an image with a comic book effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/comicEffect()
func (fc _FilterClass) ComicEffectFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("comicEffectFilter"))
	return rv
}

// Converts an image from CIELAB to RGB color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convertLabToRGB()
func (fc _FilterClass) ConvertLabToRGBFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convertLabToRGBFilter"))
	return rv
}

// Converts an image from RGB to CIELAB color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convertRGBtoLab()
func (fc _FilterClass) ConvertRGBtoLabFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convertRGBtoLabFilter"))
	return rv
}

// Applies a convolution 3 x 3 filter to the components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution3X3()
func (fc _FilterClass) Convolution3X3Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolution3X3Filter"))
	return rv
}

// Applies a convolution 5 x 5 filter to the components image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution5X5()
func (fc _FilterClass) Convolution5X5Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolution5X5Filter"))
	return rv
}

// Applies a convolution 7 x 7 filter to the color components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution7X7()
func (fc _FilterClass) Convolution7X7Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolution7X7Filter"))
	return rv
}

// Applies a convolution-9 horizontal filter to the components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution9Horizontal()
func (fc _FilterClass) Convolution9HorizontalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolution9HorizontalFilter"))
	return rv
}

// Applies a convolution-9 vertical filter to the components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution9Vertical()
func (fc _FilterClass) Convolution9VerticalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolution9VerticalFilter"))
	return rv
}

// Applies a convolution 3 x 3 filter to the components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB3X3()
func (fc _FilterClass) ConvolutionRGB3X3Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolutionRGB3X3Filter"))
	return rv
}

// Applies a convolution 5 x 5 filter to the components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB5X5()
func (fc _FilterClass) ConvolutionRGB5X5Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolutionRGB5X5Filter"))
	return rv
}

// Applies a convolution 7 x 7 filter to the RGB components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB7X7()
func (fc _FilterClass) ConvolutionRGB7X7Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolutionRGB7X7Filter"))
	return rv
}

// Applies a convolution 9 x 1 filter to the RGB components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB9Horizontal()
func (fc _FilterClass) ConvolutionRGB9HorizontalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolutionRGB9HorizontalFilter"))
	return rv
}

// Applies a convolution 1 x 9 filter to the RGB components of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB9Vertical()
func (fc _FilterClass) ConvolutionRGB9VerticalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("convolutionRGB9VerticalFilter"))
	return rv
}

// Simulates the effect of a copy machine scanner light to transiton between two images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/copyMachineTransition()
func (fc _FilterClass) CopyMachineTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("copyMachineTransitionFilter"))
	return rv
}

// Filters an image with a Core ML model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/coreMLModel()
func (fc _FilterClass) CoreMLModelFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("coreMLModelFilter"))
	return rv
}

// Creates an image made with a series of colorful polygons.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/crystallize()
func (fc _FilterClass) CrystallizeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("crystallizeFilter"))
	return rv
}

// Blends colors from two images while darkening lighter pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/darkenBlendMode()
func (fc _FilterClass) DarkenBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("darkenBlendModeFilter"))
	return rv
}

// Simulates a depth of field effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/depthOfField()
func (fc _FilterClass) DepthOfFieldFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("depthOfFieldFilter"))
	return rv
}

// Converts from an image containing depth data to an image containing disparity data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/depthToDisparity()
func (fc _FilterClass) DepthToDisparityFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("depthToDisparityFilter"))
	return rv
}

// Subtracts color values to blend colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/differenceBlendMode()
func (fc _FilterClass) DifferenceBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("differenceBlendModeFilter"))
	return rv
}

// Applies a circle-shaped blur to an area of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/discBlur()
func (fc _FilterClass) DiscBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("discBlurFilter"))
	return rv
}

// Transitions between two images using a mask image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/disintegrateWithMaskTransition()
func (fc _FilterClass) DisintegrateWithMaskTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("disintegrateWithMaskTransitionFilter"))
	return rv
}

// Creates depth data from an image containing disparity data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/disparityToDepth()
func (fc _FilterClass) DisparityToDepthFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("disparityToDepthFilter"))
	return rv
}

// Applies the grayscale values of the second image to the first image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/displacementDistortion()
func (fc _FilterClass) DisplacementDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("displacementDistortionFilter"))
	return rv
}

// Transitions between two images with a fade effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dissolveTransition()
func (fc _FilterClass) DissolveTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("dissolveTransitionFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/distanceGradientFromRedMask()
func (fc _FilterClass) DistanceGradientFromRedMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("distanceGradientFromRedMaskFilter"))
	return rv
}

// Applies randomized noise to produce a processed look.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dither()
func (fc _FilterClass) DitherFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("ditherFilter"))
	return rv
}

// Divides color values to blend colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/divideBlendMode()
func (fc _FilterClass) DivideBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("divideBlendModeFilter"))
	return rv
}

// Adjusts an image’s shadows and contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/documentEnhancer()
func (fc _FilterClass) DocumentEnhancerFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("documentEnhancerFilter"))
	return rv
}

// Creates a monochrome image with a series of dots to add detail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dotScreen()
func (fc _FilterClass) DotScreenFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("dotScreenFilter"))
	return rv
}

// Stylizes an image with the Droste effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/droste()
func (fc _FilterClass) DrosteFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("drosteFilter"))
	return rv
}

// Creates a high-quality upscaled image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edgePreserveUpsample()
func (fc _FilterClass) EdgePreserveUpsampleFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("edgePreserveUpsampleFilter"))
	return rv
}

// Produces a black-and-white image that looks similar to a woodblock print.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edgeWork()
func (fc _FilterClass) EdgeWorkFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("edgeWorkFilter"))
	return rv
}

// Hilghlights edges of objects found within an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edges()
func (fc _FilterClass) EdgesFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("edgesFilter"))
	return rv
}

// Creates an eight-way reflected pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/eightfoldReflectedTile()
func (fc _FilterClass) EightfoldReflectedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("eightfoldReflectedTileFilter"))
	return rv
}

// Subtracts color values to blend colors with less contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/exclusionBlendMode()
func (fc _FilterClass) ExclusionBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("exclusionBlendModeFilter"))
	return rv
}

// Adjusts an image’s exposure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/exposureAdjust()
func (fc _FilterClass) ExposureAdjustFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("exposureAdjustFilter"))
	return rv
}

// Replaces an image’s colors with specified colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/falseColor()
func (fc _FilterClass) FalseColorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("falseColorFilter"))
	return rv
}

// Returns an array of filter objects de-serialized from XMP data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterArray(fromSerializedXMP:inputImageExtent:error:)
func (fc _FilterClass) FilterArrayFromSerializedXMPInputImageExtentError(xmpData unsafe.Pointer, extent unsafe.Pointer, outError unsafe.Pointer) []Filter {
	rv := objc.Send[[]Filter](objc.ID(fc.class), objc.Sel("filterArrayFromSerializedXMP:inputImageExtent:error:"), xmpData, extent, outError)
	return rv
}

// Returns an array of all published filter names that match all the specified categories.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterNames(inCategories:)
func (fc _FilterClass) FilterNamesInCategories(categories unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(fc.class), objc.Sel("filterNamesInCategories:"), categories)
	return rv
}

// Returns an array of all published filter names in the specified category.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterNames(inCategory:)
func (fc _FilterClass) FilterNamesInCategory(category string) []string {
	rv := objc.Send[[]string](objc.ID(fc.class), objc.Sel("filterNamesInCategory:"), objc.String(category))
	return rv
}

// Creates a object for a specific kind of filter and initializes the input values with a -terminated list of arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterWithName:keysAndValues:
func (fc _FilterClass) FilterWithNameKeysAndValues(name string, key0 objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithName:keysAndValues:"), objc.String(name), key0)
	return rv
}

// Creates a flash of light to transition between two images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/flashTransition()
func (fc _FilterClass) FlashTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("flashTransitionFilter"))
	return rv
}

// Creates a four-way reflected pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldReflectedTile()
func (fc _FilterClass) FourfoldReflectedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fourfoldReflectedTileFilter"))
	return rv
}

// Creates a tiled image by rotating a tile in increments of 90 degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldRotatedTile()
func (fc _FilterClass) FourfoldRotatedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fourfoldRotatedTileFilter"))
	return rv
}

// Creates a tiled image by applying four translation operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldTranslatedTile()
func (fc _FilterClass) FourfoldTranslatedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fourfoldTranslatedTileFilter"))
	return rv
}

// Highlights textures in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaborGradients()
func (fc _FilterClass) GaborGradientsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("gaborGradientsFilter"))
	return rv
}

// Alters an image’s transition between black and white.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gammaAdjust()
func (fc _FilterClass) GammaAdjustFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("gammaAdjustFilter"))
	return rv
}

// Blurs an image with a Gaussian distribution pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaussianBlur()
func (fc _FilterClass) GaussianBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("gaussianBlurFilter"))
	return rv
}

// Generates a gradient that varies from one color to another using a Gaussian distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaussianGradient()
func (fc _FilterClass) GaussianGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("gaussianGradientFilter"))
	return rv
}

// Distorts an image by applying a glass-like texture.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glassDistortion()
func (fc _FilterClass) GlassDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("glassDistortionFilter"))
	return rv
}

// Creates a lozenge-shaped lens and distorts the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glassLozenge()
func (fc _FilterClass) GlassLozengeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("glassLozengeFilter"))
	return rv
}

// Tiles an image by rotating and reflecting a tile from the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glideReflectedTile()
func (fc _FilterClass) GlideReflectedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("glideReflectedTileFilter"))
	return rv
}

// Adjusts an image’s color by applying a gloom filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gloom()
func (fc _FilterClass) GloomFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("gloomFilter"))
	return rv
}

// Blends colors of two images by screening and multiplying.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hardLightBlendMode()
func (fc _FilterClass) HardLightBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hardLightBlendModeFilter"))
	return rv
}

// Creates a monochrome image with a series of lines to add detail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hatchedScreen()
func (fc _FilterClass) HatchedScreenFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hatchedScreenFilter"))
	return rv
}

// Creates a realistic shaded height-field image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/heightFieldFromMask()
func (fc _FilterClass) HeightFieldFromMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("heightFieldFromMaskFilter"))
	return rv
}

// Creates an image made of a series of colorful hexagons.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hexagonalPixellate()
func (fc _FilterClass) HexagonalPixellateFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hexagonalPixellateFilter"))
	return rv
}

// Adjusts the highlights of colors to reduce shadows.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/highlightShadowAdjust()
func (fc _FilterClass) HighlightShadowAdjustFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("highlightShadowAdjustFilter"))
	return rv
}

// Generates a histogram map from the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/histogramDisplay()
func (fc _FilterClass) HistogramDisplayFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("histogramDisplayFilter"))
	return rv
}

// Distorts an image with a circular area that pushes the image outward.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/holeDistortion()
func (fc _FilterClass) HoleDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("holeDistortionFilter"))
	return rv
}

// Modifies an image’s hue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueAdjust()
func (fc _FilterClass) HueAdjustFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hueAdjustFilter"))
	return rv
}

// Blends colors of two images by computing the sum of image color values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueBlendMode()
func (fc _FilterClass) HueBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hueBlendModeFilter"))
	return rv
}

// Generates a gradient representing a specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueSaturationValueGradient()
func (fc _FilterClass) HueSaturationValueGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("hueSaturationValueGradientFilter"))
	return rv
}

// Creates a filter from a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(CVPixelBuffer:properties:options:)
func (fc _FilterClass) FilterWithCVPixelBufferPropertiesOptions(pixelBuffer unsafe.Pointer, properties unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithCVPixelBuffer:properties:options:"), pixelBuffer, properties, options)
	return rv
}

// Creates a filter that allows the processing of RAW images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(imageData:options:)
func (fc _FilterClass) FilterWithImageDataOptions(data unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithImageData:options:"), data, options)
	return rv
}

// Creates a filter that allows the processing of RAW images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(imageURL:options:)
func (fc _FilterClass) FilterWithImageURLOptions(url unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithImageURL:options:"), url, options)
	return rv
}

// Creates a object for a specific kind of filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:)
func (fc _FilterClass) FilterWithName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithName:"), objc.String(name))
	return rv
}

// Creates a object for a specific kind of filter and initializes the input values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:withInputParameters:)
func (fc _FilterClass) FilterWithNameWithInputParameters(name string, params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterWithName:withInputParameters:"), objc.String(name), params)
	return rv
}

// Applies the k-means algorithm to find the most common colors in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/kMeans()
func (fc _FilterClass) KMeansFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("KMeansFilter"))
	return rv
}

// Creates a 12-way kaleidoscopic image from an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/kaleidoscope()
func (fc _FilterClass) KaleidoscopeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("kaleidoscopeFilter"))
	return rv
}

// Adjusts the image vertically and horizontally to remove distortion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionCombined()
func (fc _FilterClass) KeystoneCorrectionCombinedFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("keystoneCorrectionCombinedFilter"))
	return rv
}

// Horizontally adjusts an image to remove distortion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionHorizontal()
func (fc _FilterClass) KeystoneCorrectionHorizontalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("keystoneCorrectionHorizontalFilter"))
	return rv
}

// Vertically adjusts an image to remove distortion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionVertical()
func (fc _FilterClass) KeystoneCorrectionVerticalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("keystoneCorrectionVerticalFilter"))
	return rv
}

// Compares an image’s color values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/labDeltaE()
func (fc _FilterClass) LabDeltaE() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("LabDeltaE"))
	return rv
}

// Creates a high-quality, scaled version of a source image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lanczosScaleTransform()
func (fc _FilterClass) LanczosScaleTransformFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lanczosScaleTransformFilter"))
	return rv
}

// Generates a lenticular halo image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lenticularHaloGenerator()
func (fc _FilterClass) LenticularHaloGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lenticularHaloGeneratorFilter"))
	return rv
}

// Distorts an image by generating a light tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lightTunnel()
func (fc _FilterClass) LightTunnelFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lightTunnelFilter"))
	return rv
}

// Blends colors from two images by brightening colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lightenBlendMode()
func (fc _FilterClass) LightenBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lightenBlendModeFilter"))
	return rv
}

// Creates an image that resembles a sketch of the outlines of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lineOverlay()
func (fc _FilterClass) LineOverlayFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lineOverlayFilter"))
	return rv
}

// Creates a monochrome image with a series of small lines to add detail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lineScreen()
func (fc _FilterClass) LineScreenFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("lineScreenFilter"))
	return rv
}

// Blends color from two images while increasing contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearBurnBlendMode()
func (fc _FilterClass) LinearBurnBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("linearBurnBlendModeFilter"))
	return rv
}

// Blends colors of two images with dodging.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearDodgeBlendMode()
func (fc _FilterClass) LinearDodgeBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("linearDodgeBlendModeFilter"))
	return rv
}

// Generates a color gradient that varies along a linear axis between two defined endpoints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearGradient()
func (fc _FilterClass) LinearGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("linearGradientFilter"))
	return rv
}

// A combination of linear burn and linear dodge blend modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearLightBlendMode()
func (fc _FilterClass) LinearLightBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("linearLightBlendModeFilter"))
	return rv
}

// Alters an image’s color intensity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearToSRGBToneCurve()
func (fc _FilterClass) LinearToSRGBToneCurveFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("linearToSRGBToneCurveFilter"))
	return rv
}

// Returns the localized description of a filter for display in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedDescription(forFilterName:)
func (fc _FilterClass) LocalizedDescriptionForFilterName(filterName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("localizedDescriptionForFilterName:"), objc.String(filterName))
	return rv
}

// Returns the localized name for the specified filter category.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedName(forCategory:)
func (fc _FilterClass) LocalizedNameForCategory(category string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("localizedNameForCategory:"), objc.String(category))
	return rv
}

// Returns the localized name for the specified filter name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedName(forFilterName:)
func (fc _FilterClass) LocalizedNameForFilterName(filterName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("localizedNameForFilterName:"), objc.String(filterName))
	return rv
}

// Returns the location of the localized reference documentation that describes the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedReferenceDocumentation(forFilterName:)
func (fc _FilterClass) LocalizedReferenceDocumentationForFilterName(filterName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("localizedReferenceDocumentationForFilterName:"), objc.String(filterName))
	return rv
}

// Blends color from two images by calculating the color, hue, and saturation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/luminosityBlendMode()
func (fc _FilterClass) LuminosityBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("luminosityBlendModeFilter"))
	return rv
}

// Converts an image to a white image with an alpha component.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maskToAlpha()
func (fc _FilterClass) MaskToAlphaFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("maskToAlphaFilter"))
	return rv
}

// Blurs a specified portion of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maskedVariableBlur()
func (fc _FilterClass) MaskedVariableBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("maskedVariableBlurFilter"))
	return rv
}

// Creates a maximum RGB grayscale image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumComponent()
func (fc _FilterClass) MaximumComponentFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("maximumComponentFilter"))
	return rv
}

// Applies a maximum compositing filter to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumCompositing()
func (fc _FilterClass) MaximumCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("maximumCompositingFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumScaleTransform()
func (fc _FilterClass) MaximumScaleTransformFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("maximumScaleTransformFilter"))
	return rv
}

// Calculates the median of an image to refine detail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/median()
func (fc _FilterClass) MedianFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("medianFilter"))
	return rv
}

// Generates a pattern made from an array of line segments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/meshGenerator()
func (fc _FilterClass) MeshGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("meshGeneratorFilter"))
	return rv
}

// Creates a minimum RGB grayscale image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/minimumComponent()
func (fc _FilterClass) MinimumComponentFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("minimumComponentFilter"))
	return rv
}

// Blends colors from two images by computing minimum values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/minimumCompositing()
func (fc _FilterClass) MinimumCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("minimumCompositingFilter"))
	return rv
}

// Blends two images together.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/mix()
func (fc _FilterClass) MixFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("mixFilter"))
	return rv
}

// Transitions between two images by applying irregularly shaped holes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/modTransition()
func (fc _FilterClass) ModTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("modTransitionFilter"))
	return rv
}

// Detects and highlights edges of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyGradient()
func (fc _FilterClass) MorphologyGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("morphologyGradientFilter"))
	return rv
}

// Blurs a circular area by enlarging contrasting pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyMaximum()
func (fc _FilterClass) MorphologyMaximumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("morphologyMaximumFilter"))
	return rv
}

// Blurs a circular area by reducing contrasting pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyMinimum()
func (fc _FilterClass) MorphologyMinimumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("morphologyMinimumFilter"))
	return rv
}

// Blurs a rectangular area by enlarging contrasting pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyRectangleMaximum()
func (fc _FilterClass) MorphologyRectangleMaximumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("morphologyRectangleMaximumFilter"))
	return rv
}

// Blurs a rectangular area by reducing contrasting pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyRectangleMinimum()
func (fc _FilterClass) MorphologyRectangleMinimumFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("morphologyRectangleMinimumFilter"))
	return rv
}

// Creates motion blur on an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/motionBlur()
func (fc _FilterClass) MotionBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("motionBlurFilter"))
	return rv
}

// Blends colors from two images by multiplying color components.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/multiplyBlendMode()
func (fc _FilterClass) MultiplyBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("multiplyBlendModeFilter"))
	return rv
}

// Blurs the colors of two images by multiplying color components.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/multiplyCompositing()
func (fc _FilterClass) MultiplyCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("multiplyCompositingFilter"))
	return rv
}

// Distorts an image by stretching it between two breakpoints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/ninePartStretched()
func (fc _FilterClass) NinePartStretchedFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("ninePartStretchedFilter"))
	return rv
}

// Distorts an image by tiling portions of it.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/ninePartTiled()
func (fc _FilterClass) NinePartTiledFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("ninePartTiledFilter"))
	return rv
}

// Reduces noise by sharpening the edges of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/noiseReduction()
func (fc _FilterClass) NoiseReductionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("noiseReductionFilter"))
	return rv
}

// Produces an effect that mimics a style of visual art that uses optical illusions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/opTile()
func (fc _FilterClass) OpTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("opTileFilter"))
	return rv
}

// Blends colors by overlaying images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/overlayBlendMode()
func (fc _FilterClass) OverlayBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("overlayBlendModeFilter"))
	return rv
}

// Simulates the curl of a page, revealing the target image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pageCurlTransition()
func (fc _FilterClass) PageCurlTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pageCurlTransitionFilter"))
	return rv
}

// Simulates the curl of a page, revealing the target image with added shadow.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pageCurlWithShadowTransition()
func (fc _FilterClass) PageCurlWithShadowTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pageCurlWithShadowTransitionFilter"))
	return rv
}

// Calculates the location of an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/paletteCentroid()
func (fc _FilterClass) PaletteCentroidFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("paletteCentroidFilter"))
	return rv
}

// Replaces colors with colors from a palette image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/palettize()
func (fc _FilterClass) PalettizeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("palettizeFilter"))
	return rv
}

// Warps the image to create a parallelogram and tiles the result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/parallelogramTile()
func (fc _FilterClass) ParallelogramTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("parallelogramTileFilter"))
	return rv
}

// Generates a high-density linear barcode.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pdf417BarcodeGenerator()
func (fc _FilterClass) PDF417BarcodeGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("PDF417BarcodeGenerator"))
	return rv
}

// Creates a mask where red pixels indicate areas of the image that are likely to contain a person.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/personSegmentation()
func (fc _FilterClass) PersonSegmentationFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("personSegmentationFilter"))
	return rv
}

// Transforms an image’s perspective.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveCorrection()
func (fc _FilterClass) PerspectiveCorrectionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("perspectiveCorrectionFilter"))
	return rv
}

// Rotates an image in a 3D space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveRotate()
func (fc _FilterClass) PerspectiveRotateFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("perspectiveRotateFilter"))
	return rv
}

// Tiles an image by adjusting the perspective of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTile()
func (fc _FilterClass) PerspectiveTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("perspectiveTileFilter"))
	return rv
}

// Alters an image’s geometry to adjust the perspective.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTransform()
func (fc _FilterClass) PerspectiveTransformFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("perspectiveTransformFilter"))
	return rv
}

// Alters an image’s geometry to adjust the perspective while applying constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTransformWithExtent()
func (fc _FilterClass) PerspectiveTransformWithExtentFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("perspectiveTransformWithExtentFilter"))
	return rv
}

// Exaggerates an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectChrome()
func (fc _FilterClass) PhotoEffectChromeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectChromeFilter"))
	return rv
}

// Diminishes an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectFade()
func (fc _FilterClass) PhotoEffectFadeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectFadeFilter"))
	return rv
}

// Desaturates an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectInstant()
func (fc _FilterClass) PhotoEffectInstantFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectInstantFilter"))
	return rv
}

// Adjust an image’s colors to black and white.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectMono()
func (fc _FilterClass) PhotoEffectMonoFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectMonoFilter"))
	return rv
}

// Adjusts an image’s colors to black and white and intensifies the contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectNoir()
func (fc _FilterClass) PhotoEffectNoirFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectNoirFilter"))
	return rv
}

// Lowers the contrast of the input image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectProcess()
func (fc _FilterClass) PhotoEffectProcessFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectProcessFilter"))
	return rv
}

// Adjusts an image’s colors to black and white.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectTonal()
func (fc _FilterClass) PhotoEffectTonalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectTonalFilter"))
	return rv
}

// Brightens an image’s colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectTransfer()
func (fc _FilterClass) PhotoEffectTransferFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("photoEffectTransferFilter"))
	return rv
}

// Blends colors of two images by replacing brighter colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pinLightBlendMode()
func (fc _FilterClass) PinLightBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pinLightBlendModeFilter"))
	return rv
}

// Distorts an image by creating a pinch effect with stronger distortion in the center.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pinchDistortion()
func (fc _FilterClass) PinchDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pinchDistortionFilter"))
	return rv
}

// Enlarges the colors of the pixels to create a blurred effect.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pixellate()
func (fc _FilterClass) PixellateFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pixellateFilter"))
	return rv
}

// Applies a pointillize effect to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pointillize()
func (fc _FilterClass) PointillizeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("pointillizeFilter"))
	return rv
}

// Generates a quick response (QR) code image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/qrCodeGenerator()
func (fc _FilterClass) QRCodeGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("QRCodeGenerator"))
	return rv
}

// Generates a gradient that varies radially between two circles having the same center.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/radialGradient()
func (fc _FilterClass) RadialGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("radialGradientFilter"))
	return rv
}

// Generates a random filter image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/randomGenerator()
func (fc _FilterClass) RandomGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("randomGeneratorFilter"))
	return rv
}

// Publishes a custom filter that is not packaged as an image unit.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/registerName(_:constructor:classAttributes:)
func (fc _FilterClass) RegisterFilterNameConstructorClassAttributes(name string, anObject unsafe.Pointer, attributes unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("registerFilterName:constructor:classAttributes:"), objc.String(name), anObject, attributes)
}

// Simulates a ripple in a pond to transiton from one image to another.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/rippleTransition()
func (fc _FilterClass) RippleTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("rippleTransitionFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedQRCodeGenerator()
func (fc _FilterClass) RoundedQRCodeGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("roundedQRCodeGeneratorFilter"))
	return rv
}

// Generates a rounded rectangle image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedRectangleGenerator()
func (fc _FilterClass) RoundedRectangleGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("roundedRectangleGeneratorFilter"))
	return rv
}

// Creates an image containing the outline of a rounded rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedRectangleStrokeGenerator()
func (fc _FilterClass) RoundedRectangleStrokeGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("roundedRectangleStrokeGeneratorFilter"))
	return rv
}

// Calculates the average color for the specified row of pixels in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/rowAverage()
func (fc _FilterClass) RowAverageFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("rowAverageFilter"))
	return rv
}

// Converts the colors in an image from sRGB to linear.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sRGBToneCurveToLinear()
func (fc _FilterClass) SRGBToneCurveToLinearFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sRGBToneCurveToLinearFilter"))
	return rv
}

// Creates a saliency map from an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/saliencyMap()
func (fc _FilterClass) SaliencyMapFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("saliencyMapFilter"))
	return rv
}

// Blends the colors and saturation values of two images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/saturationBlendMode()
func (fc _FilterClass) SaturationBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("saturationBlendModeFilter"))
	return rv
}

// Blends colors of two images by multiplying colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/screenBlendMode()
func (fc _FilterClass) ScreenBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("screenBlendModeFilter"))
	return rv
}

// Adjusts an image’s colors to shades of brown.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sepiaTone()
func (fc _FilterClass) SepiaToneFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sepiaToneFilter"))
	return rv
}

// Serializes filter parameters into XMP form that is suitable for embedding in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/serializedXMP(from:inputImageExtent:)
func (fc _FilterClass) SerializedXMPFromFiltersInputImageExtent(filters unsafe.Pointer, extent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("serializedXMPFromFilters:inputImageExtent:"), filters, extent)
	return rv
}

// Creates a shaded image from a height-field image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/shadedMaterial()
func (fc _FilterClass) ShadedMaterialFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("shadedMaterialFilter"))
	return rv
}

// Applies a sharpening effect to an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sharpenLuminance()
func (fc _FilterClass) SharpenLuminanceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sharpenLuminanceFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/signedDistanceGradientFromRedMask()
func (fc _FilterClass) SignedDistanceGradientFromRedMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("signedDistanceGradientFromRedMaskFilter"))
	return rv
}

// Produces a tiled image from a source image by applying a six-way reflected symmetry.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sixfoldReflectedTile()
func (fc _FilterClass) SixfoldReflectedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sixfoldReflectedTileFilter"))
	return rv
}

// Creates a tiled image by rotating in increments of 60 degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sixfoldRotatedTile()
func (fc _FilterClass) SixfoldRotatedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sixfoldRotatedTileFilter"))
	return rv
}

// Generates a gradient that blends colors along a linear axis between two defined endpoints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/smoothLinearGradient()
func (fc _FilterClass) SmoothLinearGradientFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("smoothLinearGradientFilter"))
	return rv
}

// Calculates the Sobel gradients for an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sobelGradients()
func (fc _FilterClass) SobelGradientsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sobelGradientsFilter"))
	return rv
}

// Blurs the colors of two images by calculating luminance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/softLightBlendMode()
func (fc _FilterClass) SoftLightBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("softLightBlendModeFilter"))
	return rv
}

// Overlaps two images to create one cropped image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceAtopCompositing()
func (fc _FilterClass) SourceAtopCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sourceAtopCompositingFilter"))
	return rv
}

// Subtracts non-overlapping areas of two images, resulting in one image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceInCompositing()
func (fc _FilterClass) SourceInCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sourceInCompositingFilter"))
	return rv
}

// Subtracts overlapping area of two images to create the output image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceOutCompositing()
func (fc _FilterClass) SourceOutCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sourceOutCompositingFilter"))
	return rv
}

// Places one image over a second image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceOverCompositing()
func (fc _FilterClass) SourceOverCompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sourceOverCompositingFilter"))
	return rv
}

// Replaces colors of an image with specifed colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/spotColor()
func (fc _FilterClass) SpotColorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("spotColorFilter"))
	return rv
}

// Highlights a definined area of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/spotLight()
func (fc _FilterClass) SpotLightFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("spotLightFilter"))
	return rv
}

// Generates a star-shine image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/starShineGenerator()
func (fc _FilterClass) StarShineGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("starShineGeneratorFilter"))
	return rv
}

// Rotates and crops an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/straighten()
func (fc _FilterClass) StraightenFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("straightenFilter"))
	return rv
}

// Distorts an image by stretching or cropping to fit a specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/stretchCrop()
func (fc _FilterClass) StretchCropFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("stretchCropFilter"))
	return rv
}

// Generates a line of stripes as an image
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/stripesGenerator()
func (fc _FilterClass) StripesGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("stripesGeneratorFilter"))
	return rv
}

// Blends colors by subtracting color values from two images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/subtractBlendMode()
func (fc _FilterClass) SubtractBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("subtractBlendModeFilter"))
	return rv
}

// Generates an image resembling the sun.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sunbeamsGenerator()
func (fc _FilterClass) SunbeamsGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("sunbeamsGeneratorFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/supportedRawCameraModels()
func (fc _FilterClass) SupportedRawCameraModels() []string {
	rv := objc.Send[[]string](objc.ID(fc.class), objc.Sel("supportedRawCameraModels"))
	return rv
}

// Gradually transitions from one image to another with a swiping motion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/swipeTransition()
func (fc _FilterClass) SwipeTransitionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("swipeTransitionFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/systemToneMap()
func (fc _FilterClass) SystemToneMapFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("systemToneMapFilter"))
	return rv
}

// Alters an image’s temperature and tint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/temperatureAndTint()
func (fc _FilterClass) TemperatureAndTintFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("temperatureAndTintFilter"))
	return rv
}

// Generates a text image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/textImageGenerator()
func (fc _FilterClass) TextImageGeneratorFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("textImageGeneratorFilter"))
	return rv
}

// Alters the image to make it look like it was taken by a thermal camera.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/thermal()
func (fc _FilterClass) ThermalFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("thermalFilter"))
	return rv
}

// Alters an image’s tone curve according to a series of data points.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/toneCurve()
func (fc _FilterClass) ToneCurveFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("toneCurveFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/toneMapHeadroom()
func (fc _FilterClass) ToneMapHeadroomFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("toneMapHeadroomFilter"))
	return rv
}

// Creates a torus-shaped lens to distort the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/torusLensDistortion()
func (fc _FilterClass) TorusLensDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("torusLensDistortionFilter"))
	return rv
}

// Create a triangular kaleidoscope effect and then tiles the result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/triangleKaleidoscope()
func (fc _FilterClass) TriangleKaleidoscopeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("triangleKaleidoscopeFilter"))
	return rv
}

// Tiles a triangular area of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/triangleTile()
func (fc _FilterClass) TriangleTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("triangleTileFilter"))
	return rv
}

// Creates a tiled image by rotating in increments of 30 degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/twelvefoldReflectedTile()
func (fc _FilterClass) TwelvefoldReflectedTileFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("twelvefoldReflectedTileFilter"))
	return rv
}

// Distorts an image by rotating pixels around a center point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/twirlDistortion()
func (fc _FilterClass) TwirlDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("twirlDistortionFilter"))
	return rv
}

// Increases an image’s contrast between two colors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/unsharpMask()
func (fc _FilterClass) UnsharpMaskFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("unsharpMaskFilter"))
	return rv
}

// Adjusts an image’s vibrancy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vibrance()
func (fc _FilterClass) VibranceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("vibranceFilter"))
	return rv
}

// Gradually darkens an image’s edges.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vignette()
func (fc _FilterClass) VignetteFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("vignetteFilter"))
	return rv
}

// Gradually darkens a specified area of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vignetteEffect()
func (fc _FilterClass) VignetteEffectFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("vignetteEffectFilter"))
	return rv
}

// A combination of color-burn and color-dodge blend modes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vividLightBlendMode()
func (fc _FilterClass) VividLightBlendModeFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("vividLightBlendModeFilter"))
	return rv
}

// Distorts an image by using a vortex effect created by rotating pixels around a point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vortexDistortion()
func (fc _FilterClass) VortexDistortionFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("vortexDistortionFilter"))
	return rv
}

// Adjusts the image’s white-point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/whitePointAdjust()
func (fc _FilterClass) WhitePointAdjustFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("whitePointAdjustFilter"))
	return rv
}

// Alters an image to make it look like an X-ray image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/xRay()
func (fc _FilterClass) XRayFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("xRayFilter"))
	return rv
}

// Creates a zoom blur centered around a single point on the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/zoomBlur()
func (fc _FilterClass) ZoomBlurFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("zoomBlurFilter"))
	return rv
}

// Produces a object by applying arguments to a kernel function and using options to control how the kernel function is evaluated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/apply(_:arguments:options:)
func (f_ Filter) ApplyArgumentsOptions(k unsafe.Pointer, args unsafe.Pointer, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("apply:arguments:options:"), k, args, dict)
	return rv
}

// Produces a object by applying a kernel function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/apply:
func (f_ Filter) Apply(k unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("apply:"), k)
	return rv
}

// Sets all input values for a filter to default values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/setDefaults()
func (f_ Filter) SetDefaults() {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDefaults"))
}

// Returns a filter view for the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/view(forUIConfiguration:excludedKeys:)
func (f_ Filter) ViewForUIConfigurationExcludedKeys(inUIConfiguration unsafe.Pointer, inKeys unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("viewForUIConfiguration:excludedKeys:"), inUIConfiguration, inKeys)
	return rv
}

// A dictionary of key-value pairs that describe the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/attributes
func (f_ Filter) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributes"))
	return rv
}
// The names of all input parameters to the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/inputKeys
func (f_ Filter) InputKeys() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("inputKeys"))
	return rv
}
// A Boolean value that determines whether the filter is enabled. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/isEnabled
func (f_ Filter) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}

// SetEnabled sets the value of the enabled property.
// A Boolean value that determines whether the filter is enabled. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/isEnabled
func (f_ Filter) SetEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEnabled:"), value)
}
// A name associated with a filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/name
func (f_ Filter) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("name"))
	return rv
}

// SetName sets the value of the name property.
// A name associated with a filter.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/name
func (f_ Filter) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}
// Returns a object that encapsulates the operations configured in the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/outputImage
func (f_ Filter) OutputImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("outputImage"))
	return rv
}
// The names of all output parameters from the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/outputKeys
func (f_ Filter) OutputKeys() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("outputKeys"))
	return rv
}

