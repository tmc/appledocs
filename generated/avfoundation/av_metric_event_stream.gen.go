// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMetricEventStream */


/* debug [class_header]: Header for AVMetricEventStream */
// The class instance for the [MetricEventStream] class.
var (
	MetricEventStreamClass     _MetricEventStreamClass
	MetricEventStreamClassOnce sync.Once
)

func getMetricEventStreamClass() _MetricEventStreamClass {
	MetricEventStreamClassOnce.Do(func() {
		MetricEventStreamClass = _MetricEventStreamClass{objc.GetClass("AVMetricEventStream")}
	})
	return MetricEventStreamClass
}

type _MetricEventStreamClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricEventStream */
// An interface definition for the [MetricEventStream] class.
type IMetricEventStream interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MetricEventStream */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricEventStream */
	// methods:
	AddPublisher(publisher unsafe.Pointer) bool
	SetSubscriberQueue(subscriber unsafe.Pointer, queue objectivec.IObject) bool
	SubscribeToAllMetricEvents()
	SubscribeToMetricEvent(metricEventClass objc.Class)
	SubscribeToMetricEvents(metricEventClasses []objc.Class)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricEventStream */
// Alloc allocates a new instance without initialization.
func (mc _MetricEventStreamClass) Alloc() MetricEventStream {
	rv := objc.Send[MetricEventStream](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricEventStreamClass) New() MetricEventStream {
	rv := objc.Send[MetricEventStream](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricEventStream) Init() MetricEventStream {
	rv := objc.Send[MetricEventStream](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricEventStream) Autorelease() MetricEventStream {
	rv := objc.Send[MetricEventStream](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricEventStream creates a new MetricEventStream instance.
func NewMetricEventStream() MetricEventStream {
	return getMetricEventStreamClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricEventStream */
// An object that allows clients to add publishers and then subscribe to specific metric event classes from those publishers.
//
// Publishers are types that adopt . The protocol allows clients to receive metric events with a subscriber delegate which adopts the protocol.


// An object that allows clients to add publishers and then subscribe to specific metric event classes from those publishers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream
type MetricEventStream struct {
	objectivec.Object
}

// MetricEventStreamFrom constructs a [MetricEventStream] from an unsafe.Pointer.
//
// An object that allows clients to add publishers and then subscribe to specific metric event classes from those publishers.
func MetricEventStreamFrom(ptr unsafe.Pointer) MetricEventStream {
	return MetricEventStream{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricEventStream *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricEventStream */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/eventStream
func (mc _MetricEventStreamClass) EventStream() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("eventStream"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EventStream) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricEventStream */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricEventStream */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/addPublisher:
func (m_ MetricEventStream) AddPublisher(publisher unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("addPublisher:"), publisher)
	return rv
}/* debug [instance_methods/method]: AddPublisher */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/setSubscriber:queue:
func (m_ MetricEventStream) SetSubscriberQueue(subscriber unsafe.Pointer, queue objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setSubscriber:queue:"), subscriber, queue)
	return rv
}/* debug [instance_methods/method]: SetSubscriberQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToAllMetricEvents
func (m_ MetricEventStream) SubscribeToAllMetricEvents() {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToAllMetricEvents"))
}/* debug [instance_methods/method]: SubscribeToAllMetricEvents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToMetricEvent:
func (m_ MetricEventStream) SubscribeToMetricEvent(metricEventClass objc.Class) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToMetricEvent:"), metricEventClass)
}/* debug [instance_methods/method]: SubscribeToMetricEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToMetricEvents:
func (m_ MetricEventStream) SubscribeToMetricEvents(metricEventClasses []objc.Class) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToMetricEvents:"), metricEventClasses)
}/* debug [instance_methods/method]: SubscribeToMetricEvents */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricEventStream */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricEventStream */



