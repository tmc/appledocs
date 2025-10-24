// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPasswordProvider */


/* debug [class_header]: Header for ASAuthorizationPasswordProvider */
// The class instance for the [AuthorizationPasswordProvider] class.
var (
	AuthorizationPasswordProviderClass     _AuthorizationPasswordProviderClass
	AuthorizationPasswordProviderClassOnce sync.Once
)

func getAuthorizationPasswordProviderClass() _AuthorizationPasswordProviderClass {
	AuthorizationPasswordProviderClassOnce.Do(func() {
		AuthorizationPasswordProviderClass = _AuthorizationPasswordProviderClass{objc.GetClass("ASAuthorizationPasswordProvider")}
	})
	return AuthorizationPasswordProviderClass
}

type _AuthorizationPasswordProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPasswordProvider */
// An interface definition for the [AuthorizationPasswordProvider] class.
type IAuthorizationPasswordProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPasswordProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPasswordProvider */
	// methods:
	CreateRequest() IAuthorizationPasswordRequest
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPasswordProvider */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPasswordProviderClass) Alloc() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPasswordProviderClass) New() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPasswordProvider) Init() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPasswordProvider) Autorelease() AuthorizationPasswordProvider {
	rv := objc.Send[AuthorizationPasswordProvider](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPasswordProvider creates a new AuthorizationPasswordProvider instance.
func NewAuthorizationPasswordProvider() AuthorizationPasswordProvider {
	return getAuthorizationPasswordProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPasswordProvider */
// A mechanism for generating requests to perform keychain credential sharing.


// A mechanism for generating requests to perform keychain credential sharing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordProvider
type AuthorizationPasswordProvider struct {
	objectivec.Object
}

// AuthorizationPasswordProviderFrom constructs a [AuthorizationPasswordProvider] from an unsafe.Pointer.
//
// A mechanism for generating requests to perform keychain credential sharing.
func AuthorizationPasswordProviderFrom(ptr unsafe.Pointer) AuthorizationPasswordProvider {
	return AuthorizationPasswordProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPasswordProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPasswordProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPasswordProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPasswordProvider */

// Creates a new password authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordProvider/createRequest()
func (a_ AuthorizationPasswordProvider) CreateRequest() IAuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("createRequest"))
	return rv
}/* debug [instance_methods/method]: CreateRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPasswordProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPasswordProvider */



