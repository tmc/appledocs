// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricErrorEvent] class.
var (
	MetricErrorEventClass     _MetricErrorEventClass
	MetricErrorEventClassOnce sync.Once
)

func getMetricErrorEventClass() _MetricErrorEventClass {
	MetricErrorEventClassOnce.Do(func() {
		MetricErrorEventClass = _MetricErrorEventClass{objc.GetClass("AVMetricErrorEvent")}
	})
	return MetricErrorEventClass
}

type _MetricErrorEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricErrorEvent] class.
type IMetricErrorEvent interface {
	IMetricEvent
	

	// properties:
	DidRecover() bool
	Error() foundation.foundation.INSError


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricErrorEventClass) Alloc() MetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricErrorEventClass) New() MetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricErrorEvent) Init() MetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricErrorEvent) Autorelease() MetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricErrorEvent creates a new MetricErrorEvent instance.
func NewMetricErrorEvent() MetricErrorEvent {
	return getMetricErrorEventClass().New()
}





// An object that represents a metric event when an error occurs.


// An object that represents a metric event when an error occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent
type MetricErrorEvent struct {
	MetricEvent
}

// MetricErrorEventFrom constructs a [MetricErrorEvent] from an unsafe.Pointer.
//
// An object that represents a metric event when an error occurs.
func MetricErrorEventFrom(ptr unsafe.Pointer) MetricErrorEvent {
	return MetricErrorEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}

























// A Boolean value that indicates whether the error was recoverable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/didRecover
func (m_ MetricErrorEvent) DidRecover() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didRecover"))
	return rv
}


// Returns the error event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/error
func (m_ MetricErrorEvent) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](m_.ID, objc.Sel("error"))
	return rv
}








