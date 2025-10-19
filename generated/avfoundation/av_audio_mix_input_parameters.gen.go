// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAudioMixInputParameters] class.
var aVAudioMixInputParametersClass = _AVAudioMixInputParametersClass{objc.GetClass("AVAudioMixInputParameters")}

type _AVAudioMixInputParametersClass struct {
	class objc.Class
}

// An object that represents the parameters that you apply when adding an audio track to a mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters

type AVAudioMixInputParameters struct {
	objectivec.Object
}

// AVAudioMixInputParametersFrom constructs a [AVAudioMixInputParameters] from an unsafe.Pointer.
//
// An object that represents the parameters that you apply when adding an audio track to a mix.
func AVAudioMixInputParametersFrom(ptr unsafe.Pointer) AVAudioMixInputParameters {
	return AVAudioMixInputParameters{objectivec.Object{objc.ID(ptr)}}
}



