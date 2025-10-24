// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASESoundEvent */


/* debug [class_header]: Header for PHASESoundEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESoundEvent */
// An interface definition for the [PHASESoundEvent] class.
type IPHASESoundEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASESoundEvent */
	// properties:
	Indefinite() bool
	MetaParameters() foundation.IDictionary
	Mixers() foundation.IDictionary
	PrepareState() PHASESoundEventPrepareState
	PullStreamNodes() foundation.IDictionary
	PushStreamNodes() foundation.IDictionary
	RenderingState() PHASERenderingState
	PlaybackMode() PHASEPlaybackMode
	SetPlaybackMode(value PHASEPlaybackMode)
	IsIndefinite() bool
	SetIsIndefinite(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESoundEvent */
	// methods:
	Pause()
	PrepareWithCompletion(handler unsafe.Pointer)
	Resume()
	ResumeAtTime(time avfaudio.AudioTime)
	SeekToTimeCompletion(time float64, handler unsafe.Pointer)
	SeekToTimeResumeAtEngineTimeCompletion(time float64, engineTime avfaudio.AudioTime, handler unsafe.Pointer)
	StartAtTimeCompletion(when avfaudio.AudioTime, handler unsafe.Pointer)
	StartWithCompletion(handler unsafe.Pointer)
	StopAndInvalidate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESoundEvent */
// Alloc allocates a new instance without initialization.
func (pc _PHASESoundEventClass) Alloc() PHASESoundEvent {
	rv := objc.Send[PHASESoundEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESoundEvent */
// An object that determines which audio to play.
//
// A sound event represents a logic tree, or hierarchy, that defines what, when, and how the framework plays a sound at runtime. You configure the tree with conditions based on your app’s state. When you invoke a sound event’s root node at runtime, the framework navigates the tree by branching based on the logic, landing on a playable node that sends the right audio to the output device: To invoke a specific one-time sound, create a sound event from a single sampler node. To invoke a sound event that tailors its sound based on your app’s state, define a sound event hierarchy containing one or more control nodes; see . For example, to play either footsteps or a jumping noise depending on the hero’s state, you configure a switch node that navigates based on the hero’s hypothetical metaparameter. For sound event nodes that play audio, the asset’s determines whether the audio loops. One-time sound events stop automatically at the end of the audio data. Looping sound events (those with ) require you to explicitly call to stop the audio.


// An object that determines which audio to play.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESoundEvent */

// Creates a sound event node with the given asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/init(engine:assetIdentifier:)
func NewPHASESoundEventWithEngineAssetIdentifierError(engine IPHASEEngine, assetIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) PHASESoundEvent {
	instance := getPHASESoundEventClass().Alloc()
	rv := objc.Send[PHASESoundEvent](instance.ID, objc.Sel("initWithEngine:assetIdentifier:error:"), engine, assetIdentifier, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESoundEventWithEngineAssetIdentifierError */


// Creates a sound event node with the given asset and mixer parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/init(engine:assetIdentifier:mixerParameters:)
func NewPHASESoundEventWithEngineAssetIdentifierMixerParametersError(engine IPHASEEngine, assetIdentifier objc.IObject /* cross-framework: NSString */, mixerParameters IPHASEMixerParameters, error_ unsafe.Pointer) PHASESoundEvent {
	instance := getPHASESoundEventClass().Alloc()
	rv := objc.Send[PHASESoundEvent](instance.ID, objc.Sel("initWithEngine:assetIdentifier:mixerParameters:error:"), engine, assetIdentifier, mixerParameters, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESoundEventWithEngineAssetIdentifierMixerParametersError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESoundEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESoundEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESoundEvent */

// Pauses the sound event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pause()
func (p_ PHASESoundEvent) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Enables a sound event to play and runs the argument code when the sound event plays back.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/prepare(completion:)
func (p_ PHASESoundEvent) PrepareWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("prepareWithCompletion:"), handler)
}/* debug [instance_methods/method]: PrepareWithCompletion */


// Resumes the sound event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/resume()
func (p_ PHASESoundEvent) Resume() {
	objc.Send[objc.ID](p_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/resume(at:)
func (p_ PHASESoundEvent) ResumeAtTime(time avfaudio.AudioTime) {
	objc.Send[objc.ID](p_.ID, objc.Sel("resumeAtTime:"), time)
}/* debug [instance_methods/method]: ResumeAtTime */


// Advances the sound event’s playback position to a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/seek(to:completion:)
func (p_ PHASESoundEvent) SeekToTimeCompletion(time float64, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:completion:"), time, handler)
}/* debug [instance_methods/method]: SeekToTimeCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/seek(to:resumeAt:completion:)
func (p_ PHASESoundEvent) SeekToTimeResumeAtEngineTimeCompletion(time float64, engineTime avfaudio.AudioTime, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:resumeAtEngineTime:completion:"), time, engineTime, handler)
}/* debug [instance_methods/method]: SeekToTimeResumeAtEngineTimeCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/start(at:completion:)
func (p_ PHASESoundEvent) StartAtTimeCompletion(when avfaudio.AudioTime, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startAtTime:completion:"), when, handler)
}/* debug [instance_methods/method]: StartAtTimeCompletion */


// Invokes the sound event and runs the specified code on completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/start(completion:)
func (p_ PHASESoundEvent) StartWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startWithCompletion:"), handler)
}/* debug [instance_methods/method]: StartWithCompletion */


// Stops a sound event and prevents it from resuming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/stopAndInvalidate()
func (p_ PHASESoundEvent) StopAndInvalidate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopAndInvalidate"))
}/* debug [instance_methods/method]: StopAndInvalidate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESoundEvent */

// A Boolean value that indicates whether the sound loops or stops on its own.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/isIndefinite
func (p_ PHASESoundEvent) Indefinite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indefinite"))
	return rv
}/* debug [instance_properties/getter]: indefinite */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/metaParameters
func (p_ PHASESoundEvent) MetaParameters() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// Nodes in the event tree that control the volume of their child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/mixers
func (p_ PHASESoundEvent) Mixers() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("mixers"))
	return rv
}/* debug [instance_properties/getter]: mixers */


// The status of sound-event preparation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/prepareState-swift.property
func (p_ PHASESoundEvent) PrepareState() PHASESoundEventPrepareState {
	rv := objc.Send[PHASESoundEventPrepareState](p_.ID, objc.Sel("prepareState"))
	return rv
}/* debug [instance_properties/getter]: prepareState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pullStreamNodes
func (p_ PHASESoundEvent) PullStreamNodes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("pullStreamNodes"))
	return rv
}/* debug [instance_properties/getter]: pullStreamNodes */


// A collection of audio streams for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/pushStreamNodes
func (p_ PHASESoundEvent) PushStreamNodes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("pushStreamNodes"))
	return rv
}/* debug [instance_properties/getter]: pushStreamNodes */


// The sound event’s playback status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/renderingState-swift.property
func (p_ PHASESoundEvent) RenderingState() PHASERenderingState {
	rv := objc.Send[PHASERenderingState](p_.ID, objc.Sel("renderingState"))
	return rv
}/* debug [instance_properties/getter]: renderingState */


// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/playbackmode
func (p_ PHASESoundEvent) PlaybackMode() PHASEPlaybackMode {
	rv := objc.Send[PHASEPlaybackMode](p_.ID, objc.Sel("playbackMode"))
	return rv
}/* debug [instance_properties/getter]: playbackMode */


// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/playbackmode
func (p_ PHASESoundEvent) SetPlaybackMode(value PHASEPlaybackMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackMode:"), value)
}/* debug [instance_properties/setter]: playbackMode */


// A Boolean value that indicates whether the sound loops or stops on its own.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/isindefinite
func (p_ PHASESoundEvent) IsIndefinite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndefinite"))
	return rv
}/* debug [instance_properties/getter]: isIndefinite */


// A Boolean value that indicates whether the sound loops or stops on its own.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/isindefinite
func (p_ PHASESoundEvent) SetIsIndefinite(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndefinite:"), value)
}/* debug [instance_properties/setter]: isIndefinite */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESoundEvent */


