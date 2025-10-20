// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AudioOutputNode] class.
type IAudioOutputNode interface {
	IAudioIONode
}

// An object that connects to the system’s audio output.
//
// This node connects to the system’s audio output when rendering to or from an audio device. This node performs output in response to client’s requests when the engine is in manual rendering mode. This audio node has one element. The format of the output scope reflects: The audio hardware sample rate and channel count when it connects to the hardware. The engine’s manual rendering mode output format (see ). The format of the input scope is initially the same as that of the output, but you may set it to a different format, in which case the audio node converts.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioOutputNodeClass) Alloc() AudioOutputNode {
	rv := objc.Send[AudioOutputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




