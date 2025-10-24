// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSessionRouteDescription */


/* debug [class_header]: Header for AVAudioSessionRouteDescription */
// The class instance for the [AudioSessionRouteDescription] class.
var (
	AudioSessionRouteDescriptionClass     _AudioSessionRouteDescriptionClass
	AudioSessionRouteDescriptionClassOnce sync.Once
)

func getAudioSessionRouteDescriptionClass() _AudioSessionRouteDescriptionClass {
	AudioSessionRouteDescriptionClassOnce.Do(func() {
		AudioSessionRouteDescriptionClass = _AudioSessionRouteDescriptionClass{objc.GetClass("AVAudioSessionRouteDescription")}
	})
	return AudioSessionRouteDescriptionClass
}

type _AudioSessionRouteDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSessionRouteDescription */
// An interface definition for the [AudioSessionRouteDescription] class.
type IAudioSessionRouteDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSessionRouteDescription */
	// properties:
	CurrentRoute() IAVAudioSessionRouteDescription
	SetCurrentRoute(value IAVAudioSessionRouteDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSessionRouteDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSessionRouteDescription */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionRouteDescriptionClass) Alloc() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionRouteDescriptionClass) New() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionRouteDescription) Init() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionRouteDescription) Autorelease() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionRouteDescription creates a new AudioSessionRouteDescription instance.
func NewAudioSessionRouteDescription() AudioSessionRouteDescription {
	return getAudioSessionRouteDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSessionRouteDescription */
// An object that describes the input and output ports associated with a session’s audio route.
//
// You don’t create instances of this class yourself. Instead, you retrieve the current audio route from your app’s object.


// An object that describes the input and output ports associated with a session’s audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionRouteDescription
type AudioSessionRouteDescription struct {
	objectivec.Object
}

// AudioSessionRouteDescriptionFrom constructs a [AudioSessionRouteDescription] from an unsafe.Pointer.
//
// An object that describes the input and output ports associated with a session’s audio route.
func AudioSessionRouteDescriptionFrom(ptr unsafe.Pointer) AudioSessionRouteDescription {
	return AudioSessionRouteDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSessionRouteDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSessionRouteDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSessionRouteDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSessionRouteDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSessionRouteDescription */

// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionRouteDescription) CurrentRoute() IAVAudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("currentRoute"))
	return rv
}/* debug [instance_properties/getter]: currentRoute */


// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionRouteDescription) SetCurrentRoute(value IAVAudioSessionRouteDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentRoute:"), value)
}/* debug [instance_properties/setter]: currentRoute */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSessionRouteDescription */


