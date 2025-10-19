// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerPlaybackCoordinator] class.
var aVPlayerPlaybackCoordinatorClass = _AVPlayerPlaybackCoordinatorClass{objc.GetClass("AVPlayerPlaybackCoordinator")}

type _AVPlayerPlaybackCoordinatorClass struct {
	class objc.Class
}

// A playback coordinator subclass that coordinates the playback of player objects in a connected group. [Full Topic]
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



