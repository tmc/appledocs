// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BinaryImageKernel] class.
var (
	BinaryImageKernelClass     _BinaryImageKernelClass
	BinaryImageKernelClassOnce sync.Once
)

func getBinaryImageKernelClass() _BinaryImageKernelClass {
	BinaryImageKernelClassOnce.Do(func() {
		BinaryImageKernelClass = _BinaryImageKernelClass{objc.GetClass("MPSBinaryImageKernel")}
	})
	return BinaryImageKernelClass
}

type _BinaryImageKernelClass struct {
	class objc.Class
}

// An interface definition for the [BinaryImageKernel] class.
type IBinaryImageKernel interface {
	IKernel
	// properties:
	ClipRect() objc.IObject /* cross-framework: MTLRegion */
	SetClipRect(value objc.IObject /* cross-framework: MTLRegion */)
	PrimaryEdgeMode() ImageEdgeMode
	SetPrimaryEdgeMode(value ImageEdgeMode)
	PrimaryOffset() MPSOffset /* not a class type */
	SetPrimaryOffset(value MPSOffset /* not a class type */)
	SecondaryEdgeMode() ImageEdgeMode
	SetSecondaryEdgeMode(value ImageEdgeMode)
	SecondaryOffset() MPSOffset /* not a class type */
	SetSecondaryOffset(value MPSOffset /* not a class type */)
	// methods:
	EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer objectivec.IObject, primaryTexture objectivec.IObject, inPlaceSecondaryTexture objectivec.IObject, copyAllocator CopyAllocator /* not a class type */) bool
	SecondarySourceRegionForDestinationSize(destinationSize Size /* not a class type */) MPSRegion /* not a class type */
}

// A kernel that consumes two textures and produces one texture.
//
// defines shared behavior for most image processing kernels (filters) such as edging modes, clipping, and tiling support for image operations that consume two source textures. It is not meant to be used directly, but provides API abstraction and in some cases may allow some level of polymorphic manipulation of image kernel objects.


// A kernel that consumes two textures and produces one texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel
type BinaryImageKernel struct {
	Kernel
}

// BinaryImageKernelFrom constructs a [BinaryImageKernel] from an unsafe.Pointer.
//
// A kernel that consumes two textures and produces one texture.
func BinaryImageKernelFrom(ptr unsafe.Pointer) BinaryImageKernel {
	return BinaryImageKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BinaryImageKernelClass) Alloc() BinaryImageKernel {
	rv := objc.Send[BinaryImageKernel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BinaryImageKernelClass) New() BinaryImageKernel {
	rv := objc.Send[BinaryImageKernel](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryImageKernel) Init() BinaryImageKernel {
	rv := objc.Send[BinaryImageKernel](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryImageKernel) Autorelease() BinaryImageKernel {
	rv := objc.Send[BinaryImageKernel](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryImageKernel creates a new BinaryImageKernel instance.
func NewBinaryImageKernel() BinaryImageKernel {
	return getBinaryImageKernelClass().New()
}



// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/encode(commandBuffer:primaryTexture:inPlaceSecondaryTexture:fallbackCopyAllocator:)
func (b_ BinaryImageKernel) EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer objectivec.IObject, primaryTexture objectivec.IObject, inPlaceSecondaryTexture objectivec.IObject, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("encodeToCommandBuffer:primaryTexture:inPlaceSecondaryTexture:fallbackCopyAllocator:"), commandBuffer, primaryTexture, inPlaceSecondaryTexture, copyAllocator)
	return rv
}


// Determines the region of the secondary source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/secondarySourceRegion(forDestinationSize:)
func (b_ BinaryImageKernel) SecondarySourceRegionForDestinationSize(destinationSize Size /* not a class type */) MPSRegion /* not a class type */ {
	rv := objc.Send[Region](b_.ID, objc.Sel("secondarySourceRegionForDestinationSize:"), destinationSize)
	return rv
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/cliprect
func (b_ BinaryImageKernel) ClipRect() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](b_.ID, objc.Sel("clipRect"))
	return rv
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/cliprect
func (b_ BinaryImageKernel) SetClipRect(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setClipRect:"), value)
}


// The edge mode to use when texture reads stray off the edge of the primary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/primaryedgemode
func (b_ BinaryImageKernel) PrimaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](b_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of the primary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/primaryedgemode
func (b_ BinaryImageKernel) SetPrimaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}


// The position of the destination clip rectangle origin relative to the primary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/primaryoffset
func (b_ BinaryImageKernel) PrimaryOffset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](b_.ID, objc.Sel("primaryOffset"))
	return rv
}


// The position of the destination clip rectangle origin relative to the primary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/primaryoffset
func (b_ BinaryImageKernel) SetPrimaryOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryOffset:"), value)
}


// The edge mode to use when texture reads stray off the edge of the secondary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/secondaryedgemode
func (b_ BinaryImageKernel) SecondaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](b_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of the secondary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/secondaryedgemode
func (b_ BinaryImageKernel) SetSecondaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}


// The position of the destination clip rectangle origin relative to the secondary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/secondaryoffset
func (b_ BinaryImageKernel) SecondaryOffset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](b_.ID, objc.Sel("secondaryOffset"))
	return rv
}


// The position of the destination clip rectangle origin relative to the secondary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/secondaryoffset
func (b_ BinaryImageKernel) SetSecondaryOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryOffset:"), value)
}



