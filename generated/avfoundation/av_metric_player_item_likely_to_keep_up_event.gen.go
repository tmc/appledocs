// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetricPlayerItemLikelyToKeepUpEvent] class.
type IMetricPlayerItemLikelyToKeepUpEvent interface {
	IMetricEvent
	

	// properties:
	LoadedTimeRanges() []foundation.Value
	TimeTaken() float64
	Variant() IAVAssetVariant


	

	// methods:


}





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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/loadedTimeRanges-960vi
func (m_ MetricPlayerItemLikelyToKeepUpEvent) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/timeTaken
func (m_ MetricPlayerItemLikelyToKeepUpEvent) TimeTaken() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("timeTaken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/variant
func (m_ MetricPlayerItemLikelyToKeepUpEvent) Variant() IAVAssetVariant {
	rv := objc.Send[AssetVariant](m_.ID, objc.Sel("variant"))
	return rv
}








