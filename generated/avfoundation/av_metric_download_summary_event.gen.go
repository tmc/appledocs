// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricDownloadSummaryEvent */


/* debug [class_header]: Header for AVMetricDownloadSummaryEvent */
// The class instance for the [MetricDownloadSummaryEvent] class.
var (
	MetricDownloadSummaryEventClass     _MetricDownloadSummaryEventClass
	MetricDownloadSummaryEventClassOnce sync.Once
)

func getMetricDownloadSummaryEventClass() _MetricDownloadSummaryEventClass {
	MetricDownloadSummaryEventClassOnce.Do(func() {
		MetricDownloadSummaryEventClass = _MetricDownloadSummaryEventClass{objc.GetClass("AVMetricDownloadSummaryEvent")}
	})
	return MetricDownloadSummaryEventClass
}

type _MetricDownloadSummaryEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricDownloadSummaryEvent */
// An interface definition for the [MetricDownloadSummaryEvent] class.
type IMetricDownloadSummaryEvent interface {
	IMetricEvent
	
/* debug [class_interface_properties]: Properties for MetricDownloadSummaryEvent */
	// properties:
	BytesDownloadedCount() int
	DownloadDuration() float64
	ErrorEvent() IAVMetricErrorEvent
	MediaResourceRequestCount() int
	RecoverableErrorCount() int
	Variants() []AssetVariant
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricDownloadSummaryEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricDownloadSummaryEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricDownloadSummaryEventClass) Alloc() MetricDownloadSummaryEvent {
	rv := objc.Send[MetricDownloadSummaryEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricDownloadSummaryEventClass) New() MetricDownloadSummaryEvent {
	rv := objc.Send[MetricDownloadSummaryEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricDownloadSummaryEvent) Init() MetricDownloadSummaryEvent {
	rv := objc.Send[MetricDownloadSummaryEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricDownloadSummaryEvent) Autorelease() MetricDownloadSummaryEvent {
	rv := objc.Send[MetricDownloadSummaryEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricDownloadSummaryEvent creates a new MetricDownloadSummaryEvent instance.
func NewMetricDownloadSummaryEvent() MetricDownloadSummaryEvent {
	return getMetricDownloadSummaryEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricDownloadSummaryEvent */
// Represents a summary metric event with aggregated metrics for the entire download task.
//
// Subclasses of this type that are used from Swift must fulfill the requirements of a Sendable type.


// Represents a summary metric event with aggregated metrics for the entire download task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent
type MetricDownloadSummaryEvent struct {
	MetricEvent
}

// MetricDownloadSummaryEventFrom constructs a [MetricDownloadSummaryEvent] from an unsafe.Pointer.
//
// Represents a summary metric event with aggregated metrics for the entire download task.
func MetricDownloadSummaryEventFrom(ptr unsafe.Pointer) MetricDownloadSummaryEvent {
	return MetricDownloadSummaryEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricDownloadSummaryEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricDownloadSummaryEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricDownloadSummaryEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricDownloadSummaryEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricDownloadSummaryEvent */

// Returns the total number of bytes downloaded by the download task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/bytesDownloadedCount
func (m_ MetricDownloadSummaryEvent) BytesDownloadedCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("bytesDownloadedCount"))
	return rv
}/* debug [instance_properties/getter]: bytesDownloadedCount */


// Returns the total duration of the download in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/downloadDuration
func (m_ MetricDownloadSummaryEvent) DownloadDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("downloadDuration"))
	return rv
}/* debug [instance_properties/getter]: downloadDuration */


// Returns the error event if any. If no value is available, returns nil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/errorEvent
func (m_ MetricDownloadSummaryEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("errorEvent"))
	return rv
}/* debug [instance_properties/getter]: errorEvent */


// Returns the total number of media requests performed by the download task. This includes playlist requests, media segment requests, and content key requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/mediaResourceRequestCount
func (m_ MetricDownloadSummaryEvent) MediaResourceRequestCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("mediaResourceRequestCount"))
	return rv
}/* debug [instance_properties/getter]: mediaResourceRequestCount */


// Returns the total count of recoverable errors encountered during the download. If no errors were encountered, returns 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/recoverableErrorCount
func (m_ MetricDownloadSummaryEvent) RecoverableErrorCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("recoverableErrorCount"))
	return rv
}/* debug [instance_properties/getter]: recoverableErrorCount */


// Returns the variants that were downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/variants
func (m_ MetricDownloadSummaryEvent) Variants() []AssetVariant {
	rv := objc.Send[[]AssetVariant](m_.ID, objc.Sel("variants"))
	return rv
}/* debug [instance_properties/getter]: variants */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricDownloadSummaryEvent */



