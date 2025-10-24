//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for RoutingPlaybackArbiter


// iOS-only properties

// The participant that has priority to play on external playback interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForExternalPlayback
func (r_ RoutingPlaybackArbiter) PreferredParticipantForExternalPlayback() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("preferredParticipantForExternalPlayback"))
	return rv
}
func (r_ RoutingPlaybackArbiter) SetPreferredParticipantForExternalPlayback(value unsafe.Pointer) {
	r_.ID.Send(objc.RegisterName("setPreferredParticipantForExternalPlayback:"), value)
}

// The participant that has priority to play audio when it’s not possible to play multiple audio sources concurrently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForNonMixableAudioRoutes
func (r_ RoutingPlaybackArbiter) PreferredParticipantForNonMixableAudioRoutes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("preferredParticipantForNonMixableAudioRoutes"))
	return rv
}
func (r_ RoutingPlaybackArbiter) SetPreferredParticipantForNonMixableAudioRoutes(value unsafe.Pointer) {
	r_.ID.Send(objc.RegisterName("setPreferredParticipantForNonMixableAudioRoutes:"), value)
}





