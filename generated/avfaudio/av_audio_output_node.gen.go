// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioOutputNode */


/* debug [class_header]: Header for AVAudioOutputNode */
// The class instance for the [AudioOutputNode] class.
var (
	AudioOutputNodeClass     _AudioOutputNodeClass
	AudioOutputNodeClassOnce sync.Once
)

func getAudioOutputNodeClass() _AudioOutputNodeClass {
	AudioOutputNodeClassOnce.Do(func() {
		AudioOutputNodeClass = _AudioOutputNodeClass{objc.GetClass("AVAudioOutputNode")}
	})
	return AudioOutputNodeClass
}

type _AudioOutputNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioOutputNode */
// An interface definition for the [AudioOutputNode] class.
type IAudioOutputNode interface {
	IAudioIONode
	
/* debug [class_interface_properties]: Properties for AudioOutputNode */
	// properties:
	ManualRenderingFormat() IAVAudioFormat
	SetManualRenderingFormat(value IAVAudioFormat)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioOutputNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioOutputNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioOutputNodeClass) Alloc() AudioOutputNode {
	rv := objc.Send[AudioOutputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioOutputNodeClass) New() AudioOutputNode {
	rv := objc.Send[AudioOutputNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioOutputNode) Init() AudioOutputNode {
	rv := objc.Send[AudioOutputNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioOutputNode) Autorelease() AudioOutputNode {
	rv := objc.Send[AudioOutputNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioOutputNode creates a new AudioOutputNode instance.
func NewAudioOutputNode() AudioOutputNode {
	return getAudioOutputNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioOutputNode */
// An object that connects to the system’s audio output.
//
// This node connects to the system’s audio output when rendering to or from an audio device. This node performs output in response to client’s requests when the engine is in manual rendering mode. This audio node has one element. The format of the output scope reflects: The audio hardware sample rate and channel count when it connects to the hardware. The engine’s manual rendering mode output format (see ). The format of the input scope is initially the same as that of the output, but you may set it to a different format, in which case the audio node converts.


// An object that connects to the system’s audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode
type AudioOutputNode struct {
	AudioIONode
}

// AudioOutputNodeFrom constructs a [AudioOutputNode] from an unsafe.Pointer.
//
// An object that connects to the system’s audio output.
func AudioOutputNodeFrom(ptr unsafe.Pointer) AudioOutputNode {
	return AudioOutputNode{
		AudioIONode: AudioIONodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioOutputNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioOutputNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioOutputNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioOutputNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioOutputNode */

// The render format of the engine in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingformat
func (a_ AudioOutputNode) ManualRenderingFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("manualRenderingFormat"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingFormat */


// The render format of the engine in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingformat
func (a_ AudioOutputNode) SetManualRenderingFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingFormat:"), value)
}/* debug [instance_properties/setter]: manualRenderingFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioOutputNode */


