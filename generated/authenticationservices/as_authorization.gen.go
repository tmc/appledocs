// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorization */


/* debug [class_header]: Header for ASAuthorization */
// The class instance for the [Authorization] class.
var (
	AuthorizationClass     _AuthorizationClass
	AuthorizationClassOnce sync.Once
)

func getAuthorizationClass() _AuthorizationClass {
	AuthorizationClassOnce.Do(func() {
		AuthorizationClass = _AuthorizationClass{objc.GetClass("ASAuthorization")}
	})
	return AuthorizationClass
}

type _AuthorizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Authorization */
// An interface definition for the [Authorization] class.
type IAuthorization interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Authorization */
	// properties:
	Credential() unsafe.Pointer
	Provider() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Authorization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Authorization */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationClass) Alloc() Authorization {
	rv := objc.Send[Authorization](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationClass) New() Authorization {
	rv := objc.Send[Authorization](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Authorization) Init() Authorization {
	rv := objc.Send[Authorization](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Authorization) Autorelease() Authorization {
	rv := objc.Send[Authorization](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorization creates a new Authorization instance.
func NewAuthorization() Authorization {
	return getAuthorizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Authorization */
// The encapsulation of a successful authorization by a controller.


// The encapsulation of a successful authorization by a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorization
type Authorization struct {
	objectivec.Object
}

// AuthorizationFrom constructs a [Authorization] from an unsafe.Pointer.
//
// The encapsulation of a successful authorization by a controller.
func AuthorizationFrom(ptr unsafe.Pointer) Authorization {
	return Authorization{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Authorization *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Authorization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Authorization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Authorization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Authorization */

// Information provided about a user after successful authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorization/credential
func (a_ Authorization) Credential() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("credential"))
	return rv
}/* debug [instance_properties/getter]: credential */


// The provider that created the request that resulted in the successful authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorization/provider
func (a_ Authorization) Provider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("provider"))
	return rv
}/* debug [instance_properties/getter]: provider */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorization */



