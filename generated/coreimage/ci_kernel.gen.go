// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Kernel] class.
var (
	KernelClass     _KernelClass
	KernelClassOnce sync.Once
)

func getKernelClass() _KernelClass {
	KernelClassOnce.Do(func() {
		KernelClass = _KernelClass{objc.GetClass("CIKernel")}
	})
	return KernelClass
}

type _KernelClass struct {
	class objc.Class
}

// An interface definition for the [Kernel] class.
type IKernel interface {
	objectivec.IObject
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	// methods:
	ApplyWithExtentRoiCallbackArguments(extent objc.IObject /* cross-framework: Rect */, callback KernelROICallback /* not a class type */, args []objc.ID) IImage
	SetROISelector(method objc.SEL)
}

// A GPU-based image-processing routine used to create custom Core Image filters.
//
// The kernel language routine for a general-purpose filter kernel has the following characteristics: Its return type is (Core Image Kernel Language) or (Metal Shading Language); that is, it returns a pixel color for the output image. It may use zero or more input images. Each input image is represented by a parameter of type . A kernel routine typically produces its output by calculating source image coordinates (using the and functions or the function), samples from the source images (using the function), and computes a final pixel color (output using the keyword). For example, the Metal Shading Language source below implements a filter that passes through its input image unchanged. The equivalent code in Core Image Kernel Language is: The Core Image Kernel Language is a dialect of the OpenGL Shading Language. See and for more details.


// A GPU-based image-processing routine used to create custom Core Image filters.
//
// [Full Topic]
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



// Creates a single kernel object using a Metal Shading Language (MSL) kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
func NewKernelWithFunctionNameFromMetalLibraryDataError(name objc.IObject /* cross-framework: NSString */, data objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), name, data, error_)
	return rv
}


// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:outputPixelFormat:)
func NewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name objc.IObject /* cross-framework: NSString */, data objc.IObject /* cross-framework: NSData */, format objc.IObject /* cross-framework: Format */, error_ unsafe.Pointer) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, error_)
	return rv
}


// Creates a single kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(source:)
func NewKernelWithString(string_ objc.IObject /* cross-framework: NSString */) Kernel {
	rv := objc.Send[Kernel](objc.ID(getKernelClass().class), objc.Sel("kernelWithString:"), string_)
	return rv
}



// Creates a single kernel object using a Metal Shading Language (MSL) kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
func (kc _KernelClass) KernelWithFunctionNameFromMetalLibraryDataError(name objc.IObject /* cross-framework: NSString */, data objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), name, data, error_)
	return rv
}


// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:outputPixelFormat:)
func (kc _KernelClass) KernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name objc.IObject /* cross-framework: NSString */, data objc.IObject /* cross-framework: NSData */, format objc.IObject /* cross-framework: Format */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), name, data, format, error_)
	return rv
}


// Creates a single kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(source:)
func (kc _KernelClass) KernelWithString(string_ objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(kc.class), objc.Sel("kernelWithString:"), string_)
	return rv
}


// Return an array of strings containing the names of all of the kernels contained in the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/kernelNames(fromMetalLibraryData:)
func (kc _KernelClass) KernelNamesFromMetalLibraryData(data objc.IObject /* cross-framework: NSData */) []string {
	rv := objc.Send[[]string](objc.ID(kc.class), objc.Sel("kernelNamesFromMetalLibraryData:"), data)
	return rv
}


// Load kernels from a Metal language string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/kernels(withMetalString:)
func (kc _KernelClass) KernelsWithMetalStringError(source objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) []IKernel {
	rv := objc.Send[[]Kernel](objc.ID(kc.class), objc.Sel("kernelsWithMetalString:error:"), source, error_)
	return rv
}


// Creates and returns and array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/makeKernels(source:)
func (kc _KernelClass) KernelsWithString(string_ objc.IObject /* cross-framework: NSString */) []IKernel {
	rv := objc.Send[[]Kernel](objc.ID(kc.class), objc.Sel("kernelsWithString:"), string_)
	return rv
}


// Creates a new image using the kernel and specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/apply(extent:roiCallback:arguments:)
func (k_ Kernel) ApplyWithExtentRoiCallbackArguments(extent objc.IObject /* cross-framework: Rect */, callback KernelROICallback /* not a class type */, args []objc.ID) IImage {
	rv := objc.Send[Image](k_.ID, objc.Sel("applyWithExtent:roiCallback:arguments:"), extent, callback, args)
	return rv
}


// Sets the selector Core Image uses to query the region of interest for image processing with the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/setROISelector(_:)
func (k_ Kernel) SetROISelector(method objc.SEL) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setROISelector:"), method)
}


// The name of the kernel routine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIKernel/name
func (k_ Kernel) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](k_.ID, objc.Sel("name"))
	return rv
}


