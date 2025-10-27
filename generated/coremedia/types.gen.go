// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia
import (
	"unsafe"
)


// C struct types
// CMBlockBufferCustomBlockSource - A structure to support custom memory allocation and deallocation for a block used in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCustomBlockSource
type CMBlockBufferCustomBlockSource struct {
	AllocateBlock unsafe.Pointer // The function to allocate memory.
	FreeBlock unsafe.Pointer // A function to call once when the   is disposed.
	RefCon unsafe.Pointer // Contextual information passed to both the allocate and free function calls.
	Version uint32
}

// CMBufferCallbacks - A structure that stores the callbacks that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferCallbacks
type CMBufferCallbacks struct {
	Compare BufferCompareCallback // This callback is called multiple times from  , to perform an insertion sort.
	DataBecameReadyNotification StringRef // If triggers of type   are installed, the queue will listen for this notification on the head buffer.
	GetDecodeTimeStamp BufferGetTimeCallback // Client callback that returns a   from a  .
	GetDuration BufferGetTimeCallback // This callback is called (once) during enqueue and dequeue operations to update the total duration of the queue.
	GetPresentationTimeStamp BufferGetTimeCallback // Client callback that returns a   from a  .
	GetSize BufferGetSizeCallback
	IsDataReady BufferGetBooleanCallback // This callback is called from  , to ask if the buffer that is about to be dequeued is ready.
	Refcon unsafe.Pointer // Contextual data to be passed to all callbacks.
	Version uint32 // The callback version.
}

// CMBufferHandlers - A structure that stores the handlers that perform buffer operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferHandlers
type CMBufferHandlers struct {
	Compare BufferCompareHandler // A handler callback the queue uses to perform an insertion sort of the queue.
	DataBecameReadyNotification StringRef
	GetDecodeTimeStamp BufferGetTimeHandler
	GetDuration BufferGetTimeHandler
	GetPresentationTimeStamp BufferGetTimeHandler
	GetSize BufferGetSizeHandler
	IsDataReady BufferGetBooleanHandler
	Version unsafe.Pointer // The version number.
}

// CMSampleTimingInfo - A collection of timing information for a sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleTimingInfo
type CMSampleTimingInfo struct {
	DecodeTimeStamp Time // The time at which the sample will be decoded.
	Duration Time // The duration of the sample.
	PresentationTimeStamp Time // The time at which the sample will be presented.
}

// CMTag - A tag representing additional metadata on tagged media buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTag-c.struct
type CMTag struct {
	Category TagCategory // The category assigned to a tag.
	DataType TagDataType // The data type for the value stored in the tag.
	Value TagValue // The value of the tag.
}

// CMTime - A structure that represents time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTime
type CMTime struct {
	Epoch TimeEpoch // The epoch of the time.
	Flags TimeFlags // The flags associated with a time.
	Timescale TimeScale // A timescale that represents the denominator of a rational time.
	Value TimeValue // A time value that represents the numerator of a rational time.
}

// CMTimeMapping - A structure that maps a segment of a source time range to a target time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapping
type CMTimeMapping struct {
	Source TimeRange // A time range on the source timeline.
	Target TimeRange // A time range on the target timeline.
}

// CMTimeRange - A structure that represents a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
type CMTimeRange struct {
	Duration Time // The duration of the time range.
	Start Time // The start time of the time range.
}

// CMVideoDimensions - A structure that represents video dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoDimensions
type CMVideoDimensions struct {
	Height int32 // The height of the video.
	Width int32 // The width of the video.
}





