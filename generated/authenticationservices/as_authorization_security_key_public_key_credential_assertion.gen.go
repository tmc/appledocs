// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialAssertion */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialAssertion */
// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialAssertion] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialAssertionClass     _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass
	AuthorizationSecurityKeyPublicKeyCredentialAssertionClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialAssertionClass() _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass {
	AuthorizationSecurityKeyPublicKeyCredentialAssertionClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialAssertionClass = _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialAssertion")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialAssertionClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialAssertion] class.
type IAuthorizationSecurityKeyPublicKeyCredentialAssertion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
	// properties:
	AppID() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSecurityKeyPublicKeyCredentialAssertionClass) New() AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertion](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertion) Init() AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertion](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertion) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialAssertion](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialAssertion creates a new AuthorizationSecurityKeyPublicKeyCredentialAssertion instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialAssertion() AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	return getAuthorizationSecurityKeyPublicKeyCredentialAssertionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
// A class that represents the security key credential assertion type.
//
// The security key creates an assertion when signing in with an existing credential. Use this class to verify the security key credential assertion when the authorization controller calls .


// A class that represents the security key credential assertion type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertion
type AuthorizationSecurityKeyPublicKeyCredentialAssertion struct {
	objectivec.Object
}

// AuthorizationSecurityKeyPublicKeyCredentialAssertionFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialAssertion] from an unsafe.Pointer.
//
// A class that represents the security key credential assertion type.
func AuthorizationSecurityKeyPublicKeyCredentialAssertionFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialAssertion {
	return AuthorizationSecurityKeyPublicKeyCredentialAssertion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialAssertion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialAssertion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialAssertion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialAssertion/appID
func (a_ AuthorizationSecurityKeyPublicKeyCredentialAssertion) AppID() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appID"))
	return rv
}/* debug [instance_properties/getter]: appID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialAssertion */



