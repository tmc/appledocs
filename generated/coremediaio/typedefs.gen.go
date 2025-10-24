// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio
import (
"unsafe"
)

// Type aliases and typedefs
// IOExtensionProperty - A structure that defines the properties that providers, devices, and streams support.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionProperty
// CMIOExtensionProperty is a string typedef
type IOExtensionProperty = string
// IOClassID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOClassID
// CMIOClassID has base type: UInt32
type IOClassID uintptr
// IOControlID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOControlID
// CMIOControlID has base type: CMIOObjectID
type IOControlID uintptr
// IODeviceGetSMPTETimeProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceGetSMPTETimeProc
// CMIODeviceGetSMPTETimeProc is a callback function
// C type: int (*)(void *, unsigned long long *, unsigned char *, unsigned int *)
type IODeviceGetSMPTETimeProc = func(unsafe.Pointer, uint64, uint8, uint32) int32
// IODeviceID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceID
// CMIODeviceID has base type: CMIOObjectID
type IODeviceID uintptr
// IODevicePropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODevicePropertyID
// CMIODevicePropertyID has base type: CMIOObjectPropertySelector
type IODevicePropertyID uintptr
// IODeviceStreamQueueAlteredProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStreamQueueAlteredProc
// CMIODeviceStreamQueueAlteredProc is a callback function
// C type: void (*)(unsigned int, void *, void *)
type IODeviceStreamQueueAlteredProc = func(uint32, unsafe.Pointer, unsafe.Pointer)
// IOHardwarePropertyID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePropertyID
// CMIOHardwarePropertyID has base type: CMIOObjectPropertySelector
type IOHardwarePropertyID uintptr
// IOObjectID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectID
// CMIOObjectID has base type: UInt32
type IOObjectID uintptr
// IOObjectPropertyElement type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyElement
// CMIOObjectPropertyElement has base type: UInt32
type IOObjectPropertyElement uintptr
// IOObjectPropertyListenerProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyListenerProc
// CMIOObjectPropertyListenerProc is a callback function
// C type: int (*)(unsigned int, unsigned int, const struct CMIOObjectPropertyAddress *, void *)
type IOObjectPropertyListenerProc = func(uint32, uint32, unsafe.Pointer, unsafe.Pointer) int32
// IOObjectPropertyScope type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyScope
// CMIOObjectPropertyScope has base type: UInt32
type IOObjectPropertyScope uintptr
// IOObjectPropertySelector type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertySelector
// CMIOObjectPropertySelector has base type: UInt32
type IOObjectPropertySelector uintptr
// IOStreamID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamID
// CMIOStreamID has base type: CMIOObjectID
type IOStreamID uintptr
// IOStreamScheduledOutputNotificationProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamScheduledOutputNotificationProc
// CMIOStreamScheduledOutputNotificationProc is a callback function
// C type: void (*)(unsigned long long, unsigned long long, void *)
type IOStreamScheduledOutputNotificationProc = func(uint64, uint64, unsafe.Pointer)

