// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

// Enum types and constants
// PHASEAssetType - Options that determine how PHASE manages sound assets in memory.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset/AssetType
type PHASEAssetType uint

const (
// PHASEAssetTypeResident - A sound asset that plays after fully loading in memory.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset/AssetType/resident
PHASEAssetTypeResident PHASEAssetType = 0
// PHASEAssetTypeStreamed - A sound asset that streams from disk into memory as it plays.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset/AssetType/streamed
PHASEAssetTypeStreamed PHASEAssetType = 0
)

// PHASEAssetError - Codes that identify framework asset errors.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code
type PHASEAssetError uint

// PHASEAutomaticHeadTrackingFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAutomaticHeadTrackingFlags
type PHASEAutomaticHeadTrackingFlags uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAutomaticHeadTrackingFlags/orientation
PHASEAutomaticHeadTrackingFlagOrientation PHASEAutomaticHeadTrackingFlags = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAutomaticHeadTrackingFlags/position
PHASEAutomaticHeadTrackingFlagPosition PHASEAutomaticHeadTrackingFlags = 0
)

// PHASECalibrationMode - Calibration options for sound pressure level.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECalibrationMode
type PHASECalibrationMode uint

// PHASECullOption - The actions the engine takes when it culls sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption
type PHASECullOption uint

// PHASECurveType - Options that apply a mathematical function to an input value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType
type PHASECurveType uint

const (
// PHASECurveTypeCubed - A curve that increases at a rate that cubes its input.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/cubed
PHASECurveTypeCubed PHASECurveType = 0
// PHASECurveTypeHoldStartValue - A curve that equals its start value for the entire duration.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/holdStartValue
PHASECurveTypeHoldStartValue PHASECurveType = 0
// PHASECurveTypeInverseCubed - A curve that increases at a rate of one divided by the input’s cube.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/inverseCubed
PHASECurveTypeInverseCubed PHASECurveType = 0
// PHASECurveTypeInverseSigmoid - An inverse sigmoid curve.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/inverseSigmoid
PHASECurveTypeInverseSigmoid PHASECurveType = 0
// PHASECurveTypeInverseSine - An inverse sine curve.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/inverseSine
PHASECurveTypeInverseSine PHASECurveType = 0
// PHASECurveTypeInverseSquared - A curve that increases at a rate of one divided by the input’s square.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/inverseSquared
PHASECurveTypeInverseSquared PHASECurveType = 0
// PHASECurveTypeJumpToEndValue - A curve that equals its end value for the entire duration.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/jumpToEndValue
PHASECurveTypeJumpToEndValue PHASECurveType = 0
// PHASECurveTypeLinear - A curve that increases uniformly with its input.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/linear
PHASECurveTypeLinear PHASECurveType = 0
// PHASECurveTypeSigmoid - A sigmoid curve.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/sigmoid
PHASECurveTypeSigmoid PHASECurveType = 0
// PHASECurveTypeSine - A sine curve.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/sine
PHASECurveTypeSine PHASECurveType = 0
// PHASECurveTypeSquared - A curve that increases at a rate that squares its input.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECurveType/squared
PHASECurveTypeSquared PHASECurveType = 0
)

// PHASERenderingMode - Modes that determine whether the system renders audio in process or out of process.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/RenderingMode
type PHASERenderingMode uint

const (
// PHASERenderingModeClient - A mode that instructs the system to render audio in a secure process.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/RenderingMode/client
PHASERenderingModeClient PHASERenderingMode = 0
// PHASERenderingModeLocal - A mode that indicates that the system renders audio in process.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/RenderingMode/local
PHASERenderingModeLocal PHASERenderingMode = 0
)

// PHASEUpdateMode - Modes that determine when the framework consumes API calls and updates internal state.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/UpdateMode
type PHASEUpdateMode uint

const (
// PHASEUpdateModeAutomatic - A mode that indicates PHASE sets the timing of state adjustments.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/UpdateMode/automatic
PHASEUpdateModeAutomatic PHASEUpdateMode = 0
// PHASEUpdateModeManual - A mode that indicates the app controls when the framework adjusts state.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/UpdateMode/manual
PHASEUpdateModeManual PHASEUpdateMode = 0
)

// PHASEError - Codes that identify errors in PHASE.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEError-swift.struct/Code
type PHASEError uint

// PHASEMaterialPreset - A collection of physical surfaces that each add a unique acoustic quality to your app’s audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset
type PHASEMaterialPreset uint

const (
// PHASEMaterialPresetBrick - A surface characteristic that produces the acoustic quality of brick.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/brick
PHASEMaterialPresetBrick PHASEMaterialPreset = 0
// PHASEMaterialPresetCardboard - A surface characteristic that produces the acoustic quality of cardboard.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/cardboard
PHASEMaterialPresetCardboard PHASEMaterialPreset = 0
// PHASEMaterialPresetConcrete - A surface characteristic that produces the acoustic quality of concrete.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/concrete
PHASEMaterialPresetConcrete PHASEMaterialPreset = 0
// PHASEMaterialPresetDrywall - A surface characteristic that produces the acoustic quality of drywall.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/drywall
PHASEMaterialPresetDrywall PHASEMaterialPreset = 0
// PHASEMaterialPresetGlass - A surface characteristic that produces the acoustic quality of glass.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/glass
PHASEMaterialPresetGlass PHASEMaterialPreset = 0
// PHASEMaterialPresetWood - A surface characteristic that produces the acoustic quality of wood.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterialPreset/wood
PHASEMaterialPresetWood PHASEMaterialPreset = 0
)

// PHASEMediumPreset - Predetermined qualities of an environment that affect how sound transmits.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMedium/Preset
type PHASEMediumPreset uint

const (
// PHASEMediumPresetAir - A medium that simulates sound traveling through air.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMedium/Preset/air
PHASEMediumPresetAir PHASEMediumPreset = 0
)

// PHASENormalizationMode - Options that determine whether the framework adjusts a sound asset’s loudness for the user’s output device.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENormalizationMode
type PHASENormalizationMode uint

const (
// PHASENormalizationModeDynamic - A mode that instructs the framework to adjust a sound’s volume according to the user’s output device.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENormalizationMode/dynamic
PHASENormalizationModeDynamic PHASENormalizationMode = 0
// PHASENormalizationModeNone - A mode that instructs the framework not to adjust a sound’s volume according to the user’s output device.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENormalizationMode/none
PHASENormalizationModeNone PHASENormalizationMode = 0
)

// PHASEPlaybackMode - Loop options for audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPlaybackMode
type PHASEPlaybackMode uint

const (
// PHASEPlaybackModeLooping - An option that restarts a sound from the begining after it finishes.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPlaybackMode/looping
PHASEPlaybackModeLooping PHASEPlaybackMode = 0
)

// PHASEPushStreamBufferOptions - Options that inform PHASE of an audio-stream buffer’s playback priority.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions
type PHASEPushStreamBufferOptions uint

// PHASEReverbPreset - The manner in which PHASE diffuses resonating sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset
type PHASEReverbPreset uint

const (
// PHASEReverbPresetCathedral - A resonation that simulates the experience of hearing a sound in a cathedral.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/cathedral
PHASEReverbPresetCathedral PHASEReverbPreset = 0
// PHASEReverbPresetLargeChamber - A resonation that simulates the experience of hearing a sound in a large chamber with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/largeChamber
PHASEReverbPresetLargeChamber PHASEReverbPreset = 0
// PHASEReverbPresetLargeHall - A resonation that simulates the experience of hearing a sound in a large hall with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/largeHall
PHASEReverbPresetLargeHall PHASEReverbPreset = 0
// PHASEReverbPresetLargeHall2 - A resonation that simulates the experience of hearing a sound in one kind of large hall with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/largeHall2
PHASEReverbPresetLargeHall2 PHASEReverbPreset = 0
// PHASEReverbPresetLargeRoom - A resonation that simulates the experience of hearing a sound in a large room with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/largeRoom
PHASEReverbPresetLargeRoom PHASEReverbPreset = 0
// PHASEReverbPresetLargeRoom2 - A resonation that simulates the experience of hearing a sound in one kind of large room with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/largeRoom2
PHASEReverbPresetLargeRoom2 PHASEReverbPreset = 0
// PHASEReverbPresetMediumChamber - A resonation that simulates the experience of hearing a sound in a medium-size chamber with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/mediumChamber
PHASEReverbPresetMediumChamber PHASEReverbPreset = 0
// PHASEReverbPresetMediumHall - A resonation that simulates the experience of hearing a sound in a medium-size hall with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/mediumHall
PHASEReverbPresetMediumHall PHASEReverbPreset = 0
// PHASEReverbPresetMediumHall2 - A resonation that simulates the experience of hearing a sound in one kind of medium-size hall with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/mediumHall2
PHASEReverbPresetMediumHall2 PHASEReverbPreset = 0
// PHASEReverbPresetMediumHall3 - A resonation that simulates the experience of hearing a sound in another kind of medium-size hall with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/mediumHall3
PHASEReverbPresetMediumHall3 PHASEReverbPreset = 0
// PHASEReverbPresetMediumRoom - A resonation that simulates the experience of hearing a sound in a medium-size room with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/mediumRoom
PHASEReverbPresetMediumRoom PHASEReverbPreset = 0
// PHASEReverbPresetNone - An option that adds no reverberation to a sound.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/none
PHASEReverbPresetNone PHASEReverbPreset = 0
// PHASEReverbPresetSmallRoom - A resonation that simulates the experience of hearing a sound in a small room with specific dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEReverbPreset/smallRoom
PHASEReverbPresetSmallRoom PHASEReverbPreset = 0
)

// PHASESoundEventPrepareHandlerReason - Indicates the results of sound-event preparation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareHandlerReason
type PHASESoundEventPrepareHandlerReason uint

// PHASESoundEventPrepareState - Indicates the state of sound-event preparation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareState-swift.enum
type PHASESoundEventPrepareState uint

// PHASERenderingState - The playback status of audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/RenderingState-swift.enum
type PHASERenderingState uint

const (
// PHASERenderingStatePaused - A state in which sound event playback pauses.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/RenderingState-swift.enum/paused
PHASERenderingStatePaused PHASERenderingState = 0
// PHASERenderingStateStarted - A state in which sound event playback starts.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/RenderingState-swift.enum/started
PHASERenderingStateStarted PHASERenderingState = 0
// PHASERenderingStateStopped - A state in which sound event playback stops.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/RenderingState-swift.enum/stopped
PHASERenderingStateStopped PHASERenderingState = 0
)

// PHASESoundEventSeekHandlerReason - Indicates the status after a sound event changes its playback position.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/SeekHandlerReason
type PHASESoundEventSeekHandlerReason uint

// PHASESoundEventStartHandlerReason - Indicates the status after starting a sound event.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/StartHandlerReason
type PHASESoundEventStartHandlerReason uint

// PHASESoundEventError - Codes that identify sound event errors.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code
type PHASESoundEventError uint

// PHASESpatialPipelineFlags - Sound resonance options for a spatial pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/Flags-swift.struct
type PHASESpatialPipelineFlags uint

const (
// PHASESpatialPipelineFlagDirectPathTransmission - A spatial property that refers to the original audio signal.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/Flags-swift.struct/directPathTransmission
PHASESpatialPipelineFlagDirectPathTransmission PHASESpatialPipelineFlags = 0
// PHASESpatialPipelineFlagEarlyReflections - A spatial property that refers to the earlier echoes along the duration of sound resonance.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/Flags-swift.struct/earlyReflections
PHASESpatialPipelineFlagEarlyReflections PHASESpatialPipelineFlags = 0
// PHASESpatialPipelineFlagLateReverb - A spatial property that refers to the later echoes along the duration of sound resonance.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/Flags-swift.struct/lateReverb
PHASESpatialPipelineFlagLateReverb PHASESpatialPipelineFlags = 0
)

// PHASESpatializationMode - The manner in which PHASE outputs spatial audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatializationMode
type PHASESpatializationMode uint

const (
// PHASESpatializationModeAlwaysUseBinaural - A mode that introduces special processing to replicate a realistic spatial listening experience.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatializationMode/alwaysUseBinaural
PHASESpatializationModeAlwaysUseBinaural PHASESpatializationMode = 0
// PHASESpatializationModeAlwaysUseChannelBased - A mode that adds a 3D position and orientation to sound by panning across the available output channels.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatializationMode/alwaysUseChannelBased
PHASESpatializationModeAlwaysUseChannelBased PHASESpatializationMode = 0
// PHASESpatializationModeAutomatic - A mode that indicates that the framework chooses the spatialization mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatializationMode/automatic
PHASESpatializationModeAutomatic PHASESpatializationMode = 0
)


