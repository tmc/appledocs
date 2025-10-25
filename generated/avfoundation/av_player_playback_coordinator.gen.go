// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerPlaybackCoordinator */


/* debug [class_header]: Header for AVPlayerPlaybackCoordinator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerPlaybackCoordinator */
// An interface definition for the [PlayerPlaybackCoordinator] class.
type IPlayerPlaybackCoordinator interface {
	IPlaybackCoordinator
	
/* debug [class_interface_properties]: Properties for PlayerPlaybackCoordinator */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	PlaybackCoordinationMedium() IAVPlaybackCoordinationMedium
	Player() IAVPlayer
	TimeControlStatus() objectivec.IObject
	SetTimeControlStatus(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerPlaybackCoordinator */
	// methods:
	CoordinateUsingCoordinationMediumError(coordinationMedium IAVPlaybackCoordinationMedium, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerPlaybackCoordinator */
// Alloc allocates a new instance without initialization.
func (pc _PlayerPlaybackCoordinatorClass) Alloc() PlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerPlaybackCoordinator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerPlaybackCoordinator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerPlaybackCoordinator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerPlaybackCoordinator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerPlaybackCoordinator */

// Connects the playback coordinator to the coordination medium
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/coordinate(using:)
func (p_ PlayerPlaybackCoordinator) CoordinateUsingCoordinationMediumError(coordinationMedium IAVPlaybackCoordinationMedium, outError objectivec.IObject) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("coordinateUsingCoordinationMedium:error:"), coordinationMedium, outError)
	return rv
}/* debug [instance_methods/method]: CoordinateUsingCoordinationMediumError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerPlaybackCoordinator */

// A delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/delegate
func (p_ PlayerPlaybackCoordinator) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/delegate
func (p_ PlayerPlaybackCoordinator) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/playbackCoordinationMedium
func (p_ PlayerPlaybackCoordinator) PlaybackCoordinationMedium() IAVPlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("playbackCoordinationMedium"))
	return rv
}/* debug [instance_properties/getter]: playbackCoordinationMedium */


// A player that participates in coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/player
func (p_ PlayerPlaybackCoordinator) Player() IAVPlayer {
	rv := objc.Send[Player](p_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) TimeControlStatus() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}/* debug [instance_properties/getter]: timeControlStatus */


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerPlaybackCoordinator) SetTimeControlStatus(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeControlStatus:"), value)
}/* debug [instance_properties/setter]: timeControlStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerPlaybackCoordinator */



