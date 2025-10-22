// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureSpatialAudioMetadataSampleGenerator] class.
var (
	CaptureSpatialAudioMetadataSampleGeneratorClass     _CaptureSpatialAudioMetadataSampleGeneratorClass
	CaptureSpatialAudioMetadataSampleGeneratorClassOnce sync.Once
)

func getCaptureSpatialAudioMetadataSampleGeneratorClass() _CaptureSpatialAudioMetadataSampleGeneratorClass {
	CaptureSpatialAudioMetadataSampleGeneratorClassOnce.Do(func() {
		CaptureSpatialAudioMetadataSampleGeneratorClass = _CaptureSpatialAudioMetadataSampleGeneratorClass{objc.GetClass("AVCaptureSpatialAudioMetadataSampleGenerator")}
	})
	return CaptureSpatialAudioMetadataSampleGeneratorClass
}

type _CaptureSpatialAudioMetadataSampleGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [CaptureSpatialAudioMetadataSampleGenerator] class.
type ICaptureSpatialAudioMetadataSampleGenerator interface {
	objectivec.IObject
	AnalyzeAudioSample(sbuf unsafe.Pointer) unsafe.Pointer
	NewTimedMetadataSampleBufferAndResetAnalyzer() unsafe.Pointer
	TimedMetadataSampleBufferFormatDescription() unsafe.Pointer
	SetTimedMetadataSampleBufferFormatDescription(value unsafe.Pointer)
}

// An interface for generating a spatial audio timed metadata sample.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator
type CaptureSpatialAudioMetadataSampleGenerator struct {
	objectivec.Object
}

// CaptureSpatialAudioMetadataSampleGeneratorFrom constructs a [CaptureSpatialAudioMetadataSampleGenerator] from an unsafe.Pointer.
//
// An interface for generating a spatial audio timed metadata sample.
func CaptureSpatialAudioMetadataSampleGeneratorFrom(ptr unsafe.Pointer) CaptureSpatialAudioMetadataSampleGenerator {
	return CaptureSpatialAudioMetadataSampleGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureSpatialAudioMetadataSampleGeneratorClass) Alloc() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureSpatialAudioMetadataSampleGeneratorClass) New() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSpatialAudioMetadataSampleGenerator) Init() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSpatialAudioMetadataSampleGenerator) Autorelease() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSpatialAudioMetadataSampleGenerator creates a new CaptureSpatialAudioMetadataSampleGenerator instance.
func NewCaptureSpatialAudioMetadataSampleGenerator() CaptureSpatialAudioMetadataSampleGenerator {
	return getCaptureSpatialAudioMetadataSampleGeneratorClass().New()
}


// Analyzes the provided audio sample buffer for its contribution to the spatial audio timed metadata value.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/analyzeAudioSample(_:)
func (c_ CaptureSpatialAudioMetadataSampleGenerator) AnalyzeAudioSample(sbuf unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("analyzeAudioSample:"), sbuf)
	return rv
}

// Creates a sample buffer containing a spatial audio timed metadata sample computed from all analyzed audio buffers, and resets the analyzer to its initial state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/newTimedMetadataSampleBufferAndResetAnalyzer()
func (c_ CaptureSpatialAudioMetadataSampleGenerator) NewTimedMetadataSampleBufferAndResetAnalyzer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newTimedMetadataSampleBufferAndResetAnalyzer"))
	return rv
}

// Returns the format description of the sample buffer returned from the
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturespatialaudiometadatasamplegenerator/timedmetadatasamplebufferformatdescription
func (c_ CaptureSpatialAudioMetadataSampleGenerator) TimedMetadataSampleBufferFormatDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timedMetadataSampleBufferFormatDescription"))
	return rv
}


// SetTimedMetadataSampleBufferFormatDescription sets the value of the timedMetadataSampleBufferFormatDescription property.
// Returns the format description of the sample buffer returned from the

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturespatialaudiometadatasamplegenerator/timedmetadatasamplebufferformatdescription
func (c_ CaptureSpatialAudioMetadataSampleGenerator) SetTimedMetadataSampleBufferFormatDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimedMetadataSampleBufferFormatDescription:"), value)
}



