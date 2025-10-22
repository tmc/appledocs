// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SampleBufferGenerator] class.
var (
	SampleBufferGeneratorClass     _SampleBufferGeneratorClass
	SampleBufferGeneratorClassOnce sync.Once
)

func getSampleBufferGeneratorClass() _SampleBufferGeneratorClass {
	SampleBufferGeneratorClassOnce.Do(func() {
		SampleBufferGeneratorClass = _SampleBufferGeneratorClass{objc.GetClass("AVSampleBufferGenerator")}
	})
	return SampleBufferGeneratorClass
}

type _SampleBufferGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferGenerator] class.
type ISampleBufferGenerator interface {
	objectivec.IObject
}

// An object that creates sample buffers.
//
// Each request for creation is described in an object. The opaque objects are returned synchronously. If requested, sample data may be loaded asynchronously (depending on file format support).


// An object that creates sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator

type SampleBufferGenerator struct {
	objectivec.Object
}

// SampleBufferGeneratorFrom constructs a [SampleBufferGenerator] from an unsafe.Pointer.
//
// An object that creates sample buffers.
func SampleBufferGeneratorFrom(ptr unsafe.Pointer) SampleBufferGenerator {
	return SampleBufferGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferGeneratorClass) Alloc() SampleBufferGenerator {
	rv := objc.Send[SampleBufferGenerator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferGeneratorClass) New() SampleBufferGenerator {
	rv := objc.Send[SampleBufferGenerator](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferGenerator) Init() SampleBufferGenerator {
	rv := objc.Send[SampleBufferGenerator](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferGenerator) Autorelease() SampleBufferGenerator {
	rv := objc.Send[SampleBufferGenerator](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferGenerator creates a new SampleBufferGenerator instance.
func NewSampleBufferGenerator() SampleBufferGenerator {
	return getSampleBufferGeneratorClass().New()
}




