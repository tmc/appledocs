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
	

	// properties:
	MatchPrivateLTENetworks() INEPrivateLTENetwork
	SetMatchPrivateLTENetworks(value INEPrivateLTENetwork)
	MatchSSIDs() foundation.foundation.INSString
	SetMatchSSIDs(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEPrivateLTENetworkClass) Alloc() NEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The parameters of a private LTE network.
//
// Populate your manager’s with an array of objects of this type. The system starts the provider when the device’s current private LTE provider matches the properties of any member of the array.


// The parameters of a private LTE network.
//
// [Full Topic]
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

























// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEPrivateLTENetwork) MatchPrivateLTENetworks() INEPrivateLTENetwork {
	rv := objc.Send[NEPrivateLTENetwork](n_.ID, objc.Sel("matchPrivateLTENetworks"))
	return rv
}


// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEPrivateLTENetwork) SetMatchPrivateLTENetworks(value INEPrivateLTENetwork) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPrivateLTENetworks:"), value)
}


// An array of Wi-Fi SSID strings that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchssids
func (n_ NEPrivateLTENetwork) MatchSSIDs() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchSSIDs"))
	return rv
}


// An array of Wi-Fi SSID strings that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchssids
func (n_ NEPrivateLTENetwork) SetMatchSSIDs(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchSSIDs:"), value)
}







