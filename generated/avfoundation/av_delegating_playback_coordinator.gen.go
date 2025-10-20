// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DelegatingPlaybackCoordinator] class.
var (
	DelegatingPlaybackCoordinatorClass     _DelegatingPlaybackCoordinatorClass
	DelegatingPlaybackCoordinatorClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorClass() _DelegatingPlaybackCoordinatorClass {
	DelegatingPlaybackCoordinatorClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorClass = _DelegatingPlaybackCoordinatorClass{objc.GetClass("AVDelegatingPlaybackCoordinator")}
	})
	return DelegatingPlaybackCoordinatorClass
}

type _DelegatingPlaybackCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [DelegatingPlaybackCoordinator] class.
type IDelegatingPlaybackCoordinator interface {
	IPlaybackCoordinator
}

// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group.
//
// This object coordinates the state of custom player objects, such as those that render media using and , or that play audio using . Adopt the protocol so that your app responds to playback commands from the coordinator. The commands provide the details of a requested state change so you can control your player object accordingly.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator
type DelegatingPlaybackCoordinator struct {
	PlaybackCoordinator
}

// DelegatingPlaybackCoordinatorFrom constructs a [DelegatingPlaybackCoordinator] from an unsafe.Pointer.
//
// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group.
func DelegatingPlaybackCoordinatorFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinator {
	return DelegatingPlaybackCoordinator{
		PlaybackCoordinator: PlaybackCoordinatorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorClass) Alloc() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DelegatingPlaybackCoordinatorClass) New() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinator) Init() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinator) Autorelease() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinator creates a new DelegatingPlaybackCoordinator instance.
func NewDelegatingPlaybackCoordinator() DelegatingPlaybackCoordinator {
	return getDelegatingPlaybackCoordinatorClass().New()
}




