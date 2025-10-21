// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SampleBufferGeneratorBatch] class.
var (
	SampleBufferGeneratorBatchClass     _SampleBufferGeneratorBatchClass
	SampleBufferGeneratorBatchClassOnce sync.Once
)

func getSampleBufferGeneratorBatchClass() _SampleBufferGeneratorBatchClass {
	SampleBufferGeneratorBatchClassOnce.Do(func() {
		SampleBufferGeneratorBatchClass = _SampleBufferGeneratorBatchClass{objc.GetClass("AVSampleBufferGeneratorBatch")}
	})
	return SampleBufferGeneratorBatchClass
}

type _SampleBufferGeneratorBatchClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferGeneratorBatch] class.
type ISampleBufferGeneratorBatch interface {
	objectivec.IObject
	MakeDataReadyWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that generates sample buffers in a batch.
//
// The benefit of batching is it aggregates adjacent I/O requests and overlaps them when possible for all sample buffers within the batch.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch
type SampleBufferGeneratorBatch struct {
	objectivec.Object
}

// SampleBufferGeneratorBatchFrom constructs a [SampleBufferGeneratorBatch] from an unsafe.Pointer.
//
// An object that generates sample buffers in a batch.
func SampleBufferGeneratorBatchFrom(ptr unsafe.Pointer) SampleBufferGeneratorBatch {
	return SampleBufferGeneratorBatch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferGeneratorBatchClass) Alloc() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferGeneratorBatchClass) New() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferGeneratorBatch) Init() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferGeneratorBatch) Autorelease() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferGeneratorBatch creates a new SampleBufferGeneratorBatch instance.
func NewSampleBufferGeneratorBatch() SampleBufferGeneratorBatch {
	return getSampleBufferGeneratorBatchClass().New()
}


// Loads sample data asynchronously for all sample buffers within a batch.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch/makeDataReady(completionHandler:)
func (s_ SampleBufferGeneratorBatch) MakeDataReadyWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("makeDataReadyWithCompletionHandler:"), completionHandler)
}



