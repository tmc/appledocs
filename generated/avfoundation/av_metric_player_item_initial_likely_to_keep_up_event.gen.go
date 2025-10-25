// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemInitialLikelyToKeepUpEvent */


/* debug [class_header]: Header for AVMetricPlayerItemInitialLikelyToKeepUpEvent */
// The class instance for the [MetricPlayerItemInitialLikelyToKeepUpEvent] class.
var (
	MetricPlayerItemInitialLikelyToKeepUpEventClass     _MetricPlayerItemInitialLikelyToKeepUpEventClass
	MetricPlayerItemInitialLikelyToKeepUpEventClassOnce sync.Once
)

func getMetricPlayerItemInitialLikelyToKeepUpEventClass() _MetricPlayerItemInitialLikelyToKeepUpEventClass {
	MetricPlayerItemInitialLikelyToKeepUpEventClassOnce.Do(func() {
		MetricPlayerItemInitialLikelyToKeepUpEventClass = _MetricPlayerItemInitialLikelyToKeepUpEventClass{objc.GetClass("AVMetricPlayerItemInitialLikelyToKeepUpEvent")}
	})
	return MetricPlayerItemInitialLikelyToKeepUpEventClass
}

type _MetricPlayerItemInitialLikelyToKeepUpEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemInitialLikelyToKeepUpEvent */
// An interface definition for the [MetricPlayerItemInitialLikelyToKeepUpEvent] class.
type IMetricPlayerItemInitialLikelyToKeepUpEvent interface {
	IMetricPlayerItemLikelyToKeepUpEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemInitialLikelyToKeepUpEvent */
	// properties:
	ContentKeyRequestEvents() []MetricContentKeyRequestEvent
	MediaSegmentRequestEvents() []MetricHLSMediaSegmentRequestEvent
	PlaylistRequestEvents() []MetricHLSPlaylistRequestEvent
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemInitialLikelyToKeepUpEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemInitialLikelyToKeepUpEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemInitialLikelyToKeepUpEventClass) Alloc() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemInitialLikelyToKeepUpEventClass) New() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) Init() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) Autorelease() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemInitialLikelyToKeepUpEvent creates a new MetricPlayerItemInitialLikelyToKeepUpEvent instance.
func NewMetricPlayerItemInitialLikelyToKeepUpEvent() MetricPlayerItemInitialLikelyToKeepUpEvent {
	return getMetricPlayerItemInitialLikelyToKeepUpEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemInitialLikelyToKeepUpEvent */
// An event that represents the initial state for whether playback is likely to continue without stalling.


// An event that represents the initial state for whether playback is likely to continue without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent
type MetricPlayerItemInitialLikelyToKeepUpEvent struct {
	MetricPlayerItemLikelyToKeepUpEvent
}

// MetricPlayerItemInitialLikelyToKeepUpEventFrom constructs a [MetricPlayerItemInitialLikelyToKeepUpEvent] from an unsafe.Pointer.
//
// An event that represents the initial state for whether playback is likely to continue without stalling.
func MetricPlayerItemInitialLikelyToKeepUpEventFrom(ptr unsafe.Pointer) MetricPlayerItemInitialLikelyToKeepUpEvent {
	return MetricPlayerItemInitialLikelyToKeepUpEvent{
		MetricPlayerItemLikelyToKeepUpEvent: MetricPlayerItemLikelyToKeepUpEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemInitialLikelyToKeepUpEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemInitialLikelyToKeepUpEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemInitialLikelyToKeepUpEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemInitialLikelyToKeepUpEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemInitialLikelyToKeepUpEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/contentKeyRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) ContentKeyRequestEvents() []MetricContentKeyRequestEvent {
	rv := objc.Send[[]MetricContentKeyRequestEvent](m_.ID, objc.Sel("contentKeyRequestEvents"))
	return rv
}/* debug [instance_properties/getter]: contentKeyRequestEvents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/mediaSegmentRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) MediaSegmentRequestEvents() []MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[[]MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("mediaSegmentRequestEvents"))
	return rv
}/* debug [instance_properties/getter]: mediaSegmentRequestEvents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/playlistRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) PlaylistRequestEvents() []MetricHLSPlaylistRequestEvent {
	rv := objc.Send[[]MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("playlistRequestEvents"))
	return rv
}/* debug [instance_properties/getter]: playlistRequestEvents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemInitialLikelyToKeepUpEvent */



