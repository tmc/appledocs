// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class DDDevice */


/* debug [class_header]: Header for DDDevice */
// The class instance for the [DDDevice] class.
var (
	DDDeviceClass     _DDDeviceClass
	DDDeviceClassOnce sync.Once
)

func getDDDeviceClass() _DDDeviceClass {
	DDDeviceClassOnce.Do(func() {
		DDDeviceClass = _DDDeviceClass{objc.GetClass("DDDevice")}
	})
	return DDDeviceClass
}

type _DDDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDDevice */
// An interface definition for the [DDDevice] class.
type IDDDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DDDevice */
	// properties:
	BluetoothIdentifier() foundation.UUID
	SetBluetoothIdentifier(value foundation.UUID)
	Category() DDDeviceCategory
	SetCategory(value DDDeviceCategory)
	DeviceSupports() DDDeviceSupports
	SetDeviceSupports(value DDDeviceSupports)
	DisplayImageName() objc.IObject /* cross-framework: NSString */
	SetDisplayImageName(value objc.IObject /* cross-framework: NSString */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	MediaContentSubtitle() objc.IObject /* cross-framework: NSString */
	SetMediaContentSubtitle(value objc.IObject /* cross-framework: NSString */)
	MediaContentTitle() objc.IObject /* cross-framework: NSString */
	SetMediaContentTitle(value objc.IObject /* cross-framework: NSString */)
	MediaPlaybackState() DDDeviceMediaPlaybackState
	SetMediaPlaybackState(value DDDeviceMediaPlaybackState)
	NetworkEndpoint() unsafe.Pointer
	SetNetworkEndpoint(value unsafe.Pointer)
	Protocol() DDDeviceProtocol
	SetProtocol(value DDDeviceProtocol)
	ProtocolType() uniformtypeidentifiers.UTType
	SetProtocolType(value uniformtypeidentifiers.UTType)
	SSID() objc.IObject /* cross-framework: NSString */
	SetSSID(value objc.IObject /* cross-framework: NSString */)
	State() DDDeviceState
	SetState(value DDDeviceState)
	SupportsGrouping() bool
	SetSupportsGrouping(value bool)
	TxtRecordData() objc.IObject /* cross-framework: NSData */
	SetTxtRecordData(value objc.IObject /* cross-framework: NSData */)
	Url() objc.IObject /* cross-framework: NSURL */
	SetUrl(value objc.IObject /* cross-framework: NSURL */)
	WifiAwareModelName() objc.IObject /* cross-framework: NSString */
	SetWifiAwareModelName(value objc.IObject /* cross-framework: NSString */)
	WifiAwareServiceName() objc.IObject /* cross-framework: NSString */
	SetWifiAwareServiceName(value objc.IObject /* cross-framework: NSString */)
	WifiAwareServiceRole() DDDeviceWiFiAwareServiceRole
	SetWifiAwareServiceRole(value DDDeviceWiFiAwareServiceRole)
	WifiAwareVendorName() objc.IObject /* cross-framework: NSString */
	SetWifiAwareVendorName(value objc.IObject /* cross-framework: NSString */)
	TxtRecord() unsafe.Pointer
	SetTxtRecord(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDDevice */
// Alloc allocates a new instance without initialization.
func (dc _DDDeviceClass) Alloc() DDDevice {
	rv := objc.Send[DDDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDDeviceClass) New() DDDevice {
	rv := objc.Send[DDDevice](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDDevice) Init() DDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDDevice) Autorelease() DDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDDevice creates a new DDDevice instance.
func NewDDDevice() DDDevice {
	return getDDDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDDevice */
// An object that describes a discovered device of interest.
//
// The extension creates an instance of this class for a discovered device of interest and passes it to the system for display in the device picker UI ( ). The extension discovers devices through either Core Bluetooth or the local network (that is, using ). For device discovery extensions of third-party media receivers, an instance of this class corresponds to the media receiver of interest. The extension reports the status of discovered devices to the system using the function, and it receives status updates about the device from the system by implementing .


// An object that describes a discovered device of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice
type DDDevice struct {
	objectivec.Object
}

// DDDeviceFrom constructs a [DDDevice] from an unsafe.Pointer.
//
// An object that describes a discovered device of interest.
func DDDeviceFrom(ptr unsafe.Pointer) DDDevice {
	return DDDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDDevice */

// Creates an object that describes a discovered device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/init(displayName:category:protocolType:identifier:)
func NewDDDeviceWithDisplayNameCategoryProtocolTypeIdentifier(displayName objc.IObject /* cross-framework: NSString */, category DDDeviceCategory, protocolType uniformtypeidentifiers.UTType, identifier objc.IObject /* cross-framework: NSString */) DDDevice {
	instance := getDDDeviceClass().Alloc()
	rv := objc.Send[DDDevice](instance.ID, objc.Sel("initWithDisplayName:category:protocolType:identifier:"), displayName, category, protocolType, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDDDeviceWithDisplayNameCategoryProtocolTypeIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDDevice */

// An identifier to communicate with the device through Bluetooth wireless technology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/bluetoothIdentifier
func (d_ DDDevice) BluetoothIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](d_.ID, objc.Sel("bluetoothIdentifier"))
	return rv
}/* debug [instance_properties/getter]: bluetoothIdentifier */


// An identifier to communicate with the device through Bluetooth wireless technology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/bluetoothIdentifier
func (d_ DDDevice) SetBluetoothIdentifier(value foundation.UUID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBluetoothIdentifier:"), value)
}/* debug [instance_properties/setter]: bluetoothIdentifier */


// An option that determies the icon that the picker UI displays for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/category-swift.property
func (d_ DDDevice) Category() DDDeviceCategory {
	rv := objc.Send[DDDeviceCategory](d_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// An option that determies the icon that the picker UI displays for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/category-swift.property
func (d_ DDDevice) SetCategory(value DDDeviceCategory) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/deviceSupports
func (d_ DDDevice) DeviceSupports() DDDeviceSupports {
	rv := objc.Send[DDDeviceSupports](d_.ID, objc.Sel("deviceSupports"))
	return rv
}/* debug [instance_properties/getter]: deviceSupports */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/deviceSupports
func (d_ DDDevice) SetDeviceSupports(value DDDeviceSupports) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDeviceSupports:"), value)
}/* debug [instance_properties/setter]: deviceSupports */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayImageName
func (d_ DDDevice) DisplayImageName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayImageName"))
	return rv
}/* debug [instance_properties/getter]: displayImageName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayImageName
func (d_ DDDevice) SetDisplayImageName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayImageName:"), value)
}/* debug [instance_properties/setter]: displayImageName */


// A name for the device to display to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayName
func (d_ DDDevice) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A name for the device to display to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayName
func (d_ DDDevice) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), value)
}/* debug [instance_properties/setter]: displayName */


// A unique identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/identifier
func (d_ DDDevice) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A unique identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/identifier
func (d_ DDDevice) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A subtitle for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentSubtitle
func (d_ DDDevice) MediaContentSubtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("mediaContentSubtitle"))
	return rv
}/* debug [instance_properties/getter]: mediaContentSubtitle */


