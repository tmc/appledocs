// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustKerberosPrincipal */


/* debug [class_header]: Header for trustKerberosPrincipal */
// The class instance for the [trustKerberosPrincipal] class.
var (
	TrustKerberosPrincipalClass     _trustKerberosPrincipalClass
	TrustKerberosPrincipalClassOnce sync.Once
)

func gettrustKerberosPrincipalClass() _trustKerberosPrincipalClass {
	TrustKerberosPrincipalClassOnce.Do(func() {
		TrustKerberosPrincipalClass = _trustKerberosPrincipalClass{objc.GetClass("trustKerberosPrincipal")}
	})
	return TrustKerberosPrincipalClass
}

type _trustKerberosPrincipalClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustKerberosPrincipal */
// An interface definition for the [trustKerberosPrincipal] class.
type ItrustKerberosPrincipal interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustKerberosPrincipal */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustKerberosPrincipal */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustKerberosPrincipal */
// Alloc allocates a new instance without initialization.
func (tc _trustKerberosPrincipalClass) Alloc() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustKerberosPrincipalClass) New() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustKerberosPrincipal) Init() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustKerberosPrincipal) Autorelease() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustKerberosPrincipal creates a new trustKerberosPrincipal instance.
func NewtrustKerberosPrincipal() trustKerberosPrincipal {
	return gettrustKerberosPrincipalClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustKerberosPrincipal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-c.ivar
type trustKerberosPrincipal struct {
	objectivec.Object
}

// trustKerberosPrincipalFrom constructs a [trustKerberosPrincipal] from an unsafe.Pointer.
func trustKerberosPrincipalFrom(ptr unsafe.Pointer) trustKerberosPrincipal {
	return trustKerberosPrincipal{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustKerberosPrincipal *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustKerberosPrincipal */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustKerberosPrincipal */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustKerberosPrincipal */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustKerberosPrincipal */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustKerberosPrincipal */



