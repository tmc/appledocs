// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageAdd] class.
var (
	ImageAddClass     _ImageAddClass
	ImageAddClassOnce sync.Once
)

func getImageAddClass() _ImageAddClass {
	ImageAddClassOnce.Do(func() {
		ImageAddClass = _ImageAddClass{objc.GetClass("MPSImageAdd")}
	})
	return ImageAddClass
}

type _ImageAddClass struct {
	class objc.Class
}





// An interface definition for the [ImageAdd] class.
type IImageAdd interface {
	IImageArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageAddClass) Alloc() ImageAdd {
	rv := objc.Send[ImageAdd](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAddClass) New() ImageAdd {
	rv := objc.Send[ImageAdd](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAdd) Init() ImageAdd {
	rv := objc.Send[ImageAdd](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAdd) Autorelease() ImageAdd {
	rv := objc.Send[ImageAdd](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAdd creates a new ImageAdd instance.
func NewImageAdd() ImageAdd {
	return getImageAddClass().New()
}





// A filter that returns the element-wise sum of its two input images.


// A filter that returns the element-wise sum of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAdd
type ImageAdd struct {
	ImageArithmetic
}

// ImageAddFrom constructs a [ImageAdd] from an unsafe.Pointer.
//
// A filter that returns the element-wise sum of its two input images.
func ImageAddFrom(ptr unsafe.Pointer) ImageAdd {
	return ImageAdd{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageadd/2866610-initwithdevice
func NewImageAddWithDevice(device unsafe.Pointer) ImageAdd {
	instance := getImageAddClass().Alloc()
	rv := objc.Send[ImageAdd](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























