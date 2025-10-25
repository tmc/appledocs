// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemVariantSwitchEvent */


/* debug [class_header]: Header for AVMetricPlayerItemVariantSwitchEvent */
// The class instance for the [MetricPlayerItemVariantSwitchEvent] class.
var (
	MetricPlayerItemVariantSwitchEventClass     _MetricPlayerItemVariantSwitchEventClass
	MetricPlayerItemVariantSwitchEventClassOnce sync.Once
)

func getMetricPlayerItemVariantSwitchEventClass() _MetricPlayerItemVariantSwitchEventClass {
	MetricPlayerItemVariantSwitchEventClassOnce.Do(func() {
		MetricPlayerItemVariantSwitchEventClass = _MetricPlayerItemVariantSwitchEventClass{objc.GetClass("AVMetricPlayerItemVariantSwitchEvent")}
	})
	return MetricPlayerItemVariantSwitchEventClass
}

type _MetricPlayerItemVariantSwitchEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemVariantSwitchEvent */
// An interface definition for the [MetricPlayerItemVariantSwitchEvent] class.
type IMetricPlayerItemVariantSwitchEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemVariantSwitchEvent */
	// properties:
	AudioRendition() IAVMetricMediaRendition
	DidSucceed() bool
	FromVariant() IAVAssetVariant
	LoadedTimeRanges() []foundation.Value
	SubtitleRendition() IAVMetricMediaRendition
	ToVariant() IAVAssetVariant
	VideoRendition() IAVMetricMediaRendition
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemVariantSwitchEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemVariantSwitchEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemVariantSwitchEventClass) Alloc() MetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemVariantSwitchEventClass) New() MetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemVariantSwitchEvent) Init() MetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemVariantSwitchEvent) Autorelease() MetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[MetricPlayerItemVariantSwitchEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemVariantSwitchEvent creates a new MetricPlayerItemVariantSwitchEvent instance.
func NewMetricPlayerItemVariantSwitchEvent() MetricPlayerItemVariantSwitchEvent {
	return getMetricPlayerItemVariantSwitchEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemVariantSwitchEvent */
// An event that represents when the player completes a variant switch.


// An event that represents when the player completes a variant switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent
type MetricPlayerItemVariantSwitchEvent struct {
	MetricEvent
}

// MetricPlayerItemVariantSwitchEventFrom constructs a [MetricPlayerItemVariantSwitchEvent] from an unsafe.Pointer.
//
// An event that represents when the player completes a variant switch.
func MetricPlayerItemVariantSwitchEventFrom(ptr unsafe.Pointer) MetricPlayerItemVariantSwitchEvent {
	return MetricPlayerItemVariantSwitchEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemVariantSwitchEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemVariantSwitchEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemVariantSwitchEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemVariantSwitchEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemVariantSwitchEvent */

// Represents the currently selected video rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/audioRendition
func (m_ MetricPlayerItemVariantSwitchEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("audioRendition"))
	return rv
}/* debug [instance_properties/getter]: audioRendition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/didSucceed
func (m_ MetricPlayerItemVariantSwitchEvent) DidSucceed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didSucceed"))
	return rv
}/* debug [instance_properties/getter]: didSucceed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/fromVariant
func (m_ MetricPlayerItemVariantSwitchEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("fromVariant"))
	return rv
}/* debug [instance_properties/getter]: fromVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/loadedTimeRanges-4rhjw
func (m_ MetricPlayerItemVariantSwitchEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}/* debug [instance_properties/getter]: loadedTimeRanges */


// Represents the currently selected audio rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/subtitleRendition
func (m_ MetricPlayerItemVariantSwitchEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("subtitleRendition"))
	return rv
}/* debug [instance_properties/getter]: subtitleRendition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/toVariant
func (m_ MetricPlayerItemVariantSwitchEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("toVariant"))
	return rv
}/* debug [instance_properties/getter]: toVariant */


// Represents the currently selected video rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/videoRendition
func (m_ MetricPlayerItemVariantSwitchEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("videoRendition"))
	return rv
}/* debug [instance_properties/getter]: videoRendition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemVariantSwitchEvent */



