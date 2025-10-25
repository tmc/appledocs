// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemVariantSwitchStartEvent */


/* debug [class_header]: Header for AVMetricPlayerItemVariantSwitchStartEvent */
// The class instance for the [MetricPlayerItemVariantSwitchStartEvent] class.
var (
	MetricPlayerItemVariantSwitchStartEventClass     _MetricPlayerItemVariantSwitchStartEventClass
	MetricPlayerItemVariantSwitchStartEventClassOnce sync.Once
)

func getMetricPlayerItemVariantSwitchStartEventClass() _MetricPlayerItemVariantSwitchStartEventClass {
	MetricPlayerItemVariantSwitchStartEventClassOnce.Do(func() {
		MetricPlayerItemVariantSwitchStartEventClass = _MetricPlayerItemVariantSwitchStartEventClass{objc.GetClass("AVMetricPlayerItemVariantSwitchStartEvent")}
	})
	return MetricPlayerItemVariantSwitchStartEventClass
}

type _MetricPlayerItemVariantSwitchStartEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemVariantSwitchStartEvent */
// An interface definition for the [MetricPlayerItemVariantSwitchStartEvent] class.
type IMetricPlayerItemVariantSwitchStartEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemVariantSwitchStartEvent */
	// properties:
	AudioRendition() IAVMetricMediaRendition
	FromVariant() IAVAssetVariant
	LoadedTimeRanges() []foundation.Value
	SubtitleRendition() IAVMetricMediaRendition
	ToVariant() IAVAssetVariant
	VideoRendition() IAVMetricMediaRendition
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemVariantSwitchStartEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemVariantSwitchStartEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemVariantSwitchStartEventClass) Alloc() MetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchStartEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemVariantSwitchStartEventClass) New() MetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchStartEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemVariantSwitchStartEvent) Init() MetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchStartEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemVariantSwitchStartEvent) Autorelease() MetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchStartEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemVariantSwitchStartEvent creates a new MetricPlayerItemVariantSwitchStartEvent instance.
func NewMetricPlayerItemVariantSwitchStartEvent() MetricPlayerItemVariantSwitchStartEvent {
	return getMetricPlayerItemVariantSwitchStartEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemVariantSwitchStartEvent */
// An event that represents when the player attempts a variant switch.


// An event that represents when the player attempts a variant switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent
type MetricPlayerItemVariantSwitchStartEvent struct {
	MetricEvent
}

// MetricPlayerItemVariantSwitchStartEventFrom constructs a [MetricPlayerItemVariantSwitchStartEvent] from an unsafe.Pointer.
//
// An event that represents when the player attempts a variant switch.
func MetricPlayerItemVariantSwitchStartEventFrom(ptr unsafe.Pointer) MetricPlayerItemVariantSwitchStartEvent {
	return MetricPlayerItemVariantSwitchStartEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemVariantSwitchStartEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemVariantSwitchStartEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemVariantSwitchStartEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemVariantSwitchStartEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemVariantSwitchStartEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/audioRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("audioRendition"))
	return rv
}/* debug [instance_properties/getter]: audioRendition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/fromVariant
func (m_ MetricPlayerItemVariantSwitchStartEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("fromVariant"))
	return rv
}/* debug [instance_properties/getter]: fromVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/loadedTimeRanges-3svh3
func (m_ MetricPlayerItemVariantSwitchStartEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}/* debug [instance_properties/getter]: loadedTimeRanges */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/subtitleRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("subtitleRendition"))
	return rv
}/* debug [instance_properties/getter]: subtitleRendition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/toVariant
func (m_ MetricPlayerItemVariantSwitchStartEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("toVariant"))
	return rv
}/* debug [instance_properties/getter]: toVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/videoRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("videoRendition"))
	return rv
}/* debug [instance_properties/getter]: videoRendition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemVariantSwitchStartEvent */



