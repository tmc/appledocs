// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemLikelyToKeepUpEvent */


/* debug [class_header]: Header for AVMetricPlayerItemLikelyToKeepUpEvent */
// The class instance for the [MetricPlayerItemLikelyToKeepUpEvent] class.
var (
	MetricPlayerItemLikelyToKeepUpEventClass     _MetricPlayerItemLikelyToKeepUpEventClass
	MetricPlayerItemLikelyToKeepUpEventClassOnce sync.Once
)

func getMetricPlayerItemLikelyToKeepUpEventClass() _MetricPlayerItemLikelyToKeepUpEventClass {
	MetricPlayerItemLikelyToKeepUpEventClassOnce.Do(func() {
		MetricPlayerItemLikelyToKeepUpEventClass = _MetricPlayerItemLikelyToKeepUpEventClass{objc.GetClass("AVMetricPlayerItemLikelyToKeepUpEvent")}
	})
	return MetricPlayerItemLikelyToKeepUpEventClass
}

type _MetricPlayerItemLikelyToKeepUpEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemLikelyToKeepUpEvent */
// An interface definition for the [MetricPlayerItemLikelyToKeepUpEvent] class.
type IMetricPlayerItemLikelyToKeepUpEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemLikelyToKeepUpEvent */
	// properties:
	LoadedTimeRanges() []foundation.Value
	TimeTaken() float64
	Variant() IAVAssetVariant
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemLikelyToKeepUpEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemLikelyToKeepUpEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemLikelyToKeepUpEventClass) Alloc() MetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemLikelyToKeepUpEventClass) New() MetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemLikelyToKeepUpEvent) Init() MetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemLikelyToKeepUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemLikelyToKeepUpEvent) Autorelease() MetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemLikelyToKeepUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemLikelyToKeepUpEvent creates a new MetricPlayerItemLikelyToKeepUpEvent instance.
func NewMetricPlayerItemLikelyToKeepUpEvent() MetricPlayerItemLikelyToKeepUpEvent {
	return getMetricPlayerItemLikelyToKeepUpEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemLikelyToKeepUpEvent */
// An event that represents when playback is likely to continue without stalling.


// An event that represents when playback is likely to continue without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent
type MetricPlayerItemLikelyToKeepUpEvent struct {
	MetricEvent
}

// MetricPlayerItemLikelyToKeepUpEventFrom constructs a [MetricPlayerItemLikelyToKeepUpEvent] from an unsafe.Pointer.
//
// An event that represents when playback is likely to continue without stalling.
func MetricPlayerItemLikelyToKeepUpEventFrom(ptr unsafe.Pointer) MetricPlayerItemLikelyToKeepUpEvent {
	return MetricPlayerItemLikelyToKeepUpEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemLikelyToKeepUpEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemLikelyToKeepUpEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemLikelyToKeepUpEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemLikelyToKeepUpEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemLikelyToKeepUpEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/loadedTimeRanges-960vi
func (m_ MetricPlayerItemLikelyToKeepUpEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}/* debug [instance_properties/getter]: loadedTimeRanges */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/timeTaken
func (m_ MetricPlayerItemLikelyToKeepUpEvent) TimeTaken() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("timeTaken"))
	return rv
}/* debug [instance_properties/getter]: timeTaken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/variant
func (m_ MetricPlayerItemLikelyToKeepUpEvent) Variant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("variant"))
	return rv
}/* debug [instance_properties/getter]: variant */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemLikelyToKeepUpEvent */



