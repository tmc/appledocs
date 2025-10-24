// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass     _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass() _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass {
	AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass = _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] class.
type IAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
	// properties:
	IsSupported() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass) New() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) Init() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput creates a new AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput() AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	return getAuthorizationPublicKeyCredentialLargeBlobRegistrationOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput-c.class
type AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobRegistrationOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput-c.class/isSupported
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}/* debug [instance_properties/getter]: isSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput */



