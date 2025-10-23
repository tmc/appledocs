// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	InputNode() AVAudioInputNode
	MainMixerNode() AVAudioMixerNode
	MusicSequence() unsafe.Pointer
	SetMusicSequence(value unsafe.Pointer)
	OutputNode() AVAudioOutputNode
	AttachedNodes() AVAudioNode
	SetAttachedNodes(value IAVAudioNode)
	IsAutoShutdownEnabled() bool
	SetIsAutoShutdownEnabled(value bool)
	IsInManualRenderingMode() bool
	SetIsInManualRenderingMode(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
	ManualRenderingBlock() unsafe.Pointer
	SetManualRenderingBlock(value unsafe.Pointer)
	ManualRenderingFormat() AVAudioFormat
	SetManualRenderingFormat(value IAVAudioFormat)
	ManualRenderingMaximumFrameCount() AudioFrameCount
	SetManualRenderingMaximumFrameCount(value IAudioFrameCount)
	ManualRenderingMode() unsafe.Pointer
	SetManualRenderingMode(value unsafe.Pointer)
	ManualRenderingSampleTime() AudioFramePosition
	SetManualRenderingSampleTime(value IAudioFramePosition)
}

// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//
// An audio engine object contains a group of instances that you attach to form an audio processing chain. You can connect, disconnect, and remove audio nodes during runtime with minor limitations. Removing an audio node that has differing channel counts, or that’s a mixer, can break the graph. Reconnect audio nodes only when they’re upstream of a mixer. By default, Audio Engine renders to a connected audio device in real time. You can configure the engine to operate in manual rendering mode when you need to render at, or faster than, real time. In that mode, the engine disconnects from audio devices and your app drives the rendering.


// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/prepare()
func (a_ AudioEngine) Prepare() {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepare"))
}


// Starts the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/start()
func (a_ AudioEngine) StartAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startAndReturnError:"), outError)
	return rv
}


// Stops the audio engine and releases any previously prepared resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/stop()
func (a_ AudioEngine) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}


// The audio engine’s singleton input audio node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputNode
func (a_ AudioEngine) InputNode() AVAudioInputNode {
	rv := objc.Send[AVAudioInputNode](a_.ID, objc.Sel("inputNode"))
	return rv
}


// The audio engine’s optional singleton main mixer node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/mainMixerNode
func (a_ AudioEngine) MainMixerNode() AVAudioMixerNode {
	rv := objc.Send[AVAudioMixerNode](a_.ID, objc.Sel("mainMixerNode"))
	return rv
}


// The music sequence instance that you attach to the audio engine, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) MusicSequence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("musicSequence"))
	return rv
}


// The music sequence instance that you attach to the audio engine, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) SetMusicSequence(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMusicSequence:"), value)
}


// The audio engine’s singleton output audio node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputNode
func (a_ AudioEngine) OutputNode() AVAudioOutputNode {
	rv := objc.Send[AVAudioOutputNode](a_.ID, objc.Sel("outputNode"))
	return rv
}


// A read-only set that contains the nodes you attach to the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/attachednodes
func (a_ AudioEngine) AttachedNodes() AVAudioNode {
	rv := objc.Send[AVAudioNode](a_.ID, objc.Sel("attachedNodes"))
	return rv
}


// A read-only set that contains the nodes you attach to the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/attachednodes
func (a_ AudioEngine) SetAttachedNodes(value IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttachedNodes:"), value)
}


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isautoshutdownenabled
func (a_ AudioEngine) IsAutoShutdownEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isAutoShutdownEnabled"))
	return rv
}


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isautoshutdownenabled
func (a_ AudioEngine) SetIsAutoShutdownEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsAutoShutdownEnabled:"), value)
}


// A Boolean value that indicates whether the engine is operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isinmanualrenderingmode
func (a_ AudioEngine) IsInManualRenderingMode() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInManualRenderingMode"))
	return rv
}


// A Boolean value that indicates whether the engine is operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isinmanualrenderingmode
func (a_ AudioEngine) SetIsInManualRenderingMode(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInManualRenderingMode:"), value)
}


// A Boolean value that indicates whether the audio engine is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isrunning
func (a_ AudioEngine) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}


// A Boolean value that indicates whether the audio engine is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isrunning
func (a_ AudioEngine) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}


// The block that renders the engine when operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingblock
func (a_ AudioEngine) ManualRenderingBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("manualRenderingBlock"))
	return rv
}


// The block that renders the engine when operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingblock
func (a_ AudioEngine) SetManualRenderingBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingBlock:"), value)
}


// The render format of the engine in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingformat
func (a_ AudioEngine) ManualRenderingFormat() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a_.ID, objc.Sel("manualRenderingFormat"))
	return rv
}


// The render format of the engine in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingformat
func (a_ AudioEngine) SetManualRenderingFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingFormat:"), value)
}


// The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingmaximumframecount
func (a_ AudioEngine) ManualRenderingMaximumFrameCount() AudioFrameCount {
	rv := objc.Send[AudioFrameCount](a_.ID, objc.Sel("manualRenderingMaximumFrameCount"))
	return rv
}


// The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingmaximumframecount
func (a_ AudioEngine) SetManualRenderingMaximumFrameCount(value IAudioFrameCount) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingMaximumFrameCount:"), value)
}


// The manual rendering mode configured on the engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingmode
func (a_ AudioEngine) ManualRenderingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("manualRenderingMode"))
	return rv
}


// The manual rendering mode configured on the engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingmode
func (a_ AudioEngine) SetManualRenderingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingMode:"), value)
}


// An indication of where the engine is on its render timeline in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingsampletime
func (a_ AudioEngine) ManualRenderingSampleTime() AudioFramePosition {
	rv := objc.Send[AudioFramePosition](a_.ID, objc.Sel("manualRenderingSampleTime"))
	return rv
}


// An indication of where the engine is on its render timeline in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/manualrenderingsampletime
func (a_ AudioEngine) SetManualRenderingSampleTime(value IAudioFramePosition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManualRenderingSampleTime:"), value)
}



