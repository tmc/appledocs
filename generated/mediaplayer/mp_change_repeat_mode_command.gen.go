// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangeRepeatModeCommand] class.
var (
	ChangeRepeatModeCommandClass     _ChangeRepeatModeCommandClass
	ChangeRepeatModeCommandClassOnce sync.Once
)

func getChangeRepeatModeCommandClass() _ChangeRepeatModeCommandClass {
	ChangeRepeatModeCommandClassOnce.Do(func() {
		ChangeRepeatModeCommandClass = _ChangeRepeatModeCommandClass{objc.GetClass("MPChangeRepeatModeCommand")}
	})
	return ChangeRepeatModeCommandClass
}

type _ChangeRepeatModeCommandClass struct {
	class objc.Class
}

// An interface definition for the [ChangeRepeatModeCommand] class.
type IChangeRepeatModeCommand interface {
	IRemoteCommand
}

// An object that responds to requests to change the current repeat mode used during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommand
type ChangeRepeatModeCommand struct {
	RemoteCommand
}

// ChangeRepeatModeCommandFrom constructs a [ChangeRepeatModeCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current repeat mode used during playback.
func ChangeRepeatModeCommandFrom(ptr unsafe.Pointer) ChangeRepeatModeCommand {
	return ChangeRepeatModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangeRepeatModeCommandClass) Alloc() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangeRepeatModeCommandClass) New() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeRepeatModeCommand) Init() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeRepeatModeCommand) Autorelease() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeRepeatModeCommand creates a new ChangeRepeatModeCommand instance.
func NewChangeRepeatModeCommand() ChangeRepeatModeCommand {
	return getChangeRepeatModeCommandClass().New()
}


// The current repeat option for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommand/currentrepeattype
func (c_ ChangeRepeatModeCommand) CurrentRepeatType() RepeatType {
	rv := objc.Send[RepeatType](c_.ID, objc.Sel("currentRepeatType"))
	return rv
}


// SetCurrentRepeatType sets the value of the currentRepeatType property.
// The current repeat option for a media item.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommand/currentrepeattype
func (c_ ChangeRepeatModeCommand) SetCurrentRepeatType(value RepeatType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentRepeatType:"), value)
}



