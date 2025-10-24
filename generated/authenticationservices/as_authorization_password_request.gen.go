// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPasswordRequest */


/* debug [class_header]: Header for ASAuthorizationPasswordRequest */
// The class instance for the [AuthorizationPasswordRequest] class.
var (
	AuthorizationPasswordRequestClass     _AuthorizationPasswordRequestClass
	AuthorizationPasswordRequestClassOnce sync.Once
)

func getAuthorizationPasswordRequestClass() _AuthorizationPasswordRequestClass {
	AuthorizationPasswordRequestClassOnce.Do(func() {
		AuthorizationPasswordRequestClass = _AuthorizationPasswordRequestClass{objc.GetClass("ASAuthorizationPasswordRequest")}
	})
	return AuthorizationPasswordRequestClass
}

type _AuthorizationPasswordRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPasswordRequest */
// An interface definition for the [AuthorizationPasswordRequest] class.
type IAuthorizationPasswordRequest interface {
	IAuthorizationRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationPasswordRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPasswordRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPasswordRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPasswordRequestClass) Alloc() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPasswordRequestClass) New() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPasswordRequest) Init() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPasswordRequest) Autorelease() AuthorizationPasswordRequest {
	rv := objc.Send[AuthorizationPasswordRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPasswordRequest creates a new AuthorizationPasswordRequest instance.
func NewAuthorizationPasswordRequest() AuthorizationPasswordRequest {
	return getAuthorizationPasswordRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPasswordRequest */
// An authorization request that uses credentials stored in the keychain.


// An authorization request that uses credentials stored in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPasswordRequest
type AuthorizationPasswordRequest struct {
	AuthorizationRequest
}

// AuthorizationPasswordRequestFrom constructs a [AuthorizationPasswordRequest] from an unsafe.Pointer.
//
// An authorization request that uses credentials stored in the keychain.
func AuthorizationPasswordRequestFrom(ptr unsafe.Pointer) AuthorizationPasswordRequest {
	return AuthorizationPasswordRequest{
		AuthorizationRequest: AuthorizationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPasswordRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPasswordRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPasswordRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPasswordRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPasswordRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPasswordRequest */



