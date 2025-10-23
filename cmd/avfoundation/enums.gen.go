// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// Enum types and constants
// AVAssetWriterStatus - Values that indicate the state of an asset writer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum
type AVAssetWriterStatus uint

const (
	// AVAssetWriterStatusCompleted - The asset writer finishes writing successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/completed
	AVAssetWriterStatusCompleted AVAssetWriterStatus = 0
	// AVAssetWriterStatusFailed - The asset writer fails to write the output file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/failed
	AVAssetWriterStatusFailed AVAssetWriterStatus = 0
)

// AVAuthorizationStatus - Constants that indicate the status of an app’s authorization to capture media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AVAuthorizationStatus uint

const (
	// AVAuthorizationStatusAuthorized - A status that indicates the user has explicitly granted an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/authorized
	AVAuthorizationStatusAuthorized AVAuthorizationStatus = 0
	// AVAuthorizationStatusDenied - A status that indicates the user has explicitly denied an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/denied
	AVAuthorizationStatusDenied AVAuthorizationStatus = 0
	// AVAuthorizationStatusNotDetermined - A status that indicates the user hasn’t yet granted or denied authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/notDetermined
	AVAuthorizationStatusNotDetermined AVAuthorizationStatus = 0
	// AVAuthorizationStatusRestricted - A status that indicates the app isn’t permitted to use media capture devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/restricted
	AVAuthorizationStatusRestricted AVAuthorizationStatus = 0
)

// AVCaptureColorSpace - An enumeration of color spaces a device can support.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
type AVCaptureColorSpace uint

// AVCaptureCinematicVideoFocusMode - Constants indicating the focus behavior when recording a Cinematic Video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
type AVCaptureCinematicVideoFocusMode uint

// AVCaptureLensStabilizationStatus - Constants that indicate the status of optical image stabilization hardware during a bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type AVCaptureLensStabilizationStatus uint

// AVCaptureDevicePosition - Constants that indicate the physical position of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type AVCaptureDevicePosition uint

const (
	// AVCaptureDevicePositionBack - A position on the subject-facing side of an iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/back
	AVCaptureDevicePositionBack AVCaptureDevicePosition = 0
	// AVCaptureDevicePositionUnspecified - A position that’s unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/unspecified
	AVCaptureDevicePositionUnspecified AVCaptureDevicePosition = 0
)

// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions - A structure that defines the conditions in which to restrict camera switching.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct
type AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions uint

const (
	// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone - Disallow switching to a fallback camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone
	AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
)

// AVCapturePrimaryConstituentDeviceSwitchingBehavior - Constants that control when to allow a virtual device to switch its active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum
type AVCapturePrimaryConstituentDeviceSwitchingBehavior uint

const (
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto - The device automatically selects the best camera for the current scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/auto
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto AVCapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked - The device locks camera switching to the active primary constituent device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/locked
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked AVCapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported - The device doesn’t support constituent device switching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/unsupported
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported AVCapturePrimaryConstituentDeviceSwitchingBehavior = 0
)

// AVCaptureOutputDataDroppedReason - Constants that define reasons for why the system dropped a frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type AVCaptureOutputDataDroppedReason uint

// AVCapturePhotoOutputCaptureReadiness - Constants that indicate whether the output is ready to receive capture requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum
type AVCapturePhotoOutputCaptureReadiness uint

// AVCaptureSessionInterruptionReason - Constants identifying the reason a capture session was interrupted, found in an 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type AVCaptureSessionInterruptionReason uint

const (
	// AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient - An interruption caused by the audio hardware temporarily being made unavailable (for example, for a phone call or alarm).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/audioDeviceInUseByAnotherClient
	AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient AVCaptureSessionInterruptionReason = 0
	// AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated - An interruption caused by a   when it detects sensitive content on an associated  .  To resume your capture session, call your analyzer’s   method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/sensitiveContentMitigationActivated
	AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated AVCaptureSessionInterruptionReason = 0
	// AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient - An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceInUseByAnotherClient
	AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient AVCaptureSessionInterruptionReason = 0
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure - An interruption due to system pressure, such as thermal duress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableDueToSystemPressure
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure AVCaptureSessionInterruptionReason = 0
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground - An interruption caused by the app being sent to the background while using a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableInBackground
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground AVCaptureSessionInterruptionReason = 0
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps - An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableWithMultipleForegroundApps
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps AVCaptureSessionInterruptionReason = 0
)

// AVError - An enumeration that defines the errors that framework operations can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type AVError uint

const (
	// AVErrorApplicationIsNotAuthorizedToUseDevice - The user denied this app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/applicationIsNotAuthorizedToUseDevice
	AVErrorApplicationIsNotAuthorizedToUseDevice AVError = 0
	// AVErrorContentIsUnavailable - The captured content is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentIsUnavailable
	AVErrorContentIsUnavailable AVError = 0
	// AVErrorDeviceAlreadyUsedByAnotherSession - Your app can’t access the device because another session is currently using it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceAlreadyUsedByAnotherSession
	AVErrorDeviceAlreadyUsedByAnotherSession AVError = 0
	// AVErrorIncorrectlyConfigured - The system is incorrectly configured for the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incorrectlyConfigured
	AVErrorIncorrectlyConfigured AVError = 0
	// AVErrorOutOfMemory - The operation couldn’t finish because there isn’t enough memory available to process the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/outOfMemory
	AVErrorOutOfMemory AVError = 0
	// AVErrorScreenCaptureFailed - An unexpected problem occurred that prevented screen capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/screenCaptureFailed
	AVErrorScreenCaptureFailed AVError = 0
)

// AVPlayerActionAtItemEnd - The actions a player can take when it finishes playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
type AVPlayerActionAtItemEnd uint

// AVPlayerHDRMode - A bitfield type that specifies an HDR mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode
type AVPlayerHDRMode uint

// AVPlayerNetworkResourcePriority - This defines the network resource priority for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type AVPlayerNetworkResourcePriority uint

// AVPlayerStatus - Status values that indicate whether a player can successfully play media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type AVPlayerStatus uint

// AVPlayerTimeControlStatus - Constants that indicate the state of playback control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum
type AVPlayerTimeControlStatus uint

const (
	// AVPlayerTimeControlStatusPaused - A state that indicates the player paused playback indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/paused
	AVPlayerTimeControlStatusPaused AVPlayerTimeControlStatus = 0
)

// AVPlayerAudiovisualBackgroundPlaybackPolicy - Policies that describe playback behavior when an app transitions to the background while playing video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type AVPlayerAudiovisualBackgroundPlaybackPolicy uint

// AVPlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type AVPlayerItemStatus uint

const (
	// AVPlayerItemStatusFailed - The item no longer plays due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/failed
	AVPlayerItemStatusFailed AVPlayerItemStatus = 0
	// AVPlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	AVPlayerItemStatusReadyToPlay AVPlayerItemStatus = 0
	// AVPlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	AVPlayerItemStatusUnknown AVPlayerItemStatus = 0
)


