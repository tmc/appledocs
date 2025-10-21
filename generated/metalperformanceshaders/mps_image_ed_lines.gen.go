// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageEDLines] class.
var (
	ImageEDLinesClass     _ImageEDLinesClass
	ImageEDLinesClassOnce sync.Once
)

func getImageEDLinesClass() _ImageEDLinesClass {
	ImageEDLinesClassOnce.Do(func() {
		ImageEDLinesClass = _ImageEDLinesClass{objc.GetClass("MPSImageEDLines")}
	})
	return ImageEDLinesClass
}

type _ImageEDLinesClass struct {
	class objc.Class
}

// An interface definition for the [ImageEDLines] class.
type IImageEDLines interface {
	IKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines
type ImageEDLines struct {
	Kernel
}

// ImageEDLinesFrom constructs a [ImageEDLines] from an unsafe.Pointer.
func ImageEDLinesFrom(ptr unsafe.Pointer) ImageEDLines {
	return ImageEDLines{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageEDLinesClass) Alloc() ImageEDLines {
	rv := objc.Send[ImageEDLines](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageEDLinesClass) New() ImageEDLines {
	rv := objc.Send[ImageEDLines](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageEDLines) Init() ImageEDLines {
	rv := objc.Send[ImageEDLines](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageEDLines) Autorelease() ImageEDLines {
	rv := objc.Send[ImageEDLines](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageEDLines creates a new ImageEDLines instance.
func NewImageEDLines() ImageEDLines {
	return getImageEDLinesClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/minlinelength
func (i_ ImageEDLines) MinLineLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("minLineLength"))
	return rv
}


// SetMinLineLength sets the value of the minLineLength property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/minlinelength
func (i_ ImageEDLines) SetMinLineLength(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinLineLength:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/detailratio
func (i_ ImageEDLines) DetailRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("detailRatio"))
	return rv
}


// SetDetailRatio sets the value of the detailRatio property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/detailratio
func (i_ ImageEDLines) SetDetailRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDetailRatio:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/maxlines
func (i_ ImageEDLines) MaxLines() int {
	rv := objc.Send[int](i_.ID, objc.Sel("maxLines"))
	return rv
}


// SetMaxLines sets the value of the maxLines property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/maxlines
func (i_ ImageEDLines) SetMaxLines(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxLines:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/gaussiansigma
func (i_ ImageEDLines) GaussianSigma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("gaussianSigma"))
	return rv
}


// SetGaussianSigma sets the value of the gaussianSigma property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/gaussiansigma
func (i_ ImageEDLines) SetGaussianSigma(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGaussianSigma:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/lineerrorthreshold
func (i_ ImageEDLines) LineErrorThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("lineErrorThreshold"))
	return rv
}


// SetLineErrorThreshold sets the value of the lineErrorThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/lineerrorthreshold
func (i_ ImageEDLines) SetLineErrorThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLineErrorThreshold:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/clipRectSource
func (i_ ImageEDLines) ClipRectSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// SetClipRectSource sets the value of the clipRectSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/clipRectSource
func (i_ ImageEDLines) SetClipRectSource(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gradientThreshold
func (i_ ImageEDLines) GradientThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("gradientThreshold"))
	return rv
}


// SetGradientThreshold sets the value of the gradientThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gradientThreshold
func (i_ ImageEDLines) SetGradientThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGradientThreshold:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/mergeLocalityThreshold
func (i_ ImageEDLines) MergeLocalityThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mergeLocalityThreshold"))
	return rv
}


// SetMergeLocalityThreshold sets the value of the mergeLocalityThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/mergeLocalityThreshold
func (i_ ImageEDLines) SetMergeLocalityThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMergeLocalityThreshold:"), value)
}



