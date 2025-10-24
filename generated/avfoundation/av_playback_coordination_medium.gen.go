// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlaybackCoordinationMedium */


/* debug [class_header]: Header for AVPlaybackCoordinationMedium */
// The class instance for the [PlaybackCoordinationMedium] class.
var (
	PlaybackCoordinationMediumClass     _PlaybackCoordinationMediumClass
	PlaybackCoordinationMediumClassOnce sync.Once
)

func getPlaybackCoordinationMediumClass() _PlaybackCoordinationMediumClass {
	PlaybackCoordinationMediumClassOnce.Do(func() {
		PlaybackCoordinationMediumClass = _PlaybackCoordinationMediumClass{objc.GetClass("AVPlaybackCoordinationMedium")}
	})
	return PlaybackCoordinationMediumClass
}

type _PlaybackCoordinationMediumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlaybackCoordinationMedium */
// An interface definition for the [PlaybackCoordinationMedium] class.
type IPlaybackCoordinationMedium interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlaybackCoordinationMedium */
	// properties:
	ConnectedPlaybackCoordinators() []PlayerPlaybackCoordinator
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlaybackCoordinationMedium */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlaybackCoordinationMedium */
// Alloc allocates a new instance without initialization.
func (pc _PlaybackCoordinationMediumClass) Alloc() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlaybackCoordinationMediumClass) New() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlaybackCoordinationMedium) Init() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlaybackCoordinationMedium) Autorelease() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlaybackCoordinationMedium creates a new PlaybackCoordinationMedium instance.
func NewPlaybackCoordinationMedium() PlaybackCoordinationMedium {
	return getPlaybackCoordinationMediumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlaybackCoordinationMedium */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium
type PlaybackCoordinationMedium struct {
	objectivec.Object
}

// PlaybackCoordinationMediumFrom constructs a [PlaybackCoordinationMedium] from an unsafe.Pointer.
func PlaybackCoordinationMediumFrom(ptr unsafe.Pointer) PlaybackCoordinationMedium {
	return PlaybackCoordinationMedium{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlaybackCoordinationMedium */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlaybackCoordinationMedium */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlaybackCoordinationMedium */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlaybackCoordinationMedium */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlaybackCoordinationMedium */

// All playback coordinators that are connected to the coordination medium.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium/connectedPlaybackCoordinators
func (p_ PlaybackCoordinationMedium) ConnectedPlaybackCoordinators() []PlayerPlaybackCoordinator {
	rv := objc.Send[[]PlayerPlaybackCoordinator](p_.ID, objc.Sel("connectedPlaybackCoordinators"))
	return rv
}/* debug [instance_properties/getter]: connectedPlaybackCoordinators */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlaybackCoordinationMedium */


