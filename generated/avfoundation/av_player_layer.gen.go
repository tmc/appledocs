// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerLayer] class.
var aVPlayerLayerClass = _AVPlayerLayerClass{objc.GetClass("AVPlayerLayer")}

type _AVPlayerLayerClass struct {
	class objc.Class
}

// An object that presents the visual contents of a player object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer

type AVPlayerLayer struct {
	Layer
}

// AVPlayerLayerFrom constructs a [AVPlayerLayer] from an unsafe.Pointer.
//
// An object that presents the visual contents of a player object.
func AVPlayerLayerFrom(ptr unsafe.Pointer) AVPlayerLayer {
	return AVPlayerLayer{
		Layer: LayerFrom(ptr),
	}
}



