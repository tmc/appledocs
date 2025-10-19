// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Kernel] class.
var (
	kernelClass     _KernelClass
	kernelClassOnce sync.Once
)

func getKernelClass() _KernelClass {
	kernelClassOnce.Do(func() {
		kernelClass = _KernelClass{objc.GetClass("CIKernel")}
	})
	return kernelClass
}

type _KernelClass struct {
	class objc.Class
}

// An interface definition for the [Kernel] class.
type IKernel interface {
	objectivec.IObject
	ApplyWithExtentRoiCallbackArguments(extent unsafe.Pointer, callback unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer
	SetROISelector(method objc.SEL)
}

// A GPU-based image-processing routine used to create custom Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel
type Kernel struct {
	objectivec.Object
}

// KernelFrom constructs a [Kernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine used to create custom Core Image filters.
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


// Creates a single kernel object using a Metal Shading Language (MSL) kernel function. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
func NewKernelWithFunctionNameFromMetalLibraryDataError(name string, data unsafe.Pointer, error unsafe.Pointer) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, error)
	return rv
}
// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:outputPixelFormat:)
func NewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data unsafe.Pointer, format unsafe.Pointer, error unsafe.Pointer) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, error)
	return rv
}
// Creates a single kernel object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(source:)
func NewKernelWithString(string string) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}


// Creates a single kernel object using a Metal Shading Language (MSL) kernel function. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
func (kc _KernelClass) KernelWithFunctionNameFromMetalLibraryDataError(name string, data unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, error)
	return rv
}
// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:outputPixelFormat:)
func (kc _KernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data unsafe.Pointer, format unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, error)
	return rv
}
// Creates a single kernel object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(source:)
func (kc _KernelClass) KernelWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}
// Return an array of strings containing the names of all of the kernels contained in the Metal library. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/kernelNames(fromMetalLibraryData:)
func (kc _KernelClass) KernelNamesFromMetalLibraryData(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelNamesFromMetalLibraryData:"), data)
	return rv
}
// Load kernels from a Metal language string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/kernels(withMetalString:)
func (kc _KernelClass) KernelsWithMetalStringError(source string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelsWithMetalString:error:"), objc.String(source), error)
	return rv
}
// Creates and returns and array of objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/makeKernels(source:)
func (kc _KernelClass) KernelsWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelsWithString:"), objc.String(string))
	return rv
}
// Creates a new image using the kernel and specified arguments. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/apply(extent:roiCallback:arguments:)
func (k_ Kernel) ApplyWithExtentRoiCallbackArguments(extent unsafe.Pointer, callback unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("applyWithExtent:roiCallback:arguments:"), extent, callback, args)
	return rv
}
// Sets the selector Core Image uses to query the region of interest for image processing with the kernel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/setROISelector(_:)
func (k_ Kernel) SetROISelector(method objc.SEL) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setROISelector:"), method)
}

