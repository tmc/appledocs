// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureSynchronizedSampleBufferData] class.
var (
	CaptureSynchronizedSampleBufferDataClass     _CaptureSynchronizedSampleBufferDataClass
	CaptureSynchronizedSampleBufferDataClassOnce sync.Once
)

func getCaptureSynchronizedSampleBufferDataClass() _CaptureSynchronizedSampleBufferDataClass {
	CaptureSynchronizedSampleBufferDataClassOnce.Do(func() {
		CaptureSynchronizedSampleBufferDataClass = _CaptureSynchronizedSampleBufferDataClass{objc.GetClass("AVCaptureSynchronizedSampleBufferData")}
	})
	return CaptureSynchronizedSampleBufferDataClass
}

type _CaptureSynchronizedSampleBufferDataClass struct {
	class objc.Class
}

// An interface definition for the [CaptureSynchronizedSampleBufferData] class.
type ICaptureSynchronizedSampleBufferData interface {
	ICaptureSynchronizedData
	DroppedReason() unsafe.Pointer
	SetDroppedReason(value unsafe.Pointer)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	SampleBufferWasDropped() bool
	SetSampleBufferWasDropped(value bool)
}

// A container for video or audio samples collected using synchronized capture.


// A container for video or audio samples collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData
type CaptureSynchronizedSampleBufferData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedSampleBufferDataFrom constructs a [CaptureSynchronizedSampleBufferData] from an unsafe.Pointer.
//
// A container for video or audio samples collected using synchronized capture.
func CaptureSynchronizedSampleBufferDataFrom(ptr unsafe.Pointer) CaptureSynchronizedSampleBufferData {
	return CaptureSynchronizedSampleBufferData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedSampleBufferDataClass) Alloc() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureSynchronizedSampleBufferDataClass) New() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedSampleBufferData) Init() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedSampleBufferData) Autorelease() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedSampleBufferData creates a new CaptureSynchronizedSampleBufferData instance.
func NewCaptureSynchronizedSampleBufferData() CaptureSynchronizedSampleBufferData {
	return getCaptureSynchronizedSampleBufferDataClass().New()
}



// A value indicating why the capture output failed to deliver sample buffers, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/droppedreason
func (c_ CaptureSynchronizedSampleBufferData) DroppedReason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("droppedReason"))
	return rv
}


// A value indicating why the capture output failed to deliver sample buffers, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/droppedreason
func (c_ CaptureSynchronizedSampleBufferData) SetDroppedReason(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDroppedReason:"), value)
}


// The depth data captured at this synchronization point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/samplebuffer
func (c_ CaptureSynchronizedSampleBufferData) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// The depth data captured at this synchronization point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/samplebuffer
func (c_ CaptureSynchronizedSampleBufferData) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBuffer:"), value)
}


// A Boolean value indicating whether sample buffers were discarded between capture and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/samplebufferwasdropped
func (c_ CaptureSynchronizedSampleBufferData) SampleBufferWasDropped() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sampleBufferWasDropped"))
	return rv
}


// A Boolean value indicating whether sample buffers were discarded between capture and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesynchronizedsamplebufferdata/samplebufferwasdropped
func (c_ CaptureSynchronizedSampleBufferData) SetSampleBufferWasDropped(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferWasDropped:"), value)
}



