// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ColorKernel] class.
var (
	ColorKernelClass     _ColorKernelClass
	ColorKernelClassOnce sync.Once
)

func getColorKernelClass() _ColorKernelClass {
	ColorKernelClassOnce.Do(func() {
		ColorKernelClass = _ColorKernelClass{objc.GetClass("CIColorKernel")}
	})
	return ColorKernelClass
}

type _ColorKernelClass struct {
	class objc.Class
}

// An interface definition for the [ColorKernel] class.
type IColorKernel interface {
	IKernel
	ApplyWithExtentArguments(extent coregraphics.CGRect, args []objc.ID) IImage
}

// A GPU-based image-processing routine that processes only the color information in images, used to create custom Core Image filters.
//
// The kernel language routine for a color kernel has the following characteristics: Its return type is (Core Image Kernel Language) or (Metal Shading Language); that is, it returns a pixel color for the output image. It may use zero or more input images. Each input image is represented by a parameter of type (Core Image Kernel Language) or (Metal Shading Language), which can be treated as a single pixel color of type (Core Image Kernel Language) or (Metal Shading Language);. A color kernel routine receives as input single-pixel colors (one sampled from each input image) and computes a final pixel color (output using the keyword). For example, the Metal Shading Language source below implements a filter that passes through its input image unchanged. The equivalent code in Core Image Kernel Language is: The Core Image Kernel Language is a dialect of the OpenGL Shading Language. See and for more details.


// A GPU-based image-processing routine that processes only the color information in images, used to create custom Core Image filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel
type ColorKernel struct {
	Kernel
}

// ColorKernelFrom constructs a [ColorKernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine that processes only the color information in images, used to create custom Core Image filters.
func ColorKernelFrom(ptr unsafe.Pointer) ColorKernel {
	return ColorKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorKernelClass) Alloc() ColorKernel {
	rv := objc.Send[ColorKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorKernelClass) New() ColorKernel {
	rv := objc.Send[ColorKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorKernel) Init() ColorKernel {
	rv := objc.Send[ColorKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorKernel) Autorelease() ColorKernel {
	rv := objc.Send[ColorKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorKernel creates a new ColorKernel instance.
func NewColorKernel() ColorKernel {
	return getColorKernelClass().New()
}



// Creates a color kernel object from the specified kernel source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/init(source:)
func NewColorKernelWithString(string_ string) ColorKernel {
	rv := objc.Send[ColorKernel](objc.ID(getColorKernelClass().class), objc.Sel("kernelWithString:"), objc.String(string_))
	return rv
}



// Creates a color kernel object from the specified kernel source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/init(source:)
func (cc _ColorKernelClass) KernelWithString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("kernelWithString:"), objc.String(string_))
	return rv
}


// Creates a new image using the kernel and specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/apply(extent:arguments:)
func (c_ ColorKernel) ApplyWithExtentArguments(extent coregraphics.CGRect, args []objc.ID) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("applyWithExtent:arguments:"), extent, args)
	return rv
}


