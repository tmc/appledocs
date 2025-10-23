// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSession] class.
var (
	AudioSessionClass     _AudioSessionClass
	AudioSessionClassOnce sync.Once
)

func getAudioSessionClass() _AudioSessionClass {
	AudioSessionClassOnce.Do(func() {
		AudioSessionClass = _AudioSessionClass{objc.GetClass("AVAudioSession")}
	})
	return AudioSessionClass
}

type _AudioSessionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSession] class.
type IAudioSession interface {
	objectivec.IObject
	ActivateWithOptionsCompletionHandler(options AudioSessionActivationOptions, handler unsafe.Pointer)
	PrepareRouteSelectionForPlaybackWithCompletionHandler(completionHandler unsafe.Pointer)
	RequestRecordPermission(response unsafe.Pointer)
	SetActiveWithOptionsError(active bool, options AudioSessionSetActiveOptions, outError unsafe.Pointer) bool
	SetActiveWithFlagsError(active bool, flags int, outError unsafe.Pointer) bool
	SetActiveError(active bool, outError unsafe.Pointer) bool
	SetAggregatedIOPreferenceError(inIOType AudioSessionIOType, outError unsafe.Pointer) bool
	SetAllowHapticsAndSystemSoundsDuringRecordingError(inValue bool, outError unsafe.Pointer) bool
	SetCategoryError(category IAudioSessionCategory, outError unsafe.Pointer) bool
	SetCategoryModeOptionsError(category IAudioSessionCategory, mode AudioSessionMode, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool
	SetCategoryModeRouteSharingPolicyOptionsError(category IAudioSessionCategory, mode AudioSessionMode, policy AudioSessionRouteSharingPolicy, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool
	SetCategoryWithOptionsError(category IAudioSessionCategory, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool
	SetIntendedSpatialExperienceOptionsError(intendedSpatialExperience IAudioSessionSpatialExperience, options unsafe.Pointer, error_ unsafe.Pointer) bool
	SetIsNowPlayingCandidateError(inValue bool, outError unsafe.Pointer) bool
	SetModeError(mode AudioSessionMode, outError unsafe.Pointer) bool
	SetOutputMutedError(muted bool, outError unsafe.Pointer) bool
	SetPreferredInputError(inPort IAVAudioSessionPortDescription, outError unsafe.Pointer) bool
	SetPreferredInputOrientationError(orientation IAudioStereoOrientation, outError unsafe.Pointer) bool
	SetPreferredMicrophoneInjectionModeError(inValue AudioSessionMicrophoneInjectionMode, outError unsafe.Pointer) bool
	SetPrefersEchoCancelledInputError(value bool, error_ unsafe.Pointer) bool
	SetPrefersInterruptionOnRouteDisconnectError(inValue bool, outError unsafe.Pointer) bool
	SetPrefersNoInterruptionsFromSystemAlertsError(inValue bool, outError unsafe.Pointer) bool
	AllowHapticsAndSystemSoundsDuringRecording() bool
	AvailableCategories() []string
	AvailableInputs() []AudioSessionPortDescription
	AvailableModes() []string
	Category() AudioSessionCategory
	CategoryOptions() AudioSessionCategoryOptions
	CurrentRoute() AVAudioSessionRouteDescription
	InputDataSource() AVAudioSessionDataSourceDescription
	InputOrientation() AudioStereoOrientation
	IntendedSpatialExperience() AudioSessionSpatialExperience
	IntendedSpatialExperienceOptions() unsafe.Pointer
	IsEchoCancelledInputAvailable() bool
	IsEchoCancelledInputEnabled() bool
	InputAvailable() bool
	IsMicrophoneInjectionAvailable() bool
	IsNowPlayingCandidate() bool
	OtherAudioPlaying() bool
	OutputMuted() bool
	Mode() AudioSessionMode
	PreferredInputOrientation() AudioStereoOrientation
	PreferredMicrophoneInjectionMode() AudioSessionMicrophoneInjectionMode
	PrefersEchoCancelledInput() bool
	PrefersInterruptionOnRouteDisconnect() bool
	PrefersNoInterruptionsFromSystemAlerts() bool
	PromptStyle() AudioSessionPromptStyle
	RenderingMode() AudioSessionRenderingMode
	RouteSharingPolicy() AudioSessionRouteSharingPolicy
	SecondaryAudioShouldBeSilencedHint() bool
	SupportedOutputChannelLayouts() []AudioChannelLayout
	IsOtherAudioPlaying() bool
	SetIsOtherAudioPlaying(value bool)
	IsOutputMuted() bool
	SetIsOutputMuted(value bool)
}

// An object that communicates to the system how you intend to use audio in your app.
//
// An audio session acts as an intermediary between your app and the operating system — and, in turn, the underlying audio hardware. You use an audio session to communicate to the operating system the general nature of your app’s audio without detailing the specific behavior or required interactions with the audio hardware. You delegate the management of those details to the audio session, which ensures that the operating system can best manage the user’s audio experience. All iOS, tvOS, and watchOS apps have a default audio session that comes preconfigured with the following behavior: It supports audio playback, but disallows audio recording. When the app plays audio, it silences any other background audio. In iOS, setting the Ring/Silent switch to silent mode silences any audio the app is playing. In iOS, locking a device silences the app’s audio. Although the default audio session provides useful behavior, it generally doesn’t provide the audio behavior a media app needs. To change the default behavior, you configure your app’s audio session category. There are six possible categories you can use, but is the one that playback apps most commonly use. This category indicates that audio playback is a central feature of your app. When you specify this category, your app’s audio continues with the Ring/Silent switch set to silent mode (iOS only). Using this category, you can also play background audio if you’re using the Audio, AirPlay, and Picture in Picture background mode. For more information, see . You use an object to configure your app’s audio session. This class is a singleton object used to set the audio session’s category, mode, and other configurations. You can interact with the audio session throughout your app’s life cycle, but it’s often useful to perform this configuration at app launch, as shown in the following example. The audio session uses this configuration when you activate the session using the or method.


// An object that communicates to the system how you intend to use audio in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
type AudioSession struct {
	objectivec.Object
}

// AudioSessionFrom constructs a [AudioSession] from an unsafe.Pointer.
//
// An object that communicates to the system how you intend to use audio in your app.
func AudioSessionFrom(ptr unsafe.Pointer) AudioSession {
	return AudioSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionClass) Alloc() AudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionClass) New() AudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSession) Init() AudioSession {
	rv := objc.Send[AudioSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSession) Autorelease() AudioSession {
	rv := objc.Send[AudioSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSession creates a new AudioSession instance.
func NewAudioSession() AudioSession {
	return getAudioSessionClass().New()
}



// Returns the shared audio session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/sharedInstance()
func (ac _AudioSessionClass) SharedInstance() AudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}


// Activates an audio session asynchronously on watchOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/activate(options:completionHandler:)
func (a_ AudioSession) ActivateWithOptionsCompletionHandler(options AudioSessionActivationOptions, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("activateWithOptions:completionHandler:"), options, handler)
}


// Prepares the route selection for long-form video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/prepareRouteSelectionForPlayback(completionHandler:)
func (a_ AudioSession) PrepareRouteSelectionForPlaybackWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareRouteSelectionForPlaybackWithCompletionHandler:"), completionHandler)
}


// Requests the user’s permission to record audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/requestRecordPermission(_:)
func (a_ AudioSession) RequestRecordPermission(response unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestRecordPermission:"), response)
}


// Activates or deactivates your app’s audio session using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive(_:options:)
func (a_ AudioSession) SetActiveWithOptionsError(active bool, options AudioSessionSetActiveOptions, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActive:withOptions:error:"), active, options, outError)
	return rv
}


// Activates or deactivates your app’s audio session; provides flags for use by other audio sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive(_:withFlags:)
func (a_ AudioSession) SetActiveWithFlagsError(active bool, flags int, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActive:withFlags:error:"), active, flags, outError)
	return rv
}


// Activates or deactivates your app’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive:error:
func (a_ AudioSession) SetActiveError(active bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActive:error:"), active, outError)
	return rv
}


// Sets the audio session’s aggregated I/O configuration preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setAggregatedIOPreference(_:)
func (a_ AudioSession) SetAggregatedIOPreferenceError(inIOType AudioSessionIOType, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setAggregatedIOPreference:error:"), inIOType, outError)
	return rv
}


// Sets a Boolean value that indicates whether system sounds and haptics play while recording from audio input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setAllowHapticsAndSystemSoundsDuringRecording(_:)
func (a_ AudioSession) SetAllowHapticsAndSystemSoundsDuringRecordingError(inValue bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setAllowHapticsAndSystemSoundsDuringRecording:error:"), inValue, outError)
	return rv
}


// Sets the audio session’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:)
func (a_ AudioSession) SetCategoryError(category IAudioSessionCategory, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:error:"), category, outError)
	return rv
}


// Sets the audio session’s category, mode, and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:mode:options:)
func (a_ AudioSession) SetCategoryModeOptionsError(category IAudioSessionCategory, mode AudioSessionMode, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:mode:options:error:"), category, mode, options, outError)
	return rv
}


// Sets the session category, mode, route-sharing policy, and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:mode:policy:options:)
func (a_ AudioSession) SetCategoryModeRouteSharingPolicyOptionsError(category IAudioSessionCategory, mode AudioSessionMode, policy AudioSessionRouteSharingPolicy, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:mode:routeSharingPolicy:options:error:"), category, mode, policy, options, outError)
	return rv
}


// Sets the audio session’s category with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:options:)
func (a_ AudioSession) SetCategoryWithOptionsError(category IAudioSessionCategory, options AudioSessionCategoryOptions, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:withOptions:error:"), category, options, outError)
	return rv
}


// Sets the spatial audio experience your app intends to provide the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setIntendedSpatialExperience:options:error:
func (a_ AudioSession) SetIntendedSpatialExperienceOptionsError(intendedSpatialExperience IAudioSessionSpatialExperience, options unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setIntendedSpatialExperience:options:error:"), intendedSpatialExperience, options, error_)
	return rv
}


// Sets a Boolean value that indicates whether the audio session is a candidate to be the Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setIsNowPlayingCandidate(_:)
func (a_ AudioSession) SetIsNowPlayingCandidateError(inValue bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setIsNowPlayingCandidate:error:"), inValue, outError)
	return rv
}


// Sets the audio session’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setMode(_:)
func (a_ AudioSession) SetModeError(mode AudioSessionMode, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setMode:error:"), mode, outError)
	return rv
}


// Sets a Boolean value to inform the system to mute the session’s output audio. The default value is false (unmuted).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setOutputMuted(_:)
func (a_ AudioSession) SetOutputMutedError(muted bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setOutputMuted:error:"), muted, outError)
	return rv
}


// Sets the preferred input port for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredInput(_:)
func (a_ AudioSession) SetPreferredInputError(inPort IAVAudioSessionPortDescription, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredInput:error:"), inPort, outError)
	return rv
}


// Sets the audio session’s preferred stereo input orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredInputOrientation(_:)
func (a_ AudioSession) SetPreferredInputOrientationError(orientation IAudioStereoOrientation, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredInputOrientation:error:"), orientation, outError)
	return rv
}


// Sets the preferred mode of injecting audio into another app’s input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredMicrophoneInjectionMode(_:)
func (a_ AudioSession) SetPreferredMicrophoneInjectionModeError(inValue AudioSessionMicrophoneInjectionMode, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredMicrophoneInjectionMode:error:"), inValue, outError)
	return rv
}


// Sets a preference to enable echo-canceled input on supported hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersEchoCancelledInput(_:)
func (a_ AudioSession) SetPrefersEchoCancelledInputError(value bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersEchoCancelledInput:error:"), value, error_)
	return rv
}


// Sets a preference to interrupt the audio session when the active route disconnects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersInterruptionOnRouteDisconnect(_:)
func (a_ AudioSession) SetPrefersInterruptionOnRouteDisconnectError(inValue bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersInterruptionOnRouteDisconnect:error:"), inValue, outError)
	return rv
}


// Sets the preference for not interrupting the audio session with system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersNoInterruptionsFromSystemAlerts(_:)
func (a_ AudioSession) SetPrefersNoInterruptionsFromSystemAlertsError(inValue bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersNoInterruptionsFromSystemAlerts:error:"), inValue, outError)
	return rv
}


// A Boolean value that indicates whether system sounds and haptics play while recording from audio input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/allowHapticsAndSystemSoundsDuringRecording
func (a_ AudioSession) AllowHapticsAndSystemSoundsDuringRecording() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowHapticsAndSystemSoundsDuringRecording"))
	return rv
}


