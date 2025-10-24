// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBIdentityAuthority */


/* debug [class_header]: Header for CBIdentityAuthority */
// The class instance for the [CBIdentityAuthority] class.
var (
	CBIdentityAuthorityClass     _CBIdentityAuthorityClass
	CBIdentityAuthorityClassOnce sync.Once
)

func getCBIdentityAuthorityClass() _CBIdentityAuthorityClass {
	CBIdentityAuthorityClassOnce.Do(func() {
		CBIdentityAuthorityClass = _CBIdentityAuthorityClass{objc.GetClass("CBIdentityAuthority")}
	})
	return CBIdentityAuthorityClass
}

type _CBIdentityAuthorityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBIdentityAuthority */
// An interface definition for the [CBIdentityAuthority] class.
type ICBIdentityAuthority interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBIdentityAuthority */
	// properties:
	CSIdentityAuthority() unsafe.Pointer
	LocalizedName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBIdentityAuthority */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBIdentityAuthority */
// Alloc allocates a new instance without initialization.
func (cc _CBIdentityAuthorityClass) Alloc() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBIdentityAuthorityClass) New() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBIdentityAuthority) Init() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBIdentityAuthority) Autorelease() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBIdentityAuthority creates a new CBIdentityAuthority instance.
func NewCBIdentityAuthority() CBIdentityAuthority {
	return getCBIdentityAuthorityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBIdentityAuthority */
// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.


// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority
type CBIdentityAuthority struct {
	objectivec.Object
}

// CBIdentityAuthorityFrom constructs a [CBIdentityAuthority] from an unsafe.Pointer.
//
// An identity authority is a database that stores information about identities. The class defines one or more identity authorities. You can search this database for identities in conjunction with the class factory methods.
func CBIdentityAuthorityFrom(ptr unsafe.Pointer) CBIdentityAuthority {
	return CBIdentityAuthority{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBIdentityAuthority *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBIdentityAuthority */

// Returns an identity authority that contains the identities in both the local and the network-bound authorities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/default()
func (cc _CBIdentityAuthorityClass) DefaultIdentityAuthority() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("defaultIdentityAuthority"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultIdentityAuthority) */


// Returns an identity authority specified by a given Core Services Identity authority object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/identityAuthorityWithCSIdentityAuthority:
func (cc _CBIdentityAuthorityClass) IdentityAuthorityWithCSIdentityAuthority(CSIdentityAuthority unsafe.Pointer) CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("identityAuthorityWithCSIdentityAuthority:"), CSIdentityAuthority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityAuthorityWithCSIdentityAuthority) */


// Returns the identity authority on the local system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/local()
func (cc _CBIdentityAuthorityClass) LocalIdentityAuthority() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("localIdentityAuthority"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalIdentityAuthority) */


// Returns the identity authority that contains all the identities in bound network directory servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/managed()
func (cc _CBIdentityAuthorityClass) ManagedIdentityAuthority() CBIdentityAuthority {
	rv := objc.Send[CBIdentityAuthority](objc.ID(cc.class), objc.Sel("managedIdentityAuthority"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ManagedIdentityAuthority) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBIdentityAuthority */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBIdentityAuthority */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBIdentityAuthority */

// Returns an identity authority for use with the Core Services Identity API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/CSIdentityAuthority
func (c_ CBIdentityAuthority) CSIdentityAuthority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("CSIdentityAuthority"))
	return rv
}/* debug [instance_properties/getter]: CSIdentityAuthority */


// Returns the localized name of the identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityAuthority/localizedName
func (c_ CBIdentityAuthority) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBIdentityAuthority */



