// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

// Enum types and constants
// AVAudio3DMixingPointSourceInHeadMode - The in-head modes for a point source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode
type Audio3DMixingPointSourceInHeadMode uint

// AVAudio3DMixingRenderingAlgorithm - The types of rendering algorithms available per input bus of the environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm
type Audio3DMixingRenderingAlgorithm uint

const (
	// Audio3DMixingRenderingAlgorithmEqualPowerPanning - An algorithm that pans the data of the mixer bus into a stereo field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/equalPowerPanning
	Audio3DMixingRenderingAlgorithmEqualPowerPanning Audio3DMixingRenderingAlgorithm = 0
)

// AVAudio3DMixingSourceMode - The source modes for the input bus of the audio environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode
type Audio3DMixingSourceMode uint

const (
	// Audio3DMixingSourceModeBypass - A mode that does no spatial rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode/bypass
	Audio3DMixingSourceModeBypass Audio3DMixingSourceMode = 0
)

// AVAudioApplicationMicrophoneInjectionPermission - Constants that indicate an app’s permission to add audio to calls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum
type AudioApplicationMicrophoneInjectionPermission uint

const (
	// AudioApplicationMicrophoneInjectionPermissionDenied - A person denies the app permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/denied
	AudioApplicationMicrophoneInjectionPermissionDenied AudioApplicationMicrophoneInjectionPermission = 0
	// AudioApplicationMicrophoneInjectionPermissionGranted - A person grants the app permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/granted
	AudioApplicationMicrophoneInjectionPermissionGranted AudioApplicationMicrophoneInjectionPermission = 0
	// AudioApplicationMicrophoneInjectionPermissionServiceDisabled - A person disables this service for all apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/serviceDisabled
	AudioApplicationMicrophoneInjectionPermissionServiceDisabled AudioApplicationMicrophoneInjectionPermission = 0
	// AudioApplicationMicrophoneInjectionPermissionUndetermined - The app hasn’t requested a person’s permission to add audio to calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/MicrophoneInjectionPermission-swift.enum/undetermined
	AudioApplicationMicrophoneInjectionPermissionUndetermined AudioApplicationMicrophoneInjectionPermission = 0
)

// AVAudioApplicationRecordPermission - Constants that indicate the app’s permission to record audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum
type AudioApplicationRecordPermission uint

const (
	// AudioApplicationRecordPermissionDenied - Indicates the user denies the app permission to record audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum/denied
	AudioApplicationRecordPermissionDenied AudioApplicationRecordPermission = 0
	// AudioApplicationRecordPermissionGranted - Indicates the user grants the app permission to record audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum/granted
	AudioApplicationRecordPermissionGranted AudioApplicationRecordPermission = 0
)

// AVAudioCommonFormat - The format options that describe common audio formats.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
type AudioCommonFormat uint

// AVAudioContentSource enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource
type AudioContentSource uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Spatial_Live
	AudioContentSource_AppleAV_Spatial_Live AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Spatial_Offline
	AudioContentSource_AppleAV_Spatial_Offline AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Traditional_Live
	AudioContentSource_AppleAV_Traditional_Live AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleAV_Traditional_Offline
	AudioContentSource_AppleAV_Traditional_Offline AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Spatial
	AudioContentSource_AppleCapture_Spatial AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Spatial_Enhanced
	AudioContentSource_AppleCapture_Spatial_Enhanced AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleCapture_Traditional
	AudioContentSource_AppleCapture_Traditional AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleMusic_Spatial
	AudioContentSource_AppleMusic_Spatial AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/appleMusic_Traditional
	AudioContentSource_AppleMusic_Traditional AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/applePassthrough
	AudioContentSource_ApplePassthrough AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Spatial_Live
	AudioContentSource_AV_Spatial_Live AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Spatial_Offline
	AudioContentSource_AV_Spatial_Offline AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Traditional_Live
	AudioContentSource_AV_Traditional_Live AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/av_Traditional_Offline
	AudioContentSource_AV_Traditional_Offline AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Spatial
	AudioContentSource_Capture_Spatial AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Spatial_Enhanced
	AudioContentSource_Capture_Spatial_Enhanced AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/capture_Traditional
	AudioContentSource_Capture_Traditional AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/music_Spatial
	AudioContentSource_Music_Spatial AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/music_Traditional
	AudioContentSource_Music_Traditional AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/passthrough
	AudioContentSource_Passthrough AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/reserved
	AudioContentSource_Reserved AudioContentSource = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioContentSource/unspecified
	AudioContentSource_Unspecified AudioContentSource = 0
)

// AVAudioConverterInputStatus - An option that indicates the status of an audio converter input block.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus
type AudioConverterInputStatus uint

const (
	// AudioConverterInputStatus_EndOfStream - A status that indicates you’re at the end of an audio stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/endOfStream
	AudioConverterInputStatus_EndOfStream AudioConverterInputStatus = 0
	// AudioConverterInputStatus_HaveData - A status that indicates the normal case where you supply data to the converter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/haveData
	AudioConverterInputStatus_HaveData AudioConverterInputStatus = 0
	// AudioConverterInputStatus_NoDataNow - A status that indicates you’re out of data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/noDataNow
	AudioConverterInputStatus_NoDataNow AudioConverterInputStatus = 0
)

// AVAudioDynamicRangeControlConfiguration enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration
type AudioDynamicRangeControlConfiguration uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/capture
	AudioDynamicRangeControlConfiguration_Capture AudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/movie
	AudioDynamicRangeControlConfiguration_Movie AudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/music
	AudioDynamicRangeControlConfiguration_Music AudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/none
	AudioDynamicRangeControlConfiguration_None AudioDynamicRangeControlConfiguration = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioDynamicRangeControlConfiguration/speech
	AudioDynamicRangeControlConfiguration_Speech AudioDynamicRangeControlConfiguration = 0
)

// AVAudioEngineManualRenderingStatus - Status codes that return from the render call to the engine operating in manual rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus
type AudioEngineManualRenderingStatus uint

const (
	// AudioEngineManualRenderingStatusError - A problem that occurs during rendering and results in no data returning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/error
	AudioEngineManualRenderingStatusError AudioEngineManualRenderingStatus = 0
	// AudioEngineManualRenderingStatusInsufficientDataFromInputNode - A condition that occurs when the input node doesn’t return enough input data to satisfy the render request at the time of the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/insufficientDataFromInputNode
	AudioEngineManualRenderingStatusInsufficientDataFromInputNode AudioEngineManualRenderingStatus = 0
)

// AVAudioEnvironmentOutputType - The output types for using with the automatic 3D mixing rendering algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType
type AudioEnvironmentOutputType uint

// AVAudioQuality - The values that specify the sample rate audio quality for encoding and conversion.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality
type AudioQuality uint

// AVAudioRoutingArbitrationCategory - Categories that describe the general nature of your app’s audio use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category
type AudioRoutingArbitrationCategory uint

const (
	// AudioRoutingArbitrationCategoryPlayAndRecord - The app plays and records audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playAndRecord
	AudioRoutingArbitrationCategoryPlayAndRecord AudioRoutingArbitrationCategory = 0
	// AudioRoutingArbitrationCategoryPlayAndRecordVoice - The app uses Voice over IP (VoIP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playAndRecordVoice
	AudioRoutingArbitrationCategoryPlayAndRecordVoice AudioRoutingArbitrationCategory = 0
	// AudioRoutingArbitrationCategoryPlayback - The app plays audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playback
	AudioRoutingArbitrationCategoryPlayback AudioRoutingArbitrationCategory = 0
)

// AVAudioSessionCategoryOptions - Constants that specify optional audio behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct
type AudioSessionCategoryOptions uint

const (
	// AudioSessionCategoryOptionAllowAirPlay - An option that determines whether you can stream audio from this session to AirPlay devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowAirPlay
	AudioSessionCategoryOptionAllowAirPlay AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionAllowBluetooth - An option that determines whether Bluetooth hands-free devices appear as available input routes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetooth
	AudioSessionCategoryOptionAllowBluetooth AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionAllowBluetoothA2DP - An option that determines whether you can stream audio from this session to Bluetooth devices that support the Advanced Audio Distribution Profile (A2DP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetoothA2DP
	AudioSessionCategoryOptionAllowBluetoothA2DP AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionAllowBluetoothHFP - An option that makes Bluetooth Hands-Free Profile (HFP) devices available for audio input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetoothHFP
	AudioSessionCategoryOptionAllowBluetoothHFP AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionBluetoothHighQualityRecording - An option that indicates to enable high-quality audio for input and output routes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/bluetoothHighQualityRecording
	AudioSessionCategoryOptionBluetoothHighQualityRecording AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionDefaultToSpeaker - An option that determines whether audio from the session defaults to the built-in speaker instead of the receiver.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/defaultToSpeaker
	AudioSessionCategoryOptionDefaultToSpeaker AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionDuckOthers - An option that reduces the volume of other audio sessions while audio from this session plays.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/duckOthers
	AudioSessionCategoryOptionDuckOthers AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers - An option that determines whether to pause spoken audio content from other sessions when your app plays its audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/interruptSpokenAudioAndMixWithOthers
	AudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionMixWithOthers - An option that indicates whether audio from this session mixes with audio from active sessions in other audio apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/mixWithOthers
	AudioSessionCategoryOptionMixWithOthers AudioSessionCategoryOptions = 0
	// AudioSessionCategoryOptionOverrideMutedMicrophoneInterruption - An option that indicates whether the system interrupts the audio session when it mutes the built-in microphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/overrideMutedMicrophoneInterruption
	AudioSessionCategoryOptionOverrideMutedMicrophoneInterruption AudioSessionCategoryOptions = 0
)

// AVAudioSessionIOType - Constant values used to specify the audio session’s aggregated I/O behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType
type AudioSessionIOType uint

// AVAudioSessionInterruptionOptions - Constants that indicate the state of an audio session after an interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionOptions
type AudioSessionInterruptionOptions uint

// AVAudioSessionInterruptionReason - Constants that define the reasons for an audio session interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason
type AudioSessionInterruptionReason uint

// AVAudioSessionInterruptionType - Constants that describe the type of an audio interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType
type AudioSessionInterruptionType uint

const (
	// AudioSessionInterruptionTypeBegan - A type that indicates that the operating system began interrupting the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType/began
	AudioSessionInterruptionTypeBegan AudioSessionInterruptionType = 0
	// AudioSessionInterruptionTypeEnded - A type that indicates that the operating system ended interrupting the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType/ended
	AudioSessionInterruptionTypeEnded AudioSessionInterruptionType = 0
)

// AVAudioSessionMicrophoneInjectionMode - The modes of injecting audio into another app’s input stream.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode
type AudioSessionMicrophoneInjectionMode uint

const (
	// AudioSessionMicrophoneInjectionModeNone - A mode that indicates not to use spoken audio injection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode/none
	AudioSessionMicrophoneInjectionModeNone AudioSessionMicrophoneInjectionMode = 0
	// AudioSessionMicrophoneInjectionModeSpokenAudio - A mode that indicates to inject spoken audio, like synthesized speech, along with microphone audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/MicrophoneInjectionMode/spokenAudio
	AudioSessionMicrophoneInjectionModeSpokenAudio AudioSessionMicrophoneInjectionMode = 0
)

// AVAudioSessionPromptStyle - Constants that indicate the prompt style to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum
type AudioSessionPromptStyle uint

// AVAudioSessionRenderingMode - Audio session rendering mode identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum
type AudioSessionRenderingMode uint

// AVAudioSessionRouteChangeReason - Constants that indicate the reason for an audio route change.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason
type AudioSessionRouteChangeReason uint

const (
	// AudioSessionRouteChangeReasonCategoryChange - A value that indicates that the category of the session object changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/categoryChange
	AudioSessionRouteChangeReasonCategoryChange AudioSessionRouteChangeReason = 0
	// AudioSessionRouteChangeReasonNewDeviceAvailable - A value that indicates a user action, such as plugging in a headset, has made a preferred audio route available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/newDeviceAvailable
	AudioSessionRouteChangeReasonNewDeviceAvailable AudioSessionRouteChangeReason = 0
	// AudioSessionRouteChangeReasonOldDeviceUnavailable - A value that indicates that the previous audio output path is no longer available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/oldDeviceUnavailable
	AudioSessionRouteChangeReasonOldDeviceUnavailable AudioSessionRouteChangeReason = 0
)

// AVAudioSessionRouteSelection - Constants used to define the active route selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSelection
type AudioSessionRouteSelection uint

// AVAudioSessionRouteSharingPolicy - Cases that indicate the possible route-sharing policies for an audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum
type AudioSessionRouteSharingPolicy uint

const (
	// AudioSessionRouteSharingPolicyLongFormAudio - A policy that routes output to the shared long-form audio output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/longFormAudio
	AudioSessionRouteSharingPolicyLongFormAudio AudioSessionRouteSharingPolicy = 0
	// AudioSessionRouteSharingPolicyLongFormVideo - A policy that routes output to the shared long-form video output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/longFormVideo
	AudioSessionRouteSharingPolicyLongFormVideo AudioSessionRouteSharingPolicy = 0
)

// AVAudioSessionSetActiveOptions - Options that provide additional information about your app’s audio intentions upon session deactivation.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SetActiveOptions
type AudioSessionSetActiveOptions uint

const (
	// AudioSessionSetActiveOptionNotifyOthersOnDeactivation - An option that indicates that the system should notify other apps that you’ve deactivated your app’s audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SetActiveOptions/notifyOthersOnDeactivation
	AudioSessionSetActiveOptionNotifyOthersOnDeactivation AudioSessionSetActiveOptions = 0
)

// AVAudioSessionSoundStageSize - Constants that specify the perceived size of sounds the audio session plays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize
type AudioSessionSoundStageSize uint

// AVAudioStereoOrientation - Constants that define the supported stereo orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation
type AudioStereoOrientation uint

// AVAudioSessionActivationOptions - Constants that describe the options to pass when activating the audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionActivationOptions
type AudioSessionActivationOptions uint

// AVAudioSessionAnchoringStrategy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy
type AudioSessionAnchoringStrategy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyAutomatic
	AudioSessionAnchoringStrategyAutomatic AudioSessionAnchoringStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyFront
	AudioSessionAnchoringStrategyFront AudioSessionAnchoringStrategy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionAnchoringStrategy/AVAudioSessionAnchoringStrategyScene
	AudioSessionAnchoringStrategyScene AudioSessionAnchoringStrategy = 0
)

// AVAudioSessionSpatialExperience enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum
type AudioSessionSpatialExperience uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceBypassed
	AudioSessionSpatialExperienceBypassed AudioSessionSpatialExperience = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceFixed
	AudioSessionSpatialExperienceFixed AudioSessionSpatialExperience = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionSpatialExperience-c.enum/AVAudioSessionSpatialExperienceHeadTracked
	AudioSessionSpatialExperienceHeadTracked AudioSessionSpatialExperience = 0
)

// AVSpeechSynthesisVoiceTraits - Traits that describe a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits
type SpeechSynthesisVoiceTraits uint

const (
	// SpeechSynthesisVoiceTraitIsPersonalVoice - The trait that indicates a voice is a personal voice.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits/isPersonalVoice
	SpeechSynthesisVoiceTraitIsPersonalVoice SpeechSynthesisVoiceTraits = 0
)

// AVSpeechSynthesisPersonalVoiceAuthorizationStatus - An enumeration that models the personal voices authorization status.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum
type SpeechSynthesisPersonalVoiceAuthorizationStatus uint


