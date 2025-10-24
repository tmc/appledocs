// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

/* debug [enums.gen.go]: Generating 28 enums for CoreMIDI */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MIDICICategoryOptions (4 cases) */
// MIDICICategoryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions
type MIDICICategoryOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions/processInquirySupported
	kMIDICICategoryOptionsProcessInquirySupported MIDICICategoryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions/profileConfigurationSupported
	kMIDICICategoryOptionsProfileConfigurationSupported MIDICICategoryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions/propertyExchangeSupported
	kMIDICICategoryOptionsPropertyExchangeSupported MIDICICategoryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions/protocolNegotiation
	kMIDICICategoryOptionsProtocolNegotiation MIDICICategoryOptions = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIDeviceType (4 cases) */
// MIDICIDeviceType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType
type MIDICIDeviceType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType/legacyMIDI1
	kMIDICIDeviceTypeLegacyMIDI1 MIDICIDeviceType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType/unknown
	kMIDICIDeviceTypeUnknown MIDICIDeviceType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType/usbMIDI
	kMIDICIDeviceTypeUSBMIDI MIDICIDeviceType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType/virtual
	kMIDICIDeviceTypeVirtual MIDICIDeviceType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIManagementMessageType (7 cases) */
// MIDICIManagementMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType
type MIDICIManagementMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/discovery
	kMIDICIManagementMessageTypeDiscovery MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/inquiryEndpointInformation
	kMIDICIManagementMessageTypeInquiryEndpointInformation MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/invalidateMUID
	kMIDICIManagementMessageTypeInvalidateMUID MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/midiCIACK
	kMIDICIManagementMessageTypeMIDICIACK MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/midiNAK
	kMIDICIManagementMessageTypeMIDICINAK MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/replyToDiscovery
	kMIDICIManagementMessageTypeReplyToDiscovery MIDICIManagementMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType/replyToEndpointInformation
	kMIDICIManagementMessageTypeReplyToEndpointInformation MIDICIManagementMessageType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIProcessInquiryMessageType (5 cases) */
// MIDICIProcessInquiryMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType
type MIDICIProcessInquiryMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType/endOfMIDIMessageReport
	kMIDICIProcessInquiryMessageTypeEndOfMIDIMessageReport MIDICIProcessInquiryMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType/inquiryMIDIMessageReport
	kMIDICIProcessInquiryMessageTypeInquiryMIDIMessageReport MIDICIProcessInquiryMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType/inquiryProcessInquiryCapabilities
	kMIDICIProcessInquiryMessageTypeInquiryProcessInquiryCapabilities MIDICIProcessInquiryMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType/replyToMIDIMessageReport
	kMIDICIProcessInquiryMessageTypeReplyToMIDIMessageReport MIDICIProcessInquiryMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType/replyToProcessInquiryCapabilities
	kMIDICIProcessInquiryMessageTypeReplyToProcessInquiryCapabilities MIDICIProcessInquiryMessageType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIProfileMessageType (11 cases) */
// MIDICIProfileMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType
type MIDICIProfileMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/detailsInquiry
	kMIDICIProfileMessageTypeDetailsInquiry MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileAdded
	kMIDICIProfileMessageTypeProfileAdded MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileDisabledReport
	kMIDICIProfileMessageTypeProfileDisabledReport MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileEnabledReport
	kMIDICIProfileMessageTypeProfileEnabledReport MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileInquiry
	kMIDICIProfileMessageTypeProfileInquiry MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileRemoved
	kMIDICIProfileMessageTypeProfileRemoved MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/profileSpecificData
	kMIDICIProfileMessageTypeProfileSpecificData MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/replyToDetailsInquiry
	kMIDICIProfileMessageTypeReplyToDetailsInquiry MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/replyToProfileInquiry
	kMIDICIProfileMessageTypeReplyToProfileInquiry MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/setProfileOff
	kMIDICIProfileMessageTypeSetProfileOff MIDICIProfileMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType/setProfileOn
	kMIDICIProfileMessageTypeSetProfileOn MIDICIProfileMessageType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIProfileType (4 cases) */
// MIDICIProfileType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType
type MIDICIProfileType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType/functionBlock
	kMIDICIProfileTypeFunctionBlock MIDICIProfileType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType/group
	kMIDICIProfileTypeGroup MIDICIProfileType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType/multichannel
	kMIDICIProfileTypeMultichannel MIDICIProfileType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType/singleChannel
	kMIDICIProfileTypeSingleChannel MIDICIProfileType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICIPropertyExchangeMessageType (11 cases) */
// MIDICIPropertyExchangeMessageType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType
type MIDICIPropertyExchangeMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/inquiryGetPropertyData
	kMIDICIPropertyExchangeMessageTypeInquiryGetPropertyData MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/inquiryHasPropertyData_Reserved
	kMIDICIPropertyExchangeMessageTypeInquiryHasPropertyData_Reserved MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/inquiryPropertyExchangeCapabilities
	kMIDICIPropertyExchangeMessageTypeInquiryPropertyExchangeCapabilities MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/inquiryReplyToHasPropertyData_Reserved
	kMIDICIPropertyExchangeMessageTypeInquiryReplyToHasPropertyData_Reserved MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/inquirySetPropertyData
	kMIDICIPropertyExchangeMessageTypeInquirySetPropertyData MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/notify
	kMIDICIPropertyExchangeMessageTypeNotify MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/replyToGetProperty
	kMIDICIPropertyExchangeMessageTypeReplyToGetProperty MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/replyToPropertyExchangeCapabilities
	kMIDICIPropertyExchangeMessageTypeReplyToPropertyExchangeCapabilities MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/replyToSetPropertyData
	kMIDICIPropertyExchangeMessageTypeReplyToSetPropertyData MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/replyToSubscription
	kMIDICIPropertyExchangeMessageTypeReplyToSubscription MIDICIPropertyExchangeMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType/subscription
	kMIDICIPropertyExchangeMessageTypeSubscription MIDICIPropertyExchangeMessageType = 0
)

/* debug [enums.gen.go]: Processing enum MIDICVStatus (15 cases) */
// MIDICVStatus - MIDI status types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus
type MIDICVStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/assignableControl
	kMIDICVStatusAssignableControl MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/assignablePNC
	kMIDICVStatusAssignablePNC MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/channelPressure
	kMIDICVStatusChannelPressure MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/controlChange
	kMIDICVStatusControlChange MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/noteOff
	kMIDICVStatusNoteOff MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/noteOn
	kMIDICVStatusNoteOn MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/perNoteMgmt
	kMIDICVStatusPerNoteMgmt MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/perNotePitchBend
	kMIDICVStatusPerNotePitchBend MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/pitchBend
	kMIDICVStatusPitchBend MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/polyPressure
	kMIDICVStatusPolyPressure MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/programChange
	kMIDICVStatusProgramChange MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/registeredControl
	kMIDICVStatusRegisteredControl MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/registeredPNC
	kMIDICVStatusRegisteredPNC MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/relAssignableControl
	kMIDICVStatusRelAssignableControl MIDICVStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus/relRegisteredControl
	kMIDICVStatusRelRegisteredControl MIDICVStatus = 0
)

/* debug [enums.gen.go]: Processing enum MIDIMessageType (10 cases) */
// MIDIMessageType - Supported MIDI message types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType
type MIDIMessageType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/channelVoice1
	kMIDIMessageTypeChannelVoice1 MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/channelVoice2
	kMIDIMessageTypeChannelVoice2 MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/data128
	kMIDIMessageTypeData128 MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/flexData
	kMIDIMessageTypeFlexData MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/invalid
	kMIDIMessageTypeInvalid MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/stream
	kMIDIMessageTypeStream MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/sysEx
	kMIDIMessageTypeSysEx MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/system
	kMIDIMessageTypeSystem MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/unknownF
	kMIDIMessageTypeUnknownF MIDIMessageType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType/utility
	kMIDIMessageTypeUtility MIDIMessageType = 0
)

/* debug [enums.gen.go]: Processing enum MIDINetworkConnectionPolicy (3 cases) */
// MIDINetworkConnectionPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy
type MIDINetworkConnectionPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy/anyone
	MIDINetworkConnectionPolicy_Anyone MIDINetworkConnectionPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy/hostsInContactList
	MIDINetworkConnectionPolicy_HostsInContactList MIDINetworkConnectionPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy/noOne
	MIDINetworkConnectionPolicy_NoOne MIDINetworkConnectionPolicy = 0
)

