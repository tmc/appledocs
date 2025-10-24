// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageTranspose] class.
var (
	ImageTransposeClass     _ImageTransposeClass
	ImageTransposeClassOnce sync.Once
)

func getImageTransposeClass() _ImageTransposeClass {
	ImageTransposeClassOnce.Do(func() {
		ImageTransposeClass = _ImageTransposeClass{objc.GetClass("MPSImageTranspose")}
	})
	return ImageTransposeClass
}

type _ImageTransposeClass struct {
	class objc.Class
}

// An interface definition for the [ImageTranspose] class.
type IImageTranspose interface {
	IUnaryImageKernel
	// properties:
	// methods:
}

// A filter that transposes an image.
//
// An filter applies a matrix transposition to the source image by exchanging its rows with its columns.


// A filter that transposes an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTranspose
type ImageTranspose struct {
	UnaryImageKernel
}

// ImageTransposeFrom constructs a [ImageTranspose] from an unsafe.Pointer.
//
// A filter that transposes an image.
func ImageTransposeFrom(ptr unsafe.Pointer) ImageTranspose {
	return ImageTranspose{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageTransposeClass) Alloc() ImageTranspose {
	rv := objc.Send[ImageTranspose](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageTransposeClass) New() ImageTranspose {
	rv := objc.Send[ImageTranspose](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageTranspose) Init() ImageTranspose {
	rv := objc.Send[ImageTranspose](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageTranspose) Autorelease() ImageTranspose {
	rv := objc.Send[ImageTranspose](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageTranspose creates a new ImageTranspose instance.
func NewImageTranspose() ImageTranspose {
	return getImageTransposeClass().New()
}




