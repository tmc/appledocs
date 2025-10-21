// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SamplerDescriptor] class.
type ISamplerDescriptor interface {
	objectivec.IObject
}

// An object that you use to configure a texture sampler.
//
// To make a sampler, create and configure an instance and then call an instance’s method. After you create the sampler, you can release the descriptor or reconfigure its properties to create other samplers.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SamplerDescriptorClass) Alloc() SamplerDescriptor {
	rv := objc.Send[SamplerDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodaverage
func (s_ SamplerDescriptor) LodAverage() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("lodAverage"))
	return rv
}


// SetLodAverage sets the value of the lodAverage property.
// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodaverage
func (s_ SamplerDescriptor) SetLodAverage(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodAverage:"), value)
}

// The address mode for the texture height (t) coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/taddressmode
func (s_ SamplerDescriptor) TAddressMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("tAddressMode"))
	return rv
}


// SetTAddressMode sets the value of the tAddressMode property.
// The address mode for the texture height (t) coordinate.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/taddressmode
func (s_ SamplerDescriptor) SetTAddressMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTAddressMode:"), value)
}

// The sampler comparison function used when performing a sample compare operation on a depth texture.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/comparefunction
func (s_ SamplerDescriptor) CompareFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compareFunction"))
	return rv
}


// SetCompareFunction sets the value of the compareFunction property.
// The sampler comparison function used when performing a sample compare operation on a depth texture.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/comparefunction
func (s_ SamplerDescriptor) SetCompareFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCompareFunction:"), value)
}

// Sets the reduction mode for filtering contributing samples.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/reductionmode
func (s_ SamplerDescriptor) ReductionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("reductionMode"))
	return rv
}


// SetReductionMode sets the value of the reductionMode property.
// Sets the reduction mode for filtering contributing samples.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/reductionmode
func (s_ SamplerDescriptor) SetReductionMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReductionMode:"), value)
}

// A Boolean value that indicates whether you can reference a sampler, that you make
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/supportargumentbuffers
func (s_ SamplerDescriptor) SupportArgumentBuffers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("supportArgumentBuffers"))
	return rv
}


// SetSupportArgumentBuffers sets the value of the supportArgumentBuffers property.
// A Boolean value that indicates whether you can reference a sampler, that you make

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/supportargumentbuffers
func (s_ SamplerDescriptor) SetSupportArgumentBuffers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSupportArgumentBuffers:"), value)
}

// Sets the level-of-detail (lod) bias when sampling from a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodbias
func (s_ SamplerDescriptor) LodBias() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lodBias"))
	return rv
}


// SetLodBias sets the value of the lodBias property.
// Sets the level-of-detail (lod) bias when sampling from a texture.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodbias
func (s_ SamplerDescriptor) SetLodBias(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodBias:"), value)
}

// The border color for clamped texture values.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/bordercolor
func (s_ SamplerDescriptor) BorderColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("borderColor"))
	return rv
}


// SetBorderColor sets the value of the borderColor property.
// The border color for clamped texture values.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/bordercolor
func (s_ SamplerDescriptor) SetBorderColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderColor:"), value)
}

// The address mode for the texture width (s) coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/saddressmode
func (s_ SamplerDescriptor) SAddressMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sAddressMode"))
	return rv
}


// SetSAddressMode sets the value of the sAddressMode property.
// The address mode for the texture width (s) coordinate.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/saddressmode
func (s_ SamplerDescriptor) SetSAddressMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSAddressMode:"), value)
}

// A Boolean value that indicates whether texture coordinates are normalized to the range
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/normalizedcoordinates
func (s_ SamplerDescriptor) NormalizedCoordinates() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("normalizedCoordinates"))
	return rv
}


// SetNormalizedCoordinates sets the value of the normalizedCoordinates property.
// A Boolean value that indicates whether texture coordinates are normalized to the range

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/normalizedcoordinates
func (s_ SamplerDescriptor) SetNormalizedCoordinates(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNormalizedCoordinates:"), value)
}

// The address mode for the texture depth (r) coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/raddressmode
func (s_ SamplerDescriptor) RAddressMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rAddressMode"))
	return rv
}


// SetRAddressMode sets the value of the rAddressMode property.
// The address mode for the texture depth (r) coordinate.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/raddressmode
func (s_ SamplerDescriptor) SetRAddressMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRAddressMode:"), value)
}

// The filtering option for combining pixels between two mipmap levels.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/mipfilter
func (s_ SamplerDescriptor) MipFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("mipFilter"))
	return rv
}


// SetMipFilter sets the value of the mipFilter property.
// The filtering option for combining pixels between two mipmap levels.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/mipfilter
func (s_ SamplerDescriptor) SetMipFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMipFilter:"), value)
}

// The maximum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodmaxclamp
func (s_ SamplerDescriptor) LodMaxClamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lodMaxClamp"))
	return rv
}


// SetLodMaxClamp sets the value of the lodMaxClamp property.
// The maximum level of detail (LOD) to use when sampling from a texture.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/lodmaxclamp
func (s_ SamplerDescriptor) SetLodMaxClamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodMaxClamp:"), value)
}

// A string that identifies the sampler.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/label
func (s_ SamplerDescriptor) Label() string {
	rv := objc.Send[string](s_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the sampler.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/label
func (s_ SamplerDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/minfilter
func (s_ SamplerDescriptor) MinFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("minFilter"))
	return rv
}


// SetMinFilter sets the value of the minFilter property.
// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/minfilter
func (s_ SamplerDescriptor) SetMinFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinFilter:"), value)
}

// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/magfilter
func (s_ SamplerDescriptor) MagFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("magFilter"))
	return rv
}


// SetMagFilter sets the value of the magFilter property.
// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerdescriptor/magfilter
func (s_ SamplerDescriptor) SetMagFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagFilter:"), value)
}

// The minimum level of detail (LOD) to use when sampling from a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMinClamp
func (s_ SamplerDescriptor) LodMinClamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lodMinClamp"))
	return rv
}


// SetLodMinClamp sets the value of the lodMinClamp property.
// The minimum level of detail (LOD) to use when sampling from a texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMinClamp
func (s_ SamplerDescriptor) SetLodMinClamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLodMinClamp:"), value)
}

// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/maxAnisotropy
func (s_ SamplerDescriptor) MaxAnisotropy() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("maxAnisotropy"))
	return rv
}


// SetMaxAnisotropy sets the value of the maxAnisotropy property.
// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/maxAnisotropy
func (s_ SamplerDescriptor) SetMaxAnisotropy(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxAnisotropy:"), value)
}



