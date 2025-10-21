// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASESoundEvent] class.
var (
	PHASESoundEventClass     _PHASESoundEventClass
	PHASESoundEventClassOnce sync.Once
)

func getPHASESoundEventClass() _PHASESoundEventClass {
	PHASESoundEventClassOnce.Do(func() {
		PHASESoundEventClass = _PHASESoundEventClass{objc.GetClass("PHASESoundEvent")}
	})
	return PHASESoundEventClass
}

type _PHASESoundEventClass struct {
	class objc.Class
}

// An interface definition for the [PHASESoundEvent] class.
type IPHASESoundEvent interface {
	objectivec.IObject
	Pause()
	PrepareWithCompletion(handler unsafe.Pointer)
	Resume()
	ResumeAtTime(time avfaudio.IAudioTime)
	SeekToTimeCompletion(time unsafe.Pointer, handler unsafe.Pointer)
	SeekToTimeResumeAtEngineTimeCompletion(time unsafe.Pointer, engineTime avfaudio.IAudioTime, handler unsafe.Pointer)
	StartAtTimeCompletion(when avfaudio.IAudioTime, handler unsafe.Pointer)
	StartWithCompletion(handler unsafe.Pointer)
	StopAndInvalidate()
}

// An object that determines which audio to play.
//
// A sound event represents a logic tree, or hierarchy, that defines what, when, and how the framework plays a sound at runtime. You configure the tree with conditions based on your app’s state. When you invoke a sound event’s root node at runtime, the framework navigates the tree by branching based on the logic, landing on a playable node that sends the right audio to the output device: To invoke a specific one-time sound, create a sound event from a single sampler node. To invoke a sound event that tailors its sound based on your app’s state, define a sound event hierarchy containing one or more control nodes; see . For example, to play either footsteps or a jumping noise depending on the hero’s state, you configure a switch node that navigates based on the hero’s hypothetical metaparameter. For sound event nodes that play audio, the asset’s determines whether the audio loops. One-time sound events stop automatically at the end of the audio data. Looping sound events (those with ) require you to explicitly call to stop the audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent
type PHASESoundEvent struct {
	objectivec.Object
}

// PHASESoundEventFrom constructs a [PHASESoundEvent] from an unsafe.Pointer.
//
// An object that determines which audio to play.
func PHASESoundEventFrom(ptr unsafe.Pointer) PHASESoundEvent {
	return PHASESoundEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESoundEventClass) Alloc() PHASESoundEvent {
	rv := objc.Send[PHASESoundEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESoundEventClass) New() PHASESoundEvent {
	rv := objc.Send[PHASESoundEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESoundEvent) Init() PHASESoundEvent {
	rv := objc.Send[PHASESoundEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESoundEvent) Autorelease() PHASESoundEvent {
	rv := objc.Send[PHASESoundEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESoundEvent creates a new PHASESoundEvent instance.
func NewPHASESoundEvent() PHASESoundEvent {
	return getPHASESoundEventClass().New()
}




// Creates a sound event node with the given asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/init(engine:assetIdentifier:)
func NewPHASESoundEventWithEngineAssetIdentifierError(engine IPHASEEngine, assetIdentifier appkit.string, error_ unsafe.Pointer) PHASESoundEvent {
	instance := getPHASESoundEventClass().Alloc()
	rv := objc.Send[PHASESoundEvent](instance.ID, objc.Sel("initWithEngine:assetIdentifier:error:"), engine, assetIdentifier, error_)
	rv.Autorelease()
	return rv
}



// Creates a sound event node with the given asset and mixer parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/init(engine:assetIdentifier:mixerParameters:)
func NewPHASESoundEventWithEngineAssetIdentifierMixerParametersError(engine IPHASEEngine, assetIdentifier appkit.string, mixerParameters IPHASEMixerParameters, error_ unsafe.Pointer) PHASESoundEvent {
	instance := getPHASESoundEventClass().Alloc()
	rv := objc.Send[PHASESoundEvent](instance.ID, objc.Sel("initWithEngine:assetIdentifier:mixerParameters:error:"), engine, assetIdentifier, mixerParameters, error_)
	rv.Autorelease()
	return rv
}


// Pauses the sound event.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pause()
func (p_ PHASESoundEvent) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}

// Enables a sound event to play and runs the argument code when the sound event plays back.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/prepare(completion:)
func (p_ PHASESoundEvent) PrepareWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("prepareWithCompletion:"), handler)
}

// Resumes the sound event.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/resume()
func (p_ PHASESoundEvent) Resume() {
	objc.Send[objc.ID](p_.ID, objc.Sel("resume"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/resume(at:)
func (p_ PHASESoundEvent) ResumeAtTime(time avfaudio.IAudioTime) {
	objc.Send[objc.ID](p_.ID, objc.Sel("resumeAtTime:"), time)
}

// Advances the sound event’s playback position to a specific time.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/seek(to:completion:)
func (p_ PHASESoundEvent) SeekToTimeCompletion(time unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:completion:"), time, handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/seek(to:resumeAt:completion:)
func (p_ PHASESoundEvent) SeekToTimeResumeAtEngineTimeCompletion(time unsafe.Pointer, engineTime avfaudio.IAudioTime, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:resumeAtEngineTime:completion:"), time, engineTime, handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/start(at:completion:)
func (p_ PHASESoundEvent) StartAtTimeCompletion(when avfaudio.IAudioTime, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAtTime:completion:"), when, handler)
}

// Invokes the sound event and runs the specified code on completion.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/start(completion:)
func (p_ PHASESoundEvent) StartWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startWithCompletion:"), handler)
}

// Stops a sound event and prevents it from resuming.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/stopAndInvalidate()
func (p_ PHASESoundEvent) StopAndInvalidate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAndInvalidate"))
}

// A Boolean value that indicates whether the sound loops or stops on its own.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/isIndefinite
func (p_ PHASESoundEvent) Indefinite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indefinite"))
	return rv
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/metaParameters
func (p_ PHASESoundEvent) MetaParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("metaParameters"))
	return rv
}

// Nodes in the event tree that control the volume of their child nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/mixers
func (p_ PHASESoundEvent) Mixers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mixers"))
	return rv
}

// The status of sound-event preparation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/prepareState-swift.property
func (p_ PHASESoundEvent) PrepareState() PHASESoundEventPrepareState {
	rv := objc.Send[PHASESoundEventPrepareState](p_.ID, objc.Sel("prepareState"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pullStreamNodes
func (p_ PHASESoundEvent) PullStreamNodes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pullStreamNodes"))
	return rv
}

// A collection of audio streams for playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pushStreamNodes
func (p_ PHASESoundEvent) PushStreamNodes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pushStreamNodes"))
	return rv
}

// The sound event’s playback status.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/renderingState-swift.property
func (p_ PHASESoundEvent) RenderingState() PHASERenderingState {
	rv := objc.Send[PHASERenderingState](p_.ID, objc.Sel("renderingState"))
	return rv
}

// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/playbackmode
func (p_ PHASESoundEvent) PlaybackMode() PHASEPlaybackMode {
	rv := objc.Send[PHASEPlaybackMode](p_.ID, objc.Sel("playbackMode"))
	return rv
}


// SetPlaybackMode sets the value of the playbackMode property.
// An option that determines whether the node’s audio plays in a loop.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/playbackmode
func (p_ PHASESoundEvent) SetPlaybackMode(value PHASEPlaybackMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackMode:"), value)
}

// A Boolean value that indicates whether the sound loops or stops on its own.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/isindefinite
func (p_ PHASESoundEvent) IsIndefinite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndefinite"))
	return rv
}


// SetIsIndefinite sets the value of the isIndefinite property.
// A Boolean value that indicates whether the sound loops or stops on its own.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/isindefinite
func (p_ PHASESoundEvent) SetIsIndefinite(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndefinite:"), value)
}