// The audio session categories available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/availableCategories
func (a_ AudioSession) AvailableCategories() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableCategories"))
	return rv
}


// An array of input ports available for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/availableInputs
func (a_ AudioSession) AvailableInputs() []AudioSessionPortDescription {
	rv := objc.Send[[]AudioSessionPortDescription](a_.ID, objc.Sel("availableInputs"))
	return rv
}


// The audio session modes available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/availableModes
func (a_ AudioSession) AvailableModes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableModes"))
	return rv
}


// The current audio session category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/category-swift.property
func (a_ AudioSession) Category() AudioSessionCategory {
	rv := objc.Send[AudioSessionCategory](a_.ID, objc.Sel("category"))
	return rv
}


// The set of options associated with the current audio session category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/categoryOptions-swift.property
func (a_ AudioSession) CategoryOptions() AudioSessionCategoryOptions {
	rv := objc.Send[AudioSessionCategoryOptions](a_.ID, objc.Sel("categoryOptions"))
	return rv
}


// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/currentRoute
func (a_ AudioSession) CurrentRoute() AVAudioSessionRouteDescription {
	rv := objc.Send[AVAudioSessionRouteDescription](a_.ID, objc.Sel("currentRoute"))
	return rv
}


// The currently selected input data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputDataSource
func (a_ AudioSession) InputDataSource() AVAudioSessionDataSourceDescription {
	rv := objc.Send[AVAudioSessionDataSourceDescription](a_.ID, objc.Sel("inputDataSource"))
	return rv
}


