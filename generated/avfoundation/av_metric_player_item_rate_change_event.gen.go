// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemRateChangeEvent */


/* debug [class_header]: Header for AVMetricPlayerItemRateChangeEvent */
// The class instance for the [MetricPlayerItemRateChangeEvent] class.
var (
	MetricPlayerItemRateChangeEventClass     _MetricPlayerItemRateChangeEventClass
	MetricPlayerItemRateChangeEventClassOnce sync.Once
)

func getMetricPlayerItemRateChangeEventClass() _MetricPlayerItemRateChangeEventClass {
	MetricPlayerItemRateChangeEventClassOnce.Do(func() {
		MetricPlayerItemRateChangeEventClass = _MetricPlayerItemRateChangeEventClass{objc.GetClass("AVMetricPlayerItemRateChangeEvent")}
	})
	return MetricPlayerItemRateChangeEventClass
}

type _MetricPlayerItemRateChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemRateChangeEvent */
// An interface definition for the [MetricPlayerItemRateChangeEvent] class.
type IMetricPlayerItemRateChangeEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemRateChangeEvent */
	// properties:
	PreviousRate() float64
	Rate() float64
	Variant() IAVAssetVariant
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemRateChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemRateChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemRateChangeEventClass) Alloc() MetricPlayerItemRateChangeEvent {
	rv := objc.Send[MetricPlayerItemRateChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemRateChangeEventClass) New() MetricPlayerItemRateChangeEvent {
	rv := objc.Send[MetricPlayerItemRateChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemRateChangeEvent) Init() MetricPlayerItemRateChangeEvent {
	rv := objc.Send[MetricPlayerItemRateChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemRateChangeEvent) Autorelease() MetricPlayerItemRateChangeEvent {
	rv := objc.Send[MetricPlayerItemRateChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemRateChangeEvent creates a new MetricPlayerItemRateChangeEvent instance.
func NewMetricPlayerItemRateChangeEvent() MetricPlayerItemRateChangeEvent {
	return getMetricPlayerItemRateChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemRateChangeEvent */
// An event that represents when the playback rate changes.


// An event that represents when the playback rate changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent
type MetricPlayerItemRateChangeEvent struct {
	MetricEvent
}

// MetricPlayerItemRateChangeEventFrom constructs a [MetricPlayerItemRateChangeEvent] from an unsafe.Pointer.
//
// An event that represents when the playback rate changes.
func MetricPlayerItemRateChangeEventFrom(ptr unsafe.Pointer) MetricPlayerItemRateChangeEvent {
	return MetricPlayerItemRateChangeEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemRateChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemRateChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemRateChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemRateChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemRateChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/previousRate
func (m_ MetricPlayerItemRateChangeEvent) PreviousRate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("previousRate"))
	return rv
}/* debug [instance_properties/getter]: previousRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/rate
func (m_ MetricPlayerItemRateChangeEvent) Rate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/variant
func (m_ MetricPlayerItemRateChangeEvent) Variant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("variant"))
	return rv
}/* debug [instance_properties/getter]: variant */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemRateChangeEvent */



