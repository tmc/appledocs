// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MetricEvent] class.
var (
	MetricEventClass     _MetricEventClass
	MetricEventClassOnce sync.Once
)

func getMetricEventClass() _MetricEventClass {
	MetricEventClassOnce.Do(func() {
		MetricEventClass = _MetricEventClass{objc.GetClass("AVMetricEvent")}
	})
	return MetricEventClass
}

type _MetricEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricEvent] class.
type IMetricEvent interface {
	objectivec.IObject
	

	// properties:
	Date() foundation.foundation.INSDate
	MediaTime() objectivec.IObject
	SessionID() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricEventClass) Alloc() MetricEvent {
	rv := objc.Send[MetricEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricEventClass) New() MetricEvent {
	rv := objc.Send[MetricEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricEvent) Init() MetricEvent {
	rv := objc.Send[MetricEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricEvent) Autorelease() MetricEvent {
	rv := objc.Send[MetricEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricEvent creates a new MetricEvent instance.
func NewMetricEvent() MetricEvent {
	return getMetricEventClass().New()
}





// A base class that represents a metric event.


// A base class that represents a metric event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent
type MetricEvent struct {
	objectivec.Object
}

// MetricEventFrom constructs a [MetricEvent] from an unsafe.Pointer.
//
// A base class that represents a metric event.
func MetricEventFrom(ptr unsafe.Pointer) MetricEvent {
	return MetricEvent{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/date
func (m_ MetricEvent) Date() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("date"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/mediaTime
func (m_ MetricEvent) MediaTime() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("mediaTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/sessionID
func (m_ MetricEvent) SessionID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("sessionID"))
	return rv
}








