// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Sigma() float32
	UseFastMode() bool
	SetUseFastMode(value bool)
	ColorTransform() float32
	SetColorTransform(value float32)
	HighThreshold() float32
	SetHighThreshold(value float32)
	LowThreshold() float32
	SetLowThreshold(value float32)
}

//
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

// Alloc allocates a new instance without initialization.
func (ic _ImageCannyClass) Alloc() ImageCanny {
	rv := objc.Send[ImageCanny](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/sigma
func (i_ ImageCanny) Sigma() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("sigma"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/useFastMode
func (i_ ImageCanny) UseFastMode() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("useFastMode"))
	return rv
}


// SetUseFastMode sets the value of the useFastMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/useFastMode
func (i_ ImageCanny) SetUseFastMode(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUseFastMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/colortransform
func (i_ ImageCanny) ColorTransform() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("colorTransform"))
	return rv
}


// SetColorTransform sets the value of the colorTransform property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/colortransform
func (i_ ImageCanny) SetColorTransform(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorTransform:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/highthreshold
func (i_ ImageCanny) HighThreshold() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("highThreshold"))
	return rv
}


// SetHighThreshold sets the value of the highThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/highthreshold
func (i_ ImageCanny) SetHighThreshold(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighThreshold:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/lowthreshold
func (i_ ImageCanny) LowThreshold() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("lowThreshold"))
	return rv
}


// SetLowThreshold sets the value of the lowThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/lowthreshold
func (i_ ImageCanny) SetLowThreshold(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLowThreshold:"), value)
}



