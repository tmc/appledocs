// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemStallEvent */


/* debug [class_header]: Header for AVMetricPlayerItemStallEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemStallEvent */
// An interface definition for the [MetricPlayerItemStallEvent] class.
type IMetricPlayerItemStallEvent interface {
	IMetricPlayerItemRateChangeEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemStallEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemStallEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemStallEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemStallEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemStallEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemStallEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemStallEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemStallEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemStallEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemStallEvent */



