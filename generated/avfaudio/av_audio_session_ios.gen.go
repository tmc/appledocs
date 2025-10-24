//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSession


// Activates an audio session asynchronously on watchOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/activate(options:completionHandler:)
func (a_ AudioSession) ActivateWithOptionsCompletionHandler(options AudioSessionActivationOptions, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("activateWithOptions:completionHandler:"), options, handler)
}

// Temporarily changes the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/overrideOutputAudioPort(_:)
func (a_ AudioSession) OverrideOutputAudioPortError(portOverride AudioSessionPortOverride, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("overrideOutputAudioPort:error:"), portOverride, outError)
	return rv
}

// Prepares the route selection for long-form video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/prepareRouteSelectionForPlayback(completionHandler:)
func (a_ AudioSession) PrepareRouteSelectionForPlaybackWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareRouteSelectionForPlaybackWithCompletionHandler:"), completionHandler)
}

// Activates or deactivates your app’s audio session using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive(_:options:)
func (a_ AudioSession) SetActiveWithOptionsError(active bool, options AudioSessionSetActiveOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActive:withOptions:error:"), active, options, outError)
	return rv
}

// Activates or deactivates your app’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive:error:
func (a_ AudioSession) SetActiveError(active bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActive:error:"), active, outError)
	return rv
}

// Sets the audio session’s aggregated I/O configuration preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setAggregatedIOPreference(_:)
func (a_ AudioSession) SetAggregatedIOPreferenceError(inIOType AudioSessionIOType, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setAggregatedIOPreference:error:"), inIOType, outError)
	return rv
}

// Sets a Boolean value that indicates whether system sounds and haptics play while recording from audio input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setAllowHapticsAndSystemSoundsDuringRecording(_:)
func (a_ AudioSession) SetAllowHapticsAndSystemSoundsDuringRecordingError(inValue bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setAllowHapticsAndSystemSoundsDuringRecording:error:"), inValue, outError)
	return rv
}

// Sets the audio session’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:)
func (a_ AudioSession) SetCategoryError(category AudioSessionCategory /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:error:"), category, outError)
	return rv
}

// Sets the audio session’s category, mode, and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:mode:options:)
func (a_ AudioSession) SetCategoryModeOptionsError(category AudioSessionCategory /* typedef */, mode AudioSessionMode /* typedef */, options AudioSessionCategoryOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:mode:options:error:"), category, mode, options, outError)
	return rv
}

// Sets the session category, mode, route-sharing policy, and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:mode:policy:options:)
func (a_ AudioSession) SetCategoryModeRouteSharingPolicyOptionsError(category AudioSessionCategory /* typedef */, mode AudioSessionMode /* typedef */, policy AudioSessionRouteSharingPolicy, options AudioSessionCategoryOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:mode:routeSharingPolicy:options:error:"), category, mode, policy, options, outError)
	return rv
}

// Sets the audio session’s category with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setCategory(_:options:)
func (a_ AudioSession) SetCategoryWithOptionsError(category AudioSessionCategory /* typedef */, options AudioSessionCategoryOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setCategory:withOptions:error:"), category, options, outError)
	return rv
}

// Selects a data source for the audio session’s current input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setInputDataSource(_:)
func (a_ AudioSession) SetInputDataSourceError(dataSource IAVAudioSessionDataSourceDescription, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputDataSource:error:"), dataSource, outError)
	return rv
}

// Changes the input gain to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setInputGain(_:)
func (a_ AudioSession) SetInputGainError(gain float32, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputGain:error:"), gain, outError)
	return rv
}

// Sets the spatial audio experience your app intends to provide the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setIntendedSpatialExperience:options:error:
func (a_ AudioSession) SetIntendedSpatialExperienceOptionsError(intendedSpatialExperience AudioSessionSpatialExperience, options foundation.IDictionary, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setIntendedSpatialExperience:options:error:"), intendedSpatialExperience, options, error_)
	return rv
}

// Sets a Boolean value that indicates whether the audio session is a candidate to be the Now Playing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setIsNowPlayingCandidate(_:)
func (a_ AudioSession) SetIsNowPlayingCandidateError(inValue bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setIsNowPlayingCandidate:error:"), inValue, outError)
	return rv
}

// Sets the audio session’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setMode(_:)
func (a_ AudioSession) SetModeError(mode AudioSessionMode /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setMode:error:"), mode, outError)
	return rv
}

// Sets the output data source for an audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setOutputDataSource(_:)
func (a_ AudioSession) SetOutputDataSourceError(dataSource IAVAudioSessionDataSourceDescription, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setOutputDataSource:error:"), dataSource, outError)
	return rv
}

// Sets a Boolean value to inform the system to mute the session’s output audio. The default value is false (unmuted).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setOutputMuted(_:)
func (a_ AudioSession) SetOutputMutedError(muted bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setOutputMuted:error:"), muted, outError)
	return rv
}

// Sets the preferred input port for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredInput(_:)
func (a_ AudioSession) SetPreferredInputError(inPort IAVAudioSessionPortDescription, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredInput:error:"), inPort, outError)
	return rv
}

// Sets the preferred number of input channels for the current route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredInputNumberOfChannels(_:)
func (a_ AudioSession) SetPreferredInputNumberOfChannelsError(count int, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredInputNumberOfChannels:error:"), count, outError)
	return rv
}

// Sets the audio session’s preferred stereo input orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredInputOrientation(_:)
func (a_ AudioSession) SetPreferredInputOrientationError(orientation AudioStereoOrientation, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredInputOrientation:error:"), orientation, outError)
	return rv
}

// Sets the preferred audio I/O buffer duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredIOBufferDuration(_:)
func (a_ AudioSession) SetPreferredIOBufferDurationError(duration float64, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredIOBufferDuration:error:"), duration, outError)
	return rv
}

// Sets the preferred mode of injecting audio into another app’s input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredMicrophoneInjectionMode(_:)
func (a_ AudioSession) SetPreferredMicrophoneInjectionModeError(inValue AudioSessionMicrophoneInjectionMode, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredMicrophoneInjectionMode:error:"), inValue, outError)
	return rv
}

// Sets the preferred number of output channels for the current route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredOutputNumberOfChannels(_:)
func (a_ AudioSession) SetPreferredOutputNumberOfChannelsError(count int, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredOutputNumberOfChannels:error:"), count, outError)
	return rv
}

// Sets the preferred sample rate for audio input and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPreferredSampleRate(_:)
func (a_ AudioSession) SetPreferredSampleRateError(sampleRate float64, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredSampleRate:error:"), sampleRate, outError)
	return rv
}

// Sets a preference to enable echo-canceled input on supported hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersEchoCancelledInput(_:)
func (a_ AudioSession) SetPrefersEchoCancelledInputError(value bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersEchoCancelledInput:error:"), value, error_)
	return rv
}

// Sets a preference to interrupt the audio session when the active route disconnects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersInterruptionOnRouteDisconnect(_:)
func (a_ AudioSession) SetPrefersInterruptionOnRouteDisconnectError(inValue bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersInterruptionOnRouteDisconnect:error:"), inValue, outError)
	return rv
}

// Sets the preference for not interrupting the audio session with system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setPrefersNoInterruptionsFromSystemAlerts(_:)
func (a_ AudioSession) SetPrefersNoInterruptionsFromSystemAlertsError(inValue bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPrefersNoInterruptionsFromSystemAlerts:error:"), inValue, outError)
	return rv
}

// Sets whether your app supplies multichannel audio content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setSupportsMultichannelContent(_:)
func (a_ AudioSession) SetSupportsMultichannelContentError(inValue bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setSupportsMultichannelContent:error:"), inValue, outError)
	return rv
}

// iOS-only properties

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
func (a_ AudioSession) Category() AudioSessionCategory /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("category"))
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

// The number of audio hardware input channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/currentHardwareInputNumberOfChannels
func (a_ AudioSession) CurrentHardwareInputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("currentHardwareInputNumberOfChannels"))
	return rv
}

// The number of audio hardware output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/currentHardwareOutputNumberOfChannels
func (a_ AudioSession) CurrentHardwareOutputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("currentHardwareOutputNumberOfChannels"))
	return rv
}

// The audio hardware sample rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/currentHardwareSampleRate
func (a_ AudioSession) CurrentHardwareSampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("currentHardwareSampleRate"))
	return rv
}

// A description of the current audio route’s input and output ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/currentRoute
func (a_ AudioSession) CurrentRoute() IAVAudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("currentRoute"))
	return rv
}

// The delegate object for the audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/delegate
func (a_ AudioSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}
func (a_ AudioSession) SetDelegate(value unsafe.Pointer) {
	a_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The currently selected input data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputDataSource
func (a_ AudioSession) InputDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("inputDataSource"))
	return rv
}

// An array of available data sources for the audio session’s current input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputDataSources
func (a_ AudioSession) InputDataSources() []AudioSessionDataSourceDescription {
	rv := objc.Send[[]AudioSessionDataSourceDescription](a_.ID, objc.Sel("inputDataSources"))
	return rv
}

// The gain applied to inputs associated with the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputGain
func (a_ AudioSession) InputGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("inputGain"))
	return rv
}

// A Boolean value that indicates whether a hardware audio input path is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputIsAvailable
func (a_ AudioSession) InputIsAvailable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputIsAvailable"))
	return rv
}

// The latency for audio input, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputLatency
func (a_ AudioSession) InputLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("inputLatency"))
	return rv
}

// The number of audio input channels for the current route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/inputNumberOfChannels
func (a_ AudioSession) InputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("inputNumberOfChannels"))
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
func (a_ AudioSession) IntendedSpatialExperienceOptions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("intendedSpatialExperienceOptions"))
	return rv
}

// The current I/O buffer duration, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/ioBufferDuration
func (a_ AudioSession) IOBufferDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("IOBufferDuration"))
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

// A Boolean value that indicates whether you can set the input gain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/isInputGainSettable
func (a_ AudioSession) InputGainSettable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputGainSettable"))
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

// The maximum number of input channels available for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/maximumInputNumberOfChannels
func (a_ AudioSession) MaximumInputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumInputNumberOfChannels"))
	return rv
}

// The maximum number of output channels available for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/maximumOutputNumberOfChannels
func (a_ AudioSession) MaximumOutputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumOutputNumberOfChannels"))
	return rv
}

// The current audio session’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/mode-swift.property
func (a_ AudioSession) Mode() AudioSessionMode /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mode"))
	return rv
}

// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/outputDataSource
func (a_ AudioSession) OutputDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSource"))
	return rv
}

// An array of available output data sources for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/outputDataSources
func (a_ AudioSession) OutputDataSources() []AudioSessionDataSourceDescription {
	rv := objc.Send[[]AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSources"))
	return rv
}

// The latency for audio output, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/outputLatency
func (a_ AudioSession) OutputLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("outputLatency"))
	return rv
}

// The number of audio output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/outputNumberOfChannels
func (a_ AudioSession) OutputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("outputNumberOfChannels"))
	return rv
}

// The systemwide output volume set by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/outputVolume
func (a_ AudioSession) OutputVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("outputVolume"))
	return rv
}

// The preferred hardware sample rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredHardwareSampleRate
func (a_ AudioSession) PreferredHardwareSampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("preferredHardwareSampleRate"))
	return rv
}

// The preferred input port for audio routing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredInput
func (a_ AudioSession) PreferredInput() IAVAudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("preferredInput"))
	return rv
}

// The preferred number of input channels for the current route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredInputNumberOfChannels
func (a_ AudioSession) PreferredInputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("preferredInputNumberOfChannels"))
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

// The preferred I/O buffer duration, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredIOBufferDuration
func (a_ AudioSession) PreferredIOBufferDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("preferredIOBufferDuration"))
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

// The preferred number of output channels for the current route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredOutputNumberOfChannels
func (a_ AudioSession) PreferredOutputNumberOfChannels() int {
	rv := objc.Send[int](a_.ID, objc.Sel("preferredOutputNumberOfChannels"))
	return rv
}

// The preferred sample rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/preferredSampleRate
func (a_ AudioSession) PreferredSampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("preferredSampleRate"))
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

// The current recording permission status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/recordPermission-swift.property
func (a_ AudioSession) RecordPermission() AudioSessionRecordPermission {
	rv := objc.Send[AudioSessionRecordPermission](a_.ID, objc.Sel("recordPermission"))
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

// The current audio sample rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/sampleRate
func (a_ AudioSession) SampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("sampleRate"))
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

// A Boolean value that indicates whether your app supplies multichannel audio content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/supportsMultichannelContent
func (a_ AudioSession) SupportsMultichannelContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsMultichannelContent"))
	return rv
}




