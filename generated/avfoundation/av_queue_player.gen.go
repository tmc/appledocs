// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVQueuePlayer] class.
var (
	aVQueuePlayerClass     _AVQueuePlayerClass
	aVQueuePlayerClassOnce sync.Once
)

func getAVQueuePlayerClass() _AVQueuePlayerClass {
	aVQueuePlayerClassOnce.Do(func() {
		aVQueuePlayerClass = _AVQueuePlayerClass{objc.GetClass("AVQueuePlayer")}
	})
	return aVQueuePlayerClass
}

type _AVQueuePlayerClass struct {
	class objc.Class
}

// An interface definition for the [AVQueuePlayer] class.
type IAVQueuePlayer interface {
	IAVPlayer
}

// An object that plays a sequence of player items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer
type AVQueuePlayer struct {
	AVPlayer
}

// AVQueuePlayerFrom constructs a [AVQueuePlayer] from an unsafe.Pointer.
//
// An object that plays a sequence of player items.
func AVQueuePlayerFrom(ptr unsafe.Pointer) AVQueuePlayer {
	return AVQueuePlayer{
		AVPlayer: AVPlayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVQueuePlayerClass) Alloc() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVQueuePlayerClass) New() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVQueuePlayer) Init() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVQueuePlayer) Autorelease() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVQueuePlayer creates a new AVQueuePlayer instance.
func NewAVQueuePlayer() AVQueuePlayer {
	return getAVQueuePlayerClass().New()
}




