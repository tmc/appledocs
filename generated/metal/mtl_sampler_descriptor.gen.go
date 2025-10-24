// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLSamplerDescriptor */


/* debug [class_header]: Header for MTLSamplerDescriptor */
// The class instance for the [SamplerDescriptor] class.
var (
	SamplerDescriptorClass     _SamplerDescriptorClass
	SamplerDescriptorClassOnce sync.Once
)

func getSamplerDescriptorClass() _SamplerDescriptorClass {
	SamplerDescriptorClassOnce.Do(func() {
		SamplerDescriptorClass = _SamplerDescriptorClass{objc.GetClass("MTLSamplerDescriptor")}
	})
	return SamplerDescriptorClass
}

type _SamplerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SamplerDescriptor */
// An interface definition for the [SamplerDescriptor] class.
type ISamplerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SamplerDescriptor */
	// properties:
	BorderColor() SamplerBorderColor
	SetBorderColor(value SamplerBorderColor)
	CompareFunction() CompareFunction
	SetCompareFunction(value CompareFunction)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LodAverage() bool
	SetLodAverage(value bool)
	LodBias() float32
	SetLodBias(value float32)
	LodMaxClamp() float32
	SetLodMaxClamp(value float32)
	LodMinClamp() float32
	SetLodMinClamp(value float32)
	MagFilter() SamplerMinMagFilter
	SetMagFilter(value SamplerMinMagFilter)
	MaxAnisotropy() uint
	SetMaxAnisotropy(value uint)
	MinFilter() SamplerMinMagFilter
	SetMinFilter(value SamplerMinMagFilter)
	MipFilter() SamplerMipFilter
	SetMipFilter(value SamplerMipFilter)
	NormalizedCoordinates() bool
	SetNormalizedCoordinates(value bool)
	RAddressMode() SamplerAddressMode
	SetRAddressMode(value SamplerAddressMode)
	ReductionMode() SamplerReductionMode
	SetReductionMode(value SamplerReductionMode)
	SAddressMode() SamplerAddressMode
	SetSAddressMode(value SamplerAddressMode)
	SupportArgumentBuffers() bool
	SetSupportArgumentBuffers(value bool)
	TAddressMode() SamplerAddressMode
	SetTAddressMode(value SamplerAddressMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SamplerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SamplerDescriptor */
// Alloc allocates a new instance without initialization.
func (sc _SamplerDescriptorClass) Alloc() SamplerDescriptor {
	rv := objc.Send[SamplerDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SamplerDescriptorClass) New() SamplerDescriptor {
	rv := objc.Send[SamplerDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SamplerDescriptor) Init() SamplerDescriptor {
	rv := objc.Send[SamplerDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SamplerDescriptor) Autorelease() SamplerDescriptor {
	rv := objc.Send[SamplerDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSamplerDescriptor creates a new SamplerDescriptor instance.
func NewSamplerDescriptor() SamplerDescriptor {
	return getSamplerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SamplerDescriptor */
// An object that you use to configure a texture sampler.
//
// To make a sampler, create and configure an instance and then call an instance’s method. After you create the sampler, you can release the descriptor or reconfigure its properties to create other samplers.


// An object that you use to configure a texture sampler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor
type SamplerDescriptor struct {
	objectivec.Object
}

// SamplerDescriptorFrom constructs a [SamplerDescriptor] from an unsafe.Pointer.
//
// An object that you use to configure a texture sampler.
func SamplerDescriptorFrom(ptr unsafe.Pointer) SamplerDescriptor {
	return SamplerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SamplerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SamplerDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SamplerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SamplerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SamplerDescriptor */

// The border color for clamped texture values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/borderColor
func (s_ SamplerDescriptor) BorderColor() SamplerBorderColor {
	rv := objc.Send[SamplerBorderColor](s_.ID, objc.Sel("borderColor"))
	return rv
}/* debug [instance_properties/getter]: borderColor */


// The border color for clamped texture values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/borderColor
func (s_ SamplerDescriptor) SetBorderColor(value SamplerBorderColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderColor:"), value)
}/* debug [instance_properties/setter]: borderColor */


// The sampler comparison function used when performing a sample compare operation on a depth texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/compareFunction
func (s_ SamplerDescriptor) CompareFunction() CompareFunction {
	rv := objc.Send[CompareFunction](s_.ID, objc.Sel("compareFunction"))
	return rv
}/* debug [instance_properties/getter]: compareFunction */


// The sampler comparison function used when performing a sample compare operation on a depth texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/compareFunction
func (s_ SamplerDescriptor) SetCompareFunction(value CompareFunction) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCompareFunction:"), value)
}/* debug [instance_properties/setter]: compareFunction */


// A string that identifies the sampler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/label
func (s_ SamplerDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifies the sampler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/label
func (s_ SamplerDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodAverage
func (s_ SamplerDescriptor) LodAverage() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("lodAverage"))
	return rv
}/* debug [instance_properties/getter]: lodAverage */


// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodAverage
func (s_ SamplerDescriptor) SetLodAverage(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodAverage:"), value)
}/* debug [instance_properties/setter]: lodAverage */


// Sets the level-of-detail (lod) bias when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodBias
func (s_ SamplerDescriptor) LodBias() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("lodBias"))
	return rv
}/* debug [instance_properties/getter]: lodBias */


// Sets the level-of-detail (lod) bias when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodBias
func (s_ SamplerDescriptor) SetLodBias(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodBias:"), value)
}/* debug [instance_properties/setter]: lodBias */


// The maximum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMaxClamp
func (s_ SamplerDescriptor) LodMaxClamp() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("lodMaxClamp"))
	return rv
}/* debug [instance_properties/getter]: lodMaxClamp */


// The maximum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMaxClamp
func (s_ SamplerDescriptor) SetLodMaxClamp(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodMaxClamp:"), value)
}/* debug [instance_properties/setter]: lodMaxClamp */


// The minimum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMinClamp
func (s_ SamplerDescriptor) LodMinClamp() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("lodMinClamp"))
	return rv
}/* debug [instance_properties/getter]: lodMinClamp */


// The minimum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMinClamp
func (s_ SamplerDescriptor) SetLodMinClamp(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodMinClamp:"), value)
}/* debug [instance_properties/setter]: lodMinClamp */


// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/magFilter
func (s_ SamplerDescriptor) MagFilter() SamplerMinMagFilter {
	rv := objc.Send[SamplerMinMagFilter](s_.ID, objc.Sel("magFilter"))
	return rv
}/* debug [instance_properties/getter]: magFilter */


// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/magFilter
func (s_ SamplerDescriptor) SetMagFilter(value SamplerMinMagFilter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagFilter:"), value)
}/* debug [instance_properties/setter]: magFilter */


// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/maxAnisotropy
func (s_ SamplerDescriptor) MaxAnisotropy() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("maxAnisotropy"))
	return rv
}/* debug [instance_properties/getter]: maxAnisotropy */


// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/maxAnisotropy
func (s_ SamplerDescriptor) SetMaxAnisotropy(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxAnisotropy:"), value)
}/* debug [instance_properties/setter]: maxAnisotropy */


// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/minFilter
func (s_ SamplerDescriptor) MinFilter() SamplerMinMagFilter {
	rv := objc.Send[SamplerMinMagFilter](s_.ID, objc.Sel("minFilter"))
	return rv
}/* debug [instance_properties/getter]: minFilter */


// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/minFilter
func (s_ SamplerDescriptor) SetMinFilter(value SamplerMinMagFilter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinFilter:"), value)
}/* debug [instance_properties/setter]: minFilter */


// The filtering option for combining pixels between two mipmap levels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/mipFilter
func (s_ SamplerDescriptor) MipFilter() SamplerMipFilter {
	rv := objc.Send[SamplerMipFilter](s_.ID, objc.Sel("mipFilter"))
	return rv
}/* debug [instance_properties/getter]: mipFilter */


// The filtering option for combining pixels between two mipmap levels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/mipFilter
func (s_ SamplerDescriptor) SetMipFilter(value SamplerMipFilter) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMipFilter:"), value)
}/* debug [instance_properties/setter]: mipFilter */


// A Boolean value that indicates whether texture coordinates are normalized to the range .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/normalizedCoordinates
func (s_ SamplerDescriptor) NormalizedCoordinates() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("normalizedCoordinates"))
	return rv
}/* debug [instance_properties/getter]: normalizedCoordinates */


