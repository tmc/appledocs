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
	AddressString() string /* primitive/slice/pointer. */
	ClassOfDevice() BluetoothClassOfDevice /* typedef */
	ConnectionHandle() BluetoothConnectionHandle /* typedef */
	DeviceClassMajor() BluetoothDeviceClassMajor /* typedef */
	DeviceClassMinor() BluetoothDeviceClassMinor /* typedef */
	HandsFreeAudioGateway() bool /* primitive/slice/pointer. */
	HandsFreeDevice() bool /* primitive/slice/pointer. */
	LastNameUpdate() foundation.objc.IObject /* cross-framework: NSDate */
	Name() string /* primitive/slice/pointer. */
	NameOrAddress() string /* primitive/slice/pointer. */
	ServiceClassMajor() BluetoothServiceClassMajor /* typedef */
	Services() objc.ID
	IsHandsFreeAudioGateway() bool /* primitive/slice/pointer. */
	SetIsHandsFreeAudioGateway(value bool /* primitive/slice/pointer. */)
	IsHandsFreeDevice() bool /* primitive/slice/pointer. */
	SetIsHandsFreeDevice(value bool /* primitive/slice/pointer. */)
	// methods:
	AddToFavorites() Return /* not a class type */
	CloseConnection() Return /* not a class type */
	GetAddress() unsafe.Pointer
	GetClockOffset() BluetoothClockOffset /* typedef */
	GetEncryptionMode() BluetoothHCIEncryptionMode /* typedef */
	GetLastInquiryUpdate() objc.IObject /* cross-framework: Date */
	GetLastServicesUpdate() objc.IObject /* cross-framework: Date */
	GetLinkType() BluetoothLinkType /* typedef */
	GetPageScanMode() BluetoothPageScanMode /* typedef */
	GetPageScanPeriodMode() BluetoothPageScanPeriodMode /* typedef */
	GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode /* typedef */
	GetServiceRecordForUUID(sdpUUID BluetoothSDPUUID /* already interface */) IBluetoothSDPServiceRecord
	HandsFreeAudioGatewayServiceRecord() IBluetoothSDPServiceRecord
	HandsFreeDeviceServiceRecord() IBluetoothSDPServiceRecord
	IsConnected() bool /* primitive/slice/pointer. */
	IsFavorite() bool /* primitive/slice/pointer. */
	IsIncoming() bool /* primitive/slice/pointer. */
	IsPaired() bool /* primitive/slice/pointer. */
	OpenConnection() Return /* not a class type */
	OpenConnectionWithPageTimeoutAuthenticationRequired(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */, authenticationRequired bool /* primitive/slice/pointer. */) Return /* not a class type */
	OpenL2CAPChannelAsyncWithPSMDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */
	OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) Return /* not a class type */
	OpenL2CAPChannelSyncWithPSMDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */
	OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) Return /* not a class type */
	OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel BluetoothRFCOMMChannel /* already interface */, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */
	OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel BluetoothRFCOMMChannel /* already interface */, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */
	PerformSDPQuery(target objectivec.IObject) Return /* not a class type */
	PerformSDPQueryUuids(target objectivec.IObject, uuidArray objectivec.IObject) Return /* not a class type */
	RawRSSI() BluetoothHCIRSSIValue /* typedef */
	RecentAccessDate() objc.IObject /* cross-framework: Date */
	RegisterForDisconnectNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification
	RemoteNameRequest(target objectivec.IObject) Return /* not a class type */
	RemoteNameRequestWithPageTimeout(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */) Return /* not a class type */
	RemoveFromFavorites() Return /* not a class type */
	RequestAuthentication() Return /* not a class type */
	RSSI() BluetoothHCIRSSIValue /* typedef */
	SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) Return /* not a class type */
	SetSupervisionTimeout(timeout unsafe.Pointer) Return /* not a class type */
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
func NewBluetoothDeviceWithAddressString(address string /* primitive/slice/pointer. */) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}



// Gets an array of the user’s favorite devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/favoriteDevices()
func (bc _BluetoothDeviceClass) FavoriteDevices() objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](objc.ID(bc.class), objc.Sel("favoriteDevices"))
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
func (bc _BluetoothDeviceClass) DeviceWithAddressString(address string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddressString:"), objc.String(address))
	return rv
}


// Gets an array of all of the paired devices on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/pairedDevices()
func (bc _BluetoothDeviceClass) PairedDevices() objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](objc.ID(bc.class), objc.Sel("pairedDevices"))
	return rv
}


// Gets an array of recently used Bluetooth devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentDevices(_:)
func (bc _BluetoothDeviceClass) RecentDevices(numDevices unsafe.Pointer) objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](objc.ID(bc.class), objc.Sel("recentDevices:"), numDevices)
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
func (bc _BluetoothDeviceClass) WithDeviceRef(deviceRef objc.IObject /* cross-framework BluetoothDeviceRef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDeviceRef:"), deviceRef)
	return rv
}


// Adds the target device to the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addToFavorites()
func (b_ BluetoothDevice) AddToFavorites() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("addToFavorites"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/awakeAfter(using:)
func (b_ BluetoothDevice) AwakeAfterUsingCoder(coder Coder /* not a class type */) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}


// Close down the baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/closeConnection()
func (b_ BluetoothDevice) CloseConnection() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("closeConnection"))
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
func (b_ BluetoothDevice) GetClockOffset() BluetoothClockOffset /* typedef */ {
	rv := objc.Send[BluetoothClockOffset](b_.ID, objc.Sel("getClockOffset"))
	return rv
}


// Get the encryption mode for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getEncryptionMode()
func (b_ BluetoothDevice) GetEncryptionMode() BluetoothHCIEncryptionMode /* typedef */ {
	rv := objc.Send[BluetoothHCIEncryptionMode](b_.ID, objc.Sel("getEncryptionMode"))
	return rv
}


// Get the date/time of the last time the device was returned during an inquiry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastInquiryUpdate()
func (b_ BluetoothDevice) GetLastInquiryUpdate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[Date](b_.ID, objc.Sel("getLastInquiryUpdate"))
	return rv
}


// Get the date/time of the last SDP query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastServicesUpdate()
func (b_ BluetoothDevice) GetLastServicesUpdate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[Date](b_.ID, objc.Sel("getLastServicesUpdate"))
	return rv
}


// Get the link type for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLinkType()
func (b_ BluetoothDevice) GetLinkType() BluetoothLinkType /* typedef */ {
	rv := objc.Send[BluetoothLinkType](b_.ID, objc.Sel("getLinkType"))
	return rv
}


// Get the page scan mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanMode()
func (b_ BluetoothDevice) GetPageScanMode() BluetoothPageScanMode /* typedef */ {
	rv := objc.Send[BluetoothPageScanMode](b_.ID, objc.Sel("getPageScanMode"))
	return rv
}


// Get the value of the page scan period mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanPeriodMode()
func (b_ BluetoothDevice) GetPageScanPeriodMode() BluetoothPageScanPeriodMode /* typedef */ {
	rv := objc.Send[BluetoothPageScanPeriodMode](b_.ID, objc.Sel("getPageScanPeriodMode"))
	return rv
}


// Get the value of the page scan repetition mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanRepetitionMode()
func (b_ BluetoothDevice) GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode /* typedef */ {
	rv := objc.Send[BluetoothPageScanRepetitionMode](b_.ID, objc.Sel("getPageScanRepetitionMode"))
	return rv
}


// Search for a service record containing the given UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServiceRecord(for:)
func (b_ BluetoothDevice) GetServiceRecordForUUID(sdpUUID BluetoothSDPUUID /* already interface */) IBluetoothSDPServiceRecord {
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
func (b_ BluetoothDevice) IsConnected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}


// Reports whether the target device is a favorite for the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isFavorite()
func (b_ BluetoothDevice) IsFavorite() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isFavorite"))
	return rv
}


// Returns TRUE if the device connection was generated by the remote host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isIncoming()
func (b_ BluetoothDevice) IsIncoming() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}


// Returns whether the target device is paired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isPaired()
func (b_ BluetoothDevice) IsPaired() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPaired"))
	return rv
}


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection()
func (b_ BluetoothDevice) OpenConnection() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openConnection"))
	return rv
}


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection(_:withPageTimeout:authenticationRequired:)
func (b_ BluetoothDevice) OpenConnectionWithPageTimeoutAuthenticationRequired(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */, authenticationRequired bool /* primitive/slice/pointer. */) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openConnection:withPageTimeout:authenticationRequired:"), target, pageTimeoutValue, authenticationRequired)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel BluetoothL2CAPChannel /* already interface */, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objectivec.IObject, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}


// Opens a new RFCOMM channel to the target device. Returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelAsync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel BluetoothRFCOMMChannel /* already interface */, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openRFCOMMChannelAsync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}


// Opens a new RFCOMM channel to the target device. Returns only once the channel is open or failed to open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelSync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel BluetoothRFCOMMChannel /* already interface */, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("openRFCOMMChannelSync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}


// Performs an SDP query on the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:)
func (b_ BluetoothDevice) PerformSDPQuery(target objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("performSDPQuery:"), target)
	return rv
}


// Performs an SDP query on the target device with the specified service UUIDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:uuids:)
func (b_ BluetoothDevice) PerformSDPQueryUuids(target objectivec.IObject, uuidArray objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("performSDPQuery:uuids:"), target, uuidArray)
	return rv
}


// Get the raw RSSI device (if connected).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rawRSSI()
func (b_ BluetoothDevice) RawRSSI() BluetoothHCIRSSIValue /* typedef */ {
	rv := objc.Send[BluetoothHCIRSSIValue](b_.ID, objc.Sel("rawRSSI"))
	return rv
}


// Returns the date/time of the most recent access of the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentAccessDate()
func (b_ BluetoothDevice) RecentAccessDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[Date](b_.ID, objc.Sel("recentAccessDate"))
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
func (b_ BluetoothDevice) RemoteNameRequest(target objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("remoteNameRequest:"), target)
	return rv
}


// Issues a remote name request to the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:withPageTimeout:)
func (b_ BluetoothDevice) RemoteNameRequestWithPageTimeout(target objectivec.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("remoteNameRequest:withPageTimeout:"), target, pageTimeoutValue)
	return rv
}


// Removes the target device from the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/removeFromFavorites()
func (b_ BluetoothDevice) RemoveFromFavorites() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("removeFromFavorites"))
	return rv
}


// Requests that the existing baseband connection be authenticated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/requestAuthentication()
func (b_ BluetoothDevice) RequestAuthentication() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("requestAuthentication"))
	return rv
}


// Get the RSSI device (if connected), above or below the golden range. If the RSSI is within the golden range, a value of 0 is returned. For the actual RSSI value, use getRawRSSI. For more information, see the Bluetooth 4.0 Core Specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rssi()
func (b_ BluetoothDevice) RSSI() BluetoothHCIRSSIValue /* typedef */ {
	rv := objc.Send[BluetoothHCIRSSIValue](b_.ID, objc.Sel("RSSI"))
	return rv
}


// Send an echo request over the L2CAP connection to a remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/sendL2CAPEchoRequest(_:length:)
func (b_ BluetoothDevice) SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("sendL2CAPEchoRequest:length:"), data, length)
	return rv
}


// Sets the connection supervision timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/setSupervisionTimeout(_:)
func (b_ BluetoothDevice) SetSupervisionTimeout(timeout unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("setSupervisionTimeout:"), timeout)
	return rv
}


// Get a string representation of the Bluetooth device address for the target device. The format of the string is the same as returned by IOBluetoothNSStringFromDeviceAddress().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addressString
func (b_ BluetoothDevice) AddressString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("addressString"))
	return rv
}


// Gets the full class of device value for the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/classOfDevice
func (b_ BluetoothDevice) ClassOfDevice() BluetoothClassOfDevice /* typedef */ {
	rv := objc.Send[BluetoothClassOfDevice](b_.ID, objc.Sel("classOfDevice"))
	return rv
}


// Get the connection handle for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/connectionHandle
func (b_ BluetoothDevice) ConnectionHandle() BluetoothConnectionHandle /* typedef */ {
	rv := objc.Send[BluetoothConnectionHandle](b_.ID, objc.Sel("connectionHandle"))
	return rv
}


// Get the major device class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMajor
func (b_ BluetoothDevice) DeviceClassMajor() BluetoothDeviceClassMajor /* typedef */ {
	rv := objc.Send[BluetoothDeviceClassMajor](b_.ID, objc.Sel("deviceClassMajor"))
	return rv
}


// Get the minor service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMinor
func (b_ BluetoothDevice) DeviceClassMinor() BluetoothDeviceClassMinor /* typedef */ {
	rv := objc.Send[BluetoothDeviceClassMinor](b_.ID, objc.Sel("deviceClassMinor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeAudioGateway
func (b_ BluetoothDevice) HandsFreeAudioGateway() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeAudioGateway"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeDevice
func (b_ BluetoothDevice) HandsFreeDevice() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeDevice"))
	return rv
}


// Get the date/time of the last successful remote name request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/lastNameUpdate
func (b_ BluetoothDevice) LastNameUpdate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](b_.ID, objc.Sel("lastNameUpdate"))
	return rv
}


// Get the human readable name of the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/name
func (b_ BluetoothDevice) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("name"))
	return rv
}


// Get the human readable name of the remote device. If the name is not present, it will return a string containing the device’s address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/nameOrAddress
func (b_ BluetoothDevice) NameOrAddress() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("nameOrAddress"))
	return rv
}


// Get the major service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/serviceClassMajor
func (b_ BluetoothDevice) ServiceClassMajor() BluetoothServiceClassMajor /* typedef */ {
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
func (b_ BluetoothDevice) IsHandsFreeAudioGateway() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeAudioGateway"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) SetIsHandsFreeAudioGateway(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeAudioGateway:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) IsHandsFreeDevice() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeDevice"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) SetIsHandsFreeDevice(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeDevice:"), value)
}


