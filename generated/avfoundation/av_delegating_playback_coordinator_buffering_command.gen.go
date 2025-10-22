// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AnticipatedPlaybackRate() float32
	SetAnticipatedPlaybackRate(value float32)
	CompletionDueDate() foundation.Date
	SetCompletionDueDate(value foundation.IDate)
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

// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorBufferingCommandClass) Alloc() DelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorBufferingCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/anticipatedplaybackrate

func (d_ DelegatingPlaybackCoordinatorBufferingCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}


// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/anticipatedplaybackrate

func (d_ DelegatingPlaybackCoordinatorBufferingCommand) SetAnticipatedPlaybackRate(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAnticipatedPlaybackRate:"), value)
}


// The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/completionduedate

func (d_ DelegatingPlaybackCoordinatorBufferingCommand) CompletionDueDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("completionDueDate"))
	return rv
}


// The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdelegatingplaybackcoordinatorbufferingcommand/completionduedate

func (d_ DelegatingPlaybackCoordinatorBufferingCommand) SetCompletionDueDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompletionDueDate:"), value)
}



