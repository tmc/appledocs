// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVQueuePlayer] class.
var aVQueuePlayerClass = _AVQueuePlayerClass{objc.GetClass("AVQueuePlayer")}

type _AVQueuePlayerClass struct {
	class objc.Class
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



