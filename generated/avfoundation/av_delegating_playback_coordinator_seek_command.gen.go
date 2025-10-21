// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A command that indicates to seek to a new time in the item timeline.
//
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

// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorSeekCommandClass) Alloc() DelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorSeekCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The rate at which the coordinator expects playback to resume.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/anticipatedplaybackrate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) AnticipatedPlaybackRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}


// SetAnticipatedPlaybackRate sets the value of the anticipatedPlaybackRate property.
// The rate at which the coordinator expects playback to resume.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/anticipatedplaybackrate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) SetAnticipatedPlaybackRate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAnticipatedPlaybackRate:"), value)
}

// The deadline by which the coordinator expects the delegate to handle the command.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/completionduedate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) CompletionDueDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("completionDueDate"))
	return rv
}


// SetCompletionDueDate sets the value of the completionDueDate property.
// The deadline by which the coordinator expects the delegate to handle the command.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/completionduedate
func (d_ DelegatingPlaybackCoordinatorSeekCommand) SetCompletionDueDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompletionDueDate:"), value)
}

// The time to seek to in the item timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/itemtime
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ItemTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("itemTime"))
	return rv
}


// SetItemTime sets the value of the itemTime property.
// The time to seek to in the item timeline.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/itemtime
func (d_ DelegatingPlaybackCoordinatorSeekCommand) SetItemTime(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setItemTime:"), value)
}

// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/shouldbufferinanticipationofplayback
func (d_ DelegatingPlaybackCoordinatorSeekCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}


// SetShouldBufferInAnticipationOfPlayback sets the value of the shouldBufferInAnticipationOfPlayback property.
// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorseekcommand/shouldbufferinanticipationofplayback
func (d_ DelegatingPlaybackCoordinatorSeekCommand) SetShouldBufferInAnticipationOfPlayback(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShouldBufferInAnticipationOfPlayback:"), value)
}



