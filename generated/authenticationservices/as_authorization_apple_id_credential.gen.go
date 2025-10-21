// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (a_ AuthorizationAppleIDCredential) FullName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fullName"))
	return rv
}

// A value that indicates whether the user appears to be a real person.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDCredential/realUserStatus
func (a_ AuthorizationAppleIDCredential) RealUserStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("realUserStatus"))
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



