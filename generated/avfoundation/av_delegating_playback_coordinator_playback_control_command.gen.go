// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
var (
	DelegatingPlaybackCoordinatorPlaybackControlCommandClass     _DelegatingPlaybackCoordinatorPlaybackControlCommandClass
	DelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorPlaybackControlCommandClass() _DelegatingPlaybackCoordinatorPlaybackControlCommandClass {
	DelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorPlaybackControlCommandClass = _DelegatingPlaybackCoordinatorPlaybackControlCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPlaybackControlCommand")}
	})
	return DelegatingPlaybackCoordinatorPlaybackControlCommandClass
}

type _DelegatingPlaybackCoordinatorPlaybackControlCommandClass struct {
	class objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
type IDelegatingPlaybackCoordinatorPlaybackControlCommand interface {
	objectivec.IObject
}

// An abstract superclass for playback commands.
//
// Playback commands inherit state that identifies their originator and applicable item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand
type DelegatingPlaybackCoordinatorPlaybackControlCommand struct {
	objectivec.Object
}

// DelegatingPlaybackCoordinatorPlaybackControlCommandFrom constructs a [DelegatingPlaybackCoordinatorPlaybackControlCommand] from an unsafe.Pointer.
//
// An abstract superclass for playback commands.
func DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return DelegatingPlaybackCoordinatorPlaybackControlCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) Alloc() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) New() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Init() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Autorelease() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorPlaybackControlCommand creates a new DelegatingPlaybackCoordinatorPlaybackControlCommand instance.
func NewDelegatingPlaybackCoordinatorPlaybackControlCommand() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return getDelegatingPlaybackCoordinatorPlaybackControlCommandClass().New()
}




