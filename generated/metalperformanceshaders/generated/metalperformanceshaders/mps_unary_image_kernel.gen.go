// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSUnaryImageKernel */


/* debug [class_header]: Header for MPSUnaryImageKernel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnaryImageKernel */
// An interface definition for the [UnaryImageKernel] class.
type IUnaryImageKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for UnaryImageKernel */
	// properties:
	EdgeMode() ImageEdgeMode get set /* not a class type */
	SetEdgeMode(value ImageEdgeMode get set /* not a class type */)
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	Offset() Offset get set /* not a class type */
	SetOffset(value Offset get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnaryImageKernel */
	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer)
	SourceRegion()
	SourceRegionForDestinationSize(destinationSize Size /* not a class type */) objc.IObject /* cross-framework: MPSRegion */
	EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, texture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnaryImageKernel */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnaryImageKernel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnaryImageKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866329-initwithcoder
func NewUnaryImageKernelWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) UnaryImageKernel {
	instance := getUnaryImageKernelClass().Alloc()
	rv := objc.Send[UnaryImageKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUnaryImageKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice
func NewUnaryImageKernelWithDevice(device unsafe.Pointer) UnaryImageKernel {
	instance := getUnaryImageKernelClass().Alloc()
	rv := objc.Send[UnaryImageKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUnaryImageKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnaryImageKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnaryImageKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnaryImageKernel */

// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618741-encode
func (u_ UnaryImageKernel) Encode() {
	objc.Send[objc.ID](u_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// Encodes a kernel into a command buffer, out of place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618741-encodetocommandbuffer
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceTextureDestinationTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:"), commandBuffer, sourceTexture, destinationTexture)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTextureDestinationTexture */


// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618754-sourceregion
func (u_ UnaryImageKernel) SourceRegion() {
	objc.Send[objc.ID](u_.ID, objc.Sel("sourceRegion"))
}/* debug [instance_methods/method]: SourceRegion */


// Determines the region of the source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618754-sourceregionfordestinationsize
func (u_ UnaryImageKernel) SourceRegionForDestinationSize(destinationSize Size /* not a class type */) objc.IObject /* cross-framework: MPSRegion */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return rv
}/* debug [instance_methods/method]: SourceRegionForDestinationSize */


// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618873-encodetocommandbuffer
func (u_ UnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, texture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator:"), commandBuffer, texture, copyAllocator)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866328-encodetocommandbuffer
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnaryImageKernel */

// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode
func (u_ UnaryImageKernel) EdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode
func (u_ UnaryImageKernel) SetEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect
func (u_ UnaryImageKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect
func (u_ UnaryImageKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset
func (u_ UnaryImageKernel) Offset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset
func (u_ UnaryImageKernel) SetOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSUnaryImageKernel */


