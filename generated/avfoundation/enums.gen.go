// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// Enum types and constants
// AssetWriterStatus - Values that indicate the state of an asset writer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum
type AssetWriterStatus uint

// AudioSpatializationFormats - A structure that defines the spatialization formats that a player item supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats
type AudioSpatializationFormats uint

// AuthorizationStatus - Constants that indicate the status of an app’s authorization to capture media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AuthorizationStatus uint

// CaptureLensStabilizationStatus - Constants that indicate the status of optical image stabilization hardware during a bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type CaptureLensStabilizationStatus uint

// CaptureDevicePosition - Constants that indicate the physical position of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type CaptureDevicePosition uint

// CaptureMultichannelAudioMode - Constants that indicate the modes of multichannel audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode
type CaptureMultichannelAudioMode uint

// CaptureOutputDataDroppedReason - Constants that define reasons for why the system dropped a frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type CaptureOutputDataDroppedReason uint

// CaptureSessionInterruptionReason - Constants identifying the reason a capture session was interrupted, found in an 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type CaptureSessionInterruptionReason uint

const (
	// CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient - An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground - An interruption caused by the app being sent to the background while using a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableInBackground
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps - An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableWithMultipleForegroundApps
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureSessionInterruptionReason = 0
)

// ContentAuthorizationStatus - A value representing the status of a content authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus
type ContentAuthorizationStatus uint

// Error - An enumeration that defines the errors that framework operations can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type Error uint

const (
	// ErrorIncorrectlyConfigured - The system is incorrectly configured for the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incorrectlyConfigured
	ErrorIncorrectlyConfigured Error = 0
	// ErrorUndecodableMediaData - The system couldn’t decode the media data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/undecodableMediaData
	ErrorUndecodableMediaData Error = 0
)

// PlayerActionAtItemEnd - The actions a player can take when it finishes playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
type PlayerActionAtItemEnd uint

// PlayerHDRMode - A bitfield type that specifies an HDR mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode
type PlayerHDRMode uint

// PlayerNetworkResourcePriority - This defines the network resource priority for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type PlayerNetworkResourcePriority uint

// PlayerStatus - Status values that indicate whether a player can successfully play media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type PlayerStatus uint

// PlayerTimeControlStatus - Constants that indicate the state of playback control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum
type PlayerTimeControlStatus uint

const (
	// PlayerTimeControlStatusPlaying - A state that indicates that the player is currently playing media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/playing
	PlayerTimeControlStatusPlaying PlayerTimeControlStatus = 0
	// PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate - A state that indicates that the player is waiting for network conditions to improve before it can start or resume playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/waitingToPlayAtSpecifiedRate
	PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate PlayerTimeControlStatus = 0
)

// PlayerAudiovisualBackgroundPlaybackPolicy - Policies that describe playback behavior when an app transitions to the background while playing video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type PlayerAudiovisualBackgroundPlaybackPolicy uint

// PlayerInterstitialEventAssetListResponseStatus - Constants that describe the status of the asset list response for an interstitial event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus
type PlayerInterstitialEventAssetListResponseStatus uint

// PlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type PlayerItemStatus uint

const (
	// PlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	PlayerItemStatusReadyToPlay PlayerItemStatus = 0
	// PlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	PlayerItemStatusUnknown PlayerItemStatus = 0
)

// VariantPreferences - Defines the preferences the player item uses when selecting variant playlists.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences
type VariantPreferences uint


