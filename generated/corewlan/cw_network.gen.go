// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CWNetwork */


/* debug [class_header]: Header for CWNetwork */
// The class instance for the [CWNetwork] class.
var (
	CWNetworkClass     _CWNetworkClass
	CWNetworkClassOnce sync.Once
)

func getCWNetworkClass() _CWNetworkClass {
	CWNetworkClassOnce.Do(func() {
		CWNetworkClass = _CWNetworkClass{objc.GetClass("CWNetwork")}
	})
	return CWNetworkClass
}

type _CWNetworkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWNetwork */
// An interface definition for the [CWNetwork] class.
type ICWNetwork interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWNetwork */
	// properties:
	BeaconInterval() int
	Bssid() objc.IObject /* cross-framework: NSString */
	CountryCode() objc.IObject /* cross-framework: NSString */
	Ibss() bool
	InformationElementData() objc.IObject /* cross-framework: NSData */
	NoiseMeasurement() int
	RssiValue() int
	Ssid() objc.IObject /* cross-framework: NSString */
	SsidData() objc.IObject /* cross-framework: NSData */
	WlanChannel() ICWChannel
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWNetwork */
	// methods:
	IsEqualToNetwork(network ICWNetwork) bool
	SupportsPHYMode(phyMode CWPHYMode) bool
	SupportsSecurity(security CWSecurity) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWNetwork */
// Alloc allocates a new instance without initialization.
func (cc _CWNetworkClass) Alloc() CWNetwork {
	rv := objc.Send[CWNetwork](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CWNetworkClass) New() CWNetwork {
	rv := objc.Send[CWNetwork](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWNetwork) Init() CWNetwork {
	rv := objc.Send[CWNetwork](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWNetwork) Autorelease() CWNetwork {
	rv := objc.Send[CWNetwork](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWNetwork creates a new CWNetwork instance.
func NewCWNetwork() CWNetwork {
	return getCWNetworkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWNetwork */
// Encapsulates an IEEE 802.11 network, providing read-only accessors to various properties of the network.


// Encapsulates an IEEE 802.11 network, providing read-only accessors to various properties of the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork
type CWNetwork struct {
	objectivec.Object
}

// CWNetworkFrom constructs a [CWNetwork] from an unsafe.Pointer.
//
// Encapsulates an IEEE 802.11 network, providing read-only accessors to various properties of the network.
func CWNetworkFrom(ptr unsafe.Pointer) CWNetwork {
	return CWNetwork{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWNetwork *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWNetwork */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWNetwork */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWNetwork */

// Method for determining CWNetwork object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/isEqual(to:)
func (c_ CWNetwork) IsEqualToNetwork(network ICWNetwork) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToNetwork:"), network)
	return rv
}/* debug [instance_methods/method]: IsEqualToNetwork */


// Method for determining which PHY modes a network supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsPHYMode(_:)
func (c_ CWNetwork) SupportsPHYMode(phyMode CWPHYMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsPHYMode:"), phyMode)
	return rv
}/* debug [instance_methods/method]: SupportsPHYMode */


// Method for determining which security types a network supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsSecurity(_:)
func (c_ CWNetwork) SupportsSecurity(security CWSecurity) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsSecurity:"), security)
	return rv
}/* debug [instance_methods/method]: SupportsSecurity */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWNetwork */

// The beacon interval (ms) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/beaconInterval
func (c_ CWNetwork) BeaconInterval() int {
	rv := objc.Send[int](c_.ID, objc.Sel("beaconInterval"))
	return rv
}/* debug [instance_properties/getter]: beaconInterval */


// The basic service set identifier (BSSID) for the network, returned as UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/bssid
func (c_ CWNetwork) Bssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("bssid"))
	return rv
}/* debug [instance_properties/getter]: bssid */


// The country code (ISO/IEC 3166-1:1997) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/countryCode
func (c_ CWNetwork) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_properties/getter]: countryCode */


// The network is an IBSS network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ibss
func (c_ CWNetwork) Ibss() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ibss"))
	return rv
}/* debug [instance_properties/getter]: ibss */


// Information element data included in beacon or probe response frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/informationElementData
func (c_ CWNetwork) InformationElementData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("informationElementData"))
	return rv
}/* debug [instance_properties/getter]: informationElementData */


// The aggregate noise measurement (dBm) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/noiseMeasurement
func (c_ CWNetwork) NoiseMeasurement() int {
	rv := objc.Send[int](c_.ID, objc.Sel("noiseMeasurement"))
	return rv
}/* debug [instance_properties/getter]: noiseMeasurement */


// The aggregate received signal strength indication (RSSI) measurement (dBm) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/rssiValue
func (c_ CWNetwork) RssiValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("rssiValue"))
	return rv
}/* debug [instance_properties/getter]: rssiValue */


// The service set identifier (SSID) for the network, encoded as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssid
func (c_ CWNetwork) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ssid"))
	return rv
}/* debug [instance_properties/getter]: ssid */


// The service set identifier (SSID) for the network, returned as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssidData
func (c_ CWNetwork) SsidData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("ssidData"))
	return rv
}/* debug [instance_properties/getter]: ssidData */


// The channel for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/wlanChannel
func (c_ CWNetwork) WlanChannel() ICWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("wlanChannel"))
	return rv
}/* debug [instance_properties/getter]: wlanChannel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWNetwork */



