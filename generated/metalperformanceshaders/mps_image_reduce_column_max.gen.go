// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageReduceColumnMax] class.
var (
	ImageReduceColumnMaxClass     _ImageReduceColumnMaxClass
	ImageReduceColumnMaxClassOnce sync.Once
)

func getImageReduceColumnMaxClass() _ImageReduceColumnMaxClass {
	ImageReduceColumnMaxClassOnce.Do(func() {
		ImageReduceColumnMaxClass = _ImageReduceColumnMaxClass{objc.GetClass("MPSImageReduceColumnMax")}
	})
	return ImageReduceColumnMaxClass
}

type _ImageReduceColumnMaxClass struct {
	class objc.Class
}

// An interface definition for the [ImageReduceColumnMax] class.
type IImageReduceColumnMax interface {
	IImageReduceUnary
	// properties:
	// methods:
}

// A filter that returns the maximum value for each column in an image.


// A filter that returns the maximum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMax
type ImageReduceColumnMax struct {
	ImageReduceUnary
}

// ImageReduceColumnMaxFrom constructs a [ImageReduceColumnMax] from an unsafe.Pointer.
//
// A filter that returns the maximum value for each column in an image.
func ImageReduceColumnMaxFrom(ptr unsafe.Pointer) ImageReduceColumnMax {
	return ImageReduceColumnMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnMaxClass) Alloc() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageReduceColumnMaxClass) New() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnMax) Init() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnMax) Autorelease() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnMax creates a new ImageReduceColumnMax instance.
func NewImageReduceColumnMax() ImageReduceColumnMax {
	return getImageReduceColumnMaxClass().New()
}




