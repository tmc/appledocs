// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RoutingPlaybackArbiter] class.
var (
	RoutingPlaybackArbiterClass     _RoutingPlaybackArbiterClass
	RoutingPlaybackArbiterClassOnce sync.Once
)

func getRoutingPlaybackArbiterClass() _RoutingPlaybackArbiterClass {
	RoutingPlaybackArbiterClassOnce.Do(func() {
		RoutingPlaybackArbiterClass = _RoutingPlaybackArbiterClass{objc.GetClass("AVRoutingPlaybackArbiter")}
	})
	return RoutingPlaybackArbiterClass
}

type _RoutingPlaybackArbiterClass struct {
	class objc.Class
}

// An interface definition for the [RoutingPlaybackArbiter] class.
type IRoutingPlaybackArbiter interface {
	objectivec.IObject
	PreferredParticipantForExternalPlayback() objc.ID
	SetPreferredParticipantForExternalPlayback(value objc.ID)
	PreferredParticipantForNonMixableAudioRoutes() objc.ID
	SetPreferredParticipantForNonMixableAudioRoutes(value objc.ID)
}

// An object that manages playback routing preferences.
//
// This object manages instances of for arbitration of media playback routing priorities and preferences on restricted playback interfaces. The playback routing arbiter is responsible for collecting and applying preferences, such as priorities in non-mixable audio routes and external playback states where the number of allowed players is limited.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter
type RoutingPlaybackArbiter struct {
	objectivec.Object
}

// RoutingPlaybackArbiterFrom constructs a [RoutingPlaybackArbiter] from an unsafe.Pointer.
//
// An object that manages playback routing preferences.
func RoutingPlaybackArbiterFrom(ptr unsafe.Pointer) RoutingPlaybackArbiter {
	return RoutingPlaybackArbiter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RoutingPlaybackArbiterClass) Alloc() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RoutingPlaybackArbiterClass) New() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RoutingPlaybackArbiter) Init() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RoutingPlaybackArbiter) Autorelease() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRoutingPlaybackArbiter creates a new RoutingPlaybackArbiter instance.
func NewRoutingPlaybackArbiter() RoutingPlaybackArbiter {
	return getRoutingPlaybackArbiterClass().New()
}


// Returns the singleton playback arbiter instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/shared()
func (rc _RoutingPlaybackArbiterClass) SharedRoutingPlaybackArbiter() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("sharedRoutingPlaybackArbiter"))
	return rv
}

// The participant that has priority to play on external playback interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForExternalPlayback
func (r_ RoutingPlaybackArbiter) PreferredParticipantForExternalPlayback() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("preferredParticipantForExternalPlayback"))
	return rv
}


// SetPreferredParticipantForExternalPlayback sets the value of the preferredParticipantForExternalPlayback property.
// The participant that has priority to play on external playback interfaces.

//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForExternalPlayback
func (r_ RoutingPlaybackArbiter) SetPreferredParticipantForExternalPlayback(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreferredParticipantForExternalPlayback:"), value)
}

// The participant that has priority to play audio when it’s not possible to play multiple audio sources concurrently.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForNonMixableAudioRoutes
func (r_ RoutingPlaybackArbiter) PreferredParticipantForNonMixableAudioRoutes() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("preferredParticipantForNonMixableAudioRoutes"))
	return rv
}


// SetPreferredParticipantForNonMixableAudioRoutes sets the value of the preferredParticipantForNonMixableAudioRoutes property.
// The participant that has priority to play audio when it’s not possible to play multiple audio sources concurrently.

//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/preferredParticipantForNonMixableAudioRoutes
func (r_ RoutingPlaybackArbiter) SetPreferredParticipantForNonMixableAudioRoutes(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreferredParticipantForNonMixableAudioRoutes:"), value)
}




