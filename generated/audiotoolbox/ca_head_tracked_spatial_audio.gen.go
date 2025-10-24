// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAHeadTrackedSpatialAudio */


/* debug [class_header]: Header for CAHeadTrackedSpatialAudio */
// The class instance for the [HeadTrackedSpatialAudio] class.
var (
	HeadTrackedSpatialAudioClass     _HeadTrackedSpatialAudioClass
	HeadTrackedSpatialAudioClassOnce sync.Once
)

func getHeadTrackedSpatialAudioClass() _HeadTrackedSpatialAudioClass {
	HeadTrackedSpatialAudioClassOnce.Do(func() {
		HeadTrackedSpatialAudioClass = _HeadTrackedSpatialAudioClass{objc.GetClass("CAHeadTrackedSpatialAudio")}
	})
	return HeadTrackedSpatialAudioClass
}

type _HeadTrackedSpatialAudioClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HeadTrackedSpatialAudio */
// An interface definition for the [HeadTrackedSpatialAudio] class.
type IHeadTrackedSpatialAudio interface {
	ISpatialAudioExperience
	
/* debug [class_interface_properties]: Properties for HeadTrackedSpatialAudio */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HeadTrackedSpatialAudio */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HeadTrackedSpatialAudio */
// Alloc allocates a new instance without initialization.
func (hc _HeadTrackedSpatialAudioClass) Alloc() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HeadTrackedSpatialAudioClass) New() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HeadTrackedSpatialAudio) Init() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HeadTrackedSpatialAudio) Autorelease() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeadTrackedSpatialAudio creates a new HeadTrackedSpatialAudio instance.
func NewHeadTrackedSpatialAudio() HeadTrackedSpatialAudio {
	return getHeadTrackedSpatialAudioClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HeadTrackedSpatialAudio */
// A spatial experience that takes user motion into account.
//
// The Objective-C version of the Swift type.


// A spatial experience that takes user motion into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio
type HeadTrackedSpatialAudio struct {
	SpatialAudioExperience
}

// HeadTrackedSpatialAudioFrom constructs a [HeadTrackedSpatialAudio] from an unsafe.Pointer.
//
// A spatial experience that takes user motion into account.
func HeadTrackedSpatialAudioFrom(ptr unsafe.Pointer) HeadTrackedSpatialAudio {
	return HeadTrackedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HeadTrackedSpatialAudio */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio/initWithSoundStageSize:anchoringStrategy:
func NewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy(soundStageSize SoundStageSize, anchoringStrategy IAnchoringStrategy) HeadTrackedSpatialAudio {
	instance := getHeadTrackedSpatialAudioClass().Alloc()
	rv := objc.Send[HeadTrackedSpatialAudio](instance.ID, objc.Sel("initWithSoundStageSize:anchoringStrategy:"), soundStageSize, anchoringStrategy)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HeadTrackedSpatialAudio */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HeadTrackedSpatialAudio */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HeadTrackedSpatialAudio */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HeadTrackedSpatialAudio */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAHeadTrackedSpatialAudio */


