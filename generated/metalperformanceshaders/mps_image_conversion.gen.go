// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageConversion] class.
var (
	ImageConversionClass     _ImageConversionClass
	ImageConversionClassOnce sync.Once
)

func getImageConversionClass() _ImageConversionClass {
	ImageConversionClassOnce.Do(func() {
		ImageConversionClass = _ImageConversionClass{objc.GetClass("MPSImageConversion")}
	})
	return ImageConversionClass
}

type _ImageConversionClass struct {
	class objc.Class
}

// An interface definition for the [ImageConversion] class.
type IImageConversion interface {
	IUnaryImageKernel
	// properties:
	DestinationAlpha() AlphaType
	SourceAlpha() AlphaType
	// methods:
}

// A filter that performs a conversion of color space, alpha, or pixel format.
//
// An filter allows you to change the alpha encoding or color space of an image. For example, you can convert an image with a premultiplied alpha to non-premultiplied, or change the color space from one variant to another. As with all Metal Performance Shaders filters, the conversion filter allows for source and destination textures with different pixel formats and, in that case, will convert the source texture’s format to the destination texture’s format. See for a list of supported pixel formats. The following listing shows how you can create an image conversion filter to map the color intensity from the sRGB color space to a linear gamma curve. Listing 1. Mapping color intensity from the sRGB color space to a linear gamma curve.


// A filter that performs a conversion of color space, alpha, or pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion
type ImageConversion struct {
	UnaryImageKernel
}

// ImageConversionFrom constructs a [ImageConversion] from an unsafe.Pointer.
//
// A filter that performs a conversion of color space, alpha, or pixel format.
func ImageConversionFrom(ptr unsafe.Pointer) ImageConversion {
	return ImageConversion{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageConversionClass) Alloc() ImageConversion {
	rv := objc.Send[ImageConversion](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageConversionClass) New() ImageConversion {
	rv := objc.Send[ImageConversion](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageConversion) Init() ImageConversion {
	rv := objc.Send[ImageConversion](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageConversion) Autorelease() ImageConversion {
	rv := objc.Send[ImageConversion](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageConversion creates a new ImageConversion instance.
func NewImageConversion() ImageConversion {
	return getImageConversionClass().New()
}



// Initializes a filter that can convert texture color space, alpha, and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/init(device:srcAlpha:destAlpha:backgroundColor:conversionInfo:)
func NewImageConversionWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device objectivec.IObject, srcAlpha AlphaType, destAlpha AlphaType, backgroundColor corefoundation.CGFloat, conversionInfo ColorConversionInfoRef /* not a class type */) ImageConversion {
	instance := getImageConversionClass().Alloc()
	rv := objc.Send[ImageConversion](instance.ID, objc.Sel("initWithDevice:srcAlpha:destAlpha:backgroundColor:conversionInfo:"), device, srcAlpha, destAlpha, backgroundColor, conversionInfo)
	rv.Autorelease()
	return rv
}



// Premultiplication description for the destination texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/destinationAlpha
func (i_ ImageConversion) DestinationAlpha() AlphaType {
	rv := objc.Send[AlphaType](i_.ID, objc.Sel("destinationAlpha"))
	return rv
}


// Premultiplication description for the source texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/sourceAlpha
func (i_ ImageConversion) SourceAlpha() AlphaType {
	rv := objc.Send[AlphaType](i_.ID, objc.Sel("sourceAlpha"))
	return rv
}


