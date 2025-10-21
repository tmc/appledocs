// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PlayerPlaybackCoordinator] class.
var (
	PlayerPlaybackCoordinatorClass     _PlayerPlaybackCoordinatorClass
	PlayerPlaybackCoordinatorClassOnce sync.Once
)

func getPlayerPlaybackCoordinatorClass() _PlayerPlaybackCoordinatorClass {
	PlayerPlaybackCoordinatorClassOnce.Do(func() {
		PlayerPlaybackCoordinatorClass = _PlayerPlaybackCoordinatorClass{objc.GetClass("AVPlayerPlaybackCoordinator")}
	})
	return PlayerPlaybackCoordinatorClass
}

type _PlayerPlaybackCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [PlayerPlaybackCoordinator] class.
type IPlayerPlaybackCoordinator interface {
	IPlaybackCoordinator
	CoordinateUsingCoordinationMediumError(coordinationMedium unsafe.Pointer, outError unsafe.Pointer) bool
}

// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
//
// This object coordinates the state of objects. You don’t create an instance of the coordinator, but instead access the player’s instance through its property. Use the standard interfaces of to control playback in your app. The coordinator automatically intercepts calls that affect transport control state, like , , and , and propagates them to other participants in the group when appropriate. Similarly, the coordinator observes rate and time changes from other participants and imposes them on the player. If this occurs, the player item posts notifications that identify the originating participant. This object may automatically suspend coordinated playback when a system state change causes the player’s value to change from a playing state to a waiting or paused state. A suspension that begins because the player enters a waiting state due to an event like a network stall or interstitial playback, ends automatically when the player finishes waiting. However, if the system pauses playback due to a system state change, such as an audio session interruption, the suspension ends only after the player’s rate changes back to nonzero.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator
type PlayerPlaybackCoordinator struct {
	PlaybackCoordinator
}

// PlayerPlaybackCoordinatorFrom constructs a [PlayerPlaybackCoordinator] from an unsafe.Pointer.
//
// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
func PlayerPlaybackCoordinatorFrom(ptr unsafe.Pointer) PlayerPlaybackCoordinator {
	return PlayerPlaybackCoordinator{
		PlaybackCoordinator: PlaybackCoordinatorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerPlaybackCoordinatorClass) Alloc() PlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerPlaybackCoordinatorClass) New() PlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerPlaybackCoordinator) Init() PlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerPlaybackCoordinator) Autorelease() PlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerPlaybackCoordinator creates a new PlayerPlaybackCoordinator instance.
func NewPlayerPlaybackCoordinator() PlayerPlaybackCoordinator {
	return getPlayerPlaybackCoordinatorClass().New()
}


// Connects the playback coordinator to the coordination medium
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/coordinate(using:)
func (p_ PlayerPlaybackCoordinator) CoordinateUsingCoordinationMediumError(coordinationMedium unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("coordinateUsingCoordinationMedium:error:"), coordinationMedium, outError)
	return rv
}

// A delegate object for the playback coordinator.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/delegate
func (p_ PlayerPlaybackCoordinator) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate object for the playback coordinator.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/delegate
func (p_ PlayerPlaybackCoordinator) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/playbackCoordinationMedium
func (p_ PlayerPlaybackCoordinator) PlaybackCoordinationMedium() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackCoordinationMedium"))
	return rv
}

// A player that participates in coordinated playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/player
func (p_ PlayerPlaybackCoordinator) Player() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("player"))
	return rv
}

// The playback coordinator for the player.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/playbackcoordinator
func (p_ PlayerPlaybackCoordinator) PlaybackCoordinator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackCoordinator"))
	return rv
}


// SetPlaybackCoordinator sets the value of the playbackCoordinator property.
// The playback coordinator for the player.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/playbackcoordinator
func (p_ PlayerPlaybackCoordinator) SetPlaybackCoordinator(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackCoordinator:"), value)
}

// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) TimeControlStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}


// SetTimeControlStatus sets the value of the timeControlStatus property.
// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) SetTimeControlStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeControlStatus:"), value)
}



