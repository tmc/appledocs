// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth
import (
	"unsafe"
)


// C struct types
// BluetoothAFHHostChannelClassification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothAFHHostChannelClassification
type BluetoothAFHHostChannelClassification struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothAFHHostChannelClassification */

// BluetoothAFHResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothAFHResults
type BluetoothAFHResults struct {
	AfhMap uint8
	Handle BluetoothConnectionHandle
	Mode BluetoothAFHMode
}/* debug [types.gen.go/struct]: BluetoothAFHResults */

// BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothDeviceAddress
type BluetoothDeviceAddress struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothDeviceAddress */

// BluetoothEnhancedSynchronousConnectionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothEnhancedSynchronousConnectionInfo
type BluetoothEnhancedSynchronousConnectionInfo struct {
	InputBandwidth BluetoothHCIInputBandwidth
	InputCodedDataSize BluetoothHCIInputCodedDataSize
	InputCodingFormat BluetoothHCIInputCodingFormat
	InputDataPath BluetoothHCIInputDataPath
	InputPCMDataFormat BluetoothHCIInputPCMDataFormat
	InputPCMSampelPayloadMSBPosition BluetoothHCIInputPCMSamplePayloadMSBPosition
	InputTransportUnitSize BluetoothHCIInputTransportUnitSize
	MaxLatency BluetoothHCIMaxLatency
	OutputBandwidth BluetoothHCIOutputBandwidth
	OutputCodedDataSize BluetoothHCIOutputCodedDataSize
	OutputCodingFormat BluetoothHCIOutputCodingFormat
	OutputDataPath BluetoothHCIOutputDataPath
	OutputPCMDataFormat BluetoothHCIOutputPCMDataFormat
	OutputPCMSampelPayloadMSBPosition BluetoothHCIOutputPCMSamplePayloadMSBPosition
	OutputTransportUnitSize BluetoothHCIOutputTransportUnitSize
	PacketType BluetoothPacketType
	ReceiveBandWidth BluetoothHCIReceiveBandwidth
	ReceiveCodecFrameSize BluetoothHCIReceiveCodecFrameSize
	ReceiveCodingFormat BluetoothHCIReceiveCodingFormat
	RetransmissionEffort BluetoothHCIRetransmissionEffort
	TransmitBandWidth BluetoothHCITransmitBandwidth
	TransmitCodecFrameSize BluetoothHCITransmitCodecFrameSize
	TransmitCodingFormat BluetoothHCITransmitCodingFormat
	VoiceSetting BluetoothHCIVoiceSetting
}/* debug [types.gen.go/struct]: BluetoothEnhancedSynchronousConnectionInfo */

// BluetoothEventFilterCondition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothEventFilterCondition
type BluetoothEventFilterCondition struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothEventFilterCondition */

// BluetoothHCIAcceptSynchronousConnectionRequestParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAcceptSynchronousConnectionRequestParams
type BluetoothHCIAcceptSynchronousConnectionRequestParams struct {
	ContentFormat uint16
	MaxLatency uint16
	PacketType uint16
	ReceiveBandwidth uint32
	RetransmissionEffort uint8
	TransmitBandwidth uint32
}/* debug [types.gen.go/struct]: BluetoothHCIAcceptSynchronousConnectionRequestParams */

// BluetoothHCIAutomaticFlushTimeoutInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAutomaticFlushTimeoutInfo
type BluetoothHCIAutomaticFlushTimeoutInfo struct {
	Handle BluetoothConnectionHandle
	Timeout BluetoothHCIAutomaticFlushTimeout
}/* debug [types.gen.go/struct]: BluetoothHCIAutomaticFlushTimeoutInfo */

// BluetoothHCIBufferSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIBufferSize
type BluetoothHCIBufferSize struct {
	ACLDataPacketLength uint16
	SCODataPacketLength uint8
	TotalNumACLDataPackets uint16
	TotalNumSCODataPackets uint16
}/* debug [types.gen.go/struct]: BluetoothHCIBufferSize */

// BluetoothHCICurrentInquiryAccessCodes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCICurrentInquiryAccessCodes
type BluetoothHCICurrentInquiryAccessCodes struct {
	Codes BluetoothHCIInquiryAccessCode
	Count BluetoothHCIInquiryAccessCodeCount
}/* debug [types.gen.go/struct]: BluetoothHCICurrentInquiryAccessCodes */

// BluetoothHCICurrentInquiryAccessCodesForWrite
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCICurrentInquiryAccessCodesForWrite
type BluetoothHCICurrentInquiryAccessCodesForWrite struct {
	Codes uint8
	Count BluetoothHCIInquiryAccessCodeCount
}/* debug [types.gen.go/struct]: BluetoothHCICurrentInquiryAccessCodesForWrite */

// BluetoothHCIEncryptionKeySizeInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEncryptionKeySizeInfo
type BluetoothHCIEncryptionKeySizeInfo struct {
	Handle BluetoothConnectionHandle
	KeySize BluetoothHCIEncryptionKeySize
}/* debug [types.gen.go/struct]: BluetoothHCIEncryptionKeySizeInfo */

// BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams
type BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams struct {
	InputBandwidth uint32
	InputCodedDataSize uint16
	InputCodingFormat uint64
	InputDataPath uint8
	InputPCMDataFormat uint8
	InputPCMSamplePayloadMSBPosition uint8
	InputTransportUnitSize uint8
	MaxLatency uint16
	OutputBandwidth uint32
	OutputCodedDataSize uint16
	OutputCodingFormat uint64
	OutputDataPath uint8
	OutputPCMDataFormat uint8
	OutputPCMSamplePayloadMSBPosition uint8
	OutputTransportUnitSize uint8
	PacketType uint16
	ReceiveBandwidth uint32
	ReceiveCodecFrameSize uint16
	ReceiveCodingFormat uint64
	RetransmissionEffort uint8
	TransmitBandwidth uint32
	TransmitCodecFrameSize uint16
	TransmitCodingFormat uint64
}/* debug [types.gen.go/struct]: BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams */

// BluetoothHCIEnhancedSetupSynchronousConnectionParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEnhancedSetupSynchronousConnectionParams
type BluetoothHCIEnhancedSetupSynchronousConnectionParams struct {
	InputBandwidth uint32
	InputCodedDataSize uint16
	InputCodingFormat uint64
	InputDataPath uint8
	InputPCMDataFormat uint8
	InputPCMSamplePayloadMSBPosition uint8
	InputTransportUnitSize uint8
	MaxLatency uint16
	OutputBandwidth uint32
	OutputCodedDataSize uint16
	OutputCodingFormat uint64
	OutputDataPath uint8
	OutputPCMDataFormat uint8
	OutputPCMSamplePayloadMSBPosition uint8
	OutputTransportUnitSize uint8
	PacketType uint16
	ReceiveBandwidth uint32
	ReceiveCodecFrameSize uint16
	ReceiveCodingFormat uint64
	RetransmissionEffort uint8
	TransmitBandwidth uint32
	TransmitCodecFrameSize uint16
	TransmitCodingFormat uint64
}/* debug [types.gen.go/struct]: BluetoothHCIEnhancedSetupSynchronousConnectionParams */

// BluetoothHCIEventAuthenticationCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventAuthenticationCompleteResults
type BluetoothHCIEventAuthenticationCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventAuthenticationCompleteResults */

// BluetoothHCIEventChangeConnectionLinkKeyCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventChangeConnectionLinkKeyCompleteResults
type BluetoothHCIEventChangeConnectionLinkKeyCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventChangeConnectionLinkKeyCompleteResults */

// BluetoothHCIEventConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionCompleteResults
type BluetoothHCIEventConnectionCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	DeviceAddress BluetoothDeviceAddress
	EncryptionMode BluetoothHCIEncryptionMode
	LinkType BluetoothLinkType
}/* debug [types.gen.go/struct]: BluetoothHCIEventConnectionCompleteResults */

// BluetoothHCIEventConnectionPacketTypeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionPacketTypeResults
type BluetoothHCIEventConnectionPacketTypeResults struct {
	ConnectionHandle BluetoothConnectionHandle
	PacketType BluetoothPacketType
}/* debug [types.gen.go/struct]: BluetoothHCIEventConnectionPacketTypeResults */

// BluetoothHCIEventConnectionRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionRequestResults
type BluetoothHCIEventConnectionRequestResults struct {
	ClassOfDevice BluetoothClassOfDevice
	DeviceAddress BluetoothDeviceAddress
	LinkType BluetoothLinkType
}/* debug [types.gen.go/struct]: BluetoothHCIEventConnectionRequestResults */

// BluetoothHCIEventDataBufferOverflowResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventDataBufferOverflowResults
type BluetoothHCIEventDataBufferOverflowResults struct {
	LinkType BluetoothLinkType
}/* debug [types.gen.go/struct]: BluetoothHCIEventDataBufferOverflowResults */

// BluetoothHCIEventDisconnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventDisconnectionCompleteResults
type BluetoothHCIEventDisconnectionCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Reason BluetoothReasonCode
}/* debug [types.gen.go/struct]: BluetoothHCIEventDisconnectionCompleteResults */

// BluetoothHCIEventEncryptionChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventEncryptionChangeResults
type BluetoothHCIEventEncryptionChangeResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Enable BluetoothEncryptionEnable
}/* debug [types.gen.go/struct]: BluetoothHCIEventEncryptionChangeResults */

// BluetoothHCIEventEncryptionKeyRefreshCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventEncryptionKeyRefreshCompleteResults
type BluetoothHCIEventEncryptionKeyRefreshCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventEncryptionKeyRefreshCompleteResults */

// BluetoothHCIEventFlowSpecificationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventFlowSpecificationData
type BluetoothHCIEventFlowSpecificationData struct {
	AccessLatency uint32
	ConnectionHandle BluetoothConnectionHandle
	Flags uint8
	FlowDirection uint8
	PeakBandwidth uint32
	ServiceType uint8
	TokenBucketSize uint32
	TokenRate uint32
}/* debug [types.gen.go/struct]: BluetoothHCIEventFlowSpecificationData */

// BluetoothHCIEventFlushOccurredResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventFlushOccurredResults
type BluetoothHCIEventFlushOccurredResults struct {
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventFlushOccurredResults */

// BluetoothHCIEventHardwareErrorResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventHardwareErrorResults
type BluetoothHCIEventHardwareErrorResults struct {
	Error BluetoothHCIStatus
}/* debug [types.gen.go/struct]: BluetoothHCIEventHardwareErrorResults */

// BluetoothHCIEventLEConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEConnectionCompleteResults
type BluetoothHCIEventLEConnectionCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	ConnInterval uint16
	ConnLatency uint16
	MasterClockAccuracy uint8
	PeerAddress BluetoothDeviceAddress
	PeerAddressType uint8
	Role uint8
	SupervisionTimeout uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventLEConnectionCompleteResults */

// BluetoothHCIEventLEConnectionUpdateCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEConnectionUpdateCompleteResults
type BluetoothHCIEventLEConnectionUpdateCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	ConnInterval uint16
	ConnLatency uint16
	SupervisionTimeout uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventLEConnectionUpdateCompleteResults */

// BluetoothHCIEventLEEnhancedConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEEnhancedConnectionCompleteResults
type BluetoothHCIEventLEEnhancedConnectionCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	ConnInterval uint16
	ConnLatency uint16
	LocalResolvablePrivateAddress BluetoothDeviceAddress
	MasterClockAccuracy uint8
	PeerAddress BluetoothDeviceAddress
	PeerAddressType uint8
	PeerResolvablePrivateAddress BluetoothDeviceAddress
	Role uint8
	SupervisionTimeout uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventLEEnhancedConnectionCompleteResults */

// BluetoothHCIEventLELongTermKeyRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLELongTermKeyRequestResults
type BluetoothHCIEventLELongTermKeyRequestResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Ediv uint16
	RandomNumber uint8
}/* debug [types.gen.go/struct]: BluetoothHCIEventLELongTermKeyRequestResults */

// BluetoothHCIEventLEMetaResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEMetaResults
type BluetoothHCIEventLEMetaResults struct {
	Data uint8
	Length uint8
}/* debug [types.gen.go/struct]: BluetoothHCIEventLEMetaResults */

// BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults
type BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	UsedFeatures BluetoothHCISupportedFeatures
}/* debug [types.gen.go/struct]: BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults */

// BluetoothHCIEventLinkKeyNotificationResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLinkKeyNotificationResults
type BluetoothHCIEventLinkKeyNotificationResults struct {
	DeviceAddress BluetoothDeviceAddress
	KeyType BluetoothKeyType
	LinkKey BluetoothKey
}/* debug [types.gen.go/struct]: BluetoothHCIEventLinkKeyNotificationResults */

// BluetoothHCIEventMasterLinkKeyCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventMasterLinkKeyCompleteResults
type BluetoothHCIEventMasterLinkKeyCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	KeyFlag BluetoothKeyFlag
}/* debug [types.gen.go/struct]: BluetoothHCIEventMasterLinkKeyCompleteResults */

// BluetoothHCIEventMaxSlotsChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventMaxSlotsChangeResults
type BluetoothHCIEventMaxSlotsChangeResults struct {
	ConnectionHandle BluetoothConnectionHandle
	MaxSlots BluetoothMaxSlots
}/* debug [types.gen.go/struct]: BluetoothHCIEventMaxSlotsChangeResults */

// BluetoothHCIEventModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventModeChangeResults
type BluetoothHCIEventModeChangeResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Mode BluetoothHCIConnectionMode
	ModeInterval BluetoothHCIModeInterval
}/* debug [types.gen.go/struct]: BluetoothHCIEventModeChangeResults */

// BluetoothHCIEventPageScanModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventPageScanModeChangeResults
type BluetoothHCIEventPageScanModeChangeResults struct {
	DeviceAddress BluetoothDeviceAddress
	PageScanMode BluetoothPageScanMode
}/* debug [types.gen.go/struct]: BluetoothHCIEventPageScanModeChangeResults */

// BluetoothHCIEventPageScanRepetitionModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventPageScanRepetitionModeChangeResults
type BluetoothHCIEventPageScanRepetitionModeChangeResults struct {
	DeviceAddress BluetoothDeviceAddress
	PageScanRepetitionMode BluetoothPageScanRepetitionMode
}/* debug [types.gen.go/struct]: BluetoothHCIEventPageScanRepetitionModeChangeResults */

// BluetoothHCIEventQoSSetupCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventQoSSetupCompleteResults
type BluetoothHCIEventQoSSetupCompleteResults struct {
	ConnectionHandle BluetoothConnectionHandle
	SetupParams BluetoothHCIQualityOfServiceSetupParams
}/* debug [types.gen.go/struct]: BluetoothHCIEventQoSSetupCompleteResults */

// BluetoothHCIEventQoSViolationResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventQoSViolationResults
type BluetoothHCIEventQoSViolationResults struct {
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventQoSViolationResults */

// BluetoothHCIEventReadClockOffsetResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadClockOffsetResults
type BluetoothHCIEventReadClockOffsetResults struct {
	ClockOffset BluetoothClockOffset
	ConnectionHandle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadClockOffsetResults */

// BluetoothHCIEventReadExtendedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadExtendedFeaturesResults
type BluetoothHCIEventReadExtendedFeaturesResults struct {
	ConnectionHandle BluetoothConnectionHandle
	SupportedFeaturesInfo BluetoothHCIExtendedFeaturesInfo
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadExtendedFeaturesResults */

// BluetoothHCIEventReadRemoteExtendedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteExtendedFeaturesResults
type BluetoothHCIEventReadRemoteExtendedFeaturesResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Error BluetoothHCIStatus
	LmpFeatures BluetoothHCISupportedFeatures
	MaxPage BluetoothHCIPageNumber
	Page BluetoothHCIPageNumber
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadRemoteExtendedFeaturesResults */

// BluetoothHCIEventReadRemoteSupportedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteSupportedFeaturesResults
type BluetoothHCIEventReadRemoteSupportedFeaturesResults struct {
	ConnectionHandle BluetoothConnectionHandle
	Error BluetoothHCIStatus
	LmpFeatures BluetoothHCISupportedFeatures
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadRemoteSupportedFeaturesResults */

// BluetoothHCIEventReadRemoteVersionInfoResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteVersionInfoResults
type BluetoothHCIEventReadRemoteVersionInfoResults struct {
	ConnectionHandle BluetoothConnectionHandle
	LmpSubversion BluetoothLMPSubversion
	LmpVersion BluetoothLMPVersion
	ManufacturerName BluetoothManufacturerName
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadRemoteVersionInfoResults */

// BluetoothHCIEventReadSupportedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadSupportedFeaturesResults
type BluetoothHCIEventReadSupportedFeaturesResults struct {
	ConnectionHandle BluetoothConnectionHandle
	SupportedFeatures BluetoothHCISupportedFeatures
}/* debug [types.gen.go/struct]: BluetoothHCIEventReadSupportedFeaturesResults */

// BluetoothHCIEventRemoteNameRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventRemoteNameRequestResults
type BluetoothHCIEventRemoteNameRequestResults struct {
	DeviceAddress BluetoothDeviceAddress
	DeviceName BluetoothDeviceName
}/* debug [types.gen.go/struct]: BluetoothHCIEventRemoteNameRequestResults */

// BluetoothHCIEventReturnLinkKeysResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReturnLinkKeysResults
type BluetoothHCIEventReturnLinkKeysResults struct {
	LinkKeys unsafe.Pointer
	DeviceAddress BluetoothDeviceAddress
	LinkKey BluetoothKey
	NumLinkKeys uint8
}/* debug [types.gen.go/struct]: BluetoothHCIEventReturnLinkKeysResults */

// BluetoothHCIEventRoleChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventRoleChangeResults
type BluetoothHCIEventRoleChangeResults struct {
	ConnectionHandle BluetoothConnectionHandle
	DeviceAddress BluetoothDeviceAddress
	Role BluetoothRole
}/* debug [types.gen.go/struct]: BluetoothHCIEventRoleChangeResults */

// BluetoothHCIEventSimplePairingCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSimplePairingCompleteResults
type BluetoothHCIEventSimplePairingCompleteResults struct {
	DeviceAddress BluetoothDeviceAddress
}/* debug [types.gen.go/struct]: BluetoothHCIEventSimplePairingCompleteResults */

// BluetoothHCIEventSniffSubratingResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSniffSubratingResults
type BluetoothHCIEventSniffSubratingResults struct {
	ConnectionHandle BluetoothConnectionHandle
	MaxReceiveLatency uint16
	MaxTransmitLatency uint16
	MinLocalTimeout uint16
	MinRemoteTimeout uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventSniffSubratingResults */

// BluetoothHCIEventSynchronousConnectionChangedResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSynchronousConnectionChangedResults
type BluetoothHCIEventSynchronousConnectionChangedResults struct {
	ConnectionHandle BluetoothConnectionHandle
	ReceivePacketLength uint16
	RetransmissionWindow uint8
	TransmissionInterval uint8
	TransmitPacketLength uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventSynchronousConnectionChangedResults */

// BluetoothHCIEventSynchronousConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSynchronousConnectionCompleteResults
type BluetoothHCIEventSynchronousConnectionCompleteResults struct {
	AirMode BluetoothAirMode
	ConnectionHandle BluetoothConnectionHandle
	DeviceAddress BluetoothDeviceAddress
	LinkType BluetoothLinkType
	ReceivePacketLength uint16
	RetransmissionWindow uint8
	TransmissionInterval uint8
	TransmitPacketLength uint16
}/* debug [types.gen.go/struct]: BluetoothHCIEventSynchronousConnectionCompleteResults */

// BluetoothHCIEventVendorSpecificResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventVendorSpecificResults
type BluetoothHCIEventVendorSpecificResults struct {
	Data uint8
	Length uint8
}/* debug [types.gen.go/struct]: BluetoothHCIEventVendorSpecificResults */

// BluetoothHCIExtendedFeaturesInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedFeaturesInfo
type BluetoothHCIExtendedFeaturesInfo struct {
	Data uint8
	MaxPage BluetoothHCIPageNumber
	Page BluetoothHCIPageNumber
}/* debug [types.gen.go/struct]: BluetoothHCIExtendedFeaturesInfo */

// BluetoothHCIExtendedInquiryResponse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedInquiryResponse
type BluetoothHCIExtendedInquiryResponse struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothHCIExtendedInquiryResponse */

// BluetoothHCIExtendedInquiryResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedInquiryResult
type BluetoothHCIExtendedInquiryResult struct {
	ClassOfDevice BluetoothClassOfDevice
	ClockOffset BluetoothClockOffset
	DeviceAddress BluetoothDeviceAddress
	ExtendedInquiryResponse BluetoothHCIExtendedInquiryResponse
	NumberOfReponses uint8
	PageScanRepetitionMode BluetoothPageScanRepetitionMode
	Reserved uint8
	RSSIValue BluetoothHCIRSSIValue
}/* debug [types.gen.go/struct]: BluetoothHCIExtendedInquiryResult */

// BluetoothHCIFailedContactInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIFailedContactInfo
type BluetoothHCIFailedContactInfo struct {
	Count BluetoothHCIFailedContactCount
	Handle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothHCIFailedContactInfo */

// BluetoothHCIInquiryAccessCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryAccessCode
type BluetoothHCIInquiryAccessCode struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothHCIInquiryAccessCode */

// BluetoothHCIInquiryResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryResult
type BluetoothHCIInquiryResult struct {
	ClassOfDevice BluetoothClassOfDevice
	ClockOffset BluetoothClockOffset
	DeviceAddress BluetoothDeviceAddress
	PageScanMode BluetoothHCIPageScanMode
	PageScanPeriodMode BluetoothHCIPageScanPeriodMode
	PageScanRepetitionMode BluetoothPageScanRepetitionMode
}/* debug [types.gen.go/struct]: BluetoothHCIInquiryResult */

// BluetoothHCIInquiryResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryResults
type BluetoothHCIInquiryResults struct {
	Count ItemCount
	Results BluetoothHCIInquiryResult
}/* debug [types.gen.go/struct]: BluetoothHCIInquiryResults */

// BluetoothHCIInquiryWithRSSIResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryWithRSSIResult
type BluetoothHCIInquiryWithRSSIResult struct {
	ClassOfDevice BluetoothClassOfDevice
	ClockOffset BluetoothClockOffset
	DeviceAddress BluetoothDeviceAddress
	PageScanRepetitionMode BluetoothPageScanRepetitionMode
	Reserved uint8
	RSSIValue BluetoothHCIRSSIValue
}/* debug [types.gen.go/struct]: BluetoothHCIInquiryWithRSSIResult */

// BluetoothHCIInquiryWithRSSIResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryWithRSSIResults
type BluetoothHCIInquiryWithRSSIResults struct {
	Count ItemCount
	Results BluetoothHCIInquiryWithRSSIResult
}/* debug [types.gen.go/struct]: BluetoothHCIInquiryWithRSSIResults */

// BluetoothHCILEBufferSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILEBufferSize
type BluetoothHCILEBufferSize struct {
	ACLDataPacketLength uint16
	TotalNumACLDataPackets uint8
}/* debug [types.gen.go/struct]: BluetoothHCILEBufferSize */

// BluetoothHCILinkPolicySettingsInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkPolicySettingsInfo
type BluetoothHCILinkPolicySettingsInfo struct {
	Handle BluetoothConnectionHandle
	Settings BluetoothHCILinkPolicySettings
}/* debug [types.gen.go/struct]: BluetoothHCILinkPolicySettingsInfo */

// BluetoothHCILinkQualityInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkQualityInfo
type BluetoothHCILinkQualityInfo struct {
	Handle BluetoothConnectionHandle
	QualityValue BluetoothHCILinkQuality
}/* debug [types.gen.go/struct]: BluetoothHCILinkQualityInfo */

// BluetoothHCILinkSupervisionTimeout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkSupervisionTimeout
type BluetoothHCILinkSupervisionTimeout struct {
	Handle BluetoothConnectionHandle
	Timeout uint16
}/* debug [types.gen.go/struct]: BluetoothHCILinkSupervisionTimeout */

// BluetoothHCIQualityOfServiceSetupParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIQualityOfServiceSetupParams
type BluetoothHCIQualityOfServiceSetupParams struct {
	DelayVariation uint32
	Flags uint8
	Latency uint32
	PeakBandwidth uint32
	ServiceType uint8
	TokenRate uint32
}/* debug [types.gen.go/struct]: BluetoothHCIQualityOfServiceSetupParams */

// BluetoothHCIReadExtendedInquiryResponseResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadExtendedInquiryResponseResults
type BluetoothHCIReadExtendedInquiryResponseResults struct {
	ExtendedInquiryResponse BluetoothHCIExtendedInquiryResponse
	OutFECRequired BluetoothHCIFECRequired
}/* debug [types.gen.go/struct]: BluetoothHCIReadExtendedInquiryResponseResults */

// BluetoothHCIReadLMPHandleResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadLMPHandleResults
type BluetoothHCIReadLMPHandleResults struct {
	Handle BluetoothConnectionHandle
	Lmp_handle BluetoothLMPHandle
	Reserved uint32
}/* debug [types.gen.go/struct]: BluetoothHCIReadLMPHandleResults */

// BluetoothHCIReadLocalOOBDataResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadLocalOOBDataResults
type BluetoothHCIReadLocalOOBDataResults struct {
	Hash BluetoothHCISimplePairingOOBData
	Randomizer BluetoothHCISimplePairingOOBData
}/* debug [types.gen.go/struct]: BluetoothHCIReadLocalOOBDataResults */

// BluetoothHCIRequestCallbackInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRequestCallbackInfo
type BluetoothHCIRequestCallbackInfo struct {
	AsyncIDRefCon unsafe.Pointer
	InternalRefCon unsafe.Pointer
	Reserved unsafe.Pointer
	UserCallback unsafe.Pointer
	UserRefCon unsafe.Pointer
}/* debug [types.gen.go/struct]: BluetoothHCIRequestCallbackInfo */

// BluetoothHCIRoleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRoleInfo
type BluetoothHCIRoleInfo struct {
	Handle BluetoothConnectionHandle
	Role uint8
}/* debug [types.gen.go/struct]: BluetoothHCIRoleInfo */

// BluetoothHCIRSSIInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRSSIInfo
type BluetoothHCIRSSIInfo struct {
	Handle BluetoothConnectionHandle
	RSSIValue BluetoothHCIRSSIValue
}/* debug [types.gen.go/struct]: BluetoothHCIRSSIInfo */

// BluetoothHCIScanActivity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIScanActivity
type BluetoothHCIScanActivity struct {
	ScanInterval uint16
	ScanWindow uint16
}/* debug [types.gen.go/struct]: BluetoothHCIScanActivity */

// BluetoothHCISetupSynchronousConnectionParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISetupSynchronousConnectionParams
type BluetoothHCISetupSynchronousConnectionParams struct {
	MaxLatency uint16
	PacketType uint16
	ReceiveBandwidth uint32
	RetransmissionEffort uint8
	TransmitBandwidth uint32
	VoiceSetting uint16
}/* debug [types.gen.go/struct]: BluetoothHCISetupSynchronousConnectionParams */

// BluetoothHCISimplePairingOOBData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISimplePairingOOBData
type BluetoothHCISimplePairingOOBData struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothHCISimplePairingOOBData */

// BluetoothHCIStoredLinkKeysInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIStoredLinkKeysInfo
type BluetoothHCIStoredLinkKeysInfo struct {
	MaxNumLinkKeysAllowedInDevice uint16
	NumLinkKeysRead uint16
}/* debug [types.gen.go/struct]: BluetoothHCIStoredLinkKeysInfo */

// BluetoothHCISupportedCommands
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISupportedCommands
type BluetoothHCISupportedCommands struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothHCISupportedCommands */

// BluetoothHCISupportedFeatures
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISupportedFeatures
type BluetoothHCISupportedFeatures struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothHCISupportedFeatures */

// BluetoothHCITransmitPowerLevelInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCITransmitPowerLevelInfo
type BluetoothHCITransmitPowerLevelInfo struct {
	Handle BluetoothConnectionHandle
	Level BluetoothHCITransmitPowerLevel
}/* debug [types.gen.go/struct]: BluetoothHCITransmitPowerLevelInfo */

// BluetoothHCIVersionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIVersionInfo
type BluetoothHCIVersionInfo struct {
	HciRevision uint16
	HciVersion uint8
	LmpSubVersion BluetoothLMPSubversion
	LmpVersion BluetoothLMPVersion
	ManufacturerName BluetoothManufacturerName
}/* debug [types.gen.go/struct]: BluetoothHCIVersionInfo */

// BluetoothIOCapabilityResponse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothIOCapabilityResponse
type BluetoothIOCapabilityResponse struct {
	AuthenticationRequirements BluetoothAuthenticationRequirements
	DeviceAddress BluetoothDeviceAddress
	IoCapability BluetoothIOCapability
	OOBDataPresence BluetoothOOBDataPresence
}/* debug [types.gen.go/struct]: BluetoothIOCapabilityResponse */

// BluetoothIRK
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothIRK
type BluetoothIRK struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothIRK */

// BluetoothKey
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothKey
type BluetoothKey struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothKey */

// BluetoothKeypressNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothKeypressNotification
type BluetoothKeypressNotification struct {
	DeviceAddress BluetoothDeviceAddress
	NotificationType BluetoothKeypressNotificationType
}/* debug [types.gen.go/struct]: BluetoothKeypressNotification */

// BluetoothL2CAPQualityOfServiceOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPQualityOfServiceOptions
type BluetoothL2CAPQualityOfServiceOptions struct {
	DelayVariation uint32
	Flags uint8
	Latency uint32
	PeakBandwidth uint32
	ServiceType uint8
	TokenBucketSize uint32
	TokenRate uint32
}/* debug [types.gen.go/struct]: BluetoothL2CAPQualityOfServiceOptions */

// BluetoothL2CAPRetransmissionAndFlowControlOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPRetransmissionAndFlowControlOptions
type BluetoothL2CAPRetransmissionAndFlowControlOptions struct {
	Flags uint8
	MaxPDUPayloadSize uint16
	MaxTransmit uint8
	MonitorTimeout uint16
	RetransmissionTimeout uint16
	TxWindowSize uint8
}/* debug [types.gen.go/struct]: BluetoothL2CAPRetransmissionAndFlowControlOptions */

// BluetoothPINCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothPINCode
type BluetoothPINCode struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothPINCode */

// BluetoothReadClockInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothReadClockInfo
type BluetoothReadClockInfo struct {
	Accuracy uint16
	Clock uint32
	Handle BluetoothConnectionHandle
}/* debug [types.gen.go/struct]: BluetoothReadClockInfo */

// BluetoothRemoteHostSupportedFeaturesNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothRemoteHostSupportedFeaturesNotification
type BluetoothRemoteHostSupportedFeaturesNotification struct {
	DeviceAddress BluetoothDeviceAddress
	HostSupportedFeatures BluetoothHCISupportedFeatures
}/* debug [types.gen.go/struct]: BluetoothRemoteHostSupportedFeaturesNotification */

// BluetoothSetEventMask
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothSetEventMask
type BluetoothSetEventMask struct {
	Data uint8
}/* debug [types.gen.go/struct]: BluetoothSetEventMask */

// BluetoothSynchronousConnectionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothSynchronousConnectionInfo
type BluetoothSynchronousConnectionInfo struct {
	MaxLatency BluetoothHCIMaxLatency
	PacketType BluetoothPacketType
	ReceiveBandWidth BluetoothHCIReceiveBandwidth
	RetransmissionEffort BluetoothHCIRetransmissionEffort
	TransmitBandWidth BluetoothHCITransmitBandwidth
	VoiceSetting BluetoothHCIVoiceSetting
}/* debug [types.gen.go/struct]: BluetoothSynchronousConnectionInfo */

// BluetoothTransportInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothTransportInfo
type BluetoothTransportInfo struct {
	ProductID uint32
	ProductName unsafe.Pointer
	TotalDataBytesReceived uint64
	TotalDataBytesSent uint64
	TotalSCOBytesReceived uint64
	TotalSCOBytesSent uint64
	Type uint32
	VendorID uint32
	VendorName unsafe.Pointer
}/* debug [types.gen.go/struct]: BluetoothTransportInfo */

// BluetoothUserConfirmationRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothUserConfirmationRequest
type BluetoothUserConfirmationRequest struct {
	DeviceAddress BluetoothDeviceAddress
	NumericValue BluetoothNumericValue
}/* debug [types.gen.go/struct]: BluetoothUserConfirmationRequest */

// BluetoothUserPasskeyNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothUserPasskeyNotification
type BluetoothUserPasskeyNotification struct {
	DeviceAddress BluetoothDeviceAddress
	Passkey BluetoothPasskey
}/* debug [types.gen.go/struct]: BluetoothUserPasskeyNotification */

// IOBluetoothDeviceSearchAttributes - Structure used to search for particular devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchAttributes
type IOBluetoothDeviceSearchAttributes struct {
	AttributeList IOBluetoothDeviceSearchDeviceAttributes
	DeviceAttributeCount ItemCount
	MaxResults ItemCount
	Options BluetoothDeviceSearchOptions
}/* debug [types.gen.go/struct]: IOBluetoothDeviceSearchAttributes */

// IOBluetoothDeviceSearchDeviceAttributes - Structure used to search for particular devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchDeviceAttributes
type IOBluetoothDeviceSearchDeviceAttributes struct {
	Address BluetoothDeviceAddress
	DeviceClassMajor BluetoothDeviceClassMajor
	DeviceClassMinor BluetoothDeviceClassMinor
	Name BluetoothDeviceName
	ServiceClassMajor BluetoothServiceClassMajor
}/* debug [types.gen.go/struct]: IOBluetoothDeviceSearchDeviceAttributes */

// IOBluetoothL2CAPChannelDataBlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDataBlock
type IOBluetoothL2CAPChannelDataBlock struct {
	DataPtr unsafe.Pointer
	DataSize uintptr
}/* debug [types.gen.go/struct]: IOBluetoothL2CAPChannelDataBlock */

// IOBluetoothL2CAPChannelEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelEvent
type IOBluetoothL2CAPChannelEvent struct {
	EventType BluetoothL2CAPChannelEventType
	Status int
	U unsafe.Pointer
	Data BluetoothL2CAPChannelDataBlock
	Padding unsafe.Pointer
	WriteRefCon unsafe.Pointer
}/* debug [types.gen.go/struct]: IOBluetoothL2CAPChannelEvent */

// OBEXAbortCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAbortCommandData
type OBEXAbortCommandData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
}/* debug [types.gen.go/struct]: OBEXAbortCommandData */

// OBEXAbortCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAbortCommandResponseData
type OBEXAbortCommandResponseData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	ServerResponseOpCode OBEXOpCode
}/* debug [types.gen.go/struct]: OBEXAbortCommandResponseData */

// OBEXConnectCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXConnectCommandData
type OBEXConnectCommandData struct {
	Flags OBEXFlags
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	MaxPacketSize OBEXMaxPacketLength
	Version OBEXVersion
}/* debug [types.gen.go/struct]: OBEXConnectCommandData */

// OBEXConnectCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXConnectCommandResponseData
type OBEXConnectCommandResponseData struct {
	Flags OBEXFlags
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	MaxPacketSize OBEXMaxPacketLength
	ServerResponseOpCode OBEXOpCode
	Version OBEXVersion
}/* debug [types.gen.go/struct]: OBEXConnectCommandResponseData */

// OBEXDisconnectCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXDisconnectCommandData
type OBEXDisconnectCommandData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
}/* debug [types.gen.go/struct]: OBEXDisconnectCommandData */

// OBEXDisconnectCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXDisconnectCommandResponseData
type OBEXDisconnectCommandResponseData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	ServerResponseOpCode OBEXOpCode
}/* debug [types.gen.go/struct]: OBEXDisconnectCommandResponseData */

// OBEXErrorData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXErrorData
type OBEXErrorData struct {
	DataLength uintptr
	DataPtr unsafe.Pointer
	Error OBEXError
}/* debug [types.gen.go/struct]: OBEXErrorData */

// OBEXGetCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetCommandData
type OBEXGetCommandData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
}/* debug [types.gen.go/struct]: OBEXGetCommandData */

// OBEXGetCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetCommandResponseData
type OBEXGetCommandResponseData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	ServerResponseOpCode OBEXOpCode
}/* debug [types.gen.go/struct]: OBEXGetCommandResponseData */

// OBEXPutCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXPutCommandData
type OBEXPutCommandData struct {
	BodyDataLeftToSend uintptr
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
}/* debug [types.gen.go/struct]: OBEXPutCommandData */

// OBEXPutCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXPutCommandResponseData
type OBEXPutCommandResponseData struct {
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	ServerResponseOpCode OBEXOpCode
}/* debug [types.gen.go/struct]: OBEXPutCommandResponseData */

// OBEXSessionEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionEvent
type OBEXSessionEvent struct {
	IsEndOfEventData unsafe.Pointer
	RefCon unsafe.Pointer
	Reserved1 unsafe.Pointer
	Reserved2 unsafe.Pointer
	Session OBEXSessionRef
	Type OBEXSessionEventType
	U unsafe.Pointer
	AbortCommandData OBEXAbortCommandData
	AbortCommandResponseData OBEXAbortCommandResponseData
	ConnectCommandData OBEXConnectCommandData
	ConnectCommandResponseData OBEXConnectCommandResponseData
	DisconnectCommandData OBEXDisconnectCommandData
	DisconnectCommandResponseData OBEXDisconnectCommandResponseData
	ErrorData OBEXErrorData
	GetCommandData OBEXGetCommandData
	GetCommandResponseData OBEXGetCommandResponseData
	PutCommandData OBEXPutCommandData
	PutCommandResponseData OBEXPutCommandResponseData
	SetPathCommandData OBEXSetPathCommandData
	SetPathCommandResponseData OBEXSetPathCommandResponseData
}/* debug [types.gen.go/struct]: OBEXSessionEvent */

// OBEXSetPathCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSetPathCommandData
type OBEXSetPathCommandData struct {
	Constants OBEXConstants
	Flags OBEXFlags
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
}/* debug [types.gen.go/struct]: OBEXSetPathCommandData */

// OBEXSetPathCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSetPathCommandResponseData
type OBEXSetPathCommandResponseData struct {
	Constants OBEXConstants
	Flags OBEXFlags
	HeaderDataLength uintptr
	HeaderDataPtr unsafe.Pointer
	ServerResponseOpCode OBEXOpCode
}/* debug [types.gen.go/struct]: OBEXSetPathCommandResponseData */

// OBEXTransportEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXTransportEvent
type OBEXTransportEvent struct {
	DataLength uintptr
	DataPtr unsafe.Pointer
	Status OBEXError
	Type OBEXTransportEventType
}/* debug [types.gen.go/struct]: OBEXTransportEvent */