// An orientation value that dictates which directions represent left and right when capturing audio from a built-in microphone configured for stereo recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputOrientation
func (a_ AudioSession) InputOrientation() AudioStereoOrientation {
	rv := objc.Send[AudioStereoOrientation](a_.ID, objc.Sel("inputOrientation"))
	return rv
}


// The spatial audio experience your app intends to provide the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/intendedSpatialExperience-qlty
func (a_ AudioSession) IntendedSpatialExperience() AudioSessionSpatialExperience {
	rv := objc.Send[AudioSessionSpatialExperience](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}


// A dictionary of options that customize the spatial experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/intendedSpatialExperienceOptions
func (a_ AudioSession) IntendedSpatialExperienceOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("intendedSpatialExperienceOptions"))
	return rv
}


// A Boolean value that indicates whether the built-in microphone and speaker route supports echo cancellation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isEchoCancelledInputAvailable
func (a_ AudioSession) IsEchoCancelledInputAvailable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEchoCancelledInputAvailable"))
	return rv
}


// A Boolean value that indicates whether an echo-canceled input is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isEchoCancelledInputEnabled
func (a_ AudioSession) IsEchoCancelledInputEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEchoCancelledInputEnabled"))
	return rv
}


// A Boolean value that indicates whether an audio input path is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isInputAvailable
func (a_ AudioSession) InputAvailable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputAvailable"))
	return rv
}


