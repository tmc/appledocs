// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

// Type aliases and typedefs
// AUAudioChannelCount - A number of audio channels.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioChannelCount
// AUAudioChannelCount has base type: uint32_t
type AUAudioChannelCount uintptr
// AUAudioFrameCount - A number of audio sample frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioFrameCount
// AUAudioFrameCount has base type: uint32_t
type AUAudioFrameCount uintptr
// AUAudioObjectID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioObjectID
// AUAudioObjectID has base type: UInt32
type AUAudioObjectID uintptr
// AUAudioUnitStatus - A result code returned from an audio unit’s render function.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitStatus
// AUAudioUnitStatus has base type: OSStatus
type AUAudioUnitStatus uintptr
// AUEventListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerProc
// AUEventListenerProc has base type: void (*)(void *, void *, const struct AudioUnitEvent *, unsigned long long, float)
type AUEventListenerProc uintptr
// AUEventListenerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRef
// AUEventListenerRef has base type: AUParameterListenerRef
type AUEventListenerRef uintptr
// AUEventSampleTime - Expresses time as a sample count.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventSampleTime
// AUEventSampleTime has base type: int64_t
type AUEventSampleTime uintptr
// AUInputSamplesInOutputCallback - Called by the system when an audio unit has provided a buffer of output samples.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallback
// AUInputSamplesInOutputCallback has base type: void (*)(void *, const struct AudioTimeStamp *, double, double) __attribute__((nonblocking))
type AUInputSamplesInOutputCallback uintptr
// AUMIDIOutputCallback - When called by a host application, gets MIDI data from an audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallback
// AUMIDIOutputCallback has base type: int (*)(void *, const struct AudioTimeStamp *, unsigned int, const struct MIDIPacketList *) __attribute__((nonblocking))
type AUMIDIOutputCallback uintptr
// AUNode - A member of an audio processing graph, associated with an audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNode
// AUNode has base type: SInt32
type AUNode uintptr
// AUNodeConnection type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeConnection
// AUNodeConnection has base type: struct AudioUnitNodeConnection
type AUNodeConnection uintptr
// AUParameterAddress - A numeric identifier for an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAddress
// AUParameterAddress has base type: uint64_t
type AUParameterAddress uintptr
// AUParameterListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerProc
// AUParameterListenerProc has base type: void (*)(void *, void *, const struct AudioUnitParameter *, float)
type AUParameterListenerProc uintptr
// AUParameterListenerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerRef
// AUParameterListenerRef has base type: struct AUListenerBase *
type AUParameterListenerRef uintptr
// AUParameterObserverToken - A token representing an installed parameter observer block.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterObserverToken
// AUParameterObserverToken has base type: void *
type AUParameterObserverToken uintptr
// AURenderCallback - Called by the system when an audio unit requires input samples, or before and after a render operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderCallback
// AURenderCallback has base type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type AURenderCallback uintptr
// AUValue - A value of an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUValue
// AUValue has base type: float
type AUValue uintptr
// AudioCodec - An instance of a Component Manager component.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodec
// AudioCodec has base type: AudioComponentInstance
type AudioCodec uintptr
// AudioCodecAppendInputBufferListProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferListProc
// AudioCodecAppendInputBufferListProc has base type: int (*)(void *, const struct AudioBufferList *, unsigned int *, const struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecAppendInputBufferListProc uintptr
// AudioCodecAppendInputDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputDataProc
// AudioCodecAppendInputDataProc has base type: int (*)(void *, const void *, unsigned int *, unsigned int *, const struct AudioStreamPacketDescription *)
type AudioCodecAppendInputDataProc uintptr
// AudioCodecGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfoProc
// AudioCodecGetPropertyInfoProc has base type: int (*)(void *, unsigned int, unsigned int *, unsigned char *)
type AudioCodecGetPropertyInfoProc uintptr
// AudioCodecGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyProc
// AudioCodecGetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int *, void *)
type AudioCodecGetPropertyProc uintptr
// AudioCodecInitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitializeProc
// AudioCodecInitializeProc has base type: int (*)(void *, const struct AudioStreamBasicDescription *, const struct AudioStreamBasicDescription *, const void *, unsigned int)
type AudioCodecInitializeProc uintptr
// AudioCodecProduceOutputBufferListProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferListProc
// AudioCodecProduceOutputBufferListProc has base type: int (*)(void *, struct AudioBufferList *, unsigned int *, struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecProduceOutputBufferListProc uintptr
// AudioCodecProduceOutputPacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPacketsProc
// AudioCodecProduceOutputPacketsProc has base type: int (*)(void *, void *, unsigned int *, unsigned int *, struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecProduceOutputPacketsProc uintptr
// AudioCodecPropertyID - An integer identifying an audio codec property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPropertyID
// AudioCodecPropertyID has base type: UInt32
type AudioCodecPropertyID uintptr
// AudioCodecResetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecResetProc
// AudioCodecResetProc has base type: int (*)(void *)
type AudioCodecResetProc uintptr
// AudioCodecSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetPropertyProc
// AudioCodecSetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int, const void *)
type AudioCodecSetPropertyProc uintptr
// AudioCodecUninitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitializeProc
// AudioCodecUninitializeProc has base type: int (*)(void *)
type AudioCodecUninitializeProc uintptr
// AudioComponent - An audio component.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponent
// AudioComponent has base type: struct OpaqueAudioComponent *
type AudioComponent uintptr
// AudioComponentFactoryFunction type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFactoryFunction
// AudioComponentFactoryFunction has base type: struct AudioComponentPlugInInterface *(*)(const struct AudioComponentDescription *)
type AudioComponentFactoryFunction uintptr
// AudioComponentInstance - A component instance, or object, is an audio unit or audio codec.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstance
// AudioComponentInstance has base type: struct OpaqueAudioComponentInstance *
type AudioComponentInstance uintptr
// AudioComponentMethod type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentMethod
// AudioComponentMethod has base type: int (*)(void *, ...)
type AudioComponentMethod uintptr
// AudioConverterComplexInputDataProc - Supplies input data to the   function.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProc
// AudioConverterComplexInputDataProc has base type: int (*)(struct OpaqueAudioConverter *, unsigned int *, struct AudioBufferList *, struct AudioStreamPacketDescription **, void *)
type AudioConverterComplexInputDataProc uintptr
// AudioConverterComplexInputDataProcRealtimeSafe type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProcRealtimeSafe
// AudioConverterComplexInputDataProcRealtimeSafe has base type: int (*)(struct OpaqueAudioConverter *, unsigned int *, struct AudioBufferList *, struct AudioStreamPacketDescription **, void *) __attribute__((nonblocking))
type AudioConverterComplexInputDataProcRealtimeSafe uintptr
// AudioConverterInputDataProc - Deprecated. Use   instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterInputDataProc
// AudioConverterInputDataProc has base type: int (*)(struct OpaqueAudioConverter *, unsigned int *, void **, void *)
type AudioConverterInputDataProc uintptr
// AudioConverterPropertyID - An audio converter property identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPropertyID
// AudioConverterPropertyID has base type: UInt32
type AudioConverterPropertyID uintptr
// AudioConverterRef - A reference to an audio converter object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterRef
// AudioConverterRef has base type: struct OpaqueAudioConverter *
type AudioConverterRef uintptr
// AudioFileComponent type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponent
// AudioFileComponent has base type: AudioComponentInstance
type AudioFileComponent uintptr
// AudioFileComponentCloseProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCloseProc
// AudioFileComponentCloseProc has base type: int (*)(void *)
type AudioFileComponentCloseProc uintptr
// AudioFileComponentCountUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserDataProc
// AudioFileComponentCountUserDataProc has base type: int (*)(void *, unsigned int, unsigned int *)
type AudioFileComponentCountUserDataProc uintptr
// AudioFileComponentCreateURLProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURLProc
// AudioFileComponentCreateURLProc has base type: int (*)(void *, const struct __CFURL *, const struct AudioStreamBasicDescription *, unsigned int)
type AudioFileComponentCreateURLProc uintptr
// AudioFileComponentExtensionIsThisFormatProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormatProc
// AudioFileComponentExtensionIsThisFormatProc has base type: int (*)(void *, const struct __CFString *, unsigned int *)
type AudioFileComponentExtensionIsThisFormatProc uintptr
// AudioFileComponentFileDataIsThisFormatProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormatProc
// AudioFileComponentFileDataIsThisFormatProc has base type: int (*)(void *, unsigned int, const void *, unsigned int *)
type AudioFileComponentFileDataIsThisFormatProc uintptr
// AudioFileComponentGetGlobalInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoProc
// AudioFileComponentGetGlobalInfoProc has base type: int (*)(void *, unsigned int, unsigned int, const void *, unsigned int *, void *)
type AudioFileComponentGetGlobalInfoProc uintptr
// AudioFileComponentGetGlobalInfoSizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSizeProc
// AudioFileComponentGetGlobalInfoSizeProc has base type: int (*)(void *, unsigned int, unsigned int, const void *, unsigned int *)
type AudioFileComponentGetGlobalInfoSizeProc uintptr
// AudioFileComponentGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfoProc
// AudioFileComponentGetPropertyInfoProc has base type: int (*)(void *, unsigned int, unsigned int *, unsigned int *)
type AudioFileComponentGetPropertyInfoProc uintptr
// AudioFileComponentGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyProc
// AudioFileComponentGetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int *, void *)
type AudioFileComponentGetPropertyProc uintptr
// AudioFileComponentGetUserDataAtOffsetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffsetProc
// AudioFileComponentGetUserDataAtOffsetProc has base type: int (*)(void *, unsigned int, unsigned int, long long, unsigned int *, void *)
type AudioFileComponentGetUserDataAtOffsetProc uintptr
// AudioFileComponentGetUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataProc
// AudioFileComponentGetUserDataProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int *, void *)
type AudioFileComponentGetUserDataProc uintptr
// AudioFileComponentGetUserDataSize64Proc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64Proc
// AudioFileComponentGetUserDataSize64Proc has base type: int (*)(void *, unsigned int, unsigned int, unsigned long long *)
type AudioFileComponentGetUserDataSize64Proc uintptr
// AudioFileComponentGetUserDataSizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSizeProc
// AudioFileComponentGetUserDataSizeProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int *)
type AudioFileComponentGetUserDataSizeProc uintptr
// AudioFileComponentInitializeWithCallbacksProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacksProc
// AudioFileComponentInitializeWithCallbacksProc has base type: int (*)(void *, void *, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long), unsigned int, const struct AudioStreamBasicDescription *, unsigned int)
type AudioFileComponentInitializeWithCallbacksProc uintptr
// AudioFileComponentOpenURLProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURLProc
// AudioFileComponentOpenURLProc has base type: int (*)(void *, const struct __CFURL *, signed char, int)
type AudioFileComponentOpenURLProc uintptr
// AudioFileComponentOpenWithCallbacksProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacksProc
// AudioFileComponentOpenWithCallbacksProc has base type: int (*)(void *, void *, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long))
type AudioFileComponentOpenWithCallbacksProc uintptr
// AudioFileComponentOptimizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimizeProc
// AudioFileComponentOptimizeProc has base type: int (*)(void *)
type AudioFileComponentOptimizeProc uintptr
// AudioFileComponentPropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentPropertyID
// AudioFileComponentPropertyID has base type: UInt32
type AudioFileComponentPropertyID uintptr
// AudioFileComponentReadBytesProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytesProc
// AudioFileComponentReadBytesProc has base type: int (*)(void *, unsigned char, long long, unsigned int *, void *)
type AudioFileComponentReadBytesProc uintptr
// AudioFileComponentReadPacketDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketDataProc
// AudioFileComponentReadPacketDataProc has base type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type AudioFileComponentReadPacketDataProc uintptr
// AudioFileComponentReadPacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketsProc
// AudioFileComponentReadPacketsProc has base type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type AudioFileComponentReadPacketsProc uintptr
// AudioFileComponentRemoveUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserDataProc
// AudioFileComponentRemoveUserDataProc has base type: int (*)(void *, unsigned int, unsigned int)
type AudioFileComponentRemoveUserDataProc uintptr
// AudioFileComponentSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetPropertyProc
// AudioFileComponentSetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int, const void *)
type AudioFileComponentSetPropertyProc uintptr
// AudioFileComponentSetUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserDataProc
// AudioFileComponentSetUserDataProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *)
type AudioFileComponentSetUserDataProc uintptr
// AudioFileComponentWriteBytesProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytesProc
// AudioFileComponentWriteBytesProc has base type: int (*)(void *, unsigned char, long long, unsigned int *, const void *)
type AudioFileComponentWriteBytesProc uintptr
// AudioFileComponentWritePacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePacketsProc
// AudioFileComponentWritePacketsProc has base type: int (*)(void *, unsigned char, unsigned int, const struct AudioStreamPacketDescription *, long long, unsigned int *, const void *)
type AudioFileComponentWritePacketsProc uintptr
// AudioFileID - An opaque data type that represents an audio file object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileID
// AudioFileID has base type: struct OpaqueAudioFileID *
type AudioFileID uintptr
// AudioFilePropertyID - An audio file property identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePropertyID
// AudioFilePropertyID has base type: UInt32
type AudioFilePropertyID uintptr
// AudioFileStreamID - Defines an opaque data type that represents an audio file stream parser.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamID
// AudioFileStreamID has base type: struct OpaqueAudioFileStreamID *
type AudioFileStreamID uintptr
// AudioFileStreamPropertyID - Uniquely identifies an audio file stream property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamPropertyID
// AudioFileStreamPropertyID has base type: UInt32
type AudioFileStreamPropertyID uintptr
// AudioFileStream_PacketsProc - Invoked by an audio file stream parser when it finds audio data in the audio file stream.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStream_PacketsProc
// AudioFileStream_PacketsProc has base type: void (*)(void *, unsigned int, unsigned int, const void *, struct AudioStreamPacketDescription *)
type AudioFileStream_PacketsProc uintptr
// AudioFileStream_PropertyListenerProc - Invoked by an audio file stream parser when it finds a property value in the audio file stream.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStream_PropertyListenerProc
// AudioFileStream_PropertyListenerProc has base type: void (*)(void *, struct OpaqueAudioFileStreamID *, unsigned int, enum AudioFileStreamPropertyFlags *)
type AudioFileStream_PropertyListenerProc uintptr
// AudioFileTypeID - Operating system constants that indicate the type of file to be written or a hint about what type of file to expect from data provided.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeID
// AudioFileTypeID has base type: UInt32
type AudioFileTypeID uintptr
// AudioFile_GetSizeProc - Gets file data size.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_GetSizeProc
// AudioFile_GetSizeProc has base type: long long (*)(void *)
type AudioFile_GetSizeProc uintptr
// AudioFile_ReadProc - Reads audio data when used in conjunction with the   or   functions.)
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_ReadProc
// AudioFile_ReadProc has base type: int (*)(void *, long long, unsigned int, void *, unsigned int *)
type AudioFile_ReadProc uintptr
// AudioFile_SetSizeProc - Sets file data size.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SetSizeProc
// AudioFile_SetSizeProc has base type: int (*)(void *, long long)
type AudioFile_SetSizeProc uintptr
// AudioFile_WriteProc - A callback for writing file data when used in conjunction with the   or   functions.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_WriteProc
// AudioFile_WriteProc has base type: int (*)(void *, long long, unsigned int, const void *, unsigned int *)
type AudioFile_WriteProc uintptr
// AudioFormatPropertyID - A type for four-char codes for audio format property identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatPropertyID
// AudioFormatPropertyID has base type: UInt32
type AudioFormatPropertyID uintptr
// AudioOutputUnitStartProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartProc
// AudioOutputUnitStartProc has base type: int (*)(void *)
type AudioOutputUnitStartProc uintptr
// AudioOutputUnitStopProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStopProc
// AudioOutputUnitStopProc has base type: int (*)(void *)
type AudioOutputUnitStopProc uintptr
// AudioQueueBufferRef - A pointer to an audio queue buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBufferRef
// AudioQueueBufferRef has base type: AudioQueueBuffer *
type AudioQueueBufferRef uintptr
// AudioQueueInputCallback - Called by the system when a recording audio queue has finished filling an audio queue buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueInputCallback
// AudioQueueInputCallback has base type: void (*)(void *, struct OpaqueAudioQueue *, struct AudioQueueBuffer *, const struct AudioTimeStamp *, unsigned int, const struct AudioStreamPacketDescription *)
type AudioQueueInputCallback uintptr
// AudioQueueOutputCallback - Called by the system when an audio queue buffer is available for reuse.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOutputCallback
// AudioQueueOutputCallback has base type: void (*)(void *, struct OpaqueAudioQueue *, struct AudioQueueBuffer *)
type AudioQueueOutputCallback uintptr
// AudioQueueParameterID - A   value that uniquely identifies an audio queue parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterID
// AudioQueueParameterID has base type: UInt32
type AudioQueueParameterID uintptr
// AudioQueueParameterValue - A   value for an audio queue parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterValue
// AudioQueueParameterValue has base type: Float32
type AudioQueueParameterValue uintptr
// AudioQueueProcessingTapCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapCallback
// AudioQueueProcessingTapCallback has base type: void (*)(void *, struct OpaqueAudioQueueProcessingTap *, unsigned int, struct AudioTimeStamp *, enum AudioQueueProcessingTapFlags *, unsigned int *, struct AudioBufferList *)
type AudioQueueProcessingTapCallback uintptr
// AudioQueueProcessingTapRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapRef
// AudioQueueProcessingTapRef has base type: struct OpaqueAudioQueueProcessingTap *
type AudioQueueProcessingTapRef uintptr
// AudioQueuePropertyID - Identifiers for audio queue properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePropertyID
// AudioQueuePropertyID has base type: UInt32
type AudioQueuePropertyID uintptr
// AudioQueuePropertyListenerProc - Called by the system when a specified audio queue property changes value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePropertyListenerProc
// AudioQueuePropertyListenerProc has base type: void (*)(void *, struct OpaqueAudioQueue *, unsigned int)
type AudioQueuePropertyListenerProc uintptr
// AudioQueueRef - Defines an opaque data type that represents an audio queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueRef
// AudioQueueRef has base type: struct OpaqueAudioQueue *
type AudioQueueRef uintptr
// AudioQueueTimelineRef - Defines an opaque data type that represents an audio queue timeline object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueTimelineRef
// AudioQueueTimelineRef has base type: struct OpaqueAudioQueueTimeline *
type AudioQueueTimelineRef uintptr
// AudioServicesPropertyID - The data type for a system sound property identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPropertyID
// AudioServicesPropertyID has base type: UInt32
type AudioServicesPropertyID uintptr
// AudioServicesSystemSoundCompletionProc - A function the system invokes when a system sound finishes playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesSystemSoundCompletionProc
// AudioServicesSystemSoundCompletionProc has base type: void (*)(unsigned int, void *)
type AudioServicesSystemSoundCompletionProc uintptr
// AudioSessionInterruptionListener - Invoked when an audio interruption in iOS begins or ends.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInterruptionListener
// AudioSessionInterruptionListener has base type: void (*)(void *, unsigned int)
type AudioSessionInterruptionListener uintptr
// AudioSessionInterruptionType - Values that indicate the nature of the interruption that ended.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInterruptionType
// AudioSessionInterruptionType has base type: UInt32
type AudioSessionInterruptionType uintptr
// AudioSessionPropertyID - The data type for an audio session property identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionPropertyID
// AudioSessionPropertyID has base type: UInt32
type AudioSessionPropertyID uintptr
// AudioSessionPropertyListener - Invoked when an audio session property changes in iOS.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionPropertyListener
// AudioSessionPropertyListener has base type: void (*)(void *, unsigned int, unsigned int, const void *)
type AudioSessionPropertyListener uintptr
// AudioUnit - The data type for a plug-in component that provides audio processing or audio data generation.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnit
// AudioUnit has base type: AudioComponentInstance
type AudioUnit uintptr
// AudioUnitAddPropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListenerProc
// AudioUnitAddPropertyListenerProc has base type: int (*)(void *, unsigned int, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int), void *)
type AudioUnitAddPropertyListenerProc uintptr
// AudioUnitAddRenderNotifyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotifyProc
// AudioUnitAddRenderNotifyProc has base type: int (*)(void *, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), void *)
type AudioUnitAddRenderNotifyProc uintptr
// AudioUnitComplexRenderProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitComplexRenderProc
// AudioUnitComplexRenderProc has base type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, unsigned int *, struct AudioStreamPacketDescription *, struct AudioBufferList *, void *, unsigned int *) __attribute__((nonblocking))
type AudioUnitComplexRenderProc uintptr
// AudioUnitElement - The data type for an audio unit element identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitElement
// AudioUnitElement has base type: UInt32
type AudioUnitElement uintptr
// AudioUnitGetParameterProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameterProc
// AudioUnitGetParameterProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, float *) __attribute__((nonblocking))
type AudioUnitGetParameterProc uintptr
// AudioUnitGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfoProc
// AudioUnitGetPropertyInfoProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, unsigned int *, unsigned char *)
type AudioUnitGetPropertyInfoProc uintptr
// AudioUnitGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyProc
// AudioUnitGetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, void *, unsigned int *)
type AudioUnitGetPropertyProc uintptr
// AudioUnitInitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitializeProc
// AudioUnitInitializeProc has base type: int (*)(void *)
type AudioUnitInitializeProc uintptr
// AudioUnitParameterID - The data type for an audio unit parameter identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterID
// AudioUnitParameterID has base type: UInt32
type AudioUnitParameterID uintptr
// AudioUnitParameterIDName - A type definition for a data type that defines the short version of the name for an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterIDName
// AudioUnitParameterIDName has base type: struct AudioUnitParameterNameInfo
type AudioUnitParameterIDName uintptr
// AudioUnitParameterValue - The data type for an audio unit parameter value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValue
// AudioUnitParameterValue has base type: Float32
type AudioUnitParameterValue uintptr
// AudioUnitProcessMultipleProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessMultipleProc
// AudioUnitProcessMultipleProc has base type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, const struct AudioBufferList **, unsigned int, struct AudioBufferList **) __attribute__((nonblocking))
type AudioUnitProcessMultipleProc uintptr
// AudioUnitProcessProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessProc
// AudioUnitProcessProc has base type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type AudioUnitProcessProc uintptr
// AudioUnitPropertyID - The data type for audio unit property keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyID
// AudioUnitPropertyID has base type: UInt32
type AudioUnitPropertyID uintptr
// AudioUnitPropertyListenerProc - Called by the system when the value of a specified audio unit property has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyListenerProc
// AudioUnitPropertyListenerProc has base type: void (*)(void *, struct OpaqueAudioComponentInstance *, unsigned int, unsigned int, unsigned int)
type AudioUnitPropertyListenerProc uintptr
// AudioUnitRemovePropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerProc
// AudioUnitRemovePropertyListenerProc has base type: int (*)(void *, unsigned int, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int))
type AudioUnitRemovePropertyListenerProc uintptr
// AudioUnitRemovePropertyListenerWithUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserDataProc
// AudioUnitRemovePropertyListenerWithUserDataProc has base type: int (*)(void *, unsigned int, void (*)(void *, struct OpaqueAudioComponentInstance *, unsigned int, unsigned int, unsigned int), void *)
type AudioUnitRemovePropertyListenerWithUserDataProc uintptr
// AudioUnitRemoveRenderNotifyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotifyProc
// AudioUnitRemoveRenderNotifyProc has base type: int (*)(void *, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), void *)
type AudioUnitRemoveRenderNotifyProc uintptr
// AudioUnitRenderProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderProc
// AudioUnitRenderProc has base type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type AudioUnitRenderProc uintptr
// AudioUnitResetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitResetProc
// AudioUnitResetProc has base type: int (*)(void *, unsigned int, unsigned int)
type AudioUnitResetProc uintptr
// AudioUnitScheduleParametersProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParametersProc
// AudioUnitScheduleParametersProc has base type: int (*)(void *, const struct AudioUnitParameterEvent *, unsigned int) __attribute__((nonblocking))
type AudioUnitScheduleParametersProc uintptr
// AudioUnitScope - The data type for audio unit scope identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScope
// AudioUnitScope has base type: UInt32
type AudioUnitScope uintptr
// AudioUnitSetParameterProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameterProc
// AudioUnitSetParameterProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, float, unsigned int) __attribute__((nonblocking))
type AudioUnitSetParameterProc uintptr
// AudioUnitSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetPropertyProc
// AudioUnitSetPropertyProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *, unsigned int)
type AudioUnitSetPropertyProc uintptr
// AudioUnitUninitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitializeProc
// AudioUnitUninitializeProc has base type: int (*)(void *)
type AudioUnitUninitializeProc uintptr
// CAClockBeats type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBeats
// CAClockBeats has base type: Float64
type CAClockBeats uintptr
// CAClockListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockListenerProc
// CAClockListenerProc has base type: void (*)(void *, enum CAClockMessage, const void *)
type CAClockListenerProc uintptr
// CAClockRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockRef
// CAClockRef has base type: struct OpaqueCAClock *
type CAClockRef uintptr
// CAClockSMPTEFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTEFormat
// CAClockSMPTEFormat has base type: SMPTETimeType
type CAClockSMPTEFormat uintptr
// CAClockSamples type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSamples
// CAClockSamples has base type: Float64
type CAClockSamples uintptr
// CAClockSeconds type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSeconds
// CAClockSeconds has base type: Float64
type CAClockSeconds uintptr
// CAClockTempo type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTempo
// CAClockTempo has base type: Float64
type CAClockTempo uintptr
// CountUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CountUserDataFDF
// CountUserDataFDF has base type: int (*)(void *, unsigned int, unsigned int *)
type CountUserDataFDF uintptr
// ExtAudioFilePacketTableInfoOverride type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFilePacketTableInfoOverride
// ExtAudioFilePacketTableInfoOverride has base type: SInt32
type ExtAudioFilePacketTableInfoOverride uintptr
// ExtAudioFilePropertyID - An audio file object property identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFilePropertyID
// ExtAudioFilePropertyID has base type: UInt32
type ExtAudioFilePropertyID uintptr
// ExtAudioFileRef - An opaque structure representing an extended audio file object.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileRef
// ExtAudioFileRef has base type: struct OpaqueExtAudioFile *
type ExtAudioFileRef uintptr
// GetPropertyFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetPropertyFDF
// GetPropertyFDF has base type: int (*)(void *, unsigned int, unsigned int *, void *)
type GetPropertyFDF uintptr
// GetPropertyInfoFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetPropertyInfoFDF
// GetPropertyInfoFDF has base type: int (*)(void *, unsigned int, unsigned int *, unsigned int *)
type GetPropertyInfoFDF uintptr
// GetUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetUserDataFDF
// GetUserDataFDF has base type: int (*)(void *, unsigned int, unsigned int, unsigned int *, void *)
type GetUserDataFDF uintptr
// GetUserDataSizeFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetUserDataSizeFDF
// GetUserDataSizeFDF has base type: int (*)(void *, unsigned int, unsigned int, unsigned int *)
type GetUserDataSizeFDF uintptr
// HostCallback_GetBeatAndTempo - When called by the system, provides beat and tempo information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetBeatAndTempo
// HostCallback_GetBeatAndTempo has base type: int (*)(void *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetBeatAndTempo uintptr
// HostCallback_GetMusicalTimeLocation - When called by the system, provides musical timing information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetMusicalTimeLocation
// HostCallback_GetMusicalTimeLocation has base type: int (*)(void *, unsigned int *, float *, unsigned int *, double *) __attribute__((nonblocking))
type HostCallback_GetMusicalTimeLocation uintptr
// HostCallback_GetTransportState - When called by the system, provides audio transport state and timeline information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState
// HostCallback_GetTransportState has base type: int (*)(void *, unsigned char *, unsigned char *, double *, unsigned char *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetTransportState uintptr
// HostCallback_GetTransportState2 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState2
// HostCallback_GetTransportState2 has base type: int (*)(void *, unsigned char *, unsigned char *, unsigned char *, double *, unsigned char *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetTransportState2 uintptr
// MIDIChannelNumber - MIDI Channel, 0~15 (channels 1 through 16, respectively).
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIChannelNumber
// MIDIChannelNumber has base type: uint8_t
type MIDIChannelNumber uintptr
// MagicCookieInfo - A structure holding magic cookie information.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MagicCookieInfo
// MagicCookieInfo has base type: struct AudioCodecMagicCookieInfo
type MagicCookieInfo uintptr
// MusicDeviceComponent type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceComponent
// MusicDeviceComponent has base type: AudioComponentInstance
type MusicDeviceComponent uintptr
// MusicDeviceGroupID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceGroupID
// MusicDeviceGroupID has base type: UInt32
type MusicDeviceGroupID uintptr
// MusicDeviceInstrumentID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceInstrumentID
// MusicDeviceInstrumentID has base type: UInt32
type MusicDeviceInstrumentID uintptr
// MusicDeviceMIDIEventProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEventProc
// MusicDeviceMIDIEventProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, unsigned int) __attribute__((nonblocking))
type MusicDeviceMIDIEventProc uintptr
// MusicDeviceStartNoteProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNoteProc
// MusicDeviceStartNoteProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int *, unsigned int, const struct MusicDeviceNoteParams *) __attribute__((nonblocking))
type MusicDeviceStartNoteProc uintptr
// MusicDeviceStopNoteProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNoteProc
// MusicDeviceStopNoteProc has base type: int (*)(void *, unsigned int, unsigned int, unsigned int) __attribute__((nonblocking))
type MusicDeviceStopNoteProc uintptr
// MusicDeviceSysExProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysExProc
// MusicDeviceSysExProc has base type: int (*)(void *, const unsigned char *, unsigned int) __attribute__((nonblocking))
type MusicDeviceSysExProc uintptr
// MusicEventIterator - A music event iterator sequentially handles events on a music track.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIterator
// MusicEventIterator has base type: struct OpaqueMusicEventIterator *
type MusicEventIterator uintptr
// MusicEventType - MIDI and other music event types, used by music event iterator functions.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventType
// MusicEventType has base type: UInt32
type MusicEventType uintptr
// MusicPlayer - A music player plays a music sequence (of type  ).
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayer
// MusicPlayer has base type: struct OpaqueMusicPlayer *
type MusicPlayer uintptr
// MusicSequence - A music sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequence
// MusicSequence has base type: struct OpaqueMusicSequence *
type MusicSequence uintptr
// MusicSequenceUserCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceUserCallback
// MusicSequenceUserCallback has base type: void (*)(void *, struct OpaqueMusicSequence *, struct OpaqueMusicTrack *, double, const struct MusicEventUserData *, double, double) __attribute__((nonblocking))
type MusicSequenceUserCallback uintptr
// MusicTimeStamp - A timestamp for use by a music sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTimeStamp
// MusicTimeStamp has base type: Float64
type MusicTimeStamp uintptr
// MusicTrack - A music track consists of a series of music events, each timestamped using units of beats.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrack
// MusicTrack has base type: struct OpaqueMusicTrack *
type MusicTrack uintptr
// NoteInstanceID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NoteInstanceID
// NoteInstanceID has base type: UInt32
type NoteInstanceID uintptr
// ReadBytesFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ReadBytesFDF
// ReadBytesFDF has base type: int (*)(void *, unsigned char, long long, unsigned int *, void *)
type ReadBytesFDF uintptr
// ReadPacketDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ReadPacketDataFDF
// ReadPacketDataFDF has base type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type ReadPacketDataFDF uintptr
// ReadPacketsFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ReadPacketsFDF
// ReadPacketsFDF has base type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type ReadPacketsFDF uintptr
// ScheduledAudioFileRegionCompletionProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegionCompletionProc
// ScheduledAudioFileRegionCompletionProc has base type: void (*)(void *, struct ScheduledAudioFileRegion *, int)
type ScheduledAudioFileRegionCompletionProc uintptr
// ScheduledAudioSliceCompletionProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSliceCompletionProc
// ScheduledAudioSliceCompletionProc has base type: void (*)(void *, struct ScheduledAudioSlice *) __attribute__((nonblocking))
type ScheduledAudioSliceCompletionProc uintptr
// SetPropertyFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetPropertyFDF
// SetPropertyFDF has base type: int (*)(void *, unsigned int, unsigned int, const void *)
type SetPropertyFDF uintptr
// SetUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetUserDataFDF
// SetUserDataFDF has base type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *)
type SetUserDataFDF uintptr
// SystemSoundID - A system sound object, identified with a sound file you want to play.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SystemSoundID
// SystemSoundID has base type: UInt32
type SystemSoundID uintptr
// WriteBytesFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/WriteBytesFDF
// WriteBytesFDF has base type: int (*)(void *, unsigned char, long long, unsigned int *, const void *)
type WriteBytesFDF uintptr
// WritePacketsFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/WritePacketsFDF
// WritePacketsFDF has base type: int (*)(void *, unsigned char, unsigned int, const struct AudioStreamPacketDescription *, long long, unsigned int *, const void *)
type WritePacketsFDF uintptr

