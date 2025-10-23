// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

// Enum types and constants
// AU3DMixerAttenuationCurve enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve
type AU3DMixerAttenuationCurve uint

const (
	// k3DMixerAttenuationCurve_Exponential - An exponential attenuation curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve/k3DMixerAttenuationCurve_Exponential
	k3DMixerAttenuationCurve_Exponential AU3DMixerAttenuationCurve = 0
	// k3DMixerAttenuationCurve_Inverse - An inverse attenuation curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve/k3DMixerAttenuationCurve_Inverse
	k3DMixerAttenuationCurve_Inverse AU3DMixerAttenuationCurve = 0
	// k3DMixerAttenuationCurve_Linear - A linear attenuation curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve/k3DMixerAttenuationCurve_Linear
	k3DMixerAttenuationCurve_Linear AU3DMixerAttenuationCurve = 0
	// k3DMixerAttenuationCurve_Power - An equal-power-based attenuation curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerAttenuationCurve/k3DMixerAttenuationCurve_Power
	k3DMixerAttenuationCurve_Power AU3DMixerAttenuationCurve = 0
)

// AU3DMixerRenderingFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AU3DMixerRenderingFlags
type AU3DMixerRenderingFlags uint

// AUAudioMixRenderingStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle
type AudioMixRenderingStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_Cinematic
	kAudioMixRenderingStyle_Cinematic AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_CinematicBackgroundStem
	kAudioMixRenderingStyle_CinematicBackgroundStem AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_CinematicForegroundStem
	kAudioMixRenderingStyle_CinematicForegroundStem AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_InFrame
	kAudioMixRenderingStyle_InFrame AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_InFrameBackgroundStem
	kAudioMixRenderingStyle_InFrameBackgroundStem AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_InFrameForegroundStem
	kAudioMixRenderingStyle_InFrameForegroundStem AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_Standard
	kAudioMixRenderingStyle_Standard AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_Studio
	kAudioMixRenderingStyle_Studio AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_StudioBackgroundStem
	kAudioMixRenderingStyle_StudioBackgroundStem AudioMixRenderingStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioMixRenderingStyle/audioMixRenderingStyle_StudioForegroundStem
	kAudioMixRenderingStyle_StudioForegroundStem AudioMixRenderingStyle = 0
)

// AUAudioUnitBusType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusType
type AudioUnitBusType uint

const (
	// AudioUnitBusTypeInput - An input bus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusType/input
	AudioUnitBusTypeInput AudioUnitBusType = 0
	// AudioUnitBusTypeOutput - An output bus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusType/output
	AudioUnitBusTypeOutput AudioUnitBusType = 0
)

// AUHostTransportStateFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostTransportStateFlags
type HostTransportStateFlags uint

// AUParameterAutomationEventType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEventType
type ParameterAutomationEventType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEventType/release
	ParameterAutomationEventTypeRelease ParameterAutomationEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEventType/touch
	ParameterAutomationEventTypeTouch ParameterAutomationEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEventType/value
	ParameterAutomationEventTypeValue ParameterAutomationEventType = 0
)

// AUParameterEventType - Audio unit parameter event types.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEventType
type ParameterEventType uint

const (
	// kParameterEvent_Immediate - An immediate change from the parameter’s previous value to a new value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEventType/parameterEvent_Immediate
	kParameterEvent_Immediate ParameterEventType = 0
	// kParameterEvent_Ramped - A gradual change from the parameter’s previous value to a new value, applied linearly over a specified period of time
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEventType/parameterEvent_Ramped
	kParameterEvent_Ramped ParameterEventType = 0
)

// AUParameterMIDIMappingFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMappingFlags
type ParameterMIDIMappingFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMappingFlags/bipolar_On
	kAUParameterMIDIMapping_Bipolar_On ParameterMIDIMappingFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMappingFlags/subRange
	kAUParameterMIDIMapping_SubRange ParameterMIDIMappingFlags = 0
)

// AURenderEventType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderEventType
type RenderEventType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderEventType/midiEventList
	RenderEventMIDIEventList RenderEventType = 0
)

