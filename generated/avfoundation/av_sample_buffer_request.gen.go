// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SampleBufferRequest] class.
var (
	SampleBufferRequestClass     _SampleBufferRequestClass
	SampleBufferRequestClassOnce sync.Once
)

func getSampleBufferRequestClass() _SampleBufferRequestClass {
	SampleBufferRequestClassOnce.Do(func() {
		SampleBufferRequestClass = _SampleBufferRequestClass{objc.GetClass("AVSampleBufferRequest")}
	})
	return SampleBufferRequestClass
}

type _SampleBufferRequestClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferRequest] class.
type ISampleBufferRequest interface {
	objectivec.IObject
}

// An object that describes a sample buffer creation request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest
type SampleBufferRequest struct {
	objectivec.Object
}

// SampleBufferRequestFrom constructs a [SampleBufferRequest] from an unsafe.Pointer.
//
// An object that describes a sample buffer creation request.
func SampleBufferRequestFrom(ptr unsafe.Pointer) SampleBufferRequest {
	return SampleBufferRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferRequestClass) Alloc() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferRequestClass) New() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferRequest) Init() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferRequest) Autorelease() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferRequest creates a new SampleBufferRequest instance.
func NewSampleBufferRequest() SampleBufferRequest {
	return getSampleBufferRequestClass().New()
}


// The maximum number of samples to load.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/maxSampleCount
func (s_ SampleBufferRequest) MaxSampleCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maxSampleCount"))
	return rv
}


// SetMaxSampleCount sets the value of the maxSampleCount property.
// The maximum number of samples to load.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/maxSampleCount
func (s_ SampleBufferRequest) SetMaxSampleCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxSampleCount:"), value)
}



