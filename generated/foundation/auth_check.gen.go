// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class authCheck */


/* debug [class_header]: Header for authCheck */
// The class instance for the [authCheck] class.
var (
	AuthCheckClass     _authCheckClass
	AuthCheckClassOnce sync.Once
)

func getauthCheckClass() _authCheckClass {
	AuthCheckClassOnce.Do(func() {
		AuthCheckClass = _authCheckClass{objc.GetClass("authCheck")}
	})
	return AuthCheckClass
}

type _authCheckClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for authCheck */
// An interface definition for the [authCheck] class.
type IauthCheck interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for authCheck */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for authCheck */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for authCheck */
// Alloc allocates a new instance without initialization.
func (ac _authCheckClass) Alloc() authCheck {
	rv := objc.Send[authCheck](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _authCheckClass) New() authCheck {
	rv := objc.Send[authCheck](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authCheck) Init() authCheck {
	rv := objc.Send[authCheck](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authCheck) Autorelease() authCheck {
	rv := objc.Send[authCheck](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthCheck creates a new authCheck instance.
func NewauthCheck() authCheck {
	return getauthCheckClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for authCheck */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/authCheck
type authCheck struct {
	objectivec.Object
}

// authCheckFrom constructs a [authCheck] from an unsafe.Pointer.
func authCheckFrom(ptr unsafe.Pointer) authCheck {
	return authCheck{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for authCheck *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for authCheck */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for authCheck */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for authCheck */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for authCheck */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class authCheck */



