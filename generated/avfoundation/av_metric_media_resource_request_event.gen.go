// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricMediaResourceRequestEvent */


/* debug [class_header]: Header for AVMetricMediaResourceRequestEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricMediaResourceRequestEvent */
// An interface definition for the [MetricMediaResourceRequestEvent] class.
type IMetricMediaResourceRequestEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricMediaResourceRequestEvent */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricMediaResourceRequestEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricMediaResourceRequestEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricMediaResourceRequestEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricMediaResourceRequestEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricMediaResourceRequestEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricMediaResourceRequestEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricMediaResourceRequestEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricMediaResourceRequestEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/byteRange
func (m_ MetricMediaResourceRequestEvent) ByteRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](m_.ID, objc.Sel("byteRange"))
	return rv
}/* debug [instance_properties/getter]: byteRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/errorEvent
func (m_ MetricMediaResourceRequestEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("errorEvent"))
	return rv
}/* debug [instance_properties/getter]: errorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/networkTransactionMetrics
func (m_ MetricMediaResourceRequestEvent) NetworkTransactionMetrics() foundation.URLSessionTaskMetrics {
	rv := objc.Send[foundation.URLSessionTaskMetrics](m_.ID, objc.Sel("networkTransactionMetrics"))
	return rv
}/* debug [instance_properties/getter]: networkTransactionMetrics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestEndTime
func (m_ MetricMediaResourceRequestEvent) RequestEndTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("requestEndTime"))
	return rv
}/* debug [instance_properties/getter]: requestEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestStartTime
func (m_ MetricMediaResourceRequestEvent) RequestStartTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("requestStartTime"))
	return rv
}/* debug [instance_properties/getter]: requestStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseEndTime
func (m_ MetricMediaResourceRequestEvent) ResponseEndTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("responseEndTime"))
	return rv
}/* debug [instance_properties/getter]: responseEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseStartTime
func (m_ MetricMediaResourceRequestEvent) ResponseStartTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("responseStartTime"))
	return rv
}/* debug [instance_properties/getter]: responseStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/serverAddress
func (m_ MetricMediaResourceRequestEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serverAddress"))
	return rv
}/* debug [instance_properties/getter]: serverAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/url
func (m_ MetricMediaResourceRequestEvent) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/wasReadFromCache
func (m_ MetricMediaResourceRequestEvent) ReadFromCache() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readFromCache"))
	return rv
}/* debug [instance_properties/getter]: readFromCache */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetricmediaresourcerequestevent/wasreadfromcache
func (m_ MetricMediaResourceRequestEvent) WasReadFromCache() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("wasReadFromCache"))
	return rv
}/* debug [instance_properties/getter]: wasReadFromCache */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetricmediaresourcerequestevent/wasreadfromcache
func (m_ MetricMediaResourceRequestEvent) SetWasReadFromCache(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWasReadFromCache:"), value)
}/* debug [instance_properties/setter]: wasReadFromCache */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricMediaResourceRequestEvent */



