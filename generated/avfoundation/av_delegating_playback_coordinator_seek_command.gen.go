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




