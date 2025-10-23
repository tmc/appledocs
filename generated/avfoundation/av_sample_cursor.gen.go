// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	CurrentChunkInfo() unsafe.Pointer
	SetCurrentChunkInfo(value unsafe.Pointer)
	CurrentChunkStorageRange() unsafe.Pointer
	SetCurrentChunkStorageRange(value unsafe.Pointer)
	CurrentChunkStorageURL() foundation.URL
	SetCurrentChunkStorageURL(value foundation.IURL)
	CurrentSampleAudioDependencyInfo() unsafe.Pointer
	SetCurrentSampleAudioDependencyInfo(value unsafe.Pointer)
	CurrentSampleDependencyAttachments() unsafe.Pointer
	SetCurrentSampleDependencyAttachments(value unsafe.Pointer)
	CurrentSampleDependencyInfo() unsafe.Pointer
	SetCurrentSampleDependencyInfo(value unsafe.Pointer)
	CurrentSampleDuration() unsafe.Pointer
	SetCurrentSampleDuration(value unsafe.Pointer)
	CurrentSampleIndexInChunk() unsafe.Pointer
	SetCurrentSampleIndexInChunk(value unsafe.Pointer)
	CurrentSampleStorageRange() unsafe.Pointer
	SetCurrentSampleStorageRange(value unsafe.Pointer)
	CurrentSampleSyncInfo() unsafe.Pointer
	SetCurrentSampleSyncInfo(value unsafe.Pointer)
	DecodeTimeStamp() unsafe.Pointer
	SetDecodeTimeStamp(value unsafe.Pointer)
	PresentationTimeStamp() unsafe.Pointer
	SetPresentationTimeStamp(value unsafe.Pointer)
	SamplesRequiredForDecoderRefresh() int
	SetSamplesRequiredForDecoderRefresh(value int)
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

// Alloc allocates a new instance without initialization.
func (sc _SampleCursorClass) Alloc() SampleCursor {
	rv := objc.Send[SampleCursor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A value that provides information about the chunk of samples to which the current sample belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkinfo
func (s_ SampleCursor) CurrentChunkInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentChunkInfo"))
	return rv
}


// A value that provides information about the chunk of samples to which the current sample belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkinfo
func (s_ SampleCursor) SetCurrentChunkInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentChunkInfo:"), value)
}


// The sample range in the storage container to load together with the current sample as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkstoragerange
func (s_ SampleCursor) CurrentChunkStorageRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentChunkStorageRange"))
	return rv
}


// The sample range in the storage container to load together with the current sample as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkstoragerange
func (s_ SampleCursor) SetCurrentChunkStorageRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentChunkStorageRange:"), value)
}


// The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkstorageurl
func (s_ SampleCursor) CurrentChunkStorageURL() foundation.URL {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("currentChunkStorageURL"))
	return rv
}


// The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentchunkstorageurl
func (s_ SampleCursor) SetCurrentChunkStorageURL(value foundation.IURL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentChunkStorageURL:"), value)
}


// The independent decodability information for the audio sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleaudiodependencyinfo
func (s_ SampleCursor) CurrentSampleAudioDependencyInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleAudioDependencyInfo"))
	return rv
}


// The independent decodability information for the audio sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleaudiodependencyinfo
func (s_ SampleCursor) SetCurrentSampleAudioDependencyInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleAudioDependencyInfo:"), value)
}


// A dictionary of dependency-related sample buffer attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampledependencyattachments
func (s_ SampleCursor) CurrentSampleDependencyAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleDependencyAttachments"))
	return rv
}


// A dictionary of dependency-related sample buffer attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampledependencyattachments
func (s_ SampleCursor) SetCurrentSampleDependencyAttachments(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleDependencyAttachments:"), value)
}


// The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampledependencyinfo
func (s_ SampleCursor) CurrentSampleDependencyInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleDependencyInfo"))
	return rv
}


// The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampledependencyinfo
func (s_ SampleCursor) SetCurrentSampleDependencyInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleDependencyInfo:"), value)
}


// The decode duration of the sample at the cursor’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleduration
func (s_ SampleCursor) CurrentSampleDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleDuration"))
	return rv
}


// The decode duration of the sample at the cursor’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleduration
func (s_ SampleCursor) SetCurrentSampleDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleDuration:"), value)
}


// The index of the current sample within the chunk to which it belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleindexinchunk
func (s_ SampleCursor) CurrentSampleIndexInChunk() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleIndexInChunk"))
	return rv
}


// The index of the current sample within the chunk to which it belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsampleindexinchunk
func (s_ SampleCursor) SetCurrentSampleIndexInChunk(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleIndexInChunk:"), value)
}


// The offset and length of the current sample in the current chunk storage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsamplestoragerange
func (s_ SampleCursor) CurrentSampleStorageRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleStorageRange"))
	return rv
}


// The offset and length of the current sample in the current chunk storage URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsamplestoragerange
func (s_ SampleCursor) SetCurrentSampleStorageRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleStorageRange:"), value)
}


// The synchronization information for the current sample for consideration when resynchronizing a decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsamplesyncinfo
func (s_ SampleCursor) CurrentSampleSyncInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentSampleSyncInfo"))
	return rv
}


// The synchronization information for the current sample for consideration when resynchronizing a decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/currentsamplesyncinfo
func (s_ SampleCursor) SetCurrentSampleSyncInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentSampleSyncInfo:"), value)
}


// The decode timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/decodetimestamp
func (s_ SampleCursor) DecodeTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("decodeTimeStamp"))
	return rv
}


// The decode timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/decodetimestamp
func (s_ SampleCursor) SetDecodeTimeStamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDecodeTimeStamp:"), value)
}


// The presentation timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/presentationtimestamp
func (s_ SampleCursor) PresentationTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("presentationTimeStamp"))
	return rv
}


// The presentation timestamp of the sample at the current position of the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/presentationtimestamp
func (s_ SampleCursor) SetPresentationTimeStamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresentationTimeStamp:"), value)
}


// The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/samplesrequiredfordecoderrefresh
func (s_ SampleCursor) SamplesRequiredForDecoderRefresh() int {
	rv := objc.Send[int](s_.ID, objc.Sel("samplesRequiredForDecoderRefresh"))
	return rv
}


// The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplecursor/samplesrequiredfordecoderrefresh
func (s_ SampleCursor) SetSamplesRequiredForDecoderRefresh(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSamplesRequiredForDecoderRefresh:"), value)
}



