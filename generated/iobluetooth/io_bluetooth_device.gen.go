// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothDevice */


/* debug [class_header]: Header for IOBluetoothDevice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothDevice */
// An interface definition for the [BluetoothDevice] class.
type IBluetoothDevice interface {
	IBluetoothObject
	
/* debug [class_interface_properties]: Properties for BluetoothDevice */
	// properties:
	AddressString() objc.IObject /* cross-framework: NSString */
	ClassOfDevice() BluetoothClassOfDevice /* typedef */
	ConnectionHandle() BluetoothConnectionHandle /* typedef */
	DeviceClassMajor() BluetoothDeviceClassMajor /* typedef */
	DeviceClassMinor() BluetoothDeviceClassMinor /* typedef */
	HandsFreeAudioGateway() bool
	HandsFreeDevice() bool
	LastNameUpdate() objc.IObject /* cross-framework: NSDate */
	Name() objc.IObject /* cross-framework: NSString */
	NameOrAddress() objc.IObject /* cross-framework: NSString */
	ServiceClassMajor() BluetoothServiceClassMajor /* typedef */
	Services() objc.IObject /* cross-framework: NSArray */
	IsHandsFreeAudioGateway() bool
	SetIsHandsFreeAudioGateway(value bool)
	IsHandsFreeDevice() bool
	SetIsHandsFreeDevice(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothDevice */
	// methods:
	AddToFavorites() int
	CloseConnection() int
	GetAddress() objc.IObject /* cross-framework: BluetoothDeviceAddress */
	GetClockOffset() BluetoothClockOffset /* typedef */
	GetEncryptionMode() BluetoothHCIEncryptionMode /* typedef */
	GetLastInquiryUpdate() foundation.Date
	GetLastServicesUpdate() foundation.Date
	GetLinkType() BluetoothLinkType /* typedef */
	GetPageScanMode() BluetoothPageScanMode /* typedef */
	GetPageScanPeriodMode() BluetoothPageScanPeriodMode /* typedef */
	GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode /* typedef */
	GetServiceRecordForUUID(sdpUUID IOBluetoothSDPUUID) IBluetoothSDPServiceRecord
	HandsFreeAudioGatewayServiceRecord() IBluetoothSDPServiceRecord
	HandsFreeDeviceServiceRecord() IBluetoothSDPServiceRecord
	IsConnected() bool
	IsFavorite() bool
	IsIncoming() bool
	IsPaired() bool
	OpenConnection() int
	OpenConnectionWithTarget(target objc.IObject) int
	OpenConnectionWithPageTimeoutAuthenticationRequired(target objc.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */, authenticationRequired bool) int
	OpenL2CAPChannelAsyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objc.IObject) int
	OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objc.IObject /* cross-framework: NSDictionary */, channelDelegate objc.IObject) int
	OpenL2CAPChannelSyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objc.IObject) int
	OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objc.IObject /* cross-framework: NSDictionary */, channelDelegate objc.IObject) int
	OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objc.IObject) int
	OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objc.IObject) int
	PerformSDPQuery(target objc.IObject) int
	PerformSDPQueryUuids(target objc.IObject, uuidArray objc.IObject /* cross-framework: NSArray */) int
	RawRSSI() BluetoothHCIRSSIValue /* typedef */
	RecentAccessDate() foundation.Date
	RegisterForDisconnectNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification
	RemoteNameRequest(target objc.IObject) int
	RemoteNameRequestWithPageTimeout(target objc.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */) int
	RemoveFromFavorites() int
	RequestAuthentication() int
	RSSI() BluetoothHCIRSSIValue /* typedef */
	SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) int
	SetSupervisionTimeout(timeout unsafe.Pointer) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothDevice */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceClass) Alloc() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothDevice */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothDevice */

// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func NewBluetoothDeviceWithAddress(address objc.IObject /* cross-framework: BluetoothDeviceAddress */) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddress:"), address)
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothDeviceWithAddress */


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func NewBluetoothDeviceWithAddressString(address objc.IObject /* cross-framework: NSString */) BluetoothDevice {
	rv := objc.Send[BluetoothDevice](objc.ID(getBluetoothDeviceClass().class), objc.Sel("deviceWithAddressString:"), address)
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothDeviceWithAddressString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothDevice */

// Gets an array of the user’s favorite devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/favoriteDevices()
func (bc _BluetoothDeviceClass) FavoriteDevices() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("favoriteDevices"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FavoriteDevices) */


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(address:)
func (bc _BluetoothDeviceClass) DeviceWithAddress(address objc.IObject /* cross-framework: BluetoothDeviceAddress */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddress:"), address)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithAddress) */


// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/init(addressString:)
func (bc _BluetoothDeviceClass) DeviceWithAddressString(address objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("deviceWithAddressString:"), address)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithAddressString) */


// Gets an array of all of the paired devices on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/pairedDevices()
func (bc _BluetoothDeviceClass) PairedDevices() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("pairedDevices"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PairedDevices) */


// Gets an array of recently used Bluetooth devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentDevices(_:)
func (bc _BluetoothDeviceClass) RecentDevices(numDevices unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(bc.class), objc.Sel("recentDevices:"), numDevices)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecentDevices) */


// Allows a client to register for device connect notifications for any connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forConnectNotifications:selector:)
func (bc _BluetoothDeviceClass) RegisterForConnectNotificationsSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForConnectNotifications:selector:"), observer, inSelector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForConnectNotificationsSelector) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withAddress:
func (bc _BluetoothDeviceClass) WithAddress(address objc.IObject /* cross-framework: BluetoothDeviceAddress */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withAddress:"), address)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithAddress) */


// Method call to convert an IOBluetoothDeviceRef into an IOBluetoothDevice *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/withDeviceRef:
func (bc _BluetoothDeviceClass) WithDeviceRef(deviceRef BluetoothDeviceRef /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDeviceRef:"), deviceRef)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithDeviceRef) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothDevice */

// Adds the target device to the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addToFavorites()
func (b_ BluetoothDevice) AddToFavorites() int {
	rv := objc.Send[int](b_.ID, objc.Sel("addToFavorites"))
	return rv
}/* debug [instance_methods/method]: AddToFavorites */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/awakeAfter(using:)
func (b_ BluetoothDevice) AwakeAfterUsingCoder(coder foundation.Coder) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}/* debug [instance_methods/method]: AwakeAfterUsingCoder */


// Close down the baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/closeConnection()
func (b_ BluetoothDevice) CloseConnection() int {
	rv := objc.Send[int](b_.ID, objc.Sel("closeConnection"))
	return rv
}/* debug [instance_methods/method]: CloseConnection */


// Get the Bluetooth device address for the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getAddress()
func (b_ BluetoothDevice) GetAddress() objc.IObject /* cross-framework: BluetoothDeviceAddress */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("getAddress"))
	return rv
}/* debug [instance_methods/method]: GetAddress */


// Get the clock offset value of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getClockOffset()
func (b_ BluetoothDevice) GetClockOffset() BluetoothClockOffset /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("getClockOffset"))
	return rv
}/* debug [instance_methods/method]: GetClockOffset */


// Get the encryption mode for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getEncryptionMode()
func (b_ BluetoothDevice) GetEncryptionMode() BluetoothHCIEncryptionMode /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getEncryptionMode"))
	return rv
}/* debug [instance_methods/method]: GetEncryptionMode */


// Get the date/time of the last time the device was returned during an inquiry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastInquiryUpdate()
func (b_ BluetoothDevice) GetLastInquiryUpdate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("getLastInquiryUpdate"))
	return rv
}/* debug [instance_methods/method]: GetLastInquiryUpdate */


// Get the date/time of the last SDP query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLastServicesUpdate()
func (b_ BluetoothDevice) GetLastServicesUpdate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("getLastServicesUpdate"))
	return rv
}/* debug [instance_methods/method]: GetLastServicesUpdate */


// Get the link type for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getLinkType()
func (b_ BluetoothDevice) GetLinkType() BluetoothLinkType /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getLinkType"))
	return rv
}/* debug [instance_methods/method]: GetLinkType */


// Get the page scan mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanMode()
func (b_ BluetoothDevice) GetPageScanMode() BluetoothPageScanMode /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getPageScanMode"))
	return rv
}/* debug [instance_methods/method]: GetPageScanMode */


// Get the value of the page scan period mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanPeriodMode()
func (b_ BluetoothDevice) GetPageScanPeriodMode() BluetoothPageScanPeriodMode /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getPageScanPeriodMode"))
	return rv
}/* debug [instance_methods/method]: GetPageScanPeriodMode */


// Get the value of the page scan repetition mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getPageScanRepetitionMode()
func (b_ BluetoothDevice) GetPageScanRepetitionMode() BluetoothPageScanRepetitionMode /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getPageScanRepetitionMode"))
	return rv
}/* debug [instance_methods/method]: GetPageScanRepetitionMode */


// Search for a service record containing the given UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/getServiceRecord(for:)
func (b_ BluetoothDevice) GetServiceRecordForUUID(sdpUUID IOBluetoothSDPUUID) IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("getServiceRecordForUUID:"), sdpUUID)
	return rv
}/* debug [instance_methods/method]: GetServiceRecordForUUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeAudioGatewayServiceRecord()
func (b_ BluetoothDevice) HandsFreeAudioGatewayServiceRecord() IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("handsFreeAudioGatewayServiceRecord"))
	return rv
}/* debug [instance_methods/method]: HandsFreeAudioGatewayServiceRecord */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/handsFreeDeviceServiceRecord()
func (b_ BluetoothDevice) HandsFreeDeviceServiceRecord() IBluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("handsFreeDeviceServiceRecord"))
	return rv
}/* debug [instance_methods/method]: HandsFreeDeviceServiceRecord */


// Indicates whether a baseband connection to the device exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isConnected()
func (b_ BluetoothDevice) IsConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}/* debug [instance_methods/method]: IsConnected */


// Reports whether the target device is a favorite for the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isFavorite()
func (b_ BluetoothDevice) IsFavorite() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isFavorite"))
	return rv
}/* debug [instance_methods/method]: IsFavorite */


// Returns TRUE if the device connection was generated by the remote host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isIncoming()
func (b_ BluetoothDevice) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}/* debug [instance_methods/method]: IsIncoming */


// Returns whether the target device is paired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isPaired()
func (b_ BluetoothDevice) IsPaired() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPaired"))
	return rv
}/* debug [instance_methods/method]: IsPaired */


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection()
func (b_ BluetoothDevice) OpenConnection() int {
	rv := objc.Send[int](b_.ID, objc.Sel("openConnection"))
	return rv
}/* debug [instance_methods/method]: OpenConnection */


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection(_:)
func (b_ BluetoothDevice) OpenConnectionWithTarget(target objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openConnection:"), target)
	return rv
}/* debug [instance_methods/method]: OpenConnectionWithTarget */


// Create a baseband connection to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openConnection(_:withPageTimeout:authenticationRequired:)
func (b_ BluetoothDevice) OpenConnectionWithPageTimeoutAuthenticationRequired(target objc.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */, authenticationRequired bool) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openConnection:withPageTimeout:authenticationRequired:"), target, pageTimeoutValue, authenticationRequired)
	return rv
}/* debug [instance_methods/method]: OpenConnectionWithPageTimeoutAuthenticationRequired */


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenL2CAPChannelAsyncWithPSMDelegate */


// Opens a new L2CAP channel to the target device. Returns immediately after starting the opening process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelAsync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objc.IObject /* cross-framework: NSDictionary */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openL2CAPChannelAsync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenL2CAPChannelAsyncWithPSMWithConfigurationDelegate */


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:delegate:"), newChannel, psm, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenL2CAPChannelSyncWithPSMDelegate */


