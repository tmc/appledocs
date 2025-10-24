// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasswordCredentialIdentity */


/* debug [class_header]: Header for ASPasswordCredentialIdentity */
// The class instance for the [PasswordCredentialIdentity] class.
var (
	PasswordCredentialIdentityClass     _PasswordCredentialIdentityClass
	PasswordCredentialIdentityClassOnce sync.Once
)

func getPasswordCredentialIdentityClass() _PasswordCredentialIdentityClass {
	PasswordCredentialIdentityClassOnce.Do(func() {
		PasswordCredentialIdentityClass = _PasswordCredentialIdentityClass{objc.GetClass("ASPasswordCredentialIdentity")}
	})
	return PasswordCredentialIdentityClass
}

type _PasswordCredentialIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasswordCredentialIdentity */
// An interface definition for the [PasswordCredentialIdentity] class.
type IPasswordCredentialIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasswordCredentialIdentity */
	// properties:
	Rank() int
	SetRank(value int)
	RecordIdentifier() objc.IObject /* cross-framework: NSString */
	ServiceIdentifier() IASCredentialServiceIdentifier
	User() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasswordCredentialIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasswordCredentialIdentity */
// Alloc allocates a new instance without initialization.
func (pc _PasswordCredentialIdentityClass) Alloc() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasswordCredentialIdentityClass) New() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasswordCredentialIdentity) Init() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasswordCredentialIdentity) Autorelease() PasswordCredentialIdentity {
	rv := objc.Send[PasswordCredentialIdentity](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasswordCredentialIdentity creates a new PasswordCredentialIdentity instance.
func NewPasswordCredentialIdentity() PasswordCredentialIdentity {
	return getPasswordCredentialIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasswordCredentialIdentity */
// A description that uniquely identifies a particular password credential.


// A description that uniquely identifies a particular password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity
type PasswordCredentialIdentity struct {
	objectivec.Object
}

// PasswordCredentialIdentityFrom constructs a [PasswordCredentialIdentity] from an unsafe.Pointer.
//
// A description that uniquely identifies a particular password credential.
func PasswordCredentialIdentityFrom(ptr unsafe.Pointer) PasswordCredentialIdentity {
	return PasswordCredentialIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasswordCredentialIdentity */

// Initializes a password credential identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/init(serviceIdentifier:user:recordIdentifier:)
func NewPasswordCredentialIdentityWithServiceIdentifierUserRecordIdentifier(serviceIdentifier IASCredentialServiceIdentifier, user objc.IObject /* cross-framework: NSString */, recordIdentifier objc.IObject /* cross-framework: NSString */) PasswordCredentialIdentity {
	instance := getPasswordCredentialIdentityClass().Alloc()
	rv := objc.Send[PasswordCredentialIdentity](instance.ID, objc.Sel("initWithServiceIdentifier:user:recordIdentifier:"), serviceIdentifier, user, recordIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasswordCredentialIdentityWithServiceIdentifierUserRecordIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasswordCredentialIdentity */

// Creates and returns a password credential identity object with a service identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/identityWithServiceIdentifier:user:recordIdentifier:
func (pc _PasswordCredentialIdentityClass) IdentityWithServiceIdentifierUserRecordIdentifier(serviceIdentifier IASCredentialServiceIdentifier, user objc.IObject /* cross-framework: NSString */, recordIdentifier objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("identityWithServiceIdentifier:user:recordIdentifier:"), serviceIdentifier, user, recordIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IdentityWithServiceIdentifierUserRecordIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasswordCredentialIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasswordCredentialIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasswordCredentialIdentity */

// An indicator that enables you to prioritze credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/rank
func (p_ PasswordCredentialIdentity) Rank() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rank"))
	return rv
}/* debug [instance_properties/getter]: rank */


// An indicator that enables you to prioritze credential identities relative to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/rank
func (p_ PasswordCredentialIdentity) SetRank(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRank:"), value)
}/* debug [instance_properties/setter]: rank */


// A string used to correlate this identity to a record in your app’s own database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/recordIdentifier
func (p_ PasswordCredentialIdentity) RecordIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("recordIdentifier"))
	return rv
}/* debug [instance_properties/getter]: recordIdentifier */


// An identifier that helps the system know with which apps or websites to associate this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/serviceIdentifier
func (p_ PasswordCredentialIdentity) ServiceIdentifier() IASCredentialServiceIdentifier {
	rv := objc.Send[CredentialServiceIdentifier](p_.ID, objc.Sel("serviceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: serviceIdentifier */


// The username associated with the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialIdentity/user
func (p_ PasswordCredentialIdentity) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasswordCredentialIdentity */


