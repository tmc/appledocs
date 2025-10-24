// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTextureDescriptor */


/* debug [class_header]: Header for MTLTextureDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextureDescriptor */
// An interface definition for the [TextureDescriptor] class.
type ITextureDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextureDescriptor */
	// properties:
	AllowGPUOptimizedContents() bool
	SetAllowGPUOptimizedContents(value bool)
	ArrayLength() uint
	SetArrayLength(value uint)
	CompressionType() TextureCompressionType
	SetCompressionType(value TextureCompressionType)
	CpuCacheMode() CPUCacheMode
	SetCpuCacheMode(value CPUCacheMode)
	Depth() uint
	SetDepth(value uint)
	HazardTrackingMode() HazardTrackingMode
	SetHazardTrackingMode(value HazardTrackingMode)
	Height() uint
	SetHeight(value uint)
	MipmapLevelCount() uint
	SetMipmapLevelCount(value uint)
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
	PlacementSparsePageSize() SparsePageSize
	SetPlacementSparsePageSize(value SparsePageSize)
	ResourceOptions() ResourceOptions
	SetResourceOptions(value ResourceOptions)
	SampleCount() uint
	SetSampleCount(value uint)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	Swizzle() objc.IObject /* cross-framework: MTLTextureSwizzleChannels */
	SetSwizzle(value objc.IObject /* cross-framework: MTLTextureSwizzleChannels */)
	TextureType() TextureType
	SetTextureType(value TextureType)
	Usage() TextureUsage
	SetUsage(value TextureUsage)
	Width() uint
	SetWidth(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextureDescriptor */
// Alloc allocates a new instance without initialization.
func (tc _TextureDescriptorClass) Alloc() TextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextureDescriptor */
// An instance that you use to configure new Metal texture instances.
//
// To create a new texture, first create an instance and set its property values. Then, call either the or method of an instance, or the method of an instance. When you create a texture, Metal copies property values from the descriptor into the new texture. You can reuse an instance, modifying its property values as needed, to create more instances, without affecting any textures you already created.


// An instance that you use to configure new Metal texture instances.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextureDescriptor */

// Creates a texture descriptor object for a 2D texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/texture2DDescriptor(pixelFormat:width:height:mipmapped:)
func (tc _TextureDescriptorClass) Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(pixelFormat PixelFormat, width uint, height uint, mipmapped bool) ITextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("texture2DDescriptorWithPixelFormat:width:height:mipmapped:"), pixelFormat, width, height, mipmapped)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Texture2DDescriptorWithPixelFormatWidthHeightMipmapped) */


// Creates a texture descriptor object for a texture buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureBufferDescriptor(with:width:resourceOptions:usage:)
func (tc _TextureDescriptorClass) TextureBufferDescriptorWithPixelFormatWidthResourceOptionsUsage(pixelFormat PixelFormat, width uint, resourceOptions ResourceOptions, usage TextureUsage) ITextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("textureBufferDescriptorWithPixelFormat:width:resourceOptions:usage:"), pixelFormat, width, resourceOptions, usage)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextureBufferDescriptorWithPixelFormatWidthResourceOptionsUsage) */


// Creates a texture descriptor object for a cube texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureCubeDescriptor(pixelFormat:size:mipmapped:)
func (tc _TextureDescriptorClass) TextureCubeDescriptorWithPixelFormatSizeMipmapped(pixelFormat PixelFormat, size uint, mipmapped bool) ITextureDescriptor {
	rv := objc.Send[TextureDescriptor](objc.ID(tc.class), objc.Sel("textureCubeDescriptorWithPixelFormat:size:mipmapped:"), pixelFormat, size, mipmapped)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextureCubeDescriptorWithPixelFormatSizeMipmapped) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextureDescriptor */

// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/allowGPUOptimizedContents
func (t_ TextureDescriptor) AllowGPUOptimizedContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
}/* debug [instance_properties/getter]: allowGPUOptimizedContents */


// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/allowGPUOptimizedContents
func (t_ TextureDescriptor) SetAllowGPUOptimizedContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowGPUOptimizedContents:"), value)
}/* debug [instance_properties/setter]: allowGPUOptimizedContents */


// The number of array elements for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/arrayLength
func (t_ TextureDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("arrayLength"))
	return rv
}/* debug [instance_properties/getter]: arrayLength */


// The number of array elements for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/arrayLength
func (t_ TextureDescriptor) SetArrayLength(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArrayLength:"), value)
}/* debug [instance_properties/setter]: arrayLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/compressionType
func (t_ TextureDescriptor) CompressionType() TextureCompressionType {
	rv := objc.Send[TextureCompressionType](t_.ID, objc.Sel("compressionType"))
	return rv
}/* debug [instance_properties/getter]: compressionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/compressionType
func (t_ TextureDescriptor) SetCompressionType(value TextureCompressionType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompressionType:"), value)
}/* debug [instance_properties/setter]: compressionType */


// The CPU cache mode used for the CPU mapping of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/cpuCacheMode
func (t_ TextureDescriptor) CpuCacheMode() CPUCacheMode {
	rv := objc.Send[CPUCacheMode](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}/* debug [instance_properties/getter]: cpuCacheMode */


// The CPU cache mode used for the CPU mapping of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/cpuCacheMode
func (t_ TextureDescriptor) SetCpuCacheMode(value CPUCacheMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}/* debug [instance_properties/setter]: cpuCacheMode */


// The depth of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/depth
func (t_ TextureDescriptor) Depth() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("depth"))
	return rv
}/* debug [instance_properties/getter]: depth */


// The depth of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/depth
func (t_ TextureDescriptor) SetDepth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDepth:"), value)
}/* debug [instance_properties/setter]: depth */


// The texture’s hazard tracking mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/hazardTrackingMode
func (t_ TextureDescriptor) HazardTrackingMode() HazardTrackingMode {
	rv := objc.Send[HazardTrackingMode](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}/* debug [instance_properties/getter]: hazardTrackingMode */


// The texture’s hazard tracking mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/hazardTrackingMode
func (t_ TextureDescriptor) SetHazardTrackingMode(value HazardTrackingMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}/* debug [instance_properties/setter]: hazardTrackingMode */


// The height of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/height
func (t_ TextureDescriptor) Height() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The height of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/height
func (t_ TextureDescriptor) SetHeight(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// The number of mipmap levels for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/mipmapLevelCount
func (t_ TextureDescriptor) MipmapLevelCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("mipmapLevelCount"))
	return rv
}/* debug [instance_properties/getter]: mipmapLevelCount */


// The number of mipmap levels for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/mipmapLevelCount
func (t_ TextureDescriptor) SetMipmapLevelCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMipmapLevelCount:"), value)
}/* debug [instance_properties/setter]: mipmapLevelCount */


// The size and bit layout of all pixels in the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/pixelFormat
func (t_ TextureDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](t_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The size and bit layout of all pixels in the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/pixelFormat
func (t_ TextureDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// Determines the page size for a placement sparse texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/placementSparsePageSize
func (t_ TextureDescriptor) PlacementSparsePageSize() SparsePageSize {
	rv := objc.Send[SparsePageSize](t_.ID, objc.Sel("placementSparsePageSize"))
	return rv
}/* debug [instance_properties/getter]: placementSparsePageSize */


// Determines the page size for a placement sparse texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/placementSparsePageSize
func (t_ TextureDescriptor) SetPlacementSparsePageSize(value SparsePageSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlacementSparsePageSize:"), value)
}/* debug [instance_properties/setter]: placementSparsePageSize */


// The behavior of a new memory allocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/resourceOptions
func (t_ TextureDescriptor) ResourceOptions() ResourceOptions {
	rv := objc.Send[ResourceOptions](t_.ID, objc.Sel("resourceOptions"))
	return rv
}/* debug [instance_properties/getter]: resourceOptions */


// The behavior of a new memory allocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/resourceOptions
func (t_ TextureDescriptor) SetResourceOptions(value ResourceOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}/* debug [instance_properties/setter]: resourceOptions */


// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SampleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSampleCount:"), value)
}/* debug [instance_properties/setter]: sampleCount */


// The location and access permissions of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/storageMode
func (t_ TextureDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](t_.ID, objc.Sel("storageMode"))
	return rv
}/* debug [instance_properties/getter]: storageMode */


// The location and access permissions of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/storageMode
func (t_ TextureDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}/* debug [instance_properties/setter]: storageMode */


// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/swizzle
func (t_ TextureDescriptor) Swizzle() objc.IObject /* cross-framework: MTLTextureSwizzleChannels */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("swizzle"))
	return rv
}/* debug [instance_properties/getter]: swizzle */


// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/swizzle
func (t_ TextureDescriptor) SetSwizzle(value objc.IObject /* cross-framework: MTLTextureSwizzleChannels */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSwizzle:"), value)
}/* debug [instance_properties/setter]: swizzle */


// The dimension and arrangement of texture image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureType
func (t_ TextureDescriptor) TextureType() TextureType {
	rv := objc.Send[TextureType](t_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */


// The dimension and arrangement of texture image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureType
func (t_ TextureDescriptor) SetTextureType(value TextureType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}/* debug [instance_properties/setter]: textureType */


// Options that determine how you can use the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) Usage() TextureUsage {
	rv := objc.Send[TextureUsage](t_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// Options that determine how you can use the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) SetUsage(value TextureUsage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */


// The width of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/width
func (t_ TextureDescriptor) Width() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The width of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/width
func (t_ TextureDescriptor) SetWidth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTextureDescriptor */



