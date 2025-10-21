// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioRoutingArbiter] class.
type IAudioRoutingArbiter interface {
	objectivec.IObject
	BeginArbitrationWithCategoryCompletionHandler(category unsafe.Pointer, handler unsafe.Pointer)
	LeaveArbitration()
}

// An object for configuring macOS apps to participate in AirPods Automatic Switching.
//
// AirPods Automatic Switching is a feature of Apple operating systems that intelligently connects wireless headphones to the most appropriate audio device in a multidevice environment. For example, if a user plays a movie on iPad, and then locks the device and starts playing music on iPhone, the system automatically switches the source audio device from iPad to iPhone. iOS apps automatically participate in AirPods Automatic Switching. To enable your macOS app to participate in this behavior, use to indicate when your app starts and finishes playing or recording audio. For example, a Voice over IP (VoIP) app might request arbitration before starting a call, and when the arbitration completes, begin the VoIP session. Likewise, when the call ends, the app would end the VoIP session and leave arbitration.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioRoutingArbiterClass) Alloc() AudioRoutingArbiter {
	rv := objc.Send[AudioRoutingArbiter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The shared routing arbiter object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/shared
func (ac _AudioRoutingArbiterClass) SharedRoutingArbiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("sharedRoutingArbiter"))
	return rv
}
// Begins routing arbitration to take ownership of a nearby Bluetooth audio route.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/begin(category:completionHandler:)
func (a_ AudioRoutingArbiter) BeginArbitrationWithCategoryCompletionHandler(category unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginArbitrationWithCategory:completionHandler:"), category, handler)
}

// Stops an app’s participation in audio routing arbitration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/leave()
func (a_ AudioRoutingArbiter) LeaveArbitration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("leaveArbitration"))
}

// The shared routing arbiter object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/shared
func (a_ AudioRoutingArbiter) SharedRoutingArbiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sharedRoutingArbiter"))
	return rv
}



