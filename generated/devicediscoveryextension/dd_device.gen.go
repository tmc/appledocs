// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [DDDevice] class.
type IDDDevice interface {
	objectivec.IObject
	// properties:
	BluetoothIdentifier() objc.IObject /* cross-framework: UUID */
	SetBluetoothIdentifier(value objc.IObject /* cross-framework: UUID */)
	Category() DDDeviceCategory
	SetCategory(value DDDeviceCategory)
	DeviceSupports() DDDeviceSupports
	SetDeviceSupports(value DDDeviceSupports)
	DisplayImageName() string /* primitive/slice/pointer. */
	SetDisplayImageName(value string /* primitive/slice/pointer. */)
	DisplayName() string /* primitive/slice/pointer. */
	SetDisplayName(value string /* primitive/slice/pointer. */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	MediaContentSubtitle() string /* primitive/slice/pointer. */
	SetMediaContentSubtitle(value string /* primitive/slice/pointer. */)
	MediaContentTitle() string /* primitive/slice/pointer. */
	SetMediaContentTitle(value string /* primitive/slice/pointer. */)
	MediaPlaybackState() DDDeviceMediaPlaybackState
	SetMediaPlaybackState(value DDDeviceMediaPlaybackState)
	NetworkEndpoint() unsafe.Pointer
	SetNetworkEndpoint(value unsafe.Pointer)
	Protocol() DDDeviceProtocol
	SetProtocol(value DDDeviceProtocol)
	ProtocolType() objectivec.IObject
	SetProtocolType(value objectivec.IObject)
	SSID() string /* primitive/slice/pointer. */
	SetSSID(value string /* primitive/slice/pointer. */)
	State() DDDeviceState
	SetState(value DDDeviceState)
	SupportsGrouping() bool /* primitive/slice/pointer. */
	SetSupportsGrouping(value bool /* primitive/slice/pointer. */)
	TxtRecordData() foundation.objc.IObject /* cross-framework: NSData */
	SetTxtRecordData(value foundation.objc.IObject /* cross-framework: NSData */)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	WifiAwareModelName() string /* primitive/slice/pointer. */
	SetWifiAwareModelName(value string /* primitive/slice/pointer. */)
	WifiAwareServiceName() string /* primitive/slice/pointer. */
	SetWifiAwareServiceName(value string /* primitive/slice/pointer. */)
	WifiAwareServiceRole() DDDeviceWiFiAwareServiceRole
	SetWifiAwareServiceRole(value DDDeviceWiFiAwareServiceRole)
	WifiAwareVendorName() string /* primitive/slice/pointer. */
	SetWifiAwareVendorName(value string /* primitive/slice/pointer. */)
	TxtRecord() unsafe.Pointer
	SetTxtRecord(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (dc _DDDeviceClass) Alloc() DDDevice {
	rv := objc.Send[DDDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an object that describes a discovered device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/init(displayName:category:protocolType:identifier:)
func NewDDDeviceWithDisplayNameCategoryProtocolTypeIdentifier(displayName string /* primitive/slice/pointer. */, category DDDeviceCategory, protocolType objectivec.IObject, identifier string /* primitive/slice/pointer. */) DDDevice {
	instance := getDDDeviceClass().Alloc()
	rv := objc.Send[DDDevice](instance.ID, objc.Sel("initWithDisplayName:category:protocolType:identifier:"), objc.String(displayName), category, protocolType, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// An identifier to communicate with the device through Bluetooth wireless technology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/bluetoothIdentifier
func (d_ DDDevice) BluetoothIdentifier() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[UUID](d_.ID, objc.Sel("bluetoothIdentifier"))
	return rv
}


// An identifier to communicate with the device through Bluetooth wireless technology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/bluetoothIdentifier
func (d_ DDDevice) SetBluetoothIdentifier(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBluetoothIdentifier:"), value)
}


// An option that determies the icon that the picker UI displays for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/category-swift.property
func (d_ DDDevice) Category() DDDeviceCategory {
	rv := objc.Send[DDDeviceCategory](d_.ID, objc.Sel("category"))
	return rv
}


// An option that determies the icon that the picker UI displays for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/category-swift.property
func (d_ DDDevice) SetCategory(value DDDeviceCategory) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCategory:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/deviceSupports
func (d_ DDDevice) DeviceSupports() DDDeviceSupports {
	rv := objc.Send[DDDeviceSupports](d_.ID, objc.Sel("deviceSupports"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/deviceSupports
func (d_ DDDevice) SetDeviceSupports(value DDDeviceSupports) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDeviceSupports:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayImageName
func (d_ DDDevice) DisplayImageName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("displayImageName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayImageName
func (d_ DDDevice) SetDisplayImageName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayImageName:"), objc.String(value))
}


// A name for the device to display to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayName
func (d_ DDDevice) DisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("displayName"))
	return rv
}


// A name for the device to display to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayName
func (d_ DDDevice) SetDisplayName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// A unique identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/identifier
func (d_ DDDevice) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("identifier"))
	return rv
}


// A unique identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/identifier
func (d_ DDDevice) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// A subtitle for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentSubtitle
func (d_ DDDevice) MediaContentSubtitle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("mediaContentSubtitle"))
	return rv
}


// A subtitle for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentSubtitle
func (d_ DDDevice) SetMediaContentSubtitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaContentSubtitle:"), objc.String(value))
}


// A title for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentTitle
func (d_ DDDevice) MediaContentTitle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("mediaContentTitle"))
	return rv
}


