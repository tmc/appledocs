// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi
import (
	"unsafe"
)


// C struct types
// MIDI2DeviceManufacturer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceManufacturer
type MIDI2DeviceManufacturer struct {
	SysExIDByte unsafe.Pointer
}// MIDI2DeviceRevisionLevel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceRevisionLevel
type MIDI2DeviceRevisionLevel struct {
}// MIDICIProfileIDManufacturerSpecific
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDManufacturerSpecific
type MIDICIProfileIDManufacturerSpecific struct {
}// MIDICIProfileIDStandard
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDStandard
type MIDICIProfileIDStandard struct {
}// MIDIControlTransform - A structure that describes the transformation of MIDI control change events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIControlTransform
type MIDIControlTransform struct {
}// MIDIDriverInterface - The interface to a MIDI driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverInterface
type MIDIDriverInterface struct {
}// MIDIEventList - A variable-length list of MIDI event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventList
type MIDIEventList struct {
}// MIDIEventPacket - A series of simultaneous MIDI events in Universal MIDI Packets (UMP) format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventPacket
type MIDIEventPacket struct {
}// MIDIIOErrorNotification - A general I/O error notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIIOErrorNotification
type MIDIIOErrorNotification struct {
}// MIDIMessage_128 - A 128-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_128
type MIDIMessage_128 struct {
}// MIDIMessage_64 - A 64-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_64
type MIDIMessage_64 struct {
}// MIDIMessage_96 - A 96-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_96
type MIDIMessage_96 struct {
}// MIDIPacket - A collection of simultaneous MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacket
type MIDIPacket struct {
}// MIDIPacketList - A list of MIDI events the system sends to or receives from an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacketList
type MIDIPacketList struct {
}// MIDISysexSendRequest - A request to asynchronously send a single system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequest
type MIDISysexSendRequest struct {
}// MIDISysexSendRequestUMP - A request to asynchronously send a single universal MIDI packet (UMP) system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequestUMP
type MIDISysexSendRequestUMP struct {
}// MIDIThruConnectionEndpoint - A source or destination in a MIDI thru connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionEndpoint
type MIDIThruConnectionEndpoint struct {
}// MIDIThruConnectionParams - A set of MIDI routings and transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParams
type MIDIThruConnectionParams struct {
}// MIDITransform - The transformation of a single type of MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransform
type MIDITransform struct {
}// MIDIUniversalMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage
type MIDIUniversalMessage struct {
}// MIDIValueMap - A custom lookup table to transform MIDI 7-bit values, as contained in note numbers, velocities, control values, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIValueMap
type MIDIValueMap struct {
}



