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
