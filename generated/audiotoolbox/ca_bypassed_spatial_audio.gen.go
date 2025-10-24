// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CABypassedSpatialAudio */


/* debug [class_header]: Header for CABypassedSpatialAudio */
// The class instance for the [BypassedSpatialAudio] class.
var (
	BypassedSpatialAudioClass     _BypassedSpatialAudioClass
	BypassedSpatialAudioClassOnce sync.Once
)

func getBypassedSpatialAudioClass() _BypassedSpatialAudioClass {
	BypassedSpatialAudioClassOnce.Do(func() {
		BypassedSpatialAudioClass = _BypassedSpatialAudioClass{objc.GetClass("CABypassedSpatialAudio")}
	})
	return BypassedSpatialAudioClass
}

type _BypassedSpatialAudioClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BypassedSpatialAudio */
// An interface definition for the [BypassedSpatialAudio] class.
type IBypassedSpatialAudio interface {
	ISpatialAudioExperience
	
/* debug [class_interface_properties]: Properties for BypassedSpatialAudio */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BypassedSpatialAudio */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BypassedSpatialAudio */
// Alloc allocates a new instance without initialization.
func (bc _BypassedSpatialAudioClass) Alloc() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BypassedSpatialAudioClass) New() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BypassedSpatialAudio) Init() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BypassedSpatialAudio) Autorelease() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBypassedSpatialAudio creates a new BypassedSpatialAudio instance.
func NewBypassedSpatialAudio() BypassedSpatialAudio {
	return getBypassedSpatialAudioClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BypassedSpatialAudio */
// An experience in which the system does not apply spatial processing to the audio stream.
//
// The Objective-C version of the Swift type.


// An experience in which the system does not apply spatial processing to the audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CABypassedSpatialAudio
type BypassedSpatialAudio struct {
	SpatialAudioExperience
}

// BypassedSpatialAudioFrom constructs a [BypassedSpatialAudio] from an unsafe.Pointer.
//
// An experience in which the system does not apply spatial processing to the audio stream.
func BypassedSpatialAudioFrom(ptr unsafe.Pointer) BypassedSpatialAudio {
	return BypassedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BypassedSpatialAudio */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BypassedSpatialAudio */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BypassedSpatialAudio */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BypassedSpatialAudio */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BypassedSpatialAudio */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CABypassedSpatialAudio */


