// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [DelegatingPlaybackCoordinatorBufferingCommand] class.
type IDelegatingPlaybackCoordinatorBufferingCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	

	// properties:
	AnticipatedPlaybackRate() float32
	CompletionDueDate() objc.IObject /* cross-framework: NSDate */


	

	// methods:


}





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

























// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}


// The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/completionDueDate
func (d_ DelegatingPlaybackCoordinatorBufferingCommand) CompletionDueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("completionDueDate"))
	return rv
}








