// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlaybackCoordinator] class.
var (
	aVPlaybackCoordinatorClass     _AVPlaybackCoordinatorClass
	aVPlaybackCoordinatorClassOnce sync.Once
)

func getAVPlaybackCoordinatorClass() _AVPlaybackCoordinatorClass {
	aVPlaybackCoordinatorClassOnce.Do(func() {
		aVPlaybackCoordinatorClass = _AVPlaybackCoordinatorClass{objc.GetClass("AVPlaybackCoordinator")}
	})
	return aVPlaybackCoordinatorClass
}

type _AVPlaybackCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [AVPlaybackCoordinator] class.
type IAVPlaybackCoordinator interface {
	objectivec.IObject
	ParticipantLimitForWaitingOutSuspensionsWithReason(reason unsafe.Pointer) int
	SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason unsafe.Pointer)
}

// An object that coordinates the playback of players in a connected group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator
type AVPlaybackCoordinator struct {
	objectivec.Object
}

// AVPlaybackCoordinatorFrom constructs a [AVPlaybackCoordinator] from an unsafe.Pointer.
//
// An object that coordinates the playback of players in a connected group.
func AVPlaybackCoordinatorFrom(ptr unsafe.Pointer) AVPlaybackCoordinator {
	return AVPlaybackCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlaybackCoordinatorClass) Alloc() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlaybackCoordinatorClass) New() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlaybackCoordinator) Init() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlaybackCoordinator) Autorelease() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlaybackCoordinator creates a new AVPlaybackCoordinator instance.
func NewAVPlaybackCoordinator() AVPlaybackCoordinator {
	return getAVPlaybackCoordinatorClass().New()
}


// Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/participantLimitForWaitingOutSuspensions(withReason:)
func (a_ AVPlaybackCoordinator) ParticipantLimitForWaitingOutSuspensionsWithReason(reason unsafe.Pointer) int {
	rv := objc.Send[int](a_.ID, objc.Sel("participantLimitForWaitingOutSuspensionsWithReason:"), reason)
	return rv
}
// Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/setParticipantLimit(_:forWaitingOutSuspensionsWithReason:)
func (a_ AVPlaybackCoordinator) SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParticipantLimit:forWaitingOutSuspensionsWithReason:"), participantLimit, reason)
}


