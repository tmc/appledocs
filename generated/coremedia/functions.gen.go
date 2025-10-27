// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia


import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
	foundation "github.com/tmc/appledocs/generated/foundation"
)


// CoreMedia Functions (370 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(AllocatorRef, TaggedBufferGroupRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioClockCreate func(AllocatorRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockCreate func(AllocatorRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockCreateFromAudioDeviceID func(AllocatorRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockGetAudioDevice func(ClockRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockSetAudioDeviceID func(ClockRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioDeviceClockSetAudioDeviceUID func(ClockRef, StringRef) unsafe.Pointer
	_CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer func(AllocatorRef, AudioFormatDescriptionRef, SoundDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionCreate func(AllocatorRef, unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, SoundDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, SoundDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionCreateSummary func(AllocatorRef, ArrayRef, uint32, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionEqual func(AudioFormatDescriptionRef, AudioFormatDescriptionRef, AudioFormatDescriptionMask, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionGetChannelLayout func(AudioFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionGetFormatList func(AudioFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionGetMagicCookie func(AudioFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMAudioFormatDescriptionGetMostCompatibleFormat func(AudioFormatDescriptionRef) unsafe.Pointer
	_CMAudioFormatDescriptionGetRichestDecodableFormat func(AudioFormatDescriptionRef) unsafe.Pointer
	_CMAudioFormatDescriptionGetStreamBasicDescription func(AudioFormatDescriptionRef) unsafe.Pointer
	_CMAudioSampleBufferCreateReadyWithPacketDescriptions func(AllocatorRef, BlockBufferRef, FormatDescriptionRef, ItemCount, Time, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioSampleBufferCreateWithPacketDescriptions func(AllocatorRef, BlockBufferRef, unsafe.Pointer, SampleBufferMakeDataReadyCallback, unsafe.Pointer, FormatDescriptionRef, ItemCount, Time, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler func(AllocatorRef, BlockBufferRef, unsafe.Pointer, FormatDescriptionRef, ItemCount, Time, unsafe.Pointer, unsafe.Pointer, SampleBufferMakeDataReadyHandler) unsafe.Pointer
	_CMBlockBufferAccessDataBytes func(BlockBufferRef, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferAppendBufferReference func(BlockBufferRef, BlockBufferRef, uintptr, uintptr, BlockBufferFlags) unsafe.Pointer
	_CMBlockBufferAppendMemoryBlock func(BlockBufferRef, unsafe.Pointer, uintptr, AllocatorRef, unsafe.Pointer, uintptr, uintptr, BlockBufferFlags) unsafe.Pointer
	_CMBlockBufferAssureBlockMemory func(BlockBufferRef) unsafe.Pointer
	_CMBlockBufferCopyDataBytes func(BlockBufferRef, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateContiguous func(AllocatorRef, BlockBufferRef, AllocatorRef, unsafe.Pointer, uintptr, uintptr, BlockBufferFlags, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateEmpty func(AllocatorRef, uint32, BlockBufferFlags, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateWithBufferReference func(AllocatorRef, BlockBufferRef, uintptr, uintptr, BlockBufferFlags, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferCreateWithMemoryBlock func(AllocatorRef, unsafe.Pointer, uintptr, AllocatorRef, unsafe.Pointer, uintptr, uintptr, BlockBufferFlags, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferFillDataBytes func(unsafe.Pointer, BlockBufferRef, uintptr, uintptr) unsafe.Pointer
	_CMBlockBufferGetDataLength func(BlockBufferRef) uintptr
	_CMBlockBufferGetDataPointer func(BlockBufferRef, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBlockBufferGetTypeID func() TypeID
	_CMBlockBufferIsEmpty func(BlockBufferRef) unsafe.Pointer
	_CMBlockBufferIsRangeContiguous func(BlockBufferRef, uintptr, uintptr) unsafe.Pointer
	_CMBlockBufferReplaceDataBytes func(unsafe.Pointer, BlockBufferRef, uintptr, uintptr) unsafe.Pointer
	_CMBufferQueueCallForEachBuffer func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueContainsEndOfData func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueCopyHead func(BufferQueueRef) BufferRef
	_CMBufferQueueCreate func(AllocatorRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueCreateWithHandlers func(AllocatorRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueDequeueAndRetain func(BufferQueueRef) BufferRef
	_CMBufferQueueDequeueIfDataReadyAndRetain func(BufferQueueRef) BufferRef
	_CMBufferQueueEnqueue func(BufferQueueRef, BufferRef) unsafe.Pointer
	_CMBufferQueueGetBufferCount func(BufferQueueRef) ItemCount
	_CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS func() unsafe.Pointer
	_CMBufferQueueGetCallbacksForUnsortedSampleBuffers func() unsafe.Pointer
	_CMBufferQueueGetDuration func(BufferQueueRef) Time
	_CMBufferQueueGetEndPresentationTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetFirstDecodeTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetFirstPresentationTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetHead func(BufferQueueRef) BufferRef
	_CMBufferQueueGetMaxPresentationTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetMinDecodeTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetMinPresentationTimeStamp func(BufferQueueRef) Time
	_CMBufferQueueGetTotalSize func(BufferQueueRef) uintptr
	_CMBufferQueueGetTypeID func() TypeID
	_CMBufferQueueInstallTrigger func(BufferQueueRef, BufferQueueTriggerCallback, unsafe.Pointer, BufferQueueTriggerCondition, Time, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueInstallTriggerHandler func(BufferQueueRef, BufferQueueTriggerCondition, Time, unsafe.Pointer, BufferQueueTriggerHandler) unsafe.Pointer
	_CMBufferQueueInstallTriggerHandlerWithIntegerThreshold func(BufferQueueRef, BufferQueueTriggerCondition, ItemCount, unsafe.Pointer, BufferQueueTriggerHandler) unsafe.Pointer
	_CMBufferQueueInstallTriggerWithIntegerThreshold func(BufferQueueRef, BufferQueueTriggerCallback, unsafe.Pointer, BufferQueueTriggerCondition, ItemCount, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueIsAtEndOfData func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueIsEmpty func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueMarkEndOfData func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueRemoveTrigger func(BufferQueueRef, BufferQueueTriggerToken) unsafe.Pointer
	_CMBufferQueueReset func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueResetWithCallback func(BufferQueueRef) unsafe.Pointer
	_CMBufferQueueSetValidationCallback func(BufferQueueRef, BufferValidationCallback, unsafe.Pointer) unsafe.Pointer
	_CMBufferQueueSetValidationHandler func(BufferQueueRef, BufferValidationHandler) unsafe.Pointer
	_CMBufferQueueTestTrigger func(BufferQueueRef, BufferQueueTriggerToken) unsafe.Pointer
	_CMClockConvertHostTimeToSystemUnits func(Time) uint64
	_CMClockGetAnchorTime func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMClockGetHostTimeClock func() ClockRef
	_CMClockGetTime func(ClockRef) Time
	_CMClockGetTypeID func() TypeID
	_CMClockInvalidate func(ClockRef)
	_CMClockMakeHostTimeFromSystemUnits func(uint64) Time
	_CMClockMightDrift func(ClockRef, ClockRef) unsafe.Pointer
	_CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer func(AllocatorRef, ClosedCaptionFormatDescriptionRef, ClosedCaptionDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, ClosedCaptionDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, ClosedCaptionDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMCopyDictionaryOfAttachments func(AllocatorRef, AttachmentBearerRef, AttachmentMode) DictionaryRef
	_CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout func(BlockBufferRef, SoundDescriptionFlavor) unsafe.Pointer
	_CMFormatDescriptionCreate func(AllocatorRef, MediaType, unsafe.Pointer, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMFormatDescriptionEqual func(FormatDescriptionRef, FormatDescriptionRef) unsafe.Pointer
	_CMFormatDescriptionEqualIgnoringExtensionKeys func(FormatDescriptionRef, FormatDescriptionRef, TypeRef, TypeRef) unsafe.Pointer
	_CMFormatDescriptionGetExtension func(FormatDescriptionRef, StringRef) PropertyListRef
	_CMFormatDescriptionGetExtensions func(FormatDescriptionRef) DictionaryRef
	_CMFormatDescriptionGetMediaSubType func(FormatDescriptionRef) unsafe.Pointer
	_CMFormatDescriptionGetMediaType func(FormatDescriptionRef) MediaType
	_CMFormatDescriptionGetTypeID func() TypeID
	_CMGetAttachment func(AttachmentBearerRef, StringRef, unsafe.Pointer) TypeRef
	_CMMemoryPoolCreate func(DictionaryRef) MemoryPoolRef
	_CMMemoryPoolFlush func(MemoryPoolRef)
	_CMMemoryPoolGetAllocator func(MemoryPoolRef) AllocatorRef
	_CMMemoryPoolGetTypeID func() TypeID
	_CMMemoryPoolInvalidate func(MemoryPoolRef)
	_CMMetadataCreateIdentifierForKeyAndKeySpace func(AllocatorRef, TypeRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataCreateKeyFromIdentifier func(AllocatorRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataCreateKeyFromIdentifierAsCFData func(AllocatorRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataCreateKeySpaceFromIdentifier func(AllocatorRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeConformsToDataType func(StringRef, StringRef) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeIsBaseDataType func(StringRef) unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeIsRegistered func(StringRef) unsafe.Pointer
	_CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType func(StringRef) StringRef
	_CMMetadataDataTypeRegistryGetBaseDataTypes func() ArrayRef
	_CMMetadataDataTypeRegistryGetConformingDataTypes func(StringRef) ArrayRef
	_CMMetadataDataTypeRegistryGetDataTypeDescription func(StringRef) StringRef
	_CMMetadataDataTypeRegistryRegisterDataType func(StringRef, StringRef, ArrayRef) unsafe.Pointer
	_CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer func(AllocatorRef, MetadataFormatDescriptionRef, MetadataDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions func(AllocatorRef, MetadataFormatDescriptionRef, MetadataFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, MetadataDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, MetadataDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateWithKeys func(AllocatorRef, MetadataFormatType, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications func(AllocatorRef, MetadataFormatDescriptionRef, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionCreateWithMetadataSpecifications func(AllocatorRef, MetadataFormatType, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CMMetadataFormatDescriptionGetIdentifiers func(MetadataFormatDescriptionRef) ArrayRef
	_CMMetadataFormatDescriptionGetKeyWithLocalID func(MetadataFormatDescriptionRef, unsafe.Pointer) DictionaryRef
	_CMMuxedFormatDescriptionCreate func(AllocatorRef, MuxedStreamType, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMPropagateAttachments func(AttachmentBearerRef, AttachmentBearerRef)
	_CMRemoveAllAttachments func(AttachmentBearerRef)
	_CMRemoveAttachment func(AttachmentBearerRef, StringRef)
	_CMSampleBufferCallBlockForEachSample func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferCallForEachSample func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferCopyPCMDataIntoAudioBufferList func(SampleBufferRef, int32, int32, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCopySampleBufferForRange func(AllocatorRef, SampleBufferRef, foundation.Range, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreate func(AllocatorRef, BlockBufferRef, unsafe.Pointer, SampleBufferMakeDataReadyCallback, unsafe.Pointer, FormatDescriptionRef, ItemCount, ItemCount, unsafe.Pointer, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateCopy func(AllocatorRef, SampleBufferRef, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateCopyWithNewTiming func(AllocatorRef, SampleBufferRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateForImageBuffer func(AllocatorRef, ImageBufferRef, unsafe.Pointer, SampleBufferMakeDataReadyCallback, unsafe.Pointer, VideoFormatDescriptionRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler func(AllocatorRef, ImageBufferRef, unsafe.Pointer, VideoFormatDescriptionRef, unsafe.Pointer, unsafe.Pointer, SampleBufferMakeDataReadyHandler) unsafe.Pointer
	_CMSampleBufferCreateForTaggedBufferGroup func(AllocatorRef, TaggedBufferGroupRef, Time, Time, TaggedBufferGroupFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateReady func(AllocatorRef, BlockBufferRef, FormatDescriptionRef, ItemCount, ItemCount, unsafe.Pointer, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateReadyWithImageBuffer func(AllocatorRef, ImageBufferRef, VideoFormatDescriptionRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferCreateWithMakeDataReadyHandler func(AllocatorRef, BlockBufferRef, unsafe.Pointer, FormatDescriptionRef, ItemCount, ItemCount, unsafe.Pointer, ItemCount, unsafe.Pointer, unsafe.Pointer, SampleBufferMakeDataReadyHandler) unsafe.Pointer
	_CMSampleBufferDataIsReady func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer func(SampleBufferRef, unsafe.Pointer, unsafe.Pointer, uintptr, AllocatorRef, AllocatorRef, uint32, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetAudioStreamPacketDescriptions func(SampleBufferRef, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetAudioStreamPacketDescriptionsPtr func(SampleBufferRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetDataBuffer func(SampleBufferRef) BlockBufferRef
	_CMSampleBufferGetDecodeTimeStamp func(SampleBufferRef) Time
	_CMSampleBufferGetDuration func(SampleBufferRef) Time
	_CMSampleBufferGetFormatDescription func(SampleBufferRef) FormatDescriptionRef
	_CMSampleBufferGetImageBuffer func(SampleBufferRef) ImageBufferRef
	_CMSampleBufferGetNumSamples func(SampleBufferRef) ItemCount
	_CMSampleBufferGetOutputDecodeTimeStamp func(SampleBufferRef) Time
	_CMSampleBufferGetOutputDuration func(SampleBufferRef) Time
	_CMSampleBufferGetOutputPresentationTimeStamp func(SampleBufferRef) Time
	_CMSampleBufferGetOutputSampleTimingInfoArray func(SampleBufferRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetPresentationTimeStamp func(SampleBufferRef) Time
	_CMSampleBufferGetSampleAttachmentsArray func(SampleBufferRef, unsafe.Pointer) ArrayRef
	_CMSampleBufferGetSampleSize func(SampleBufferRef, ItemIndex) uintptr
	_CMSampleBufferGetSampleSizeArray func(SampleBufferRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleTimingInfo func(SampleBufferRef, ItemIndex, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleTimingInfoArray func(SampleBufferRef, ItemCount, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetTaggedBufferGroup func(SampleBufferRef) TaggedBufferGroupRef
	_CMSampleBufferGetTotalSampleSize func(SampleBufferRef) uintptr
	_CMSampleBufferGetTypeID func() TypeID
	_CMSampleBufferHasDataFailed func(SampleBufferRef, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferInvalidate func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferIsValid func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferMakeDataReady func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferSetDataBuffer func(SampleBufferRef, BlockBufferRef) unsafe.Pointer
	_CMSampleBufferSetDataBufferFromAudioBufferList func(SampleBufferRef, AllocatorRef, AllocatorRef, uint32, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataFailed func(SampleBufferRef, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferSetDataReady func(SampleBufferRef) unsafe.Pointer
	_CMSampleBufferSetInvalidateCallback func(SampleBufferRef, SampleBufferInvalidateCallback, uint64) unsafe.Pointer
	_CMSampleBufferSetInvalidateHandler func(SampleBufferRef, SampleBufferInvalidateHandler) unsafe.Pointer
	_CMSampleBufferSetOutputPresentationTimeStamp func(SampleBufferRef, Time) unsafe.Pointer
	_CMSampleBufferTrackDataReadiness func(SampleBufferRef, SampleBufferRef) unsafe.Pointer
	_CMSetAttachment func(AttachmentBearerRef, StringRef, TypeRef, AttachmentMode)
	_CMSetAttachments func(AttachmentBearerRef, DictionaryRef, AttachmentMode)
	_CMSimpleQueueCreate func(AllocatorRef, int32, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueDequeue func(SimpleQueueRef) unsafe.Pointer
	_CMSimpleQueueEnqueue func(SimpleQueueRef, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetCapacity func(SimpleQueueRef) int32
	_CMSimpleQueueGetCount func(SimpleQueueRef) int32
	_CMSimpleQueueGetHead func(SimpleQueueRef) unsafe.Pointer
	_CMSimpleQueueGetTypeID func() TypeID
	_CMSimpleQueueReset func(SimpleQueueRef) unsafe.Pointer
	_CMSwapBigEndianClosedCaptionDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapBigEndianImageDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapBigEndianMetadataDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapBigEndianSoundDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapBigEndianTextDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapBigEndianTimeCodeDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianClosedCaptionDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianImageDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianMetadataDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianSoundDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianTextDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSwapHostEndianTimeCodeDescriptionToBig func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSyncConvertTime func(Time, ClockOrTimebaseRef, ClockOrTimebaseRef) Time
	_CMSyncGetRelativeRate func(ClockOrTimebaseRef, ClockOrTimebaseRef) unsafe.Pointer
	_CMSyncGetRelativeRateAndAnchorTime func(ClockOrTimebaseRef, ClockOrTimebaseRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSyncGetTime func(ClockOrTimebaseRef) Time
	_CMSyncMightDrift func(ClockOrTimebaseRef, ClockOrTimebaseRef) unsafe.Pointer
	_CMTagCollectionAddTag func(MutableTagCollectionRef, Tag) unsafe.Pointer
	_CMTagCollectionAddTagsFromArray func(MutableTagCollectionRef, unsafe.Pointer, ItemCount) unsafe.Pointer
	_CMTagCollectionAddTagsFromCollection func(MutableTagCollectionRef, TagCollectionRef) unsafe.Pointer
	_CMTagCollectionApply func(TagCollectionRef, TagCollectionApplierFunction, unsafe.Pointer)
	_CMTagCollectionApplyUntil func(TagCollectionRef, TagCollectionTagFilterFunction, unsafe.Pointer) Tag
	_CMTagCollectionContainsCategory func(TagCollectionRef, TagCategory) unsafe.Pointer
	_CMTagCollectionContainsSpecifiedTags func(TagCollectionRef, unsafe.Pointer, ItemCount) unsafe.Pointer
	_CMTagCollectionContainsTag func(TagCollectionRef, Tag) unsafe.Pointer
	_CMTagCollectionContainsTagsOfCollection func(TagCollectionRef, TagCollectionRef) unsafe.Pointer
	_CMTagCollectionCopyAsData func(TagCollectionRef, AllocatorRef) DataRef
	_CMTagCollectionCopyAsDictionary func(TagCollectionRef, AllocatorRef) DictionaryRef
	_CMTagCollectionCopyDescription func(AllocatorRef, TagCollectionRef) StringRef
	_CMTagCollectionCopyTagsOfCategories func(AllocatorRef, TagCollectionRef, unsafe.Pointer, ItemCount, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCountTagsWithFilterFunction func(TagCollectionRef, TagCollectionTagFilterFunction, unsafe.Pointer) ItemCount
	_CMTagCollectionCreate func(AllocatorRef, unsafe.Pointer, ItemCount, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateCopy func(TagCollectionRef, AllocatorRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateDifference func(TagCollectionRef, TagCollectionRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateExclusiveOr func(TagCollectionRef, TagCollectionRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateFromData func(DataRef, AllocatorRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateFromDictionary func(DictionaryRef, AllocatorRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateIntersection func(TagCollectionRef, TagCollectionRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateMutable func(AllocatorRef, Index, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateMutableCopy func(TagCollectionRef, AllocatorRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionCreateUnion func(TagCollectionRef, TagCollectionRef, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetCount func(TagCollectionRef) ItemCount
	_CMTagCollectionGetCountOfCategory func(TagCollectionRef, TagCategory) ItemCount
	_CMTagCollectionGetTags func(TagCollectionRef, unsafe.Pointer, ItemCount, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetTagsWithCategory func(TagCollectionRef, TagCategory, unsafe.Pointer, ItemCount, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetTagsWithFilterFunction func(TagCollectionRef, unsafe.Pointer, ItemCount, unsafe.Pointer, TagCollectionTagFilterFunction, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionGetTypeID func() TypeID
	_CMTagCollectionIsEmpty func(TagCollectionRef) unsafe.Pointer
	_CMTagCollectionRemoveAllTags func(MutableTagCollectionRef) unsafe.Pointer
	_CMTagCollectionRemoveAllTagsOfCategory func(MutableTagCollectionRef, TagCategory) unsafe.Pointer
	_CMTagCollectionRemoveTag func(MutableTagCollectionRef, Tag) unsafe.Pointer
	_CMTagCompare func(Tag, Tag) ComparisonResult
	_CMTagCopyAsDictionary func(Tag, AllocatorRef) DictionaryRef
	_CMTagCopyDescription func(AllocatorRef, Tag) StringRef
	_CMTagEqualToTag func(Tag, Tag) unsafe.Pointer
	_CMTaggedBufferGroupCreate func(AllocatorRef, ArrayRef, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupCreateCombined func(AllocatorRef, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup func(AllocatorRef, TaggedBufferGroupRef, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup func(TaggedBufferGroupFormatDescriptionRef, TaggedBufferGroupRef) unsafe.Pointer
	_CMTaggedBufferGroupGetCMSampleBufferAtIndex func(TaggedBufferGroupRef, Index) SampleBufferRef
	_CMTaggedBufferGroupGetCMSampleBufferForTag func(TaggedBufferGroupRef, Tag, unsafe.Pointer) SampleBufferRef
	_CMTaggedBufferGroupGetCMSampleBufferForTagCollection func(TaggedBufferGroupRef, TagCollectionRef, unsafe.Pointer) SampleBufferRef
	_CMTaggedBufferGroupGetCount func(TaggedBufferGroupRef) ItemCount
	_CMTaggedBufferGroupGetCVPixelBufferAtIndex func(TaggedBufferGroupRef, Index) PixelBufferRef
	_CMTaggedBufferGroupGetCVPixelBufferForTag func(TaggedBufferGroupRef, Tag, unsafe.Pointer) PixelBufferRef
	_CMTaggedBufferGroupGetCVPixelBufferForTagCollection func(TaggedBufferGroupRef, TagCollectionRef, unsafe.Pointer) PixelBufferRef
	_CMTaggedBufferGroupGetNumberOfMatchesForTagCollection func(TaggedBufferGroupRef, TagCollectionRef) ItemCount
	_CMTaggedBufferGroupGetTagCollectionAtIndex func(TaggedBufferGroupRef, Index) TagCollectionRef
	_CMTaggedBufferGroupGetTypeID func() TypeID
	_CMTagGetFlagsValue func(Tag) uint64
	_CMTagGetFloat64Value func(Tag) unsafe.Pointer
	_CMTagGetOSTypeValue func(Tag) unsafe.Pointer
	_CMTagGetSInt64Value func(Tag) int64
	_CMTagGetValueDataType func(Tag) TagDataType
	_CMTagHasFlagsValue func(Tag) unsafe.Pointer
	_CMTagHasFloat64Value func(Tag) unsafe.Pointer
	_CMTagHash func(Tag) HashCode
	_CMTagHasOSTypeValue func(Tag) unsafe.Pointer
	_CMTagHasSInt64Value func(Tag) unsafe.Pointer
	_CMTagMakeFromDictionary func(DictionaryRef) Tag
	_CMTagMakeWithFlagsValue func(TagCategory, uint64) Tag
	_CMTagMakeWithFloat64Value func(TagCategory, unsafe.Pointer) Tag
	_CMTagMakeWithOSTypeValue func(TagCategory, unsafe.Pointer) Tag
	_CMTagMakeWithSInt64Value func(TagCategory, int64) Tag
	_CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer func(AllocatorRef, TextFormatDescriptionRef, TextDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, TextDescriptionFlavor, MediaType, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, TextDescriptionFlavor, MediaType, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionGetDefaultStyle func(FormatDescriptionRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []float64, float64, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionGetDefaultTextBox func(FormatDescriptionRef, unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionGetDisplayFlags func(FormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionGetFontName func(FormatDescriptionRef, uint16, unsafe.Pointer) unsafe.Pointer
	_CMTextFormatDescriptionGetJustification func(FormatDescriptionRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeAbsoluteValue func(Time) Time
	_CMTimeAdd func(Time, Time) Time
	_CMTimebaseAddTimer func(TimebaseRef, RunLoopTimerRef, RunLoopRef) unsafe.Pointer
	_CMTimebaseAddTimerDispatchSource func(TimebaseRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCopyMasterClock func(TimebaseRef) ClockRef
	_CMTimebaseCopyMasterTimebase func(TimebaseRef) TimebaseRef
	_CMTimebaseCopySource func(TimebaseRef) ClockOrTimebaseRef
	_CMTimebaseCopySourceClock func(TimebaseRef) ClockRef
	_CMTimebaseCopyUltimateMasterClock func(TimebaseRef) ClockRef
	_CMTimebaseCopyUltimateSourceClock func(TimebaseRef) ClockRef
	_CMTimebaseCreateWithMasterClock func(AllocatorRef, ClockRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCreateWithMasterTimebase func(AllocatorRef, TimebaseRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseCreateWithSourceClock func(AllocatorRef, ClockRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetEffectiveRate func(TimebaseRef) unsafe.Pointer
	_CMTimebaseGetMaster func(TimebaseRef) ClockOrTimebaseRef
	_CMTimebaseGetMasterClock func(TimebaseRef) ClockRef
	_CMTimebaseGetMasterTimebase func(TimebaseRef) TimebaseRef
	_CMTimebaseGetRate func(TimebaseRef) unsafe.Pointer
	_CMTimebaseGetTime func(TimebaseRef) Time
	_CMTimebaseGetTimeAndRate func(TimebaseRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeWithTimeScale func(TimebaseRef, TimeScale, TimeRoundingMethod) Time
	_CMTimebaseGetTypeID func() TypeID
	_CMTimebaseGetUltimateMasterClock func(TimebaseRef) ClockRef
	_CMTimebaseNotificationBarrier func(TimebaseRef) unsafe.Pointer
	_CMTimebaseRemoveTimer func(TimebaseRef, RunLoopTimerRef) unsafe.Pointer
	_CMTimebaseRemoveTimerDispatchSource func(TimebaseRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetAnchorTime func(TimebaseRef, Time, Time) unsafe.Pointer
	_CMTimebaseSetRate func(TimebaseRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetRateAndAnchorTime func(TimebaseRef, unsafe.Pointer, Time, Time) unsafe.Pointer
	_CMTimebaseSetSourceClock func(TimebaseRef, ClockRef) unsafe.Pointer
	_CMTimebaseSetSourceTimebase func(TimebaseRef, TimebaseRef) unsafe.Pointer
	_CMTimebaseSetTime func(TimebaseRef, Time) unsafe.Pointer
	_CMTimebaseSetTimerDispatchSourceNextFireTime func(TimebaseRef, unsafe.Pointer, Time, uint32) unsafe.Pointer
	_CMTimebaseSetTimerDispatchSourceToFireImmediately func(TimebaseRef, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseSetTimerNextFireTime func(TimebaseRef, RunLoopTimerRef, Time, uint32) unsafe.Pointer
	_CMTimebaseSetTimerToFireImmediately func(TimebaseRef, RunLoopTimerRef) unsafe.Pointer
	_CMTimeClampToRange func(Time, TimeRange) Time
	_CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer func(AllocatorRef, TimeCodeFormatDescriptionRef, TimeCodeDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionCreate func(AllocatorRef, TimeCodeFormatType, Time, uint32, uint32, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, TimeCodeDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, TimeCodeDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMTimeCodeFormatDescriptionGetFrameDuration func(TimeCodeFormatDescriptionRef) Time
	_CMTimeCodeFormatDescriptionGetFrameQuanta func(TimeCodeFormatDescriptionRef) uint32
	_CMTimeCodeFormatDescriptionGetTimeCodeFlags func(TimeCodeFormatDescriptionRef) uint32
	_CMTimeCompare func(Time, Time) int32
	_CMTimeConvertScale func(Time, int32, TimeRoundingMethod) Time
	_CMTimeCopyAsDictionary func(Time, AllocatorRef) DictionaryRef
	_CMTimeCopyDescription func(AllocatorRef, Time) StringRef
	_CMTimeFoldIntoRange func(Time, TimeRange) Time
	_CMTimeGetSeconds func(Time) unsafe.Pointer
	_CMTimeMake func(int64, int32) Time
	_CMTimeMakeFromDictionary func(DictionaryRef) Time
	_CMTimeMakeWithEpoch func(int64, int32, int64) Time
	_CMTimeMakeWithSeconds func(unsafe.Pointer, int32) Time
	_CMTimeMapDurationFromRangeToRange func(Time, TimeRange, TimeRange) Time
	_CMTimeMappingCopyAsDictionary func(TimeMapping, AllocatorRef) DictionaryRef
	_CMTimeMappingCopyDescription func(AllocatorRef, TimeMapping) StringRef
	_CMTimeMappingMake func(TimeRange, TimeRange) TimeMapping
	_CMTimeMappingMakeEmpty func(TimeRange) TimeMapping
	_CMTimeMappingMakeFromDictionary func(DictionaryRef) TimeMapping
	_CMTimeMappingShow func(TimeMapping)
	_CMTimeMapTimeFromRangeToRange func(Time, TimeRange, TimeRange) Time
	_CMTimeMaximum func(Time, Time) Time
	_CMTimeMinimum func(Time, Time) Time
	_CMTimeMultiply func(Time, int32) Time
	_CMTimeMultiplyByFloat64 func(Time, unsafe.Pointer) Time
	_CMTimeMultiplyByRatio func(Time, int32, int32) Time
	_CMTimeRangeContainsTime func(TimeRange, Time) unsafe.Pointer
	_CMTimeRangeContainsTimeRange func(TimeRange, TimeRange) unsafe.Pointer
	_CMTimeRangeCopyAsDictionary func(TimeRange, AllocatorRef) DictionaryRef
	_CMTimeRangeCopyDescription func(AllocatorRef, TimeRange) StringRef
	_CMTimeRangeEqual func(TimeRange, TimeRange) unsafe.Pointer
	_CMTimeRangeFromTimeToTime func(Time, Time) TimeRange
	_CMTimeRangeGetEnd func(TimeRange) Time
	_CMTimeRangeGetIntersection func(TimeRange, TimeRange) TimeRange
	_CMTimeRangeGetUnion func(TimeRange, TimeRange) TimeRange
	_CMTimeRangeMake func(Time, Time) TimeRange
	_CMTimeRangeMakeFromDictionary func(DictionaryRef) TimeRange
	_CMTimeRangeShow func(TimeRange)
	_CMTimeShow func(Time)
	_CMTimeSubtract func(Time, Time) Time
	_CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer func(AllocatorRef, VideoFormatDescriptionRef, StringEncoding, ImageDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCopyTagCollectionArray func(VideoFormatDescriptionRef, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreate func(AllocatorRef, VideoCodecType, int32, int32, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreateForImageBuffer func(AllocatorRef, ImageBufferRef, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer func(AllocatorRef, BlockBufferRef, StringEncoding, ImageDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData func(AllocatorRef, unsafe.Pointer, uintptr, StringEncoding, ImageDescriptionFlavor, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreateFromH264ParameterSets func(AllocatorRef, uintptr, unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionCreateFromHEVCParameterSets func(AllocatorRef, uintptr, unsafe.Pointer, unsafe.Pointer, int, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CMVideoFormatDescriptionGetCleanAperture func(VideoFormatDescriptionRef, unsafe.Pointer) corefoundation.CGRect
	_CMVideoFormatDescriptionGetDimensions func(VideoFormatDescriptionRef) VideoDimensions
	_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers func() ArrayRef
	_CMVideoFormatDescriptionGetH264ParameterSetAtIndex func(FormatDescriptionRef, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []int) unsafe.Pointer
	_CMVideoFormatDescriptionGetHEVCParameterSetAtIndex func(FormatDescriptionRef, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []int) unsafe.Pointer
	_CMVideoFormatDescriptionGetPresentationDimensions func(VideoFormatDescriptionRef, unsafe.Pointer, unsafe.Pointer) corefoundation.CGSize
	_CMVideoFormatDescriptionMatchesImageBuffer func(VideoFormatDescriptionRef, ImageBufferRef) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, lib, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions")
	tryRegister(&_CMAudioClockCreate, lib, "CMAudioClockCreate")
	tryRegister(&_CMAudioDeviceClockCreate, lib, "CMAudioDeviceClockCreate")
	tryRegister(&_CMAudioDeviceClockCreateFromAudioDeviceID, lib, "CMAudioDeviceClockCreateFromAudioDeviceID")
	tryRegister(&_CMAudioDeviceClockGetAudioDevice, lib, "CMAudioDeviceClockGetAudioDevice")
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceID, lib, "CMAudioDeviceClockSetAudioDeviceID")
	tryRegister(&_CMAudioDeviceClockSetAudioDeviceUID, lib, "CMAudioDeviceClockSetAudioDeviceUID")
	tryRegister(&_CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer, lib, "CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer")
	tryRegister(&_CMAudioFormatDescriptionCreate, lib, "CMAudioFormatDescriptionCreate")
	tryRegister(&_CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer, lib, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer")
	tryRegister(&_CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData, lib, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData")
	tryRegister(&_CMAudioFormatDescriptionCreateSummary, lib, "CMAudioFormatDescriptionCreateSummary")
	tryRegister(&_CMAudioFormatDescriptionEqual, lib, "CMAudioFormatDescriptionEqual")
	tryRegister(&_CMAudioFormatDescriptionGetChannelLayout, lib, "CMAudioFormatDescriptionGetChannelLayout")
	tryRegister(&_CMAudioFormatDescriptionGetFormatList, lib, "CMAudioFormatDescriptionGetFormatList")
	tryRegister(&_CMAudioFormatDescriptionGetMagicCookie, lib, "CMAudioFormatDescriptionGetMagicCookie")
	tryRegister(&_CMAudioFormatDescriptionGetMostCompatibleFormat, lib, "CMAudioFormatDescriptionGetMostCompatibleFormat")
	tryRegister(&_CMAudioFormatDescriptionGetRichestDecodableFormat, lib, "CMAudioFormatDescriptionGetRichestDecodableFormat")
	tryRegister(&_CMAudioFormatDescriptionGetStreamBasicDescription, lib, "CMAudioFormatDescriptionGetStreamBasicDescription")
	tryRegister(&_CMAudioSampleBufferCreateReadyWithPacketDescriptions, lib, "CMAudioSampleBufferCreateReadyWithPacketDescriptions")
	tryRegister(&_CMAudioSampleBufferCreateWithPacketDescriptions, lib, "CMAudioSampleBufferCreateWithPacketDescriptions")
	tryRegister(&_CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler, lib, "CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler")
	tryRegister(&_CMBlockBufferAccessDataBytes, lib, "CMBlockBufferAccessDataBytes")
	tryRegister(&_CMBlockBufferAppendBufferReference, lib, "CMBlockBufferAppendBufferReference")
	tryRegister(&_CMBlockBufferAppendMemoryBlock, lib, "CMBlockBufferAppendMemoryBlock")
	tryRegister(&_CMBlockBufferAssureBlockMemory, lib, "CMBlockBufferAssureBlockMemory")
	tryRegister(&_CMBlockBufferCopyDataBytes, lib, "CMBlockBufferCopyDataBytes")
	tryRegister(&_CMBlockBufferCreateContiguous, lib, "CMBlockBufferCreateContiguous")
	tryRegister(&_CMBlockBufferCreateEmpty, lib, "CMBlockBufferCreateEmpty")
	tryRegister(&_CMBlockBufferCreateWithBufferReference, lib, "CMBlockBufferCreateWithBufferReference")
	tryRegister(&_CMBlockBufferCreateWithMemoryBlock, lib, "CMBlockBufferCreateWithMemoryBlock")
	tryRegister(&_CMBlockBufferFillDataBytes, lib, "CMBlockBufferFillDataBytes")
	tryRegister(&_CMBlockBufferGetDataLength, lib, "CMBlockBufferGetDataLength")
	tryRegister(&_CMBlockBufferGetDataPointer, lib, "CMBlockBufferGetDataPointer")
	tryRegister(&_CMBlockBufferGetTypeID, lib, "CMBlockBufferGetTypeID")
	tryRegister(&_CMBlockBufferIsEmpty, lib, "CMBlockBufferIsEmpty")
	tryRegister(&_CMBlockBufferIsRangeContiguous, lib, "CMBlockBufferIsRangeContiguous")
	tryRegister(&_CMBlockBufferReplaceDataBytes, lib, "CMBlockBufferReplaceDataBytes")
	tryRegister(&_CMBufferQueueCallForEachBuffer, lib, "CMBufferQueueCallForEachBuffer")
	tryRegister(&_CMBufferQueueContainsEndOfData, lib, "CMBufferQueueContainsEndOfData")
	tryRegister(&_CMBufferQueueCopyHead, lib, "CMBufferQueueCopyHead")
	tryRegister(&_CMBufferQueueCreate, lib, "CMBufferQueueCreate")
	tryRegister(&_CMBufferQueueCreateWithHandlers, lib, "CMBufferQueueCreateWithHandlers")
	tryRegister(&_CMBufferQueueDequeueAndRetain, lib, "CMBufferQueueDequeueAndRetain")
	tryRegister(&_CMBufferQueueDequeueIfDataReadyAndRetain, lib, "CMBufferQueueDequeueIfDataReadyAndRetain")
	tryRegister(&_CMBufferQueueEnqueue, lib, "CMBufferQueueEnqueue")
	tryRegister(&_CMBufferQueueGetBufferCount, lib, "CMBufferQueueGetBufferCount")
	tryRegister(&_CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS, lib, "CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS")
	tryRegister(&_CMBufferQueueGetCallbacksForUnsortedSampleBuffers, lib, "CMBufferQueueGetCallbacksForUnsortedSampleBuffers")
	tryRegister(&_CMBufferQueueGetDuration, lib, "CMBufferQueueGetDuration")
	tryRegister(&_CMBufferQueueGetEndPresentationTimeStamp, lib, "CMBufferQueueGetEndPresentationTimeStamp")
	tryRegister(&_CMBufferQueueGetFirstDecodeTimeStamp, lib, "CMBufferQueueGetFirstDecodeTimeStamp")
	tryRegister(&_CMBufferQueueGetFirstPresentationTimeStamp, lib, "CMBufferQueueGetFirstPresentationTimeStamp")
	tryRegister(&_CMBufferQueueGetHead, lib, "CMBufferQueueGetHead")
	tryRegister(&_CMBufferQueueGetMaxPresentationTimeStamp, lib, "CMBufferQueueGetMaxPresentationTimeStamp")
	tryRegister(&_CMBufferQueueGetMinDecodeTimeStamp, lib, "CMBufferQueueGetMinDecodeTimeStamp")
	tryRegister(&_CMBufferQueueGetMinPresentationTimeStamp, lib, "CMBufferQueueGetMinPresentationTimeStamp")
	tryRegister(&_CMBufferQueueGetTotalSize, lib, "CMBufferQueueGetTotalSize")
	tryRegister(&_CMBufferQueueGetTypeID, lib, "CMBufferQueueGetTypeID")
	tryRegister(&_CMBufferQueueInstallTrigger, lib, "CMBufferQueueInstallTrigger")
	tryRegister(&_CMBufferQueueInstallTriggerHandler, lib, "CMBufferQueueInstallTriggerHandler")
	tryRegister(&_CMBufferQueueInstallTriggerHandlerWithIntegerThreshold, lib, "CMBufferQueueInstallTriggerHandlerWithIntegerThreshold")
	tryRegister(&_CMBufferQueueInstallTriggerWithIntegerThreshold, lib, "CMBufferQueueInstallTriggerWithIntegerThreshold")
	tryRegister(&_CMBufferQueueIsAtEndOfData, lib, "CMBufferQueueIsAtEndOfData")
	tryRegister(&_CMBufferQueueIsEmpty, lib, "CMBufferQueueIsEmpty")
	tryRegister(&_CMBufferQueueMarkEndOfData, lib, "CMBufferQueueMarkEndOfData")
	tryRegister(&_CMBufferQueueRemoveTrigger, lib, "CMBufferQueueRemoveTrigger")
	tryRegister(&_CMBufferQueueReset, lib, "CMBufferQueueReset")
	tryRegister(&_CMBufferQueueResetWithCallback, lib, "CMBufferQueueResetWithCallback")
	tryRegister(&_CMBufferQueueSetValidationCallback, lib, "CMBufferQueueSetValidationCallback")
	tryRegister(&_CMBufferQueueSetValidationHandler, lib, "CMBufferQueueSetValidationHandler")
	tryRegister(&_CMBufferQueueTestTrigger, lib, "CMBufferQueueTestTrigger")
	tryRegister(&_CMClockConvertHostTimeToSystemUnits, lib, "CMClockConvertHostTimeToSystemUnits")
	tryRegister(&_CMClockGetAnchorTime, lib, "CMClockGetAnchorTime")
	tryRegister(&_CMClockGetHostTimeClock, lib, "CMClockGetHostTimeClock")
	tryRegister(&_CMClockGetTime, lib, "CMClockGetTime")
	tryRegister(&_CMClockGetTypeID, lib, "CMClockGetTypeID")
	tryRegister(&_CMClockInvalidate, lib, "CMClockInvalidate")
	tryRegister(&_CMClockMakeHostTimeFromSystemUnits, lib, "CMClockMakeHostTimeFromSystemUnits")
	tryRegister(&_CMClockMightDrift, lib, "CMClockMightDrift")
	tryRegister(&_CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer, lib, "CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer")
	tryRegister(&_CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer, lib, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer")
	tryRegister(&_CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData, lib, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData")
	tryRegister(&_CMCopyDictionaryOfAttachments, lib, "CMCopyDictionaryOfAttachments")
	tryRegister(&_CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout, lib, "CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout")
	tryRegister(&_CMFormatDescriptionCreate, lib, "CMFormatDescriptionCreate")
	tryRegister(&_CMFormatDescriptionEqual, lib, "CMFormatDescriptionEqual")
	tryRegister(&_CMFormatDescriptionEqualIgnoringExtensionKeys, lib, "CMFormatDescriptionEqualIgnoringExtensionKeys")
	tryRegister(&_CMFormatDescriptionGetExtension, lib, "CMFormatDescriptionGetExtension")
	tryRegister(&_CMFormatDescriptionGetExtensions, lib, "CMFormatDescriptionGetExtensions")
	tryRegister(&_CMFormatDescriptionGetMediaSubType, lib, "CMFormatDescriptionGetMediaSubType")
	tryRegister(&_CMFormatDescriptionGetMediaType, lib, "CMFormatDescriptionGetMediaType")
	tryRegister(&_CMFormatDescriptionGetTypeID, lib, "CMFormatDescriptionGetTypeID")
	tryRegister(&_CMGetAttachment, lib, "CMGetAttachment")
	tryRegister(&_CMMemoryPoolCreate, lib, "CMMemoryPoolCreate")
	tryRegister(&_CMMemoryPoolFlush, lib, "CMMemoryPoolFlush")
	tryRegister(&_CMMemoryPoolGetAllocator, lib, "CMMemoryPoolGetAllocator")
	tryRegister(&_CMMemoryPoolGetTypeID, lib, "CMMemoryPoolGetTypeID")
	tryRegister(&_CMMemoryPoolInvalidate, lib, "CMMemoryPoolInvalidate")
	tryRegister(&_CMMetadataCreateIdentifierForKeyAndKeySpace, lib, "CMMetadataCreateIdentifierForKeyAndKeySpace")
	tryRegister(&_CMMetadataCreateKeyFromIdentifier, lib, "CMMetadataCreateKeyFromIdentifier")
	tryRegister(&_CMMetadataCreateKeyFromIdentifierAsCFData, lib, "CMMetadataCreateKeyFromIdentifierAsCFData")
	tryRegister(&_CMMetadataCreateKeySpaceFromIdentifier, lib, "CMMetadataCreateKeySpaceFromIdentifier")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeConformsToDataType, lib, "CMMetadataDataTypeRegistryDataTypeConformsToDataType")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeIsBaseDataType, lib, "CMMetadataDataTypeRegistryDataTypeIsBaseDataType")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeIsRegistered, lib, "CMMetadataDataTypeRegistryDataTypeIsRegistered")
	tryRegister(&_CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType, lib, "CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType")
	tryRegister(&_CMMetadataDataTypeRegistryGetBaseDataTypes, lib, "CMMetadataDataTypeRegistryGetBaseDataTypes")
	tryRegister(&_CMMetadataDataTypeRegistryGetConformingDataTypes, lib, "CMMetadataDataTypeRegistryGetConformingDataTypes")
	tryRegister(&_CMMetadataDataTypeRegistryGetDataTypeDescription, lib, "CMMetadataDataTypeRegistryGetDataTypeDescription")
	tryRegister(&_CMMetadataDataTypeRegistryRegisterDataType, lib, "CMMetadataDataTypeRegistryRegisterDataType")
	tryRegister(&_CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer, lib, "CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer")
	tryRegister(&_CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions, lib, "CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions")
	tryRegister(&_CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer, lib, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer")
	tryRegister(&_CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData, lib, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData")
	tryRegister(&_CMMetadataFormatDescriptionCreateWithKeys, lib, "CMMetadataFormatDescriptionCreateWithKeys")
	tryRegister(&_CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications, lib, "CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications")
	tryRegister(&_CMMetadataFormatDescriptionCreateWithMetadataSpecifications, lib, "CMMetadataFormatDescriptionCreateWithMetadataSpecifications")
	tryRegister(&_CMMetadataFormatDescriptionGetIdentifiers, lib, "CMMetadataFormatDescriptionGetIdentifiers")
	tryRegister(&_CMMetadataFormatDescriptionGetKeyWithLocalID, lib, "CMMetadataFormatDescriptionGetKeyWithLocalID")
	tryRegister(&_CMMuxedFormatDescriptionCreate, lib, "CMMuxedFormatDescriptionCreate")
	tryRegister(&_CMPropagateAttachments, lib, "CMPropagateAttachments")
	tryRegister(&_CMRemoveAllAttachments, lib, "CMRemoveAllAttachments")
	tryRegister(&_CMRemoveAttachment, lib, "CMRemoveAttachment")
	tryRegister(&_CMSampleBufferCallBlockForEachSample, lib, "CMSampleBufferCallBlockForEachSample")
	tryRegister(&_CMSampleBufferCallForEachSample, lib, "CMSampleBufferCallForEachSample")
	tryRegister(&_CMSampleBufferCopyPCMDataIntoAudioBufferList, lib, "CMSampleBufferCopyPCMDataIntoAudioBufferList")
	tryRegister(&_CMSampleBufferCopySampleBufferForRange, lib, "CMSampleBufferCopySampleBufferForRange")
	tryRegister(&_CMSampleBufferCreate, lib, "CMSampleBufferCreate")
	tryRegister(&_CMSampleBufferCreateCopy, lib, "CMSampleBufferCreateCopy")
	tryRegister(&_CMSampleBufferCreateCopyWithNewTiming, lib, "CMSampleBufferCreateCopyWithNewTiming")
	tryRegister(&_CMSampleBufferCreateForImageBuffer, lib, "CMSampleBufferCreateForImageBuffer")
	tryRegister(&_CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler, lib, "CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler")
	tryRegister(&_CMSampleBufferCreateForTaggedBufferGroup, lib, "CMSampleBufferCreateForTaggedBufferGroup")
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
	tryRegister(&_CMSampleBufferGetPresentationTimeStamp, lib, "CMSampleBufferGetPresentationTimeStamp")
	tryRegister(&_CMSampleBufferGetSampleAttachmentsArray, lib, "CMSampleBufferGetSampleAttachmentsArray")
	tryRegister(&_CMSampleBufferGetSampleSize, lib, "CMSampleBufferGetSampleSize")
	tryRegister(&_CMSampleBufferGetSampleSizeArray, lib, "CMSampleBufferGetSampleSizeArray")
	tryRegister(&_CMSampleBufferGetSampleTimingInfo, lib, "CMSampleBufferGetSampleTimingInfo")
	tryRegister(&_CMSampleBufferGetSampleTimingInfoArray, lib, "CMSampleBufferGetSampleTimingInfoArray")
	tryRegister(&_CMSampleBufferGetTaggedBufferGroup, lib, "CMSampleBufferGetTaggedBufferGroup")
	tryRegister(&_CMSampleBufferGetTotalSampleSize, lib, "CMSampleBufferGetTotalSampleSize")
	tryRegister(&_CMSampleBufferGetTypeID, lib, "CMSampleBufferGetTypeID")
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
	tryRegister(&_CMSetAttachment, lib, "CMSetAttachment")
	tryRegister(&_CMSetAttachments, lib, "CMSetAttachments")
	tryRegister(&_CMSimpleQueueCreate, lib, "CMSimpleQueueCreate")
	tryRegister(&_CMSimpleQueueDequeue, lib, "CMSimpleQueueDequeue")
	tryRegister(&_CMSimpleQueueEnqueue, lib, "CMSimpleQueueEnqueue")
	tryRegister(&_CMSimpleQueueGetCapacity, lib, "CMSimpleQueueGetCapacity")
	tryRegister(&_CMSimpleQueueGetCount, lib, "CMSimpleQueueGetCount")
	tryRegister(&_CMSimpleQueueGetHead, lib, "CMSimpleQueueGetHead")
	tryRegister(&_CMSimpleQueueGetTypeID, lib, "CMSimpleQueueGetTypeID")
	tryRegister(&_CMSimpleQueueReset, lib, "CMSimpleQueueReset")
	tryRegister(&_CMSwapBigEndianClosedCaptionDescriptionToHost, lib, "CMSwapBigEndianClosedCaptionDescriptionToHost")
	tryRegister(&_CMSwapBigEndianImageDescriptionToHost, lib, "CMSwapBigEndianImageDescriptionToHost")
	tryRegister(&_CMSwapBigEndianMetadataDescriptionToHost, lib, "CMSwapBigEndianMetadataDescriptionToHost")
	tryRegister(&_CMSwapBigEndianSoundDescriptionToHost, lib, "CMSwapBigEndianSoundDescriptionToHost")
	tryRegister(&_CMSwapBigEndianTextDescriptionToHost, lib, "CMSwapBigEndianTextDescriptionToHost")
	tryRegister(&_CMSwapBigEndianTimeCodeDescriptionToHost, lib, "CMSwapBigEndianTimeCodeDescriptionToHost")
	tryRegister(&_CMSwapHostEndianClosedCaptionDescriptionToBig, lib, "CMSwapHostEndianClosedCaptionDescriptionToBig")
	tryRegister(&_CMSwapHostEndianImageDescriptionToBig, lib, "CMSwapHostEndianImageDescriptionToBig")
	tryRegister(&_CMSwapHostEndianMetadataDescriptionToBig, lib, "CMSwapHostEndianMetadataDescriptionToBig")
	tryRegister(&_CMSwapHostEndianSoundDescriptionToBig, lib, "CMSwapHostEndianSoundDescriptionToBig")
	tryRegister(&_CMSwapHostEndianTextDescriptionToBig, lib, "CMSwapHostEndianTextDescriptionToBig")
	tryRegister(&_CMSwapHostEndianTimeCodeDescriptionToBig, lib, "CMSwapHostEndianTimeCodeDescriptionToBig")
	tryRegister(&_CMSyncConvertTime, lib, "CMSyncConvertTime")
	tryRegister(&_CMSyncGetRelativeRate, lib, "CMSyncGetRelativeRate")
	tryRegister(&_CMSyncGetRelativeRateAndAnchorTime, lib, "CMSyncGetRelativeRateAndAnchorTime")
	tryRegister(&_CMSyncGetTime, lib, "CMSyncGetTime")
	tryRegister(&_CMSyncMightDrift, lib, "CMSyncMightDrift")
	tryRegister(&_CMTagCollectionAddTag, lib, "CMTagCollectionAddTag")
	tryRegister(&_CMTagCollectionAddTagsFromArray, lib, "CMTagCollectionAddTagsFromArray")
	tryRegister(&_CMTagCollectionAddTagsFromCollection, lib, "CMTagCollectionAddTagsFromCollection")
	tryRegister(&_CMTagCollectionApply, lib, "CMTagCollectionApply")
	tryRegister(&_CMTagCollectionApplyUntil, lib, "CMTagCollectionApplyUntil")
	tryRegister(&_CMTagCollectionContainsCategory, lib, "CMTagCollectionContainsCategory")
	tryRegister(&_CMTagCollectionContainsSpecifiedTags, lib, "CMTagCollectionContainsSpecifiedTags")
	tryRegister(&_CMTagCollectionContainsTag, lib, "CMTagCollectionContainsTag")
	tryRegister(&_CMTagCollectionContainsTagsOfCollection, lib, "CMTagCollectionContainsTagsOfCollection")
	tryRegister(&_CMTagCollectionCopyAsData, lib, "CMTagCollectionCopyAsData")
	tryRegister(&_CMTagCollectionCopyAsDictionary, lib, "CMTagCollectionCopyAsDictionary")
	tryRegister(&_CMTagCollectionCopyDescription, lib, "CMTagCollectionCopyDescription")
	tryRegister(&_CMTagCollectionCopyTagsOfCategories, lib, "CMTagCollectionCopyTagsOfCategories")
	tryRegister(&_CMTagCollectionCountTagsWithFilterFunction, lib, "CMTagCollectionCountTagsWithFilterFunction")
	tryRegister(&_CMTagCollectionCreate, lib, "CMTagCollectionCreate")
	tryRegister(&_CMTagCollectionCreateCopy, lib, "CMTagCollectionCreateCopy")
	tryRegister(&_CMTagCollectionCreateDifference, lib, "CMTagCollectionCreateDifference")
	tryRegister(&_CMTagCollectionCreateExclusiveOr, lib, "CMTagCollectionCreateExclusiveOr")
	tryRegister(&_CMTagCollectionCreateFromData, lib, "CMTagCollectionCreateFromData")
	tryRegister(&_CMTagCollectionCreateFromDictionary, lib, "CMTagCollectionCreateFromDictionary")
	tryRegister(&_CMTagCollectionCreateIntersection, lib, "CMTagCollectionCreateIntersection")
	tryRegister(&_CMTagCollectionCreateMutable, lib, "CMTagCollectionCreateMutable")
	tryRegister(&_CMTagCollectionCreateMutableCopy, lib, "CMTagCollectionCreateMutableCopy")
	tryRegister(&_CMTagCollectionCreateUnion, lib, "CMTagCollectionCreateUnion")
	tryRegister(&_CMTagCollectionGetCount, lib, "CMTagCollectionGetCount")
	tryRegister(&_CMTagCollectionGetCountOfCategory, lib, "CMTagCollectionGetCountOfCategory")
	tryRegister(&_CMTagCollectionGetTags, lib, "CMTagCollectionGetTags")
	tryRegister(&_CMTagCollectionGetTagsWithCategory, lib, "CMTagCollectionGetTagsWithCategory")
	tryRegister(&_CMTagCollectionGetTagsWithFilterFunction, lib, "CMTagCollectionGetTagsWithFilterFunction")
	tryRegister(&_CMTagCollectionGetTypeID, lib, "CMTagCollectionGetTypeID")
	tryRegister(&_CMTagCollectionIsEmpty, lib, "CMTagCollectionIsEmpty")
	tryRegister(&_CMTagCollectionRemoveAllTags, lib, "CMTagCollectionRemoveAllTags")
	tryRegister(&_CMTagCollectionRemoveAllTagsOfCategory, lib, "CMTagCollectionRemoveAllTagsOfCategory")
	tryRegister(&_CMTagCollectionRemoveTag, lib, "CMTagCollectionRemoveTag")
	tryRegister(&_CMTagCompare, lib, "CMTagCompare")
	tryRegister(&_CMTagCopyAsDictionary, lib, "CMTagCopyAsDictionary")
	tryRegister(&_CMTagCopyDescription, lib, "CMTagCopyDescription")
	tryRegister(&_CMTagEqualToTag, lib, "CMTagEqualToTag")
	tryRegister(&_CMTaggedBufferGroupCreate, lib, "CMTaggedBufferGroupCreate")
	tryRegister(&_CMTaggedBufferGroupCreateCombined, lib, "CMTaggedBufferGroupCreateCombined")
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup, lib, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup")
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup, lib, "CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferAtIndex, lib, "CMTaggedBufferGroupGetCMSampleBufferAtIndex")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferForTag, lib, "CMTaggedBufferGroupGetCMSampleBufferForTag")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferForTagCollection, lib, "CMTaggedBufferGroupGetCMSampleBufferForTagCollection")
	tryRegister(&_CMTaggedBufferGroupGetCount, lib, "CMTaggedBufferGroupGetCount")
	tryRegister(&_CMTaggedBufferGroupGetCVPixelBufferAtIndex, lib, "CMTaggedBufferGroupGetCVPixelBufferAtIndex")
	tryRegister(&_CMTaggedBufferGroupGetCVPixelBufferForTag, lib, "CMTaggedBufferGroupGetCVPixelBufferForTag")
	tryRegister(&_CMTaggedBufferGroupGetCVPixelBufferForTagCollection, lib, "CMTaggedBufferGroupGetCVPixelBufferForTagCollection")
	tryRegister(&_CMTaggedBufferGroupGetNumberOfMatchesForTagCollection, lib, "CMTaggedBufferGroupGetNumberOfMatchesForTagCollection")
	tryRegister(&_CMTaggedBufferGroupGetTagCollectionAtIndex, lib, "CMTaggedBufferGroupGetTagCollectionAtIndex")
	tryRegister(&_CMTaggedBufferGroupGetTypeID, lib, "CMTaggedBufferGroupGetTypeID")
	tryRegister(&_CMTagGetFlagsValue, lib, "CMTagGetFlagsValue")
	tryRegister(&_CMTagGetFloat64Value, lib, "CMTagGetFloat64Value")
	tryRegister(&_CMTagGetOSTypeValue, lib, "CMTagGetOSTypeValue")
	tryRegister(&_CMTagGetSInt64Value, lib, "CMTagGetSInt64Value")
	tryRegister(&_CMTagGetValueDataType, lib, "CMTagGetValueDataType")
	tryRegister(&_CMTagHasFlagsValue, lib, "CMTagHasFlagsValue")
	tryRegister(&_CMTagHasFloat64Value, lib, "CMTagHasFloat64Value")
	tryRegister(&_CMTagHash, lib, "CMTagHash")
	tryRegister(&_CMTagHasOSTypeValue, lib, "CMTagHasOSTypeValue")
	tryRegister(&_CMTagHasSInt64Value, lib, "CMTagHasSInt64Value")
	tryRegister(&_CMTagMakeFromDictionary, lib, "CMTagMakeFromDictionary")
	tryRegister(&_CMTagMakeWithFlagsValue, lib, "CMTagMakeWithFlagsValue")
	tryRegister(&_CMTagMakeWithFloat64Value, lib, "CMTagMakeWithFloat64Value")
	tryRegister(&_CMTagMakeWithOSTypeValue, lib, "CMTagMakeWithOSTypeValue")
	tryRegister(&_CMTagMakeWithSInt64Value, lib, "CMTagMakeWithSInt64Value")
	tryRegister(&_CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer, lib, "CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer")
	tryRegister(&_CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer, lib, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer")
	tryRegister(&_CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData, lib, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData")
	tryRegister(&_CMTextFormatDescriptionGetDefaultStyle, lib, "CMTextFormatDescriptionGetDefaultStyle")
	tryRegister(&_CMTextFormatDescriptionGetDefaultTextBox, lib, "CMTextFormatDescriptionGetDefaultTextBox")
	tryRegister(&_CMTextFormatDescriptionGetDisplayFlags, lib, "CMTextFormatDescriptionGetDisplayFlags")
	tryRegister(&_CMTextFormatDescriptionGetFontName, lib, "CMTextFormatDescriptionGetFontName")
	tryRegister(&_CMTextFormatDescriptionGetJustification, lib, "CMTextFormatDescriptionGetJustification")
	tryRegister(&_CMTimeAbsoluteValue, lib, "CMTimeAbsoluteValue")
	tryRegister(&_CMTimeAdd, lib, "CMTimeAdd")
	tryRegister(&_CMTimebaseAddTimer, lib, "CMTimebaseAddTimer")
	tryRegister(&_CMTimebaseAddTimerDispatchSource, lib, "CMTimebaseAddTimerDispatchSource")
	tryRegister(&_CMTimebaseCopyMasterClock, lib, "CMTimebaseCopyMasterClock")
	tryRegister(&_CMTimebaseCopyMasterTimebase, lib, "CMTimebaseCopyMasterTimebase")
	tryRegister(&_CMTimebaseCopySource, lib, "CMTimebaseCopySource")
	tryRegister(&_CMTimebaseCopySourceClock, lib, "CMTimebaseCopySourceClock")
	tryRegister(&_CMTimebaseCopyUltimateMasterClock, lib, "CMTimebaseCopyUltimateMasterClock")
	tryRegister(&_CMTimebaseCopyUltimateSourceClock, lib, "CMTimebaseCopyUltimateSourceClock")
	tryRegister(&_CMTimebaseCreateWithMasterClock, lib, "CMTimebaseCreateWithMasterClock")
	tryRegister(&_CMTimebaseCreateWithMasterTimebase, lib, "CMTimebaseCreateWithMasterTimebase")
	tryRegister(&_CMTimebaseCreateWithSourceClock, lib, "CMTimebaseCreateWithSourceClock")
	tryRegister(&_CMTimebaseGetEffectiveRate, lib, "CMTimebaseGetEffectiveRate")
	tryRegister(&_CMTimebaseGetMaster, lib, "CMTimebaseGetMaster")
	tryRegister(&_CMTimebaseGetMasterClock, lib, "CMTimebaseGetMasterClock")
	tryRegister(&_CMTimebaseGetMasterTimebase, lib, "CMTimebaseGetMasterTimebase")
	tryRegister(&_CMTimebaseGetRate, lib, "CMTimebaseGetRate")
	tryRegister(&_CMTimebaseGetTime, lib, "CMTimebaseGetTime")
	tryRegister(&_CMTimebaseGetTimeAndRate, lib, "CMTimebaseGetTimeAndRate")
	tryRegister(&_CMTimebaseGetTimeWithTimeScale, lib, "CMTimebaseGetTimeWithTimeScale")
	tryRegister(&_CMTimebaseGetTypeID, lib, "CMTimebaseGetTypeID")
	tryRegister(&_CMTimebaseGetUltimateMasterClock, lib, "CMTimebaseGetUltimateMasterClock")
	tryRegister(&_CMTimebaseNotificationBarrier, lib, "CMTimebaseNotificationBarrier")
	tryRegister(&_CMTimebaseRemoveTimer, lib, "CMTimebaseRemoveTimer")
	tryRegister(&_CMTimebaseRemoveTimerDispatchSource, lib, "CMTimebaseRemoveTimerDispatchSource")
	tryRegister(&_CMTimebaseSetAnchorTime, lib, "CMTimebaseSetAnchorTime")
	tryRegister(&_CMTimebaseSetRate, lib, "CMTimebaseSetRate")
	tryRegister(&_CMTimebaseSetRateAndAnchorTime, lib, "CMTimebaseSetRateAndAnchorTime")
	tryRegister(&_CMTimebaseSetSourceClock, lib, "CMTimebaseSetSourceClock")
	tryRegister(&_CMTimebaseSetSourceTimebase, lib, "CMTimebaseSetSourceTimebase")
	tryRegister(&_CMTimebaseSetTime, lib, "CMTimebaseSetTime")
	tryRegister(&_CMTimebaseSetTimerDispatchSourceNextFireTime, lib, "CMTimebaseSetTimerDispatchSourceNextFireTime")
	tryRegister(&_CMTimebaseSetTimerDispatchSourceToFireImmediately, lib, "CMTimebaseSetTimerDispatchSourceToFireImmediately")
	tryRegister(&_CMTimebaseSetTimerNextFireTime, lib, "CMTimebaseSetTimerNextFireTime")
	tryRegister(&_CMTimebaseSetTimerToFireImmediately, lib, "CMTimebaseSetTimerToFireImmediately")
	tryRegister(&_CMTimeClampToRange, lib, "CMTimeClampToRange")
	tryRegister(&_CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer, lib, "CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer")
	tryRegister(&_CMTimeCodeFormatDescriptionCreate, lib, "CMTimeCodeFormatDescriptionCreate")
	tryRegister(&_CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer, lib, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer")
	tryRegister(&_CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData, lib, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData")
	tryRegister(&_CMTimeCodeFormatDescriptionGetFrameDuration, lib, "CMTimeCodeFormatDescriptionGetFrameDuration")
	tryRegister(&_CMTimeCodeFormatDescriptionGetFrameQuanta, lib, "CMTimeCodeFormatDescriptionGetFrameQuanta")
	tryRegister(&_CMTimeCodeFormatDescriptionGetTimeCodeFlags, lib, "CMTimeCodeFormatDescriptionGetTimeCodeFlags")
	tryRegister(&_CMTimeCompare, lib, "CMTimeCompare")
	tryRegister(&_CMTimeConvertScale, lib, "CMTimeConvertScale")
	tryRegister(&_CMTimeCopyAsDictionary, lib, "CMTimeCopyAsDictionary")
	tryRegister(&_CMTimeCopyDescription, lib, "CMTimeCopyDescription")
	tryRegister(&_CMTimeFoldIntoRange, lib, "CMTimeFoldIntoRange")
	tryRegister(&_CMTimeGetSeconds, lib, "CMTimeGetSeconds")
	tryRegister(&_CMTimeMake, lib, "CMTimeMake")
	tryRegister(&_CMTimeMakeFromDictionary, lib, "CMTimeMakeFromDictionary")
	tryRegister(&_CMTimeMakeWithEpoch, lib, "CMTimeMakeWithEpoch")
	tryRegister(&_CMTimeMakeWithSeconds, lib, "CMTimeMakeWithSeconds")
	tryRegister(&_CMTimeMapDurationFromRangeToRange, lib, "CMTimeMapDurationFromRangeToRange")
	tryRegister(&_CMTimeMappingCopyAsDictionary, lib, "CMTimeMappingCopyAsDictionary")
	tryRegister(&_CMTimeMappingCopyDescription, lib, "CMTimeMappingCopyDescription")
	tryRegister(&_CMTimeMappingMake, lib, "CMTimeMappingMake")
	tryRegister(&_CMTimeMappingMakeEmpty, lib, "CMTimeMappingMakeEmpty")
	tryRegister(&_CMTimeMappingMakeFromDictionary, lib, "CMTimeMappingMakeFromDictionary")
	tryRegister(&_CMTimeMappingShow, lib, "CMTimeMappingShow")
	tryRegister(&_CMTimeMapTimeFromRangeToRange, lib, "CMTimeMapTimeFromRangeToRange")
	tryRegister(&_CMTimeMaximum, lib, "CMTimeMaximum")
	tryRegister(&_CMTimeMinimum, lib, "CMTimeMinimum")
	tryRegister(&_CMTimeMultiply, lib, "CMTimeMultiply")
	tryRegister(&_CMTimeMultiplyByFloat64, lib, "CMTimeMultiplyByFloat64")
	tryRegister(&_CMTimeMultiplyByRatio, lib, "CMTimeMultiplyByRatio")
	tryRegister(&_CMTimeRangeContainsTime, lib, "CMTimeRangeContainsTime")
	tryRegister(&_CMTimeRangeContainsTimeRange, lib, "CMTimeRangeContainsTimeRange")
	tryRegister(&_CMTimeRangeCopyAsDictionary, lib, "CMTimeRangeCopyAsDictionary")
	tryRegister(&_CMTimeRangeCopyDescription, lib, "CMTimeRangeCopyDescription")
	tryRegister(&_CMTimeRangeEqual, lib, "CMTimeRangeEqual")
	tryRegister(&_CMTimeRangeFromTimeToTime, lib, "CMTimeRangeFromTimeToTime")
	tryRegister(&_CMTimeRangeGetEnd, lib, "CMTimeRangeGetEnd")
	tryRegister(&_CMTimeRangeGetIntersection, lib, "CMTimeRangeGetIntersection")
	tryRegister(&_CMTimeRangeGetUnion, lib, "CMTimeRangeGetUnion")
	tryRegister(&_CMTimeRangeMake, lib, "CMTimeRangeMake")
	tryRegister(&_CMTimeRangeMakeFromDictionary, lib, "CMTimeRangeMakeFromDictionary")
	tryRegister(&_CMTimeRangeShow, lib, "CMTimeRangeShow")
	tryRegister(&_CMTimeShow, lib, "CMTimeShow")
	tryRegister(&_CMTimeSubtract, lib, "CMTimeSubtract")
	tryRegister(&_CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer, lib, "CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer")
	tryRegister(&_CMVideoFormatDescriptionCopyTagCollectionArray, lib, "CMVideoFormatDescriptionCopyTagCollectionArray")
	tryRegister(&_CMVideoFormatDescriptionCreate, lib, "CMVideoFormatDescriptionCreate")
	tryRegister(&_CMVideoFormatDescriptionCreateForImageBuffer, lib, "CMVideoFormatDescriptionCreateForImageBuffer")
	tryRegister(&_CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer, lib, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer")
	tryRegister(&_CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData, lib, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData")
	tryRegister(&_CMVideoFormatDescriptionCreateFromH264ParameterSets, lib, "CMVideoFormatDescriptionCreateFromH264ParameterSets")
	tryRegister(&_CMVideoFormatDescriptionCreateFromHEVCParameterSets, lib, "CMVideoFormatDescriptionCreateFromHEVCParameterSets")
	tryRegister(&_CMVideoFormatDescriptionGetCleanAperture, lib, "CMVideoFormatDescriptionGetCleanAperture")
	tryRegister(&_CMVideoFormatDescriptionGetDimensions, lib, "CMVideoFormatDescriptionGetDimensions")
	tryRegister(&_CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers, lib, "CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers")
	tryRegister(&_CMVideoFormatDescriptionGetH264ParameterSetAtIndex, lib, "CMVideoFormatDescriptionGetH264ParameterSetAtIndex")
	tryRegister(&_CMVideoFormatDescriptionGetHEVCParameterSetAtIndex, lib, "CMVideoFormatDescriptionGetHEVCParameterSetAtIndex")
	tryRegister(&_CMVideoFormatDescriptionGetPresentationDimensions, lib, "CMVideoFormatDescriptionGetPresentationDimensions")
	tryRegister(&_CMVideoFormatDescriptionMatchesImageBuffer, lib, "CMVideoFormatDescriptionMatchesImageBuffer")
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



// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions is a CoreMedia function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator AllocatorRef, taggedBufferGroup TaggedBufferGroupRef, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
}

// Creates a clock that advances at the same rate as audio output.

// Creates a clock that advances at the same rate as audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioClockCreate(allocator:clockOut:)
func CMAudioClockCreate(allocator AllocatorRef, clockOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioClockCreate(allocator, clockOut)
}

// Creates a clock that tracks playback through a Core Audio device with the specified unique identifier.
//
// Added in macOS 10.8.
// Creates a clock that tracks playback through a Core Audio device with the specified unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreate(allocator:deviceUID:clockOut:)
func CMAudioDeviceClockCreate(allocator AllocatorRef, deviceUID StringRef, clockOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockCreate(allocator, deviceUID, clockOut)
}

// Creates a clock that tracks playback through a Core Audio device with the specified identifier.
//
// Added in macOS 10.8.
// Creates a clock that tracks playback through a Core Audio device with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreateFromAudioDeviceID(allocator:deviceID:clockOut:)
func CMAudioDeviceClockCreateFromAudioDeviceID(allocator AllocatorRef, deviceID unsafe.Pointer, clockOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockCreateFromAudioDeviceID(allocator, deviceID, clockOut)
}

// Returns the Core Audio device the clock is tracking.
//
// Added in macOS 10.8.
// Returns the Core Audio device the clock is tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockGetAudioDevice(_:deviceUIDOut:deviceIDOut:trackingDefaultDeviceOut:)
func CMAudioDeviceClockGetAudioDevice(clock ClockRef, deviceUIDOut unsafe.Pointer, deviceIDOut unsafe.Pointer, trackingDefaultDeviceOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockGetAudioDevice(clock, deviceUIDOut, deviceIDOut, trackingDefaultDeviceOut)
}

// Changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// Added in macOS 10.8.
// Changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceID(_:deviceID:)
func CMAudioDeviceClockSetAudioDeviceID(clock ClockRef, deviceID unsafe.Pointer) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceID(clock, deviceID)
}

// Changes the Core Audio device the clock is tracking by specifying a new device unique identifier.
//
// Added in macOS 10.8.
// Changes the Core Audio device the clock is tracking by specifying a new device unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceUID(_:deviceUID:)
func CMAudioDeviceClockSetAudioDeviceUID(clock ClockRef, deviceUID StringRef) unsafe.Pointer {
	return _CMAudioDeviceClockSetAudioDeviceUID(clock, deviceUID)
}

// Copies the contents of an audio format description to a buffer in big-endian byte ordering.
//
// Added in macOS 10.10.
// Copies the contents of an audio format description to a buffer in big-endian byte ordering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator:audioFormatDescription:flavor:blockBufferOut:)
func CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator AllocatorRef, audioFormatDescription AudioFormatDescriptionRef, flavor SoundDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator, audioFormatDescription, flavor, blockBufferOut)
}

// Creates a format description for an audio media stream.
//
// Added in macOS 10.7.
// Creates a format description for an audio media stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreate(allocator:asbd:layoutSize:layout:magicCookieSize:magicCookie:extensions:formatDescriptionOut:)
func CMAudioFormatDescriptionCreate(allocator AllocatorRef, asbd unsafe.Pointer, layoutSize uintptr, layout unsafe.Pointer, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionCreate(allocator, asbd, layoutSize, layout, magicCookieSize, magicCookie, extensions, formatDescriptionOut)
}

// Creates an audio format description from a big-endian sound description data structure in a buffer.
//
// Added in macOS 10.10.
// Creates an audio format description from a big-endian sound description data structure in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator:bigEndianSoundDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator AllocatorRef, soundDescriptionBlockBuffer BlockBufferRef, flavor SoundDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator, soundDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

// Creates an audio format description from a big-endian sound description data structure.
//
// Added in macOS 10.10.
// Creates an audio format description from a big-endian sound description data structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator:bigEndianSoundDescriptionData:size:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator AllocatorRef, soundDescriptionData unsafe.Pointer, size uintptr, flavor SoundDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator, soundDescriptionData, size, flavor, formatDescriptionOut)
}

// Creates a summary audio format description from an array of descriptions.
//
// Added in macOS 10.7.
// Creates a summary audio format description from an array of descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateSummary(allocator:formatDescriptionArray:flags:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateSummary(allocator AllocatorRef, formatDescriptionArray ArrayRef, flags uint32, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionCreateSummary(allocator, formatDescriptionArray, flags, formatDescriptionOut)
}

// Returns a Boolean value that indicates whether the two audio format descriptions are equal.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether the two audio format descriptions are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionEqual(_:otherFormatDescription:equalityMask:equalityMaskOut:)
func CMAudioFormatDescriptionEqual(formatDescription AudioFormatDescriptionRef, otherFormatDescription AudioFormatDescriptionRef, equalityMask AudioFormatDescriptionMask, equalityMaskOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionEqual(formatDescription, otherFormatDescription, equalityMask, equalityMaskOut)
}

// Returns a read-only pointer to, and the size of, the audio channel layout inside an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to, and the size of, the audio channel layout inside an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetChannelLayout(_:sizeOut:)
func CMAudioFormatDescriptionGetChannelLayout(desc AudioFormatDescriptionRef, sizeOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetChannelLayout(desc, sizeOut)
}

// Returns a read-only pointer to, and size of, the array of audio format list item structures in an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to, and size of, the array of audio format list item structures in an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetFormatList(_:sizeOut:)
func CMAudioFormatDescriptionGetFormatList(desc AudioFormatDescriptionRef, sizeOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetFormatList(desc, sizeOut)
}

// Returns a read-only pointer to, and size of, the magic cookie in an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to, and size of, the magic cookie in an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMagicCookie(_:sizeOut:)
func CMAudioFormatDescriptionGetMagicCookie(desc AudioFormatDescriptionRef, sizeOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetMagicCookie(desc, sizeOut)
}

// Returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMostCompatibleFormat(_:)
func CMAudioFormatDescriptionGetMostCompatibleFormat(desc AudioFormatDescriptionRef) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetMostCompatibleFormat(desc)
}

// Returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetRichestDecodableFormat(_:)
func CMAudioFormatDescriptionGetRichestDecodableFormat(desc AudioFormatDescriptionRef) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetRichestDecodableFormat(desc)
}

// Returns a read-only pointer to the audio stream description in an audio format description.
//
// Added in macOS 10.7.
// Returns a read-only pointer to the audio stream description in an audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetStreamBasicDescription(_:)
func CMAudioFormatDescriptionGetStreamBasicDescription(desc AudioFormatDescriptionRef) unsafe.Pointer {
	return _CMAudioFormatDescriptionGetStreamBasicDescription(desc)
}

// Creates a sample buffer with packet descriptions.
//
// Added in macOS 10.10.
// Creates a sample buffer with packet descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator:dataBuffer:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator AllocatorRef, dataBuffer BlockBufferRef, formatDescription FormatDescriptionRef, numSamples ItemCount, presentationTimeStamp Time, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator, dataBuffer, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
}

// Creates a sample buffer with packet descriptions and a callback to make the data ready for use.
//
// Added in macOS 10.7.
// Creates a sample buffer with packet descriptions and a callback to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptions(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateWithPacketDescriptions(allocator AllocatorRef, dataBuffer BlockBufferRef, dataReady unsafe.Pointer, makeDataReadyCallback SampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription FormatDescriptionRef, numSamples ItemCount, presentationTimeStamp Time, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMAudioSampleBufferCreateWithPacketDescriptions(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
}

// Creates a sample buffer with packet descriptions and a handler to make the data ready for use.
//
// Added in macOS 10.14.4.
// Creates a sample buffer with packet descriptions and a handler to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:)
func CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator AllocatorRef, dataBuffer BlockBufferRef, dataReady unsafe.Pointer, formatDescription FormatDescriptionRef, numSamples ItemCount, presentationTimeStamp Time, packetDescriptions unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler SampleBufferMakeDataReadyHandler) unsafe.Pointer {
	return _CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut, makeDataReadyHandler)
}

// Accesses potentially noncontiguous data in a block buffer.
//
// Added in macOS 10.7.
// Accesses potentially noncontiguous data in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAccessDataBytes(_:atOffset:length:temporaryBlock:returnedPointerOut:)
func CMBlockBufferAccessDataBytes(theBuffer BlockBufferRef, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferAccessDataBytes(theBuffer, offset, length, temporaryBlock, returnedPointerOut)
}

// Adds a reference to an existing block buffer.
//
// Added in macOS 10.7.
// Adds a reference to an existing block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendBufferReference(_:targetBBuf:offsetToData:dataLength:flags:)
func CMBlockBufferAppendBufferReference(theBuffer BlockBufferRef, targetBBuf BlockBufferRef, offsetToData uintptr, dataLength uintptr, flags BlockBufferFlags) unsafe.Pointer {
	return _CMBlockBufferAppendBufferReference(theBuffer, targetBBuf, offsetToData, dataLength, flags)
}

// Adds a memory block to an existing block buffer.
//
// Added in macOS 10.7.
// Adds a memory block to an existing block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendMemoryBlock(_:memoryBlock:length:blockAllocator:customBlockSource:offsetToData:dataLength:flags:)
func CMBlockBufferAppendMemoryBlock(theBuffer BlockBufferRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator AllocatorRef, customBlockSource unsafe.Pointer, offsetToData uintptr, dataLength uintptr, flags BlockBufferFlags) unsafe.Pointer {
	return _CMBlockBufferAppendMemoryBlock(theBuffer, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags)
}

// Assures that the system allocates memory for all memory blocks in a block buffer.
//
// Added in macOS 10.7.
// Assures that the system allocates memory for all memory blocks in a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAssureBlockMemory(_:)
func CMBlockBufferAssureBlockMemory(theBuffer BlockBufferRef) unsafe.Pointer {
	return _CMBlockBufferAssureBlockMemory(theBuffer)
}

// Copies bytes from a block buffer into a provided memory area.
//
// Added in macOS 10.7.
// Copies bytes from a block buffer into a provided memory area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCopyDataBytes(_:atOffset:dataLength:destination:)
func CMBlockBufferCopyDataBytes(theSourceBuffer BlockBufferRef, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCopyDataBytes(theSourceBuffer, offsetToData, dataLength, destination)
}

// Creates a block buffer that contains a contiguous copy of, or reference to, the data specified by the parameters.
//
// Added in macOS 10.7.
// Creates a block buffer that contains a contiguous copy of, or reference to, the data specified by the parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateContiguous(allocator:sourceBuffer:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateContiguous(structureAllocator AllocatorRef, sourceBuffer BlockBufferRef, blockAllocator AllocatorRef, customBlockSource unsafe.Pointer, offsetToData uintptr, dataLength uintptr, flags BlockBufferFlags, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateContiguous(structureAllocator, sourceBuffer, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
}

// Creates an empty block buffer.
//
// Added in macOS 10.7.
// Creates an empty block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateEmpty(allocator:capacity:flags:blockBufferOut:)
func CMBlockBufferCreateEmpty(structureAllocator AllocatorRef, subBlockCapacity uint32, flags BlockBufferFlags, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateEmpty(structureAllocator, subBlockCapacity, flags, blockBufferOut)
}

// Creates a block buffer that refers to another block buffer object.
//
// Added in macOS 10.7.
// Creates a block buffer that refers to another block buffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithBufferReference(allocator:referenceBuffer:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithBufferReference(structureAllocator AllocatorRef, bufferReference BlockBufferRef, offsetToData uintptr, dataLength uintptr, flags BlockBufferFlags, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateWithBufferReference(structureAllocator, bufferReference, offsetToData, dataLength, flags, blockBufferOut)
}

// Creates a block buffer that’s backed by a memory block.
//
// Added in macOS 10.7.
// Creates a block buffer that’s backed by a memory block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithMemoryBlock(allocator:memoryBlock:blockLength:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithMemoryBlock(structureAllocator AllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator AllocatorRef, customBlockSource unsafe.Pointer, offsetToData uintptr, dataLength uintptr, flags BlockBufferFlags, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferCreateWithMemoryBlock(structureAllocator, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
}

// Fills the destination buffer with the specified data byte.
//
// Added in macOS 10.7.
// Fills the destination buffer with the specified data byte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferFillDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferFillDataBytes(fillByte unsafe.Pointer, destinationBuffer BlockBufferRef, offsetIntoDestination uintptr, dataLength uintptr) unsafe.Pointer {
	return _CMBlockBufferFillDataBytes(fillByte, destinationBuffer, offsetIntoDestination, dataLength)
}

// Returns the total length of data that’s accessible by a block buffer.
//
// Added in macOS 10.7.
// Returns the total length of data that’s accessible by a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataLength(_:)
func CMBlockBufferGetDataLength(theBuffer BlockBufferRef) uintptr {
	return _CMBlockBufferGetDataLength(theBuffer)
}

// Gains access to the data represented by a block buffer.
//
// Added in macOS 10.7.
// Gains access to the data represented by a block buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataPointer(_:atOffset:lengthAtOffsetOut:totalLengthOut:dataPointerOut:)
func CMBlockBufferGetDataPointer(theBuffer BlockBufferRef, offset uintptr, lengthAtOffsetOut unsafe.Pointer, totalLengthOut unsafe.Pointer, dataPointerOut unsafe.Pointer) unsafe.Pointer {
	return _CMBlockBufferGetDataPointer(theBuffer, offset, lengthAtOffsetOut, totalLengthOut, dataPointerOut)
}

// Returns the type identifier for block buffer objects.
//
// Added in macOS 10.7.
// Returns the type identifier for block buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetTypeID()
func CMBlockBufferGetTypeID() TypeID {
	return _CMBlockBufferGetTypeID()
}

// Returns a Boolean value that indicates whether the buffer is empty.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether the buffer is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsEmpty(_:)
func CMBlockBufferIsEmpty(theBuffer BlockBufferRef) unsafe.Pointer {
	return _CMBlockBufferIsEmpty(theBuffer)
}

// Returns a Boolean value that indicates whether the specified range within a block buffer is contiguous.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether the specified range within a block buffer is contiguous.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsRangeContiguous(_:atOffset:length:)
func CMBlockBufferIsRangeContiguous(theBuffer BlockBufferRef, offset uintptr, length uintptr) unsafe.Pointer {
	return _CMBlockBufferIsRangeContiguous(theBuffer, offset, length)
}

// Copies bytes from a given memory block into a block buffer replacing bytes in the underlying data blocks.
//
// Added in macOS 10.7.
// Copies bytes from a given memory block into a block buffer replacing bytes in the underlying data blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferReplaceDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferReplaceDataBytes(sourceBytes unsafe.Pointer, destinationBuffer BlockBufferRef, offsetIntoDestination uintptr, dataLength uintptr) unsafe.Pointer {
	return _CMBlockBufferReplaceDataBytes(sourceBytes, destinationBuffer, offsetIntoDestination, dataLength)
}

// Calls a function for every buffer in a queue.
//
// Added in macOS 10.7.
// Calls a function for every buffer in a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCallForEachBuffer(_:callback:refcon:)
func CMBufferQueueCallForEachBuffer(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueCallForEachBuffer(queue)
}

// Returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueContainsEndOfData(_:)
func CMBufferQueueContainsEndOfData(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueContainsEndOfData(queue)
}

// CMBufferQueueCopyHead is a CoreMedia function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue BufferQueueRef) BufferRef {
	return _CMBufferQueueCopyHead(queue)
}

// Creates a buffer queue with callbacks to inspect buffers.
//
// Added in macOS 10.7.
// Creates a buffer queue with callbacks to inspect buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreate(allocator:capacity:callbacks:queueOut:)
func CMBufferQueueCreate(allocator AllocatorRef, capacity ItemCount, callbacks unsafe.Pointer, queueOut unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueCreate(allocator, capacity, callbacks, queueOut)
}

// Creates a buffer queue with handlers to inspect buffers.
//
// Added in macOS 10.14.4.
// Creates a buffer queue with handlers to inspect buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreateWithHandlers(_:_:_:_:)
func CMBufferQueueCreateWithHandlers(allocator AllocatorRef, capacity ItemCount, handlers unsafe.Pointer, queueOut unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueCreateWithHandlers(allocator, capacity, handlers, queueOut)
}

// Dequeues a buffer from a queue.
//
// Added in macOS 10.7.
// Dequeues a buffer from a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeueAndRetain(queue BufferQueueRef) BufferRef {
	return _CMBufferQueueDequeueAndRetain(queue)
}

// Dequeues a buffer from a queue, if it’s ready.
//
// Added in macOS 10.7.
// Dequeues a buffer from a queue, if it’s ready.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReadyAndRetain(queue BufferQueueRef) BufferRef {
	return _CMBufferQueueDequeueIfDataReadyAndRetain(queue)
}

// Enqueues a buffer onto a queue.
//
// Added in macOS 10.7.
// Enqueues a buffer onto a queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)
func CMBufferQueueEnqueue(queue BufferQueueRef, buf BufferRef) unsafe.Pointer {
	return _CMBufferQueueEnqueue(queue, buf)
}

// Gets the number of buffers in the queue.
//
// Added in macOS 10.7.
// Gets the number of buffers in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetBufferCount(_:)
func CMBufferQueueGetBufferCount(queue BufferQueueRef) ItemCount {
	return _CMBufferQueueGetBufferCount(queue)
}

// Returns a pointer to a structure that contains callbacks to sort sample buffers by output presentation timestamp.
//
// Added in macOS 10.7.
// Returns a pointer to a structure that contains callbacks to sort sample buffers by output presentation timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
func CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS() unsafe.Pointer {
	return _CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
}

// Returns a pointer to a callback structure for unsorted sample buffers.
//
// Added in macOS 10.7.
// Returns a pointer to a callback structure for unsorted sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForUnsortedSampleBuffers()
func CMBufferQueueGetCallbacksForUnsortedSampleBuffers() unsafe.Pointer {
	return _CMBufferQueueGetCallbacksForUnsortedSampleBuffers()
}

// Gets the duration of a buffer queue.
//
// Added in macOS 10.7.
// Gets the duration of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetDuration(_:)
func CMBufferQueueGetDuration(queue BufferQueueRef) Time {
	return _CMBufferQueueGetDuration(queue)
}

// Gets the greatest end presentation timestamp of a buffer queue.
//
// Added in macOS 10.7.
// Gets the greatest end presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetEndPresentationTimeStamp(_:)
func CMBufferQueueGetEndPresentationTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetEndPresentationTimeStamp(queue)
}

// Gets the decode timestamp of the first buffer in a buffer queue.
//
// Added in macOS 10.7.
// Gets the decode timestamp of the first buffer in a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstDecodeTimeStamp(_:)
func CMBufferQueueGetFirstDecodeTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetFirstDecodeTimeStamp(queue)
}

// Gets the presentation timestamp of the first buffer in a buffer queue.
//
// Added in macOS 10.7.
// Gets the presentation timestamp of the first buffer in a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstPresentationTimeStamp(_:)
func CMBufferQueueGetFirstPresentationTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetFirstPresentationTimeStamp(queue)
}

// Retrieves the next buffer from a queue, but doesn’t remove it.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.7.
// Retrieves the next buffer from a queue, but doesn’t remove it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetHead(_:)
func CMBufferQueueGetHead(queue BufferQueueRef) BufferRef {
	return _CMBufferQueueGetHead(queue)
}

// Gets the greatest presentation timestamp of a buffer queue.
//
// Added in macOS 10.7.
// Gets the greatest presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMaxPresentationTimeStamp(_:)
func CMBufferQueueGetMaxPresentationTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetMaxPresentationTimeStamp(queue)
}

// Gets the earliest decode timestamp of a buffer queue.
//
// Added in macOS 10.7.
// Gets the earliest decode timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinDecodeTimeStamp(_:)
func CMBufferQueueGetMinDecodeTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetMinDecodeTimeStamp(queue)
}

// Gets the earliest presentation timestamp of a buffer queue.
//
// Added in macOS 10.7.
// Gets the earliest presentation timestamp of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinPresentationTimeStamp(_:)
func CMBufferQueueGetMinPresentationTimeStamp(queue BufferQueueRef) Time {
	return _CMBufferQueueGetMinPresentationTimeStamp(queue)
}

// Gets the total size of all sample buffers of a buffer queue.
//
// Added in macOS 10.10.
// Gets the total size of all sample buffers of a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTotalSize(_:)
func CMBufferQueueGetTotalSize(queue BufferQueueRef) uintptr {
	return _CMBufferQueueGetTotalSize(queue)
}

// Returns the type identifier of buffer queue objects.
//
// Added in macOS 10.7.
// Returns the type identifier of buffer queue objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTypeID()
func CMBufferQueueGetTypeID() TypeID {
	return _CMBufferQueueGetTypeID()
}

// Installs a trigger with a callback on a buffer queue.
//
// Added in macOS 10.7.
// Installs a trigger with a callback on a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTrigger(_:callback:refcon:condition:time:triggerTokenOut:)
func CMBufferQueueInstallTrigger(queue BufferQueueRef, callback BufferQueueTriggerCallback, refcon unsafe.Pointer, condition BufferQueueTriggerCondition, time Time, triggerTokenOut unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueInstallTrigger(queue, callback, refcon, condition, time, triggerTokenOut)
}

// Installs a trigger with a handler on a buffer queue.
//
// Added in macOS 10.14.4.
// Installs a trigger with a handler on a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandler(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandler(queue BufferQueueRef, condition BufferQueueTriggerCondition, time Time, triggerTokenOut unsafe.Pointer, handler BufferQueueTriggerHandler) unsafe.Pointer {
	return _CMBufferQueueInstallTriggerHandler(queue, condition, time, triggerTokenOut, handler)
}

// Installs a trigger with a handler and threshold on a buffer queue.
//
// Added in macOS 10.14.4.
// Installs a trigger with a handler and threshold on a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue BufferQueueRef, condition BufferQueueTriggerCondition, threshold ItemCount, triggerTokenOut unsafe.Pointer, handler BufferQueueTriggerHandler) unsafe.Pointer {
	return _CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue, condition, threshold, triggerTokenOut, handler)
}

// Installs a trigger with a callback and threshold on a buffer queue.
//
// Added in macOS 10.7.
// Installs a trigger with a callback and threshold on a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerWithIntegerThreshold(_:callback:refcon:condition:threshold:triggerTokenOut:)
func CMBufferQueueInstallTriggerWithIntegerThreshold(queue BufferQueueRef, callback BufferQueueTriggerCallback, refcon unsafe.Pointer, condition BufferQueueTriggerCondition, threshold ItemCount, triggerTokenOut unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueInstallTriggerWithIntegerThreshold(queue, callback, refcon, condition, threshold, triggerTokenOut)
}

// Returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set, and is now empty.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set, and is now empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsAtEndOfData(_:)
func CMBufferQueueIsAtEndOfData(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueIsAtEndOfData(queue)
}

// Returns a Boolean value that indicates whether a buffer queue is empty.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a buffer queue is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsEmpty(_:)
func CMBufferQueueIsEmpty(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueIsEmpty(queue)
}

// Sets a marker to indicate this queue doesn’t allow enqueuing new buffers.
//
// Added in macOS 10.7.
// Sets a marker to indicate this queue doesn’t allow enqueuing new buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueMarkEndOfData(_:)
func CMBufferQueueMarkEndOfData(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueMarkEndOfData(queue)
}

// Removes a previously installed trigger from a buffer queue.
//
// Added in macOS 10.7.
// Removes a previously installed trigger from a buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueRemoveTrigger(_:triggerToken:)
func CMBufferQueueRemoveTrigger(queue BufferQueueRef, triggerToken BufferQueueTriggerToken) unsafe.Pointer {
	return _CMBufferQueueRemoveTrigger(queue, triggerToken)
}

// Resets a buffer queue, which allows it to enqueue new buffers.
//
// Added in macOS 10.7.
// Resets a buffer queue, which allows it to enqueue new buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueReset(_:)
func CMBufferQueueReset(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueReset(queue)
}

// A callback that invokes a function for every buffer in a queue and then resets the queue.
//
// Added in macOS 10.7.
// A callback that invokes a function for every buffer in a queue and then resets the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueResetWithCallback(_:callback:refcon:)
func CMBufferQueueResetWithCallback(queue BufferQueueRef) unsafe.Pointer {
	return _CMBufferQueueResetWithCallback(queue)
}

// A validation callback for the queue to call before enqueuing buffers.
//
// Added in macOS 10.7.
// A validation callback for the queue to call before enqueuing buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationCallback(_:callback:refcon:)
func CMBufferQueueSetValidationCallback(queue BufferQueueRef, callback BufferValidationCallback, refcon unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueSetValidationCallback(queue, callback, refcon)
}

// A validation handler for the queue to call before enqueuing buffers.
//
// Added in macOS 10.14.4.
// A validation handler for the queue to call before enqueuing buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationHandler(_:_:)
func CMBufferQueueSetValidationHandler(queue BufferQueueRef, handler BufferValidationHandler) unsafe.Pointer {
	return _CMBufferQueueSetValidationHandler(queue, handler)
}

// Tests whether the trigger condition is true for the specified buffer queue.
//
// Added in macOS 10.7.
// Tests whether the trigger condition is true for the specified buffer queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTestTrigger(_:triggerToken:)
func CMBufferQueueTestTrigger(queue BufferQueueRef, triggerToken BufferQueueTriggerToken) unsafe.Pointer {
	return _CMBufferQueueTestTrigger(queue, triggerToken)
}

// Converts a host time from a core media time structure to the host time’s native units.
//
// Added in macOS 10.8.
// Converts a host time from a core media time structure to the host time’s native units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockConvertHostTimeToSystemUnits(_:)
func CMClockConvertHostTimeToSystemUnits(hostTime Time) uint64 {
	return _CMClockConvertHostTimeToSystemUnits(hostTime)
}

// Returns the current time from a clock and the matching time from the clock’s reference clock.
//
// Added in macOS 10.8.
// Returns the current time from a clock and the matching time from the clock’s reference clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetAnchorTime(_:clockTimeOut:referenceClockTimeOut:)
func CMClockGetAnchorTime(clock ClockRef, clockTimeOut unsafe.Pointer, referenceClockTimeOut unsafe.Pointer) unsafe.Pointer {
	return _CMClockGetAnchorTime(clock, clockTimeOut, referenceClockTimeOut)
}

// Returns a reference to the singleton clock that reflects the host time.
//
// Added in macOS 10.8.
// Returns a reference to the singleton clock that reflects the host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetHostTimeClock()
func CMClockGetHostTimeClock() ClockRef {
	return _CMClockGetHostTimeClock()
}

// Returns the current time from a clock.
//
// Added in macOS 10.8.
// Returns the current time from a clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTime(_:)
func CMClockGetTime(clock ClockRef) Time {
	return _CMClockGetTime(clock)
}

// Returns the core foundation type identifier of a clock type.
//
// Added in macOS 10.8.
// Returns the core foundation type identifier of a clock type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockGetTypeID()
func CMClockGetTypeID() TypeID {
	return _CMClockGetTypeID()
}

// Stops the clock.
//
// Added in macOS 10.8.
// Stops the clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockInvalidate(_:)
func CMClockInvalidate(clock ClockRef) {
	_CMClockInvalidate(clock)
}

// Converts a host time from native units to a core media time structure.
//
// Added in macOS 10.8.
// Converts a host time from native units to a core media time structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockMakeHostTimeFromSystemUnits(_:)
func CMClockMakeHostTimeFromSystemUnits(hostTime uint64) Time {
	return _CMClockMakeHostTimeFromSystemUnits(hostTime)
}

// Returns a Boolean value that indicates whether it’s possible for two clocks to drift relative to each other.
//
// Added in macOS 10.8.
// Returns a Boolean value that indicates whether it’s possible for two clocks to drift relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClockMightDrift(_:otherClock:)
func CMClockMightDrift(clock ClockRef, otherClock ClockRef) unsafe.Pointer {
	return _CMClockMightDrift(clock, otherClock)
}

// Copies the contents of a closed caption format description to a buffer in big-endian byte order.
//
// Added in macOS 10.10.
// Copies the contents of a closed caption format description to a buffer in big-endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator:closedCaptionFormatDescription:flavor:blockBufferOut:)
func CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator AllocatorRef, closedCaptionFormatDescription ClosedCaptionFormatDescriptionRef, flavor ClosedCaptionDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionFormatDescription, flavor, blockBufferOut)
}

// Creates a closed caption format description from a big-endian closed caption description structure in a buffer.
//
// Added in macOS 10.10.
// Creates a closed caption format description from a big-endian closed caption description structure in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator:bigEndianClosedCaptionDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator AllocatorRef, closedCaptionDescriptionBlockBuffer BlockBufferRef, flavor ClosedCaptionDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

// Creates a closed caption format description from a big-endian closed caption description structure.
//
// Added in macOS 10.10.
// Creates a closed caption format description from a big-endian closed caption description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator:bigEndianClosedCaptionDescriptionData:size:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator AllocatorRef, closedCaptionDescriptionData unsafe.Pointer, size uintptr, flavor ClosedCaptionDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator, closedCaptionDescriptionData, size, flavor, formatDescriptionOut)
}

// Returns a dictionary of all attachments for an attachment bearer object.
//
// Added in macOS 10.7.
// Returns a dictionary of all attachments for an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMCopyDictionaryOfAttachments(allocator:target:attachmentMode:)
func CMCopyDictionaryOfAttachments(allocator AllocatorRef, target AttachmentBearerRef, attachmentMode AttachmentMode) DictionaryRef {
	return _CMCopyDictionaryOfAttachments(allocator, target, attachmentMode)
}

// Returns a Boolean value that indicates whether the sample tables need to use the legacy constant bit-rate encoding layout.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether the sample tables need to use the legacy constant bit-rate encoding layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(_:flavor:)
func CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer BlockBufferRef, flavor SoundDescriptionFlavor) unsafe.Pointer {
	return _CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer, flavor)
}

// Creates a format description for general use.
//
// Added in macOS 10.7.
// Creates a format description for general use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionCreate(allocator:mediaType:mediaSubType:extensions:formatDescriptionOut:)
func CMFormatDescriptionCreate(allocator AllocatorRef, mediaType MediaType, mediaSubType unsafe.Pointer, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMFormatDescriptionCreate(allocator, mediaType, mediaSubType, extensions, formatDescriptionOut)
}

// Returns a Boolean value that indicates whether two format descriptions are equal.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether two format descriptions are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqual(_:otherFormatDescription:)
func CMFormatDescriptionEqual(formatDescription FormatDescriptionRef, otherFormatDescription FormatDescriptionRef) unsafe.Pointer {
	return _CMFormatDescriptionEqual(formatDescription, otherFormatDescription)
}

// Returns a Boolean value that indicates whether two format descriptions are equal, ignoring differences in the extension keys you specify.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether two format descriptions are equal, ignoring differences in the extension keys you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqualIgnoringExtensionKeys(_:otherFormatDescription:extensionKeysToIgnore:sampleDescriptionExtensionAtomKeysToIgnore:)
func CMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription FormatDescriptionRef, otherFormatDescription FormatDescriptionRef, formatDescriptionExtensionKeysToIgnore TypeRef, sampleDescriptionExtensionAtomKeysToIgnore TypeRef) unsafe.Pointer {
	return _CMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription, otherFormatDescription, formatDescriptionExtensionKeysToIgnore, sampleDescriptionExtensionAtomKeysToIgnore)
}

// Returns an extension from the format description by using an extension key.
//
// Added in macOS 10.7.
// Returns an extension from the format description by using an extension key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtension(_:extensionKey:)
func CMFormatDescriptionGetExtension(desc FormatDescriptionRef, extensionKey StringRef) PropertyListRef {
	return _CMFormatDescriptionGetExtension(desc, extensionKey)
}

// Returns all of the extensions for a format description.
//
// Added in macOS 10.7.
// Returns all of the extensions for a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtensions(_:)
func CMFormatDescriptionGetExtensions(desc FormatDescriptionRef) DictionaryRef {
	return _CMFormatDescriptionGetExtensions(desc)
}

// Returns the media subtype of a format description.
//
// Added in macOS 10.7.
// Returns the media subtype of a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
func CMFormatDescriptionGetMediaSubType(desc FormatDescriptionRef) unsafe.Pointer {
	return _CMFormatDescriptionGetMediaSubType(desc)
}

// Returns the media type of a format description.
//
// Added in macOS 10.7.
// Returns the media type of a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaType(_:)
func CMFormatDescriptionGetMediaType(desc FormatDescriptionRef) MediaType {
	return _CMFormatDescriptionGetMediaType(desc)
}

// Returns the Core Foundation type identifier that identifies format description objects.
//
// Added in macOS 10.7.
// Returns the Core Foundation type identifier that identifies format description objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetTypeID()
func CMFormatDescriptionGetTypeID() TypeID {
	return _CMFormatDescriptionGetTypeID()
}

// Returns an attachment from an attachment bearer object.
//
// Added in macOS 10.7.
// Returns an attachment from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMGetAttachment(_:key:attachmentModeOut:)
func CMGetAttachment(target AttachmentBearerRef, key StringRef, attachmentModeOut unsafe.Pointer) TypeRef {
	return _CMGetAttachment(target, key, attachmentModeOut)
}

// Creates a memory pool.
//
// Added in macOS 10.8.
// Creates a memory pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolCreate(options:)
func CMMemoryPoolCreate(options DictionaryRef) MemoryPoolRef {
	return _CMMemoryPoolCreate(options)
}

// Deallocates all memory the pool holds.
//
// Added in macOS 10.8.
// Deallocates all memory the pool holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolFlush(_:)
func CMMemoryPoolFlush(pool MemoryPoolRef) {
	_CMMemoryPoolFlush(pool)
}

// Returns the allocator for the memory pool.
//
// Added in macOS 10.8.
// Returns the allocator for the memory pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetAllocator(_:)
func CMMemoryPoolGetAllocator(pool MemoryPoolRef) AllocatorRef {
	return _CMMemoryPoolGetAllocator(pool)
}

// Returns the type identifier of memory pool objects.
//
// Added in macOS 10.8.
// Returns the type identifier of memory pool objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() TypeID {
	return _CMMemoryPoolGetTypeID()
}

// Invalidates the memory pool, which causes its allocator to stop recycling memory.
//
// Added in macOS 10.8.
// Invalidates the memory pool, which causes its allocator to stop recycling memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolInvalidate(_:)
func CMMemoryPoolInvalidate(pool MemoryPoolRef) {
	_CMMemoryPoolInvalidate(pool)
}

// Creates a URL-like string identifier that represents a key or keyspace tuple.
//
// Added in macOS 10.10.
// Creates a URL-like string identifier that represents a key or keyspace tuple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateIdentifierForKeyAndKeySpace(allocator:key:keySpace:identifierOut:)
func CMMetadataCreateIdentifierForKeyAndKeySpace(allocator AllocatorRef, key TypeRef, keySpace StringRef, identifierOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateIdentifierForKeyAndKeySpace(allocator, key, keySpace, identifierOut)
}

// Creates a copy of the key by using an identifier.
//
// Added in macOS 10.10.
// Creates a copy of the key by using an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifier(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifier(allocator AllocatorRef, identifier StringRef, keyOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateKeyFromIdentifier(allocator, identifier, keyOut)
}

// Creates a copy of the key by using an identifier, and results in a core foundation data object.
//
// Added in macOS 10.10.
// Creates a copy of the key by using an identifier, and results in a core foundation data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifierAsCFData(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifierAsCFData(allocator AllocatorRef, identifier StringRef, keyOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateKeyFromIdentifierAsCFData(allocator, identifier, keyOut)
}

// Creates a copy of the keyspace by using an identifier.
//
// Added in macOS 10.10.
// Creates a copy of the keyspace by using an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeySpaceFromIdentifier(allocator:identifier:keySpaceOut:)
func CMMetadataCreateKeySpaceFromIdentifier(allocator AllocatorRef, identifier StringRef, keySpaceOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataCreateKeySpaceFromIdentifier(allocator, identifier, keySpaceOut)
}

// Returns a Boolean value that indicates whether a data type conforms to another data type.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether a data type conforms to another data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType StringRef, conformsToDataType StringRef) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
}

// Returns a Boolean value that indicates whether a data type identifier represents a base data type.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether a data type identifier represents a base data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsBaseDataType(_:)
func CMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType StringRef) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType)
}

// Returns a Boolean value that indicates the registration status of a data type identifier.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates the registration status of a data type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsRegistered(_:)
func CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType StringRef) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType)
}

// Returns the base data type identifier that a data type conforms to.
//
// Added in macOS 10.10.
// Returns the base data type identifier that a data type conforms to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(_:)
func CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType StringRef) StringRef {
	return _CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType)
}

// Returns an array of base data type identifiers.
//
// Added in macOS 10.10.
// Returns an array of base data type identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypes()
func CMMetadataDataTypeRegistryGetBaseDataTypes() ArrayRef {
	return _CMMetadataDataTypeRegistryGetBaseDataTypes()
}

// Returns the conforming data types for the data type, if any.
//
// Added in macOS 10.10.
// Returns the conforming data types for the data type, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetConformingDataTypes(_:)
func CMMetadataDataTypeRegistryGetConformingDataTypes(dataType StringRef) ArrayRef {
	return _CMMetadataDataTypeRegistryGetConformingDataTypes(dataType)
}

// Returns the data type description if it exists.
//
// Added in macOS 10.10.
// Returns the data type description if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetDataTypeDescription(_:)
func CMMetadataDataTypeRegistryGetDataTypeDescription(dataType StringRef) StringRef {
	return _CMMetadataDataTypeRegistryGetDataTypeDescription(dataType)
}

// Register a data type with the data type registry.
//
// Added in macOS 10.10.
// Register a data type with the data type registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryRegisterDataType(_:description:conformingDataTypes:)
func CMMetadataDataTypeRegistryRegisterDataType(dataType StringRef, description StringRef, conformingDataTypes ArrayRef) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryRegisterDataType(dataType, description, conformingDataTypes)
}

// Copies the contents of a metadata format description to a buffer in big-endian byte order.
//
// Added in macOS 10.10.
// Copies the contents of a metadata format description to a buffer in big-endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator:metadataFormatDescription:flavor:blockBufferOut:)
func CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator AllocatorRef, metadataFormatDescription MetadataFormatDescriptionRef, flavor MetadataDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator, metadataFormatDescription, flavor, blockBufferOut)
}

// Creates a metadata format description object by merging with another description.
//
// Added in macOS 10.10.
// Creates a metadata format description object by merging with another description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator:sourceDescription:otherSourceDescription:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator AllocatorRef, sourceDescription MetadataFormatDescriptionRef, otherSourceDescription MetadataFormatDescriptionRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator, sourceDescription, otherSourceDescription, formatDescriptionOut)
}

// Creates a metadata format description from a big-endian metadata description structure inside a buffer.
//
// Added in macOS 10.10.
// Creates a metadata format description from a big-endian metadata description structure inside a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator:bigEndianMetadataDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator AllocatorRef, metadataDescriptionBlockBuffer BlockBufferRef, flavor MetadataDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator, metadataDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

// Creates a metadata format description from a big-endian metadata description structure.
//
// Added in macOS 10.10.
// Creates a metadata format description from a big-endian metadata description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator:bigEndianMetadataDescriptionData:size:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator AllocatorRef, metadataDescriptionData unsafe.Pointer, size uintptr, flavor MetadataDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator, metadataDescriptionData, size, flavor, formatDescriptionOut)
}

// Creates a metadata format description with the metadata keys you specify.
//
// Added in macOS 10.7.
// Creates a metadata format description with the metadata keys you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithKeys(allocator:metadataType:keys:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithKeys(allocator AllocatorRef, metadataType MetadataFormatType, keys ArrayRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateWithKeys(allocator, metadataType, keys, formatDescriptionOut)
}

// Creates a metadata format description by extending an existing description with the values you specify.
//
// Added in macOS 10.10.
// Creates a metadata format description by extending an existing description with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator:sourceDescription:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator AllocatorRef, sourceDescription MetadataFormatDescriptionRef, metadataSpecifications ArrayRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator, sourceDescription, metadataSpecifications, formatDescriptionOut)
}

// Creates a metadata format description with the specifications you specify.
//
// Added in macOS 10.10.
// Creates a metadata format description with the specifications you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator:metadataType:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator AllocatorRef, metadataType MetadataFormatType, metadataSpecifications ArrayRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator, metadataType, metadataSpecifications, formatDescriptionOut)
}

// Returns an array of metadata identifiers from a metadata format description.
//
// Added in macOS 10.10.
// Returns an array of metadata identifiers from a metadata format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetIdentifiers(_:)
func CMMetadataFormatDescriptionGetIdentifiers(desc MetadataFormatDescriptionRef) ArrayRef {
	return _CMMetadataFormatDescriptionGetIdentifiers(desc)
}

// Returns the key for the local identifier.
//
// Added in macOS 10.7.
// Returns the key for the local identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetKeyWithLocalID(_:localKeyID:)
func CMMetadataFormatDescriptionGetKeyWithLocalID(desc MetadataFormatDescriptionRef, localKeyID unsafe.Pointer) DictionaryRef {
	return _CMMetadataFormatDescriptionGetKeyWithLocalID(desc, localKeyID)
}

// Creates a format description for a muxed media stream.
//
// Added in macOS 10.7.
// Creates a format description for a muxed media stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMMuxedFormatDescriptionCreate(allocator:muxType:extensions:formatDescriptionOut:)
func CMMuxedFormatDescriptionCreate(allocator AllocatorRef, muxType MuxedStreamType, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMMuxedFormatDescriptionCreate(allocator, muxType, extensions, formatDescriptionOut)
}

// Copies all propagable attachments from one attachment bearer object to another.
//
// Added in macOS 10.7.
// Copies all propagable attachments from one attachment bearer object to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPropagateAttachments(_:destination:)
func CMPropagateAttachments(source AttachmentBearerRef, destination AttachmentBearerRef) {
	_CMPropagateAttachments(source, destination)
}

// Removes all attachments from an attachment bearer object.
//
// Added in macOS 10.7.
// Removes all attachments from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAllAttachments(_:)
func CMRemoveAllAttachments(target AttachmentBearerRef) {
	_CMRemoveAllAttachments(target)
}

// Removes a specific attachment from an attachment bearer object.
//
// Added in macOS 10.7.
// Removes a specific attachment from an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMRemoveAttachment(_:key:)
func CMRemoveAttachment(target AttachmentBearerRef, key StringRef) {
	_CMRemoveAttachment(target, key)
}

// Calls a block for every individual sample in a sample buffer.
//
// Added in macOS 10.10.
// Calls a block for every individual sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallBlockForEachSample(_:_:)
func CMSampleBufferCallBlockForEachSample(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferCallBlockForEachSample(sbuf)
}

// Calls a function for every individual sample in a sample buffer.
//
// Added in macOS 10.7.
// Calls a function for every individual sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallForEachSample(_:callback:refcon:)
func CMSampleBufferCallForEachSample(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferCallForEachSample(sbuf)
}

// Copies PCM audio data from a sample buffer into an audio buffer list.
//
// Added in macOS 10.9.
// Copies PCM audio data from a sample buffer into an audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopyPCMDataIntoAudioBufferList(_:at:frameCount:into:)
func CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf SampleBufferRef, frameOffset int32, numFrames int32, bufferList unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf, frameOffset, numFrames, bufferList)
}

// Creates a sample buffer that contains a range of samples from an existing sample buffer.
//
// Added in macOS 10.7.
// Creates a sample buffer that contains a range of samples from an existing sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopySampleBufferForRange(allocator:sampleBuffer:sampleRange:sampleBufferOut:)
func CMSampleBufferCopySampleBufferForRange(allocator AllocatorRef, sbuf SampleBufferRef, sampleRange foundation.Range, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCopySampleBufferForRange(allocator, sbuf, sampleRange, sampleBufferOut)
}

// Creates a sample buffer with a callback to make the data ready for use.
//
// Added in macOS 10.7.
// Creates a sample buffer with a callback to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreate(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreate(allocator AllocatorRef, dataBuffer BlockBufferRef, dataReady unsafe.Pointer, makeDataReadyCallback SampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription FormatDescriptionRef, numSamples ItemCount, numSampleTimingEntries ItemCount, sampleTimingArray unsafe.Pointer, numSampleSizeEntries ItemCount, sampleSizeArray unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreate(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
}

// Creates a copy of a sample buffer.
//
// Added in macOS 10.7.
// Creates a copy of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopy(allocator:sampleBuffer:sampleBufferOut:)
func CMSampleBufferCreateCopy(allocator AllocatorRef, sbuf SampleBufferRef, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateCopy(allocator, sbuf, sampleBufferOut)
}

// Creates a copy of a sample buffer with new timing information.
//
// Added in macOS 10.7.
// Creates a copy of a sample buffer with new timing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopyWithNewTiming(allocator:sampleBuffer:sampleTimingEntryCount:sampleTimingArray:sampleBufferOut:)
func CMSampleBufferCreateCopyWithNewTiming(allocator AllocatorRef, originalSBuf SampleBufferRef, numSampleTimingEntries ItemCount, sampleTimingArray unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateCopyWithNewTiming(allocator, originalSBuf, numSampleTimingEntries, sampleTimingArray, sampleBufferOut)
}

// Creates a sample buffer with an image buffer and a callback to make the data ready for use.
//
// Added in macOS 10.7.
// Creates a sample buffer with an image buffer and a callback to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBuffer(allocator:imageBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateForImageBuffer(allocator AllocatorRef, imageBuffer ImageBufferRef, dataReady unsafe.Pointer, makeDataReadyCallback SampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription VideoFormatDescriptionRef, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateForImageBuffer(allocator, imageBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, sampleTiming, sampleBufferOut)
}

// Creates a sample buffer with an image buffer and a handler to make the data ready for use.
//
// Added in macOS 10.14.4.
// Creates a sample buffer with an image buffer and a handler to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(_:_:_:_:_:_:_:)
func CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator AllocatorRef, imageBuffer ImageBufferRef, dataReady unsafe.Pointer, formatDescription VideoFormatDescriptionRef, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler SampleBufferMakeDataReadyHandler) unsafe.Pointer {
	return _CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator, imageBuffer, dataReady, formatDescription, sampleTiming, sampleBufferOut, makeDataReadyHandler)
}

// Creates a new sample buffer from a tagged buffer group.
//
// Added in macOS 14.0.
// Creates a new sample buffer from a tagged buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForTaggedBufferGroup
func CMSampleBufferCreateForTaggedBufferGroup(allocator AllocatorRef, taggedBufferGroup TaggedBufferGroupRef, sbufPTS Time, sbufDuration Time, formatDescription TaggedBufferGroupFormatDescriptionRef, sBufOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateForTaggedBufferGroup(allocator, taggedBufferGroup, sbufPTS, sbufDuration, formatDescription, sBufOut)
}

// Creates a sample buffer with media data.
//
// Added in macOS 10.10.
// Creates a sample buffer with media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReady(allocator:dataBuffer:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreateReady(allocator AllocatorRef, dataBuffer BlockBufferRef, formatDescription FormatDescriptionRef, numSamples ItemCount, numSampleTimingEntries ItemCount, sampleTimingArray unsafe.Pointer, numSampleSizeEntries ItemCount, sampleSizeArray unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateReady(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
}

// Creates a sample buffer with image data.
//
// Added in macOS 10.10.
// Creates a sample buffer with image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReadyWithImageBuffer(allocator:imageBuffer:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateReadyWithImageBuffer(allocator AllocatorRef, imageBuffer ImageBufferRef, formatDescription VideoFormatDescriptionRef, sampleTiming unsafe.Pointer, sampleBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferCreateReadyWithImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sampleBufferOut)
}

// Creates a sample buffer with a handler to make the data ready for use.
//
// Added in macOS 10.14.4.
// Creates a sample buffer with a handler to make the data ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateWithMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:_:_:)
func CMSampleBufferCreateWithMakeDataReadyHandler(allocator AllocatorRef, dataBuffer BlockBufferRef, dataReady unsafe.Pointer, formatDescription FormatDescriptionRef, numSamples ItemCount, numSampleTimingEntries ItemCount, sampleTimingArray unsafe.Pointer, numSampleSizeEntries ItemCount, sampleSizeArray unsafe.Pointer, sampleBufferOut unsafe.Pointer, makeDataReadyHandler SampleBufferMakeDataReadyHandler) unsafe.Pointer {
	return _CMSampleBufferCreateWithMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut, makeDataReadyHandler)
}

// Returns a Boolean value that indicates whether the sample buffer’s data is ready for use.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether the sample buffer’s data is ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferDataIsReady(_:)
func CMSampleBufferDataIsReady(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferDataIsReady(sbuf)
}

// Returns an audio buffer list that contains the media data.
//
// Added in macOS 10.7.
// Returns an audio buffer list that contains the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(_:bufferListSizeNeededOut:bufferListOut:bufferListSize:blockBufferAllocator:blockBufferMemoryAllocator:flags:blockBufferOut:)
func CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf SampleBufferRef, bufferListSizeNeededOut unsafe.Pointer, bufferListOut unsafe.Pointer, bufferListSize uintptr, blockBufferStructureAllocator AllocatorRef, blockBufferBlockAllocator AllocatorRef, flags uint32, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf, bufferListSizeNeededOut, bufferListOut, bufferListSize, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, blockBufferOut)
}

// Creates an array of audio stream packet descriptions.
//
// Added in macOS 10.7.
// Creates an array of audio stream packet descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptions(_:allocatedSize:packetDescriptionsOut:packetDescriptionsSizeNeededOut:)
func CMSampleBufferGetAudioStreamPacketDescriptions(sbuf SampleBufferRef, packetDescriptionsSize uintptr, packetDescriptionsOut unsafe.Pointer, packetDescriptionsSizeNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioStreamPacketDescriptions(sbuf, packetDescriptionsSize, packetDescriptionsOut, packetDescriptionsSizeNeededOut)
}

// Returns a pointer to a constant array of audio stream packet descriptions.
//
// Added in macOS 10.7.
// Returns a pointer to a constant array of audio stream packet descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptionsPtr(_:packetDescriptionsPointerOut:sizeOut:)
func CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf SampleBufferRef, packetDescriptionsPointerOut unsafe.Pointer, packetDescriptionsSizeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf, packetDescriptionsPointerOut, packetDescriptionsSizeOut)
}

// Returns a block buffer that contains the media data.
//
// Added in macOS 10.7.
// Returns a block buffer that contains the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDataBuffer(_:)
func CMSampleBufferGetDataBuffer(sbuf SampleBufferRef) BlockBufferRef {
	return _CMSampleBufferGetDataBuffer(sbuf)
}

// Returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// Added in macOS 10.7.
// Returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDecodeTimeStamp(_:)
func CMSampleBufferGetDecodeTimeStamp(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetDecodeTimeStamp(sbuf)
}

// Returns the total duration of a sample buffer.
//
// Added in macOS 10.7.
// Returns the total duration of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDuration(_:)
func CMSampleBufferGetDuration(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetDuration(sbuf)
}

// Returns the format description of the samples in a sample buffer.
//
// Added in macOS 10.7.
// Returns the format description of the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetFormatDescription(_:)
func CMSampleBufferGetFormatDescription(sbuf SampleBufferRef) FormatDescriptionRef {
	return _CMSampleBufferGetFormatDescription(sbuf)
}

// Returns an image buffer that contains the media data.
//
// Added in macOS 10.7.
// Returns an image buffer that contains the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetImageBuffer(_:)
func CMSampleBufferGetImageBuffer(sbuf SampleBufferRef) ImageBufferRef {
	return _CMSampleBufferGetImageBuffer(sbuf)
}

// Returns the number of media samples in a sample buffer.
//
// Added in macOS 10.7.
// Returns the number of media samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetNumSamples(_:)
func CMSampleBufferGetNumSamples(sbuf SampleBufferRef) ItemCount {
	return _CMSampleBufferGetNumSamples(sbuf)
}

// Returns the output decode timestamp of a sample buffer.
//
// Added in macOS 10.7.
// Returns the output decode timestamp of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDecodeTimeStamp(_:)
func CMSampleBufferGetOutputDecodeTimeStamp(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetOutputDecodeTimeStamp(sbuf)
}

// Returns the output duration of a sample buffer.
//
// Added in macOS 10.7.
// Returns the output duration of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDuration(_:)
func CMSampleBufferGetOutputDuration(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetOutputDuration(sbuf)
}

// Returns the output presentation timestamp of a sample buffer.
//
// Added in macOS 10.7.
// Returns the output presentation timestamp of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputPresentationTimeStamp(_:)
func CMSampleBufferGetOutputPresentationTimeStamp(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetOutputPresentationTimeStamp(sbuf)
}

// Retrieves an array of output timing information structures that represents each sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves an array of output timing information structures that represents each sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetOutputSampleTimingInfoArray(sbuf SampleBufferRef, timingArrayEntries ItemCount, timingArrayOut unsafe.Pointer, timingArrayEntriesNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetOutputSampleTimingInfoArray(sbuf, timingArrayEntries, timingArrayOut, timingArrayEntriesNeededOut)
}

// Returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// Added in macOS 10.7.
// Returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetPresentationTimeStamp(_:)
func CMSampleBufferGetPresentationTimeStamp(sbuf SampleBufferRef) Time {
	return _CMSampleBufferGetPresentationTimeStamp(sbuf)
}

// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf SampleBufferRef, createIfNecessary unsafe.Pointer) ArrayRef {
	return _CMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
}

// Returns the size in bytes of a specified sample in a sample buffer.
//
// Added in macOS 10.7.
// Returns the size in bytes of a specified sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSize(_:at:)
func CMSampleBufferGetSampleSize(sbuf SampleBufferRef, sampleIndex ItemIndex) uintptr {
	return _CMSampleBufferGetSampleSize(sbuf, sampleIndex)
}

// Retrieves an array of sample sizes that represents each sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves an array of sample sizes that represents each sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSizeArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleSizeArray(sbuf SampleBufferRef, sizeArrayEntries ItemCount, sizeArrayOut unsafe.Pointer, sizeArrayEntriesNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleSizeArray(sbuf, sizeArrayEntries, sizeArrayOut, sizeArrayEntriesNeededOut)
}

// Retrieves a timing information structure that describes a specified sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves a timing information structure that describes a specified sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfo(_:at:timingInfoOut:)
func CMSampleBufferGetSampleTimingInfo(sbuf SampleBufferRef, sampleIndex ItemIndex, timingInfoOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleTimingInfo(sbuf, sampleIndex, timingInfoOut)
}

// Retrieves an array of sample timing information structures that represents each sample in a sample buffer.
//
// Added in macOS 10.7.
// Retrieves an array of sample timing information structures that represents each sample in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleTimingInfoArray(sbuf SampleBufferRef, numSampleTimingEntries ItemCount, timingArrayOut unsafe.Pointer, timingArrayEntriesNeededOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleTimingInfoArray(sbuf, numSampleTimingEntries, timingArrayOut, timingArrayEntriesNeededOut)
}

// Gets the tagged buffer group of a sample buffer.
//
// Added in macOS 14.0.
// Gets the tagged buffer group of a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTaggedBufferGroup
func CMSampleBufferGetTaggedBufferGroup(sbuf SampleBufferRef) TaggedBufferGroupRef {
	return _CMSampleBufferGetTaggedBufferGroup(sbuf)
}

// Returns the total size in bytes of sample data in a sample buffer.
//
// Added in macOS 10.7.
// Returns the total size in bytes of sample data in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTotalSampleSize(_:)
func CMSampleBufferGetTotalSampleSize(sbuf SampleBufferRef) uintptr {
	return _CMSampleBufferGetTotalSampleSize(sbuf)
}

// Returns the type identifier of sample buffer objects.
//
// Added in macOS 10.7.
// Returns the type identifier of sample buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTypeID()
func CMSampleBufferGetTypeID() TypeID {
	return _CMSampleBufferGetTypeID()
}

// Returns a Boolean value that indicates whether the sample buffer’s data loading request failed.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether the sample buffer’s data loading request failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferHasDataFailed(_:statusOut:)
func CMSampleBufferHasDataFailed(sbuf SampleBufferRef, statusOut unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferHasDataFailed(sbuf, statusOut)
}

// Invalidates a sample buffer by calling its invalidation callback.
//
// Added in macOS 10.7.
// Invalidates a sample buffer by calling its invalidation callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidate(_:)
func CMSampleBufferInvalidate(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferInvalidate(sbuf)
}

// Returns a Boolean value that indicates whether a sample buffer is valid.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a sample buffer is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferIsValid(_:)
func CMSampleBufferIsValid(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferIsValid(sbuf)
}

// Makes the sample buffer’s data ready for use by invoking its callback to load the data.
//
// Added in macOS 10.7.
// Makes the sample buffer’s data ready for use by invoking its callback to load the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReady(_:)
func CMSampleBufferMakeDataReady(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferMakeDataReady(sbuf)
}

// Sets a block buffer of media data on a sample buffer.
//
// Added in macOS 10.7.
// Sets a block buffer of media data on a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBuffer(_:newValue:)
func CMSampleBufferSetDataBuffer(sbuf SampleBufferRef, dataBuffer BlockBufferRef) unsafe.Pointer {
	return _CMSampleBufferSetDataBuffer(sbuf, dataBuffer)
}

// Creates a block buffer that contains a copy of the data from an audio buffer list.
//
// Added in macOS 10.7.
// Creates a block buffer that contains a copy of the data from an audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBufferFromAudioBufferList(_:blockBufferAllocator:blockBufferMemoryAllocator:flags:bufferList:)
func CMSampleBufferSetDataBufferFromAudioBufferList(sbuf SampleBufferRef, blockBufferStructureAllocator AllocatorRef, blockBufferBlockAllocator AllocatorRef, flags uint32, bufferList unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataBufferFromAudioBufferList(sbuf, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, bufferList)
}

// Marks the sample buffer’s data as failed to indicate that it won’t become ready.
//
// Added in macOS 10.10.
// Marks the sample buffer’s data as failed to indicate that it won’t become ready.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataFailed(_:status:)
func CMSampleBufferSetDataFailed(sbuf SampleBufferRef, status unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferSetDataFailed(sbuf, status)
}

// Marks a sample buffer’s data as ready for use.
//
// Added in macOS 10.7.
// Marks a sample buffer’s data as ready for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataReady(_:)
func CMSampleBufferSetDataReady(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferSetDataReady(sbuf)
}

// Sets the sample buffer’s invalidation callback.
//
// Added in macOS 10.7.
// Sets the sample buffer’s invalidation callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateCallback(_:callback:refcon:)
func CMSampleBufferSetInvalidateCallback(sbuf SampleBufferRef, invalidateCallback SampleBufferInvalidateCallback, invalidateRefCon uint64) unsafe.Pointer {
	return _CMSampleBufferSetInvalidateCallback(sbuf, invalidateCallback, invalidateRefCon)
}

// Sets the sample buffer’s invalidation handler.
//
// Added in macOS 10.10.
// Sets the sample buffer’s invalidation handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateHandler(_:invalidateHandler:)
func CMSampleBufferSetInvalidateHandler(sbuf SampleBufferRef, invalidateHandler SampleBufferInvalidateHandler) unsafe.Pointer {
	return _CMSampleBufferSetInvalidateHandler(sbuf, invalidateHandler)
}

// Sets an output presentation timestamp to use in place of a calculated value.
//
// Added in macOS 10.7.
// Sets an output presentation timestamp to use in place of a calculated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetOutputPresentationTimeStamp(_:newValue:)
func CMSampleBufferSetOutputPresentationTimeStamp(sbuf SampleBufferRef, outputPresentationTimeStamp Time) unsafe.Pointer {
	return _CMSampleBufferSetOutputPresentationTimeStamp(sbuf, outputPresentationTimeStamp)
}

// Associates a sample buffer’s data readiness with that of another sample buffer.
//
// Added in macOS 10.7.
// Associates a sample buffer’s data readiness with that of another sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferTrackDataReadiness(_:sampleBufferToTrack:)
func CMSampleBufferTrackDataReadiness(sbuf SampleBufferRef, sampleBufferToTrack SampleBufferRef) unsafe.Pointer {
	return _CMSampleBufferTrackDataReadiness(sbuf, sampleBufferToTrack)
}

// Sets or adds an attachment to an attachment bearer object.
//
// Added in macOS 10.7.
// Sets or adds an attachment to an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachment(_:key:value:attachmentMode:)
func CMSetAttachment(target AttachmentBearerRef, key StringRef, value TypeRef, attachmentMode AttachmentMode) {
	_CMSetAttachment(target, key, value, attachmentMode)
}

// Sets a dictionary of attachments on an attachment bearer object.
//
// Added in macOS 10.7.
// Sets a dictionary of attachments on an attachment bearer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSetAttachments(_:attachments:attachmentMode:)
func CMSetAttachments(target AttachmentBearerRef, theAttachments DictionaryRef, attachmentMode AttachmentMode) {
	_CMSetAttachments(target, theAttachments, attachmentMode)
}

// Creates a queue that has the specified capacity.
//
// Added in macOS 10.7.
// Creates a queue that has the specified capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueCreate(allocator:capacity:queueOut:)
func CMSimpleQueueCreate(allocator AllocatorRef, capacity int32, queueOut unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueCreate(allocator, capacity, queueOut)
}

// Dequeues an element from the queue.
//
// Added in macOS 10.7.
// Dequeues an element from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueDequeue(_:)
func CMSimpleQueueDequeue(queue SimpleQueueRef) unsafe.Pointer {
	return _CMSimpleQueueDequeue(queue)
}

// Enqueues an element in the queue.
//
// Added in macOS 10.7.
// Enqueues an element in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueEnqueue(_:element:)
func CMSimpleQueueEnqueue(queue SimpleQueueRef, element unsafe.Pointer) unsafe.Pointer {
	return _CMSimpleQueueEnqueue(queue, element)
}

// Returns the number of elements that the queue can hold.
//
// Added in macOS 10.7.
// Returns the number of elements that the queue can hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCapacity(_:)
func CMSimpleQueueGetCapacity(queue SimpleQueueRef) int32 {
	return _CMSimpleQueueGetCapacity(queue)
}

// Returns the number of elements currently in the queue.
//
// Added in macOS 10.7.
// Returns the number of elements currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCount(_:)
func CMSimpleQueueGetCount(queue SimpleQueueRef) int32 {
	return _CMSimpleQueueGetCount(queue)
}

// Returns the element at the head of the queue.
//
// Added in macOS 10.7.
// Returns the element at the head of the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetHead(_:)
func CMSimpleQueueGetHead(queue SimpleQueueRef) unsafe.Pointer {
	return _CMSimpleQueueGetHead(queue)
}

// Returns the type identifier of sample buffer objects.
//
// Added in macOS 10.7.
// Returns the type identifier of sample buffer objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() TypeID {
	return _CMSimpleQueueGetTypeID()
}

// Resets the queue.
//
// Added in macOS 10.7.
// Resets the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueReset(_:)
func CMSimpleQueueReset(queue SimpleQueueRef) unsafe.Pointer {
	return _CMSimpleQueueReset(queue)
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

// Converts an image description data structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts an image description data structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianImageDescriptionToHost(_:_:)
func CMSwapBigEndianImageDescriptionToHost(imageDescriptionData unsafe.Pointer, imageDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianImageDescriptionToHost(imageDescriptionData, imageDescriptionSize)
}

// Converts a metadata description data structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts a metadata description data structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianMetadataDescriptionToHost(_:_:)
func CMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData unsafe.Pointer, metadataDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData, metadataDescriptionSize)
}

// Converts a sound description data structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts a sound description data structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianSoundDescriptionToHost(_:_:)
func CMSwapBigEndianSoundDescriptionToHost(soundDescriptionData unsafe.Pointer, soundDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianSoundDescriptionToHost(soundDescriptionData, soundDescriptionSize)
}

// Converts a text description structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts a text description structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTextDescriptionToHost(_:_:)
func CMSwapBigEndianTextDescriptionToHost(textDescriptionData unsafe.Pointer, textDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianTextDescriptionToHost(textDescriptionData, textDescriptionSize)
}

// Converts a time code description data structure from big-endian to host-endian, in place.
//
// Added in macOS 10.10.
// Converts a time code description data structure from big-endian to host-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTimeCodeDescriptionToHost(_:_:)
func CMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData unsafe.Pointer, timeCodeDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData, timeCodeDescriptionSize)
}

// Converts a closed caption description structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts a closed caption description structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianClosedCaptionDescriptionToBig(_:_:)
func CMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData unsafe.Pointer, closedCaptionDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData, closedCaptionDescriptionSize)
}

// Converts an image description data structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts an image description data structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianImageDescriptionToBig(_:_:)
func CMSwapHostEndianImageDescriptionToBig(imageDescriptionData unsafe.Pointer, imageDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianImageDescriptionToBig(imageDescriptionData, imageDescriptionSize)
}

// Converts a metadata description data structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts a metadata description data structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianMetadataDescriptionToBig(_:_:)
func CMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData unsafe.Pointer, metadataDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData, metadataDescriptionSize)
}

// Converts a sound description data structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts a sound description data structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianSoundDescriptionToBig(_:_:)
func CMSwapHostEndianSoundDescriptionToBig(soundDescriptionData unsafe.Pointer, soundDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianSoundDescriptionToBig(soundDescriptionData, soundDescriptionSize)
}

// Converts a text description structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts a text description structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTextDescriptionToBig(_:_:)
func CMSwapHostEndianTextDescriptionToBig(textDescriptionData unsafe.Pointer, textDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianTextDescriptionToBig(textDescriptionData, textDescriptionSize)
}

// Converts a time code description data structure from host-endian to big-endian, in place.
//
// Added in macOS 10.10.
// Converts a time code description data structure from host-endian to big-endian, in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTimeCodeDescriptionToBig(_:_:)
func CMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData unsafe.Pointer, timeCodeDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData, timeCodeDescriptionSize)
}

// Converts a time from one timebase or clock to another timebase or clock.
//
// Added in macOS 10.8.
// Converts a time from one timebase or clock to another timebase or clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time Time, fromClockOrTimebase ClockOrTimebaseRef, toClockOrTimebase ClockOrTimebaseRef) Time {
	return _CMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
}

// Returns the relative rate of one timebase or clock relative to another timebase or clock.
//
// Added in macOS 10.8.
// Returns the relative rate of one timebase or clock relative to another timebase or clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRate(_:relativeTo:)
func CMSyncGetRelativeRate(ofClockOrTimebase ClockOrTimebaseRef, relativeToClockOrTimebase ClockOrTimebaseRef) unsafe.Pointer {
	return _CMSyncGetRelativeRate(ofClockOrTimebase, relativeToClockOrTimebase)
}

// Returns the relative rate of one timebase or clock relative to another timebase or clock and the times of each timebase or clock at which the relative rate went into effect.
//
// Added in macOS 10.8.
// Returns the relative rate of one timebase or clock relative to another timebase or clock and the times of each timebase or clock at which the relative rate went into effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRateAndAnchorTime(_:relativeTo:relativeRateOut:anchorTimeOut:relativeToAnchorTimeOut:)
func CMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase ClockOrTimebaseRef, relativeToClockOrTimebase ClockOrTimebaseRef, outRelativeRate unsafe.Pointer, outOfClockOrTimebaseAnchorTime unsafe.Pointer, outRelativeToClockOrTimebaseAnchorTime unsafe.Pointer) unsafe.Pointer {
	return _CMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase, relativeToClockOrTimebase, outRelativeRate, outOfClockOrTimebaseAnchorTime, outRelativeToClockOrTimebaseAnchorTime)
}

// Returns the time from a clock or timebase.
//
// Added in macOS 10.8.
// Returns the time from a clock or timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncGetTime(_:)
func CMSyncGetTime(clockOrTimebase ClockOrTimebaseRef) Time {
	return _CMSyncGetTime(clockOrTimebase)
}

// Returns a Boolean value that indicates whether it’s possible for one timebase or clock to drift relative to the other.
//
// Added in macOS 10.8.
// Returns a Boolean value that indicates whether it’s possible for one timebase or clock to drift relative to the other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMSyncMightDrift(_:_:)
func CMSyncMightDrift(clockOrTimebase1 ClockOrTimebaseRef, clockOrTimebase2 ClockOrTimebaseRef) unsafe.Pointer {
	return _CMSyncMightDrift(clockOrTimebase1, clockOrTimebase2)
}

// Adds a new tag to an existing collection.
//
// Added in macOS 14.0.
// Adds a new tag to an existing collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTag
func CMTagCollectionAddTag(tagCollection MutableTagCollectionRef, tagToAdd Tag) unsafe.Pointer {
	return _CMTagCollectionAddTag(tagCollection, tagToAdd)
}

// Adds the tags contained in a C-style array to a tag collection.
//
// Added in macOS 14.0.
// Adds the tags contained in a C-style array to a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromArray
func CMTagCollectionAddTagsFromArray(tagCollection MutableTagCollectionRef, tags unsafe.Pointer, tagCount ItemCount) unsafe.Pointer {
	return _CMTagCollectionAddTagsFromArray(tagCollection, tags, tagCount)
}

// Add the tags contained in one collection to another.
//
// Added in macOS 14.0.
// Add the tags contained in one collection to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromCollection
func CMTagCollectionAddTagsFromCollection(tagCollection MutableTagCollectionRef, collectionWithTagsToAdd TagCollectionRef) unsafe.Pointer {
	return _CMTagCollectionAddTagsFromCollection(tagCollection, collectionWithTagsToAdd)
}

// Applies a function to all tags in a collection.
//
// Added in macOS 14.0.
// Applies a function to all tags in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApply
func CMTagCollectionApply(tagCollection TagCollectionRef, applier TagCollectionApplierFunction, context unsafe.Pointer) {
	_CMTagCollectionApply(tagCollection, applier, context)
}

// Applies a Boolean function to tags in a collection, stopping when it returns true.
//
// Added in macOS 14.0.
// Applies a Boolean function to tags in a collection, stopping when it returns true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApplyUntil
func CMTagCollectionApplyUntil(tagCollection TagCollectionRef, applier TagCollectionTagFilterFunction, context unsafe.Pointer) Tag {
	return _CMTagCollectionApplyUntil(tagCollection, applier, context)
}

// Determines if a tag collection contains tags for a given category.
//
// Added in macOS 14.0.
// Determines if a tag collection contains tags for a given category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsCategory
func CMTagCollectionContainsCategory(tagCollection TagCollectionRef, category TagCategory) unsafe.Pointer {
	return _CMTagCollectionContainsCategory(tagCollection, category)
}

// Determines if a tag collection contains a subset of tags.
//
// Added in macOS 14.0.
// Determines if a tag collection contains a subset of tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsSpecifiedTags
func CMTagCollectionContainsSpecifiedTags(tagCollection TagCollectionRef, containedTags unsafe.Pointer, containedTagCount ItemCount) unsafe.Pointer {
	return _CMTagCollectionContainsSpecifiedTags(tagCollection, containedTags, containedTagCount)
}

// Determines if a tag collection contains a specific tag.
//
// Added in macOS 14.0.
// Determines if a tag collection contains a specific tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTag
func CMTagCollectionContainsTag(tagCollection TagCollectionRef, tag Tag) unsafe.Pointer {
	return _CMTagCollectionContainsTag(tagCollection, tag)
}

// Determines if one collection of tags contains every tag from another collection.
//
// Added in macOS 14.0.
// Determines if one collection of tags contains every tag from another collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTagsOfCollection
func CMTagCollectionContainsTagsOfCollection(tagCollection TagCollectionRef, containedTagCollection TagCollectionRef) unsafe.Pointer {
	return _CMTagCollectionContainsTagsOfCollection(tagCollection, containedTagCollection)
}

// Creates a new Core Foundation data instance from a tag collection.
//
// Added in macOS 14.0.
// Creates a new Core Foundation data instance from a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsData
func CMTagCollectionCopyAsData(tagCollection TagCollectionRef, allocator AllocatorRef) DataRef {
	return _CMTagCollectionCopyAsData(tagCollection, allocator)
}

// Creates a new Core Foundation dictionary from a tag collection.
//
// Added in macOS 14.0.
// Creates a new Core Foundation dictionary from a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsDictionary
func CMTagCollectionCopyAsDictionary(tagCollection TagCollectionRef, allocator AllocatorRef) DictionaryRef {
	return _CMTagCollectionCopyAsDictionary(tagCollection, allocator)
}

// Retrieves a copy of the tag collection’s description.
//
// Added in macOS 14.0.
// Retrieves a copy of the tag collection’s description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyDescription
func CMTagCollectionCopyDescription(allocator AllocatorRef, tagCollection TagCollectionRef) StringRef {
	return _CMTagCollectionCopyDescription(allocator, tagCollection)
}

// Creates a new tag collection from an existing collection, copying all tags which match a list of categories.
//
// Added in macOS 14.0.
// Creates a new tag collection from an existing collection, copying all tags which match a list of categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyTagsOfCategories
func CMTagCollectionCopyTagsOfCategories(allocator AllocatorRef, tagCollection TagCollectionRef, categories unsafe.Pointer, categoriesCount ItemCount, collectionWithTagsOfCategories unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCopyTagsOfCategories(allocator, tagCollection, categories, categoriesCount, collectionWithTagsOfCategories)
}

// Counts the number of tags in a collection matching an evaluation function.
//
// Added in macOS 14.0.
// Counts the number of tags in a collection matching an evaluation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCountTagsWithFilterFunction
func CMTagCollectionCountTagsWithFilterFunction(tagCollection TagCollectionRef, filterApplier TagCollectionTagFilterFunction, context unsafe.Pointer) ItemCount {
	return _CMTagCollectionCountTagsWithFilterFunction(tagCollection, filterApplier, context)
}

// Creates a new tag collection from an existing C-style array of tags.
//
// Added in macOS 14.0.
// Creates a new tag collection from an existing C-style array of tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreate
func CMTagCollectionCreate(allocator AllocatorRef, tags unsafe.Pointer, tagCount ItemCount, newCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreate(allocator, tags, tagCount, newCollectionOut)
}

// Creates a copy of a tag collection.
//
// Added in macOS 14.0.
// Creates a copy of a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateCopy
func CMTagCollectionCreateCopy(tagCollection TagCollectionRef, allocator AllocatorRef, newCollectionCopyOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateCopy(tagCollection, allocator, newCollectionCopyOut)
}

// Creates a new tag collection with the difference of two existing collections.
//
// Added in macOS 14.0.
// Creates a new tag collection with the difference of two existing collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateDifference
func CMTagCollectionCreateDifference(tagCollectionMinuend TagCollectionRef, tagCollectionSubtrahend TagCollectionRef, tagCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateDifference(tagCollectionMinuend, tagCollectionSubtrahend, tagCollectionOut)
}

// Creates a new tag collection from two existing tag collections, copying elements which are in one collection or the other, but not both.
//
// Added in macOS 14.0.
// Creates a new tag collection from two existing tag collections, copying elements which are in one collection or the other, but not both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateExclusiveOr
func CMTagCollectionCreateExclusiveOr(tagCollection1 TagCollectionRef, tagCollection2 TagCollectionRef, tagCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateExclusiveOr(tagCollection1, tagCollection2, tagCollectionOut)
}

// Creates a new tag collection from an existing Core Foundation data instance.
//
// Added in macOS 14.0.
// Creates a new tag collection from an existing Core Foundation data instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromData
func CMTagCollectionCreateFromData(data DataRef, allocator AllocatorRef, newCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateFromData(data, allocator, newCollectionOut)
}

// Creates a new tag collection from an existing Core Foundation dictionary.
//
// Added in macOS 14.0.
// Creates a new tag collection from an existing Core Foundation dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromDictionary
func CMTagCollectionCreateFromDictionary(dict DictionaryRef, allocator AllocatorRef, newCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateFromDictionary(dict, allocator, newCollectionOut)
}

// Creates a new tag collection containing only the tags from two existing collections which match.
//
// Added in macOS 14.0.
// Creates a new tag collection containing only the tags from two existing collections which match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateIntersection
func CMTagCollectionCreateIntersection(tagCollection1 TagCollectionRef, tagCollection2 TagCollectionRef, tagCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateIntersection(tagCollection1, tagCollection2, tagCollectionOut)
}

// Creates a new, mutable, empty tag collection.
//
// Added in macOS 14.0.
// Creates a new, mutable, empty tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutable
func CMTagCollectionCreateMutable(allocator AllocatorRef, capacity Index, newMutableCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateMutable(allocator, capacity, newMutableCollectionOut)
}

// Creates a new mutable copy from an existing tag collection.
//
// Added in macOS 14.0.
// Creates a new mutable copy from an existing tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutableCopy
func CMTagCollectionCreateMutableCopy(tagCollection TagCollectionRef, allocator AllocatorRef, newMutableCollectionCopyOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateMutableCopy(tagCollection, allocator, newMutableCollectionCopyOut)
}

// Creates a new tag collection containing all tags from two collections without duplicates.
//
// Added in macOS 14.0.
// Creates a new tag collection containing all tags from two collections without duplicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateUnion
func CMTagCollectionCreateUnion(tagCollection1 TagCollectionRef, tagCollection2 TagCollectionRef, tagCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateUnion(tagCollection1, tagCollection2, tagCollectionOut)
}

// Gets the number of tags in a tag collection.
//
// Added in macOS 14.0.
// Gets the number of tags in a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCount
func CMTagCollectionGetCount(tagCollection TagCollectionRef) ItemCount {
	return _CMTagCollectionGetCount(tagCollection)
}

// Retrieves the number of tags in the collection matching a given category.
//
// Added in macOS 14.0.
// Retrieves the number of tags in the collection matching a given category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCountOfCategory
func CMTagCollectionGetCountOfCategory(tagCollection TagCollectionRef, category TagCategory) ItemCount {
	return _CMTagCollectionGetCountOfCategory(tagCollection, category)
}

// Retrieves an arbitrary number of tags from the collection.
//
// Added in macOS 14.0.
// Retrieves an arbitrary number of tags from the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTags
func CMTagCollectionGetTags(tagCollection TagCollectionRef, tagBuffer unsafe.Pointer, tagBufferCount ItemCount, numberOfTagsCopied unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionGetTags(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied)
}

// Retrieves a C-style array of tags with a given category from a tag collection.
//
// Added in macOS 14.0.
// Retrieves a C-style array of tags with a given category from a tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithCategory
func CMTagCollectionGetTagsWithCategory(tagCollection TagCollectionRef, category TagCategory, tagBuffer unsafe.Pointer, tagBufferCount ItemCount, numberOfTagsCopied unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionGetTagsWithCategory(tagCollection, category, tagBuffer, tagBufferCount, numberOfTagsCopied)
}

// Gets all tags in a collection matching an evaluation function.
//
// Added in macOS 14.0.
// Gets all tags in a collection matching an evaluation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithFilterFunction
func CMTagCollectionGetTagsWithFilterFunction(tagCollection TagCollectionRef, tagBuffer unsafe.Pointer, tagBufferCount ItemCount, numberOfTagsCopied unsafe.Pointer, filter TagCollectionTagFilterFunction, context unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionGetTagsWithFilterFunction(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied, filter, context)
}

// Retrieves the internal type ID for tag collections.
//
// Added in macOS 14.0.
// Retrieves the internal type ID for tag collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTypeID
func CMTagCollectionGetTypeID() TypeID {
	return _CMTagCollectionGetTypeID()
}

// Checks if a tag collection has no elements.
//
// Added in macOS 14.0.
// Checks if a tag collection has no elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionIsEmpty
func CMTagCollectionIsEmpty(tagCollection TagCollectionRef) unsafe.Pointer {
	return _CMTagCollectionIsEmpty(tagCollection)
}

// Removes all tags from a collection.
//
// Added in macOS 14.0.
// Removes all tags from a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection MutableTagCollectionRef) unsafe.Pointer {
	return _CMTagCollectionRemoveAllTags(tagCollection)
}

// Removes all tags of a given category from a collection.
//
// Added in macOS 14.0.
// Removes all tags of a given category from a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTagsOfCategory
func CMTagCollectionRemoveAllTagsOfCategory(tagCollection MutableTagCollectionRef, category TagCategory) unsafe.Pointer {
	return _CMTagCollectionRemoveAllTagsOfCategory(tagCollection, category)
}

// Removes a specific tag from a collection.
//
// Added in macOS 14.0.
// Removes a specific tag from a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveTag
func CMTagCollectionRemoveTag(tagCollection MutableTagCollectionRef, tagToRemove Tag) unsafe.Pointer {
	return _CMTagCollectionRemoveTag(tagCollection, tagToRemove)
}

// Compares two tags in terms of partial equality.
//
// Added in macOS 14.0.
// Compares two tags in terms of partial equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCompare
func CMTagCompare(tag1 Tag, tag2 Tag) ComparisonResult {
	return _CMTagCompare(tag1, tag2)
}

// Copies an existing tag to a new dictionary object.
//
// Added in macOS 14.0.
// Copies an existing tag to a new dictionary object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCopyAsDictionary
func CMTagCopyAsDictionary(tag Tag, allocator AllocatorRef) DictionaryRef {
	return _CMTagCopyAsDictionary(tag, allocator)
}

// Copies the description of a tag to a new string.
//
// Added in macOS 14.0.
// Copies the description of a tag to a new string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCopyDescription
func CMTagCopyDescription(allocator AllocatorRef, tag Tag) StringRef {
	return _CMTagCopyDescription(allocator, tag)
}

// Compares two tags for strict equality.
//
// Added in macOS 14.0.
// Compares two tags for strict equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagEqualToTag
func CMTagEqualToTag(tag1 Tag, tag2 Tag) unsafe.Pointer {
	return _CMTagEqualToTag(tag1, tag2)
}

// Creates a new tagged buffer group from a pair of buffers and the tags to associate with them.
//
// Added in macOS 14.0.
// Creates a new tagged buffer group from a pair of buffers and the tags to associate with them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreate
func CMTaggedBufferGroupCreate(allocator AllocatorRef, tagCollections ArrayRef, buffers ArrayRef, groupOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupCreate(allocator, tagCollections, buffers, groupOut)
}

// Creates a new tagged buffer group from an array of existing tagged buffer groups.
//
// Added in macOS 14.0.
// Creates a new tagged buffer group from an array of existing tagged buffer groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreateCombined
func CMTaggedBufferGroupCreateCombined(allocator AllocatorRef, taggedBufferGroups ArrayRef, groupOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupCreateCombined(allocator, taggedBufferGroups, groupOut)
}

// Creates a new format description for a tagged buffer group.
//
// Added in macOS 14.0.
// Creates a new format description for a tagged buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator AllocatorRef, taggedBufferGroup TaggedBufferGroupRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator, taggedBufferGroup, formatDescriptionOut)
}

// Checks to see if a tagged buffer group’s format matches an existing format description.
//
// Added in macOS 14.0.
// Checks to see if a tagged buffer group’s format matches an existing format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc TaggedBufferGroupFormatDescriptionRef, taggedBufferGroup TaggedBufferGroupRef) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc, taggedBufferGroup)
}

// Gets the sample buffer at a given index in the buffer group.
//
// Added in macOS 14.0.
// Gets the sample buffer at a given index in the buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group TaggedBufferGroupRef, index Index) SampleBufferRef {
	return _CMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
}

// Gets the single sample buffer in a group which contains a given tag, if present.
//
// Added in macOS 14.0.
// Gets the single sample buffer in a group which contains a given tag, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTag
func CMTaggedBufferGroupGetCMSampleBufferForTag(group TaggedBufferGroupRef, tag Tag, indexOut unsafe.Pointer) SampleBufferRef {
	return _CMTaggedBufferGroupGetCMSampleBufferForTag(group, tag, indexOut)
}

// Gets the single sample buffer in a group which contains a given tag collection, if present.
//
// Added in macOS 14.0.
// Gets the single sample buffer in a group which contains a given tag collection, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTagCollection
func CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group TaggedBufferGroupRef, tagCollection TagCollectionRef, indexOut unsafe.Pointer) SampleBufferRef {
	return _CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group, tagCollection, indexOut)
}

// Gets the number of buffers contained within a tagged buffer group.
//
// Added in macOS 14.0.
// Gets the number of buffers contained within a tagged buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCount
func CMTaggedBufferGroupGetCount(group TaggedBufferGroupRef) ItemCount {
	return _CMTaggedBufferGroupGetCount(group)
}

// Gets the pixel buffer at a given index in the buffer group.
//
// Added in macOS 14.0.
// Gets the pixel buffer at a given index in the buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferAtIndex
func CMTaggedBufferGroupGetCVPixelBufferAtIndex(group TaggedBufferGroupRef, index Index) PixelBufferRef {
	return _CMTaggedBufferGroupGetCVPixelBufferAtIndex(group, index)
}

// Gets the single pixel buffer in a group which contains a given tag, if present.
//
// Added in macOS 14.0.
// Gets the single pixel buffer in a group which contains a given tag, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTag
func CMTaggedBufferGroupGetCVPixelBufferForTag(group TaggedBufferGroupRef, tag Tag, indexOut unsafe.Pointer) PixelBufferRef {
	return _CMTaggedBufferGroupGetCVPixelBufferForTag(group, tag, indexOut)
}

// Gets the single pixel buffer in a group which contains a given tag collection, if present.
//
// Added in macOS 14.0.
// Gets the single pixel buffer in a group which contains a given tag collection, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTagCollection
func CMTaggedBufferGroupGetCVPixelBufferForTagCollection(group TaggedBufferGroupRef, tagCollection TagCollectionRef, indexOut unsafe.Pointer) PixelBufferRef {
	return _CMTaggedBufferGroupGetCVPixelBufferForTagCollection(group, tagCollection, indexOut)
}

// Gets the number of buffers in the group associated with a given tag collection.
//
// Added in macOS 14.0.
// Gets the number of buffers in the group associated with a given tag collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetNumberOfMatchesForTagCollection
func CMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group TaggedBufferGroupRef, tagCollection TagCollectionRef) ItemCount {
	return _CMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group, tagCollection)
}

// Gets the collection of tags for a buffer at a given index in the group.
//
// Added in macOS 14.0.
// Gets the collection of tags for a buffer at a given index in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTagCollectionAtIndex
func CMTaggedBufferGroupGetTagCollectionAtIndex(group TaggedBufferGroupRef, index Index) TagCollectionRef {
	return _CMTaggedBufferGroupGetTagCollectionAtIndex(group, index)
}

// Gets the internal type ID for a tagged buffer group.
//
// Added in macOS 14.0.
// Gets the internal type ID for a tagged buffer group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTypeID
func CMTaggedBufferGroupGetTypeID() TypeID {
	return _CMTaggedBufferGroupGetTypeID()
}

// Retrieves a tag’s value as a 64-bit field flag.
//
// Added in macOS 14.0.
// Retrieves a tag’s value as a 64-bit field flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetFlagsValue
func CMTagGetFlagsValue(tag Tag) uint64 {
	return _CMTagGetFlagsValue(tag)
}

// Retrieves a tag’s value as a 64-bit floating point number.
//
// Added in macOS 14.0.
// Retrieves a tag’s value as a 64-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetFloat64Value
func CMTagGetFloat64Value(tag Tag) unsafe.Pointer {
	return _CMTagGetFloat64Value(tag)
}

// Retrieves a tag’s value for use by the operating system.
//
// Added in macOS 14.0.
// Retrieves a tag’s value for use by the operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetOSTypeValue
func CMTagGetOSTypeValue(tag Tag) unsafe.Pointer {
	return _CMTagGetOSTypeValue(tag)
}

// Retrieves a tag’s value as a signed 64-bit integer.
//
// Added in macOS 14.0.
// Retrieves a tag’s value as a signed 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetSInt64Value
func CMTagGetSInt64Value(tag Tag) int64 {
	return _CMTagGetSInt64Value(tag)
}

// Retrieves the data type of a tag.
//
// Added in macOS 14.0.
// Retrieves the data type of a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagGetValueDataType
func CMTagGetValueDataType(tag Tag) TagDataType {
	return _CMTagGetValueDataType(tag)
}

// Whether a given tag contains a value for a 64-bit flag field.
//
// Added in macOS 14.0.
// Whether a given tag contains a value for a 64-bit flag field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagHasFlagsValue
func CMTagHasFlagsValue(tag Tag) unsafe.Pointer {
	return _CMTagHasFlagsValue(tag)
}

// Whether a given tag contains a value for a 64-bit floating point number.
//
// Added in macOS 14.0.
// Whether a given tag contains a value for a 64-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagHasFloat64Value
func CMTagHasFloat64Value(tag Tag) unsafe.Pointer {
	return _CMTagHasFloat64Value(tag)
}

// Generates a hash identifier for a tag.
//
// Added in macOS 14.0.
// Generates a hash identifier for a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagHash
func CMTagHash(tag Tag) HashCode {
	return _CMTagHash(tag)
}

// Whether a given tag contains a value for use by the operating system.
//
// Added in macOS 14.0.
// Whether a given tag contains a value for use by the operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagHasOSTypeValue
func CMTagHasOSTypeValue(tag Tag) unsafe.Pointer {
	return _CMTagHasOSTypeValue(tag)
}

// Whether a given tag contains a value for a signed 64-bit integer.
//
// Added in macOS 14.0.
// Whether a given tag contains a value for a signed 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagHasSInt64Value
func CMTagHasSInt64Value(tag Tag) unsafe.Pointer {
	return _CMTagHasSInt64Value(tag)
}

// Create a new tag from a dictionary object.
//
// Added in macOS 14.0.
// Create a new tag from a dictionary object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeFromDictionary
func CMTagMakeFromDictionary(dict DictionaryRef) Tag {
	return _CMTagMakeFromDictionary(dict)
}

// Creates a new tag with a given category and a value interpreted as a 64-bit flag field.
//
// Added in macOS 14.0.
// Creates a new tag with a given category and a value interpreted as a 64-bit flag field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFlagsValue
func CMTagMakeWithFlagsValue(category TagCategory, flagsForTag uint64) Tag {
	return _CMTagMakeWithFlagsValue(category, flagsForTag)
}

// Creates a new tag with a given category and a 64-bit floating point value.
//
// Added in macOS 14.0.
// Creates a new tag with a given category and a 64-bit floating point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFloat64Value
func CMTagMakeWithFloat64Value(category TagCategory, value unsafe.Pointer) Tag {
	return _CMTagMakeWithFloat64Value(category, value)
}

// Creates a new tag with a given category and a 64-bit value for use by the framework.
//
// Added in macOS 14.0.
// Creates a new tag with a given category and a 64-bit value for use by the framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithOSTypeValue
func CMTagMakeWithOSTypeValue(category TagCategory, value unsafe.Pointer) Tag {
	return _CMTagMakeWithOSTypeValue(category, value)
}

// Creates a new tag with a given category and a 64-bit signed integer.
//
// Added in macOS 14.0.
// Creates a new tag with a given category and a 64-bit signed integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithSInt64Value
func CMTagMakeWithSInt64Value(category TagCategory, value int64) Tag {
	return _CMTagMakeWithSInt64Value(category, value)
}

// Copies the contents of a text format description to a buffer in big-endian byte order.
//
// Added in macOS 10.10.
// Copies the contents of a text format description to a buffer in big-endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator:textFormatDescription:flavor:blockBufferOut:)
func CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator AllocatorRef, textFormatDescription TextFormatDescriptionRef, flavor TextDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator, textFormatDescription, flavor, blockBufferOut)
}

// Creates a text format description from a big-endian text description structure inside a buffer.
//
// Added in macOS 10.10.
// Creates a text format description from a big-endian text description structure inside a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator:bigEndianTextDescriptionBlockBuffer:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator AllocatorRef, textDescriptionBlockBuffer BlockBufferRef, flavor TextDescriptionFlavor, mediaType MediaType, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator, textDescriptionBlockBuffer, flavor, mediaType, formatDescriptionOut)
}

// Creates a text format description from a big-endian text description structure.
//
// Added in macOS 10.10.
// Creates a text format description from a big-endian text description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator:bigEndianTextDescriptionData:size:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator AllocatorRef, textDescriptionData unsafe.Pointer, size uintptr, flavor TextDescriptionFlavor, mediaType MediaType, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator, textDescriptionData, size, flavor, mediaType, formatDescriptionOut)
}

// Returns the default text style.
//
// Added in macOS 10.7.
// Returns the default text style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultStyle(_:localFontIDOut:boldOut:italicOut:underlineOut:fontSizeOut:colorComponentsOut:)
func CMTextFormatDescriptionGetDefaultStyle(desc FormatDescriptionRef, localFontIDOut unsafe.Pointer, boldOut unsafe.Pointer, italicOut unsafe.Pointer, underlineOut unsafe.Pointer, fontSizeOut []float64, colorComponentsOut float64, p7 unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionGetDefaultStyle(desc, localFontIDOut, boldOut, italicOut, underlineOut, fontSizeOut, colorComponentsOut, p7)
}

// Returns the default text box.
//
// Added in macOS 10.7.
// Returns the default text box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultTextBox(_:originIsAtTopLeft:heightOfTextTrack:defaultTextBoxOut:)
func CMTextFormatDescriptionGetDefaultTextBox(desc FormatDescriptionRef, originIsAtTopLeft unsafe.Pointer, heightOfTextTrack float64, defaultTextBoxOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionGetDefaultTextBox(desc, originIsAtTopLeft, heightOfTextTrack, defaultTextBoxOut)
}

// Returns the display flags.
//
// Added in macOS 10.7.
// Returns the display flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDisplayFlags(_:displayFlagsOut:)
func CMTextFormatDescriptionGetDisplayFlags(desc FormatDescriptionRef, displayFlagsOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionGetDisplayFlags(desc, displayFlagsOut)
}

// Returns a font name for a local font identifier.
//
// Added in macOS 10.7.
// Returns a font name for a local font identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetFontName(_:localFontID:fontNameOut:)
func CMTextFormatDescriptionGetFontName(desc FormatDescriptionRef, localFontID uint16, fontNameOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionGetFontName(desc, localFontID, fontNameOut)
}

// Returns the horizontal and vertical justification.
//
// Added in macOS 10.7.
// Returns the horizontal and vertical justification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetJustification(_:horizontalOut:verticalOut:)
func CMTextFormatDescriptionGetJustification(desc FormatDescriptionRef, horizontaJustificationlOut unsafe.Pointer, verticalJustificationOut unsafe.Pointer) unsafe.Pointer {
	return _CMTextFormatDescriptionGetJustification(desc, horizontaJustificationlOut, verticalJustificationOut)
}

// Returns the absolute value of a time.
//
// Added in macOS 10.7.
// Returns the absolute value of a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAbsoluteValue(_:)
func CMTimeAbsoluteValue(time Time) Time {
	return _CMTimeAbsoluteValue(time)
}

// Returns the sum of two times.
//
// Added in macOS 10.7.
// Returns the sum of two times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeAdd(_:_:)
func CMTimeAdd(lhs Time, rhs Time) Time {
	return _CMTimeAdd(lhs, rhs)
}

// Adds the timer to the list of timers the timebase manages.
//
// Added in macOS 10.8.
// Adds the timer to the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimer(_:timer:runloop:)
func CMTimebaseAddTimer(timebase TimebaseRef, timer RunLoopTimerRef, runloop RunLoopRef) unsafe.Pointer {
	return _CMTimebaseAddTimer(timebase, timer, runloop)
}

// Adds the timer dispatch source to the list of timers the timebase manages.
//
// Added in macOS 10.8.
// Adds the timer dispatch source to the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimerDispatchSource(_:timerSource:)
func CMTimebaseAddTimerDispatchSource(timebase TimebaseRef, timerSource unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseAddTimerDispatchSource(timebase, timerSource)
}

// Returns the immediate host clock of a timebase.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.11.
// Returns the immediate host clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterClock(_:)
func CMTimebaseCopyMasterClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseCopyMasterClock(timebase)
}

// Returns the immediate host timebase of a timebase.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.11.
// Returns the immediate host timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyMasterTimebase(_:)
func CMTimebaseCopyMasterTimebase(timebase TimebaseRef) TimebaseRef {
	return _CMTimebaseCopyMasterTimebase(timebase)
}

// Returns the immediate source — either a clock or timebase — of a timebase.
//
// Added in macOS 10.11.
// Returns the immediate source — either a clock or timebase — of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySource(_:)
func CMTimebaseCopySource(timebase TimebaseRef) ClockOrTimebaseRef {
	return _CMTimebaseCopySource(timebase)
}

// Returns the immediate source clock of a timebase.
//
// Added in macOS 10.11.
// Returns the immediate source clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceClock(_:)
func CMTimebaseCopySourceClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseCopySourceClock(timebase)
}

// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.11.
// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateMasterClock(_:)
func CMTimebaseCopyUltimateMasterClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseCopyUltimateMasterClock(timebase)
}

// Returns the source clock that’s the source of all of a timebase’s source timebases.
//
// Added in macOS 10.11.
// Returns the source clock that’s the source of all of a timebase’s source timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateSourceClock(_:)
func CMTimebaseCopyUltimateSourceClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseCopyUltimateSourceClock(timebase)
}

// Creates a timebase by using a primary clock.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.8.
// Creates a timebase by using a primary clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithMasterClock(allocator:masterClock:timebaseOut:)
func CMTimebaseCreateWithMasterClock(allocator AllocatorRef, masterClock ClockRef, timebaseOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCreateWithMasterClock(allocator, masterClock, timebaseOut)
}

// Creates a timebase by using a host timebase.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.8.
// Creates a timebase by using a host timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithMasterTimebase(allocator:masterTimebase:timebaseOut:)
func CMTimebaseCreateWithMasterTimebase(allocator AllocatorRef, masterTimebase TimebaseRef, timebaseOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCreateWithMasterTimebase(allocator, masterTimebase, timebaseOut)
}

// Creates a timebase by using a source clock.
//
// Added in macOS 10.8.
// Creates a timebase by using a source clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)
func CMTimebaseCreateWithSourceClock(allocator AllocatorRef, sourceClock ClockRef, timebaseOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseCreateWithSourceClock(allocator, sourceClock, timebaseOut)
}

// Returns the effective rate of a timebase, which combines its rate with the rates of all its host timebases.
//
// Added in macOS 10.8.
// Returns the effective rate of a timebase, which combines its rate with the rates of all its host timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetEffectiveRate(_:)
func CMTimebaseGetEffectiveRate(timebase TimebaseRef) unsafe.Pointer {
	return _CMTimebaseGetEffectiveRate(timebase)
}

// Returns the immediate host (either timebase or clock) of a timebase.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.8.
// Returns the immediate host (either timebase or clock) of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMaster(_:)
func CMTimebaseGetMaster(timebase TimebaseRef) ClockOrTimebaseRef {
	return _CMTimebaseGetMaster(timebase)
}

// Returns the immediate host clock of a timebase.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.8.
// Returns the immediate host clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterClock(_:)
func CMTimebaseGetMasterClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseGetMasterClock(timebase)
}

// Returns the immediate host timebase of a timebase.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.8.
// Returns the immediate host timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetMasterTimebase(_:)
func CMTimebaseGetMasterTimebase(timebase TimebaseRef) TimebaseRef {
	return _CMTimebaseGetMasterTimebase(timebase)
}

// Returns the current rate of a timebase.
//
// Added in macOS 10.8.
// Returns the current rate of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetRate(_:)
func CMTimebaseGetRate(timebase TimebaseRef) unsafe.Pointer {
	return _CMTimebaseGetRate(timebase)
}

// Returns the current time from a timebase.
//
// Added in macOS 10.8.
// Returns the current time from a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTime(_:)
func CMTimebaseGetTime(timebase TimebaseRef) Time {
	return _CMTimebaseGetTime(timebase)
}

// Returns the current time and rate of a timebase.
//
// Added in macOS 10.8.
// Returns the current time and rate of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase TimebaseRef, timeOut unsafe.Pointer, rateOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
}

// Returns the current time from a timebase in the specified timescale.
//
// Added in macOS 10.8.
// Returns the current time from a timebase in the specified timescale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeWithTimeScale(_:timescale:method:)
func CMTimebaseGetTimeWithTimeScale(timebase TimebaseRef, timescale TimeScale, method TimeRoundingMethod) Time {
	return _CMTimebaseGetTimeWithTimeScale(timebase, timescale, method)
}

// Returns the Core Foundation type identifier that identifies a timebase object.
//
// Added in macOS 10.8.
// Returns the Core Foundation type identifier that identifies a timebase object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() TypeID {
	return _CMTimebaseGetTypeID()
}

// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.8.
// Returns the host clock that is the host of all of a timebase’s host timebases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetUltimateMasterClock(_:)
func CMTimebaseGetUltimateMasterClock(timebase TimebaseRef) ClockRef {
	return _CMTimebaseGetUltimateMasterClock(timebase)
}

// Requests that the timebase wait until it isn’t posting notifications.
//
// Added in macOS 10.8.
// Requests that the timebase wait until it isn’t posting notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseNotificationBarrier(_:)
func CMTimebaseNotificationBarrier(timebase TimebaseRef) unsafe.Pointer {
	return _CMTimebaseNotificationBarrier(timebase)
}

// Removes the timer from the list of timers the timebase manages.
//
// Added in macOS 10.8.
// Removes the timer from the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimer(_:timer:)
func CMTimebaseRemoveTimer(timebase TimebaseRef, timer RunLoopTimerRef) unsafe.Pointer {
	return _CMTimebaseRemoveTimer(timebase, timer)
}

// Removes the timer dispatch source from the list of timers the timebase manages.
//
// Added in macOS 10.8.
// Removes the timer dispatch source from the list of timers the timebase manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimerDispatchSource(_:timerSource:)
func CMTimebaseRemoveTimerDispatchSource(timebase TimebaseRef, timerSource unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseRemoveTimerDispatchSource(timebase, timerSource)
}

// Sets the time of a timebase at a particular host time.
//
// Added in macOS 10.8.
// Sets the time of a timebase at a particular host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetAnchorTime(_:timebaseTime:immediateSourceTime:)
func CMTimebaseSetAnchorTime(timebase TimebaseRef, timebaseTime Time, immediateSourceTime Time) unsafe.Pointer {
	return _CMTimebaseSetAnchorTime(timebase, timebaseTime, immediateSourceTime)
}

// Sets the rate of a timebase.
//
// Added in macOS 10.8.
// Sets the rate of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRate(_:rate:)
func CMTimebaseSetRate(timebase TimebaseRef, rate unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetRate(timebase, rate)
}

// Sets the time of a timebase at a particular host time, and changes the rate at exactly that time.
//
// Added in macOS 10.8.
// Sets the time of a timebase at a particular host time, and changes the rate at exactly that time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRateAndAnchorTime(_:rate:anchorTime:immediateSourceTime:)
func CMTimebaseSetRateAndAnchorTime(timebase TimebaseRef, rate unsafe.Pointer, timebaseTime Time, immediateSourceTime Time) unsafe.Pointer {
	return _CMTimebaseSetRateAndAnchorTime(timebase, rate, timebaseTime, immediateSourceTime)
}

// Sets the source clock of a timebase.
//
// Added in macOS 10.8.
// Sets the source clock of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceClock(_:_:)
func CMTimebaseSetSourceClock(timebase TimebaseRef, newSourceClock ClockRef) unsafe.Pointer {
	return _CMTimebaseSetSourceClock(timebase, newSourceClock)
}

// Sets the source timebase of a timebase.
//
// Added in macOS 10.8.
// Sets the source timebase of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceTimebase(_:_:)
func CMTimebaseSetSourceTimebase(timebase TimebaseRef, newSourceTimebase TimebaseRef) unsafe.Pointer {
	return _CMTimebaseSetSourceTimebase(timebase, newSourceTimebase)
}

// Sets the current time of a timebase.
//
// Added in macOS 10.8.
// Sets the current time of a timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTime(_:time:)
func CMTimebaseSetTime(timebase TimebaseRef, time Time) unsafe.Pointer {
	return _CMTimebaseSetTime(timebase, time)
}

// Sets the time on the timebase’s timeline at which the timer dispatch source should fire next.
//
// Added in macOS 10.8.
// Sets the time on the timebase’s timeline at which the timer dispatch source should fire next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceNextFireTime(_:timerSource:fireTime:flags:)
func CMTimebaseSetTimerDispatchSourceNextFireTime(timebase TimebaseRef, timerSource unsafe.Pointer, fireTime Time, flags uint32) unsafe.Pointer {
	return _CMTimebaseSetTimerDispatchSourceNextFireTime(timebase, timerSource, fireTime, flags)
}

// Sets the timer dispatch source to fire immediately once, overriding any previous timer call.
//
// Added in macOS 10.8.
// Sets the timer dispatch source to fire immediately once, overriding any previous timer call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceToFireImmediately(_:timerSource:)
func CMTimebaseSetTimerDispatchSourceToFireImmediately(timebase TimebaseRef, timerSource unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseSetTimerDispatchSourceToFireImmediately(timebase, timerSource)
}

// Sets the time on the timebase’s timeline at which the timer should fire next.
//
// Added in macOS 10.8.
// Sets the time on the timebase’s timeline at which the timer should fire next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerNextFireTime(_:timer:fireTime:flags:)
func CMTimebaseSetTimerNextFireTime(timebase TimebaseRef, timer RunLoopTimerRef, fireTime Time, flags uint32) unsafe.Pointer {
	return _CMTimebaseSetTimerNextFireTime(timebase, timer, fireTime, flags)
}

// Sets the timer to fire immediately once, overriding any previous timer calls.
//
// Added in macOS 10.8.
// Sets the timer to fire immediately once, overriding any previous timer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerToFireImmediately(_:timer:)
func CMTimebaseSetTimerToFireImmediately(timebase TimebaseRef, timer RunLoopTimerRef) unsafe.Pointer {
	return _CMTimebaseSetTimerToFireImmediately(timebase, timer)
}

// Returns the nearest time value inside the time range.
//
// Added in macOS 10.7.
// Returns the nearest time value inside the time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeClampToRange(_:range:)
func CMTimeClampToRange(time Time, range_ TimeRange) Time {
	return _CMTimeClampToRange(time, range_)
}

// Copies the contents of a time code format description to a buffer in big-endian byte order.
//
// Added in macOS 10.10.
// Copies the contents of a time code format description to a buffer in big-endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator:timeCodeFormatDescription:flavor:blockBufferOut:)
func CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator AllocatorRef, timeCodeFormatDescription TimeCodeFormatDescriptionRef, flavor TimeCodeDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeFormatDescription, flavor, blockBufferOut)
}

// Creates a format description for time code media.
//
// Added in macOS 10.7.
// Creates a format description for time code media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreate(allocator:timeCodeFormatType:frameDuration:frameQuanta:flags:extensions:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreate(allocator AllocatorRef, timeCodeFormatType TimeCodeFormatType, frameDuration Time, frameQuanta uint32, flags uint32, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionCreate(allocator, timeCodeFormatType, frameDuration, frameQuanta, flags, extensions, formatDescriptionOut)
}

// Creates a time code format description from a big-endian time code description data structure in a buffer.
//
// Added in macOS 10.10.
// Creates a time code format description from a big-endian time code description data structure in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator:bigEndianTimeCodeDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator AllocatorRef, timeCodeDescriptionBlockBuffer BlockBufferRef, flavor TimeCodeDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

// Creates a time code format description from a big-endian time code description structure.
//
// Added in macOS 10.10.
// Creates a time code format description from a big-endian time code description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator:bigEndianTimeCodeDescriptionData:size:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator AllocatorRef, timeCodeDescriptionData unsafe.Pointer, size uintptr, flavor TimeCodeDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator, timeCodeDescriptionData, size, flavor, formatDescriptionOut)
}

// Returns the duration of each frame.
//
// Added in macOS 10.7.
// Returns the duration of each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameDuration(_:)
func CMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription TimeCodeFormatDescriptionRef) Time {
	return _CMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription)
}

// Returns the frames per second for a time code, or frames per tick in counter mode.
//
// Added in macOS 10.7.
// Returns the frames per second for a time code, or frames per tick in counter mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameQuanta(_:)
func CMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription TimeCodeFormatDescriptionRef) uint32 {
	return _CMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription)
}

// Returns time code flags.
//
// Added in macOS 10.7.
// Returns time code flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetTimeCodeFlags(_:)
func CMTimeCodeFormatDescriptionGetTimeCodeFlags(desc TimeCodeFormatDescriptionRef) uint32 {
	return _CMTimeCodeFormatDescriptionGetTimeCodeFlags(desc)
}

// Returns the numerical relationship of two times.
//
// Added in macOS 10.7.
// Returns the numerical relationship of two times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCompare(_:_:)
func CMTimeCompare(time1 Time, time2 Time) int32 {
	return _CMTimeCompare(time1, time2)
}

// Converts the source time to a new timescale using the specified rounding method.
//
// Added in macOS 10.7.
// Converts the source time to a new timescale using the specified rounding method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeConvertScale(_:timescale:method:)
func CMTimeConvertScale(time Time, newTimescale int32, method TimeRoundingMethod) Time {
	return _CMTimeConvertScale(time, newTimescale, method)
}

// Creates a dictionary representation of the time.
//
// Added in macOS 10.7.
// Creates a dictionary representation of the time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyAsDictionary(_:allocator:)
func CMTimeCopyAsDictionary(time Time, allocator AllocatorRef) DictionaryRef {
	return _CMTimeCopyAsDictionary(time, allocator)
}

// Creates a string representation of the time.
//
// Added in macOS 10.7.
// Creates a string representation of the time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyDescription(allocator:time:)
func CMTimeCopyDescription(allocator AllocatorRef, time Time) StringRef {
	return _CMTimeCopyDescription(allocator, time)
}

// Folds a time into a time range.
//
// Added in macOS 10.14.
// Folds a time into a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFoldIntoRange(_:foldRange:)
func CMTimeFoldIntoRange(time Time, foldRange TimeRange) Time {
	return _CMTimeFoldIntoRange(time, foldRange)
}

// Returns a representation of the time in seconds.
//
// Added in macOS 10.7.
// Returns a representation of the time in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time Time) unsafe.Pointer {
	return _CMTimeGetSeconds(time)
}

// Creates a time with a value and timescale.
//
// Added in macOS 10.7.
// Creates a time with a value and timescale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMake(value:timescale:)
func CMTimeMake(value int64, timescale int32) Time {
	return _CMTimeMake(value, timescale)
}

// Creates a time from a dictionary representation of its fields.
//
// Added in macOS 10.7.
// Creates a time from a dictionary representation of its fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeFromDictionary(_:)
func CMTimeMakeFromDictionary(dictionaryRepresentation DictionaryRef) Time {
	return _CMTimeMakeFromDictionary(dictionaryRepresentation)
}

// Creates a time with a value, timescale, and epoch.
//
// Added in macOS 10.7.
// Creates a time with a value, timescale, and epoch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithEpoch(value:timescale:epoch:)
func CMTimeMakeWithEpoch(value int64, timescale int32, epoch int64) Time {
	return _CMTimeMakeWithEpoch(value, timescale, epoch)
}

// Creates a time that represents a number of seconds in a preferred timescale.
//
// Added in macOS 10.7.
// Creates a time that represents a number of seconds in a preferred timescale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithSeconds(_:preferredTimescale:)
func CMTimeMakeWithSeconds(seconds unsafe.Pointer, preferredTimescale int32) Time {
	return _CMTimeMakeWithSeconds(seconds, preferredTimescale)
}

// Translates a duration through a mapping from two time ranges.
//
// Added in macOS 10.7.
// Translates a duration through a mapping from two time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapDurationFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapDurationFromRangeToRange(dur Time, fromRange TimeRange, toRange TimeRange) Time {
	return _CMTimeMapDurationFromRangeToRange(dur, fromRange, toRange)
}

// Returns a dictionary representation of a time mapping.
//
// Added in macOS 10.11.
// Returns a dictionary representation of a time mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyAsDictionary(_:allocator:)
func CMTimeMappingCopyAsDictionary(mapping TimeMapping, allocator AllocatorRef) DictionaryRef {
	return _CMTimeMappingCopyAsDictionary(mapping, allocator)
}

// Copies a string description of a time mapping.
//
// Added in macOS 10.11.
// Copies a string description of a time mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator AllocatorRef, mapping TimeMapping) StringRef {
	return _CMTimeMappingCopyDescription(allocator, mapping)
}

// Creates a time mapping with a source and target time range.
//
// Added in macOS 10.11.
// Creates a time mapping with a source and target time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMake(source:target:)
func CMTimeMappingMake(source TimeRange, target TimeRange) TimeMapping {
	return _CMTimeMappingMake(source, target)
}

// Creates a valid time mapping with an empty source.
//
// Added in macOS 10.11.
// Creates a valid time mapping with an empty source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeEmpty(target:)
func CMTimeMappingMakeEmpty(target TimeRange) TimeMapping {
	return _CMTimeMappingMakeEmpty(target)
}

// Creates a time mapping from a dictionary representation.
//
// Added in macOS 10.11.
// Creates a time mapping from a dictionary representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeFromDictionary(_:)
func CMTimeMappingMakeFromDictionary(dictionaryRepresentation DictionaryRef) TimeMapping {
	return _CMTimeMappingMakeFromDictionary(dictionaryRepresentation)
}

// Prints a description of a time mapping to standard output.
//
// Added in macOS 10.11.
// Prints a description of a time mapping to standard output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingShow(_:)
func CMTimeMappingShow(mapping TimeMapping) {
	_CMTimeMappingShow(mapping)
}

// Translates a time through a mapping from two time ranges.
//
// Added in macOS 10.7.
// Translates a time through a mapping from two time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMapTimeFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapTimeFromRangeToRange(t Time, fromRange TimeRange, toRange TimeRange) Time {
	return _CMTimeMapTimeFromRangeToRange(t, fromRange, toRange)
}

// Returns the greater of two time values.
//
// Added in macOS 10.7.
// Returns the greater of two time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMaximum(_:_:)
func CMTimeMaximum(time1 Time, time2 Time) Time {
	return _CMTimeMaximum(time1, time2)
}

// Returns the lesser of two time values.
//
// Added in macOS 10.7.
// Returns the lesser of two time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMinimum(_:_:)
func CMTimeMinimum(time1 Time, time2 Time) Time {
	return _CMTimeMinimum(time1, time2)
}

// Returns the result of multiplying a time by an integer multiplier.
//
// Added in macOS 10.7.
// Returns the result of multiplying a time by an integer multiplier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiply(_:multiplier:)
func CMTimeMultiply(time Time, multiplier int32) Time {
	return _CMTimeMultiply(time, multiplier)
}

// Returns the result of multiplying a time by a floating-point multiplier.
//
// Added in macOS 10.7.
// Returns the result of multiplying a time by a floating-point multiplier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByFloat64(_:multiplier:)
func CMTimeMultiplyByFloat64(time Time, multiplier unsafe.Pointer) Time {
	return _CMTimeMultiplyByFloat64(time, multiplier)
}

// Returns the result of multiplying a time by an integer multiplier, and then dividing the result by the divisor.
//
// Added in macOS 10.10.
// Returns the result of multiplying a time by an integer multiplier, and then dividing the result by the divisor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByRatio(_:multiplier:divisor:)
func CMTimeMultiplyByRatio(time Time, multiplier int32, divisor int32) Time {
	return _CMTimeMultiplyByRatio(time, multiplier, divisor)
}

// Returns a Boolean value that indicates whether a time range contains a time.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a time range contains a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTime(_:time:)
func CMTimeRangeContainsTime(range_ TimeRange, time Time) unsafe.Pointer {
	return _CMTimeRangeContainsTime(range_, time)
}

// Returns a Boolean value that indicates whether a time range contains another time range.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a time range contains another time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTimeRange(_:otherRange:)
func CMTimeRangeContainsTimeRange(range_ TimeRange, otherRange TimeRange) unsafe.Pointer {
	return _CMTimeRangeContainsTimeRange(range_, otherRange)
}

// Returns a dictionary representation of a time range.
//
// Added in macOS 10.7.
// Returns a dictionary representation of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ TimeRange, allocator AllocatorRef) DictionaryRef {
	return _CMTimeRangeCopyAsDictionary(range_, allocator)
}

// Returns a string with a description of a time range.
//
// Added in macOS 10.7.
// Returns a string with a description of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyDescription(allocator:range:)
func CMTimeRangeCopyDescription(allocator AllocatorRef, range_ TimeRange) StringRef {
	return _CMTimeRangeCopyDescription(allocator, range_)
}

// Returns a Boolean value that indicates whether two time ranges are equal.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether two time ranges are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeEqual(_:_:)
func CMTimeRangeEqual(range1 TimeRange, range2 TimeRange) unsafe.Pointer {
	return _CMTimeRangeEqual(range1, range2)
}

// Creates a valid time range from a start and end time.
//
// Added in macOS 10.7.
// Creates a valid time range from a start and end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeFromTimeToTime(start:end:)
func CMTimeRangeFromTimeToTime(start Time, end Time) TimeRange {
	return _CMTimeRangeFromTimeToTime(start, end)
}

// Returns a time value that represents the end of a time range.
//
// Added in macOS 10.7.
// Returns a time value that represents the end of a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetEnd(_:)
func CMTimeRangeGetEnd(range_ TimeRange) Time {
	return _CMTimeRangeGetEnd(range_)
}

// Returns a new time range with the time elements that are common between the input.
//
// Added in macOS 10.7.
// Returns a new time range with the time elements that are common between the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetIntersection(_:otherRange:)
func CMTimeRangeGetIntersection(range_ TimeRange, otherRange TimeRange) TimeRange {
	return _CMTimeRangeGetIntersection(range_, otherRange)
}

// Returns a new time range with the time elements of the input.
//
// Added in macOS 10.7.
// Returns a new time range with the time elements of the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetUnion(_:otherRange:)
func CMTimeRangeGetUnion(range_ TimeRange, otherRange TimeRange) TimeRange {
	return _CMTimeRangeGetUnion(range_, otherRange)
}

// Creates a valid time range with a start time and duration.
//
// Added in macOS 10.7.
// Creates a valid time range with a start time and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMake(start:duration:)
func CMTimeRangeMake(start Time, duration Time) TimeRange {
	return _CMTimeRangeMake(start, duration)
}

// Creates a time range from a dictionary representation of its fields.
//
// Added in macOS 10.7.
// Creates a time range from a dictionary representation of its fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMakeFromDictionary(_:)
func CMTimeRangeMakeFromDictionary(dictionaryRepresentation DictionaryRef) TimeRange {
	return _CMTimeRangeMakeFromDictionary(dictionaryRepresentation)
}

// Prints a description of the time range to standard error.
//
// Added in macOS 10.7.
// Prints a description of the time range to standard error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeShow(_:)
func CMTimeRangeShow(range_ TimeRange) {
	_CMTimeRangeShow(range_)
}

// Prints a description of the time to the console.
//
// Added in macOS 10.7.
// Prints a description of the time to the console.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeShow(_:)
func CMTimeShow(time Time) {
	_CMTimeShow(time)
}

// Returns the difference between two times.
//
// Added in macOS 10.7.
// Returns the difference between two times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeSubtract(_:_:)
func CMTimeSubtract(lhs Time, rhs Time) Time {
	return _CMTimeSubtract(lhs, rhs)
}

// Copies the contents of a video format description to a buffer in big-endian byte ordering.
//
// Added in macOS 10.10.
// Copies the contents of a video format description to a buffer in big-endian byte ordering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator:videoFormatDescription:stringEncoding:flavor:blockBufferOut:)
func CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator AllocatorRef, videoFormatDescription VideoFormatDescriptionRef, stringEncoding StringEncoding, flavor ImageDescriptionFlavor, blockBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator, videoFormatDescription, stringEncoding, flavor, blockBufferOut)
}

// Copies the multi-image encoding properties as an array of CMTagCollections.
//
// Added in macOS 14.0.
// Copies the multi-image encoding properties as an array of CMTagCollections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyTagCollectionArray
func CMVideoFormatDescriptionCopyTagCollectionArray(formatDescription VideoFormatDescriptionRef, tagCollectionsOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCopyTagCollectionArray(formatDescription, tagCollectionsOut)
}

// Creates a format description for a video media stream.
//
// Added in macOS 10.7.
// Creates a format description for a video media stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreate(allocator:codecType:width:height:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreate(allocator AllocatorRef, codecType VideoCodecType, width int32, height int32, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreate(allocator, codecType, width, height, extensions, formatDescriptionOut)
}

// Creates a format description for a video media stream by using an image buffer.
//
// Added in macOS 10.7.
// Creates a format description for a video media stream by using an image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateForImageBuffer(allocator:imageBuffer:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateForImageBuffer(allocator AllocatorRef, imageBuffer ImageBufferRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreateForImageBuffer(allocator, imageBuffer, formatDescriptionOut)
}

// Creates a video format description from a big-endian image description inside a buffer.
//
// Added in macOS 10.10.
// Creates a video format description from a big-endian image description inside a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator:bigEndianImageDescriptionBlockBuffer:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator AllocatorRef, imageDescriptionBlockBuffer BlockBufferRef, stringEncoding StringEncoding, flavor ImageDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator, imageDescriptionBlockBuffer, stringEncoding, flavor, formatDescriptionOut)
}

// Creates a video format description from a big-endian image description structure.
//
// Added in macOS 10.10.
// Creates a video format description from a big-endian image description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator:bigEndianImageDescriptionData:size:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator AllocatorRef, imageDescriptionData unsafe.Pointer, size uintptr, stringEncoding StringEncoding, flavor ImageDescriptionFlavor, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator, imageDescriptionData, size, stringEncoding, flavor, formatDescriptionOut)
}

// Creates a format description for a video media stream that the parameter set describes.
//
// Added in macOS 10.9.
// Creates a format description for a video media stream that the parameter set describes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator AllocatorRef, parameterSetCount uintptr, parameterSetPointers unsafe.Pointer, parameterSetSizes unsafe.Pointer, NALUnitHeaderLength int, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, formatDescriptionOut)
}

// Creates a format description for a video media stream using HEVC (H.265) parameter set NAL units.
//
// Added in macOS 10.13.
// Creates a format description for a video media stream using HEVC (H.265) parameter set NAL units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator AllocatorRef, parameterSetCount uintptr, parameterSetPointers unsafe.Pointer, parameterSetSizes unsafe.Pointer, NALUnitHeaderLength int, extensions DictionaryRef, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, extensions, formatDescriptionOut)
}

// Returns a rectangle that defines the portion of the encoded pixel dimensions that represent the image data that’s valid for displaying.
//
// Added in macOS 10.7.
// Returns a rectangle that defines the portion of the encoded pixel dimensions that represent the image data that’s valid for displaying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetCleanAperture(_:originIsAtTopLeft:)
func CMVideoFormatDescriptionGetCleanAperture(videoDesc VideoFormatDescriptionRef, originIsAtTopLeft unsafe.Pointer) corefoundation.CGRect {
	return _CMVideoFormatDescriptionGetCleanAperture(videoDesc, originIsAtTopLeft)
}

// Returns the video dimensions, in encoded pixels.
//
// Added in macOS 10.7.
// Returns the video dimensions, in encoded pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetDimensions(_:)
func CMVideoFormatDescriptionGetDimensions(videoDesc VideoFormatDescriptionRef) VideoDimensions {
	return _CMVideoFormatDescriptionGetDimensions(videoDesc)
}

// Returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.
//
// Added in macOS 10.7.
// Returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
func CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() ArrayRef {
	return _CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
}

// Returns a parameter set that an H.264 format description contains.
//
// Added in macOS 10.9.
// Returns a parameter set that an H.264 format description contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetH264ParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc FormatDescriptionRef, parameterSetIndex uintptr, parameterSetPointerOut unsafe.Pointer, parameterSetSizeOut unsafe.Pointer, parameterSetCountOut unsafe.Pointer, NALUnitHeaderLengthOut []int) unsafe.Pointer {
	return _CMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, NALUnitHeaderLengthOut)
}

// Returns a parameter set contained in an HEVC (H.265) format description.
//
// Added in macOS 10.13.
// Returns a parameter set contained in an HEVC (H.265) format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc FormatDescriptionRef, parameterSetIndex uintptr, parameterSetPointerOut unsafe.Pointer, parameterSetSizeOut unsafe.Pointer, parameterSetCountOut unsafe.Pointer, NALUnitHeaderLengthOut []int) unsafe.Pointer {
	return _CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, NALUnitHeaderLengthOut)
}

// Returns the dimensions after taking the pixel aspect ratio and clean aperture into account.
//
// Added in macOS 10.7.
// Returns the dimensions after taking the pixel aspect ratio and clean aperture into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetPresentationDimensions(_:usePixelAspectRatio:useCleanAperture:)
func CMVideoFormatDescriptionGetPresentationDimensions(videoDesc VideoFormatDescriptionRef, usePixelAspectRatio unsafe.Pointer, useCleanAperture unsafe.Pointer) corefoundation.CGSize {
	return _CMVideoFormatDescriptionGetPresentationDimensions(videoDesc, usePixelAspectRatio, useCleanAperture)
}

// Returns a Boolean value that indicates whether a format description matches an image buffer.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether a format description matches an image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionMatchesImageBuffer(_:imageBuffer:)
func CMVideoFormatDescriptionMatchesImageBuffer(desc VideoFormatDescriptionRef, imageBuffer ImageBufferRef) unsafe.Pointer {
	return _CMVideoFormatDescriptionMatchesImageBuffer(desc, imageBuffer)
}




