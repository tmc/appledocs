// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox
import (
"unsafe"
)

// Type aliases and typedefs
// AudioChannelCount - A number of audio channels.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioChannelCount
// AUAudioChannelCount has base type: uint32_t
type AudioChannelCount uintptr
// AudioFrameCount - A number of audio sample frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioFrameCount
// AUAudioFrameCount has base type: uint32_t
type AudioFrameCount uintptr
// AudioObjectID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioObjectID
// AUAudioObjectID has base type: UInt32
type AudioObjectID uintptr
// AudioUnitStatus - A result code returned from an audio unit’s render function.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitStatus
// AUAudioUnitStatus has base type: OSStatus
type AudioUnitStatus uintptr
// EventListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerProc
// AUEventListenerProc is a callback function
// C type: void (*)(void *, void *, const struct AudioUnitEvent *, unsigned long long, float)
type EventListenerProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint64, float32)
// EventListenerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRef
// AUEventListenerRef has base type: AUParameterListenerRef
type EventListenerRef uintptr
// EventSampleTime - Expresses time as a sample count.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventSampleTime
// AUEventSampleTime has base type: int64_t
type EventSampleTime uintptr
// InputSamplesInOutputCallback - Called by the system when an audio unit has provided a buffer of output samples.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallback
// AUInputSamplesInOutputCallback is a callback function
// C type: void (*)(void *, const struct AudioTimeStamp *, double, double) __attribute__((nonblocking))
type InputSamplesInOutputCallback = func(unsafe.Pointer, unsafe.Pointer, float64, double) __attribute__((nonblocking))
// MIDIOutputCallback - When called by a host application, gets MIDI data from an audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallback
// AUMIDIOutputCallback is a callback function
// C type: int (*)(void *, const struct AudioTimeStamp *, unsigned int, const struct MIDIPacketList *) __attribute__((nonblocking))
type MIDIOutputCallback = func(unsafe.Pointer, unsafe.Pointer, uint32, MIDIPacketList *) __attribute__((nonblocking)) int32
// Node - A member of an audio processing graph, associated with an audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNode
// AUNode has base type: SInt32
type Node uintptr
// NodeConnection type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeConnection
// AUNodeConnection has base type: struct AudioUnitNodeConnection
type NodeConnection uintptr
// ParameterAddress - A numeric identifier for an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAddress
// AUParameterAddress has base type: uint64_t
type ParameterAddress uintptr
// ParameterListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerProc
// AUParameterListenerProc is a callback function
// C type: void (*)(void *, void *, const struct AudioUnitParameter *, float)
type ParameterListenerProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float32)
// ParameterListenerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerRef
// AUParameterListenerRef has base type: struct AUListenerBase *
type ParameterListenerRef uintptr
// ParameterObserverToken - A token representing an installed parameter observer block.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterObserverToken
// AUParameterObserverToken has base type: void *
type ParameterObserverToken uintptr
// RenderCallback - Called by the system when an audio unit requires input samples, or before and after a render operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderCallback
// AURenderCallback is a callback function
// C type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type RenderCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, uint32, AudioBufferList *) __attribute__((nonblocking)) int32
// Value - A value of an audio unit parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUValue
// AUValue has base type: float
type Value uintptr
// AudioCodec - An instance of a Component Manager component.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodec
// AudioCodec has base type: AudioComponentInstance
type AudioCodec uintptr
// AudioCodecAppendInputBufferListProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferListProc
// AudioCodecAppendInputBufferListProc is a callback function
// C type: int (*)(void *, const struct AudioBufferList *, unsigned int *, const struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecAppendInputBufferListProc = func(unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer, uint32) int32
// AudioCodecAppendInputDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputDataProc
// AudioCodecAppendInputDataProc is a callback function
// C type: int (*)(void *, const void *, unsigned int *, unsigned int *, const struct AudioStreamPacketDescription *)
type AudioCodecAppendInputDataProc = func(unsafe.Pointer, unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// AudioCodecGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfoProc
// AudioCodecGetPropertyInfoProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, unsigned char *)
type AudioCodecGetPropertyInfoProc = func(unsafe.Pointer, uint32, uint32, uint8) int32
// AudioCodecGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyProc
// AudioCodecGetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, void *)
type AudioCodecGetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// AudioCodecInitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitializeProc
// AudioCodecInitializeProc is a callback function
// C type: int (*)(void *, const struct AudioStreamBasicDescription *, const struct AudioStreamBasicDescription *, const void *, unsigned int)
type AudioCodecInitializeProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32) int32
// AudioCodecProduceOutputBufferListProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferListProc
// AudioCodecProduceOutputBufferListProc is a callback function
// C type: int (*)(void *, struct AudioBufferList *, unsigned int *, struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecProduceOutputBufferListProc = func(unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer, uint32) int32
// AudioCodecProduceOutputPacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPacketsProc
// AudioCodecProduceOutputPacketsProc is a callback function
// C type: int (*)(void *, void *, unsigned int *, unsigned int *, struct AudioStreamPacketDescription *, unsigned int *)
type AudioCodecProduceOutputPacketsProc = func(unsafe.Pointer, unsafe.Pointer, uint32, uint32, unsafe.Pointer, uint32) int32
// AudioCodecPropertyID - An integer identifying an audio codec property.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPropertyID
// AudioCodecPropertyID has base type: UInt32
type AudioCodecPropertyID uintptr
// AudioCodecResetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecResetProc
// AudioCodecResetProc is a callback function
// C type: int (*)(void *)
type AudioCodecResetProc = func(unsafe.Pointer) int32
// AudioCodecSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetPropertyProc
// AudioCodecSetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, const void *)
type AudioCodecSetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// AudioCodecUninitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitializeProc
// AudioCodecUninitializeProc is a callback function
// C type: int (*)(void *)
type AudioCodecUninitializeProc = func(unsafe.Pointer) int32
// AudioComponent - An audio component.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponent
// AudioComponent has base type: struct OpaqueAudioComponent *
type AudioComponent uintptr
// AudioComponentFactoryFunction type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFactoryFunction
// AudioComponentFactoryFunction is a callback function
// C type: struct AudioComponentPlugInInterface *(*)(const struct AudioComponentDescription *)
type AudioComponentFactoryFunction = func(unsafe.Pointer) unsafe.Pointer
// AudioComponentInstance - A component instance, or object, is an audio unit or audio codec.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstance
// AudioComponentInstance has base type: struct OpaqueAudioComponentInstance *
type AudioComponentInstance uintptr
// AudioComponentMethod type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentMethod
// AudioComponentMethod is a callback function
// C type: int (*)(void *, ...)
type AudioComponentMethod = func(unsafe.Pointer, ...) int32
// AudioConverterComplexInputDataProc - Supplies input data to the   function.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProc
// AudioConverterComplexInputDataProc is a callback function
// C type: int (*)(struct OpaqueAudioConverter *, unsigned int *, struct AudioBufferList *, struct AudioStreamPacketDescription **, void *)
type AudioConverterComplexInputDataProc = func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32
// AudioConverterComplexInputDataProcRealtimeSafe type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterComplexInputDataProcRealtimeSafe
// AudioConverterComplexInputDataProcRealtimeSafe is a callback function
// C type: int (*)(struct OpaqueAudioConverter *, unsigned int *, struct AudioBufferList *, struct AudioStreamPacketDescription **, void *) __attribute__((nonblocking))
type AudioConverterComplexInputDataProcRealtimeSafe = func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, void *) __attribute__((nonblocking)) int32
// AudioConverterInputDataProc - Deprecated. Use   instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterInputDataProc
// AudioConverterInputDataProc is a callback function
// C type: int (*)(struct OpaqueAudioConverter *, unsigned int *, void **, void *)
type AudioConverterInputDataProc = func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) int32
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
// AudioFileComponentCloseProc is a callback function
// C type: int (*)(void *)
type AudioFileComponentCloseProc = func(unsafe.Pointer) int32
// AudioFileComponentCountUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserDataProc
// AudioFileComponentCountUserDataProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *)
type AudioFileComponentCountUserDataProc = func(unsafe.Pointer, uint32, uint32) int32
// AudioFileComponentCreateURLProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURLProc
// AudioFileComponentCreateURLProc is a callback function
// C type: int (*)(void *, const struct __CFURL *, const struct AudioStreamBasicDescription *, unsigned int)
type AudioFileComponentCreateURLProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32) int32
// AudioFileComponentExtensionIsThisFormatProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormatProc
// AudioFileComponentExtensionIsThisFormatProc is a callback function
// C type: int (*)(void *, const struct __CFString *, unsigned int *)
type AudioFileComponentExtensionIsThisFormatProc = func(unsafe.Pointer, unsafe.Pointer, uint32) int32
// AudioFileComponentFileDataIsThisFormatProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormatProc
// AudioFileComponentFileDataIsThisFormatProc is a callback function
// C type: int (*)(void *, unsigned int, const void *, unsigned int *)
type AudioFileComponentFileDataIsThisFormatProc = func(unsafe.Pointer, uint32, unsafe.Pointer, uint32) int32
// AudioFileComponentGetGlobalInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoProc
// AudioFileComponentGetGlobalInfoProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, const void *, unsigned int *, void *)
type AudioFileComponentGetGlobalInfoProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, uint32, unsafe.Pointer) int32
// AudioFileComponentGetGlobalInfoSizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSizeProc
// AudioFileComponentGetGlobalInfoSizeProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, const void *, unsigned int *)
type AudioFileComponentGetGlobalInfoSizeProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, uint32) int32
// AudioFileComponentGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfoProc
// AudioFileComponentGetPropertyInfoProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, unsigned int *)
type AudioFileComponentGetPropertyInfoProc = func(unsafe.Pointer, uint32, uint32, uint32) int32
// AudioFileComponentGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyProc
// AudioFileComponentGetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, void *)
type AudioFileComponentGetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// AudioFileComponentGetUserDataAtOffsetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffsetProc
// AudioFileComponentGetUserDataAtOffsetProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, long long, unsigned int *, void *)
type AudioFileComponentGetUserDataAtOffsetProc = func(unsafe.Pointer, uint32, uint32, int64, uint32, unsafe.Pointer) int32
// AudioFileComponentGetUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataProc
// AudioFileComponentGetUserDataProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int *, void *)
type AudioFileComponentGetUserDataProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32
// AudioFileComponentGetUserDataSize64Proc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64Proc
// AudioFileComponentGetUserDataSize64Proc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned long long *)
type AudioFileComponentGetUserDataSize64Proc = func(unsafe.Pointer, uint32, uint32, uint64) int32
// AudioFileComponentGetUserDataSizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSizeProc
// AudioFileComponentGetUserDataSizeProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int *)
type AudioFileComponentGetUserDataSizeProc = func(unsafe.Pointer, uint32, uint32, uint32) int32
// AudioFileComponentInitializeWithCallbacksProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacksProc
// AudioFileComponentInitializeWithCallbacksProc is a callback function
// C type: int (*)(void *, void *, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long), unsigned int, const struct AudioStreamBasicDescription *, unsigned int)
type AudioFileComponentInitializeWithCallbacksProc = func(unsafe.Pointer, unsafe.Pointer, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long), uint32, unsafe.Pointer, uint32) int32
// AudioFileComponentOpenURLProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURLProc
// AudioFileComponentOpenURLProc is a callback function
// C type: int (*)(void *, const struct __CFURL *, signed char, int)
type AudioFileComponentOpenURLProc = func(unsafe.Pointer, unsafe.Pointer, signed char, int32) int32
// AudioFileComponentOpenWithCallbacksProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacksProc
// AudioFileComponentOpenWithCallbacksProc is a callback function
// C type: int (*)(void *, void *, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long))
type AudioFileComponentOpenWithCallbacksProc = func(unsafe.Pointer, unsafe.Pointer, int (*)(void *, long long, unsigned int, void *, unsigned int *), int (*)(void *, long long, unsigned int, const void *, unsigned int *), long long (*)(void *), int (*)(void *, long long)) int32
// AudioFileComponentOptimizeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimizeProc
// AudioFileComponentOptimizeProc is a callback function
// C type: int (*)(void *)
type AudioFileComponentOptimizeProc = func(unsafe.Pointer) int32
// AudioFileComponentPropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentPropertyID
// AudioFileComponentPropertyID has base type: UInt32
type AudioFileComponentPropertyID uintptr
// AudioFileComponentReadBytesProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytesProc
// AudioFileComponentReadBytesProc is a callback function
// C type: int (*)(void *, unsigned char, long long, unsigned int *, void *)
type AudioFileComponentReadBytesProc = func(unsafe.Pointer, uint8, int64, uint32, unsafe.Pointer) int32
// AudioFileComponentReadPacketDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketDataProc
// AudioFileComponentReadPacketDataProc is a callback function
// C type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type AudioFileComponentReadPacketDataProc = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32
// AudioFileComponentReadPacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketsProc
// AudioFileComponentReadPacketsProc is a callback function
// C type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type AudioFileComponentReadPacketsProc = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32
// AudioFileComponentRemoveUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserDataProc
// AudioFileComponentRemoveUserDataProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int)
type AudioFileComponentRemoveUserDataProc = func(unsafe.Pointer, uint32, uint32) int32
// AudioFileComponentSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetPropertyProc
// AudioFileComponentSetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, const void *)
type AudioFileComponentSetPropertyProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// AudioFileComponentSetUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserDataProc
// AudioFileComponentSetUserDataProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *)
type AudioFileComponentSetUserDataProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32
// AudioFileComponentWriteBytesProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytesProc
// AudioFileComponentWriteBytesProc is a callback function
// C type: int (*)(void *, unsigned char, long long, unsigned int *, const void *)
type AudioFileComponentWriteBytesProc = func(unsafe.Pointer, uint8, int64, uint32, unsafe.Pointer) int32
// AudioFileComponentWritePacketsProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePacketsProc
// AudioFileComponentWritePacketsProc is a callback function
// C type: int (*)(void *, unsigned char, unsigned int, const struct AudioStreamPacketDescription *, long long, unsigned int *, const void *)
type AudioFileComponentWritePacketsProc = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32
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
// AudioFileStream_PacketsProc is a callback function
// C type: void (*)(void *, unsigned int, unsigned int, const void *, struct AudioStreamPacketDescription *)
type AudioFileStream_PacketsProc = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, unsafe.Pointer)
// AudioFileStream_PropertyListenerProc - Invoked by an audio file stream parser when it finds a property value in the audio file stream.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStream_PropertyListenerProc
// AudioFileStream_PropertyListenerProc is a callback function
// C type: void (*)(void *, struct OpaqueAudioFileStreamID *, unsigned int, enum AudioFileStreamPropertyFlags *)
type AudioFileStream_PropertyListenerProc = func(unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer)
// AudioFileTypeID - Operating system constants that indicate the type of file to be written or a hint about what type of file to expect from data provided.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeID
// AudioFileTypeID has base type: UInt32
type AudioFileTypeID uintptr
// AudioFile_GetSizeProc - Gets file data size.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_GetSizeProc
// AudioFile_GetSizeProc is a callback function
// C type: long long (*)(void *)
type AudioFile_GetSizeProc = func(unsafe.Pointer) int64
// AudioFile_ReadProc - Reads audio data when used in conjunction with the   or   functions.)
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_ReadProc
// AudioFile_ReadProc is a callback function
// C type: int (*)(void *, long long, unsigned int, void *, unsigned int *)
type AudioFile_ReadProc = func(unsafe.Pointer, int64, uint32, unsafe.Pointer, uint32) int32
// AudioFile_SetSizeProc - Sets file data size.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SetSizeProc
// AudioFile_SetSizeProc is a callback function
// C type: int (*)(void *, long long)
type AudioFile_SetSizeProc = func(unsafe.Pointer, int64) int32
// AudioFile_WriteProc - A callback for writing file data when used in conjunction with the   or   functions.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_WriteProc
// AudioFile_WriteProc is a callback function
// C type: int (*)(void *, long long, unsigned int, const void *, unsigned int *)
type AudioFile_WriteProc = func(unsafe.Pointer, int64, uint32, unsafe.Pointer, uint32) int32
// AudioFormatPropertyID - A type for four-char codes for audio format property identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatPropertyID
// AudioFormatPropertyID has base type: UInt32
type AudioFormatPropertyID uintptr
// AudioOutputUnitStartProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartProc
// AudioOutputUnitStartProc is a callback function
// C type: int (*)(void *)
type AudioOutputUnitStartProc = func(unsafe.Pointer) int32
// AudioOutputUnitStopProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStopProc
// AudioOutputUnitStopProc is a callback function
// C type: int (*)(void *)
type AudioOutputUnitStopProc = func(unsafe.Pointer) int32
// AudioQueueBufferRef - A pointer to an audio queue buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBufferRef
// AudioQueueBufferRef has base type: AudioQueueBuffer *
type AudioQueueBufferRef uintptr
// AudioQueueInputCallback - Called by the system when a recording audio queue has finished filling an audio queue buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueInputCallback
// AudioQueueInputCallback is a callback function
// C type: void (*)(void *, struct OpaqueAudioQueue *, struct AudioQueueBuffer *, const struct AudioTimeStamp *, unsigned int, const struct AudioStreamPacketDescription *)
type AudioQueueInputCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer)
// AudioQueueOutputCallback - Called by the system when an audio queue buffer is available for reuse.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOutputCallback
// AudioQueueOutputCallback is a callback function
// C type: void (*)(void *, struct OpaqueAudioQueue *, struct AudioQueueBuffer *)
type AudioQueueOutputCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
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
// AudioQueueProcessingTapCallback is a callback function
// C type: void (*)(void *, struct OpaqueAudioQueueProcessingTap *, unsigned int, struct AudioTimeStamp *, enum AudioQueueProcessingTapFlags *, unsigned int *, struct AudioBufferList *)
type AudioQueueProcessingTapCallback = func(unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, uint32, unsafe.Pointer)
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
// AudioQueuePropertyListenerProc is a callback function
// C type: void (*)(void *, struct OpaqueAudioQueue *, unsigned int)
type AudioQueuePropertyListenerProc = func(unsafe.Pointer, unsafe.Pointer, uint32)
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
// AudioServicesSystemSoundCompletionProc is a callback function
// C type: void (*)(unsigned int, void *)
type AudioServicesSystemSoundCompletionProc = func(uint32, unsafe.Pointer)
// AudioSessionInterruptionListener - Invoked when an audio interruption in iOS begins or ends.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInterruptionListener
// AudioSessionInterruptionListener is a callback function
// C type: void (*)(void *, unsigned int)
type AudioSessionInterruptionListener = func(unsafe.Pointer, uint32)
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
// AudioSessionPropertyListener is a callback function
// C type: void (*)(void *, unsigned int, unsigned int, const void *)
type AudioSessionPropertyListener = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer)
// AudioUnit - The data type for a plug-in component that provides audio processing or audio data generation.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnit
// AudioUnit has base type: AudioComponentInstance
type AudioUnit uintptr
// AudioUnitAddPropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListenerProc
// AudioUnitAddPropertyListenerProc is a callback function
// C type: int (*)(void *, unsigned int, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int), void *)
type AudioUnitAddPropertyListenerProc = func(unsafe.Pointer, uint32, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int), unsafe.Pointer) int32
// AudioUnitAddRenderNotifyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotifyProc
// AudioUnitAddRenderNotifyProc is a callback function
// C type: int (*)(void *, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), void *)
type AudioUnitAddRenderNotifyProc = func(unsafe.Pointer, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), unsafe.Pointer) int32
// AudioUnitComplexRenderProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitComplexRenderProc
// AudioUnitComplexRenderProc is a callback function
// C type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, unsigned int *, struct AudioStreamPacketDescription *, struct AudioBufferList *, void *, unsigned int *) __attribute__((nonblocking))
type AudioUnitComplexRenderProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsigned int *) __attribute__((nonblocking)) int32
// AudioUnitElement - The data type for an audio unit element identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitElement
// AudioUnitElement has base type: UInt32
type AudioUnitElement uintptr
// AudioUnitGetParameterProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameterProc
// AudioUnitGetParameterProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, float *) __attribute__((nonblocking))
type AudioUnitGetParameterProc = func(unsafe.Pointer, uint32, uint32, uint32, float *) __attribute__((nonblocking)) int32
// AudioUnitGetPropertyInfoProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfoProc
// AudioUnitGetPropertyInfoProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, unsigned int *, unsigned char *)
type AudioUnitGetPropertyInfoProc = func(unsafe.Pointer, uint32, uint32, uint32, uint32, uint8) int32
// AudioUnitGetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyProc
// AudioUnitGetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, void *, unsigned int *)
type AudioUnitGetPropertyProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer, uint32) int32
// AudioUnitInitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitializeProc
// AudioUnitInitializeProc is a callback function
// C type: int (*)(void *)
type AudioUnitInitializeProc = func(unsafe.Pointer) int32
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
// AudioUnitProcessMultipleProc is a callback function
// C type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, const struct AudioBufferList **, unsigned int, struct AudioBufferList **) __attribute__((nonblocking))
type AudioUnitProcessMultipleProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, uint32, unsafe.Pointer, uint32, AudioBufferList **) __attribute__((nonblocking)) int32
// AudioUnitProcessProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessProc
// AudioUnitProcessProc is a callback function
// C type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type AudioUnitProcessProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, AudioBufferList *) __attribute__((nonblocking)) int32
// AudioUnitPropertyID - The data type for audio unit property keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyID
// AudioUnitPropertyID has base type: UInt32
type AudioUnitPropertyID uintptr
// AudioUnitPropertyListenerProc - Called by the system when the value of a specified audio unit property has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPropertyListenerProc
// AudioUnitPropertyListenerProc is a callback function
// C type: void (*)(void *, struct OpaqueAudioComponentInstance *, unsigned int, unsigned int, unsigned int)
type AudioUnitPropertyListenerProc = func(unsafe.Pointer, unsafe.Pointer, uint32, uint32, uint32)
// AudioUnitRemovePropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerProc
// AudioUnitRemovePropertyListenerProc is a callback function
// C type: int (*)(void *, unsigned int, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int))
type AudioUnitRemovePropertyListenerProc = func(unsafe.Pointer, uint32, void (*)(void *, struct ComponentInstanceRecord *, unsigned int, unsigned int, unsigned int)) int32
// AudioUnitRemovePropertyListenerWithUserDataProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserDataProc
// AudioUnitRemovePropertyListenerWithUserDataProc is a callback function
// C type: int (*)(void *, unsigned int, void (*)(void *, struct OpaqueAudioComponentInstance *, unsigned int, unsigned int, unsigned int), void *)
type AudioUnitRemovePropertyListenerWithUserDataProc = func(unsafe.Pointer, uint32, void (*)(void *, struct OpaqueAudioComponentInstance *, unsigned int, unsigned int, unsigned int), unsafe.Pointer) int32
// AudioUnitRemoveRenderNotifyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotifyProc
// AudioUnitRemoveRenderNotifyProc is a callback function
// C type: int (*)(void *, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), void *)
type AudioUnitRemoveRenderNotifyProc = func(unsafe.Pointer, int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking)), unsafe.Pointer) int32
// AudioUnitRenderProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderProc
// AudioUnitRenderProc is a callback function
// C type: int (*)(void *, enum AudioUnitRenderActionFlags *, const struct AudioTimeStamp *, unsigned int, unsigned int, struct AudioBufferList *) __attribute__((nonblocking))
type AudioUnitRenderProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, uint32, AudioBufferList *) __attribute__((nonblocking)) int32
// AudioUnitResetProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitResetProc
// AudioUnitResetProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int)
type AudioUnitResetProc = func(unsafe.Pointer, uint32, uint32) int32
// AudioUnitScheduleParametersProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParametersProc
// AudioUnitScheduleParametersProc is a callback function
// C type: int (*)(void *, const struct AudioUnitParameterEvent *, unsigned int) __attribute__((nonblocking))
type AudioUnitScheduleParametersProc = func(unsafe.Pointer, unsafe.Pointer, unsigned int) __attribute__((nonblocking)) int32
// AudioUnitScope - The data type for audio unit scope identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScope
// AudioUnitScope has base type: UInt32
type AudioUnitScope uintptr
// AudioUnitSetParameterProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameterProc
// AudioUnitSetParameterProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, float, unsigned int) __attribute__((nonblocking))
type AudioUnitSetParameterProc = func(unsafe.Pointer, uint32, uint32, uint32, float32, unsigned int) __attribute__((nonblocking)) int32
// AudioUnitSetPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetPropertyProc
// AudioUnitSetPropertyProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *, unsigned int)
type AudioUnitSetPropertyProc = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer, uint32) int32
// AudioUnitUninitializeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitializeProc
// AudioUnitUninitializeProc is a callback function
// C type: int (*)(void *)
type AudioUnitUninitializeProc = func(unsafe.Pointer) int32
// ClockBeats type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBeats
// CAClockBeats has base type: Float64
type ClockBeats uintptr
// ClockListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockListenerProc
// CAClockListenerProc is a callback function
// C type: void (*)(void *, enum CAClockMessage, const void *)
type ClockListenerProc = func(unsafe.Pointer, ClockMessage, unsafe.Pointer)
// ClockRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockRef
// CAClockRef has base type: struct OpaqueCAClock *
type ClockRef uintptr
// ClockSMPTEFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTEFormat
// CAClockSMPTEFormat has base type: SMPTETimeType
type ClockSMPTEFormat uintptr
// ClockSamples type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSamples
// CAClockSamples has base type: Float64
type ClockSamples uintptr
// ClockSeconds type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSeconds
// CAClockSeconds has base type: Float64
type ClockSeconds uintptr
// ClockTempo type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTempo
// CAClockTempo has base type: Float64
type ClockTempo uintptr
// CountUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CountUserDataFDF
// CountUserDataFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *)
type CountUserDataFDF = func(unsafe.Pointer, uint32, uint32) int32
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
// GetPropertyFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, void *)
type GetPropertyFDF = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// GetPropertyInfoFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetPropertyInfoFDF
// GetPropertyInfoFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int *, unsigned int *)
type GetPropertyInfoFDF = func(unsafe.Pointer, uint32, uint32, uint32) int32
// GetUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetUserDataFDF
// GetUserDataFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int *, void *)
type GetUserDataFDF = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32
// GetUserDataSizeFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetUserDataSizeFDF
// GetUserDataSizeFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int *)
type GetUserDataSizeFDF = func(unsafe.Pointer, uint32, uint32, uint32) int32
// HostCallback_GetBeatAndTempo - When called by the system, provides beat and tempo information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetBeatAndTempo
// HostCallback_GetBeatAndTempo is a callback function
// C type: int (*)(void *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetBeatAndTempo = func(unsafe.Pointer, float64, double *) __attribute__((nonblocking)) int32
// HostCallback_GetMusicalTimeLocation - When called by the system, provides musical timing information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetMusicalTimeLocation
// HostCallback_GetMusicalTimeLocation is a callback function
// C type: int (*)(void *, unsigned int *, float *, unsigned int *, double *) __attribute__((nonblocking))
type HostCallback_GetMusicalTimeLocation = func(unsafe.Pointer, uint32, float32, uint32, double *) __attribute__((nonblocking)) int32
// HostCallback_GetTransportState - When called by the system, provides audio transport state and timeline information to an audio unit from a host application.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState
// HostCallback_GetTransportState is a callback function
// C type: int (*)(void *, unsigned char *, unsigned char *, double *, unsigned char *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetTransportState = func(unsafe.Pointer, uint8, uint8, float64, uint8, float64, double *) __attribute__((nonblocking)) int32
// HostCallback_GetTransportState2 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallback_GetTransportState2
// HostCallback_GetTransportState2 is a callback function
// C type: int (*)(void *, unsigned char *, unsigned char *, unsigned char *, double *, unsigned char *, double *, double *) __attribute__((nonblocking))
type HostCallback_GetTransportState2 = func(unsafe.Pointer, uint8, uint8, uint8, float64, uint8, float64, double *) __attribute__((nonblocking)) int32
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
// MusicDeviceMIDIEventProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, unsigned int) __attribute__((nonblocking))
type MusicDeviceMIDIEventProc = func(unsafe.Pointer, uint32, uint32, uint32, unsigned int) __attribute__((nonblocking)) int32
// MusicDeviceStartNoteProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNoteProc
// MusicDeviceStartNoteProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int *, unsigned int, const struct MusicDeviceNoteParams *) __attribute__((nonblocking))
type MusicDeviceStartNoteProc = func(unsafe.Pointer, uint32, uint32, uint32, uint32, MusicDeviceNoteParams *) __attribute__((nonblocking)) int32
// MusicDeviceStopNoteProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNoteProc
// MusicDeviceStopNoteProc is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int) __attribute__((nonblocking))
type MusicDeviceStopNoteProc = func(unsafe.Pointer, uint32, uint32, unsigned int) __attribute__((nonblocking)) int32
// MusicDeviceSysExProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysExProc
// MusicDeviceSysExProc is a callback function
// C type: int (*)(void *, const unsigned char *, unsigned int) __attribute__((nonblocking))
type MusicDeviceSysExProc = func(unsafe.Pointer, uint8, unsigned int) __attribute__((nonblocking)) int32
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
// MusicSequenceUserCallback is a callback function
// C type: void (*)(void *, struct OpaqueMusicSequence *, struct OpaqueMusicTrack *, double, const struct MusicEventUserData *, double, double) __attribute__((nonblocking))
type MusicSequenceUserCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float64, unsafe.Pointer, float64, double) __attribute__((nonblocking))
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
// ReadBytesFDF is a callback function
// C type: int (*)(void *, unsigned char, long long, unsigned int *, void *)
type ReadBytesFDF = func(unsafe.Pointer, uint8, int64, uint32, unsafe.Pointer) int32
// ReadPacketDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ReadPacketDataFDF
// ReadPacketDataFDF is a callback function
// C type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type ReadPacketDataFDF = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32
// ReadPacketsFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ReadPacketsFDF
// ReadPacketsFDF is a callback function
// C type: int (*)(void *, unsigned char, unsigned int *, struct AudioStreamPacketDescription *, long long, unsigned int *, void *)
type ReadPacketsFDF = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32
// ScheduledAudioFileRegionCompletionProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegionCompletionProc
// ScheduledAudioFileRegionCompletionProc is a callback function
// C type: void (*)(void *, struct ScheduledAudioFileRegion *, int)
type ScheduledAudioFileRegionCompletionProc = func(unsafe.Pointer, unsafe.Pointer, int32)
// ScheduledAudioSliceCompletionProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSliceCompletionProc
// ScheduledAudioSliceCompletionProc is a callback function
// C type: void (*)(void *, struct ScheduledAudioSlice *) __attribute__((nonblocking))
type ScheduledAudioSliceCompletionProc = func(unsafe.Pointer, ScheduledAudioSlice *) __attribute__((nonblocking))
// SetPropertyFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetPropertyFDF
// SetPropertyFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, const void *)
type SetPropertyFDF = func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
// SetUserDataFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetUserDataFDF
// SetUserDataFDF is a callback function
// C type: int (*)(void *, unsigned int, unsigned int, unsigned int, const void *)
type SetUserDataFDF = func(unsafe.Pointer, uint32, uint32, uint32, unsafe.Pointer) int32
// SystemSoundID - A system sound object, identified with a sound file you want to play.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SystemSoundID
// SystemSoundID has base type: UInt32
type SystemSoundID uintptr
// WriteBytesFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/WriteBytesFDF
// WriteBytesFDF is a callback function
// C type: int (*)(void *, unsigned char, long long, unsigned int *, const void *)
type WriteBytesFDF = func(unsafe.Pointer, uint8, int64, uint32, unsafe.Pointer) int32
// WritePacketsFDF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/WritePacketsFDF
// WritePacketsFDF is a callback function
// C type: int (*)(void *, unsigned char, unsigned int, const struct AudioStreamPacketDescription *, long long, unsigned int *, const void *)
type WritePacketsFDF = func(unsafe.Pointer, uint8, uint32, unsafe.Pointer, int64, uint32, unsafe.Pointer) int32

