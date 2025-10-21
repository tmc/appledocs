// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlaybackCoordinator] class.
var (
	PlaybackCoordinatorClass     _PlaybackCoordinatorClass
	PlaybackCoordinatorClassOnce sync.Once
)

func getPlaybackCoordinatorClass() _PlaybackCoordinatorClass {
	PlaybackCoordinatorClassOnce.Do(func() {
		PlaybackCoordinatorClass = _PlaybackCoordinatorClass{objc.GetClass("AVPlaybackCoordinator")}
	})
	return PlaybackCoordinatorClass
}

type _PlaybackCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [PlaybackCoordinator] class.
type IPlaybackCoordinator interface {
	objectivec.IObject
	ParticipantLimitForWaitingOutSuspensionsWithReason(reason unsafe.Pointer) int
	SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason unsafe.Pointer)
}

// An object that coordinates the playback of players in a connected group.
//
// The framework provides two playback coordinator subclasses that manage different types of player objects: coordinates the state of objects. If your app uses , continue to use its standard interfaces to control playback. The coordinator intercepts changes to the player’s rate and time, and propagates them to other players in the group. coordinates the state of custom player objects. If your app uses a custom player, such as one that renders media using and , use this object to coordinate group playback. Adopt the coordinator’s delegate protocol so that your player responds to the commands that the coordinator issues.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator
type PlaybackCoordinator struct {
	objectivec.Object
}

// PlaybackCoordinatorFrom constructs a [PlaybackCoordinator] from an unsafe.Pointer.
//
// An object that coordinates the playback of players in a connected group.
func PlaybackCoordinatorFrom(ptr unsafe.Pointer) PlaybackCoordinator {
	return PlaybackCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlaybackCoordinatorClass) Alloc() PlaybackCoordinator {
	rv := objc.Send[PlaybackCoordinator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlaybackCoordinatorClass) New() PlaybackCoordinator {
	rv := objc.Send[PlaybackCoordinator](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlaybackCoordinator) Init() PlaybackCoordinator {
	rv := objc.Send[PlaybackCoordinator](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlaybackCoordinator) Autorelease() PlaybackCoordinator {
	rv := objc.Send[PlaybackCoordinator](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlaybackCoordinator creates a new PlaybackCoordinator instance.
func NewPlaybackCoordinator() PlaybackCoordinator {
	return getPlaybackCoordinatorClass().New()
}


// Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/participantLimitForWaitingOutSuspensions(withReason:)
func (p_ PlaybackCoordinator) ParticipantLimitForWaitingOutSuspensionsWithReason(reason unsafe.Pointer) int {
	rv := objc.Send[int](p_.ID, objc.Sel("participantLimitForWaitingOutSuspensionsWithReason:"), reason)
	return rv
}

// Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/setParticipantLimit(_:forWaitingOutSuspensionsWithReason:)
func (p_ PlaybackCoordinator) SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParticipantLimit:forWaitingOutSuspensionsWithReason:"), participantLimit, reason)
}

// The identifiers of the other participants in a group.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/otherParticipants
func (p_ PlaybackCoordinator) OtherParticipants() []CoordinatedPlaybackParticipant {
	rv := objc.Send[[]CoordinatedPlaybackParticipant](p_.ID, objc.Sel("otherParticipants"))
	return rv
}

// The reasons a coordinator is currently unable to participate in a group playback activity.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/suspensionReasons
func (p_ PlaybackCoordinator) SuspensionReasons() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("suspensionReasons"))
	return rv
}

// A Boolean value that indicates whether participants mirror the originator’s stop time when they pause.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/pausesnapstomediatimeoforiginator
func (p_ PlaybackCoordinator) PauseSnapsToMediaTimeOfOriginator() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pauseSnapsToMediaTimeOfOriginator"))
	return rv
}


// SetPauseSnapsToMediaTimeOfOriginator sets the value of the pauseSnapsToMediaTimeOfOriginator property.
// A Boolean value that indicates whether participants mirror the originator’s stop time when they pause.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/pausesnapstomediatimeoforiginator
func (p_ PlaybackCoordinator) SetPauseSnapsToMediaTimeOfOriginator(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPauseSnapsToMediaTimeOfOriginator:"), value)
}

// The reasons that cause a coordinator to suspend playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/suspensionreasonsthattriggerwaiting
func (p_ PlaybackCoordinator) SuspensionReasonsThatTriggerWaiting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("suspensionReasonsThatTriggerWaiting"))
	return rv
}


// SetSuspensionReasonsThatTriggerWaiting sets the value of the suspensionReasonsThatTriggerWaiting property.
// The reasons that cause a coordinator to suspend playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/suspensionreasonsthattriggerwaiting
func (p_ PlaybackCoordinator) SetSuspensionReasonsThatTriggerWaiting(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSuspensionReasonsThatTriggerWaiting:"), value)
}



