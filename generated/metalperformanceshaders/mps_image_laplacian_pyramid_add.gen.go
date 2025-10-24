// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageLaplacianPyramidAdd] class.
var (
	ImageLaplacianPyramidAddClass     _ImageLaplacianPyramidAddClass
	ImageLaplacianPyramidAddClassOnce sync.Once
)

func getImageLaplacianPyramidAddClass() _ImageLaplacianPyramidAddClass {
	ImageLaplacianPyramidAddClassOnce.Do(func() {
		ImageLaplacianPyramidAddClass = _ImageLaplacianPyramidAddClass{objc.GetClass("MPSImageLaplacianPyramidAdd")}
	})
	return ImageLaplacianPyramidAddClass
}

type _ImageLaplacianPyramidAddClass struct {
	class objc.Class
}

// An interface definition for the [ImageLaplacianPyramidAdd] class.
type IImageLaplacianPyramidAdd interface {
	IImageLaplacianPyramid
	// properties:
	// methods:
}

// A filter that convolves an image with an additive Laplacian pyramid.


// A filter that convolves an image with an additive Laplacian pyramid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramidAdd
type ImageLaplacianPyramidAdd struct {
	ImageLaplacianPyramid
}

// ImageLaplacianPyramidAddFrom constructs a [ImageLaplacianPyramidAdd] from an unsafe.Pointer.
//
// A filter that convolves an image with an additive Laplacian pyramid.
func ImageLaplacianPyramidAddFrom(ptr unsafe.Pointer) ImageLaplacianPyramidAdd {
	return ImageLaplacianPyramidAdd{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidAddClass) Alloc() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageLaplacianPyramidAddClass) New() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramidAdd) Init() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramidAdd) Autorelease() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramidAdd creates a new ImageLaplacianPyramidAdd instance.
func NewImageLaplacianPyramidAdd() ImageLaplacianPyramidAdd {
	return getImageLaplacianPyramidAddClass().New()
}




