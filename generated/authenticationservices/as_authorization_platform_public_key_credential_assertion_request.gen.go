// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialAssertionRequest */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialAssertionRequest */
// The class instance for the [AuthorizationPlatformPublicKeyCredentialAssertionRequest] class.
var (
	AuthorizationPlatformPublicKeyCredentialAssertionRequestClass     _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass
	AuthorizationPlatformPublicKeyCredentialAssertionRequestClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialAssertionRequestClass() _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass {
	AuthorizationPlatformPublicKeyCredentialAssertionRequestClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialAssertionRequestClass = _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialAssertionRequest")}
	})
	return AuthorizationPlatformPublicKeyCredentialAssertionRequestClass
}

type _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialAssertionRequest] class.
type IAuthorizationPlatformPublicKeyCredentialAssertionRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
	// properties:
	AllowedCredentials() []AuthorizationPlatformPublicKeyCredentialDescriptor
	SetAllowedCredentials(value []AuthorizationPlatformPublicKeyCredentialDescriptor)
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput
	SetLargeBlob(value IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput)
	Prf() IASAuthorizationPublicKeyCredentialPRFAssertionInput
	SetPrf(value IASAuthorizationPublicKeyCredentialPRFAssertionInput)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass) Alloc() AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPlatformPublicKeyCredentialAssertionRequestClass) New() AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertionRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) Init() AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) Autorelease() AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertionRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialAssertionRequest creates a new AuthorizationPlatformPublicKeyCredentialAssertionRequest instance.
func NewAuthorizationPlatformPublicKeyCredentialAssertionRequest() AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	return getAuthorizationPlatformPublicKeyCredentialAssertionRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
// The concrete assertion request type for platform credentials.
//
// Use this class to sign in with an existing credential that the system stores in iCloud Keychain.


// The concrete assertion request type for platform credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest
type AuthorizationPlatformPublicKeyCredentialAssertionRequest struct {
	AuthorizationRequest
}

// AuthorizationPlatformPublicKeyCredentialAssertionRequestFrom constructs a [AuthorizationPlatformPublicKeyCredentialAssertionRequest] from an unsafe.Pointer.
//
// The concrete assertion request type for platform credentials.
func AuthorizationPlatformPublicKeyCredentialAssertionRequestFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialAssertionRequest {
	return AuthorizationPlatformPublicKeyCredentialAssertionRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialAssertionRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialAssertionRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialAssertionRequest */

// The array of allowed credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/allowedCredentials
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) AllowedCredentials() []AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[[]AuthorizationPlatformPublicKeyCredentialDescriptor](a_.ID, objc.Sel("allowedCredentials"))
	return rv
}/* debug [instance_properties/getter]: allowedCredentials */


// The array of allowed credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/allowedCredentials
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) SetAllowedCredentials(value []AuthorizationPlatformPublicKeyCredentialDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowedCredentials:"), nsArray)
}/* debug [instance_properties/setter]: allowedCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/largeBlob-5mg1q
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](a_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/largeBlob-5mg1q
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) SetLargeBlob(value IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLargeBlob:"), value)
}/* debug [instance_properties/setter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/prf-60tle
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) Prf() IASAuthorizationPublicKeyCredentialPRFAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInput](a_.ID, objc.Sel("prf"))
	return rv
}/* debug [instance_properties/getter]: prf */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertionRequest/prf-60tle
func (a_ AuthorizationPlatformPublicKeyCredentialAssertionRequest) SetPrf(value IASAuthorizationPublicKeyCredentialPRFAssertionInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrf:"), value)
}/* debug [instance_properties/setter]: prf */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialAssertionRequest */



