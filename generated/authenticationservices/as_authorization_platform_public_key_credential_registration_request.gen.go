// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialRegistrationRequest] class.
type IAuthorizationPlatformPublicKeyCredentialRegistrationRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
	// properties:
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput
	SetLargeBlob(value IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput)
	Prf() IASAuthorizationPublicKeyCredentialPRFRegistrationInput
	SetPrf(value IASAuthorizationPublicKeyCredentialPRFRegistrationInput)
	RequestStyle() AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle
	SetRequestStyle(value AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialRegistrationRequestClass) Alloc() AuthorizationPlatformPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialRegistrationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialRegistrationRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/largeBlob-28v2m
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/largeBlob-28v2m
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetLargeBlob(value IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLargeBlob:"), value)
}/* debug [instance_properties/setter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/prf-8fus5
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) Prf() IASAuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("prf"))
	return rv
}/* debug [instance_properties/getter]: prf */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/prf-8fus5
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetPrf(value IASAuthorizationPublicKeyCredentialPRFRegistrationInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrf:"), value)
}/* debug [instance_properties/setter]: prf */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/requestStyle-swift.property
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) RequestStyle() AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle](a_.ID, objc.Sel("requestStyle"))
	return rv
}/* debug [instance_properties/getter]: requestStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/requestStyle-swift.property
func (a_ AuthorizationPlatformPublicKeyCredentialRegistrationRequest) SetRequestStyle(value AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestStyle:"), value)
}/* debug [instance_properties/setter]: requestStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest */



