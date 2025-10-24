// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetricDownloadSummaryEvent] class.
type IMetricDownloadSummaryEvent interface {
	IMetricEvent
	

	// properties:
	BytesDownloadedCount() int
	DownloadDuration() float64
	ErrorEvent() IAVMetricErrorEvent
	MediaResourceRequestCount() int
	RecoverableErrorCount() int
	Variants() []AssetVariant


	

	// methods:


}





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

























// Returns the total number of bytes downloaded by the download task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/bytesDownloadedCount
func (m_ MetricDownloadSummaryEvent) BytesDownloadedCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("bytesDownloadedCount"))
	return rv
}


// Returns the total duration of the download in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/downloadDuration
func (m_ MetricDownloadSummaryEvent) DownloadDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("downloadDuration"))
	return rv
}


// Returns the error event if any. If no value is available, returns nil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/errorEvent
func (m_ MetricDownloadSummaryEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[MetricErrorEvent](m_.ID, objc.Sel("errorEvent"))
	return rv
}


// Returns the total number of media requests performed by the download task. This includes playlist requests, media segment requests, and content key requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/mediaResourceRequestCount
func (m_ MetricDownloadSummaryEvent) MediaResourceRequestCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("mediaResourceRequestCount"))
	return rv
}


// Returns the total count of recoverable errors encountered during the download. If no errors were encountered, returns 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/recoverableErrorCount
func (m_ MetricDownloadSummaryEvent) RecoverableErrorCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("recoverableErrorCount"))
	return rv
}


// Returns the variants that were downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/variants
func (m_ MetricDownloadSummaryEvent) Variants() []AssetVariant {
	rv := objc.Send[[]AssetVariant](m_.ID, objc.Sel("variants"))
	return rv
}








