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
	// properties:
	SampleCount() uint
	SetSampleCount(value uint)
	Usage() TextureUsage
	SetUsage(value TextureUsage)
	AllowGPUOptimizedContents() bool
	SetAllowGPUOptimizedContents(value bool)
	ArrayLength() int
	SetArrayLength(value int)
	CompressionType() TextureCompressionType /* not a class type */
	SetCompressionType(value TextureCompressionType /* not a class type */)
	CpuCacheMode() CPUCacheMode /* not a class type */
	SetCpuCacheMode(value CPUCacheMode /* not a class type */)
	Depth() int
	SetDepth(value int)
	HazardTrackingMode() HazardTrackingMode /* not a class type */
	SetHazardTrackingMode(value HazardTrackingMode /* not a class type */)
	Height() int
	SetHeight(value int)
	MipmapLevelCount() int
	SetMipmapLevelCount(value int)
	PixelFormat() PixelFormat /* not a class type */
	SetPixelFormat(value PixelFormat /* not a class type */)
	PlacementSparsePageSize() SparsePageSize /* not a class type */
	SetPlacementSparsePageSize(value SparsePageSize /* not a class type */)
	ResourceOptions() ResourceOptions /* not a class type */
	SetResourceOptions(value ResourceOptions /* not a class type */)
	StorageMode() StorageMode /* not a class type */
	SetStorageMode(value StorageMode /* not a class type */)
	Swizzle() TextureSwizzleChannels /* not a class type */
	SetSwizzle(value TextureSwizzleChannels /* not a class type */)
	TextureType() TextureType
	SetTextureType(value TextureType)
	Width() int
	SetWidth(value int)
	// methods:
}

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



// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SampleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("sampleCount"))
	return rv
}


// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t_ TextureDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSampleCount:"), value)
}


// Options that determine how you can use the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) Usage() TextureUsage {
	rv := objc.Send[TextureUsage](t_.ID, objc.Sel("usage"))
	return rv
}


// Options that determine how you can use the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t_ TextureDescriptor) SetUsage(value TextureUsage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsage:"), value)
}


// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/allowgpuoptimizedcontents
func (t_ TextureDescriptor) AllowGPUOptimizedContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
}


// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/allowgpuoptimizedcontents
func (t_ TextureDescriptor) SetAllowGPUOptimizedContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowGPUOptimizedContents:"), value)
}


// The number of array elements for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/arraylength
func (t_ TextureDescriptor) ArrayLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("arrayLength"))
	return rv
}


// The number of array elements for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/arraylength
func (t_ TextureDescriptor) SetArrayLength(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArrayLength:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/compressiontype
func (t_ TextureDescriptor) CompressionType() TextureCompressionType /* not a class type */ {
	rv := objc.Send[TextureCompressionType](t_.ID, objc.Sel("compressionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/compressiontype
func (t_ TextureDescriptor) SetCompressionType(value TextureCompressionType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompressionType:"), value)
}


// The CPU cache mode used for the CPU mapping of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/cpucachemode
func (t_ TextureDescriptor) CpuCacheMode() CPUCacheMode /* not a class type */ {
	rv := objc.Send[CPUCacheMode](t_.ID, objc.Sel("cpuCacheMode"))
	return rv
}


// The CPU cache mode used for the CPU mapping of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/cpucachemode
func (t_ TextureDescriptor) SetCpuCacheMode(value CPUCacheMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCpuCacheMode:"), value)
}


// The depth of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/depth
func (t_ TextureDescriptor) Depth() int {
	rv := objc.Send[int](t_.ID, objc.Sel("depth"))
	return rv
}


// The depth of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/depth
func (t_ TextureDescriptor) SetDepth(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDepth:"), value)
}


// The texture’s hazard tracking mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/hazardtrackingmode
func (t_ TextureDescriptor) HazardTrackingMode() HazardTrackingMode /* not a class type */ {
	rv := objc.Send[HazardTrackingMode](t_.ID, objc.Sel("hazardTrackingMode"))
	return rv
}


// The texture’s hazard tracking mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/hazardtrackingmode
func (t_ TextureDescriptor) SetHazardTrackingMode(value HazardTrackingMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHazardTrackingMode:"), value)
}


// The height of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/height
func (t_ TextureDescriptor) Height() int {
	rv := objc.Send[int](t_.ID, objc.Sel("height"))
	return rv
}


// The height of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/height
func (t_ TextureDescriptor) SetHeight(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeight:"), value)
}


// The number of mipmap levels for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/mipmaplevelcount
func (t_ TextureDescriptor) MipmapLevelCount() int {
	rv := objc.Send[int](t_.ID, objc.Sel("mipmapLevelCount"))
	return rv
}


// The number of mipmap levels for this texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/mipmaplevelcount
func (t_ TextureDescriptor) SetMipmapLevelCount(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMipmapLevelCount:"), value)
}


// The size and bit layout of all pixels in the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/pixelformat
func (t_ TextureDescriptor) PixelFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](t_.ID, objc.Sel("pixelFormat"))
	return rv
}


// The size and bit layout of all pixels in the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/pixelformat
func (t_ TextureDescriptor) SetPixelFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}


// Determines the page size for a placement sparse texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/placementsparsepagesize
func (t_ TextureDescriptor) PlacementSparsePageSize() SparsePageSize /* not a class type */ {
	rv := objc.Send[SparsePageSize](t_.ID, objc.Sel("placementSparsePageSize"))
	return rv
}


// Determines the page size for a placement sparse texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/placementsparsepagesize
func (t_ TextureDescriptor) SetPlacementSparsePageSize(value SparsePageSize /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlacementSparsePageSize:"), value)
}


// The behavior of a new memory allocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/resourceoptions
func (t_ TextureDescriptor) ResourceOptions() ResourceOptions /* not a class type */ {
	rv := objc.Send[ResourceOptions](t_.ID, objc.Sel("resourceOptions"))
	return rv
}


// The behavior of a new memory allocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/resourceoptions
func (t_ TextureDescriptor) SetResourceOptions(value ResourceOptions /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResourceOptions:"), value)
}


// The location and access permissions of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/storagemode
func (t_ TextureDescriptor) StorageMode() StorageMode /* not a class type */ {
	rv := objc.Send[StorageMode](t_.ID, objc.Sel("storageMode"))
	return rv
}


// The location and access permissions of the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/storagemode
func (t_ TextureDescriptor) SetStorageMode(value StorageMode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorageMode:"), value)
}


// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/swizzle
func (t_ TextureDescriptor) Swizzle() TextureSwizzleChannels /* not a class type */ {
	rv := objc.Send[TextureSwizzleChannels](t_.ID, objc.Sel("swizzle"))
	return rv
}


// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/swizzle
func (t_ TextureDescriptor) SetSwizzle(value TextureSwizzleChannels /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSwizzle:"), value)
}


// The dimension and arrangement of texture image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/texturetype
func (t_ TextureDescriptor) TextureType() TextureType {
	rv := objc.Send[TextureType](t_.ID, objc.Sel("textureType"))
	return rv
}


// The dimension and arrangement of texture image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/texturetype
func (t_ TextureDescriptor) SetTextureType(value TextureType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}


// The width of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/width
func (t_ TextureDescriptor) Width() int {
	rv := objc.Send[int](t_.ID, objc.Sel("width"))
	return rv
}


// The width of the texture image for the base level mipmap, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturedescriptor/width
func (t_ TextureDescriptor) SetWidth(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}



