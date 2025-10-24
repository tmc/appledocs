//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Player


// iOS-only properties

// A Boolean value that indicates whether the player allows AirPlay video playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/allowsAirPlayVideo
func (p_ Player) AllowsAirPlayVideo() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsAirPlayVideo"))
	return rv
}
func (p_ Player) SetAllowsAirPlayVideo(value bool) {
	p_.ID.Send(objc.RegisterName("setAllowsAirPlayVideo:"), value)
}

// Whether the player’s audio output is suppressed due to being on a non-mixable audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audioOutputSuppressedDueToNonMixableAudioRoute
func (p_ Player) AudioOutputSuppressedDueToNonMixableAudioRoute() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("audioOutputSuppressedDueToNonMixableAudioRoute"))
	return rv
}

// The video gravity of the player for external playback mode only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/externalPlaybackVideoGravity
func (p_ Player) ExternalPlaybackVideoGravity() LayerVideoGravity /* not a class type */ {
	rv := objc.Send[LayerVideoGravity](p_.ID, objc.Sel("externalPlaybackVideoGravity"))
	return rv
}
func (p_ Player) SetExternalPlaybackVideoGravity(value LayerVideoGravity /* not a class type */) {
	p_.ID.Send(objc.RegisterName("setExternalPlaybackVideoGravity:"), value)
}

// The AVPlayer’s intended spatial audio experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/intendedSpatialAudioExperience-3uy8g
func (p_ Player) IntendedSpatialAudioExperience() objc.IObject /* cross-framework: SpatialAudioExperience */ {
	rv := objc.Send[audiotoolbox.SpatialAudioExperience](p_.ID, objc.Sel("intendedSpatialAudioExperience"))
	return rv
}
func (p_ Player) SetIntendedSpatialAudioExperience(value objc.IObject /* cross-framework: SpatialAudioExperience */) {
	p_.ID.Send(objc.RegisterName("setIntendedSpatialAudioExperience:"), value)
}

// A Boolean value that indicates whether the player is playing video through AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isAirPlayVideoActive
func (p_ Player) AirPlayVideoActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("airPlayVideoActive"))
	return rv
}

// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preventsAutomaticBackgroundingDuringVideoPlayback
func (p_ Player) PreventsAutomaticBackgroundingDuringVideoPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}
func (p_ Player) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool) {
	p_.ID.Send(objc.RegisterName("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}

// A Boolean value that indicates whether the player automatically switches to AirPlay Video while AirPlay Screen is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/usesAirPlayVideoWhileAirPlayScreenIsActive
func (p_ Player) UsesAirPlayVideoWhileAirPlayScreenIsActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesAirPlayVideoWhileAirPlayScreenIsActive"))
	return rv
}
func (p_ Player) SetUsesAirPlayVideoWhileAirPlayScreenIsActive(value bool) {
	p_.ID.Send(objc.RegisterName("setUsesAirPlayVideoWhileAirPlayScreenIsActive:"), value)
}

// A Boolean value that indicates whether the player should automatically switch to external playback mode while the external screen mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/usesExternalPlaybackWhileExternalScreenIsActive
func (p_ Player) UsesExternalPlaybackWhileExternalScreenIsActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesExternalPlaybackWhileExternalScreenIsActive"))
	return rv
}
func (p_ Player) SetUsesExternalPlaybackWhileExternalScreenIsActive(value bool) {
	p_.ID.Send(objc.RegisterName("setUsesExternalPlaybackWhileExternalScreenIsActive:"), value)
}




