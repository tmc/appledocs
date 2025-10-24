// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetricPlayerItemVariantSwitchEvent] class.
type IMetricPlayerItemVariantSwitchEvent interface {
	IMetricEvent
	

	// properties:
	AudioRendition() IAVMetricMediaRendition
	DidSucceed() bool
	FromVariant() IAVAssetVariant
	LoadedTimeRanges() []foundation.Value
	SubtitleRendition() IAVMetricMediaRendition
	ToVariant() IAVAssetVariant
	VideoRendition() IAVMetricMediaRendition


	

	// methods:


}





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

























// Represents the currently selected video rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/audioRendition
func (m_ MetricPlayerItemVariantSwitchEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("audioRendition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/didSucceed
func (m_ MetricPlayerItemVariantSwitchEvent) DidSucceed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didSucceed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/fromVariant
func (m_ MetricPlayerItemVariantSwitchEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("fromVariant"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/loadedTimeRanges-4rhjw
func (m_ MetricPlayerItemVariantSwitchEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// Represents the currently selected audio rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/subtitleRendition
func (m_ MetricPlayerItemVariantSwitchEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("subtitleRendition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/toVariant
func (m_ MetricPlayerItemVariantSwitchEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("toVariant"))
	return rv
}


// Represents the currently selected video rendition’s identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/videoRendition
func (m_ MetricPlayerItemVariantSwitchEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("videoRendition"))
	return rv
}








