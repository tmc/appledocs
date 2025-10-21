// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass     _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass() _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass {
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass = _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] class.
type IAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest interface {
	IAuthorizationRequest
}

// The object for registering a new security key credential.
//
// Create an instance of this class when registering for a new credential using security key authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest
type AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest struct {
	AuthorizationRequest
}

// AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] from an unsafe.Pointer.
//
// The object for registering a new security key credential.
func AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	return AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass) New() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) Init() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest creates a new AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	return getAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass().New()
}


// An array of parameters for the credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/credentialparameters
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) CredentialParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("credentialParameters"))
	return rv
}


// SetCredentialParameters sets the value of the credentialParameters property.
// An array of parameters for the credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/credentialparameters
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetCredentialParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCredentialParameters:"), value)
}

// An array of excluded parameters for the credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/excludedcredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) ExcludedCredentials() ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor](a_.ID, objc.Sel("excludedCredentials"))
	return rv
}


// SetExcludedCredentials sets the value of the excludedCredentials property.
// An array of excluded parameters for the credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/excludedcredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetExcludedCredentials(value IASAuthorizationSecurityKeyPublicKeyCredentialDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExcludedCredentials:"), value)
}

// The preference that indicates where the resident key resides.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/residentkeypreference
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) ResidentKeyPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("residentKeyPreference"))
	return rv
}


// SetResidentKeyPreference sets the value of the residentKeyPreference property.
// The preference that indicates where the resident key resides.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsecuritykeypublickeycredentialregistrationrequest/residentkeypreference
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetResidentKeyPreference(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setResidentKeyPreference:"), value)
}



