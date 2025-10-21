// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PlaybackSpeed] class.
var (
	PlaybackSpeedClass     _PlaybackSpeedClass
	PlaybackSpeedClassOnce sync.Once
)

func getPlaybackSpeedClass() _PlaybackSpeedClass {
	PlaybackSpeedClassOnce.Do(func() {
		PlaybackSpeedClass = _PlaybackSpeedClass{objc.GetClass("AVPlaybackSpeed")}
	})
	return PlaybackSpeedClass
}

type _PlaybackSpeedClass struct {
	class objc.Class
}

// An interface definition for the [PlaybackSpeed] class.
type IPlaybackSpeed interface {
	objectivec.IObject
}

// An object that represents a user-selectable playback speed in a playback user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed
type PlaybackSpeed struct {
	objectivec.Object
}

// PlaybackSpeedFrom constructs a [PlaybackSpeed] from an unsafe.Pointer.
//
// An object that represents a user-selectable playback speed in a playback user interface.
func PlaybackSpeedFrom(ptr unsafe.Pointer) PlaybackSpeed {
	return PlaybackSpeed{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlaybackSpeedClass) Alloc() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlaybackSpeedClass) New() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlaybackSpeed) Init() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlaybackSpeed) Autorelease() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlaybackSpeed creates a new PlaybackSpeed instance.
func NewPlaybackSpeed() PlaybackSpeed {
	return getPlaybackSpeedClass().New()
}




