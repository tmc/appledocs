// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricPlayerItemSeekDidCompleteEvent] class.
var (
	MetricPlayerItemSeekDidCompleteEventClass     _MetricPlayerItemSeekDidCompleteEventClass
	MetricPlayerItemSeekDidCompleteEventClassOnce sync.Once
)

func getMetricPlayerItemSeekDidCompleteEventClass() _MetricPlayerItemSeekDidCompleteEventClass {
	MetricPlayerItemSeekDidCompleteEventClassOnce.Do(func() {
		MetricPlayerItemSeekDidCompleteEventClass = _MetricPlayerItemSeekDidCompleteEventClass{objc.GetClass("AVMetricPlayerItemSeekDidCompleteEvent")}
	})
	return MetricPlayerItemSeekDidCompleteEventClass
}

type _MetricPlayerItemSeekDidCompleteEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricPlayerItemSeekDidCompleteEvent] class.
type IMetricPlayerItemSeekDidCompleteEvent interface {
	IMetricPlayerItemRateChangeEvent
	

	// properties:
	DidSeekInBuffer() bool


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemSeekDidCompleteEventClass) Alloc() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemSeekDidCompleteEventClass) New() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemSeekDidCompleteEvent) Init() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemSeekDidCompleteEvent) Autorelease() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemSeekDidCompleteEvent creates a new MetricPlayerItemSeekDidCompleteEvent instance.
func NewMetricPlayerItemSeekDidCompleteEvent() MetricPlayerItemSeekDidCompleteEvent {
	return getMetricPlayerItemSeekDidCompleteEventClass().New()
}





// An event that represents when the playback seek completes.


// An event that represents when the playback seek completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent
type MetricPlayerItemSeekDidCompleteEvent struct {
	MetricPlayerItemRateChangeEvent
}

// MetricPlayerItemSeekDidCompleteEventFrom constructs a [MetricPlayerItemSeekDidCompleteEvent] from an unsafe.Pointer.
//
// An event that represents when the playback seek completes.
func MetricPlayerItemSeekDidCompleteEventFrom(ptr unsafe.Pointer) MetricPlayerItemSeekDidCompleteEvent {
	return MetricPlayerItemSeekDidCompleteEvent{
		MetricPlayerItemRateChangeEvent: MetricPlayerItemRateChangeEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent/didSeekInBuffer
func (m_ MetricPlayerItemSeekDidCompleteEvent) DidSeekInBuffer() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didSeekInBuffer"))
	return rv
}








