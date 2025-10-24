// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangePlaybackPositionCommand] class.
var (
	ChangePlaybackPositionCommandClass     _ChangePlaybackPositionCommandClass
	ChangePlaybackPositionCommandClassOnce sync.Once
)

func getChangePlaybackPositionCommandClass() _ChangePlaybackPositionCommandClass {
	ChangePlaybackPositionCommandClassOnce.Do(func() {
		ChangePlaybackPositionCommandClass = _ChangePlaybackPositionCommandClass{objc.GetClass("MPChangePlaybackPositionCommand")}
	})
	return ChangePlaybackPositionCommandClass
}

type _ChangePlaybackPositionCommandClass struct {
	class objc.Class
}

// An interface definition for the [ChangePlaybackPositionCommand] class.
type IChangePlaybackPositionCommand interface {
	IRemoteCommand
	// properties:
	// methods:
}

// An object that responds to requests to change the current playback position of the playing item.


// An object that responds to requests to change the current playback position of the playing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommand
type ChangePlaybackPositionCommand struct {
	RemoteCommand
}

// ChangePlaybackPositionCommandFrom constructs a [ChangePlaybackPositionCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current playback position of the playing item.
func ChangePlaybackPositionCommandFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommand {
	return ChangePlaybackPositionCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackPositionCommandClass) Alloc() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangePlaybackPositionCommandClass) New() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackPositionCommand) Init() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackPositionCommand) Autorelease() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackPositionCommand creates a new ChangePlaybackPositionCommand instance.
func NewChangePlaybackPositionCommand() ChangePlaybackPositionCommand {
	return getChangePlaybackPositionCommandClass().New()
}




