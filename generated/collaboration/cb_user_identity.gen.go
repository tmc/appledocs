// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CBUserIdentity */


/* debug [class_header]: Header for CBUserIdentity */
// The class instance for the [CBUserIdentity] class.
var (
	CBUserIdentityClass     _CBUserIdentityClass
	CBUserIdentityClassOnce sync.Once
)

func getCBUserIdentityClass() _CBUserIdentityClass {
	CBUserIdentityClassOnce.Do(func() {
		CBUserIdentityClass = _CBUserIdentityClass{objc.GetClass("CBUserIdentity")}
	})
	return CBUserIdentityClass
}

type _CBUserIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBUserIdentity */
// An interface definition for the [CBUserIdentity] class.
type ICBUserIdentity interface {
	ICBIdentity
	
/* debug [class_interface_properties]: Properties for CBUserIdentity */
	// properties:
	Certificate() unsafe.Pointer
	Enabled() bool
	PosixUID() unsafe.Pointer
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBUserIdentity */
	// methods:
	AuthenticateWithPassword(password objc.IObject /* cross-framework: NSString */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBUserIdentity */
// Alloc allocates a new instance without initialization.
func (cc _CBUserIdentityClass) Alloc() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBUserIdentityClass) New() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBUserIdentity) Init() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBUserIdentity) Autorelease() CBUserIdentity {
	rv := objc.Send[CBUserIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBUserIdentity creates a new CBUserIdentity instance.
func NewCBUserIdentity() CBUserIdentity {
	return getCBUserIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBUserIdentity */
// An object of the class represents a user identity and is used for accessing the attributes of a user identity from an identity authority. The principal attributes of are a POSIX user identifier (UID), password, and certificate.


// An object of the class represents a user identity and is used for accessing the attributes of a user identity from an identity authority. The principal attributes of are a POSIX user identifier (UID), password, and certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity
type CBUserIdentity struct {
	CBIdentity
}

// CBUserIdentityFrom constructs a [CBUserIdentity] from an unsafe.Pointer.
//
// An object of the class represents a user identity and is used for accessing the attributes of a user identity from an identity authority. The principal attributes of are a POSIX user identifier (UID), password, and certificate.
func CBUserIdentityFrom(ptr unsafe.Pointer) CBUserIdentity {
	return CBUserIdentity{
		CBIdentity: CBIdentityFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBUserIdentity */

// Returns the user identity with the given POSIX UID in the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/init(posixUID:authority:)
func NewCBUserIdentityWithPosixUIDAuthority(uid unsafe.Pointer, authority ICBIdentityAuthority) CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(getCBUserIdentityClass().class), objc.Sel("userIdentityWithPosixUID:authority:"), uid, authority)
	return rv
}/* debug [class_init_methods/constructor]: NewCBUserIdentityWithPosixUIDAuthority */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBUserIdentity */

// Returns the user identity with the given POSIX UID in the specified identity authority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/init(posixUID:authority:)
func (cc _CBUserIdentityClass) UserIdentityWithPosixUIDAuthority(uid unsafe.Pointer, authority ICBIdentityAuthority) CBUserIdentity {
	rv := objc.Send[CBUserIdentity](objc.ID(cc.class), objc.Sel("userIdentityWithPosixUID:authority:"), uid, authority)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UserIdentityWithPosixUIDAuthority) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBUserIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBUserIdentity */

// Returns a Boolean value indicating whether the given password is correct for the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/authenticate(withPassword:)
func (c_ CBUserIdentity) AuthenticateWithPassword(password objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("authenticateWithPassword:"), password)
	return rv
}/* debug [instance_methods/method]: AuthenticateWithPassword */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBUserIdentity */

// Returns the public authentication certificate associated with a user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/certificate
func (c_ CBUserIdentity) Certificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("certificate"))
	return rv
}/* debug [instance_properties/getter]: certificate */


// Returns a Boolean value indicating whether the identity is allowed to authenticate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/isEnabled
func (c_ CBUserIdentity) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// Returns the POSIX UID of the identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBUserIdentity/posixUID
func (c_ CBUserIdentity) PosixUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("posixUID"))
	return rv
}/* debug [instance_properties/getter]: posixUID */


// Returns a Boolean value indicating whether the identity is allowed to authenticate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbuseridentity/isenabled
func (c_ CBUserIdentity) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// Returns a Boolean value indicating whether the identity is allowed to authenticate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/collaboration/cbuseridentity/isenabled
func (c_ CBUserIdentity) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBUserIdentity */


