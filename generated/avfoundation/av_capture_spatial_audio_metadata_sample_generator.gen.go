// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureSpatialAudioMetadataSampleGenerator] class.
var (
	aVCaptureSpatialAudioMetadataSampleGeneratorClass     _AVCaptureSpatialAudioMetadataSampleGeneratorClass
	aVCaptureSpatialAudioMetadataSampleGeneratorClassOnce sync.Once
)

func getAVCaptureSpatialAudioMetadataSampleGeneratorClass() _AVCaptureSpatialAudioMetadataSampleGeneratorClass {
	aVCaptureSpatialAudioMetadataSampleGeneratorClassOnce.Do(func() {
		aVCaptureSpatialAudioMetadataSampleGeneratorClass = _AVCaptureSpatialAudioMetadataSampleGeneratorClass{objc.GetClass("AVCaptureSpatialAudioMetadataSampleGenerator")}
	})
	return aVCaptureSpatialAudioMetadataSampleGeneratorClass
}

type _AVCaptureSpatialAudioMetadataSampleGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureSpatialAudioMetadataSampleGenerator] class.
type IAVCaptureSpatialAudioMetadataSampleGenerator interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureSpatialAudioMetadataSampleGeneratorClass) Alloc() AVCaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[AVCaptureSpatialAudioMetadataSampleGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureSpatialAudioMetadataSampleGeneratorClass) New() AVCaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[AVCaptureSpatialAudioMetadataSampleGenerator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureSpatialAudioMetadataSampleGenerator) Init() AVCaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[AVCaptureSpatialAudioMetadataSampleGenerator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureSpatialAudioMetadataSampleGenerator) Autorelease() AVCaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[AVCaptureSpatialAudioMetadataSampleGenerator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSpatialAudioMetadataSampleGenerator creates a new AVCaptureSpatialAudioMetadataSampleGenerator instance.
func NewAVCaptureSpatialAudioMetadataSampleGenerator() AVCaptureSpatialAudioMetadataSampleGenerator {
	return getAVCaptureSpatialAudioMetadataSampleGeneratorClass().New()
}




