// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureSpatialAudioMetadataSampleGenerator] class.
var aVCaptureSpatialAudioMetadataSampleGeneratorClass = _AVCaptureSpatialAudioMetadataSampleGeneratorClass{objc.GetClass("AVCaptureSpatialAudioMetadataSampleGenerator")}

type _AVCaptureSpatialAudioMetadataSampleGeneratorClass struct {
	class objc.Class
}

// An interface for generating a spatial audio timed metadata sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator

type AVCaptureSpatialAudioMetadataSampleGenerator struct {
	objectivec.Object
}

// AVCaptureSpatialAudioMetadataSampleGeneratorFrom constructs a [AVCaptureSpatialAudioMetadataSampleGenerator] from an unsafe.Pointer.
//
// An interface for generating a spatial audio timed metadata sample.
func AVCaptureSpatialAudioMetadataSampleGeneratorFrom(ptr unsafe.Pointer) AVCaptureSpatialAudioMetadataSampleGenerator {
	return AVCaptureSpatialAudioMetadataSampleGenerator{objectivec.Object{objc.ID(ptr)}}
}



