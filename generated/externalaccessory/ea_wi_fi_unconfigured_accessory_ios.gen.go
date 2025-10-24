//go:build darwin && ios

// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for EAWiFiUnconfiguredAccessory


// iOS-only properties

// The primary MAC address of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/macAddress
func (e_ EAWiFiUnconfiguredAccessory) MacAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("macAddress"))
	return rv
}

// The name of the accessory’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/manufacturer
func (e_ EAWiFiUnconfiguredAccessory) Manufacturer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("manufacturer"))
	return rv
}

// The model name of accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/model
func (e_ EAWiFiUnconfiguredAccessory) Model() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("model"))
	return rv
}

// The name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/name
func (e_ EAWiFiUnconfiguredAccessory) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}

// The properties the accessory supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/properties
func (e_ EAWiFiUnconfiguredAccessory) Properties() EAWiFiUnconfiguredAccessoryProperties {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryProperties](e_.ID, objc.Sel("properties"))
	return rv
}

// The Wi-Fi SSID of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/ssid
func (e_ EAWiFiUnconfiguredAccessory) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("ssid"))
	return rv
}





