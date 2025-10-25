// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVDelegatingPlaybackCoordinatorBufferingCommand */


/* debug [class_header]: Header for AVDelegatingPlaybackCoordinatorBufferingCommand */
// The class instance for the [DelegatingPlaybackCoordinatorBufferingCommand] class.
var (
	DelegatingPlaybackCoordinatorBufferingCommandClass     _DelegatingPlaybackCoordinatorBufferingCommandClass
	DelegatingPlaybackCoordinatorBufferingCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorBufferingCommandClass() _DelegatingPlaybackCoordinatorBufferingCommandClass {
	DelegatingPlaybackCoordinatorBufferingCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorBufferingCommandClass = _DelegatingPlaybackCoordinatorBufferingCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorBufferingCommand")}
	})
	return DelegatingPlaybackCoordinatorBufferingCommandClass
}

type _DelegatingPlaybackCoordinatorBufferingCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DelegatingPlaybackCoordinatorBufferingCommand */
// An interface definition for the [DelegatingPlaybackCoordinatorBufferingCommand] class.
type IDelegatingPlaybackCoordinatorBufferingCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	
/* debug [class_interface_properties]: Properties for DelegatingPlaybackCoordinatorBufferingCommand */
	// properties:
	AnticipatedPlaybackRate() float32
	CompletionDueDate() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DelegatingPlaybackCoordinatorBufferingCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DelegatingPlaybackCoordinatorBufferingCommand */
// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorBufferingCommandClass) Alloc() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorBufferingCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorBufferingCommandClass) New() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorBufferingCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) Init() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorBufferingCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) Autorelease() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorBufferingCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorBufferingCommand creates a new DelegatingPlaybackCoordinatorBufferingCommand instance.
func NewDelegatingPlaybackCoordinatorBufferingCommand() DelegatingPlaybackCoordinatorBufferingCommand {
	return getDelegatingPlaybackCoordinatorBufferingCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DelegatingPlaybackCoordinatorBufferingCommand */
// A command that indicates to start buffering data in preparation for playback.
//
// When your app receives this command, update its user interface to indicate that playback is buffering.


// A command that indicates to start buffering data in preparation for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand
type DelegatingPlaybackCoordinatorBufferingCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

// DelegatingPlaybackCoordinatorBufferingCommandFrom constructs a [DelegatingPlaybackCoordinatorBufferingCommand] from an unsafe.Pointer.
//
// A command that indicates to start buffering data in preparation for playback.
func DelegatingPlaybackCoordinatorBufferingCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorBufferingCommand {
	return DelegatingPlaybackCoordinatorBufferingCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DelegatingPlaybackCoordinatorBufferingCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DelegatingPlaybackCoordinatorBufferingCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DelegatingPlaybackCoordinatorBufferingCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DelegatingPlaybackCoordinatorBufferingCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DelegatingPlaybackCoordinatorBufferingCommand */

// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}/* debug [instance_properties/getter]: anticipatedPlaybackRate */


// The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/completionDueDate
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) CompletionDueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("completionDueDate"))
	return rv
}/* debug [instance_properties/getter]: completionDueDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDelegatingPlaybackCoordinatorBufferingCommand */



