// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	EdgeMode() ImageEdgeMode get set /* not a class type */
	SetEdgeMode(value ImageEdgeMode get set /* not a class type */)
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	Offset() Offset get set /* not a class type */
	SetOffset(value Offset get set /* not a class type */)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer)
	SourceRegion()
	SourceRegionForDestinationSize(destinationSize metal.IMTLSize) MPSRegion
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage)
	EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, texture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool


}





// Alloc allocates a new instance without initialization.
func (uc _UnaryImageKernelClass) Alloc() UnaryImageKernel {
	rv := objc.Send[UnaryImageKernel](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866329-initwithcoder
func NewUnaryImageKernelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) UnaryImageKernel {
	instance := getUnaryImageKernelClass().Alloc()
	rv := objc.Send[UnaryImageKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice
func NewUnaryImageKernelWithDevice(device unsafe.Pointer) UnaryImageKernel {
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


// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618741-encodetocommandbuffer
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:"), commandBuffer, sourceTexture, destinationTexture)
}


// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618754-sourceregion
func (u_ UnaryImageKernel) SourceRegion() {
	objc.Send[objc.ID](u_.ID, objc.Sel("sourceRegion"))
}


// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618754-sourceregionfordestinationsize
func (u_ UnaryImageKernel) SourceRegionForDestinationSize(destinationSize metal.IMTLSize) MPSRegion {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866328-encodetocommandbuffer
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}


// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/encode(commandBuffer:inPlaceTexture:fallbackCopyAllocator:)
func (u_ UnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, texture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator:"), commandBuffer, texture, copyAllocator)
	return rv
}







// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode
func (u_ UnaryImageKernel) EdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("edgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode
func (u_ UnaryImageKernel) SetEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEdgeMode:"), value)
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect
func (u_ UnaryImageKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("clipRect"))
	return rv
}


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect
func (u_ UnaryImageKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRect:"), value)
}


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset
func (u_ UnaryImageKernel) Offset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("offset"))
	return rv
}


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset
func (u_ UnaryImageKernel) SetOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOffset:"), value)
}







