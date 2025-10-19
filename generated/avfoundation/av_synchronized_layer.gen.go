// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVSynchronizedLayer] class.
var aVSynchronizedLayerClass = _AVSynchronizedLayerClass{objc.GetClass("AVSynchronizedLayer")}

type _AVSynchronizedLayerClass struct {
	class objc.Class
}

// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer

type AVSynchronizedLayer struct {
	Layer
}

// AVSynchronizedLayerFrom constructs a [AVSynchronizedLayer] from an unsafe.Pointer.
//
// A Core Animation layer that derives its timing from a player item so that you can synchronize layer animations with media playback.
func AVSynchronizedLayerFrom(ptr unsafe.Pointer) AVSynchronizedLayer {
	return AVSynchronizedLayer{
		Layer: LayerFrom(ptr),
	}
}



