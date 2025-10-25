// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricErrorEvent */


/* debug [class_header]: Header for AVMetricErrorEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricErrorEvent */
// An interface definition for the [MetricErrorEvent] class.
type IMetricErrorEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricErrorEvent */
	// properties:
	DidRecover() bool
	Error() Error
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricErrorEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricErrorEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricErrorEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricErrorEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricErrorEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricErrorEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricErrorEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricErrorEvent */

// A Boolean value that indicates whether the error was recoverable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/didRecover
func (m_ MetricErrorEvent) DidRecover() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didRecover"))
	return rv
}/* debug [instance_properties/getter]: didRecover */


// Returns the error event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/error
func (m_ MetricErrorEvent) Error() Error {
	rv := objc.Send[Error](m_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricErrorEvent */



