// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioMixerNode] class.
var (
	AudioMixerNodeClass     _AudioMixerNodeClass
	AudioMixerNodeClassOnce sync.Once
)

func getAudioMixerNodeClass() _AudioMixerNodeClass {
	AudioMixerNodeClassOnce.Do(func() {
		AudioMixerNodeClass = _AudioMixerNodeClass{objc.GetClass("AVAudioMixerNode")}
	})
	return AudioMixerNodeClass
}

type _AudioMixerNodeClass struct {
	class objc.Class
}





// An interface definition for the [AudioMixerNode] class.
type IAudioMixerNode interface {
	IAudioNode
	

	// properties:
	NextAvailableInputBus() AudioNodeBus /* typedef */
	OutputVolume() float32
	SetOutputVolume(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioMixerNodeClass) Alloc() AudioMixerNode {
	rv := objc.Send[AudioMixerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioMixerNodeClass) New() AudioMixerNode {
	rv := objc.Send[AudioMixerNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioMixerNode) Init() AudioMixerNode {
	rv := objc.Send[AudioMixerNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioMixerNode) Autorelease() AudioMixerNode {
	rv := objc.Send[AudioMixerNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioMixerNode creates a new AudioMixerNode instance.
func NewAudioMixerNode() AudioMixerNode {
	return getAudioMixerNodeClass().New()
}





// An object that takes any number of inputs and converts them into a single output.
//
// The mixer accepts input at any sample rate and efficiently combines sample rate conversions. It also accepts any channel count and correctly upmixes or downmixes to the output channel count.


// An object that takes any number of inputs and converts them into a single output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode
type AudioMixerNode struct {
	AudioNode
}

// AudioMixerNodeFrom constructs a [AudioMixerNode] from an unsafe.Pointer.
//
// An object that takes any number of inputs and converts them into a single output.
func AudioMixerNodeFrom(ptr unsafe.Pointer) AudioMixerNode {
	return AudioMixerNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}


























// An audio bus that isn’t in a connected state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode/nextAvailableInputBus
func (a_ AudioMixerNode) NextAvailableInputBus() AudioNodeBus /* typedef */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("nextAvailableInputBus"))
	return rv
}


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode/outputVolume
func (a_ AudioMixerNode) OutputVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("outputVolume"))
	return rv
}


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixerNode/outputVolume
func (a_ AudioMixerNode) SetOutputVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputVolume:"), value)
}







