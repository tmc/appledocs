// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Kernel] class.
var (
	KernelClass     _KernelClass
	KernelClassOnce sync.Once
)

func getKernelClass() _KernelClass {
	KernelClassOnce.Do(func() {
		KernelClass = _KernelClass{objc.GetClass("MPSKernel")}
	})
	return KernelClass
}

type _KernelClass struct {
	class objc.Class
}

// An interface definition for the [Kernel] class.
type IKernel interface {
	objectivec.IObject
}

// A standard interface for Metal Performance Shaders kernels.
//
// You should not use the class directly. Instead, a number of subclasses are available that define specific high-performance data-parallel operations. The basic sequence for applying a kernel to an image is as follows: Initialize a kernel corresponding to the operation you wish to perform: Encode the kernel into a command buffer. Encoding the kernel merely encodes the operation into a command buffer. It does not modify any pixels, yet. All kernel state has been copied to the command buffer. Kernels may be reused. If the texture was previously operated on by another command encoder (e.g. a render command encoder), you should call the method on the other encoder before encoding the filter. Some kernels work in place, even in situations where Metal might not normally allow in-place operation on textures. If in-place operation is desired, you may attempt to call the method. If the operation cannot be completed in place, then will be returned and you will have to create a new result texture and try again. To make an in-place image filter reliable, pass a fallback block to the method to create a new texture to write to in the event that a filter cannot operate in place. You may repeat step 2 to encode more kernels, as desired. 3. After encoding any additional work to the command buffer using other encoders, submit the command buffer to your command queue, using:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel
type Kernel struct {
	objectivec.Object
}

// KernelFrom constructs a [Kernel] from an unsafe.Pointer.
//
// A standard interface for Metal Performance Shaders kernels.
func KernelFrom(ptr unsafe.Pointer) Kernel {
	return Kernel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (kc _KernelClass) Alloc() Kernel {
	rv := objc.Send[Kernel](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KernelClass) New() Kernel {
	rv := objc.Send[Kernel](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ Kernel) Init() Kernel {
	rv := objc.Send[Kernel](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ Kernel) Autorelease() Kernel {
	rv := objc.Send[Kernel](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKernel creates a new Kernel instance.
func NewKernel() Kernel {
	return getKernelClass().New()
}
