// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAFixedSpatialAudio */


/* debug [class_header]: Header for CAFixedSpatialAudio */
// The class instance for the [FixedSpatialAudio] class.
var (
	FixedSpatialAudioClass     _FixedSpatialAudioClass
	FixedSpatialAudioClassOnce sync.Once
)

func getFixedSpatialAudioClass() _FixedSpatialAudioClass {
	FixedSpatialAudioClassOnce.Do(func() {
		FixedSpatialAudioClass = _FixedSpatialAudioClass{objc.GetClass("CAFixedSpatialAudio")}
	})
	return FixedSpatialAudioClass
}

type _FixedSpatialAudioClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FixedSpatialAudio */
// An interface definition for the [FixedSpatialAudio] class.
type IFixedSpatialAudio interface {
	ISpatialAudioExperience
	
/* debug [class_interface_properties]: Properties for FixedSpatialAudio */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FixedSpatialAudio */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FixedSpatialAudio */
// Alloc allocates a new instance without initialization.
func (fc _FixedSpatialAudioClass) Alloc() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FixedSpatialAudioClass) New() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FixedSpatialAudio) Init() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FixedSpatialAudio) Autorelease() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFixedSpatialAudio creates a new FixedSpatialAudio instance.
func NewFixedSpatialAudio() FixedSpatialAudio {
	return getFixedSpatialAudioClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FixedSpatialAudio */
// A spatial experience that does not take user motion into account.
//
// The Objective-C version of the Swift type.


// A spatial experience that does not take user motion into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio
type FixedSpatialAudio struct {
	SpatialAudioExperience
}

// FixedSpatialAudioFrom constructs a [FixedSpatialAudio] from an unsafe.Pointer.
//
// A spatial experience that does not take user motion into account.
func FixedSpatialAudioFrom(ptr unsafe.Pointer) FixedSpatialAudio {
	return FixedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FixedSpatialAudio */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio/initWithSoundStageSize:
func NewFixedSpatialAudioWithSoundStageSize(soundStageSize SoundStageSize) FixedSpatialAudio {
	instance := getFixedSpatialAudioClass().Alloc()
	rv := objc.Send[FixedSpatialAudio](instance.ID, objc.Sel("initWithSoundStageSize:"), soundStageSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFixedSpatialAudioWithSoundStageSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FixedSpatialAudio */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FixedSpatialAudio */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FixedSpatialAudio */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FixedSpatialAudio */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAFixedSpatialAudio */


