// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

/* debug [enums.gen.go]: Generating 24 enums for PHASE */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum PHASEAssetType (2 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEAutomaticHeadTrackingFlags (2 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASECurveType (11 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASERenderingMode (2 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEUpdateMode (2 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEMaterialPreset (6 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEMediumPreset (1 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASENormalizationMode (2 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEReverbPreset (13 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASERenderingState (3 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASESoundEventPrepareHandlerReason (3 cases) */
// PHASESoundEventPrepareHandlerReason - Indicates the results of sound-event preparation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareHandlerReason
type PHASESoundEventPrepareHandlerReason uint

const (
	// PHASESoundEventPrepareHandlerReasonFailure - Indicates an error occurs during sound-event preparation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareHandlerReason/failure
	PHASESoundEventPrepareHandlerReasonFailure PHASESoundEventPrepareHandlerReason = 0
	// PHASESoundEventPrepareHandlerReasonPrepared - Indicates the completion of sound-event preparation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareHandlerReason/prepared
	PHASESoundEventPrepareHandlerReasonPrepared PHASESoundEventPrepareHandlerReason = 0
	// PHASESoundEventPrepareHandlerReasonTerminated - Indicates sound-event preparation stops abruptly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareHandlerReason/terminated
	PHASESoundEventPrepareHandlerReasonTerminated PHASESoundEventPrepareHandlerReason = 0
)

/* debug [enums.gen.go]: Processing enum PHASESoundEventPrepareState (3 cases) */
// PHASESoundEventPrepareState - Indicates the state of sound-event preparation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareState-swift.enum
type PHASESoundEventPrepareState uint

const (
	// PHASESoundEventPrepareStatePrepared - Indicates that the sound event preparation is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareState-swift.enum/prepared
	PHASESoundEventPrepareStatePrepared PHASESoundEventPrepareState = 0
	// PHASESoundEventPrepareStatePrepareInProgress - Indicates that the sound event prepares for playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareState-swift.enum/prepareInProgress
	PHASESoundEventPrepareStatePrepareInProgress PHASESoundEventPrepareState = 0
	// PHASESoundEventPrepareStatePrepareNotStarted - Indicates that the sound event awaits preparation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/PrepareState-swift.enum/prepareNotStarted
	PHASESoundEventPrepareStatePrepareNotStarted PHASESoundEventPrepareState = 0
)

/* debug [enums.gen.go]: Processing enum PHASESoundEventSeekHandlerReason (3 cases) */
// PHASESoundEventSeekHandlerReason - Indicates the status after a sound event changes its playback position.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/SeekHandlerReason
type PHASESoundEventSeekHandlerReason uint

const (
	// PHASESoundEventSeekHandlerReasonFailure - Indicates the sound event fails to update its playback position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/SeekHandlerReason/failure
	PHASESoundEventSeekHandlerReasonFailure PHASESoundEventSeekHandlerReason = 0
	// PHASESoundEventSeekHandlerReasonFailureSeekAlreadyInProgress - Indicates the sound event is still updating its playback position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/SeekHandlerReason/failureSeekAlreadyInProgress
	PHASESoundEventSeekHandlerReasonFailureSeekAlreadyInProgress PHASESoundEventSeekHandlerReason = 0
	// PHASESoundEventSeekHandlerReasonSeekSuccessful - Indicates the sound event successfully updated its playback position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/SeekHandlerReason/seekSuccessful
	PHASESoundEventSeekHandlerReasonSeekSuccessful PHASESoundEventSeekHandlerReason = 0
)

/* debug [enums.gen.go]: Processing enum PHASESoundEventStartHandlerReason (3 cases) */
// PHASESoundEventStartHandlerReason - Indicates the status after starting a sound event.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/StartHandlerReason
type PHASESoundEventStartHandlerReason uint

const (
	// PHASESoundEventStartHandlerReasonFailure - Indicates an error occurred while starting the sound event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/StartHandlerReason/failure
	PHASESoundEventStartHandlerReasonFailure PHASESoundEventStartHandlerReason = 0
	// PHASESoundEventStartHandlerReasonFinishedPlaying - Indicates the framework successfully started the sound event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/StartHandlerReason/finishedPlaying
	PHASESoundEventStartHandlerReasonFinishedPlaying PHASESoundEventStartHandlerReason = 0
	// PHASESoundEventStartHandlerReasonTerminated - Indicates the framework terminated the sound event abruptly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEvent/StartHandlerReason/terminated
	PHASESoundEventStartHandlerReasonTerminated PHASESoundEventStartHandlerReason = 0
)

/* debug [enums.gen.go]: Processing enum PHASESpatializationMode (3 cases) */
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

/* debug [enums.gen.go]: Processing enum PHASEAssetError (6 cases) */
// PHASEAssetError - Codes that identify framework asset errors.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code
type PHASEAssetError uint

const (
	// PHASEAssetErrorAlreadyExists - An error the asset registry throws when the app registers an asset twice by the same name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/alreadyExists
	PHASEAssetErrorAlreadyExists PHASEAssetError = 0
	// PHASEAssetErrorBadParameters - An error that indicates an asset registry call contains invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/badParameters
	PHASEAssetErrorBadParameters PHASEAssetError = 0
	// PHASEAssetErrorFailedToLoad - An error that indicates an asset failed to load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/failedToLoad
	PHASEAssetErrorFailedToLoad PHASEAssetError = 0
	// PHASEAssetErrorGeneralError - An error the asset registry throws when an unspecified problem occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/generalError
	PHASEAssetErrorGeneralError PHASEAssetError = 0
	// PHASEAssetErrorInvalidEngineInstance - An error that indicates an asset registry call references an invalid engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/invalidEngineInstance
	PHASEAssetErrorInvalidEngineInstance PHASEAssetError = 0
	// PHASEAssetErrorMemoryAllocation - An error the framework throws when an asset depletes system memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetError-swift.struct/Code/memoryAllocation
	PHASEAssetErrorMemoryAllocation PHASEAssetError = 0
)

/* debug [enums.gen.go]: Processing enum PHASECalibrationMode (3 cases) */
// PHASECalibrationMode - Calibration options for sound pressure level.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECalibrationMode
type PHASECalibrationMode uint

const (
	// PHASECalibrationModeAbsoluteSpl - A sound pressure level based on the current output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECalibrationMode/absoluteSpl
	PHASECalibrationModeAbsoluteSpl PHASECalibrationMode = 0
	// PHASECalibrationModeNone - An option that specifies no loudness calibration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECalibrationMode/none
	PHASECalibrationModeNone PHASECalibrationMode = 0
	// PHASECalibrationModeRelativeSpl - A sound pressure level that’s tuned for the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECalibrationMode/relativeSpl
	PHASECalibrationModeRelativeSpl PHASECalibrationMode = 0
)

/* debug [enums.gen.go]: Processing enum PHASECullOption (5 cases) */
// PHASECullOption - The actions the engine takes when it culls sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption
type PHASECullOption uint

const (
	// PHASECullOptionDoNotCull - An option that indicates the framework takes no action to cull sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption/doNotCull
	PHASECullOptionDoNotCull PHASECullOption = 0
	// PHASECullOptionSleepWakeAtRandomOffset - An option that pauses playback and resumes at a random position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption/sleepWakeAtRandomOffset
	PHASECullOptionSleepWakeAtRandomOffset PHASECullOption = 0
	// PHASECullOptionSleepWakeAtRealtimeOffset - An option that pauses playback and resumes where it left off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption/sleepWakeAtRealtimeOffset
	PHASECullOptionSleepWakeAtRealtimeOffset PHASECullOption = 0
	// PHASECullOptionSleepWakeAtZero - An option that pauses playback and resumes at the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption/sleepWakeAtZero
	PHASECullOptionSleepWakeAtZero PHASECullOption = 0
	// PHASECullOptionTerminate - An option that culls sound by stopping playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECullOption/terminate
	PHASECullOptionTerminate PHASECullOption = 0
)

/* debug [enums.gen.go]: Processing enum PHASEError (2 cases) */
// PHASEError - Codes that identify errors in PHASE.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEError-swift.struct/Code
type PHASEError uint

const (
	// PHASEErrorInitializeFailed - An error that indicates the engine failed to initialize.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEError-swift.struct/Code/initializeFailed
	PHASEErrorInitializeFailed PHASEError = 0
	// PHASEErrorInvalidObject - An error that indicates an object is invalid in a specific context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEError-swift.struct/Code/invalidObject
	PHASEErrorInvalidObject PHASEError = 0
)

/* debug [enums.gen.go]: Processing enum PHASEPlaybackMode (2 cases) */
// PHASEPlaybackMode - Loop options for audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPlaybackMode
type PHASEPlaybackMode uint

const (
	// PHASEPlaybackModeLooping - An option that restarts a sound from the begining after it finishes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPlaybackMode/looping
	PHASEPlaybackModeLooping PHASEPlaybackMode = 0
	// PHASEPlaybackModeOneShot - An option that plays a sound only once.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPlaybackMode/oneShot
	PHASEPlaybackModeOneShot PHASEPlaybackMode = 0
)

/* debug [enums.gen.go]: Processing enum PHASEPushStreamBufferOptions (4 cases) */
// PHASEPushStreamBufferOptions - Options that inform PHASE of an audio-stream buffer’s playback priority.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions
type PHASEPushStreamBufferOptions uint

const (
	// PHASEPushStreamBufferDefault - Indicates a buffer processes after existing buffers in the queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions/default
	PHASEPushStreamBufferDefault PHASEPushStreamBufferOptions = 0
	// PHASEPushStreamBufferInterrupts - Indicates a buffer begins processing immediately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions/interrupts
	PHASEPushStreamBufferInterrupts PHASEPushStreamBufferOptions = 0
	// PHASEPushStreamBufferInterruptsAtLoop - Indicates a buffer begins processing when an existing buffer loops.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions/interruptsAtLoop
	PHASEPushStreamBufferInterruptsAtLoop PHASEPushStreamBufferOptions = 0
	// PHASEPushStreamBufferLoops - Indicates a buffer restarts after it finishes processing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamBufferOptions/loops
	PHASEPushStreamBufferLoops PHASEPushStreamBufferOptions = 0
)

/* debug [enums.gen.go]: Processing enum PHASEPushStreamCompletionCallbackCondition (1 cases) */
// PHASEPushStreamCompletionCallbackCondition - A status that describes the results after the app schedules a push-stream buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamCompletionCallbackCondition
type PHASEPushStreamCompletionCallbackCondition uint

const (
	// PHASEPushStreamCompletionDataRendered - Indicates the framework invokes the callback when the engine processes the audio for output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamCompletionCallbackCondition/dataRendered
	PHASEPushStreamCompletionDataRendered PHASEPushStreamCompletionCallbackCondition = 0
)

/* debug [enums.gen.go]: Processing enum PHASESoundEventError (6 cases) */
// PHASESoundEventError - Codes that identify sound event errors.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code
type PHASESoundEventError uint

const (
	// PHASESoundEventErrorAPIMisuse - An error that indicates the app misconfigures data or calls the framework in unsupported succession.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/apiMisuse
	PHASESoundEventErrorAPIMisuse PHASESoundEventError = 0
	// PHASESoundEventErrorBadData - An error that indicates a sound event contains invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/badData
	PHASESoundEventErrorBadData PHASESoundEventError = 0
	// PHASESoundEventErrorInvalidInstance - An error that indicates a sound event object is no longer valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/invalidInstance
	PHASESoundEventErrorInvalidInstance PHASESoundEventError = 0
	// PHASESoundEventErrorNotFound - An error the framework throws when it fails to find a particular sound event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/notFound
	PHASESoundEventErrorNotFound PHASESoundEventError = 0
	// PHASESoundEventErrorOutOfMemory - An error the framework throws when a sound event depletes system memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/outOfMemory
	PHASESoundEventErrorOutOfMemory PHASESoundEventError = 0
	// PHASESoundEventErrorSystemNotInitialized - An error the framework throws when engine initialization interrupts sound event playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventError-swift.struct/Code/systemNotInitialized
	PHASESoundEventErrorSystemNotInitialized PHASESoundEventError = 0
)

/* debug [enums.gen.go]: Processing enum PHASESpatialPipelineFlags (3 cases) */
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


