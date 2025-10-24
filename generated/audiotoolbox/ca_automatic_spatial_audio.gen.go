// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAAutomaticSpatialAudio */


/* debug [class_header]: Header for CAAutomaticSpatialAudio */
// The class instance for the [AutomaticSpatialAudio] class.
var (
	AutomaticSpatialAudioClass     _AutomaticSpatialAudioClass
	AutomaticSpatialAudioClassOnce sync.Once
)

func getAutomaticSpatialAudioClass() _AutomaticSpatialAudioClass {
	AutomaticSpatialAudioClassOnce.Do(func() {
		AutomaticSpatialAudioClass = _AutomaticSpatialAudioClass{objc.GetClass("CAAutomaticSpatialAudio")}
	})
	return AutomaticSpatialAudioClass
}

type _AutomaticSpatialAudioClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AutomaticSpatialAudio */
// An interface definition for the [AutomaticSpatialAudio] class.
type IAutomaticSpatialAudio interface {
	ISpatialAudioExperience
	
/* debug [class_interface_properties]: Properties for AutomaticSpatialAudio */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AutomaticSpatialAudio */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AutomaticSpatialAudio */
// Alloc allocates a new instance without initialization.
func (ac _AutomaticSpatialAudioClass) Alloc() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AutomaticSpatialAudioClass) New() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutomaticSpatialAudio) Init() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutomaticSpatialAudio) Autorelease() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutomaticSpatialAudio creates a new AutomaticSpatialAudio instance.
func NewAutomaticSpatialAudio() AutomaticSpatialAudio {
	return getAutomaticSpatialAudioClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AutomaticSpatialAudio */
// A spatial audio experience determined by the system.
//
// The Objective-C version of the Swift type.


// A spatial audio experience determined by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAutomaticSpatialAudio
type AutomaticSpatialAudio struct {
	SpatialAudioExperience
}

// AutomaticSpatialAudioFrom constructs a [AutomaticSpatialAudio] from an unsafe.Pointer.
//
// A spatial audio experience determined by the system.
func AutomaticSpatialAudioFrom(ptr unsafe.Pointer) AutomaticSpatialAudio {
	return AutomaticSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AutomaticSpatialAudio */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AutomaticSpatialAudio */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AutomaticSpatialAudio */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AutomaticSpatialAudio */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AutomaticSpatialAudio */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAAutomaticSpatialAudio */


