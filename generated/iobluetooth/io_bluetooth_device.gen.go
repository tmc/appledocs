// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AddToFavorites() unsafe.Pointer
	CloseConnection() unsafe.Pointer
	GetAddress() unsafe.Pointer
	GetAddressString() string
	GetClassOfDevice() unsafe.Pointer
	GetClockOffset() unsafe.Pointer
	GetConnectionHandle() unsafe.Pointer
	GetDeviceClassMajor() unsafe.Pointer
	GetDeviceClassMinor() unsafe.Pointer
	GetDeviceRef() unsafe.Pointer
	GetEncryptionMode() unsafe.Pointer
	GetLastInquiryUpdate() unsafe.Pointer
	GetLastNameUpdate() unsafe.Pointer
	GetLastServicesUpdate() unsafe.Pointer
	GetLinkType() unsafe.Pointer
	GetName() string
	GetNameOrAddress() string
	GetPageScanMode() unsafe.Pointer
	GetPageScanPeriodMode() unsafe.Pointer
	GetPageScanRepetitionMode() unsafe.Pointer
	GetServiceClassMajor() unsafe.Pointer
	GetServiceRecordForUUID(sdpUUID unsafe.Pointer) unsafe.Pointer
	GetServices() unsafe.Pointer
	HandsFreeAudioGatewayDriverID() string
	HandsFreeAudioGatewayServiceRecord() unsafe.Pointer
	HandsFreeDeviceDriverID() string
	HandsFreeDeviceServiceRecord() unsafe.Pointer
	IsConnected() bool
	IsFavorite() bool
	IsIncoming() bool
	IsPaired() bool
	OpenConnection() unsafe.Pointer
	OpenConnectionWithPageTimeoutAuthenticationRequired(target objc.ID, pageTimeoutValue unsafe.Pointer, authenticationRequired bool) unsafe.Pointer
	OpenL2CAPChannelFindExistingNewChannel(psm unsafe.Pointer, findExisting bool, newChannel unsafe.Pointer) unsafe.Pointer
	OpenL2CAPChannelAsyncWithPSMDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer
	OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelConfiguration objc.ID, channelDelegate objc.ID) unsafe.Pointer
	OpenL2CAPChannelSyncWithPSMDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer
	OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelConfiguration objc.ID, channelDelegate objc.ID) unsafe.Pointer
	OpenRFCOMMChannelChannel(channelID unsafe.Pointer, rfcommChannel unsafe.Pointer) unsafe.Pointer
	OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel unsafe.Pointer, channelID unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer
	OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel unsafe.Pointer, channelID unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer
	PerformSDPQuery(target objc.ID) unsafe.Pointer
	PerformSDPQueryUuids(target objc.ID, uuidArray objc.ID) unsafe.Pointer
	RawRSSI() unsafe.Pointer
	RecentAccessDate() unsafe.Pointer
	RegisterForDisconnectNotificationSelector(observer objc.ID, inSelector objc.SEL) unsafe.Pointer
	RemoteNameRequest(target objc.ID) unsafe.Pointer
	RemoteNameRequestWithPageTimeout(target objc.ID, pageTimeoutValue unsafe.Pointer) unsafe.Pointer
	RemoveFromFavorites() unsafe.Pointer
	RequestAuthentication() unsafe.Pointer
	RSSI() unsafe.Pointer
	SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer
	SetSupervisionTimeout(timeout unsafe.Pointer) unsafe.Pointer
}

// An instance of IOBluetoothDevice represents a single remote Bluetooth device.
//
// An IOBluetoothDevice object may exist independent of the existence of a baseband connection with the target device. Using this object, a client can request creation and destruction of baseband connections, and request the opening of L2CAP and RFCOMM channels on the remote device. Many of the other APIs in the IOBluetooth framework will return this object, or it’s C counterpart (IOBluetoothDeviceRef).
//
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
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func NewBluetoothDeviceWithAddress(address unsafe.Pointer) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddress:"), address)
	return rv
}



// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func NewBluetoothDeviceWithAddressString(address string) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}


// Gets an array of the user’s favorite devices.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/favoriteDevices()
func (bc _BluetoothDeviceClass) FavoriteDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("favoriteDevices"))
	return rv
}

// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func (bc _BluetoothDeviceClass) DeviceWithAddress(address unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddress:"), address)
	return rv
}

// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func (bc _BluetoothDeviceClass) DeviceWithAddressString(address string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}

