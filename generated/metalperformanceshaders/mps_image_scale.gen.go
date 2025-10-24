// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageScale] class.
var (
	ImageScaleClass     _ImageScaleClass
	ImageScaleClassOnce sync.Once
)

func getImageScaleClass() _ImageScaleClass {
	ImageScaleClassOnce.Do(func() {
		ImageScaleClass = _ImageScaleClass{objc.GetClass("MPSImageScale")}
	})
	return ImageScaleClass
}

type _ImageScaleClass struct {
	class objc.Class
}

// An interface definition for the [ImageScale] class.
type IImageScale interface {
	IUnaryImageKernel
	// properties:
	ScaleTransform() MPSScaleTransform /* not a class type */
	SetScaleTransform(value MPSScaleTransform /* not a class type */)
	// methods:
}

// A filter that resizes and changes the aspect ratio of an image.


// A filter that resizes and changes the aspect ratio of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale
type ImageScale struct {
	UnaryImageKernel
}

// ImageScaleFrom constructs a [ImageScale] from an unsafe.Pointer.
//
// A filter that resizes and changes the aspect ratio of an image.
func ImageScaleFrom(ptr unsafe.Pointer) ImageScale {
	return ImageScale{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageScaleClass) Alloc() ImageScale {
	rv := objc.Send[ImageScale](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageScaleClass) New() ImageScale {
	rv := objc.Send[ImageScale](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageScale) Init() ImageScale {
	rv := objc.Send[ImageScale](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageScale) Autorelease() ImageScale {
	rv := objc.Send[ImageScale](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageScale creates a new ImageScale instance.
func NewImageScale() ImageScale {
	return getImageScaleClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/init(coder:device:)
func NewImageScaleWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImageScale {
	instance := getImageScaleClass().Alloc()
	rv := objc.Send[ImageScale](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/init(device:)
func NewImageScaleWithDevice(device objectivec.IObject) ImageScale {
	instance := getImageScaleClass().Alloc()
	rv := objc.Send[ImageScale](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/scaleTransform
func (i_ ImageScale) ScaleTransform() MPSScaleTransform /* not a class type */ {
	rv := objc.Send[ScaleTransform](i_.ID, objc.Sel("scaleTransform"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/scaleTransform
func (i_ ImageScale) SetScaleTransform(value MPSScaleTransform /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScaleTransform:"), value)
}


