// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotNetwork] class.
var (
	NEHotspotNetworkClass     _NEHotspotNetworkClass
	NEHotspotNetworkClassOnce sync.Once
)

func getNEHotspotNetworkClass() _NEHotspotNetworkClass {
	NEHotspotNetworkClassOnce.Do(func() {
		NEHotspotNetworkClass = _NEHotspotNetworkClass{objc.GetClass("NEHotspotNetwork")}
	})
	return NEHotspotNetworkClass
}

type _NEHotspotNetworkClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotNetwork] class.
type INEHotspotNetwork interface {
	objectivec.IObject
	// properties:
	Bssid() objc.IObject /* cross-framework: NSString */
	SetBssid(value objc.IObject /* cross-framework: NSString */)
	DidAutoJoin() bool
	SetDidAutoJoin(value bool)
	DidJustJoin() bool
	SetDidJustJoin(value bool)
	IsChosenHelper() bool
	SetIsChosenHelper(value bool)
	IsSecure() bool
	SetIsSecure(value bool)
	SecurityType() unsafe.Pointer
	SetSecurityType(value unsafe.Pointer)
	SignalStrength() float64
	SetSignalStrength(value float64)
	Ssid() objc.IObject /* cross-framework: NSString */
	SetSsid(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// Information about a Wi-Fi network associated with a command or a response.
//
// When the Hotspot Helper app is asked to evaluate the a network or filter the Wi-Fi scan list, it annotates the object via the method.


// Information about a Wi-Fi network associated with a command or a response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork
type NEHotspotNetwork struct {
	objectivec.Object
}

// NEHotspotNetworkFrom constructs a [NEHotspotNetwork] from an unsafe.Pointer.
//
// Information about a Wi-Fi network associated with a command or a response.
func NEHotspotNetworkFrom(ptr unsafe.Pointer) NEHotspotNetwork {
	return NEHotspotNetwork{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotNetworkClass) Alloc() NEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotNetworkClass) New() NEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotNetwork) Init() NEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotNetwork) Autorelease() NEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotNetwork creates a new NEHotspotNetwork instance.
func NewNEHotspotNetwork() NEHotspotNetwork {
	return getNEHotspotNetworkClass().New()
}



// Fetches information about the current Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/fetchCurrent(completionHandler:)
func (nc _NEHotspotNetworkClass) FetchCurrentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("fetchCurrentWithCompletionHandler:"), completionHandler)
}


// The BSSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/bssid
func (n_ NEHotspotNetwork) Bssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("bssid"))
	return rv
}


// The BSSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/bssid
func (n_ NEHotspotNetwork) SetBssid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBssid:"), value)
}


// Indicates whether the network was joined automatically or was joined explicitly by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didautojoin
func (n_ NEHotspotNetwork) DidAutoJoin() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("didAutoJoin"))
	return rv
}


// Indicates whether the network was joined automatically or was joined explicitly by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didautojoin
func (n_ NEHotspotNetwork) SetDidAutoJoin(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDidAutoJoin:"), value)
}


// Indicates whether the network was just joined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didjustjoin
func (n_ NEHotspotNetwork) DidJustJoin() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("didJustJoin"))
	return rv
}


// Indicates whether the network was just joined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didjustjoin
func (n_ NEHotspotNetwork) SetDidJustJoin(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDidJustJoin:"), value)
}


// Indicates whether the calling Hotspot Helper is the chosen helper for this network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ischosenhelper
func (n_ NEHotspotNetwork) IsChosenHelper() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isChosenHelper"))
	return rv
}


// Indicates whether the calling Hotspot Helper is the chosen helper for this network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ischosenhelper
func (n_ NEHotspotNetwork) SetIsChosenHelper(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsChosenHelper:"), value)
}


// Indicates whether the network is secure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/issecure
func (n_ NEHotspotNetwork) IsSecure() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isSecure"))
	return rv
}


// Indicates whether the network is secure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/issecure
func (n_ NEHotspotNetwork) SetIsSecure(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsSecure:"), value)
}


// The type of security used by the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/securitytype
func (n_ NEHotspotNetwork) SecurityType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("securityType"))
	return rv
}


// The type of security used by the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/securitytype
func (n_ NEHotspotNetwork) SetSecurityType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecurityType:"), value)
}


// The recent signal strength for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/signalstrength
func (n_ NEHotspotNetwork) SignalStrength() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("signalStrength"))
	return rv
}


// The recent signal strength for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/signalstrength
func (n_ NEHotspotNetwork) SetSignalStrength(value float64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSignalStrength:"), value)
}


// The SSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ssid
func (n_ NEHotspotNetwork) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("ssid"))
	return rv
}


// The SSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ssid
func (n_ NEHotspotNetwork) SetSsid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSsid:"), value)
}


