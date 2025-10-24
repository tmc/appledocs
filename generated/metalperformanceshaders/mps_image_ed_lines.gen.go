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
	// properties:
	ClipRectSource() objc.IObject /* cross-framework: MTLRegion */
	SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */)
	GradientThreshold() float32
	SetGradientThreshold(value float32)
	MergeLocalityThreshold() float32
	SetMergeLocalityThreshold(value float32)
	DetailRatio() unsafe.Pointer
	SetDetailRatio(value unsafe.Pointer)
	GaussianSigma() float32
	SetGaussianSigma(value float32)
	LineErrorThreshold() float32
	SetLineErrorThreshold(value float32)
	MaxLines() int
	SetMaxLines(value int)
	MinLineLength() unsafe.Pointer
	SetMinLineLength(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/clipRectSource
func (i_ ImageEDLines) ClipRectSource() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/clipRectSource
func (i_ ImageEDLines) SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gradientThreshold
func (i_ ImageEDLines) GradientThreshold() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("gradientThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gradientThreshold
func (i_ ImageEDLines) SetGradientThreshold(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGradientThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/mergeLocalityThreshold
func (i_ ImageEDLines) MergeLocalityThreshold() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("mergeLocalityThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/mergeLocalityThreshold
func (i_ ImageEDLines) SetMergeLocalityThreshold(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMergeLocalityThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/detailratio
func (i_ ImageEDLines) DetailRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("detailRatio"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/detailratio
func (i_ ImageEDLines) SetDetailRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDetailRatio:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/gaussiansigma
func (i_ ImageEDLines) GaussianSigma() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("gaussianSigma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/gaussiansigma
func (i_ ImageEDLines) SetGaussianSigma(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGaussianSigma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/lineerrorthreshold
func (i_ ImageEDLines) LineErrorThreshold() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("lineErrorThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/lineerrorthreshold
func (i_ ImageEDLines) SetLineErrorThreshold(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLineErrorThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/maxlines
func (i_ ImageEDLines) MaxLines() int {
	rv := objc.Send[int](i_.ID, objc.Sel("maxLines"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/maxlines
func (i_ ImageEDLines) SetMaxLines(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxLines:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/minlinelength
func (i_ ImageEDLines) MinLineLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("minLineLength"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/minlinelength
func (i_ ImageEDLines) SetMinLineLength(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinLineLength:"), value)
}



