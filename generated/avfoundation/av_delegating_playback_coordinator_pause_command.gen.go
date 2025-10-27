// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [DelegatingPlaybackCoordinatorPauseCommand] class.
type IDelegatingPlaybackCoordinatorPauseCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	

	// properties:
	AnticipatedPlaybackRate() float32
	ShouldBufferInAnticipationOfPlayback() bool


	

	// methods:


}





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

























// The rate at which the coordinator expects the current item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/anticipatedPlaybackRate
func (d_ DelegatingPlaybackCoordinatorPauseCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}


// A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/shouldBufferInAnticipationOfPlayback
func (d_ DelegatingPlaybackCoordinatorPauseCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}








