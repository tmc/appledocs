// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// Enum types and constants
// AVAssetWriterStatus - Values that indicate the state of an asset writer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum
type AssetWriterStatus uint

const (
	// AssetWriterStatusCompleted - The asset writer finishes writing successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/completed
	AssetWriterStatusCompleted AssetWriterStatus = 2
	// AssetWriterStatusFailed - The asset writer fails to write the output file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/failed
	AssetWriterStatusFailed AssetWriterStatus = 3
)

// AVAuthorizationStatus - Constants that indicate the status of an app’s authorization to capture media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AuthorizationStatus uint

const (
	// AuthorizationStatusAuthorized - A status that indicates the user has explicitly granted an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/authorized
	AuthorizationStatusAuthorized AuthorizationStatus = 3
	// AuthorizationStatusDenied - A status that indicates the user has explicitly denied an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/denied
	AuthorizationStatusDenied AuthorizationStatus = 2
	// AuthorizationStatusNotDetermined - A status that indicates the user hasn’t yet granted or denied authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/notDetermined
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	// AuthorizationStatusRestricted - A status that indicates the app isn’t permitted to use media capture devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/restricted
	AuthorizationStatusRestricted AuthorizationStatus = 1
)

// AVCaptureColorSpace - An enumeration of color spaces a device can support.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
type CaptureColorSpace uint

const (
	CaptureColorSpace_sRGB CaptureColorSpace = 0
	CaptureColorSpace_P3_D65 CaptureColorSpace = 1
	CaptureColorSpace_HLG_BT2020 CaptureColorSpace = 2
	CaptureColorSpace_AppleLog CaptureColorSpace = 3
	CaptureCenterStageControlModeUser CaptureColorSpace = 0
	CaptureCenterStageControlModeApp CaptureColorSpace = 1
	CaptureCenterStageControlModeCooperative CaptureColorSpace = 2
	CaptureMicrophoneModeStandard CaptureColorSpace = 0
	CaptureMicrophoneModeWideSpectrum CaptureColorSpace = 1
	CaptureMicrophoneModeVoiceIsolation CaptureColorSpace = 2
	CaptureSystemUserInterfaceVideoEffects CaptureColorSpace = 1
	CaptureSystemUserInterfaceMicrophoneModes CaptureColorSpace = 2
	CaptureVideoStabilizationModeOff CaptureColorSpace = 0
	CaptureVideoStabilizationModeStandard CaptureColorSpace = 1
	CaptureVideoStabilizationModeCinematic CaptureColorSpace = 2
	CaptureVideoStabilizationModeCinematicExtended CaptureColorSpace = 3
	CaptureVideoStabilizationModePreviewOptimized CaptureColorSpace = 4
	CaptureVideoStabilizationModeCinematicExtendedEnhanced CaptureColorSpace = 5
	CaptureAutoFocusSystemNone CaptureColorSpace = 0
	CaptureAutoFocusSystemContrastDetection CaptureColorSpace = 1
	CaptureAutoFocusSystemPhaseDetection CaptureColorSpace = 2
	CaptureCameraLensSmudgeDetectionStatusDisabled CaptureColorSpace = 0
	CaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected CaptureColorSpace = 1
	CaptureCameraLensSmudgeDetectionStatusSmudged CaptureColorSpace = 2
	CaptureCameraLensSmudgeDetectionStatusUnknown CaptureColorSpace = 3
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureColorSpace = 1
	CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient CaptureColorSpace = 2
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureColorSpace = 3
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureColorSpace = 4
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure CaptureColorSpace = 5
	CaptureSessionInterruptionReasonSensitiveContentMitigationActivated CaptureColorSpace = 6
	CaptureVideoOrientationPortrait CaptureColorSpace = 1
	CaptureVideoOrientationPortraitUpsideDown CaptureColorSpace = 2
	CaptureVideoOrientationLandscapeRight CaptureColorSpace = 3
	CaptureVideoOrientationLandscapeLeft CaptureColorSpace = 4
	VideoFieldModeBoth CaptureColorSpace = 0
	VideoFieldModeTopOnly CaptureColorSpace = 1
	VideoFieldModeBottomOnly CaptureColorSpace = 2
	VideoFieldModeDeinterlace CaptureColorSpace = 3
	CaptureOutputDataDroppedReasonNone CaptureColorSpace = 0
	CaptureOutputDataDroppedReasonLateData CaptureColorSpace = 1
	CaptureOutputDataDroppedReasonOutOfBuffers CaptureColorSpace = 2
	CaptureOutputDataDroppedReasonDiscontinuity CaptureColorSpace = 3
	CapturePhotoQualityPrioritizationSpeed CaptureColorSpace = 1
	CapturePhotoQualityPrioritizationBalanced CaptureColorSpace = 2
	CapturePhotoQualityPrioritizationQuality CaptureColorSpace = 3
	CapturePhotoOutputCaptureReadinessSessionNotRunning CaptureColorSpace = 0
	CapturePhotoOutputCaptureReadinessReady CaptureColorSpace = 1
	CapturePhotoOutputCaptureReadinessNotReadyMomentarily CaptureColorSpace = 2
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture CaptureColorSpace = 3
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing CaptureColorSpace = 4
	CaptureLensStabilizationStatusUnsupported CaptureColorSpace = 0
	CaptureLensStabilizationStatusOff CaptureColorSpace = 1
	CaptureLensStabilizationStatusActive CaptureColorSpace = 2
	CaptureLensStabilizationStatusOutOfRange CaptureColorSpace = 3
	CaptureLensStabilizationStatusUnavailable CaptureColorSpace = 4
	CaptureMultichannelAudioModeNone CaptureColorSpace = 0
	CaptureMultichannelAudioModeStereo CaptureColorSpace = 1
	CaptureMultichannelAudioModeFirstOrderAmbisonics CaptureColorSpace = 2
	CaptureSystemPressureFactorNone CaptureColorSpace = 0
	CaptureSystemPressureFactorCameraTemperature CaptureColorSpace = 1
	DepthDataQualityLow CaptureColorSpace = 0
	DepthDataQualityHigh CaptureColorSpace = 1
	DepthDataAccuracyRelative CaptureColorSpace = 0
	DepthDataAccuracyAbsolute CaptureColorSpace = 1
)

// AVCaptureCinematicVideoFocusMode - Constants indicating the focus behavior when recording a Cinematic Video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
type CaptureCinematicVideoFocusMode uint

const (
	CaptureCinematicVideoFocusModeNone CaptureCinematicVideoFocusMode = 0
	CaptureCinematicVideoFocusModeStrong CaptureCinematicVideoFocusMode = 1
	CaptureCinematicVideoFocusModeWeak CaptureCinematicVideoFocusMode = 2
	CaptureExposureModeLocked CaptureCinematicVideoFocusMode = 0
	CaptureExposureModeAutoExpose CaptureCinematicVideoFocusMode = 1
	CaptureExposureModeContinuousAutoExposure CaptureCinematicVideoFocusMode = 2
	CaptureExposureModeCustom CaptureCinematicVideoFocusMode = 3
	CaptureWhiteBalanceModeLocked CaptureCinematicVideoFocusMode = 0
	CaptureWhiteBalanceModeAutoWhiteBalance CaptureCinematicVideoFocusMode = 1
	CaptureWhiteBalanceModeContinuousAutoWhiteBalance CaptureCinematicVideoFocusMode = 2
	AuthorizationStatusNotDetermined CaptureCinematicVideoFocusMode = 0
	AuthorizationStatusRestricted CaptureCinematicVideoFocusMode = 1
	AuthorizationStatusDenied CaptureCinematicVideoFocusMode = 2
	AuthorizationStatusAuthorized CaptureCinematicVideoFocusMode = 3
	CaptureDeviceTransportControlsNotPlayingMode CaptureCinematicVideoFocusMode = 0
	CaptureDeviceTransportControlsPlayingMode CaptureCinematicVideoFocusMode = 1
	CaptureColorSpace_sRGB CaptureCinematicVideoFocusMode = 0
	CaptureColorSpace_P3_D65 CaptureCinematicVideoFocusMode = 1
	CaptureColorSpace_HLG_BT2020 CaptureCinematicVideoFocusMode = 2
	CaptureColorSpace_AppleLog CaptureCinematicVideoFocusMode = 3
	CaptureCenterStageControlModeUser CaptureCinematicVideoFocusMode = 0
	CaptureCenterStageControlModeApp CaptureCinematicVideoFocusMode = 1
	CaptureCenterStageControlModeCooperative CaptureCinematicVideoFocusMode = 2
	CaptureMicrophoneModeStandard CaptureCinematicVideoFocusMode = 0
	CaptureMicrophoneModeWideSpectrum CaptureCinematicVideoFocusMode = 1
	CaptureMicrophoneModeVoiceIsolation CaptureCinematicVideoFocusMode = 2
	CaptureSystemUserInterfaceVideoEffects CaptureCinematicVideoFocusMode = 1
	CaptureSystemUserInterfaceMicrophoneModes CaptureCinematicVideoFocusMode = 2
	CaptureVideoStabilizationModeOff CaptureCinematicVideoFocusMode = 0
	CaptureVideoStabilizationModeStandard CaptureCinematicVideoFocusMode = 1
	CaptureVideoStabilizationModeCinematic CaptureCinematicVideoFocusMode = 2
	CaptureVideoStabilizationModeCinematicExtended CaptureCinematicVideoFocusMode = 3
	CaptureVideoStabilizationModePreviewOptimized CaptureCinematicVideoFocusMode = 4
	CaptureVideoStabilizationModeCinematicExtendedEnhanced CaptureCinematicVideoFocusMode = 5
	CaptureAutoFocusSystemNone CaptureCinematicVideoFocusMode = 0
	CaptureAutoFocusSystemContrastDetection CaptureCinematicVideoFocusMode = 1
	CaptureAutoFocusSystemPhaseDetection CaptureCinematicVideoFocusMode = 2
	CaptureCameraLensSmudgeDetectionStatusDisabled CaptureCinematicVideoFocusMode = 0
	CaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected CaptureCinematicVideoFocusMode = 1
	CaptureCameraLensSmudgeDetectionStatusSmudged CaptureCinematicVideoFocusMode = 2
	CaptureCameraLensSmudgeDetectionStatusUnknown CaptureCinematicVideoFocusMode = 3
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureCinematicVideoFocusMode = 1
	CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient CaptureCinematicVideoFocusMode = 2
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureCinematicVideoFocusMode = 3
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureCinematicVideoFocusMode = 4
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure CaptureCinematicVideoFocusMode = 5
	CaptureSessionInterruptionReasonSensitiveContentMitigationActivated CaptureCinematicVideoFocusMode = 6
	CaptureVideoOrientationPortrait CaptureCinematicVideoFocusMode = 1
	CaptureVideoOrientationPortraitUpsideDown CaptureCinematicVideoFocusMode = 2
	CaptureVideoOrientationLandscapeRight CaptureCinematicVideoFocusMode = 3
	CaptureVideoOrientationLandscapeLeft CaptureCinematicVideoFocusMode = 4
	VideoFieldModeBoth CaptureCinematicVideoFocusMode = 0
	VideoFieldModeTopOnly CaptureCinematicVideoFocusMode = 1
	VideoFieldModeBottomOnly CaptureCinematicVideoFocusMode = 2
	VideoFieldModeDeinterlace CaptureCinematicVideoFocusMode = 3
	CaptureOutputDataDroppedReasonNone CaptureCinematicVideoFocusMode = 0
	CaptureOutputDataDroppedReasonLateData CaptureCinematicVideoFocusMode = 1
	CaptureOutputDataDroppedReasonOutOfBuffers CaptureCinematicVideoFocusMode = 2
	CaptureOutputDataDroppedReasonDiscontinuity CaptureCinematicVideoFocusMode = 3
	CapturePhotoQualityPrioritizationSpeed CaptureCinematicVideoFocusMode = 1
	CapturePhotoQualityPrioritizationBalanced CaptureCinematicVideoFocusMode = 2
	CapturePhotoQualityPrioritizationQuality CaptureCinematicVideoFocusMode = 3
	CapturePhotoOutputCaptureReadinessSessionNotRunning CaptureCinematicVideoFocusMode = 0
	CapturePhotoOutputCaptureReadinessReady CaptureCinematicVideoFocusMode = 1
	CapturePhotoOutputCaptureReadinessNotReadyMomentarily CaptureCinematicVideoFocusMode = 2
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture CaptureCinematicVideoFocusMode = 3
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing CaptureCinematicVideoFocusMode = 4
	CaptureLensStabilizationStatusUnsupported CaptureCinematicVideoFocusMode = 0
	CaptureLensStabilizationStatusOff CaptureCinematicVideoFocusMode = 1
	CaptureLensStabilizationStatusActive CaptureCinematicVideoFocusMode = 2
	CaptureLensStabilizationStatusOutOfRange CaptureCinematicVideoFocusMode = 3
	CaptureLensStabilizationStatusUnavailable CaptureCinematicVideoFocusMode = 4
	CaptureMultichannelAudioModeNone CaptureCinematicVideoFocusMode = 0
	CaptureMultichannelAudioModeStereo CaptureCinematicVideoFocusMode = 1
	CaptureMultichannelAudioModeFirstOrderAmbisonics CaptureCinematicVideoFocusMode = 2
	CaptureSystemPressureFactorNone CaptureCinematicVideoFocusMode = 0
	CaptureSystemPressureFactorCameraTemperature CaptureCinematicVideoFocusMode = 1
	DepthDataQualityLow CaptureCinematicVideoFocusMode = 0
	DepthDataQualityHigh CaptureCinematicVideoFocusMode = 1
	DepthDataAccuracyRelative CaptureCinematicVideoFocusMode = 0
	DepthDataAccuracyAbsolute CaptureCinematicVideoFocusMode = 1
)

// AVCaptureLensStabilizationStatus - Constants that indicate the status of optical image stabilization hardware during a bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type CaptureLensStabilizationStatus uint

const (
	CaptureLensStabilizationStatusUnsupported CaptureLensStabilizationStatus = 0
	CaptureLensStabilizationStatusOff CaptureLensStabilizationStatus = 1
	CaptureLensStabilizationStatusActive CaptureLensStabilizationStatus = 2
	CaptureLensStabilizationStatusOutOfRange CaptureLensStabilizationStatus = 3
	CaptureLensStabilizationStatusUnavailable CaptureLensStabilizationStatus = 4
	CaptureMultichannelAudioModeNone CaptureLensStabilizationStatus = 0
	CaptureMultichannelAudioModeStereo CaptureLensStabilizationStatus = 1
	CaptureMultichannelAudioModeFirstOrderAmbisonics CaptureLensStabilizationStatus = 2
	CaptureSystemPressureFactorNone CaptureLensStabilizationStatus = 0
	CaptureSystemPressureFactorCameraTemperature CaptureLensStabilizationStatus = 1
	DepthDataQualityLow CaptureLensStabilizationStatus = 0
	DepthDataQualityHigh CaptureLensStabilizationStatus = 1
	DepthDataAccuracyRelative CaptureLensStabilizationStatus = 0
	DepthDataAccuracyAbsolute CaptureLensStabilizationStatus = 1
)

// AVCaptureDevicePosition - Constants that indicate the physical position of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type CaptureDevicePosition uint

const (
	// CaptureDevicePositionBack - A position on the subject-facing side of an iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/back
	CaptureDevicePositionBack CaptureDevicePosition = 1
	// CaptureDevicePositionUnspecified - A position that’s unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/unspecified
	CaptureDevicePositionUnspecified CaptureDevicePosition = 0
)

// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions - A structure that defines the conditions in which to restrict camera switching.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct
type CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions uint

const (
	// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone - Disallow switching to a fallback camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
)

// AVCapturePrimaryConstituentDeviceSwitchingBehavior - Constants that control when to allow a virtual device to switch its active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum
type CapturePrimaryConstituentDeviceSwitchingBehavior uint

const (
	// CapturePrimaryConstituentDeviceSwitchingBehaviorAuto - The device automatically selects the best camera for the current scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/auto
	CapturePrimaryConstituentDeviceSwitchingBehaviorAuto CapturePrimaryConstituentDeviceSwitchingBehavior = 1
	// CapturePrimaryConstituentDeviceSwitchingBehaviorLocked - The device locks camera switching to the active primary constituent device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/locked
	CapturePrimaryConstituentDeviceSwitchingBehaviorLocked CapturePrimaryConstituentDeviceSwitchingBehavior = 3
	// CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported - The device doesn’t support constituent device switching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/unsupported
	CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported CapturePrimaryConstituentDeviceSwitchingBehavior = 0
)

// AVCaptureOutputDataDroppedReason - Constants that define reasons for why the system dropped a frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type CaptureOutputDataDroppedReason uint

const (
	CaptureOutputDataDroppedReasonNone CaptureOutputDataDroppedReason = 0
	CaptureOutputDataDroppedReasonLateData CaptureOutputDataDroppedReason = 1
	CaptureOutputDataDroppedReasonOutOfBuffers CaptureOutputDataDroppedReason = 2
	CaptureOutputDataDroppedReasonDiscontinuity CaptureOutputDataDroppedReason = 3
	CapturePhotoQualityPrioritizationSpeed CaptureOutputDataDroppedReason = 1
	CapturePhotoQualityPrioritizationBalanced CaptureOutputDataDroppedReason = 2
	CapturePhotoQualityPrioritizationQuality CaptureOutputDataDroppedReason = 3
	CapturePhotoOutputCaptureReadinessSessionNotRunning CaptureOutputDataDroppedReason = 0
	CapturePhotoOutputCaptureReadinessReady CaptureOutputDataDroppedReason = 1
	CapturePhotoOutputCaptureReadinessNotReadyMomentarily CaptureOutputDataDroppedReason = 2
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture CaptureOutputDataDroppedReason = 3
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing CaptureOutputDataDroppedReason = 4
	CaptureLensStabilizationStatusUnsupported CaptureOutputDataDroppedReason = 0
	CaptureLensStabilizationStatusOff CaptureOutputDataDroppedReason = 1
	CaptureLensStabilizationStatusActive CaptureOutputDataDroppedReason = 2
	CaptureLensStabilizationStatusOutOfRange CaptureOutputDataDroppedReason = 3
	CaptureLensStabilizationStatusUnavailable CaptureOutputDataDroppedReason = 4
	CaptureMultichannelAudioModeNone CaptureOutputDataDroppedReason = 0
	CaptureMultichannelAudioModeStereo CaptureOutputDataDroppedReason = 1
	CaptureMultichannelAudioModeFirstOrderAmbisonics CaptureOutputDataDroppedReason = 2
	CaptureSystemPressureFactorNone CaptureOutputDataDroppedReason = 0
	CaptureSystemPressureFactorCameraTemperature CaptureOutputDataDroppedReason = 1
	DepthDataQualityLow CaptureOutputDataDroppedReason = 0
	DepthDataQualityHigh CaptureOutputDataDroppedReason = 1
	DepthDataAccuracyRelative CaptureOutputDataDroppedReason = 0
	DepthDataAccuracyAbsolute CaptureOutputDataDroppedReason = 1
)

// AVCapturePhotoOutputCaptureReadiness - Constants that indicate whether the output is ready to receive capture requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum
type CapturePhotoOutputCaptureReadiness uint

const (
	CapturePhotoOutputCaptureReadinessSessionNotRunning CapturePhotoOutputCaptureReadiness = 0
	CapturePhotoOutputCaptureReadinessReady CapturePhotoOutputCaptureReadiness = 1
	CapturePhotoOutputCaptureReadinessNotReadyMomentarily CapturePhotoOutputCaptureReadiness = 2
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture CapturePhotoOutputCaptureReadiness = 3
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing CapturePhotoOutputCaptureReadiness = 4
	CaptureLensStabilizationStatusUnsupported CapturePhotoOutputCaptureReadiness = 0
	CaptureLensStabilizationStatusOff CapturePhotoOutputCaptureReadiness = 1
	CaptureLensStabilizationStatusActive CapturePhotoOutputCaptureReadiness = 2
	CaptureLensStabilizationStatusOutOfRange CapturePhotoOutputCaptureReadiness = 3
	CaptureLensStabilizationStatusUnavailable CapturePhotoOutputCaptureReadiness = 4
	CaptureMultichannelAudioModeNone CapturePhotoOutputCaptureReadiness = 0
	CaptureMultichannelAudioModeStereo CapturePhotoOutputCaptureReadiness = 1
	CaptureMultichannelAudioModeFirstOrderAmbisonics CapturePhotoOutputCaptureReadiness = 2
	CaptureSystemPressureFactorNone CapturePhotoOutputCaptureReadiness = 0
	CaptureSystemPressureFactorCameraTemperature CapturePhotoOutputCaptureReadiness = 1
	DepthDataQualityLow CapturePhotoOutputCaptureReadiness = 0
	DepthDataQualityHigh CapturePhotoOutputCaptureReadiness = 1
	DepthDataAccuracyRelative CapturePhotoOutputCaptureReadiness = 0
	DepthDataAccuracyAbsolute CapturePhotoOutputCaptureReadiness = 1
)

// AVCaptureSessionInterruptionReason - Constants identifying the reason a capture session was interrupted, found in an 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type CaptureSessionInterruptionReason uint

const (
	// CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient - An interruption caused by the audio hardware temporarily being made unavailable (for example, for a phone call or alarm).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/audioDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 2
	// CaptureSessionInterruptionReasonSensitiveContentMitigationActivated - An interruption caused by a   when it detects sensitive content on an associated  .  To resume your capture session, call your analyzer’s   method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/sensitiveContentMitigationActivated
	CaptureSessionInterruptionReasonSensitiveContentMitigationActivated CaptureSessionInterruptionReason = 6
	// CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient - An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 3
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure - An interruption due to system pressure, such as thermal duress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableDueToSystemPressure
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure CaptureSessionInterruptionReason = 5
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground - An interruption caused by the app being sent to the background while using a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableInBackground
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureSessionInterruptionReason = 1
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps - An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableWithMultipleForegroundApps
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureSessionInterruptionReason = 4
)

// AVError - An enumeration that defines the errors that framework operations can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type Error uint

const (
	// ErrorApplicationIsNotAuthorizedToUseDevice - The user denied this app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/applicationIsNotAuthorizedToUseDevice
	ErrorApplicationIsNotAuthorizedToUseDevice Error = 0
	// ErrorContentIsUnavailable - The captured content is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentIsUnavailable
	ErrorContentIsUnavailable Error = 0
	// ErrorDeviceAlreadyUsedByAnotherSession - Your app can’t access the device because another session is currently using it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceAlreadyUsedByAnotherSession
	ErrorDeviceAlreadyUsedByAnotherSession Error = 0
	// ErrorIncorrectlyConfigured - The system is incorrectly configured for the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incorrectlyConfigured
	ErrorIncorrectlyConfigured Error = 0
	// ErrorOutOfMemory - The operation couldn’t finish because there isn’t enough memory available to process the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/outOfMemory
	ErrorOutOfMemory Error = 0
	// ErrorScreenCaptureFailed - An unexpected problem occurred that prevented screen capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/screenCaptureFailed
	ErrorScreenCaptureFailed Error = 0
)

// AVPlayerActionAtItemEnd - The actions a player can take when it finishes playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
type PlayerActionAtItemEnd uint

// AVPlayerHDRMode - A bitfield type that specifies an HDR mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode
type PlayerHDRMode uint

const (
	PlayerHDRModeHLG PlayerHDRMode = 0
	PlayerHDRModeHDR10 PlayerHDRMode = 0
	PlayerHDRModeDolbyVision PlayerHDRMode = 0
	PlayerAudiovisualBackgroundPlaybackPolicyAutomatic PlayerHDRMode = 1
	PlayerAudiovisualBackgroundPlaybackPolicyPauses PlayerHDRMode = 2
	PlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible PlayerHDRMode = 3
	PlayerNetworkResourcePriorityDefault PlayerHDRMode = 0
	PlayerNetworkResourcePriorityLow PlayerHDRMode = 1
	PlayerNetworkResourcePriorityHigh PlayerHDRMode = 2
)

// AVPlayerNetworkResourcePriority - This defines the network resource priority for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type PlayerNetworkResourcePriority uint

const (
	PlayerNetworkResourcePriorityDefault PlayerNetworkResourcePriority = 0
	PlayerNetworkResourcePriorityLow PlayerNetworkResourcePriority = 1
	PlayerNetworkResourcePriorityHigh PlayerNetworkResourcePriority = 2
)

// AVPlayerStatus - Status values that indicate whether a player can successfully play media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type PlayerStatus uint

const (
	PlayerStatusUnknown PlayerStatus = 0
	PlayerStatusReadyToPlay PlayerStatus = 1
	PlayerStatusFailed PlayerStatus = 2
)

// AVPlayerTimeControlStatus - Constants that indicate the state of playback control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum
type PlayerTimeControlStatus uint

const (
	// PlayerTimeControlStatusPaused - A state that indicates the player paused playback indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/paused
	PlayerTimeControlStatusPaused PlayerTimeControlStatus = 0
)

// AVPlayerAudiovisualBackgroundPlaybackPolicy - Policies that describe playback behavior when an app transitions to the background while playing video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type PlayerAudiovisualBackgroundPlaybackPolicy uint

const (
	PlayerAudiovisualBackgroundPlaybackPolicyAutomatic PlayerAudiovisualBackgroundPlaybackPolicy = 1
	PlayerAudiovisualBackgroundPlaybackPolicyPauses PlayerAudiovisualBackgroundPlaybackPolicy = 2
	PlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible PlayerAudiovisualBackgroundPlaybackPolicy = 3
	PlayerNetworkResourcePriorityDefault PlayerAudiovisualBackgroundPlaybackPolicy = 0
	PlayerNetworkResourcePriorityLow PlayerAudiovisualBackgroundPlaybackPolicy = 1
	PlayerNetworkResourcePriorityHigh PlayerAudiovisualBackgroundPlaybackPolicy = 2
)

// AVPlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type PlayerItemStatus uint

const (
	// PlayerItemStatusFailed - The item no longer plays due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/failed
	PlayerItemStatusFailed PlayerItemStatus = 2
	// PlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	PlayerItemStatusReadyToPlay PlayerItemStatus = 1
	// PlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	PlayerItemStatusUnknown PlayerItemStatus = 0
)


