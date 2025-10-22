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
	AssetWriterStatusCompleted AssetWriterStatus = 0
	// AssetWriterStatusFailed - The asset writer fails to write the output file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/failed
	AssetWriterStatusFailed AssetWriterStatus = 0
)

// AVAuthorizationStatus - Constants that indicate the status of an app’s authorization to capture media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AuthorizationStatus uint

const (
	// AuthorizationStatusAuthorized - A status that indicates the user has explicitly granted an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/authorized
	AuthorizationStatusAuthorized AuthorizationStatus = 0
	// AuthorizationStatusDenied - A status that indicates the user has explicitly denied an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/denied
	AuthorizationStatusDenied AuthorizationStatus = 0
	// AuthorizationStatusNotDetermined - A status that indicates the user hasn’t yet granted or denied authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/notDetermined
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	// AuthorizationStatusRestricted - A status that indicates the app isn’t permitted to use media capture devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/restricted
	AuthorizationStatusRestricted AuthorizationStatus = 0
)

// AVCaptureColorSpace - An enumeration of color spaces a device can support.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
type CaptureColorSpace uint

// AVCaptureCinematicVideoFocusMode - Constants indicating the focus behavior when recording a Cinematic Video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
type CaptureCinematicVideoFocusMode uint

// AVCaptureLensStabilizationStatus - Constants that indicate the status of optical image stabilization hardware during a bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type CaptureLensStabilizationStatus uint

// AVCaptureDevicePosition - Constants that indicate the physical position of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type CaptureDevicePosition uint

const (
	// CaptureDevicePositionBack - A position on the subject-facing side of an iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/back
	CaptureDevicePositionBack CaptureDevicePosition = 0
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
	CapturePrimaryConstituentDeviceSwitchingBehaviorAuto CapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// CapturePrimaryConstituentDeviceSwitchingBehaviorLocked - The device locks camera switching to the active primary constituent device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/locked
	CapturePrimaryConstituentDeviceSwitchingBehaviorLocked CapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported - The device doesn’t support constituent device switching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/unsupported
	CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported CapturePrimaryConstituentDeviceSwitchingBehavior = 0
)

// AVCaptureOutputDataDroppedReason - Constants that define reasons for why the system dropped a frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type CaptureOutputDataDroppedReason uint

// AVCapturePhotoOutputCaptureReadiness - Constants that indicate whether the output is ready to receive capture requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum
type CapturePhotoOutputCaptureReadiness uint

// AVCaptureSessionInterruptionReason - Constants identifying the reason a capture session was interrupted, found in an 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type CaptureSessionInterruptionReason uint

const (
	// CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient - An interruption caused by the audio hardware temporarily being made unavailable (for example, for a phone call or alarm).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/audioDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonSensitiveContentMitigationActivated - An interruption caused by a   when it detects sensitive content on an associated  .  To resume your capture session, call your analyzer’s   method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/sensitiveContentMitigationActivated
	CaptureSessionInterruptionReasonSensitiveContentMitigationActivated CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient - An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure - An interruption due to system pressure, such as thermal duress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableDueToSystemPressure
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground - An interruption caused by the app being sent to the background while using a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableInBackground
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps - An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableWithMultipleForegroundApps
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureSessionInterruptionReason = 0
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

// AVPlayerNetworkResourcePriority - This defines the network resource priority for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type PlayerNetworkResourcePriority uint

// AVPlayerStatus - Status values that indicate whether a player can successfully play media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type PlayerStatus uint

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

// AVPlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type PlayerItemStatus uint

const (
	// PlayerItemStatusFailed - The item no longer plays due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/failed
	PlayerItemStatusFailed PlayerItemStatus = 0
	// PlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	PlayerItemStatusReadyToPlay PlayerItemStatus = 0
	// PlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	PlayerItemStatusUnknown PlayerItemStatus = 0
)


