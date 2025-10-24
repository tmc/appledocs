// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotNetwork */


/* debug [class_header]: Header for NEHotspotNetwork */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotNetwork */
// An interface definition for the [NEHotspotNetwork] class.
type INEHotspotNetwork interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotNetwork */
	// properties:
	DidAutoJoin() bool
	SetDidAutoJoin(value bool)
	DidJustJoin() bool
	SetDidJustJoin(value bool)
	IsChosenHelper() bool
	SetIsChosenHelper(value bool)
	IsSecure() bool
	SetIsSecure(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotNetwork */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotNetwork */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotNetworkClass) Alloc() NEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotNetwork */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotNetwork *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotNetwork */

// Fetches information about the current Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/fetchCurrent(completionHandler:)
func (nc _NEHotspotNetworkClass) FetchCurrentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("fetchCurrentWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FetchCurrentWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotNetwork */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotNetwork */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotNetwork */

// Indicates whether the network was joined automatically or was joined explicitly by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didautojoin
func (n_ NEHotspotNetwork) DidAutoJoin() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("didAutoJoin"))
	return rv
}/* debug [instance_properties/getter]: didAutoJoin */


// Indicates whether the network was joined automatically or was joined explicitly by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didautojoin
func (n_ NEHotspotNetwork) SetDidAutoJoin(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDidAutoJoin:"), value)
}/* debug [instance_properties/setter]: didAutoJoin */


// Indicates whether the network was just joined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didjustjoin
func (n_ NEHotspotNetwork) DidJustJoin() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("didJustJoin"))
	return rv
}/* debug [instance_properties/getter]: didJustJoin */


// Indicates whether the network was just joined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/didjustjoin
func (n_ NEHotspotNetwork) SetDidJustJoin(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDidJustJoin:"), value)
}/* debug [instance_properties/setter]: didJustJoin */


// Indicates whether the calling Hotspot Helper is the chosen helper for this network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ischosenhelper
func (n_ NEHotspotNetwork) IsChosenHelper() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isChosenHelper"))
	return rv
}/* debug [instance_properties/getter]: isChosenHelper */


// Indicates whether the calling Hotspot Helper is the chosen helper for this network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/ischosenhelper
func (n_ NEHotspotNetwork) SetIsChosenHelper(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsChosenHelper:"), value)
}/* debug [instance_properties/setter]: isChosenHelper */


// Indicates whether the network is secure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/issecure
func (n_ NEHotspotNetwork) IsSecure() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isSecure"))
	return rv
}/* debug [instance_properties/getter]: isSecure */


// Indicates whether the network is secure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotnetwork/issecure
func (n_ NEHotspotNetwork) SetIsSecure(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsSecure:"), value)
}/* debug [instance_properties/setter]: isSecure */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotNetwork */


