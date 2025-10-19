// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerItemRenderedLegibleOutput] class.
var aVPlayerItemRenderedLegibleOutputClass = _AVPlayerItemRenderedLegibleOutputClass{objc.GetClass("AVPlayerItemRenderedLegibleOutput")}

type _AVPlayerItemRenderedLegibleOutputClass struct {
	class objc.Class
}

// A player item output that vends media with a legible characteristic as rendered pixel buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput

type AVPlayerItemRenderedLegibleOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemRenderedLegibleOutputFrom constructs a [AVPlayerItemRenderedLegibleOutput] from an unsafe.Pointer.
//
// A player item output that vends media with a legible characteristic as rendered pixel buffers.
func AVPlayerItemRenderedLegibleOutputFrom(ptr unsafe.Pointer) AVPlayerItemRenderedLegibleOutput {
	return AVPlayerItemRenderedLegibleOutput{
		AVPlayerItemOutput: AVPlayerItemOutputFrom(ptr),
	}
}



