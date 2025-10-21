// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

// Enum types and constants
// MIDICICategoryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions
type MIDICICategoryOptions uint

// MIDICIDeviceType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType
type MIDICIDeviceType uint

// MIDICIManagementMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType
type MIDICIManagementMessageType uint

// MIDICIProcessInquiryMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType
type MIDICIProcessInquiryMessageType uint

// MIDICIProfileMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType
type MIDICIProfileMessageType uint

// MIDICIProfileType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType
type MIDICIProfileType uint

// MIDICIPropertyExchangeMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType
type MIDICIPropertyExchangeMessageType uint

// MIDICVStatus - MIDI status types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus
type MIDICVStatus uint

// MIDIMessageType - Supported MIDI message types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType
type MIDIMessageType uint

// MIDINetworkConnectionPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy
type MIDINetworkConnectionPolicy uint

// MIDINoteAttribute enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute
type MIDINoteAttribute uint

// MIDIObjectType - The MIDI object types that the system supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType
type MIDIObjectType uint

const (
// kMIDIObjectType_Destination - A MIDI destination.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/destination
kMIDIObjectType_Destination MIDIObjectType = 0
// kMIDIObjectType_Device - A MIDI device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/device
kMIDIObjectType_Device MIDIObjectType = 0
// kMIDIObjectType_Entity - A MIDI entity.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/entity
kMIDIObjectType_Entity MIDIObjectType = 0
// kMIDIObjectType_ExternalDestination - An external destination.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/externalDestination
kMIDIObjectType_ExternalDestination MIDIObjectType = 0
// kMIDIObjectType_ExternalDevice - An external device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/externalDevice
kMIDIObjectType_ExternalDevice MIDIObjectType = 0
// kMIDIObjectType_ExternalEntity - An external entity.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/externalEntity
kMIDIObjectType_ExternalEntity MIDIObjectType = 0
// kMIDIObjectType_ExternalSource - An external source.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/externalSource
kMIDIObjectType_ExternalSource MIDIObjectType = 0
// kMIDIObjectType_Other - A MIDI object with an undefined type.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/other
kMIDIObjectType_Other MIDIObjectType = 0
// kMIDIObjectType_Source - A MIDI source.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType/source
kMIDIObjectType_Source MIDIObjectType = 0
)

// MIDIPerNoteManagementOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPerNoteManagementOptions
type MIDIPerNoteManagementOptions uint

// MIDIProgramChangeOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProgramChangeOptions
type MIDIProgramChangeOptions uint

// MIDIProtocolID - Specifies a MIDI protocol variant.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProtocolID
type MIDIProtocolID uint

// MIDISysExStatus - MIDI System Exclusive (SysEx) types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus
type MIDISysExStatus uint

// MIDISystemStatus - MIDI System status types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus
type MIDISystemStatus uint

// MIDITransformControlType - A set of values that indicate how to interpret control numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType
type MIDITransformControlType uint

// MIDITransformType - Values that specify the type of MIDI transformation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType
type MIDITransformType uint

// MIDIUMPCIObjectBackingType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType
type MIDIUMPCIObjectBackingType uint

// MIDIUMPFunctionBlockDirection enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection
type MIDIUMPFunctionBlockDirection uint

// MIDIUMPFunctionBlockMIDI1Info enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info
type MIDIUMPFunctionBlockMIDI1Info uint

// MIDIUMPFunctionBlockUIHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint
type MIDIUMPFunctionBlockUIHint uint

// MIDIUMPProtocolOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPProtocolOptions
type MIDIUMPProtocolOptions uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPProtocolOptions/midi1
kMIDIUMPProtocolOptionsMIDI1 MIDIUMPProtocolOptions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPProtocolOptions/midi2
kMIDIUMPProtocolOptionsMIDI2 MIDIUMPProtocolOptions = 0
)

// MIDIUtilityStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus
type MIDIUtilityStatus uint

// UMPStreamMessageFormat enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat
type UMPStreamMessageFormat uint

// UMPStreamMessageStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus
type UMPStreamMessageStatus uint


