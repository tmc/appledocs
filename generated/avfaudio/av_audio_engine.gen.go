// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioEngine] class.
var (
	AudioEngineClass     _AudioEngineClass
	AudioEngineClassOnce sync.Once
)

func getAudioEngineClass() _AudioEngineClass {
	AudioEngineClassOnce.Do(func() {
		AudioEngineClass = _AudioEngineClass{objc.GetClass("AVAudioEngine")}
	})
	return AudioEngineClass
}

type _AudioEngineClass struct {
	class objc.Class
}

// An interface definition for the [AudioEngine] class.
type IAudioEngine interface {
	objectivec.IObject
	Prepare()
	StartAndReturnError(outError unsafe.Pointer) bool
	Stop()
}

// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//
// An audio engine object contains a group of instances that you attach to form an audio processing chain. You can connect, disconnect, and remove audio nodes during runtime with minor limitations. Removing an audio node that has differing channel counts, or that’s a mixer, can break the graph. Reconnect audio nodes only when they’re upstream of a mixer. By default, Audio Engine renders to a connected audio device in real time. You can configure the engine to operate in manual rendering mode when you need to render at, or faster than, real time. In that mode, the engine disconnects from audio devices and your app drives the rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type AudioEngine struct {
	objectivec.Object
}

// AudioEngineFrom constructs a [AudioEngine] from an unsafe.Pointer.
//
// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
func AudioEngineFrom(ptr unsafe.Pointer) AudioEngine {
	return AudioEngine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioEngineClass) Alloc() AudioEngine {
	rv := objc.Send[AudioEngine](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioEngineClass) New() AudioEngine {
	rv := objc.Send[AudioEngine](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEngine) Init() AudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEngine) Autorelease() AudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEngine creates a new AudioEngine instance.
func NewAudioEngine() AudioEngine {
	return getAudioEngineClass().New()
}


// Prepares the audio engine for starting.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/prepare()
func (a_ AudioEngine) Prepare() {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepare"))
}

// Starts the audio engine.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/start()
func (a_ AudioEngine) StartAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startAndReturnError:"), outError)
	return rv
}

// Stops the audio engine and releases any previously prepared resources.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/stop()
func (a_ AudioEngine) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}

// The audio engine’s singleton input audio node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputNode
func (a_ AudioEngine) InputNode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputNode"))
	return rv
}

// The audio engine’s optional singleton main mixer node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/mainMixerNode
func (a_ AudioEngine) MainMixerNode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mainMixerNode"))
	return rv
}

// The music sequence instance that you attach to the audio engine, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) MusicSequence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("musicSequence"))
	return rv
}


// SetMusicSequence sets the value of the musicSequence property.
// The music sequence instance that you attach to the audio engine, if any.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) SetMusicSequence(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMusicSequence:"), value)
}
// The audio engine’s singleton output audio node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputNode
func (a_ AudioEngine) OutputNode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputNode"))
	return rv
}