// Opens a new L2CAP channel to the target device. Returns only after the channel is opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openL2CAPChannelSync(_:withPSM:withConfiguration:delegate:)
func (b_ BluetoothDevice) OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate(newChannel IOBluetoothL2CAPChannel, psm BluetoothL2CAPPSM /* typedef */, channelConfiguration objc.IObject /* cross-framework: NSDictionary */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openL2CAPChannelSync:withPSM:withConfiguration:delegate:"), newChannel, psm, channelConfiguration, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenL2CAPChannelSyncWithPSMWithConfigurationDelegate */


// Opens a new RFCOMM channel to the target device. Returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelAsync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelAsyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openRFCOMMChannelAsync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenRFCOMMChannelAsyncWithChannelIDDelegate */


// Opens a new RFCOMM channel to the target device. Returns only once the channel is open or failed to open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/openRFCOMMChannelSync(_:withChannelID:delegate:)
func (b_ BluetoothDevice) OpenRFCOMMChannelSyncWithChannelIDDelegate(rfcommChannel IOBluetoothRFCOMMChannel, channelID BluetoothRFCOMMChannelID /* typedef */, channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("openRFCOMMChannelSync:withChannelID:delegate:"), rfcommChannel, channelID, channelDelegate)
	return rv
}/* debug [instance_methods/method]: OpenRFCOMMChannelSyncWithChannelIDDelegate */


// Performs an SDP query on the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:)
func (b_ BluetoothDevice) PerformSDPQuery(target objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("performSDPQuery:"), target)
	return rv
}/* debug [instance_methods/method]: PerformSDPQuery */


// Performs an SDP query on the target device with the specified service UUIDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/performSDPQuery(_:uuids:)
func (b_ BluetoothDevice) PerformSDPQueryUuids(target objc.IObject, uuidArray objc.IObject /* cross-framework: NSArray */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("performSDPQuery:uuids:"), target, uuidArray)
	return rv
}/* debug [instance_methods/method]: PerformSDPQueryUuids */


// Get the raw RSSI device (if connected).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rawRSSI()
func (b_ BluetoothDevice) RawRSSI() BluetoothHCIRSSIValue /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("rawRSSI"))
	return rv
}/* debug [instance_methods/method]: RawRSSI */


// Returns the date/time of the most recent access of the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/recentAccessDate()
func (b_ BluetoothDevice) RecentAccessDate() foundation.Date {
	rv := objc.Send[foundation.Date](b_.ID, objc.Sel("recentAccessDate"))
	return rv
}/* debug [instance_methods/method]: RecentAccessDate */


// Allows a client to register for device disconnect notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/register(forDisconnectNotification:selector:)
func (b_ BluetoothDevice) RegisterForDisconnectNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("registerForDisconnectNotification:selector:"), observer, inSelector)
	return rv
}/* debug [instance_methods/method]: RegisterForDisconnectNotificationSelector */


// Issues a remote name request to the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:)
func (b_ BluetoothDevice) RemoteNameRequest(target objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("remoteNameRequest:"), target)
	return rv
}/* debug [instance_methods/method]: RemoteNameRequest */


// Issues a remote name request to the target device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/remoteNameRequest(_:withPageTimeout:)
func (b_ BluetoothDevice) RemoteNameRequestWithPageTimeout(target objc.IObject, pageTimeoutValue BluetoothHCIPageTimeout /* typedef */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("remoteNameRequest:withPageTimeout:"), target, pageTimeoutValue)
	return rv
}/* debug [instance_methods/method]: RemoteNameRequestWithPageTimeout */


// Removes the target device from the user’s favorite devices list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/removeFromFavorites()
func (b_ BluetoothDevice) RemoveFromFavorites() int {
	rv := objc.Send[int](b_.ID, objc.Sel("removeFromFavorites"))
	return rv
}/* debug [instance_methods/method]: RemoveFromFavorites */


// Requests that the existing baseband connection be authenticated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/requestAuthentication()
func (b_ BluetoothDevice) RequestAuthentication() int {
	rv := objc.Send[int](b_.ID, objc.Sel("requestAuthentication"))
	return rv
}/* debug [instance_methods/method]: RequestAuthentication */


// Get the RSSI device (if connected), above or below the golden range. If the RSSI is within the golden range, a value of 0 is returned. For the actual RSSI value, use getRawRSSI. For more information, see the Bluetooth 4.0 Core Specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/rssi()
func (b_ BluetoothDevice) RSSI() BluetoothHCIRSSIValue /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("RSSI"))
	return rv
}/* debug [instance_methods/method]: RSSI */


// Send an echo request over the L2CAP connection to a remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/sendL2CAPEchoRequest(_:length:)
func (b_ BluetoothDevice) SendL2CAPEchoRequestLength(data unsafe.Pointer, length unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("sendL2CAPEchoRequest:length:"), data, length)
	return rv
}/* debug [instance_methods/method]: SendL2CAPEchoRequestLength */


// Sets the connection supervision timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/setSupervisionTimeout(_:)
func (b_ BluetoothDevice) SetSupervisionTimeout(timeout unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setSupervisionTimeout:"), timeout)
	return rv
}/* debug [instance_methods/method]: SetSupervisionTimeout */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothDevice */

// Get a string representation of the Bluetooth device address for the target device. The format of the string is the same as returned by IOBluetoothNSStringFromDeviceAddress().
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/addressString
func (b_ BluetoothDevice) AddressString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("addressString"))
	return rv
}/* debug [instance_properties/getter]: addressString */


// Gets the full class of device value for the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/classOfDevice
func (b_ BluetoothDevice) ClassOfDevice() BluetoothClassOfDevice /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("classOfDevice"))
	return rv
}/* debug [instance_properties/getter]: classOfDevice */


// Get the connection handle for the baseband connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/connectionHandle
func (b_ BluetoothDevice) ConnectionHandle() BluetoothConnectionHandle /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("connectionHandle"))
	return rv
}/* debug [instance_properties/getter]: connectionHandle */


// Get the major device class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMajor
func (b_ BluetoothDevice) DeviceClassMajor() BluetoothDeviceClassMajor /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceClassMajor"))
	return rv
}/* debug [instance_properties/getter]: deviceClassMajor */


// Get the minor service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/deviceClassMinor
func (b_ BluetoothDevice) DeviceClassMinor() BluetoothDeviceClassMinor /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceClassMinor"))
	return rv
}/* debug [instance_properties/getter]: deviceClassMinor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeAudioGateway
func (b_ BluetoothDevice) HandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeAudioGateway"))
	return rv
}/* debug [instance_properties/getter]: handsFreeAudioGateway */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/isHandsFreeDevice
func (b_ BluetoothDevice) HandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("handsFreeDevice"))
	return rv
}/* debug [instance_properties/getter]: handsFreeDevice */


// Get the date/time of the last successful remote name request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/lastNameUpdate
func (b_ BluetoothDevice) LastNameUpdate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](b_.ID, objc.Sel("lastNameUpdate"))
	return rv
}/* debug [instance_properties/getter]: lastNameUpdate */


// Get the human readable name of the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/name
func (b_ BluetoothDevice) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Get the human readable name of the remote device. If the name is not present, it will return a string containing the device’s address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/nameOrAddress
func (b_ BluetoothDevice) NameOrAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("nameOrAddress"))
	return rv
}/* debug [instance_properties/getter]: nameOrAddress */


// Get the major service class of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/serviceClassMajor
func (b_ BluetoothDevice) ServiceClassMajor() BluetoothServiceClassMajor /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("serviceClassMajor"))
	return rv
}/* debug [instance_properties/getter]: serviceClassMajor */


// Gets an array of service records for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/services
func (b_ BluetoothDevice) Services() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](b_.ID, objc.Sel("services"))
	return rv
}/* debug [instance_properties/getter]: services */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) IsHandsFreeAudioGateway() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeAudioGateway"))
	return rv
}/* debug [instance_properties/getter]: isHandsFreeAudioGateway */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreeaudiogateway
func (b_ BluetoothDevice) SetIsHandsFreeAudioGateway(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeAudioGateway:"), value)
}/* debug [instance_properties/setter]: isHandsFreeAudioGateway */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) IsHandsFreeDevice() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isHandsFreeDevice"))
	return rv
}/* debug [instance_properties/getter]: isHandsFreeDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdevice/ishandsfreedevice
func (b_ BluetoothDevice) SetIsHandsFreeDevice(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsHandsFreeDevice:"), value)
}/* debug [instance_properties/setter]: isHandsFreeDevice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothDevice */


