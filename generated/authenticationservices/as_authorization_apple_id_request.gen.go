// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAuthorizationAppleIDRequest */


/* debug [class_header]: Header for ASAuthorizationAppleIDRequest */
// The class instance for the [AuthorizationAppleIDRequest] class.
var (
	AuthorizationAppleIDRequestClass     _AuthorizationAppleIDRequestClass
	AuthorizationAppleIDRequestClassOnce sync.Once
)

func getAuthorizationAppleIDRequestClass() _AuthorizationAppleIDRequestClass {
	AuthorizationAppleIDRequestClassOnce.Do(func() {
		AuthorizationAppleIDRequestClass = _AuthorizationAppleIDRequestClass{objc.GetClass("ASAuthorizationAppleIDRequest")}
	})
	return AuthorizationAppleIDRequestClass
}

type _AuthorizationAppleIDRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationAppleIDRequest */
// An interface definition for the [AuthorizationAppleIDRequest] class.
type IAuthorizationAppleIDRequest interface {
	IAuthorizationOpenIDRequest
	
/* debug [class_interface_properties]: Properties for AuthorizationAppleIDRequest */
	// properties:
	User() objc.IObject /* cross-framework: NSString */
	SetUser(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationAppleIDRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationAppleIDRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationAppleIDRequestClass) Alloc() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationAppleIDRequestClass) New() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationAppleIDRequest) Init() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationAppleIDRequest) Autorelease() AuthorizationAppleIDRequest {
	rv := objc.Send[AuthorizationAppleIDRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationAppleIDRequest creates a new AuthorizationAppleIDRequest instance.
func NewAuthorizationAppleIDRequest() AuthorizationAppleIDRequest {
	return getAuthorizationAppleIDRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationAppleIDRequest */
// An OpenID authorization request that relies on the user’s Apple ID.


// An OpenID authorization request that relies on the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDRequest
type AuthorizationAppleIDRequest struct {
	AuthorizationOpenIDRequest
}

// AuthorizationAppleIDRequestFrom constructs a [AuthorizationAppleIDRequest] from an unsafe.Pointer.
//
// An OpenID authorization request that relies on the user’s Apple ID.
func AuthorizationAppleIDRequestFrom(ptr unsafe.Pointer) AuthorizationAppleIDRequest {
	return AuthorizationAppleIDRequest{
		AuthorizationOpenIDRequest: AuthorizationOpenIDRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationAppleIDRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationAppleIDRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationAppleIDRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationAppleIDRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationAppleIDRequest */

// An identifier associated with the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDRequest/user
func (a_ AuthorizationAppleIDRequest) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */


// An identifier associated with the user’s Apple ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDRequest/user
func (a_ AuthorizationAppleIDRequest) SetUser(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUser:"), value)
}/* debug [instance_properties/setter]: user */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationAppleIDRequest */



