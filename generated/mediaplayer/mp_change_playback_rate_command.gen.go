// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangePlaybackRateCommand] class.
var (
	ChangePlaybackRateCommandClass     _ChangePlaybackRateCommandClass
	ChangePlaybackRateCommandClassOnce sync.Once
)

func getChangePlaybackRateCommandClass() _ChangePlaybackRateCommandClass {
	ChangePlaybackRateCommandClassOnce.Do(func() {
		ChangePlaybackRateCommandClass = _ChangePlaybackRateCommandClass{objc.GetClass("MPChangePlaybackRateCommand")}
	})
	return ChangePlaybackRateCommandClass
}

type _ChangePlaybackRateCommandClass struct {
	class objc.Class
}

// An interface definition for the [ChangePlaybackRateCommand] class.
type IChangePlaybackRateCommand interface {
	IRemoteCommand
}

// An object that responds to requests to change the playback rate of the playing item.
//
// Apps can change the current playback rate of a media item to one of the supported rates defined by the property.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand
type ChangePlaybackRateCommand struct {
	RemoteCommand
}

// ChangePlaybackRateCommandFrom constructs a [ChangePlaybackRateCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the playback rate of the playing item.
func ChangePlaybackRateCommandFrom(ptr unsafe.Pointer) ChangePlaybackRateCommand {
	return ChangePlaybackRateCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackRateCommandClass) Alloc() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangePlaybackRateCommandClass) New() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackRateCommand) Init() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackRateCommand) Autorelease() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackRateCommand creates a new ChangePlaybackRateCommand instance.
func NewChangePlaybackRateCommand() ChangePlaybackRateCommand {
	return getChangePlaybackRateCommandClass().New()
}


// The supported playback rates for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand/supportedPlaybackRates
func (c_ ChangePlaybackRateCommand) SupportedPlaybackRates() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](c_.ID, objc.Sel("supportedPlaybackRates"))
	return rv
}


// SetSupportedPlaybackRates sets the value of the supportedPlaybackRates property.
// The supported playback rates for a media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand/supportedPlaybackRates
func (c_ ChangePlaybackRateCommand) SetSupportedPlaybackRates(value []unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedPlaybackRates:"), value)
}


