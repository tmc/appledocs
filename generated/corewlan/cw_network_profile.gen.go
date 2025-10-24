// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CWNetworkProfile */


/* debug [class_header]: Header for CWNetworkProfile */
// The class instance for the [CWNetworkProfile] class.
var (
	CWNetworkProfileClass     _CWNetworkProfileClass
	CWNetworkProfileClassOnce sync.Once
)

func getCWNetworkProfileClass() _CWNetworkProfileClass {
	CWNetworkProfileClassOnce.Do(func() {
		CWNetworkProfileClass = _CWNetworkProfileClass{objc.GetClass("CWNetworkProfile")}
	})
	return CWNetworkProfileClass
}

type _CWNetworkProfileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWNetworkProfile */
// An interface definition for the [CWNetworkProfile] class.
type ICWNetworkProfile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWNetworkProfile */
	// properties:
	Security() CWSecurity
	Ssid() objc.IObject /* cross-framework: NSString */
	SsidData() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWNetworkProfile */
	// methods:
	IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWNetworkProfile */
// Alloc allocates a new instance without initialization.
func (cc _CWNetworkProfileClass) Alloc() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CWNetworkProfileClass) New() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWNetworkProfile) Init() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWNetworkProfile) Autorelease() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWNetworkProfile creates a new CWNetworkProfile instance.
func NewCWNetworkProfile() CWNetworkProfile {
	return getCWNetworkProfileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWNetworkProfile */
// Encapsulates an immutable network profile entry.


// Encapsulates an immutable network profile entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile
type CWNetworkProfile struct {
	objectivec.Object
}

// CWNetworkProfileFrom constructs a [CWNetworkProfile] from an unsafe.Pointer.
//
// Encapsulates an immutable network profile entry.
func CWNetworkProfileFrom(ptr unsafe.Pointer) CWNetworkProfile {
	return CWNetworkProfile{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWNetworkProfile */

// Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(networkProfile:)
func NewCWNetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) CWNetworkProfile {
	instance := getCWNetworkProfileClass().Alloc()
	rv := objc.Send[CWNetworkProfile](instance.ID, objc.Sel("initWithNetworkProfile:"), networkProfile)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCWNetworkProfileWithNetworkProfile */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWNetworkProfile */

// Convenience method for getting a CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/networkProfile
func (cc _CWNetworkProfileClass) NetworkProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("networkProfile"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NetworkProfile) */


// Convenience method for getting a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/networkProfileWithNetworkProfile:
func (cc _CWNetworkProfileClass) NetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("networkProfileWithNetworkProfile:"), networkProfile)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NetworkProfileWithNetworkProfile) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWNetworkProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWNetworkProfile */

// Determine CWNetworkProfile object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/isEqual(to:)
func (c_ CWNetworkProfile) IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToNetworkProfile:"), networkProfile)
	return rv
}/* debug [instance_methods/method]: IsEqualToNetworkProfile */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWNetworkProfile */

// The security mode for the network profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/security
func (c_ CWNetworkProfile) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c_.ID, objc.Sel("security"))
	return rv
}/* debug [instance_properties/getter]: security */


// The service set identifier (SSID) for the network profile, encoded as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssid
func (c_ CWNetworkProfile) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ssid"))
	return rv
}/* debug [instance_properties/getter]: ssid */


// The service set identifier (SSID) for the network profile, returned as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssidData
func (c_ CWNetworkProfile) SsidData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("ssidData"))
	return rv
}/* debug [instance_properties/getter]: ssidData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWNetworkProfile */


