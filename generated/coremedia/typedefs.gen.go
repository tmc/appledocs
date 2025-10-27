// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia
import (
"unsafe"
)

// Type aliases and typedefs
// AttachmentBearerRef - An object that can carry attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAttachmentBearer
// CMAttachmentBearerRef has base type: CFTypeRef
type AttachmentBearerRef uintptr
// AttachmentMode - The mode to use when propagating attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAttachmentMode
// CMAttachmentMode has base type: uint32_t
type AttachmentMode uintptr
// AudioCodecType - An audio codec type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioCodecType
// CMAudioCodecType has base type: FourCharCode
type AudioCodecType uintptr
// AudioFormatDescriptionRef - A type you use to interact with audio format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescription
// CMAudioFormatDescriptionRef has base type: CMFormatDescriptionRef
type AudioFormatDescriptionRef uintptr
// AudioFormatDescriptionMask - A type for mask bits that represent parts of an audio format description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionMask
// CMAudioFormatDescriptionMask has base type: uint32_t
type AudioFormatDescriptionMask uintptr
// BaseClassVersion type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBaseClassVersion
// CMBaseClassVersion has base type: uint32_t
type BaseClassVersion uintptr
// BlockBufferRef - A reference to a block buffer instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBuffer
// CMBlockBufferRef has base type: struct OpaqueCMBlockBuffer *
type BlockBufferRef uintptr
// BlockBufferFlags - A type for flags that control behaviors and features of block buffer APIs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferFlags
// CMBlockBufferFlags has base type: uint32_t
type BlockBufferFlags uintptr
// BufferRef - A reference to a buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBuffer
// CMBufferRef has base type: CFTypeRef
type BufferRef uintptr
// BufferCompareCallback - Callback that compares one   with another.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferCompareCallback
// CMBufferCompareCallback is a callback function
// C type: enum CFComparisonResult (*)(const void *, const void *, void *)
type BufferCompareCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ComparisonResult
// BufferGetBooleanCallback - Callback that returns a Boolean value from a  .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferGetBooleanCallback
// CMBufferGetBooleanCallback is a callback function
// C type: unsigned char (*)(const void *, void *)
type BufferGetBooleanCallback = func(unsafe.Pointer, unsafe.Pointer) uint8
// BufferGetSizeCallback - A client callback that returns a size.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferGetSizeCallback
// CMBufferGetSizeCallback is a callback function
// C type: unsigned long (*)(const void *, void *)
type BufferGetSizeCallback = func(unsafe.Pointer, unsafe.Pointer) uint
// BufferGetTimeCallback - Callback that returns a   from a  .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferGetTimeCallback
// CMBufferGetTimeCallback is a callback function
// C type: CMTime (*)(const void *, void *)
type BufferGetTimeCallback = func(unsafe.Pointer, unsafe.Pointer) CMTime
// BufferQueueRef - A reference to a buffer queue instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueue
// CMBufferQueueRef has base type: struct opaqueCMBufferQueue *
type BufferQueueRef uintptr
// BufferQueueTriggerCallback - A callback for the system to invoke when a trigger condition becomes true.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerCallback
// CMBufferQueueTriggerCallback is a callback function
// C type: void (*)(void *, struct opaqueCMBufferQueueTriggerToken *)
type BufferQueueTriggerCallback = func(unsafe.Pointer, unsafe.Pointer)
// BufferQueueTriggerCondition - A type to specify conditions to associate with a buffer queue trigger.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerCondition
// CMBufferQueueTriggerCondition has base type: int32_t
type BufferQueueTriggerCondition uintptr
// BufferQueueTriggerToken - A type alias for a trigger token.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTriggerToken
// CMBufferQueueTriggerToken has base type: struct opaqueCMBufferQueueTriggerToken *
type BufferQueueTriggerToken uintptr
// BufferValidationCallback - A type alias for a callback that tests whether a buffer is in a valid state to add to a queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferValidationCallback
// CMBufferValidationCallback is a callback function
// C type: int (*)(struct opaqueCMBufferQueue *, const void *, void *)
type BufferValidationCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32
// ClockRef - An object that represents a source of time.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClock
// CMClockRef has base type: struct OpaqueCMClock *
type ClockRef uintptr
// ClockOrTimebaseRef - A type you use in argument lists and function results to indicate that you can pass either a clock or timebase.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockOrTimebase
// CMClockOrTimebaseRef has base type: CFTypeRef
type ClockOrTimebaseRef uintptr
// ClosedCaptionDescriptionFlavor - Types that represent closed caption format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionDescriptionFlavor
// CMClosedCaptionDescriptionFlavor has base type: CFStringRef
type ClosedCaptionDescriptionFlavor uintptr
// ClosedCaptionFormatDescriptionRef - A type you use to interact with closed caption format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescription
// CMClosedCaptionFormatDescriptionRef has base type: CMFormatDescriptionRef
type ClosedCaptionFormatDescriptionRef uintptr
// ClosedCaptionFormatType - A closed caption format type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatType
// CMClosedCaptionFormatType has base type: FourCharCode
type ClosedCaptionFormatType uintptr
// FormatDescriptionRef - An object that describes a media format descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
// CMFormatDescriptionRef has base type: const struct opaqueCMFormatDescription *
type FormatDescriptionRef uintptr
// ImageDescriptionFlavor - Types that represent image format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMImageDescriptionFlavor
// CMImageDescriptionFlavor has base type: CFStringRef
type ImageDescriptionFlavor uintptr
// ItemCount - A datatype that represents an item count.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMItemCount
// CMItemCount has base type: CFIndex
type ItemCount uintptr
// ItemIndex - A datatype that represents an item index.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMItemIndex
// CMItemIndex has base type: CFIndex
type ItemIndex uintptr
// MediaType - Constants that represent media types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMediaType
// CMMediaType has base type: FourCharCode
type MediaType uintptr
// MemoryPoolRef - An instance that optimizes memory allocation when working with large blocks of memory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPool
// CMMemoryPoolRef has base type: struct OpaqueCMMemoryPool *
type MemoryPoolRef uintptr
// MetadataDescriptionFlavor - Types that represent metadata format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDescriptionFlavor
// CMMetadataDescriptionFlavor has base type: CFStringRef
type MetadataDescriptionFlavor uintptr
// MetadataFormatDescriptionRef - A type you use to interact with metadata format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescription
// CMMetadataFormatDescriptionRef has base type: CMFormatDescriptionRef
type MetadataFormatDescriptionRef uintptr
// MetadataFormatType - A metadata format type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatType
// CMMetadataFormatType has base type: FourCharCode
type MetadataFormatType uintptr
// MuxedFormatDescriptionRef - A type you use to interact with muxed format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMuxedFormatDescription
// CMMuxedFormatDescriptionRef has base type: CMFormatDescriptionRef
type MuxedFormatDescriptionRef uintptr
// MuxedStreamType - A datatype that represents a muxed stream of data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMuxedStreamType
// CMMuxedStreamType has base type: FourCharCode
type MuxedStreamType uintptr
// PersistentTrackID - A datatype that represents a persistent track identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPersistentTrackID
// CMPersistentTrackID has base type: int32_t
type PersistentTrackID uintptr
// PixelFormatType - A pixel format type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPixelFormatType
// CMPixelFormatType has base type: FourCharCode
type PixelFormatType uintptr
// SampleBufferRef - A reference to a buffer of media data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
// CMSampleBufferRef has base type: struct opaqueCMSampleBuffer *
type SampleBufferRef uintptr
// SampleBufferInvalidateCallback - Client callback called by  .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidateCallback
// CMSampleBufferInvalidateCallback is a callback function
// C type: void (*)(struct opaqueCMSampleBuffer *, unsigned long long)
type SampleBufferInvalidateCallback = func(unsafe.Pointer, uint64)
// SampleBufferMakeDataReadyCallback - Client callback called by  .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReadyCallback
// CMSampleBufferMakeDataReadyCallback is a callback function
// C type: int (*)(struct opaqueCMSampleBuffer *, void *)
type SampleBufferMakeDataReadyCallback = func(unsafe.Pointer, unsafe.Pointer) int32
// SimpleQueueRef - A reference to an instance that provides a simple lockless queue of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueue
// CMSimpleQueueRef has base type: struct opaqueCMSimpleQueue *
type SimpleQueueRef uintptr
// SoundDescriptionFlavor - Types that represent sound format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSoundDescriptionFlavor
// CMSoundDescriptionFlavor has base type: CFStringRef
type SoundDescriptionFlavor uintptr
// StructVersion type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStructVersion
// CMStructVersion has base type: uintptr_t
type StructVersion uintptr
// SubtitleFormatType - A type that represents a text subtitle format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSubtitleFormatType
// CMSubtitleFormatType has base type: FourCharCode
type SubtitleFormatType uintptr
// TaggedBufferGroupFormatDescriptionRef - A type for tagged buffer format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescription
// CMTaggedBufferGroupFormatDescriptionRef has base type: CMFormatDescriptionRef
type TaggedBufferGroupFormatDescriptionRef uintptr
// TaggedBufferGroupFormatType - A type for tagged buffer format information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatType
// CMTaggedBufferGroupFormatType has base type: FourCharCode
type TaggedBufferGroupFormatType uintptr
// TextDescriptionFlavor - Types that represent text format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextDescriptionFlavor
// CMTextDescriptionFlavor has base type: CFStringRef
type TextDescriptionFlavor uintptr
// TextDisplayFlags - An integer value that describes the display mode flags for text media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextDisplayFlags
// CMTextDisplayFlags has base type: uint32_t
type TextDisplayFlags uintptr
// TextFormatDescriptionRef - A type you use to interact with text format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescription
// CMTextFormatDescriptionRef has base type: CMFormatDescriptionRef
type TextFormatDescriptionRef uintptr
// TextFormatType - A text format type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatType
// CMTextFormatType has base type: FourCharCode
type TextFormatType uintptr
// TextJustificationValue - An integer value that describes the justification modes for text media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextJustificationValue
// CMTextJustificationValue has base type: int8_t
type TextJustificationValue uintptr
// TimebaseRef - A model of a timeline under application control.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebase
// CMTimebaseRef has base type: struct OpaqueCMTimebase *
type TimebaseRef uintptr
// TimeCodeDescriptionFlavor - Types that represent time code format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeDescriptionFlavor
// CMTimeCodeDescriptionFlavor has base type: CFStringRef
type TimeCodeDescriptionFlavor uintptr
// TimeCodeFormatDescriptionRef - A type you use to interact with time code format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescription
// CMTimeCodeFormatDescriptionRef has base type: CMFormatDescriptionRef
type TimeCodeFormatDescriptionRef uintptr
// TimeCodeFormatType - A time code format type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatType
// CMTimeCodeFormatType has base type: FourCharCode
type TimeCodeFormatType uintptr
// TimeEpoch - An epoch for a time.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeEpoch
// CMTimeEpoch has base type: int64_t
type TimeEpoch uintptr
// TimeScale - An integer timescale.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeScale
// CMTimeScale has base type: int32_t
type TimeScale uintptr
// TimeValue - An integer time value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeValue
// CMTimeValue has base type: int64_t
type TimeValue uintptr
// VideoCodecType - A video codec type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoCodecType
// CMVideoCodecType has base type: FourCharCode
type VideoCodecType uintptr
// VideoFormatDescriptionRef - A type you use to interact with video format descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescription
// CMVideoFormatDescriptionRef has base type: CMFormatDescriptionRef
type VideoFormatDescriptionRef uintptr

