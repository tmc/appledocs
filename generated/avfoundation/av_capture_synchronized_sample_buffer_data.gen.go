// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureSynchronizedSampleBufferData] class.
var (
	aVCaptureSynchronizedSampleBufferDataClass     _AVCaptureSynchronizedSampleBufferDataClass
	aVCaptureSynchronizedSampleBufferDataClassOnce sync.Once
)

func getAVCaptureSynchronizedSampleBufferDataClass() _AVCaptureSynchronizedSampleBufferDataClass {
	aVCaptureSynchronizedSampleBufferDataClassOnce.Do(func() {
		aVCaptureSynchronizedSampleBufferDataClass = _AVCaptureSynchronizedSampleBufferDataClass{objc.GetClass("AVCaptureSynchronizedSampleBufferData")}
	})
	return aVCaptureSynchronizedSampleBufferDataClass
}

type _AVCaptureSynchronizedSampleBufferDataClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureSynchronizedSampleBufferData] class.
type IAVCaptureSynchronizedSampleBufferData interface {
	IAVCaptureSynchronizedData
}

// A container for video or audio samples collected using synchronized capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData
type AVCaptureSynchronizedSampleBufferData struct {
	AVCaptureSynchronizedData
}

// AVCaptureSynchronizedSampleBufferDataFrom constructs a [AVCaptureSynchronizedSampleBufferData] from an unsafe.Pointer.
//
// A container for video or audio samples collected using synchronized capture.
func AVCaptureSynchronizedSampleBufferDataFrom(ptr unsafe.Pointer) AVCaptureSynchronizedSampleBufferData {
	return AVCaptureSynchronizedSampleBufferData{
		AVCaptureSynchronizedData: AVCaptureSynchronizedDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureSynchronizedSampleBufferDataClass) Alloc() AVCaptureSynchronizedSampleBufferData {
	rv := objc.Send[AVCaptureSynchronizedSampleBufferData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureSynchronizedSampleBufferDataClass) New() AVCaptureSynchronizedSampleBufferData {
	rv := objc.Send[AVCaptureSynchronizedSampleBufferData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureSynchronizedSampleBufferData) Init() AVCaptureSynchronizedSampleBufferData {
	rv := objc.Send[AVCaptureSynchronizedSampleBufferData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureSynchronizedSampleBufferData) Autorelease() AVCaptureSynchronizedSampleBufferData {
	rv := objc.Send[AVCaptureSynchronizedSampleBufferData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSynchronizedSampleBufferData creates a new AVCaptureSynchronizedSampleBufferData instance.
func NewAVCaptureSynchronizedSampleBufferData() AVCaptureSynchronizedSampleBufferData {
	return getAVCaptureSynchronizedSampleBufferDataClass().New()
}




