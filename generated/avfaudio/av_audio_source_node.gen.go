// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AudioSourceNode] class.
type IAudioSourceNode interface {
	IAudioNode
	// properties:
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _AudioSourceNodeClass) Alloc() AudioSourceNode {
	rv := objc.Send[AudioSourceNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




