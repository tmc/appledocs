// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricPlayerItemStallEvent] class.
var (
	MetricPlayerItemStallEventClass     _MetricPlayerItemStallEventClass
	MetricPlayerItemStallEventClassOnce sync.Once
)

func getMetricPlayerItemStallEventClass() _MetricPlayerItemStallEventClass {
	MetricPlayerItemStallEventClassOnce.Do(func() {
		MetricPlayerItemStallEventClass = _MetricPlayerItemStallEventClass{objc.GetClass("AVMetricPlayerItemStallEvent")}
	})
	return MetricPlayerItemStallEventClass
}

type _MetricPlayerItemStallEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricPlayerItemStallEvent] class.
type IMetricPlayerItemStallEvent interface {
	IMetricPlayerItemRateChangeEvent
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemStallEventClass) Alloc() MetricPlayerItemStallEvent {
	rv := objc.Send[MetricPlayerItemStallEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemStallEventClass) New() MetricPlayerItemStallEvent {
	rv := objc.Send[MetricPlayerItemStallEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemStallEvent) Init() MetricPlayerItemStallEvent {
	rv := objc.Send[MetricPlayerItemStallEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemStallEvent) Autorelease() MetricPlayerItemStallEvent {
	rv := objc.Send[MetricPlayerItemStallEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemStallEvent creates a new MetricPlayerItemStallEvent instance.
func NewMetricPlayerItemStallEvent() MetricPlayerItemStallEvent {
	return getMetricPlayerItemStallEventClass().New()
}





// An event that represents when playback stalls.


// An event that represents when playback stalls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemStallEvent
type MetricPlayerItemStallEvent struct {
	MetricPlayerItemRateChangeEvent
}

// MetricPlayerItemStallEventFrom constructs a [MetricPlayerItemStallEvent] from an unsafe.Pointer.
//
// An event that represents when playback stalls.
func MetricPlayerItemStallEventFrom(ptr unsafe.Pointer) MetricPlayerItemStallEvent {
	return MetricPlayerItemStallEvent{
		MetricPlayerItemRateChangeEvent: MetricPlayerItemRateChangeEventFrom(ptr),
	}
}































