// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricHLSPlaylistRequestEvent */


/* debug [class_header]: Header for AVMetricHLSPlaylistRequestEvent */
// The class instance for the [MetricHLSPlaylistRequestEvent] class.
var (
	MetricHLSPlaylistRequestEventClass     _MetricHLSPlaylistRequestEventClass
	MetricHLSPlaylistRequestEventClassOnce sync.Once
)

func getMetricHLSPlaylistRequestEventClass() _MetricHLSPlaylistRequestEventClass {
	MetricHLSPlaylistRequestEventClassOnce.Do(func() {
		MetricHLSPlaylistRequestEventClass = _MetricHLSPlaylistRequestEventClass{objc.GetClass("AVMetricHLSPlaylistRequestEvent")}
	})
	return MetricHLSPlaylistRequestEventClass
}

type _MetricHLSPlaylistRequestEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricHLSPlaylistRequestEvent */
// An interface definition for the [MetricHLSPlaylistRequestEvent] class.
type IMetricHLSPlaylistRequestEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricHLSPlaylistRequestEvent */
	// properties:
	IsMultivariantPlaylist() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() MediaType /* typedef */
	Url() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricHLSPlaylistRequestEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricHLSPlaylistRequestEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricHLSPlaylistRequestEventClass) Alloc() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricHLSPlaylistRequestEventClass) New() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricHLSPlaylistRequestEvent) Init() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricHLSPlaylistRequestEvent) Autorelease() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricHLSPlaylistRequestEvent creates a new MetricHLSPlaylistRequestEvent instance.
func NewMetricHLSPlaylistRequestEvent() MetricHLSPlaylistRequestEvent {
	return getMetricHLSPlaylistRequestEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricHLSPlaylistRequestEvent */
// An event that represents a live streaming playlist resource request.


// An event that represents a live streaming playlist resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent
type MetricHLSPlaylistRequestEvent struct {
	MetricEvent
}

// MetricHLSPlaylistRequestEventFrom constructs a [MetricHLSPlaylistRequestEvent] from an unsafe.Pointer.
//
// An event that represents a live streaming playlist resource request.
func MetricHLSPlaylistRequestEventFrom(ptr unsafe.Pointer) MetricHLSPlaylistRequestEvent {
	return MetricHLSPlaylistRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricHLSPlaylistRequestEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricHLSPlaylistRequestEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricHLSPlaylistRequestEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricHLSPlaylistRequestEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricHLSPlaylistRequestEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/isMultivariantPlaylist
func (m_ MetricHLSPlaylistRequestEvent) IsMultivariantPlaylist() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMultivariantPlaylist"))
	return rv
}/* debug [instance_properties/getter]: isMultivariantPlaylist */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaResourceRequestEvent
func (m_ MetricHLSPlaylistRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("mediaResourceRequestEvent"))
	return rv
}/* debug [instance_properties/getter]: mediaResourceRequestEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaType
func (m_ MetricHLSPlaylistRequestEvent) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/url
func (m_ MetricHLSPlaylistRequestEvent) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricHLSPlaylistRequestEvent */



