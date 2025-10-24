// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	MacAddress() objc.IObject /* cross-framework: NSString */
	SetMacAddress(value objc.IObject /* cross-framework: NSString */)
	Model() objc.IObject /* cross-framework: NSString */
	SetModel(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Properties() EAWiFiUnconfiguredAccessoryProperties
	SetProperties(value EAWiFiUnconfiguredAccessoryProperties)
	Ssid() objc.IObject /* cross-framework: NSString */
	SetSsid(value objc.IObject /* cross-framework: NSString */)
	// methods:
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



// The primary MAC address of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/macaddress
func (e_ EAWiFiUnconfiguredAccessory) MacAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("macAddress"))
	return rv
}


// The primary MAC address of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/macaddress
func (e_ EAWiFiUnconfiguredAccessory) SetMacAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMacAddress:"), value)
}


// The model name of accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/model
func (e_ EAWiFiUnconfiguredAccessory) Model() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("model"))
	return rv
}


// The model name of accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/model
func (e_ EAWiFiUnconfiguredAccessory) SetModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setModel:"), value)
}


// The name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/name
func (e_ EAWiFiUnconfiguredAccessory) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}


// The name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/name
func (e_ EAWiFiUnconfiguredAccessory) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
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
func (e_ EAWiFiUnconfiguredAccessory) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("ssid"))
	return rv
}


// The Wi-Fi SSID of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eawifiunconfiguredaccessory/ssid
func (e_ EAWiFiUnconfiguredAccessory) SetSsid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSsid:"), value)
}


