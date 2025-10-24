// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialRegistration */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialRegistration */
// The class instance for the [AuthorizationPlatformPublicKeyCredentialRegistration] class.
var (
	AuthorizationPlatformPublicKeyCredentialRegistrationClass     _AuthorizationPlatformPublicKeyCredentialRegistrationClass
	AuthorizationPlatformPublicKeyCredentialRegistrationClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialRegistrationClass() _AuthorizationPlatformPublicKeyCredentialRegistrationClass {
	AuthorizationPlatformPublicKeyCredentialRegistrationClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialRegistrationClass = _AuthorizationPlatformPublicKeyCredentialRegistrationClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialRegistration")}
	})
	return AuthorizationPlatformPublicKeyCredentialRegistrationClass
}

type _AuthorizationPlatformPublicKeyCredentialRegistrationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialRegistration */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialRegistration] class.
type IAuthorizationPlatformPublicKeyCredentialRegistration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialRegistration */
	// properties:
	Attachment() AuthorizationPublicKeyCredentialAttachment
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput
	Prf() IASAuthorizationPublicKeyCredentialPRFRegistrationOutput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialRegistration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialRegistration */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialRegistrationClass) Alloc() AuthorizationPlatformPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPlatformPublicKeyCredentialRegistrationClass) New() AuthorizationPlatformPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialRegistration) Init() AuthorizationPlatformPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialRegistration) Autorelease() AuthorizationPlatformPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialRegistration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialRegistration creates a new AuthorizationPlatformPublicKeyCredentialRegistration instance.
func NewAuthorizationPlatformPublicKeyCredentialRegistration() AuthorizationPlatformPublicKeyCredentialRegistration {
	return getAuthorizationPlatformPublicKeyCredentialRegistrationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialRegistration */
// A newly created platform credential that results from a credential registration request.
//
// Use this class to verify a successful platform authorization request in .


// A newly created platform credential that results from a credential registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistration
type AuthorizationPlatformPublicKeyCredentialRegistration struct {
	objectivec.Object
}

// AuthorizationPlatformPublicKeyCredentialRegistrationFrom constructs a [AuthorizationPlatformPublicKeyCredentialRegistration] from an unsafe.Pointer.
//
// A newly created platform credential that results from a credential registration request.
func AuthorizationPlatformPublicKeyCredentialRegistrationFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialRegistration {
	return AuthorizationPlatformPublicKeyCredentialRegistration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialRegistration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialRegistration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialRegistration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialRegistration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialRegistration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistration/attachment
func (a_ AuthorizationPlatformPublicKeyCredentialRegistration) Attachment() AuthorizationPublicKeyCredentialAttachment {
	rv := objc.Send[AuthorizationPublicKeyCredentialAttachment](a_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistration/largeBlob-jhnw
func (a_ AuthorizationPlatformPublicKeyCredentialRegistration) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](a_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistration/prf-49gbm
func (a_ AuthorizationPlatformPublicKeyCredentialRegistration) Prf() IASAuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](a_.ID, objc.Sel("prf"))
	return rv
}/* debug [instance_properties/getter]: prf */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialRegistration */



