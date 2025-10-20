// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioIONode] class.
var (
	AudioIONodeClass     _AudioIONodeClass
	AudioIONodeClassOnce sync.Once
)

func getAudioIONodeClass() _AudioIONodeClass {
	AudioIONodeClassOnce.Do(func() {
		AudioIONodeClass = _AudioIONodeClass{objc.GetClass("AVAudioIONode")}
	})
	return AudioIONodeClass
}

type _AudioIONodeClass struct {
	class objc.Class
}

// An interface definition for the [AudioIONode] class.
type IAudioIONode interface {
	IAudioNode
}

// An object that performs audio input or output in the engine.
//
// When rendering to and from an audio device in macOS, and communicate with the system’s default input and output devices. In iOS, they communicate with the devices appropriate to the app’s category, configurations, and user actions, such as connecting or disconnecting external devices. In the manual rendering mode, and perform the input and output in the engine in response to the client’s request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
type AudioIONode struct {
	AudioNode
}

// AudioIONodeFrom constructs a [AudioIONode] from an unsafe.Pointer.
//
// An object that performs audio input or output in the engine.
func AudioIONodeFrom(ptr unsafe.Pointer) AudioIONode {
	return AudioIONode{
		AudioNode: AudioNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioIONodeClass) Alloc() AudioIONode {
	rv := objc.Send[AudioIONode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioIONodeClass) New() AudioIONode {
	rv := objc.Send[AudioIONode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioIONode) Init() AudioIONode {
	rv := objc.Send[AudioIONode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioIONode) Autorelease() AudioIONode {
	rv := objc.Send[AudioIONode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioIONode creates a new AudioIONode instance.
func NewAudioIONode() AudioIONode {
	return getAudioIONodeClass().New()
}