// A Boolean value that indicates whether microphone injection is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isMicrophoneInjectionAvailable
func (a_ AudioSession) IsMicrophoneInjectionAvailable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isMicrophoneInjectionAvailable"))
	return rv
}


// A Boolean value that indicates whether the audio session is a candidate to be the Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isNowPlayingCandidate
func (a_ AudioSession) IsNowPlayingCandidate() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isNowPlayingCandidate"))
	return rv
}


// A Boolean value that indicates whether another app is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isOtherAudioPlaying
func (a_ AudioSession) OtherAudioPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("otherAudioPlaying"))
	return rv
}


// A Boolean value that indicates whether audio output is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isOutputMuted
func (a_ AudioSession) OutputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("outputMuted"))
	return rv
}


// The current audio session’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/mode-swift.property
func (a_ AudioSession) Mode() AudioSessionMode {
	rv := objc.Send[AudioSessionMode](a_.ID, objc.Sel("mode"))
	return rv
}


// The audio session’s preferred stereo input orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredInputOrientation
func (a_ AudioSession) PreferredInputOrientation() AudioStereoOrientation {
	rv := objc.Send[AudioStereoOrientation](a_.ID, objc.Sel("preferredInputOrientation"))
	return rv
}


// The preferred mode of injecting audio into another app’s input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredMicrophoneInjectionMode
func (a_ AudioSession) PreferredMicrophoneInjectionMode() AudioSessionMicrophoneInjectionMode {
	rv := objc.Send[AudioSessionMicrophoneInjectionMode](a_.ID, objc.Sel("preferredMicrophoneInjectionMode"))
	return rv
}


// A Boolean value that indicates the audio session’s preference for using an echo-canceled input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/prefersEchoCancelledInput
func (a_ AudioSession) PrefersEchoCancelledInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersEchoCancelledInput"))
	return rv
}


// A Boolean value that indicates whether the system interrupts the audio session when the active route disconnects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/prefersInterruptionOnRouteDisconnect
func (a_ AudioSession) PrefersInterruptionOnRouteDisconnect() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersInterruptionOnRouteDisconnect"))
	return rv
}


// A Boolean value that indicates a preference for not interrupting the session with system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/prefersNoInterruptionsFromSystemAlerts
func (a_ AudioSession) PrefersNoInterruptionsFromSystemAlerts() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("prefersNoInterruptionsFromSystemAlerts"))
	return rv
}


// A hint to audio sessions that use voice prompt mode to alter the type of prompts they issue in response to other system audio, such as Siri and phone calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/promptStyle-swift.property
func (a_ AudioSession) PromptStyle() AudioSessionPromptStyle {
	rv := objc.Send[AudioSessionPromptStyle](a_.ID, objc.Sel("promptStyle"))
	return rv
}


// The current audio session’s rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/renderingMode-swift.property
func (a_ AudioSession) RenderingMode() AudioSessionRenderingMode {
	rv := objc.Send[AudioSessionRenderingMode](a_.ID, objc.Sel("renderingMode"))
	return rv
}


// The active route-sharing policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/routeSharingPolicy-swift.property
func (a_ AudioSession) RouteSharingPolicy() AudioSessionRouteSharingPolicy {
	rv := objc.Send[AudioSessionRouteSharingPolicy](a_.ID, objc.Sel("routeSharingPolicy"))
	return rv
}


// A Boolean value that indicates whether another app, with a nonmixable audio session, is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/secondaryAudioShouldBeSilencedHint
func (a_ AudioSession) SecondaryAudioShouldBeSilencedHint() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("secondaryAudioShouldBeSilencedHint"))
	return rv
}


// The array of channel layouts that the current route supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/supportedOutputChannelLayouts
func (a_ AudioSession) SupportedOutputChannelLayouts() []AudioChannelLayout {
	rv := objc.Send[[]AudioChannelLayout](a_.ID, objc.Sel("supportedOutputChannelLayouts"))
	return rv
}


// A Boolean value that indicates whether another app is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isotheraudioplaying
func (a_ AudioSession) IsOtherAudioPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOtherAudioPlaying"))
	return rv
}


// A Boolean value that indicates whether another app is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isotheraudioplaying
func (a_ AudioSession) SetIsOtherAudioPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOtherAudioPlaying:"), value)
}


// A Boolean value that indicates whether audio output is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isoutputmuted
func (a_ AudioSession) IsOutputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOutputMuted"))
	return rv
}


// A Boolean value that indicates whether audio output is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isoutputmuted
func (a_ AudioSession) SetIsOutputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOutputMuted:"), value)
}



