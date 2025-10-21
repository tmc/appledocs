// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnaryImageKernel] class.
var (
	UnaryImageKernelClass     _UnaryImageKernelClass
	UnaryImageKernelClassOnce sync.Once
)

func getUnaryImageKernelClass() _UnaryImageKernelClass {
	UnaryImageKernelClassOnce.Do(func() {
		UnaryImageKernelClass = _UnaryImageKernelClass{objc.GetClass("MPSUnaryImageKernel")}
	})
	return UnaryImageKernelClass
}

type _UnaryImageKernelClass struct {
	class objc.Class
}

// An interface definition for the [UnaryImageKernel] class.
type IUnaryImageKernel interface {
	IKernel
	Encode()
	SourceRegionForDestinationSize(destinationSize unsafe.Pointer) unsafe.Pointer
}

// A kernel that consumes one texture and produces one texture.
//
// defines shared behavior for most image processing kernels (filters) such as edging modes, clipping, and tiling support for image operations that consumes a single source textures. It is not meant to be used directly, but provides API abstraction and in some cases may allow some level of polymorphic manipulation of image kernel objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel
type UnaryImageKernel struct {
	Kernel
}

// UnaryImageKernelFrom constructs a [UnaryImageKernel] from an unsafe.Pointer.
//
// A kernel that consumes one texture and produces one texture.
func UnaryImageKernelFrom(ptr unsafe.Pointer) UnaryImageKernel {
	return UnaryImageKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnaryImageKernelClass) Alloc() UnaryImageKernel {
	rv := objc.Send[UnaryImageKernel](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnaryImageKernelClass) New() UnaryImageKernel {
	rv := objc.Send[UnaryImageKernel](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnaryImageKernel) Init() UnaryImageKernel {
	rv := objc.Send[UnaryImageKernel](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnaryImageKernel) Autorelease() UnaryImageKernel {
	rv := objc.Send[UnaryImageKernel](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnaryImageKernel creates a new UnaryImageKernel instance.
func NewUnaryImageKernel() UnaryImageKernel {
	return getUnaryImageKernelClass().New()
}


// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618741-encode
func (u_ UnaryImageKernel) Encode() {
	objc.Send[objc.ID](u_.ID, objc.Sel("encode"))
}

// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/sourceRegion(destinationSize:)
func (u_ UnaryImageKernel) SourceRegionForDestinationSize(destinationSize unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return rv
}

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/cliprect
func (u_ UnaryImageKernel) ClipRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("clipRect"))
	return rv
}


// SetClipRect sets the value of the clipRect property.
// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/cliprect
func (u_ UnaryImageKernel) SetClipRect(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRect:"), value)
}

// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/offset
func (u_ UnaryImageKernel) Offset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
// The position of the destination clip rectangle origin relative to the source buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/offset
func (u_ UnaryImageKernel) SetOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOffset:"), value)
}

// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (u_ UnaryImageKernel) EdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("edgeMode"))
	return rv
}


// SetEdgeMode sets the value of the edgeMode property.
// The edge mode to use when texture reads stray off the edge of an image.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (u_ UnaryImageKernel) SetEdgeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEdgeMode:"), value)
}




