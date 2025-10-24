// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageCanny] class.
var (
	ImageCannyClass     _ImageCannyClass
	ImageCannyClassOnce sync.Once
)

func getImageCannyClass() _ImageCannyClass {
	ImageCannyClassOnce.Do(func() {
		ImageCannyClass = _ImageCannyClass{objc.GetClass("MPSImageCanny")}
	})
	return ImageCannyClass
}

type _ImageCannyClass struct {
	class objc.Class
}





// An interface definition for the [ImageCanny] class.
type IImageCanny interface {
	IUnaryImageKernel
	

	// properties:
	ColorTransform() objectivec.IObject
	SetColorTransform(value objectivec.IObject)
	HighThreshold() objectivec.IObject
	SetHighThreshold(value objectivec.IObject)
	LowThreshold() objectivec.IObject
	SetLowThreshold(value objectivec.IObject)
	Sigma() objectivec.IObject
	SetSigma(value objectivec.IObject)
	UseFastMode() objectivec.IObject
	SetUseFastMode(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageCannyClass) Alloc() ImageCanny {
	rv := objc.Send[ImageCanny](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageCannyClass) New() ImageCanny {
	rv := objc.Send[ImageCanny](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageCanny) Init() ImageCanny {
	rv := objc.Send[ImageCanny](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageCanny) Autorelease() ImageCanny {
	rv := objc.Send[ImageCanny](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageCanny creates a new ImageCanny instance.
func NewImageCanny() ImageCanny {
	return getImageCannyClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny
type ImageCanny struct {
	UnaryImageKernel
}

// ImageCannyFrom constructs a [ImageCanny] from an unsafe.Pointer.
func ImageCannyFrom(ptr unsafe.Pointer) ImageCanny {
	return ImageCanny{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547971-initwithcoder
func NewImageCannyWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547972-initwithdevice
func NewImageCannyWithDevice(device unsafe.Pointer) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547973-initwithdevice
func NewImageCannyWithDeviceLinearToGrayScaleTransformSigma(device unsafe.Pointer, transform objectivec.IObject, sigma float32) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithDevice:linearToGrayScaleTransform:sigma:"), device, transform, sigma)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547969-colortransform
func (i_ ImageCanny) ColorTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("colorTransform"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547969-colortransform
func (i_ ImageCanny) SetColorTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorTransform:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold
func (i_ ImageCanny) HighThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("highThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold
func (i_ ImageCanny) SetHighThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold
func (i_ ImageCanny) LowThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("lowThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold
func (i_ ImageCanny) SetLowThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLowThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547975-sigma
func (i_ ImageCanny) Sigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("sigma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547975-sigma
func (i_ ImageCanny) SetSigma(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSigma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode
func (i_ ImageCanny) UseFastMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("useFastMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode
func (i_ ImageCanny) SetUseFastMode(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUseFastMode:"), value)
}







