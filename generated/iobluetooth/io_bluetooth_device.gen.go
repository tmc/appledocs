// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothDevice] class.
var (
	BluetoothDeviceClass     _BluetoothDeviceClass
	BluetoothDeviceClassOnce sync.Once
)

func getBluetoothDeviceClass() _BluetoothDeviceClass {
	BluetoothDeviceClassOnce.Do(func() {
		BluetoothDeviceClass = _BluetoothDeviceClass{objc.GetClass("IOBluetoothDevice")}
	})
	return BluetoothDeviceClass
}

type _BluetoothDeviceClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothDevice] class.
type IBluetoothDevice interface {
	IBluetoothObject
	// properties:
	AddressString() string
	ClassOfDevice() BluetoothClassOfDevice
	ConnectionHandle() BluetoothConnectionHandle
	DeviceClassMajor() BluetoothDeviceClassMajor
	DeviceClassMinor() BluetoothDeviceClassMinor
	HandsFreeAudioGateway() bool
	HandsFreeDevice() bool
	LastNameUpdate() foundation.NSDate
	Name() string
	NameOrAddress() string
	ServiceClassMajor() BluetoothServiceClassMajor
	Services() objc.ID
	IsHandsFreeAudioGateway() bool
	SetIsHandsFreeAudioGateway(value bool)
	IsHandsFreeDevice() bool
	SetIsHandsFreeDevice(value bool)
	// methods:
	AddToFavorites() unsafe.Pointer
	CloseConnection() unsafe.Pointer
	GetAddress() unsafe.Pointer
	GetClockOffset() BluetoothClockOffset
	GetEncryptionMode() BluetoothHCIEncryptionMode
	GetLastInquiryUpdate() foundation.Date
	GetLastServicesUpdate() foundation.Date
	GetLinkType() BluetoothLinkType
	GetPageScanMode() BluetoothPageScanMode
	GetPageScanPeriodMode() BluetoothPageScanPeriodMode
	GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode
	GetServiceRecordForUUID(sdpUUID IOBluetoothSDPUUID) IBluetoothSDPServiceRecord
	HandsFreeAudioGatewayServiceRecord() IBluetoothSDPServiceRecord
	HandsFreeDeviceServiceRecord() IBluetoothSDPServiceRecord
	IsConnected() bool
	IsFavorite() bool
	IsIncoming() bool
	IsPaired() bool
	OpenConnection() unsafe.Pointer
	OpenConnectionWithPageTimeoutAuthenticationRequired(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout, authenticationRequired bool) unsafe.Pointer
	OpenL2CAPChannelAsyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelDelegate objectivec.IObject) unsafe.Pointer
	OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) unsafe.Pointer
	OpenL2CAPChannelSyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelDelegate objectivec.IObject) unsafe.Pointer
	OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) unsafe.Pointer
	OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID, channelDelegate objectivec.IObject) unsafe.Pointer
	OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID, channelDelegate objectivec.IObject) unsafe.Pointer
	PerformSDPQuery(target objectivec.IObject) unsafe.Pointer
	PerformSDPQueryUuids(target objectivec.IObject, uuidArray objectivec.IObject) unsafe.Pointer
	RawRSSI() BluetoothHCIRSSIValue
	RecentAccessDate() foundation.Date
	RegisterForDisconnectNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification
	RemoteNameRequest(target objectivec.IObject) unsafe.Pointer
	RemoteNameRequestWithPageTimeout(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout) unsafe.Pointer
	RemoveFromFavorites() unsafe.Pointer
	RequestAuthentication() unsafe.Pointer
	RSSI() BluetoothHCIRSSIValue
	SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer
	SetSupervisionTimeout(timeout unsafe.Pointer) unsafe.Pointer
}

// An instance of IOBluetoothDevice represents a single remote Bluetooth device.
//
// An IOBluetoothDevice object may exist independent of the existence of a baseband connection with the target device. Using this object, a client can request creation and destruction of baseband connections, and request the opening of L2CAP and RFCOMM channels on the remote device. Many of the other APIs in the IOBluetooth framework will return this object, or it’s C counterpart (IOBluetoothDeviceRef).


// An instance of IOBluetoothDevice represents a single remote Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice
type BluetoothDevice struct {
	BluetoothObject
}

// BluetoothDeviceFrom constructs a [BluetoothDevice] from an unsafe.Pointer.
//
// An instance of IOBluetoothDevice represents a single remote Bluetooth device.
func BluetoothDeviceFrom(ptr unsafe.Pointer) BluetoothDevice {
	return BluetoothDevice{
		BluetoothObject: BluetoothObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceClass) Alloc() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothDeviceClass) New() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothDevice) Init() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothDevice) Autorelease() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothDevice creates a new BluetoothDevice instance.
func NewBluetoothDevice() BluetoothDevice {
	return getBluetoothDeviceClass().New()
}



// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func NewBluetoothDeviceWithAddress(address unsafe.Pointer) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddress:"), address)
	return rv
}


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func NewBluetoothDeviceWithAddressString(address string) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}



// Gets an array of the user’s favorite devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/favoriteDevices()
func (bc _BluetoothDeviceClass) FavoriteDevices() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("favoriteDevices"))
	return rv
}


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func (bc _BluetoothDeviceClass) DeviceWithAddress(address unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddress:"), address)
	return rv
}


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func (bc _BluetoothDeviceClass) DeviceWithAddressString(address string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}


// Gets an array of all of the paired devices on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/pairedDevices()
func (bc _BluetoothDeviceClass) PairedDevices() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("pairedDevices"))
	return rv
}


// Gets an array of recently used Bluetooth devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentDevices(_:)
func (bc _BluetoothDeviceClass) RecentDevices(numDevices unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("recentDevices:"), numDevices)
	return rv
}


// Allows a client to register for device connect notifications for any connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forConnectNotifications:selector:)
func (bc _BluetoothDeviceClass) RegisterForConnectNotificationsSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForConnectNotifications:selector:"), observer, inSelector)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withAddress:
func (bc _BluetoothDeviceClass) WithAddress(address unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withAddress:"), address)
	return rv
}


// Method call to convert an IOBluetoothDeviceRef into an IOBluetoothDevice *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withDeviceRef:
func (bc _BluetoothDeviceClass) WithDeviceRef(deviceRef BluetoothDeviceRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDeviceRef:"), deviceRef)
	return rv
}


// Adds the target device to the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addToFavorites()
func (b_ BluetoothDevice) AddToFavorites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("addToFavorites"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/awakeAfter(using:)
func (b_ BluetoothDevice) AwakeAfterUsingCoder(coder foundation.Coder) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}


// Close down the baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/closeConnection()
func (b_ BluetoothDevice) CloseConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("closeConnection"))
	return rv
}


// Get the Bluetooth device address for the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getAddress()
func (b_ BluetoothDevice) GetAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getAddress"))
	return rv
}


// Get the clock offset value of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getClockOffset()
func (b_ BluetoothDevice) GetClockOffset() BluetoothClockOffset {
	rv := objc.Send[BluetoothClockOffset](b_.ID, objc.Sel("getClockOffset"))
	return rv
}


// Get the encryption mode for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getEncryptionMode()
func (b_ BluetoothDevice) GetEncryptionMode() BluetoothHCIEncryptionMode {
	rv := objc.Send[BluetoothHCIEncryptionMode](b_.ID, objc.Sel("getEncryptionMode"))
	return rv
}


// Get the date/time of the last time the device was returned during an inquiry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastInquiryUpdate()
func (b_ BluetoothDevice) GetLastInquiryUpdate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("getLastInquiryUpdate"))
	return rv
}


// Get the date/time of the last SDP query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastServicesUpdate()
func (b_ BluetoothDevice) GetLastServicesUpdate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("getLastServicesUpdate"))
	return rv
}


// Get the link type for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLinkType()
func (b_ BluetoothDevice) GetLinkType() BluetoothLinkType {
	rv := objc.Send[BluetoothLinkType](b_.ID, objc.Sel("getLinkType"))
	return rv
}


// Get the page scan mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanMode()
func (b_ BluetoothDevice) GetPageScanMode() BluetoothPageScanMode {
	rv := objc.Send[BluetoothPageScanMode](b_.ID, objc.Sel("getPageScanMode"))
	return rv
}


// Get the value of the page scan period mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanPeriodMode()
func (b_ BluetoothDevice) GetPageScanPeriodMode() BluetoothPageScanPeriodMode {
	rv := objc.Send[BluetoothPageScanPeriodMode](b_.ID, objc.Sel("getPageScanPeriodMode"))
	return rv
}


// Get the value of the page scan repetition mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanRepetitionMode()
func (b_ BluetoothDevice) GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode {
	rv := objc.Send[BluetoothPageScanRepetitionMode](b_.ID, objc.Sel("getPageScanRepetitionMode"))
	return rv
}


// Search for a service record containing the given UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServiceRecord(for:)
func (b_ BluetoothDevice) GetServiceRecordForUUID(sdpUUID IOBluetoothSDPUUID) IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("getServiceRecordForUUID:"), sdpUUID)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeAudioGatewayServiceRecord()
func (b_ BluetoothDevice) HandsFreeAudioGatewayServiceRecord() IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("handsFreeAudioGatewayServiceRecord"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeDeviceServiceRecord()
func (b_ BluetoothDevice) HandsFreeDeviceServiceRecord() IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("handsFreeDeviceServiceRecord"))
	return rv
}


// Indicates whether a baseband connection to the device exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isConnected()
func (b_ BluetoothDevice) IsConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}


// Reports whether the target device is a favorite for the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isFavorite()
func (b_ BluetoothDevice) IsFavorite() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isFavorite"))
	return rv
}


// Returns TRUE if the device connection was generated by the remote host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isIncoming()
func (b_ BluetoothDevice) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}


// Returns whether the target device is paired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isPaired()
func (b_ BluetoothDevice) IsPaired() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPaired"))
	return rv
}


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection()
func (b_ BluetoothDevice) OpenConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openConnection"))
	return rv
}


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection(_:withPageTimeout:authenticationRequired:)
func (b_ BluetoothDevice) OpenConnectionWithPageTimeoutAuthenticationRequired(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout, authenticationRequired bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openConnection:withPageTimeout:authenticationRequired:"), target, pageTimeoutValue, authenticationRequired)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}


// Opens a new RFCOMM channel to the target device. Returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelAsync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openRFCOMMChannelAsync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}


// Opens a new RFCOMM channel to the target device. Returns only once the channel is open or failed to open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelSync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID, channelDelegate objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openRFCOMMChannelSync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}


// Performs an SDP query on the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:)
func (b_ BluetoothDevice) PerformSDPQuery(target objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("performSDPQuery:"), target)
	return rv
}


// Performs an SDP query on the target device with the specified service UUIDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:uuids:)
func (b_ BluetoothDevice) PerformSDPQueryUuids(target objectivec.IObject, uuidArray objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("performSDPQuery:uuids:"), target, uuidArray)
	return rv
}


// Get the raw RSSI device (if connected).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rawRSSI()
func (b_ BluetoothDevice) RawRSSI() BluetoothHCIRSSIValue {
	rv := objc.Send[BluetoothHCIRSSIValue](b_.ID, objc.Sel("rawRSSI"))
	return rv
}


// Returns the date/time of the most recent access of the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentAccessDate()
func (b_ BluetoothDevice) RecentAccessDate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("recentAccessDate"))
	return rv
}


// Allows a client to register for device disconnect notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forDisconnectNotification:selector:)
func (b_ BluetoothDevice) RegisterForDisconnectNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("registerForDisconnectNotification:selector:"), observer, inSelector)
	return rv
}


// Issues a remote name request to the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:)
func (b_ BluetoothDevice) RemoteNameRequest(target objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("remoteNameRequest:"), target)
	return rv
}


// Issues a remote name request to the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:withPageTimeout:)
func (b_ BluetoothDevice) RemoteNameRequestWithPageTimeout(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("remoteNameRequest:withPageTimeout:"), target, pageTimeoutValue)
	return rv
}


// Removes the target device from the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/removeFromFavorites()
func (b_ BluetoothDevice) RemoveFromFavorites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("removeFromFavorites"))
	return rv
}


// Requests that the existing baseband connection be authenticated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/requestAuthentication()
func (b_ BluetoothDevice) RequestAuthentication() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("requestAuthentication"))
	return rv
}


// Get the RSSI device (if connected), above or below the golden range. If the RSSI is within the golden range, a value of 0 is returned. For the actual RSSI value, use getRawRSSI. For more information, see the Bluetooth 4.0 Core Specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rssi()
func (b_ BluetoothDevice) RSSI() BluetoothHCIRSSIValue {
	rv := objc.Send[BluetoothHCIRSSIValue](b_.ID, objc.Sel("RSSI"))
	return rv
}


// Send an echo request over the L2CAP connection to a remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/sendL2CAPEchoRequest(_:length:)
func (b_ BluetoothDevice) SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sendL2CAPEchoRequest:length:"), data, length)
	return rv
}


// Sets the connection supervision timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/setSupervisionTimeout(_:)
func (b_ BluetoothDevice) SetSupervisionTimeout(timeout unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("setSupervisionTimeout:"), timeout)
	return rv
}


// Get a string representation of the Bluetooth device address for the target device. The format of the string is the same as returned by IOBluetoothNSStringFromDeviceAddress().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addressString
func (b_ BluetoothDevice) AddressString() string {
	rv := objc.Send[string](b_.ID, objc.Sel("addressString"))
	return rv
}


// Gets the full class of device value for the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/classOfDevice
func (b_ BluetoothDevice) ClassOfDevice() BluetoothClassOfDevice {
	rv := objc.Send[BluetoothClassOfDevice](b_.ID, objc.Sel("classOfDevice"))
	return rv
}


// Get the connection handle for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/connectionHandle
func (b_ BluetoothDevice) ConnectionHandle() BluetoothConnectionHandle {
	rv := objc.Send[BluetoothConnectionHandle](b_.ID, objc.Sel("connectionHandle"))
	return rv
}


// Get the major device class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMajor
func (b_ BluetoothDevice) DeviceClassMajor() BluetoothDeviceClassMajor {
	rv := objc.Send[BluetoothDeviceClassMajor](b_.ID, objc.Sel("deviceClassMajor"))
	return rv
}


// Get the minor service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMinor
func (b_ BluetoothDevice) DeviceClassMinor() BluetoothDeviceClassMinor {
	rv := objc.Send[BluetoothDeviceClassMinor](b_.ID, objc.Sel("deviceClassMinor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeAudioGateway
func (b_ BluetoothDevice) HandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeAudioGateway"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeDevice
func (b_ BluetoothDevice) HandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeDevice"))
	return rv
}


// Get the date/time of the last successful remote name request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/lastNameUpdate
func (b_ BluetoothDevice) LastNameUpdate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](b_.ID, objc.Sel("lastNameUpdate"))
	return rv
}


// Get the human readable name of the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/name
func (b_ BluetoothDevice) Name() string {
	rv := objc.Send[string](b_.ID, objc.Sel("name"))
	return rv
}


// Get the human readable name of the remote device. If the name is not present, it will return a string containing the device’s address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/nameOrAddress
func (b_ BluetoothDevice) NameOrAddress() string {
	rv := objc.Send[string](b_.ID, objc.Sel("nameOrAddress"))
	return rv
}


// Get the major service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/serviceClassMajor
func (b_ BluetoothDevice) ServiceClassMajor() BluetoothServiceClassMajor {
	rv := objc.Send[BluetoothServiceClassMajor](b_.ID, objc.Sel("serviceClassMajor"))
	return rv
}


// Gets an array of service records for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/services
func (b_ BluetoothDevice) Services() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("services"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) IsHandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeAudioGateway"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) SetIsHandsFreeAudioGateway(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeAudioGateway:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) IsHandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeDevice"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) SetIsHandsFreeDevice(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeDevice:"), value)
}


