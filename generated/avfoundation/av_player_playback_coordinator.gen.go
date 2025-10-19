// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerPlaybackCoordinator] class.
var (
	aVPlayerPlaybackCoordinatorClass     _AVPlayerPlaybackCoordinatorClass
	aVPlayerPlaybackCoordinatorClassOnce sync.Once
)

func getAVPlayerPlaybackCoordinatorClass() _AVPlayerPlaybackCoordinatorClass {
	aVPlayerPlaybackCoordinatorClassOnce.Do(func() {
		aVPlayerPlaybackCoordinatorClass = _AVPlayerPlaybackCoordinatorClass{objc.GetClass("AVPlayerPlaybackCoordinator")}
	})
	return aVPlayerPlaybackCoordinatorClass
}

type _AVPlayerPlaybackCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerPlaybackCoordinator] class.
type IAVPlayerPlaybackCoordinator interface {
	IAVPlaybackCoordinator
}

// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator
type AVPlayerPlaybackCoordinator struct {
	AVPlaybackCoordinator
}

// AVPlayerPlaybackCoordinatorFrom constructs a [AVPlayerPlaybackCoordinator] from an unsafe.Pointer.
//
// A playback coordinator subclass that coordinates the playback of player objects in a connected group.
func AVPlayerPlaybackCoordinatorFrom(ptr unsafe.Pointer) AVPlayerPlaybackCoordinator {
	return AVPlayerPlaybackCoordinator{
		AVPlaybackCoordinator: AVPlaybackCoordinatorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerPlaybackCoordinatorClass) Alloc() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerPlaybackCoordinatorClass) New() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerPlaybackCoordinator) Init() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerPlaybackCoordinator) Autorelease() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerPlaybackCoordinator creates a new AVPlayerPlaybackCoordinator instance.
func NewAVPlayerPlaybackCoordinator() AVPlayerPlaybackCoordinator {
	return getAVPlayerPlaybackCoordinatorClass().New()
}




