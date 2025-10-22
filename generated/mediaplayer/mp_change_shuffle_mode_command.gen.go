// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangeShuffleModeCommand] class.
var (
	ChangeShuffleModeCommandClass     _ChangeShuffleModeCommandClass
	ChangeShuffleModeCommandClassOnce sync.Once
)

func getChangeShuffleModeCommandClass() _ChangeShuffleModeCommandClass {
	ChangeShuffleModeCommandClassOnce.Do(func() {
		ChangeShuffleModeCommandClass = _ChangeShuffleModeCommandClass{objc.GetClass("MPChangeShuffleModeCommand")}
	})
	return ChangeShuffleModeCommandClass
}

type _ChangeShuffleModeCommandClass struct {
	class objc.Class
}

// An interface definition for the [ChangeShuffleModeCommand] class.
type IChangeShuffleModeCommand interface {
	IRemoteCommand
	CurrentShuffleType() ShuffleType
	SetCurrentShuffleType(value ShuffleType)
}

// An object that responds to requests to change the current shuffle mode used during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand
type ChangeShuffleModeCommand struct {
	RemoteCommand
}

// ChangeShuffleModeCommandFrom constructs a [ChangeShuffleModeCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current shuffle mode used during playback.
func ChangeShuffleModeCommandFrom(ptr unsafe.Pointer) ChangeShuffleModeCommand {
	return ChangeShuffleModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangeShuffleModeCommandClass) Alloc() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangeShuffleModeCommandClass) New() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeShuffleModeCommand) Init() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeShuffleModeCommand) Autorelease() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeShuffleModeCommand creates a new ChangeShuffleModeCommand instance.
func NewChangeShuffleModeCommand() ChangeShuffleModeCommand {
	return getChangeShuffleModeCommandClass().New()
}


// The current shuffle mode for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand/currentShuffleType
func (c_ ChangeShuffleModeCommand) CurrentShuffleType() ShuffleType {
	rv := objc.Send[ShuffleType](c_.ID, objc.Sel("currentShuffleType"))
	return rv
}


// SetCurrentShuffleType sets the value of the currentShuffleType property.
// The current shuffle mode for a media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand/currentShuffleType
func (c_ ChangeShuffleModeCommand) SetCurrentShuffleType(value ShuffleType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentShuffleType:"), value)
}



