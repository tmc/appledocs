// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

// Type aliases and typedefs
// AudioClassID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioClassID
// AudioClassID has base type: UInt32
type AudioClassID uintptr
// AudioDeviceID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceID
// AudioDeviceID has base type: AudioObjectID
type AudioDeviceID uintptr
// AudioDeviceIOProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProc
// AudioDeviceIOProc has base type: int (*)(unsigned int, const struct AudioTimeStamp *, const struct AudioBufferList *, const struct AudioTimeStamp *, struct AudioBufferList *, const struct AudioTimeStamp *, void *) __attribute__((nonblocking))
type AudioDeviceIOProc uintptr
// AudioDeviceIOProcID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProcID
// AudioDeviceIOProcID has base type: AudioDeviceIOProc
type AudioDeviceIOProcID uintptr
// AudioDevicePropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyID
// AudioDevicePropertyID has base type: AudioObjectPropertySelector
type AudioDevicePropertyID uintptr
// AudioDevicePropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyListenerProc
// AudioDevicePropertyListenerProc has base type: int (*)(unsigned int, unsigned int, unsigned char, unsigned int, void *)
type AudioDevicePropertyListenerProc uintptr
// AudioHardwarePropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyID
// AudioHardwarePropertyID has base type: AudioObjectPropertySelector
type AudioHardwarePropertyID uintptr
// AudioHardwarePropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyListenerProc
// AudioHardwarePropertyListenerProc has base type: int (*)(unsigned int, void *)
type AudioHardwarePropertyListenerProc uintptr
// AudioObjectID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectID
// AudioObjectID has base type: UInt32
type AudioObjectID uintptr
// AudioObjectPropertyElement type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyElement
// AudioObjectPropertyElement has base type: UInt32
type AudioObjectPropertyElement uintptr
// AudioObjectPropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyListenerProc
// AudioObjectPropertyListenerProc has base type: int (*)(unsigned int, unsigned int, const struct AudioObjectPropertyAddress *, void *)
type AudioObjectPropertyListenerProc uintptr
// AudioObjectPropertyScope type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyScope
// AudioObjectPropertyScope has base type: UInt32
type AudioObjectPropertyScope uintptr
// AudioObjectPropertySelector type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertySelector
// AudioObjectPropertySelector has base type: UInt32
type AudioObjectPropertySelector uintptr
// AudioStreamID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamID
// AudioStreamID has base type: AudioObjectID
type AudioStreamID uintptr
// AudioStreamPropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamPropertyListenerProc
// AudioStreamPropertyListenerProc has base type: int (*)(unsigned int, unsigned int, unsigned int, void *)
type AudioStreamPropertyListenerProc uintptr

