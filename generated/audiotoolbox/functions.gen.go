// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AudioToolbox Functions (345 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AudioWorkIntervalCreate func(unsafe.Pointer) unsafe.Pointer
	_AUEventListenerAddEventType func(EventListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerCreate func(EventListenerProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerCreateWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerNotify func(EventListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerRemoveEventType func(EventListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphAddNode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetInteractionInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetNodeInfo func(unsafe.Pointer, Node, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetNodeInteractions func(unsafe.Pointer, Node, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphNewNode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerAddParameter func(ParameterListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerCreate func(ParameterListenerProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerCreateWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerDispose func(ParameterListenerRef) unsafe.Pointer
	_AUListenerRemoveParameter func(ParameterListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterFormatValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterListenerNotify func(ParameterListenerRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterSet func(ParameterListenerRef, unsafe.Pointer, unsafe.Pointer, AudioUnitParameterValue, unsafe.Pointer) unsafe.Pointer
	_AUParameterValueFromLinear func(unsafe.Pointer, unsafe.Pointer) AudioUnitParameterValue
	_AUParameterValueToLinear func(AudioUnitParameterValue, unsafe.Pointer) unsafe.Pointer
	_AudioCodecAppendInputBufferList func(AudioCodec, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecAppendInputData func(AudioCodec, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecGetProperty func(AudioCodec, AudioCodecPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecGetPropertyInfo func(AudioCodec, AudioCodecPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecInitialize func(AudioCodec, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecProduceOutputBufferList func(AudioCodec, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecProduceOutputPackets func(AudioCodec, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecReset func(AudioCodec) unsafe.Pointer
	_AudioCodecSetProperty func(AudioCodec, AudioCodecPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecUninitialize func(AudioCodec) unsafe.Pointer
	_AudioComponentCopyConfigurationInfo func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioComponentCopyIcon func(AudioComponent) unsafe.Pointer
	_AudioComponentCopyName func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioComponentCount func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentFindNext func(AudioComponent, unsafe.Pointer) AudioComponent
	_AudioComponentGetDescription func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioComponentGetIcon func(AudioComponent, float32) unsafe.Pointer
	_AudioComponentGetLastActiveTime func(AudioComponent) unsafe.Pointer
	_AudioComponentGetVersion func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceCanDo func(AudioComponentInstance, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceDispose func(AudioComponentInstance) unsafe.Pointer
	_AudioComponentInstanceGetComponent func(AudioComponentInstance) AudioComponent
	_AudioComponentInstanceNew func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstantiate func(AudioComponent, unsafe.Pointer)
	_AudioComponentRegister func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, AudioComponentFactoryFunction) AudioComponent
	_AudioComponentValidate func(AudioComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentValidateWithResults func(AudioComponent, unsafe.Pointer) unsafe.Pointer
	_AudioConverterConvertBuffer func(AudioConverterRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterConvertComplexBuffer func(AudioConverterRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterDispose func(AudioConverterRef) unsafe.Pointer
	_AudioConverterFillBuffer func(AudioConverterRef, AudioConverterInputDataProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBuffer func(AudioConverterRef, AudioConverterComplexInputDataProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBufferRealtimeSafe func(AudioConverterRef, AudioConverterComplexInputDataProcRealtimeSafe, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBufferWithPacketDependencies func(AudioConverterRef, AudioConverterComplexInputDataProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterGetProperty func(AudioConverterRef, AudioConverterPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterGetPropertyInfo func(AudioConverterRef, AudioConverterPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNew func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNewSpecific func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNewWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterPrepare func(unsafe.Pointer, unsafe.Pointer)
	_AudioConverterReset func(AudioConverterRef) unsafe.Pointer
	_AudioConverterSetProperty func(AudioConverterRef, AudioConverterPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileClose func(AudioFileID) unsafe.Pointer
	_AudioFileComponentCloseFile func(AudioFileComponent) unsafe.Pointer
	_AudioFileComponentCountUserData func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCreate func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCreateURL func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentDataIsThisFormat func(AudioFileComponent, unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentExtensionIsThisFormat func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentFileDataIsThisFormat func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentFileIsThisFormat func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetGlobalInfo func(AudioFileComponent, AudioFileComponentPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetGlobalInfoSize func(AudioFileComponent, AudioFileComponentPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetProperty func(AudioFileComponent, AudioFileComponentPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetPropertyInfo func(AudioFileComponent, AudioFileComponentPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserData func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataAtOffset func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataSize func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataSize64 func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentInitialize func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentInitializeWithCallbacks func(AudioFileComponent, unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOpenFile func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOpenURL func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AudioFileComponentOpenWithCallbacks func(AudioFileComponent, unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc) unsafe.Pointer
	_AudioFileComponentOptimize func(AudioFileComponent) unsafe.Pointer
	_AudioFileComponentReadBytes func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentReadPacketData func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentReadPackets func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentRemoveUserData func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentSetProperty func(AudioFileComponent, AudioFileComponentPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentSetUserData func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentWriteBytes func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentWritePackets func(AudioFileComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCountUserData func(AudioFileID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCreate func(unsafe.Pointer, unsafe.Pointer, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCreateWithURL func(unsafe.Pointer, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetGlobalInfo func(AudioFilePropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetGlobalInfoSize func(AudioFilePropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetProperty func(AudioFileID, AudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetPropertyInfo func(AudioFileID, AudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserData func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataAtOffset func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataSize func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataSize64 func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileInitialize func(unsafe.Pointer, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileInitializeWithCallbacks func(unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpen func(unsafe.Pointer, unsafe.Pointer, AudioFileTypeID, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpenURL func(unsafe.Pointer, unsafe.Pointer, AudioFileTypeID, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpenWithCallbacks func(unsafe.Pointer, AudioFile_ReadProc, AudioFile_WriteProc, AudioFile_GetSizeProc, AudioFile_SetSizeProc, AudioFileTypeID, unsafe.Pointer) unsafe.Pointer
	_AudioFileOptimize func(AudioFileID) unsafe.Pointer
	_AudioFileReadBytes func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileReadPacketData func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileReadPackets func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileRemoveUserData func(AudioFileID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileSetProperty func(AudioFileID, AudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileSetUserData func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamClose func(AudioFileStreamID) unsafe.Pointer
	_AudioFileStreamGetProperty func(AudioFileStreamID, AudioFileStreamPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamGetPropertyInfo func(AudioFileStreamID, AudioFileStreamPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamOpen func(unsafe.Pointer, AudioFileStream_PropertyListenerProc, AudioFileStream_PacketsProc, AudioFileTypeID, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamParseBytes func(AudioFileStreamID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamSeek func(AudioFileStreamID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamSetProperty func(AudioFileStreamID, AudioFileStreamPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWriteBytes func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWritePackets func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWritePacketsWithDependencies func(AudioFileID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFormatGetProperty func(AudioFormatPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFormatGetPropertyInfo func(AudioFormatPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceAddPropertyListener func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceGetPropertyData func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceGetPropertyDataSize func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceHasProperty func(AudioObjectID, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceIsPropertySettable func(AudioObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceRemovePropertyListener func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceSetPropertyData func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioOutputUnitGetHostIcon func(AudioUnit, float32) unsafe.Pointer
	_AudioOutputUnitPublish func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, AudioUnit) unsafe.Pointer
	_AudioOutputUnitStart func(AudioUnit) unsafe.Pointer
	_AudioOutputUnitStop func(AudioUnit) unsafe.Pointer
	_AudioQueueAddPropertyListener func(AudioQueueRef, AudioQueuePropertyID, AudioQueuePropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioQueueAllocateBuffer func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueAllocateBufferWithPacketDescriptions func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueCreateTimeline func(AudioQueueRef, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceGetCurrentTime func(AudioQueueRef, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceGetNearestStartTime func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceTranslateTime func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDispose func(AudioQueueRef, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDisposeTimeline func(AudioQueueRef, AudioQueueTimelineRef) unsafe.Pointer
	_AudioQueueEnqueueBuffer func(AudioQueueRef, AudioQueueBufferRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueEnqueueBufferWithParameters func(AudioQueueRef, AudioQueueBufferRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueFlush func(AudioQueueRef) unsafe.Pointer
	_AudioQueueFreeBuffer func(AudioQueueRef, AudioQueueBufferRef) unsafe.Pointer
	_AudioQueueGetCurrentTime func(AudioQueueRef, AudioQueueTimelineRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetParameter func(AudioQueueRef, AudioQueueParameterID, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetProperty func(AudioQueueRef, AudioQueuePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetPropertySize func(AudioQueueRef, AudioQueuePropertyID, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewInput func(unsafe.Pointer, AudioQueueInputCallback, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewInputWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewOutput func(unsafe.Pointer, AudioQueueOutputCallback, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewOutputWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueOfflineRender func(AudioQueueRef, unsafe.Pointer, AudioQueueBufferRef, unsafe.Pointer) unsafe.Pointer
	_AudioQueuePause func(AudioQueueRef) unsafe.Pointer
	_AudioQueuePrime func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapDispose func(AudioQueueProcessingTapRef) unsafe.Pointer
	_AudioQueueProcessingTapGetQueueTime func(AudioQueueProcessingTapRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapGetSourceAudio func(AudioQueueProcessingTapRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapNew func(AudioQueueRef, AudioQueueProcessingTapCallback, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueRemovePropertyListener func(AudioQueueRef, AudioQueuePropertyID, AudioQueuePropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioQueueReset func(AudioQueueRef) unsafe.Pointer
	_AudioQueueSetOfflineRenderFormat func(AudioQueueRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueSetParameter func(AudioQueueRef, AudioQueueParameterID, AudioQueueParameterValue) unsafe.Pointer
	_AudioQueueSetProperty func(AudioQueueRef, AudioQueuePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueStart func(AudioQueueRef, unsafe.Pointer) unsafe.Pointer
	_AudioQueueStop func(AudioQueueRef, unsafe.Pointer) unsafe.Pointer
	_AudioServicesAddSystemSoundCompletion func(SystemSoundID, unsafe.Pointer, unsafe.Pointer, AudioServicesSystemSoundCompletionProc, unsafe.Pointer) unsafe.Pointer
	_AudioServicesCreateSystemSoundID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesDisposeSystemSoundID func(SystemSoundID) unsafe.Pointer
	_AudioServicesGetProperty func(AudioServicesPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesGetPropertyInfo func(AudioServicesPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlayAlertSound func(SystemSoundID)
	_AudioServicesPlayAlertSoundWithCompletion func(SystemSoundID)
	_AudioServicesPlayAlertSoundWithDetails func(SystemSoundID, unsafe.Pointer)
	_AudioServicesPlaySystemSound func(SystemSoundID)
	_AudioServicesPlaySystemSoundWithCompletion func(SystemSoundID)
	_AudioServicesPlaySystemSoundWithDetails func(SystemSoundID, unsafe.Pointer)
	_AudioServicesRemoveSystemSoundCompletion func(SystemSoundID)
	_AudioServicesSetProperty func(AudioServicesPropertyID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionAddPropertyListener func(AudioSessionPropertyID, AudioSessionPropertyListener, unsafe.Pointer) unsafe.Pointer
	_AudioSessionGetProperty func(AudioSessionPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionGetPropertySize func(AudioSessionPropertyID, unsafe.Pointer) unsafe.Pointer
	_AudioSessionInitialize func(unsafe.Pointer, unsafe.Pointer, AudioSessionInterruptionListener, unsafe.Pointer) unsafe.Pointer
	_AudioSessionRemovePropertyListener func(AudioSessionPropertyID) unsafe.Pointer
	_AudioSessionRemovePropertyListenerWithUserData func(AudioSessionPropertyID, AudioSessionPropertyListener, unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetActive func(unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetActiveWithFlags func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetProperty func(AudioSessionPropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitAddPropertyListener func(AudioUnit, AudioUnitPropertyID, AudioUnitPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioUnitAddRenderNotify func(AudioUnit, RenderCallback, unsafe.Pointer) unsafe.Pointer
	_AudioUnitExtensionCopyComponentList func(unsafe.Pointer) unsafe.Pointer
	_AudioUnitExtensionSetComponentList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetParameter func(AudioUnit, AudioUnitParameterID, AudioUnitScope, AudioUnitElement, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetProperty func(AudioUnit, AudioUnitPropertyID, AudioUnitScope, AudioUnitElement, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetPropertyInfo func(AudioUnit, AudioUnitPropertyID, AudioUnitScope, AudioUnitElement, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitInitialize func(AudioUnit) unsafe.Pointer
	_AudioUnitProcess func(AudioUnit, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitProcessMultiple func(AudioUnit, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRemovePropertyListenerWithUserData func(AudioUnit, AudioUnitPropertyID, AudioUnitPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRemoveRenderNotify func(AudioUnit, RenderCallback, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRender func(AudioUnit, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitReset func(AudioUnit, AudioUnitScope, AudioUnitElement) unsafe.Pointer
	_AudioUnitScheduleParameters func(AudioUnit, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitSetParameter func(AudioUnit, AudioUnitParameterID, AudioUnitScope, AudioUnitElement, AudioUnitParameterValue, unsafe.Pointer) unsafe.Pointer
	_AudioUnitSetProperty func(AudioUnit, AudioUnitPropertyID, AudioUnitScope, AudioUnitElement, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitUninitialize func(AudioUnit) unsafe.Pointer
	_CAClockAddListener func(ClockRef, ClockListenerProc, unsafe.Pointer) unsafe.Pointer
	_CAClockArm func(ClockRef) unsafe.Pointer
	_CAClockBarBeatTimeToBeats func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockBeatsToBarBeatTime func(ClockRef, ClockBeats, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockDisarm func(ClockRef) unsafe.Pointer
	_CAClockDispose func(ClockRef) unsafe.Pointer
	_CAClockGetCurrentTempo func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetCurrentTime func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetPlayRate func(ClockRef, unsafe.Pointer) unsafe.Pointer
	_CAClockGetProperty func(ClockRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetPropertyInfo func(ClockRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetStartTime func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockNew func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockParseMIDI func(ClockRef, unsafe.Pointer) unsafe.Pointer
	_CAClockRemoveListener func(ClockRef, ClockListenerProc, unsafe.Pointer) unsafe.Pointer
	_CAClockSMPTETimeToSeconds func(ClockRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSecondsToSMPTETime func(ClockRef, ClockSeconds, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSetCurrentTempo func(ClockRef, ClockTempo, unsafe.Pointer) unsafe.Pointer
	_CAClockSetCurrentTime func(ClockRef, unsafe.Pointer) unsafe.Pointer
	_CAClockSetPlayRate func(ClockRef, unsafe.Pointer) unsafe.Pointer
	_CAClockSetProperty func(ClockRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockStart func(ClockRef) unsafe.Pointer
	_CAClockStop func(ClockRef) unsafe.Pointer
	_CAClockTranslateTime func(ClockRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAShow func(unsafe.Pointer)
	_CAShowFile func(unsafe.Pointer, unsafe.Pointer)
	_CopyInstrumentInfoFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CopyNameFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DisposeMusicEventIterator func(MusicEventIterator) unsafe.Pointer
	_DisposeMusicPlayer func(MusicPlayer) unsafe.Pointer
	_DisposeMusicSequence func(MusicSequence) unsafe.Pointer
	_ExtAudioFileCreateNew func(unsafe.Pointer, unsafe.Pointer, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileCreateWithURL func(unsafe.Pointer, AudioFileTypeID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileDispose func(ExtAudioFileRef) unsafe.Pointer
	_ExtAudioFileGetProperty func(ExtAudioFileRef, ExtAudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileGetPropertyInfo func(ExtAudioFileRef, ExtAudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileOpen func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileOpenURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileRead func(ExtAudioFileRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileSeek func(ExtAudioFileRef, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileSetProperty func(ExtAudioFileRef, ExtAudioFilePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileTell func(ExtAudioFileRef, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWrapAudioFileID func(AudioFileID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWrite func(ExtAudioFileRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWriteAsync func(ExtAudioFileRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GetNameFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceMIDIEvent func(MusicDeviceComponent, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceMIDIEventList func(MusicDeviceComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDevicePrepareInstrument func(MusicDeviceComponent, MusicDeviceInstrumentID) unsafe.Pointer
	_MusicDeviceReleaseInstrument func(MusicDeviceComponent, MusicDeviceInstrumentID) unsafe.Pointer
	_MusicDeviceStartNote func(MusicDeviceComponent, MusicDeviceInstrumentID, MusicDeviceGroupID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceStopNote func(MusicDeviceComponent, MusicDeviceGroupID, NoteInstanceID, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceSysEx func(MusicDeviceComponent, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorDeleteEvent func(MusicEventIterator) unsafe.Pointer
	_MusicEventIteratorGetEventInfo func(MusicEventIterator, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasCurrentEvent func(MusicEventIterator, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasNextEvent func(MusicEventIterator, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasPreviousEvent func(MusicEventIterator, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorNextEvent func(MusicEventIterator) unsafe.Pointer
	_MusicEventIteratorPreviousEvent func(MusicEventIterator) unsafe.Pointer
	_MusicEventIteratorSeek func(MusicEventIterator, MusicTimeStamp) unsafe.Pointer
	_MusicEventIteratorSetEventInfo func(MusicEventIterator, MusicEventType, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorSetEventTime func(MusicEventIterator, MusicTimeStamp) unsafe.Pointer
	_MusicPlayerGetBeatsForHostTime func(MusicPlayer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetHostTimeForBeats func(MusicPlayer, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetPlayRateScalar func(MusicPlayer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetSequence func(MusicPlayer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetTime func(MusicPlayer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerIsPlaying func(MusicPlayer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerPreroll func(MusicPlayer) unsafe.Pointer
	_MusicPlayerSetPlayRateScalar func(MusicPlayer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerSetSequence func(MusicPlayer, MusicSequence) unsafe.Pointer
	_MusicPlayerSetTime func(MusicPlayer, MusicTimeStamp) unsafe.Pointer
	_MusicPlayerStart func(MusicPlayer) unsafe.Pointer
	_MusicPlayerStop func(MusicPlayer) unsafe.Pointer
	_MusicSequenceBarBeatTimeToBeats func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceBeatsToBarBeatTime func(MusicSequence, MusicTimeStamp, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceDisposeTrack func(MusicSequence, MusicTrack) unsafe.Pointer
	_MusicSequenceFileCreate func(MusicSequence, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileCreateData func(MusicSequence, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileLoad func(MusicSequence, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileLoadData func(MusicSequence, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetAUGraph func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetBeatsForSeconds func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetIndTrack func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetInfoDictionary func(MusicSequence) unsafe.Pointer
	_MusicSequenceGetSecondsForBeats func(MusicSequence, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetSequenceType func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTempoTrack func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTrackCount func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTrackIndex func(MusicSequence, MusicTrack, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceLoadSMFDataWithFlags func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceLoadSMFWithFlags func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceNewTrack func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceReverse func(MusicSequence) unsafe.Pointer
	_MusicSequenceSaveMIDIFile func(MusicSequence, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSaveSMFData func(MusicSequence, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetAUGraph func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetMIDIEndpoint func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetSequenceType func(MusicSequence, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetUserCallback func(MusicSequence, MusicSequenceUserCallback, unsafe.Pointer) unsafe.Pointer
	_MusicTrackClear func(MusicTrack, MusicTimeStamp, MusicTimeStamp) unsafe.Pointer
	_MusicTrackCopyInsert func(MusicTrack, MusicTimeStamp, MusicTimeStamp, MusicTrack, MusicTimeStamp) unsafe.Pointer
	_MusicTrackCut func(MusicTrack, MusicTimeStamp, MusicTimeStamp) unsafe.Pointer
	_MusicTrackGetDestMIDIEndpoint func(MusicTrack, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetDestNode func(MusicTrack, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetProperty func(MusicTrack, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetSequence func(MusicTrack, unsafe.Pointer) unsafe.Pointer
	_MusicTrackMerge func(MusicTrack, MusicTimeStamp, MusicTimeStamp, MusicTrack, MusicTimeStamp) unsafe.Pointer
	_MusicTrackMoveEvents func(MusicTrack, MusicTimeStamp, MusicTimeStamp, MusicTimeStamp) unsafe.Pointer
	_MusicTrackNewAUPresetEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedControlEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedNoteEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedTempoEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDIChannelEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDINoteEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDIRawDataEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMetaEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewParameterEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewUserEvent func(MusicTrack, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_MusicTrackSetDestMIDIEndpoint func(MusicTrack, unsafe.Pointer) unsafe.Pointer
	_MusicTrackSetDestNode func(MusicTrack, Node) unsafe.Pointer
	_MusicTrackSetProperty func(MusicTrack, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NewAUGraph func(unsafe.Pointer) unsafe.Pointer
	_NewMusicEventIterator func(MusicTrack, unsafe.Pointer) unsafe.Pointer
	_NewMusicPlayer func(unsafe.Pointer) unsafe.Pointer
	_NewMusicSequence func(unsafe.Pointer) unsafe.Pointer
	_NewMusicTrackFrom func(MusicTrack, MusicTimeStamp, MusicTimeStamp, unsafe.Pointer) unsafe.Pointer
	_GetAudioUnitParameterDisplayType func(unsafe.Pointer) unsafe.Pointer
	_SetAudioUnitParameterDisplayType func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AudioWorkIntervalCreate, lib, "AudioWorkIntervalCreate")
	tryRegister(&_AUEventListenerAddEventType, lib, "AUEventListenerAddEventType")
	tryRegister(&_AUEventListenerCreate, lib, "AUEventListenerCreate")
	tryRegister(&_AUEventListenerCreateWithDispatchQueue, lib, "AUEventListenerCreateWithDispatchQueue")
	tryRegister(&_AUEventListenerNotify, lib, "AUEventListenerNotify")
	tryRegister(&_AUEventListenerRemoveEventType, lib, "AUEventListenerRemoveEventType")
	tryRegister(&_AUGraphAddNode, lib, "AUGraphAddNode")
	tryRegister(&_AUGraphGetInteractionInfo, lib, "AUGraphGetInteractionInfo")
	tryRegister(&_AUGraphGetNodeInfo, lib, "AUGraphGetNodeInfo")
	tryRegister(&_AUGraphGetNodeInteractions, lib, "AUGraphGetNodeInteractions")
	tryRegister(&_AUGraphNewNode, lib, "AUGraphNewNode")
	tryRegister(&_AUListenerAddParameter, lib, "AUListenerAddParameter")
	tryRegister(&_AUListenerCreate, lib, "AUListenerCreate")
	tryRegister(&_AUListenerCreateWithDispatchQueue, lib, "AUListenerCreateWithDispatchQueue")
	tryRegister(&_AUListenerDispose, lib, "AUListenerDispose")
	tryRegister(&_AUListenerRemoveParameter, lib, "AUListenerRemoveParameter")
	tryRegister(&_AUParameterFormatValue, lib, "AUParameterFormatValue")
	tryRegister(&_AUParameterListenerNotify, lib, "AUParameterListenerNotify")
	tryRegister(&_AUParameterSet, lib, "AUParameterSet")
	tryRegister(&_AUParameterValueFromLinear, lib, "AUParameterValueFromLinear")
	tryRegister(&_AUParameterValueToLinear, lib, "AUParameterValueToLinear")
	tryRegister(&_AudioCodecAppendInputBufferList, lib, "AudioCodecAppendInputBufferList")
	tryRegister(&_AudioCodecAppendInputData, lib, "AudioCodecAppendInputData")
	tryRegister(&_AudioCodecGetProperty, lib, "AudioCodecGetProperty")
	tryRegister(&_AudioCodecGetPropertyInfo, lib, "AudioCodecGetPropertyInfo")
	tryRegister(&_AudioCodecInitialize, lib, "AudioCodecInitialize")
	tryRegister(&_AudioCodecProduceOutputBufferList, lib, "AudioCodecProduceOutputBufferList")
	tryRegister(&_AudioCodecProduceOutputPackets, lib, "AudioCodecProduceOutputPackets")
	tryRegister(&_AudioCodecReset, lib, "AudioCodecReset")
	tryRegister(&_AudioCodecSetProperty, lib, "AudioCodecSetProperty")
	tryRegister(&_AudioCodecUninitialize, lib, "AudioCodecUninitialize")
	tryRegister(&_AudioComponentCopyConfigurationInfo, lib, "AudioComponentCopyConfigurationInfo")
	tryRegister(&_AudioComponentCopyIcon, lib, "AudioComponentCopyIcon")
	tryRegister(&_AudioComponentCopyName, lib, "AudioComponentCopyName")
	tryRegister(&_AudioComponentCount, lib, "AudioComponentCount")
	tryRegister(&_AudioComponentFindNext, lib, "AudioComponentFindNext")
	tryRegister(&_AudioComponentGetDescription, lib, "AudioComponentGetDescription")
	tryRegister(&_AudioComponentGetIcon, lib, "AudioComponentGetIcon")
	tryRegister(&_AudioComponentGetLastActiveTime, lib, "AudioComponentGetLastActiveTime")
	tryRegister(&_AudioComponentGetVersion, lib, "AudioComponentGetVersion")
	tryRegister(&_AudioComponentInstanceCanDo, lib, "AudioComponentInstanceCanDo")
	tryRegister(&_AudioComponentInstanceDispose, lib, "AudioComponentInstanceDispose")
	tryRegister(&_AudioComponentInstanceGetComponent, lib, "AudioComponentInstanceGetComponent")
	tryRegister(&_AudioComponentInstanceNew, lib, "AudioComponentInstanceNew")
	tryRegister(&_AudioComponentInstantiate, lib, "AudioComponentInstantiate")
	tryRegister(&_AudioComponentRegister, lib, "AudioComponentRegister")
	tryRegister(&_AudioComponentValidate, lib, "AudioComponentValidate")
	tryRegister(&_AudioComponentValidateWithResults, lib, "AudioComponentValidateWithResults")
	tryRegister(&_AudioConverterConvertBuffer, lib, "AudioConverterConvertBuffer")
	tryRegister(&_AudioConverterConvertComplexBuffer, lib, "AudioConverterConvertComplexBuffer")
	tryRegister(&_AudioConverterDispose, lib, "AudioConverterDispose")
	tryRegister(&_AudioConverterFillBuffer, lib, "AudioConverterFillBuffer")
	tryRegister(&_AudioConverterFillComplexBuffer, lib, "AudioConverterFillComplexBuffer")
	tryRegister(&_AudioConverterFillComplexBufferRealtimeSafe, lib, "AudioConverterFillComplexBufferRealtimeSafe")
	tryRegister(&_AudioConverterFillComplexBufferWithPacketDependencies, lib, "AudioConverterFillComplexBufferWithPacketDependencies")
	tryRegister(&_AudioConverterGetProperty, lib, "AudioConverterGetProperty")
	tryRegister(&_AudioConverterGetPropertyInfo, lib, "AudioConverterGetPropertyInfo")
	tryRegister(&_AudioConverterNew, lib, "AudioConverterNew")
	tryRegister(&_AudioConverterNewSpecific, lib, "AudioConverterNewSpecific")
	tryRegister(&_AudioConverterNewWithOptions, lib, "AudioConverterNewWithOptions")
	tryRegister(&_AudioConverterPrepare, lib, "AudioConverterPrepare")
	tryRegister(&_AudioConverterReset, lib, "AudioConverterReset")
	tryRegister(&_AudioConverterSetProperty, lib, "AudioConverterSetProperty")
	tryRegister(&_AudioFileClose, lib, "AudioFileClose")
	tryRegister(&_AudioFileComponentCloseFile, lib, "AudioFileComponentCloseFile")
	tryRegister(&_AudioFileComponentCountUserData, lib, "AudioFileComponentCountUserData")
	tryRegister(&_AudioFileComponentCreate, lib, "AudioFileComponentCreate")
	tryRegister(&_AudioFileComponentCreateURL, lib, "AudioFileComponentCreateURL")
	tryRegister(&_AudioFileComponentDataIsThisFormat, lib, "AudioFileComponentDataIsThisFormat")
	tryRegister(&_AudioFileComponentExtensionIsThisFormat, lib, "AudioFileComponentExtensionIsThisFormat")
	tryRegister(&_AudioFileComponentFileDataIsThisFormat, lib, "AudioFileComponentFileDataIsThisFormat")
	tryRegister(&_AudioFileComponentFileIsThisFormat, lib, "AudioFileComponentFileIsThisFormat")
	tryRegister(&_AudioFileComponentGetGlobalInfo, lib, "AudioFileComponentGetGlobalInfo")
	tryRegister(&_AudioFileComponentGetGlobalInfoSize, lib, "AudioFileComponentGetGlobalInfoSize")
	tryRegister(&_AudioFileComponentGetProperty, lib, "AudioFileComponentGetProperty")
	tryRegister(&_AudioFileComponentGetPropertyInfo, lib, "AudioFileComponentGetPropertyInfo")
	tryRegister(&_AudioFileComponentGetUserData, lib, "AudioFileComponentGetUserData")
	tryRegister(&_AudioFileComponentGetUserDataAtOffset, lib, "AudioFileComponentGetUserDataAtOffset")
	tryRegister(&_AudioFileComponentGetUserDataSize, lib, "AudioFileComponentGetUserDataSize")
	tryRegister(&_AudioFileComponentGetUserDataSize64, lib, "AudioFileComponentGetUserDataSize64")
	tryRegister(&_AudioFileComponentInitialize, lib, "AudioFileComponentInitialize")
	tryRegister(&_AudioFileComponentInitializeWithCallbacks, lib, "AudioFileComponentInitializeWithCallbacks")
	tryRegister(&_AudioFileComponentOpenFile, lib, "AudioFileComponentOpenFile")
	tryRegister(&_AudioFileComponentOpenURL, lib, "AudioFileComponentOpenURL")
	tryRegister(&_AudioFileComponentOpenWithCallbacks, lib, "AudioFileComponentOpenWithCallbacks")
	tryRegister(&_AudioFileComponentOptimize, lib, "AudioFileComponentOptimize")
	tryRegister(&_AudioFileComponentReadBytes, lib, "AudioFileComponentReadBytes")
	tryRegister(&_AudioFileComponentReadPacketData, lib, "AudioFileComponentReadPacketData")
	tryRegister(&_AudioFileComponentReadPackets, lib, "AudioFileComponentReadPackets")
	tryRegister(&_AudioFileComponentRemoveUserData, lib, "AudioFileComponentRemoveUserData")
	tryRegister(&_AudioFileComponentSetProperty, lib, "AudioFileComponentSetProperty")
	tryRegister(&_AudioFileComponentSetUserData, lib, "AudioFileComponentSetUserData")
	tryRegister(&_AudioFileComponentWriteBytes, lib, "AudioFileComponentWriteBytes")
	tryRegister(&_AudioFileComponentWritePackets, lib, "AudioFileComponentWritePackets")
	tryRegister(&_AudioFileCountUserData, lib, "AudioFileCountUserData")
	tryRegister(&_AudioFileCreate, lib, "AudioFileCreate")
	tryRegister(&_AudioFileCreateWithURL, lib, "AudioFileCreateWithURL")
	tryRegister(&_AudioFileGetGlobalInfo, lib, "AudioFileGetGlobalInfo")
	tryRegister(&_AudioFileGetGlobalInfoSize, lib, "AudioFileGetGlobalInfoSize")
	tryRegister(&_AudioFileGetProperty, lib, "AudioFileGetProperty")
	tryRegister(&_AudioFileGetPropertyInfo, lib, "AudioFileGetPropertyInfo")
	tryRegister(&_AudioFileGetUserData, lib, "AudioFileGetUserData")
	tryRegister(&_AudioFileGetUserDataAtOffset, lib, "AudioFileGetUserDataAtOffset")
	tryRegister(&_AudioFileGetUserDataSize, lib, "AudioFileGetUserDataSize")
	tryRegister(&_AudioFileGetUserDataSize64, lib, "AudioFileGetUserDataSize64")
	tryRegister(&_AudioFileInitialize, lib, "AudioFileInitialize")
	tryRegister(&_AudioFileInitializeWithCallbacks, lib, "AudioFileInitializeWithCallbacks")
	tryRegister(&_AudioFileOpen, lib, "AudioFileOpen")
	tryRegister(&_AudioFileOpenURL, lib, "AudioFileOpenURL")
	tryRegister(&_AudioFileOpenWithCallbacks, lib, "AudioFileOpenWithCallbacks")
	tryRegister(&_AudioFileOptimize, lib, "AudioFileOptimize")
	tryRegister(&_AudioFileReadBytes, lib, "AudioFileReadBytes")
	tryRegister(&_AudioFileReadPacketData, lib, "AudioFileReadPacketData")
	tryRegister(&_AudioFileReadPackets, lib, "AudioFileReadPackets")
	tryRegister(&_AudioFileRemoveUserData, lib, "AudioFileRemoveUserData")
	tryRegister(&_AudioFileSetProperty, lib, "AudioFileSetProperty")
	tryRegister(&_AudioFileSetUserData, lib, "AudioFileSetUserData")
	tryRegister(&_AudioFileStreamClose, lib, "AudioFileStreamClose")
	tryRegister(&_AudioFileStreamGetProperty, lib, "AudioFileStreamGetProperty")
	tryRegister(&_AudioFileStreamGetPropertyInfo, lib, "AudioFileStreamGetPropertyInfo")
	tryRegister(&_AudioFileStreamOpen, lib, "AudioFileStreamOpen")
	tryRegister(&_AudioFileStreamParseBytes, lib, "AudioFileStreamParseBytes")
	tryRegister(&_AudioFileStreamSeek, lib, "AudioFileStreamSeek")
	tryRegister(&_AudioFileStreamSetProperty, lib, "AudioFileStreamSetProperty")
	tryRegister(&_AudioFileWriteBytes, lib, "AudioFileWriteBytes")
	tryRegister(&_AudioFileWritePackets, lib, "AudioFileWritePackets")
	tryRegister(&_AudioFileWritePacketsWithDependencies, lib, "AudioFileWritePacketsWithDependencies")
	tryRegister(&_AudioFormatGetProperty, lib, "AudioFormatGetProperty")
	tryRegister(&_AudioFormatGetPropertyInfo, lib, "AudioFormatGetPropertyInfo")
	tryRegister(&_AudioHardwareServiceAddPropertyListener, lib, "AudioHardwareServiceAddPropertyListener")
	tryRegister(&_AudioHardwareServiceGetPropertyData, lib, "AudioHardwareServiceGetPropertyData")
	tryRegister(&_AudioHardwareServiceGetPropertyDataSize, lib, "AudioHardwareServiceGetPropertyDataSize")
	tryRegister(&_AudioHardwareServiceHasProperty, lib, "AudioHardwareServiceHasProperty")
	tryRegister(&_AudioHardwareServiceIsPropertySettable, lib, "AudioHardwareServiceIsPropertySettable")
	tryRegister(&_AudioHardwareServiceRemovePropertyListener, lib, "AudioHardwareServiceRemovePropertyListener")
	tryRegister(&_AudioHardwareServiceSetPropertyData, lib, "AudioHardwareServiceSetPropertyData")
	tryRegister(&_AudioOutputUnitGetHostIcon, lib, "AudioOutputUnitGetHostIcon")
	tryRegister(&_AudioOutputUnitPublish, lib, "AudioOutputUnitPublish")
	tryRegister(&_AudioOutputUnitStart, lib, "AudioOutputUnitStart")
	tryRegister(&_AudioOutputUnitStop, lib, "AudioOutputUnitStop")
	tryRegister(&_AudioQueueAddPropertyListener, lib, "AudioQueueAddPropertyListener")
	tryRegister(&_AudioQueueAllocateBuffer, lib, "AudioQueueAllocateBuffer")
	tryRegister(&_AudioQueueAllocateBufferWithPacketDescriptions, lib, "AudioQueueAllocateBufferWithPacketDescriptions")
	tryRegister(&_AudioQueueCreateTimeline, lib, "AudioQueueCreateTimeline")
	tryRegister(&_AudioQueueDeviceGetCurrentTime, lib, "AudioQueueDeviceGetCurrentTime")
	tryRegister(&_AudioQueueDeviceGetNearestStartTime, lib, "AudioQueueDeviceGetNearestStartTime")
	tryRegister(&_AudioQueueDeviceTranslateTime, lib, "AudioQueueDeviceTranslateTime")
	tryRegister(&_AudioQueueDispose, lib, "AudioQueueDispose")
	tryRegister(&_AudioQueueDisposeTimeline, lib, "AudioQueueDisposeTimeline")
	tryRegister(&_AudioQueueEnqueueBuffer, lib, "AudioQueueEnqueueBuffer")
	tryRegister(&_AudioQueueEnqueueBufferWithParameters, lib, "AudioQueueEnqueueBufferWithParameters")
	tryRegister(&_AudioQueueFlush, lib, "AudioQueueFlush")
	tryRegister(&_AudioQueueFreeBuffer, lib, "AudioQueueFreeBuffer")
	tryRegister(&_AudioQueueGetCurrentTime, lib, "AudioQueueGetCurrentTime")
	tryRegister(&_AudioQueueGetParameter, lib, "AudioQueueGetParameter")
	tryRegister(&_AudioQueueGetProperty, lib, "AudioQueueGetProperty")
	tryRegister(&_AudioQueueGetPropertySize, lib, "AudioQueueGetPropertySize")
	tryRegister(&_AudioQueueNewInput, lib, "AudioQueueNewInput")
	tryRegister(&_AudioQueueNewInputWithDispatchQueue, lib, "AudioQueueNewInputWithDispatchQueue")
	tryRegister(&_AudioQueueNewOutput, lib, "AudioQueueNewOutput")
	tryRegister(&_AudioQueueNewOutputWithDispatchQueue, lib, "AudioQueueNewOutputWithDispatchQueue")
	tryRegister(&_AudioQueueOfflineRender, lib, "AudioQueueOfflineRender")
	tryRegister(&_AudioQueuePause, lib, "AudioQueuePause")
	tryRegister(&_AudioQueuePrime, lib, "AudioQueuePrime")
	tryRegister(&_AudioQueueProcessingTapDispose, lib, "AudioQueueProcessingTapDispose")
	tryRegister(&_AudioQueueProcessingTapGetQueueTime, lib, "AudioQueueProcessingTapGetQueueTime")
	tryRegister(&_AudioQueueProcessingTapGetSourceAudio, lib, "AudioQueueProcessingTapGetSourceAudio")
	tryRegister(&_AudioQueueProcessingTapNew, lib, "AudioQueueProcessingTapNew")
	tryRegister(&_AudioQueueRemovePropertyListener, lib, "AudioQueueRemovePropertyListener")
	tryRegister(&_AudioQueueReset, lib, "AudioQueueReset")
	tryRegister(&_AudioQueueSetOfflineRenderFormat, lib, "AudioQueueSetOfflineRenderFormat")
	tryRegister(&_AudioQueueSetParameter, lib, "AudioQueueSetParameter")
	tryRegister(&_AudioQueueSetProperty, lib, "AudioQueueSetProperty")
	tryRegister(&_AudioQueueStart, lib, "AudioQueueStart")
	tryRegister(&_AudioQueueStop, lib, "AudioQueueStop")
	tryRegister(&_AudioServicesAddSystemSoundCompletion, lib, "AudioServicesAddSystemSoundCompletion")
	tryRegister(&_AudioServicesCreateSystemSoundID, lib, "AudioServicesCreateSystemSoundID")
	tryRegister(&_AudioServicesDisposeSystemSoundID, lib, "AudioServicesDisposeSystemSoundID")
	tryRegister(&_AudioServicesGetProperty, lib, "AudioServicesGetProperty")
	tryRegister(&_AudioServicesGetPropertyInfo, lib, "AudioServicesGetPropertyInfo")
	tryRegister(&_AudioServicesPlayAlertSound, lib, "AudioServicesPlayAlertSound")
	tryRegister(&_AudioServicesPlayAlertSoundWithCompletion, lib, "AudioServicesPlayAlertSoundWithCompletion")
	tryRegister(&_AudioServicesPlayAlertSoundWithDetails, lib, "AudioServicesPlayAlertSoundWithDetails")
	tryRegister(&_AudioServicesPlaySystemSound, lib, "AudioServicesPlaySystemSound")
	tryRegister(&_AudioServicesPlaySystemSoundWithCompletion, lib, "AudioServicesPlaySystemSoundWithCompletion")
	tryRegister(&_AudioServicesPlaySystemSoundWithDetails, lib, "AudioServicesPlaySystemSoundWithDetails")
	tryRegister(&_AudioServicesRemoveSystemSoundCompletion, lib, "AudioServicesRemoveSystemSoundCompletion")
	tryRegister(&_AudioServicesSetProperty, lib, "AudioServicesSetProperty")
	tryRegister(&_AudioSessionAddPropertyListener, lib, "AudioSessionAddPropertyListener")
	tryRegister(&_AudioSessionGetProperty, lib, "AudioSessionGetProperty")
	tryRegister(&_AudioSessionGetPropertySize, lib, "AudioSessionGetPropertySize")
	tryRegister(&_AudioSessionInitialize, lib, "AudioSessionInitialize")
	tryRegister(&_AudioSessionRemovePropertyListener, lib, "AudioSessionRemovePropertyListener")
	tryRegister(&_AudioSessionRemovePropertyListenerWithUserData, lib, "AudioSessionRemovePropertyListenerWithUserData")
	tryRegister(&_AudioSessionSetActive, lib, "AudioSessionSetActive")
	tryRegister(&_AudioSessionSetActiveWithFlags, lib, "AudioSessionSetActiveWithFlags")
	tryRegister(&_AudioSessionSetProperty, lib, "AudioSessionSetProperty")
	tryRegister(&_AudioUnitAddPropertyListener, lib, "AudioUnitAddPropertyListener")
	tryRegister(&_AudioUnitAddRenderNotify, lib, "AudioUnitAddRenderNotify")
	tryRegister(&_AudioUnitExtensionCopyComponentList, lib, "AudioUnitExtensionCopyComponentList")
	tryRegister(&_AudioUnitExtensionSetComponentList, lib, "AudioUnitExtensionSetComponentList")
	tryRegister(&_AudioUnitGetParameter, lib, "AudioUnitGetParameter")
	tryRegister(&_AudioUnitGetProperty, lib, "AudioUnitGetProperty")
	tryRegister(&_AudioUnitGetPropertyInfo, lib, "AudioUnitGetPropertyInfo")
	tryRegister(&_AudioUnitInitialize, lib, "AudioUnitInitialize")
	tryRegister(&_AudioUnitProcess, lib, "AudioUnitProcess")
	tryRegister(&_AudioUnitProcessMultiple, lib, "AudioUnitProcessMultiple")
	tryRegister(&_AudioUnitRemovePropertyListenerWithUserData, lib, "AudioUnitRemovePropertyListenerWithUserData")
	tryRegister(&_AudioUnitRemoveRenderNotify, lib, "AudioUnitRemoveRenderNotify")
	tryRegister(&_AudioUnitRender, lib, "AudioUnitRender")
	tryRegister(&_AudioUnitReset, lib, "AudioUnitReset")
	tryRegister(&_AudioUnitScheduleParameters, lib, "AudioUnitScheduleParameters")
	tryRegister(&_AudioUnitSetParameter, lib, "AudioUnitSetParameter")
	tryRegister(&_AudioUnitSetProperty, lib, "AudioUnitSetProperty")
	tryRegister(&_AudioUnitUninitialize, lib, "AudioUnitUninitialize")
	tryRegister(&_CAClockAddListener, lib, "CAClockAddListener")
	tryRegister(&_CAClockArm, lib, "CAClockArm")
	tryRegister(&_CAClockBarBeatTimeToBeats, lib, "CAClockBarBeatTimeToBeats")
	tryRegister(&_CAClockBeatsToBarBeatTime, lib, "CAClockBeatsToBarBeatTime")
	tryRegister(&_CAClockDisarm, lib, "CAClockDisarm")
	tryRegister(&_CAClockDispose, lib, "CAClockDispose")
	tryRegister(&_CAClockGetCurrentTempo, lib, "CAClockGetCurrentTempo")
	tryRegister(&_CAClockGetCurrentTime, lib, "CAClockGetCurrentTime")
	tryRegister(&_CAClockGetPlayRate, lib, "CAClockGetPlayRate")
	tryRegister(&_CAClockGetProperty, lib, "CAClockGetProperty")
	tryRegister(&_CAClockGetPropertyInfo, lib, "CAClockGetPropertyInfo")
	tryRegister(&_CAClockGetStartTime, lib, "CAClockGetStartTime")
	tryRegister(&_CAClockNew, lib, "CAClockNew")
	tryRegister(&_CAClockParseMIDI, lib, "CAClockParseMIDI")
	tryRegister(&_CAClockRemoveListener, lib, "CAClockRemoveListener")
	tryRegister(&_CAClockSMPTETimeToSeconds, lib, "CAClockSMPTETimeToSeconds")
	tryRegister(&_CAClockSecondsToSMPTETime, lib, "CAClockSecondsToSMPTETime")
	tryRegister(&_CAClockSetCurrentTempo, lib, "CAClockSetCurrentTempo")
	tryRegister(&_CAClockSetCurrentTime, lib, "CAClockSetCurrentTime")
	tryRegister(&_CAClockSetPlayRate, lib, "CAClockSetPlayRate")
	tryRegister(&_CAClockSetProperty, lib, "CAClockSetProperty")
	tryRegister(&_CAClockStart, lib, "CAClockStart")
	tryRegister(&_CAClockStop, lib, "CAClockStop")
	tryRegister(&_CAClockTranslateTime, lib, "CAClockTranslateTime")
	tryRegister(&_CAShow, lib, "CAShow")
	tryRegister(&_CAShowFile, lib, "CAShowFile")
	tryRegister(&_CopyInstrumentInfoFromSoundBank, lib, "CopyInstrumentInfoFromSoundBank")
	tryRegister(&_CopyNameFromSoundBank, lib, "CopyNameFromSoundBank")
	tryRegister(&_DisposeMusicEventIterator, lib, "DisposeMusicEventIterator")
	tryRegister(&_DisposeMusicPlayer, lib, "DisposeMusicPlayer")
	tryRegister(&_DisposeMusicSequence, lib, "DisposeMusicSequence")
	tryRegister(&_ExtAudioFileCreateNew, lib, "ExtAudioFileCreateNew")
	tryRegister(&_ExtAudioFileCreateWithURL, lib, "ExtAudioFileCreateWithURL")
	tryRegister(&_ExtAudioFileDispose, lib, "ExtAudioFileDispose")
	tryRegister(&_ExtAudioFileGetProperty, lib, "ExtAudioFileGetProperty")
	tryRegister(&_ExtAudioFileGetPropertyInfo, lib, "ExtAudioFileGetPropertyInfo")
	tryRegister(&_ExtAudioFileOpen, lib, "ExtAudioFileOpen")
	tryRegister(&_ExtAudioFileOpenURL, lib, "ExtAudioFileOpenURL")
	tryRegister(&_ExtAudioFileRead, lib, "ExtAudioFileRead")
	tryRegister(&_ExtAudioFileSeek, lib, "ExtAudioFileSeek")
	tryRegister(&_ExtAudioFileSetProperty, lib, "ExtAudioFileSetProperty")
	tryRegister(&_ExtAudioFileTell, lib, "ExtAudioFileTell")
	tryRegister(&_ExtAudioFileWrapAudioFileID, lib, "ExtAudioFileWrapAudioFileID")
	tryRegister(&_ExtAudioFileWrite, lib, "ExtAudioFileWrite")
	tryRegister(&_ExtAudioFileWriteAsync, lib, "ExtAudioFileWriteAsync")
	tryRegister(&_GetNameFromSoundBank, lib, "GetNameFromSoundBank")
	tryRegister(&_MusicDeviceMIDIEvent, lib, "MusicDeviceMIDIEvent")
	tryRegister(&_MusicDeviceMIDIEventList, lib, "MusicDeviceMIDIEventList")
	tryRegister(&_MusicDevicePrepareInstrument, lib, "MusicDevicePrepareInstrument")
	tryRegister(&_MusicDeviceReleaseInstrument, lib, "MusicDeviceReleaseInstrument")
	tryRegister(&_MusicDeviceStartNote, lib, "MusicDeviceStartNote")
	tryRegister(&_MusicDeviceStopNote, lib, "MusicDeviceStopNote")
	tryRegister(&_MusicDeviceSysEx, lib, "MusicDeviceSysEx")
	tryRegister(&_MusicEventIteratorDeleteEvent, lib, "MusicEventIteratorDeleteEvent")
	tryRegister(&_MusicEventIteratorGetEventInfo, lib, "MusicEventIteratorGetEventInfo")
	tryRegister(&_MusicEventIteratorHasCurrentEvent, lib, "MusicEventIteratorHasCurrentEvent")
	tryRegister(&_MusicEventIteratorHasNextEvent, lib, "MusicEventIteratorHasNextEvent")
	tryRegister(&_MusicEventIteratorHasPreviousEvent, lib, "MusicEventIteratorHasPreviousEvent")
	tryRegister(&_MusicEventIteratorNextEvent, lib, "MusicEventIteratorNextEvent")
	tryRegister(&_MusicEventIteratorPreviousEvent, lib, "MusicEventIteratorPreviousEvent")
	tryRegister(&_MusicEventIteratorSeek, lib, "MusicEventIteratorSeek")
	tryRegister(&_MusicEventIteratorSetEventInfo, lib, "MusicEventIteratorSetEventInfo")
	tryRegister(&_MusicEventIteratorSetEventTime, lib, "MusicEventIteratorSetEventTime")
	tryRegister(&_MusicPlayerGetBeatsForHostTime, lib, "MusicPlayerGetBeatsForHostTime")
	tryRegister(&_MusicPlayerGetHostTimeForBeats, lib, "MusicPlayerGetHostTimeForBeats")
	tryRegister(&_MusicPlayerGetPlayRateScalar, lib, "MusicPlayerGetPlayRateScalar")
	tryRegister(&_MusicPlayerGetSequence, lib, "MusicPlayerGetSequence")
	tryRegister(&_MusicPlayerGetTime, lib, "MusicPlayerGetTime")
	tryRegister(&_MusicPlayerIsPlaying, lib, "MusicPlayerIsPlaying")
	tryRegister(&_MusicPlayerPreroll, lib, "MusicPlayerPreroll")
	tryRegister(&_MusicPlayerSetPlayRateScalar, lib, "MusicPlayerSetPlayRateScalar")
	tryRegister(&_MusicPlayerSetSequence, lib, "MusicPlayerSetSequence")
	tryRegister(&_MusicPlayerSetTime, lib, "MusicPlayerSetTime")
	tryRegister(&_MusicPlayerStart, lib, "MusicPlayerStart")
	tryRegister(&_MusicPlayerStop, lib, "MusicPlayerStop")
	tryRegister(&_MusicSequenceBarBeatTimeToBeats, lib, "MusicSequenceBarBeatTimeToBeats")
	tryRegister(&_MusicSequenceBeatsToBarBeatTime, lib, "MusicSequenceBeatsToBarBeatTime")
	tryRegister(&_MusicSequenceDisposeTrack, lib, "MusicSequenceDisposeTrack")
	tryRegister(&_MusicSequenceFileCreate, lib, "MusicSequenceFileCreate")
	tryRegister(&_MusicSequenceFileCreateData, lib, "MusicSequenceFileCreateData")
	tryRegister(&_MusicSequenceFileLoad, lib, "MusicSequenceFileLoad")
	tryRegister(&_MusicSequenceFileLoadData, lib, "MusicSequenceFileLoadData")
	tryRegister(&_MusicSequenceGetAUGraph, lib, "MusicSequenceGetAUGraph")
	tryRegister(&_MusicSequenceGetBeatsForSeconds, lib, "MusicSequenceGetBeatsForSeconds")
	tryRegister(&_MusicSequenceGetIndTrack, lib, "MusicSequenceGetIndTrack")
	tryRegister(&_MusicSequenceGetInfoDictionary, lib, "MusicSequenceGetInfoDictionary")
	tryRegister(&_MusicSequenceGetSecondsForBeats, lib, "MusicSequenceGetSecondsForBeats")
	tryRegister(&_MusicSequenceGetSequenceType, lib, "MusicSequenceGetSequenceType")
	tryRegister(&_MusicSequenceGetTempoTrack, lib, "MusicSequenceGetTempoTrack")
	tryRegister(&_MusicSequenceGetTrackCount, lib, "MusicSequenceGetTrackCount")
	tryRegister(&_MusicSequenceGetTrackIndex, lib, "MusicSequenceGetTrackIndex")
	tryRegister(&_MusicSequenceLoadSMFDataWithFlags, lib, "MusicSequenceLoadSMFDataWithFlags")
	tryRegister(&_MusicSequenceLoadSMFWithFlags, lib, "MusicSequenceLoadSMFWithFlags")
	tryRegister(&_MusicSequenceNewTrack, lib, "MusicSequenceNewTrack")
	tryRegister(&_MusicSequenceReverse, lib, "MusicSequenceReverse")
	tryRegister(&_MusicSequenceSaveMIDIFile, lib, "MusicSequenceSaveMIDIFile")
	tryRegister(&_MusicSequenceSaveSMFData, lib, "MusicSequenceSaveSMFData")
	tryRegister(&_MusicSequenceSetAUGraph, lib, "MusicSequenceSetAUGraph")
	tryRegister(&_MusicSequenceSetMIDIEndpoint, lib, "MusicSequenceSetMIDIEndpoint")
	tryRegister(&_MusicSequenceSetSequenceType, lib, "MusicSequenceSetSequenceType")
	tryRegister(&_MusicSequenceSetUserCallback, lib, "MusicSequenceSetUserCallback")
	tryRegister(&_MusicTrackClear, lib, "MusicTrackClear")
	tryRegister(&_MusicTrackCopyInsert, lib, "MusicTrackCopyInsert")
	tryRegister(&_MusicTrackCut, lib, "MusicTrackCut")
	tryRegister(&_MusicTrackGetDestMIDIEndpoint, lib, "MusicTrackGetDestMIDIEndpoint")
	tryRegister(&_MusicTrackGetDestNode, lib, "MusicTrackGetDestNode")
	tryRegister(&_MusicTrackGetProperty, lib, "MusicTrackGetProperty")
	tryRegister(&_MusicTrackGetSequence, lib, "MusicTrackGetSequence")
	tryRegister(&_MusicTrackMerge, lib, "MusicTrackMerge")
	tryRegister(&_MusicTrackMoveEvents, lib, "MusicTrackMoveEvents")
	tryRegister(&_MusicTrackNewAUPresetEvent, lib, "MusicTrackNewAUPresetEvent")
	tryRegister(&_MusicTrackNewExtendedControlEvent, lib, "MusicTrackNewExtendedControlEvent")
	tryRegister(&_MusicTrackNewExtendedNoteEvent, lib, "MusicTrackNewExtendedNoteEvent")
	tryRegister(&_MusicTrackNewExtendedTempoEvent, lib, "MusicTrackNewExtendedTempoEvent")
	tryRegister(&_MusicTrackNewMIDIChannelEvent, lib, "MusicTrackNewMIDIChannelEvent")
	tryRegister(&_MusicTrackNewMIDINoteEvent, lib, "MusicTrackNewMIDINoteEvent")
	tryRegister(&_MusicTrackNewMIDIRawDataEvent, lib, "MusicTrackNewMIDIRawDataEvent")
	tryRegister(&_MusicTrackNewMetaEvent, lib, "MusicTrackNewMetaEvent")
	tryRegister(&_MusicTrackNewParameterEvent, lib, "MusicTrackNewParameterEvent")
	tryRegister(&_MusicTrackNewUserEvent, lib, "MusicTrackNewUserEvent")
	tryRegister(&_MusicTrackSetDestMIDIEndpoint, lib, "MusicTrackSetDestMIDIEndpoint")
	tryRegister(&_MusicTrackSetDestNode, lib, "MusicTrackSetDestNode")
	tryRegister(&_MusicTrackSetProperty, lib, "MusicTrackSetProperty")
	tryRegister(&_NewAUGraph, lib, "NewAUGraph")
	tryRegister(&_NewMusicEventIterator, lib, "NewMusicEventIterator")
	tryRegister(&_NewMusicPlayer, lib, "NewMusicPlayer")
	tryRegister(&_NewMusicSequence, lib, "NewMusicSequence")
	tryRegister(&_NewMusicTrackFrom, lib, "NewMusicTrackFrom")
	tryRegister(&_GetAudioUnitParameterDisplayType, lib, "GetAudioUnitParameterDisplayType")
	tryRegister(&_SetAudioUnitParameterDisplayType, lib, "SetAudioUnitParameterDisplayType")
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



// Creates a new interval workgroup for managing real-time audio threads.
//
// Added in macOS 11.0.
// Creates a new interval workgroup for managing real-time audio threads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/3547073-audioworkintervalcreate
func AudioWorkIntervalCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _AudioWorkIntervalCreate(p0)
}

// AUEventListenerAddEventType is a AudioToolbox function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerAddEventType(_:_:_:)
func AUEventListenerAddEventType(inListener EventListenerRef, inObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerAddEventType(inListener, inObject, inEvent)
}

// AUEventListenerCreate is a AudioToolbox function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreate(_:_:_:_:_:_:_:)
func AUEventListenerCreate(inProc EventListenerProc, inUserData unsafe.Pointer, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inNotificationInterval unsafe.Pointer, inValueChangeGranularity unsafe.Pointer, outListener unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, inValueChangeGranularity, outListener)
}

// AUEventListenerCreateWithDispatchQueue is a AudioToolbox function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreateWithDispatchQueue(_:_:_:_:_:)
func AUEventListenerCreateWithDispatchQueue(outListener unsafe.Pointer, inNotificationInterval unsafe.Pointer, inValueChangeGranularity unsafe.Pointer, inDispatchQueue unsafe.Pointer, inBlock unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inValueChangeGranularity, inDispatchQueue, inBlock)
}

// AUEventListenerNotify is a AudioToolbox function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerNotify(_:_:_:)
func AUEventListenerNotify(inSendingListener EventListenerRef, inSendingObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerNotify(inSendingListener, inSendingObject, inEvent)
}

// AUEventListenerRemoveEventType is a AudioToolbox function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRemoveEventType(_:_:_:)
func AUEventListenerRemoveEventType(inListener EventListenerRef, inObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerRemoveEventType(inListener, inObject, inEvent)
}

// Adds a node to an audio processing graph.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Adds a node to an audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphAddNode(_:_:_:)
func AUGraphAddNode(inGraph unsafe.Pointer, inDescription unsafe.Pointer, outNode unsafe.Pointer) unsafe.Pointer {
	return _AUGraphAddNode(inGraph, inDescription, outNode)
}

// Retrieves information about a particular interaction in an audio processing graph.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Retrieves information about a particular interaction in an audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetInteractionInfo(_:_:_:)
func AUGraphGetInteractionInfo(inGraph unsafe.Pointer, inInteractionIndex unsafe.Pointer, outInteraction unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetInteractionInfo(inGraph, inInteractionIndex, outInteraction)
}

// Deprecated in OS X v10.5. Instead, use .
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Deprecated in OS X v10.5. Instead, use .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInfo
func AUGraphGetNodeInfo(inGraph unsafe.Pointer, inNode Node, outDescription unsafe.Pointer, outClassDataSize unsafe.Pointer, outClassData unsafe.Pointer, outAudioUnit unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetNodeInfo(inGraph, inNode, outDescription, outClassDataSize, outClassData, outAudioUnit)
}

// Retrieves information about the interactions in an audio processing graph for a given node.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Retrieves information about the interactions in an audio processing graph for a given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInteractions(_:_:_:_:)
func AUGraphGetNodeInteractions(inGraph unsafe.Pointer, inNode Node, ioNumInteractions unsafe.Pointer, outInteractions unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetNodeInteractions(inGraph, inNode, ioNumInteractions, outInteractions)
}

// Deprecated in OS X v10.5. Instead, use .
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Deprecated in OS X v10.5. Instead, use .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphNewNode
func AUGraphNewNode(inGraph unsafe.Pointer, inDescription unsafe.Pointer, inClassDataSize unsafe.Pointer, inClassData unsafe.Pointer, outNode unsafe.Pointer) unsafe.Pointer {
	return _AUGraphNewNode(inGraph, inDescription, inClassDataSize, inClassData, outNode)
}

// AUListenerAddParameter is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerAddParameter(_:_:_:)
func AUListenerAddParameter(inListener ParameterListenerRef, inObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUListenerAddParameter(inListener, inObject, inParameter)
}

// AUListenerCreate is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreate(_:_:_:_:_:_:)
func AUListenerCreate(inProc ParameterListenerProc, inUserData unsafe.Pointer, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inNotificationInterval unsafe.Pointer, outListener unsafe.Pointer) unsafe.Pointer {
	return _AUListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, outListener)
}

// AUListenerCreateWithDispatchQueue is a AudioToolbox function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreateWithDispatchQueue(_:_:_:_:)
func AUListenerCreateWithDispatchQueue(outListener unsafe.Pointer, inNotificationInterval unsafe.Pointer, inDispatchQueue unsafe.Pointer, inBlock unsafe.Pointer) unsafe.Pointer {
	return _AUListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inDispatchQueue, inBlock)
}

// AUListenerDispose is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerDispose(_:)
func AUListenerDispose(inListener ParameterListenerRef) unsafe.Pointer {
	return _AUListenerDispose(inListener)
}

// AUListenerRemoveParameter is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerRemoveParameter(_:_:_:)
func AUListenerRemoveParameter(inListener ParameterListenerRef, inObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUListenerRemoveParameter(inListener, inObject, inParameter)
}

// AUParameterFormatValue is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterFormatValue(_:_:_:_:)
func AUParameterFormatValue(inParameterValue unsafe.Pointer, inParameter unsafe.Pointer, inTextBuffer unsafe.Pointer, inDigits unsafe.Pointer) unsafe.Pointer {
	return _AUParameterFormatValue(inParameterValue, inParameter, inTextBuffer, inDigits)
}

// AUParameterListenerNotify is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerNotify(_:_:_:)
func AUParameterListenerNotify(inSendingListener ParameterListenerRef, inSendingObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUParameterListenerNotify(inSendingListener, inSendingObject, inParameter)
}

// AUParameterSet is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterSet(_:_:_:_:_:)
func AUParameterSet(inSendingListener ParameterListenerRef, inSendingObject unsafe.Pointer, inParameter unsafe.Pointer, inValue AudioUnitParameterValue, inBufferOffsetInFrames unsafe.Pointer) unsafe.Pointer {
	return _AUParameterSet(inSendingListener, inSendingObject, inParameter, inValue, inBufferOffsetInFrames)
}

// AUParameterValueFromLinear is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueFromLinear(_:_:)
func AUParameterValueFromLinear(inLinearValue unsafe.Pointer, inParameter unsafe.Pointer) AudioUnitParameterValue {
	return _AUParameterValueFromLinear(inLinearValue, inParameter)
}

// AUParameterValueToLinear is a AudioToolbox function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueToLinear(_:_:)
func AUParameterValueToLinear(inParameterValue AudioUnitParameterValue, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUParameterValueToLinear(inParameterValue, inParameter)
}

// AudioCodecAppendInputBufferList is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferList(_:_:_:_:_:)
func AudioCodecAppendInputBufferList(inCodec AudioCodec, inBufferList unsafe.Pointer, ioNumberPackets unsafe.Pointer, inPacketDescription unsafe.Pointer, outBytesConsumed unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecAppendInputBufferList(inCodec, inBufferList, ioNumberPackets, inPacketDescription, outBytesConsumed)
}

// Appends audio data to the codec’s input buffer.
//
// Added in macOS 10.2.
// Appends audio data to the codec’s input buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputData(_:_:_:_:_:)
func AudioCodecAppendInputData(inCodec AudioCodec, inInputData unsafe.Pointer, ioInputDataByteSize unsafe.Pointer, ioNumberPackets unsafe.Pointer, inPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecAppendInputData(inCodec, inInputData, ioInputDataByteSize, ioNumberPackets, inPacketDescription)
}

// Retrieves the value of a codec property.
//
// Added in macOS 10.2.
// Retrieves the value of a codec property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetProperty(_:_:_:_:)
func AudioCodecGetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecGetProperty(inCodec, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// Retrieves information about a codec property.
//
// Added in macOS 10.2.
// Retrieves information about a codec property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfo(_:_:_:_:)
func AudioCodecGetPropertyInfo(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecGetPropertyInfo(inCodec, inPropertyID, outSize, outWritable)
}

// Sets up the specified codec to perform a data format translation.
//
// Added in macOS 10.2.
// Sets up the specified codec to perform a data format translation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitialize(_:_:_:_:_:)
func AudioCodecInitialize(inCodec AudioCodec, inInputFormat unsafe.Pointer, inOutputFormat unsafe.Pointer, inMagicCookie unsafe.Pointer, inMagicCookieByteSize unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecInitialize(inCodec, inInputFormat, inOutputFormat, inMagicCookie, inMagicCookieByteSize)
}

// AudioCodecProduceOutputBufferList is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferList(_:_:_:_:_:)
func AudioCodecProduceOutputBufferList(inCodec AudioCodec, ioBufferList unsafe.Pointer, ioNumberPackets unsafe.Pointer, outPacketDescription unsafe.Pointer, outStatus unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecProduceOutputBufferList(inCodec, ioBufferList, ioNumberPackets, outPacketDescription, outStatus)
}

// Retrieves output data from a codec.
//
// Added in macOS 10.2.
// Retrieves output data from a codec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPackets(_:_:_:_:_:_:)
func AudioCodecProduceOutputPackets(inCodec AudioCodec, outOutputData unsafe.Pointer, ioOutputDataByteSize unsafe.Pointer, ioNumberPackets unsafe.Pointer, outPacketDescription unsafe.Pointer, outStatus unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecProduceOutputPackets(inCodec, outOutputData, ioOutputDataByteSize, ioNumberPackets, outPacketDescription, outStatus)
}

// Flushes all the audio data in the codec and clears the input buffer.
//
// Added in macOS 10.2.
// Flushes all the audio data in the codec and clears the input buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecReset(_:)
func AudioCodecReset(inCodec AudioCodec) unsafe.Pointer {
	return _AudioCodecReset(inCodec)
}

// Sets the value of a codec property.
//
// Added in macOS 10.2.
// Sets the value of a codec property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetProperty(_:_:_:_:)
func AudioCodecSetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecSetProperty(inCodec, inPropertyID, inPropertyDataSize, inPropertyData)
}

// Moves the codec from the initialized state back to the uninitialized state.
//
// Added in macOS 10.2.
// Moves the codec from the initialized state back to the uninitialized state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitialize(_:)
func AudioCodecUninitialize(inCodec AudioCodec) unsafe.Pointer {
	return _AudioCodecUninitialize(inCodec)
}

// AudioComponentCopyConfigurationInfo is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyConfigurationInfo(_:_:)
func AudioComponentCopyConfigurationInfo(inComponent AudioComponent, outConfigurationInfo unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCopyConfigurationInfo(inComponent, outConfigurationInfo)
}

// AudioComponentCopyIcon is a AudioToolbox function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyIcon(_:)
func AudioComponentCopyIcon(comp AudioComponent) unsafe.Pointer {
	return _AudioComponentCopyIcon(comp)
}

// Returns the generic name of an audio component.
//
// Added in macOS 10.6.
// Returns the generic name of an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyName(_:_:)
func AudioComponentCopyName(inComponent AudioComponent, outName unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCopyName(inComponent, outName)
}

// Returns the number of audio components that match a specified structure.
//
// Added in macOS 10.6.
// Returns the number of audio components that match a specified structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCount(_:)
func AudioComponentCount(inDesc unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCount(inDesc)
}

// Finds the next component that matches a specified structure after a specified audio component.
//
// Added in macOS 10.6.
// Finds the next component that matches a specified structure after a specified audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFindNext(_:_:)
func AudioComponentFindNext(inComponent AudioComponent, inDesc unsafe.Pointer) AudioComponent {
	return _AudioComponentFindNext(inComponent, inDesc)
}

// Gets the class description, as an structure, of an audio component.
//
// Added in macOS 10.6.
// Gets the class description, as an structure, of an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetDescription(_:_:)
func AudioComponentGetDescription(inComponent AudioComponent, outDesc unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetDescription(inComponent, outDesc)
}

// The UIImage of the audio component’s icon.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.11.
// The UIImage of the audio component’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetIcon(_:_:)
func AudioComponentGetIcon(comp AudioComponent, desiredPointSize float32) unsafe.Pointer {
	return _AudioComponentGetIcon(comp, desiredPointSize)
}

// The time at which the application publishing the component was last active.

// The time at which the application publishing the component was last active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetLastActiveTime(_:)
func AudioComponentGetLastActiveTime(comp AudioComponent) unsafe.Pointer {
	return _AudioComponentGetLastActiveTime(comp)
}

// Gets the version of an audio component in hexadecimal form as (major, minor, dot).
//
// Added in macOS 10.6.
// Gets the version of an audio component in hexadecimal form as (major, minor, dot).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetVersion(_:_:)
func AudioComponentGetVersion(inComponent AudioComponent, outVersion unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetVersion(inComponent, outVersion)
}

// Determines if an audio component instance implements a particular function.
//
// Added in macOS 10.6.
// Determines if an audio component instance implements a particular function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceCanDo(_:_:)
func AudioComponentInstanceCanDo(inInstance AudioComponentInstance, inSelectorID unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceCanDo(inInstance, inSelectorID)
}

// Disposes of an audio component instance.
//
// Added in macOS 10.6.
// Disposes of an audio component instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceDispose(_:)
func AudioComponentInstanceDispose(inInstance AudioComponentInstance) unsafe.Pointer {
	return _AudioComponentInstanceDispose(inInstance)
}

// Retrieves a reference to an audio component from an instance of that audio component.
//
// Added in macOS 10.6.
// Retrieves a reference to an audio component from an instance of that audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceGetComponent(_:)
func AudioComponentInstanceGetComponent(inInstance AudioComponentInstance) AudioComponent {
	return _AudioComponentInstanceGetComponent(inInstance)
}

// Creates a new instance of an audio component.
//
// Added in macOS 10.6.
// Creates a new instance of an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceNew(_:_:)
func AudioComponentInstanceNew(inComponent AudioComponent, outInstance unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceNew(inComponent, outInstance)
}

// AudioComponentInstantiate is a AudioToolbox function.
//
// Added in macOS 10.11.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiate(_:_:_:)
func AudioComponentInstantiate(inComponent AudioComponent, inOptions unsafe.Pointer) {
	_AudioComponentInstantiate(inComponent, inOptions)
}

// AudioComponentRegister is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentRegister(_:_:_:_:)
func AudioComponentRegister(inDesc unsafe.Pointer, inName unsafe.Pointer, inVersion unsafe.Pointer, inFactory AudioComponentFactoryFunction) AudioComponent {
	return _AudioComponentRegister(inDesc, inName, inVersion, inFactory)
}

// AudioComponentValidate is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidate(_:_:_:)
func AudioComponentValidate(inComponent AudioComponent, inValidationParameters unsafe.Pointer, outValidationResult unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentValidate(inComponent, inValidationParameters, outValidationResult)
}

// AudioComponentValidateWithResults is a AudioToolbox function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidateWithResults(_:_:_:)
func AudioComponentValidateWithResults(inComponent AudioComponent, inValidationParameters unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentValidateWithResults(inComponent, inValidationParameters)
}

// Converts audio data from one linear PCM format to another.
//
// Added in macOS 10.1.
// Converts audio data from one linear PCM format to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertBuffer(_:_:_:_:_:)
func AudioConverterConvertBuffer(inAudioConverter AudioConverterRef, inInputDataSize unsafe.Pointer, inInputData unsafe.Pointer, ioOutputDataSize unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterConvertBuffer(inAudioConverter, inInputDataSize, inInputData, ioOutputDataSize, outOutputData)
}

// Converts audio data from one linear PCM format to another, where both use the same sample rate.
//
// Added in macOS 10.7.
// Converts audio data from one linear PCM format to another, where both use the same sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertComplexBuffer(_:_:_:_:)
func AudioConverterConvertComplexBuffer(inAudioConverter AudioConverterRef, inNumberPCMFrames unsafe.Pointer, inInputData unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterConvertComplexBuffer(inAudioConverter, inNumberPCMFrames, inInputData, outOutputData)
}

// Disposes of an audio converter object.
//
// Added in macOS 10.1.
// Disposes of an audio converter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterDispose(_:)
func AudioConverterDispose(inAudioConverter AudioConverterRef) unsafe.Pointer {
	return _AudioConverterDispose(inAudioConverter)
}

// AudioConverterFillBuffer is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillBuffer
func AudioConverterFillBuffer(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataSize unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataSize, outOutputData)
}

// Converts audio data supplied by a callback function, supporting non-interleaved and packetized formats.
//
// Added in macOS 10.2.
// Converts audio data supplied by a callback function, supporting non-interleaved and packetized formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBuffer(_:_:_:_:_:_:)
func AudioConverterFillComplexBuffer(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
}

// AudioConverterFillComplexBufferRealtimeSafe is a AudioToolbox function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferRealtimeSafe(_:_:_:_:_:_:)
func AudioConverterFillComplexBufferRealtimeSafe(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProcRealtimeSafe, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBufferRealtimeSafe(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
}

// AudioConverterFillComplexBufferWithPacketDependencies is a AudioToolbox function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferWithPacketDependencies(_:_:_:_:_:_:_:)
func AudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescriptions unsafe.Pointer, outPacketDependencies unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescriptions, outPacketDependencies)
}

// Gets an audio converter property value.
//
// Added in macOS 10.1.
// Gets an audio converter property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetProperty(_:_:_:_:)
func AudioConverterGetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterGetProperty(inAudioConverter, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// Gets information about an audio converter property.
//
// Added in macOS 10.1.
// Gets information about an audio converter property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetPropertyInfo(_:_:_:_:)
func AudioConverterGetPropertyInfo(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterGetPropertyInfo(inAudioConverter, inPropertyID, outSize, outWritable)
}

// Creates a new audio converter object based on specified audio formats.
//
// Added in macOS 10.1.
// Creates a new audio converter object based on specified audio formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNew(_:_:_:)
func AudioConverterNew(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNew(inSourceFormat, inDestinationFormat, outAudioConverter)
}

// Creates a new audio converter object using a specified codec.
//
// Added in macOS 10.4.
// Creates a new audio converter object using a specified codec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewSpecific(_:_:_:_:_:)
func AudioConverterNewSpecific(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, inNumberClassDescriptions unsafe.Pointer, inClassDescriptions unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNewSpecific(inSourceFormat, inDestinationFormat, inNumberClassDescriptions, inClassDescriptions, outAudioConverter)
}

// AudioConverterNewWithOptions is a AudioToolbox function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewWithOptions(_:_:_:_:)
func AudioConverterNewWithOptions(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, inOptions unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNewWithOptions(inSourceFormat, inDestinationFormat, inOptions, outAudioConverter)
}

// AudioConverterPrepare is a AudioToolbox function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrepare(_:_:_:)
func AudioConverterPrepare(inFlags unsafe.Pointer, ioReserved unsafe.Pointer) {
	_AudioConverterPrepare(inFlags, ioReserved)
}

// Resets an audio converter object, clearing and flushing its buffers.
//
// Added in macOS 10.1.
// Resets an audio converter object, clearing and flushing its buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterReset(_:)
func AudioConverterReset(inAudioConverter AudioConverterRef) unsafe.Pointer {
	return _AudioConverterReset(inAudioConverter)
}

// Sets the value of an audio converter object property.
//
// Added in macOS 10.1.
// Sets the value of an audio converter object property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterSetProperty(_:_:_:_:)
func AudioConverterSetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterSetProperty(inAudioConverter, inPropertyID, inPropertyDataSize, inPropertyData)
}

// Closes an audio file.
//
// Added in macOS 10.2.
// Closes an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileClose(_:)
func AudioFileClose(inAudioFile AudioFileID) unsafe.Pointer {
	return _AudioFileClose(inAudioFile)
}

// AudioFileComponentCloseFile is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCloseFile(_:)
func AudioFileComponentCloseFile(inComponent AudioFileComponent) unsafe.Pointer {
	return _AudioFileComponentCloseFile(inComponent)
}

// AudioFileComponentCountUserData is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserData(_:_:_:)
func AudioFileComponentCountUserData(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, outNumberItems unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCountUserData(inComponent, inUserDataID, outNumberItems)
}

// AudioFileComponentCreate is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreate
func AudioFileComponentCreate(inComponent AudioFileComponent, inParentRef unsafe.Pointer, inFileName unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outNewFileRef unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCreate(inComponent, inParentRef, inFileName, inFormat, inFlags, outNewFileRef)
}

// AudioFileComponentCreateURL is a AudioToolbox function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURL(_:_:_:_:)
func AudioFileComponentCreateURL(inComponent AudioFileComponent, inFileRef unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCreateURL(inComponent, inFileRef, inFormat, inFlags)
}

// AudioFileComponentDataIsThisFormat is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentDataIsThisFormat
func AudioFileComponentDataIsThisFormat(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentDataIsThisFormat(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, outResult)
}

// AudioFileComponentExtensionIsThisFormat is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormat(_:_:_:)
func AudioFileComponentExtensionIsThisFormat(inComponent AudioFileComponent, inExtension unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentExtensionIsThisFormat(inComponent, inExtension, outResult)
}

// AudioFileComponentFileDataIsThisFormat is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormat(_:_:_:_:)
func AudioFileComponentFileDataIsThisFormat(inComponent AudioFileComponent, inDataByteSize unsafe.Pointer, inData unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentFileDataIsThisFormat(inComponent, inDataByteSize, inData, outResult)
}

// AudioFileComponentFileIsThisFormat is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileIsThisFormat
func AudioFileComponentFileIsThisFormat(inComponent AudioFileComponent, inFileRefNum unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentFileIsThisFormat(inComponent, inFileRefNum, outResult)
}

// AudioFileComponentGetGlobalInfo is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfo(_:_:_:_:_:_:)
func AudioFileComponentGetGlobalInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetGlobalInfo(inComponent, inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
}

// AudioFileComponentGetGlobalInfoSize is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSize(_:_:_:_:_:)
func AudioFileComponentGetGlobalInfoSize(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertySize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetGlobalInfoSize(inComponent, inPropertyID, inSpecifierSize, inSpecifier, outPropertySize)
}

// AudioFileComponentGetProperty is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetProperty(_:_:_:_:)
func AudioFileComponentGetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetProperty(inComponent, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioFileComponentGetPropertyInfo is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfo(_:_:_:_:)
func AudioFileComponentGetPropertyInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, outPropertySize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetPropertyInfo(inComponent, inPropertyID, outPropertySize, outWritable)
}

// AudioFileComponentGetUserData is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserData(_:_:_:_:_:)
func AudioFileComponentGetUserData(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserData(inComponent, inUserDataID, inIndex, ioUserDataSize, outUserData)
}

// AudioFileComponentGetUserDataAtOffset is a AudioToolbox function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileComponentGetUserDataAtOffset(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inOffset unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataAtOffset(inComponent, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
}

// AudioFileComponentGetUserDataSize is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize(_:_:_:_:)
func AudioFileComponentGetUserDataSize(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataSize(inComponent, inUserDataID, inIndex, outUserDataSize)
}

// AudioFileComponentGetUserDataSize64 is a AudioToolbox function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64(_:_:_:_:)
func AudioFileComponentGetUserDataSize64(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataSize64(inComponent, inUserDataID, inIndex, outUserDataSize)
}

// AudioFileComponentInitialize is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitialize
func AudioFileComponentInitialize(inComponent AudioFileComponent, inFileRef unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentInitialize(inComponent, inFileRef, inFormat, inFlags)
}

// AudioFileComponentInitializeWithCallbacks is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileComponentInitializeWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentInitializeWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags)
}

// AudioFileComponentOpenFile is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenFile
func AudioFileComponentOpenFile(inComponent AudioFileComponent, inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inRefNum unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentOpenFile(inComponent, inFileRef, inPermissions, inRefNum)
}

// AudioFileComponentOpenURL is a AudioToolbox function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURL(_:_:_:_:)
func AudioFileComponentOpenURL(inComponent AudioFileComponent, inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileDescriptor int) unsafe.Pointer {
	return _AudioFileComponentOpenURL(inComponent, inFileRef, inPermissions, inFileDescriptor)
}

// AudioFileComponentOpenWithCallbacks is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacks(_:_:_:_:_:_:)
func AudioFileComponentOpenWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc) unsafe.Pointer {
	return _AudioFileComponentOpenWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc)
}

// AudioFileComponentOptimize is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimize(_:)
func AudioFileComponentOptimize(inComponent AudioFileComponent) unsafe.Pointer {
	return _AudioFileComponentOptimize(inComponent)
}

// AudioFileComponentReadBytes is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytes(_:_:_:_:_:)
func AudioFileComponentReadBytes(inComponent AudioFileComponent, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, outBuffer)
}

// AudioFileComponentReadPacketData is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketData(_:_:_:_:_:_:_:)
func AudioFileComponentReadPacketData(inComponent AudioFileComponent, inUseCache unsafe.Pointer, ioNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadPacketData(inComponent, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
}

// AudioFileComponentReadPackets is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPackets(_:_:_:_:_:_:_:)
func AudioFileComponentReadPackets(inComponent AudioFileComponent, inUseCache unsafe.Pointer, outNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadPackets(inComponent, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
}

// AudioFileComponentRemoveUserData is a AudioToolbox function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserData(_:_:_:)
func AudioFileComponentRemoveUserData(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentRemoveUserData(inComponent, inUserDataID, inIndex)
}

// AudioFileComponentSetProperty is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetProperty(_:_:_:_:)
func AudioFileComponentSetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentSetProperty(inComponent, inPropertyID, inPropertyDataSize, inPropertyData)
}

// AudioFileComponentSetUserData is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserData(_:_:_:_:_:)
func AudioFileComponentSetUserData(inComponent AudioFileComponent, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inUserDataSize unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentSetUserData(inComponent, inUserDataID, inIndex, inUserDataSize, inUserData)
}

// AudioFileComponentWriteBytes is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytes(_:_:_:_:_:)
func AudioFileComponentWriteBytes(inComponent AudioFileComponent, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentWriteBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, inBuffer)
}

// AudioFileComponentWritePackets is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePackets(_:_:_:_:_:_:_:)
func AudioFileComponentWritePackets(inComponent AudioFileComponent, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentWritePackets(inComponent, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
}

// Gets the number of user data items with a specified ID in a file.
//
// Added in macOS 10.4.
// Gets the number of user data items with a specified ID in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCountUserData(_:_:_:)
func AudioFileCountUserData(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, outNumberItems unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCountUserData(inAudioFile, inUserDataID, outNumberItems)
}

// AudioFileCreate is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCreate
func AudioFileCreate(inParentRef unsafe.Pointer, inFileName unsafe.Pointer, inFileType AudioFileTypeID, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outNewFileRef unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCreate(inParentRef, inFileName, inFileType, inFormat, inFlags, outNewFileRef, outAudioFile)
}

// Creates a new audio file, or initializes an existing file, specified by a URL.
//
// Added in macOS 10.5.
// Creates a new audio file, or initializes an existing file, specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCreateWithURL(_:_:_:_:_:)
func AudioFileCreateWithURL(inFileRef unsafe.Pointer, inFileType AudioFileTypeID, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCreateWithURL(inFileRef, inFileType, inFormat, inFlags, outAudioFile)
}

// Copies the value of a global property into a buffer.
//
// Added in macOS 10.3.
// Copies the value of a global property into a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfo(_:_:_:_:_:)
func AudioFileGetGlobalInfo(inPropertyID AudioFilePropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetGlobalInfo(inPropertyID, inSpecifierSize, inSpecifier, ioDataSize, outPropertyData)
}

// Gets the size of a global audio file property.
//
// Added in macOS 10.3.
// Gets the size of a global audio file property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfoSize(_:_:_:_:)
func AudioFileGetGlobalInfoSize(inPropertyID AudioFilePropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetGlobalInfoSize(inPropertyID, inSpecifierSize, inSpecifier, outDataSize)
}

// Gets the value of an audio file property.
//
// Added in macOS 10.2.
// Gets the value of an audio file property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetProperty(_:_:_:_:)
func AudioFileGetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, ioDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetProperty(inAudioFile, inPropertyID, ioDataSize, outPropertyData)
}

// Gets information about an audio file property, including the size of the property value and whether the value is writable.
//
// Added in macOS 10.2.
// Gets information about an audio file property, including the size of the property value and whether the value is writable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetPropertyInfo(_:_:_:_:)
func AudioFileGetPropertyInfo(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, outDataSize unsafe.Pointer, isWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetPropertyInfo(inAudioFile, inPropertyID, outDataSize, isWritable)
}

// Gets a chunk from an audio file.
//
// Added in macOS 10.4.
// Gets a chunk from an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserData(_:_:_:_:_:)
func AudioFileGetUserData(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserData(inAudioFile, inUserDataID, inIndex, ioUserDataSize, outUserData)
}

// Gets part of the data from a chunk in an audio file.
//
// Added in macOS 14.0.
// Gets part of the data from a chunk in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileGetUserDataAtOffset(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inOffset unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataAtOffset(inAudioFile, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
}

// Gets the size of a user data item in an audio file.
//
// Added in macOS 10.4.
// Gets the size of a user data item in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize(_:_:_:_:)
func AudioFileGetUserDataSize(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataSize(inAudioFile, inUserDataID, inIndex, outUserDataSize)
}

// Gets the size of a user data item in an audio file.
//
// Added in macOS 14.0.
// Gets the size of a user data item in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize64(_:_:_:_:)
func AudioFileGetUserDataSize64(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataSize64(inAudioFile, inUserDataID, inIndex, outUserDataSize)
}

// AudioFileInitialize is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileInitialize
func AudioFileInitialize(inFileRef unsafe.Pointer, inFileType AudioFileTypeID, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileInitialize(inFileRef, inFileType, inFormat, inFlags, outAudioFile)
}

// Deletes the content of an existing file and assigns callbacks to the audio file object.
//
// Added in macOS 10.3.
// Deletes the content of an existing file and assigns callbacks to the audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileInitializeWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType AudioFileTypeID, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileInitializeWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags, outAudioFile)
}

// AudioFileOpen is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpen
func AudioFileOpen(inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileTypeHint AudioFileTypeID, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpen(inFileRef, inPermissions, inFileTypeHint, outAudioFile)
}

// Open an existing audio file specified by a URL.
//
// Added in macOS 10.5.
// Open an existing audio file specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenURL(_:_:_:_:)
func AudioFileOpenURL(inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileTypeHint AudioFileTypeID, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpenURL(inFileRef, inPermissions, inFileTypeHint, outAudioFile)
}

// Opens an existing file with callbacks you provide.
//
// Added in macOS 10.3.
// Opens an existing file with callbacks you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenWithCallbacks(_:_:_:_:_:_:_:)
func AudioFileOpenWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileTypeHint AudioFileTypeID, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpenWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileTypeHint, outAudioFile)
}

// Consolidates audio data and performs other internal optimizations of the file structure.
//
// Added in macOS 10.2.
// Consolidates audio data and performs other internal optimizations of the file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOptimize(_:)
func AudioFileOptimize(inAudioFile AudioFileID) unsafe.Pointer {
	return _AudioFileOptimize(inAudioFile)
}

// Reads bytes of audio data from an audio file.
//
// Added in macOS 10.2.
// Reads bytes of audio data from an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadBytes(_:_:_:_:_:)
func AudioFileReadBytes(inAudioFile AudioFileID, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, outBuffer)
}

// Reads packets of audio data from an audio file.
//
// Added in macOS 10.6.
// Reads packets of audio data from an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadPacketData(_:_:_:_:_:_:_:)
func AudioFileReadPacketData(inAudioFile AudioFileID, inUseCache unsafe.Pointer, ioNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadPacketData(inAudioFile, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
}

// Reads a fixed duration of audio data from an audio file.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Reads a fixed duration of audio data from an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadPackets(_:_:_:_:_:_:_:)
func AudioFileReadPackets(inAudioFile AudioFileID, inUseCache unsafe.Pointer, outNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadPackets(inAudioFile, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
}

// Removes a user data item from an audio file.
//
// Added in macOS 10.5.
// Removes a user data item from an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRemoveUserData(_:_:_:)
func AudioFileRemoveUserData(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer) unsafe.Pointer {
	return _AudioFileRemoveUserData(inAudioFile, inUserDataID, inIndex)
}

// Sets the value of an audio file property
//
// Added in macOS 10.2.
// Sets the value of an audio file property
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetProperty(_:_:_:_:)
func AudioFileSetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, inDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileSetProperty(inAudioFile, inPropertyID, inDataSize, inPropertyData)
}

// Sets a user data item in an audio file.
//
// Added in macOS 10.4.
// Sets a user data item in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetUserData(_:_:_:_:_:)
func AudioFileSetUserData(inAudioFile AudioFileID, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inUserDataSize unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileSetUserData(inAudioFile, inUserDataID, inIndex, inUserDataSize, inUserData)
}

// Closes and deallocates the specified audio file stream parser.
//
// Added in macOS 10.5.
// Closes and deallocates the specified audio file stream parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamClose(_:)
func AudioFileStreamClose(inAudioFileStream AudioFileStreamID) unsafe.Pointer {
	return _AudioFileStreamClose(inAudioFileStream)
}

// Retrieves the value of the specified property.
//
// Added in macOS 10.5.
// Retrieves the value of the specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetProperty(_:_:_:_:)
func AudioFileStreamGetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamGetProperty(inAudioFileStream, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// Retrieves information about a property value.
//
// Added in macOS 10.5.
// Retrieves information about a property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetPropertyInfo(_:_:_:_:)
func AudioFileStreamGetPropertyInfo(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, outPropertyDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamGetPropertyInfo(inAudioFileStream, inPropertyID, outPropertyDataSize, outWritable)
}

// Creates and opens a new audio file stream parser.
//
// Added in macOS 10.5.
// Creates and opens a new audio file stream parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamOpen(_:_:_:_:_:)
func AudioFileStreamOpen(inClientData unsafe.Pointer, inPropertyListenerProc AudioFileStream_PropertyListenerProc, inPacketsProc AudioFileStream_PacketsProc, inFileTypeHint AudioFileTypeID, outAudioFileStream unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamOpen(inClientData, inPropertyListenerProc, inPacketsProc, inFileTypeHint, outAudioFileStream)
}

// Passes audio file stream data to the parser.
//
// Added in macOS 10.5.
// Passes audio file stream data to the parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseBytes(_:_:_:_:)
func AudioFileStreamParseBytes(inAudioFileStream AudioFileStreamID, inDataByteSize unsafe.Pointer, inData unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamParseBytes(inAudioFileStream, inDataByteSize, inData, inFlags)
}

// Provides a byte offset for a specified packet in the data stream.
//
// Added in macOS 10.5.
// Provides a byte offset for a specified packet in the data stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeek(_:_:_:_:)
func AudioFileStreamSeek(inAudioFileStream AudioFileStreamID, inPacketOffset unsafe.Pointer, outDataByteOffset unsafe.Pointer, ioFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamSeek(inAudioFileStream, inPacketOffset, outDataByteOffset, ioFlags)
}

// Sets the value of the specified property.
//
// Added in macOS 10.5.
// Sets the value of the specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSetProperty(_:_:_:_:)
func AudioFileStreamSetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamSetProperty(inAudioFileStream, inPropertyID, inPropertyDataSize, inPropertyData)
}

// Writes bytes of audio data to an audio file.
//
// Added in macOS 10.2.
// Writes bytes of audio data to an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWriteBytes(_:_:_:_:_:)
func AudioFileWriteBytes(inAudioFile AudioFileID, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWriteBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, inBuffer)
}

// Writes packets of audio data to an audio data file.
//
// Added in macOS 10.2.
// Writes packets of audio data to an audio data file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePackets(_:_:_:_:_:_:_:)
func AudioFileWritePackets(inAudioFile AudioFileID, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWritePackets(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
}

// AudioFileWritePacketsWithDependencies is a AudioToolbox function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePacketsWithDependencies(_:_:_:_:_:_:_:_:)
func AudioFileWritePacketsWithDependencies(inAudioFile AudioFileID, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inPacketDependencies unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWritePacketsWithDependencies(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inPacketDependencies, inStartingPacket, ioNumPackets, inBuffer)
}

// Gets the value of an audio format property.
//
// Added in macOS 10.3.
// Gets the value of an audio format property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetProperty(_:_:_:_:_:)
func AudioFormatGetProperty(inPropertyID AudioFormatPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFormatGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
}

// Gets information about an audio format property.
//
// Added in macOS 10.3.
// Gets information about an audio format property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetPropertyInfo(_:_:_:_:)
func AudioFormatGetPropertyInfo(inPropertyID AudioFormatPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertyDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFormatGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize)
}

// Registers a HAL audio object property listener callback function to be invoked when a specified property changes.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Registers a HAL audio object property listener callback function to be invoked when a specified property changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceAddPropertyListener(_:_:_:_:)
func AudioHardwareServiceAddPropertyListener(inObjectID AudioObjectID, inAddress unsafe.Pointer, inListener unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceAddPropertyListener(inObjectID, inAddress, inListener, inClientData)
}

// Gets the value for a specified property.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Gets the value for a specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceGetPropertyData(_:_:_:_:_:_:)
func AudioHardwareServiceGetPropertyData(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, ioDataSize unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData)
}

// Gets the payload size for a given property.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Gets the payload size for a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceGetPropertyDataSize(_:_:_:_:_:)
func AudioHardwareServiceGetPropertyDataSize(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize)
}

// Queries a HAL audio object about whether or not it has a specified property.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Queries a HAL audio object about whether or not it has a specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceHasProperty(_:_:)
func AudioHardwareServiceHasProperty(inObjectID AudioObjectID, inAddress unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceHasProperty(inObjectID, inAddress)
}

// Queries a HAL audio object about whether a specified property is settable.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Queries a HAL audio object about whether a specified property is settable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceIsPropertySettable(_:_:_:)
func AudioHardwareServiceIsPropertySettable(inObjectID AudioObjectID, inAddress unsafe.Pointer, outIsSettable unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceIsPropertySettable(inObjectID, inAddress, outIsSettable)
}

// Unregisters a HAL audio object property listener callback function.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Unregisters a HAL audio object property listener callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceRemovePropertyListener(_:_:_:_:)
func AudioHardwareServiceRemovePropertyListener(inObjectID AudioObjectID, inAddress unsafe.Pointer, inListener unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceRemovePropertyListener(inObjectID, inAddress, inListener, inClientData)
}

// Asks a HAL audio object to change the value of a specified property.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
// Asks a HAL audio object to change the value of a specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceSetPropertyData(_:_:_:_:_:_:)
func AudioHardwareServiceSetPropertyData(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, inDataSize unsafe.Pointer, inData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData)
}

// The host app’s icon.

// The host app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitGetHostIcon(_:_:)
func AudioOutputUnitGetHostIcon(au AudioUnit, desiredPointSize float32) unsafe.Pointer {
	return _AudioOutputUnitGetHostIcon(au, desiredPointSize)
}

// Registers an audio output unit for use by other applications.

// Registers an audio output unit for use by other applications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitPublish(_:_:_:_:)
func AudioOutputUnitPublish(inDesc unsafe.Pointer, inName unsafe.Pointer, inVersion unsafe.Pointer, inOutputUnit AudioUnit) unsafe.Pointer {
	return _AudioOutputUnitPublish(inDesc, inName, inVersion, inOutputUnit)
}

// Starts an I/O audio unit, which in turn starts the audio unit processing graph that it is connected to.
//
// Added in macOS 10.0.
// Starts an I/O audio unit, which in turn starts the audio unit processing graph that it is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStart(_:)
func AudioOutputUnitStart(ci AudioUnit) unsafe.Pointer {
	return _AudioOutputUnitStart(ci)
}

// Stops an I/O audio unit, which in turn stops the audio unit processing graph that it is connected to.
//
// Added in macOS 10.0.
// Stops an I/O audio unit, which in turn stops the audio unit processing graph that it is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStop(_:)
func AudioOutputUnitStop(ci AudioUnit) unsafe.Pointer {
	return _AudioOutputUnitStop(ci)
}

// Adds a property listener callback to an audio queue.
//
// Added in macOS 10.5.
// Adds a property listener callback to an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAddPropertyListener(_:_:_:_:)
func AudioQueueAddPropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAddPropertyListener(inAQ, inID, inProc, inUserData)
}

// Asks an audio queue object to allocate an audio queue buffer.
//
// Added in macOS 10.5.
// Asks an audio queue object to allocate an audio queue buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBuffer(_:_:_:)
func AudioQueueAllocateBuffer(inAQ AudioQueueRef, inBufferByteSize unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAllocateBuffer(inAQ, inBufferByteSize, outBuffer)
}

// Asks an audio queue object to allocate an audio queue buffer with space for packet descriptions.
//
// Added in macOS 10.6.
// Asks an audio queue object to allocate an audio queue buffer with space for packet descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBufferWithPacketDescriptions(_:_:_:_:)
func AudioQueueAllocateBufferWithPacketDescriptions(inAQ AudioQueueRef, inBufferByteSize unsafe.Pointer, inNumberPacketDescriptions unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAllocateBufferWithPacketDescriptions(inAQ, inBufferByteSize, inNumberPacketDescriptions, outBuffer)
}

// Creates a timeline object for an audio queue.
//
// Added in macOS 10.5.
// Creates a timeline object for an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueCreateTimeline(_:_:)
func AudioQueueCreateTimeline(inAQ AudioQueueRef, outTimeline unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueCreateTimeline(inAQ, outTimeline)
}

// Gets the current time of the audio hardware device associated with an audio queue.
//
// Added in macOS 10.5.
// Gets the current time of the audio hardware device associated with an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetCurrentTime(_:_:)
func AudioQueueDeviceGetCurrentTime(inAQ AudioQueueRef, outTimeStamp unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceGetCurrentTime(inAQ, outTimeStamp)
}

// Gets the start time, for an audio hardware device, that is closest to a requested start time.
//
// Added in macOS 10.5.
// Gets the start time, for an audio hardware device, that is closest to a requested start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetNearestStartTime(_:_:_:)
func AudioQueueDeviceGetNearestStartTime(inAQ AudioQueueRef, ioRequestedStartTime unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceGetNearestStartTime(inAQ, ioRequestedStartTime, inFlags)
}

// Converts the time for an audio queue’s associated audio hardware device from one time base representation to another.
//
// Added in macOS 10.5.
// Converts the time for an audio queue’s associated audio hardware device from one time base representation to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceTranslateTime(_:_:_:)
func AudioQueueDeviceTranslateTime(inAQ AudioQueueRef, inTime unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceTranslateTime(inAQ, inTime, outTime)
}

// Disposes of an audio queue.
//
// Added in macOS 10.5.
// Disposes of an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDispose(_:_:)
func AudioQueueDispose(inAQ AudioQueueRef, inImmediate unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDispose(inAQ, inImmediate)
}

// Disposes of an audio queue’s timeline object.
//
// Added in macOS 10.5.
// Disposes of an audio queue’s timeline object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDisposeTimeline(_:_:)
func AudioQueueDisposeTimeline(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef) unsafe.Pointer {
	return _AudioQueueDisposeTimeline(inAQ, inTimeline)
}

// Adds a buffer to the buffer queue of a recording or playback audio queue.
//
// Added in macOS 10.5.
// Adds a buffer to the buffer queue of a recording or playback audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBuffer(_:_:_:_:)
func AudioQueueEnqueueBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs unsafe.Pointer, inPacketDescs unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueEnqueueBuffer(inAQ, inBuffer, inNumPacketDescs, inPacketDescs)
}

// Adds a buffer to the buffer queue of a playback audio queue object, specifying start time and other settings.
//
// Added in macOS 10.5.
// Adds a buffer to the buffer queue of a playback audio queue object, specifying start time and other settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBufferWithParameters(_:_:_:_:_:_:_:_:_:_:)
func AudioQueueEnqueueBufferWithParameters(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs unsafe.Pointer, inPacketDescs unsafe.Pointer, inTrimFramesAtStart unsafe.Pointer, inTrimFramesAtEnd unsafe.Pointer, inNumParamValues unsafe.Pointer, inParamValues unsafe.Pointer, inStartTime unsafe.Pointer, outActualStartTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueEnqueueBufferWithParameters(inAQ, inBuffer, inNumPacketDescs, inPacketDescs, inTrimFramesAtStart, inTrimFramesAtEnd, inNumParamValues, inParamValues, inStartTime, outActualStartTime)
}

// Resets an audio queue’s decoder state.
//
// Added in macOS 10.5.
// Resets an audio queue’s decoder state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFlush(_:)
func AudioQueueFlush(inAQ AudioQueueRef) unsafe.Pointer {
	return _AudioQueueFlush(inAQ)
}

// Asks an audio queue to dispose of an audio queue buffer.
//
// Added in macOS 10.5.
// Asks an audio queue to dispose of an audio queue buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFreeBuffer(_:_:)
func AudioQueueFreeBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef) unsafe.Pointer {
	return _AudioQueueFreeBuffer(inAQ, inBuffer)
}

// Gets the current audio queue time.
//
// Added in macOS 10.5.
// Gets the current audio queue time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetCurrentTime(_:_:_:_:)
func AudioQueueGetCurrentTime(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef, outTimeStamp unsafe.Pointer, outTimelineDiscontinuity unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetCurrentTime(inAQ, inTimeline, outTimeStamp, outTimelineDiscontinuity)
}

// Gets an audio queue parameter value.
//
// Added in macOS 10.5.
// Gets an audio queue parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetParameter(_:_:_:)
func AudioQueueGetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, outValue unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetParameter(inAQ, inParamID, outValue)
}

// Gets an audio queue property value.
//
// Added in macOS 10.5.
// Gets an audio queue property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetProperty(_:_:_:_:)
func AudioQueueGetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, outData unsafe.Pointer, ioDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetProperty(inAQ, inID, outData, ioDataSize)
}

// Gets the size of the value of an audio queue property.
//
// Added in macOS 10.5.
// Gets the size of the value of an audio queue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetPropertySize(_:_:_:)
func AudioQueueGetPropertySize(inAQ AudioQueueRef, inID AudioQueuePropertyID, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetPropertySize(inAQ, inID, outDataSize)
}

// Creates a new recording audio queue object.
//
// Added in macOS 10.5.
// Creates a new recording audio queue object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInput(_:_:_:_:_:_:_:)
func AudioQueueNewInput(inFormat unsafe.Pointer, inCallbackProc AudioQueueInputCallback, inUserData unsafe.Pointer, inCallbackRunLoop unsafe.Pointer, inCallbackRunLoopMode unsafe.Pointer, inFlags unsafe.Pointer, outAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewInput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
}

// AudioQueueNewInputWithDispatchQueue is a AudioToolbox function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewInputWithDispatchQueue(outAQ unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, inCallbackDispatchQueue unsafe.Pointer, inCallbackBlock unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewInputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
}

// Creates a new playback audio queue object.
//
// Added in macOS 10.5.
// Creates a new playback audio queue object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutput(_:_:_:_:_:_:_:)
func AudioQueueNewOutput(inFormat unsafe.Pointer, inCallbackProc AudioQueueOutputCallback, inUserData unsafe.Pointer, inCallbackRunLoop unsafe.Pointer, inCallbackRunLoopMode unsafe.Pointer, inFlags unsafe.Pointer, outAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewOutput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
}

// AudioQueueNewOutputWithDispatchQueue is a AudioToolbox function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewOutputWithDispatchQueue(outAQ unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, inCallbackDispatchQueue unsafe.Pointer, inCallbackBlock unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewOutputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
}

// Exports audio to a buffer, instead of to a device, using a playback audio queue.
//
// Added in macOS 10.5.
// Exports audio to a buffer, instead of to a device, using a playback audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOfflineRender(_:_:_:_:)
func AudioQueueOfflineRender(inAQ AudioQueueRef, inTimestamp unsafe.Pointer, ioBuffer AudioQueueBufferRef, inNumberFrames unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueOfflineRender(inAQ, inTimestamp, ioBuffer, inNumberFrames)
}

// Pauses audio playback or recording.
//
// Added in macOS 10.5.
// Pauses audio playback or recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePause(_:)
func AudioQueuePause(inAQ AudioQueueRef) unsafe.Pointer {
	return _AudioQueuePause(inAQ)
}

// Decodes enqueued buffers in preparation for playback.
//
// Added in macOS 10.5.
// Decodes enqueued buffers in preparation for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePrime(_:_:_:)
func AudioQueuePrime(inAQ AudioQueueRef, inNumberOfFramesToPrepare unsafe.Pointer, outNumberOfFramesPrepared unsafe.Pointer) unsafe.Pointer {
	return _AudioQueuePrime(inAQ, inNumberOfFramesToPrepare, outNumberOfFramesPrepared)
}

// AudioQueueProcessingTapDispose is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapDispose(_:)
func AudioQueueProcessingTapDispose(inAQTap AudioQueueProcessingTapRef) unsafe.Pointer {
	return _AudioQueueProcessingTapDispose(inAQTap)
}

// AudioQueueProcessingTapGetQueueTime is a AudioToolbox function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetQueueTime(_:_:_:)
func AudioQueueProcessingTapGetQueueTime(inAQTap AudioQueueProcessingTapRef, outQueueSampleTime unsafe.Pointer, outQueueFrameCount unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapGetQueueTime(inAQTap, outQueueSampleTime, outQueueFrameCount)
}

// AudioQueueProcessingTapGetSourceAudio is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetSourceAudio(_:_:_:_:_:_:)
func AudioQueueProcessingTapGetSourceAudio(inAQTap AudioQueueProcessingTapRef, inNumberFrames unsafe.Pointer, ioTimeStamp unsafe.Pointer, outFlags unsafe.Pointer, outNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapGetSourceAudio(inAQTap, inNumberFrames, ioTimeStamp, outFlags, outNumberFrames, ioData)
}

// AudioQueueProcessingTapNew is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapNew(_:_:_:_:_:_:_:)
func AudioQueueProcessingTapNew(inAQ AudioQueueRef, inCallback AudioQueueProcessingTapCallback, inClientData unsafe.Pointer, inFlags unsafe.Pointer, outMaxFrames unsafe.Pointer, outProcessingFormat unsafe.Pointer, outAQTap unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapNew(inAQ, inCallback, inClientData, inFlags, outMaxFrames, outProcessingFormat, outAQTap)
}

// Removes a property listener callback from an audio queue.
//
// Added in macOS 10.5.
// Removes a property listener callback from an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueRemovePropertyListener(_:_:_:_:)
func AudioQueueRemovePropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueRemovePropertyListener(inAQ, inID, inProc, inUserData)
}

// Resets an audio queue.
//
// Added in macOS 10.5.
// Resets an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueReset(_:)
func AudioQueueReset(inAQ AudioQueueRef) unsafe.Pointer {
	return _AudioQueueReset(inAQ)
}

// Sets the rendering mode and audio format for a playback audio queue.
//
// Added in macOS 10.5.
// Sets the rendering mode and audio format for a playback audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetOfflineRenderFormat(_:_:_:)
func AudioQueueSetOfflineRenderFormat(inAQ AudioQueueRef, inFormat unsafe.Pointer, inLayout unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueSetOfflineRenderFormat(inAQ, inFormat, inLayout)
}

// Sets a playback audio queue parameter value.
//
// Added in macOS 10.5.
// Sets a playback audio queue parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetParameter(_:_:_:)
func AudioQueueSetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, inValue AudioQueueParameterValue) unsafe.Pointer {
	return _AudioQueueSetParameter(inAQ, inParamID, inValue)
}

// Sets an audio queue property value.
//
// Added in macOS 10.5.
// Sets an audio queue property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetProperty(_:_:_:_:)
func AudioQueueSetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, inData unsafe.Pointer, inDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueSetProperty(inAQ, inID, inData, inDataSize)
}

// Begins playing or recording audio.
//
// Added in macOS 10.5.
// Begins playing or recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStart(_:_:)
func AudioQueueStart(inAQ AudioQueueRef, inStartTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueStart(inAQ, inStartTime)
}

// Stops playing or recording audio.
//
// Added in macOS 10.5.
// Stops playing or recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStop(_:_:)
func AudioQueueStop(inAQ AudioQueueRef, inImmediate unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueStop(inAQ, inImmediate)
}

// Registers a callback function that is invoked when a specified system sound finishes playing.
//
// Added in macOS 10.5.
// Registers a callback function that is invoked when a specified system sound finishes playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesAddSystemSoundCompletion(_:_:_:_:_:)
func AudioServicesAddSystemSoundCompletion(inSystemSoundID SystemSoundID, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inCompletionRoutine AudioServicesSystemSoundCompletionProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesAddSystemSoundCompletion(inSystemSoundID, inRunLoop, inRunLoopMode, inCompletionRoutine, inClientData)
}

// Creates a system sound object.
//
// Added in macOS 10.5.
// Creates a system sound object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesCreateSystemSoundID(_:_:)
func AudioServicesCreateSystemSoundID(inFileURL unsafe.Pointer, outSystemSoundID unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesCreateSystemSoundID(inFileURL, outSystemSoundID)
}

// Disposes of a system sound object and associated resources.
//
// Added in macOS 10.5.
// Disposes of a system sound object and associated resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesDisposeSystemSoundID(_:)
func AudioServicesDisposeSystemSoundID(inSystemSoundID SystemSoundID) unsafe.Pointer {
	return _AudioServicesDisposeSystemSoundID(inSystemSoundID)
}

// Gets a specified System Sound Services property value.
//
// Added in macOS 10.5.
// Gets a specified System Sound Services property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetProperty(_:_:_:_:_:)
func AudioServicesGetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
}

// Gets information about a System Sound Services property.
//
// Added in macOS 10.5.
// Gets information about a System Sound Services property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetPropertyInfo(_:_:_:_:_:)
func AudioServicesGetPropertyInfo(inPropertyID AudioServicesPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertyDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize, outWritable)
}

// Plays a system sound as an alert.
//
// Added in macOS 10.5.
// Plays a system sound as an alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSound(_:)
func AudioServicesPlayAlertSound(inSystemSoundID SystemSoundID) {
	_AudioServicesPlayAlertSound(inSystemSoundID)
}

// AudioServicesPlayAlertSoundWithCompletion is a AudioToolbox function.
//
// Added in macOS 10.11.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSoundWithCompletion(_:_:)
func AudioServicesPlayAlertSoundWithCompletion(inSystemSoundID SystemSoundID) {
	_AudioServicesPlayAlertSoundWithCompletion(inSystemSoundID)
}

// AudioServicesPlayAlertSoundWithDetails is a AudioToolbox function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSoundWithDetails
func AudioServicesPlayAlertSoundWithDetails(inSystemSoundID SystemSoundID, inDetails unsafe.Pointer) {
	_AudioServicesPlayAlertSoundWithDetails(inSystemSoundID, inDetails)
}

// Plays a system sound object.
//
// Added in macOS 10.5.
// Plays a system sound object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSound(_:)
func AudioServicesPlaySystemSound(inSystemSoundID SystemSoundID) {
	_AudioServicesPlaySystemSound(inSystemSoundID)
}

// AudioServicesPlaySystemSoundWithCompletion is a AudioToolbox function.
//
// Added in macOS 10.11.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSoundWithCompletion(_:_:)
func AudioServicesPlaySystemSoundWithCompletion(inSystemSoundID SystemSoundID) {
	_AudioServicesPlaySystemSoundWithCompletion(inSystemSoundID)
}

// AudioServicesPlaySystemSoundWithDetails is a AudioToolbox function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSoundWithDetails
func AudioServicesPlaySystemSoundWithDetails(inSystemSoundID SystemSoundID, inDetails unsafe.Pointer) {
	_AudioServicesPlaySystemSoundWithDetails(inSystemSoundID, inDetails)
}

// Unregisters any completion callback functions that were registered for a specified system sound.
//
// Added in macOS 10.5.
// Unregisters any completion callback functions that were registered for a specified system sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesRemoveSystemSoundCompletion(_:)
func AudioServicesRemoveSystemSoundCompletion(inSystemSoundID SystemSoundID) {
	_AudioServicesRemoveSystemSoundCompletion(inSystemSoundID)
}

// Sets the value for a specified System Sound Services property.
//
// Added in macOS 10.5.
// Sets the value for a specified System Sound Services property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesSetProperty(_:_:_:_:_:)
func AudioServicesSetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesSetProperty(inPropertyID, inSpecifierSize, inSpecifier, inPropertyDataSize, inPropertyData)
}

// Adds a property listener callback function to your application’s audio session object.

// Adds a property listener callback function to your application’s audio session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionAddPropertyListener(_:_:_:)
func AudioSessionAddPropertyListener(inID AudioSessionPropertyID, inProc AudioSessionPropertyListener, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionAddPropertyListener(inID, inProc, inClientData)
}

// Gets the value of a specified audio session property.

// Gets the value of a specified audio session property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionGetProperty(_:_:_:)
func AudioSessionGetProperty(inID AudioSessionPropertyID, ioDataSize unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionGetProperty(inID, ioDataSize, outData)
}

// Gets the size of the value for a specified audio session property.

// Gets the size of the value for a specified audio session property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionGetPropertySize(_:_:)
func AudioSessionGetPropertySize(inID AudioSessionPropertyID, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionGetPropertySize(inID, outDataSize)
}

// Initializes an iOS application’s audio session object.

// Initializes an iOS application’s audio session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInitialize(_:_:_:_:)
func AudioSessionInitialize(inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inInterruptionListener AudioSessionInterruptionListener, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionInitialize(inRunLoop, inRunLoopMode, inInterruptionListener, inClientData)
}

// Removes an audio session property listener callback function.

// Removes an audio session property listener callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionRemovePropertyListener(_:)
func AudioSessionRemovePropertyListener(inID AudioSessionPropertyID) unsafe.Pointer {
	return _AudioSessionRemovePropertyListener(inID)
}

// Removes a property listener callback function from your application’s audio session object.

// Removes a property listener callback function from your application’s audio session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionRemovePropertyListenerWithUserData(_:_:_:)
func AudioSessionRemovePropertyListenerWithUserData(inID AudioSessionPropertyID, inProc AudioSessionPropertyListener, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionRemovePropertyListenerWithUserData(inID, inProc, inClientData)
}

// Actives or deactivates your application’s audio session.

// Actives or deactivates your application’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetActive(_:)
func AudioSessionSetActive(active unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetActive(active)
}

// Activates or deactivates your application’s audio session; provides flags for use by other audio sessions.

// Activates or deactivates your application’s audio session; provides flags for use by other audio sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetActiveWithFlags(_:_:)
func AudioSessionSetActiveWithFlags(active unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetActiveWithFlags(active, inFlags)
}

// Sets the value of a specified audio session property.

// Sets the value of a specified audio session property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetProperty(_:_:_:)
func AudioSessionSetProperty(inID AudioSessionPropertyID, inDataSize unsafe.Pointer, inData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetProperty(inID, inDataSize, inData)
}

// Registers a callback to receive audio unit property change notifications.
//
// Added in macOS 10.0.
// Registers a callback to receive audio unit property change notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListener(_:_:_:_:)
func AudioUnitAddPropertyListener(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitAddPropertyListener(inUnit, inID, inProc, inProcUserData)
}

// Registers a callback to receive audio unit render notifications.
//
// Added in macOS 10.2.
// Registers a callback to receive audio unit render notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotify(_:_:_:)
func AudioUnitAddRenderNotify(inUnit AudioUnit, inProc RenderCallback, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitAddRenderNotify(inUnit, inProc, inProcUserData)
}

// Returns the component registrations for a given audio unit extension.
//
// Added in macOS 10.13.
// Returns the component registrations for a given audio unit extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionCopyComponentList(_:)
func AudioUnitExtensionCopyComponentList(extensionIdentifier unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitExtensionCopyComponentList(extensionIdentifier)
}

// Allows the implementor of an audio unit extension to dynamically modify the list of component registrations for the extension.
//
// Added in macOS 10.13.
// Allows the implementor of an audio unit extension to dynamically modify the list of component registrations for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionSetComponentList(_:_:)
func AudioUnitExtensionSetComponentList(extensionIdentifier unsafe.Pointer, audioComponentInfo unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitExtensionSetComponentList(extensionIdentifier, audioComponentInfo)
}

// Gets the value of an audio unit parameter.
//
// Added in macOS 10.0.
// Gets the value of an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameter(_:_:_:_:_:)
func AudioUnitGetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, outValue unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetParameter(inUnit, inID, inScope, inElement, outValue)
}

// Gets the value of an audio unit property.
//
// Added in macOS 10.0.
// Gets the value of an audio unit property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetProperty(_:_:_:_:_:_:)
func AudioUnitGetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outData unsafe.Pointer, ioDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetProperty(inUnit, inID, inScope, inElement, outData, ioDataSize)
}

// Gets information about an audio unit property.
//
// Added in macOS 10.0.
// Gets information about an audio unit property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfo(_:_:_:_:_:_:)
func AudioUnitGetPropertyInfo(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetPropertyInfo(inUnit, inID, inScope, inElement, outDataSize, outWritable)
}

// Initializes an audio unit
//
// Added in macOS 10.0.
// Initializes an audio unit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitialize(_:)
func AudioUnitInitialize(inUnit AudioUnit) unsafe.Pointer {
	return _AudioUnitInitialize(inUnit)
}

// AudioUnitProcess is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcess(_:_:_:_:_:)
func AudioUnitProcess(inUnit AudioUnit, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitProcess(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, ioData)
}

// AudioUnitProcessMultiple is a AudioToolbox function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessMultiple(_:_:_:_:_:_:_:_:)
func AudioUnitProcessMultiple(inUnit AudioUnit, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inNumberFrames unsafe.Pointer, inNumberInputBufferLists unsafe.Pointer, inInputBufferLists unsafe.Pointer, inNumberOutputBufferLists unsafe.Pointer, ioOutputBufferLists unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitProcessMultiple(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, inNumberInputBufferLists, inInputBufferLists, inNumberOutputBufferLists, ioOutputBufferLists)
}

// Unregisters a previously-registered property listener callback function.
//
// Added in macOS 10.5.
// Unregisters a previously-registered property listener callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserData(_:_:_:_:)
func AudioUnitRemovePropertyListenerWithUserData(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRemovePropertyListenerWithUserData(inUnit, inID, inProc, inProcUserData)
}

// Unregisters a previously-registered render listener callback function.
//
// Added in macOS 10.2.
// Unregisters a previously-registered render listener callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotify(_:_:_:)
func AudioUnitRemoveRenderNotify(inUnit AudioUnit, inProc RenderCallback, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRemoveRenderNotify(inUnit, inProc, inProcUserData)
}

// Initiates a rendering cycle for an audio unit.
//
// Added in macOS 10.2.
// Initiates a rendering cycle for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRender(_:_:_:_:_:_:)
func AudioUnitRender(inUnit AudioUnit, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inOutputBusNumber unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRender(inUnit, ioActionFlags, inTimeStamp, inOutputBusNumber, inNumberFrames, ioData)
}

// Resets an audio unit’s render state.
//
// Added in macOS 10.0.
// Resets an audio unit’s render state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitReset(_:_:_:)
func AudioUnitReset(inUnit AudioUnit, inScope AudioUnitScope, inElement AudioUnitElement) unsafe.Pointer {
	return _AudioUnitReset(inUnit, inScope, inElement)
}

// Schedules changes to the value of an audio unit parameter.
//
// Added in macOS 10.2.
// Schedules changes to the value of an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParameters(_:_:_:)
func AudioUnitScheduleParameters(inUnit AudioUnit, inParameterEvent unsafe.Pointer, inNumParamEvents unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitScheduleParameters(inUnit, inParameterEvent, inNumParamEvents)
}

// Sets the value of an audio unit parameter.
//
// Added in macOS 10.0.
// Sets the value of an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameter(_:_:_:_:_:_:)
func AudioUnitSetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, inValue AudioUnitParameterValue, inBufferOffsetInFrames unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitSetParameter(inUnit, inID, inScope, inElement, inValue, inBufferOffsetInFrames)
}

// Sets the value of an audio unit property.
//
// Added in macOS 10.0.
// Sets the value of an audio unit property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetProperty(_:_:_:_:_:_:)
func AudioUnitSetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, inData unsafe.Pointer, inDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitSetProperty(inUnit, inID, inScope, inElement, inData, inDataSize)
}

// Uninitializes an audio unit.
//
// Added in macOS 10.0.
// Uninitializes an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitialize(_:)
func AudioUnitUninitialize(inUnit AudioUnit) unsafe.Pointer {
	return _AudioUnitUninitialize(inUnit)
}

// CAClockAddListener is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockAddListener(_:_:_:)
func CAClockAddListener(inCAClock ClockRef, inListenerProc ClockListenerProc, inUserData unsafe.Pointer) unsafe.Pointer {
	return _CAClockAddListener(inCAClock, inListenerProc, inUserData)
}

// CAClockArm is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockArm(_:)
func CAClockArm(inCAClock ClockRef) unsafe.Pointer {
	return _CAClockArm(inCAClock)
}

// CAClockBarBeatTimeToBeats is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBarBeatTimeToBeats(_:_:_:)
func CAClockBarBeatTimeToBeats(inCAClock ClockRef, inBarBeatTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _CAClockBarBeatTimeToBeats(inCAClock, inBarBeatTime, outBeats)
}

// CAClockBeatsToBarBeatTime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBeatsToBarBeatTime(_:_:_:_:)
func CAClockBeatsToBarBeatTime(inCAClock ClockRef, inBeats ClockBeats, inSubbeatDivisor unsafe.Pointer, outBarBeatTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockBeatsToBarBeatTime(inCAClock, inBeats, inSubbeatDivisor, outBarBeatTime)
}

// CAClockDisarm is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockDisarm(_:)
func CAClockDisarm(inCAClock ClockRef) unsafe.Pointer {
	return _CAClockDisarm(inCAClock)
}

// CAClockDispose is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockDispose(_:)
func CAClockDispose(inCAClock ClockRef) unsafe.Pointer {
	return _CAClockDispose(inCAClock)
}

// CAClockGetCurrentTempo is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTempo(_:_:_:)
func CAClockGetCurrentTempo(inCAClock ClockRef, outTempo unsafe.Pointer, outTimestamp unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetCurrentTempo(inCAClock, outTempo, outTimestamp)
}

// CAClockGetCurrentTime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTime(_:_:_:)
func CAClockGetCurrentTime(inCAClock ClockRef, inTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetCurrentTime(inCAClock, inTimeFormat, outTime)
}

// CAClockGetPlayRate is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPlayRate(_:_:)
func CAClockGetPlayRate(inCAClock ClockRef, outPlayRate unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetPlayRate(inCAClock, outPlayRate)
}

// CAClockGetProperty is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetProperty(_:_:_:_:)
func CAClockGetProperty(inCAClock ClockRef, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetProperty(inCAClock, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// CAClockGetPropertyInfo is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPropertyInfo(_:_:_:_:)
func CAClockGetPropertyInfo(inCAClock ClockRef, inPropertyID unsafe.Pointer, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetPropertyInfo(inCAClock, inPropertyID, outSize, outWritable)
}

// CAClockGetStartTime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetStartTime(_:_:_:)
func CAClockGetStartTime(inCAClock ClockRef, inTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetStartTime(inCAClock, inTimeFormat, outTime)
}

// CAClockNew is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockNew(_:_:)
func CAClockNew(inReservedFlags unsafe.Pointer, outCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockNew(inReservedFlags, outCAClock)
}

// CAClockParseMIDI is a AudioToolbox function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockParseMIDI(_:_:)
func CAClockParseMIDI(inCAClock ClockRef, inMIDIPacketList unsafe.Pointer) unsafe.Pointer {
	return _CAClockParseMIDI(inCAClock, inMIDIPacketList)
}

// CAClockRemoveListener is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockRemoveListener(_:_:_:)
func CAClockRemoveListener(inCAClock ClockRef, inListenerProc ClockListenerProc, inUserData unsafe.Pointer) unsafe.Pointer {
	return _CAClockRemoveListener(inCAClock, inListenerProc, inUserData)
}

// CAClockSMPTETimeToSeconds is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTETimeToSeconds(_:_:_:)
func CAClockSMPTETimeToSeconds(inCAClock ClockRef, inSMPTETime unsafe.Pointer, outSeconds unsafe.Pointer) unsafe.Pointer {
	return _CAClockSMPTETimeToSeconds(inCAClock, inSMPTETime, outSeconds)
}

// CAClockSecondsToSMPTETime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSecondsToSMPTETime(_:_:_:_:)
func CAClockSecondsToSMPTETime(inCAClock ClockRef, inSeconds ClockSeconds, inSubframeDivisor unsafe.Pointer, outSMPTETime unsafe.Pointer) unsafe.Pointer {
	return _CAClockSecondsToSMPTETime(inCAClock, inSeconds, inSubframeDivisor, outSMPTETime)
}

// CAClockSetCurrentTempo is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTempo(_:_:_:)
func CAClockSetCurrentTempo(inCAClock ClockRef, inTempo ClockTempo, inTimestamp unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetCurrentTempo(inCAClock, inTempo, inTimestamp)
}

// CAClockSetCurrentTime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTime(_:_:)
func CAClockSetCurrentTime(inCAClock ClockRef, inTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetCurrentTime(inCAClock, inTime)
}

// CAClockSetPlayRate is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetPlayRate(_:_:)
func CAClockSetPlayRate(inCAClock ClockRef, inPlayRate unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetPlayRate(inCAClock, inPlayRate)
}

// CAClockSetProperty is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetProperty(_:_:_:_:)
func CAClockSetProperty(inCAClock ClockRef, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetProperty(inCAClock, inPropertyID, inPropertyDataSize, inPropertyData)
}

// CAClockStart is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockStart(_:)
func CAClockStart(inCAClock ClockRef) unsafe.Pointer {
	return _CAClockStart(inCAClock)
}

// CAClockStop is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockStop(_:)
func CAClockStop(inCAClock ClockRef) unsafe.Pointer {
	return _CAClockStop(inCAClock)
}

// CAClockTranslateTime is a AudioToolbox function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTranslateTime(_:_:_:_:)
func CAClockTranslateTime(inCAClock ClockRef, inTime unsafe.Pointer, inOutputTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockTranslateTime(inCAClock, inTime, inOutputTimeFormat, outTime)
}

// Prints the internal state of an object to .
//
// Added in macOS 10.2.
// Prints the internal state of an object to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAShow(_:)
func CAShow(inObject unsafe.Pointer) {
	_CAShow(inObject)
}

// Prints the internal state of an object to a file.
//
// Added in macOS 10.2.
// Prints the internal state of an object to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAShowFile(_:_:)
func CAShowFile(inObject unsafe.Pointer, inFile unsafe.Pointer) {
	_CAShowFile(inObject, inFile)
}

// CopyInstrumentInfoFromSoundBank is a AudioToolbox function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CopyInstrumentInfoFromSoundBank(_:_:)
func CopyInstrumentInfoFromSoundBank(inURL unsafe.Pointer, outInstrumentInfo unsafe.Pointer) unsafe.Pointer {
	return _CopyInstrumentInfoFromSoundBank(inURL, outInstrumentInfo)
}

// Copies the name of a sound bank from a sound bank file at a specified URL.
//
// Added in macOS 10.5.
// Copies the name of a sound bank from a sound bank file at a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CopyNameFromSoundBank(_:_:)
func CopyNameFromSoundBank(inURL unsafe.Pointer, outName unsafe.Pointer) unsafe.Pointer {
	return _CopyNameFromSoundBank(inURL, outName)
}

// Disposes of a music event iterator.
//
// Added in macOS 10.0.
// Disposes of a music event iterator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicEventIterator(_:)
func DisposeMusicEventIterator(inIterator MusicEventIterator) unsafe.Pointer {
	return _DisposeMusicEventIterator(inIterator)
}

// Disposes of a music player.
//
// Added in macOS 10.0.
// Disposes of a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicPlayer(_:)
func DisposeMusicPlayer(inPlayer MusicPlayer) unsafe.Pointer {
	return _DisposeMusicPlayer(inPlayer)
}

// Disposes of a music sequence.
//
// Added in macOS 10.0.
// Disposes of a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicSequence(_:)
func DisposeMusicSequence(inSequence MusicSequence) unsafe.Pointer {
	return _DisposeMusicSequence(inSequence)
}

// Deprecated. Use the function instead.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
// Deprecated. Use the function instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateNew
func ExtAudioFileCreateNew(inParentDir unsafe.Pointer, inFileName unsafe.Pointer, inFileType AudioFileTypeID, inStreamDesc unsafe.Pointer, inChannelLayout unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileCreateNew(inParentDir, inFileName, inFileType, inStreamDesc, inChannelLayout, outExtAudioFile)
}

// Creates a new audio file and associates it with a new extended audio file object.
//
// Added in macOS 10.5.
// Creates a new audio file and associates it with a new extended audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateWithURL(_:_:_:_:_:_:)
func ExtAudioFileCreateWithURL(inURL unsafe.Pointer, inFileType AudioFileTypeID, inStreamDesc unsafe.Pointer, inChannelLayout unsafe.Pointer, inFlags unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileCreateWithURL(inURL, inFileType, inStreamDesc, inChannelLayout, inFlags, outExtAudioFile)
}

// Disposes of an extended audio file object and closes the associated file.
//
// Added in macOS 10.4.
// Disposes of an extended audio file object and closes the associated file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileDispose(_:)
func ExtAudioFileDispose(inExtAudioFile ExtAudioFileRef) unsafe.Pointer {
	return _ExtAudioFileDispose(inExtAudioFile)
}

// Gets a property value from an extended audio file object.
//
// Added in macOS 10.4.
// Gets a property value from an extended audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetProperty(_:_:_:_:)
func ExtAudioFileGetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileGetProperty(inExtAudioFile, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// Gets information about an extended audio file object property.
//
// Added in macOS 10.4.
// Gets information about an extended audio file object property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetPropertyInfo(_:_:_:_:)
func ExtAudioFileGetPropertyInfo(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileGetPropertyInfo(inExtAudioFile, inPropertyID, outSize, outWritable)
}

// Deprecated. Use the function instead.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
// Deprecated. Use the function instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpen
func ExtAudioFileOpen(inFSRef unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileOpen(inFSRef, outExtAudioFile)
}

// Opens an existing audio file for reading, and associates it with a new extended audio file object.
//
// Added in macOS 10.5.
// Opens an existing audio file for reading, and associates it with a new extended audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpenURL(_:_:)
func ExtAudioFileOpenURL(inURL unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileOpenURL(inURL, outExtAudioFile)
}

// Performs a synchronous, sequential read operation on an audio file.
//
// Added in macOS 10.4.
// Performs a synchronous, sequential read operation on an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileRead(_:_:_:)
func ExtAudioFileRead(inExtAudioFile ExtAudioFileRef, ioNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileRead(inExtAudioFile, ioNumberFrames, ioData)
}

// Seeks to a specified frame in a file.
//
// Added in macOS 10.4.
// Seeks to a specified frame in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSeek(_:_:)
func ExtAudioFileSeek(inExtAudioFile ExtAudioFileRef, inFrameOffset unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileSeek(inExtAudioFile, inFrameOffset)
}

// Sets a property value for an extended audio file object.
//
// Added in macOS 10.4.
// Sets a property value for an extended audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSetProperty(_:_:_:_:)
func ExtAudioFileSetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileSetProperty(inExtAudioFile, inPropertyID, inPropertyDataSize, inPropertyData)
}

// Gets an audio file’s read/write position.
//
// Added in macOS 10.4.
// Gets an audio file’s read/write position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileTell(_:_:)
func ExtAudioFileTell(inExtAudioFile ExtAudioFileRef, outFrameOffset unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileTell(inExtAudioFile, outFrameOffset)
}

// Wraps an audio file object in an extended audio file object.
//
// Added in macOS 10.4.
// Wraps an audio file object in an extended audio file object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrapAudioFileID(_:_:_:)
func ExtAudioFileWrapAudioFileID(inFileID AudioFileID, inForWriting unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWrapAudioFileID(inFileID, inForWriting, outExtAudioFile)
}

// Performs a synchronous, sequential write operation on an audio file.
//
// Added in macOS 10.4.
// Performs a synchronous, sequential write operation on an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrite(_:_:_:)
func ExtAudioFileWrite(inExtAudioFile ExtAudioFileRef, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWrite(inExtAudioFile, inNumberFrames, ioData)
}

// Perform an asynchronous, sequential write operation on an audio file.
//
// Added in macOS 10.4.
// Perform an asynchronous, sequential write operation on an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWriteAsync(_:_:_:)
func ExtAudioFileWriteAsync(inExtAudioFile ExtAudioFileRef, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWriteAsync(inExtAudioFile, inNumberFrames, ioData)
}

// Gets the name of a sound bank from a sound bank file.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.2.
// Gets the name of a sound bank from a sound bank file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetNameFromSoundBank
func GetNameFromSoundBank(inSoundBankRef unsafe.Pointer, outName unsafe.Pointer) unsafe.Pointer {
	return _GetNameFromSoundBank(inSoundBankRef, outName)
}

// MusicDeviceMIDIEvent is a AudioToolbox function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEvent(_:_:_:_:_:)
func MusicDeviceMIDIEvent(inUnit MusicDeviceComponent, inStatus unsafe.Pointer, inData1 unsafe.Pointer, inData2 unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceMIDIEvent(inUnit, inStatus, inData1, inData2, inOffsetSampleFrame)
}

// MusicDeviceMIDIEventList is a AudioToolbox function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEventList(_:_:_:)
func MusicDeviceMIDIEventList(inUnit MusicDeviceComponent, inOffsetSampleFrame unsafe.Pointer, evtList unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceMIDIEventList(inUnit, inOffsetSampleFrame, evtList)
}

// MusicDevicePrepareInstrument is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDevicePrepareInstrument
func MusicDevicePrepareInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) unsafe.Pointer {
	return _MusicDevicePrepareInstrument(inUnit, inInstrument)
}

// MusicDeviceReleaseInstrument is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceReleaseInstrument
func MusicDeviceReleaseInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) unsafe.Pointer {
	return _MusicDeviceReleaseInstrument(inUnit, inInstrument)
}

// MusicDeviceStartNote is a AudioToolbox function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNote(_:_:_:_:_:_:)
func MusicDeviceStartNote(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID, inGroupID MusicDeviceGroupID, outNoteInstanceID unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer, inParams unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceStartNote(inUnit, inInstrument, inGroupID, outNoteInstanceID, inOffsetSampleFrame, inParams)
}

// MusicDeviceStopNote is a AudioToolbox function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNote(_:_:_:_:)
func MusicDeviceStopNote(inUnit MusicDeviceComponent, inGroupID MusicDeviceGroupID, inNoteInstanceID NoteInstanceID, inOffsetSampleFrame unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceStopNote(inUnit, inGroupID, inNoteInstanceID, inOffsetSampleFrame)
}

// MusicDeviceSysEx is a AudioToolbox function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysEx(_:_:_:)
func MusicDeviceSysEx(inUnit MusicDeviceComponent, inData unsafe.Pointer, inLength unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceSysEx(inUnit, inData, inLength)
}

// Deletes the event at a music event iterator’s current position.
//
// Added in macOS 10.0.
// Deletes the event at a music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorDeleteEvent(_:)
func MusicEventIteratorDeleteEvent(inIterator MusicEventIterator) unsafe.Pointer {
	return _MusicEventIteratorDeleteEvent(inIterator)
}

// Gets information about the event at a music event iterator’s current position.
//
// Added in macOS 10.0.
// Gets information about the event at a music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorGetEventInfo(_:_:_:_:_:)
func MusicEventIteratorGetEventInfo(inIterator MusicEventIterator, outTimeStamp unsafe.Pointer, outEventType unsafe.Pointer, outEventData unsafe.Pointer, outEventDataSize unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorGetEventInfo(inIterator, outTimeStamp, outEventType, outEventData, outEventDataSize)
}

// Indicates whether or not a music track contains an event at the music event iterator’s current position.
//
// Added in macOS 10.2.
// Indicates whether or not a music track contains an event at the music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasCurrentEvent(_:_:)
func MusicEventIteratorHasCurrentEvent(inIterator MusicEventIterator, outHasCurEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasCurrentEvent(inIterator, outHasCurEvent)
}

// Indicates whether or not a music track contains an event beyond the music event iterator’s current position.
//
// Added in macOS 10.0.
// Indicates whether or not a music track contains an event beyond the music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasNextEvent(_:_:)
func MusicEventIteratorHasNextEvent(inIterator MusicEventIterator, outHasNextEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasNextEvent(inIterator, outHasNextEvent)
}

// Indicates whether or not a music track contains an event before the music event iterator’s current position.
//
// Added in macOS 10.0.
// Indicates whether or not a music track contains an event before the music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasPreviousEvent(_:_:)
func MusicEventIteratorHasPreviousEvent(inIterator MusicEventIterator, outHasPrevEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasPreviousEvent(inIterator, outHasPrevEvent)
}

// Positions a music event iterator at the next event on a music track.
//
// Added in macOS 10.0.
// Positions a music event iterator at the next event on a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorNextEvent(_:)
func MusicEventIteratorNextEvent(inIterator MusicEventIterator) unsafe.Pointer {
	return _MusicEventIteratorNextEvent(inIterator)
}

// Positions a music event iterator at the previous event on a music track.
//
// Added in macOS 10.0.
// Positions a music event iterator at the previous event on a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorPreviousEvent(_:)
func MusicEventIteratorPreviousEvent(inIterator MusicEventIterator) unsafe.Pointer {
	return _MusicEventIteratorPreviousEvent(inIterator)
}

// Positions a music event iterator at a specified timestamp, in beats.
//
// Added in macOS 10.0.
// Positions a music event iterator at a specified timestamp, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSeek(_:_:)
func MusicEventIteratorSeek(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) unsafe.Pointer {
	return _MusicEventIteratorSeek(inIterator, inTimeStamp)
}

// Sets information for the event at a music event iterator’s current position.
//
// Added in macOS 10.2.
// Sets information for the event at a music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventInfo(_:_:_:)
func MusicEventIteratorSetEventInfo(inIterator MusicEventIterator, inEventType MusicEventType, inEventData unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorSetEventInfo(inIterator, inEventType, inEventData)
}

// Sets the timestamp for the event at a music event iterator’s current position.
//
// Added in macOS 10.0.
// Sets the timestamp for the event at a music event iterator’s current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventTime(_:_:)
func MusicEventIteratorSetEventTime(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) unsafe.Pointer {
	return _MusicEventIteratorSetEventTime(inIterator, inTimeStamp)
}

// Gets the beat number associated a specified host time.
//
// Added in macOS 10.2.
// Gets the beat number associated a specified host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetBeatsForHostTime(_:_:_:)
func MusicPlayerGetBeatsForHostTime(inPlayer MusicPlayer, inHostTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetBeatsForHostTime(inPlayer, inHostTime, outBeats)
}

// Gets the host time associated with a specified beat.
//
// Added in macOS 10.2.
// Gets the host time associated with a specified beat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetHostTimeForBeats(_:_:_:)
func MusicPlayerGetHostTimeForBeats(inPlayer MusicPlayer, inBeats MusicTimeStamp, outHostTime unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetHostTimeForBeats(inPlayer, inBeats, outHostTime)
}

// Gets the playback rate multiplier for a music player.
//
// Added in macOS 10.3.
// Gets the playback rate multiplier for a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetPlayRateScalar(_:_:)
func MusicPlayerGetPlayRateScalar(inPlayer MusicPlayer, outScaleRate unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetPlayRateScalar(inPlayer, outScaleRate)
}

// Gets the music sequence associated with a music player.
//
// Added in macOS 10.3.
// Gets the music sequence associated with a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetSequence(_:_:)
func MusicPlayerGetSequence(inPlayer MusicPlayer, outSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetSequence(inPlayer, outSequence)
}

// Gets the playback point for a music player, in beats.
//
// Added in macOS 10.0.
// Gets the playback point for a music player, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetTime(_:_:)
func MusicPlayerGetTime(inPlayer MusicPlayer, outTime unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetTime(inPlayer, outTime)
}

// Indicates whether or not a music player is playing.
//
// Added in macOS 10.2.
// Indicates whether or not a music player is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerIsPlaying(_:_:)
func MusicPlayerIsPlaying(inPlayer MusicPlayer, outIsPlaying unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerIsPlaying(inPlayer, outIsPlaying)
}

// Prepares a music player to play.
//
// Added in macOS 10.0.
// Prepares a music player to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerPreroll(_:)
func MusicPlayerPreroll(inPlayer MusicPlayer) unsafe.Pointer {
	return _MusicPlayerPreroll(inPlayer)
}

// Sets a playback rate multiplier for a music player.
//
// Added in macOS 10.3.
// Sets a playback rate multiplier for a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetPlayRateScalar(_:_:)
func MusicPlayerSetPlayRateScalar(inPlayer MusicPlayer, inScaleRate unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerSetPlayRateScalar(inPlayer, inScaleRate)
}

// Sets the music sequence for the music player to play.
//
// Added in macOS 10.0.
// Sets the music sequence for the music player to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetSequence(_:_:)
func MusicPlayerSetSequence(inPlayer MusicPlayer, inSequence MusicSequence) unsafe.Pointer {
	return _MusicPlayerSetSequence(inPlayer, inSequence)
}

// Sets the playback point for a music player, in beats.
//
// Added in macOS 10.0.
// Sets the playback point for a music player, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetTime(_:_:)
func MusicPlayerSetTime(inPlayer MusicPlayer, inTime MusicTimeStamp) unsafe.Pointer {
	return _MusicPlayerSetTime(inPlayer, inTime)
}

// Starts playback of a music player.
//
// Added in macOS 10.0.
// Starts playback of a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStart(_:)
func MusicPlayerStart(inPlayer MusicPlayer) unsafe.Pointer {
	return _MusicPlayerStart(inPlayer)
}

// Stops playback of a music player.
//
// Added in macOS 10.0.
// Stops playback of a music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStop(_:)
func MusicPlayerStop(inPlayer MusicPlayer) unsafe.Pointer {
	return _MusicPlayerStop(inPlayer)
}

// Formats a music sequence’s bar-beat time to its beat time.
//
// Added in macOS 10.5.
// Formats a music sequence’s bar-beat time to its beat time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBarBeatTimeToBeats(_:_:_:)
func MusicSequenceBarBeatTimeToBeats(inSequence MusicSequence, inBarBeatTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceBarBeatTimeToBeats(inSequence, inBarBeatTime, outBeats)
}

// Formats a music sequence’s beat time to its bar-beat time.
//
// Added in macOS 10.5.
// Formats a music sequence’s beat time to its bar-beat time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBeatsToBarBeatTime(_:_:_:_:)
func MusicSequenceBeatsToBarBeatTime(inSequence MusicSequence, inBeats MusicTimeStamp, inSubbeatDivisor unsafe.Pointer, outBarBeatTime unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceBeatsToBarBeatTime(inSequence, inBeats, inSubbeatDivisor, outBarBeatTime)
}

// Removes a music track from a music sequence, and disposes of the track.
//
// Added in macOS 10.0.
// Removes a music track from a music sequence, and disposes of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceDisposeTrack(_:_:)
func MusicSequenceDisposeTrack(inSequence MusicSequence, inTrack MusicTrack) unsafe.Pointer {
	return _MusicSequenceDisposeTrack(inSequence, inTrack)
}

// Creates a MIDI file from the events in a music sequence.
//
// Added in macOS 10.5.
// Creates a MIDI file from the events in a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreate(_:_:_:_:_:)
func MusicSequenceFileCreate(inSequence MusicSequence, inFileRef unsafe.Pointer, inFileType unsafe.Pointer, inFlags unsafe.Pointer, inResolution unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileCreate(inSequence, inFileRef, inFileType, inFlags, inResolution)
}

// Creates a data object containing the events from a music sequence.
//
// Added in macOS 10.5.
// Creates a data object containing the events from a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreateData(_:_:_:_:_:)
func MusicSequenceFileCreateData(inSequence MusicSequence, inFileType unsafe.Pointer, inFlags unsafe.Pointer, inResolution unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileCreateData(inSequence, inFileType, inFlags, inResolution, outData)
}

// Loads data into a music sequence from a URL reference.
//
// Added in macOS 10.5.
// Loads data into a music sequence from a URL reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoad(_:_:_:_:)
func MusicSequenceFileLoad(inSequence MusicSequence, inFileRef unsafe.Pointer, inFileTypeHint unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileLoad(inSequence, inFileRef, inFileTypeHint, inFlags)
}

// Load data into a music sequence from a data reference.
//
// Added in macOS 10.5.
// Load data into a music sequence from a data reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoadData(_:_:_:_:)
func MusicSequenceFileLoadData(inSequence MusicSequence, inData unsafe.Pointer, inFileTypeHint unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileLoadData(inSequence, inData, inFileTypeHint, inFlags)
}

// Gets the audio processing graph associated with a music sequence.
//
// Added in macOS 10.0.
// Gets the audio processing graph associated with a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetAUGraph(_:_:)
func MusicSequenceGetAUGraph(inSequence MusicSequence, outGraph unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetAUGraph(inSequence, outGraph)
}

// Calculates the number of beats that correspond to a number of seconds.
//
// Added in macOS 10.2.
// Calculates the number of beats that correspond to a number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetBeatsForSeconds(_:_:_:)
func MusicSequenceGetBeatsForSeconds(inSequence MusicSequence, inSeconds unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetBeatsForSeconds(inSequence, inSeconds, outBeats)
}

// Gets the music track at the specified track index.
//
// Added in macOS 10.0.
// Gets the music track at the specified track index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetIndTrack(_:_:_:)
func MusicSequenceGetIndTrack(inSequence MusicSequence, inTrackIndex unsafe.Pointer, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetIndTrack(inSequence, inTrackIndex, outTrack)
}

// Returns a dictionary containing music sequence information.
//
// Added in macOS 10.5.
// Returns a dictionary containing music sequence information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetInfoDictionary(_:)
func MusicSequenceGetInfoDictionary(inSequence MusicSequence) unsafe.Pointer {
	return _MusicSequenceGetInfoDictionary(inSequence)
}

// Calculates the number of seconds that correspond to a number of beats.
//
// Added in macOS 10.2.
// Calculates the number of seconds that correspond to a number of beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSecondsForBeats(_:_:_:)
func MusicSequenceGetSecondsForBeats(inSequence MusicSequence, inBeats MusicTimeStamp, outSeconds unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetSecondsForBeats(inSequence, inBeats, outSeconds)
}

// Gets the sequence type for a music sequence.
//
// Added in macOS 10.5.
// Gets the sequence type for a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSequenceType(_:_:)
func MusicSequenceGetSequenceType(inSequence MusicSequence, outType unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetSequenceType(inSequence, outType)
}

// Gets the tempo track for a music sequence.
//
// Added in macOS 10.1.
// Gets the tempo track for a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTempoTrack(_:_:)
func MusicSequenceGetTempoTrack(inSequence MusicSequence, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTempoTrack(inSequence, outTrack)
}

// Gets the number of music tracks owned by a music sequence.
//
// Added in macOS 10.0.
// Gets the number of music tracks owned by a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackCount(_:_:)
func MusicSequenceGetTrackCount(inSequence MusicSequence, outNumberOfTracks unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTrackCount(inSequence, outNumberOfTracks)
}

// Gets the index number for a specified music track.
//
// Added in macOS 10.0.
// Gets the index number for a specified music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackIndex(_:_:_:)
func MusicSequenceGetTrackIndex(inSequence MusicSequence, inTrack MusicTrack, outTrackIndex unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTrackIndex(inSequence, inTrack, outTrackIndex)
}

// MusicSequenceLoadSMFDataWithFlags is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFDataWithFlags
func MusicSequenceLoadSMFDataWithFlags(inSequence MusicSequence, inData unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceLoadSMFDataWithFlags(inSequence, inData, inFlags)
}

// MusicSequenceLoadSMFWithFlags is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFWithFlags
func MusicSequenceLoadSMFWithFlags(inSequence MusicSequence, inFileRef unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceLoadSMFWithFlags(inSequence, inFileRef, inFlags)
}

// Add a new, empty music track to a music sequence.
//
// Added in macOS 10.0.
// Add a new, empty music track to a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceNewTrack(_:_:)
func MusicSequenceNewTrack(inSequence MusicSequence, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceNewTrack(inSequence, outTrack)
}

// Reverses the MIDI and tempo events in a music sequence, so the start becomes the end.
//
// Added in macOS 10.0.
// Reverses the MIDI and tempo events in a music sequence, so the start becomes the end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceReverse(_:)
func MusicSequenceReverse(inSequence MusicSequence) unsafe.Pointer {
	return _MusicSequenceReverse(inSequence)
}

// MusicSequenceSaveMIDIFile is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveMIDIFile
func MusicSequenceSaveMIDIFile(inSequence MusicSequence, inParentDirectory unsafe.Pointer, inFileName unsafe.Pointer, inResolution unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSaveMIDIFile(inSequence, inParentDirectory, inFileName, inResolution, inFlags)
}

// MusicSequenceSaveSMFData is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveSMFData
func MusicSequenceSaveSMFData(inSequence MusicSequence, outData unsafe.Pointer, inResolution unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSaveSMFData(inSequence, outData, inResolution)
}

// Associates an audio processing graph with a music sequence.
//
// Added in macOS 10.0.
// Associates an audio processing graph with a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetAUGraph(_:_:)
func MusicSequenceSetAUGraph(inSequence MusicSequence, inGraph unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetAUGraph(inSequence, inGraph)
}

// Associates a specified MIDI endpoint with all music tracks in a music sequence.
//
// Added in macOS 10.1.
// Associates a specified MIDI endpoint with all music tracks in a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetMIDIEndpoint(_:_:)
func MusicSequenceSetMIDIEndpoint(inSequence MusicSequence, inEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetMIDIEndpoint(inSequence, inEndpoint)
}

// Sets the sequence type for a music sequence.
//
// Added in macOS 10.5.
// Sets the sequence type for a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetSequenceType(_:_:)
func MusicSequenceSetSequenceType(inSequence MusicSequence, inType unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetSequenceType(inSequence, inType)
}

// Registers a user callback function with a music sequence.
//
// Added in macOS 10.3.
// Registers a user callback function with a music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetUserCallback(_:_:_:)
func MusicSequenceSetUserCallback(inSequence MusicSequence, inCallback MusicSequenceUserCallback, inClientData unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetUserCallback(inSequence, inCallback, inClientData)
}

// Removes a specified range of music track events.
//
// Added in macOS 10.0.
// Removes a specified range of music track events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackClear(_:_:_:)
func MusicTrackClear(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) unsafe.Pointer {
	return _MusicTrackClear(inTrack, inStartTime, inEndTime)
}

// Copies a range of events from one music track and inserts them into another music track.
//
// Added in macOS 10.0.
// Copies a range of events from one music track and inserts them into another music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCopyInsert(_:_:_:_:_:)
func MusicTrackCopyInsert(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) unsafe.Pointer {
	return _MusicTrackCopyInsert(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
}

// Removes a specified range of music track events, and shifts later events toward the start of the track to fill in the gap.
//
// Added in macOS 10.0.
// Removes a specified range of music track events, and shifts later events toward the start of the track to fill in the gap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCut(_:_:_:)
func MusicTrackCut(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) unsafe.Pointer {
	return _MusicTrackCut(inTrack, inStartTime, inEndTime)
}

// Gets the MIDI endpoint that is the event target for a music track.
//
// Added in macOS 10.1.
// Gets the MIDI endpoint that is the event target for a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestMIDIEndpoint(_:_:)
func MusicTrackGetDestMIDIEndpoint(inTrack MusicTrack, outEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetDestMIDIEndpoint(inTrack, outEndpoint)
}

// Gets the audio unit node that is the event target for a music track.
//
// Added in macOS 10.1.
// Gets the audio unit node that is the event target for a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestNode(_:_:)
func MusicTrackGetDestNode(inTrack MusicTrack, outNode unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetDestNode(inTrack, outNode)
}

// Gets a music track property value.
//
// Added in macOS 10.0.
// Gets a music track property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetProperty(_:_:_:_:)
func MusicTrackGetProperty(inTrack MusicTrack, inPropertyID unsafe.Pointer, outData unsafe.Pointer, ioLength unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetProperty(inTrack, inPropertyID, outData, ioLength)
}

// Gets the music sequence that the music track is a member of.
//
// Added in macOS 10.0.
// Gets the music sequence that the music track is a member of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetSequence(_:_:)
func MusicTrackGetSequence(inTrack MusicTrack, outSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetSequence(inTrack, outSequence)
}

// Copies a range of events from one music track and merges them into another music track.
//
// Added in macOS 10.0.
// Copies a range of events from one music track and merges them into another music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMerge(_:_:_:_:_:)
func MusicTrackMerge(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) unsafe.Pointer {
	return _MusicTrackMerge(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
}

// Shifts music track events forward or backward in time, in terms of beats.
//
// Added in macOS 10.0.
// Shifts music track events forward or backward in time, in terms of beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMoveEvents(_:_:_:_:)
func MusicTrackMoveEvents(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp, inMoveTime MusicTimeStamp) unsafe.Pointer {
	return _MusicTrackMoveEvents(inTrack, inStartTime, inEndTime, inMoveTime)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.3.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewAUPresetEvent(_:_:_:)
func MusicTrackNewAUPresetEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inPresetEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewAUPresetEvent(inTrack, inTimeStamp, inPresetEvent)
}

// MusicTrackNewExtendedControlEvent is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedControlEvent
func MusicTrackNewExtendedControlEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedControlEvent(inTrack, inTimeStamp, inInfo)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedNoteEvent(_:_:_:)
func MusicTrackNewExtendedNoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedNoteEvent(inTrack, inTimeStamp, inInfo)
}

// Adds a tempo to a music track.
//
// Added in macOS 10.0.
// Adds a tempo to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedTempoEvent(_:_:_:)
func MusicTrackNewExtendedTempoEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inBPM unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedTempoEvent(inTrack, inTimeStamp, inBPM)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIChannelEvent(_:_:_:)
func MusicTrackNewMIDIChannelEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDIChannelEvent(inTrack, inTimeStamp, inMessage)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDINoteEvent(_:_:_:)
func MusicTrackNewMIDINoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDINoteEvent(inTrack, inTimeStamp, inMessage)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIRawDataEvent(_:_:_:)
func MusicTrackNewMIDIRawDataEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inRawData unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDIRawDataEvent(inTrack, inTimeStamp, inRawData)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMetaEvent(_:_:_:)
func MusicTrackNewMetaEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMetaEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMetaEvent(inTrack, inTimeStamp, inMetaEvent)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.2.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewParameterEvent(_:_:_:)
func MusicTrackNewParameterEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewParameterEvent(inTrack, inTimeStamp, inInfo)
}

// Adds an event of type to a music track.
//
// Added in macOS 10.0.
// Adds an event of type to a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewUserEvent(_:_:_:)
func MusicTrackNewUserEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inUserData unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewUserEvent(inTrack, inTimeStamp, inUserData)
}

// Sets the music track’s event target to a MIDI endpoint.
//
// Added in macOS 10.1.
// Sets the music track’s event target to a MIDI endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestMIDIEndpoint(_:_:)
func MusicTrackSetDestMIDIEndpoint(inTrack MusicTrack, inEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackSetDestMIDIEndpoint(inTrack, inEndpoint)
}

// Sets the music track’s event target to an audio unit node.
//
// Added in macOS 10.0.
// Sets the music track’s event target to an audio unit node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestNode(_:_:)
func MusicTrackSetDestNode(inTrack MusicTrack, inNode Node) unsafe.Pointer {
	return _MusicTrackSetDestNode(inTrack, inNode)
}

// Sets a music track property value.
//
// Added in macOS 10.0.
// Sets a music track property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetProperty(_:_:_:_:)
func MusicTrackSetProperty(inTrack MusicTrack, inPropertyID unsafe.Pointer, inData unsafe.Pointer, inLength unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackSetProperty(inTrack, inPropertyID, inData, inLength)
}

// Creates a new audio processing graph.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
// Creates a new audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewAUGraph(_:)
func NewAUGraph(outGraph unsafe.Pointer) unsafe.Pointer {
	return _NewAUGraph(outGraph)
}

// Creates a new music event iterator.
//
// Added in macOS 10.0.
// Creates a new music event iterator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicEventIterator(_:_:)
func NewMusicEventIterator(inTrack MusicTrack, outIterator unsafe.Pointer) unsafe.Pointer {
	return _NewMusicEventIterator(inTrack, outIterator)
}

// Creates a new music player.
//
// Added in macOS 10.0.
// Creates a new music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicPlayer(_:)
func NewMusicPlayer(outPlayer unsafe.Pointer) unsafe.Pointer {
	return _NewMusicPlayer(outPlayer)
}

// Creates a new empty music sequence.
//
// Added in macOS 10.0.
// Creates a new empty music sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicSequence(_:)
func NewMusicSequence(outSequence unsafe.Pointer) unsafe.Pointer {
	return _NewMusicSequence(outSequence)
}

// NewMusicTrackFrom is a AudioToolbox function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicTrackFrom
func NewMusicTrackFrom(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, outNewTrack unsafe.Pointer) unsafe.Pointer {
	return _NewMusicTrackFrom(inSourceTrack, inSourceStartTime, inSourceEndTime, outNewTrack)
}

// GetAudioUnitParameterDisplayType is a AudioToolbox function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetAudioUnitParameterDisplayType(_:)
func GetAudioUnitParameterDisplayType(p0 unsafe.Pointer) unsafe.Pointer {
	return _GetAudioUnitParameterDisplayType(p0)
}

// SetAudioUnitParameterDisplayType is a AudioToolbox function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetAudioUnitParameterDisplayType(_:_:)
func SetAudioUnitParameterDisplayType(p0 unsafe.Pointer) unsafe.Pointer {
	return _SetAudioUnitParameterDisplayType(p0)
}



