// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVDelegatingPlaybackCoordinatorPauseCommand */


/* debug [class_header]: Header for AVDelegatingPlaybackCoordinatorPauseCommand */
// The class instance for the [DelegatingPlaybackCoordinatorPauseCommand] class.
var (
	DelegatingPlaybackCoordinatorPauseCommandClass     _DelegatingPlaybackCoordinatorPauseCommandClass
	DelegatingPlaybackCoordinatorPauseCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorPauseCommandClass() _DelegatingPlaybackCoordinatorPauseCommandClass {
	DelegatingPlaybackCoordinatorPauseCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorPauseCommandClass = _DelegatingPlaybackCoordinatorPauseCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPauseCommand")}
	})
	return DelegatingPlaybackCoordinatorPauseCommandClass
}

type _DelegatingPlaybackCoordinatorPauseCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DelegatingPlaybackCoordinatorPauseCommand */
// An interface definition for the [DelegatingPlaybackCoordinatorPauseCommand] class.
type IDelegatingPlaybackCoordinatorPauseCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	
/* debug [class_interface_properties]: Properties for DelegatingPlaybackCoordinatorPauseCommand */
	// properties:
	AnticipatedPlaybackRate() float32
	ShouldBufferInAnticipationOfPlayback() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DelegatingPlaybackCoordinatorPauseCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DelegatingPlaybackCoordinatorPauseCommand */
// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorPauseCommandClass) Alloc() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPauseCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorPauseCommandClass) New() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPauseCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorPauseCommand) Init() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPauseCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorPauseCommand) Autorelease() DelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPauseCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorPauseCommand creates a new DelegatingPlaybackCoordinatorPauseCommand instance.
func NewDelegatingPlaybackCoordinatorPauseCommand() DelegatingPlaybackCoordinatorPauseCommand {
	return getDelegatingPlaybackCoordinatorPauseCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DelegatingPlaybackCoordinatorPauseCommand */
// A command that indicates to pause playback.


// A command that indicates to pause playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand
type DelegatingPlaybackCoordinatorPauseCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

// DelegatingPlaybackCoordinatorPauseCommandFrom constructs a [DelegatingPlaybackCoordinatorPauseCommand] from an unsafe.Pointer.
//
// A command that indicates to pause playback.
func DelegatingPlaybackCoordinatorPauseCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPauseCommand {
	return DelegatingPlaybackCoordinatorPauseCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DelegatingPlaybackCoordinatorPauseCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DelegatingPlaybackCoordinatorPauseCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DelegatingPlaybackCoordinatorPauseCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DelegatingPlaybackCoordinatorPauseCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DelegatingPlaybackCoordinatorPauseCommand */

// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorPauseCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}/* debug [instance_properties/getter]: anticipatedPlaybackRate */


// A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/shouldBufferInAnticipationOfPlayback
func (d_ DelegatingPlaybackCoordinatorPauseCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}/* debug [instance_properties/getter]: shouldBufferInAnticipationOfPlayback */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDelegatingPlaybackCoordinatorPauseCommand */



