// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerLooper] class.
var aVPlayerLooperClass = _AVPlayerLooperClass{objc.GetClass("AVPlayerLooper")}

type _AVPlayerLooperClass struct {
	class objc.Class
}

// An object that loops media content using a queue player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper

type AVPlayerLooper struct {
	objectivec.Object
}

// AVPlayerLooperFrom constructs a [AVPlayerLooper] from an unsafe.Pointer.
//
// An object that loops media content using a queue player.
func AVPlayerLooperFrom(ptr unsafe.Pointer) AVPlayerLooper {
	return AVPlayerLooper{objectivec.Object{objc.ID(ptr)}}
}



