// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [MetricMediaResourceRequestEvent] class.
var (
	MetricMediaResourceRequestEventClass     _MetricMediaResourceRequestEventClass
	MetricMediaResourceRequestEventClassOnce sync.Once
)

func getMetricMediaResourceRequestEventClass() _MetricMediaResourceRequestEventClass {
	MetricMediaResourceRequestEventClassOnce.Do(func() {
		MetricMediaResourceRequestEventClass = _MetricMediaResourceRequestEventClass{objc.GetClass("AVMetricMediaResourceRequestEvent")}
	})
	return MetricMediaResourceRequestEventClass
}

type _MetricMediaResourceRequestEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricMediaResourceRequestEvent] class.
type IMetricMediaResourceRequestEvent interface {
	IMetricEvent
	

	// properties:
	ByteRange() corefoundation.Range
	ErrorEvent() IAVMetricErrorEvent
	NetworkTransactionMetrics() foundation.URLSessionTaskMetrics
	RequestEndTime() objc.IObject /* cross-framework: NSDate */
	RequestStartTime() objc.IObject /* cross-framework: NSDate */
	ResponseEndTime() objc.IObject /* cross-framework: NSDate */
	ResponseStartTime() objc.IObject /* cross-framework: NSDate */
	ServerAddress() objc.IObject /* cross-framework: NSString */
	Url() objc.IObject /* cross-framework: NSURL */
	ReadFromCache() bool
	WasReadFromCache() bool
	SetWasReadFromCache(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricMediaResourceRequestEventClass) Alloc() MetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricMediaResourceRequestEventClass) New() MetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricMediaResourceRequestEvent) Init() MetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricMediaResourceRequestEvent) Autorelease() MetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricMediaResourceRequestEvent creates a new MetricMediaResourceRequestEvent instance.
func NewMetricMediaResourceRequestEvent() MetricMediaResourceRequestEvent {
	return getMetricMediaResourceRequestEventClass().New()
}





// An event that represents a media resource request.


// An event that represents a media resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent
type MetricMediaResourceRequestEvent struct {
	MetricEvent
}

// MetricMediaResourceRequestEventFrom constructs a [MetricMediaResourceRequestEvent] from an unsafe.Pointer.
//
// An event that represents a media resource request.
func MetricMediaResourceRequestEventFrom(ptr unsafe.Pointer) MetricMediaResourceRequestEvent {
	return MetricMediaResourceRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/byteRange
func (m_ MetricMediaResourceRequestEvent) ByteRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](m_.ID, objc.Sel("byteRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/errorEvent
func (m_ MetricMediaResourceRequestEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("errorEvent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/networkTransactionMetrics
func (m_ MetricMediaResourceRequestEvent) NetworkTransactionMetrics() foundation.URLSessionTaskMetrics {
	rv := objc.Send[foundation.URLSessionTaskMetrics](m_.ID, objc.Sel("networkTransactionMetrics"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestEndTime
func (m_ MetricMediaResourceRequestEvent) RequestEndTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("requestEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestStartTime
func (m_ MetricMediaResourceRequestEvent) RequestStartTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("requestStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseEndTime
func (m_ MetricMediaResourceRequestEvent) ResponseEndTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("responseEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseStartTime
func (m_ MetricMediaResourceRequestEvent) ResponseStartTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("responseStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/serverAddress
func (m_ MetricMediaResourceRequestEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serverAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/url
func (m_ MetricMediaResourceRequestEvent) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/wasReadFromCache
func (m_ MetricMediaResourceRequestEvent) ReadFromCache() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromCache"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetricmediaresourcerequestevent/wasreadfromcache
func (m_ MetricMediaResourceRequestEvent) WasReadFromCache() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("wasReadFromCache"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetricmediaresourcerequestevent/wasreadfromcache
func (m_ MetricMediaResourceRequestEvent) SetWasReadFromCache(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasReadFromCache:"), value)
}








