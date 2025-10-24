// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNMultiaryKernel] class.
var (
	CNNMultiaryKernelClass     _CNNMultiaryKernelClass
	CNNMultiaryKernelClassOnce sync.Once
)

func getCNNMultiaryKernelClass() _CNNMultiaryKernelClass {
	CNNMultiaryKernelClassOnce.Do(func() {
		CNNMultiaryKernelClass = _CNNMultiaryKernelClass{objc.GetClass("MPSCNNMultiaryKernel")}
	})
	return CNNMultiaryKernelClass
}

type _CNNMultiaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [CNNMultiaryKernel] class.
type ICNNMultiaryKernel interface {
	IKernel
	// properties:
	ClipRect() objc.IObject /* cross-framework: MTLRegion */
	SetClipRect(value objc.IObject /* cross-framework: MTLRegion */)
	DestinationFeatureChannelOffset() int
	SetDestinationFeatureChannelOffset(value int)
	DestinationImageAllocator() ImageAllocator /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator /* not a class type */)
	IsBackwards() bool
	SetIsBackwards(value bool)
	IsStateModified() bool
	SetIsStateModified(value bool)
	Padding() Padding /* not a class type */
	SetPadding(value Padding /* not a class type */)
	SourceCount() int
	SetSourceCount(value int)
	// methods:
	DilationRateYatIndex(index uint) uint
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel
type CNNMultiaryKernel struct {
	Kernel
}

// CNNMultiaryKernelFrom constructs a [CNNMultiaryKernel] from an unsafe.Pointer.
func CNNMultiaryKernelFrom(ptr unsafe.Pointer) CNNMultiaryKernel {
	return CNNMultiaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNMultiaryKernelClass) Alloc() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNMultiaryKernelClass) New() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiaryKernel) Init() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiaryKernel) Autorelease() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiaryKernel creates a new CNNMultiaryKernel instance.
func NewCNNMultiaryKernel() CNNMultiaryKernel {
	return getCNNMultiaryKernelClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateYatIndex(_:)
func (c_ CNNMultiaryKernel) DilationRateYatIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateYatIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/cliprect
func (c_ CNNMultiaryKernel) ClipRect() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](c_.ID, objc.Sel("clipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/cliprect
func (c_ CNNMultiaryKernel) SetClipRect(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationimageallocator
func (c_ CNNMultiaryKernel) DestinationImageAllocator() ImageAllocator /* not a class type */ {
	rv := objc.Send[ImageAllocator](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationimageallocator
func (c_ CNNMultiaryKernel) SetDestinationImageAllocator(value ImageAllocator /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isbackwards
func (c_ CNNMultiaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isbackwards
func (c_ CNNMultiaryKernel) SetIsBackwards(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isstatemodified
func (c_ CNNMultiaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isstatemodified
func (c_ CNNMultiaryKernel) SetIsStateModified(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/padding
func (c_ CNNMultiaryKernel) Padding() Padding /* not a class type */ {
	rv := objc.Send[Padding](c_.ID, objc.Sel("padding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/padding
func (c_ CNNMultiaryKernel) SetPadding(value Padding /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/sourcecount
func (c_ CNNMultiaryKernel) SourceCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sourceCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/sourcecount
func (c_ CNNMultiaryKernel) SetSourceCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceCount:"), value)
}



