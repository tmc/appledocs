// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationPlatformPublicKeyCredentialRegistrationRequest] class.
var (
	AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass     _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass
	AuthorizationPlatformPublicKeyCredentialRegistrationRequestClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialRegistrationRequestClass() _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass {
	AuthorizationPlatformPublicKeyCredentialRegistrationRequestClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass = _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest")}
	})
	return AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass
}

type _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPlatformPublicKeyCredentialRegistrationRequest] class.
type IAuthorizationPlatformPublicKeyCredentialRegistrationRequest interface {
	IAuthorizationRequest
	// properties:
	LargeBlob() IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput
	SetLargeBlob(value IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput)
	Prf() IAuthorizationPublicKeyCredentialPRFRegistrationInput
	SetPrf(value IAuthorizationPublicKeyCredentialPRFRegistrationInput)
	RequestStyle() unsafe.Pointer
	SetRequestStyle(value unsafe.Pointer)
	// methods:
}

// The object for registering a new platform public key credential.
//
// Create an instance of this class when registering for a new credential using platform authorization.


// The object for registering a new platform public key credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest
type AuthorizationPlatformPublicKeyCredentialRegistrationRequest struct {
	AuthorizationRequest
}

// AuthorizationPlatformPublicKeyCredentialRegistrationRequestFrom constructs a [AuthorizationPlatformPublicKeyCredentialRegistrationRequest] from an unsafe.Pointer.
//
// The object for registering a new platform public key credential.
func AuthorizationPlatformPublicKeyCredentialRegistrationRequestFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	return AuthorizationPlatformPublicKeyCredentialRegistrationRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass) Alloc() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass) New() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) Init() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) Autorelease() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialRegistrationRequest creates a new AuthorizationPlatformPublicKeyCredentialRegistrationRequest instance.
func NewAuthorizationPlatformPublicKeyCredentialRegistrationRequest() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	return getAuthorizationPlatformPublicKeyCredentialRegistrationRequestClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/largeblob-5ismm
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) LargeBlob() IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("largeBlob"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/largeblob-5ismm
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetLargeBlob(value IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLargeBlob:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/prf-3d9iw
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) Prf() IAuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("prf"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/prf-3d9iw
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetPrf(value IAuthorizationPublicKeyCredentialPRFRegistrationInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrf:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/requeststyle-swift.property
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) RequestStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("requestStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationplatformpublickeycredentialregistrationrequest/requeststyle-swift.property
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetRequestStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestStyle:"), value)
}



