// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ColorKernel] class.
var colorKernelClass = _ColorKernelClass{objc.GetClass("CIColorKernel")}

type _ColorKernelClass struct {
	class objc.Class
}

// An interface definition for the [ColorKernel] class.
type IColorKernel interface {
	IKernel
	ApplyWithExtentArguments(extent unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer
}

// A GPU-based image-processing routine that processes only the color information in images, used to create custom Core Image filters. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return colorKernelClass.New()
}


// Creates a color kernel object from the specified kernel source code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/init(source:)
func NewColorKernelWithString(string string) ColorKernel {
	rv := objc.Send[ColorKernel](objc.ID(colorKernelClass.class), objc.Sel("kernelWithString:"), string)
	rv.Autorelease()
	return rv
}


// Creates a color kernel object from the specified kernel source code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/init(source:)
func (cc _ColorKernelClass) KernelWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("kernelWithString:"), string)
	return rv
}
// Creates a new image using the kernel and specified arguments. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColorKernel/apply(extent:arguments:)
func (c_ ColorKernel) ApplyWithExtentArguments(extent unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("applyWithExtent:arguments:"), extent, args)
	return rv
}

