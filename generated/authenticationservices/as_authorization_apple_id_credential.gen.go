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
	Email() string
	FullName() foundation.PersonNameComponents
	RealUserStatus() UserDetectionStatus
	State() string
	User() string
	AuthorizationCode() foundation.Data
	SetAuthorizationCode(value foundation.IData)
	AuthorizedScopes() unsafe.Pointer
	SetAuthorizedScopes(value unsafe.Pointer)
	IdentityToken() foundation.Data
	SetIdentityToken(value foundation.IData)
	UserAgeRange() UserAgeRange
	SetUserAgeRange(value IUserAgeRange)
}

// A credential that results from a successful Apple ID authentication.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/email
func (a_ AuthorizationAppleIDCredential) Email() string {
	rv := objc.Send[string](a_.ID, objc.Sel("email"))
	return rv
}

// The user’s full name from their Apple ID or a user-submitted value provided from the Sign in with Apple UI.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/fullName
func (a_ AuthorizationAppleIDCredential) FullName() foundation.PersonNameComponents {
	rv := objc.Send[foundation.PersonNameComponents](a_.ID, objc.Sel("fullName"))
	return rv
}

// A value that indicates whether the user appears to be a real person.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/realUserStatus
func (a_ AuthorizationAppleIDCredential) RealUserStatus() UserDetectionStatus {
	rv := objc.Send[UserDetectionStatus](a_.ID, objc.Sel("realUserStatus"))
	return rv
}

// An arbitrary string that your app provides to the request that generates the credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/state
func (a_ AuthorizationAppleIDCredential) State() string {
	rv := objc.Send[string](a_.ID, objc.Sel("state"))
	return rv
}

// An identifier for the authenticated user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/user
func (a_ AuthorizationAppleIDCredential) User() string {
	rv := objc.Send[string](a_.ID, objc.Sel("user"))
	return rv
}

// A token that the app uses to interact with the server.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizationcode
func (a_ AuthorizationAppleIDCredential) AuthorizationCode() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("authorizationCode"))
	return rv
}


// SetAuthorizationCode sets the value of the authorizationCode property.
// A token that the app uses to interact with the server.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizationcode
func (a_ AuthorizationAppleIDCredential) SetAuthorizationCode(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizationCode:"), value)
}

// The contact information the user authorized your app to access.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizedscopes
func (a_ AuthorizationAppleIDCredential) AuthorizedScopes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("authorizedScopes"))
	return rv
}


// SetAuthorizedScopes sets the value of the authorizedScopes property.
// The contact information the user authorized your app to access.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/authorizedscopes
func (a_ AuthorizationAppleIDCredential) SetAuthorizedScopes(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthorizedScopes:"), value)
}

// A JSON Web Token (JWT) that securely communicates information about the user to the app.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/identitytoken
func (a_ AuthorizationAppleIDCredential) IdentityToken() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("identityToken"))
	return rv
}


// SetIdentityToken sets the value of the identityToken property.
// A JSON Web Token (JWT) that securely communicates information about the user to the app.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/identitytoken
func (a_ AuthorizationAppleIDCredential) SetIdentityToken(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentityToken:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/useragerange
func (a_ AuthorizationAppleIDCredential) UserAgeRange() UserAgeRange {
	rv := objc.Send[UserAgeRange](a_.ID, objc.Sel("userAgeRange"))
	return rv
}


// SetUserAgeRange sets the value of the userAgeRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationappleidcredential/useragerange
func (a_ AuthorizationAppleIDCredential) SetUserAgeRange(value IUserAgeRange) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserAgeRange:"), value)
}



