// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemPlaybackSummaryEvent */


/* debug [class_header]: Header for AVMetricPlayerItemPlaybackSummaryEvent */
// The class instance for the [MetricPlayerItemPlaybackSummaryEvent] class.
var (
	MetricPlayerItemPlaybackSummaryEventClass     _MetricPlayerItemPlaybackSummaryEventClass
	MetricPlayerItemPlaybackSummaryEventClassOnce sync.Once
)

func getMetricPlayerItemPlaybackSummaryEventClass() _MetricPlayerItemPlaybackSummaryEventClass {
	MetricPlayerItemPlaybackSummaryEventClassOnce.Do(func() {
		MetricPlayerItemPlaybackSummaryEventClass = _MetricPlayerItemPlaybackSummaryEventClass{objc.GetClass("AVMetricPlayerItemPlaybackSummaryEvent")}
	})
	return MetricPlayerItemPlaybackSummaryEventClass
}

type _MetricPlayerItemPlaybackSummaryEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemPlaybackSummaryEvent */
// An interface definition for the [MetricPlayerItemPlaybackSummaryEvent] class.
type IMetricPlayerItemPlaybackSummaryEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemPlaybackSummaryEvent */
	// properties:
	ErrorEvent() IAVMetricErrorEvent
	MediaResourceRequestCount() int
	PlaybackDuration() int
	RecoverableErrorCount() int
	StallCount() int
	TimeSpentInInitialStartup() float64
	TimeSpentRecoveringFromStall() float64
	TimeWeightedAverageBitrate() int
	TimeWeightedPeakBitrate() int
	VariantSwitchCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemPlaybackSummaryEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemPlaybackSummaryEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemPlaybackSummaryEventClass) Alloc() MetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[MetricPlayerItemPlaybackSummaryEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemPlaybackSummaryEventClass) New() MetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[MetricPlayerItemPlaybackSummaryEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemPlaybackSummaryEvent) Init() MetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[MetricPlayerItemPlaybackSummaryEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemPlaybackSummaryEvent) Autorelease() MetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[MetricPlayerItemPlaybackSummaryEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemPlaybackSummaryEvent creates a new MetricPlayerItemPlaybackSummaryEvent instance.
func NewMetricPlayerItemPlaybackSummaryEvent() MetricPlayerItemPlaybackSummaryEvent {
	return getMetricPlayerItemPlaybackSummaryEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemPlaybackSummaryEvent */
// An event that represents the combined metrics for the entire playback session.


// An event that represents the combined metrics for the entire playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent
type MetricPlayerItemPlaybackSummaryEvent struct {
	MetricEvent
}

// MetricPlayerItemPlaybackSummaryEventFrom constructs a [MetricPlayerItemPlaybackSummaryEvent] from an unsafe.Pointer.
//
// An event that represents the combined metrics for the entire playback session.
func MetricPlayerItemPlaybackSummaryEventFrom(ptr unsafe.Pointer) MetricPlayerItemPlaybackSummaryEvent {
	return MetricPlayerItemPlaybackSummaryEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemPlaybackSummaryEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemPlaybackSummaryEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemPlaybackSummaryEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemPlaybackSummaryEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemPlaybackSummaryEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/errorEvent
func (m_ MetricPlayerItemPlaybackSummaryEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("errorEvent"))
	return rv
}/* debug [instance_properties/getter]: errorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/mediaResourceRequestCount
func (m_ MetricPlayerItemPlaybackSummaryEvent) MediaResourceRequestCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("mediaResourceRequestCount"))
	return rv
}/* debug [instance_properties/getter]: mediaResourceRequestCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/playbackDuration
func (m_ MetricPlayerItemPlaybackSummaryEvent) PlaybackDuration() int {
	rv := objc.Send[int](m_.ID, objc.Sel("playbackDuration"))
	return rv
}/* debug [instance_properties/getter]: playbackDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/recoverableErrorCount
func (m_ MetricPlayerItemPlaybackSummaryEvent) RecoverableErrorCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("recoverableErrorCount"))
	return rv
}/* debug [instance_properties/getter]: recoverableErrorCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/stallCount
func (m_ MetricPlayerItemPlaybackSummaryEvent) StallCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("stallCount"))
	return rv
}/* debug [instance_properties/getter]: stallCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeSpentInInitialStartup
func (m_ MetricPlayerItemPlaybackSummaryEvent) TimeSpentInInitialStartup() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("timeSpentInInitialStartup"))
	return rv
}/* debug [instance_properties/getter]: timeSpentInInitialStartup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeSpentRecoveringFromStall
func (m_ MetricPlayerItemPlaybackSummaryEvent) TimeSpentRecoveringFromStall() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("timeSpentRecoveringFromStall"))
	return rv
}/* debug [instance_properties/getter]: timeSpentRecoveringFromStall */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeWeightedAverageBitrate
func (m_ MetricPlayerItemPlaybackSummaryEvent) TimeWeightedAverageBitrate() int {
	rv := objc.Send[int](m_.ID, objc.Sel("timeWeightedAverageBitrate"))
	return rv
}/* debug [instance_properties/getter]: timeWeightedAverageBitrate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeWeightedPeakBitrate
func (m_ MetricPlayerItemPlaybackSummaryEvent) TimeWeightedPeakBitrate() int {
	rv := objc.Send[int](m_.ID, objc.Sel("timeWeightedPeakBitrate"))
	return rv
}/* debug [instance_properties/getter]: timeWeightedPeakBitrate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/variantSwitchCount
func (m_ MetricPlayerItemPlaybackSummaryEvent) VariantSwitchCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("variantSwitchCount"))
	return rv
}/* debug [instance_properties/getter]: variantSwitchCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemPlaybackSummaryEvent */



