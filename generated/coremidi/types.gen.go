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
}/* debug [types.gen.go/struct]: MIDI2DeviceManufacturer */

// MIDI2DeviceRevisionLevel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceRevisionLevel
type MIDI2DeviceRevisionLevel struct {
	RevisionLevel unsafe.Pointer
}/* debug [types.gen.go/struct]: MIDI2DeviceRevisionLevel */

// MIDICIDeviceIdentification - A structure that describes a MIDI-CI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceIdentification
type MIDICIDeviceIdentification struct {
	Family uint8 // The group of familes to which the device belongs.
	Manufacturer uint8 // The MIDI System Exclusive (SysEx) ID of the device manufacturer.
	ModelNumber uint8 // The device model number.
	Reserved uint8 // A reserved field whose value is always zero.
	RevisionLevel uint8 // The revision number of the device model number.
}/* debug [types.gen.go/struct]: MIDICIDeviceIdentification */

// MIDICIProfileIDManufacturerSpecific
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDManufacturerSpecific
type MIDICIProfileIDManufacturerSpecific struct {
	Info1 MIDIUInteger7
	Info2 MIDIUInteger7
	SysExID1 MIDIUInteger7
	SysExID2 MIDIUInteger7
	SysExID3 MIDIUInteger7
}/* debug [types.gen.go/struct]: MIDICIProfileIDManufacturerSpecific */

// MIDICIProfileIDStandard
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDStandard
type MIDICIProfileIDStandard struct {
	ProfileBank MIDIUInteger7
	ProfileIDByte1 MIDIUInteger7
	ProfileLevel MIDIUInteger7
	ProfileNumber MIDIUInteger7
	ProfileVersion MIDIUInteger7
}/* debug [types.gen.go/struct]: MIDICIProfileIDStandard */

// MIDIControlTransform - A structure that describes the transformation of MIDI control change events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIControlTransform
type MIDIControlTransform struct {
	ControlNumber unsafe.Pointer // The control number to affect.
	ControlType MIDITransformControlType // The type of control specified by the control number.
	Param unsafe.Pointer // An argument to the transformation method.
	RemappedControlType MIDITransformControlType // The remapped control type.
	Transform MIDITransformType // The type of transformation to apply to the event values.
}/* debug [types.gen.go/struct]: MIDIControlTransform */

// MIDIDriverInterface - The interface to a MIDI driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverInterface
type MIDIDriverInterface struct {
	AddRef unsafe.Pointer
	Configure unsafe.Pointer // The system doesn’t currently use this method.
	EnableSource unsafe.Pointer // Tells the driver whether input from a particular source has listeners.
	FindDevices unsafe.Pointer // Finds the available devices.
	Flush unsafe.Pointer // Unschedules all pending output to the specified destination.
	Monitor unsafe.Pointer // Enables monitoring of MIDI packet lists by the specified driver.
	MonitorEvents unsafe.Pointer // Enables monitoring of MIDI event lists by the specified driver.
	QueryInterface unsafe.Pointer
	Release unsafe.Pointer
	Send unsafe.Pointer // Sends a MIDI packet list to the specified destination endpoints.
	SendPackets unsafe.Pointer // Sends a MIDI event list to the specified destination endpoints.
	Start unsafe.Pointer // Starts MIDI I/O.
	Stop unsafe.Pointer // Stops MIDI I/O.
}/* debug [types.gen.go/struct]: MIDIDriverInterface */

// MIDIEventList - A variable-length list of MIDI event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventList
type MIDIEventList struct {
	NumPackets unsafe.Pointer // The number of MIDI event packet structures in the list.
	Packet MIDIEventPacket // An array of variable-length MIDI event packet structures.
	Protocol MIDIProtocolID // The MIDI protocol variant of the events in the list.
}/* debug [types.gen.go/struct]: MIDIEventList */

// MIDIEventPacket - A series of simultaneous MIDI events in Universal MIDI Packets (UMP) format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventPacket
type MIDIEventPacket struct {
	TimeStamp MIDITimeStamp // The event packet timestamp.
	WordCount unsafe.Pointer // The number of valid MIDI 32-bit words in this event packet.
	Words unsafe.Pointer // A variable-length stream of native-endian 32-bit Universal MIDI Packets (UMP).
}/* debug [types.gen.go/struct]: MIDIEventPacket */

