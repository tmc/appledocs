// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMedia Functions (142 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMAudioDeviceClockSetAudioDeviceID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferGetDataLength func(unsafe.Pointer) uintptr
	_CMBufferQueueCopyHead func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueAndRetain func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetHostTimeClock func() unsafe.Pointer
	_CMGetAttachment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolGetTypeID func() unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeConformsToDataType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeIsRegistered func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleAttachmentsArray func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetCapacity func(unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetTypeID func() unsafe.Pointer
	_CMSwapBigEndianClosedCaptionDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSyncConvertTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionRemoveAllTags func(unsafe.Pointer) unsafe.Pointer
	_CMTagCopyDescription func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetCMSampleBufferAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeGetSeconds func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingCopyDescription func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeCopyAsDictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeGetEnd func(unsafe.Pointer) unsafe.Pointer
	_CMTimeShow func(unsafe.Pointer)
	_CMTimebaseGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeAndRate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTypeID func() unsafe.Pointer
	_CMRemoveAttachment func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDataBuffer func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCopyAsDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetUltimateMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeMake func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueEnqueue func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetBufferCount func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolFlush func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataFailed func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseRemoveTimer func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMaximum func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMaster func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseRemoveTimerDispatchSource func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMake func(unsafe.Pointer) unsafe.Pointer
	_CMRemoveAllAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingCopyAsDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingShow func(unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferGetTypeID func() unsafe.Pointer
	_CMTimeRangeShow func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueSetValidationHandler func(unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionGetKeyWithLocalID func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetNumSamples func(unsafe.Pointer) unsafe.Pointer
	_CMTimeSubtract func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTimerToFireImmediately func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeGetUnion func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetMediaType func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockSetAudioDeviceUID func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetDuration func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputDuration func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySourceTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolCreate func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySource func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyUltimateMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetAnchorTime func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMinimum func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetMaster func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMinPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMapDurationFromRangeToRange func(unsafe.Pointer) unsafe.Pointer
	_CMClockInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferTrackDataReadiness func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetSourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMapTimeFromRangeToRange func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetHead func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolGetAllocator func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataBuffer func(unsafe.Pointer) unsafe.Pointer
	_CMTimeFoldIntoRange func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTime func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionGetFrameDuration func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetExtensions func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetFormatDescription func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetImageBuffer func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetSourceTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMakeEmpty func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyUltimateSourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDuration func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferMakeDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMTimeAdd func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetTypeID func() unsafe.Pointer
	_CMFormatDescriptionGetExtension func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTimerDispatchSourceToFireImmediately func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMinDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeue func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseAddTimerDispatchSource func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeWithTimeScale func(unsafe.Pointer) unsafe.Pointer
	_CMSyncGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetMediaSubType func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeFromTimeToTime func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueReset func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseAddTimer func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseNotificationBarrier func(unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionGetIdentifiers func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMPropagateAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetFirstDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSetAttachment func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetTypeID func() unsafe.Pointer
	_CMFormatDescriptionGetTypeID func() unsafe.Pointer
	_CMTimeRangeGetIntersection func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMaxPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueRemoveTrigger func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetTypeID func() unsafe.Pointer
	_CMSampleBufferSetOutputPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueReset func(unsafe.Pointer) unsafe.Pointer
	_CMCopyDictionaryOfAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferAssureBlockMemory func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueIfDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers func() unsafe.Pointer
	_CMSetAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetEndPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetFirstPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueMarkEndOfData func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeClampToRange func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionGetDimensions func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCallBlockForEachSample func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetInvalidateHandler func(unsafe.Pointer) unsafe.Pointer
	_CMTimeAbsoluteValue func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMasterClock func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceID, lib, "CMAudioDeviceClockSetAudioDeviceID")
	tryRegister(&_CMBlockBufferGetDataLength, lib, "CMBlockBufferGetDataLength")
	tryRegister(&_CMBufferQueueCopyHead, lib, "CMBufferQueueCopyHead")
	tryRegister(&_CMBufferQueueDequeueAndRetain, lib, "CMBufferQueueDequeueAndRetain")
	tryRegister(&_CMClockGetHostTimeClock, lib, "CMClockGetHostTimeClock")
	tryRegister(&_CMGetAttachment, lib, "CMGetAttachment")
	tryRegister(&_CMMemoryPoolGetTypeID, lib, "CMMemoryPoolGetTypeID")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeConformsToDataType, lib, "CMMetadataDataTypeRegistryDataTypeConformsToDataType")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeIsRegistered, lib, "CMMetadataDataTypeRegistryDataTypeIsRegistered")
	tryRegister(&_CMSampleBufferGetSampleAttachmentsArray, lib, "CMSampleBufferGetSampleAttachmentsArray")
	tryRegister(&_CMSimpleQueueGetCapacity, lib, "CMSimpleQueueGetCapacity")
	tryRegister(&_CMSimpleQueueGetTypeID, lib, "CMSimpleQueueGetTypeID")
	tryRegister(&_CMSwapBigEndianClosedCaptionDescriptionToHost, lib, "CMSwapBigEndianClosedCaptionDescriptionToHost")
	tryRegister(&_CMSyncConvertTime, lib, "CMSyncConvertTime")
	tryRegister(&_CMTagCollectionRemoveAllTags, lib, "CMTagCollectionRemoveAllTags")
	tryRegister(&_CMTagCopyDescription, lib, "CMTagCopyDescription")
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, lib, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferAtIndex, lib, "CMTaggedBufferGroupGetCMSampleBufferAtIndex")
	tryRegister(&_CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer, lib, "CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer")
	tryRegister(&_CMTimeGetSeconds, lib, "CMTimeGetSeconds")
	tryRegister(&_CMTimeMappingCopyDescription, lib, "CMTimeMappingCopyDescription")
	tryRegister(&_CMTimeRangeCopyAsDictionary, lib, "CMTimeRangeCopyAsDictionary")
	tryRegister(&_CMTimeRangeGetEnd, lib, "CMTimeRangeGetEnd")
	tryRegister(&_CMTimeShow, lib, "CMTimeShow")
	tryRegister(&_CMTimebaseGetTime, lib, "CMTimebaseGetTime")
	tryRegister(&_CMTimebaseGetTimeAndRate, lib, "CMTimebaseGetTimeAndRate")
	tryRegister(&_CMTimebaseGetTypeID, lib, "CMTimebaseGetTypeID")
	tryRegister(&_CMRemoveAttachment, lib, "CMRemoveAttachment")
	tryRegister(&_CMSampleBufferGetDataBuffer, lib, "CMSampleBufferGetDataBuffer")
	tryRegister(&_CMTimeCopyAsDictionary, lib, "CMTimeCopyAsDictionary")
	tryRegister(&_CMTimebaseGetUltimateMasterClock, lib, "CMTimebaseGetUltimateMasterClock")
	tryRegister(&_CMTimeRangeMake, lib, "CMTimeRangeMake")
	tryRegister(&_CMTimeRangeMakeFromDictionary, lib, "CMTimeRangeMakeFromDictionary")
	tryRegister(&_CMBufferQueueEnqueue, lib, "CMBufferQueueEnqueue")
	tryRegister(&_CMBufferQueueGetBufferCount, lib, "CMBufferQueueGetBufferCount")
	tryRegister(&_CMMemoryPoolFlush, lib, "CMMemoryPoolFlush")
	tryRegister(&_CMSampleBufferGetDecodeTimeStamp, lib, "CMSampleBufferGetDecodeTimeStamp")
	tryRegister(&_CMSampleBufferSetDataFailed, lib, "CMSampleBufferSetDataFailed")
	tryRegister(&_CMTimebaseRemoveTimer, lib, "CMTimebaseRemoveTimer")
	tryRegister(&_CMTimeMaximum, lib, "CMTimeMaximum")
	tryRegister(&_CMTimebaseCopyMaster, lib, "CMTimebaseCopyMaster")
	tryRegister(&_CMTimebaseRemoveTimerDispatchSource, lib, "CMTimebaseRemoveTimerDispatchSource")
	tryRegister(&_CMTimeMappingMake, lib, "CMTimeMappingMake")
	tryRegister(&_CMRemoveAllAttachments, lib, "CMRemoveAllAttachments")
	tryRegister(&_CMSampleBufferSetDataReady, lib, "CMSampleBufferSetDataReady")
	tryRegister(&_CMTimeMappingCopyAsDictionary, lib, "CMTimeMappingCopyAsDictionary")
	tryRegister(&_CMTimeMappingShow, lib, "CMTimeMappingShow")
	tryRegister(&_CMBlockBufferGetTypeID, lib, "CMBlockBufferGetTypeID")
	tryRegister(&_CMTimeRangeShow, lib, "CMTimeRangeShow")
	tryRegister(&_CMBufferQueueSetValidationHandler, lib, "CMBufferQueueSetValidationHandler")
	tryRegister(&_CMMetadataFormatDescriptionGetKeyWithLocalID, lib, "CMMetadataFormatDescriptionGetKeyWithLocalID")
	tryRegister(&_CMSampleBufferGetNumSamples, lib, "CMSampleBufferGetNumSamples")
	tryRegister(&_CMTimeSubtract, lib, "CMTimeSubtract")
	tryRegister(&_CMTimebaseSetTimerToFireImmediately, lib, "CMTimebaseSetTimerToFireImmediately")
	tryRegister(&_CMTimeRangeGetUnion, lib, "CMTimeRangeGetUnion")
	tryRegister(&_CMFormatDescriptionGetMediaType, lib, "CMFormatDescriptionGetMediaType")
	tryRegister(&_CMSampleBufferGetOutputDecodeTimeStamp, lib, "CMSampleBufferGetOutputDecodeTimeStamp")
	tryRegister(&_CMSampleBufferInvalidate, lib, "CMSampleBufferInvalidate")
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceUID, lib, "CMAudioDeviceClockSetAudioDeviceUID")
	tryRegister(&_CMBufferQueueGetDuration, lib, "CMBufferQueueGetDuration")
	tryRegister(&_CMSampleBufferGetOutputDuration, lib, "CMSampleBufferGetOutputDuration")
	tryRegister(&_CMTimebaseCopySourceTimebase, lib, "CMTimebaseCopySourceTimebase")
	tryRegister(&_CMMemoryPoolCreate, lib, "CMMemoryPoolCreate")
	tryRegister(&_CMSampleBufferGetPresentationTimeStamp, lib, "CMSampleBufferGetPresentationTimeStamp")
	tryRegister(&_CMTimebaseCopySource, lib, "CMTimebaseCopySource")
	tryRegister(&_CMTimebaseCopyUltimateMasterClock, lib, "CMTimebaseCopyUltimateMasterClock")
	tryRegister(&_CMTimebaseSetAnchorTime, lib, "CMTimebaseSetAnchorTime")
	tryRegister(&_CMMemoryPoolInvalidate, lib, "CMMemoryPoolInvalidate")
	tryRegister(&_CMSampleBufferGetOutputPresentationTimeStamp, lib, "CMSampleBufferGetOutputPresentationTimeStamp")
	tryRegister(&_CMTimeMakeFromDictionary, lib, "CMTimeMakeFromDictionary")
	tryRegister(&_CMTimeMinimum, lib, "CMTimeMinimum")
	tryRegister(&_CMTimebaseCopyMasterTimebase, lib, "CMTimebaseCopyMasterTimebase")
	tryRegister(&_CMTimebaseGetMaster, lib, "CMTimebaseGetMaster")
	tryRegister(&_CMBufferQueueGetMinPresentationTimeStamp, lib, "CMBufferQueueGetMinPresentationTimeStamp")
	tryRegister(&_CMTimeMapDurationFromRangeToRange, lib, "CMTimeMapDurationFromRangeToRange")
	tryRegister(&_CMClockInvalidate, lib, "CMClockInvalidate")
	tryRegister(&_CMSampleBufferTrackDataReadiness, lib, "CMSampleBufferTrackDataReadiness")
	tryRegister(&_CMTimebaseSetSourceClock, lib, "CMTimebaseSetSourceClock")
	tryRegister(&_CMTimeMapTimeFromRangeToRange, lib, "CMTimeMapTimeFromRangeToRange")
	tryRegister(&_CMBufferQueueGetHead, lib, "CMBufferQueueGetHead")
	tryRegister(&_CMMemoryPoolGetAllocator, lib, "CMMemoryPoolGetAllocator")
	tryRegister(&_CMSampleBufferSetDataBuffer, lib, "CMSampleBufferSetDataBuffer")
	tryRegister(&_CMTimeFoldIntoRange, lib, "CMTimeFoldIntoRange")
	tryRegister(&_CMTimebaseSetTime, lib, "CMTimebaseSetTime")
	tryRegister(&_CMTimeCodeFormatDescriptionGetFrameDuration, lib, "CMTimeCodeFormatDescriptionGetFrameDuration")
	tryRegister(&_CMTimebaseGetMasterTimebase, lib, "CMTimebaseGetMasterTimebase")
	tryRegister(&_CMFormatDescriptionGetExtensions, lib, "CMFormatDescriptionGetExtensions")
	tryRegister(&_CMSampleBufferGetFormatDescription, lib, "CMSampleBufferGetFormatDescription")
	tryRegister(&_CMSampleBufferGetImageBuffer, lib, "CMSampleBufferGetImageBuffer")
	tryRegister(&_CMTimebaseSetSourceTimebase, lib, "CMTimebaseSetSourceTimebase")
	tryRegister(&_CMTimeMappingMakeEmpty, lib, "CMTimeMappingMakeEmpty")
	tryRegister(&_CMTimebaseCopyUltimateSourceClock, lib, "CMTimebaseCopyUltimateSourceClock")
	tryRegister(&_CMTimebaseSetMasterClock, lib, "CMTimebaseSetMasterClock")
	tryRegister(&_CMSampleBufferGetDuration, lib, "CMSampleBufferGetDuration")
	tryRegister(&_CMSampleBufferMakeDataReady, lib, "CMSampleBufferMakeDataReady")
	tryRegister(&_CMTimeAdd, lib, "CMTimeAdd")
	tryRegister(&_CMBufferQueueGetTypeID, lib, "CMBufferQueueGetTypeID")
	tryRegister(&_CMFormatDescriptionGetExtension, lib, "CMFormatDescriptionGetExtension")
	tryRegister(&_CMTimebaseSetTimerDispatchSourceToFireImmediately, lib, "CMTimebaseSetTimerDispatchSourceToFireImmediately")
	tryRegister(&_CMBufferQueueGetMinDecodeTimeStamp, lib, "CMBufferQueueGetMinDecodeTimeStamp")
	tryRegister(&_CMBufferQueueDequeue, lib, "CMBufferQueueDequeue")
	tryRegister(&_CMTimebaseAddTimerDispatchSource, lib, "CMTimebaseAddTimerDispatchSource")
	tryRegister(&_CMTimebaseGetTimeWithTimeScale, lib, "CMTimebaseGetTimeWithTimeScale")
	tryRegister(&_CMSyncGetTime, lib, "CMSyncGetTime")
	tryRegister(&_CMFormatDescriptionGetMediaSubType, lib, "CMFormatDescriptionGetMediaSubType")
	tryRegister(&_CMTimeMappingMakeFromDictionary, lib, "CMTimeMappingMakeFromDictionary")
	tryRegister(&_CMTimeRangeFromTimeToTime, lib, "CMTimeRangeFromTimeToTime")
	tryRegister(&_CMBufferQueueReset, lib, "CMBufferQueueReset")
	tryRegister(&_CMTimebaseAddTimer, lib, "CMTimebaseAddTimer")
	tryRegister(&_CMTimebaseNotificationBarrier, lib, "CMTimebaseNotificationBarrier")
	tryRegister(&_CMMetadataFormatDescriptionGetIdentifiers, lib, "CMMetadataFormatDescriptionGetIdentifiers")
	tryRegister(&_CMTimebaseSetMasterTimebase, lib, "CMTimebaseSetMasterTimebase")
	tryRegister(&_CMPropagateAttachments, lib, "CMPropagateAttachments")
	tryRegister(&_CMTimeCopyDescription, lib, "CMTimeCopyDescription")
	tryRegister(&_CMTimebaseGetMasterClock, lib, "CMTimebaseGetMasterClock")
	tryRegister(&_CMBufferQueueGetFirstDecodeTimeStamp, lib, "CMBufferQueueGetFirstDecodeTimeStamp")
	tryRegister(&_CMSetAttachment, lib, "CMSetAttachment")
	tryRegister(&_CMClockGetTypeID, lib, "CMClockGetTypeID")
	tryRegister(&_CMFormatDescriptionGetTypeID, lib, "CMFormatDescriptionGetTypeID")
	tryRegister(&_CMTimeRangeGetIntersection, lib, "CMTimeRangeGetIntersection")
	tryRegister(&_CMBufferQueueGetMaxPresentationTimeStamp, lib, "CMBufferQueueGetMaxPresentationTimeStamp")
	tryRegister(&_CMBufferQueueRemoveTrigger, lib, "CMBufferQueueRemoveTrigger")
	tryRegister(&_CMSampleBufferGetTypeID, lib, "CMSampleBufferGetTypeID")
	tryRegister(&_CMSampleBufferSetOutputPresentationTimeStamp, lib, "CMSampleBufferSetOutputPresentationTimeStamp")
	tryRegister(&_CMSimpleQueueReset, lib, "CMSimpleQueueReset")
	tryRegister(&_CMCopyDictionaryOfAttachments, lib, "CMCopyDictionaryOfAttachments")
	tryRegister(&_CMBlockBufferAssureBlockMemory, lib, "CMBlockBufferAssureBlockMemory")
	tryRegister(&_CMBufferQueueDequeueIfDataReady, lib, "CMBufferQueueDequeueIfDataReady")
	tryRegister(&_CMClockGetTime, lib, "CMClockGetTime")
	tryRegister(&_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers, lib, "CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers")
	tryRegister(&_CMSetAttachments, lib, "CMSetAttachments")
	tryRegister(&_CMBufferQueueGetEndPresentationTimeStamp, lib, "CMBufferQueueGetEndPresentationTimeStamp")
	tryRegister(&_CMBufferQueueGetFirstPresentationTimeStamp, lib, "CMBufferQueueGetFirstPresentationTimeStamp")
	tryRegister(&_CMBufferQueueMarkEndOfData, lib, "CMBufferQueueMarkEndOfData")
	tryRegister(&_CMTimebaseCopySourceClock, lib, "CMTimebaseCopySourceClock")
	tryRegister(&_CMTimeClampToRange, lib, "CMTimeClampToRange")
	tryRegister(&_CMTimeRangeCopyDescription, lib, "CMTimeRangeCopyDescription")
	tryRegister(&_CMVideoFormatDescriptionGetDimensions, lib, "CMVideoFormatDescriptionGetDimensions")
	tryRegister(&_CMSampleBufferCallBlockForEachSample, lib, "CMSampleBufferCallBlockForEachSample")
	tryRegister(&_CMSampleBufferSetInvalidateHandler, lib, "CMSampleBufferSetInvalidateHandler")
	tryRegister(&_CMTimeAbsoluteValue, lib, "CMTimeAbsoluteValue")
	tryRegister(&_CMTimebaseCopyMasterClock, lib, "CMTimebaseCopyMasterClock")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// Added in macOS 10.8.
// Changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceID(_:deviceID:)
func CMAudioDeviceClockSetAudioDeviceID(clock unsafe.Pointer, deviceID unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceID(clock, deviceID)
}

// Returns the total length of data that’s accessible by a block buffer.
//
// Added in macOS 10.7.
// Returns the total length of data that’s accessible by a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataLength(_:)
func CMBlockBufferGetDataLength(theBuffer unsafe.Pointer) uintptr {
	return _CMBlockBufferGetDataLength(theBuffer)
}

// CMBufferQueueCopyHead is a CoreMedia function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueCopyHead(queue)
}

// Dequeues a buffer from a queue.
//
// Added in macOS 10.7.
// Dequeues a buffer from a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeueAndRetain(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeueAndRetain(queue)
}

// Returns a reference to the singleton clock that reflects the host time.
//
// Added in macOS 10.8.
// Returns a reference to the singleton clock that reflects the host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetHostTimeClock()
func CMClockGetHostTimeClock() unsafe.Pointer {
	return _CMClockGetHostTimeClock()
}

// Returns an attachment from an attachment bearer object.
//
// Added in macOS 10.7.
// Returns an attachment from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMGetAttachment(_:key:attachmentModeOut:)
func CMGetAttachment(target unsafe.Pointer, key unsafe.Pointer, attachmentModeOut unsafe.Pointer) unsafe.Pointer {
	return _CMGetAttachment(target, key, attachmentModeOut)
}

// Returns the type identifier of memory pool objects.
//
// Added in macOS 10.8.
// Returns the type identifier of memory pool objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() unsafe.Pointer {
	return _CMMemoryPoolGetTypeID()
}

// Returns a Boolean value that indicates whether a data type conforms to another data type.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether a data type conforms to another data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType unsafe.Pointer, conformsToDataType unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
}

// Returns a Boolean value that indicates the registration status of a data type identifier.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates the registration status of a data type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsRegistered(_:)
func CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType)
}

// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf unsafe.Pointer, createIfNecessary unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
}

// Returns the number of elements that the queue can hold.
//
// Added in macOS 10.7.
// Returns the number of elements that the queue can hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCapacity(_:)
func CMSimpleQueueGetCapacity(queue unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueGetCapacity(queue)
}

// Returns the type identifier of sample buffer objects.
//
// Added in macOS 10.7.
// Returns the type identifier of sample buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() unsafe.Pointer {
	return _CMSimpleQueueGetTypeID()
}

// Converts a closed caption description structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts a closed caption description structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianClosedCaptionDescriptionToHost(_:_:)
func CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData unsafe.Pointer, closedCaptionDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize)
}

// Converts a time from one timebase or clock to another timebase or clock.
//
// Added in macOS 10.8.
// Converts a time from one timebase or clock to another timebase or clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time unsafe.Pointer, fromClockOrTimebase unsafe.Pointer, toClockOrTimebase unsafe.Pointer) unsafe.Pointer {
	return _CMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
}

// Removes all tags from a collection.
//
// Added in macOS 14.0.
// Removes all tags from a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionRemoveAllTags(tagCollection)
}

// Copies the description of a tag to a new string.
//
// Added in macOS 14.0.
// Copies the description of a tag to a new string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCopyDescription
func CMTagCopyDescription(allocator unsafe.Pointer, tag unsafe.Pointer) unsafe.Pointer {
	return _CMTagCopyDescription(allocator, tag)
}

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions is a CoreMedia function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator unsafe.Pointer, taggedBufferGroup unsafe.Pointer, extensions unsafe.Pointer, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
}

// Gets the sample buffer at a given index in the buffer group.
//
// Added in macOS 14.0.
// Gets the sample buffer at a given index in the buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
}

// Copies the contents of a time code format description to a buffer in big-endian byte order.
//
// Added in macOS 10.10.
// Copies the contents of a time code format description to a buffer in big-endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator:timeCodeFormatDescription:flavor:blockBufferOut:)
func CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator unsafe.Pointer, timeCodeFormatDescription unsafe.Pointer, flavor unsafe.Pointer, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeFormatDescription, flavor, blockBufferOut)
}

// Returns a representation of the time in seconds.
//
// Added in macOS 10.7.
// Returns a representation of the time in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time unsafe.Pointer) unsafe.Pointer {
	return _CMTimeGetSeconds(time)
}

// Copies a string description of a time mapping.
//
// Added in macOS 10.11.
// Copies a string description of a time mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator unsafe.Pointer, mapping unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingCopyDescription(allocator, mapping)
}

// Returns a dictionary representation of a time range.
//
// Added in macOS 10.7.
// Returns a dictionary representation of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ unsafe.Pointer, allocator unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeCopyAsDictionary(range_, allocator)
}

// Returns a time value that represents the end of a time range.
//
// Added in macOS 10.7.
// Returns a time value that represents the end of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetEnd(_:)
func CMTimeRangeGetEnd(range_ unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetEnd(range_)
}

// Prints a description of the time to the console.
//
// Added in macOS 10.7.
// Prints a description of the time to the console.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeShow(_:)
func CMTimeShow(time unsafe.Pointer) {
	_CMTimeShow(time)
}

// Returns the current time from a timebase.
//
// Added in macOS 10.8.
// Returns the current time from a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTime(_:)
func CMTimebaseGetTime(timebase unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTime(timebase)
}

// Returns the current time and rate of a timebase.
//
// Added in macOS 10.8.
// Returns the current time and rate of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase unsafe.Pointer, timeOut unsafe.Pointer, rateOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
}

// Returns the Core Foundation type identifier that identifies a timebase object.
//
// Added in macOS 10.8.
// Returns the Core Foundation type identifier that identifies a timebase object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() unsafe.Pointer {
	return _CMTimebaseGetTypeID()
}

// Removes a specific attachment from an attachment bearer object.

// Removes a specific attachment from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAttachment(_:key:)
func CMRemoveAttachment(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMRemoveAttachment(p0)
}

// Returns a block buffer that contains the media data.

// Returns a block buffer that contains the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDataBuffer(_:)
func CMSampleBufferGetDataBuffer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDataBuffer(p0)
}

// Creates a dictionary representation of the time.

// Creates a dictionary representation of the time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyAsDictionary(_:allocator:)
func CMTimeCopyAsDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCopyAsDictionary(p0)
}

// Returns the host clock that is the host of all of a timebase’s host timebases.

// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetUltimateMasterClock(_:)
func CMTimebaseGetUltimateMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetUltimateMasterClock(p0)
}

// Creates a valid time range with a start time and duration.

// Creates a valid time range with a start time and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMake(start:duration:)
func CMTimeRangeMake(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeMake(p0)
}

// Creates a time range from a dictionary representation of its fields.

// Creates a time range from a dictionary representation of its fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMakeFromDictionary(_:)
func CMTimeRangeMakeFromDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeMakeFromDictionary(p0)
}

// Enqueues a buffer onto a queue.

// Enqueues a buffer onto a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)
func CMBufferQueueEnqueue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueEnqueue(p0)
}

// Gets the number of buffers in the queue.

// Gets the number of buffers in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetBufferCount(_:)
func CMBufferQueueGetBufferCount(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetBufferCount(p0)
}

// Deallocates all memory the pool holds.

// Deallocates all memory the pool holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolFlush(_:)
func CMMemoryPoolFlush(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolFlush(p0)
}

// Returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.

// Returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDecodeTimeStamp(_:)
func CMSampleBufferGetDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDecodeTimeStamp(p0)
}

// Marks the sample buffer’s data as failed to indicate that it won’t become ready.

// Marks the sample buffer’s data as failed to indicate that it won’t become ready.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataFailed(_:status:)
func CMSampleBufferSetDataFailed(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataFailed(p0)
}

// Removes the timer from the list of timers the timebase manages.

// Removes the timer from the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimer(_:timer:)
func CMTimebaseRemoveTimer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseRemoveTimer(p0)
}

// Returns the greater of two time values.

// Returns the greater of two time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMaximum(_:_:)
func CMTimeMaximum(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMaximum(p0)
}

// Returns the immediate host timebase of a timebase.

// Returns the immediate host timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMaster(_:)
func CMTimebaseCopyMaster(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMaster(p0)
}

// Removes the timer dispatch source from the list of timers the timebase manages.

// Removes the timer dispatch source from the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimerDispatchSource(_:timerSource:)
func CMTimebaseRemoveTimerDispatchSource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseRemoveTimerDispatchSource(p0)
}

// Creates a time mapping with a source and target time range.

// Creates a time mapping with a source and target time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMake(source:target:)
func CMTimeMappingMake(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMake(p0)
}

// Removes all attachments from an attachment bearer object.

// Removes all attachments from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAllAttachments(_:)
func CMRemoveAllAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMRemoveAllAttachments(p0)
}

// Marks a sample buffer’s data as ready for use.

// Marks a sample buffer’s data as ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataReady(_:)
func CMSampleBufferSetDataReady(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataReady(p0)
}

// Returns a dictionary representation of a time mapping.

// Returns a dictionary representation of a time mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyAsDictionary(_:allocator:)
func CMTimeMappingCopyAsDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingCopyAsDictionary(p0)
}

// Prints a description of a time mapping to standard output.

// Prints a description of a time mapping to standard output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingShow(_:)
func CMTimeMappingShow(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingShow(p0)
}

// Returns the type identifier for block buffer objects.

// Returns the type identifier for block buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetTypeID()
func CMBlockBufferGetTypeID() unsafe.Pointer {
	return _CMBlockBufferGetTypeID()
}

// Prints a description of the time range to standard error.

// Prints a description of the time range to standard error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeShow(_:)
func CMTimeRangeShow(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeShow(p0)
}

// A validation handler for the queue to call before enqueuing buffers.

// A validation handler for the queue to call before enqueuing buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationHandler(_:_:)
func CMBufferQueueSetValidationHandler(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueSetValidationHandler(p0)
}

// Returns the key for the local identifier.

// Returns the key for the local identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetKeyWithLocalID(_:localKeyID:)
func CMMetadataFormatDescriptionGetKeyWithLocalID(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionGetKeyWithLocalID(p0)
}

// Returns the number of media samples in a sample buffer.

// Returns the number of media samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetNumSamples(_:)
func CMSampleBufferGetNumSamples(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetNumSamples(p0)
}

// Returns the difference between two times.

// Returns the difference between two times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeSubtract(_:_:)
func CMTimeSubtract(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeSubtract(p0)
}

// Sets the timer to fire immediately once, overriding any previous timer calls.

// Sets the timer to fire immediately once, overriding any previous timer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerToFireImmediately(_:timer:)
func CMTimebaseSetTimerToFireImmediately(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTimerToFireImmediately(p0)
}

// Returns a new time range with the time elements of the input.

// Returns a new time range with the time elements of the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetUnion(_:otherRange:)
func CMTimeRangeGetUnion(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetUnion(p0)
}

// Returns the media type of a format description.

// Returns the media type of a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaType(_:)
func CMFormatDescriptionGetMediaType(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetMediaType(p0)
}

// Returns the output decode timestamp of a sample buffer.

// Returns the output decode timestamp of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDecodeTimeStamp(_:)
func CMSampleBufferGetOutputDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputDecodeTimeStamp(p0)
}

// Invalidates a sample buffer by calling its invalidation callback.

// Invalidates a sample buffer by calling its invalidation callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidate(_:)
func CMSampleBufferInvalidate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferInvalidate(p0)
}

// Changes the Core Audio device the clock is tracking by specifying a new device unique identifier.

// Changes the Core Audio device the clock is tracking by specifying a new device unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceUID(_:deviceUID:)
func CMAudioDeviceClockSetAudioDeviceUID(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceUID(p0)
}

// Gets the duration of a buffer queue.

// Gets the duration of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetDuration(_:)
func CMBufferQueueGetDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetDuration(p0)
}

// Returns the output duration of a sample buffer.

// Returns the output duration of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDuration(_:)
func CMSampleBufferGetOutputDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputDuration(p0)
}

// Returns the immediate source timebase of a timebase.

// Returns the immediate source timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceTimebase(_:)
func CMTimebaseCopySourceTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySourceTimebase(p0)
}

// Creates a memory pool.

// Creates a memory pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolCreate(options:)
func CMMemoryPoolCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolCreate(p0)
}

// Returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.

// Returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetPresentationTimeStamp(_:)
func CMSampleBufferGetPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetPresentationTimeStamp(p0)
}

// Returns the immediate source — either a clock or timebase — of a timebase.

// Returns the immediate source — either a clock or timebase — of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySource(_:)
func CMTimebaseCopySource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySource(p0)
}

// Returns the host clock that is the host of all of a timebase’s host timebases.

// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateMasterClock(_:)
func CMTimebaseCopyUltimateMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyUltimateMasterClock(p0)
}

// Sets the time of a timebase at a particular host time.

// Sets the time of a timebase at a particular host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetAnchorTime(_:timebaseTime:immediateSourceTime:)
func CMTimebaseSetAnchorTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetAnchorTime(p0)
}

// Invalidates the memory pool, which causes its allocator to stop recycling memory.

// Invalidates the memory pool, which causes its allocator to stop recycling memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolInvalidate(_:)
func CMMemoryPoolInvalidate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolInvalidate(p0)
}

// Returns the output presentation timestamp of a sample buffer.

// Returns the output presentation timestamp of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputPresentationTimeStamp(_:)
func CMSampleBufferGetOutputPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputPresentationTimeStamp(p0)
}

// Creates a time from a dictionary representation of its fields.

// Creates a time from a dictionary representation of its fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeFromDictionary(_:)
func CMTimeMakeFromDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMakeFromDictionary(p0)
}

// Returns the lesser of two time values.

// Returns the lesser of two time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMinimum(_:_:)
func CMTimeMinimum(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMinimum(p0)
}

// Returns the immediate host timebase of a timebase.

// Returns the immediate host timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterTimebase(_:)
func CMTimebaseCopyMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMasterTimebase(p0)
}

// Returns the immediate host (either timebase or clock) of a timebase.

// Returns the immediate host (either timebase or clock) of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMaster(_:)
func CMTimebaseGetMaster(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMaster(p0)
}

// Gets the earliest presentation timestamp of a buffer queue.

// Gets the earliest presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinPresentationTimeStamp(_:)
func CMBufferQueueGetMinPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMinPresentationTimeStamp(p0)
}

// Translates a duration through a mapping from two time ranges.

// Translates a duration through a mapping from two time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapDurationFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapDurationFromRangeToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMapDurationFromRangeToRange(p0)
}

// Stops the clock.

// Stops the clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockInvalidate(_:)
func CMClockInvalidate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMClockInvalidate(p0)
}

// Associates a sample buffer’s data readiness with that of another sample buffer.

// Associates a sample buffer’s data readiness with that of another sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferTrackDataReadiness(_:sampleBufferToTrack:)
func CMSampleBufferTrackDataReadiness(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferTrackDataReadiness(p0)
}

// Sets the source clock of a timebase.

// Sets the source clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceClock(_:_:)
func CMTimebaseSetSourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetSourceClock(p0)
}

// Translates a time through a mapping from two time ranges.

// Translates a time through a mapping from two time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapTimeFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapTimeFromRangeToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMapTimeFromRangeToRange(p0)
}

// Retrieves the next buffer from a queue, but doesn’t remove it.

// Retrieves the next buffer from a queue, but doesn’t remove it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetHead(_:)
func CMBufferQueueGetHead(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetHead(p0)
}

// Returns the allocator for the memory pool.

// Returns the allocator for the memory pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetAllocator(_:)
func CMMemoryPoolGetAllocator(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolGetAllocator(p0)
}

// Sets a block buffer of media data on a sample buffer.

// Sets a block buffer of media data on a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBuffer(_:newValue:)
func CMSampleBufferSetDataBuffer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataBuffer(p0)
}

// Folds a time into a time range.

// Folds a time into a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFoldIntoRange(_:foldRange:)
func CMTimeFoldIntoRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeFoldIntoRange(p0)
}

// Sets the current time of a timebase.

// Sets the current time of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTime(_:time:)
func CMTimebaseSetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTime(p0)
}

// Returns the duration of each frame.

// Returns the duration of each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameDuration(_:)
func CMTimeCodeFormatDescriptionGetFrameDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionGetFrameDuration(p0)
}

// Returns the immediate host timebase of a timebase.

// Returns the immediate host timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterTimebase(_:)
func CMTimebaseGetMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMasterTimebase(p0)
}

// Returns all of the extensions for a format description.

// Returns all of the extensions for a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtensions(_:)
func CMFormatDescriptionGetExtensions(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetExtensions(p0)
}

// Returns the format description of the samples in a sample buffer.

// Returns the format description of the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetFormatDescription(_:)
func CMSampleBufferGetFormatDescription(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetFormatDescription(p0)
}

// Returns an image buffer that contains the media data.

// Returns an image buffer that contains the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetImageBuffer(_:)
func CMSampleBufferGetImageBuffer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetImageBuffer(p0)
}

// Sets the source timebase of a timebase.

// Sets the source timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceTimebase(_:_:)
func CMTimebaseSetSourceTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetSourceTimebase(p0)
}

// Creates a valid time mapping with an empty source.

// Creates a valid time mapping with an empty source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeEmpty(target:)
func CMTimeMappingMakeEmpty(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMakeEmpty(p0)
}

// Returns the source clock that’s the source of all of a timebase’s source timebases.

// Returns the source clock that’s the source of all of a timebase’s source timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateSourceClock(_:)
func CMTimebaseCopyUltimateSourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyUltimateSourceClock(p0)
}

// Sets the time of a timebase at a particular source time.

// Sets the time of a timebase at a particular source time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetMasterClock(_:_:)
func CMTimebaseSetMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetMasterClock(p0)
}

// Returns the total duration of a sample buffer.

// Returns the total duration of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDuration(_:)
func CMSampleBufferGetDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDuration(p0)
}

// Makes the sample buffer’s data ready for use by invoking its callback to load the data.

// Makes the sample buffer’s data ready for use by invoking its callback to load the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReady(_:)
func CMSampleBufferMakeDataReady(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferMakeDataReady(p0)
}

// Returns the sum of two times.

// Returns the sum of two times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAdd(_:_:)
func CMTimeAdd(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeAdd(p0)
}

// Returns the type identifier of buffer queue objects.

// Returns the type identifier of buffer queue objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTypeID()
func CMBufferQueueGetTypeID() unsafe.Pointer {
	return _CMBufferQueueGetTypeID()
}

// Returns an extension from the format description by using an extension key.

// Returns an extension from the format description by using an extension key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtension(_:extensionKey:)
func CMFormatDescriptionGetExtension(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetExtension(p0)
}

// Sets the timer dispatch source to fire immediately once, overriding any previous timer call.

// Sets the timer dispatch source to fire immediately once, overriding any previous timer call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceToFireImmediately(_:timerSource:)
func CMTimebaseSetTimerDispatchSourceToFireImmediately(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTimerDispatchSourceToFireImmediately(p0)
}

// Gets the earliest decode timestamp of a buffer queue.

// Gets the earliest decode timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinDecodeTimeStamp(_:)
func CMBufferQueueGetMinDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMinDecodeTimeStamp(p0)
}

// Dequeues a buffer from a queue.

// Dequeues a buffer from a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeue(p0)
}

// Adds the timer dispatch source to the list of timers the timebase manages.

// Adds the timer dispatch source to the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimerDispatchSource(_:timerSource:)
func CMTimebaseAddTimerDispatchSource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseAddTimerDispatchSource(p0)
}

// Returns the current time from a timebase in the specified timescale.

// Returns the current time from a timebase in the specified timescale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeWithTimeScale(_:timescale:method:)
func CMTimebaseGetTimeWithTimeScale(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeWithTimeScale(p0)
}

// Returns the time from a clock or timebase.

// Returns the time from a clock or timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncGetTime(_:)
func CMSyncGetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSyncGetTime(p0)
}

// Returns the media subtype of a format description.

// Returns the media subtype of a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
func CMFormatDescriptionGetMediaSubType(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetMediaSubType(p0)
}

// Creates a time mapping from a dictionary representation.

// Creates a time mapping from a dictionary representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeFromDictionary(_:)
func CMTimeMappingMakeFromDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMakeFromDictionary(p0)
}

// Creates a valid time range from a start and end time.

// Creates a valid time range from a start and end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeFromTimeToTime(start:end:)
func CMTimeRangeFromTimeToTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeFromTimeToTime(p0)
}

// Resets a buffer queue, which allows it to enqueue new buffers.

// Resets a buffer queue, which allows it to enqueue new buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueReset(_:)
func CMBufferQueueReset(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueReset(p0)
}

// Adds the timer to the list of timers the timebase manages.

// Adds the timer to the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimer(_:timer:runloop:)
func CMTimebaseAddTimer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseAddTimer(p0)
}

// Requests that the timebase wait until it isn’t posting notifications.

// Requests that the timebase wait until it isn’t posting notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseNotificationBarrier(_:)
func CMTimebaseNotificationBarrier(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseNotificationBarrier(p0)
}

// Returns an array of metadata identifiers from a metadata format description.

// Returns an array of metadata identifiers from a metadata format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetIdentifiers(_:)
func CMMetadataFormatDescriptionGetIdentifiers(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionGetIdentifiers(p0)
}

// CMTimebaseSetMasterTimebase is a CoreMedia function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetMasterTimebase(_:_:)
func CMTimebaseSetMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetMasterTimebase(p0)
}

// Copies all propagable attachments from one attachment bearer object to another.

// Copies all propagable attachments from one attachment bearer object to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPropagateAttachments(_:destination:)
func CMPropagateAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMPropagateAttachments(p0)
}

// Creates a string representation of the time.

// Creates a string representation of the time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyDescription(allocator:time:)
func CMTimeCopyDescription(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCopyDescription(p0)
}

// Returns the immediate host clock of a timebase.

// Returns the immediate host clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterClock(_:)
func CMTimebaseGetMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMasterClock(p0)
}

// Gets the decode timestamp of the first buffer in a buffer queue.

// Gets the decode timestamp of the first buffer in a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstDecodeTimeStamp(_:)
func CMBufferQueueGetFirstDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetFirstDecodeTimeStamp(p0)
}

// Sets or adds an attachment to an attachment bearer object.

// Sets or adds an attachment to an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachment(_:key:value:attachmentMode:)
func CMSetAttachment(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSetAttachment(p0)
}

// Returns the core foundation type identifier of a clock type.

// Returns the core foundation type identifier of a clock type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTypeID()
func CMClockGetTypeID() unsafe.Pointer {
	return _CMClockGetTypeID()
}

// Returns the Core Foundation type identifier that identifies format description objects.

// Returns the Core Foundation type identifier that identifies format description objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetTypeID()
func CMFormatDescriptionGetTypeID() unsafe.Pointer {
	return _CMFormatDescriptionGetTypeID()
}

// Returns a new time range with the time elements that are common between the input.

// Returns a new time range with the time elements that are common between the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetIntersection(_:otherRange:)
func CMTimeRangeGetIntersection(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetIntersection(p0)
}

// Gets the greatest presentation timestamp of a buffer queue.

// Gets the greatest presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMaxPresentationTimeStamp(_:)
func CMBufferQueueGetMaxPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMaxPresentationTimeStamp(p0)
}

// Removes a previously installed trigger from a buffer queue.

// Removes a previously installed trigger from a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueRemoveTrigger(_:triggerToken:)
func CMBufferQueueRemoveTrigger(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueRemoveTrigger(p0)
}

// Returns the type identifier of sample buffer objects.

// Returns the type identifier of sample buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTypeID()
func CMSampleBufferGetTypeID() unsafe.Pointer {
	return _CMSampleBufferGetTypeID()
}

// Sets an output presentation timestamp to use in place of a calculated value.

// Sets an output presentation timestamp to use in place of a calculated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetOutputPresentationTimeStamp(_:newValue:)
func CMSampleBufferSetOutputPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetOutputPresentationTimeStamp(p0)
}

// Resets the queue.

// Resets the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueReset(_:)
func CMSimpleQueueReset(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueReset(p0)
}

// Returns a dictionary of all attachments for an attachment bearer object.

// Returns a dictionary of all attachments for an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMCopyDictionaryOfAttachments(allocator:target:attachmentMode:)
func CMCopyDictionaryOfAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMCopyDictionaryOfAttachments(p0)
}

// Assures that the system allocates memory for all memory blocks in a block buffer.

// Assures that the system allocates memory for all memory blocks in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAssureBlockMemory(_:)
func CMBlockBufferAssureBlockMemory(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferAssureBlockMemory(p0)
}

// Dequeues a buffer from a queue, if it’s ready.

// Dequeues a buffer from a queue, if it’s ready.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReady(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeueIfDataReady(p0)
}

// Returns the current time from a clock.

// Returns the current time from a clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTime(_:)
func CMClockGetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMClockGetTime(p0)
}

// Returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.

// Returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
func CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() unsafe.Pointer {
	return _CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
}

// Sets a dictionary of attachments on an attachment bearer object.

// Sets a dictionary of attachments on an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachments(_:attachments:attachmentMode:)
func CMSetAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSetAttachments(p0)
}

// Gets the greatest end presentation timestamp of a buffer queue.

// Gets the greatest end presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetEndPresentationTimeStamp(_:)
func CMBufferQueueGetEndPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetEndPresentationTimeStamp(p0)
}

// Gets the presentation timestamp of the first buffer in a buffer queue.

// Gets the presentation timestamp of the first buffer in a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstPresentationTimeStamp(_:)
func CMBufferQueueGetFirstPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetFirstPresentationTimeStamp(p0)
}

// Sets a marker to indicate this queue doesn’t allow enqueuing new buffers.

// Sets a marker to indicate this queue doesn’t allow enqueuing new buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueMarkEndOfData(_:)
func CMBufferQueueMarkEndOfData(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueMarkEndOfData(p0)
}

// Returns the immediate source clock of a timebase.

// Returns the immediate source clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceClock(_:)
func CMTimebaseCopySourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySourceClock(p0)
}

// Returns the nearest time value inside the time range.

// Returns the nearest time value inside the time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeClampToRange(_:range:)
func CMTimeClampToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeClampToRange(p0)
}

// Returns a string with a description of a time range.

// Returns a string with a description of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyDescription(allocator:range:)
func CMTimeRangeCopyDescription(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeCopyDescription(p0)
}

// Returns the video dimensions, in encoded pixels.

// Returns the video dimensions, in encoded pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetDimensions(_:)
func CMVideoFormatDescriptionGetDimensions(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionGetDimensions(p0)
}

// Calls a block for every individual sample in a sample buffer.

// Calls a block for every individual sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallBlockForEachSample(_:_:)
func CMSampleBufferCallBlockForEachSample(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCallBlockForEachSample(p0)
}

// Sets the sample buffer’s invalidation handler.

// Sets the sample buffer’s invalidation handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateHandler(_:invalidateHandler:)
func CMSampleBufferSetInvalidateHandler(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetInvalidateHandler(p0)
}

// Returns the absolute value of a time.

// Returns the absolute value of a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAbsoluteValue(_:)
func CMTimeAbsoluteValue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeAbsoluteValue(p0)
}

// Returns the immediate host clock of a timebase.

// Returns the immediate host clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterClock(_:)
func CMTimebaseCopyMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMasterClock(p0)
}



