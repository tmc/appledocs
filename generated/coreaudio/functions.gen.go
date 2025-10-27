// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio


import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreAudio Functions (57 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AudioConvertHostTimeToNanos func(unsafe.Pointer) unsafe.Pointer
	_AudioConvertNanosToHostTime func(unsafe.Pointer) unsafe.Pointer
	_AudioDeviceAddIOProc func(AudioDeviceID, AudioDeviceIOProc, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceAddPropertyListener func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, AudioDevicePropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceCreateIOProcID func(AudioObjectID, AudioDeviceIOProc, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceCreateIOProcIDWithBlock func(unsafe.Pointer, AudioObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceDestroyIOProcID func(AudioObjectID, AudioDeviceIOProcID) unsafe.Pointer
	_AudioDeviceGetCurrentTime func(AudioObjectID, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceGetNearestStartTime func(AudioObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceGetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceGetPropertyInfo func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceRead func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceRemoveIOProc func(AudioDeviceID, AudioDeviceIOProc) unsafe.Pointer
	_AudioDeviceRemovePropertyListener func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, AudioDevicePropertyListenerProc) unsafe.Pointer
	_AudioDeviceSetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceStart func(AudioObjectID, AudioDeviceIOProcID) unsafe.Pointer
	_AudioDeviceStartAtTime func(AudioObjectID, AudioDeviceIOProcID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDeviceStop func(AudioObjectID, AudioDeviceIOProcID) unsafe.Pointer
	_AudioDeviceTranslateTime func(AudioObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInClose func(AudioDeviceID) unsafe.Pointer
	_AudioDriverPlugInDeviceGetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInDeviceGetPropertyInfo func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInDeviceSetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInOpen func(unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInStreamGetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInStreamGetPropertyInfo func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioDriverPlugInStreamSetProperty func(AudioDeviceID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioGetCurrentHostTime func() unsafe.Pointer
	_AudioGetHostClockFrequency func() unsafe.Pointer
	_AudioGetHostClockMinimumTimeDelta func() unsafe.Pointer
	_AudioHardwareAddPropertyListener func(AudioHardwarePropertyID, AudioHardwarePropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareAddRunLoopSource func(RunLoopSourceRef) unsafe.Pointer
	_AudioHardwareCreateAggregateDevice func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareCreateProcessTap func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareDestroyAggregateDevice func(AudioObjectID) unsafe.Pointer
	_AudioHardwareDestroyProcessTap func(AudioObjectID) unsafe.Pointer
	_AudioHardwareGetProperty func(AudioHardwarePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareGetPropertyInfo func(AudioHardwarePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareRemovePropertyListener func(AudioHardwarePropertyID, AudioHardwarePropertyListenerProc) unsafe.Pointer
	_AudioHardwareRemoveRunLoopSource func(RunLoopSourceRef) unsafe.Pointer
	_AudioHardwareSetProperty func(AudioHardwarePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioHardwareUnload func() unsafe.Pointer
	_AudioObjectAddPropertyListener func(AudioObjectID, unsafe.Pointer, AudioObjectPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioObjectAddPropertyListenerBlock func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectGetPropertyData func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectGetPropertyDataSize func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectHasProperty func(AudioObjectID, unsafe.Pointer) unsafe.Pointer
	_AudioObjectIsPropertySettable func(AudioObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectRemovePropertyListener func(AudioObjectID, unsafe.Pointer, AudioObjectPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioObjectRemovePropertyListenerBlock func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectSetPropertyData func(AudioObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioObjectShow func(AudioObjectID)
	_AudioStreamAddPropertyListener func(AudioStreamID, unsafe.Pointer, AudioDevicePropertyID, AudioStreamPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_AudioStreamGetProperty func(AudioStreamID, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioStreamGetPropertyInfo func(AudioStreamID, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AudioStreamRemovePropertyListener func(AudioStreamID, unsafe.Pointer, AudioDevicePropertyID, AudioStreamPropertyListenerProc) unsafe.Pointer
	_AudioStreamSetProperty func(AudioStreamID, unsafe.Pointer, unsafe.Pointer, AudioDevicePropertyID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AudioConvertHostTimeToNanos, lib, "AudioConvertHostTimeToNanos")
	tryRegister(&_AudioConvertNanosToHostTime, lib, "AudioConvertNanosToHostTime")
	tryRegister(&_AudioDeviceAddIOProc, lib, "AudioDeviceAddIOProc")
	tryRegister(&_AudioDeviceAddPropertyListener, lib, "AudioDeviceAddPropertyListener")
	tryRegister(&_AudioDeviceCreateIOProcID, lib, "AudioDeviceCreateIOProcID")
	tryRegister(&_AudioDeviceCreateIOProcIDWithBlock, lib, "AudioDeviceCreateIOProcIDWithBlock")
	tryRegister(&_AudioDeviceDestroyIOProcID, lib, "AudioDeviceDestroyIOProcID")
	tryRegister(&_AudioDeviceGetCurrentTime, lib, "AudioDeviceGetCurrentTime")
	tryRegister(&_AudioDeviceGetNearestStartTime, lib, "AudioDeviceGetNearestStartTime")
	tryRegister(&_AudioDeviceGetProperty, lib, "AudioDeviceGetProperty")
	tryRegister(&_AudioDeviceGetPropertyInfo, lib, "AudioDeviceGetPropertyInfo")
	tryRegister(&_AudioDeviceRead, lib, "AudioDeviceRead")
	tryRegister(&_AudioDeviceRemoveIOProc, lib, "AudioDeviceRemoveIOProc")
	tryRegister(&_AudioDeviceRemovePropertyListener, lib, "AudioDeviceRemovePropertyListener")
	tryRegister(&_AudioDeviceSetProperty, lib, "AudioDeviceSetProperty")
	tryRegister(&_AudioDeviceStart, lib, "AudioDeviceStart")
	tryRegister(&_AudioDeviceStartAtTime, lib, "AudioDeviceStartAtTime")
	tryRegister(&_AudioDeviceStop, lib, "AudioDeviceStop")
	tryRegister(&_AudioDeviceTranslateTime, lib, "AudioDeviceTranslateTime")
	tryRegister(&_AudioDriverPlugInClose, lib, "AudioDriverPlugInClose")
	tryRegister(&_AudioDriverPlugInDeviceGetProperty, lib, "AudioDriverPlugInDeviceGetProperty")
	tryRegister(&_AudioDriverPlugInDeviceGetPropertyInfo, lib, "AudioDriverPlugInDeviceGetPropertyInfo")
	tryRegister(&_AudioDriverPlugInDeviceSetProperty, lib, "AudioDriverPlugInDeviceSetProperty")
	tryRegister(&_AudioDriverPlugInOpen, lib, "AudioDriverPlugInOpen")
	tryRegister(&_AudioDriverPlugInStreamGetProperty, lib, "AudioDriverPlugInStreamGetProperty")
	tryRegister(&_AudioDriverPlugInStreamGetPropertyInfo, lib, "AudioDriverPlugInStreamGetPropertyInfo")
	tryRegister(&_AudioDriverPlugInStreamSetProperty, lib, "AudioDriverPlugInStreamSetProperty")
	tryRegister(&_AudioGetCurrentHostTime, lib, "AudioGetCurrentHostTime")
	tryRegister(&_AudioGetHostClockFrequency, lib, "AudioGetHostClockFrequency")
	tryRegister(&_AudioGetHostClockMinimumTimeDelta, lib, "AudioGetHostClockMinimumTimeDelta")
	tryRegister(&_AudioHardwareAddPropertyListener, lib, "AudioHardwareAddPropertyListener")
	tryRegister(&_AudioHardwareAddRunLoopSource, lib, "AudioHardwareAddRunLoopSource")
	tryRegister(&_AudioHardwareCreateAggregateDevice, lib, "AudioHardwareCreateAggregateDevice")
	tryRegister(&_AudioHardwareCreateProcessTap, lib, "AudioHardwareCreateProcessTap")
	tryRegister(&_AudioHardwareDestroyAggregateDevice, lib, "AudioHardwareDestroyAggregateDevice")
	tryRegister(&_AudioHardwareDestroyProcessTap, lib, "AudioHardwareDestroyProcessTap")
	tryRegister(&_AudioHardwareGetProperty, lib, "AudioHardwareGetProperty")
	tryRegister(&_AudioHardwareGetPropertyInfo, lib, "AudioHardwareGetPropertyInfo")
	tryRegister(&_AudioHardwareRemovePropertyListener, lib, "AudioHardwareRemovePropertyListener")
	tryRegister(&_AudioHardwareRemoveRunLoopSource, lib, "AudioHardwareRemoveRunLoopSource")
	tryRegister(&_AudioHardwareSetProperty, lib, "AudioHardwareSetProperty")
	tryRegister(&_AudioHardwareUnload, lib, "AudioHardwareUnload")
	tryRegister(&_AudioObjectAddPropertyListener, lib, "AudioObjectAddPropertyListener")
	tryRegister(&_AudioObjectAddPropertyListenerBlock, lib, "AudioObjectAddPropertyListenerBlock")
	tryRegister(&_AudioObjectGetPropertyData, lib, "AudioObjectGetPropertyData")
	tryRegister(&_AudioObjectGetPropertyDataSize, lib, "AudioObjectGetPropertyDataSize")
	tryRegister(&_AudioObjectHasProperty, lib, "AudioObjectHasProperty")
	tryRegister(&_AudioObjectIsPropertySettable, lib, "AudioObjectIsPropertySettable")
	tryRegister(&_AudioObjectRemovePropertyListener, lib, "AudioObjectRemovePropertyListener")
	tryRegister(&_AudioObjectRemovePropertyListenerBlock, lib, "AudioObjectRemovePropertyListenerBlock")
	tryRegister(&_AudioObjectSetPropertyData, lib, "AudioObjectSetPropertyData")
	tryRegister(&_AudioObjectShow, lib, "AudioObjectShow")
	tryRegister(&_AudioStreamAddPropertyListener, lib, "AudioStreamAddPropertyListener")
	tryRegister(&_AudioStreamGetProperty, lib, "AudioStreamGetProperty")
	tryRegister(&_AudioStreamGetPropertyInfo, lib, "AudioStreamGetPropertyInfo")
	tryRegister(&_AudioStreamRemovePropertyListener, lib, "AudioStreamRemovePropertyListener")
	tryRegister(&_AudioStreamSetProperty, lib, "AudioStreamSetProperty")
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



// AudioConvertHostTimeToNanos is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioConvertHostTimeToNanos(_:)
func AudioConvertHostTimeToNanos(inHostTime unsafe.Pointer) unsafe.Pointer {
	return _AudioConvertHostTimeToNanos(inHostTime)
}

// AudioConvertNanosToHostTime is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioConvertNanosToHostTime(_:)
func AudioConvertNanosToHostTime(inNanos unsafe.Pointer) unsafe.Pointer {
	return _AudioConvertNanosToHostTime(inNanos)
}

// AudioDeviceAddIOProc is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddIOProc
func AudioDeviceAddIOProc(inDevice AudioDeviceID, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceAddIOProc(inDevice, inProc, inClientData)
}

// AudioDeviceAddPropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddPropertyListener
func AudioDeviceAddPropertyListener(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceAddPropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc, inClientData)
}

// AudioDeviceCreateIOProcID is a CoreAudio function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcID(_:_:_:_:)
func AudioDeviceCreateIOProcID(inDevice AudioObjectID, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceCreateIOProcID(inDevice, inProc, inClientData, outIOProcID)
}

// AudioDeviceCreateIOProcIDWithBlock is a CoreAudio function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcIDWithBlock(_:_:_:_:)
func AudioDeviceCreateIOProcIDWithBlock(outIOProcID unsafe.Pointer, inDevice AudioObjectID, inDispatchQueue unsafe.Pointer, inIOBlock unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceCreateIOProcIDWithBlock(outIOProcID, inDevice, inDispatchQueue, inIOBlock)
}

// AudioDeviceDestroyIOProcID is a CoreAudio function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceDestroyIOProcID(_:_:)
func AudioDeviceDestroyIOProcID(inDevice AudioObjectID, inIOProcID AudioDeviceIOProcID) unsafe.Pointer {
	return _AudioDeviceDestroyIOProcID(inDevice, inIOProcID)
}

// AudioDeviceGetCurrentTime is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetCurrentTime(_:_:)
func AudioDeviceGetCurrentTime(inDevice AudioObjectID, outTime unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceGetCurrentTime(inDevice, outTime)
}

// AudioDeviceGetNearestStartTime is a CoreAudio function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetNearestStartTime(_:_:_:)
func AudioDeviceGetNearestStartTime(inDevice AudioObjectID, ioRequestedStartTime unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceGetNearestStartTime(inDevice, ioRequestedStartTime, inFlags)
}

// AudioDeviceGetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetProperty
func AudioDeviceGetProperty(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioDeviceGetPropertyInfo is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetPropertyInfo
func AudioDeviceGetPropertyInfo(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
}

// AudioDeviceRead is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRead
func AudioDeviceRead(inDevice AudioDeviceID, inStartTime unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceRead(inDevice, inStartTime, outData)
}

// AudioDeviceRemoveIOProc is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemoveIOProc
func AudioDeviceRemoveIOProc(inDevice AudioDeviceID, inProc AudioDeviceIOProc) unsafe.Pointer {
	return _AudioDeviceRemoveIOProc(inDevice, inProc)
}

// AudioDeviceRemovePropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemovePropertyListener
func AudioDeviceRemovePropertyListener(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, inProc AudioDevicePropertyListenerProc) unsafe.Pointer {
	return _AudioDeviceRemovePropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc)
}

// AudioDeviceSetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceSetProperty
func AudioDeviceSetProperty(inDevice AudioDeviceID, inWhen unsafe.Pointer, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
}

// AudioDeviceStart is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStart(_:_:)
func AudioDeviceStart(inDevice AudioObjectID, inProcID AudioDeviceIOProcID) unsafe.Pointer {
	return _AudioDeviceStart(inDevice, inProcID)
}

// AudioDeviceStartAtTime is a CoreAudio function.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStartAtTime(_:_:_:_:)
func AudioDeviceStartAtTime(inDevice AudioObjectID, inProcID AudioDeviceIOProcID, ioRequestedStartTime unsafe.Pointer, inFlags unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceStartAtTime(inDevice, inProcID, ioRequestedStartTime, inFlags)
}

// AudioDeviceStop is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStop(_:_:)
func AudioDeviceStop(inDevice AudioObjectID, inProcID AudioDeviceIOProcID) unsafe.Pointer {
	return _AudioDeviceStop(inDevice, inProcID)
}

// AudioDeviceTranslateTime is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceTranslateTime(_:_:_:)
func AudioDeviceTranslateTime(inDevice AudioObjectID, inTime unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _AudioDeviceTranslateTime(inDevice, inTime, outTime)
}

// AudioDriverPlugInClose is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInClose
func AudioDriverPlugInClose(inDevice AudioDeviceID) unsafe.Pointer {
	return _AudioDriverPlugInClose(inDevice)
}

// AudioDriverPlugInDeviceGetProperty is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetProperty
func AudioDriverPlugInDeviceGetProperty(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioDriverPlugInDeviceGetPropertyInfo is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetPropertyInfo
func AudioDriverPlugInDeviceGetPropertyInfo(inDevice AudioDeviceID, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
}

// AudioDriverPlugInDeviceSetProperty is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceSetProperty
func AudioDriverPlugInDeviceSetProperty(inDevice AudioDeviceID, inWhen unsafe.Pointer, inChannel unsafe.Pointer, isInput unsafe.Pointer, inPropertyID AudioDevicePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
}

// AudioDriverPlugInOpen is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInOpen
func AudioDriverPlugInOpen(inHostInfo unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInOpen(inHostInfo)
}

// AudioDriverPlugInStreamGetProperty is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetProperty
func AudioDriverPlugInStreamGetProperty(inDevice AudioDeviceID, inIOAudioStream unsafe.Pointer, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInStreamGetProperty(inDevice, inIOAudioStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioDriverPlugInStreamGetPropertyInfo is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetPropertyInfo
func AudioDriverPlugInStreamGetPropertyInfo(inDevice AudioDeviceID, inIOAudioStream unsafe.Pointer, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInStreamGetPropertyInfo(inDevice, inIOAudioStream, inChannel, inPropertyID, outSize, outWritable)
}

// AudioDriverPlugInStreamSetProperty is a CoreAudio function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamSetProperty
func AudioDriverPlugInStreamSetProperty(inDevice AudioDeviceID, inIOAudioStream unsafe.Pointer, inWhen unsafe.Pointer, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioDriverPlugInStreamSetProperty(inDevice, inIOAudioStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
}

// AudioGetCurrentHostTime is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioGetCurrentHostTime()
func AudioGetCurrentHostTime() unsafe.Pointer {
	return _AudioGetCurrentHostTime()
}

// AudioGetHostClockFrequency is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockFrequency()
func AudioGetHostClockFrequency() unsafe.Pointer {
	return _AudioGetHostClockFrequency()
}

// AudioGetHostClockMinimumTimeDelta is a CoreAudio function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockMinimumTimeDelta()
func AudioGetHostClockMinimumTimeDelta() unsafe.Pointer {
	return _AudioGetHostClockMinimumTimeDelta()
}

// AudioHardwareAddPropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddPropertyListener
func AudioHardwareAddPropertyListener(inPropertyID AudioHardwarePropertyID, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareAddPropertyListener(inPropertyID, inProc, inClientData)
}

// AudioHardwareAddRunLoopSource is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddRunLoopSource
func AudioHardwareAddRunLoopSource(inRunLoopSource RunLoopSourceRef) unsafe.Pointer {
	return _AudioHardwareAddRunLoopSource(inRunLoopSource)
}

// AudioHardwareCreateAggregateDevice is a CoreAudio function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateAggregateDevice(_:_:)
func AudioHardwareCreateAggregateDevice(inDescription DictionaryRef, outDeviceID unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareCreateAggregateDevice(inDescription, outDeviceID)
}

// AudioHardwareCreateProcessTap is a CoreAudio function.
//
// Added in macOS 14.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateProcessTap(_:_:)
func AudioHardwareCreateProcessTap(inDescription unsafe.Pointer, outTapID unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareCreateProcessTap(inDescription, outTapID)
}

// AudioHardwareDestroyAggregateDevice is a CoreAudio function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyAggregateDevice(_:)
func AudioHardwareDestroyAggregateDevice(inDeviceID AudioObjectID) unsafe.Pointer {
	return _AudioHardwareDestroyAggregateDevice(inDeviceID)
}

// AudioHardwareDestroyProcessTap is a CoreAudio function.
//
// Added in macOS 14.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyProcessTap(_:)
func AudioHardwareDestroyProcessTap(inTapID AudioObjectID) unsafe.Pointer {
	return _AudioHardwareDestroyProcessTap(inTapID)
}

// AudioHardwareGetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetProperty
func AudioHardwareGetProperty(inPropertyID AudioHardwarePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareGetProperty(inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioHardwareGetPropertyInfo is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetPropertyInfo
func AudioHardwareGetPropertyInfo(inPropertyID AudioHardwarePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareGetPropertyInfo(inPropertyID, outSize, outWritable)
}

// AudioHardwareRemovePropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemovePropertyListener
func AudioHardwareRemovePropertyListener(inPropertyID AudioHardwarePropertyID, inProc AudioHardwarePropertyListenerProc) unsafe.Pointer {
	return _AudioHardwareRemovePropertyListener(inPropertyID, inProc)
}

// AudioHardwareRemoveRunLoopSource is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemoveRunLoopSource
func AudioHardwareRemoveRunLoopSource(inRunLoopSource RunLoopSourceRef) unsafe.Pointer {
	return _AudioHardwareRemoveRunLoopSource(inRunLoopSource)
}

// AudioHardwareSetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareSetProperty
func AudioHardwareSetProperty(inPropertyID AudioHardwarePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioHardwareSetProperty(inPropertyID, inPropertyDataSize, inPropertyData)
}

// AudioHardwareUnload is a CoreAudio function.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareUnload()
func AudioHardwareUnload() unsafe.Pointer {
	return _AudioHardwareUnload()
}

// AudioObjectAddPropertyListener is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListener(_:_:_:_:)
func AudioObjectAddPropertyListener(inObjectID AudioObjectID, inAddress unsafe.Pointer, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectAddPropertyListener(inObjectID, inAddress, inListener, inClientData)
}

// AudioObjectAddPropertyListenerBlock is a CoreAudio function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListenerBlock(_:_:_:_:)
func AudioObjectAddPropertyListenerBlock(inObjectID AudioObjectID, inAddress unsafe.Pointer, inDispatchQueue unsafe.Pointer, inListener unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectAddPropertyListenerBlock(inObjectID, inAddress, inDispatchQueue, inListener)
}

// AudioObjectGetPropertyData is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyData(_:_:_:_:_:_:)
func AudioObjectGetPropertyData(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, ioDataSize unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData)
}

// AudioObjectGetPropertyDataSize is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyDataSize(_:_:_:_:_:)
func AudioObjectGetPropertyDataSize(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, outDataSize unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize)
}

// AudioObjectHasProperty is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectHasProperty(_:_:)
func AudioObjectHasProperty(inObjectID AudioObjectID, inAddress unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectHasProperty(inObjectID, inAddress)
}

// AudioObjectIsPropertySettable is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectIsPropertySettable(_:_:_:)
func AudioObjectIsPropertySettable(inObjectID AudioObjectID, inAddress unsafe.Pointer, outIsSettable unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectIsPropertySettable(inObjectID, inAddress, outIsSettable)
}

// AudioObjectRemovePropertyListener is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListener(_:_:_:_:)
func AudioObjectRemovePropertyListener(inObjectID AudioObjectID, inAddress unsafe.Pointer, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectRemovePropertyListener(inObjectID, inAddress, inListener, inClientData)
}

// AudioObjectRemovePropertyListenerBlock is a CoreAudio function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListenerBlock(_:_:_:_:)
func AudioObjectRemovePropertyListenerBlock(inObjectID AudioObjectID, inAddress unsafe.Pointer, inDispatchQueue unsafe.Pointer, inListener unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectRemovePropertyListenerBlock(inObjectID, inAddress, inDispatchQueue, inListener)
}

// AudioObjectSetPropertyData is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectSetPropertyData(_:_:_:_:_:_:)
func AudioObjectSetPropertyData(inObjectID AudioObjectID, inAddress unsafe.Pointer, inQualifierDataSize unsafe.Pointer, inQualifierData unsafe.Pointer, inDataSize unsafe.Pointer, inData unsafe.Pointer) unsafe.Pointer {
	return _AudioObjectSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData)
}

// AudioObjectShow is a CoreAudio function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectShow(_:)
func AudioObjectShow(inObjectID AudioObjectID) {
	_AudioObjectShow(inObjectID)
}

// AudioStreamAddPropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamAddPropertyListener
func AudioStreamAddPropertyListener(inStream AudioStreamID, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) unsafe.Pointer {
	return _AudioStreamAddPropertyListener(inStream, inChannel, inPropertyID, inProc, inClientData)
}

// AudioStreamGetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetProperty
func AudioStreamGetProperty(inStream AudioStreamID, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, ioPropertyDataSize unsafe.Pointer, outPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioStreamGetProperty(inStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
}

// AudioStreamGetPropertyInfo is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetPropertyInfo
func AudioStreamGetPropertyInfo(inStream AudioStreamID, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, outSize unsafe.Pointer, outWritable unsafe.Pointer) unsafe.Pointer {
	return _AudioStreamGetPropertyInfo(inStream, inChannel, inPropertyID, outSize, outWritable)
}

// AudioStreamRemovePropertyListener is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamRemovePropertyListener
func AudioStreamRemovePropertyListener(inStream AudioStreamID, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, inProc AudioStreamPropertyListenerProc) unsafe.Pointer {
	return _AudioStreamRemovePropertyListener(inStream, inChannel, inPropertyID, inProc)
}

// AudioStreamSetProperty is a CoreAudio function.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamSetProperty
func AudioStreamSetProperty(inStream AudioStreamID, inWhen unsafe.Pointer, inChannel unsafe.Pointer, inPropertyID AudioDevicePropertyID, inPropertyDataSize unsafe.Pointer, inPropertyData unsafe.Pointer) unsafe.Pointer {
	return _AudioStreamSetProperty(inStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
}




