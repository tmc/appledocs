// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi
import (
"unsafe"
)

// Type aliases and typedefs
// MIDICIDeviceID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceID
// MIDICIDeviceID has base type: MIDIUInteger7
type MIDICIDeviceID uintptr
// MIDICIDeviceManagerDictionaryKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/DictionaryKey
// MIDICIDeviceManagerDictionaryKey has base type: NSString * const
type MIDICIDeviceManagerDictionaryKey uintptr
// MIDICIMUID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIMUID
// MIDICIMUID has base type: MIDIUInteger28
type MIDICIMUID uintptr
// MIDICIPropertyExchangeRequestID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeRequestID
// MIDICIPropertyExchangeRequestID has base type: UInt8
type MIDICIPropertyExchangeRequestID uintptr
// MIDIChannelNumber type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIChannelNumber
// MIDIChannelNumber has base type: MIDIUInteger4
type MIDIChannelNumber uintptr
// MIDIClientRef - An object that maintains per-client state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIClientRef
// MIDIClientRef has base type: MIDIObjectRef
type MIDIClientRef uintptr
// MIDIDeviceListRef - A list of MIDI devices.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListRef
// MIDIDeviceListRef has base type: MIDIObjectRef
type MIDIDeviceListRef uintptr
// MIDIDeviceRef - A MIDI device that contains entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceRef
// MIDIDeviceRef has base type: MIDIObjectRef
type MIDIDeviceRef uintptr
// MIDIDriverRef - A MIDI driver object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverRef
// MIDIDriverRef has base type: MIDIDriverInterface * *
type MIDIDriverRef uintptr
// MIDIEndpointRef - A MIDI source or destination an entity owns.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointRef
// MIDIEndpointRef has base type: MIDIObjectRef
type MIDIEndpointRef uintptr
// MIDIEntityRef - An entity that a device owns and that contains endpoints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityRef
// MIDIEntityRef has base type: MIDIObjectRef
type MIDIEntityRef uintptr
// MIDIEventVisitor type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventVisitor
// MIDIEventVisitor is a callback function
// C type: void (*)(void *, unsigned long long, struct MIDIUniversalMessage)
type MIDIEventVisitor = func(unsafe.Pointer, uint64, MIDIUniversalMessage)
// MIDIMessage_32 - A 32-bit MIDI message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_32
// MIDIMessage_32 has base type: UInt32
type MIDIMessage_32 uintptr
// MIDIObjectRef - The common base class for many of the framework’s objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectRef
// MIDIObjectRef has base type: UInt32
type MIDIObjectRef uintptr
// MIDIPortRef - A MIDI connection that a client maintains.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPortRef
// MIDIPortRef has base type: MIDIObjectRef
type MIDIPortRef uintptr
// MIDISetupRef - A type that represents the global state of the MIDI system, that contains lists of the devices and serial port owners.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRef
// MIDISetupRef has base type: MIDIObjectRef
type MIDISetupRef uintptr
// MIDIThruConnectionRef - An opaque reference to a play-through connection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionRef
// MIDIThruConnectionRef has base type: MIDIObjectRef
type MIDIThruConnectionRef uintptr
// MIDITimeStamp - The time on the host clock when the event occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITimeStamp
// MIDITimeStamp has base type: UInt64
type MIDITimeStamp uintptr
// MIDIUInteger14 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger14
// MIDIUInteger14 has base type: UInt16
type MIDIUInteger14 uintptr
// MIDIUInteger2 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger2
// MIDIUInteger2 has base type: UInt8
type MIDIUInteger2 uintptr
// MIDIUInteger28 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger28
// MIDIUInteger28 has base type: UInt32
type MIDIUInteger28 uintptr
// MIDIUInteger4 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger4
// MIDIUInteger4 has base type: UInt8
type MIDIUInteger4 uintptr
// MIDIUInteger7 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger7
// MIDIUInteger7 has base type: UInt8
type MIDIUInteger7 uintptr
// MIDIUMPEndpointManagerDictionaryKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/DictionaryKey
// MIDIUMPEndpointManagerDictionaryKey has base type: NSString * const
type MIDIUMPEndpointManagerDictionaryKey uintptr
// MIDIUMPFunctionBlockID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockID
// MIDIUMPFunctionBlockID has base type: MIDIUInteger7
type MIDIUMPFunctionBlockID uintptr
// MIDIUMPGroupNumber type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPGroupNumber
// MIDIUMPGroupNumber has base type: MIDIUInteger4
type MIDIUMPGroupNumber uintptr

