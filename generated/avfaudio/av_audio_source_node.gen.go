// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSourceNode */


/* debug [class_header]: Header for AVAudioSourceNode */
// The class instance for the [AudioSourceNode] class.
var (
	AudioSourceNodeClass     _AudioSourceNodeClass
	AudioSourceNodeClassOnce sync.Once
)

func getAudioSourceNodeClass() _AudioSourceNodeClass {
	AudioSourceNodeClassOnce.Do(func() {
		AudioSourceNodeClass = _AudioSourceNodeClass{objc.GetClass("AVAudioSourceNode")}
	})
	return AudioSourceNodeClass
}

type _AudioSourceNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSourceNode */
// An interface definition for the [AudioSourceNode] class.
type IAudioSourceNode interface {
	IAudioNode
	
/* debug [class_interface_properties]: Properties for AudioSourceNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSourceNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSourceNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioSourceNodeClass) Alloc() AudioSourceNode {
	rv := objc.Send[AudioSourceNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSourceNodeClass) New() AudioSourceNode {
	rv := objc.Send[AudioSourceNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSourceNode) Init() AudioSourceNode {
	rv := objc.Send[AudioSourceNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSourceNode) Autorelease() AudioSourceNode {
	rv := objc.Send[AudioSourceNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSourceNode creates a new AudioSourceNode instance.
func NewAudioSourceNode() AudioSourceNode {
	return getAudioSourceNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSourceNode */
// An object that supplies audio data.
//
// The class allows for supplying audio data for rendering through . It’s a convenient method for delievering audio data instead of setting the input callback on an audio unit with .


// An object that supplies audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode
type AudioSourceNode struct {
	AudioNode
}

// AudioSourceNodeFrom constructs a [AudioSourceNode] from an unsafe.Pointer.
//
// An object that supplies audio data.
func AudioSourceNodeFrom(ptr unsafe.Pointer) AudioSourceNode {
	return AudioSourceNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSourceNode */

// Creates an audio source node with the audio format and a block that supplies audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(format:renderBlock:)
func NewAudioSourceNodeWithFormatRenderBlock(format IAVAudioFormat, block AudioSourceNodeRenderBlock /* not a class type */) AudioSourceNode {
	instance := getAudioSourceNodeClass().Alloc()
	rv := objc.Send[AudioSourceNode](instance.ID, objc.Sel("initWithFormat:renderBlock:"), format, block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioSourceNodeWithFormatRenderBlock */


// Creates an audio source node with a block that supplies audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSourceNode/init(renderBlock:)
func NewAudioSourceNodeWithRenderBlock(block AudioSourceNodeRenderBlock /* not a class type */) AudioSourceNode {
	instance := getAudioSourceNodeClass().Alloc()
	rv := objc.Send[AudioSourceNode](instance.ID, objc.Sel("initWithRenderBlock:"), block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioSourceNodeWithRenderBlock */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSourceNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSourceNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSourceNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSourceNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSourceNode */


