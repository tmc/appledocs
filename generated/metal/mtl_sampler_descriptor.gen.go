// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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


