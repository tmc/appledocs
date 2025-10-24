// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CBGroupIdentity */


/* debug [class_header]: Header for CBGroupIdentity */
// The class instance for the [CBGroupIdentity] class.
var (
	CBGroupIdentityClass     _CBGroupIdentityClass
	CBGroupIdentityClassOnce sync.Once
)

func getCBGroupIdentityClass() _CBGroupIdentityClass {
	CBGroupIdentityClassOnce.Do(func() {
		CBGroupIdentityClass = _CBGroupIdentityClass{objc.GetClass("CBGroupIdentity")}
	})
	return CBGroupIdentityClass
}

type _CBGroupIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBGroupIdentity */
// An interface definition for the [CBGroupIdentity] class.
type ICBGroupIdentity interface {
	ICBIdentity
	
/* debug [class_interface_properties]: Properties for CBGroupIdentity */
	// properties:
	MemberIdentities() []CBIdentity
	Members() objc.IObject /* cross-framework: NSArray */
	PosixGID() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBGroupIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBGroupIdentity */
// Alloc allocates a new instance without initialization.
func (cc _CBGroupIdentityClass) Alloc() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBGroupIdentityClass) New() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBGroupIdentity) Init() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBGroupIdentity) Autorelease() CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBGroupIdentity creates a new CBGroupIdentity instance.
func NewCBGroupIdentity() CBGroupIdentity {
	return getCBGroupIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBGroupIdentity */
// An object of the class represents a group identity and is used for viewing the attributes of group identities from an identity authority. The principal attributes of a object are a POSIX group identifier (GID) and a list of members.


// An object of the class represents a group identity and is used for viewing the attributes of group identities from an identity authority. The principal attributes of a object are a POSIX group identifier (GID) and a list of members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity
type CBGroupIdentity struct {
	CBIdentity
}

// CBGroupIdentityFrom constructs a [CBGroupIdentity] from an unsafe.Pointer.
//
// An object of the class represents a group identity and is used for viewing the attributes of group identities from an identity authority. The principal attributes of a object are a POSIX group identifier (GID) and a list of members.
func CBGroupIdentityFrom(ptr unsafe.Pointer) CBGroupIdentity {
	return CBGroupIdentity{
		CBIdentity: CBIdentityFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBGroupIdentity */

// Returns the group identity with the given POSIX GID in the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/init(posixGID:authority:)
func NewCBGroupIdentityWithPosixGIDAuthority(gid unsafe.Pointer, authority ICBIdentityAuthority) CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(getCBGroupIdentityClass().class), objc.Sel("groupIdentityWithPosixGID:authority:"), gid, authority)
	return rv
}/* debug [class_init_methods/constructor]: NewCBGroupIdentityWithPosixGIDAuthority */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBGroupIdentity */

// Returns the group identity with the given POSIX GID in the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/init(posixGID:authority:)
func (cc _CBGroupIdentityClass) GroupIdentityWithPosixGIDAuthority(gid unsafe.Pointer, authority ICBIdentityAuthority) CBGroupIdentity {
	rv := objc.Send[CBGroupIdentity](objc.ID(cc.class), objc.Sel("groupIdentityWithPosixGID:authority:"), gid, authority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GroupIdentityWithPosixGIDAuthority) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBGroupIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBGroupIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBGroupIdentity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/memberIdentities
func (c_ CBGroupIdentity) MemberIdentities() []CBIdentity {
	rv := objc.Send[[]CBIdentity](c_.ID, objc.Sel("memberIdentities"))
	return rv
}/* debug [instance_properties/getter]: memberIdentities */


// Returns the members of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/members
func (c_ CBGroupIdentity) Members() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](c_.ID, objc.Sel("members"))
	return rv
}/* debug [instance_properties/getter]: members */


// Returns the POSIX GID of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBGroupIdentity/posixGID
func (c_ CBGroupIdentity) PosixGID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("posixGID"))
	return rv
}/* debug [instance_properties/getter]: posixGID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBGroupIdentity */


