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
	

	// properties:


	

	// methods:
	MakeBatch() ISampleBufferGeneratorBatch
	CreateSampleBufferForRequestError(request IAVSampleBufferRequest, outError objectivec.IObject) SampleBufferRef /* not a class type */
	CreateSampleBufferForRequestAddingToBatchError(request IAVSampleBufferRequest, batch IAVSampleBufferGeneratorBatch, outError objectivec.IObject) SampleBufferRef /* not a class type */


}





// Alloc allocates a new instance without initialization.
func (sc _SampleBufferGeneratorClass) Alloc() SampleBufferGenerator {
	rv := objc.Send[SampleBufferGenerator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a new sample buffer generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/init(asset:timebase:)
func NewSampleBufferGeneratorWithAssetTimebase(asset IAVAsset, timebase TimebaseRef /* not a class type */) SampleBufferGenerator {
	instance := getSampleBufferGeneratorClass().Alloc()
	rv := objc.Send[SampleBufferGenerator](instance.ID, objc.Sel("initWithAsset:timebase:"), asset, timebase)
	rv.Autorelease()
	return rv
}







// Notifies the sample buffer generator when data is ready for the sample buffer reference or an error has occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/notifyOfDataReady(for:completionHandler:)
func (sc _SampleBufferGeneratorClass) NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf SampleBufferRef /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("notifyOfDataReadyForSampleBuffer:completionHandler:"), sbuf, completionHandler)
}












// Creates a batch object to handle generating multiple sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeBatch()
func (s_ SampleBufferGenerator) MakeBatch() ISampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](s_.ID, objc.Sel("makeBatch"))
	return rv
}


// Creates a sample buffer, and attempts to load its data asynchronously if requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeSampleBuffer(for:)
func (s_ SampleBufferGenerator) CreateSampleBufferForRequestError(request IAVSampleBufferRequest, outError objectivec.IObject) SampleBufferRef /* not a class type */ {
	rv := objc.Send[SampleBufferRef](s_.ID, objc.Sel("createSampleBufferForRequest:error:"), request, outError)
	return rv
}


// Creates a sample buffer and attempts to defer I/O for its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeSampleBuffer(for:addTo:)
func (s_ SampleBufferGenerator) CreateSampleBufferForRequestAddingToBatchError(request IAVSampleBufferRequest, batch IAVSampleBufferGeneratorBatch, outError objectivec.IObject) SampleBufferRef /* not a class type */ {
	rv := objc.Send[SampleBufferRef](s_.ID, objc.Sel("createSampleBufferForRequest:addingToBatch:error:"), request, batch, outError)
	return rv
}












