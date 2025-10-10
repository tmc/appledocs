// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

// CoreAudio Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (57 total):

// AudioConvertHostTimeToNanos(inHostTime _, :  UInt64) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioConvertNanosToHostTime(inNanos _, :  UInt64) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioDeviceAddIOProc(inDevice AudioDeviceID, inProc ,  AudioDeviceIOProc, inClientData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.


// AudioDeviceAddPropertyListener(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, inProc ,  AudioDevicePropertyListenerProc, inClientData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioDeviceCreateIOProcID(inDevice _, inProc :  AudioObjectID,  _, inClientData :  AudioDeviceIOProc,  _, outIOProcID :  UnsafeMutableRawPointer?,  _, :  UnsafeMutablePointer< AudioDeviceIOProcID?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+

// AudioDeviceCreateIOProcIDWithBlock(outIOProcID _, inDevice :  UnsafeMutablePointer< AudioDeviceIOProcID?>,  _, inDispatchQueue :  AudioObjectID,  _, inIOBlock :  dispatch_queue_t?,  _, :  @escaping  AudioDeviceIOBlock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.7+


// AudioDeviceDestroyIOProcID(inDevice _, inIOProcID :  AudioObjectID,  _, :  AudioDeviceIOProcID) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+

// AudioDeviceGetCurrentTime(inDevice _, outTime :  AudioObjectID,  _, :  UnsafeMutablePointer< AudioTimeStamp>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioDeviceGetNearestStartTime(inDevice _, ioRequestedStartTime :  AudioObjectID,  _, inFlags :  UnsafeMutablePointer< AudioTimeStamp>,  _, :  UInt32) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+


// AudioDeviceGetProperty(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, ioPropertyDataSize ,  UInt32 *, outPropertyData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioDeviceGetPropertyInfo(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, outSize ,  UInt32 *, outWritable ,  Boolean *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioDeviceRead(inDevice AudioDeviceID, inStartTime ,  const  AudioTimeStamp *, outData ,  AudioBufferList *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.


// AudioDeviceRemoveIOProc(inDevice AudioDeviceID, inProc ,  AudioDeviceIOProc, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.

// AudioDeviceRemovePropertyListener(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, inProc ,  AudioDevicePropertyListenerProc, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioDeviceSetProperty(inDevice AudioDeviceID, inWhen ,  const  AudioTimeStamp *, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, inPropertyDataSize ,  UInt32, inPropertyData ,  const  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// AudioDeviceStart(inDevice _, inProcID :  AudioObjectID,  _, :  AudioDeviceIOProcID?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioDeviceStartAtTime(inDevice _, inProcID :  AudioObjectID,  _, ioRequestedStartTime :  AudioDeviceIOProcID?,  _, inFlags :  UnsafeMutablePointer< AudioTimeStamp>,  _, :  UInt32) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+

// AudioDeviceStop(inDevice _, inProcID :  AudioObjectID,  _, :  AudioDeviceIOProcID?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+


// AudioDeviceTranslateTime(inDevice _, inTime :  AudioObjectID,  _, outTime :  UnsafePointer< AudioTimeStamp>,  _, :  UnsafeMutablePointer< AudioTimeStamp>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioDriverPlugInClose(inDevice AudioDeviceID, );) extern   OSStatus

// AudioDriverPlugInDeviceGetProperty(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, ioPropertyDataSize ,  UInt32 *, outPropertyData ,  void *, );) extern   OSStatus


// AudioDriverPlugInDeviceGetPropertyInfo(inDevice AudioDeviceID, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, outSize ,  UInt32 *, outWritable ,  Boolean *, );) extern   OSStatus

// AudioDriverPlugInDeviceSetProperty(inDevice AudioDeviceID, inWhen ,  const  AudioTimeStamp *, inChannel ,  UInt32, isInput ,  Boolean, inPropertyID ,  AudioDevicePropertyID, inPropertyDataSize ,  UInt32, inPropertyData ,  const  void *, );) extern   OSStatus

// AudioDriverPlugInOpen(inHostInfo AudioDriverPlugInHostInfo *, );) extern   OSStatus


// AudioDriverPlugInStreamGetProperty(inDevice AudioDeviceID, inIOAudioStream ,  io_object_t, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, ioPropertyDataSize ,  UInt32 *, outPropertyData ,  void *, );) extern   OSStatus

// AudioDriverPlugInStreamGetPropertyInfo(inDevice AudioDeviceID, inIOAudioStream ,  io_object_t, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, outSize ,  UInt32 *, outWritable ,  Boolean *, );) extern   OSStatus

// AudioDriverPlugInStreamSetProperty(inDevice AudioDeviceID, inIOAudioStream ,  io_object_t, inWhen ,  const  AudioTimeStamp *, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, inPropertyDataSize ,  UInt32, inPropertyData ,  const  void *, );) extern   OSStatus


// AudioGetCurrentHostTime() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioGetHostClockFrequency() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+

// AudioGetHostClockMinimumTimeDelta() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+


