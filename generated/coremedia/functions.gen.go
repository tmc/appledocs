// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMedia Functions (192 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMAudioDeviceClockCreateFromAudioDeviceID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockSetAudioDeviceID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockSetAudioDeviceUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioSampleBufferCreateReadyWithPacketDescriptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioSampleBufferCreateWithPacketDescriptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateContiguous func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateWithBufferReference func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateWithMemoryBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferGetDataLength func(unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferIsRangeContiguous func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueCopyHead func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueAndRetain func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueIfDataReadyAndRetain func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetHostTimeClock func() unsafe.Pointer
	_CMClockMightDrift func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolGetTypeID func() unsafe.Pointer
	_CMMetadataCreateKeyFromIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMetadataCreateKeySpaceFromIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeConformsToDataType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeIsRegistered func(unsafe.Pointer) unsafe.Pointer
	_CMMetadataDataTypeRegistryRegisterDataType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCallBlockForEachSample func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCallForEachSample func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCopyPCMDataIntoAudioBufferList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCopySampleBufferForRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateCopyWithNewTiming func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateForImageBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateReady func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateReadyWithImageBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateWithMakeDataReadyHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferDataIsReady func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetAudioStreamPacketDescriptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetAudioStreamPacketDescriptionsPtr func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDataBuffer func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDuration func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetFormatDescription func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetImageBuffer func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetNumSamples func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputDuration func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetOutputSampleTimingInfoArray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleAttachmentsArray func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleSizeArray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleTimingInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetTaggedBufferGroup func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetTotalSampleSize func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferHasDataFailed func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferIsValid func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferMakeDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataBuffer func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataBufferFromAudioBufferList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataFailed func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetInvalidateCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetInvalidateHandler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetOutputPresentationTimeStamp func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferTrackDataReadiness func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetCapacity func(unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetHead func(unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetTypeID func() unsafe.Pointer
	_CMSwapBigEndianClosedCaptionDescriptionToHost func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSyncConvertTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateUnion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetTagsWithCategory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetTypeID func() unsafe.Pointer
	_CMTagCollectionRemoveAllTags func(unsafe.Pointer) unsafe.Pointer
	_CMTagGetFlagsValue func(unsafe.Pointer) unsafe.Pointer
	_CMTagGetSInt64Value func(unsafe.Pointer) unsafe.Pointer
	_CMTagMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetCMSampleBufferAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetCMSampleBufferForTagCollection func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetCount func(unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetTypeID func() unsafe.Pointer
	_CMTimeFoldIntoRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeGetSeconds func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMakeWithSeconds func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingCopyDescription func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeCopyAsDictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCreateWithMasterTimebase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCreateWithSourceTimebase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetEffectiveRate func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeAndRate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTypeID func() unsafe.Pointer
	_CMVideoFormatDescriptionGetHEVCParameterSetAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMPropagateAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueReset func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetAnchorTime func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeMake func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMapDurationFromRangeToRange func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueMarkEndOfData func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionGetFrameDuration func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySourceTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeFromTimeToTime func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeue func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetBufferCount func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMaxPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseAddTimer func(unsafe.Pointer) unsafe.Pointer
	_CMClockInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMTimeAbsoluteValue func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMake func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeGetUnion func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingShow func(unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferGetTypeID func() unsafe.Pointer
	_CMTimebaseGetMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTime func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetTypeID func() unsafe.Pointer
	_CMTimeMinimum func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseNotificationBarrier func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueIfDataReady func(unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferAssureBlockMemory func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetFirstPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCopyAsDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetSourceTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMapTimeFromRangeToRange func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetMaster func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseAddTimerDispatchSource func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyUltimateMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTimerDispatchSourceToFireImmediately func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueRemoveTrigger func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetTypeID func() unsafe.Pointer
	_CMMemoryPoolInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetTypeID func() unsafe.Pointer
	_CMTimeMaximum func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueSetValidationHandler func(unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionGetIdentifiers func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeGetIntersection func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolFlush func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMTimeCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMasterTimebase func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetExtensions func(unsafe.Pointer) unsafe.Pointer
	_CMCopyDictionaryOfAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMSetAttachment func(unsafe.Pointer) unsafe.Pointer
	_CMSetAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueEnqueue func(unsafe.Pointer) unsafe.Pointer
	_CMTimeSubtract func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopySource func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetHead func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetMediaSubType func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMinPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetExtension func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolCreate func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyUltimateSourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolGetAllocator func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseRemoveTimerDispatchSource func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueReset func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetUltimateMasterClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingMakeEmpty func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeWithTimeScale func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetFirstDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetDuration func(unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionGetKeyWithLocalID func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetSourceClock func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetMinDecodeTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMTimeShow func(unsafe.Pointer) unsafe.Pointer
	_CMRemoveAllAttachments func(unsafe.Pointer) unsafe.Pointer
	_CMRemoveAttachment func(unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueGetEndPresentationTimeStamp func(unsafe.Pointer) unsafe.Pointer
	_CMSyncGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseRemoveTimer func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeMakeFromDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimeClampToRange func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeGetEnd func(unsafe.Pointer) unsafe.Pointer
	_CMTimeAdd func(unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionGetMediaType func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingCopyAsDictionary func(unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTimerToFireImmediately func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetTime func(unsafe.Pointer) unsafe.Pointer
	_CMClockGetTypeID func() unsafe.Pointer
	_CMVideoFormatDescriptionGetDimensions func(unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers func() unsafe.Pointer
	_CMTimebaseCopyMaster func(unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeShow func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMAudioDeviceClockCreateFromAudioDeviceID, lib, "CMAudioDeviceClockCreateFromAudioDeviceID")
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceID, lib, "CMAudioDeviceClockSetAudioDeviceID")
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceUID, lib, "CMAudioDeviceClockSetAudioDeviceUID")
	tryRegister(&_CMAudioSampleBufferCreateReadyWithPacketDescriptions, lib, "CMAudioSampleBufferCreateReadyWithPacketDescriptions")
	tryRegister(&_CMAudioSampleBufferCreateWithPacketDescriptions, lib, "CMAudioSampleBufferCreateWithPacketDescriptions")
	tryRegister(&_CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler, lib, "CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler")
	tryRegister(&_CMBlockBufferCreateContiguous, lib, "CMBlockBufferCreateContiguous")
	tryRegister(&_CMBlockBufferCreateWithBufferReference, lib, "CMBlockBufferCreateWithBufferReference")
	tryRegister(&_CMBlockBufferCreateWithMemoryBlock, lib, "CMBlockBufferCreateWithMemoryBlock")
	tryRegister(&_CMBlockBufferGetDataLength, lib, "CMBlockBufferGetDataLength")
	tryRegister(&_CMBlockBufferIsRangeContiguous, lib, "CMBlockBufferIsRangeContiguous")
	tryRegister(&_CMBufferQueueCopyHead, lib, "CMBufferQueueCopyHead")
	tryRegister(&_CMBufferQueueDequeueAndRetain, lib, "CMBufferQueueDequeueAndRetain")
	tryRegister(&_CMBufferQueueDequeueIfDataReadyAndRetain, lib, "CMBufferQueueDequeueIfDataReadyAndRetain")
	tryRegister(&_CMClockGetHostTimeClock, lib, "CMClockGetHostTimeClock")
	tryRegister(&_CMClockMightDrift, lib, "CMClockMightDrift")
	tryRegister(&_CMMemoryPoolGetTypeID, lib, "CMMemoryPoolGetTypeID")
	tryRegister(&_CMMetadataCreateKeyFromIdentifier, lib, "CMMetadataCreateKeyFromIdentifier")
	tryRegister(&_CMMetadataCreateKeySpaceFromIdentifier, lib, "CMMetadataCreateKeySpaceFromIdentifier")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeConformsToDataType, lib, "CMMetadataDataTypeRegistryDataTypeConformsToDataType")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeIsRegistered, lib, "CMMetadataDataTypeRegistryDataTypeIsRegistered")
	tryRegister(&_CMMetadataDataTypeRegistryRegisterDataType, lib, "CMMetadataDataTypeRegistryRegisterDataType")
	tryRegister(&_CMSampleBufferCallBlockForEachSample, lib, "CMSampleBufferCallBlockForEachSample")
	tryRegister(&_CMSampleBufferCallForEachSample, lib, "CMSampleBufferCallForEachSample")
	tryRegister(&_CMSampleBufferCopyPCMDataIntoAudioBufferList, lib, "CMSampleBufferCopyPCMDataIntoAudioBufferList")
	tryRegister(&_CMSampleBufferCopySampleBufferForRange, lib, "CMSampleBufferCopySampleBufferForRange")
	tryRegister(&_CMSampleBufferCreateCopy, lib, "CMSampleBufferCreateCopy")
	tryRegister(&_CMSampleBufferCreateCopyWithNewTiming, lib, "CMSampleBufferCreateCopyWithNewTiming")
	tryRegister(&_CMSampleBufferCreateForImageBuffer, lib, "CMSampleBufferCreateForImageBuffer")
	tryRegister(&_CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler, lib, "CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler")
	tryRegister(&_CMSampleBufferCreateReady, lib, "CMSampleBufferCreateReady")
	tryRegister(&_CMSampleBufferCreateReadyWithImageBuffer, lib, "CMSampleBufferCreateReadyWithImageBuffer")
	tryRegister(&_CMSampleBufferCreateWithMakeDataReadyHandler, lib, "CMSampleBufferCreateWithMakeDataReadyHandler")
	tryRegister(&_CMSampleBufferDataIsReady, lib, "CMSampleBufferDataIsReady")
	tryRegister(&_CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer, lib, "CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer")
	tryRegister(&_CMSampleBufferGetAudioStreamPacketDescriptions, lib, "CMSampleBufferGetAudioStreamPacketDescriptions")
	tryRegister(&_CMSampleBufferGetAudioStreamPacketDescriptionsPtr, lib, "CMSampleBufferGetAudioStreamPacketDescriptionsPtr")
	tryRegister(&_CMSampleBufferGetDataBuffer, lib, "CMSampleBufferGetDataBuffer")
	tryRegister(&_CMSampleBufferGetDecodeTimeStamp, lib, "CMSampleBufferGetDecodeTimeStamp")
	tryRegister(&_CMSampleBufferGetDuration, lib, "CMSampleBufferGetDuration")
	tryRegister(&_CMSampleBufferGetFormatDescription, lib, "CMSampleBufferGetFormatDescription")
	tryRegister(&_CMSampleBufferGetImageBuffer, lib, "CMSampleBufferGetImageBuffer")
	tryRegister(&_CMSampleBufferGetNumSamples, lib, "CMSampleBufferGetNumSamples")
	tryRegister(&_CMSampleBufferGetOutputDecodeTimeStamp, lib, "CMSampleBufferGetOutputDecodeTimeStamp")
	tryRegister(&_CMSampleBufferGetOutputDuration, lib, "CMSampleBufferGetOutputDuration")
	tryRegister(&_CMSampleBufferGetOutputPresentationTimeStamp, lib, "CMSampleBufferGetOutputPresentationTimeStamp")
	tryRegister(&_CMSampleBufferGetOutputSampleTimingInfoArray, lib, "CMSampleBufferGetOutputSampleTimingInfoArray")
	tryRegister(&_CMSampleBufferGetSampleAttachmentsArray, lib, "CMSampleBufferGetSampleAttachmentsArray")
	tryRegister(&_CMSampleBufferGetSampleSize, lib, "CMSampleBufferGetSampleSize")
	tryRegister(&_CMSampleBufferGetSampleSizeArray, lib, "CMSampleBufferGetSampleSizeArray")
	tryRegister(&_CMSampleBufferGetSampleTimingInfo, lib, "CMSampleBufferGetSampleTimingInfo")
	tryRegister(&_CMSampleBufferGetTaggedBufferGroup, lib, "CMSampleBufferGetTaggedBufferGroup")
	tryRegister(&_CMSampleBufferGetTotalSampleSize, lib, "CMSampleBufferGetTotalSampleSize")
	tryRegister(&_CMSampleBufferHasDataFailed, lib, "CMSampleBufferHasDataFailed")
	tryRegister(&_CMSampleBufferInvalidate, lib, "CMSampleBufferInvalidate")
	tryRegister(&_CMSampleBufferIsValid, lib, "CMSampleBufferIsValid")
	tryRegister(&_CMSampleBufferMakeDataReady, lib, "CMSampleBufferMakeDataReady")
	tryRegister(&_CMSampleBufferSetDataBuffer, lib, "CMSampleBufferSetDataBuffer")
	tryRegister(&_CMSampleBufferSetDataBufferFromAudioBufferList, lib, "CMSampleBufferSetDataBufferFromAudioBufferList")
	tryRegister(&_CMSampleBufferSetDataFailed, lib, "CMSampleBufferSetDataFailed")
	tryRegister(&_CMSampleBufferSetDataReady, lib, "CMSampleBufferSetDataReady")
	tryRegister(&_CMSampleBufferSetInvalidateCallback, lib, "CMSampleBufferSetInvalidateCallback")
	tryRegister(&_CMSampleBufferSetInvalidateHandler, lib, "CMSampleBufferSetInvalidateHandler")
	tryRegister(&_CMSampleBufferSetOutputPresentationTimeStamp, lib, "CMSampleBufferSetOutputPresentationTimeStamp")
	tryRegister(&_CMSampleBufferTrackDataReadiness, lib, "CMSampleBufferTrackDataReadiness")
	tryRegister(&_CMSimpleQueueCreate, lib, "CMSimpleQueueCreate")
	tryRegister(&_CMSimpleQueueGetCapacity, lib, "CMSimpleQueueGetCapacity")
	tryRegister(&_CMSimpleQueueGetHead, lib, "CMSimpleQueueGetHead")
	tryRegister(&_CMSimpleQueueGetTypeID, lib, "CMSimpleQueueGetTypeID")
	tryRegister(&_CMSwapBigEndianClosedCaptionDescriptionToHost, lib, "CMSwapBigEndianClosedCaptionDescriptionToHost")
	tryRegister(&_CMSyncConvertTime, lib, "CMSyncConvertTime")
	tryRegister(&_CMTagCollectionCreateUnion, lib, "CMTagCollectionCreateUnion")
	tryRegister(&_CMTagCollectionGetTagsWithCategory, lib, "CMTagCollectionGetTagsWithCategory")
	tryRegister(&_CMTagCollectionGetTypeID, lib, "CMTagCollectionGetTypeID")
	tryRegister(&_CMTagCollectionRemoveAllTags, lib, "CMTagCollectionRemoveAllTags")
	tryRegister(&_CMTagGetFlagsValue, lib, "CMTagGetFlagsValue")
	tryRegister(&_CMTagGetSInt64Value, lib, "CMTagGetSInt64Value")
	tryRegister(&_CMTagMakeFromDictionary, lib, "CMTagMakeFromDictionary")
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, lib, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferAtIndex, lib, "CMTaggedBufferGroupGetCMSampleBufferAtIndex")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferForTagCollection, lib, "CMTaggedBufferGroupGetCMSampleBufferForTagCollection")
	tryRegister(&_CMTaggedBufferGroupGetCount, lib, "CMTaggedBufferGroupGetCount")
	tryRegister(&_CMTaggedBufferGroupGetTypeID, lib, "CMTaggedBufferGroupGetTypeID")
	tryRegister(&_CMTimeFoldIntoRange, lib, "CMTimeFoldIntoRange")
	tryRegister(&_CMTimeGetSeconds, lib, "CMTimeGetSeconds")
	tryRegister(&_CMTimeMakeWithSeconds, lib, "CMTimeMakeWithSeconds")
	tryRegister(&_CMTimeMappingCopyDescription, lib, "CMTimeMappingCopyDescription")
	tryRegister(&_CMTimeMappingMakeFromDictionary, lib, "CMTimeMappingMakeFromDictionary")
	tryRegister(&_CMTimeRangeCopyAsDictionary, lib, "CMTimeRangeCopyAsDictionary")
	tryRegister(&_CMTimebaseCreateWithMasterTimebase, lib, "CMTimebaseCreateWithMasterTimebase")
	tryRegister(&_CMTimebaseCreateWithSourceTimebase, lib, "CMTimebaseCreateWithSourceTimebase")
	tryRegister(&_CMTimebaseGetEffectiveRate, lib, "CMTimebaseGetEffectiveRate")
	tryRegister(&_CMTimebaseGetTimeAndRate, lib, "CMTimebaseGetTimeAndRate")
	tryRegister(&_CMTimebaseGetTypeID, lib, "CMTimebaseGetTypeID")
	tryRegister(&_CMVideoFormatDescriptionGetHEVCParameterSetAtIndex, lib, "CMVideoFormatDescriptionGetHEVCParameterSetAtIndex")
	tryRegister(&_CMTimebaseGetMasterClock, lib, "CMTimebaseGetMasterClock")
	tryRegister(&_CMPropagateAttachments, lib, "CMPropagateAttachments")
	tryRegister(&_CMTimebaseCopyMasterClock, lib, "CMTimebaseCopyMasterClock")
	tryRegister(&_CMSimpleQueueReset, lib, "CMSimpleQueueReset")
	tryRegister(&_CMTimeMakeFromDictionary, lib, "CMTimeMakeFromDictionary")
	tryRegister(&_CMTimebaseSetAnchorTime, lib, "CMTimebaseSetAnchorTime")
	tryRegister(&_CMTimeRangeMake, lib, "CMTimeRangeMake")
	tryRegister(&_CMTimeMapDurationFromRangeToRange, lib, "CMTimeMapDurationFromRangeToRange")
	tryRegister(&_CMBufferQueueMarkEndOfData, lib, "CMBufferQueueMarkEndOfData")
	tryRegister(&_CMTimeCodeFormatDescriptionGetFrameDuration, lib, "CMTimeCodeFormatDescriptionGetFrameDuration")
	tryRegister(&_CMTimebaseCopySourceTimebase, lib, "CMTimebaseCopySourceTimebase")
	tryRegister(&_CMTimeRangeFromTimeToTime, lib, "CMTimeRangeFromTimeToTime")
	tryRegister(&_CMBufferQueueDequeue, lib, "CMBufferQueueDequeue")
	tryRegister(&_CMBufferQueueGetBufferCount, lib, "CMBufferQueueGetBufferCount")
	tryRegister(&_CMBufferQueueGetMaxPresentationTimeStamp, lib, "CMBufferQueueGetMaxPresentationTimeStamp")
	tryRegister(&_CMTimebaseAddTimer, lib, "CMTimebaseAddTimer")
	tryRegister(&_CMClockInvalidate, lib, "CMClockInvalidate")
	tryRegister(&_CMTimeAbsoluteValue, lib, "CMTimeAbsoluteValue")
	tryRegister(&_CMTimeMappingMake, lib, "CMTimeMappingMake")
	tryRegister(&_CMTimeRangeGetUnion, lib, "CMTimeRangeGetUnion")
	tryRegister(&_CMTimebaseCopySourceClock, lib, "CMTimebaseCopySourceClock")
	tryRegister(&_CMTimeMappingShow, lib, "CMTimeMappingShow")
	tryRegister(&_CMBlockBufferGetTypeID, lib, "CMBlockBufferGetTypeID")
	tryRegister(&_CMTimebaseGetMasterTimebase, lib, "CMTimebaseGetMasterTimebase")
	tryRegister(&_CMTimebaseSetMasterClock, lib, "CMTimebaseSetMasterClock")
	tryRegister(&_CMTimebaseSetTime, lib, "CMTimebaseSetTime")
	tryRegister(&_CMBufferQueueGetTypeID, lib, "CMBufferQueueGetTypeID")
	tryRegister(&_CMTimeMinimum, lib, "CMTimeMinimum")
	tryRegister(&_CMTimebaseNotificationBarrier, lib, "CMTimebaseNotificationBarrier")
	tryRegister(&_CMBufferQueueDequeueIfDataReady, lib, "CMBufferQueueDequeueIfDataReady")
	tryRegister(&_CMBlockBufferAssureBlockMemory, lib, "CMBlockBufferAssureBlockMemory")
	tryRegister(&_CMBufferQueueGetFirstPresentationTimeStamp, lib, "CMBufferQueueGetFirstPresentationTimeStamp")
	tryRegister(&_CMTimeCopyAsDictionary, lib, "CMTimeCopyAsDictionary")
	tryRegister(&_CMTimebaseSetSourceTimebase, lib, "CMTimebaseSetSourceTimebase")
	tryRegister(&_CMTimeMapTimeFromRangeToRange, lib, "CMTimeMapTimeFromRangeToRange")
	tryRegister(&_CMTimebaseGetMaster, lib, "CMTimebaseGetMaster")
	tryRegister(&_CMTimebaseAddTimerDispatchSource, lib, "CMTimebaseAddTimerDispatchSource")
	tryRegister(&_CMTimebaseCopyUltimateMasterClock, lib, "CMTimebaseCopyUltimateMasterClock")
	tryRegister(&_CMTimebaseSetTimerDispatchSourceToFireImmediately, lib, "CMTimebaseSetTimerDispatchSourceToFireImmediately")
	tryRegister(&_CMBufferQueueRemoveTrigger, lib, "CMBufferQueueRemoveTrigger")
	tryRegister(&_CMFormatDescriptionGetTypeID, lib, "CMFormatDescriptionGetTypeID")
	tryRegister(&_CMMemoryPoolInvalidate, lib, "CMMemoryPoolInvalidate")
	tryRegister(&_CMTimebaseSetMasterTimebase, lib, "CMTimebaseSetMasterTimebase")
	tryRegister(&_CMSampleBufferGetTypeID, lib, "CMSampleBufferGetTypeID")
	tryRegister(&_CMTimeMaximum, lib, "CMTimeMaximum")
	tryRegister(&_CMBufferQueueSetValidationHandler, lib, "CMBufferQueueSetValidationHandler")
	tryRegister(&_CMMetadataFormatDescriptionGetIdentifiers, lib, "CMMetadataFormatDescriptionGetIdentifiers")
	tryRegister(&_CMTimeRangeGetIntersection, lib, "CMTimeRangeGetIntersection")
	tryRegister(&_CMMemoryPoolFlush, lib, "CMMemoryPoolFlush")
	tryRegister(&_CMTimebaseGetTime, lib, "CMTimebaseGetTime")
	tryRegister(&_CMTimeCopyDescription, lib, "CMTimeCopyDescription")
	tryRegister(&_CMTimebaseCopyMasterTimebase, lib, "CMTimebaseCopyMasterTimebase")
	tryRegister(&_CMFormatDescriptionGetExtensions, lib, "CMFormatDescriptionGetExtensions")
	tryRegister(&_CMCopyDictionaryOfAttachments, lib, "CMCopyDictionaryOfAttachments")
	tryRegister(&_CMSetAttachment, lib, "CMSetAttachment")
	tryRegister(&_CMSetAttachments, lib, "CMSetAttachments")
	tryRegister(&_CMBufferQueueEnqueue, lib, "CMBufferQueueEnqueue")
	tryRegister(&_CMTimeSubtract, lib, "CMTimeSubtract")
	tryRegister(&_CMTimebaseCopySource, lib, "CMTimebaseCopySource")
	tryRegister(&_CMBufferQueueGetHead, lib, "CMBufferQueueGetHead")
	tryRegister(&_CMFormatDescriptionGetMediaSubType, lib, "CMFormatDescriptionGetMediaSubType")
	tryRegister(&_CMBufferQueueGetMinPresentationTimeStamp, lib, "CMBufferQueueGetMinPresentationTimeStamp")
	tryRegister(&_CMFormatDescriptionGetExtension, lib, "CMFormatDescriptionGetExtension")
	tryRegister(&_CMMemoryPoolCreate, lib, "CMMemoryPoolCreate")
	tryRegister(&_CMTimebaseCopyUltimateSourceClock, lib, "CMTimebaseCopyUltimateSourceClock")
	tryRegister(&_CMMemoryPoolGetAllocator, lib, "CMMemoryPoolGetAllocator")
	tryRegister(&_CMTimebaseRemoveTimerDispatchSource, lib, "CMTimebaseRemoveTimerDispatchSource")
	tryRegister(&_CMBufferQueueReset, lib, "CMBufferQueueReset")
	tryRegister(&_CMTimebaseGetUltimateMasterClock, lib, "CMTimebaseGetUltimateMasterClock")
	tryRegister(&_CMTimeMappingMakeEmpty, lib, "CMTimeMappingMakeEmpty")
	tryRegister(&_CMTimebaseGetTimeWithTimeScale, lib, "CMTimebaseGetTimeWithTimeScale")
	tryRegister(&_CMBufferQueueGetFirstDecodeTimeStamp, lib, "CMBufferQueueGetFirstDecodeTimeStamp")
	tryRegister(&_CMBufferQueueGetDuration, lib, "CMBufferQueueGetDuration")
	tryRegister(&_CMMetadataFormatDescriptionGetKeyWithLocalID, lib, "CMMetadataFormatDescriptionGetKeyWithLocalID")
	tryRegister(&_CMTimebaseSetSourceClock, lib, "CMTimebaseSetSourceClock")
	tryRegister(&_CMTimeRangeCopyDescription, lib, "CMTimeRangeCopyDescription")
	tryRegister(&_CMBufferQueueGetMinDecodeTimeStamp, lib, "CMBufferQueueGetMinDecodeTimeStamp")
	tryRegister(&_CMSampleBufferGetPresentationTimeStamp, lib, "CMSampleBufferGetPresentationTimeStamp")
	tryRegister(&_CMTimeShow, lib, "CMTimeShow")
	tryRegister(&_CMRemoveAllAttachments, lib, "CMRemoveAllAttachments")
	tryRegister(&_CMRemoveAttachment, lib, "CMRemoveAttachment")
	tryRegister(&_CMBufferQueueGetEndPresentationTimeStamp, lib, "CMBufferQueueGetEndPresentationTimeStamp")
	tryRegister(&_CMSyncGetTime, lib, "CMSyncGetTime")
	tryRegister(&_CMTimebaseRemoveTimer, lib, "CMTimebaseRemoveTimer")
	tryRegister(&_CMTimeRangeMakeFromDictionary, lib, "CMTimeRangeMakeFromDictionary")
	tryRegister(&_CMTimeClampToRange, lib, "CMTimeClampToRange")
	tryRegister(&_CMTimeRangeGetEnd, lib, "CMTimeRangeGetEnd")
	tryRegister(&_CMTimeAdd, lib, "CMTimeAdd")
	tryRegister(&_CMFormatDescriptionGetMediaType, lib, "CMFormatDescriptionGetMediaType")
	tryRegister(&_CMTimeMappingCopyAsDictionary, lib, "CMTimeMappingCopyAsDictionary")
	tryRegister(&_CMTimebaseSetTimerToFireImmediately, lib, "CMTimebaseSetTimerToFireImmediately")
	tryRegister(&_CMClockGetTime, lib, "CMClockGetTime")
	tryRegister(&_CMClockGetTypeID, lib, "CMClockGetTypeID")
	tryRegister(&_CMVideoFormatDescriptionGetDimensions, lib, "CMVideoFormatDescriptionGetDimensions")
	tryRegister(&_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers, lib, "CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers")
	tryRegister(&_CMTimebaseCopyMaster, lib, "CMTimebaseCopyMaster")
	tryRegister(&_CMTimeRangeShow, lib, "CMTimeRangeShow")
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



// Creates a clock that tracks playback through a Core Audio device with the specified identifier. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreateFromAudioDeviceID(allocator:deviceID:clockOut:)
func CMAudioDeviceClockCreateFromAudioDeviceID(allocator unsafe.Pointer, deviceID unsafe.Pointer, clockOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockCreateFromAudioDeviceID(allocator, deviceID, clockOut)
	}


// Changes the Core Audio device the clock is tracking by specifying a new device identifier. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceID(_:deviceID:)
func CMAudioDeviceClockSetAudioDeviceID(clock unsafe.Pointer, deviceID unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceID(clock, deviceID)
	}


// Changes the Core Audio device the clock is tracking by specifying a new device unique identifier. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceUID(_:deviceUID:)
func CMAudioDeviceClockSetAudioDeviceUID(clock unsafe.Pointer, deviceUID unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceUID(clock, deviceUID)
	}


// Creates a sample buffer with packet descriptions. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator:dataBuffer:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator unsafe.Pointer, dataBuffer unsafe.Pointer, formatDescription unsafe.Pointer, numSamples unsafe.Pointer, presentationTimeStamp unsafe.Pointer, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator, dataBuffer, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
	}


// Creates a sample buffer with packet descriptions and a callback to make the data ready for use. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptions(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateWithPacketDescriptions(allocator unsafe.Pointer, dataBuffer unsafe.Pointer, dataReady unsafe.Pointer, makeDataReadyCallback unsafe.Pointer, makeDataReadyRefcon unsafe.Pointer, formatDescription unsafe.Pointer, numSamples unsafe.Pointer, presentationTimeStamp unsafe.Pointer, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioSampleBufferCreateWithPacketDescriptions(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
	}


// Creates a sample buffer with packet descriptions and a handler to make the data ready for use. [Full Topic]
//
// Added in macOS 10.14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:)
func CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator unsafe.Pointer, dataBuffer unsafe.Pointer, dataReady unsafe.Pointer, formatDescription unsafe.Pointer, numSamples unsafe.Pointer, presentationTimeStamp unsafe.Pointer, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler unsafe.Pointer) unsafe.Pointer {
	return _CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut, makeDataReadyHandler)
	}


// Creates a block buffer that contains a contiguous copy of, or reference to, the data specified by the parameters. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateContiguous(allocator:sourceBuffer:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateContiguous(structureAllocator unsafe.Pointer, sourceBuffer unsafe.Pointer, blockAllocator unsafe.Pointer, customBlockSource unsafe.Pointer, offsetToData unsafe.Pointer, dataLength unsafe.Pointer, flags unsafe.Pointer, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateContiguous(structureAllocator, sourceBuffer, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
	}


// Creates a block buffer that refers to another block buffer object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithBufferReference(allocator:referenceBuffer:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithBufferReference(structureAllocator unsafe.Pointer, bufferReference unsafe.Pointer, offsetToData unsafe.Pointer, dataLength unsafe.Pointer, flags unsafe.Pointer, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateWithBufferReference(structureAllocator, bufferReference, offsetToData, dataLength, flags, blockBufferOut)
	}


// Creates a block buffer that’s backed by a memory block. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithMemoryBlock(allocator:memoryBlock:blockLength:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithMemoryBlock(structureAllocator unsafe.Pointer, memoryBlock unsafe.Pointer, blockLength unsafe.Pointer, blockAllocator unsafe.Pointer, customBlockSource unsafe.Pointer, offsetToData unsafe.Pointer, dataLength unsafe.Pointer, flags unsafe.Pointer, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateWithMemoryBlock(structureAllocator, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
	}


// Returns the total length of data that’s accessible by a block buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataLength(_:)
func CMBlockBufferGetDataLength(theBuffer unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferGetDataLength(theBuffer)
	}


// Returns a Boolean value that indicates whether the specified range within a block buffer is contiguous. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsRangeContiguous(_:atOffset:length:)
func CMBlockBufferIsRangeContiguous(theBuffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferIsRangeContiguous(theBuffer, offset, length)
	}


// CMBufferQueueCopyHead is a CoreMedia function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueCopyHead(queue)
	}


// Dequeues a buffer from a queue. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeueAndRetain(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeueAndRetain(queue)
	}


// Dequeues a buffer from a queue, if it’s ready. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReadyAndRetain(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeueIfDataReadyAndRetain(queue)
	}


// Returns a reference to the singleton clock that reflects the host time. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetHostTimeClock()
func CMClockGetHostTimeClock() unsafe.Pointer {
	return _CMClockGetHostTimeClock()
	}


// Returns a Boolean value that indicates whether it’s possible for two clocks to drift relative to each other. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockMightDrift(_:otherClock:)
func CMClockMightDrift(clock unsafe.Pointer, otherClock unsafe.Pointer) unsafe.Pointer {
	return _CMClockMightDrift(clock, otherClock)
	}


// Returns the type identifier of memory pool objects. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() unsafe.Pointer {
	return _CMMemoryPoolGetTypeID()
	}


// Creates a copy of the key by using an identifier. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifier(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifier(allocator unsafe.Pointer, identifier unsafe.Pointer, keyOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateKeyFromIdentifier(allocator, identifier, keyOut)
	}


// Creates a copy of the keyspace by using an identifier. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeySpaceFromIdentifier(allocator:identifier:keySpaceOut:)
func CMMetadataCreateKeySpaceFromIdentifier(allocator unsafe.Pointer, identifier unsafe.Pointer, keySpaceOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateKeySpaceFromIdentifier(allocator, identifier, keySpaceOut)
	}


// Returns a Boolean value that indicates whether a data type conforms to another data type. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType unsafe.Pointer, conformsToDataType unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
	}


// Returns a Boolean value that indicates the registration status of a data type identifier. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsRegistered(_:)
func CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType)
	}


// Register a data type with the data type registry. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryRegisterDataType(_:description:conformingDataTypes:)
func CMMetadataDataTypeRegistryRegisterDataType(dataType unsafe.Pointer, description unsafe.Pointer, conformingDataTypes unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryRegisterDataType(dataType, description, conformingDataTypes)
	}


// Calls a block for every individual sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallBlockForEachSample(_:_:)
func CMSampleBufferCallBlockForEachSample(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCallBlockForEachSample(sbuf)
	}


// Calls a function for every individual sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallForEachSample(_:callback:refcon:)
func CMSampleBufferCallForEachSample(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCallForEachSample(sbuf)
	}


// Copies PCM audio data from a sample buffer into an audio buffer list. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopyPCMDataIntoAudioBufferList(_:at:frameCount:into:)
func CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf unsafe.Pointer, frameOffset unsafe.Pointer, numFrames unsafe.Pointer, bufferList unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf, frameOffset, numFrames, bufferList)
	}


// Creates a sample buffer that contains a range of samples from an existing sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopySampleBufferForRange(allocator:sampleBuffer:sampleRange:sampleBufferOut:)
func CMSampleBufferCopySampleBufferForRange(allocator unsafe.Pointer, sbuf unsafe.Pointer, sampleRange unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCopySampleBufferForRange(allocator, sbuf, sampleRange, sampleBufferOut)
	}


// Creates a copy of a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopy(allocator:sampleBuffer:sampleBufferOut:)
func CMSampleBufferCreateCopy(allocator unsafe.Pointer, sbuf unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateCopy(allocator, sbuf, sampleBufferOut)
	}


// Creates a copy of a sample buffer with new timing information. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopyWithNewTiming(allocator:sampleBuffer:sampleTimingEntryCount:sampleTimingArray:sampleBufferOut:)
func CMSampleBufferCreateCopyWithNewTiming(allocator unsafe.Pointer, originalSBuf unsafe.Pointer, numSampleTimingEntries unsafe.Pointer, sampleTimingArray unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateCopyWithNewTiming(allocator, originalSBuf, numSampleTimingEntries, sampleTimingArray, sampleBufferOut)
	}


// Creates a sample buffer with an image buffer and a callback to make the data ready for use. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBuffer(allocator:imageBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateForImageBuffer(allocator unsafe.Pointer, imageBuffer unsafe.Pointer, dataReady unsafe.Pointer, makeDataReadyCallback unsafe.Pointer, makeDataReadyRefcon unsafe.Pointer, formatDescription unsafe.Pointer, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateForImageBuffer(allocator, imageBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, sampleTiming, sampleBufferOut)
	}


// Creates a sample buffer with an image buffer and a handler to make the data ready for use. [Full Topic]
//
// Added in macOS 10.14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(_:_:_:_:_:_:_:)
func CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator unsafe.Pointer, imageBuffer unsafe.Pointer, dataReady unsafe.Pointer, formatDescription unsafe.Pointer, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator, imageBuffer, dataReady, formatDescription, sampleTiming, sampleBufferOut, makeDataReadyHandler)
	}


// Creates a sample buffer with media data. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReady(allocator:dataBuffer:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreateReady(allocator unsafe.Pointer, dataBuffer unsafe.Pointer, formatDescription unsafe.Pointer, numSamples unsafe.Pointer, numSampleTimingEntries unsafe.Pointer, sampleTimingArray unsafe.Pointer, numSampleSizeEntries unsafe.Pointer, sampleSizeArray unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateReady(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
	}


// Creates a sample buffer with image data. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReadyWithImageBuffer(allocator:imageBuffer:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateReadyWithImageBuffer(allocator unsafe.Pointer, imageBuffer unsafe.Pointer, formatDescription unsafe.Pointer, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateReadyWithImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sampleBufferOut)
	}


// Creates a sample buffer with a handler to make the data ready for use. [Full Topic]
//
// Added in macOS 10.14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateWithMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:_:_:)
func CMSampleBufferCreateWithMakeDataReadyHandler(allocator unsafe.Pointer, dataBuffer unsafe.Pointer, dataReady unsafe.Pointer, formatDescription unsafe.Pointer, numSamples unsafe.Pointer, numSampleTimingEntries unsafe.Pointer, sampleTimingArray unsafe.Pointer, numSampleSizeEntries unsafe.Pointer, sampleSizeArray unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateWithMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut, makeDataReadyHandler)
	}


// Returns a Boolean value that indicates whether the sample buffer’s data is ready for use. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferDataIsReady(_:)
func CMSampleBufferDataIsReady(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferDataIsReady(sbuf)
	}


// Returns an audio buffer list that contains the media data. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(_:bufferListSizeNeededOut:bufferListOut:bufferListSize:blockBufferAllocator:blockBufferMemoryAllocator:flags:blockBufferOut:)
func CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf unsafe.Pointer, bufferListSizeNeededOut unsafe.Pointer, bufferListOut unsafe.Pointer, bufferListSize unsafe.Pointer, blockBufferStructureAllocator unsafe.Pointer, blockBufferBlockAllocator unsafe.Pointer, flags unsafe.Pointer, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf, bufferListSizeNeededOut, bufferListOut, bufferListSize, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, blockBufferOut)
	}


// Creates an array of audio stream packet descriptions. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptions(_:allocatedSize:packetDescriptionsOut:packetDescriptionsSizeNeededOut:)
func CMSampleBufferGetAudioStreamPacketDescriptions(sbuf unsafe.Pointer, packetDescriptionsSize unsafe.Pointer, packetDescriptionsOut unsafe.Pointer, packetDescriptionsSizeNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioStreamPacketDescriptions(sbuf, packetDescriptionsSize, packetDescriptionsOut, packetDescriptionsSizeNeededOut)
	}


// Returns a pointer to a constant array of audio stream packet descriptions. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptionsPtr(_:packetDescriptionsPointerOut:sizeOut:)
func CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf unsafe.Pointer, packetDescriptionsPointerOut unsafe.Pointer, packetDescriptionsSizeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf, packetDescriptionsPointerOut, packetDescriptionsSizeOut)
	}


// Returns a block buffer that contains the media data. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDataBuffer(_:)
func CMSampleBufferGetDataBuffer(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDataBuffer(sbuf)
	}


// Returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDecodeTimeStamp(_:)
func CMSampleBufferGetDecodeTimeStamp(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDecodeTimeStamp(sbuf)
	}


// Returns the total duration of a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDuration(_:)
func CMSampleBufferGetDuration(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetDuration(sbuf)
	}


// Returns the format description of the samples in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetFormatDescription(_:)
func CMSampleBufferGetFormatDescription(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetFormatDescription(sbuf)
	}


// Returns an image buffer that contains the media data. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetImageBuffer(_:)
func CMSampleBufferGetImageBuffer(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetImageBuffer(sbuf)
	}


// Returns the number of media samples in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetNumSamples(_:)
func CMSampleBufferGetNumSamples(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetNumSamples(sbuf)
	}


// Returns the output decode timestamp of a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDecodeTimeStamp(_:)
func CMSampleBufferGetOutputDecodeTimeStamp(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputDecodeTimeStamp(sbuf)
	}


// Returns the output duration of a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDuration(_:)
func CMSampleBufferGetOutputDuration(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputDuration(sbuf)
	}


// Returns the output presentation timestamp of a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputPresentationTimeStamp(_:)
func CMSampleBufferGetOutputPresentationTimeStamp(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputPresentationTimeStamp(sbuf)
	}


// Retrieves an array of output timing information structures that represents each sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetOutputSampleTimingInfoArray(sbuf unsafe.Pointer, timingArrayEntries unsafe.Pointer, timingArrayOut unsafe.Pointer, timingArrayEntriesNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputSampleTimingInfoArray(sbuf, timingArrayEntries, timingArrayOut, timingArrayEntriesNeededOut)
	}


// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf unsafe.Pointer, createIfNecessary unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
	}


// Returns the size in bytes of a specified sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSize(_:at:)
func CMSampleBufferGetSampleSize(sbuf unsafe.Pointer, sampleIndex unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleSize(sbuf, sampleIndex)
	}


// Retrieves an array of sample sizes that represents each sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSizeArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleSizeArray(sbuf unsafe.Pointer, sizeArrayEntries unsafe.Pointer, sizeArrayOut unsafe.Pointer, sizeArrayEntriesNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleSizeArray(sbuf, sizeArrayEntries, sizeArrayOut, sizeArrayEntriesNeededOut)
	}


// Retrieves a timing information structure that describes a specified sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfo(_:at:timingInfoOut:)
func CMSampleBufferGetSampleTimingInfo(sbuf unsafe.Pointer, sampleIndex unsafe.Pointer, timingInfoOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleTimingInfo(sbuf, sampleIndex, timingInfoOut)
	}


// Gets the tagged buffer group of a sample buffer. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTaggedBufferGroup
func CMSampleBufferGetTaggedBufferGroup(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetTaggedBufferGroup(sbuf)
	}


// Returns the total size in bytes of sample data in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTotalSampleSize(_:)
func CMSampleBufferGetTotalSampleSize(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetTotalSampleSize(sbuf)
	}


// Returns a Boolean value that indicates whether the sample buffer’s data loading request failed. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferHasDataFailed(_:statusOut:)
func CMSampleBufferHasDataFailed(sbuf unsafe.Pointer, statusOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferHasDataFailed(sbuf, statusOut)
	}


// Invalidates a sample buffer by calling its invalidation callback. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidate(_:)
func CMSampleBufferInvalidate(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferInvalidate(sbuf)
	}


// Returns a Boolean value that indicates whether a sample buffer is valid. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferIsValid(_:)
func CMSampleBufferIsValid(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferIsValid(sbuf)
	}


// Makes the sample buffer’s data ready for use by invoking its callback to load the data. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReady(_:)
func CMSampleBufferMakeDataReady(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferMakeDataReady(sbuf)
	}


// Sets a block buffer of media data on a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBuffer(_:newValue:)
func CMSampleBufferSetDataBuffer(sbuf unsafe.Pointer, dataBuffer unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataBuffer(sbuf, dataBuffer)
	}


// Creates a block buffer that contains a copy of the data from an audio buffer list. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBufferFromAudioBufferList(_:blockBufferAllocator:blockBufferMemoryAllocator:flags:bufferList:)
func CMSampleBufferSetDataBufferFromAudioBufferList(sbuf unsafe.Pointer, blockBufferStructureAllocator unsafe.Pointer, blockBufferBlockAllocator unsafe.Pointer, flags unsafe.Pointer, bufferList unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataBufferFromAudioBufferList(sbuf, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, bufferList)
	}


// Marks the sample buffer’s data as failed to indicate that it won’t become ready. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataFailed(_:status:)
func CMSampleBufferSetDataFailed(sbuf unsafe.Pointer, status unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataFailed(sbuf, status)
	}


// Marks a sample buffer’s data as ready for use. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataReady(_:)
func CMSampleBufferSetDataReady(sbuf unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataReady(sbuf)
	}


// Sets the sample buffer’s invalidation callback. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateCallback(_:callback:refcon:)
func CMSampleBufferSetInvalidateCallback(sbuf unsafe.Pointer, invalidateCallback unsafe.Pointer, invalidateRefCon unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetInvalidateCallback(sbuf, invalidateCallback, invalidateRefCon)
	}


// Sets the sample buffer’s invalidation handler. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateHandler(_:invalidateHandler:)
func CMSampleBufferSetInvalidateHandler(sbuf unsafe.Pointer, invalidateHandler unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetInvalidateHandler(sbuf, invalidateHandler)
	}


// Sets an output presentation timestamp to use in place of a calculated value. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetOutputPresentationTimeStamp(_:newValue:)
func CMSampleBufferSetOutputPresentationTimeStamp(sbuf unsafe.Pointer, outputPresentationTimeStamp unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetOutputPresentationTimeStamp(sbuf, outputPresentationTimeStamp)
	}


// Associates a sample buffer’s data readiness with that of another sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferTrackDataReadiness(_:sampleBufferToTrack:)
func CMSampleBufferTrackDataReadiness(sbuf unsafe.Pointer, sampleBufferToTrack unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferTrackDataReadiness(sbuf, sampleBufferToTrack)
	}


// Creates a queue that has the specified capacity. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueCreate(allocator:capacity:queueOut:)
func CMSimpleQueueCreate(allocator unsafe.Pointer, capacity unsafe.Pointer, queueOut unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueCreate(allocator, capacity, queueOut)
	}


// Returns the number of elements that the queue can hold. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCapacity(_:)
func CMSimpleQueueGetCapacity(queue unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueGetCapacity(queue)
	}


// Returns the element at the head of the queue. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetHead(_:)
func CMSimpleQueueGetHead(queue unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueGetHead(queue)
	}


// Returns the type identifier of sample buffer objects. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() unsafe.Pointer {
	return _CMSimpleQueueGetTypeID()
	}


// Converts a closed caption description structure from big-endian to host-endian, in place. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianClosedCaptionDescriptionToHost(_:_:)
func CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData unsafe.Pointer, closedCaptionDescriptionSize unsafe.Pointer) unsafe.Pointer {
	return _CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize)
	}


// Converts a time from one timebase or clock to another timebase or clock. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time unsafe.Pointer, fromClockOrTimebase unsafe.Pointer, toClockOrTimebase unsafe.Pointer) unsafe.Pointer {
	return _CMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
	}


// Creates a new tag collection containing all tags from two collections without duplicates. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateUnion
func CMTagCollectionCreateUnion(tagCollection1 unsafe.Pointer, tagCollection2 unsafe.Pointer, tagCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateUnion(tagCollection1, tagCollection2, tagCollectionOut)
	}


// Retrieves a C-style array of tags with a given category from a tag collection. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithCategory
func CMTagCollectionGetTagsWithCategory(tagCollection unsafe.Pointer, category unsafe.Pointer, tagBuffer unsafe.Pointer, tagBufferCount unsafe.Pointer, numberOfTagsCopied unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionGetTagsWithCategory(tagCollection, category, tagBuffer, tagBufferCount, numberOfTagsCopied)
	}


// Retrieves the internal type ID for tag collections. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTypeID
func CMTagCollectionGetTypeID() unsafe.Pointer {
	return _CMTagCollectionGetTypeID()
	}


// Removes all tags from a collection. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionRemoveAllTags(tagCollection)
	}


// Retrieves a tag’s value as a 64-bit field flag. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetFlagsValue
func CMTagGetFlagsValue(tag unsafe.Pointer) unsafe.Pointer {
	return _CMTagGetFlagsValue(tag)
	}


// Retrieves a tag’s value as a signed 64-bit integer. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetSInt64Value
func CMTagGetSInt64Value(tag unsafe.Pointer) unsafe.Pointer {
	return _CMTagGetSInt64Value(tag)
	}


// Create a new tag from a dictionary object. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeFromDictionary
func CMTagMakeFromDictionary(dict unsafe.Pointer) unsafe.Pointer {
	return _CMTagMakeFromDictionary(dict)
	}


// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions is a CoreMedia function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator unsafe.Pointer, taggedBufferGroup unsafe.Pointer, extensions unsafe.Pointer, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
	}


// Gets the sample buffer at a given index in the buffer group. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
	}


// Gets the single sample buffer in a group which contains a given tag collection, if present. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTagCollection
func CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group unsafe.Pointer, tagCollection unsafe.Pointer, indexOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group, tagCollection, indexOut)
	}


// Gets the number of buffers contained within a tagged buffer group. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCount
func CMTaggedBufferGroupGetCount(group unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupGetCount(group)
	}


// Gets the internal type ID for a tagged buffer group. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTypeID
func CMTaggedBufferGroupGetTypeID() unsafe.Pointer {
	return _CMTaggedBufferGroupGetTypeID()
	}


// Folds a time into a time range. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFoldIntoRange(_:foldRange:)
func CMTimeFoldIntoRange(time unsafe.Pointer, foldRange unsafe.Pointer) unsafe.Pointer {
	return _CMTimeFoldIntoRange(time, foldRange)
	}


// Returns a representation of the time in seconds. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time unsafe.Pointer) unsafe.Pointer {
	return _CMTimeGetSeconds(time)
	}


// Creates a time that represents a number of seconds in a preferred timescale. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithSeconds(_:preferredTimescale:)
func CMTimeMakeWithSeconds(seconds unsafe.Pointer, preferredTimescale unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMakeWithSeconds(seconds, preferredTimescale)
	}


// Copies a string description of a time mapping. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator unsafe.Pointer, mapping unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingCopyDescription(allocator, mapping)
	}


// Creates a time mapping from a dictionary representation. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeFromDictionary(_:)
func CMTimeMappingMakeFromDictionary(dictionaryRepresentation unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMakeFromDictionary(dictionaryRepresentation)
	}


// Returns a dictionary representation of a time range. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ unsafe.Pointer, allocator unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeCopyAsDictionary(range_, allocator)
	}


// Creates a timebase by using a host timebase. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithMasterTimebase(allocator:masterTimebase:timebaseOut:)
func CMTimebaseCreateWithMasterTimebase(allocator unsafe.Pointer, masterTimebase unsafe.Pointer, timebaseOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCreateWithMasterTimebase(allocator, masterTimebase, timebaseOut)
	}


// Creates a timebase by using a source timebase. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceTimebase(allocator:sourceTimebase:timebaseOut:)
func CMTimebaseCreateWithSourceTimebase(allocator unsafe.Pointer, sourceTimebase unsafe.Pointer, timebaseOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCreateWithSourceTimebase(allocator, sourceTimebase, timebaseOut)
	}


// Returns the effective rate of a timebase, which combines its rate with the rates of all its host timebases. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetEffectiveRate(_:)
func CMTimebaseGetEffectiveRate(timebase unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetEffectiveRate(timebase)
	}


// Returns the current time and rate of a timebase. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase unsafe.Pointer, timeOut unsafe.Pointer, rateOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
	}


// Returns the Core Foundation type identifier that identifies a timebase object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() unsafe.Pointer {
	return _CMTimebaseGetTypeID()
	}


// Returns a parameter set contained in an HEVC (H.265) format description. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc unsafe.Pointer, parameterSetIndex unsafe.Pointer, parameterSetPointerOut unsafe.Pointer, parameterSetSizeOut unsafe.Pointer, parameterSetCountOut unsafe.Pointer, NALUnitHeaderLengthOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, NALUnitHeaderLengthOut)
	}


// Returns the immediate host clock of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterClock(_:)
func CMTimebaseGetMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMasterClock(p0)
	}


// Copies all propagable attachments from one attachment bearer object to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPropagateAttachments(_:destination:)
func CMPropagateAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMPropagateAttachments(p0)
	}


// Returns the immediate host clock of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterClock(_:)
func CMTimebaseCopyMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMasterClock(p0)
	}


// Resets the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueReset(_:)
func CMSimpleQueueReset(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueReset(p0)
	}


// Creates a time from a dictionary representation of its fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeFromDictionary(_:)
func CMTimeMakeFromDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMakeFromDictionary(p0)
	}


// Sets the time of a timebase at a particular source time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetAnchorTime(_:timebaseTime:immediateMasterTime:)
func CMTimebaseSetAnchorTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetAnchorTime(p0)
	}


// Creates a valid time range with a start time and duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMake(start:duration:)
func CMTimeRangeMake(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeMake(p0)
	}


// Translates a duration through a mapping from two time ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapDurationFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapDurationFromRangeToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMapDurationFromRangeToRange(p0)
	}


// Sets a marker to indicate this queue doesn’t allow enqueuing new buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueMarkEndOfData(_:)
func CMBufferQueueMarkEndOfData(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueMarkEndOfData(p0)
	}


// Returns the duration of each frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameDuration(_:)
func CMTimeCodeFormatDescriptionGetFrameDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionGetFrameDuration(p0)
	}


// Returns the immediate source timebase of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceTimebase(_:)
func CMTimebaseCopySourceTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySourceTimebase(p0)
	}


// Creates a valid time range from a start and end time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeFromTimeToTime(start:end:)
func CMTimeRangeFromTimeToTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeFromTimeToTime(p0)
	}


// Dequeues a buffer from a queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeue(p0)
	}


// Gets the number of buffers in the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetBufferCount(_:)
func CMBufferQueueGetBufferCount(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetBufferCount(p0)
	}


// Gets the greatest presentation timestamp of a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMaxPresentationTimeStamp(_:)
func CMBufferQueueGetMaxPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMaxPresentationTimeStamp(p0)
	}


// Adds the timer to the list of timers the timebase manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimer(_:timer:runloop:)
func CMTimebaseAddTimer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseAddTimer(p0)
	}


// Stops the clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockInvalidate(_:)
func CMClockInvalidate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMClockInvalidate(p0)
	}


// Returns the absolute value of a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAbsoluteValue(_:)
func CMTimeAbsoluteValue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeAbsoluteValue(p0)
	}


// Creates a time mapping with a source and target time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMake(source:target:)
func CMTimeMappingMake(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMake(p0)
	}


// Returns a new time range with the time elements of the input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetUnion(_:otherRange:)
func CMTimeRangeGetUnion(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetUnion(p0)
	}


// Returns the immediate source clock of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceClock(_:)
func CMTimebaseCopySourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySourceClock(p0)
	}


// Prints a description of a time mapping to standard output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingShow(_:)
func CMTimeMappingShow(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingShow(p0)
	}


// Returns the type identifier for block buffer objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetTypeID()
func CMBlockBufferGetTypeID() unsafe.Pointer {
	return _CMBlockBufferGetTypeID()
	}


// Returns the immediate host timebase of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterTimebase(_:)
func CMTimebaseGetMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMasterTimebase(p0)
	}


// Sets the time of a timebase at a particular source time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetMasterClock(_:_:)
func CMTimebaseSetMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetMasterClock(p0)
	}


// Sets the current time of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTime(_:time:)
func CMTimebaseSetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTime(p0)
	}


// Returns the type identifier of buffer queue objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTypeID()
func CMBufferQueueGetTypeID() unsafe.Pointer {
	return _CMBufferQueueGetTypeID()
	}


// Returns the lesser of two time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMinimum(_:_:)
func CMTimeMinimum(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMinimum(p0)
	}


// Requests that the timebase wait until it isn’t posting notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseNotificationBarrier(_:)
func CMTimebaseNotificationBarrier(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseNotificationBarrier(p0)
	}


// Dequeues a buffer from a queue, if it’s ready. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReady(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueDequeueIfDataReady(p0)
	}


// Assures that the system allocates memory for all memory blocks in a block buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAssureBlockMemory(_:)
func CMBlockBufferAssureBlockMemory(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferAssureBlockMemory(p0)
	}


// Gets the presentation timestamp of the first buffer in a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstPresentationTimeStamp(_:)
func CMBufferQueueGetFirstPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetFirstPresentationTimeStamp(p0)
	}


// Creates a dictionary representation of the time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyAsDictionary(_:allocator:)
func CMTimeCopyAsDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCopyAsDictionary(p0)
	}


// Sets the source timebase of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceTimebase(_:_:)
func CMTimebaseSetSourceTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetSourceTimebase(p0)
	}


// Translates a time through a mapping from two time ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapTimeFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapTimeFromRangeToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMapTimeFromRangeToRange(p0)
	}


// Returns the immediate host (either timebase or clock) of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMaster(_:)
func CMTimebaseGetMaster(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetMaster(p0)
	}


// Adds the timer dispatch source to the list of timers the timebase manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimerDispatchSource(_:timerSource:)
func CMTimebaseAddTimerDispatchSource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseAddTimerDispatchSource(p0)
	}


// Returns the host clock that is the host of all of a timebase’s host timebases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateMasterClock(_:)
func CMTimebaseCopyUltimateMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyUltimateMasterClock(p0)
	}


// Sets the timer dispatch source to fire immediately once, overriding any previous timer call. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceToFireImmediately(_:timerSource:)
func CMTimebaseSetTimerDispatchSourceToFireImmediately(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTimerDispatchSourceToFireImmediately(p0)
	}


// Removes a previously installed trigger from a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueRemoveTrigger(_:triggerToken:)
func CMBufferQueueRemoveTrigger(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueRemoveTrigger(p0)
	}


// Returns the Core Foundation type identifier that identifies format description objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetTypeID()
func CMFormatDescriptionGetTypeID() unsafe.Pointer {
	return _CMFormatDescriptionGetTypeID()
	}


// Invalidates the memory pool, which causes its allocator to stop recycling memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolInvalidate(_:)
func CMMemoryPoolInvalidate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolInvalidate(p0)
	}


// CMTimebaseSetMasterTimebase is a CoreMedia function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetMasterTimebase(_:_:)
func CMTimebaseSetMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetMasterTimebase(p0)
	}


// Returns the type identifier of sample buffer objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTypeID()
func CMSampleBufferGetTypeID() unsafe.Pointer {
	return _CMSampleBufferGetTypeID()
	}


// Returns the greater of two time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMaximum(_:_:)
func CMTimeMaximum(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMaximum(p0)
	}


// A validation handler for the queue to call before enqueuing buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationHandler(_:_:)
func CMBufferQueueSetValidationHandler(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueSetValidationHandler(p0)
	}


// Returns an array of metadata identifiers from a metadata format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetIdentifiers(_:)
func CMMetadataFormatDescriptionGetIdentifiers(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionGetIdentifiers(p0)
	}


// Returns a new time range with the time elements that are common between the input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetIntersection(_:otherRange:)
func CMTimeRangeGetIntersection(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetIntersection(p0)
	}


// Deallocates all memory the pool holds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolFlush(_:)
func CMMemoryPoolFlush(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolFlush(p0)
	}


// Returns the current time from a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTime(_:)
func CMTimebaseGetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTime(p0)
	}


// Creates a string representation of the time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyDescription(allocator:time:)
func CMTimeCopyDescription(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCopyDescription(p0)
	}


// Returns the immediate host timebase of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterTimebase(_:)
func CMTimebaseCopyMasterTimebase(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMasterTimebase(p0)
	}


// Returns all of the extensions for a format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtensions(_:)
func CMFormatDescriptionGetExtensions(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetExtensions(p0)
	}


// Returns a dictionary of all attachments for an attachment bearer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMCopyDictionaryOfAttachments(allocator:target:attachmentMode:)
func CMCopyDictionaryOfAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMCopyDictionaryOfAttachments(p0)
	}


// Sets or adds an attachment to an attachment bearer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachment(_:key:value:attachmentMode:)
func CMSetAttachment(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSetAttachment(p0)
	}


// Sets a dictionary of attachments on an attachment bearer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachments(_:attachments:attachmentMode:)
func CMSetAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSetAttachments(p0)
	}


// Enqueues a buffer onto a queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)
func CMBufferQueueEnqueue(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueEnqueue(p0)
	}


// Returns the difference between two times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeSubtract(_:_:)
func CMTimeSubtract(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeSubtract(p0)
	}


// Returns the immediate source — either a clock or timebase — of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySource(_:)
func CMTimebaseCopySource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopySource(p0)
	}


// Retrieves the next buffer from a queue, but doesn’t remove it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetHead(_:)
func CMBufferQueueGetHead(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetHead(p0)
	}


// Returns the media subtype of a format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
func CMFormatDescriptionGetMediaSubType(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetMediaSubType(p0)
	}


// Gets the earliest presentation timestamp of a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinPresentationTimeStamp(_:)
func CMBufferQueueGetMinPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMinPresentationTimeStamp(p0)
	}


// Returns an extension from the format description by using an extension key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtension(_:extensionKey:)
func CMFormatDescriptionGetExtension(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetExtension(p0)
	}


// Creates a memory pool. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolCreate(options:)
func CMMemoryPoolCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolCreate(p0)
	}


// Returns the source clock that’s the source of all of a timebase’s source timebases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateSourceClock(_:)
func CMTimebaseCopyUltimateSourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyUltimateSourceClock(p0)
	}


// Returns the allocator for the memory pool. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetAllocator(_:)
func CMMemoryPoolGetAllocator(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMemoryPoolGetAllocator(p0)
	}


// Removes the timer dispatch source from the list of timers the timebase manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimerDispatchSource(_:timerSource:)
func CMTimebaseRemoveTimerDispatchSource(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseRemoveTimerDispatchSource(p0)
	}


// Resets a buffer queue, which allows it to enqueue new buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueReset(_:)
func CMBufferQueueReset(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueReset(p0)
	}


// Returns the host clock that is the host of all of a timebase’s host timebases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetUltimateMasterClock(_:)
func CMTimebaseGetUltimateMasterClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetUltimateMasterClock(p0)
	}


// Creates a valid time mapping with an empty source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeEmpty(target:)
func CMTimeMappingMakeEmpty(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingMakeEmpty(p0)
	}


// Returns the current time from a timebase in the specified timescale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeWithTimeScale(_:timescale:method:)
func CMTimebaseGetTimeWithTimeScale(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeWithTimeScale(p0)
	}


// Gets the decode timestamp of the first buffer in a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstDecodeTimeStamp(_:)
func CMBufferQueueGetFirstDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetFirstDecodeTimeStamp(p0)
	}


// Gets the duration of a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetDuration(_:)
func CMBufferQueueGetDuration(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetDuration(p0)
	}


// Returns the key for the local identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetKeyWithLocalID(_:localKeyID:)
func CMMetadataFormatDescriptionGetKeyWithLocalID(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionGetKeyWithLocalID(p0)
	}


// Sets the source clock of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceClock(_:_:)
func CMTimebaseSetSourceClock(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetSourceClock(p0)
	}


// Returns a string with a description of a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyDescription(allocator:range:)
func CMTimeRangeCopyDescription(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeCopyDescription(p0)
	}


// Gets the earliest decode timestamp of a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinDecodeTimeStamp(_:)
func CMBufferQueueGetMinDecodeTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetMinDecodeTimeStamp(p0)
	}


// Returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetPresentationTimeStamp(_:)
func CMSampleBufferGetPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetPresentationTimeStamp(p0)
	}


// Prints a description of the time to the console. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeShow(_:)
func CMTimeShow(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeShow(p0)
	}


// Removes all attachments from an attachment bearer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAllAttachments(_:)
func CMRemoveAllAttachments(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMRemoveAllAttachments(p0)
	}


// Removes a specific attachment from an attachment bearer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAttachment(_:key:)
func CMRemoveAttachment(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMRemoveAttachment(p0)
	}


// Gets the greatest end presentation timestamp of a buffer queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetEndPresentationTimeStamp(_:)
func CMBufferQueueGetEndPresentationTimeStamp(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueGetEndPresentationTimeStamp(p0)
	}


// Returns the time from a clock or timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncGetTime(_:)
func CMSyncGetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMSyncGetTime(p0)
	}


// Removes the timer from the list of timers the timebase manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimer(_:timer:)
func CMTimebaseRemoveTimer(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseRemoveTimer(p0)
	}


// Creates a time range from a dictionary representation of its fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMakeFromDictionary(_:)
func CMTimeRangeMakeFromDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeMakeFromDictionary(p0)
	}


// Returns the nearest time value inside the time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeClampToRange(_:range:)
func CMTimeClampToRange(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeClampToRange(p0)
	}


// Returns a time value that represents the end of a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetEnd(_:)
func CMTimeRangeGetEnd(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeGetEnd(p0)
	}


// Returns the sum of two times. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAdd(_:_:)
func CMTimeAdd(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeAdd(p0)
	}


// Returns the media type of a format description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaType(_:)
func CMFormatDescriptionGetMediaType(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionGetMediaType(p0)
	}


// Returns a dictionary representation of a time mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyAsDictionary(_:allocator:)
func CMTimeMappingCopyAsDictionary(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingCopyAsDictionary(p0)
	}


// Sets the timer to fire immediately once, overriding any previous timer calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerToFireImmediately(_:timer:)
func CMTimebaseSetTimerToFireImmediately(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTimerToFireImmediately(p0)
	}


// Returns the current time from a clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTime(_:)
func CMClockGetTime(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMClockGetTime(p0)
	}


// Returns the core foundation type identifier of a clock type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTypeID()
func CMClockGetTypeID() unsafe.Pointer {
	return _CMClockGetTypeID()
	}


// Returns the video dimensions, in encoded pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetDimensions(_:)
func CMVideoFormatDescriptionGetDimensions(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionGetDimensions(p0)
	}


// Returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
func CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() unsafe.Pointer {
	return _CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
	}


// Returns the immediate host timebase of a timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMaster(_:)
func CMTimebaseCopyMaster(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCopyMaster(p0)
	}


// Prints a description of the time range to standard error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeShow(_:)
func CMTimeRangeShow(p0 unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeShow(p0)
	}




