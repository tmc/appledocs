// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioEnvironmentNode] class.
var (
	AudioEnvironmentNodeClass     _AudioEnvironmentNodeClass
	AudioEnvironmentNodeClassOnce sync.Once
)

func getAudioEnvironmentNodeClass() _AudioEnvironmentNodeClass {
	AudioEnvironmentNodeClassOnce.Do(func() {
		AudioEnvironmentNodeClass = _AudioEnvironmentNodeClass{objc.GetClass("AVAudioEnvironmentNode")}
	})
	return AudioEnvironmentNodeClass
}

type _AudioEnvironmentNodeClass struct {
	class objc.Class
}

// An interface definition for the [AudioEnvironmentNode] class.
type IAudioEnvironmentNode interface {
	IAudioNode
}

// An object that simulates a 3D audio environment.
//
// The class is a mixer node that simulates a 3D audio environment. Any node that conforms to can act as a source node, such as . The environment node has an implicit listener. You set the listener’s position and orientation, and the system then controls the way the user experiences the virtual world. To help characterize the environment, this class defines properties for distance attenuation and reverberation. affects how inputs with different channel configurations render. Spatialization applies only to inputs with a mono channel connection format. This class doesn’t spatialize stereo inputs or support inputs with connection formats of more than two channels. To set the node’s output to a multichannel format, use an that has one of the following :
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode
type AudioEnvironmentNode struct {
	AudioNode
}

// AudioEnvironmentNodeFrom constructs a [AudioEnvironmentNode] from an unsafe.Pointer.
//
// An object that simulates a 3D audio environment.
func AudioEnvironmentNodeFrom(ptr unsafe.Pointer) AudioEnvironmentNode {
	return AudioEnvironmentNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentNodeClass) Alloc() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioEnvironmentNodeClass) New() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEnvironmentNode) Init() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEnvironmentNode) Autorelease() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEnvironmentNode creates a new AudioEnvironmentNode instance.
func NewAudioEnvironmentNode() AudioEnvironmentNode {
	return getAudioEnvironmentNodeClass().New()
}



// An array of rendering algorithms applicable to the environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/applicableRenderingAlgorithms
func (a_ AudioEnvironmentNode) ApplicableRenderingAlgorithms() []foundation.NSNumber {
	rv := objc.Send[[]foundation.NSNumber](a_.ID, objc.Sel("applicableRenderingAlgorithms"))
	return rv
}

// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled
func (a_ AudioEnvironmentNode) ListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("listenerHeadTrackingEnabled"))
	return rv
}


// SetListenerHeadTrackingEnabled sets the value of the listenerHeadTrackingEnabled property.
// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled
func (a_ AudioEnvironmentNode) SetListenerHeadTrackingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerHeadTrackingEnabled:"), value)
}

// The listener’s angular orientation in the environment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation
func (a_ AudioEnvironmentNode) ListenerAngularOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("listenerAngularOrientation"))
	return rv
}


// SetListenerAngularOrientation sets the value of the listenerAngularOrientation property.
// The listener’s angular orientation in the environment.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation
func (a_ AudioEnvironmentNode) SetListenerAngularOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerAngularOrientation:"), value)
}


