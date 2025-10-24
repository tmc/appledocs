// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ClipRect() objc.IObject /* cross-framework: MTLRegion */
	SetClipRect(value objc.IObject /* cross-framework: MTLRegion */)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
	Offset() MPSOffset /* not a class type */
	SetOffset(value MPSOffset /* not a class type */)
	// methods:
	Encode()
	EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer objectivec.IObject, texture objectivec.IObject, copyAllocator CopyAllocator /* not a class type */) bool
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer objectivec.IObject, sourceImage IMPSImage, destinationImage IMPSImage)
	EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, destinationTexture objectivec.IObject)
	SourceRegionForDestinationSize(destinationSize Size /* not a class type */) MPSRegion /* not a class type */
}

// A kernel that consumes one texture and produces one texture.
//
// defines shared behavior for most image processing kernels (filters) such as edging modes, clipping, and tiling support for image operations that consumes a single source textures. It is not meant to be used directly, but provides API abstraction and in some cases may allow some level of polymorphic manipulation of image kernel objects.


// A kernel that consumes one texture and produces one texture.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewUnaryImageKernelWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) UnaryImageKernel {
	instance := getUnaryImageKernelClass().Alloc()
	rv := objc.Send[UnaryImageKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewUnaryImageKernelWithDevice(device objectivec.IObject) UnaryImageKernel {
	instance := getUnaryImageKernelClass().Alloc()
	rv := objc.Send[UnaryImageKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618741-encode
func (u_ UnaryImageKernel) Encode() {
	objc.Send[objc.ID](u_.ID, objc.Sel("encode"))
}


// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:inPlaceTexture:fallbackCopyAllocator:)
func (u_ UnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer objectivec.IObject, texture objectivec.IObject, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator:"), commandBuffer, texture, copyAllocator)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:sourceImage:destinationImage:)
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer objectivec.IObject, sourceImage IMPSImage, destinationImage IMPSImage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}


// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:sourceTexture:destinationTexture:)
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, destinationTexture objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:"), commandBuffer, sourceTexture, destinationTexture)
}


// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/sourceRegion(destinationSize:)
func (u_ UnaryImageKernel) SourceRegionForDestinationSize(destinationSize Size /* not a class type */) MPSRegion /* not a class type */ {
	rv := objc.Send[Region](u_.ID, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return rv
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/clipRect
func (u_ UnaryImageKernel) ClipRect() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](u_.ID, objc.Sel("clipRect"))
	return rv
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/clipRect
func (u_ UnaryImageKernel) SetClipRect(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRect:"), value)
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/edgeMode
func (u_ UnaryImageKernel) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](u_.ID, objc.Sel("edgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/edgeMode
func (u_ UnaryImageKernel) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEdgeMode:"), value)
}


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/offset
func (u_ UnaryImageKernel) Offset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](u_.ID, objc.Sel("offset"))
	return rv
}


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/offset
func (u_ UnaryImageKernel) SetOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOffset:"), value)
}


