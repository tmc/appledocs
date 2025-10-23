// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AuthorizationAppleIDCredential] class.
type IAuthorizationAppleIDCredential interface {
	objectivec.IObject
	// properties:
	Email() string /* primitive/slice/pointer. */
	RealUserStatus() UserDetectionStatus /* not a class type */
	AuthorizationCode() foundation.objc.IObject /* cross-framework: Data */
	SetAuthorizationCode(value foundation.objc.IObject /* cross-framework: Data */)
	AuthorizedScopes() unsafe.Pointer
	SetAuthorizedScopes(value unsafe.Pointer)
	FullName() foundation.objc.IObject /* cross-framework: PersonNameComponents */
	SetFullName(value foundation.objc.IObject /* cross-framework: PersonNameComponents */)
	IdentityToken() foundation.objc.IObject /* cross-framework: Data */
	SetIdentityToken(value foundation.objc.IObject /* cross-framework: Data */)
	State() string /* primitive/slice/pointer. */
	SetState(value string /* primitive/slice/pointer. */)
	User() string /* primitive/slice/pointer. */
	SetUser(value string /* primitive/slice/pointer. */)
	UserAgeRange() UserAgeRange /* not a class type */
	SetUserAgeRange(value UserAgeRange /* not a class type */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDCredentialClass) Alloc() AuthorizationAppleIDCredential {
	rv := objc.Send[AuthorizationAppleIDCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The user’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/email
func (a_ AuthorizationAppleIDCredential) Email() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("email"))
	return rv
}


// A value that indicates whether the user appears to be a real person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/realUserStatus
func (a_ AuthorizationAppleIDCredential) RealUserStatus() UserDetectionStatus /* not a class type */ {
	rv := objc.Send[UserDetectionStatus](a_.ID, objc.Sel("realUserStatus"))
	return rv
}


// A token that the app uses to interact with the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizationcode
func (a_ AuthorizationAppleIDCredential) AuthorizationCode() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("authorizationCode"))
	return rv
}


// A token that the app uses to interact with the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizationcode
func (a_ AuthorizationAppleIDCredential) SetAuthorizationCode(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationCode:"), value)
}


// The contact information the user authorized your app to access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizedscopes
func (a_ AuthorizationAppleIDCredential) AuthorizedScopes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("authorizedScopes"))
	return rv
}


// The contact information the user authorized your app to access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizedscopes
func (a_ AuthorizationAppleIDCredential) SetAuthorizedScopes(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizedScopes:"), value)
}


// The user’s full name from their Apple ID or a user-submitted value provided from the Sign in with Apple UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/fullname
func (a_ AuthorizationAppleIDCredential) FullName() foundation.objc.IObject /* cross-framework: PersonNameComponents */ {
	rv := objc.Send[foundation.PersonNameComponents](a_.ID, objc.Sel("fullName"))
	return rv
}


// The user’s full name from their Apple ID or a user-submitted value provided from the Sign in with Apple UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/fullname
func (a_ AuthorizationAppleIDCredential) SetFullName(value foundation.objc.IObject /* cross-framework: PersonNameComponents */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullName:"), value)
}


// A JSON Web Token (JWT) that securely communicates information about the user to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/identitytoken
func (a_ AuthorizationAppleIDCredential) IdentityToken() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("identityToken"))
	return rv
}


// A JSON Web Token (JWT) that securely communicates information about the user to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/identitytoken
func (a_ AuthorizationAppleIDCredential) SetIdentityToken(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentityToken:"), value)
}


// An arbitrary string that your app provides to the request that generates the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/state
func (a_ AuthorizationAppleIDCredential) State() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("state"))
	return rv
}


// An arbitrary string that your app provides to the request that generates the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/state
func (a_ AuthorizationAppleIDCredential) SetState(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setState:"), objc.String(value))
}


// An identifier for the authenticated user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/user
func (a_ AuthorizationAppleIDCredential) User() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("user"))
	return rv
}


// An identifier for the authenticated user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/user
func (a_ AuthorizationAppleIDCredential) SetUser(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUser:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/useragerange
func (a_ AuthorizationAppleIDCredential) UserAgeRange() UserAgeRange /* not a class type */ {
	rv := objc.Send[UserAgeRange](a_.ID, objc.Sel("userAgeRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/useragerange
func (a_ AuthorizationAppleIDCredential) SetUserAgeRange(value UserAgeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserAgeRange:"), value)
}



