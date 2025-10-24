// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio


// Enum types and constants

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

const (
	// AVAudioApplicationRecordPermissionDenied - Indicates the user denies the app permission to record audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum/denied
	AVAudioApplicationRecordPermissionDenied AVAudioApplicationRecordPermission = 0
	// AVAudioApplicationRecordPermissionGranted - Indicates the user grants the app permission to record audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum/granted
	AVAudioApplicationRecordPermissionGranted AVAudioApplicationRecordPermission = 0
	// AVAudioApplicationRecordPermissionUndetermined - Indicates the app hasn’t requested recording permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.enum/undetermined
	AVAudioApplicationRecordPermissionUndetermined AVAudioApplicationRecordPermission = 0
)


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


// AVAudioRoutingArbitrationCategory - Categories that describe the general nature of your app’s audio use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category
type AVAudioRoutingArbitrationCategory uint

const (
	// AVAudioRoutingArbitrationCategoryPlayAndRecord - The app plays and records audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playAndRecord
	AVAudioRoutingArbitrationCategoryPlayAndRecord AVAudioRoutingArbitrationCategory = 0
	// AVAudioRoutingArbitrationCategoryPlayAndRecordVoice - The app uses Voice over IP (VoIP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playAndRecordVoice
	AVAudioRoutingArbitrationCategoryPlayAndRecordVoice AVAudioRoutingArbitrationCategory = 0
	// AVAudioRoutingArbitrationCategoryPlayback - The app plays audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/Category/playback
	AVAudioRoutingArbitrationCategoryPlayback AVAudioRoutingArbitrationCategory = 0
)


// AVAudioSessionCategoryOptions - Constants that specify optional audio behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct
type AVAudioSessionCategoryOptions uint

const (
	// AVAudioSessionCategoryOptionAllowAirPlay - An option that determines whether you can stream audio from this session to AirPlay devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowAirPlay
	AVAudioSessionCategoryOptionAllowAirPlay AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionAllowBluetooth - An option that determines whether Bluetooth hands-free devices appear as available input routes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetooth
	AVAudioSessionCategoryOptionAllowBluetooth AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionAllowBluetoothA2DP - An option that determines whether you can stream audio from this session to Bluetooth devices that support the Advanced Audio Distribution Profile (A2DP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetoothA2DP
	AVAudioSessionCategoryOptionAllowBluetoothA2DP AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionAllowBluetoothHFP - An option that makes Bluetooth Hands-Free Profile (HFP) devices available for audio input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/allowBluetoothHFP
	AVAudioSessionCategoryOptionAllowBluetoothHFP AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionBluetoothHighQualityRecording - An option that indicates to enable high-quality audio for input and output routes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/bluetoothHighQualityRecording
	AVAudioSessionCategoryOptionBluetoothHighQualityRecording AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionDefaultToSpeaker - An option that determines whether audio from the session defaults to the built-in speaker instead of the receiver.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/defaultToSpeaker
	AVAudioSessionCategoryOptionDefaultToSpeaker AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionDuckOthers - An option that reduces the volume of other audio sessions while audio from this session plays.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/duckOthers
	AVAudioSessionCategoryOptionDuckOthers AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers - An option that determines whether to pause spoken audio content from other sessions when your app plays its audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/interruptSpokenAudioAndMixWithOthers
	AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionMixWithOthers - An option that indicates whether audio from this session mixes with audio from active sessions in other audio apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/mixWithOthers
	AVAudioSessionCategoryOptionMixWithOthers AVAudioSessionCategoryOptions = 0
	// AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption - An option that indicates whether the system interrupts the audio session when it mutes the built-in microphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/CategoryOptions-swift.struct/overrideMutedMicrophoneInterruption
	AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption AVAudioSessionCategoryOptions = 0
)


// AVAudioSessionInterruptionOptions - Constants that indicate the state of an audio session after an interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionOptions
type AVAudioSessionInterruptionOptions uint

const (
	// AVAudioSessionInterruptionOptionShouldResume - An option that indicates the interruption by another audio session has ended and the app can resume its audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionOptions/shouldResume
	AVAudioSessionInterruptionOptionShouldResume AVAudioSessionInterruptionOptions = 0
)


// AVAudioSessionInterruptionReason - Constants that define the reasons for an audio session interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason
type AVAudioSessionInterruptionReason uint

const (
	// AVAudioSessionInterruptionReasonAppWasSuspended - The system suspends the app and interrupts the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason/appWasSuspended
	AVAudioSessionInterruptionReasonAppWasSuspended AVAudioSessionInterruptionReason = 0
	// AVAudioSessionInterruptionReasonBuiltInMicMuted - The system interrupts the audio session when the device mutes the built-in microphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason/builtInMicMuted
	AVAudioSessionInterruptionReasonBuiltInMicMuted AVAudioSessionInterruptionReason = 0
	// AVAudioSessionInterruptionReasonDefault - The system interrupts this audio session when it activates another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason/default
	AVAudioSessionInterruptionReasonDefault AVAudioSessionInterruptionReason = 0
	// AVAudioSessionInterruptionReasonRouteDisconnected - The system interrupts the audio session due to a disconnection of an audio route.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason/routeDisconnected
	AVAudioSessionInterruptionReasonRouteDisconnected AVAudioSessionInterruptionReason = 0
	// AVAudioSessionInterruptionReasonSceneWasBackgrounded - The system backgrounds the scene and interrupts the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionReason/sceneWasBackgrounded
	AVAudioSessionInterruptionReasonSceneWasBackgrounded AVAudioSessionInterruptionReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionInterruptionReason/AVAudioSessionInterruptionReasonDeviceUnauthenticated
	AVAudioSessionInterruptionReasonDeviceUnauthenticated AVAudioSessionInterruptionReason = 0
)


// AVAudioSessionInterruptionType - Constants that describe the type of an audio interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType
type AVAudioSessionInterruptionType uint

const (
	// AVAudioSessionInterruptionTypeBegan - A type that indicates that the operating system began interrupting the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType/began
	AVAudioSessionInterruptionTypeBegan AVAudioSessionInterruptionType = 0
	// AVAudioSessionInterruptionTypeEnded - A type that indicates that the operating system ended interrupting the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/InterruptionType/ended
	AVAudioSessionInterruptionTypeEnded AVAudioSessionInterruptionType = 0
)


// AVAudioSessionIOType - Constant values used to specify the audio session’s aggregated I/O behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType
type AVAudioSessionIOType uint

const (
	// AVAudioSessionIOTypeAggregated - An I/O type that indicates if audio input and output should be presented in the same realtime I/O callback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType/aggregated
	AVAudioSessionIOTypeAggregated AVAudioSessionIOType = 0
	// AVAudioSessionIOTypeNotSpecified - The default audio session I/O type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/IOType/notSpecified
	AVAudioSessionIOTypeNotSpecified AVAudioSessionIOType = 0
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


// AVAudioSessionPortOverride - Constants for use with the 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PortOverride
type AVAudioSessionPortOverride uint

const (
	// AVAudioSessionPortOverrideNone - A value that indicates not to override the output audio port.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PortOverride/none
	AVAudioSessionPortOverrideNone AVAudioSessionPortOverride = 0
	// AVAudioSessionPortOverrideSpeaker - A value that indicates to override the current inputs and outputs, and route audio to the built-in speaker and microphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PortOverride/speaker
	AVAudioSessionPortOverrideSpeaker AVAudioSessionPortOverride = 0
)


// AVAudioSessionPromptStyle - Constants that indicate the prompt style to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum
type AVAudioSessionPromptStyle uint

const (
	// AVAudioSessionPromptStyleNone - Your app shouldn’t issue prompts at this time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum/none
	AVAudioSessionPromptStyleNone AVAudioSessionPromptStyle = 0
	// AVAudioSessionPromptStyleNormal - Your app may use long, verbal prompts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum/normal
	AVAudioSessionPromptStyleNormal AVAudioSessionPromptStyle = 0
	// AVAudioSessionPromptStyleShort - Your app should issue short, nonverbal prompts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/PromptStyle-swift.enum/short
	AVAudioSessionPromptStyleShort AVAudioSessionPromptStyle = 0
)


// AVAudioSessionRecordPermission - The values that define the current state of the record permission request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RecordPermission-swift.enum
type AVAudioSessionRecordPermission uint

const (
	// AVAudioSessionRecordPermissionDenied - A value that indicates that the user has denied recording permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RecordPermission-swift.enum/denied
	AVAudioSessionRecordPermissionDenied AVAudioSessionRecordPermission = 0
	// AVAudioSessionRecordPermissionGranted - A value that indicates that the user has granted recording permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RecordPermission-swift.enum/granted
	AVAudioSessionRecordPermissionGranted AVAudioSessionRecordPermission = 0
	// AVAudioSessionRecordPermissionUndetermined - A value that indicates that the user hasn’t granted or denied recording permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RecordPermission-swift.enum/undetermined
	AVAudioSessionRecordPermissionUndetermined AVAudioSessionRecordPermission = 0
)


// AVAudioSessionRenderingMode - Audio session rendering mode identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum
type AVAudioSessionRenderingMode uint

const (
	// AVAudioSessionRenderingModeDolbyAtmos - A mode that represents Dolby Atmos.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/dolbyAtmos
	AVAudioSessionRenderingModeDolbyAtmos AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeDolbyAudio - A mode that represents Dolby audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/dolbyAudio
	AVAudioSessionRenderingModeDolbyAudio AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeMonoStereo - A mode that represents non multi-channel audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/monoStereo
	AVAudioSessionRenderingModeMonoStereo AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeNotApplicable - A mode that represents there’s no asset in a loading or playing state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/notApplicable
	AVAudioSessionRenderingModeNotApplicable AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeSpatialAudio - A mode that represents a fallback for when hardware capabilities don’t support Dolby.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/spatialAudio
	AVAudioSessionRenderingModeSpatialAudio AVAudioSessionRenderingMode = 0
	// AVAudioSessionRenderingModeSurround - A mode that represents general multi-channel audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RenderingMode-swift.enum/surround
	AVAudioSessionRenderingModeSurround AVAudioSessionRenderingMode = 0
)


// AVAudioSessionRouteChangeReason - Constants that indicate the reason for an audio route change.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason
type AVAudioSessionRouteChangeReason uint

const (
	// AVAudioSessionRouteChangeReasonCategoryChange - A value that indicates that the category of the session object changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteChangeReason/categoryChange
	AVAudioSessionRouteChangeReasonCategoryChange AVAudioSessionRouteChangeReason = 0
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
	// AVAudioSessionRouteSharingPolicyDefault - A policy that follows standard rules for routing audio output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/default
	AVAudioSessionRouteSharingPolicyDefault AVAudioSessionRouteSharingPolicy = 0
	// AVAudioSessionRouteSharingPolicyIndependent - A policy in which the route picker UI directs videos to a wireless route.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/RouteSharingPolicy-swift.enum/independent
	AVAudioSessionRouteSharingPolicyIndependent AVAudioSessionRouteSharingPolicy = 0
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

const (
	// AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation - An option that indicates that the system should notify other apps that you’ve deactivated your app’s audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SetActiveOptions/notifyOthersOnDeactivation
	AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation AVAudioSessionSetActiveOptions = 0
)


// AVAudioSessionSilenceSecondaryAudioHintType - Constants that indicate whether optional secondary audio muting should begin or end.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SilenceSecondaryAudioHintType
type AVAudioSessionSilenceSecondaryAudioHintType uint

const (
	// AVAudioSessionSilenceSecondaryAudioHintTypeBegin - A value that indicates that another application’s primary audio has started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SilenceSecondaryAudioHintType/begin
	AVAudioSessionSilenceSecondaryAudioHintTypeBegin AVAudioSessionSilenceSecondaryAudioHintType = 0
	// AVAudioSessionSilenceSecondaryAudioHintTypeEnd - A value that indicates that another application’s primary audio has stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SilenceSecondaryAudioHintType/end
	AVAudioSessionSilenceSecondaryAudioHintTypeEnd AVAudioSessionSilenceSecondaryAudioHintType = 0
)


// AVAudioSessionSoundStageSize - Constants that specify the perceived size of sounds the audio session plays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize
type AVAudioSessionSoundStageSize uint

const (
	// AVAudioSessionSoundStageSizeAutomatic - The system sets the sound stage size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize/automatic
	AVAudioSessionSoundStageSizeAutomatic AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeLarge - A large sound stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize/large
	AVAudioSessionSoundStageSizeLarge AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeMedium - A medium sound stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize/medium
	AVAudioSessionSoundStageSizeMedium AVAudioSessionSoundStageSize = 0
	// AVAudioSessionSoundStageSizeSmall - A small sound stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/SoundStageSize/small
	AVAudioSessionSoundStageSizeSmall AVAudioSessionSoundStageSize = 0
)


// AVAudioStereoOrientation - Constants that define the supported stereo orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation
type AVAudioStereoOrientation uint

const (
	// AVAudioStereoOrientationLandscapeLeft - Audio capture should be horizontally oriented, with the USB-C or Lightning connector on the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation/landscapeLeft
	AVAudioStereoOrientationLandscapeLeft AVAudioStereoOrientation = 0
	// AVAudioStereoOrientationLandscapeRight - Audio capture should be horizontally oriented, with the USB-C or Lightning connector on the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation/landscapeRight
	AVAudioStereoOrientationLandscapeRight AVAudioStereoOrientation = 0
	// AVAudioStereoOrientationNone - The audio session isn’t configured for stereo recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation/none
	AVAudioStereoOrientationNone AVAudioStereoOrientation = 0
	// AVAudioStereoOrientationPortrait - Audio capture should be vertically oriented, with the USB-C or Lightning connector on the bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation/portrait
	AVAudioStereoOrientationPortrait AVAudioStereoOrientation = 0
	// AVAudioStereoOrientationPortraitUpsideDown - Audio capture should be vertically oriented, with the USB-C or Lightning connector on the top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/StereoOrientation/portraitUpsideDown
	AVAudioStereoOrientationPortraitUpsideDown AVAudioStereoOrientation = 0
)


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


// AVAudio3DMixingPointSourceInHeadMode - The in-head modes for a point source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode
type AVAudio3DMixingPointSourceInHeadMode uint

const (
	// AVAudio3DMixingPointSourceInHeadModeBypass - The point source distributes into each output channel inside the head of the listener.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode/bypass
	AVAudio3DMixingPointSourceInHeadModeBypass AVAudio3DMixingPointSourceInHeadMode = 0
	// AVAudio3DMixingPointSourceInHeadModeMono - The point source remains a single mono source inside the head of the listener regardless of the channels it consists of.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingPointSourceInHeadMode/mono
	AVAudio3DMixingPointSourceInHeadModeMono AVAudio3DMixingPointSourceInHeadMode = 0
)


// AVAudio3DMixingRenderingAlgorithm - The types of rendering algorithms available per input bus of the environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm
type AVAudio3DMixingRenderingAlgorithm uint

const (
	// AVAudio3DMixingRenderingAlgorithmAuto - Automatically selects the highest-quality rendering algorithm available for the current playback hardware.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/auto
	AVAudio3DMixingRenderingAlgorithmAuto AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmEqualPowerPanning - An algorithm that pans the data of the mixer bus into a stereo field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/equalPowerPanning
	AVAudio3DMixingRenderingAlgorithmEqualPowerPanning AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmHRTF - A high-quality algorithm that uses filtering to emulate 3D space in headphones.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/HRTF
	AVAudio3DMixingRenderingAlgorithmHRTF AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmHRTFHQ - A higher-quality head-related transfer function rendering algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/HRTFHQ
	AVAudio3DMixingRenderingAlgorithmHRTFHQ AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmSoundField - An algorithm that renders to multichannel hardware.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/soundField
	AVAudio3DMixingRenderingAlgorithmSoundField AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmSphericalHead - An algorithm that emulates 3D space in headphones by simulating interaural time delays and other spatial cues.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/sphericalHead
	AVAudio3DMixingRenderingAlgorithmSphericalHead AVAudio3DMixingRenderingAlgorithm = 0
	// AVAudio3DMixingRenderingAlgorithmStereoPassThrough - An algorithm to use when the source data doesn’t need localization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingRenderingAlgorithm/stereoPassThrough
	AVAudio3DMixingRenderingAlgorithmStereoPassThrough AVAudio3DMixingRenderingAlgorithm = 0
)


// AVAudio3DMixingSourceMode - The source modes for the input bus of the audio environment node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode
type AVAudio3DMixingSourceMode uint

const (
	// AVAudio3DMixingSourceModeAmbienceBed - The input channels spread around the listener as far-field sources that anchor to global space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode/ambienceBed
	AVAudio3DMixingSourceModeAmbienceBed AVAudio3DMixingSourceMode = 0
	// AVAudio3DMixingSourceModeBypass - A mode that does no spatial rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode/bypass
	AVAudio3DMixingSourceModeBypass AVAudio3DMixingSourceMode = 0
	// AVAudio3DMixingSourceModePointSource - All channels of the bus render as a single source at the location of the source node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode/pointSource
	AVAudio3DMixingSourceModePointSource AVAudio3DMixingSourceMode = 0
	// AVAudio3DMixingSourceModeSpatializeIfMono - A mono input bus that renders as a point source at the location of the source node.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixingSourceMode/spatializeIfMono
	AVAudio3DMixingSourceModeSpatializeIfMono AVAudio3DMixingSourceMode = 0
)


// AVAudioCommonFormat - The format options that describe common audio formats.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
type AVAudioCommonFormat uint

const (
	// AVAudioOtherFormat - A format other than one the enumeration specifies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat/otherFormat
	AVAudioOtherFormat AVAudioCommonFormat = 0
	// AVAudioPCMFormatFloat32 - A format that represents the standard format as native-endian floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat/pcmFormatFloat32
	AVAudioPCMFormatFloat32 AVAudioCommonFormat = 0
	// AVAudioPCMFormatFloat64 - A format that represents native-endian doubles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat/pcmFormatFloat64
	AVAudioPCMFormatFloat64 AVAudioCommonFormat = 0
	// AVAudioPCMFormatInt16 - A format that represents signed 16-bit native-endian integers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat/pcmFormatInt16
	AVAudioPCMFormatInt16 AVAudioCommonFormat = 0
	// AVAudioPCMFormatInt32 - A format that represents signed 32-bit native-endian integers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat/pcmFormatInt32
	AVAudioPCMFormatInt32 AVAudioCommonFormat = 0
)


// AVAudioConverterInputStatus - An option that indicates the status of an audio converter input block.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus
type AVAudioConverterInputStatus uint

const (
	// AVAudioConverterInputStatus_EndOfStream - A status that indicates you’re at the end of an audio stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/endOfStream
	AVAudioConverterInputStatus_EndOfStream AVAudioConverterInputStatus = 0
	// AVAudioConverterInputStatus_HaveData - A status that indicates the normal case where you supply data to the converter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/haveData
	AVAudioConverterInputStatus_HaveData AVAudioConverterInputStatus = 0
	// AVAudioConverterInputStatus_NoDataNow - A status that indicates you’re out of data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterInputStatus/noDataNow
	AVAudioConverterInputStatus_NoDataNow AVAudioConverterInputStatus = 0
)


// AVAudioConverterOutputStatus - An option that indicates the return status of an audio converter method.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus
type AVAudioConverterOutputStatus uint

const (
	// AVAudioConverterOutputStatus_EndOfStream - A status that indicates the method reaches the end of the stream, and doesn’t return any data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus/endOfStream
	AVAudioConverterOutputStatus_EndOfStream AVAudioConverterOutputStatus = 0
	// AVAudioConverterOutputStatus_Error - A status that indicates the method encounters an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus/error
	AVAudioConverterOutputStatus_Error AVAudioConverterOutputStatus = 0
	// AVAudioConverterOutputStatus_HaveData - A status that indicates that the method returns all of the requested data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus/haveData
	AVAudioConverterOutputStatus_HaveData AVAudioConverterOutputStatus = 0
	// AVAudioConverterOutputStatus_InputRanDry - A status that indicates the method doesn’t have enough input available to satisfy the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterOutputStatus/inputRanDry
	AVAudioConverterOutputStatus_InputRanDry AVAudioConverterOutputStatus = 0
)


// AVAudioConverterPrimeMethod - Options for the prime method property.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeMethod
type AVAudioConverterPrimeMethod uint

const (
	// AVAudioConverterPrimeMethod_None - An option to prime the converter assumes leading and trailing frames are silence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeMethod/none
	AVAudioConverterPrimeMethod_None AVAudioConverterPrimeMethod = 0
	// AVAudioConverterPrimeMethod_Normal - An option to prime with trailing (zero latency) frames where the converter assumes the leading frames are silence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeMethod/normal
	AVAudioConverterPrimeMethod_Normal AVAudioConverterPrimeMethod = 0
	// AVAudioConverterPrimeMethod_Pre - An option to prime with leading and trailing input frames.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeMethod/pre
	AVAudioConverterPrimeMethod_Pre AVAudioConverterPrimeMethod = 0
)


// AVAudioEngineManualRenderingError - Constants that describe error codes that the framework returns from manual rendering mode methods.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingError
type AVAudioEngineManualRenderingError uint

const (
	// AVAudioEngineManualRenderingErrorInitialized - An operation that the system can’t perform because the engine is still running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingError/initialized
	AVAudioEngineManualRenderingErrorInitialized AVAudioEngineManualRenderingError = 0
	// AVAudioEngineManualRenderingErrorInvalidMode - An operation the system can’t perform because the engine isn’t in manual rendering mode or the right variant of it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingError/invalidMode
	AVAudioEngineManualRenderingErrorInvalidMode AVAudioEngineManualRenderingError = 0
	// AVAudioEngineManualRenderingErrorNotRunning - An operation the system can’t perform because the engine isn’t running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingError/notRunning
	AVAudioEngineManualRenderingErrorNotRunning AVAudioEngineManualRenderingError = 0
)


// AVAudioEngineManualRenderingMode - The two modes for manual rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingMode
type AVAudioEngineManualRenderingMode uint

const (
	// AVAudioEngineManualRenderingModeOffline - An engine that operates in an offline mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingMode/offline
	AVAudioEngineManualRenderingModeOffline AVAudioEngineManualRenderingMode = 0
	// AVAudioEngineManualRenderingModeRealtime - An engine that operates under real-time constraints and doesn’t make blocking calls while rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingMode/realtime
	AVAudioEngineManualRenderingModeRealtime AVAudioEngineManualRenderingMode = 0
)


// AVAudioEngineManualRenderingStatus - Status codes that return from the render call to the engine operating in manual rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus
type AVAudioEngineManualRenderingStatus uint

const (
	// AVAudioEngineManualRenderingStatusCannotDoInCurrentContext - An operation that the system can’t perform under the current conditions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/cannotDoInCurrentContext
	AVAudioEngineManualRenderingStatusCannotDoInCurrentContext AVAudioEngineManualRenderingStatus = 0
	// AVAudioEngineManualRenderingStatusError - A problem that occurs during rendering and results in no data returning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/error
	AVAudioEngineManualRenderingStatusError AVAudioEngineManualRenderingStatus = 0
	// AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode - A condition that occurs when the input node doesn’t return enough input data to satisfy the render request at the time of the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/insufficientDataFromInputNode
	AVAudioEngineManualRenderingStatusInsufficientDataFromInputNode AVAudioEngineManualRenderingStatus = 0
	// AVAudioEngineManualRenderingStatusSuccess - A status that indicates the successful return of the requested data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngineManualRenderingStatus/success
	AVAudioEngineManualRenderingStatusSuccess AVAudioEngineManualRenderingStatus = 0
)


// AVAudioEnvironmentDistanceAttenuationModel - Types of distance attenuation models.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationModel
type AVAudioEnvironmentDistanceAttenuationModel uint

const (
	// AVAudioEnvironmentDistanceAttenuationModelExponential - An exponential model that describes the drop-off in gain as the source moves away from the listener.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationModel/exponential
	AVAudioEnvironmentDistanceAttenuationModelExponential AVAudioEnvironmentDistanceAttenuationModel = 0
	// AVAudioEnvironmentDistanceAttenuationModelInverse - An inverse model that describes the drop-off in gain as the source moves away from the listener.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationModel/inverse
	AVAudioEnvironmentDistanceAttenuationModelInverse AVAudioEnvironmentDistanceAttenuationModel = 0
	// AVAudioEnvironmentDistanceAttenuationModelLinear - A linear model that describes the drop-off in gain as the source moves away from the listener.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationModel/linear
	AVAudioEnvironmentDistanceAttenuationModelLinear AVAudioEnvironmentDistanceAttenuationModel = 0
)


// AVAudioEnvironmentOutputType - The output types for using with the automatic 3D mixing rendering algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType
type AVAudioEnvironmentOutputType uint

const (
	// AVAudioEnvironmentOutputTypeAuto - Automatically detects the playback route and picks the correct output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType/auto
	AVAudioEnvironmentOutputTypeAuto AVAudioEnvironmentOutputType = 0
	// AVAudioEnvironmentOutputTypeBuiltInSpeakers - Renders the audio output for built-in speakers on the current hardware.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType/builtInSpeakers
	AVAudioEnvironmentOutputTypeBuiltInSpeakers AVAudioEnvironmentOutputType = 0
	// AVAudioEnvironmentOutputTypeExternalSpeakers - Renders the audio output for external speakers according to the audio environment node’s output channel layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType/externalSpeakers
	AVAudioEnvironmentOutputTypeExternalSpeakers AVAudioEnvironmentOutputType = 0
	// AVAudioEnvironmentOutputTypeHeadphones - Renders the audio output for headphones.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentOutputType/headphones
	AVAudioEnvironmentOutputTypeHeadphones AVAudioEnvironmentOutputType = 0
)


// AVAudioPlayerNodeBufferOptions - The buffer options that control the playback scheduling.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeBufferOptions
type AVAudioPlayerNodeBufferOptions uint

const (
	// AVAudioPlayerNodeBufferInterrupts - An option that indicates the buffer interrupts any buffer in a playing state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeBufferOptions/interrupts
	AVAudioPlayerNodeBufferInterrupts AVAudioPlayerNodeBufferOptions = 0
	// AVAudioPlayerNodeBufferInterruptsAtLoop - An option that indicates the buffer interrupts any buffer in a playing state at its loop point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeBufferOptions/interruptsAtLoop
	AVAudioPlayerNodeBufferInterruptsAtLoop AVAudioPlayerNodeBufferOptions = 0
	// AVAudioPlayerNodeBufferLoops - An option that indicates the buffer loops indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeBufferOptions/loops
	AVAudioPlayerNodeBufferLoops AVAudioPlayerNodeBufferOptions = 0
)


// AVAudioPlayerNodeCompletionCallbackType - Constants that specify when the framework must invoke the completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType
type AVAudioPlayerNodeCompletionCallbackType uint

const (
	// AVAudioPlayerNodeCompletionDataConsumed - A completion handler that indicates the player consumes the buffer or file data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType/dataConsumed
	AVAudioPlayerNodeCompletionDataConsumed AVAudioPlayerNodeCompletionCallbackType = 0
	// AVAudioPlayerNodeCompletionDataPlayedBack - A completion handler that indicates the player finishes the buffer or file data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType/dataPlayedBack
	AVAudioPlayerNodeCompletionDataPlayedBack AVAudioPlayerNodeCompletionCallbackType = 0
	// AVAudioPlayerNodeCompletionDataRendered - A completion handler that indicates the player renders the buffer or file data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType/dataRendered
	AVAudioPlayerNodeCompletionDataRendered AVAudioPlayerNodeCompletionCallbackType = 0
)


// AVAudioQuality - The values that specify the sample rate audio quality for encoding and conversion.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality
type AVAudioQuality uint

const (
	// AVAudioQualityHigh - A value that represents a high audio quality for encoding and conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality/high
	AVAudioQualityHigh AVAudioQuality = 0
	// AVAudioQualityLow - A value that represents a low audio quality for encoding and conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality/low
	AVAudioQualityLow AVAudioQuality = 0
	// AVAudioQualityMax - A value that represents a maximum audio quality for encoding and conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality/max
	AVAudioQualityMax AVAudioQuality = 0
	// AVAudioQualityMedium - A value that represents a medium audio quality for encoding and conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality/medium
	AVAudioQualityMedium AVAudioQuality = 0
	// AVAudioQualityMin - A value that represents a minimum audio quality for encoding and conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioQuality/min
	AVAudioQualityMin AVAudioQuality = 0
)


// AVAudioSessionActivationOptions - Constants that describe the options to pass when activating the audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionActivationOptions
type AVAudioSessionActivationOptions uint

const (
	// AVAudioSessionActivationOptionNone - A value that indicates the system should activate the audio session with no options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionActivationOptions/AVAudioSessionActivationOptionNone
	AVAudioSessionActivationOptionNone AVAudioSessionActivationOptions = 0
)


// AVAudioUnitDistortionPreset - Constants that represent preset audio distortions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset
type AVAudioUnitDistortionPreset uint

const (
	// AVAudioUnitDistortionPresetDrumsBitBrush - A preset that represents a bit brush drums distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/drumsBitBrush
	AVAudioUnitDistortionPresetDrumsBitBrush AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetDrumsBufferBeats - A preset that represents a buffer beat drums distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/drumsBufferBeats
	AVAudioUnitDistortionPresetDrumsBufferBeats AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetDrumsLoFi - A preset that represents a low fidelity drums distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/drumsLoFi
	AVAudioUnitDistortionPresetDrumsLoFi AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiBrokenSpeaker - A preset that represents a broken speaker distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiBrokenSpeaker
	AVAudioUnitDistortionPresetMultiBrokenSpeaker AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiCellphoneConcert - A preset that represents a cellphone concert distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiCellphoneConcert
	AVAudioUnitDistortionPresetMultiCellphoneConcert AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDecimated1 - A preset that represents a variant of the decimated distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDecimated1
	AVAudioUnitDistortionPresetMultiDecimated1 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDecimated2 - A preset that represents a variant of the decimated distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDecimated2
	AVAudioUnitDistortionPresetMultiDecimated2 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDecimated3 - A preset that represents a variant of the decimated distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDecimated3
	AVAudioUnitDistortionPresetMultiDecimated3 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDecimated4 - A preset that represents a variant of the decimated distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDecimated4
	AVAudioUnitDistortionPresetMultiDecimated4 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDistortedCubed - A preset that represents a distorted cubed distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDistortedCubed
	AVAudioUnitDistortionPresetMultiDistortedCubed AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDistortedFunk - A preset that represents a distorted funk distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDistortedFunk
	AVAudioUnitDistortionPresetMultiDistortedFunk AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiDistortedSquared - A preset that represents a distorted squared distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiDistortedSquared
	AVAudioUnitDistortionPresetMultiDistortedSquared AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiEcho1 - A preset that represents a variant of an echo distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiEcho1
	AVAudioUnitDistortionPresetMultiEcho1 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiEcho2 - A preset that represents a variant of an echo distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiEcho2
	AVAudioUnitDistortionPresetMultiEcho2 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiEchoTight1 - A preset that represents a variant of a tight echo distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiEchoTight1
	AVAudioUnitDistortionPresetMultiEchoTight1 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiEchoTight2 - A preset that represents a variant of a tight echo distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiEchoTight2
	AVAudioUnitDistortionPresetMultiEchoTight2 AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetMultiEverythingIsBroken - A preset that represents an everything-is-broken distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/multiEverythingIsBroken
	AVAudioUnitDistortionPresetMultiEverythingIsBroken AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetSpeechAlienChatter - A preset that represents an alien chatter distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/speechAlienChatter
	AVAudioUnitDistortionPresetSpeechAlienChatter AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetSpeechCosmicInterference - A preset that represents a cosmic interference distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/speechCosmicInterference
	AVAudioUnitDistortionPresetSpeechCosmicInterference AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetSpeechGoldenPi - A preset that represents a golden pi distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/speechGoldenPi
	AVAudioUnitDistortionPresetSpeechGoldenPi AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetSpeechRadioTower - A preset that represents a radio tower distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/speechRadioTower
	AVAudioUnitDistortionPresetSpeechRadioTower AVAudioUnitDistortionPreset = 0
	// AVAudioUnitDistortionPresetSpeechWaves - A preset that represents a speech wave distortion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset/speechWaves
	AVAudioUnitDistortionPresetSpeechWaves AVAudioUnitDistortionPreset = 0
)


// AVAudioUnitEQFilterType - Filter types available to use with the filter type property.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType
type AVAudioUnitEQFilterType uint

const (
	// AVAudioUnitEQFilterTypeBandPass - A type that represents a bandpass filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/bandPass
	AVAudioUnitEQFilterTypeBandPass AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeBandStop - A type that represents a band-stop filter, also known as a notch filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/bandStop
	AVAudioUnitEQFilterTypeBandStop AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeHighPass - A type that represents a simple Butterworth second-order high-pass filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/highPass
	AVAudioUnitEQFilterTypeHighPass AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeHighShelf - A type that represents a high-shelf filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/highShelf
	AVAudioUnitEQFilterTypeHighShelf AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeLowPass - A type that represents a simple Butterworth second-order low-pass filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/lowPass
	AVAudioUnitEQFilterTypeLowPass AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeLowShelf - A type that represents a low-shelf filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/lowShelf
	AVAudioUnitEQFilterTypeLowShelf AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeParametric - A type that represents a parametric filter that derives from a Butterworth analog prototype.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/parametric
	AVAudioUnitEQFilterTypeParametric AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeResonantHighPass - A type that represents a high-pass filter with resonance support using the bandwidth parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/resonantHighPass
	AVAudioUnitEQFilterTypeResonantHighPass AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeResonantHighShelf - A type that represents a high-shelf filter with resonance support using the bandwidth parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/resonantHighShelf
	AVAudioUnitEQFilterTypeResonantHighShelf AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeResonantLowPass - A type that represents a low-pass filter with resonance support using the bandwidth parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/resonantLowPass
	AVAudioUnitEQFilterTypeResonantLowPass AVAudioUnitEQFilterType = 0
	// AVAudioUnitEQFilterTypeResonantLowShelf - A type that represents a low-shelf filter with resonance support using the bandwidth parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterType/resonantLowShelf
	AVAudioUnitEQFilterTypeResonantLowShelf AVAudioUnitEQFilterType = 0
)


// AVAudioUnitReverbPreset - Constants that represent preset reverbs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset
type AVAudioUnitReverbPreset uint

const (
	// AVAudioUnitReverbPresetCathedral - A preset that represents a reverb with the acoustic characteristics of a cathedral environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/cathedral
	AVAudioUnitReverbPresetCathedral AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetLargeChamber - A preset that represents a reverb with the acoustic characteristics of a large-sized chamber environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/largeChamber
	AVAudioUnitReverbPresetLargeChamber AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetLargeHall - A preset that represents a reverb with the acoustic characteristics of a large-sized hall environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/largeHall
	AVAudioUnitReverbPresetLargeHall AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetLargeHall2 - A preset that represents a reverb with the acoustic characteristics of an alternative large-sized hall environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/largeHall2
	AVAudioUnitReverbPresetLargeHall2 AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetLargeRoom - A preset that represents a reverb with the acoustic characteristics of a large-sized room environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/largeRoom
	AVAudioUnitReverbPresetLargeRoom AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetLargeRoom2 - A preset that represents a reverb with the acoustic characteristics of an alternative large-sized room environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/largeRoom2
	AVAudioUnitReverbPresetLargeRoom2 AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetMediumChamber - A preset that represents a reverb with the acoustic characteristics of a medium-sized chamber environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/mediumChamber
	AVAudioUnitReverbPresetMediumChamber AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetMediumHall - A preset that represents a reverb with the acoustic characteristics of a medium-sized hall environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/mediumHall
	AVAudioUnitReverbPresetMediumHall AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetMediumHall2 - A preset that represents a reverb with the acoustic characteristics of an alternative medium-sized hall environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/mediumHall2
	AVAudioUnitReverbPresetMediumHall2 AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetMediumHall3 - A preset that represents a reverb with the acoustic characteristics of an alternative medium-sized hall environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/mediumHall3
	AVAudioUnitReverbPresetMediumHall3 AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetMediumRoom - A preset that represents a reverb with the acoustic characteristics of a medium-sized room environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/mediumRoom
	AVAudioUnitReverbPresetMediumRoom AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetPlate - A preset that represents a reverb with the acoustic characteristics of a plate environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/plate
	AVAudioUnitReverbPresetPlate AVAudioUnitReverbPreset = 0
	// AVAudioUnitReverbPresetSmallRoom - A preset that represents a reverb with the acoustic characteristics of a small-sized room environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset/smallRoom
	AVAudioUnitReverbPresetSmallRoom AVAudioUnitReverbPreset = 0
)


// AVAudioVoiceProcessingOtherAudioDuckingLevel - Constants that define the supported ducking levels.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level
type AVAudioVoiceProcessingOtherAudioDuckingLevel uint

const (
	// AVAudioVoiceProcessingOtherAudioDuckingLevelDefault - The default ducking level for typical voice chat.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level/default
	AVAudioVoiceProcessingOtherAudioDuckingLevelDefault AVAudioVoiceProcessingOtherAudioDuckingLevel = 0
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMax - Applies maximum ducking to other audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level/max
	AVAudioVoiceProcessingOtherAudioDuckingLevelMax AVAudioVoiceProcessingOtherAudioDuckingLevel = 0
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMid - Applies medium ducking to other audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level/mid
	AVAudioVoiceProcessingOtherAudioDuckingLevelMid AVAudioVoiceProcessingOtherAudioDuckingLevel = 0
	// AVAudioVoiceProcessingOtherAudioDuckingLevelMin - Applies minimum ducking to other audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration/Level/min
	AVAudioVoiceProcessingOtherAudioDuckingLevelMin AVAudioVoiceProcessingOtherAudioDuckingLevel = 0
)


// AVAudioVoiceProcessingSpeechActivityEvent - Types of speech activity events.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingSpeechActivityEvent
type AVAudioVoiceProcessingSpeechActivityEvent uint

const (
	// AVAudioVoiceProcessingSpeechActivityEnded - Indicates the end of speech activity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingSpeechActivityEvent/ended
	AVAudioVoiceProcessingSpeechActivityEnded AVAudioVoiceProcessingSpeechActivityEvent = 0
	// AVAudioVoiceProcessingSpeechActivityStarted - Indicates the start of speech activity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingSpeechActivityEvent/started
	AVAudioVoiceProcessingSpeechActivityStarted AVAudioVoiceProcessingSpeechActivityEvent = 0
)


// AVMIDIControlChangeMessageType - Constants that represents control change event types.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum
type AVMIDIControlChangeMessageType uint

const (
	// AVMIDIControlChangeMessageTypeAllNotesOff - An event type for muting all sounding notes while maintaining the release time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/allNotesOff
	AVMIDIControlChangeMessageTypeAllNotesOff AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeAllSoundOff - An event type for muting all sounding notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/allSoundOff
	AVMIDIControlChangeMessageTypeAllSoundOff AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeAttackTime - An event type for controlling the attack time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/attackTime
	AVMIDIControlChangeMessageTypeAttackTime AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeBalance - An event type for controlling the left and right channel balance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/balance
	AVMIDIControlChangeMessageTypeBalance AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeBankSelect - An event type for switching bank selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/bankSelect
	AVMIDIControlChangeMessageTypeBankSelect AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeBreath - An event type for a breath controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/breath
	AVMIDIControlChangeMessageTypeBreath AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeBrightness - An event type for controlling the brightness.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/brightness
	AVMIDIControlChangeMessageTypeBrightness AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeChorusLevel - An event type for controlling the chorus level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/chorusLevel
	AVMIDIControlChangeMessageTypeChorusLevel AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeDataEntry - An event type for controlling the data entry parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/dataEntry
	AVMIDIControlChangeMessageTypeDataEntry AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeDecayTime - An event type for controlling the decay time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/decayTime
	AVMIDIControlChangeMessageTypeDecayTime AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeExpression - An event type that represents an expression controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/expression
	AVMIDIControlChangeMessageTypeExpression AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeFilterResonance - An event type for a filter resonance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/filterResonance
	AVMIDIControlChangeMessageTypeFilterResonance AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeFoot - An event type for sending continuous stream of values when using a foot controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/foot
	AVMIDIControlChangeMessageTypeFoot AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeHold2Pedal - An event type for holding notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/hold2Pedal
	AVMIDIControlChangeMessageTypeHold2Pedal AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeLegatoPedal - An event type for switching the legato pedal on or off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/legatoPedal
	AVMIDIControlChangeMessageTypeLegatoPedal AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeModWheel - An event type for modulating a vibrato effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/modWheel
	AVMIDIControlChangeMessageTypeModWheel AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeMonoModeOff - An event type for setting the device mode to polyphonic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/monoModeOff
	AVMIDIControlChangeMessageTypeMonoModeOff AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeMonoModeOn - An event type for setting the device mode to monophonic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/monoModeOn
	AVMIDIControlChangeMessageTypeMonoModeOn AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeOmniModeOff - An event type for setting omni off mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/omniModeOff
	AVMIDIControlChangeMessageTypeOmniModeOff AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeOmniModeOn - An event type for setting omni on mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/omniModeOn
	AVMIDIControlChangeMessageTypeOmniModeOn AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypePan - An event type for controlling the left and right channel pan.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/pan
	AVMIDIControlChangeMessageTypePan AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypePortamento - An event type for switching portamento on or off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/portamento
	AVMIDIControlChangeMessageTypePortamento AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypePortamentoTime - An event type for controlling the portamento rate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/portamentoTime
	AVMIDIControlChangeMessageTypePortamentoTime AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeReleaseTime - An event type for controlling the release time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/releaseTime
	AVMIDIControlChangeMessageTypeReleaseTime AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeResetAllControllers - An event type for resetting all controllers to their default state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/resetAllControllers
	AVMIDIControlChangeMessageTypeResetAllControllers AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeReverbLevel - An event type for controlling the reverb level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/reverbLevel
	AVMIDIControlChangeMessageTypeReverbLevel AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeRPN_LSB - An event type that represents the registered parameter number LSB.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/RPN_LSB
	AVMIDIControlChangeMessageTypeRPN_LSB AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeRPN_MSB - An event type that represents the registered parameter number MSB.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/RPN_MSB
	AVMIDIControlChangeMessageTypeRPN_MSB AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeSoft - An event type for lowering the volume of the notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/soft
	AVMIDIControlChangeMessageTypeSoft AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeSostenuto - An event type for switching sostenuto on or off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/sostenuto
	AVMIDIControlChangeMessageTypeSostenuto AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeSustain - An event type for switching a damper pedal on or off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/sustain
	AVMIDIControlChangeMessageTypeSustain AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeVibratoDelay - An event type for controlling the vibrato delay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/vibratoDelay
	AVMIDIControlChangeMessageTypeVibratoDelay AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeVibratoDepth - An event type for controlling the vibrato depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/vibratoDepth
	AVMIDIControlChangeMessageTypeVibratoDepth AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeVibratoRate - An event type for controlling the vibrato rate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/vibratoRate
	AVMIDIControlChangeMessageTypeVibratoRate AVMIDIControlChangeMessageType = 0
	// AVMIDIControlChangeMessageTypeVolume - An event type for controlling the channel volume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/MessageType-swift.enum/volume
	AVMIDIControlChangeMessageTypeVolume AVMIDIControlChangeMessageType = 0
)


// AVMIDIMetaEventType - Constants that represent the types of meta events.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType
type AVMIDIMetaEventType uint

const (
	// AVMIDIMetaEventTypeCopyright - An event type that represents a copyright.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/copyright
	AVMIDIMetaEventTypeCopyright AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeCuePoint - An event type that represents a cue point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/cuePoint
	AVMIDIMetaEventTypeCuePoint AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeEndOfTrack - An event type that represents the end of the track.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/endOfTrack
	AVMIDIMetaEventTypeEndOfTrack AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeInstrument - An event type that represents an instrument.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/instrument
	AVMIDIMetaEventTypeInstrument AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeKeySignature - An event type that represents a key signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/keySignature
	AVMIDIMetaEventTypeKeySignature AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeLyric - An event type that represents a lyric.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/lyric
	AVMIDIMetaEventTypeLyric AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeMarker - An event type that represents a marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/marker
	AVMIDIMetaEventTypeMarker AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeMidiChannel - An event type that represents a MIDI channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/midiChannel
	AVMIDIMetaEventTypeMidiChannel AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeMidiPort - An event type that represents a MIDI port.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/midiPort
	AVMIDIMetaEventTypeMidiPort AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeProprietaryEvent - An event type that represents a proprietary event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/proprietaryEvent
	AVMIDIMetaEventTypeProprietaryEvent AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeSequenceNumber - An event type that represents a sequence number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/sequenceNumber
	AVMIDIMetaEventTypeSequenceNumber AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeSmpteOffset - An event type that represents a SMPTE time offset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/smpteOffset
	AVMIDIMetaEventTypeSmpteOffset AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeTempo - An event type that represents a tempo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/tempo
	AVMIDIMetaEventTypeTempo AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeText - An event type that represents text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/text
	AVMIDIMetaEventTypeText AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeTimeSignature - An event type that represents a time signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/timeSignature
	AVMIDIMetaEventTypeTimeSignature AVMIDIMetaEventType = 0
	// AVMIDIMetaEventTypeTrackName - An event type that represents a track name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType/trackName
	AVMIDIMetaEventTypeTrackName AVMIDIMetaEventType = 0
)


// AVMusicSequenceLoadOptions - A structure that defines whether data on different MIDI channels map to multiple tracks, or whether the framework preserves the tracks as they are.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicSequenceLoadOptions
type AVMusicSequenceLoadOptions uint

const (
	// AVMusicSequenceLoadSMF_PreserveTracks - An option that preserves the tracks as they are.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicSequenceLoadOptions/AVMusicSequenceLoadSMF_PreserveTracks
	AVMusicSequenceLoadSMF_PreserveTracks AVMusicSequenceLoadOptions = 0
	// AVMusicSequenceLoadSMF_ChannelsToTracks - An option that represents data on different MIDI channels mapped to multiple tracks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicSequenceLoadOptions/smf_ChannelsToTracks
	AVMusicSequenceLoadSMF_ChannelsToTracks AVMusicSequenceLoadOptions = 0
)


// AVMusicTrackLoopCount - Options that define the number of times a track loops.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackLoopCount
type AVMusicTrackLoopCount uint

const (
	// AVMusicTrackLoopCountForever - A track that loops forever.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackLoopCount/forever
	AVMusicTrackLoopCountForever AVMusicTrackLoopCount = 0
)


// AVSpeechBoundary - Specifies when to pause or stop speech.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechBoundary
type AVSpeechBoundary uint

const (
	// AVSpeechBoundaryImmediate - Indicates to pause or stop speech immediately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechBoundary/immediate
	AVSpeechBoundaryImmediate AVSpeechBoundary = 0
	// AVSpeechBoundaryWord - Indicates to pause or stop speech after the synthesizer finishes speaking the current word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechBoundary/word
	AVSpeechBoundaryWord AVSpeechBoundary = 0
)


// AVSpeechSynthesisMarkerMark - Constants that describe the type of text.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum
type AVSpeechSynthesisMarkerMark uint

const (
	// AVSpeechSynthesisMarkerMarkBookmark - A Speech Synthesis Markup Language (SSML) mark tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum/bookmark
	AVSpeechSynthesisMarkerMarkBookmark AVSpeechSynthesisMarkerMark = 0
	// AVSpeechSynthesisMarkerMarkParagraph - A type of text that represents a paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum/paragraph
	AVSpeechSynthesisMarkerMarkParagraph AVSpeechSynthesisMarkerMark = 0
	// AVSpeechSynthesisMarkerMarkPhoneme - A type of text that represents a phoneme.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum/phoneme
	AVSpeechSynthesisMarkerMarkPhoneme AVSpeechSynthesisMarkerMark = 0
	// AVSpeechSynthesisMarkerMarkSentence - A type of text that represents a sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum/sentence
	AVSpeechSynthesisMarkerMarkSentence AVSpeechSynthesisMarkerMark = 0
	// AVSpeechSynthesisMarkerMarkWord - A type of text that represents a word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/Mark-swift.enum/word
	AVSpeechSynthesisMarkerMarkWord AVSpeechSynthesisMarkerMark = 0
)


// AVSpeechSynthesisVoiceTraits - Traits that describe a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits
type AVSpeechSynthesisVoiceTraits uint

const (
	// AVSpeechSynthesisVoiceTraitIsNoveltyVoice - The trait that indicates a voice is a novelty voice.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits/isNoveltyVoice
	AVSpeechSynthesisVoiceTraitIsNoveltyVoice AVSpeechSynthesisVoiceTraits = 0
	// AVSpeechSynthesisVoiceTraitIsPersonalVoice - The trait that indicates a voice is a personal voice.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoice/Traits/isPersonalVoice
	AVSpeechSynthesisVoiceTraitIsPersonalVoice AVSpeechSynthesisVoiceTraits = 0
	// AVSpeechSynthesisVoiceTraitNone - The trait that indicates a voice is a regular voice.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceTraits/AVSpeechSynthesisVoiceTraitNone
	AVSpeechSynthesisVoiceTraitNone AVSpeechSynthesisVoiceTraits = 0
)


// AVSpeechSynthesisVoiceGender - The gender for a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceGender
type AVSpeechSynthesisVoiceGender uint

const (
	// AVSpeechSynthesisVoiceGenderFemale - The female voice option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceGender/female
	AVSpeechSynthesisVoiceGenderFemale AVSpeechSynthesisVoiceGender = 0
	// AVSpeechSynthesisVoiceGenderMale - The male voice option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceGender/male
	AVSpeechSynthesisVoiceGenderMale AVSpeechSynthesisVoiceGender = 0
	// AVSpeechSynthesisVoiceGenderUnspecified - The nonspecific gender option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceGender/unspecified
	AVSpeechSynthesisVoiceGenderUnspecified AVSpeechSynthesisVoiceGender = 0
)


// AVSpeechSynthesisVoiceQuality - The speech quality of a voice.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceQuality
type AVSpeechSynthesisVoiceQuality uint

const (
	// AVSpeechSynthesisVoiceQualityDefault - A basic quality voice that’s  available on the device by default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceQuality/default
	AVSpeechSynthesisVoiceQualityDefault AVSpeechSynthesisVoiceQuality = 0
	// AVSpeechSynthesisVoiceQualityEnhanced - An enhanced quality voice that you must download to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceQuality/enhanced
	AVSpeechSynthesisVoiceQualityEnhanced AVSpeechSynthesisVoiceQuality = 0
	// AVSpeechSynthesisVoiceQualityPremium - A premium quality voice that you must download to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisVoiceQuality/premium
	AVSpeechSynthesisVoiceQualityPremium AVSpeechSynthesisVoiceQuality = 0
)


// AVSpeechSynthesisPersonalVoiceAuthorizationStatus - An enumeration that models the personal voices authorization status.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum
type AVSpeechSynthesisPersonalVoiceAuthorizationStatus uint

const (
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized - The user granted your app’s request to use personal voices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum/authorized
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusAuthorized AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 0
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied - The user denied your app’s request to use personal voices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum/denied
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusDenied AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 0
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined - The app hasn’t requested authorization to use personal voices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum/notDetermined
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusNotDetermined AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 0
	// AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported - The device doesn’t support personal voices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/PersonalVoiceAuthorizationStatus-swift.enum/unsupported
	AVSpeechSynthesisPersonalVoiceAuthorizationStatusUnsupported AVSpeechSynthesisPersonalVoiceAuthorizationStatus = 0
)