/* debug [enums.gen.go]: Processing enum MIDINoteAttribute (4 cases) */
// MIDINoteAttribute enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute
type MIDINoteAttribute uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute/manufacturerSpecific
	kMIDINoteAttributeManufacturerSpecific MIDINoteAttribute = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute/none
	kMIDINoteAttributeNone MIDINoteAttribute = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute/pitch
	kMIDINoteAttributePitch MIDINoteAttribute = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute/profileSpecific
	kMIDINoteAttributeProfileSpecific MIDINoteAttribute = 0
)

/* debug [enums.gen.go]: Processing enum MIDINotificationMessageID (8 cases) */
// MIDINotificationMessageID - The types of state changes the system supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID
type MIDINotificationMessageID uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgInternalStart
	kMIDIMsgInternalStart MIDINotificationMessageID = 0
	// kMIDIMsgIOError - A driver I/O error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgIOError
	kMIDIMsgIOError MIDINotificationMessageID = 0
	// kMIDIMsgObjectAdded - The system added a device, entity, or endpoint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgObjectAdded
	kMIDIMsgObjectAdded MIDINotificationMessageID = 0
	// kMIDIMsgObjectRemoved - The system removed a device, entity, or endpoint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgObjectRemoved
	kMIDIMsgObjectRemoved MIDINotificationMessageID = 0
	// kMIDIMsgPropertyChanged - An object’s property value changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgPropertyChanged
	kMIDIMsgPropertyChanged MIDINotificationMessageID = 0
	// kMIDIMsgSerialPortOwnerChanged - The system changed a serial port owner.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgSerialPortOwnerChanged
	kMIDIMsgSerialPortOwnerChanged MIDINotificationMessageID = 0
	// kMIDIMsgSetupChanged - Some aspect of the current MIDI setup changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgSetupChanged
	kMIDIMsgSetupChanged MIDINotificationMessageID = 0
	// kMIDIMsgThruConnectionsChanged - The system created or disposed of a persistent MIDI Thru connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID/msgThruConnectionsChanged
	kMIDIMsgThruConnectionsChanged MIDINotificationMessageID = 0
)

/* debug [enums.gen.go]: Processing enum MIDIObjectType (9 cases) */
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

/* debug [enums.gen.go]: Processing enum MIDIPerNoteManagementOptions (2 cases) */
// MIDIPerNoteManagementOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPerNoteManagementOptions
type MIDIPerNoteManagementOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPerNoteManagementOptions/detach
	kMIDIPerNoteManagementDetach MIDIPerNoteManagementOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPerNoteManagementOptions/reset
	kMIDIPerNoteManagementReset MIDIPerNoteManagementOptions = 0
)

/* debug [enums.gen.go]: Processing enum MIDIProgramChangeOptions (1 cases) */
// MIDIProgramChangeOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProgramChangeOptions
type MIDIProgramChangeOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProgramChangeOptions/bankValid
	kMIDIProgramChangeBankValid MIDIProgramChangeOptions = 0
)

/* debug [enums.gen.go]: Processing enum MIDIProtocolID (2 cases) */
// MIDIProtocolID - Specifies a MIDI protocol variant.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProtocolID
type MIDIProtocolID uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProtocolID/kMIDIProtocol_1_0
	kMIDIProtocol_1_0 MIDIProtocolID = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIProtocolID/kMIDIProtocol_2_0
	kMIDIProtocol_2_0 MIDIProtocolID = 0
)

/* debug [enums.gen.go]: Processing enum MIDISysExStatus (6 cases) */
// MIDISysExStatus - MIDI System Exclusive (SysEx) types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus
type MIDISysExStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/complete
	kMIDISysExStatusComplete MIDISysExStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/continue
	kMIDISysExStatusContinue MIDISysExStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/end
	kMIDISysExStatusEnd MIDISysExStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/mixedDataSetHeader
	kMIDISysExStatusMixedDataSetHeader MIDISysExStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/mixedDataSetPayload
	kMIDISysExStatusMixedDataSetPayload MIDISysExStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus/start
	kMIDISysExStatusStart MIDISysExStatus = 0
)

/* debug [enums.gen.go]: Processing enum MIDISystemStatus (13 cases) */
// MIDISystemStatus - MIDI System status types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus
type MIDISystemStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusActiveSending
	kMIDIStatusActiveSending MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusActiveSensing
	kMIDIStatusActiveSensing MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusContinue
	kMIDIStatusContinue MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusEndOfExclusive
	kMIDIStatusEndOfExclusive MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusMTC
	kMIDIStatusMTC MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusSongPosPointer
	kMIDIStatusSongPosPointer MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusSongSelect
	kMIDIStatusSongSelect MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusStart
	kMIDIStatusStart MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusStartOfExclusive
	kMIDIStatusStartOfExclusive MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusStop
	kMIDIStatusStop MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusSystemReset
	kMIDIStatusSystemReset MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusTimingClock
	kMIDIStatusTimingClock MIDISystemStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus/statusTuneRequest
	kMIDIStatusTuneRequest MIDISystemStatus = 0
)

/* debug [enums.gen.go]: Processing enum MIDITransformControlType (6 cases) */
// MIDITransformControlType - A set of values that indicate how to interpret control numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType
type MIDITransformControlType uint

const (
	// kMIDIControlType_14Bit - A 14-bit control type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_14Bit
	kMIDIControlType_14Bit MIDITransformControlType = 0
	// kMIDIControlType_14BitNRPN - A 14-bit Nonregistered Parameter Number (RPN).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_14BitNRPN
	kMIDIControlType_14BitNRPN MIDITransformControlType = 0
	// kMIDIControlType_14BitRPN - A 14-bit Registered Parameter Number (RPN).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_14BitRPN
	kMIDIControlType_14BitRPN MIDITransformControlType = 0
	// kMIDIControlType_7Bit - A 7-bit control type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_7Bit
	kMIDIControlType_7Bit MIDITransformControlType = 0
	// kMIDIControlType_7BitNRPN - A 7-bit Nonregistered Parameter Number (RPN).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_7BitNRPN
	kMIDIControlType_7BitNRPN MIDITransformControlType = 0
	// kMIDIControlType_7BitRPN - A 7-bit Registered Parameter Number (RPN).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType/controlType_7BitRPN
	kMIDIControlType_7BitRPN MIDITransformControlType = 0
)

/* debug [enums.gen.go]: Processing enum MIDITransformType (8 cases) */
// MIDITransformType - Values that specify the type of MIDI transformation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType
type MIDITransformType uint

const (
	// kMIDITransform_Add - A transform that adds a parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/add
	kMIDITransform_Add MIDITransformType = 0
	// kMIDITransform_FilterOut - A transformation that filters out an event type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/filterOut
	kMIDITransform_FilterOut MIDITransformType = 0
	// kMIDITransform_MapControl - A transformation that changes a specified control number to a supplied parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/mapControl
	kMIDITransform_MapControl MIDITransformType = 0
	// kMIDITransform_MapValue - A transform that maps one value to another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/mapValue
	kMIDITransform_MapValue MIDITransformType = 0
	// kMIDITransform_MaxValue - A transform that sets the maximum value to the specified parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/maxValue
	kMIDITransform_MaxValue MIDITransformType = 0
	// kMIDITransform_MinValue - A transform that sets the minimum value to the specified parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/minValue
	kMIDITransform_MinValue MIDITransformType = 0
	// kMIDITransform_None - No transformation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/none
	kMIDITransform_None MIDITransformType = 0
	// kMIDITransform_Scale - A transform that multiplies by the specified parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType/scale
	kMIDITransform_Scale MIDITransformType = 0
)

/* debug [enums.gen.go]: Processing enum MIDIUMPCIObjectBackingType (4 cases) */
// MIDIUMPCIObjectBackingType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType
type MIDIUMPCIObjectBackingType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType/driverDevice
	kMIDIUMPCIObjectBackingTypeDriverDevice MIDIUMPCIObjectBackingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType/unknown
	kMIDIUMPCIObjectBackingTypeUnknown MIDIUMPCIObjectBackingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType/usbMIDI
	kMIDIUMPCIObjectBackingTypeUSBMIDI MIDIUMPCIObjectBackingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType/virtual
	kMIDIUMPCIObjectBackingTypeVirtual MIDIUMPCIObjectBackingType = 0
)

/* debug [enums.gen.go]: Processing enum MIDIUMPFunctionBlockDirection (4 cases) */
// MIDIUMPFunctionBlockDirection enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection
type MIDIUMPFunctionBlockDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection/bidirectional
	kMIDIUMPFunctionBlockDirectionBidirectional MIDIUMPFunctionBlockDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection/input
	kMIDIUMPFunctionBlockDirectionInput MIDIUMPFunctionBlockDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection/output
	kMIDIUMPFunctionBlockDirectionOutput MIDIUMPFunctionBlockDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection/unknown
	kMIDIUMPFunctionBlockDirectionUnknown MIDIUMPFunctionBlockDirection = 0
)

/* debug [enums.gen.go]: Processing enum MIDIUMPFunctionBlockMIDI1Info (3 cases) */
// MIDIUMPFunctionBlockMIDI1Info enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info
type MIDIUMPFunctionBlockMIDI1Info uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info/notMIDI1
	kMIDIUMPFunctionBlockMIDI1InfoNotMIDI1 MIDIUMPFunctionBlockMIDI1Info = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info/restrictedBandwidth
	kMIDIUMPFunctionBlockMIDI1InfoRestrictedBandwidth MIDIUMPFunctionBlockMIDI1Info = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info/unrestrictedBandwidth
	kMIDIUMPFunctionBlockMIDI1InfoUnrestrictedBandwidth MIDIUMPFunctionBlockMIDI1Info = 0
)

/* debug [enums.gen.go]: Processing enum MIDIUMPFunctionBlockUIHint (4 cases) */
// MIDIUMPFunctionBlockUIHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint
type MIDIUMPFunctionBlockUIHint uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint/receiver
	kMIDIUMPFunctionBlockUIHintReceiver MIDIUMPFunctionBlockUIHint = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint/sender
	kMIDIUMPFunctionBlockUIHintSender MIDIUMPFunctionBlockUIHint = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint/senderReceiver
	kMIDIUMPFunctionBlockUIHintSenderReceiver MIDIUMPFunctionBlockUIHint = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint/unknown
	kMIDIUMPFunctionBlockUIHintUnknown MIDIUMPFunctionBlockUIHint = 0
)

/* debug [enums.gen.go]: Processing enum MIDIUMPProtocolOptions (2 cases) */
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

/* debug [enums.gen.go]: Processing enum MIDIUtilityStatus (5 cases) */
// MIDIUtilityStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus
type MIDIUtilityStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus/deltaClockstampTicksPerQuarterNote
	kMIDIUtilityStatusDeltaClockstampTicksPerQuarterNote MIDIUtilityStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus/jitterReductionClock
	kMIDIUtilityStatusJitterReductionClock MIDIUtilityStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus/jitterReductionTimestamp
	kMIDIUtilityStatusJitterReductionTimestamp MIDIUtilityStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus/NOOP
	kMIDIUtilityStatusNOOP MIDIUtilityStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus/ticksSinceLastEvent
	kMIDIUtilityStatusTicksSinceLastEvent MIDIUtilityStatus = 0
)

/* debug [enums.gen.go]: Processing enum UMPStreamMessageFormat (4 cases) */
// UMPStreamMessageFormat enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat
type UMPStreamMessageFormat uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat/complete
	kUMPStreamMessageFormatComplete UMPStreamMessageFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat/continuing
	kUMPStreamMessageFormatContinuing UMPStreamMessageFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat/end
	kUMPStreamMessageFormatEnd UMPStreamMessageFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat/start
	kUMPStreamMessageFormatStart UMPStreamMessageFormat = 0
)

/* debug [enums.gen.go]: Processing enum UMPStreamMessageStatus (12 cases) */
// UMPStreamMessageStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus
type UMPStreamMessageStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/deviceIdentityNotification
	kUMPStreamMessageStatusDeviceIdentityNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/endOfClip
	kUMPStreamMessageStatusEndOfClip UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/endpointDiscovery
	kUMPStreamMessageStatusEndpointDiscovery UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/endpointInfoNotification
	kUMPStreamMessageStatusEndpointInfoNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/endpointNameNotification
	kUMPStreamMessageStatusEndpointNameNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/functionBlockDiscovery
	kUMPStreamMessageStatusFunctionBlockDiscovery UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/functionBlockInfoNotification
	kUMPStreamMessageStatusFunctionBlockInfoNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/functionBlockNameNotification
	kUMPStreamMessageStatusFunctionBlockNameNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/productInstanceIDNotification
	kUMPStreamMessageStatusProductInstanceIDNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/startOfClip
	kUMPStreamMessageStatusStartOfClip UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/streamConfigurationNotification
	kUMPStreamMessageStatusStreamConfigurationNotification UMPStreamMessageStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus/streamConfigurationRequest
	kUMPStreamMessageStatusStreamConfigurationRequest UMPStreamMessageStatus = 0
)


