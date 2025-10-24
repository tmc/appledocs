// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest] class.
type IAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
	// properties:
	CredentialParameters() []AuthorizationPublicKeyCredentialParameters
	SetCredentialParameters(value []AuthorizationPublicKeyCredentialParameters)
	ExcludedCredentials() []AuthorizationSecurityKeyPublicKeyCredentialDescriptor
	SetExcludedCredentials(value []AuthorizationSecurityKeyPublicKeyCredentialDescriptor)
	ResidentKeyPreference() AuthorizationPublicKeyCredentialResidentKeyPreference /* typedef */
	SetResidentKeyPreference(value AuthorizationPublicKeyCredentialResidentKeyPreference /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequestClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
// The object for registering a new security key credential.
//
// Create an instance of this class when registering for a new credential using security key authorization.


// The object for registering a new security key credential.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */

// An array of parameters for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/credentialParameters
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) CredentialParameters() []AuthorizationPublicKeyCredentialParameters {
	rv := objc.Send[[]AuthorizationPublicKeyCredentialParameters](a_.ID, objc.Sel("credentialParameters"))
	return rv
}/* debug [instance_properties/getter]: credentialParameters */


// An array of parameters for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/credentialParameters
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetCredentialParameters(value []AuthorizationPublicKeyCredentialParameters) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCredentialParameters:"), nsArray)
}/* debug [instance_properties/setter]: credentialParameters */


// An array of excluded parameters for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/excludedCredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) ExcludedCredentials() []AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[[]AuthorizationSecurityKeyPublicKeyCredentialDescriptor](a_.ID, objc.Sel("excludedCredentials"))
	return rv
}/* debug [instance_properties/getter]: excludedCredentials */


// An array of excluded parameters for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/excludedCredentials
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetExcludedCredentials(value []AuthorizationSecurityKeyPublicKeyCredentialDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setExcludedCredentials:"), nsArray)
}/* debug [instance_properties/setter]: excludedCredentials */


// The preference that indicates where the resident key resides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/residentKeyPreference
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) ResidentKeyPreference() AuthorizationPublicKeyCredentialResidentKeyPreference /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("residentKeyPreference"))
	return rv
}/* debug [instance_properties/getter]: residentKeyPreference */


// The preference that indicates where the resident key resides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest/residentKeyPreference
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest) SetResidentKeyPreference(value AuthorizationPublicKeyCredentialResidentKeyPreference /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setResidentKeyPreference:"), value)
}/* debug [instance_properties/setter]: residentKeyPreference */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialRegistrationRequest */



