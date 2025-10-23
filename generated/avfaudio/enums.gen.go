// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

// Enum types and constants
// AVAudio3DMixingPointSourceInHeadMode - The in-head modes for a point source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode
type AVAudio3DMixingPointSourceInHeadMode uint

// AVAudio3DMixingRenderingAlgorithm - The types of rendering algorithms available per input bus of the environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm
type AVAudio3DMixingRenderingAlgorithm uint

// AVAudio3DMixingSourceMode - The source modes for the input bus of the audio environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode
type AVAudio3DMixingSourceMode uint

// AVAudioApplicationMicrophoneInjectionPermission - Constants that indicate an app’s permission to add audio to calls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum
type AVAudioApplicationMicrophoneInjectionPermission uint

const (
	// AVAudioApplicationMicrophoneInjectionPermissionDenied - A person denies the app permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/denied
	AVAudioApplicationMicrophoneInjectionPermissionDenied AVAudioApplicationMicrophoneInjectionPermission = 0
	// AVAudioApplicationMicrophoneInjectionPermissionGranted - A person grants the app permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/granted
	AVAudioApplicationMicrophoneInjectionPermissionGranted AVAudioApplicationMicrophoneInjectionPermission = 0
	// AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled - A person disables this service for all apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/serviceDisabled
	AVAudioApplicationMicrophoneInjectionPermissionServiceDisabled AVAudioApplicationMicrophoneInjectionPermission = 0
	// AVAudioApplicationMicrophoneInjectionPermissionUndetermined - The app hasn’t requested a person’s permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/undetermined
	AVAudioApplicationMicrophoneInjectionPermissionUndetermined AVAudioApplicationMicrophoneInjectionPermission = 0
)

// AVAudioApplicationRecordPermission - Constants that indicate the app’s permission to record audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum
type AVAudioApplicationRecordPermission uint

// AVAudioContentSource enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource
type AVAudioContentSource uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Spatial_Live
	AVAudioContentSource_AppleAV_Spatial_Live AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Spatial_Offline
	AVAudioContentSource_AppleAV_Spatial_Offline AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Traditional_Live
	AVAudioContentSource_AppleAV_Traditional_Live AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Traditional_Offline
	AVAudioContentSource_AppleAV_Traditional_Offline AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Spatial
	AVAudioContentSource_AppleCapture_Spatial AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Spatial_Enhanced
	AVAudioContentSource_AppleCapture_Spatial_Enhanced AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Traditional
	AVAudioContentSource_AppleCapture_Traditional AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleMusic_Spatial
	AVAudioContentSource_AppleMusic_Spatial AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleMusic_Traditional
	AVAudioContentSource_AppleMusic_Traditional AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/applePassthrough
	AVAudioContentSource_ApplePassthrough AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Spatial_Live
	AVAudioContentSource_AV_Spatial_Live AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Spatial_Offline
	AVAudioContentSource_AV_Spatial_Offline AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Traditional_Live
	AVAudioContentSource_AV_Traditional_Live AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Traditional_Offline
	AVAudioContentSource_AV_Traditional_Offline AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Spatial
	AVAudioContentSource_Capture_Spatial AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Spatial_Enhanced
	AVAudioContentSource_Capture_Spatial_Enhanced AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Traditional
	AVAudioContentSource_Capture_Traditional AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/music_Spatial
	AVAudioContentSource_Music_Spatial AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/music_Traditional
	AVAudioContentSource_Music_Traditional AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/passthrough
	AVAudioContentSource_Passthrough AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/reserved
	AVAudioContentSource_Reserved AVAudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/unspecified
	AVAudioContentSource_Unspecified AVAudioContentSource = 0
)

// AVAudioConverterInputStatus - An option that indicates the status of an audio converter input block.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus
type AVAudioConverterInputStatus uint

// AVAudioDynamicRangeControlConfiguration enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration
type AVAudioDynamicRangeControlConfiguration uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/capture
	AVAudioDynamicRangeControlConfiguration_Capture AVAudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/movie
	AVAudioDynamicRangeControlConfiguration_Movie AVAudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/music
	AVAudioDynamicRangeControlConfiguration_Music AVAudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/none
	AVAudioDynamicRangeControlConfiguration_None AVAudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/speech
	AVAudioDynamicRangeControlConfiguration_Speech AVAudioDynamicRangeControlConfiguration = 0
)

// AVAudioEnvironmentOutputType - The output types for using with the automatic 3D mixing rendering algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType
type AVAudioEnvironmentOutputType uint

// AVAudioRoutingArbitrationCategory - Categories that describe the general nature of your app’s audio use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category
type AVAudioRoutingArbitrationCategory uint

// AVAudioSessionCategoryOptions - Constants that specify optional audio behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct
type AVAudioSessionCategoryOptions uint

const (
	// AVAudioSessionCategoryOptionAllowBluetoothHFP - An option that makes Bluetooth Hands-Free Profile (HFP) devices available for audio input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetoothHFP
	AVAudioSessionCategoryOptionAllowBluetoothHFP AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption - An option that indicates whether the system interrupts the audio session when it mutes the built-in microphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/overrideMutedMicrophoneInterruption
	AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption AVAudioSessionCategoryOptions = 0
)

// AVAudioSessionIOType - Constant values used to specify the audio session’s aggregated I/O behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType
type AVAudioSessionIOType uint

// AVAudioSessionInterruptionOptions - Constants that indicate the state of an audio session after an interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionOptions
type AVAudioSessionInterruptionOptions uint

// AVAudioSessionInterruptionType - Constants that describe the type of an audio interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType
type AVAudioSessionInterruptionType uint

const (
	// AVAudioSessionInterruptionTypeEnded - A type that indicates that the operating system ended interrupting the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType/ended
	AVAudioSessionInterruptionTypeEnded AVAudioSessionInterruptionType = 0
)

// AVAudioSessionMicrophoneInjectionMode - The modes of injecting audio into another app’s input stream.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode
type AVAudioSessionMicrophoneInjectionMode uint

const (
	// AVAudioSessionMicrophoneInjectionModeNone - A mode that indicates not to use spoken audio injection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode/none
	AVAudioSessionMicrophoneInjectionModeNone AVAudioSessionMicrophoneInjectionMode = 0
	// AVAudioSessionMicrophoneInjectionModeSpokenAudio - A mode that indicates to inject spoken audio, like synthesized speech, along with microphone audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode/spokenAudio
	AVAudioSessionMicrophoneInjectionModeSpokenAudio AVAudioSessionMicrophoneInjectionMode = 0
)

// AVAudioSessionPromptStyle - Constants that indicate the prompt style to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum
type AVAudioSessionPromptStyle uint

// AVAudioSessionRenderingMode - Audio session rendering mode identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum
type AVAudioSessionRenderingMode uint

// AVAudioSessionRouteChangeReason - Constants that indicate the reason for an audio route change.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason
type AVAudioSessionRouteChangeReason uint

const (
	// AVAudioSessionRouteChangeReasonNewDeviceAvailable - A value that indicates a user action, such as plugging in a headset, has made a preferred audio route available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/newDeviceAvailable
	AVAudioSessionRouteChangeReasonNewDeviceAvailable AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory - A value that indicates that the route changed because no suitable route is now available for the specified category.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/noSuitableRouteForCategory
	AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonOldDeviceUnavailable - A value that indicates that the previous audio output path is no longer available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/oldDeviceUnavailable
	AVAudioSessionRouteChangeReasonOldDeviceUnavailable AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonOverride - A value that indicates that the output route was overridden by the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/override
	AVAudioSessionRouteChangeReasonOverride AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonRouteConfigurationChange - A value that indicates that the configuration for a set of I/O ports has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/routeConfigurationChange
	AVAudioSessionRouteChangeReasonRouteConfigurationChange AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonUnknown - A value that indicates the reason for the change is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/unknown
	AVAudioSessionRouteChangeReasonUnknown AVAudioSessionRouteChangeReason = 0
	// AVAudioSessionRouteChangeReasonWakeFromSleep - A value that indicates that the route changed when the device woke up from sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/wakeFromSleep
	AVAudioSessionRouteChangeReasonWakeFromSleep AVAudioSessionRouteChangeReason = 0
)

// AVAudioSessionRouteSelection - Constants used to define the active route selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSelection
type AVAudioSessionRouteSelection uint

// AVAudioSessionRouteSharingPolicy - Cases that indicate the possible route-sharing policies for an audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum
type AVAudioSessionRouteSharingPolicy uint

const (
	// AVAudioSessionRouteSharingPolicyLongForm - A policy that routes output to the shared long-form audio output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/longForm
	AVAudioSessionRouteSharingPolicyLongForm AVAudioSessionRouteSharingPolicy = 0
	// AVAudioSessionRouteSharingPolicyLongFormAudio - A policy that routes output to the shared long-form audio output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/longFormAudio
	AVAudioSessionRouteSharingPolicyLongFormAudio AVAudioSessionRouteSharingPolicy = 0
	// AVAudioSessionRouteSharingPolicyLongFormVideo - A policy that routes output to the shared long-form video output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/longFormVideo
	AVAudioSessionRouteSharingPolicyLongFormVideo AVAudioSessionRouteSharingPolicy = 0
)

// AVAudioSessionSetActiveOptions - Options that provide additional information about your app’s audio intentions upon session deactivation.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SetActiveOptions
type AVAudioSessionSetActiveOptions uint

// AVAudioSessionSoundStageSize - Constants that specify the perceived size of sounds the audio session plays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize
type AVAudioSessionSoundStageSize uint

// AVAudioStereoOrientation - Constants that define the supported stereo orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation
type AVAudioStereoOrientation uint

// AVAudioSessionActivationOptions - Constants that describe the options to pass when activating the audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionActivationOptions
type AVAudioSessionActivationOptions uint

// AVAudioSessionAnchoringStrategy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy
type AVAudioSessionAnchoringStrategy int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyAutomatic
	AVAudioSessionAnchoringStrategyAutomatic AVAudioSessionAnchoringStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyFront
	AVAudioSessionAnchoringStrategyFront AVAudioSessionAnchoringStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyScene
	AVAudioSessionAnchoringStrategyScene AVAudioSessionAnchoringStrategy = 0
)

// AVAudioSessionSpatialExperience enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum
type AVAudioSessionSpatialExperience int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceBypassed
	AVAudioSessionSpatialExperienceBypassed AVAudioSessionSpatialExperience = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceFixed
	AVAudioSessionSpatialExperienceFixed AVAudioSessionSpatialExperience = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceHeadTracked
	AVAudioSessionSpatialExperienceHeadTracked AVAudioSessionSpatialExperience = 0
)

// AVSpeechBoundary - Specifies when to pause or stop speech.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechBoundary
type AVSpeechBoundary uint

// AVSpeechSynthesisPersonalVoiceAuthorizationStatus - An enumeration that models the personal voices authorization status.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum
type AVSpeechSynthesisPersonalVoiceAuthorizationStatus uint


