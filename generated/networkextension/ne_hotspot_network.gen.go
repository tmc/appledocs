// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// Information about a Wi-Fi network associated with a command or a response.
//
// When the Hotspot Helper app is asked to evaluate the a network or filter the Wi-Fi scan list, it annotates the object via the method.
//
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
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/fetchCurrent(completionHandler:)
func (nc _NEHotspotNetworkClass) FetchCurrentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("fetchCurrentWithCompletionHandler:"), completionHandler)
}



