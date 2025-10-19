// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WarpKernel] class.
var (
	warpKernelClass     _WarpKernelClass
	warpKernelClassOnce sync.Once
)

func getWarpKernelClass() _WarpKernelClass {
	warpKernelClassOnce.Do(func() {
		warpKernelClass = _WarpKernelClass{objc.GetClass("CIWarpKernel")}
	})
	return warpKernelClass
}

type _WarpKernelClass struct {
	class objc.Class
}

// An interface definition for the [WarpKernel] class.
type IWarpKernel interface {
	IKernel
	ApplyWithExtentRoiCallbackInputImageArguments(extent unsafe.Pointer, callback unsafe.Pointer, image unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer
}

// A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel
type WarpKernel struct {
	Kernel
}

// WarpKernelFrom constructs a [WarpKernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
func WarpKernelFrom(ptr unsafe.Pointer) WarpKernel {
	return WarpKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (wc _WarpKernelClass) Alloc() WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WarpKernelClass) New() WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WarpKernel) Init() WarpKernel {
	rv := objc.Send[WarpKernel](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WarpKernel) Autorelease() WarpKernel {
	rv := objc.Send[WarpKernel](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWarpKernel creates a new WarpKernel instance.
func NewWarpKernel() WarpKernel {
	return getWarpKernelClass().New()
}


// Creates a warp kernel object from the specified kernel source code.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/init(source:)
func NewWarpKernelWithString(string string) WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(getWarpKernelClass().class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}


// Creates a warp kernel object from the specified kernel source code.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/init(source:)
func (wc _WarpKernelClass) KernelWithString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("kernelWithString:"), objc.String(string))
	return rv
}
// Creates a new image using the kernel and the specified input image and arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/apply(extent:roiCallback:image:arguments:)
func (w_ WarpKernel) ApplyWithExtentRoiCallbackInputImageArguments(extent unsafe.Pointer, callback unsafe.Pointer, image unsafe.Pointer, args unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("applyWithExtent:roiCallback:inputImage:arguments:"), extent, callback, image, args)
	return rv
}

