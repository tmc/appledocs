// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NIDiscoveryToken */


/* debug [class_header]: Header for NIDiscoveryToken */
// The class instance for the [NIDiscoveryToken] class.
var (
	NIDiscoveryTokenClass     _NIDiscoveryTokenClass
	NIDiscoveryTokenClassOnce sync.Once
)

func getNIDiscoveryTokenClass() _NIDiscoveryTokenClass {
	NIDiscoveryTokenClassOnce.Do(func() {
		NIDiscoveryTokenClass = _NIDiscoveryTokenClass{objc.GetClass("NIDiscoveryToken")}
	})
	return NIDiscoveryTokenClass
}

type _NIDiscoveryTokenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NIDiscoveryToken */
// An interface definition for the [NIDiscoveryToken] class.
type INIDiscoveryToken interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NIDiscoveryToken */
	// properties:
	Configuration() INIConfiguration
	SetConfiguration(value INIConfiguration)
	DelegateQueue() unsafe.Pointer
	SetDelegateQueue(value unsafe.Pointer)
	DiscoveryToken() INIDiscoveryToken
	SetDiscoveryToken(value INIDiscoveryToken)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NIDiscoveryToken */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NIDiscoveryToken */
// Alloc allocates a new instance without initialization.
func (nc _NIDiscoveryTokenClass) Alloc() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NIDiscoveryTokenClass) New() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIDiscoveryToken) Init() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIDiscoveryToken) Autorelease() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIDiscoveryToken creates a new NIDiscoveryToken instance.
func NewNIDiscoveryToken() NIDiscoveryToken {
	return getNIDiscoveryTokenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NIDiscoveryToken */
// An object that uniquely identifies a peer that participates in an interaction session.
//
// Use to determine the peer device’s nearby interaction capabilities by examining the that describes the available capabilities on a person’s device.


// An object that uniquely identifies a peer that participates in an interaction session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDiscoveryToken
type NIDiscoveryToken struct {
	objectivec.Object
}

// NIDiscoveryTokenFrom constructs a [NIDiscoveryToken] from an unsafe.Pointer.
//
// An object that uniquely identifies a peer that participates in an interaction session.
func NIDiscoveryTokenFrom(ptr unsafe.Pointer) NIDiscoveryToken {
	return NIDiscoveryToken{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NIDiscoveryToken *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NIDiscoveryToken */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NIDiscoveryToken */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NIDiscoveryToken */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NIDiscoveryToken */

// The configuration run by the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NIDiscoveryToken) Configuration() INIConfiguration {
	rv := objc.Send[NIConfiguration](n_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The configuration run by the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NIDiscoveryToken) SetConfiguration(value INIConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NIDiscoveryToken) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */


// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NIDiscoveryToken) SetDelegateQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegateQueue:"), value)
}/* debug [instance_properties/setter]: delegateQueue */


// A temporary, random identifier for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/discoverytoken
func (n_ NIDiscoveryToken) DiscoveryToken() INIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("discoveryToken"))
	return rv
}/* debug [instance_properties/getter]: discoveryToken */


// A temporary, random identifier for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/discoverytoken
func (n_ NIDiscoveryToken) SetDiscoveryToken(value INIDiscoveryToken) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDiscoveryToken:"), value)
}/* debug [instance_properties/setter]: discoveryToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NIDiscoveryToken */


