// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Identifier() foundation.UUID /* not a class type */
	SetIdentifier(value foundation.UUID /* not a class type */)
	IsReadyToPlay() bool /* primitive/slice/pointer */
	SetIsReadyToPlay(value bool /* primitive/slice/pointer */)
	SuspensionReasons() unsafe.Pointer
	SetSuspensionReasons(value unsafe.Pointer)
	OtherParticipants() IAVCoordinatedPlaybackParticipant
	SetOtherParticipants(value IAVCoordinatedPlaybackParticipant)
	// methods:
}

// An object that represents a participant in a coordinated playback session.
//
// Access the other participants in a session through the playback coordinator’s property to determine their playback readiness and suspension reasons.


// An object that represents a participant in a coordinated playback session.
//
// [Full Topic]
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



// A unique identifier for the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/identifier
func (c_ CoordinatedPlaybackParticipant) Identifier() foundation.UUID /* not a class type */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("identifier"))
	return rv
}


// A unique identifier for the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/identifier
func (c_ CoordinatedPlaybackParticipant) SetIdentifier(value foundation.UUID /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}


// A Boolean value that indicates whether the participant is ready to start coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/isreadytoplay
func (c_ CoordinatedPlaybackParticipant) IsReadyToPlay() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isReadyToPlay"))
	return rv
}


// A Boolean value that indicates whether the participant is ready to start coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/isreadytoplay
func (c_ CoordinatedPlaybackParticipant) SetIsReadyToPlay(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsReadyToPlay:"), value)
}


// The reasons a participant isn’t currently participating in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/suspensionreasons
func (c_ CoordinatedPlaybackParticipant) SuspensionReasons() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("suspensionReasons"))
	return rv
}


// The reasons a participant isn’t currently participating in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/suspensionreasons
func (c_ CoordinatedPlaybackParticipant) SetSuspensionReasons(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSuspensionReasons:"), value)
}


// The identifiers of the other participants in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipants
func (c_ CoordinatedPlaybackParticipant) OtherParticipants() IAVCoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](c_.ID, objc.Sel("otherParticipants"))
	return rv
}


// The identifiers of the other participants in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipants
func (c_ CoordinatedPlaybackParticipant) SetOtherParticipants(value IAVCoordinatedPlaybackParticipant) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOtherParticipants:"), value)
}