// A subtitle for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentSubtitle
func (d_ DDDevice) SetMediaContentSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaContentSubtitle:"), value)
}/* debug [instance_properties/setter]: mediaContentSubtitle */


// A title for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentTitle
func (d_ DDDevice) MediaContentTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("mediaContentTitle"))
	return rv
}/* debug [instance_properties/getter]: mediaContentTitle */


// A title for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentTitle
func (d_ DDDevice) SetMediaContentTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaContentTitle:"), value)
}/* debug [instance_properties/setter]: mediaContentTitle */


// A playback status for the device’s current media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaPlaybackState-swift.property
func (d_ DDDevice) MediaPlaybackState() DDDeviceMediaPlaybackState {
	rv := objc.Send[DDDeviceMediaPlaybackState](d_.ID, objc.Sel("mediaPlaybackState"))
	return rv
}/* debug [instance_properties/getter]: mediaPlaybackState */


// A playback status for the device’s current media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaPlaybackState-swift.property
func (d_ DDDevice) SetMediaPlaybackState(value DDDeviceMediaPlaybackState) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaPlaybackState:"), value)
}/* debug [instance_properties/setter]: mediaPlaybackState */


// An object that describes a local-network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/networkEndpoint-7excg
func (d_ DDDevice) NetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("networkEndpoint"))
	return rv
}/* debug [instance_properties/getter]: networkEndpoint */


// An object that describes a local-network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/networkEndpoint-7excg
func (d_ DDDevice) SetNetworkEndpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNetworkEndpoint:"), value)
}/* debug [instance_properties/setter]: networkEndpoint */


// The manner in which the system applies your app’s device discovery extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocol-swift.property
func (d_ DDDevice) Protocol() DDDeviceProtocol {
	rv := objc.Send[DDDeviceProtocol](d_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// The manner in which the system applies your app’s device discovery extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocol-swift.property
func (d_ DDDevice) SetProtocol(value DDDeviceProtocol) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setProtocol:"), value)
}/* debug [instance_properties/setter]: protocol */


// A custom universal type that describes the device’s manner of communication with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocolType
func (d_ DDDevice) ProtocolType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](d_.ID, objc.Sel("protocolType"))
	return rv
}/* debug [instance_properties/getter]: protocolType */


// A custom universal type that describes the device’s manner of communication with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocolType
func (d_ DDDevice) SetProtocolType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setProtocolType:"), value)
}/* debug [instance_properties/setter]: protocolType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/ssid
func (d_ DDDevice) SSID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("SSID"))
	return rv
}/* debug [instance_properties/getter]: SSID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/ssid
func (d_ DDDevice) SetSSID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSSID:"), value)
}/* debug [instance_properties/setter]: SSID */


// A state that represents the level of user interaction with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/state
func (d_ DDDevice) State() DDDeviceState {
	rv := objc.Send[DDDeviceState](d_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A state that represents the level of user interaction with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/state
func (d_ DDDevice) SetState(value DDDeviceState) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// A Boolean value that indicates whether to group the device with others in the AirPlay UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/supportsGrouping
func (d_ DDDevice) SupportsGrouping() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("supportsGrouping"))
	return rv
}/* debug [instance_properties/getter]: supportsGrouping */


// A Boolean value that indicates whether to group the device with others in the AirPlay UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/supportsGrouping
func (d_ DDDevice) SetSupportsGrouping(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportsGrouping:"), value)
}/* debug [instance_properties/setter]: supportsGrouping */


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/txtRecordData
func (d_ DDDevice) TxtRecordData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("txtRecordData"))
	return rv
}/* debug [instance_properties/getter]: txtRecordData */


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/txtRecordData
func (d_ DDDevice) SetTxtRecordData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTxtRecordData:"), value)
}/* debug [instance_properties/setter]: txtRecordData */


// A resource locator for the simple service discovery protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/url
func (d_ DDDevice) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A resource locator for the simple service discovery protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/url
func (d_ DDDevice) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */


// Device’s Wi-Fi Aware model name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareModelName
func (d_ DDDevice) WifiAwareModelName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareModelName"))
	return rv
}/* debug [instance_properties/getter]: wifiAwareModelName */


// Device’s Wi-Fi Aware model name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareModelName
func (d_ DDDevice) SetWifiAwareModelName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareModelName:"), value)
}/* debug [instance_properties/setter]: wifiAwareModelName */


// Device’s Wi-Fi Aware’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceName
func (d_ DDDevice) WifiAwareServiceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareServiceName"))
	return rv
}/* debug [instance_properties/getter]: wifiAwareServiceName */


// Device’s Wi-Fi Aware’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceName
func (d_ DDDevice) SetWifiAwareServiceName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareServiceName:"), value)
}/* debug [instance_properties/setter]: wifiAwareServiceName */


// Device’s Wi-Fi Aware’s service. Default is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceRole-swift.property
func (d_ DDDevice) WifiAwareServiceRole() DDDeviceWiFiAwareServiceRole {
	rv := objc.Send[DDDeviceWiFiAwareServiceRole](d_.ID, objc.Sel("wifiAwareServiceRole"))
	return rv
}/* debug [instance_properties/getter]: wifiAwareServiceRole */


// Device’s Wi-Fi Aware’s service. Default is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceRole-swift.property
func (d_ DDDevice) SetWifiAwareServiceRole(value DDDeviceWiFiAwareServiceRole) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareServiceRole:"), value)
}/* debug [instance_properties/setter]: wifiAwareServiceRole */


// Device’s Wi-Fi Aware vendor name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareVendorName
func (d_ DDDevice) WifiAwareVendorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareVendorName"))
	return rv
}/* debug [instance_properties/getter]: wifiAwareVendorName */


// Device’s Wi-Fi Aware vendor name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareVendorName
func (d_ DDDevice) SetWifiAwareVendorName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareVendorName:"), value)
}/* debug [instance_properties/setter]: wifiAwareVendorName */


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) TxtRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("txtRecord"))
	return rv
}/* debug [instance_properties/getter]: txtRecord */


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) SetTxtRecord(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTxtRecord:"), value)
}/* debug [instance_properties/setter]: txtRecord */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDDevice */


