// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageGuidedFilter] class.
var (
	ImageGuidedFilterClass     _ImageGuidedFilterClass
	ImageGuidedFilterClassOnce sync.Once
)

func getImageGuidedFilterClass() _ImageGuidedFilterClass {
	ImageGuidedFilterClassOnce.Do(func() {
		ImageGuidedFilterClass = _ImageGuidedFilterClass{objc.GetClass("MPSImageGuidedFilter")}
	})
	return ImageGuidedFilterClass
}

type _ImageGuidedFilterClass struct {
	class objc.Class
}

// An interface definition for the [ImageGuidedFilter] class.
type IImageGuidedFilter interface {
	IKernel
}

// A filter that performs edge-aware filtering on an image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter
type ImageGuidedFilter struct {
	Kernel
}

// ImageGuidedFilterFrom constructs a [ImageGuidedFilter] from an unsafe.Pointer.
//
// A filter that performs edge-aware filtering on an image.
func ImageGuidedFilterFrom(ptr unsafe.Pointer) ImageGuidedFilter {
	return ImageGuidedFilter{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageGuidedFilterClass) Alloc() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageGuidedFilterClass) New() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageGuidedFilter) Init() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageGuidedFilter) Autorelease() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageGuidedFilter creates a new ImageGuidedFilter instance.
func NewImageGuidedFilter() ImageGuidedFilter {
	return getImageGuidedFilterClass().New()
}