// AudioHardwareAddPropertyListener(inPropertyID AudioHardwarePropertyID, inProc ,  AudioHardwarePropertyListenerProc, inClientData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioHardwareAddRunLoopSource(inRunLoopSource CFRunLoopSourceRef, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.3+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AudioHardwareCreateAggregateDevice(inDescription _, outDeviceID :  CFDictionary,  _, :  UnsafeMutablePointer< AudioObjectID>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+


// AudioHardwareCreateProcessTap(inDescription _, outTapID :  CATapDescription!,  _, :  UnsafeMutablePointer< AudioObjectID>!) ->  OSStatus) func
//
// Availability:
//   - macOS 14.2+

// AudioHardwareDestroyAggregateDevice(inDeviceID _, :  AudioObjectID) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+

// AudioHardwareDestroyProcessTap(inTapID _, :  AudioObjectID) ->  OSStatus) func
//
// Availability:
//   - macOS 14.2+


// AudioHardwareGetProperty(inPropertyID AudioHardwarePropertyID, ioPropertyDataSize ,  UInt32 *, outPropertyData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioHardwareGetPropertyInfo(inPropertyID AudioHardwarePropertyID, outSize ,  UInt32 *, outWritable ,  Boolean *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioHardwareRemovePropertyListener(inPropertyID AudioHardwarePropertyID, inProc ,  AudioHardwarePropertyListenerProc, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// AudioHardwareRemoveRunLoopSource(inRunLoopSource CFRunLoopSourceRef, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.3+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AudioHardwareSetProperty(inPropertyID AudioHardwarePropertyID, inPropertyDataSize ,  UInt32, inPropertyData ,  const  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioHardwareUnload() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.1+


// AudioObjectAddPropertyListener(inObjectID _, inAddress :  AudioObjectID,  _, inListener :  UnsafePointer< AudioObjectPropertyAddress>,  _, inClientData :  AudioObjectPropertyListenerProc,  _, :  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+

// AudioObjectAddPropertyListenerBlock(inObjectID _, inAddress :  AudioObjectID,  _, inDispatchQueue :  UnsafePointer< AudioObjectPropertyAddress>,  _, inListener :  dispatch_queue_t?,  _, :  @escaping  AudioObjectPropertyListenerBlock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.7+

// AudioObjectGetPropertyData(inObjectID _, inAddress :  AudioObjectID,  _, inQualifierDataSize :  UnsafePointer< AudioObjectPropertyAddress>,  _, inQualifierData :  UInt32,  _, ioDataSize :  UnsafeRawPointer?,  _, outData :  UnsafeMutablePointer< UInt32>,  _, :  UnsafeMutableRawPointer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+


// AudioObjectGetPropertyDataSize(inObjectID _, inAddress :  AudioObjectID,  _, inQualifierDataSize :  UnsafePointer< AudioObjectPropertyAddress>,  _, inQualifierData :  UInt32,  _, outDataSize :  UnsafeRawPointer?,  _, :  UnsafeMutablePointer< UInt32>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+

// AudioObjectHasProperty(inObjectID _, inAddress :  AudioObjectID,  _, :  UnsafePointer< AudioObjectPropertyAddress>) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+

// AudioObjectIsPropertySettable(inObjectID _, inAddress :  AudioObjectID,  _, outIsSettable :  UnsafePointer< AudioObjectPropertyAddress>,  _, :  UnsafeMutablePointer< DarwinBoolean>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+


// AudioObjectRemovePropertyListener(inObjectID _, inAddress :  AudioObjectID,  _, inListener :  UnsafePointer< AudioObjectPropertyAddress>,  _, inClientData :  AudioObjectPropertyListenerProc,  _, :  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+

// AudioObjectRemovePropertyListenerBlock(inObjectID _, inAddress :  AudioObjectID,  _, inDispatchQueue :  UnsafePointer< AudioObjectPropertyAddress>,  _, inListener :  dispatch_queue_t?,  _, :  @escaping  AudioObjectPropertyListenerBlock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.7+

// AudioObjectSetPropertyData(inObjectID _, inAddress :  AudioObjectID,  _, inQualifierDataSize :  UnsafePointer< AudioObjectPropertyAddress>,  _, inQualifierData :  UInt32,  _, inDataSize :  UnsafeRawPointer?,  _, inData :  UInt32,  _, :  UnsafeRawPointer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+


// AudioObjectShow(inObjectID _, :  AudioObjectID) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+

// AudioStreamAddPropertyListener(inStream AudioStreamID, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, inProc ,  AudioStreamPropertyListenerProc, inClientData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioStreamGetProperty(inStream AudioStreamID, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, ioPropertyDataSize ,  UInt32 *, outPropertyData ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// AudioStreamGetPropertyInfo(inStream AudioStreamID, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, outSize ,  UInt32 *, outWritable ,  Boolean *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioStreamRemovePropertyListener(inStream AudioStreamID, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, inProc ,  AudioStreamPropertyListenerProc, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// AudioStreamSetProperty(inStream AudioStreamID, inWhen ,  const  AudioTimeStamp *, inChannel ,  UInt32, inPropertyID ,  AudioDevicePropertyID, inPropertyDataSize ,  UInt32, inPropertyData ,  const  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


