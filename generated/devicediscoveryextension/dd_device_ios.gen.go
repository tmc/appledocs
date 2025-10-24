//go:build darwin && ios

// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// iOS-only methods for DDDevice


// iOS-only properties

// An identifier to communicate with the device through Bluetooth wireless technology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/bluetoothIdentifier
func (d_ DDDevice) BluetoothIdentifier() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](d_.ID, objc.Sel("bluetoothIdentifier"))
	return rv
}
func (d_ DDDevice) SetBluetoothIdentifier(value objc.IObject /* cross-framework: UUID */) {
	d_.ID.Send(objc.RegisterName("setBluetoothIdentifier:"), value)
}

// An option that determies the icon that the picker UI displays for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/category-swift.property
func (d_ DDDevice) Category() DDDeviceCategory {
	rv := objc.Send[DDDeviceCategory](d_.ID, objc.Sel("category"))
	return rv
}
func (d_ DDDevice) SetCategory(value DDDeviceCategory) {
	d_.ID.Send(objc.RegisterName("setCategory:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/deviceSupports
func (d_ DDDevice) DeviceSupports() DDDeviceSupports {
	rv := objc.Send[DDDeviceSupports](d_.ID, objc.Sel("deviceSupports"))
	return rv
}
func (d_ DDDevice) SetDeviceSupports(value DDDeviceSupports) {
	d_.ID.Send(objc.RegisterName("setDeviceSupports:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayImageName
func (d_ DDDevice) DisplayImageName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayImageName"))
	return rv
}
func (d_ DDDevice) SetDisplayImageName(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setDisplayImageName:"), value)
}

// A name for the device to display to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/displayName
func (d_ DDDevice) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayName"))
	return rv
}
func (d_ DDDevice) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setDisplayName:"), value)
}

// A unique identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/identifier
func (d_ DDDevice) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("identifier"))
	return rv
}
func (d_ DDDevice) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setIdentifier:"), value)
}

// A subtitle for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentSubtitle
func (d_ DDDevice) MediaContentSubtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("mediaContentSubtitle"))
	return rv
}
func (d_ DDDevice) SetMediaContentSubtitle(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setMediaContentSubtitle:"), value)
}

// A title for the current media that the device plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaContentTitle
func (d_ DDDevice) MediaContentTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("mediaContentTitle"))
	return rv
}
func (d_ DDDevice) SetMediaContentTitle(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setMediaContentTitle:"), value)
}

// A playback status for the device’s current media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/mediaPlaybackState-swift.property
func (d_ DDDevice) MediaPlaybackState() DDDeviceMediaPlaybackState {
	rv := objc.Send[DDDeviceMediaPlaybackState](d_.ID, objc.Sel("mediaPlaybackState"))
	return rv
}
func (d_ DDDevice) SetMediaPlaybackState(value DDDeviceMediaPlaybackState) {
	d_.ID.Send(objc.RegisterName("setMediaPlaybackState:"), value)
}

// An object that describes a local-network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/networkEndpoint-7excg
func (d_ DDDevice) NetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("networkEndpoint"))
	return rv
}
func (d_ DDDevice) SetNetworkEndpoint(value unsafe.Pointer) {
	d_.ID.Send(objc.RegisterName("setNetworkEndpoint:"), value)
}

// The manner in which the system applies your app’s device discovery extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocol-swift.property
func (d_ DDDevice) Protocol() DDDeviceProtocol {
	rv := objc.Send[DDDeviceProtocol](d_.ID, objc.Sel("protocol"))
	return rv
}
func (d_ DDDevice) SetProtocol(value DDDeviceProtocol) {
	d_.ID.Send(objc.RegisterName("setProtocol:"), value)
}

// A custom universal type that describes the device’s manner of communication with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/protocolType
func (d_ DDDevice) ProtocolType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](d_.ID, objc.Sel("protocolType"))
	return rv
}
func (d_ DDDevice) SetProtocolType(value objc.IObject /* cross-framework: UTType */) {
	d_.ID.Send(objc.RegisterName("setProtocolType:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/ssid
func (d_ DDDevice) SSID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("SSID"))
	return rv
}
func (d_ DDDevice) SetSSID(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setSSID:"), value)
}

// A state that represents the level of user interaction with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/state
func (d_ DDDevice) State() DDDeviceState {
	rv := objc.Send[DDDeviceState](d_.ID, objc.Sel("state"))
	return rv
}
func (d_ DDDevice) SetState(value DDDeviceState) {
	d_.ID.Send(objc.RegisterName("setState:"), value)
}

// A Boolean value that indicates whether to group the device with others in the AirPlay UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/supportsGrouping
func (d_ DDDevice) SupportsGrouping() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("supportsGrouping"))
	return rv
}
func (d_ DDDevice) SetSupportsGrouping(value bool) {
	d_.ID.Send(objc.RegisterName("setSupportsGrouping:"), value)
}

// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/txtRecordData
func (d_ DDDevice) TxtRecordData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("txtRecordData"))
	return rv
}
func (d_ DDDevice) SetTxtRecordData(value objc.IObject /* cross-framework: NSData */) {
	d_.ID.Send(objc.RegisterName("setTxtRecordData:"), value)
}

// A resource locator for the simple service discovery protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/url
func (d_ DDDevice) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("url"))
	return rv
}
func (d_ DDDevice) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	d_.ID.Send(objc.RegisterName("setUrl:"), value)
}

// Device’s Wi-Fi Aware model name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareModelName
func (d_ DDDevice) WifiAwareModelName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareModelName"))
	return rv
}
func (d_ DDDevice) SetWifiAwareModelName(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setWifiAwareModelName:"), value)
}

// Device’s Wi-Fi Aware’s service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceName
func (d_ DDDevice) WifiAwareServiceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareServiceName"))
	return rv
}
func (d_ DDDevice) SetWifiAwareServiceName(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setWifiAwareServiceName:"), value)
}

// Device’s Wi-Fi Aware’s service. Default is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareServiceRole-swift.property
func (d_ DDDevice) WifiAwareServiceRole() DDDeviceWiFiAwareServiceRole {
	rv := objc.Send[DDDeviceWiFiAwareServiceRole](d_.ID, objc.Sel("wifiAwareServiceRole"))
	return rv
}
func (d_ DDDevice) SetWifiAwareServiceRole(value DDDeviceWiFiAwareServiceRole) {
	d_.ID.Send(objc.RegisterName("setWifiAwareServiceRole:"), value)
}

// Device’s Wi-Fi Aware vendor name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/wifiAwareVendorName
func (d_ DDDevice) WifiAwareVendorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wifiAwareVendorName"))
	return rv
}
func (d_ DDDevice) SetWifiAwareVendorName(value objc.IObject /* cross-framework: NSString */) {
	d_.ID.Send(objc.RegisterName("setWifiAwareVendorName:"), value)
}




