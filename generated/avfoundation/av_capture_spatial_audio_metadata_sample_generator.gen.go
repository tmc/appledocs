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
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSpatialAudioMetadataSampleGeneratorClass) Alloc() CaptureSpatialAudioMetadataSampleGenerator {
	rv := objc.Send[CaptureSpatialAudioMetadataSampleGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An interface for generating a spatial audio timed metadata sample.


// An interface for generating a spatial audio timed metadata sample.
//
// [Full Topic]
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






























