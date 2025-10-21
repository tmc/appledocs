// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
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
	DilationRateYatIndex(index uint) uint
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateYatIndex(_:)
func (c_ CNNMultiaryKernel) DilationRateYatIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateYatIndex:"), index)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/cliprect
func (c_ CNNMultiaryKernel) ClipRect() corelocation.Region {
	rv := objc.Send[corelocation.Region](c_.ID, objc.Sel("clipRect"))
	return rv
}


// SetClipRect sets the value of the clipRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/cliprect
func (c_ CNNMultiaryKernel) SetClipRect(value corelocation.IRegion) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// SetDestinationFeatureChannelOffset sets the value of the destinationFeatureChannelOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationimageallocator
func (c_ CNNMultiaryKernel) DestinationImageAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// SetDestinationImageAllocator sets the value of the destinationImageAllocator property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/destinationimageallocator
func (c_ CNNMultiaryKernel) SetDestinationImageAllocator(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isbackwards
func (c_ CNNMultiaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// SetIsBackwards sets the value of the isBackwards property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isbackwards
func (c_ CNNMultiaryKernel) SetIsBackwards(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isstatemodified
func (c_ CNNMultiaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// SetIsStateModified sets the value of the isStateModified property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/isstatemodified
func (c_ CNNMultiaryKernel) SetIsStateModified(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/padding
func (c_ CNNMultiaryKernel) Padding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("padding"))
	return rv
}


// SetPadding sets the value of the padding property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/padding
func (c_ CNNMultiaryKernel) SetPadding(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/sourcecount
func (c_ CNNMultiaryKernel) SourceCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sourceCount"))
	return rv
}


// SetSourceCount sets the value of the sourceCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/sourcecount
func (c_ CNNMultiaryKernel) SetSourceCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceCount:"), value)
}



