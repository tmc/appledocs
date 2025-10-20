// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [BlendKernel] class.
var (
	blendKernelClass     _BlendKernelClass
	blendKernelClassOnce sync.Once
)

func getBlendKernelClass() _BlendKernelClass {
	blendKernelClassOnce.Do(func() {
		blendKernelClass = _BlendKernelClass{objc.GetClass("CIBlendKernel")}
	})
	return blendKernelClass
}

type _BlendKernelClass struct {
	class objc.Class
}

// An interface definition for the [BlendKernel] class.
type IBlendKernel interface {
	IColorKernel
	ApplyWithForegroundBackground(foreground unsafe.Pointer, background unsafe.Pointer) unsafe.Pointer
	ApplyWithForegroundBackgroundColorSpace(foreground unsafe.Pointer, background unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) unsafe.Pointer
}

// A GPU-based image-processing routine that is optimized for blending two images.
//
// The blend kernel function has the following characteristics: It has two arguments of type (Core Image Kernel Language) or (Metal Shading Language), representing the foreground and background images. Its return type is (Core Image Kernel Language) or (Metal Shading Language); that is, it returns a pixel color for the output image. A blend kernel routine receives as input single-pixel colors (one sampled from each input image) and computes a final pixel color (output using the return keyword). For example, the Metal Shading Language source below implements a filter that returns the average of its two input images. Generally, the extent of the output image is the union of the extents of the foreground and background images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel
type BlendKernel struct {
	ColorKernel
}

// BlendKernelFrom constructs a [BlendKernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine that is optimized for blending two images.
func BlendKernelFrom(ptr unsafe.Pointer) BlendKernel {
	return BlendKernel{
		ColorKernel: ColorKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BlendKernelClass) Alloc() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BlendKernelClass) New() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlendKernel) Init() BlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlendKernel) Autorelease() BlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlendKernel creates a new BlendKernel instance.
func NewBlendKernel() BlendKernel {
	return getBlendKernelClass().New()
}


// Creates a custom blend kernel from a program string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/init(source:)
func NewBlendKernelWithString(string string) BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(getBlendKernelClass().class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}


// Creates a custom blend kernel from a program string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/init(source:)
func (bc _BlendKernelClass) KernelWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}

// Creates a new image using the blend kernel and specified foreground and background images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:)
func (b_ BlendKernel) ApplyWithForegroundBackground(foreground unsafe.Pointer, background unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("applyWithForeground:background:"), foreground, background)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:colorSpace:)
func (b_ BlendKernel) ApplyWithForegroundBackgroundColorSpace(foreground unsafe.Pointer, background unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("applyWithForeground:background:colorSpace:"), foreground, background, colorSpace)
	return rv
}


