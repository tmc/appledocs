// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AudioSinkNode] class.
type IAudioSinkNode interface {
	IAudioNode
}

// An object that receives audio data.
//
// You use an to receive audio data through . You only use it in the input chain. An audio sink node doesn’t support format conversion. When connecting, use the output format of the input for the format for the connection. The format should match the hardware input sample rate. The voice processing I/O unit is an exception to the above because it supports sample rate conversion. The input scope format (hardware format) and output scope format (client format) of the input node can differ in that case. An audio sink node doesn’t support manual rendering mode, and doesn’t have an output bus, so you can’t install a tap on it.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioSinkNodeClass) Alloc() AudioSinkNode {
	rv := objc.Send[AudioSinkNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