// A title for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentTitle
func (d_ DDDevice) SetMediaContentTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaContentTitle:"), objc.String(value))
}


// A playback status for the device’s current media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaPlaybackState-swift.property
func (d_ DDDevice) MediaPlaybackState() DDDeviceMediaPlaybackState {
	rv := objc.Send[DDDeviceMediaPlaybackState](d_.ID, objc.Sel("mediaPlaybackState"))
	return rv
}


// A playback status for the device’s current media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaPlaybackState-swift.property
func (d_ DDDevice) SetMediaPlaybackState(value DDDeviceMediaPlaybackState) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaPlaybackState:"), value)
}


// An object that describes a local-network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/networkEndpoint-7excg
func (d_ DDDevice) NetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("networkEndpoint"))
	return rv
}


// An object that describes a local-network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/networkEndpoint-7excg
func (d_ DDDevice) SetNetworkEndpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNetworkEndpoint:"), value)
}


// The manner in which the system applies your app’s device discovery extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocol-swift.property
func (d_ DDDevice) Protocol() DDDeviceProtocol {
	rv := objc.Send[DDDeviceProtocol](d_.ID, objc.Sel("protocol"))
	return rv
}


// The manner in which the system applies your app’s device discovery extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocol-swift.property
func (d_ DDDevice) SetProtocol(value DDDeviceProtocol) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setProtocol:"), value)
}


// A custom universal type that describes the device’s manner of communication with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocolType
func (d_ DDDevice) ProtocolType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("protocolType"))
	return rv
}


// A custom universal type that describes the device’s manner of communication with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocolType
func (d_ DDDevice) SetProtocolType(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setProtocolType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/ssid
func (d_ DDDevice) SSID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("SSID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/ssid
func (d_ DDDevice) SetSSID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSSID:"), objc.String(value))
}


// A state that represents the level of user interaction with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/state
func (d_ DDDevice) State() DDDeviceState {
	rv := objc.Send[DDDeviceState](d_.ID, objc.Sel("state"))
	return rv
}


// A state that represents the level of user interaction with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/state
func (d_ DDDevice) SetState(value DDDeviceState) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setState:"), value)
}


// A Boolean value that indicates whether to group the device with others in the AirPlay UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/supportsGrouping
func (d_ DDDevice) SupportsGrouping() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("supportsGrouping"))
	return rv
}


// A Boolean value that indicates whether to group the device with others in the AirPlay UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/supportsGrouping
func (d_ DDDevice) SetSupportsGrouping(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportsGrouping:"), value)
}


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/txtRecordData
func (d_ DDDevice) TxtRecordData() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("txtRecordData"))
	return rv
}


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/txtRecordData
func (d_ DDDevice) SetTxtRecordData(value foundation.objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTxtRecordData:"), value)
}


// A resource locator for the simple service discovery protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/url
func (d_ DDDevice) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("url"))
	return rv
}


// A resource locator for the simple service discovery protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/url
func (d_ DDDevice) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUrl:"), value)
}


// Device’s Wi-Fi Aware model name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareModelName
func (d_ DDDevice) WifiAwareModelName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("wifiAwareModelName"))
	return rv
}


// Device’s Wi-Fi Aware model name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareModelName
func (d_ DDDevice) SetWifiAwareModelName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareModelName:"), objc.String(value))
}


// Device’s Wi-Fi Aware’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceName
func (d_ DDDevice) WifiAwareServiceName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("wifiAwareServiceName"))
	return rv
}


// Device’s Wi-Fi Aware’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceName
func (d_ DDDevice) SetWifiAwareServiceName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareServiceName:"), objc.String(value))
}


// Device’s Wi-Fi Aware’s service. Default is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceRole-swift.property
func (d_ DDDevice) WifiAwareServiceRole() DDDeviceWiFiAwareServiceRole {
	rv := objc.Send[DDDeviceWiFiAwareServiceRole](d_.ID, objc.Sel("wifiAwareServiceRole"))
	return rv
}


// Device’s Wi-Fi Aware’s service. Default is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceRole-swift.property
func (d_ DDDevice) SetWifiAwareServiceRole(value DDDeviceWiFiAwareServiceRole) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareServiceRole:"), value)
}


// Device’s Wi-Fi Aware vendor name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareVendorName
func (d_ DDDevice) WifiAwareVendorName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("wifiAwareVendorName"))
	return rv
}


// Device’s Wi-Fi Aware vendor name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareVendorName
func (d_ DDDevice) SetWifiAwareVendorName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWifiAwareVendorName:"), objc.String(value))
}


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) TxtRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("txtRecord"))
	return rv
}


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) SetTxtRecord(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTxtRecord:"), value)
}


