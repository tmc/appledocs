// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class AVMetricHLSMediaSegmentRequestEvent */


/* debug [class_header]: Header for AVMetricHLSMediaSegmentRequestEvent */
// The class instance for the [MetricHLSMediaSegmentRequestEvent] class.
var (
	MetricHLSMediaSegmentRequestEventClass     _MetricHLSMediaSegmentRequestEventClass
	MetricHLSMediaSegmentRequestEventClassOnce sync.Once
)

func getMetricHLSMediaSegmentRequestEventClass() _MetricHLSMediaSegmentRequestEventClass {
	MetricHLSMediaSegmentRequestEventClassOnce.Do(func() {
		MetricHLSMediaSegmentRequestEventClass = _MetricHLSMediaSegmentRequestEventClass{objc.GetClass("AVMetricHLSMediaSegmentRequestEvent")}
	})
	return MetricHLSMediaSegmentRequestEventClass
}

type _MetricHLSMediaSegmentRequestEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricHLSMediaSegmentRequestEvent */
// An interface definition for the [MetricHLSMediaSegmentRequestEvent] class.
type IMetricHLSMediaSegmentRequestEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricHLSMediaSegmentRequestEvent */
	// properties:
	ByteRange() corefoundation.Range
	IndexFileURL() objc.IObject /* cross-framework: NSURL */
	IsMapSegment() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() MediaType /* typedef */
	SegmentDuration() float64
	Url() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricHLSMediaSegmentRequestEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricHLSMediaSegmentRequestEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricHLSMediaSegmentRequestEventClass) Alloc() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricHLSMediaSegmentRequestEventClass) New() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricHLSMediaSegmentRequestEvent) Init() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricHLSMediaSegmentRequestEvent) Autorelease() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricHLSMediaSegmentRequestEvent creates a new MetricHLSMediaSegmentRequestEvent instance.
func NewMetricHLSMediaSegmentRequestEvent() MetricHLSMediaSegmentRequestEvent {
	return getMetricHLSMediaSegmentRequestEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricHLSMediaSegmentRequestEvent */
// An event that represents a live streaming media segment resource request.


// An event that represents a live streaming media segment resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent
type MetricHLSMediaSegmentRequestEvent struct {
	MetricEvent
}

// MetricHLSMediaSegmentRequestEventFrom constructs a [MetricHLSMediaSegmentRequestEvent] from an unsafe.Pointer.
//
// An event that represents a live streaming media segment resource request.
func MetricHLSMediaSegmentRequestEventFrom(ptr unsafe.Pointer) MetricHLSMediaSegmentRequestEvent {
	return MetricHLSMediaSegmentRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricHLSMediaSegmentRequestEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricHLSMediaSegmentRequestEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricHLSMediaSegmentRequestEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricHLSMediaSegmentRequestEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricHLSMediaSegmentRequestEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/byteRange
func (m_ MetricHLSMediaSegmentRequestEvent) ByteRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](m_.ID, objc.Sel("byteRange"))
	return rv
}/* debug [instance_properties/getter]: byteRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/indexFileURL
func (m_ MetricHLSMediaSegmentRequestEvent) IndexFileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("indexFileURL"))
	return rv
}/* debug [instance_properties/getter]: indexFileURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/isMapSegment
func (m_ MetricHLSMediaSegmentRequestEvent) IsMapSegment() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMapSegment"))
	return rv
}/* debug [instance_properties/getter]: isMapSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaResourceRequestEvent
func (m_ MetricHLSMediaSegmentRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("mediaResourceRequestEvent"))
	return rv
}/* debug [instance_properties/getter]: mediaResourceRequestEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaType
func (m_ MetricHLSMediaSegmentRequestEvent) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// Returns the duration of segment in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/segmentDuration
func (m_ MetricHLSMediaSegmentRequestEvent) SegmentDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("segmentDuration"))
	return rv
}/* debug [instance_properties/getter]: segmentDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/url
func (m_ MetricHLSMediaSegmentRequestEvent) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricHLSMediaSegmentRequestEvent */



