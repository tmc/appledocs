// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEPrivateLTENetwork] class.
var (
	NEPrivateLTENetworkClass     _NEPrivateLTENetworkClass
	NEPrivateLTENetworkClassOnce sync.Once
)

func getNEPrivateLTENetworkClass() _NEPrivateLTENetworkClass {
	NEPrivateLTENetworkClassOnce.Do(func() {
		NEPrivateLTENetworkClass = _NEPrivateLTENetworkClass{objc.GetClass("NEPrivateLTENetwork")}
	})
	return NEPrivateLTENetworkClass
}

type _NEPrivateLTENetworkClass struct {
	class objc.Class
}

// An interface definition for the [NEPrivateLTENetwork] class.
type INEPrivateLTENetwork interface {
	objectivec.IObject
	MatchPrivateLTENetworks() NEPrivateLTENetwork
	SetMatchPrivateLTENetworks(value INEPrivateLTENetwork)
	MatchSSIDs() string
	SetMatchSSIDs(value string)
	MobileCountryCode() string
	SetMobileCountryCode(value string)
	MobileNetworkCode() string
	SetMobileNetworkCode(value string)
	TrackingAreaCode() string
	SetTrackingAreaCode(value string)
}

// The parameters of a private LTE network.
//
// Populate your manager’s with an array of objects of this type. The system starts the provider when the device’s current private LTE provider matches the properties of any member of the array.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork
type NEPrivateLTENetwork struct {
	objectivec.Object
}

// NEPrivateLTENetworkFrom constructs a [NEPrivateLTENetwork] from an unsafe.Pointer.
//
// The parameters of a private LTE network.
func NEPrivateLTENetworkFrom(ptr unsafe.Pointer) NEPrivateLTENetwork {
	return NEPrivateLTENetwork{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEPrivateLTENetworkClass) Alloc() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEPrivateLTENetworkClass) New() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPrivateLTENetwork) Init() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPrivateLTENetwork) Autorelease() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPrivateLTENetwork creates a new NEPrivateLTENetwork instance.
func NewNEPrivateLTENetwork() NEPrivateLTENetwork {
	return getNEPrivateLTENetworkClass().New()
}


// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEPrivateLTENetwork) MatchPrivateLTENetworks() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](n_.ID, objc.Sel("matchPrivateLTENetworks"))
	return rv
}


// SetMatchPrivateLTENetworks sets the value of the matchPrivateLTENetworks property.
// An array of private LTE networks that the system matches for local push activation.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEPrivateLTENetwork) SetMatchPrivateLTENetworks(value INEPrivateLTENetwork) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPrivateLTENetworks:"), value)
}

// An array of Wi-Fi SSID strings that the system matches for local push activation.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchssids
func (n_ NEPrivateLTENetwork) MatchSSIDs() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchSSIDs"))
	return rv
}


// SetMatchSSIDs sets the value of the matchSSIDs property.
// An array of Wi-Fi SSID strings that the system matches for local push activation.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchssids
func (n_ NEPrivateLTENetwork) SetMatchSSIDs(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchSSIDs:"), objc.String(value))
}

// The Mobile Country Code (MCC) of the private LTE network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/mobilecountrycode
func (n_ NEPrivateLTENetwork) MobileCountryCode() string {
	rv := objc.Send[string](n_.ID, objc.Sel("mobileCountryCode"))
	return rv
}


// SetMobileCountryCode sets the value of the mobileCountryCode property.
// The Mobile Country Code (MCC) of the private LTE network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/mobilecountrycode
func (n_ NEPrivateLTENetwork) SetMobileCountryCode(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMobileCountryCode:"), objc.String(value))
}

// The Mobile Network Code (MNC) of the private LTE network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/mobilenetworkcode
func (n_ NEPrivateLTENetwork) MobileNetworkCode() string {
	rv := objc.Send[string](n_.ID, objc.Sel("mobileNetworkCode"))
	return rv
}


// SetMobileNetworkCode sets the value of the mobileNetworkCode property.
// The Mobile Network Code (MNC) of the private LTE network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/mobilenetworkcode
func (n_ NEPrivateLTENetwork) SetMobileNetworkCode(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMobileNetworkCode:"), objc.String(value))
}

// The Tracking Area Code of the private LTE network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/trackingareacode
func (n_ NEPrivateLTENetwork) TrackingAreaCode() string {
	rv := objc.Send[string](n_.ID, objc.Sel("trackingAreaCode"))
	return rv
}


// SetTrackingAreaCode sets the value of the trackingAreaCode property.
// The Tracking Area Code of the private LTE network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprivateltenetwork/trackingareacode
func (n_ NEPrivateLTENetwork) SetTrackingAreaCode(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTrackingAreaCode:"), objc.String(value))
}



