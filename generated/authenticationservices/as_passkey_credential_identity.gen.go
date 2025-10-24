// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyCredentialIdentity */


/* debug [class_header]: Header for ASPasskeyCredentialIdentity */
// The class instance for the [PasskeyCredentialIdentity] class.
var (
	PasskeyCredentialIdentityClass     _PasskeyCredentialIdentityClass
	PasskeyCredentialIdentityClassOnce sync.Once
)

func getPasskeyCredentialIdentityClass() _PasskeyCredentialIdentityClass {
	PasskeyCredentialIdentityClassOnce.Do(func() {
		PasskeyCredentialIdentityClass = _PasskeyCredentialIdentityClass{objc.GetClass("ASPasskeyCredentialIdentity")}
	})
	return PasskeyCredentialIdentityClass
}

type _PasskeyCredentialIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyCredentialIdentity */
// An interface definition for the [PasskeyCredentialIdentity] class.
type IPasskeyCredentialIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyCredentialIdentity */
	// properties:
	CredentialID() objc.IObject /* cross-framework: NSData */
	Rank() int
	SetRank(value int)
	RecordIdentifier() objc.IObject /* cross-framework: NSString */
	RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */
	UserHandle() objc.IObject /* cross-framework: NSData */
	UserName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyCredentialIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyCredentialIdentity */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyCredentialIdentityClass) Alloc() PasskeyCredentialIdentity {
	rv := objc.Send[PasskeyCredentialIdentity](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyCredentialIdentityClass) New() PasskeyCredentialIdentity {
	rv := objc.Send[PasskeyCredentialIdentity](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyCredentialIdentity) Init() PasskeyCredentialIdentity {
	rv := objc.Send[PasskeyCredentialIdentity](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyCredentialIdentity) Autorelease() PasskeyCredentialIdentity {
	rv := objc.Send[PasskeyCredentialIdentity](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyCredentialIdentity creates a new PasskeyCredentialIdentity instance.
func NewPasskeyCredentialIdentity() PasskeyCredentialIdentity {
	return getPasskeyCredentialIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyCredentialIdentity */
// A description that uniquely identifies a particular passkey credential.


// A description that uniquely identifies a particular passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity
type PasskeyCredentialIdentity struct {
	objectivec.Object
}

// PasskeyCredentialIdentityFrom constructs a [PasskeyCredentialIdentity] from an unsafe.Pointer.
//
// A description that uniquely identifies a particular passkey credential.
func PasskeyCredentialIdentityFrom(ptr unsafe.Pointer) PasskeyCredentialIdentity {
	return PasskeyCredentialIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyCredentialIdentity */

// Initializes a passkey credential identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/initWithRelyingPartyIdentifier:userName:credentialID:userHandle:recordIdentifier:
func NewPasskeyCredentialIdentityWithRelyingPartyIdentifierUserNameCredentialIDUserHandleRecordIdentifier(relyingPartyIdentifier objc.IObject /* cross-framework: NSString */, userName objc.IObject /* cross-framework: NSString */, credentialID objc.IObject /* cross-framework: NSData */, userHandle objc.IObject /* cross-framework: NSData */, recordIdentifier objc.IObject /* cross-framework: NSString */) PasskeyCredentialIdentity {
	instance := getPasskeyCredentialIdentityClass().Alloc()
	rv := objc.Send[PasskeyCredentialIdentity](instance.ID, objc.Sel("initWithRelyingPartyIdentifier:userName:credentialID:userHandle:recordIdentifier:"), relyingPartyIdentifier, userName, credentialID, userHandle, recordIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyCredentialIdentityWithRelyingPartyIdentifierUserNameCredentialIDUserHandleRecordIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyCredentialIdentity */

// Creates and initializes a passkey credential identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/init(relyingPartyIdentifier:userName:credentialID:userHandle:recordIdentifier:)-9iuhb
func (pc _PasskeyCredentialIdentityClass) IdentityWithRelyingPartyIdentifierUserNameCredentialIDUserHandleRecordIdentifier(relyingPartyIdentifier objc.IObject /* cross-framework: NSString */, userName objc.IObject /* cross-framework: NSString */, credentialID objc.IObject /* cross-framework: NSData */, userHandle objc.IObject /* cross-framework: NSData */, recordIdentifier objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("identityWithRelyingPartyIdentifier:userName:credentialID:userHandle:recordIdentifier:"), relyingPartyIdentifier, userName, credentialID, userHandle, recordIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithRelyingPartyIdentifierUserNameCredentialIDUserHandleRecordIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyCredentialIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyCredentialIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyCredentialIdentity */

// The credential identifier for this passkey credential identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/credentialID
func (p_ PasskeyCredentialIdentity) CredentialID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("credentialID"))
	return rv
}/* debug [instance_properties/getter]: credentialID */


// An indicator that enables you to prioritize credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/rank
func (p_ PasskeyCredentialIdentity) Rank() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rank"))
	return rv
}/* debug [instance_properties/getter]: rank */


// An indicator that enables you to prioritize credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/rank
func (p_ PasskeyCredentialIdentity) SetRank(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRank:"), value)
}/* debug [instance_properties/setter]: rank */


// A string used to correlate this identity to a record in your app’s own database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/recordIdentifier
func (p_ PasskeyCredentialIdentity) RecordIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("recordIdentifier"))
	return rv
}/* debug [instance_properties/getter]: recordIdentifier */


// A string that identifies this identity’s relying party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/relyingPartyIdentifier
func (p_ PasskeyCredentialIdentity) RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("relyingPartyIdentifier"))
	return rv
}/* debug [instance_properties/getter]: relyingPartyIdentifier */


// The user handle of this passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/userHandle
func (p_ PasskeyCredentialIdentity) UserHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("userHandle"))
	return rv
}/* debug [instance_properties/getter]: userHandle */


// The username of this passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialIdentity/userName
func (p_ PasskeyCredentialIdentity) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyCredentialIdentity */


