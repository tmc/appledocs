// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EAWiFiUnconfiguredAccessory] class.
var (
	EAWiFiUnconfiguredAccessoryClass     _EAWiFiUnconfiguredAccessoryClass
	EAWiFiUnconfiguredAccessoryClassOnce sync.Once
)

func getEAWiFiUnconfiguredAccessoryClass() _EAWiFiUnconfiguredAccessoryClass {
	EAWiFiUnconfiguredAccessoryClassOnce.Do(func() {
		EAWiFiUnconfiguredAccessoryClass = _EAWiFiUnconfiguredAccessoryClass{objc.GetClass("EAWiFiUnconfiguredAccessory")}
	})
	return EAWiFiUnconfiguredAccessoryClass
}

type _EAWiFiUnconfiguredAccessoryClass struct {
	class objc.Class
}

// An interface definition for the [EAWiFiUnconfiguredAccessory] class.
type IEAWiFiUnconfiguredAccessory interface {
	objectivec.IObject
	Manufacturer() string
	MacAddress() string
	SetMacAddress(value string)
	Model() string
	SetModel(value string)
	Name() string
	SetName(value string)
	Properties() EAWiFiUnconfiguredAccessoryProperties
	SetProperties(value EAWiFiUnconfiguredAccessoryProperties)
	Ssid() string
	SetSsid(value string)
}

// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.


// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory
type EAWiFiUnconfiguredAccessory struct {
	objectivec.Object
}

// EAWiFiUnconfiguredAccessoryFrom constructs a [EAWiFiUnconfiguredAccessory] from an unsafe.Pointer.
//
// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.
func EAWiFiUnconfiguredAccessoryFrom(ptr unsafe.Pointer) EAWiFiUnconfiguredAccessory {
	return EAWiFiUnconfiguredAccessory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EAWiFiUnconfiguredAccessoryClass) Alloc() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAWiFiUnconfiguredAccessoryClass) New() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAWiFiUnconfiguredAccessory) Init() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAWiFiUnconfiguredAccessory) Autorelease() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAWiFiUnconfiguredAccessory creates a new EAWiFiUnconfiguredAccessory instance.
func NewEAWiFiUnconfiguredAccessory() EAWiFiUnconfiguredAccessory {
	return getEAWiFiUnconfiguredAccessoryClass().New()
}



// The name of the accessory’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory/manufacturer
func (e_ EAWiFiUnconfiguredAccessory) Manufacturer() string {
	rv := objc.Send[string](e_.ID, objc.Sel("manufacturer"))
	return rv
}


// The primary MAC address of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/macaddress
func (e_ EAWiFiUnconfiguredAccessory) MacAddress() string {
	rv := objc.Send[string](e_.ID, objc.Sel("macAddress"))
	return rv
}


// The primary MAC address of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/macaddress
func (e_ EAWiFiUnconfiguredAccessory) SetMacAddress(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMacAddress:"), objc.String(value))
}


// The model name of accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/model
func (e_ EAWiFiUnconfiguredAccessory) Model() string {
	rv := objc.Send[string](e_.ID, objc.Sel("model"))
	return rv
}


// The model name of accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/model
func (e_ EAWiFiUnconfiguredAccessory) SetModel(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setModel:"), objc.String(value))
}


// The name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/name
func (e_ EAWiFiUnconfiguredAccessory) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// The name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/name
func (e_ EAWiFiUnconfiguredAccessory) SetName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), objc.String(value))
}


// The properties the accessory supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/properties
func (e_ EAWiFiUnconfiguredAccessory) Properties() EAWiFiUnconfiguredAccessoryProperties {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryProperties](e_.ID, objc.Sel("properties"))
	return rv
}


// The properties the accessory supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/properties
func (e_ EAWiFiUnconfiguredAccessory) SetProperties(value EAWiFiUnconfiguredAccessoryProperties) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProperties:"), value)
}


// The Wi-Fi SSID of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/ssid
func (e_ EAWiFiUnconfiguredAccessory) Ssid() string {
	rv := objc.Send[string](e_.ID, objc.Sel("ssid"))
	return rv
}


// The Wi-Fi SSID of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/ssid
func (e_ EAWiFiUnconfiguredAccessory) SetSsid(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSsid:"), objc.String(value))
}



