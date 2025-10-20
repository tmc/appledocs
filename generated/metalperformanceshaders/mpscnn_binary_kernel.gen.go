// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNBinaryKernel] class.
var (
	CNNBinaryKernelClass     _CNNBinaryKernelClass
	CNNBinaryKernelClassOnce sync.Once
)

func getCNNBinaryKernelClass() _CNNBinaryKernelClass {
	CNNBinaryKernelClassOnce.Do(func() {
		CNNBinaryKernelClass = _CNNBinaryKernelClass{objc.GetClass("MPSCNNBinaryKernel")}
	})
	return CNNBinaryKernelClass
}

type _CNNBinaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [CNNBinaryKernel] class.
type ICNNBinaryKernel interface {
	IKernel
}

// A convolution neural network kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel
type CNNBinaryKernel struct {
	Kernel
}

// CNNBinaryKernelFrom constructs a [CNNBinaryKernel] from an unsafe.Pointer.
//
// A convolution neural network kernel.
func CNNBinaryKernelFrom(ptr unsafe.Pointer) CNNBinaryKernel {
	return CNNBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryKernelClass) Alloc() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNBinaryKernelClass) New() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryKernel) Init() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryKernel) Autorelease() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryKernel creates a new CNNBinaryKernel instance.
func NewCNNBinaryKernel() CNNBinaryKernel {
	return getCNNBinaryKernelClass().New()
}




