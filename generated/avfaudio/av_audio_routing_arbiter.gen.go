// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioRoutingArbiter */


/* debug [class_header]: Header for AVAudioRoutingArbiter */
// The class instance for the [AudioRoutingArbiter] class.
var (
	AudioRoutingArbiterClass     _AudioRoutingArbiterClass
	AudioRoutingArbiterClassOnce sync.Once
)

func getAudioRoutingArbiterClass() _AudioRoutingArbiterClass {
	AudioRoutingArbiterClassOnce.Do(func() {
		AudioRoutingArbiterClass = _AudioRoutingArbiterClass{objc.GetClass("AVAudioRoutingArbiter")}
	})
	return AudioRoutingArbiterClass
}

type _AudioRoutingArbiterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioRoutingArbiter */
// An interface definition for the [AudioRoutingArbiter] class.
type IAudioRoutingArbiter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioRoutingArbiter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioRoutingArbiter */
	// methods:
	BeginArbitrationWithCategoryCompletionHandler(category AudioRoutingArbitrationCategory, handler unsafe.Pointer)
	LeaveArbitration()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioRoutingArbiter */
// Alloc allocates a new instance without initialization.
func (ac _AudioRoutingArbiterClass) Alloc() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioRoutingArbiterClass) New() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioRoutingArbiter) Init() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioRoutingArbiter) Autorelease() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioRoutingArbiter creates a new AudioRoutingArbiter instance.
func NewAudioRoutingArbiter() AudioRoutingArbiter {
	return getAudioRoutingArbiterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioRoutingArbiter */
// An object for configuring macOS apps to participate in AirPods Automatic Switching.
//
// AirPods Automatic Switching is a feature of Apple operating systems that intelligently connects wireless headphones to the most appropriate audio device in a multidevice environment. For example, if a user plays a movie on iPad, and then locks the device and starts playing music on iPhone, the system automatically switches the source audio device from iPad to iPhone. iOS apps automatically participate in AirPods Automatic Switching. To enable your macOS app to participate in this behavior, use to indicate when your app starts and finishes playing or recording audio. For example, a Voice over IP (VoIP) app might request arbitration before starting a call, and when the arbitration completes, begin the VoIP session. Likewise, when the call ends, the app would end the VoIP session and leave arbitration.


// An object for configuring macOS apps to participate in AirPods Automatic Switching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter
type AudioRoutingArbiter struct {
	objectivec.Object
}

// AudioRoutingArbiterFrom constructs a [AudioRoutingArbiter] from an unsafe.Pointer.
//
// An object for configuring macOS apps to participate in AirPods Automatic Switching.
func AudioRoutingArbiterFrom(ptr unsafe.Pointer) AudioRoutingArbiter {
	return AudioRoutingArbiter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioRoutingArbiter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioRoutingArbiter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioRoutingArbiter */

// The shared routing arbiter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/shared
func (ac _AudioRoutingArbiterClass) SharedRoutingArbiter() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](objc.ID(ac.class), objc.Sel("sharedRoutingArbiter"))
	return rv
}/* debug [class_properties_class/property]: sharedRoutingArbiter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioRoutingArbiter */

// Begins routing arbitration to take ownership of a nearby Bluetooth audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/begin(category:completionHandler:)
func (a_ AudioRoutingArbiter) BeginArbitrationWithCategoryCompletionHandler(category AudioRoutingArbitrationCategory, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginArbitrationWithCategory:completionHandler:"), category, handler)
}/* debug [instance_methods/method]: BeginArbitrationWithCategoryCompletionHandler */


// Stops an app’s participation in audio routing arbitration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/leave()
func (a_ AudioRoutingArbiter) LeaveArbitration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("leaveArbitration"))
}/* debug [instance_methods/method]: LeaveArbitration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioRoutingArbiter */

// The shared routing arbiter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/shared
func (a_ AudioRoutingArbiter) SharedRoutingArbiter() IAVAudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](a_.ID, objc.Sel("sharedRoutingArbiter"))
	return rv
}/* debug [instance_properties/getter]: sharedRoutingArbiter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioRoutingArbiter */



