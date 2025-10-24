// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [DelegatingPlaybackCoordinatorSeekCommand] class.
type IDelegatingPlaybackCoordinatorSeekCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	

	// properties:
	AnticipatedPlaybackRate() float32
	CompletionDueDate() objc.IObject /* cross-framework: NSDate */
	ItemTime() objc.IObject /* cross-framework: Time */
	ShouldBufferInAnticipationOfPlayback() bool


	

	// methods:


}





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

























// The rate at which the coordinator expects playback to resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}


// The deadline by which the coordinator expects the delegate to handle the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/completionDueDate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) CompletionDueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("completionDueDate"))
	return rv
}


// The time to seek to in the item timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/itemTime
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ItemTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](d_.ID, objc.Sel("itemTime"))
	return rv
}


// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/shouldBufferInAnticipationOfPlayback
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}








