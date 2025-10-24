// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageSobel] class.
var (
	ImageSobelClass     _ImageSobelClass
	ImageSobelClassOnce sync.Once
)

func getImageSobelClass() _ImageSobelClass {
	ImageSobelClassOnce.Do(func() {
		ImageSobelClass = _ImageSobelClass{objc.GetClass("MPSImageSobel")}
	})
	return ImageSobelClass
}

type _ImageSobelClass struct {
	class objc.Class
}





// An interface definition for the [ImageSobel] class.
type IImageSobel interface {
	IUnaryImageKernel
	

	// properties:
	ColorTransform() objectivec.IObject
	SetColorTransform(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageSobelClass) Alloc() ImageSobel {
	rv := objc.Send[ImageSobel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSobelClass) New() ImageSobel {
	rv := objc.Send[ImageSobel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSobel) Init() ImageSobel {
	rv := objc.Send[ImageSobel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSobel) Autorelease() ImageSobel {
	rv := objc.Send[ImageSobel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSobel creates a new ImageSobel instance.
func NewImageSobel() ImageSobel {
	return getImageSobelClass().New()
}





// A filter that convolves an image with the Sobel operator.
//
// When the color model (e.g. RGB, two-channel, grayscale, etc.) of the source and destination textures match, the filter is applied to each color channel separately. If the destination is single-channel (i.e. monochrome) but the source is multi-channel, the pixel values are converted to grayscale before applying the Sobel operator by using the linear gray color transform vector shown in the code listing below.


// A filter that convolves an image with the Sobel operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel
type ImageSobel struct {
	UnaryImageKernel
}

// ImageSobelFrom constructs a [ImageSobel] from an unsafe.Pointer.
//
// A filter that convolves an image with the Sobel operator.
func ImageSobelFrom(ptr unsafe.Pointer) ImageSobel {
	return ImageSobel{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/2866152-initwithcoder
func NewImageSobelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a Sobel filter on a given device using the default color transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618843-initwithdevice
func NewImageSobelWithDevice(device unsafe.Pointer) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Initializes a Sobel filter on a given device using a specific color transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618899-initwithdevice
func NewImageSobelWithDeviceLinearGrayColorTransform(device unsafe.Pointer, transform objectivec.IObject) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithDevice:linearGrayColorTransform:"), device, transform)
	rv.Autorelease()
	return rv
}






















// The color transform used to initialize the Sobel filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618777-colortransform
func (i_ ImageSobel) ColorTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("colorTransform"))
	return rv
}


// The color transform used to initialize the Sobel filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618777-colortransform
func (i_ ImageSobel) SetColorTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorTransform:"), value)
}







