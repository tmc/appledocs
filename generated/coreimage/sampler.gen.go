// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Sampler] class.
var (
	samplerClass     _SamplerClass
	samplerClassOnce sync.Once
)

func getSamplerClass() _SamplerClass {
	samplerClassOnce.Do(func() {
		samplerClass = _SamplerClass{objc.GetClass("CISampler")}
	})
	return samplerClass
}

type _SamplerClass struct {
	class objc.Class
}

// An interface definition for the [Sampler] class.
type ISampler interface {
	objectivec.IObject
}

// An object that retrieves pixel samples for processing by a filter kernel.
//
// The class retrieves samples of images for processing by a object. A object defines a coordinate transform, and modes for interpolation and wrapping. You use objects in conjunction with other Core Image classes, such as , , and , to create custom filters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler
type Sampler struct {
	objectivec.Object
}

// SamplerFrom constructs a [Sampler] from an unsafe.Pointer.
//
// An object that retrieves pixel samples for processing by a filter kernel.
func SamplerFrom(ptr unsafe.Pointer) Sampler {
	return Sampler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SamplerClass) Alloc() Sampler {
	rv := objc.Send[Sampler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SamplerClass) New() Sampler {
	rv := objc.Send[Sampler](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Sampler) Init() Sampler {
	rv := objc.Send[Sampler](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Sampler) Autorelease() Sampler {
	rv := objc.Send[Sampler](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampler creates a new Sampler instance.
func NewSampler() Sampler {
	return getSamplerClass().New()
}


// Initializes the sampler with an image object using options specified as key-value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/initWithImage:keysAndValues:
func NewSamplerWithImageKeysAndValues(im unsafe.Pointer, key0 objc.ID) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:keysAndValues:"), im, key0)
	rv.Autorelease()
	return rv
}

// Initializes a sampler with an image object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:)
func NewSamplerWithImage(im unsafe.Pointer) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:"), im)
	rv.Autorelease()
	return rv
}

// Initializes the sampler with an image object using options specified in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:options:)
func NewSamplerWithImageOptions(im unsafe.Pointer, dict objc.ID) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:options:"), im, dict)
	rv.Autorelease()
	return rv
}


// Creates and returns a sampler that references an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:
func (sc _SamplerClass) SamplerWithImage(im unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("samplerWithImage:"), im)
	return rv
}

// Creates and returns a sampler that references an image using options specified as key-value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:keysAndValues:
func (sc _SamplerClass) SamplerWithImageKeysAndValues(im unsafe.Pointer, key0 objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("samplerWithImage:keysAndValues:"), im, key0)
	return rv
}

// Creates and returns a sampler that references an image using options specified in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:options:
func (sc _SamplerClass) SamplerWithImageOptions(im unsafe.Pointer, dict objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("samplerWithImage:options:"), im, dict)
	return rv
}

// The domain of definition (DOD) of the sampler
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/definition
func (s_ Sampler) Definition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("definition"))
	return rv
}

// The rectangle that specifies the extent of the sampler
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/extent
func (s_ Sampler) Extent() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("extent"))
	return rv
}