// A Boolean value that indicates whether texture coordinates are normalized to the range .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/normalizedCoordinates
func (s_ SamplerDescriptor) SetNormalizedCoordinates(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNormalizedCoordinates:"), value)
}/* debug [instance_properties/setter]: normalizedCoordinates */


// The address mode for the texture depth (r) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/rAddressMode
func (s_ SamplerDescriptor) RAddressMode() SamplerAddressMode {
	rv := objc.Send[SamplerAddressMode](s_.ID, objc.Sel("rAddressMode"))
	return rv
}/* debug [instance_properties/getter]: rAddressMode */


// The address mode for the texture depth (r) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/rAddressMode
func (s_ SamplerDescriptor) SetRAddressMode(value SamplerAddressMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRAddressMode:"), value)
}/* debug [instance_properties/setter]: rAddressMode */


// Sets the reduction mode for filtering contributing samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/reductionMode
func (s_ SamplerDescriptor) ReductionMode() SamplerReductionMode {
	rv := objc.Send[SamplerReductionMode](s_.ID, objc.Sel("reductionMode"))
	return rv
}/* debug [instance_properties/getter]: reductionMode */


// Sets the reduction mode for filtering contributing samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/reductionMode
func (s_ SamplerDescriptor) SetReductionMode(value SamplerReductionMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReductionMode:"), value)
}/* debug [instance_properties/setter]: reductionMode */


// The address mode for the texture width (s) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/sAddressMode
func (s_ SamplerDescriptor) SAddressMode() SamplerAddressMode {
	rv := objc.Send[SamplerAddressMode](s_.ID, objc.Sel("sAddressMode"))
	return rv
}/* debug [instance_properties/getter]: sAddressMode */


// The address mode for the texture width (s) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/sAddressMode
func (s_ SamplerDescriptor) SetSAddressMode(value SamplerAddressMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSAddressMode:"), value)
}/* debug [instance_properties/setter]: sAddressMode */


// A Boolean value that indicates whether you can reference a sampler, that you make with this descriptor, by its resource ID from an argument buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/supportArgumentBuffers
func (s_ SamplerDescriptor) SupportArgumentBuffers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("supportArgumentBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportArgumentBuffers */


// A Boolean value that indicates whether you can reference a sampler, that you make with this descriptor, by its resource ID from an argument buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/supportArgumentBuffers
func (s_ SamplerDescriptor) SetSupportArgumentBuffers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSupportArgumentBuffers:"), value)
}/* debug [instance_properties/setter]: supportArgumentBuffers */


// The address mode for the texture height (t) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/tAddressMode
func (s_ SamplerDescriptor) TAddressMode() SamplerAddressMode {
	rv := objc.Send[SamplerAddressMode](s_.ID, objc.Sel("tAddressMode"))
	return rv
}/* debug [instance_properties/getter]: tAddressMode */


// The address mode for the texture height (t) coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/tAddressMode
func (s_ SamplerDescriptor) SetTAddressMode(value SamplerAddressMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTAddressMode:"), value)
}/* debug [instance_properties/setter]: tAddressMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLSamplerDescriptor */



