// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustUsesMutualAuthentication */


/* debug [class_header]: Header for trustUsesMutualAuthentication */
// The class instance for the [trustUsesMutualAuthentication] class.
var (
	TrustUsesMutualAuthenticationClass     _trustUsesMutualAuthenticationClass
	TrustUsesMutualAuthenticationClassOnce sync.Once
)

func gettrustUsesMutualAuthenticationClass() _trustUsesMutualAuthenticationClass {
	TrustUsesMutualAuthenticationClassOnce.Do(func() {
		TrustUsesMutualAuthenticationClass = _trustUsesMutualAuthenticationClass{objc.GetClass("trustUsesMutualAuthentication")}
	})
	return TrustUsesMutualAuthenticationClass
}

type _trustUsesMutualAuthenticationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustUsesMutualAuthentication */
// An interface definition for the [trustUsesMutualAuthentication] class.
type ItrustUsesMutualAuthentication interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustUsesMutualAuthentication */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustUsesMutualAuthentication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustUsesMutualAuthentication */
// Alloc allocates a new instance without initialization.
func (tc _trustUsesMutualAuthenticationClass) Alloc() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustUsesMutualAuthenticationClass) New() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesMutualAuthentication) Init() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesMutualAuthentication) Autorelease() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesMutualAuthentication creates a new trustUsesMutualAuthentication instance.
func NewtrustUsesMutualAuthentication() trustUsesMutualAuthentication {
	return gettrustUsesMutualAuthenticationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustUsesMutualAuthentication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-c.ivar
type trustUsesMutualAuthentication struct {
	objectivec.Object
}

// trustUsesMutualAuthenticationFrom constructs a [trustUsesMutualAuthentication] from an unsafe.Pointer.
func trustUsesMutualAuthenticationFrom(ptr unsafe.Pointer) trustUsesMutualAuthentication {
	return trustUsesMutualAuthentication{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustUsesMutualAuthentication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustUsesMutualAuthentication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustUsesMutualAuthentication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustUsesMutualAuthentication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustUsesMutualAuthentication */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustUsesMutualAuthentication */



