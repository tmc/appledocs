// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageTent] class.
var (
	ImageTentClass     _ImageTentClass
	ImageTentClassOnce sync.Once
)

func getImageTentClass() _ImageTentClass {
	ImageTentClassOnce.Do(func() {
		ImageTentClass = _ImageTentClass{objc.GetClass("MPSImageTent")}
	})
	return ImageTentClass
}

type _ImageTentClass struct {
	class objc.Class
}





// An interface definition for the [ImageTent] class.
type IImageTent interface {
	IImageBox
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageTentClass) Alloc() ImageTent {
	rv := objc.Send[ImageTent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageTentClass) New() ImageTent {
	rv := objc.Send[ImageTent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageTent) Init() ImageTent {
	rv := objc.Send[ImageTent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageTent) Autorelease() ImageTent {
	rv := objc.Send[ImageTent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageTent creates a new ImageTent instance.
func NewImageTent() ImageTent {
	return getImageTentClass().New()
}





// A filter that convolves an image with a tent filter.
//
// The kernel elements of the filter form a tent shape with increasing sides, for example: Like a box filter, this arrangement allows for much faster algorithms, especially for larger blur radii but with a more pleasing appearance. The tent blur is a separable filter and the Metal Performance Shaders framework will act accordingly to give the best performance for multi-dimensional blurs.


// A filter that convolves an image with a tent filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTent
type ImageTent struct {
	ImageBox
}

// ImageTentFrom constructs a [ImageTent] from an unsafe.Pointer.
//
// A filter that convolves an image with a tent filter.
func ImageTentFrom(ptr unsafe.Pointer) ImageTent {
	return ImageTent{
		ImageBox: ImageBoxFrom(ptr),
	}
}