// AUReverbRoomType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType
type ReverbRoomType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_Cathedral
	kReverbRoomType_Cathedral ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_LargeChamber
	kReverbRoomType_LargeChamber ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_LargeHall
	kReverbRoomType_LargeHall ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_LargeHall2
	kReverbRoomType_LargeHall2 ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_LargeRoom
	kReverbRoomType_LargeRoom ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_LargeRoom2
	kReverbRoomType_LargeRoom2 ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_MediumChamber
	kReverbRoomType_MediumChamber ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_MediumHall
	kReverbRoomType_MediumHall ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_MediumHall2
	kReverbRoomType_MediumHall2 ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_MediumHall3
	kReverbRoomType_MediumHall3 ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_MediumRoom
	kReverbRoomType_MediumRoom ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_Plate
	kReverbRoomType_Plate ReverbRoomType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUReverbRoomType/reverbRoomType_SmallRoom
	kReverbRoomType_SmallRoom ReverbRoomType = 0
)

// AUScheduledAudioSliceFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUScheduledAudioSliceFlags
type ScheduledAudioSliceFlags uint

// AUSpatialMixerAttenuationCurve enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve
type SpatialMixerAttenuationCurve uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve/spatialMixerAttenuationCurve_Exponential
	kSpatialMixerAttenuationCurve_Exponential SpatialMixerAttenuationCurve = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve/spatialMixerAttenuationCurve_Inverse
	kSpatialMixerAttenuationCurve_Inverse SpatialMixerAttenuationCurve = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve/spatialMixerAttenuationCurve_Linear
	kSpatialMixerAttenuationCurve_Linear SpatialMixerAttenuationCurve = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerAttenuationCurve/spatialMixerAttenuationCurve_Power
	kSpatialMixerAttenuationCurve_Power SpatialMixerAttenuationCurve = 0
)

// AUSpatialMixerOutputType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerOutputType
type SpatialMixerOutputType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerOutputType/spatialMixerOutputType_BuiltInSpeakers
	kSpatialMixerOutputType_BuiltInSpeakers SpatialMixerOutputType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerOutputType/spatialMixerOutputType_ExternalSpeakers
	kSpatialMixerOutputType_ExternalSpeakers SpatialMixerOutputType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerOutputType/spatialMixerOutputType_Headphones
	kSpatialMixerOutputType_Headphones SpatialMixerOutputType = 0
)

// AUSpatialMixerPersonalizedHRTFMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPersonalizedHRTFMode
type SpatialMixerPersonalizedHRTFMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPersonalizedHRTFMode/auto
	kSpatialMixerPersonalizedHRTFMode_Auto SpatialMixerPersonalizedHRTFMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPersonalizedHRTFMode/off
	kSpatialMixerPersonalizedHRTFMode_Off SpatialMixerPersonalizedHRTFMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPersonalizedHRTFMode/on
	kSpatialMixerPersonalizedHRTFMode_On SpatialMixerPersonalizedHRTFMode = 0
)

// AUSpatialMixerPointSourceInHeadMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPointSourceInHeadMode
type SpatialMixerPointSourceInHeadMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPointSourceInHeadMode/spatialMixerPointSourceInHeadMode_Bypass
	kSpatialMixerPointSourceInHeadMode_Bypass SpatialMixerPointSourceInHeadMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerPointSourceInHeadMode/spatialMixerPointSourceInHeadMode_Mono
	kSpatialMixerPointSourceInHeadMode_Mono SpatialMixerPointSourceInHeadMode = 0
)

// AUSpatialMixerRenderingFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerRenderingFlags
type SpatialMixerRenderingFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerRenderingFlags/spatialMixerRenderingFlags_DistanceAttenuation
	kSpatialMixerRenderingFlags_DistanceAttenuation SpatialMixerRenderingFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerRenderingFlags/spatialMixerRenderingFlags_InterAuralDelay
	kSpatialMixerRenderingFlags_InterAuralDelay SpatialMixerRenderingFlags = 0
)

// AUSpatialMixerSourceMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode
type SpatialMixerSourceMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode/spatialMixerSourceMode_AmbienceBed
	kSpatialMixerSourceMode_AmbienceBed SpatialMixerSourceMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode/spatialMixerSourceMode_Bypass
	kSpatialMixerSourceMode_Bypass SpatialMixerSourceMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode/spatialMixerSourceMode_PointSource
	kSpatialMixerSourceMode_PointSource SpatialMixerSourceMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatialMixerSourceMode/spatialMixerSourceMode_SpatializeIfMono
	kSpatialMixerSourceMode_SpatializeIfMono SpatialMixerSourceMode = 0
)

// AUSpatializationAlgorithm enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatializationAlgorithm
type SpatializationAlgorithm uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatializationAlgorithm/spatializationAlgorithm_SphericalHead
	kSpatializationAlgorithm_SphericalHead SpatializationAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSpatializationAlgorithm/spatializationAlgorithm_UseOutputType
	kSpatializationAlgorithm_UseOutputType SpatializationAlgorithm = 0
)

// AUVoiceIOOtherAudioDuckingLevel - The ducking level to apply to other non-voice audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingLevel
type VoiceIOOtherAudioDuckingLevel uint

const (
	// kAUVoiceIOOtherAudioDuckingLevelDefault - The default ducking level of other non-voice audio in a typical voice chat.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingLevel/default
	kAUVoiceIOOtherAudioDuckingLevelDefault VoiceIOOtherAudioDuckingLevel = 0
)

// AUVoiceIOSpeechActivityEvent - Constants that indicate the state of muted speech.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOSpeechActivityEvent
type VoiceIOSpeechActivityEvent uint

const (
	// kAUVoiceIOSpeechActivityHasEnded - A state that indicates speech ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOSpeechActivityEvent/hasEnded
	kAUVoiceIOSpeechActivityHasEnded VoiceIOSpeechActivityEvent = 0
	// kAUVoiceIOSpeechActivityHasStarted - A state that indicates speech started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOSpeechActivityEvent/hasStarted
	kAUVoiceIOSpeechActivityHasStarted VoiceIOSpeechActivityEvent = 0
)

// AudioBalanceFadeType - Identifiers for audio balance fade types.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFadeType
type AudioBalanceFadeType uint

const (
	// kAudioBalanceFadeType_EqualPower - Overall loudness remains constant, but gain can exceed 1.0. The gain value is 1.0 when the balance and fade are in the center. From there they can increase to +3dB (1.414) and decrease to silence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFadeType/equalPower
	kAudioBalanceFadeType_EqualPower AudioBalanceFadeType = 0
	// kAudioBalanceFadeType_MaxUnityGain - Ensures that the overall gain value never exceeds 1.0 by fading one channel as the other channel’s level rises. This can reduce overall loudness when the balance or fade is not in the center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFadeType/maxUnityGain
	kAudioBalanceFadeType_MaxUnityGain AudioBalanceFadeType = 0
)

// AudioBytePacketTranslationFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslationFlags
type AudioBytePacketTranslationFlags uint

const (
	// kBytePacketTranslationFlag_IsEstimate - If set, the result value is an estimate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslationFlags/bytePacketTranslationFlag_IsEstimate
	kBytePacketTranslationFlag_IsEstimate AudioBytePacketTranslationFlags = 0
)

// AudioComponentFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags
type AudioComponentFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/canLoadInProcess
	kAudioComponentFlag_CanLoadInProcess AudioComponentFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/isV3AudioUnit
	kAudioComponentFlag_IsV3AudioUnit AudioComponentFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/requiresAsyncInstantiation
	kAudioComponentFlag_RequiresAsyncInstantiation AudioComponentFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/sandboxSafe
	kAudioComponentFlag_SandboxSafe AudioComponentFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/unsearchable
	kAudioComponentFlag_Unsearchable AudioComponentFlags = 0
)

// AudioComponentInstantiationOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiationOptions
type AudioComponentInstantiationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiationOptions/loadOutOfProcess
	kAudioComponentInstantiation_LoadOutOfProcess AudioComponentInstantiationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiationOptions/loadedRemotely
	kAudioComponentInstantiation_LoadedRemotely AudioComponentInstantiationOptions = 0
)

// AudioComponentValidationResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult
type AudioComponentValidationResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/failed
	kAudioComponentValidationResult_Failed AudioComponentValidationResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/passed
	kAudioComponentValidationResult_Passed AudioComponentValidationResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/timedOut
	kAudioComponentValidationResult_TimedOut AudioComponentValidationResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/unauthorizedError_Init
	kAudioComponentValidationResult_UnauthorizedError_Init AudioComponentValidationResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/unauthorizedError_Open
	kAudioComponentValidationResult_UnauthorizedError_Open AudioComponentValidationResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidationResult/unknown
	kAudioComponentValidationResult_Unknown AudioComponentValidationResult = 0
)

// AudioConverterOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterOptions
type AudioConverterOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterOptions/unbuffered
	kAudioConverterOption_Unbuffered AudioConverterOptions = 0
)

// AudioFileFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFlags
type AudioFileFlags uint

const (
	// kAudioFileFlags_DontPageAlignAudioData - Typically, the audio data in a file is page aligned. To make reading the file data as fast as possible, you can use page-aligned data to take advantage of optimized code paths in the file system. However, when space is at a premium, you might want to avoid the additional padding required to attain alignment. To do so, set this flag when calling   or  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFlags/dontPageAlignAudioData
	kAudioFileFlags_DontPageAlignAudioData AudioFileFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFlags/eraseFile
	kAudioFileFlags_EraseFile AudioFileFlags = 0
)

// AudioFilePermissions - Flags for use when opening an audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePermissions
type AudioFilePermissions uint

const (
	// kAudioFileReadPermission - File is read-only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePermissions/readPermission
	kAudioFileReadPermission AudioFilePermissions = 0
	// kAudioFileReadWritePermission - File has read-write permission.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePermissions/readWritePermission
	kAudioFileReadWritePermission AudioFilePermissions = 0
	// kAudioFileWritePermission - File is write-only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePermissions/writePermission
	kAudioFileWritePermission AudioFilePermissions = 0
)

// AudioFileRegionFlags - Flags that specify a playback direction for an audio file region structure.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionFlags
type AudioFileRegionFlags uint

const (
	// kAudioFileRegionFlag_LoopEnable - If set, the region is looped. You must set one or both of the remaining flags must also be set for the region to be looped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionFlags/loopEnable
	kAudioFileRegionFlag_LoopEnable AudioFileRegionFlags = 0
	// kAudioFileRegionFlag_PlayForward - If set, the region is played forward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionFlags/playForward
	kAudioFileRegionFlag_PlayForward AudioFileRegionFlags = 0
)

// AudioFileStreamParseFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseFlags
type AudioFileStreamParseFlags uint

const (
	// kAudioFileStreamParseFlag_Discontinuity - Pass this flag to the   function to signal a discontinuity in the audio data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseFlags/discontinuity
	kAudioFileStreamParseFlag_Discontinuity AudioFileStreamParseFlags = 0
)

// AudioFileStreamPropertyFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyFlags
type AudioFileStreamPropertyFlags uint

const (
	// kAudioFileStreamPropertyFlag_CacheProperty - A property listener sets this flag to instruct the parser to cache the property value so that it remains available after the callback returns.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyFlags/cacheProperty
	kAudioFileStreamPropertyFlag_CacheProperty AudioFileStreamPropertyFlags = 0
	// kAudioFileStreamPropertyFlag_PropertyIsCached - This flag is set when the callback   is invoked in the case that the value of the property has been cached and can be obtained later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyFlags/propertyIsCached
	kAudioFileStreamPropertyFlag_PropertyIsCached AudioFileStreamPropertyFlags = 0
)

// AudioFileStreamSeekFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeekFlags
type AudioFileStreamSeekFlags uint

const (
	// kAudioFileStreamSeekFlag_OffsetIsEstimated - This flag is returned by the   function if the byte offset is only an estimate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeekFlags/offsetIsEstimated
	kAudioFileStreamSeekFlag_OffsetIsEstimated AudioFileStreamSeekFlags = 0
)

// AudioPanningMode - Identifiers for audio panning algorithms.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningMode
type AudioPanningMode uint

const (
	// kPanningMode_SoundField - The SoundField panning algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningMode/panningMode_SoundField
	kPanningMode_SoundField AudioPanningMode = 0
	// kPanningMode_VectorBasedPanning - A vector-based panning algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningMode/panningMode_VectorBasedPanning
	kPanningMode_VectorBasedPanning AudioPanningMode = 0
)

// AudioQueueProcessingTapFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapFlags
type AudioQueueProcessingTapFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapFlags/siphon
	kAudioQueueProcessingTap_Siphon AudioQueueProcessingTapFlags = 0
)

// AudioSettingsFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags
type AudioSettingsFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags/expertParameter
	kAudioSettingsFlags_ExpertParameter AudioSettingsFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags/invisibleParameter
	kAudioSettingsFlags_InvisibleParameter AudioSettingsFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags/metaParameter
	kAudioSettingsFlags_MetaParameter AudioSettingsFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSettingsFlags/userInterfaceParameter
	kAudioSettingsFlags_UserInterfaceParameter AudioSettingsFlags = 0
)

// AudioUnitEventType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType
type AudioUnitEventType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType/beginParameterChangeGesture
	kAudioUnitEvent_BeginParameterChangeGesture AudioUnitEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType/endParameterChangeGesture
	kAudioUnitEvent_EndParameterChangeGesture AudioUnitEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType/parameterValueChange
	kAudioUnitEvent_ParameterValueChange AudioUnitEventType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEventType/propertyChange
	kAudioUnitEvent_PropertyChange AudioUnitEventType = 0
)

// AudioUnitParameterOptions - Value options for audio unit parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterOptions
type AudioUnitParameterOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterOptions/flag_OmitFromPresets
	kAudioUnitParameterFlag_OmitFromPresets AudioUnitParameterOptions = 0
	// kAudioUnitParameterFlag_PlotHistory - If set, getting the   property fills out the   struct containing the recommended update rate and history duration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterOptions/flag_PlotHistory
	kAudioUnitParameterFlag_PlotHistory AudioUnitParameterOptions = 0
)

// AudioUnitParameterUnit - The unit-of-measure for an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit
type AudioUnitParameterUnit uint

const (
	// kAudioUnitParameterUnit_BPM - A whole-number unit of measure for musical tempo, representing beats per minute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/BPM
	kAudioUnitParameterUnit_BPM AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_AbsoluteCents - An absolute unit of measure for the musical pitch of a note.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/absoluteCents
	kAudioUnitParameterUnit_AbsoluteCents AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Beats - A time unit of measure in musical beats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/beats
	kAudioUnitParameterUnit_Beats AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Boolean - A Boolean-like unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/boolean
	kAudioUnitParameterUnit_Boolean AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Cents - A logarithmic unit of measure for a musical interval between two notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/cents
	kAudioUnitParameterUnit_Cents AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_CustomUnit - A custom unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/customUnit
	kAudioUnitParameterUnit_CustomUnit AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Decibels - A logarithmic unit of measure representing the ratio between two audio levels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/decibels
	kAudioUnitParameterUnit_Decibels AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Degrees - An angular degree unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/degrees
	kAudioUnitParameterUnit_Degrees AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_EqualPowerCrossfade - An audio power unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/equalPowerCrossfade
	kAudioUnitParameterUnit_EqualPowerCrossfade AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Generic - A generic unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/generic
	kAudioUnitParameterUnit_Generic AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Hertz - A hertz unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/hertz
	kAudioUnitParameterUnit_Hertz AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Indexed - An indexed unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/indexed
	kAudioUnitParameterUnit_Indexed AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_LinearGain - A linear unit of measure representing the difference between two audio levels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/linearGain
	kAudioUnitParameterUnit_LinearGain AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Meters - A distance unit of measure, corresponding to meters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/meters
	kAudioUnitParameterUnit_Meters AudioUnitParameterUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/midi2Controller
	kAudioUnitParameterUnit_MIDI2Controller AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_MIDIController - A whole-number unit of measure corresponding to standard MIDI control numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/midiController
	kAudioUnitParameterUnit_MIDIController AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Milliseconds - A time unit of measure representing milliseconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/milliseconds
	kAudioUnitParameterUnit_Milliseconds AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_MixerFaderCurve1 - An audio power unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/mixerFaderCurve1
	kAudioUnitParameterUnit_MixerFaderCurve1 AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Octaves - A relative unit of measure for the musical interval between two notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/octaves
	kAudioUnitParameterUnit_Octaves AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Pan - An audio position unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/pan
	kAudioUnitParameterUnit_Pan AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Percent - A percentage unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/percent
	kAudioUnitParameterUnit_Percent AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Phase - An angular degree unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/phase
	kAudioUnitParameterUnit_Phase AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Rate - A multiplication factor unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/rate
	kAudioUnitParameterUnit_Rate AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Ratio - A unitless ratio unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/ratio
	kAudioUnitParameterUnit_Ratio AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_RelativeSemiTones - A relative unit of measure for a musical interval between two notes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/relativeSemiTones
	kAudioUnitParameterUnit_RelativeSemiTones AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_SampleFrames - A sample-frame-count unit of measure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/sampleFrames
	kAudioUnitParameterUnit_SampleFrames AudioUnitParameterUnit = 0
	// kAudioUnitParameterUnit_Seconds - A whole-seconds unit of measure, indicating either absolute or relative time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterUnit/seconds
	kAudioUnitParameterUnit_Seconds AudioUnitParameterUnit = 0
)

// AudioUnitRemoteControlEvent enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEvent
type AudioUnitRemoteControlEvent uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEvent/rewind
	kAudioUnitRemoteControlEvent_Rewind AudioUnitRemoteControlEvent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEvent/togglePlayPause
	kAudioUnitRemoteControlEvent_TogglePlayPause AudioUnitRemoteControlEvent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoteControlEvent/toggleRecord
	kAudioUnitRemoteControlEvent_ToggleRecord AudioUnitRemoteControlEvent = 0
)

// AudioUnitRenderActionFlags - Flags for configuring audio unit rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderActionFlags
type AudioUnitRenderActionFlags uint

const (
	// kAudioUnitRenderAction_PreRender - Called on a render notification Proc - which is called either before or after the render operation of the audio unit. If this flag is set, the proc is being called before the render operation is performed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderActionFlags/unitRenderAction_PreRender
	kAudioUnitRenderAction_PreRender AudioUnitRenderActionFlags = 0
)

// CAClockMessage enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage
type ClockMessage uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/armed
	kCAClockMessage_Armed ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/disarmed
	kCAClockMessage_Disarmed ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/propertyChanged
	kCAClockMessage_PropertyChanged ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/startTimeSet
	kCAClockMessage_StartTimeSet ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/started
	kCAClockMessage_Started ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/stopped
	kCAClockMessage_Stopped ClockMessage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockMessage/wrongSMPTEFormat
	kCAClockMessage_WrongSMPTEFormat ClockMessage = 0
)

// CAClockPropertyID enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID
type ClockPropertyID uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/internalTimebase
	kCAClockProperty_InternalTimebase ClockPropertyID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/meterTrack
	kCAClockProperty_MeterTrack ClockPropertyID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/name
	kCAClockProperty_Name ClockPropertyID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/sendMIDISPP
	kCAClockProperty_SendMIDISPP ClockPropertyID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/syncMode
	kCAClockProperty_SyncMode ClockPropertyID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockPropertyID/timebaseSource
	kCAClockProperty_TimebaseSource ClockPropertyID = 0
)

// CAClockSyncMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSyncMode
type ClockSyncMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSyncMode/midiClockTransport
	kCAClockSyncMode_MIDIClockTransport ClockSyncMode = 0
)

// CAClockTimeFormat enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat
type ClockTimeFormat uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat/absoluteSeconds
	kCAClockTimeFormat_AbsoluteSeconds ClockTimeFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat/samples
	kCAClockTimeFormat_Samples ClockTimeFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat/seconds
	kCAClockTimeFormat_Seconds ClockTimeFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimeFormat/smpteTime
	kCAClockTimeFormat_SMPTETime ClockTimeFormat = 0
)

// CAClockTimebase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimebase
type ClockTimebase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTimebase/hostTime
	kCAClockTimebase_HostTime ClockTimebase = 0
)

// CAFFormatFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFormatFlags
type FFormatFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFormatFlags/linearPCMFormatFlagIsFloat
	kCAFLinearPCMFormatFlagIsFloat FFormatFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFormatFlags/linearPCMFormatFlagIsLittleEndian
	kCAFLinearPCMFormatFlagIsLittleEndian FFormatFlags = 0
)

// CAFRegionFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionFlags
type FRegionFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionFlags/loopEnable
	kCAFRegionFlag_LoopEnable FRegionFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionFlags/playBackward
	kCAFRegionFlag_PlayBackward FRegionFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionFlags/playForward
	kCAFRegionFlag_PlayForward FRegionFlags = 0
)

// CASoundStageSize - Configure the distribution of audio channels in 3D space.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASoundStageSize
type SoundStageSize uint

const (
	// SoundStageSizeAutomatic - A system-defined sound stage size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASoundStageSize/CASoundStageSizeAutomatic
	SoundStageSizeAutomatic SoundStageSize = 0
	// SoundStageSizeLarge - Spreads an audio stream’s channels around the user according to the   coordinates described in its channel layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASoundStageSize/CASoundStageSizeLarge
	SoundStageSizeLarge SoundStageSize = 0
	// SoundStageSizeMedium - Pulls an audio stream’s channels closer to the channel layout’s front.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASoundStageSize/CASoundStageSizeMedium
	SoundStageSizeMedium SoundStageSize = 0
	// SoundStageSizeSmall - Places all of an audio stream’s channels near the layout’s front.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASoundStageSize/CASoundStageSizeSmall
	SoundStageSizeSmall SoundStageSize = 0
)

// MusicSequenceFileFlags - Flags that configure the behavior of the 
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileFlags
type MusicSequenceFileFlags uint

const (
	// kMusicSequenceFileFlags_EraseFile - Specifies that an existing file should be erased when creating a new file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileFlags/eraseFile
	kMusicSequenceFileFlags_EraseFile MusicSequenceFileFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileFlags/kMusicSequenceFileFlags_Default
	kMusicSequenceFileFlags_Default MusicSequenceFileFlags = 0
)

// MusicSequenceFileTypeID - The various types of files that can be parsed by a music sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileTypeID
type MusicSequenceFileTypeID uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileTypeID/anyType
	kMusicSequenceFile_AnyType MusicSequenceFileTypeID = 0
	// kMusicSequenceFile_iMelodyType - An iMelody file type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileTypeID/iMelodyType
	kMusicSequenceFile_iMelodyType MusicSequenceFileTypeID = 0
	// kMusicSequenceFile_MIDIType - A MIDI file type
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileTypeID/midiType
	kMusicSequenceFile_MIDIType MusicSequenceFileTypeID = 0
)

// MusicSequenceLoadFlags - Flags used to configure the behavior of the 
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadFlags
type MusicSequenceLoadFlags uint

const (
	// kMusicSequenceLoadSMF_ChannelsToTracks - If this flag is set the resultant Sequence will contain a tempo track, 1 track for each MIDI Channel that is found in the SMF, 1 track for SysEx or MetaEvents - and this will be the last track in the sequence after the LoadSMFWithFlags calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadFlags/smf_ChannelsToTracks
	kMusicSequenceLoadSMF_ChannelsToTracks MusicSequenceLoadFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadFlags/smf_PreserveTracks
	kMusicSequenceLoadSMF_PreserveTracks MusicSequenceLoadFlags = 0
)

// MusicSequenceType - The various types of music sequences.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceType
type MusicSequenceType uint

const (
	// kMusicSequenceType_Beats - Used for a music sequence that corresponds to a normal MIDI file. The tempo track defines the number of beats per second and can have multiple tempo events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceType/beats
	kMusicSequenceType_Beats MusicSequenceType = 0
	// kMusicSequenceType_Samples - Used for audio samples; a music sequence of this type cannot be saved to a MIDI file. The tempo track contains a single tempo event that specifies an audio sample rate in samples-per-second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceType/samples
	kMusicSequenceType_Samples MusicSequenceType = 0
	// kMusicSequenceType_Seconds - Used for a music sequence that corresponds to a MIDI file, but employs SMPTE timecode. The tempo track contains a single tempo event that specifies 60 beat-per-minute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceType/seconds
	kMusicSequenceType_Seconds MusicSequenceType = 0
)


