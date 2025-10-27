// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetricPlayerItemVariantSwitchStartEvent] class.
type IMetricPlayerItemVariantSwitchStartEvent interface {
	IMetricEvent
	

	// properties:
	AudioRendition() IAVMetricMediaRendition
	FromVariant() IAVAssetVariant
	LoadedTimeRanges() []foundation.Value
	SubtitleRendition() IAVMetricMediaRendition
	ToVariant() IAVAssetVariant
	VideoRendition() IAVMetricMediaRendition


	

	// methods:


}





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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/audioRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("audioRendition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/fromVariant
func (m_ MetricPlayerItemVariantSwitchStartEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("fromVariant"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/loadedTimeRanges-3svh3
func (m_ MetricPlayerItemVariantSwitchStartEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/subtitleRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("subtitleRendition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/toVariant
func (m_ MetricPlayerItemVariantSwitchStartEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("toVariant"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/videoRendition
func (m_ MetricPlayerItemVariantSwitchStartEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("videoRendition"))
	return rv
}








