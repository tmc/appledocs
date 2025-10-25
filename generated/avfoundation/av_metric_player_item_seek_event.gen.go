// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemSeekEvent */


/* debug [class_header]: Header for AVMetricPlayerItemSeekEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemSeekEvent */
// An interface definition for the [MetricPlayerItemSeekEvent] class.
type IMetricPlayerItemSeekEvent interface {
	IMetricPlayerItemRateChangeEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemSeekEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemSeekEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemSeekEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemSeekEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemSeekEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemSeekEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemSeekEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemSeekEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemSeekEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemSeekEvent */



