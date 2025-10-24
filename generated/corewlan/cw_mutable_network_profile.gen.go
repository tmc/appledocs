// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CWMutableNetworkProfile */


/* debug [class_header]: Header for CWMutableNetworkProfile */
// The class instance for the [CWMutableNetworkProfile] class.
var (
	CWMutableNetworkProfileClass     _CWMutableNetworkProfileClass
	CWMutableNetworkProfileClassOnce sync.Once
)

func getCWMutableNetworkProfileClass() _CWMutableNetworkProfileClass {
	CWMutableNetworkProfileClassOnce.Do(func() {
		CWMutableNetworkProfileClass = _CWMutableNetworkProfileClass{objc.GetClass("CWMutableNetworkProfile")}
	})
	return CWMutableNetworkProfileClass
}

type _CWMutableNetworkProfileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWMutableNetworkProfile */
// An interface definition for the [CWMutableNetworkProfile] class.
type ICWMutableNetworkProfile interface {
	ICWNetworkProfile
	
/* debug [class_interface_properties]: Properties for CWMutableNetworkProfile */
	// properties:
	Security() CWSecurity
	SetSecurity(value CWSecurity)
	SsidData() objc.IObject /* cross-framework: NSData */
	SetSsidData(value objc.IObject /* cross-framework: NSData */)
	NetworkProfiles() foundation.OrderedSet
	SetNetworkProfiles(value foundation.OrderedSet)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWMutableNetworkProfile */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWMutableNetworkProfile */
// Alloc allocates a new instance without initialization.
func (cc _CWMutableNetworkProfileClass) Alloc() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CWMutableNetworkProfileClass) New() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWMutableNetworkProfile) Init() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWMutableNetworkProfile) Autorelease() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWMutableNetworkProfile creates a new CWMutableNetworkProfile instance.
func NewCWMutableNetworkProfile() CWMutableNetworkProfile {
	return getCWMutableNetworkProfileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWMutableNetworkProfile */
// Encapsulates a mutable network profile entry.
//
// Use this class to change profile properties. To commit Wi-Fi network profile changes, use and .


// Encapsulates a mutable network profile entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile
type CWMutableNetworkProfile struct {
	CWNetworkProfile
}

// CWMutableNetworkProfileFrom constructs a [CWMutableNetworkProfile] from an unsafe.Pointer.
//
// Encapsulates a mutable network profile entry.
func CWMutableNetworkProfileFrom(ptr unsafe.Pointer) CWMutableNetworkProfile {
	return CWMutableNetworkProfile{
		CWNetworkProfile: CWNetworkProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWMutableNetworkProfile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWMutableNetworkProfile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWMutableNetworkProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWMutableNetworkProfile */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWMutableNetworkProfile */

// The security type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/security
func (c_ CWMutableNetworkProfile) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c_.ID, objc.Sel("security"))
	return rv
}/* debug [instance_properties/getter]: security */


// The security type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/security
func (c_ CWMutableNetworkProfile) SetSecurity(value CWSecurity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurity:"), value)
}/* debug [instance_properties/setter]: security */


// The service set identifier (SSID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/ssidData
func (c_ CWMutableNetworkProfile) SsidData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("ssidData"))
	return rv
}/* debug [instance_properties/getter]: ssidData */


// The service set identifier (SSID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/ssidData
func (c_ CWMutableNetworkProfile) SetSsidData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSsidData:"), value)
}/* debug [instance_properties/setter]: ssidData */


// The preferred networks list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corewlan/cwmutableconfiguration/networkprofiles
func (c_ CWMutableNetworkProfile) NetworkProfiles() foundation.OrderedSet {
	rv := objc.Send[foundation.OrderedSet](c_.ID, objc.Sel("networkProfiles"))
	return rv
}/* debug [instance_properties/getter]: networkProfiles */


// The preferred networks list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corewlan/cwmutableconfiguration/networkprofiles
func (c_ CWMutableNetworkProfile) SetNetworkProfiles(value foundation.OrderedSet) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNetworkProfiles:"), value)
}/* debug [instance_properties/setter]: networkProfiles */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWMutableNetworkProfile */



