// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CISampler */


/* debug [class_header]: Header for CISampler */
// The class instance for the [Sampler] class.
var (
	SamplerClass     _SamplerClass
	SamplerClassOnce sync.Once
)

func getSamplerClass() _SamplerClass {
	SamplerClassOnce.Do(func() {
		SamplerClass = _SamplerClass{objc.GetClass("CISampler")}
	})
	return SamplerClass
}

type _SamplerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Sampler */
// An interface definition for the [Sampler] class.
type ISampler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Sampler */
	// properties:
	Definition() ICIFilterShape
	Extent() corefoundation.CGRect
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Sampler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Sampler */
// Alloc allocates a new instance without initialization.
func (sc _SamplerClass) Alloc() Sampler {
	rv := objc.Send[Sampler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Sampler */
// An object that retrieves pixel samples for processing by a filter kernel.
//
// The class retrieves samples of images for processing by a object. A object defines a coordinate transform, and modes for interpolation and wrapping. You use objects in conjunction with other Core Image classes, such as , , and , to create custom filters.


// An object that retrieves pixel samples for processing by a filter kernel.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Sampler */

// Initializes a sampler with an image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:)
func NewSamplerWithImage(im ICIImage) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:"), im)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSamplerWithImage */


// Initializes the sampler with an image object using options specified as key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/initWithImage:keysAndValues:
func NewSamplerWithImageKeysAndValues(im ICIImage, key0 objc.IObject) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:keysAndValues:"), im, key0)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSamplerWithImageKeysAndValues */


// Initializes the sampler with an image object using options specified in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:options:)
func NewSamplerWithImageOptions(im ICIImage, dict objc.IObject /* cross-framework: NSDictionary */) Sampler {
	instance := getSamplerClass().Alloc()
	rv := objc.Send[Sampler](instance.ID, objc.Sel("initWithImage:options:"), im, dict)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSamplerWithImageOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Sampler */

// Creates and returns a sampler that references an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:
func (sc _SamplerClass) SamplerWithImage(im ICIImage) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("samplerWithImage:"), im)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SamplerWithImage) */


// Creates and returns a sampler that references an image using options specified as key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:keysAndValues:
func (sc _SamplerClass) SamplerWithImageKeysAndValues(im ICIImage, key0 objc.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("samplerWithImage:keysAndValues:"), im, key0)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SamplerWithImageKeysAndValues) */


// Creates and returns a sampler that references an image using options specified in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:options:
func (sc _SamplerClass) SamplerWithImageOptions(im ICIImage, dict objc.IObject /* cross-framework: NSDictionary */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("samplerWithImage:options:"), im, dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SamplerWithImageOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Sampler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Sampler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Sampler */

// The domain of definition (DOD) of the sampler
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/definition
func (s_ Sampler) Definition() ICIFilterShape {
	rv := objc.Send[FilterShape](s_.ID, objc.Sel("definition"))
	return rv
}/* debug [instance_properties/getter]: definition */


// The rectangle that specifies the extent of the sampler
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CISampler/extent
func (s_ Sampler) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("extent"))
	return rv
}/* debug [instance_properties/getter]: extent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CISampler */


