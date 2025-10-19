// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItemTrack] class.
var (
	aVPlayerItemTrackClass     _AVPlayerItemTrackClass
	aVPlayerItemTrackClassOnce sync.Once
)

func getAVPlayerItemTrackClass() _AVPlayerItemTrackClass {
	aVPlayerItemTrackClassOnce.Do(func() {
		aVPlayerItemTrackClass = _AVPlayerItemTrackClass{objc.GetClass("AVPlayerItemTrack")}
	})
	return aVPlayerItemTrackClass
}

type _AVPlayerItemTrackClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerItemTrack] class.
type IAVPlayerItemTrack interface {
	objectivec.IObject
}

// An object that represents the presentation state of an asset track during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack
type AVPlayerItemTrack struct {
	objectivec.Object
}

// AVPlayerItemTrackFrom constructs a [AVPlayerItemTrack] from an unsafe.Pointer.
//
// An object that represents the presentation state of an asset track during playback.
func AVPlayerItemTrackFrom(ptr unsafe.Pointer) AVPlayerItemTrack {
	return AVPlayerItemTrack{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerItemTrackClass) Alloc() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerItemTrackClass) New() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerItemTrack) Init() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerItemTrack) Autorelease() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemTrack creates a new AVPlayerItemTrack instance.
func NewAVPlayerItemTrack() AVPlayerItemTrack {
	return getAVPlayerItemTrackClass().New()
}




