// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAudioMix] class.
var aVAudioMixClass = _AVAudioMixClass{objc.GetClass("AVAudioMix")}

type _AVAudioMixClass struct {
	class objc.Class
}

// An object that manages the input parameters for mixing audio tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMix

type AVAudioMix struct {
	objectivec.Object
}

// AVAudioMixFrom constructs a [AVAudioMix] from an unsafe.Pointer.
//
// An object that manages the input parameters for mixing audio tracks.
func AVAudioMixFrom(ptr unsafe.Pointer) AVAudioMix {
	return AVAudioMix{objectivec.Object{objc.ID(ptr)}}
}



