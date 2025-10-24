// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SkipIntervalCommand] class.
var (
	SkipIntervalCommandClass     _SkipIntervalCommandClass
	SkipIntervalCommandClassOnce sync.Once
)

func getSkipIntervalCommandClass() _SkipIntervalCommandClass {
	SkipIntervalCommandClassOnce.Do(func() {
		SkipIntervalCommandClass = _SkipIntervalCommandClass{objc.GetClass("MPSkipIntervalCommand")}
	})
	return SkipIntervalCommandClass
}

type _SkipIntervalCommandClass struct {
	class objc.Class
}

// An interface definition for the [SkipIntervalCommand] class.
type ISkipIntervalCommand interface {
	IRemoteCommand
	// properties:
	PreferredIntervals() objc.IObject /* cross-framework: NSNumber */
	SetPreferredIntervals(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// An object that defines the skip intervals for the player.
//
// You use a skip interval to move the playback of a media item, forward or backward, the indicated number of seconds. For example, a forward skip interval of 30 seconds at 2 minutes and 30 seconds into a song would immediately jump to 3 minutes into the song and continue playing. The skipped content isn’t played.


// An object that defines the skip intervals for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommand
type SkipIntervalCommand struct {
	RemoteCommand
}

// SkipIntervalCommandFrom constructs a [SkipIntervalCommand] from an unsafe.Pointer.
//
// An object that defines the skip intervals for the player.
func SkipIntervalCommandFrom(ptr unsafe.Pointer) SkipIntervalCommand {
	return SkipIntervalCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SkipIntervalCommandClass) Alloc() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SkipIntervalCommandClass) New() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SkipIntervalCommand) Init() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SkipIntervalCommand) Autorelease() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkipIntervalCommand creates a new SkipIntervalCommand instance.
func NewSkipIntervalCommand() SkipIntervalCommand {
	return getSkipIntervalCommandClass().New()
}



// The available skip intervals, in seconds, for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommand/preferredintervals
func (s_ SkipIntervalCommand) PreferredIntervals() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("preferredIntervals"))
	return rv
}


// The available skip intervals, in seconds, for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpskipintervalcommand/preferredintervals
func (s_ SkipIntervalCommand) SetPreferredIntervals(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredIntervals:"), value)
}



