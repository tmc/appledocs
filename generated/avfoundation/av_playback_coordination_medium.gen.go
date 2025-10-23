// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlaybackCoordinationMedium] class.
var (
	PlaybackCoordinationMediumClass     _PlaybackCoordinationMediumClass
	PlaybackCoordinationMediumClassOnce sync.Once
)

func getPlaybackCoordinationMediumClass() _PlaybackCoordinationMediumClass {
	PlaybackCoordinationMediumClassOnce.Do(func() {
		PlaybackCoordinationMediumClass = _PlaybackCoordinationMediumClass{objc.GetClass("AVPlaybackCoordinationMedium")}
	})
	return PlaybackCoordinationMediumClass
}

type _PlaybackCoordinationMediumClass struct {
	class objc.Class
}

// An interface definition for the [PlaybackCoordinationMedium] class.
type IPlaybackCoordinationMedium interface {
	objectivec.IObject
	// properties:
	ConnectedPlaybackCoordinators() IAVPlayerPlaybackCoordinator
	SetConnectedPlaybackCoordinators(value IAVPlayerPlaybackCoordinator)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium
type PlaybackCoordinationMedium struct {
	objectivec.Object
}

// PlaybackCoordinationMediumFrom constructs a [PlaybackCoordinationMedium] from an unsafe.Pointer.
func PlaybackCoordinationMediumFrom(ptr unsafe.Pointer) PlaybackCoordinationMedium {
	return PlaybackCoordinationMedium{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlaybackCoordinationMediumClass) Alloc() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlaybackCoordinationMediumClass) New() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlaybackCoordinationMedium) Init() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlaybackCoordinationMedium) Autorelease() PlaybackCoordinationMedium {
	rv := objc.Send[PlaybackCoordinationMedium](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlaybackCoordinationMedium creates a new PlaybackCoordinationMedium instance.
func NewPlaybackCoordinationMedium() PlaybackCoordinationMedium {
	return getPlaybackCoordinationMediumClass().New()
}



// All playback coordinators that are connected to the coordination medium.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinationmedium/connectedplaybackcoordinators
func (p_ PlaybackCoordinationMedium) ConnectedPlaybackCoordinators() IAVPlayerPlaybackCoordinator {
	rv := objc.Send[PlayerPlaybackCoordinator](p_.ID, objc.Sel("connectedPlaybackCoordinators"))
	return rv
}


// All playback coordinators that are connected to the coordination medium.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinationmedium/connectedplaybackcoordinators
func (p_ PlaybackCoordinationMedium) SetConnectedPlaybackCoordinators(value IAVPlayerPlaybackCoordinator) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConnectedPlaybackCoordinators:"), value)
}



