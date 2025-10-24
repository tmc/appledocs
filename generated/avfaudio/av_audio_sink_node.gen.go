// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSinkNode */


/* debug [class_header]: Header for AVAudioSinkNode */
// The class instance for the [AudioSinkNode] class.
var (
	AudioSinkNodeClass     _AudioSinkNodeClass
	AudioSinkNodeClassOnce sync.Once
)

func getAudioSinkNodeClass() _AudioSinkNodeClass {
	AudioSinkNodeClassOnce.Do(func() {
		AudioSinkNodeClass = _AudioSinkNodeClass{objc.GetClass("AVAudioSinkNode")}
	})
	return AudioSinkNodeClass
}

type _AudioSinkNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSinkNode */
// An interface definition for the [AudioSinkNode] class.
type IAudioSinkNode interface {
	IAudioNode
	
/* debug [class_interface_properties]: Properties for AudioSinkNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSinkNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSinkNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioSinkNodeClass) Alloc() AudioSinkNode {
	rv := objc.Send[AudioSinkNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSinkNodeClass) New() AudioSinkNode {
	rv := objc.Send[AudioSinkNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSinkNode) Init() AudioSinkNode {
	rv := objc.Send[AudioSinkNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSinkNode) Autorelease() AudioSinkNode {
	rv := objc.Send[AudioSinkNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSinkNode creates a new AudioSinkNode instance.
func NewAudioSinkNode() AudioSinkNode {
	return getAudioSinkNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSinkNode */
// An object that receives audio data.
//
// You use an to receive audio data through . You only use it in the input chain. An audio sink node doesn’t support format conversion. When connecting, use the output format of the input for the format for the connection. The format should match the hardware input sample rate. The voice processing I/O unit is an exception to the above because it supports sample rate conversion. The input scope format (hardware format) and output scope format (client format) of the input node can differ in that case. An audio sink node doesn’t support manual rendering mode, and doesn’t have an output bus, so you can’t install a tap on it.


// An object that receives audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode
type AudioSinkNode struct {
	AudioNode
}

// AudioSinkNodeFrom constructs a [AudioSinkNode] from an unsafe.Pointer.
//
// An object that receives audio data.
func AudioSinkNodeFrom(ptr unsafe.Pointer) AudioSinkNode {
	return AudioSinkNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSinkNode */

// Creates an audio sink node with a block that receives audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSinkNode/init(receiverBlock:)
func NewAudioSinkNodeWithReceiverBlock(block AudioSinkNodeReceiverBlock /* not a class type */) AudioSinkNode {
	instance := getAudioSinkNodeClass().Alloc()
	rv := objc.Send[AudioSinkNode](instance.ID, objc.Sel("initWithReceiverBlock:"), block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioSinkNodeWithReceiverBlock */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSinkNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSinkNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSinkNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSinkNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSinkNode */


