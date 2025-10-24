// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSpatialAudioMetadataSampleGenerator */


/* debug [class_header]: Header for AVCaptureSpatialAudioMetadataSampleGenerator */
// The class instance for the [CaptureSpatialAudioMetadataSampleGenerator] class.
var (
	CaptureSpatialAudioMetadataSampleGeneratorClass     _CaptureSpatialAudioMetadataSampleGeneratorClass
	CaptureSpatialAudioMetadataSampleGeneratorClassOnce sync.Once
)

func getCaptureSpatialAudioMetadataSampleGeneratorClass() _CaptureSpatialAudioMetadataSampleGeneratorClass {
	CaptureSpatialAudioMetadataSampleGeneratorClassOnce.Do(func() {
		CaptureSpatialAudioMetadataSampleGeneratorClass = _CaptureSpatialAudioMetadataSampleGeneratorClass{objc.GetClass("AVCaptureSpatialAudioMetadataSampleGenerator")}
	})
	return CaptureSpatialAudioMetadataSampleGeneratorClass
}

type _CaptureSpatialAudioMetadataSampleGeneratorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSpatialAudioMetadataSampleGenerator */
// An interface definition for the [CaptureSpatialAudioMetadataSampleGenerator] class.
type ICaptureSpatialAudioMetadataSampleGenerator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSpatialAudioMetadataSampleGenerator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSpatialAudioMetadataSampleGenerator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSpatialAudioMetadataSampleGenerator */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSpatialAudioMetadataSampleGeneratorClass) Alloc() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSpatialAudioMetadataSampleGeneratorClass) New() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSpatialAudioMetadataSampleGenerator) Init() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSpatialAudioMetadataSampleGenerator) Autorelease() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSpatialAudioMetadataSampleGenerator creates a new CaptureSpatialAudioMetadataSampleGenerator instance.
func NewCaptureSpatialAudioMetadataSampleGenerator() CaptureSpatialAudioMetadataSampleGenerator {
	return getCaptureSpatialAudioMetadataSampleGeneratorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSpatialAudioMetadataSampleGenerator */
// An interface for generating a spatial audio timed metadata sample.


// An interface for generating a spatial audio timed metadata sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator
type CaptureSpatialAudioMetadataSampleGenerator struct {
	objectivec.Object
}

// CaptureSpatialAudioMetadataSampleGeneratorFrom constructs a [CaptureSpatialAudioMetadataSampleGenerator] from an unsafe.Pointer.
//
// An interface for generating a spatial audio timed metadata sample.
func CaptureSpatialAudioMetadataSampleGeneratorFrom(ptr unsafe.Pointer) CaptureSpatialAudioMetadataSampleGenerator {
	return CaptureSpatialAudioMetadataSampleGenerator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSpatialAudioMetadataSampleGenerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSpatialAudioMetadataSampleGenerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSpatialAudioMetadataSampleGenerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSpatialAudioMetadataSampleGenerator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSpatialAudioMetadataSampleGenerator */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSpatialAudioMetadataSampleGenerator */


