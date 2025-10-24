// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialAssertion */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialAssertion */
// The class instance for the [AuthorizationPlatformPublicKeyCredentialAssertion] class.
var (
	AuthorizationPlatformPublicKeyCredentialAssertionClass     _AuthorizationPlatformPublicKeyCredentialAssertionClass
	AuthorizationPlatformPublicKeyCredentialAssertionClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialAssertionClass() _AuthorizationPlatformPublicKeyCredentialAssertionClass {
	AuthorizationPlatformPublicKeyCredentialAssertionClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialAssertionClass = _AuthorizationPlatformPublicKeyCredentialAssertionClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialAssertion")}
	})
	return AuthorizationPlatformPublicKeyCredentialAssertionClass
}

type _AuthorizationPlatformPublicKeyCredentialAssertionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialAssertion */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialAssertion] class.
type IAuthorizationPlatformPublicKeyCredentialAssertion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialAssertion */
	// properties:
	Attachment() AuthorizationPublicKeyCredentialAttachment
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput
	Prf() IASAuthorizationPublicKeyCredentialPRFAssertionOutput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialAssertion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialAssertion */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialAssertionClass) Alloc() AuthorizationPlatformPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPlatformPublicKeyCredentialAssertionClass) New() AuthorizationPlatformPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertion](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialAssertion) Init() AuthorizationPlatformPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertion](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialAssertion) Autorelease() AuthorizationPlatformPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialAssertion](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialAssertion creates a new AuthorizationPlatformPublicKeyCredentialAssertion instance.
func NewAuthorizationPlatformPublicKeyCredentialAssertion() AuthorizationPlatformPublicKeyCredentialAssertion {
	return getAuthorizationPlatformPublicKeyCredentialAssertionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialAssertion */
// A class that represents the platform credential assertion type.
//
// The device creates an assertion when signing in with an existing credential. Use this class to verify the platform credential assertion when the authorization controller calls .


// A class that represents the platform credential assertion type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertion
type AuthorizationPlatformPublicKeyCredentialAssertion struct {
	objectivec.Object
}

// AuthorizationPlatformPublicKeyCredentialAssertionFrom constructs a [AuthorizationPlatformPublicKeyCredentialAssertion] from an unsafe.Pointer.
//
// A class that represents the platform credential assertion type.
func AuthorizationPlatformPublicKeyCredentialAssertionFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialAssertion {
	return AuthorizationPlatformPublicKeyCredentialAssertion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialAssertion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialAssertion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialAssertion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialAssertion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialAssertion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertion/attachment
func (a_ AuthorizationPlatformPublicKeyCredentialAssertion) Attachment() AuthorizationPublicKeyCredentialAttachment {
	rv := objc.Send[AuthorizationPublicKeyCredentialAttachment](a_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertion/largeBlob-97tbp
func (a_ AuthorizationPlatformPublicKeyCredentialAssertion) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](a_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialAssertion/prf-8blir
func (a_ AuthorizationPlatformPublicKeyCredentialAssertion) Prf() IASAuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](a_.ID, objc.Sel("prf"))
	return rv
}/* debug [instance_properties/getter]: prf */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialAssertion */



