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
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedSampleBufferDataClass) Alloc() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






