// Gets an array of all of the paired devices on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/pairedDevices()
func (bc _BluetoothDeviceClass) PairedDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("pairedDevices"))
	return rv
}

// Gets an array of recently used Bluetooth devices.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentDevices(_:)
func (bc _BluetoothDeviceClass) RecentDevices(numDevices unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("recentDevices:"), numDevices)
	return rv
}

// Allows a client to register for device connect notifications for any connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forConnectNotifications:selector:)
func (bc _BluetoothDeviceClass) RegisterForConnectNotificationsSelector(observer objc.ID, inSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("registerForConnectNotifications:selector:"), observer, inSelector)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withAddress:
func (bc _BluetoothDeviceClass) WithAddress(address unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withAddress:"), address)
	return rv
}

// Method call to convert an IOBluetoothDeviceRef into an IOBluetoothDevice *.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withDeviceRef:
func (bc _BluetoothDeviceClass) WithDeviceRef(deviceRef unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDeviceRef:"), deviceRef)
	return rv
}

// Adds the target device to the user’s favorite devices list.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addToFavorites()
func (b_ BluetoothDevice) AddToFavorites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("addToFavorites"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/awakeAfter(using:)
func (b_ BluetoothDevice) AwakeAfterUsingCoder(coder unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}

// Close down the baseband connection to the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/closeConnection()
func (b_ BluetoothDevice) CloseConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("closeConnection"))
	return rv
}

// Get the Bluetooth device address for the target device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getAddress()
func (b_ BluetoothDevice) GetAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getAddress"))
	return rv
}

// Get a string representation of the Bluetooth device address for the target device. The format of the string is the same as returned by IOBluetoothNSStringFromDeviceAddress().
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getAddressString
func (b_ BluetoothDevice) GetAddressString() string {
	rv := objc.Send[string](b_.ID, objc.Sel("getAddressString"))
	return rv
}

// Gets the full class of device value for the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getClassOfDevice
func (b_ BluetoothDevice) GetClassOfDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getClassOfDevice"))
	return rv
}

// Get the clock offset value of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getClockOffset()
func (b_ BluetoothDevice) GetClockOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getClockOffset"))
	return rv
}

// Get the connection handle for the baseband connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getConnectionHandle
func (b_ BluetoothDevice) GetConnectionHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getConnectionHandle"))
	return rv
}

// Get the major device class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getDeviceClassMajor
func (b_ BluetoothDevice) GetDeviceClassMajor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDeviceClassMajor"))
	return rv
}

// Get the minor service class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getDeviceClassMinor
func (b_ BluetoothDevice) GetDeviceClassMinor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDeviceClassMinor"))
	return rv
}

// Returns an IOBluetoothDeviceRef representation of the target IOBluetoothDevice object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getDeviceRef
func (b_ BluetoothDevice) GetDeviceRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDeviceRef"))
	return rv
}

// Get the encryption mode for the baseband connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getEncryptionMode()
func (b_ BluetoothDevice) GetEncryptionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getEncryptionMode"))
	return rv
}

// Get the date/time of the last time the device was returned during an inquiry.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastInquiryUpdate()
func (b_ BluetoothDevice) GetLastInquiryUpdate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getLastInquiryUpdate"))
	return rv
}

// Get the date/time of the last successful remote name request.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastNameUpdate
func (b_ BluetoothDevice) GetLastNameUpdate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getLastNameUpdate"))
	return rv
}

// Get the date/time of the last SDP query.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastServicesUpdate()
func (b_ BluetoothDevice) GetLastServicesUpdate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getLastServicesUpdate"))
	return rv
}

// Get the link type for the baseband connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLinkType()
func (b_ BluetoothDevice) GetLinkType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getLinkType"))
	return rv
}

// Get the human readable name of the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getName
func (b_ BluetoothDevice) GetName() string {
	rv := objc.Send[string](b_.ID, objc.Sel("getName"))
	return rv
}

// Get the human readable name of the remote device. If the name is not present, it will return a string containing the device’s address.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getNameOrAddress
func (b_ BluetoothDevice) GetNameOrAddress() string {
	rv := objc.Send[string](b_.ID, objc.Sel("getNameOrAddress"))
	return rv
}

// Get the page scan mode for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanMode()
func (b_ BluetoothDevice) GetPageScanMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getPageScanMode"))
	return rv
}

// Get the value of the page scan period mode for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanPeriodMode()
func (b_ BluetoothDevice) GetPageScanPeriodMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getPageScanPeriodMode"))
	return rv
}

// Get the value of the page scan repetition mode for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanRepetitionMode()
func (b_ BluetoothDevice) GetPageScanRepetitionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getPageScanRepetitionMode"))
	return rv
}

// Get the major service class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServiceClassMajor
func (b_ BluetoothDevice) GetServiceClassMajor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getServiceClassMajor"))
	return rv
}

// Search for a service record containing the given UUID.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServiceRecord(for:)
func (b_ BluetoothDevice) GetServiceRecordForUUID(sdpUUID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getServiceRecordForUUID:"), sdpUUID)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServices
func (b_ BluetoothDevice) GetServices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getServices"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeAudioGatewayDriverID
func (b_ BluetoothDevice) HandsFreeAudioGatewayDriverID() string {
	rv := objc.Send[string](b_.ID, objc.Sel("handsFreeAudioGatewayDriverID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeAudioGatewayServiceRecord()
func (b_ BluetoothDevice) HandsFreeAudioGatewayServiceRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("handsFreeAudioGatewayServiceRecord"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeDeviceDriverID
func (b_ BluetoothDevice) HandsFreeDeviceDriverID() string {
	rv := objc.Send[string](b_.ID, objc.Sel("handsFreeDeviceDriverID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeDeviceServiceRecord()
func (b_ BluetoothDevice) HandsFreeDeviceServiceRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("handsFreeDeviceServiceRecord"))
	return rv
}

// Indicates whether a baseband connection to the device exists.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isConnected()
func (b_ BluetoothDevice) IsConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}

// Reports whether the target device is a favorite for the user.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isFavorite()
func (b_ BluetoothDevice) IsFavorite() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isFavorite"))
	return rv
}

// Returns TRUE if the device connection was generated by the remote host.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isIncoming()
func (b_ BluetoothDevice) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}

// Returns whether the target device is paired.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isPaired()
func (b_ BluetoothDevice) IsPaired() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPaired"))
	return rv
}

// Create a baseband connection to the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection()
func (b_ BluetoothDevice) OpenConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openConnection"))
	return rv
}

// Create a baseband connection to the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection(_:withPageTimeout:authenticationRequired:)
func (b_ BluetoothDevice) OpenConnectionWithPageTimeoutAuthenticationRequired(target objc.ID, pageTimeoutValue unsafe.Pointer, authenticationRequired bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openConnection:withPageTimeout:authenticationRequired:"), target, pageTimeoutValue, authenticationRequired)
	return rv
}

// Opens a new L2CAP channel to the target device. Returns immedialty after starting the opening process.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannel:findExisting:newChannel:
func (b_ BluetoothDevice) OpenL2CAPChannelFindExistingNewChannel(psm unsafe.Pointer, findExisting bool, newChannel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannel:findExisting:newChannel:"), psm, findExisting, newChannel)
	return rv
}

// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}

// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelConfiguration objc.ID, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}

// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}

// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel unsafe.Pointer, psm unsafe.Pointer, channelConfiguration objc.ID, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}

// Opens a new RFCOMM channel to the target device. Returns only once the channel is open or failed to open.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannel:channel:
func (b_ BluetoothDevice) OpenRFCOMMChannelChannel(channelID unsafe.Pointer, rfcommChannel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openRFCOMMChannel:channel:"), channelID, rfcommChannel)
	return rv
}

// Opens a new RFCOMM channel to the target device. Returns immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelAsync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel unsafe.Pointer, channelID unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openRFCOMMChannelAsync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}

// Opens a new RFCOMM channel to the target device. Returns only once the channel is open or failed to open.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelSync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel unsafe.Pointer, channelID unsafe.Pointer, channelDelegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("openRFCOMMChannelSync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}

// Performs an SDP query on the target device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:)
func (b_ BluetoothDevice) PerformSDPQuery(target objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("performSDPQuery:"), target)
	return rv
}

// Performs an SDP query on the target device with the specified service UUIDs.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:uuids:)
func (b_ BluetoothDevice) PerformSDPQueryUuids(target objc.ID, uuidArray objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("performSDPQuery:uuids:"), target, uuidArray)
	return rv
}

// Get the raw RSSI device (if connected).
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rawRSSI()
func (b_ BluetoothDevice) RawRSSI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("rawRSSI"))
	return rv
}

// Returns the date/time of the most recent access of the target device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentAccessDate()
func (b_ BluetoothDevice) RecentAccessDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("recentAccessDate"))
	return rv
}

// Allows a client to register for device disconnect notification.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forDisconnectNotification:selector:)
func (b_ BluetoothDevice) RegisterForDisconnectNotificationSelector(observer objc.ID, inSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("registerForDisconnectNotification:selector:"), observer, inSelector)
	return rv
}

// Issues a remote name request to the target device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:)
func (b_ BluetoothDevice) RemoteNameRequest(target objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("remoteNameRequest:"), target)
	return rv
}

// Issues a remote name request to the target device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:withPageTimeout:)
func (b_ BluetoothDevice) RemoteNameRequestWithPageTimeout(target objc.ID, pageTimeoutValue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("remoteNameRequest:withPageTimeout:"), target, pageTimeoutValue)
	return rv
}

// Removes the target device from the user’s favorite devices list.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/removeFromFavorites()
func (b_ BluetoothDevice) RemoveFromFavorites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("removeFromFavorites"))
	return rv
}

// Requests that the existing baseband connection be authenticated.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/requestAuthentication()
func (b_ BluetoothDevice) RequestAuthentication() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("requestAuthentication"))
	return rv
}

// Get the RSSI device (if connected), above or below the golden range. If the RSSI is within the golden range, a value of 0 is returned. For the actual RSSI value, use getRawRSSI. For more information, see the Bluetooth 4.0 Core Specification.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rssi()
func (b_ BluetoothDevice) RSSI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("RSSI"))
	return rv
}

// Send an echo request over the L2CAP connection to a remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/sendL2CAPEchoRequest(_:length:)
func (b_ BluetoothDevice) SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sendL2CAPEchoRequest:length:"), data, length)
	return rv
}

// Sets the connection supervision timeout.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/setSupervisionTimeout(_:)
func (b_ BluetoothDevice) SetSupervisionTimeout(timeout unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("setSupervisionTimeout:"), timeout)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) IsHandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeAudioGateway"))
	return rv
}


// SetIsHandsFreeAudioGateway sets the value of the isHandsFreeAudioGateway property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) SetIsHandsFreeAudioGateway(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeAudioGateway:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) IsHandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeDevice"))
	return rv
}


// SetIsHandsFreeDevice sets the value of the isHandsFreeDevice property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) SetIsHandsFreeDevice(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeDevice:"), value)
}

// Get a string representation of the Bluetooth device address for the target device. The format of the string is the same as returned by IOBluetoothNSStringFromDeviceAddress().
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addressString
func (b_ BluetoothDevice) AddressString() string {
	rv := objc.Send[string](b_.ID, objc.Sel("addressString"))
	return rv
}

// Gets the full class of device value for the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/classOfDevice
func (b_ BluetoothDevice) ClassOfDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("classOfDevice"))
	return rv
}

// Get the connection handle for the baseband connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/connectionHandle
func (b_ BluetoothDevice) ConnectionHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("connectionHandle"))
	return rv
}

// Get the major device class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMajor
func (b_ BluetoothDevice) DeviceClassMajor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("deviceClassMajor"))
	return rv
}

// Get the minor service class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMinor
func (b_ BluetoothDevice) DeviceClassMinor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("deviceClassMinor"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeAudioGateway
func (b_ BluetoothDevice) HandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeAudioGateway"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeDevice
func (b_ BluetoothDevice) HandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeDevice"))
	return rv
}

// Get the date/time of the last successful remote name request.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/lastNameUpdate
func (b_ BluetoothDevice) LastNameUpdate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("lastNameUpdate"))
	return rv
}

// Get the human readable name of the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/name
func (b_ BluetoothDevice) Name() string {
	rv := objc.Send[string](b_.ID, objc.Sel("name"))
	return rv
}

// Get the human readable name of the remote device. If the name is not present, it will return a string containing the device’s address.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/nameOrAddress
func (b_ BluetoothDevice) NameOrAddress() string {
	rv := objc.Send[string](b_.ID, objc.Sel("nameOrAddress"))
	return rv
}

// Get the major service class of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/serviceClassMajor
func (b_ BluetoothDevice) ServiceClassMajor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("serviceClassMajor"))
	return rv
}

// Gets an array of service records for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/services
func (b_ BluetoothDevice) Services() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("services"))
	return rv
}


