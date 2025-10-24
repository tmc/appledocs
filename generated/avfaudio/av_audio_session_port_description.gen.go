// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSessionPortDescription */


/* debug [class_header]: Header for AVAudioSessionPortDescription */
// The class instance for the [AudioSessionPortDescription] class.
var (
	AudioSessionPortDescriptionClass     _AudioSessionPortDescriptionClass
	AudioSessionPortDescriptionClassOnce sync.Once
)

func getAudioSessionPortDescriptionClass() _AudioSessionPortDescriptionClass {
	AudioSessionPortDescriptionClassOnce.Do(func() {
		AudioSessionPortDescriptionClass = _AudioSessionPortDescriptionClass{objc.GetClass("AVAudioSessionPortDescription")}
	})
	return AudioSessionPortDescriptionClass
}

type _AudioSessionPortDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSessionPortDescription */
// An interface definition for the [AudioSessionPortDescription] class.
type IAudioSessionPortDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSessionPortDescription */
	// properties:
	AvailableInputs() IAVAudioSessionPortDescription
	SetAvailableInputs(value IAVAudioSessionPortDescription)
	CurrentRoute() IAVAudioSessionRouteDescription
	SetCurrentRoute(value IAVAudioSessionRouteDescription)
	IsSpatialAudioEnabled() bool
	SetIsSpatialAudioEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSessionPortDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSessionPortDescription */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionPortDescriptionClass) Alloc() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionPortDescriptionClass) New() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionPortDescription) Init() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionPortDescription) Autorelease() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionPortDescription creates a new AudioSessionPortDescription instance.
func NewAudioSessionPortDescription() AudioSessionPortDescription {
	return getAudioSessionPortDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSessionPortDescription */
// Information about the capabilities of the port and the hardware channels it supports.
//
// A port description object describes a single input or output port associated with an audio route. Examples of audio ports include a device’s built-in speaker, a microphone on a wired headset, and a Bluetooth device supporting the Advanced Audio Distribution Profile (A2DP). You can query the audio session’s property to get information about the active set of input and output ports. To change the current audio routing, call the method. For example, on a device with a wired headset attached, the audio session’s array may contain two port descriptions: one for the headset microphone and one for the device’s built-in microphone. You can use the audio session’s method to select the headset or built-in microphone for audio input.


// Information about the capabilities of the port and the hardware channels it supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription
type AudioSessionPortDescription struct {
	objectivec.Object
}

// AudioSessionPortDescriptionFrom constructs a [AudioSessionPortDescription] from an unsafe.Pointer.
//
// Information about the capabilities of the port and the hardware channels it supports.
func AudioSessionPortDescriptionFrom(ptr unsafe.Pointer) AudioSessionPortDescription {
	return AudioSessionPortDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSessionPortDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSessionPortDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSessionPortDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSessionPortDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSessionPortDescription */

// An array of input ports available for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/availableinputs
func (a_ AudioSessionPortDescription) AvailableInputs() IAVAudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("availableInputs"))
	return rv
}/* debug [instance_properties/getter]: availableInputs */


// An array of input ports available for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/availableinputs
func (a_ AudioSessionPortDescription) SetAvailableInputs(value IAVAudioSessionPortDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableInputs:"), value)
}/* debug [instance_properties/setter]: availableInputs */


// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionPortDescription) CurrentRoute() IAVAudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("currentRoute"))
	return rv
}/* debug [instance_properties/getter]: currentRoute */


// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionPortDescription) SetCurrentRoute(value IAVAudioSessionRouteDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentRoute:"), value)
}/* debug [instance_properties/setter]: currentRoute */


// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionPortDescription) IsSpatialAudioEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSpatialAudioEnabled"))
	return rv
}/* debug [instance_properties/getter]: isSpatialAudioEnabled */


// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionPortDescription) SetIsSpatialAudioEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSpatialAudioEnabled:"), value)
}/* debug [instance_properties/setter]: isSpatialAudioEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSessionPortDescription */


