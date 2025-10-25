// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCoordinatedPlaybackParticipant */


/* debug [class_header]: Header for AVCoordinatedPlaybackParticipant */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CoordinatedPlaybackParticipant */
// An interface definition for the [CoordinatedPlaybackParticipant] class.
type ICoordinatedPlaybackParticipant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CoordinatedPlaybackParticipant */
	// properties:
	Identifier() foundation.UUID
	ReadyToPlay() bool
	SuspensionReasons() []string
	IsReadyToPlay() bool
	SetIsReadyToPlay(value bool)
	OtherParticipants() IAVCoordinatedPlaybackParticipant
	SetOtherParticipants(value IAVCoordinatedPlaybackParticipant)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CoordinatedPlaybackParticipant */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CoordinatedPlaybackParticipant */
// Alloc allocates a new instance without initialization.
func (cc _CoordinatedPlaybackParticipantClass) Alloc() CoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CoordinatedPlaybackParticipant */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CoordinatedPlaybackParticipant *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CoordinatedPlaybackParticipant */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CoordinatedPlaybackParticipant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CoordinatedPlaybackParticipant */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CoordinatedPlaybackParticipant */

// A unique identifier for the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/identifier
func (c_ CoordinatedPlaybackParticipant) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether the participant is ready to start coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/isReadyToPlay
func (c_ CoordinatedPlaybackParticipant) ReadyToPlay() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("readyToPlay"))
	return rv
}/* debug [instance_properties/getter]: readyToPlay */


// The reasons a participant isn’t currently participating in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/suspensionReasons
func (c_ CoordinatedPlaybackParticipant) SuspensionReasons() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("suspensionReasons"))
	return rv
}/* debug [instance_properties/getter]: suspensionReasons */


// A Boolean value that indicates whether the participant is ready to start coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/isreadytoplay
func (c_ CoordinatedPlaybackParticipant) IsReadyToPlay() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isReadyToPlay"))
	return rv
}/* debug [instance_properties/getter]: isReadyToPlay */


// A Boolean value that indicates whether the participant is ready to start coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybackparticipant/isreadytoplay
func (c_ CoordinatedPlaybackParticipant) SetIsReadyToPlay(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsReadyToPlay:"), value)
}/* debug [instance_properties/setter]: isReadyToPlay */


// The identifiers of the other participants in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipants
func (c_ CoordinatedPlaybackParticipant) OtherParticipants() IAVCoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](c_.ID, objc.Sel("otherParticipants"))
	return rv
}/* debug [instance_properties/getter]: otherParticipants */


// The identifiers of the other participants in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipants
func (c_ CoordinatedPlaybackParticipant) SetOtherParticipants(value IAVCoordinatedPlaybackParticipant) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOtherParticipants:"), value)
}/* debug [instance_properties/setter]: otherParticipants */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCoordinatedPlaybackParticipant */



