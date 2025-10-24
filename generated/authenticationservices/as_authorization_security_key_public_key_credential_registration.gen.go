// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialRegistration */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialRegistration */
// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialRegistration] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass     _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialRegistrationClass() _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass {
	AuthorizationSecurityKeyPublicKeyCredentialRegistrationClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass = _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialRegistration")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialRegistration] class.
type IAuthorizationSecurityKeyPublicKeyCredentialRegistration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
	// properties:
	Transports() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSecurityKeyPublicKeyCredentialRegistrationClass) New() AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistration) Init() AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistration) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialRegistration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialRegistration creates a new AuthorizationSecurityKeyPublicKeyCredentialRegistration instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialRegistration() AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	return getAuthorizationSecurityKeyPublicKeyCredentialRegistrationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
// A newly created security key credential that results from a credential registration request.
//
// Use this class to verify a successful security key authorization request in .


// A newly created security key credential that results from a credential registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistration
type AuthorizationSecurityKeyPublicKeyCredentialRegistration struct {
	objectivec.Object
}

// AuthorizationSecurityKeyPublicKeyCredentialRegistrationFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialRegistration] from an unsafe.Pointer.
//
// A newly created security key credential that results from a credential registration request.
func AuthorizationSecurityKeyPublicKeyCredentialRegistrationFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialRegistration {
	return AuthorizationSecurityKeyPublicKeyCredentialRegistration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialRegistration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialRegistration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialRegistration */

// An array of transport types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialRegistration/transports
func (a_ AuthorizationSecurityKeyPublicKeyCredentialRegistration) Transports() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("transports"))
	return rv
}/* debug [instance_properties/getter]: transports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialRegistration */



