// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioMixingDestination */


/* debug [class_header]: Header for AVAudioMixingDestination */
// The class instance for the [AudioMixingDestination] class.
var (
	AudioMixingDestinationClass     _AudioMixingDestinationClass
	AudioMixingDestinationClassOnce sync.Once
)

func getAudioMixingDestinationClass() _AudioMixingDestinationClass {
	AudioMixingDestinationClassOnce.Do(func() {
		AudioMixingDestinationClass = _AudioMixingDestinationClass{objc.GetClass("AVAudioMixingDestination")}
	})
	return AudioMixingDestinationClass
}

type _AudioMixingDestinationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioMixingDestination */
// An interface definition for the [AudioMixingDestination] class.
type IAudioMixingDestination interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioMixingDestination */
	// properties:
	ConnectionPoint() IAVAudioConnectionPoint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioMixingDestination */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioMixingDestination */
// Alloc allocates a new instance without initialization.
func (ac _AudioMixingDestinationClass) Alloc() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioMixingDestinationClass) New() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioMixingDestination) Init() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioMixingDestination) Autorelease() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioMixingDestination creates a new AudioMixingDestination instance.
func NewAudioMixingDestination() AudioMixingDestination {
	return getAudioMixingDestinationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioMixingDestination */
// An object that represents a connection to a mixer node from a node that conforms to the audio mixing protocol.
//
// You can only use a destination instance when a source node provides it. You can’t use it as a standalone instance.


// An object that represents a connection to a mixer node from a node that conforms to the audio mixing protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination
type AudioMixingDestination struct {
	objectivec.Object
}

// AudioMixingDestinationFrom constructs a [AudioMixingDestination] from an unsafe.Pointer.
//
// An object that represents a connection to a mixer node from a node that conforms to the audio mixing protocol.
func AudioMixingDestinationFrom(ptr unsafe.Pointer) AudioMixingDestination {
	return AudioMixingDestination{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioMixingDestination *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioMixingDestination */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioMixingDestination */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioMixingDestination */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioMixingDestination */

// The underlying mixer connection point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination/connectionPoint
func (a_ AudioMixingDestination) ConnectionPoint() IAVAudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](a_.ID, objc.Sel("connectionPoint"))
	return rv
}/* debug [instance_properties/getter]: connectionPoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioMixingDestination */



