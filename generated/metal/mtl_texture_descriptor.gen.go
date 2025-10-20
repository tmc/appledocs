// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextureDescriptor] class.
var (
	TextureDescriptorClass     _TextureDescriptorClass
	TextureDescriptorClassOnce sync.Once
)

func getTextureDescriptorClass() _TextureDescriptorClass {
	TextureDescriptorClassOnce.Do(func() {
		TextureDescriptorClass = _TextureDescriptorClass{objc.GetClass("MTLTextureDescriptor")}
	})
	return TextureDescriptorClass
}

type _TextureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [TextureDescriptor] class.
type ITextureDescriptor interface {
	objectivec.IObject
}

// An instance that you use to configure new Metal texture instances.
//
// To create a new texture, first create an instance and set its property values. Then, call either the or method of an instance, or the method of an instance. When you create a texture, Metal copies property values from the descriptor into the new texture. You can reuse an instance, modifying its property values as needed, to create more instances, without affecting any textures you already created.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor
type TextureDescriptor struct {
	objectivec.Object
}

// TextureDescriptorFrom constructs a [TextureDescriptor] from an unsafe.Pointer.
//
// An instance that you use to configure new Metal texture instances.
func TextureDescriptorFrom(ptr unsafe.Pointer) TextureDescriptor {
	return TextureDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextureDescriptorClass) Alloc() TextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextureDescriptorClass) New() TextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextureDescriptor) Init() TextureDescriptor {
	rv := objc.Send[TextureDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextureDescriptor) Autorelease() TextureDescriptor {
	rv := objc.Send[TextureDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextureDescriptor creates a new TextureDescriptor instance.
func NewTextureDescriptor() TextureDescriptor {
	return getTextureDescriptorClass().New()
}


// Creates a texture descriptor object for a 2D texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/texture2DDescriptor(pixelFormat:width:height:mipmapped:)
func (tc _TextureDescriptorClass) Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(pixelFormat unsafe.Pointer, width uint, height uint, mipmapped bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("texture2DDescriptorWithPixelFormat:width:height:mipmapped:"), pixelFormat, width, height, mipmapped)
	return rv
}

// Creates a texture descriptor object for a texture buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureBufferDescriptor(with:width:resourceOptions:usage:)
func (tc _TextureDescriptorClass) TextureBufferDescriptorWithPixelFormatWidthResourceOptionsUsage(pixelFormat unsafe.Pointer, width uint, resourceOptions unsafe.Pointer, usage unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textureBufferDescriptorWithPixelFormat:width:resourceOptions:usage:"), pixelFormat, width, resourceOptions, usage)
	return rv
}

// Creates a texture descriptor object for a cube texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureCubeDescriptor(pixelFormat:size:mipmapped:)
func (tc _TextureDescriptorClass) TextureCubeDescriptorWithPixelFormatSizeMipmapped(pixelFormat unsafe.Pointer, size uint, mipmapped bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textureCubeDescriptorWithPixelFormat:size:mipmapped:"), pixelFormat, size, mipmapped)
	return rv
}

// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/allowGPUOptimizedContents
func (t_ TextureDescriptor) AllowGPUOptimizedContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
}


// SetAllowGPUOptimizedContents sets the value of the allowGPUOptimizedContents property.
// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/allowGPUOptimizedContents
func (t_ TextureDescriptor) SetAllowGPUOptimizedContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowGPUOptimizedContents:"), value)
}
// The number of array elements for this texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/arrayLength
func (t_ TextureDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("arrayLength"))
	return rv
}


// SetArrayLength sets the value of the arrayLength property.
// The number of array elements for this texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/arrayLength
func (t_ TextureDescriptor) SetArrayLength(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArrayLength:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/compressionType
func (t_ TextureDescriptor) CompressionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("compressionType"))
	return rv
}


// SetCompressionType sets the value of the compressionType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/compressionType
func (t_ TextureDescriptor) SetCompressionType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompressionType:"), value)
}
// The CPU cache mode used for the CPU mapping of the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/cpuCacheMode
func (t_ TextureDescriptor) CpuCacheMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}


// SetCpuCacheMode sets the value of the cpuCacheMode property.
// The CPU cache mode used for the CPU mapping of the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/cpuCacheMode
func (t_ TextureDescriptor) SetCpuCacheMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}
// The depth of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/depth
func (t_ TextureDescriptor) Depth() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("depth"))
	return rv
}


// SetDepth sets the value of the depth property.
// The depth of the texture image for the base level mipmap, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/depth
func (t_ TextureDescriptor) SetDepth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDepth:"), value)
}
// The texture’s hazard tracking mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/hazardTrackingMode
func (t_ TextureDescriptor) HazardTrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}


// SetHazardTrackingMode sets the value of the hazardTrackingMode property.
// The texture’s hazard tracking mode.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/hazardTrackingMode
func (t_ TextureDescriptor) SetHazardTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}
// The height of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/height
func (t_ TextureDescriptor) Height() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("height"))
	return rv
}


// SetHeight sets the value of the height property.
// The height of the texture image for the base level mipmap, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/height
func (t_ TextureDescriptor) SetHeight(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeight:"), value)
}
// The number of mipmap levels for this texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/mipmapLevelCount
func (t_ TextureDescriptor) MipmapLevelCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("mipmapLevelCount"))
	return rv
}


// SetMipmapLevelCount sets the value of the mipmapLevelCount property.
// The number of mipmap levels for this texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/mipmapLevelCount
func (t_ TextureDescriptor) SetMipmapLevelCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMipmapLevelCount:"), value)
}
// The size and bit layout of all pixels in the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/pixelFormat
func (t_ TextureDescriptor) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// The size and bit layout of all pixels in the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/pixelFormat
func (t_ TextureDescriptor) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}
// Determines the page size for a placement sparse texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/placementSparsePageSize
func (t_ TextureDescriptor) PlacementSparsePageSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("placementSparsePageSize"))
	return rv
}


// SetPlacementSparsePageSize sets the value of the placementSparsePageSize property.
// Determines the page size for a placement sparse texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/placementSparsePageSize
func (t_ TextureDescriptor) SetPlacementSparsePageSize(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlacementSparsePageSize:"), value)
}
// The behavior of a new memory allocation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/resourceOptions
func (t_ TextureDescriptor) ResourceOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resourceOptions"))
	return rv
}


// SetResourceOptions sets the value of the resourceOptions property.
// The behavior of a new memory allocation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/resourceOptions
func (t_ TextureDescriptor) SetResourceOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}
// The number of samples in each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SampleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("sampleCount"))
	return rv
}


// SetSampleCount sets the value of the sampleCount property.
// The number of samples in each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSampleCount:"), value)
}
// The location and access permissions of the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/storageMode
func (t_ TextureDescriptor) StorageMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("storageMode"))
	return rv
}


// SetStorageMode sets the value of the storageMode property.
// The location and access permissions of the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/storageMode
func (t_ TextureDescriptor) SetStorageMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}
// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/swizzle
func (t_ TextureDescriptor) Swizzle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("swizzle"))
	return rv
}


// SetSwizzle sets the value of the swizzle property.
// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/swizzle
func (t_ TextureDescriptor) SetSwizzle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSwizzle:"), value)
}
// The dimension and arrangement of texture image data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureType
func (t_ TextureDescriptor) TextureType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textureType"))
	return rv
}


// SetTextureType sets the value of the textureType property.
// The dimension and arrangement of texture image data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureType
func (t_ TextureDescriptor) SetTextureType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}
// Options that determine how you can use the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) Usage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("usage"))
	return rv
}


// SetUsage sets the value of the usage property.
// Options that determine how you can use the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) SetUsage(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}
// The width of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/width
func (t_ TextureDescriptor) Width() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
// The width of the texture image for the base level mipmap, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/width
func (t_ TextureDescriptor) SetWidth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}


