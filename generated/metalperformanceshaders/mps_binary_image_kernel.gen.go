// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSBinaryImageKernel */


/* debug [class_header]: Header for MPSBinaryImageKernel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BinaryImageKernel */
// An interface definition for the [BinaryImageKernel] class.
type IBinaryImageKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for BinaryImageKernel */
	// properties:
	SecondaryOffset() Offset get set /* not a class type */
	SetSecondaryOffset(value Offset get set /* not a class type */)
	PrimaryEdgeMode() ImageEdgeMode get set /* not a class type */
	SetPrimaryEdgeMode(value ImageEdgeMode get set /* not a class type */)
	SecondaryEdgeMode() ImageEdgeMode get set /* not a class type */
	SetSecondaryEdgeMode(value ImageEdgeMode get set /* not a class type */)
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	PrimaryOffset() Offset get set /* not a class type */
	SetPrimaryOffset(value Offset get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BinaryImageKernel */
	// methods:
	Encode()
	EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, inPlacePrimaryTexture unsafe.Pointer, secondaryTexture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool
	SecondarySourceRegion()
	SecondarySourceRegionForDestinationSize(destinationSize objc.IObject /* cross-framework: MTLSize */) objc.IObject /* cross-framework: MPSRegion */
	EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture(commandBuffer unsafe.Pointer, primaryTexture unsafe.Pointer, secondaryTexture unsafe.Pointer, destinationTexture unsafe.Pointer)
	EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, primaryTexture unsafe.Pointer, inPlaceSecondaryTexture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool
	PrimarySourceRegion()
	PrimarySourceRegionForDestinationSize(destinationSize objc.IObject /* cross-framework: MTLSize */) objc.IObject /* cross-framework: MPSRegion */
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationImage IImage)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BinaryImageKernel */
// Alloc allocates a new instance without initialization.
func (bc _BinaryImageKernelClass) Alloc() BinaryImageKernel {
	rv := objc.Send[BinaryImageKernel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BinaryImageKernel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BinaryImageKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866333-initwithcoder
func NewBinaryImageKernelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) BinaryImageKernel {
	instance := getBinaryImageKernelClass().Alloc()
	rv := objc.Send[BinaryImageKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBinaryImageKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866331-initwithdevice
func NewBinaryImageKernelWithDevice(device unsafe.Pointer) BinaryImageKernel {
	instance := getBinaryImageKernelClass().Alloc()
	rv := objc.Send[BinaryImageKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBinaryImageKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BinaryImageKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BinaryImageKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BinaryImageKernel */

// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618771-encode
func (b_ BinaryImageKernel) Encode() {
	objc.Send[objc.ID](b_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618771-encodetocommandbuffer
func (b_ BinaryImageKernel) EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, inPlacePrimaryTexture unsafe.Pointer, secondaryTexture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("encodeToCommandBuffer:inPlacePrimaryTexture:secondaryTexture:fallbackCopyAllocator:"), commandBuffer, inPlacePrimaryTexture, secondaryTexture, copyAllocator)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator */


// Determines the region of the secondary source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618838-secondarysourceregion
func (b_ BinaryImageKernel) SecondarySourceRegion() {
	objc.Send[objc.ID](b_.ID, objc.Sel("secondarySourceRegion"))
}/* debug [instance_methods/method]: SecondarySourceRegion */


// Determines the region of the secondary source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618838-secondarysourceregionfordestinat
func (b_ BinaryImageKernel) SecondarySourceRegionForDestinationSize(destinationSize objc.IObject /* cross-framework: MTLSize */) objc.IObject /* cross-framework: MPSRegion */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("secondarySourceRegionForDestinationSize:"), destinationSize)
	return rv
}/* debug [instance_methods/method]: SecondarySourceRegionForDestinationSize */


// Encodes a kernel into a command buffer, out-of-place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618871-encodetocommandbuffer
func (b_ BinaryImageKernel) EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture(commandBuffer unsafe.Pointer, primaryTexture unsafe.Pointer, secondaryTexture unsafe.Pointer, destinationTexture unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("encodeToCommandBuffer:primaryTexture:secondaryTexture:destinationTexture:"), commandBuffer, primaryTexture, secondaryTexture, destinationTexture)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryTextureSecondaryTextureDestinationTexture */


// This method attempts to apply a kernel in place on a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618890-encodetocommandbuffer
func (b_ BinaryImageKernel) EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator(commandBuffer unsafe.Pointer, primaryTexture unsafe.Pointer, inPlaceSecondaryTexture unsafe.Pointer, copyAllocator CopyAllocator /* not a class type */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("encodeToCommandBuffer:primaryTexture:inPlaceSecondaryTexture:fallbackCopyAllocator:"), commandBuffer, primaryTexture, inPlaceSecondaryTexture, copyAllocator)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator */


// Determines the region of the primary source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618900-primarysourceregion
func (b_ BinaryImageKernel) PrimarySourceRegion() {
	objc.Send[objc.ID](b_.ID, objc.Sel("primarySourceRegion"))
}/* debug [instance_methods/method]: PrimarySourceRegion */


// Determines the region of the primary source texture that will be read for an encode operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618900-primarysourceregionfordestinatio
func (b_ BinaryImageKernel) PrimarySourceRegionForDestinationSize(destinationSize objc.IObject /* cross-framework: MTLSize */) objc.IObject /* cross-framework: MPSRegion */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("primarySourceRegionForDestinationSize:"), destinationSize)
	return rv
}/* debug [instance_methods/method]: PrimarySourceRegionForDestinationSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866330-encodetocommandbuffer
func (b_ BinaryImageKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BinaryImageKernel */

// The position of the destination clip rectangle origin relative to the secondary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618755-secondaryoffset
func (b_ BinaryImageKernel) SecondaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("secondaryOffset"))
	return rv
}/* debug [instance_properties/getter]: secondaryOffset */


// The position of the destination clip rectangle origin relative to the secondary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618755-secondaryoffset
func (b_ BinaryImageKernel) SetSecondaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryOffset:"), value)
}/* debug [instance_properties/setter]: secondaryOffset */


// The edge mode to use when texture reads stray off the edge of the primary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618782-primaryedgemode
func (b_ BinaryImageKernel) PrimaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: primaryEdgeMode */


// The edge mode to use when texture reads stray off the edge of the primary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618782-primaryedgemode
func (b_ BinaryImageKernel) SetPrimaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: primaryEdgeMode */


// The edge mode to use when texture reads stray off the edge of the secondary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618848-secondaryedgemode
func (b_ BinaryImageKernel) SecondaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: secondaryEdgeMode */


// The edge mode to use when texture reads stray off the edge of the secondary source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618848-secondaryedgemode
func (b_ BinaryImageKernel) SetSecondaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: secondaryEdgeMode */


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618879-cliprect
func (b_ BinaryImageKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618879-cliprect
func (b_ BinaryImageKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// The position of the destination clip rectangle origin relative to the primary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618880-primaryoffset
func (b_ BinaryImageKernel) PrimaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("primaryOffset"))
	return rv
}/* debug [instance_properties/getter]: primaryOffset */


// The position of the destination clip rectangle origin relative to the primary source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618880-primaryoffset
func (b_ BinaryImageKernel) SetPrimaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryOffset:"), value)
}/* debug [instance_properties/setter]: primaryOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSBinaryImageKernel */


