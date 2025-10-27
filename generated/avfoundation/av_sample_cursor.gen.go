// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SampleCursor] class.
var (
	SampleCursorClass     _SampleCursorClass
	SampleCursorClassOnce sync.Once
)

func getSampleCursorClass() _SampleCursorClass {
	SampleCursorClassOnce.Do(func() {
		SampleCursorClass = _SampleCursorClass{objc.GetClass("AVSampleCursor")}
	})
	return SampleCursorClass
}

type _SampleCursorClass struct {
	class objc.Class
}





// An interface definition for the [SampleCursor] class.
type ISampleCursor interface {
	objectivec.IObject
	

	// properties:
	CurrentChunkInfo() AVSampleCursorChunkInfo
	CurrentChunkStorageRange() AVSampleCursorStorageRange
	CurrentChunkStorageURL() foundation.foundation.INSURL
	CurrentSampleAudioDependencyInfo() AVSampleCursorAudioDependencyInfo
	CurrentSampleDependencyAttachments() foundation.foundation.INSDictionary
	CurrentSampleDependencyInfo() AVSampleCursorDependencyInfo
	CurrentSampleDuration() objectivec.IObject
	CurrentSampleIndexInChunk() int64
	CurrentSampleStorageRange() AVSampleCursorStorageRange
	CurrentSampleSyncInfo() AVSampleCursorSyncInfo
	DecodeTimeStamp() objectivec.IObject
	PresentationTimeStamp() objectivec.IObject
	SamplesRequiredForDecoderRefresh() int


	

	// methods:
	ComparePositionInDecodeOrderWithPositionOfCursor(cursor IAVSampleCursor) ComparisonResult /* not a class type */
	CopyCurrentSampleFormatDescription() FormatDescriptionRef /* not a class type */
	SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool
	SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool
	StepByDecodeTimeWasPinned(deltaDecodeTime objectivec.IObject, outWasPinned objectivec.IObject) objectivec.IObject
	StepByPresentationTimeWasPinned(deltaPresentationTime objectivec.IObject, outWasPinned objectivec.IObject) objectivec.IObject
	StepInDecodeOrderByCount(stepCount int64) int64
	StepInPresentationOrderByCount(stepCount int64) int64


}





// Alloc allocates a new instance without initialization.
func (sc _SampleCursorClass) Alloc() SampleCursor {
	rv := objc.Send[SampleCursor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleCursorClass) New() SampleCursor {
	rv := objc.Send[SampleCursor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleCursor) Init() SampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleCursor) Autorelease() SampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleCursor creates a new SampleCursor instance.
func NewSampleCursor() SampleCursor {
	return getSampleCursorClass().New()
}





// An object that provides information about the media sample at the cursor’s current position.
//
// You position a sample cursor at a specific media sample in a sequence of samples contained in a higher-level object, like an . You can move it to a new position in that sequence either backwards or forwards, either in decode order or in presentation order. You can also request moving it according to a count of samples or a delta in time. Use a sample cursor to get information about the media sample such as its duration, timestamps, dependency information, and so on. You can also use them to synchronously to perform I/O in order to load media data of one or more media samples into memory.


// An object that provides information about the media sample at the cursor’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor
type SampleCursor struct {
	objectivec.Object
}

// SampleCursorFrom constructs a [SampleCursor] from an unsafe.Pointer.
//
// An object that provides information about the media sample at the cursor’s current position.
func SampleCursorFrom(ptr unsafe.Pointer) SampleCursor {
	return SampleCursor{objectivec.Object{objc.ID(ptr)}}
}




















// Compares the relative positions of two sample cursors and returns their relative positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/comparePositionInDecodeOrder(withPositionOf:)
func (s_ SampleCursor) ComparePositionInDecodeOrderWithPositionOfCursor(cursor IAVSampleCursor) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("comparePositionInDecodeOrderWithPositionOfCursor:"), cursor)
	return rv
}


// Returns the format description of the sample at the cursor’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/copyCurrentSampleFormatDescription()
func (s_ SampleCursor) CopyCurrentSampleFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](s_.ID, objc.Sel("copyCurrentSampleFormatDescription"))
	return rv
}


// Determines whether a sample earlier in decode order can have a presentation timestamp later than that of the specified sample cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/maySamplesWithEarlierDecodeTimeStampsHavePresentationTimeStamps(laterThan:)
func (s_ SampleCursor) SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("samplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor:"), cursor)
	return rv
}


// Determines whether a sample later in decode order can have a presentation timestamp earlier than that of the specified sample cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/maySamplesWithLaterDecodeTimeStampsHavePresentationTimeStamps(earlierThan:)
func (s_ SampleCursor) SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("samplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor:"), cursor)
	return rv
}


// Moves the cursor by a given delta time on the decode timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/step(byDecodeTime:wasPinned:)
func (s_ SampleCursor) StepByDecodeTimeWasPinned(deltaDecodeTime objectivec.IObject, outWasPinned objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("stepByDecodeTime:wasPinned:"), deltaDecodeTime, outWasPinned)
	return rv
}


// Moves the cursor by a given delta time on the presentation timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/step(byPresentationTime:wasPinned:)
func (s_ SampleCursor) StepByPresentationTimeWasPinned(deltaPresentationTime objectivec.IObject, outWasPinned objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("stepByPresentationTime:wasPinned:"), deltaPresentationTime, outWasPinned)
	return rv
}


// Moves the cursor a given number of samples in decode order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/stepInDecodeOrder(byCount:)
func (s_ SampleCursor) StepInDecodeOrderByCount(stepCount int64) int64 {
	rv := objc.Send[int64](s_.ID, objc.Sel("stepInDecodeOrderByCount:"), stepCount)
	return rv
}


// Moves the cursor a given number of samples in presentation order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/stepInPresentationOrder(byCount:)
func (s_ SampleCursor) StepInPresentationOrderByCount(stepCount int64) int64 {
	rv := objc.Send[int64](s_.ID, objc.Sel("stepInPresentationOrderByCount:"), stepCount)
	return rv
}







// A value that provides information about the chunk of samples to which the current sample belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkInfo
func (s_ SampleCursor) CurrentChunkInfo() AVSampleCursorChunkInfo {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentChunkInfo"))
	return rv
}


// The sample range in the storage container to load together with the current sample as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkStorageRange
func (s_ SampleCursor) CurrentChunkStorageRange() AVSampleCursorStorageRange {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentChunkStorageRange"))
	return rv
}


// The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkStorageURL
func (s_ SampleCursor) CurrentChunkStorageURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("currentChunkStorageURL"))
	return rv
}


// The independent decodability information for the audio sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleAudioDependencyInfo
func (s_ SampleCursor) CurrentSampleAudioDependencyInfo() AVSampleCursorAudioDependencyInfo {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentSampleAudioDependencyInfo"))
	return rv
}


// A dictionary of dependency-related sample buffer attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDependencyAttachments
func (s_ SampleCursor) CurrentSampleDependencyAttachments() foundation.foundation.INSDictionary {
	rv := objc.Send[foundation.NSDictionary](s_.ID, objc.Sel("currentSampleDependencyAttachments"))
	return rv
}


// The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDependencyInfo
func (s_ SampleCursor) CurrentSampleDependencyInfo() AVSampleCursorDependencyInfo {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentSampleDependencyInfo"))
	return rv
}


// The decode duration of the sample at the cursor’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDuration
func (s_ SampleCursor) CurrentSampleDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("currentSampleDuration"))
	return rv
}


// The index of the current sample within the chunk to which it belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleIndexInChunk
func (s_ SampleCursor) CurrentSampleIndexInChunk() int64 {
	rv := objc.Send[int64](s_.ID, objc.Sel("currentSampleIndexInChunk"))
	return rv
}


// The offset and length of the current sample in the current chunk storage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleStorageRange
func (s_ SampleCursor) CurrentSampleStorageRange() AVSampleCursorStorageRange {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentSampleStorageRange"))
	return rv
}


// The synchronization information for the current sample for consideration when resynchronizing a decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleSyncInfo
func (s_ SampleCursor) CurrentSampleSyncInfo() AVSampleCursorSyncInfo {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("currentSampleSyncInfo"))
	return rv
}


// The decode timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/decodeTimeStamp
func (s_ SampleCursor) DecodeTimeStamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("decodeTimeStamp"))
	return rv
}


// The presentation timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/presentationTimeStamp
func (s_ SampleCursor) PresentationTimeStamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("presentationTimeStamp"))
	return rv
}


// The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/samplesRequiredForDecoderRefresh
func (s_ SampleCursor) SamplesRequiredForDecoderRefresh() int {
	rv := objc.Send[int](s_.ID, objc.Sel("samplesRequiredForDecoderRefresh"))
	return rv
}








