// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVDelegatingPlaybackCoordinatorSeekCommand */


/* debug [class_header]: Header for AVDelegatingPlaybackCoordinatorSeekCommand */
// The class instance for the [DelegatingPlaybackCoordinatorSeekCommand] class.
var (
	DelegatingPlaybackCoordinatorSeekCommandClass     _DelegatingPlaybackCoordinatorSeekCommandClass
	DelegatingPlaybackCoordinatorSeekCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorSeekCommandClass() _DelegatingPlaybackCoordinatorSeekCommandClass {
	DelegatingPlaybackCoordinatorSeekCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorSeekCommandClass = _DelegatingPlaybackCoordinatorSeekCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorSeekCommand")}
	})
	return DelegatingPlaybackCoordinatorSeekCommandClass
}

type _DelegatingPlaybackCoordinatorSeekCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DelegatingPlaybackCoordinatorSeekCommand */
// An interface definition for the [DelegatingPlaybackCoordinatorSeekCommand] class.
type IDelegatingPlaybackCoordinatorSeekCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	
/* debug [class_interface_properties]: Properties for DelegatingPlaybackCoordinatorSeekCommand */
	// properties:
	AnticipatedPlaybackRate() float32
	CompletionDueDate() objc.IObject /* cross-framework: NSDate */
	ItemTime() objc.IObject /* cross-framework: Time */
	ShouldBufferInAnticipationOfPlayback() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DelegatingPlaybackCoordinatorSeekCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DelegatingPlaybackCoordinatorSeekCommand */
// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorSeekCommandClass) Alloc() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorSeekCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorSeekCommandClass) New() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorSeekCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorSeekCommand) Init() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorSeekCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorSeekCommand) Autorelease() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorSeekCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorSeekCommand creates a new DelegatingPlaybackCoordinatorSeekCommand instance.
func NewDelegatingPlaybackCoordinatorSeekCommand() DelegatingPlaybackCoordinatorSeekCommand {
	return getDelegatingPlaybackCoordinatorSeekCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DelegatingPlaybackCoordinatorSeekCommand */
// A command that indicates to seek to a new time in the item timeline.


// A command that indicates to seek to a new time in the item timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand
type DelegatingPlaybackCoordinatorSeekCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

// DelegatingPlaybackCoordinatorSeekCommandFrom constructs a [DelegatingPlaybackCoordinatorSeekCommand] from an unsafe.Pointer.
//
// A command that indicates to seek to a new time in the item timeline.
func DelegatingPlaybackCoordinatorSeekCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorSeekCommand {
	return DelegatingPlaybackCoordinatorSeekCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DelegatingPlaybackCoordinatorSeekCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DelegatingPlaybackCoordinatorSeekCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DelegatingPlaybackCoordinatorSeekCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DelegatingPlaybackCoordinatorSeekCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DelegatingPlaybackCoordinatorSeekCommand */

// The rate at which the coordinator expects playback to resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}/* debug [instance_properties/getter]: anticipatedPlaybackRate */


// The deadline by which the coordinator expects the delegate to handle the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/completionDueDate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) CompletionDueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("completionDueDate"))
	return rv
}/* debug [instance_properties/getter]: completionDueDate */


// The time to seek to in the item timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/itemTime
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ItemTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](d_.ID, objc.Sel("itemTime"))
	return rv
}/* debug [instance_properties/getter]: itemTime */


// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/shouldBufferInAnticipationOfPlayback
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}/* debug [instance_properties/getter]: shouldBufferInAnticipationOfPlayback */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDelegatingPlaybackCoordinatorSeekCommand */



