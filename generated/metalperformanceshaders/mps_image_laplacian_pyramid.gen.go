// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageLaplacianPyramid] class.
var (
	ImageLaplacianPyramidClass     _ImageLaplacianPyramidClass
	ImageLaplacianPyramidClassOnce sync.Once
)

func getImageLaplacianPyramidClass() _ImageLaplacianPyramidClass {
	ImageLaplacianPyramidClassOnce.Do(func() {
		ImageLaplacianPyramidClass = _ImageLaplacianPyramidClass{objc.GetClass("MPSImageLaplacianPyramid")}
	})
	return ImageLaplacianPyramidClass
}

type _ImageLaplacianPyramidClass struct {
	class objc.Class
}





// An interface definition for the [ImageLaplacianPyramid] class.
type IImageLaplacianPyramid interface {
	IImagePyramid
	

	// properties:
	LaplacianBias() objectivec.IObject
	SetLaplacianBias(value objectivec.IObject)
	LaplacianScale() objectivec.IObject
	SetLaplacianScale(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidClass) Alloc() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianPyramidClass) New() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramid) Init() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramid) Autorelease() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramid creates a new ImageLaplacianPyramid instance.
func NewImageLaplacianPyramid() ImageLaplacianPyramid {
	return getImageLaplacianPyramidClass().New()
}





// A filter that convolves an image with a Laplacian filter.


// A filter that convolves an image with a Laplacian filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid
type ImageLaplacianPyramid struct {
	ImagePyramid
}

// ImageLaplacianPyramidFrom constructs a [ImageLaplacianPyramid] from an unsafe.Pointer.
//
// A filter that convolves an image with a Laplacian filter.
func ImageLaplacianPyramidFrom(ptr unsafe.Pointer) ImageLaplacianPyramid {
	return ImageLaplacianPyramid{
		ImagePyramid: ImagePyramidFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias
func (i_ ImageLaplacianPyramid) LaplacianBias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("laplacianBias"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias
func (i_ ImageLaplacianPyramid) SetLaplacianBias(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLaplacianBias:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale
func (i_ ImageLaplacianPyramid) LaplacianScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("laplacianScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale
func (i_ ImageLaplacianPyramid) SetLaplacianScale(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLaplacianScale:"), value)
}








