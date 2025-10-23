// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// Enum types and constants
// AVAudioSpatializationFormats - A structure that defines the spatialization formats that a player item supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats
type AVAudioSpatializationFormats uint

// AVContentAuthorizationStatus - A value representing the status of a content authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus
type AVContentAuthorizationStatus uint

// AVError - An enumeration that defines the errors that framework operations can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type AVError uint

const (
	// AVErrorIncorrectlyConfigured - The system is incorrectly configured for the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incorrectlyConfigured
	AVErrorIncorrectlyConfigured AVError = 0
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

// AVPlayerAudiovisualBackgroundPlaybackPolicy - Policies that describe playback behavior when an app transitions to the background while playing video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type AVPlayerAudiovisualBackgroundPlaybackPolicy uint

// AVPlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type AVPlayerItemStatus uint

const (
	// AVPlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	AVPlayerItemStatusReadyToPlay AVPlayerItemStatus = 0
	// AVPlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	AVPlayerItemStatusUnknown AVPlayerItemStatus = 0
)

// AVVariantPreferences - Defines the preferences the player item uses when selecting variant playlists.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences
type AVVariantPreferences uint


