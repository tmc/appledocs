// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageMultiply] class.
var (
	ImageMultiplyClass     _ImageMultiplyClass
	ImageMultiplyClassOnce sync.Once
)

func getImageMultiplyClass() _ImageMultiplyClass {
	ImageMultiplyClassOnce.Do(func() {
		ImageMultiplyClass = _ImageMultiplyClass{objc.GetClass("MPSImageMultiply")}
	})
	return ImageMultiplyClass
}

type _ImageMultiplyClass struct {
	class objc.Class
}





// An interface definition for the [ImageMultiply] class.
type IImageMultiply interface {
	IImageArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageMultiplyClass) Alloc() ImageMultiply {
	rv := objc.Send[ImageMultiply](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageMultiplyClass) New() ImageMultiply {
	rv := objc.Send[ImageMultiply](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageMultiply) Init() ImageMultiply {
	rv := objc.Send[ImageMultiply](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageMultiply) Autorelease() ImageMultiply {
	rv := objc.Send[ImageMultiply](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageMultiply creates a new ImageMultiply instance.
func NewImageMultiply() ImageMultiply {
	return getImageMultiplyClass().New()
}





// A filter that returns the element-wise product of its two input images.


// A filter that returns the element-wise product of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMultiply
type ImageMultiply struct {
	ImageArithmetic
}

// ImageMultiplyFrom constructs a [ImageMultiply] from an unsafe.Pointer.
//
// A filter that returns the element-wise product of its two input images.
func ImageMultiplyFrom(ptr unsafe.Pointer) ImageMultiply {
	return ImageMultiply{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemultiply/2866600-initwithdevice
func NewImageMultiplyWithDevice(device unsafe.Pointer) ImageMultiply {
	instance := getImageMultiplyClass().Alloc()
	rv := objc.Send[ImageMultiply](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























