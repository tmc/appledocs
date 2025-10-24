// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DelegatingPlaybackCoordinatorPlayCommand] class.
var (
	DelegatingPlaybackCoordinatorPlayCommandClass     _DelegatingPlaybackCoordinatorPlayCommandClass
	DelegatingPlaybackCoordinatorPlayCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorPlayCommandClass() _DelegatingPlaybackCoordinatorPlayCommandClass {
	DelegatingPlaybackCoordinatorPlayCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorPlayCommandClass = _DelegatingPlaybackCoordinatorPlayCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPlayCommand")}
	})
	return DelegatingPlaybackCoordinatorPlayCommandClass
}

type _DelegatingPlaybackCoordinatorPlayCommandClass struct {
	class objc.Class
}





// An interface definition for the [DelegatingPlaybackCoordinatorPlayCommand] class.
type IDelegatingPlaybackCoordinatorPlayCommand interface {
	IDelegatingPlaybackCoordinatorPlaybackControlCommand
	

	// properties:
	HostClockTime() objc.IObject /* cross-framework: Time */
	ItemTime() objc.IObject /* cross-framework: Time */
	Rate() float32


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorPlayCommandClass) Alloc() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlayCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorPlayCommandClass) New() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlayCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorPlayCommand) Init() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlayCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorPlayCommand) Autorelease() DelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlayCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorPlayCommand creates a new DelegatingPlaybackCoordinatorPlayCommand instance.
func NewDelegatingPlaybackCoordinatorPlayCommand() DelegatingPlaybackCoordinatorPlayCommand {
	return getDelegatingPlaybackCoordinatorPlayCommandClass().New()
}





// A command that indicates to play at a specific rate and time.


// A command that indicates to play at a specific rate and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand
type DelegatingPlaybackCoordinatorPlayCommand struct {
	DelegatingPlaybackCoordinatorPlaybackControlCommand
}

// DelegatingPlaybackCoordinatorPlayCommandFrom constructs a [DelegatingPlaybackCoordinatorPlayCommand] from an unsafe.Pointer.
//
// A command that indicates to play at a specific rate and time.
func DelegatingPlaybackCoordinatorPlayCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPlayCommand {
	return DelegatingPlaybackCoordinatorPlayCommand{
		DelegatingPlaybackCoordinatorPlaybackControlCommand: DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr),
	}
}

























// A host clock time to use to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/hostClockTime
func (d_ DelegatingPlaybackCoordinatorPlayCommand) HostClockTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](d_.ID, objc.Sel("hostClockTime"))
	return rv
}


// A time in the item timeline to use to begin playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/itemTime
func (d_ DelegatingPlaybackCoordinatorPlayCommand) ItemTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](d_.ID, objc.Sel("itemTime"))
	return rv
}


// A rate to use when starting playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/rate
func (d_ DelegatingPlaybackCoordinatorPlayCommand) Rate() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("rate"))
	return rv
}








