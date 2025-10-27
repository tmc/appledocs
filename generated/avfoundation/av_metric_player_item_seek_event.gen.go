// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricPlayerItemSeekEvent] class.
var (
	MetricPlayerItemSeekEventClass     _MetricPlayerItemSeekEventClass
	MetricPlayerItemSeekEventClassOnce sync.Once
)

func getMetricPlayerItemSeekEventClass() _MetricPlayerItemSeekEventClass {
	MetricPlayerItemSeekEventClassOnce.Do(func() {
		MetricPlayerItemSeekEventClass = _MetricPlayerItemSeekEventClass{objc.GetClass("AVMetricPlayerItemSeekEvent")}
	})
	return MetricPlayerItemSeekEventClass
}

type _MetricPlayerItemSeekEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricPlayerItemSeekEvent] class.
type IMetricPlayerItemSeekEvent interface {
	IMetricPlayerItemRateChangeEvent
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemSeekEventClass) Alloc() MetricPlayerItemSeekEvent {
	rv := objc.Send[MetricPlayerItemSeekEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemSeekEventClass) New() MetricPlayerItemSeekEvent {
	rv := objc.Send[MetricPlayerItemSeekEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemSeekEvent) Init() MetricPlayerItemSeekEvent {
	rv := objc.Send[MetricPlayerItemSeekEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemSeekEvent) Autorelease() MetricPlayerItemSeekEvent {
	rv := objc.Send[MetricPlayerItemSeekEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemSeekEvent creates a new MetricPlayerItemSeekEvent instance.
func NewMetricPlayerItemSeekEvent() MetricPlayerItemSeekEvent {
	return getMetricPlayerItemSeekEventClass().New()
}





// An event that represents when a playback seek occurs.


// An event that represents when a playback seek occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekEvent
type MetricPlayerItemSeekEvent struct {
	MetricPlayerItemRateChangeEvent
}

// MetricPlayerItemSeekEventFrom constructs a [MetricPlayerItemSeekEvent] from an unsafe.Pointer.
//
// An event that represents when a playback seek occurs.
func MetricPlayerItemSeekEventFrom(ptr unsafe.Pointer) MetricPlayerItemSeekEvent {
	return MetricPlayerItemSeekEvent{
		MetricPlayerItemRateChangeEvent: MetricPlayerItemRateChangeEventFrom(ptr),
	}
}































