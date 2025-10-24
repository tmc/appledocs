// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class LAEnvironmentMechanismUserPassword */


/* debug [class_header]: Header for LAEnvironmentMechanismUserPassword */
// The class instance for the [EnvironmentMechanismUserPassword] class.
var (
	EnvironmentMechanismUserPasswordClass     _EnvironmentMechanismUserPasswordClass
	EnvironmentMechanismUserPasswordClassOnce sync.Once
)

func getEnvironmentMechanismUserPasswordClass() _EnvironmentMechanismUserPasswordClass {
	EnvironmentMechanismUserPasswordClassOnce.Do(func() {
		EnvironmentMechanismUserPasswordClass = _EnvironmentMechanismUserPasswordClass{objc.GetClass("LAEnvironmentMechanismUserPassword")}
	})
	return EnvironmentMechanismUserPasswordClass
}

type _EnvironmentMechanismUserPasswordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EnvironmentMechanismUserPassword */
// An interface definition for the [EnvironmentMechanismUserPassword] class.
type IEnvironmentMechanismUserPassword interface {
	IEnvironmentMechanism
	
/* debug [class_interface_properties]: Properties for EnvironmentMechanismUserPassword */
	// properties:
	IsSet() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EnvironmentMechanismUserPassword */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EnvironmentMechanismUserPassword */
// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismUserPasswordClass) Alloc() EnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EnvironmentMechanismUserPasswordClass) New() EnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnvironmentMechanismUserPassword) Init() EnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnvironmentMechanismUserPassword) Autorelease() EnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironmentMechanismUserPassword creates a new EnvironmentMechanismUserPassword instance.
func NewEnvironmentMechanismUserPassword() EnvironmentMechanismUserPassword {
	return getEnvironmentMechanismUserPasswordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EnvironmentMechanismUserPassword */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword
type EnvironmentMechanismUserPassword struct {
	EnvironmentMechanism
}

// EnvironmentMechanismUserPasswordFrom constructs a [EnvironmentMechanismUserPassword] from an unsafe.Pointer.
func EnvironmentMechanismUserPasswordFrom(ptr unsafe.Pointer) EnvironmentMechanismUserPassword {
	return EnvironmentMechanismUserPassword{
		EnvironmentMechanism: EnvironmentMechanismFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EnvironmentMechanismUserPassword *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EnvironmentMechanismUserPassword */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EnvironmentMechanismUserPassword */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EnvironmentMechanismUserPassword */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EnvironmentMechanismUserPassword */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword/isSet
func (e_ EnvironmentMechanismUserPassword) IsSet() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isSet"))
	return rv
}/* debug [instance_properties/getter]: isSet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAEnvironmentMechanismUserPassword */



