// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class authGen */


/* debug [class_header]: Header for authGen */
// The class instance for the [authGen] class.
var (
	AuthGenClass     _authGenClass
	AuthGenClassOnce sync.Once
)

func getauthGenClass() _authGenClass {
	AuthGenClassOnce.Do(func() {
		AuthGenClass = _authGenClass{objc.GetClass("authGen")}
	})
	return AuthGenClass
}

type _authGenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for authGen */
// An interface definition for the [authGen] class.
type IauthGen interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for authGen */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for authGen */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for authGen */
// Alloc allocates a new instance without initialization.
func (ac _authGenClass) Alloc() authGen {
	rv := objc.Send[authGen](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _authGenClass) New() authGen {
	rv := objc.Send[authGen](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authGen) Init() authGen {
	rv := objc.Send[authGen](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authGen) Autorelease() authGen {
	rv := objc.Send[authGen](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthGen creates a new authGen instance.
func NewauthGen() authGen {
	return getauthGenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for authGen */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/authGen
type authGen struct {
	objectivec.Object
}

// authGenFrom constructs a [authGen] from an unsafe.Pointer.
func authGenFrom(ptr unsafe.Pointer) authGen {
	return authGen{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for authGen *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for authGen */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for authGen */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for authGen */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for authGen */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class authGen */