// MIDIIOErrorNotification - A general I/O error notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIIOErrorNotification
type MIDIIOErrorNotification struct {
	DriverDevice MIDIDeviceRef // The device with an I/O error.
	ErrorCode unsafe.Pointer // The error code of the generated error.
	MessageID MIDINotificationMessageID // The type of message.
	MessageSize unsafe.Pointer // The size of the message.
}/* debug [types.gen.go/struct]: MIDIIOErrorNotification */

// MIDIMessage_128 - A 128-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_128
type MIDIMessage_128 struct {
	Word0 unsafe.Pointer
	Word1 unsafe.Pointer
	Word2 unsafe.Pointer
	Word3 unsafe.Pointer
}/* debug [types.gen.go/struct]: MIDIMessage_128 */

// MIDIMessage_64 - A 64-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_64
type MIDIMessage_64 struct {
	Word0 unsafe.Pointer
	Word1 unsafe.Pointer
}/* debug [types.gen.go/struct]: MIDIMessage_64 */

// MIDIMessage_96 - A 96-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_96
type MIDIMessage_96 struct {
	Word0 unsafe.Pointer
	Word1 unsafe.Pointer
	Word2 unsafe.Pointer
}/* debug [types.gen.go/struct]: MIDIMessage_96 */

// MIDINotification - A message that describes a system state change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotification
type MIDINotification struct {
	MessageID MIDINotificationMessageID // An identifier that describes the type of state change.
	MessageSize unsafe.Pointer // The size of the message including its ID.
}/* debug [types.gen.go/struct]: MIDINotification */

// MIDIObjectAddRemoveNotification - A message that describes the addition or removal of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectAddRemoveNotification
type MIDIObjectAddRemoveNotification struct {
	Child MIDIObjectRef // The added or removed child object.
	ChildType MIDIObjectType // The child object type.
	MessageID MIDINotificationMessageID // The message type.
	MessageSize unsafe.Pointer // The message size.
	Parent MIDIObjectRef // The parent object of the added or removed child.
	ParentType MIDIObjectType // The parent object type.
}/* debug [types.gen.go/struct]: MIDIObjectAddRemoveNotification */

// MIDIObjectPropertyChangeNotification - A message that describes the change to an object property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectPropertyChangeNotification
type MIDIObjectPropertyChangeNotification struct {
	MessageID MIDINotificationMessageID // The message type.
	MessageSize unsafe.Pointer // The message size.
	Object MIDIObjectRef // The object whose property changed.
	ObjectType MIDIObjectType // The object type.
	PropertyName StringRef // The name of the modified property.
}/* debug [types.gen.go/struct]: MIDIObjectPropertyChangeNotification */

// MIDIPacket - A collection of simultaneous MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacket
type MIDIPacket struct {
	Data unsafe.Pointer // A variable-length stream of MIDI messages.
	Length unsafe.Pointer // The number of valid MIDI data bytes in this packet.
	TimeStamp MIDITimeStamp // The MIDI packet timestamp.
}/* debug [types.gen.go/struct]: MIDIPacket */

// MIDIPacketList - A list of MIDI events the system sends to or receives from an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacketList
type MIDIPacketList struct {
	NumPackets unsafe.Pointer // The number of MIDI packets in the list.
	Packet MIDIPacket // An open-ended array of variable-length MIDI packets.
}/* debug [types.gen.go/struct]: MIDIPacketList */

// MIDISysexSendRequest - A request to asynchronously send a single system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequest
type MIDISysexSendRequest struct {
	BytesToSend unsafe.Pointer // The number of bytes to send.
	Complete unsafe.Pointer // A Boolean value that indicates whether the transmission is complete.
	CompletionProc MIDICompletionProc // A function that the system calls after it sends all bytes for the request, or after the client marks the request as complete.
	CompletionRefCon unsafe.Pointer // Data to pass to the completion function.
	Data unsafe.Pointer // The request’s data.
	Destination MIDIEndpointRef // The endpoint to send the event to.
	Reserved unsafe.Pointer // A field that’s reserved for future use.
}/* debug [types.gen.go/struct]: MIDISysexSendRequest */

// MIDISysexSendRequestUMP - A request to asynchronously send a single universal MIDI packet (UMP) system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequestUMP
type MIDISysexSendRequestUMP struct {
	Complete unsafe.Pointer // A Boolean value that indicates whether the transmission is complete.
	CompletionProc MIDICompletionProcUMP // A function that the system calls after it sends all data for the request, or after the client marks the request as complete.
	CompletionRefCon unsafe.Pointer // Data to pass to the completion function.
	Destination MIDIEndpointRef // The endpoint to send the event to.
	Words unsafe.Pointer // A pointer to the event to send, which the system advances as it sends the data.
	WordsToSend unsafe.Pointer // A counter of the number of words to send.
}/* debug [types.gen.go/struct]: MIDISysexSendRequestUMP */

// MIDIThruConnectionEndpoint - A source or destination in a MIDI thru connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionEndpoint
type MIDIThruConnectionEndpoint struct {
	EndpointRef MIDIEndpointRef // The endpoint reference.
	UniqueID MIDIUniqueID // The connection’s unique identifier.
}/* debug [types.gen.go/struct]: MIDIThruConnectionEndpoint */

// MIDIThruConnectionParams - A set of MIDI routings and transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParams
type MIDIThruConnectionParams struct {
	ChannelMap unsafe.Pointer // A mapping of MIDI channels.
	ChannelPressure MIDITransform // The transformation of MIDI monophonic channel pressure events.
	Destinations MIDIThruConnectionEndpoint // All MIDI destinations for this connection.
	FilterOutAllControls unsafe.Pointer // A value that indicates whether to filter out MIDI continuous control messages.
	FilterOutBeatClock unsafe.Pointer // A value that indicates whether to filter out MIDI clock, play, stop, and resume messages.
	FilterOutMTC unsafe.Pointer // A value that indicates whether to filter out MIDI Time Code messages.
	FilterOutSysEx unsafe.Pointer // A value that indicates wheter to filter out system-exclusive messages.
	FilterOutTuneRequest unsafe.Pointer // A value that specifies whether to filter out MIDI tune request messages.
	HighNote unsafe.Pointer // The note value above which the system filters out notes.
	HighVelocity unsafe.Pointer // The velocity value above which the system filters out notes.
	KeyPressure MIDITransform // The transformation of polyphonic key pressure events.
	LowNote unsafe.Pointer // The note value below which the system filters out notes.
	LowVelocity unsafe.Pointer // The velocity value below which the system filters out notes.
	NoteNumber MIDITransform // The transformation of MIDI note numbers.
	NumControlTransforms unsafe.Pointer // The number of control transformations in the variable-length portion of the struct.
	NumDestinations unsafe.Pointer // The number of valid destinations.
	NumMaps unsafe.Pointer // The number of MIDI value maps in the variable-length portion of the struct.
	NumSources unsafe.Pointer // The number of valid sources.
	PitchBend MIDITransform // The transformation of a MIDI pitch bend event.
	ProgramChange MIDITransform // A transformation of a MIDI program change event.
	Reserved2 unsafe.Pointer // A reserved value that must be 0.
	Reserved3 unsafe.Pointer // A reserved value that must be 0.
	Sources MIDIThruConnectionEndpoint // All MIDI sources for this connection.
	Velocity MIDITransform // A note velocity transformation.
	Version unsafe.Pointer // The version number.
}/* debug [types.gen.go/struct]: MIDIThruConnectionParams */

// MIDITransform - The transformation of a single type of MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransform
type MIDITransform struct {
	Param unsafe.Pointer // An argument to the transformation method (see description of MIDITransformType).
	Transform MIDITransformType // The type of transformation to apply to the event values.
}/* debug [types.gen.go/struct]: MIDITransform */

// MIDIUniversalMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage
type MIDIUniversalMessage struct {
	ChannelVoice1 unsafe.Pointer // active when type is kMIDIMessageTypeChannelVoice1
	Channel unsafe.Pointer // MIDI channel 0-15
	ChannelPressure unsafe.Pointer // 7 bit channel pressure, active when status is kMIDICVStatusChannelPressure
	ControlChange unsafe.Pointer // active when status is kMIDICVStatusControlChange
	Data unsafe.Pointer // 7 bit value for control parameter
	Index unsafe.Pointer // 7 bit index of control parameter
	Note unsafe.Pointer // active when status is kMIDICVStatusNoteOff or kMIDICVStatusNoteOn
	Number unsafe.Pointer // 7 bit note number
	Velocity unsafe.Pointer // 7 bit note velocity
	PitchBend unsafe.Pointer // 7 bit pitch bend active when status is kMIDICVStatusPitchBend
	PolyPressure unsafe.Pointer // active when status is kMIDICVStatusPolyPressure
	NoteNumber unsafe.Pointer // 7 bit note number
	Pressure unsafe.Pointer // 7 bit poly pressure data
	Program unsafe.Pointer // 7 bit program nr, active when status is kMIDICVStatusProgramChange
	Reserved unsafe.Pointer
	Status MIDICVStatus // determines which variant is active
	ChannelVoice2 unsafe.Pointer // active when type is kMIDIMessageTypeChannelVoice2
	Controller unsafe.Pointer // active when status is any of kMIDICVStatusRegisteredControl, kMIDICVStatusAssignableControl, kMIDICVStatusRelRegisteredControl, or kMIDICVStatusRelAssignableControl
	Bank unsafe.Pointer // 7 bit bank
	Attribute unsafe.Pointer // attribute data
	AttributeType MIDINoteAttribute // attribute type
	PerNoteController unsafe.Pointer // active when status is kMIDICVStatusRegisteredPNC or kMIDICVStatusAssignablePNC
	PerNoteManagement unsafe.Pointer // active when status is kMIDICVStatusPerNoteMgmt
	Options MIDIPerNoteManagementOptions
	PerNotePitchBend unsafe.Pointer // active when status is kMIDICVStatusPerNotePitchBend
	Bend unsafe.Pointer // per note pitch bend value
	ProgramChange unsafe.Pointer // active when status is kMIDICVStatusProgramChange
	Data128 unsafe.Pointer // active when type is kMIDIMessageTypeData128
	MixedDataSet unsafe.Pointer // active when status is kMIDISysExStatusMixedDataSetHeader or kMIDISysExStatusMixedDataSetPayload
	MdsID unsafe.Pointer // mixed data set ID
	Sysex8 unsafe.Pointer // active when status any of kMIDISysExStatusComplete, kMIDISysExStatusStart, kMIDISysExStatusContinue, or kMIDISysExStatusEnd
	ByteCount unsafe.Pointer // byte count of data including stream ID (1-14 bytes)
	StreamID unsafe.Pointer
	Group unsafe.Pointer
	SysEx unsafe.Pointer // active when type is kMIDIMessageTypeSysEx
	System unsafe.Pointer // active when type is kMIDIMessageTypeSystem
	SongPositionPointer unsafe.Pointer // active when status is kMIDIStatusSongPosPointer
	SongSelect unsafe.Pointer // active when status is kMIDIStatusSongSelect
	TimeCode unsafe.Pointer // active when status is kMIDIStatusMTC
	Type MIDIMessageType
	Unknown unsafe.Pointer // active when type is unkown
	Words unsafe.Pointer // up to four 32 bit words
	Utility unsafe.Pointer // active when type is kMIDIMessageTypeUtility
	JitterReductionClock unsafe.Pointer // active when status is kMIDIUtilityStatusJitterReductionClock
	JitterReductionTimestamp unsafe.Pointer // active when status is kMIDIUtilityStatusJitterReductionTimestamp
}/* debug [types.gen.go/struct]: MIDIUniversalMessage */

// channelVoice1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/channelVoice1-3muv1
type channelVoice1 struct {
}/* debug [types.gen.go/struct]: channelVoice1 */

// channelVoice2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/channelVoice2-3muv2
type channelVoice2 struct {
}/* debug [types.gen.go/struct]: channelVoice2 */

// data128
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/data128-3jrad
type data128 struct {
}/* debug [types.gen.go/struct]: data128 */

// sysEx
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/sysEx-2jr6w
type sysEx struct {
}/* debug [types.gen.go/struct]: sysEx */

// system
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/system-6vxkw
type system struct {
}/* debug [types.gen.go/struct]: system */

// unknown
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/unknown-9rrub
type unknown struct {
}/* debug [types.gen.go/struct]: unknown */

// utility
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage/utility-1trwz
type utility struct {
}/* debug [types.gen.go/struct]: utility */

// MIDIValueMap - A custom lookup table to transform MIDI 7-bit values, as contained in note numbers, velocities, control values, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIValueMap
type MIDIValueMap struct {
	Value unsafe.Pointer // The array of unsigned 8-bit integers.
}/* debug [types.gen.go/struct]: MIDIValueMap */





