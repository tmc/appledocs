// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageReduceUnary] class.
var (
	ImageReduceUnaryClass     _ImageReduceUnaryClass
	ImageReduceUnaryClassOnce sync.Once
)

func getImageReduceUnaryClass() _ImageReduceUnaryClass {
	ImageReduceUnaryClassOnce.Do(func() {
		ImageReduceUnaryClass = _ImageReduceUnaryClass{objc.GetClass("MPSImageReduceUnary")}
	})
	return ImageReduceUnaryClass
}

type _ImageReduceUnaryClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceUnary] class.
type IImageReduceUnary interface {
	IUnaryImageKernel
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceUnaryClass) Alloc() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceUnaryClass) New() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceUnary) Init() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceUnary) Autorelease() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceUnary creates a new ImageReduceUnary instance.
func NewImageReduceUnary() ImageReduceUnary {
	return getImageReduceUnaryClass().New()
}





// The base class for reduction filters that take a single source as input.


// The base class for reduction filters that take a single source as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceUnary
type ImageReduceUnary struct {
	UnaryImageKernel
}

// ImageReduceUnaryFrom constructs a [ImageReduceUnary] from an unsafe.Pointer.
//
// The base class for reduction filters that take a single source as input.
func ImageReduceUnaryFrom(ptr unsafe.Pointer) ImageReduceUnary {
	return ImageReduceUnary{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource
func (i_ ImageReduceUnary) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource
func (i_ ImageReduceUnary) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}








