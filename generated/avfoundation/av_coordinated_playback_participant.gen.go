// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CoordinatedPlaybackParticipant] class.
var (
	CoordinatedPlaybackParticipantClass     _CoordinatedPlaybackParticipantClass
	CoordinatedPlaybackParticipantClassOnce sync.Once
)

func getCoordinatedPlaybackParticipantClass() _CoordinatedPlaybackParticipantClass {
	CoordinatedPlaybackParticipantClassOnce.Do(func() {
		CoordinatedPlaybackParticipantClass = _CoordinatedPlaybackParticipantClass{objc.GetClass("AVCoordinatedPlaybackParticipant")}
	})
	return CoordinatedPlaybackParticipantClass
}

type _CoordinatedPlaybackParticipantClass struct {
	class objc.Class
}

// An interface definition for the [CoordinatedPlaybackParticipant] class.
type ICoordinatedPlaybackParticipant interface {
	objectivec.IObject
}

// An object that represents a participant in a coordinated playback session.
//
// Access the other participants in a session through the playback coordinator’s property to determine their playback readiness and suspension reasons.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant
type CoordinatedPlaybackParticipant struct {
	objectivec.Object
}

// CoordinatedPlaybackParticipantFrom constructs a [CoordinatedPlaybackParticipant] from an unsafe.Pointer.
//
// An object that represents a participant in a coordinated playback session.
func CoordinatedPlaybackParticipantFrom(ptr unsafe.Pointer) CoordinatedPlaybackParticipant {
	return CoordinatedPlaybackParticipant{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoordinatedPlaybackParticipantClass) Alloc() CoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoordinatedPlaybackParticipantClass) New() CoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoordinatedPlaybackParticipant) Init() CoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoordinatedPlaybackParticipant) Autorelease() CoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoordinatedPlaybackParticipant creates a new CoordinatedPlaybackParticipant instance.
func NewCoordinatedPlaybackParticipant() CoordinatedPlaybackParticipant {
	return getCoordinatedPlaybackParticipantClass().New()
}




