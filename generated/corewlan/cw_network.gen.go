// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CWNetwork] class.
type ICWNetwork interface {
	objectivec.IObject
	BeaconInterval() int
	Bssid() string
	CountryCode() string
	Ibss() bool
	InformationElementData() foundation.NSData
	NoiseMeasurement() int
	RssiValue() int
	Ssid() string
	SsidData() foundation.NSData
	WlanChannel() ICWChannel
	IsEqualToNetwork(network ICWNetwork) bool
	SupportsPHYMode(phyMode CWPHYMode) bool
	SupportsSecurity(security CWSecurity) bool
}

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

// Alloc allocates a new instance without initialization.
func (cc _CWNetworkClass) Alloc() CWNetwork {
	rv := objc.Send[CWNetwork](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Method for determining CWNetwork object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/isEqual(to:)
func (c_ CWNetwork) IsEqualToNetwork(network ICWNetwork) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToNetwork:"), network)
	return rv
}


// Method for determining which PHY modes a network supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsPHYMode(_:)
func (c_ CWNetwork) SupportsPHYMode(phyMode CWPHYMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsPHYMode:"), phyMode)
	return rv
}


// Method for determining which security types a network supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/supportsSecurity(_:)
func (c_ CWNetwork) SupportsSecurity(security CWSecurity) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsSecurity:"), security)
	return rv
}


// The beacon interval (ms) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/beaconInterval
func (c_ CWNetwork) BeaconInterval() int {
	rv := objc.Send[int](c_.ID, objc.Sel("beaconInterval"))
	return rv
}


// The basic service set identifier (BSSID) for the network, returned as UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/bssid
func (c_ CWNetwork) Bssid() string {
	rv := objc.Send[string](c_.ID, objc.Sel("bssid"))
	return rv
}


// The country code (ISO/IEC 3166-1:1997) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/countryCode
func (c_ CWNetwork) CountryCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("countryCode"))
	return rv
}


// The network is an IBSS network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ibss
func (c_ CWNetwork) Ibss() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ibss"))
	return rv
}


// Information element data included in beacon or probe response frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/informationElementData
func (c_ CWNetwork) InformationElementData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("informationElementData"))
	return rv
}


// The aggregate noise measurement (dBm) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/noiseMeasurement
func (c_ CWNetwork) NoiseMeasurement() int {
	rv := objc.Send[int](c_.ID, objc.Sel("noiseMeasurement"))
	return rv
}


// The aggregate received signal strength indication (RSSI) measurement (dBm) for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/rssiValue
func (c_ CWNetwork) RssiValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("rssiValue"))
	return rv
}


// The service set identifier (SSID) for the network, encoded as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssid
func (c_ CWNetwork) Ssid() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ssid"))
	return rv
}


// The service set identifier (SSID) for the network, returned as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/ssidData
func (c_ CWNetwork) SsidData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("ssidData"))
	return rv
}


// The channel for the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetwork/wlanChannel
func (c_ CWNetwork) WlanChannel() ICWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("wlanChannel"))
	return rv
}



