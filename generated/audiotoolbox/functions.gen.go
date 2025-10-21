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
	_AUEventListenerAddEventType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerCreateWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerNotify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUEventListenerRemoveEventType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphAddNode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetInteractionInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetNodeInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphGetNodeInteractions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUGraphNewNode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerAddParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerCreateWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUListenerDispose func(unsafe.Pointer) unsafe.Pointer
	_AUListenerRemoveParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterFormatValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterListenerNotify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterSet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterValueFromLinear func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AUParameterValueToLinear func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecAppendInputBufferList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecAppendInputData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecInitialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecProduceOutputBufferList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecProduceOutputPackets func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecReset func(unsafe.Pointer) unsafe.Pointer
	_AudioCodecSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioCodecUninitialize func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentCopyConfigurationInfo func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentCopyIcon func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentCopyName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentCount func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentFindNext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentGetDescription func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentGetIcon func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentGetLastActiveTime func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentGetVersion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceCanDo func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceDispose func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceGetComponent func(unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstanceNew func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentInstantiate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentRegister func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentValidate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioComponentValidateWithResults func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterConvertBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterConvertComplexBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterDispose func(unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBufferRealtimeSafe func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterFillComplexBufferWithPacketDependencies func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNew func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNewSpecific func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterNewWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterPrepare func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioConverterReset func(unsafe.Pointer) unsafe.Pointer
	_AudioConverterSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileClose func(unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCloseFile func(unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCountUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentCreateURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentDataIsThisFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentExtensionIsThisFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentFileDataIsThisFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentFileIsThisFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetGlobalInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetGlobalInfoSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataAtOffset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentGetUserDataSize64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentInitialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentInitializeWithCallbacks func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOpenFile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOpenURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOpenWithCallbacks func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentOptimize func(unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentReadBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentReadPacketData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentReadPackets func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentRemoveUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentSetUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentWriteBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileComponentWritePackets func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCountUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileCreateWithURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetGlobalInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetGlobalInfoSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataAtOffset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileGetUserDataSize64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileInitialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileInitializeWithCallbacks func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpenURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileOpenWithCallbacks func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileOptimize func(unsafe.Pointer) unsafe.Pointer
	_AudioFileReadBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileReadPacketData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileReadPackets func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileRemoveUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileSetUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamClose func(unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamParseBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamSeek func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileStreamSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWriteBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWritePackets func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFileWritePacketsWithDependencies func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFormatGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioFormatGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceAddPropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceGetPropertyData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceGetPropertyDataSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceHasProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceIsPropertySettable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceRemovePropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareServiceSetPropertyData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioOutputUnitGetHostIcon func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioOutputUnitPublish func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioOutputUnitStart func(unsafe.Pointer) unsafe.Pointer
	_AudioOutputUnitStop func(unsafe.Pointer) unsafe.Pointer
	_AudioQueueAddPropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueAllocateBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueAllocateBufferWithPacketDescriptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueCreateTimeline func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceGetCurrentTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceGetNearestStartTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDeviceTranslateTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDispose func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueDisposeTimeline func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueEnqueueBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueEnqueueBufferWithParameters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueFlush func(unsafe.Pointer) unsafe.Pointer
	_AudioQueueFreeBuffer func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetCurrentTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueGetPropertySize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewInput func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewInputWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewOutput func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueNewOutputWithDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueOfflineRender func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueuePause func(unsafe.Pointer) unsafe.Pointer
	_AudioQueuePrime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapDispose func(unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapGetQueueTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapGetSourceAudio func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueProcessingTapNew func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueRemovePropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueReset func(unsafe.Pointer) unsafe.Pointer
	_AudioQueueSetOfflineRenderFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueSetParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueStart func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioQueueStop func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesAddSystemSoundCompletion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesCreateSystemSoundID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesDisposeSystemSoundID func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlayAlertSound func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlayAlertSoundWithCompletion func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlayAlertSoundWithDetails func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlaySystemSound func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlaySystemSoundWithCompletion func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesPlaySystemSoundWithDetails func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioServicesRemoveSystemSoundCompletion func(unsafe.Pointer) unsafe.Pointer
	_AudioServicesSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionAddPropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionGetPropertySize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionInitialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionRemovePropertyListener func(unsafe.Pointer) unsafe.Pointer
	_AudioSessionRemovePropertyListenerWithUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetActive func(unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetActiveWithFlags func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioSessionSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitAddPropertyListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitAddRenderNotify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitExtensionCopyComponentList func(unsafe.Pointer) unsafe.Pointer
	_AudioUnitExtensionSetComponentList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitInitialize func(unsafe.Pointer) unsafe.Pointer
	_AudioUnitProcess func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitProcessMultiple func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRemovePropertyListenerWithUserData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRemoveRenderNotify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitRender func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitReset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitScheduleParameters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitSetParameter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioUnitUninitialize func(unsafe.Pointer) unsafe.Pointer
	_CAClockAddListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockArm func(unsafe.Pointer) unsafe.Pointer
	_CAClockBarBeatTimeToBeats func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockBeatsToBarBeatTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockDisarm func(unsafe.Pointer) unsafe.Pointer
	_CAClockDispose func(unsafe.Pointer) unsafe.Pointer
	_CAClockGetCurrentTempo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetCurrentTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetPlayRate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockGetStartTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockNew func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockParseMIDI func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockRemoveListener func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSMPTETimeToSeconds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSecondsToSMPTETime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSetCurrentTempo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSetCurrentTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSetPlayRate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAClockStart func(unsafe.Pointer) unsafe.Pointer
	_CAClockStop func(unsafe.Pointer) unsafe.Pointer
	_CAClockTranslateTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAShow func(unsafe.Pointer) unsafe.Pointer
	_CAShowFile func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CopyInstrumentInfoFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CopyNameFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DisposeMusicEventIterator func(unsafe.Pointer) unsafe.Pointer
	_DisposeMusicPlayer func(unsafe.Pointer) unsafe.Pointer
	_DisposeMusicSequence func(unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileCreateNew func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileCreateWithURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileDispose func(unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileGetPropertyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileOpen func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileOpenURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileRead func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileSeek func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileTell func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWrapAudioFileID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWrite func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ExtAudioFileWriteAsync func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GetNameFromSoundBank func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceMIDIEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceMIDIEventList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDevicePrepareInstrument func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceReleaseInstrument func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceStartNote func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceStopNote func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicDeviceSysEx func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorDeleteEvent func(unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorGetEventInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasCurrentEvent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasNextEvent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorHasPreviousEvent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorNextEvent func(unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorPreviousEvent func(unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorSeek func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorSetEventInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicEventIteratorSetEventTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetBeatsForHostTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetHostTimeForBeats func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetPlayRateScalar func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetSequence func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerGetTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerIsPlaying func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerPreroll func(unsafe.Pointer) unsafe.Pointer
	_MusicPlayerSetPlayRateScalar func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerSetSequence func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerSetTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicPlayerStart func(unsafe.Pointer) unsafe.Pointer
	_MusicPlayerStop func(unsafe.Pointer) unsafe.Pointer
	_MusicSequenceBarBeatTimeToBeats func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceBeatsToBarBeatTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceDisposeTrack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileCreateData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileLoad func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceFileLoadData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetAUGraph func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetBeatsForSeconds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetIndTrack func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetInfoDictionary func(unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetSecondsForBeats func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetSequenceType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTempoTrack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTrackCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceGetTrackIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceLoadSMFDataWithFlags func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceLoadSMFWithFlags func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceNewTrack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceReverse func(unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSaveMIDIFile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSaveSMFData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetAUGraph func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetMIDIEndpoint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetSequenceType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicSequenceSetUserCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackClear func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackCopyInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackCut func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetDestMIDIEndpoint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetDestNode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackGetSequence func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackMerge func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackMoveEvents func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewAUPresetEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedControlEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedNoteEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewExtendedTempoEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDIChannelEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDINoteEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMIDIRawDataEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewMetaEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewParameterEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackNewUserEvent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackSetDestMIDIEndpoint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackSetDestNode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MusicTrackSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NewAUGraph func(unsafe.Pointer) unsafe.Pointer
	_NewMusicEventIterator func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NewMusicPlayer func(unsafe.Pointer) unsafe.Pointer
	_NewMusicSequence func(unsafe.Pointer) unsafe.Pointer
	_NewMusicTrackFrom func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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



// Creates a new interval workgroup for managing real-time audio threads. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/3547073-audioworkintervalcreate
func AudioWorkIntervalCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _AudioWorkIntervalCreate(p0)
	}


// AUEventListenerAddEventType is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerAddEventType(_:_:_:)
func AUEventListenerAddEventType(inListener unsafe.Pointer, inObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerAddEventType(inListener, inObject, inEvent)
	}


// AUEventListenerCreate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreate(_:_:_:_:_:_:_:)
func AUEventListenerCreate(inProc unsafe.Pointer, inUserData unsafe.Pointer, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inNotificationInterval unsafe.Pointer, inValueChangeGranularity unsafe.Pointer, outListener unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, inValueChangeGranularity, outListener)
	}


// AUEventListenerCreateWithDispatchQueue is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreateWithDispatchQueue(_:_:_:_:_:)
func AUEventListenerCreateWithDispatchQueue(outListener unsafe.Pointer, inNotificationInterval unsafe.Pointer, inValueChangeGranularity unsafe.Pointer, inDispatchQueue unsafe.Pointer, inBlock unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inValueChangeGranularity, inDispatchQueue, inBlock)
	}


// AUEventListenerNotify is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerNotify(_:_:_:)
func AUEventListenerNotify(inSendingListener unsafe.Pointer, inSendingObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerNotify(inSendingListener, inSendingObject, inEvent)
	}


// AUEventListenerRemoveEventType is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRemoveEventType(_:_:_:)
func AUEventListenerRemoveEventType(inListener unsafe.Pointer, inObject unsafe.Pointer, inEvent unsafe.Pointer) unsafe.Pointer {
	return _AUEventListenerRemoveEventType(inListener, inObject, inEvent)
	}


// Adds a node to an audio processing graph. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphAddNode(_:_:_:)
func AUGraphAddNode(inGraph unsafe.Pointer, inDescription unsafe.Pointer, outNode unsafe.Pointer) unsafe.Pointer {
	return _AUGraphAddNode(inGraph, inDescription, outNode)
	}


// Retrieves information about a particular interaction in an audio processing graph. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetInteractionInfo(_:_:_:)
func AUGraphGetInteractionInfo(inGraph unsafe.Pointer, inInteractionIndex unsafe.Pointer, outInteraction unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetInteractionInfo(inGraph, inInteractionIndex, outInteraction)
	}


// Deprecated in OS X v10.5. Instead, use . [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInfo
func AUGraphGetNodeInfo(inGraph unsafe.Pointer, inNode unsafe.Pointer, outDescription unsafe.Pointer, outClassDataSize unsafe.Pointer, outClassData unsafe.Pointer, outAudioUnit unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetNodeInfo(inGraph, inNode, outDescription, outClassDataSize, outClassData, outAudioUnit)
	}


// Retrieves information about the interactions in an audio processing graph for a given node. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInteractions(_:_:_:_:)
func AUGraphGetNodeInteractions(inGraph unsafe.Pointer, inNode unsafe.Pointer, ioNumInteractions unsafe.Pointer, outInteractions unsafe.Pointer) unsafe.Pointer {
	return _AUGraphGetNodeInteractions(inGraph, inNode, ioNumInteractions, outInteractions)
	}


// Deprecated in OS X v10.5. Instead, use . [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUGraphNewNode
func AUGraphNewNode(inGraph unsafe.Pointer, inDescription unsafe.Pointer, inClassDataSize unsafe.Pointer, inClassData unsafe.Pointer, outNode unsafe.Pointer) unsafe.Pointer {
	return _AUGraphNewNode(inGraph, inDescription, inClassDataSize, inClassData, outNode)
	}


// AUListenerAddParameter is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerAddParameter(_:_:_:)
func AUListenerAddParameter(inListener unsafe.Pointer, inObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUListenerAddParameter(inListener, inObject, inParameter)
	}


// AUListenerCreate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreate(_:_:_:_:_:_:)
func AUListenerCreate(inProc unsafe.Pointer, inUserData unsafe.Pointer, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inNotificationInterval unsafe.Pointer, outListener unsafe.Pointer) unsafe.Pointer {
	return _AUListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, outListener)
	}


// AUListenerCreateWithDispatchQueue is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreateWithDispatchQueue(_:_:_:_:)
func AUListenerCreateWithDispatchQueue(outListener unsafe.Pointer, inNotificationInterval unsafe.Pointer, inDispatchQueue unsafe.Pointer, inBlock unsafe.Pointer) unsafe.Pointer {
	return _AUListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inDispatchQueue, inBlock)
	}


// AUListenerDispose is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerDispose(_:)
func AUListenerDispose(inListener unsafe.Pointer) unsafe.Pointer {
	return _AUListenerDispose(inListener)
	}


// AUListenerRemoveParameter is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUListenerRemoveParameter(_:_:_:)
func AUListenerRemoveParameter(inListener unsafe.Pointer, inObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUListenerRemoveParameter(inListener, inObject, inParameter)
	}


// AUParameterFormatValue is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterFormatValue(_:_:_:_:)
func AUParameterFormatValue(inParameterValue unsafe.Pointer, inParameter unsafe.Pointer, inTextBuffer unsafe.Pointer, inDigits unsafe.Pointer) unsafe.Pointer {
	return _AUParameterFormatValue(inParameterValue, inParameter, inTextBuffer, inDigits)
	}


// AUParameterListenerNotify is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerNotify(_:_:_:)
func AUParameterListenerNotify(inSendingListener unsafe.Pointer, inSendingObject unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUParameterListenerNotify(inSendingListener, inSendingObject, inParameter)
	}


// AUParameterSet is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterSet(_:_:_:_:_:)
func AUParameterSet(inSendingListener unsafe.Pointer, inSendingObject unsafe.Pointer, inParameter unsafe.Pointer, inValue unsafe.Pointer, inBufferOffsetInFrames unsafe.Pointer) unsafe.Pointer {
	return _AUParameterSet(inSendingListener, inSendingObject, inParameter, inValue, inBufferOffsetInFrames)
	}


// AUParameterValueFromLinear is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueFromLinear(_:_:)
func AUParameterValueFromLinear(inLinearValue unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUParameterValueFromLinear(inLinearValue, inParameter)
	}


// AUParameterValueToLinear is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueToLinear(_:_:)
func AUParameterValueToLinear(inParameterValue unsafe.Pointer, inParameter unsafe.Pointer) unsafe.Pointer {
	return _AUParameterValueToLinear(inParameterValue, inParameter)
	}


// AudioCodecAppendInputBufferList is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferList(_:_:_:_:_:)
func AudioCodecAppendInputBufferList(inCodec unsafe.Pointer, inBufferList unsafe.Pointer, ioNumberPackets unsafe.Pointer, inPacketDescription unsafe.Pointer, outBytesConsumed unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecAppendInputBufferList(inCodec, inBufferList, ioNumberPackets, inPacketDescription, outBytesConsumed)
	}


// Appends audio data to the codec’s input buffer. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputData(_:_:_:_:_:)
func AudioCodecAppendInputData(inCodec unsafe.Pointer, inInputData unsafe.Pointer, ioInputDataByteSize unsafe.Pointer, ioNumberPackets unsafe.Pointer, inPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecAppendInputData(inCodec, inInputData, ioInputDataByteSize, ioNumberPackets, inPacketDescription)
	}


// Retrieves the value of a codec property. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetProperty(_:_:_:_:)
func AudioCodecGetProperty(inCodec unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecGetProperty(inCodec, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// Retrieves information about a codec property. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfo(_:_:_:_:)
func AudioCodecGetPropertyInfo(inCodec unsafe.Pointer, inPropertyID unsafe.Pointer, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecGetPropertyInfo(inCodec, inPropertyID, outSize, outWritable)
	}


// Sets up the specified codec to perform a data format translation. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitialize(_:_:_:_:_:)
func AudioCodecInitialize(inCodec unsafe.Pointer, inInputFormat unsafe.Pointer, inOutputFormat unsafe.Pointer, inMagicCookie unsafe.Pointer, inMagicCookieByteSize unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecInitialize(inCodec, inInputFormat, inOutputFormat, inMagicCookie, inMagicCookieByteSize)
	}


// AudioCodecProduceOutputBufferList is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferList(_:_:_:_:_:)
func AudioCodecProduceOutputBufferList(inCodec unsafe.Pointer, ioBufferList unsafe.Pointer, ioNumberPackets unsafe.Pointer, outPacketDescription unsafe.Pointer, outStatus unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecProduceOutputBufferList(inCodec, ioBufferList, ioNumberPackets, outPacketDescription, outStatus)
	}


// Retrieves output data from a codec. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPackets(_:_:_:_:_:_:)
func AudioCodecProduceOutputPackets(inCodec unsafe.Pointer, outOutputData unsafe.Pointer, ioOutputDataByteSize unsafe.Pointer, ioNumberPackets unsafe.Pointer, outPacketDescription unsafe.Pointer, outStatus unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecProduceOutputPackets(inCodec, outOutputData, ioOutputDataByteSize, ioNumberPackets, outPacketDescription, outStatus)
	}


// Flushes all the audio data in the codec and clears the input buffer. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecReset(_:)
func AudioCodecReset(inCodec unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecReset(inCodec)
	}


// Sets the value of a codec property. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetProperty(_:_:_:_:)
func AudioCodecSetProperty(inCodec unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecSetProperty(inCodec, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// Moves the codec from the initialized state back to the uninitialized state. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitialize(_:)
func AudioCodecUninitialize(inCodec unsafe.Pointer) unsafe.Pointer {
	return _AudioCodecUninitialize(inCodec)
	}


// AudioComponentCopyConfigurationInfo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyConfigurationInfo(_:_:)
func AudioComponentCopyConfigurationInfo(inComponent unsafe.Pointer, outConfigurationInfo unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCopyConfigurationInfo(inComponent, outConfigurationInfo)
	}


// AudioComponentCopyIcon is a AudioToolbox function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyIcon(_:)
func AudioComponentCopyIcon(comp unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCopyIcon(comp)
	}


// Returns the generic name of an audio component. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyName(_:_:)
func AudioComponentCopyName(inComponent unsafe.Pointer, outName unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCopyName(inComponent, outName)
	}


// Returns the number of audio components that match a specified structure. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCount(_:)
func AudioComponentCount(inDesc unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentCount(inDesc)
	}


// Finds the next component that matches a specified structure after a specified audio component. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFindNext(_:_:)
func AudioComponentFindNext(inComponent unsafe.Pointer, inDesc unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentFindNext(inComponent, inDesc)
	}


// Gets the class description, as an structure, of an audio component. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetDescription(_:_:)
func AudioComponentGetDescription(inComponent unsafe.Pointer, outDesc unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetDescription(inComponent, outDesc)
	}


// The UIImage of the audio component’s icon. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetIcon(_:_:)
func AudioComponentGetIcon(comp unsafe.Pointer, desiredPointSize unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetIcon(comp, desiredPointSize)
	}


// The time at which the application publishing the component was last active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetLastActiveTime(_:)
func AudioComponentGetLastActiveTime(comp unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetLastActiveTime(comp)
	}


// Gets the version of an audio component in hexadecimal form as (major, minor, dot). [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetVersion(_:_:)
func AudioComponentGetVersion(inComponent unsafe.Pointer, outVersion unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentGetVersion(inComponent, outVersion)
	}


// Determines if an audio component instance implements a particular function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceCanDo(_:_:)
func AudioComponentInstanceCanDo(inInstance unsafe.Pointer, inSelectorID unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceCanDo(inInstance, inSelectorID)
	}


// Disposes of an audio component instance. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceDispose(_:)
func AudioComponentInstanceDispose(inInstance unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceDispose(inInstance)
	}


// Retrieves a reference to an audio component from an instance of that audio component. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceGetComponent(_:)
func AudioComponentInstanceGetComponent(inInstance unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceGetComponent(inInstance)
	}


// Creates a new instance of an audio component. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceNew(_:_:)
func AudioComponentInstanceNew(inComponent unsafe.Pointer, outInstance unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentInstanceNew(inComponent, outInstance)
	}


// AudioComponentInstantiate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiate(_:_:_:)
func AudioComponentInstantiate(inComponent unsafe.Pointer, inOptions unsafe.Pointer) {
	_AudioComponentInstantiate(inComponent, inOptions)
	}


// AudioComponentRegister is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentRegister(_:_:_:_:)
func AudioComponentRegister(inDesc unsafe.Pointer, inName unsafe.Pointer, inVersion unsafe.Pointer, inFactory unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentRegister(inDesc, inName, inVersion, inFactory)
	}


// AudioComponentValidate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidate(_:_:_:)
func AudioComponentValidate(inComponent unsafe.Pointer, inValidationParameters unsafe.Pointer, outValidationResult unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentValidate(inComponent, inValidationParameters, outValidationResult)
	}


// AudioComponentValidateWithResults is a AudioToolbox function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidateWithResults(_:_:_:)
func AudioComponentValidateWithResults(inComponent unsafe.Pointer, inValidationParameters unsafe.Pointer) unsafe.Pointer {
	return _AudioComponentValidateWithResults(inComponent, inValidationParameters)
	}


// Converts audio data from one linear PCM format to another. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertBuffer(_:_:_:_:_:)
func AudioConverterConvertBuffer(inAudioConverter unsafe.Pointer, inInputDataSize unsafe.Pointer, inInputData unsafe.Pointer, ioOutputDataSize unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterConvertBuffer(inAudioConverter, inInputDataSize, inInputData, ioOutputDataSize, outOutputData)
	}


// Converts audio data from one linear PCM format to another, where both use the same sample rate. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertComplexBuffer(_:_:_:_:)
func AudioConverterConvertComplexBuffer(inAudioConverter unsafe.Pointer, inNumberPCMFrames unsafe.Pointer, inInputData unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterConvertComplexBuffer(inAudioConverter, inNumberPCMFrames, inInputData, outOutputData)
	}


// Disposes of an audio converter object. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterDispose(_:)
func AudioConverterDispose(inAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterDispose(inAudioConverter)
	}


// AudioConverterFillBuffer is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillBuffer
func AudioConverterFillBuffer(inAudioConverter unsafe.Pointer, inInputDataProc unsafe.Pointer, inInputDataProcUserData unsafe.Pointer, ioOutputDataSize unsafe.Pointer, outOutputData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataSize, outOutputData)
	}


// Converts audio data supplied by a callback function, supporting non-interleaved and packetized formats. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBuffer(_:_:_:_:_:_:)
func AudioConverterFillComplexBuffer(inAudioConverter unsafe.Pointer, inInputDataProc unsafe.Pointer, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
	}


// AudioConverterFillComplexBufferRealtimeSafe is a AudioToolbox function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferRealtimeSafe(_:_:_:_:_:_:)
func AudioConverterFillComplexBufferRealtimeSafe(inAudioConverter unsafe.Pointer, inInputDataProc unsafe.Pointer, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescription unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBufferRealtimeSafe(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
	}


// AudioConverterFillComplexBufferWithPacketDependencies is a AudioToolbox function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferWithPacketDependencies(_:_:_:_:_:_:_:)
func AudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter unsafe.Pointer, inInputDataProc unsafe.Pointer, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize unsafe.Pointer, outOutputData unsafe.Pointer, outPacketDescriptions unsafe.Pointer, outPacketDependencies unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescriptions, outPacketDependencies)
	}


// Gets an audio converter property value. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetProperty(_:_:_:_:)
func AudioConverterGetProperty(inAudioConverter unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterGetProperty(inAudioConverter, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// Gets information about an audio converter property. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetPropertyInfo(_:_:_:_:)
func AudioConverterGetPropertyInfo(inAudioConverter unsafe.Pointer, inPropertyID unsafe.Pointer, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterGetPropertyInfo(inAudioConverter, inPropertyID, outSize, outWritable)
	}


// Creates a new audio converter object based on specified audio formats. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNew(_:_:_:)
func AudioConverterNew(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNew(inSourceFormat, inDestinationFormat, outAudioConverter)
	}


// Creates a new audio converter object using a specified codec. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewSpecific(_:_:_:_:_:)
func AudioConverterNewSpecific(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, inNumberClassDescriptions unsafe.Pointer, inClassDescriptions unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNewSpecific(inSourceFormat, inDestinationFormat, inNumberClassDescriptions, inClassDescriptions, outAudioConverter)
	}


// AudioConverterNewWithOptions is a AudioToolbox function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewWithOptions(_:_:_:_:)
func AudioConverterNewWithOptions(inSourceFormat unsafe.Pointer, inDestinationFormat unsafe.Pointer, inOptions unsafe.Pointer, outAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterNewWithOptions(inSourceFormat, inDestinationFormat, inOptions, outAudioConverter)
	}


// AudioConverterPrepare is a AudioToolbox function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrepare(_:_:_:)
func AudioConverterPrepare(inFlags unsafe.Pointer, ioReserved unsafe.Pointer) {
	_AudioConverterPrepare(inFlags, ioReserved)
	}


// Resets an audio converter object, clearing and flushing its buffers. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterReset(_:)
func AudioConverterReset(inAudioConverter unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterReset(inAudioConverter)
	}


// Sets the value of an audio converter object property. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterSetProperty(_:_:_:_:)
func AudioConverterSetProperty(inAudioConverter unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioConverterSetProperty(inAudioConverter, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// Closes an audio file. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileClose(_:)
func AudioFileClose(inAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileClose(inAudioFile)
	}


// AudioFileComponentCloseFile is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCloseFile(_:)
func AudioFileComponentCloseFile(inComponent unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCloseFile(inComponent)
	}


// AudioFileComponentCountUserData is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserData(_:_:_:)
func AudioFileComponentCountUserData(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, outNumberItems unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCountUserData(inComponent, inUserDataID, outNumberItems)
	}


// AudioFileComponentCreate is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreate
func AudioFileComponentCreate(inComponent unsafe.Pointer, inParentRef unsafe.Pointer, inFileName unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outNewFileRef unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCreate(inComponent, inParentRef, inFileName, inFormat, inFlags, outNewFileRef)
	}


// AudioFileComponentCreateURL is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURL(_:_:_:_:)
func AudioFileComponentCreateURL(inComponent unsafe.Pointer, inFileRef unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentCreateURL(inComponent, inFileRef, inFormat, inFlags)
	}


// AudioFileComponentDataIsThisFormat is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentDataIsThisFormat
func AudioFileComponentDataIsThisFormat(inComponent unsafe.Pointer, inClientData unsafe.Pointer, inReadFunc unsafe.Pointer, inWriteFunc unsafe.Pointer, inGetSizeFunc unsafe.Pointer, inSetSizeFunc unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentDataIsThisFormat(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, outResult)
	}


// AudioFileComponentExtensionIsThisFormat is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormat(_:_:_:)
func AudioFileComponentExtensionIsThisFormat(inComponent unsafe.Pointer, inExtension unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentExtensionIsThisFormat(inComponent, inExtension, outResult)
	}


// AudioFileComponentFileDataIsThisFormat is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormat(_:_:_:_:)
func AudioFileComponentFileDataIsThisFormat(inComponent unsafe.Pointer, inDataByteSize unsafe.Pointer, inData unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentFileDataIsThisFormat(inComponent, inDataByteSize, inData, outResult)
	}


// AudioFileComponentFileIsThisFormat is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileIsThisFormat
func AudioFileComponentFileIsThisFormat(inComponent unsafe.Pointer, inFileRefNum unsafe.Pointer, outResult unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentFileIsThisFormat(inComponent, inFileRefNum, outResult)
	}


// AudioFileComponentGetGlobalInfo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfo(_:_:_:_:_:_:)
func AudioFileComponentGetGlobalInfo(inComponent unsafe.Pointer, inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetGlobalInfo(inComponent, inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	}


// AudioFileComponentGetGlobalInfoSize is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSize(_:_:_:_:_:)
func AudioFileComponentGetGlobalInfoSize(inComponent unsafe.Pointer, inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertySize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetGlobalInfoSize(inComponent, inPropertyID, inSpecifierSize, inSpecifier, outPropertySize)
	}


// AudioFileComponentGetProperty is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetProperty(_:_:_:_:)
func AudioFileComponentGetProperty(inComponent unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetProperty(inComponent, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// AudioFileComponentGetPropertyInfo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfo(_:_:_:_:)
func AudioFileComponentGetPropertyInfo(inComponent unsafe.Pointer, inPropertyID unsafe.Pointer, outPropertySize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetPropertyInfo(inComponent, inPropertyID, outPropertySize, outWritable)
	}


// AudioFileComponentGetUserData is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserData(_:_:_:_:_:)
func AudioFileComponentGetUserData(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserData(inComponent, inUserDataID, inIndex, ioUserDataSize, outUserData)
	}


// AudioFileComponentGetUserDataAtOffset is a AudioToolbox function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileComponentGetUserDataAtOffset(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inOffset unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataAtOffset(inComponent, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
	}


// AudioFileComponentGetUserDataSize is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize(_:_:_:_:)
func AudioFileComponentGetUserDataSize(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataSize(inComponent, inUserDataID, inIndex, outUserDataSize)
	}


// AudioFileComponentGetUserDataSize64 is a AudioToolbox function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64(_:_:_:_:)
func AudioFileComponentGetUserDataSize64(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentGetUserDataSize64(inComponent, inUserDataID, inIndex, outUserDataSize)
	}


// AudioFileComponentInitialize is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitialize
func AudioFileComponentInitialize(inComponent unsafe.Pointer, inFileRef unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentInitialize(inComponent, inFileRef, inFormat, inFlags)
	}


// AudioFileComponentInitializeWithCallbacks is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileComponentInitializeWithCallbacks(inComponent unsafe.Pointer, inClientData unsafe.Pointer, inReadFunc unsafe.Pointer, inWriteFunc unsafe.Pointer, inGetSizeFunc unsafe.Pointer, inSetSizeFunc unsafe.Pointer, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentInitializeWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags)
	}


// AudioFileComponentOpenFile is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenFile
func AudioFileComponentOpenFile(inComponent unsafe.Pointer, inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inRefNum unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentOpenFile(inComponent, inFileRef, inPermissions, inRefNum)
	}


// AudioFileComponentOpenURL is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURL(_:_:_:_:)
func AudioFileComponentOpenURL(inComponent unsafe.Pointer, inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileDescriptor unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentOpenURL(inComponent, inFileRef, inPermissions, inFileDescriptor)
	}


// AudioFileComponentOpenWithCallbacks is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacks(_:_:_:_:_:_:)
func AudioFileComponentOpenWithCallbacks(inComponent unsafe.Pointer, inClientData unsafe.Pointer, inReadFunc unsafe.Pointer, inWriteFunc unsafe.Pointer, inGetSizeFunc unsafe.Pointer, inSetSizeFunc unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentOpenWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc)
	}


// AudioFileComponentOptimize is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimize(_:)
func AudioFileComponentOptimize(inComponent unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentOptimize(inComponent)
	}


// AudioFileComponentReadBytes is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytes(_:_:_:_:_:)
func AudioFileComponentReadBytes(inComponent unsafe.Pointer, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, outBuffer)
	}


// AudioFileComponentReadPacketData is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketData(_:_:_:_:_:_:_:)
func AudioFileComponentReadPacketData(inComponent unsafe.Pointer, inUseCache unsafe.Pointer, ioNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadPacketData(inComponent, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	}


// AudioFileComponentReadPackets is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPackets(_:_:_:_:_:_:_:)
func AudioFileComponentReadPackets(inComponent unsafe.Pointer, inUseCache unsafe.Pointer, outNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentReadPackets(inComponent, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	}


// AudioFileComponentRemoveUserData is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserData(_:_:_:)
func AudioFileComponentRemoveUserData(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentRemoveUserData(inComponent, inUserDataID, inIndex)
	}


// AudioFileComponentSetProperty is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetProperty(_:_:_:_:)
func AudioFileComponentSetProperty(inComponent unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentSetProperty(inComponent, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// AudioFileComponentSetUserData is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserData(_:_:_:_:_:)
func AudioFileComponentSetUserData(inComponent unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inUserDataSize unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentSetUserData(inComponent, inUserDataID, inIndex, inUserDataSize, inUserData)
	}


// AudioFileComponentWriteBytes is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytes(_:_:_:_:_:)
func AudioFileComponentWriteBytes(inComponent unsafe.Pointer, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentWriteBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, inBuffer)
	}


// AudioFileComponentWritePackets is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePackets(_:_:_:_:_:_:_:)
func AudioFileComponentWritePackets(inComponent unsafe.Pointer, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileComponentWritePackets(inComponent, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
	}


// Gets the number of user data items with a specified ID in a file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCountUserData(_:_:_:)
func AudioFileCountUserData(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, outNumberItems unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCountUserData(inAudioFile, inUserDataID, outNumberItems)
	}


// AudioFileCreate is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCreate
func AudioFileCreate(inParentRef unsafe.Pointer, inFileName unsafe.Pointer, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outNewFileRef unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCreate(inParentRef, inFileName, inFileType, inFormat, inFlags, outNewFileRef, outAudioFile)
	}


// Creates a new audio file, or initializes an existing file, specified by a URL. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileCreateWithURL(_:_:_:_:_:)
func AudioFileCreateWithURL(inFileRef unsafe.Pointer, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileCreateWithURL(inFileRef, inFileType, inFormat, inFlags, outAudioFile)
	}


// Copies the value of a global property into a buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfo(_:_:_:_:_:)
func AudioFileGetGlobalInfo(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetGlobalInfo(inPropertyID, inSpecifierSize, inSpecifier, ioDataSize, outPropertyData)
	}


// Gets the size of a global audio file property. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfoSize(_:_:_:_:)
func AudioFileGetGlobalInfoSize(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetGlobalInfoSize(inPropertyID, inSpecifierSize, inSpecifier, outDataSize)
	}


// Gets the value of an audio file property. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetProperty(_:_:_:_:)
func AudioFileGetProperty(inAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, ioDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetProperty(inAudioFile, inPropertyID, ioDataSize, outPropertyData)
	}


// Gets information about an audio file property, including the size of the property value and whether the value is writable. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetPropertyInfo(_:_:_:_:)
func AudioFileGetPropertyInfo(inAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, outDataSize unsafe.Pointer, isWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetPropertyInfo(inAudioFile, inPropertyID, outDataSize, isWritable)
	}


// Gets a chunk from an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserData(_:_:_:_:_:)
func AudioFileGetUserData(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserData(inAudioFile, inUserDataID, inIndex, ioUserDataSize, outUserData)
	}


// Gets part of the data from a chunk in an audio file. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileGetUserDataAtOffset(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inOffset unsafe.Pointer, ioUserDataSize unsafe.Pointer, outUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataAtOffset(inAudioFile, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
	}


// Gets the size of a user data item in an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize(_:_:_:_:)
func AudioFileGetUserDataSize(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataSize(inAudioFile, inUserDataID, inIndex, outUserDataSize)
	}


// Gets the size of a user data item in an audio file. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize64(_:_:_:_:)
func AudioFileGetUserDataSize64(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, outUserDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFileGetUserDataSize64(inAudioFile, inUserDataID, inIndex, outUserDataSize)
	}


// AudioFileInitialize is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileInitialize
func AudioFileInitialize(inFileRef unsafe.Pointer, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileInitialize(inFileRef, inFileType, inFormat, inFlags, outAudioFile)
	}


// Deletes the content of an existing file and assigns callbacks to the audio file object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileInitializeWithCallbacks(inClientData unsafe.Pointer, inReadFunc unsafe.Pointer, inWriteFunc unsafe.Pointer, inGetSizeFunc unsafe.Pointer, inSetSizeFunc unsafe.Pointer, inFileType unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileInitializeWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags, outAudioFile)
	}


// AudioFileOpen is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpen
func AudioFileOpen(inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileTypeHint unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpen(inFileRef, inPermissions, inFileTypeHint, outAudioFile)
	}


// Open an existing audio file specified by a URL. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenURL(_:_:_:_:)
func AudioFileOpenURL(inFileRef unsafe.Pointer, inPermissions unsafe.Pointer, inFileTypeHint unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpenURL(inFileRef, inPermissions, inFileTypeHint, outAudioFile)
	}


// Opens an existing file with callbacks you provide. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenWithCallbacks(_:_:_:_:_:_:_:)
func AudioFileOpenWithCallbacks(inClientData unsafe.Pointer, inReadFunc unsafe.Pointer, inWriteFunc unsafe.Pointer, inGetSizeFunc unsafe.Pointer, inSetSizeFunc unsafe.Pointer, inFileTypeHint unsafe.Pointer, outAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOpenWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileTypeHint, outAudioFile)
	}


// Consolidates audio data and performs other internal optimizations of the file structure. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileOptimize(_:)
func AudioFileOptimize(inAudioFile unsafe.Pointer) unsafe.Pointer {
	return _AudioFileOptimize(inAudioFile)
	}


// Reads bytes of audio data from an audio file. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadBytes(_:_:_:_:_:)
func AudioFileReadBytes(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, outBuffer)
	}


// Reads packets of audio data from an audio file. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadPacketData(_:_:_:_:_:_:_:)
func AudioFileReadPacketData(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, ioNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadPacketData(inAudioFile, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	}


// Reads a fixed duration of audio data from an audio file. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadPackets(_:_:_:_:_:_:_:)
func AudioFileReadPackets(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, outNumBytes unsafe.Pointer, outPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileReadPackets(inAudioFile, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	}


// Removes a user data item from an audio file. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRemoveUserData(_:_:_:)
func AudioFileRemoveUserData(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer) unsafe.Pointer {
	return _AudioFileRemoveUserData(inAudioFile, inUserDataID, inIndex)
	}


// Sets the value of an audio file property [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetProperty(_:_:_:_:)
func AudioFileSetProperty(inAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, inDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileSetProperty(inAudioFile, inPropertyID, inDataSize, inPropertyData)
	}


// Sets a user data item in an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetUserData(_:_:_:_:_:)
func AudioFileSetUserData(inAudioFile unsafe.Pointer, inUserDataID unsafe.Pointer, inIndex unsafe.Pointer, inUserDataSize unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileSetUserData(inAudioFile, inUserDataID, inIndex, inUserDataSize, inUserData)
	}


// Closes and deallocates the specified audio file stream parser. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamClose(_:)
func AudioFileStreamClose(inAudioFileStream unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamClose(inAudioFileStream)
	}


// Retrieves the value of the specified property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetProperty(_:_:_:_:)
func AudioFileStreamGetProperty(inAudioFileStream unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamGetProperty(inAudioFileStream, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// Retrieves information about a property value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetPropertyInfo(_:_:_:_:)
func AudioFileStreamGetPropertyInfo(inAudioFileStream unsafe.Pointer, inPropertyID unsafe.Pointer, outPropertyDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamGetPropertyInfo(inAudioFileStream, inPropertyID, outPropertyDataSize, outWritable)
	}


// Creates and opens a new audio file stream parser. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamOpen(_:_:_:_:_:)
func AudioFileStreamOpen(inClientData unsafe.Pointer, inPropertyListenerProc unsafe.Pointer, inPacketsProc unsafe.Pointer, inFileTypeHint unsafe.Pointer, outAudioFileStream unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamOpen(inClientData, inPropertyListenerProc, inPacketsProc, inFileTypeHint, outAudioFileStream)
	}


// Passes audio file stream data to the parser. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseBytes(_:_:_:_:)
func AudioFileStreamParseBytes(inAudioFileStream unsafe.Pointer, inDataByteSize unsafe.Pointer, inData unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamParseBytes(inAudioFileStream, inDataByteSize, inData, inFlags)
	}


// Provides a byte offset for a specified packet in the data stream. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeek(_:_:_:_:)
func AudioFileStreamSeek(inAudioFileStream unsafe.Pointer, inPacketOffset unsafe.Pointer, outDataByteOffset unsafe.Pointer, ioFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamSeek(inAudioFileStream, inPacketOffset, outDataByteOffset, ioFlags)
	}


// Sets the value of the specified property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSetProperty(_:_:_:_:)
func AudioFileStreamSetProperty(inAudioFileStream unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFileStreamSetProperty(inAudioFileStream, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// Writes bytes of audio data to an audio file. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWriteBytes(_:_:_:_:_:)
func AudioFileWriteBytes(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, inStartingByte unsafe.Pointer, ioNumBytes unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWriteBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, inBuffer)
	}


// Writes packets of audio data to an audio data file. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePackets(_:_:_:_:_:_:_:)
func AudioFileWritePackets(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWritePackets(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
	}


// AudioFileWritePacketsWithDependencies is a AudioToolbox function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePacketsWithDependencies(_:_:_:_:_:_:_:_:)
func AudioFileWritePacketsWithDependencies(inAudioFile unsafe.Pointer, inUseCache unsafe.Pointer, inNumBytes unsafe.Pointer, inPacketDescriptions unsafe.Pointer, inPacketDependencies unsafe.Pointer, inStartingPacket unsafe.Pointer, ioNumPackets unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioFileWritePacketsWithDependencies(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inPacketDependencies, inStartingPacket, ioNumPackets, inBuffer)
	}


// Gets the value of an audio format property. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetProperty(_:_:_:_:_:)
func AudioFormatGetProperty(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioFormatGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	}


// Gets information about an audio format property. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetPropertyInfo(_:_:_:_:)
func AudioFormatGetPropertyInfo(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertyDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioFormatGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize)
	}


// Registers a HAL audio object property listener callback function to be invoked when a specified property changes. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceAddPropertyListener(_:_:_:_:)
func AudioHardwareServiceAddPropertyListener(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, inListener unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceAddPropertyListener(inObjectID, inAddress, inListener, inClientData)
	}


// Gets the value for a specified property. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceGetPropertyData(_:_:_:_:_:_:)
func AudioHardwareServiceGetPropertyData(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, ioDataSize unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData)
	}


// Gets the payload size for a given property. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceGetPropertyDataSize(_:_:_:_:_:)
func AudioHardwareServiceGetPropertyDataSize(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize)
	}


// Queries a HAL audio object about whether or not it has a specified property. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceHasProperty(_:_:)
func AudioHardwareServiceHasProperty(inObjectID unsafe.Pointer, inAddress unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceHasProperty(inObjectID, inAddress)
	}


// Queries a HAL audio object about whether a specified property is settable. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceIsPropertySettable(_:_:_:)
func AudioHardwareServiceIsPropertySettable(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, outIsSettable unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceIsPropertySettable(inObjectID, inAddress, outIsSettable)
	}


// Unregisters a HAL audio object property listener callback function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceRemovePropertyListener(_:_:_:_:)
func AudioHardwareServiceRemovePropertyListener(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, inListener unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceRemovePropertyListener(inObjectID, inAddress, inListener, inClientData)
	}


// Asks a HAL audio object to change the value of a specified property. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioHardwareServiceSetPropertyData(_:_:_:_:_:_:)
func AudioHardwareServiceSetPropertyData(inObjectID unsafe.Pointer, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, inDataSize unsafe.Pointer, inData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareServiceSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData)
	}


// The host app’s icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitGetHostIcon(_:_:)
func AudioOutputUnitGetHostIcon(au unsafe.Pointer, desiredPointSize unsafe.Pointer) unsafe.Pointer {
	return _AudioOutputUnitGetHostIcon(au, desiredPointSize)
	}


// Registers an audio output unit for use by other applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitPublish(_:_:_:_:)
func AudioOutputUnitPublish(inDesc unsafe.Pointer, inName unsafe.Pointer, inVersion unsafe.Pointer, inOutputUnit unsafe.Pointer) unsafe.Pointer {
	return _AudioOutputUnitPublish(inDesc, inName, inVersion, inOutputUnit)
	}


// Starts an I/O audio unit, which in turn starts the audio unit processing graph that it is connected to. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStart(_:)
func AudioOutputUnitStart(ci unsafe.Pointer) unsafe.Pointer {
	return _AudioOutputUnitStart(ci)
	}


// Stops an I/O audio unit, which in turn stops the audio unit processing graph that it is connected to. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStop(_:)
func AudioOutputUnitStop(ci unsafe.Pointer) unsafe.Pointer {
	return _AudioOutputUnitStop(ci)
	}


// Adds a property listener callback to an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAddPropertyListener(_:_:_:_:)
func AudioQueueAddPropertyListener(inAQ unsafe.Pointer, inID unsafe.Pointer, inProc unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAddPropertyListener(inAQ, inID, inProc, inUserData)
	}


// Asks an audio queue object to allocate an audio queue buffer. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBuffer(_:_:_:)
func AudioQueueAllocateBuffer(inAQ unsafe.Pointer, inBufferByteSize unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAllocateBuffer(inAQ, inBufferByteSize, outBuffer)
	}


// Asks an audio queue object to allocate an audio queue buffer with space for packet descriptions. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBufferWithPacketDescriptions(_:_:_:_:)
func AudioQueueAllocateBufferWithPacketDescriptions(inAQ unsafe.Pointer, inBufferByteSize unsafe.Pointer, inNumberPacketDescriptions unsafe.Pointer, outBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueAllocateBufferWithPacketDescriptions(inAQ, inBufferByteSize, inNumberPacketDescriptions, outBuffer)
	}


// Creates a timeline object for an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueCreateTimeline(_:_:)
func AudioQueueCreateTimeline(inAQ unsafe.Pointer, outTimeline unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueCreateTimeline(inAQ, outTimeline)
	}


// Gets the current time of the audio hardware device associated with an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetCurrentTime(_:_:)
func AudioQueueDeviceGetCurrentTime(inAQ unsafe.Pointer, outTimeStamp unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceGetCurrentTime(inAQ, outTimeStamp)
	}


// Gets the start time, for an audio hardware device, that is closest to a requested start time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetNearestStartTime(_:_:_:)
func AudioQueueDeviceGetNearestStartTime(inAQ unsafe.Pointer, ioRequestedStartTime unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceGetNearestStartTime(inAQ, ioRequestedStartTime, inFlags)
	}


// Converts the time for an audio queue’s associated audio hardware device from one time base representation to another. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceTranslateTime(_:_:_:)
func AudioQueueDeviceTranslateTime(inAQ unsafe.Pointer, inTime unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDeviceTranslateTime(inAQ, inTime, outTime)
	}


// Disposes of an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDispose(_:_:)
func AudioQueueDispose(inAQ unsafe.Pointer, inImmediate unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDispose(inAQ, inImmediate)
	}


// Disposes of an audio queue’s timeline object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDisposeTimeline(_:_:)
func AudioQueueDisposeTimeline(inAQ unsafe.Pointer, inTimeline unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueDisposeTimeline(inAQ, inTimeline)
	}


// Adds a buffer to the buffer queue of a recording or playback audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBuffer(_:_:_:_:)
func AudioQueueEnqueueBuffer(inAQ unsafe.Pointer, inBuffer unsafe.Pointer, inNumPacketDescs unsafe.Pointer, inPacketDescs unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueEnqueueBuffer(inAQ, inBuffer, inNumPacketDescs, inPacketDescs)
	}


// Adds a buffer to the buffer queue of a playback audio queue object, specifying start time and other settings. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBufferWithParameters(_:_:_:_:_:_:_:_:_:_:)
func AudioQueueEnqueueBufferWithParameters(inAQ unsafe.Pointer, inBuffer unsafe.Pointer, inNumPacketDescs unsafe.Pointer, inPacketDescs unsafe.Pointer, inTrimFramesAtStart unsafe.Pointer, inTrimFramesAtEnd unsafe.Pointer, inNumParamValues unsafe.Pointer, inParamValues unsafe.Pointer, inStartTime unsafe.Pointer, outActualStartTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueEnqueueBufferWithParameters(inAQ, inBuffer, inNumPacketDescs, inPacketDescs, inTrimFramesAtStart, inTrimFramesAtEnd, inNumParamValues, inParamValues, inStartTime, outActualStartTime)
	}


// Resets an audio queue’s decoder state. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFlush(_:)
func AudioQueueFlush(inAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueFlush(inAQ)
	}


// Asks an audio queue to dispose of an audio queue buffer. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFreeBuffer(_:_:)
func AudioQueueFreeBuffer(inAQ unsafe.Pointer, inBuffer unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueFreeBuffer(inAQ, inBuffer)
	}


// Gets the current audio queue time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetCurrentTime(_:_:_:_:)
func AudioQueueGetCurrentTime(inAQ unsafe.Pointer, inTimeline unsafe.Pointer, outTimeStamp unsafe.Pointer, outTimelineDiscontinuity unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetCurrentTime(inAQ, inTimeline, outTimeStamp, outTimelineDiscontinuity)
	}


// Gets an audio queue parameter value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetParameter(_:_:_:)
func AudioQueueGetParameter(inAQ unsafe.Pointer, inParamID unsafe.Pointer, outValue unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetParameter(inAQ, inParamID, outValue)
	}


// Gets an audio queue property value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetProperty(_:_:_:_:)
func AudioQueueGetProperty(inAQ unsafe.Pointer, inID unsafe.Pointer, outData unsafe.Pointer, ioDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetProperty(inAQ, inID, outData, ioDataSize)
	}


// Gets the size of the value of an audio queue property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetPropertySize(_:_:_:)
func AudioQueueGetPropertySize(inAQ unsafe.Pointer, inID unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueGetPropertySize(inAQ, inID, outDataSize)
	}


// Creates a new recording audio queue object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInput(_:_:_:_:_:_:_:)
func AudioQueueNewInput(inFormat unsafe.Pointer, inCallbackProc unsafe.Pointer, inUserData unsafe.Pointer, inCallbackRunLoop unsafe.Pointer, inCallbackRunLoopMode unsafe.Pointer, inFlags unsafe.Pointer, outAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewInput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
	}


// AudioQueueNewInputWithDispatchQueue is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewInputWithDispatchQueue(outAQ unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, inCallbackDispatchQueue unsafe.Pointer, inCallbackBlock unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewInputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
	}


// Creates a new playback audio queue object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutput(_:_:_:_:_:_:_:)
func AudioQueueNewOutput(inFormat unsafe.Pointer, inCallbackProc unsafe.Pointer, inUserData unsafe.Pointer, inCallbackRunLoop unsafe.Pointer, inCallbackRunLoopMode unsafe.Pointer, inFlags unsafe.Pointer, outAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewOutput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
	}


// AudioQueueNewOutputWithDispatchQueue is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewOutputWithDispatchQueue(outAQ unsafe.Pointer, inFormat unsafe.Pointer, inFlags unsafe.Pointer, inCallbackDispatchQueue unsafe.Pointer, inCallbackBlock unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueNewOutputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
	}


// Exports audio to a buffer, instead of to a device, using a playback audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOfflineRender(_:_:_:_:)
func AudioQueueOfflineRender(inAQ unsafe.Pointer, inTimestamp unsafe.Pointer, ioBuffer unsafe.Pointer, inNumberFrames unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueOfflineRender(inAQ, inTimestamp, ioBuffer, inNumberFrames)
	}


// Pauses audio playback or recording. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePause(_:)
func AudioQueuePause(inAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueuePause(inAQ)
	}


// Decodes enqueued buffers in preparation for playback. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePrime(_:_:_:)
func AudioQueuePrime(inAQ unsafe.Pointer, inNumberOfFramesToPrepare unsafe.Pointer, outNumberOfFramesPrepared unsafe.Pointer) unsafe.Pointer {
	return _AudioQueuePrime(inAQ, inNumberOfFramesToPrepare, outNumberOfFramesPrepared)
	}


// AudioQueueProcessingTapDispose is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapDispose(_:)
func AudioQueueProcessingTapDispose(inAQTap unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapDispose(inAQTap)
	}


// AudioQueueProcessingTapGetQueueTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetQueueTime(_:_:_:)
func AudioQueueProcessingTapGetQueueTime(inAQTap unsafe.Pointer, outQueueSampleTime unsafe.Pointer, outQueueFrameCount unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapGetQueueTime(inAQTap, outQueueSampleTime, outQueueFrameCount)
	}


// AudioQueueProcessingTapGetSourceAudio is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetSourceAudio(_:_:_:_:_:_:)
func AudioQueueProcessingTapGetSourceAudio(inAQTap unsafe.Pointer, inNumberFrames unsafe.Pointer, ioTimeStamp unsafe.Pointer, outFlags unsafe.Pointer, outNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapGetSourceAudio(inAQTap, inNumberFrames, ioTimeStamp, outFlags, outNumberFrames, ioData)
	}


// AudioQueueProcessingTapNew is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapNew(_:_:_:_:_:_:_:)
func AudioQueueProcessingTapNew(inAQ unsafe.Pointer, inCallback unsafe.Pointer, inClientData unsafe.Pointer, inFlags unsafe.Pointer, outMaxFrames unsafe.Pointer, outProcessingFormat unsafe.Pointer, outAQTap unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueProcessingTapNew(inAQ, inCallback, inClientData, inFlags, outMaxFrames, outProcessingFormat, outAQTap)
	}


// Removes a property listener callback from an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueRemovePropertyListener(_:_:_:_:)
func AudioQueueRemovePropertyListener(inAQ unsafe.Pointer, inID unsafe.Pointer, inProc unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueRemovePropertyListener(inAQ, inID, inProc, inUserData)
	}


// Resets an audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueReset(_:)
func AudioQueueReset(inAQ unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueReset(inAQ)
	}


// Sets the rendering mode and audio format for a playback audio queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetOfflineRenderFormat(_:_:_:)
func AudioQueueSetOfflineRenderFormat(inAQ unsafe.Pointer, inFormat unsafe.Pointer, inLayout unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueSetOfflineRenderFormat(inAQ, inFormat, inLayout)
	}


// Sets a playback audio queue parameter value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetParameter(_:_:_:)
func AudioQueueSetParameter(inAQ unsafe.Pointer, inParamID unsafe.Pointer, inValue unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueSetParameter(inAQ, inParamID, inValue)
	}


// Sets an audio queue property value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetProperty(_:_:_:_:)
func AudioQueueSetProperty(inAQ unsafe.Pointer, inID unsafe.Pointer, inData unsafe.Pointer, inDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueSetProperty(inAQ, inID, inData, inDataSize)
	}


// Begins playing or recording audio. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStart(_:_:)
func AudioQueueStart(inAQ unsafe.Pointer, inStartTime unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueStart(inAQ, inStartTime)
	}


// Stops playing or recording audio. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStop(_:_:)
func AudioQueueStop(inAQ unsafe.Pointer, inImmediate unsafe.Pointer) unsafe.Pointer {
	return _AudioQueueStop(inAQ, inImmediate)
	}


// Registers a callback function that is invoked when a specified system sound finishes playing. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesAddSystemSoundCompletion(_:_:_:_:_:)
func AudioServicesAddSystemSoundCompletion(inSystemSoundID unsafe.Pointer, inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inCompletionRoutine unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesAddSystemSoundCompletion(inSystemSoundID, inRunLoop, inRunLoopMode, inCompletionRoutine, inClientData)
	}


// Creates a system sound object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesCreateSystemSoundID(_:_:)
func AudioServicesCreateSystemSoundID(inFileURL unsafe.Pointer, outSystemSoundID unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesCreateSystemSoundID(inFileURL, outSystemSoundID)
	}


// Disposes of a system sound object and associated resources. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesDisposeSystemSoundID(_:)
func AudioServicesDisposeSystemSoundID(inSystemSoundID unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesDisposeSystemSoundID(inSystemSoundID)
	}


// Gets a specified System Sound Services property value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetProperty(_:_:_:_:_:)
func AudioServicesGetProperty(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	}


// Gets information about a System Sound Services property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetPropertyInfo(_:_:_:_:_:)
func AudioServicesGetPropertyInfo(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, outPropertyDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize, outWritable)
	}


// Plays a system sound as an alert. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSound(_:)
func AudioServicesPlayAlertSound(inSystemSoundID unsafe.Pointer) {
	_AudioServicesPlayAlertSound(inSystemSoundID)
	}


// AudioServicesPlayAlertSoundWithCompletion is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSoundWithCompletion(_:_:)
func AudioServicesPlayAlertSoundWithCompletion(inSystemSoundID unsafe.Pointer) {
	_AudioServicesPlayAlertSoundWithCompletion(inSystemSoundID)
	}


// AudioServicesPlayAlertSoundWithDetails is a AudioToolbox function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSoundWithDetails
func AudioServicesPlayAlertSoundWithDetails(inSystemSoundID unsafe.Pointer, inDetails unsafe.Pointer) {
	_AudioServicesPlayAlertSoundWithDetails(inSystemSoundID, inDetails)
	}


// Plays a system sound object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSound(_:)
func AudioServicesPlaySystemSound(inSystemSoundID unsafe.Pointer) {
	_AudioServicesPlaySystemSound(inSystemSoundID)
	}


// AudioServicesPlaySystemSoundWithCompletion is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSoundWithCompletion(_:_:)
func AudioServicesPlaySystemSoundWithCompletion(inSystemSoundID unsafe.Pointer) {
	_AudioServicesPlaySystemSoundWithCompletion(inSystemSoundID)
	}


// AudioServicesPlaySystemSoundWithDetails is a AudioToolbox function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSoundWithDetails
func AudioServicesPlaySystemSoundWithDetails(inSystemSoundID unsafe.Pointer, inDetails unsafe.Pointer) {
	_AudioServicesPlaySystemSoundWithDetails(inSystemSoundID, inDetails)
	}


// Unregisters any completion callback functions that were registered for a specified system sound. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesRemoveSystemSoundCompletion(_:)
func AudioServicesRemoveSystemSoundCompletion(inSystemSoundID unsafe.Pointer) {
	_AudioServicesRemoveSystemSoundCompletion(inSystemSoundID)
	}


// Sets the value for a specified System Sound Services property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioServicesSetProperty(_:_:_:_:_:)
func AudioServicesSetProperty(inPropertyID unsafe.Pointer, inSpecifierSize unsafe.Pointer, inSpecifier unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioServicesSetProperty(inPropertyID, inSpecifierSize, inSpecifier, inPropertyDataSize, inPropertyData)
	}


// Adds a property listener callback function to your application’s audio session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionAddPropertyListener(_:_:_:)
func AudioSessionAddPropertyListener(inID unsafe.Pointer, inProc unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionAddPropertyListener(inID, inProc, inClientData)
	}


// Gets the value of a specified audio session property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionGetProperty(_:_:_:)
func AudioSessionGetProperty(inID unsafe.Pointer, ioDataSize unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionGetProperty(inID, ioDataSize, outData)
	}


// Gets the size of the value for a specified audio session property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionGetPropertySize(_:_:)
func AudioSessionGetPropertySize(inID unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionGetPropertySize(inID, outDataSize)
	}


// Initializes an iOS application’s audio session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionInitialize(_:_:_:_:)
func AudioSessionInitialize(inRunLoop unsafe.Pointer, inRunLoopMode unsafe.Pointer, inInterruptionListener unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionInitialize(inRunLoop, inRunLoopMode, inInterruptionListener, inClientData)
	}


// Removes an audio session property listener callback function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionRemovePropertyListener(_:)
func AudioSessionRemovePropertyListener(inID unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionRemovePropertyListener(inID)
	}


// Removes a property listener callback function from your application’s audio session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionRemovePropertyListenerWithUserData(_:_:_:)
func AudioSessionRemovePropertyListenerWithUserData(inID unsafe.Pointer, inProc unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionRemovePropertyListenerWithUserData(inID, inProc, inClientData)
	}


// Actives or deactivates your application’s audio session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetActive(_:)
func AudioSessionSetActive(active unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetActive(active)
	}


// Activates or deactivates your application’s audio session; provides flags for use by other audio sessions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetActiveWithFlags(_:_:)
func AudioSessionSetActiveWithFlags(active unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetActiveWithFlags(active, inFlags)
	}


// Sets the value of a specified audio session property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioSessionSetProperty(_:_:_:)
func AudioSessionSetProperty(inID unsafe.Pointer, inDataSize unsafe.Pointer, inData unsafe.Pointer) unsafe.Pointer {
	return _AudioSessionSetProperty(inID, inDataSize, inData)
	}


// Registers a callback to receive audio unit property change notifications. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListener(_:_:_:_:)
func AudioUnitAddPropertyListener(inUnit unsafe.Pointer, inID unsafe.Pointer, inProc unsafe.Pointer, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitAddPropertyListener(inUnit, inID, inProc, inProcUserData)
	}


// Registers a callback to receive audio unit render notifications. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotify(_:_:_:)
func AudioUnitAddRenderNotify(inUnit unsafe.Pointer, inProc unsafe.Pointer, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitAddRenderNotify(inUnit, inProc, inProcUserData)
	}


// Returns the component registrations for a given audio unit extension. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionCopyComponentList(_:)
func AudioUnitExtensionCopyComponentList(extensionIdentifier unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitExtensionCopyComponentList(extensionIdentifier)
	}


// Allows the implementor of an audio unit extension to dynamically modify the list of component registrations for the extension. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionSetComponentList(_:_:)
func AudioUnitExtensionSetComponentList(extensionIdentifier unsafe.Pointer, audioComponentInfo unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitExtensionSetComponentList(extensionIdentifier, audioComponentInfo)
	}


// Gets the value of an audio unit parameter. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameter(_:_:_:_:_:)
func AudioUnitGetParameter(inUnit unsafe.Pointer, inID unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer, outValue unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetParameter(inUnit, inID, inScope, inElement, outValue)
	}


// Gets the value of an audio unit property. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetProperty(_:_:_:_:_:_:)
func AudioUnitGetProperty(inUnit unsafe.Pointer, inID unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer, outData unsafe.Pointer, ioDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetProperty(inUnit, inID, inScope, inElement, outData, ioDataSize)
	}


// Gets information about an audio unit property. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfo(_:_:_:_:_:_:)
func AudioUnitGetPropertyInfo(inUnit unsafe.Pointer, inID unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer, outDataSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitGetPropertyInfo(inUnit, inID, inScope, inElement, outDataSize, outWritable)
	}


// Initializes an audio unit [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitialize(_:)
func AudioUnitInitialize(inUnit unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitInitialize(inUnit)
	}


// AudioUnitProcess is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcess(_:_:_:_:_:)
func AudioUnitProcess(inUnit unsafe.Pointer, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitProcess(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, ioData)
	}


// AudioUnitProcessMultiple is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessMultiple(_:_:_:_:_:_:_:_:)
func AudioUnitProcessMultiple(inUnit unsafe.Pointer, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inNumberFrames unsafe.Pointer, inNumberInputBufferLists unsafe.Pointer, inInputBufferLists unsafe.Pointer, inNumberOutputBufferLists unsafe.Pointer, ioOutputBufferLists unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitProcessMultiple(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, inNumberInputBufferLists, inInputBufferLists, inNumberOutputBufferLists, ioOutputBufferLists)
	}


// Unregisters a previously-registered property listener callback function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserData(_:_:_:_:)
func AudioUnitRemovePropertyListenerWithUserData(inUnit unsafe.Pointer, inID unsafe.Pointer, inProc unsafe.Pointer, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRemovePropertyListenerWithUserData(inUnit, inID, inProc, inProcUserData)
	}


// Unregisters a previously-registered render listener callback function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotify(_:_:_:)
func AudioUnitRemoveRenderNotify(inUnit unsafe.Pointer, inProc unsafe.Pointer, inProcUserData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRemoveRenderNotify(inUnit, inProc, inProcUserData)
	}


// Initiates a rendering cycle for an audio unit. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRender(_:_:_:_:_:_:)
func AudioUnitRender(inUnit unsafe.Pointer, ioActionFlags unsafe.Pointer, inTimeStamp unsafe.Pointer, inOutputBusNumber unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitRender(inUnit, ioActionFlags, inTimeStamp, inOutputBusNumber, inNumberFrames, ioData)
	}


// Resets an audio unit’s render state. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitReset(_:_:_:)
func AudioUnitReset(inUnit unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitReset(inUnit, inScope, inElement)
	}


// Schedules changes to the value of an audio unit parameter. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParameters(_:_:_:)
func AudioUnitScheduleParameters(inUnit unsafe.Pointer, inParameterEvent unsafe.Pointer, inNumParamEvents unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitScheduleParameters(inUnit, inParameterEvent, inNumParamEvents)
	}


// Sets the value of an audio unit parameter. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameter(_:_:_:_:_:_:)
func AudioUnitSetParameter(inUnit unsafe.Pointer, inID unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer, inValue unsafe.Pointer, inBufferOffsetInFrames unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitSetParameter(inUnit, inID, inScope, inElement, inValue, inBufferOffsetInFrames)
	}


// Sets the value of an audio unit property. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetProperty(_:_:_:_:_:_:)
func AudioUnitSetProperty(inUnit unsafe.Pointer, inID unsafe.Pointer, inScope unsafe.Pointer, inElement unsafe.Pointer, inData unsafe.Pointer, inDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitSetProperty(inUnit, inID, inScope, inElement, inData, inDataSize)
	}


// Uninitializes an audio unit. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitialize(_:)
func AudioUnitUninitialize(inUnit unsafe.Pointer) unsafe.Pointer {
	return _AudioUnitUninitialize(inUnit)
	}


// CAClockAddListener is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockAddListener(_:_:_:)
func CAClockAddListener(inCAClock unsafe.Pointer, inListenerProc unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _CAClockAddListener(inCAClock, inListenerProc, inUserData)
	}


// CAClockArm is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockArm(_:)
func CAClockArm(inCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockArm(inCAClock)
	}


// CAClockBarBeatTimeToBeats is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBarBeatTimeToBeats(_:_:_:)
func CAClockBarBeatTimeToBeats(inCAClock unsafe.Pointer, inBarBeatTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _CAClockBarBeatTimeToBeats(inCAClock, inBarBeatTime, outBeats)
	}


// CAClockBeatsToBarBeatTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockBeatsToBarBeatTime(_:_:_:_:)
func CAClockBeatsToBarBeatTime(inCAClock unsafe.Pointer, inBeats unsafe.Pointer, inSubbeatDivisor unsafe.Pointer, outBarBeatTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockBeatsToBarBeatTime(inCAClock, inBeats, inSubbeatDivisor, outBarBeatTime)
	}


// CAClockDisarm is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockDisarm(_:)
func CAClockDisarm(inCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockDisarm(inCAClock)
	}


// CAClockDispose is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockDispose(_:)
func CAClockDispose(inCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockDispose(inCAClock)
	}


// CAClockGetCurrentTempo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTempo(_:_:_:)
func CAClockGetCurrentTempo(inCAClock unsafe.Pointer, outTempo unsafe.Pointer, outTimestamp unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetCurrentTempo(inCAClock, outTempo, outTimestamp)
	}


// CAClockGetCurrentTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTime(_:_:_:)
func CAClockGetCurrentTime(inCAClock unsafe.Pointer, inTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetCurrentTime(inCAClock, inTimeFormat, outTime)
	}


// CAClockGetPlayRate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPlayRate(_:_:)
func CAClockGetPlayRate(inCAClock unsafe.Pointer, outPlayRate unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetPlayRate(inCAClock, outPlayRate)
	}


// CAClockGetProperty is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetProperty(_:_:_:_:)
func CAClockGetProperty(inCAClock unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetProperty(inCAClock, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// CAClockGetPropertyInfo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPropertyInfo(_:_:_:_:)
func CAClockGetPropertyInfo(inCAClock unsafe.Pointer, inPropertyID unsafe.Pointer, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetPropertyInfo(inCAClock, inPropertyID, outSize, outWritable)
	}


// CAClockGetStartTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockGetStartTime(_:_:_:)
func CAClockGetStartTime(inCAClock unsafe.Pointer, inTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockGetStartTime(inCAClock, inTimeFormat, outTime)
	}


// CAClockNew is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockNew(_:_:)
func CAClockNew(inReservedFlags unsafe.Pointer, outCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockNew(inReservedFlags, outCAClock)
	}


// CAClockParseMIDI is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockParseMIDI(_:_:)
func CAClockParseMIDI(inCAClock unsafe.Pointer, inMIDIPacketList unsafe.Pointer) unsafe.Pointer {
	return _CAClockParseMIDI(inCAClock, inMIDIPacketList)
	}


// CAClockRemoveListener is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockRemoveListener(_:_:_:)
func CAClockRemoveListener(inCAClock unsafe.Pointer, inListenerProc unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _CAClockRemoveListener(inCAClock, inListenerProc, inUserData)
	}


// CAClockSMPTETimeToSeconds is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTETimeToSeconds(_:_:_:)
func CAClockSMPTETimeToSeconds(inCAClock unsafe.Pointer, inSMPTETime unsafe.Pointer, outSeconds unsafe.Pointer) unsafe.Pointer {
	return _CAClockSMPTETimeToSeconds(inCAClock, inSMPTETime, outSeconds)
	}


// CAClockSecondsToSMPTETime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSecondsToSMPTETime(_:_:_:_:)
func CAClockSecondsToSMPTETime(inCAClock unsafe.Pointer, inSeconds unsafe.Pointer, inSubframeDivisor unsafe.Pointer, outSMPTETime unsafe.Pointer) unsafe.Pointer {
	return _CAClockSecondsToSMPTETime(inCAClock, inSeconds, inSubframeDivisor, outSMPTETime)
	}


// CAClockSetCurrentTempo is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTempo(_:_:_:)
func CAClockSetCurrentTempo(inCAClock unsafe.Pointer, inTempo unsafe.Pointer, inTimestamp unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetCurrentTempo(inCAClock, inTempo, inTimestamp)
	}


// CAClockSetCurrentTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTime(_:_:)
func CAClockSetCurrentTime(inCAClock unsafe.Pointer, inTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetCurrentTime(inCAClock, inTime)
	}


// CAClockSetPlayRate is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetPlayRate(_:_:)
func CAClockSetPlayRate(inCAClock unsafe.Pointer, inPlayRate unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetPlayRate(inCAClock, inPlayRate)
	}


// CAClockSetProperty is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockSetProperty(_:_:_:_:)
func CAClockSetProperty(inCAClock unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _CAClockSetProperty(inCAClock, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// CAClockStart is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockStart(_:)
func CAClockStart(inCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockStart(inCAClock)
	}


// CAClockStop is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockStop(_:)
func CAClockStop(inCAClock unsafe.Pointer) unsafe.Pointer {
	return _CAClockStop(inCAClock)
	}


// CAClockTranslateTime is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTranslateTime(_:_:_:_:)
func CAClockTranslateTime(inCAClock unsafe.Pointer, inTime unsafe.Pointer, inOutputTimeFormat unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CAClockTranslateTime(inCAClock, inTime, inOutputTimeFormat, outTime)
	}


// Prints the internal state of an object to . [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAShow(_:)
func CAShow(inObject unsafe.Pointer) {
	_CAShow(inObject)
	}


// Prints the internal state of an object to a file. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAShowFile(_:_:)
func CAShowFile(inObject unsafe.Pointer, inFile unsafe.Pointer) {
	_CAShowFile(inObject, inFile)
	}


// CopyInstrumentInfoFromSoundBank is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CopyInstrumentInfoFromSoundBank(_:_:)
func CopyInstrumentInfoFromSoundBank(inURL unsafe.Pointer, outInstrumentInfo unsafe.Pointer) unsafe.Pointer {
	return _CopyInstrumentInfoFromSoundBank(inURL, outInstrumentInfo)
	}


// Copies the name of a sound bank from a sound bank file at a specified URL. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CopyNameFromSoundBank(_:_:)
func CopyNameFromSoundBank(inURL unsafe.Pointer, outName unsafe.Pointer) unsafe.Pointer {
	return _CopyNameFromSoundBank(inURL, outName)
	}


// Disposes of a music event iterator. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicEventIterator(_:)
func DisposeMusicEventIterator(inIterator unsafe.Pointer) unsafe.Pointer {
	return _DisposeMusicEventIterator(inIterator)
	}


// Disposes of a music player. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicPlayer(_:)
func DisposeMusicPlayer(inPlayer unsafe.Pointer) unsafe.Pointer {
	return _DisposeMusicPlayer(inPlayer)
	}


// Disposes of a music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicSequence(_:)
func DisposeMusicSequence(inSequence unsafe.Pointer) unsafe.Pointer {
	return _DisposeMusicSequence(inSequence)
	}


// Deprecated. Use the function instead. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateNew
func ExtAudioFileCreateNew(inParentDir unsafe.Pointer, inFileName unsafe.Pointer, inFileType unsafe.Pointer, inStreamDesc unsafe.Pointer, inChannelLayout unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileCreateNew(inParentDir, inFileName, inFileType, inStreamDesc, inChannelLayout, outExtAudioFile)
	}


// Creates a new audio file and associates it with a new extended audio file object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateWithURL(_:_:_:_:_:_:)
func ExtAudioFileCreateWithURL(inURL unsafe.Pointer, inFileType unsafe.Pointer, inStreamDesc unsafe.Pointer, inChannelLayout unsafe.Pointer, inFlags unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileCreateWithURL(inURL, inFileType, inStreamDesc, inChannelLayout, inFlags, outExtAudioFile)
	}


// Disposes of an extended audio file object and closes the associated file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileDispose(_:)
func ExtAudioFileDispose(inExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileDispose(inExtAudioFile)
	}


// Gets a property value from an extended audio file object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetProperty(_:_:_:_:)
func ExtAudioFileGetProperty(inExtAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileGetProperty(inExtAudioFile, inPropertyID, ioPropertyDataSize, outPropertyData)
	}


// Gets information about an extended audio file object property. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetPropertyInfo(_:_:_:_:)
func ExtAudioFileGetPropertyInfo(inExtAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileGetPropertyInfo(inExtAudioFile, inPropertyID, outSize, outWritable)
	}


// Deprecated. Use the function instead. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpen
func ExtAudioFileOpen(inFSRef unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileOpen(inFSRef, outExtAudioFile)
	}


// Opens an existing audio file for reading, and associates it with a new extended audio file object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpenURL(_:_:)
func ExtAudioFileOpenURL(inURL unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileOpenURL(inURL, outExtAudioFile)
	}


// Performs a synchronous, sequential read operation on an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileRead(_:_:_:)
func ExtAudioFileRead(inExtAudioFile unsafe.Pointer, ioNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileRead(inExtAudioFile, ioNumberFrames, ioData)
	}


// Seeks to a specified frame in a file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSeek(_:_:)
func ExtAudioFileSeek(inExtAudioFile unsafe.Pointer, inFrameOffset unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileSeek(inExtAudioFile, inFrameOffset)
	}


// Sets a property value for an extended audio file object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSetProperty(_:_:_:_:)
func ExtAudioFileSetProperty(inExtAudioFile unsafe.Pointer, inPropertyID unsafe.Pointer, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileSetProperty(inExtAudioFile, inPropertyID, inPropertyDataSize, inPropertyData)
	}


// Gets an audio file’s read/write position. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileTell(_:_:)
func ExtAudioFileTell(inExtAudioFile unsafe.Pointer, outFrameOffset unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileTell(inExtAudioFile, outFrameOffset)
	}


// Wraps an audio file object in an extended audio file object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrapAudioFileID(_:_:_:)
func ExtAudioFileWrapAudioFileID(inFileID unsafe.Pointer, inForWriting unsafe.Pointer, outExtAudioFile unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWrapAudioFileID(inFileID, inForWriting, outExtAudioFile)
	}


// Performs a synchronous, sequential write operation on an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrite(_:_:_:)
func ExtAudioFileWrite(inExtAudioFile unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWrite(inExtAudioFile, inNumberFrames, ioData)
	}


// Perform an asynchronous, sequential write operation on an audio file. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWriteAsync(_:_:_:)
func ExtAudioFileWriteAsync(inExtAudioFile unsafe.Pointer, inNumberFrames unsafe.Pointer, ioData unsafe.Pointer) unsafe.Pointer {
	return _ExtAudioFileWriteAsync(inExtAudioFile, inNumberFrames, ioData)
	}


// Gets the name of a sound bank from a sound bank file. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetNameFromSoundBank
func GetNameFromSoundBank(inSoundBankRef unsafe.Pointer, outName unsafe.Pointer) unsafe.Pointer {
	return _GetNameFromSoundBank(inSoundBankRef, outName)
	}


// MusicDeviceMIDIEvent is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEvent(_:_:_:_:_:)
func MusicDeviceMIDIEvent(inUnit unsafe.Pointer, inStatus unsafe.Pointer, inData1 unsafe.Pointer, inData2 unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceMIDIEvent(inUnit, inStatus, inData1, inData2, inOffsetSampleFrame)
	}


// MusicDeviceMIDIEventList is a AudioToolbox function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEventList(_:_:_:)
func MusicDeviceMIDIEventList(inUnit unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer, evtList unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceMIDIEventList(inUnit, inOffsetSampleFrame, evtList)
	}


// MusicDevicePrepareInstrument is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDevicePrepareInstrument
func MusicDevicePrepareInstrument(inUnit unsafe.Pointer, inInstrument unsafe.Pointer) unsafe.Pointer {
	return _MusicDevicePrepareInstrument(inUnit, inInstrument)
	}


// MusicDeviceReleaseInstrument is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceReleaseInstrument
func MusicDeviceReleaseInstrument(inUnit unsafe.Pointer, inInstrument unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceReleaseInstrument(inUnit, inInstrument)
	}


// MusicDeviceStartNote is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNote(_:_:_:_:_:_:)
func MusicDeviceStartNote(inUnit unsafe.Pointer, inInstrument unsafe.Pointer, inGroupID unsafe.Pointer, outNoteInstanceID unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer, inParams unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceStartNote(inUnit, inInstrument, inGroupID, outNoteInstanceID, inOffsetSampleFrame, inParams)
	}


// MusicDeviceStopNote is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNote(_:_:_:_:)
func MusicDeviceStopNote(inUnit unsafe.Pointer, inGroupID unsafe.Pointer, inNoteInstanceID unsafe.Pointer, inOffsetSampleFrame unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceStopNote(inUnit, inGroupID, inNoteInstanceID, inOffsetSampleFrame)
	}


// MusicDeviceSysEx is a AudioToolbox function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysEx(_:_:_:)
func MusicDeviceSysEx(inUnit unsafe.Pointer, inData unsafe.Pointer, inLength unsafe.Pointer) unsafe.Pointer {
	return _MusicDeviceSysEx(inUnit, inData, inLength)
	}


// Deletes the event at a music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorDeleteEvent(_:)
func MusicEventIteratorDeleteEvent(inIterator unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorDeleteEvent(inIterator)
	}


// Gets information about the event at a music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorGetEventInfo(_:_:_:_:_:)
func MusicEventIteratorGetEventInfo(inIterator unsafe.Pointer, outTimeStamp unsafe.Pointer, outEventType unsafe.Pointer, outEventData unsafe.Pointer, outEventDataSize unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorGetEventInfo(inIterator, outTimeStamp, outEventType, outEventData, outEventDataSize)
	}


// Indicates whether or not a music track contains an event at the music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasCurrentEvent(_:_:)
func MusicEventIteratorHasCurrentEvent(inIterator unsafe.Pointer, outHasCurEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasCurrentEvent(inIterator, outHasCurEvent)
	}


// Indicates whether or not a music track contains an event beyond the music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasNextEvent(_:_:)
func MusicEventIteratorHasNextEvent(inIterator unsafe.Pointer, outHasNextEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasNextEvent(inIterator, outHasNextEvent)
	}


// Indicates whether or not a music track contains an event before the music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasPreviousEvent(_:_:)
func MusicEventIteratorHasPreviousEvent(inIterator unsafe.Pointer, outHasPrevEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorHasPreviousEvent(inIterator, outHasPrevEvent)
	}


// Positions a music event iterator at the next event on a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorNextEvent(_:)
func MusicEventIteratorNextEvent(inIterator unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorNextEvent(inIterator)
	}


// Positions a music event iterator at the previous event on a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorPreviousEvent(_:)
func MusicEventIteratorPreviousEvent(inIterator unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorPreviousEvent(inIterator)
	}


// Positions a music event iterator at a specified timestamp, in beats. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSeek(_:_:)
func MusicEventIteratorSeek(inIterator unsafe.Pointer, inTimeStamp unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorSeek(inIterator, inTimeStamp)
	}


// Sets information for the event at a music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventInfo(_:_:_:)
func MusicEventIteratorSetEventInfo(inIterator unsafe.Pointer, inEventType unsafe.Pointer, inEventData unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorSetEventInfo(inIterator, inEventType, inEventData)
	}


// Sets the timestamp for the event at a music event iterator’s current position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventTime(_:_:)
func MusicEventIteratorSetEventTime(inIterator unsafe.Pointer, inTimeStamp unsafe.Pointer) unsafe.Pointer {
	return _MusicEventIteratorSetEventTime(inIterator, inTimeStamp)
	}


// Gets the beat number associated a specified host time. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetBeatsForHostTime(_:_:_:)
func MusicPlayerGetBeatsForHostTime(inPlayer unsafe.Pointer, inHostTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetBeatsForHostTime(inPlayer, inHostTime, outBeats)
	}


// Gets the host time associated with a specified beat. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetHostTimeForBeats(_:_:_:)
func MusicPlayerGetHostTimeForBeats(inPlayer unsafe.Pointer, inBeats unsafe.Pointer, outHostTime unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetHostTimeForBeats(inPlayer, inBeats, outHostTime)
	}


// Gets the playback rate multiplier for a music player. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetPlayRateScalar(_:_:)
func MusicPlayerGetPlayRateScalar(inPlayer unsafe.Pointer, outScaleRate unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetPlayRateScalar(inPlayer, outScaleRate)
	}


// Gets the music sequence associated with a music player. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetSequence(_:_:)
func MusicPlayerGetSequence(inPlayer unsafe.Pointer, outSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetSequence(inPlayer, outSequence)
	}


// Gets the playback point for a music player, in beats. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetTime(_:_:)
func MusicPlayerGetTime(inPlayer unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerGetTime(inPlayer, outTime)
	}


// Indicates whether or not a music player is playing. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerIsPlaying(_:_:)
func MusicPlayerIsPlaying(inPlayer unsafe.Pointer, outIsPlaying unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerIsPlaying(inPlayer, outIsPlaying)
	}


// Prepares a music player to play. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerPreroll(_:)
func MusicPlayerPreroll(inPlayer unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerPreroll(inPlayer)
	}


// Sets a playback rate multiplier for a music player. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetPlayRateScalar(_:_:)
func MusicPlayerSetPlayRateScalar(inPlayer unsafe.Pointer, inScaleRate unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerSetPlayRateScalar(inPlayer, inScaleRate)
	}


// Sets the music sequence for the music player to play. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetSequence(_:_:)
func MusicPlayerSetSequence(inPlayer unsafe.Pointer, inSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerSetSequence(inPlayer, inSequence)
	}


// Sets the playback point for a music player, in beats. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetTime(_:_:)
func MusicPlayerSetTime(inPlayer unsafe.Pointer, inTime unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerSetTime(inPlayer, inTime)
	}


// Starts playback of a music player. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStart(_:)
func MusicPlayerStart(inPlayer unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerStart(inPlayer)
	}


// Stops playback of a music player. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStop(_:)
func MusicPlayerStop(inPlayer unsafe.Pointer) unsafe.Pointer {
	return _MusicPlayerStop(inPlayer)
	}


// Formats a music sequence’s bar-beat time to its beat time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBarBeatTimeToBeats(_:_:_:)
func MusicSequenceBarBeatTimeToBeats(inSequence unsafe.Pointer, inBarBeatTime unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceBarBeatTimeToBeats(inSequence, inBarBeatTime, outBeats)
	}


// Formats a music sequence’s beat time to its bar-beat time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBeatsToBarBeatTime(_:_:_:_:)
func MusicSequenceBeatsToBarBeatTime(inSequence unsafe.Pointer, inBeats unsafe.Pointer, inSubbeatDivisor unsafe.Pointer, outBarBeatTime unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceBeatsToBarBeatTime(inSequence, inBeats, inSubbeatDivisor, outBarBeatTime)
	}


// Removes a music track from a music sequence, and disposes of the track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceDisposeTrack(_:_:)
func MusicSequenceDisposeTrack(inSequence unsafe.Pointer, inTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceDisposeTrack(inSequence, inTrack)
	}


// Creates a MIDI file from the events in a music sequence. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreate(_:_:_:_:_:)
func MusicSequenceFileCreate(inSequence unsafe.Pointer, inFileRef unsafe.Pointer, inFileType unsafe.Pointer, inFlags unsafe.Pointer, inResolution unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileCreate(inSequence, inFileRef, inFileType, inFlags, inResolution)
	}


// Creates a data object containing the events from a music sequence. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreateData(_:_:_:_:_:)
func MusicSequenceFileCreateData(inSequence unsafe.Pointer, inFileType unsafe.Pointer, inFlags unsafe.Pointer, inResolution unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileCreateData(inSequence, inFileType, inFlags, inResolution, outData)
	}


// Loads data into a music sequence from a URL reference. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoad(_:_:_:_:)
func MusicSequenceFileLoad(inSequence unsafe.Pointer, inFileRef unsafe.Pointer, inFileTypeHint unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileLoad(inSequence, inFileRef, inFileTypeHint, inFlags)
	}


// Load data into a music sequence from a data reference. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoadData(_:_:_:_:)
func MusicSequenceFileLoadData(inSequence unsafe.Pointer, inData unsafe.Pointer, inFileTypeHint unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceFileLoadData(inSequence, inData, inFileTypeHint, inFlags)
	}


// Gets the audio processing graph associated with a music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetAUGraph(_:_:)
func MusicSequenceGetAUGraph(inSequence unsafe.Pointer, outGraph unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetAUGraph(inSequence, outGraph)
	}


// Calculates the number of beats that correspond to a number of seconds. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetBeatsForSeconds(_:_:_:)
func MusicSequenceGetBeatsForSeconds(inSequence unsafe.Pointer, inSeconds unsafe.Pointer, outBeats unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetBeatsForSeconds(inSequence, inSeconds, outBeats)
	}


// Gets the music track at the specified track index. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetIndTrack(_:_:_:)
func MusicSequenceGetIndTrack(inSequence unsafe.Pointer, inTrackIndex unsafe.Pointer, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetIndTrack(inSequence, inTrackIndex, outTrack)
	}


// Returns a dictionary containing music sequence information. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetInfoDictionary(_:)
func MusicSequenceGetInfoDictionary(inSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetInfoDictionary(inSequence)
	}


// Calculates the number of seconds that correspond to a number of beats. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSecondsForBeats(_:_:_:)
func MusicSequenceGetSecondsForBeats(inSequence unsafe.Pointer, inBeats unsafe.Pointer, outSeconds unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetSecondsForBeats(inSequence, inBeats, outSeconds)
	}


// Gets the sequence type for a music sequence. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSequenceType(_:_:)
func MusicSequenceGetSequenceType(inSequence unsafe.Pointer, outType unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetSequenceType(inSequence, outType)
	}


// Gets the tempo track for a music sequence. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTempoTrack(_:_:)
func MusicSequenceGetTempoTrack(inSequence unsafe.Pointer, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTempoTrack(inSequence, outTrack)
	}


// Gets the number of music tracks owned by a music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackCount(_:_:)
func MusicSequenceGetTrackCount(inSequence unsafe.Pointer, outNumberOfTracks unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTrackCount(inSequence, outNumberOfTracks)
	}


// Gets the index number for a specified music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackIndex(_:_:_:)
func MusicSequenceGetTrackIndex(inSequence unsafe.Pointer, inTrack unsafe.Pointer, outTrackIndex unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceGetTrackIndex(inSequence, inTrack, outTrackIndex)
	}


// MusicSequenceLoadSMFDataWithFlags is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFDataWithFlags
func MusicSequenceLoadSMFDataWithFlags(inSequence unsafe.Pointer, inData unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceLoadSMFDataWithFlags(inSequence, inData, inFlags)
	}


// MusicSequenceLoadSMFWithFlags is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFWithFlags
func MusicSequenceLoadSMFWithFlags(inSequence unsafe.Pointer, inFileRef unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceLoadSMFWithFlags(inSequence, inFileRef, inFlags)
	}


// Add a new, empty music track to a music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceNewTrack(_:_:)
func MusicSequenceNewTrack(inSequence unsafe.Pointer, outTrack unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceNewTrack(inSequence, outTrack)
	}


// Reverses the MIDI and tempo events in a music sequence, so the start becomes the end. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceReverse(_:)
func MusicSequenceReverse(inSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceReverse(inSequence)
	}


// MusicSequenceSaveMIDIFile is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveMIDIFile
func MusicSequenceSaveMIDIFile(inSequence unsafe.Pointer, inParentDirectory unsafe.Pointer, inFileName unsafe.Pointer, inResolution unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSaveMIDIFile(inSequence, inParentDirectory, inFileName, inResolution, inFlags)
	}


// MusicSequenceSaveSMFData is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveSMFData
func MusicSequenceSaveSMFData(inSequence unsafe.Pointer, outData unsafe.Pointer, inResolution unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSaveSMFData(inSequence, outData, inResolution)
	}


// Associates an audio processing graph with a music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetAUGraph(_:_:)
func MusicSequenceSetAUGraph(inSequence unsafe.Pointer, inGraph unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetAUGraph(inSequence, inGraph)
	}


// Associates a specified MIDI endpoint with all music tracks in a music sequence. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetMIDIEndpoint(_:_:)
func MusicSequenceSetMIDIEndpoint(inSequence unsafe.Pointer, inEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetMIDIEndpoint(inSequence, inEndpoint)
	}


// Sets the sequence type for a music sequence. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetSequenceType(_:_:)
func MusicSequenceSetSequenceType(inSequence unsafe.Pointer, inType unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetSequenceType(inSequence, inType)
	}


// Registers a user callback function with a music sequence. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetUserCallback(_:_:_:)
func MusicSequenceSetUserCallback(inSequence unsafe.Pointer, inCallback unsafe.Pointer, inClientData unsafe.Pointer) unsafe.Pointer {
	return _MusicSequenceSetUserCallback(inSequence, inCallback, inClientData)
	}


// Removes a specified range of music track events. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackClear(_:_:_:)
func MusicTrackClear(inTrack unsafe.Pointer, inStartTime unsafe.Pointer, inEndTime unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackClear(inTrack, inStartTime, inEndTime)
	}


// Copies a range of events from one music track and inserts them into another music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCopyInsert(_:_:_:_:_:)
func MusicTrackCopyInsert(inSourceTrack unsafe.Pointer, inSourceStartTime unsafe.Pointer, inSourceEndTime unsafe.Pointer, inDestTrack unsafe.Pointer, inDestInsertTime unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackCopyInsert(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
	}


// Removes a specified range of music track events, and shifts later events toward the start of the track to fill in the gap. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCut(_:_:_:)
func MusicTrackCut(inTrack unsafe.Pointer, inStartTime unsafe.Pointer, inEndTime unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackCut(inTrack, inStartTime, inEndTime)
	}


// Gets the MIDI endpoint that is the event target for a music track. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestMIDIEndpoint(_:_:)
func MusicTrackGetDestMIDIEndpoint(inTrack unsafe.Pointer, outEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetDestMIDIEndpoint(inTrack, outEndpoint)
	}


// Gets the audio unit node that is the event target for a music track. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestNode(_:_:)
func MusicTrackGetDestNode(inTrack unsafe.Pointer, outNode unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetDestNode(inTrack, outNode)
	}


// Gets a music track property value. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetProperty(_:_:_:_:)
func MusicTrackGetProperty(inTrack unsafe.Pointer, inPropertyID unsafe.Pointer, outData unsafe.Pointer, ioLength unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetProperty(inTrack, inPropertyID, outData, ioLength)
	}


// Gets the music sequence that the music track is a member of. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetSequence(_:_:)
func MusicTrackGetSequence(inTrack unsafe.Pointer, outSequence unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackGetSequence(inTrack, outSequence)
	}


// Copies a range of events from one music track and merges them into another music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMerge(_:_:_:_:_:)
func MusicTrackMerge(inSourceTrack unsafe.Pointer, inSourceStartTime unsafe.Pointer, inSourceEndTime unsafe.Pointer, inDestTrack unsafe.Pointer, inDestInsertTime unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackMerge(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
	}


// Shifts music track events forward or backward in time, in terms of beats. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMoveEvents(_:_:_:_:)
func MusicTrackMoveEvents(inTrack unsafe.Pointer, inStartTime unsafe.Pointer, inEndTime unsafe.Pointer, inMoveTime unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackMoveEvents(inTrack, inStartTime, inEndTime, inMoveTime)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewAUPresetEvent(_:_:_:)
func MusicTrackNewAUPresetEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inPresetEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewAUPresetEvent(inTrack, inTimeStamp, inPresetEvent)
	}


// MusicTrackNewExtendedControlEvent is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedControlEvent
func MusicTrackNewExtendedControlEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedControlEvent(inTrack, inTimeStamp, inInfo)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedNoteEvent(_:_:_:)
func MusicTrackNewExtendedNoteEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedNoteEvent(inTrack, inTimeStamp, inInfo)
	}


// Adds a tempo to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedTempoEvent(_:_:_:)
func MusicTrackNewExtendedTempoEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inBPM unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewExtendedTempoEvent(inTrack, inTimeStamp, inBPM)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIChannelEvent(_:_:_:)
func MusicTrackNewMIDIChannelEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inMessage unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDIChannelEvent(inTrack, inTimeStamp, inMessage)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDINoteEvent(_:_:_:)
func MusicTrackNewMIDINoteEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inMessage unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDINoteEvent(inTrack, inTimeStamp, inMessage)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIRawDataEvent(_:_:_:)
func MusicTrackNewMIDIRawDataEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inRawData unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMIDIRawDataEvent(inTrack, inTimeStamp, inRawData)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMetaEvent(_:_:_:)
func MusicTrackNewMetaEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inMetaEvent unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewMetaEvent(inTrack, inTimeStamp, inMetaEvent)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewParameterEvent(_:_:_:)
func MusicTrackNewParameterEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inInfo unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewParameterEvent(inTrack, inTimeStamp, inInfo)
	}


// Adds an event of type to a music track. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewUserEvent(_:_:_:)
func MusicTrackNewUserEvent(inTrack unsafe.Pointer, inTimeStamp unsafe.Pointer, inUserData unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackNewUserEvent(inTrack, inTimeStamp, inUserData)
	}


// Sets the music track’s event target to a MIDI endpoint. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestMIDIEndpoint(_:_:)
func MusicTrackSetDestMIDIEndpoint(inTrack unsafe.Pointer, inEndpoint unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackSetDestMIDIEndpoint(inTrack, inEndpoint)
	}


// Sets the music track’s event target to an audio unit node. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestNode(_:_:)
func MusicTrackSetDestNode(inTrack unsafe.Pointer, inNode unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackSetDestNode(inTrack, inNode)
	}


// Sets a music track property value. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetProperty(_:_:_:_:)
func MusicTrackSetProperty(inTrack unsafe.Pointer, inPropertyID unsafe.Pointer, inData unsafe.Pointer, inLength unsafe.Pointer) unsafe.Pointer {
	return _MusicTrackSetProperty(inTrack, inPropertyID, inData, inLength)
	}


// Creates a new audio processing graph. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewAUGraph(_:)
func NewAUGraph(outGraph unsafe.Pointer) unsafe.Pointer {
	return _NewAUGraph(outGraph)
	}


// Creates a new music event iterator. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicEventIterator(_:_:)
func NewMusicEventIterator(inTrack unsafe.Pointer, outIterator unsafe.Pointer) unsafe.Pointer {
	return _NewMusicEventIterator(inTrack, outIterator)
	}


// Creates a new music player. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicPlayer(_:)
func NewMusicPlayer(outPlayer unsafe.Pointer) unsafe.Pointer {
	return _NewMusicPlayer(outPlayer)
	}


// Creates a new empty music sequence. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicSequence(_:)
func NewMusicSequence(outSequence unsafe.Pointer) unsafe.Pointer {
	return _NewMusicSequence(outSequence)
	}


// NewMusicTrackFrom is a AudioToolbox function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NewMusicTrackFrom
func NewMusicTrackFrom(inSourceTrack unsafe.Pointer, inSourceStartTime unsafe.Pointer, inSourceEndTime unsafe.Pointer, outNewTrack unsafe.Pointer) unsafe.Pointer {
	return _NewMusicTrackFrom(inSourceTrack, inSourceStartTime, inSourceEndTime, outNewTrack)
	}


// GetAudioUnitParameterDisplayType is a AudioToolbox function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/GetAudioUnitParameterDisplayType(_:)
func GetAudioUnitParameterDisplayType(p0 unsafe.Pointer) unsafe.Pointer {
	return _GetAudioUnitParameterDisplayType(p0)
	}


// SetAudioUnitParameterDisplayType is a AudioToolbox function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/SetAudioUnitParameterDisplayType(_:_:)
func SetAudioUnitParameterDisplayType(p0 unsafe.Pointer) unsafe.Pointer {
	return _SetAudioUnitParameterDisplayType(p0)
	}




