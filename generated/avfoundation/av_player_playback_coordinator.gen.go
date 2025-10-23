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
	// properties:
	TimeControlStatus() unsafe.Pointer
	SetTimeControlStatus(value unsafe.Pointer)
	Delegate() AVPlayerPlaybackCoordinatorDelegate /* foo */
	SetDelegate(value AVPlayerPlaybackCoordinatorDelegate /* foo */)
	PlaybackCoordinationMedium() AVPlaybackCoordinationMedium /* foo */
	SetPlaybackCoordinationMedium(value AVPlaybackCoordinationMedium /* foo */)
	Player() IAVPlayer
	SetPlayer(value IAVPlayer)
	// methods:
}

// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
//
// This object coordinates the state of objects. You don’t create an instance of the coordinator, but instead access the player’s instance through its property. Use the standard interfaces of to control playback in your app. The coordinator automatically intercepts calls that affect transport control state, like , , and , and propagates them to other participants in the group when appropriate. Similarly, the coordinator observes rate and time changes from other participants and imposes them on the player. If this occurs, the player item posts notifications that identify the originating participant. This object may automatically suspend coordinated playback when a system state change causes the player’s value to change from a playing state to a waiting or paused state. A suspension that begins because the player enters a waiting state due to an event like a network stall or interstitial playback, ends automatically when the player finishes waiting. However, if the system pauses playback due to a system state change, such as an audio session interruption, the suspension ends only after the player’s rate changes back to nonzero.


// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
//
// [Full Topic]
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



// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) TimeControlStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) SetTimeControlStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeControlStatus:"), value)
}


// A delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/delegate
func (p_ PlayerPlaybackCoordinator) Delegate() AVPlayerPlaybackCoordinatorDelegate /* foo */ {
	rv := objc.Send[PlayerPlaybackCoordinatorDelegate](p_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/delegate
func (p_ PlayerPlaybackCoordinator) SetDelegate(value AVPlayerPlaybackCoordinatorDelegate /* foo */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/playbackcoordinationmedium
func (p_ PlayerPlaybackCoordinator) PlaybackCoordinationMedium() AVPlaybackCoordinationMedium /* foo */ {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("playbackCoordinationMedium"))
	return rv
}


// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/playbackcoordinationmedium
func (p_ PlayerPlaybackCoordinator) SetPlaybackCoordinationMedium(value AVPlaybackCoordinationMedium /* foo */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackCoordinationMedium:"), value)
}


// A player that participates in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/player
func (p_ PlayerPlaybackCoordinator) Player() IAVPlayer {
	rv := objc.Send[Player](p_.ID, objc.Sel("player"))
	return rv
}


// A player that participates in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerplaybackcoordinator/player
func (p_ PlayerPlaybackCoordinator) SetPlayer(value IAVPlayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}



