// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CASpatialAudioExperience */


/* debug [class_header]: Header for CASpatialAudioExperience */
// The class instance for the [SpatialAudioExperience] class.
var (
	SpatialAudioExperienceClass     _SpatialAudioExperienceClass
	SpatialAudioExperienceClassOnce sync.Once
)

func getSpatialAudioExperienceClass() _SpatialAudioExperienceClass {
	SpatialAudioExperienceClassOnce.Do(func() {
		SpatialAudioExperienceClass = _SpatialAudioExperienceClass{objc.GetClass("CASpatialAudioExperience")}
	})
	return SpatialAudioExperienceClass
}

type _SpatialAudioExperienceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpatialAudioExperience */
// An interface definition for the [SpatialAudioExperience] class.
type ISpatialAudioExperience interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpatialAudioExperience */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpatialAudioExperience */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpatialAudioExperience */
// Alloc allocates a new instance without initialization.
func (sc _SpatialAudioExperienceClass) Alloc() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpatialAudioExperienceClass) New() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpatialAudioExperience) Init() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpatialAudioExperience) Autorelease() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpatialAudioExperience creates a new SpatialAudioExperience instance.
func NewSpatialAudioExperience() SpatialAudioExperience {
	return getSpatialAudioExperienceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpatialAudioExperience */
// Configure an audio stream for spatial computing.
//
// The Objective-C version of the Swift type.


// Configure an audio stream for spatial computing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASpatialAudioExperience
type SpatialAudioExperience struct {
	objectivec.Object
}

// SpatialAudioExperienceFrom constructs a [SpatialAudioExperience] from an unsafe.Pointer.
//
// Configure an audio stream for spatial computing.
func SpatialAudioExperienceFrom(ptr unsafe.Pointer) SpatialAudioExperience {
	return SpatialAudioExperience{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpatialAudioExperience *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpatialAudioExperience */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpatialAudioExperience */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpatialAudioExperience */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpatialAudioExperience */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CASpatialAudioExperience */



