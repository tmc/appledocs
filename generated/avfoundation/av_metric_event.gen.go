// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMetricEvent */


/* debug [class_header]: Header for AVMetricEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricEvent */
// An interface definition for the [MetricEvent] class.
type IMetricEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MetricEvent */
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	MediaTime() objc.IObject /* cross-framework: Time */
	SessionID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/date
func (m_ MetricEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/mediaTime
func (m_ MetricEvent) MediaTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("mediaTime"))
	return rv
}/* debug [instance_properties/getter]: mediaTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/sessionID
func (m_ MetricEvent) SessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricEvent */



