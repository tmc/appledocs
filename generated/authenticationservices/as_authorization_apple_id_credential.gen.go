// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationAppleIDCredential */


/* debug [class_header]: Header for ASAuthorizationAppleIDCredential */
// The class instance for the [AuthorizationAppleIDCredential] class.
var (
	AuthorizationAppleIDCredentialClass     _AuthorizationAppleIDCredentialClass
	AuthorizationAppleIDCredentialClassOnce sync.Once
)

func getAuthorizationAppleIDCredentialClass() _AuthorizationAppleIDCredentialClass {
	AuthorizationAppleIDCredentialClassOnce.Do(func() {
		AuthorizationAppleIDCredentialClass = _AuthorizationAppleIDCredentialClass{objc.GetClass("ASAuthorizationAppleIDCredential")}
	})
	return AuthorizationAppleIDCredentialClass
}

type _AuthorizationAppleIDCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationAppleIDCredential */
// An interface definition for the [AuthorizationAppleIDCredential] class.
type IAuthorizationAppleIDCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationAppleIDCredential */
	// properties:
	AuthorizationCode() objc.IObject /* cross-framework: NSData */
	AuthorizedScopes() []string
	Email() objc.IObject /* cross-framework: NSString */
	FullName() foundation.PersonNameComponents
	IdentityToken() objc.IObject /* cross-framework: NSData */
	RealUserStatus() UserDetectionStatus
	State() objc.IObject /* cross-framework: NSString */
	User() objc.IObject /* cross-framework: NSString */
	UserAgeRange() UserAgeRange
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationAppleIDCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationAppleIDCredential */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDCredentialClass) Alloc() AuthorizationAppleIDCredential {
	rv := objc.Send[AuthorizationAppleIDCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationAppleIDCredentialClass) New() AuthorizationAppleIDCredential {
	rv := objc.Send[AuthorizationAppleIDCredential](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDCredential) Init() AuthorizationAppleIDCredential {
	rv := objc.Send[AuthorizationAppleIDCredential](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDCredential) Autorelease() AuthorizationAppleIDCredential {
	rv := objc.Send[AuthorizationAppleIDCredential](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDCredential creates a new AuthorizationAppleIDCredential instance.
func NewAuthorizationAppleIDCredential() AuthorizationAppleIDCredential {
	return getAuthorizationAppleIDCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationAppleIDCredential */
// A credential that results from a successful Apple ID authentication.


// A credential that results from a successful Apple ID authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential
type AuthorizationAppleIDCredential struct {
	objectivec.Object
}

// AuthorizationAppleIDCredentialFrom constructs a [AuthorizationAppleIDCredential] from an unsafe.Pointer.
//
// A credential that results from a successful Apple ID authentication.
func AuthorizationAppleIDCredentialFrom(ptr unsafe.Pointer) AuthorizationAppleIDCredential {
	return AuthorizationAppleIDCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationAppleIDCredential *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationAppleIDCredential */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationAppleIDCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationAppleIDCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationAppleIDCredential */

// A token that the app uses to interact with the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/authorizationCode
func (a_ AuthorizationAppleIDCredential) AuthorizationCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("authorizationCode"))
	return rv
}/* debug [instance_properties/getter]: authorizationCode */


// The contact information the user authorized your app to access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/authorizedScopes
func (a_ AuthorizationAppleIDCredential) AuthorizedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("authorizedScopes"))
	return rv
}/* debug [instance_properties/getter]: authorizedScopes */


// The user’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/email
func (a_ AuthorizationAppleIDCredential) Email() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("email"))
	return rv
}/* debug [instance_properties/getter]: email */


// The user’s full name from their Apple ID or a user-submitted value provided from the Sign in with Apple UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/fullName
func (a_ AuthorizationAppleIDCredential) FullName() foundation.PersonNameComponents {
	rv := objc.Send[foundation.PersonNameComponents](a_.ID, objc.Sel("fullName"))
	return rv
}/* debug [instance_properties/getter]: fullName */


// A JSON Web Token (JWT) that securely communicates information about the user to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/identityToken
func (a_ AuthorizationAppleIDCredential) IdentityToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("identityToken"))
	return rv
}/* debug [instance_properties/getter]: identityToken */


// A value that indicates whether the user appears to be a real person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/realUserStatus
func (a_ AuthorizationAppleIDCredential) RealUserStatus() UserDetectionStatus {
	rv := objc.Send[UserDetectionStatus](a_.ID, objc.Sel("realUserStatus"))
	return rv
}/* debug [instance_properties/getter]: realUserStatus */


// An arbitrary string that your app provides to the request that generates the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/state
func (a_ AuthorizationAppleIDCredential) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// An identifier for the authenticated user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/user
func (a_ AuthorizationAppleIDCredential) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/userAgeRange
func (a_ AuthorizationAppleIDCredential) UserAgeRange() UserAgeRange {
	rv := objc.Send[UserAgeRange](a_.ID, objc.Sel("userAgeRange"))
	return rv
}/* debug [instance_properties/getter]: userAgeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationAppleIDCredential */



