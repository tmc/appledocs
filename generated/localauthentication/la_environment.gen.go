// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LAEnvironment */


/* debug [class_header]: Header for LAEnvironment */
// The class instance for the [Environment] class.
var (
	EnvironmentClass     _EnvironmentClass
	EnvironmentClassOnce sync.Once
)

func getEnvironmentClass() _EnvironmentClass {
	EnvironmentClassOnce.Do(func() {
		EnvironmentClass = _EnvironmentClass{objc.GetClass("LAEnvironment")}
	})
	return EnvironmentClass
}

type _EnvironmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Environment */
// An interface definition for the [Environment] class.
type IEnvironment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Environment */
	// properties:
	State() ILAEnvironmentState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Environment */
	// methods:
	AddObserver(observer unsafe.Pointer)
	RemoveObserver(observer unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Environment */
// Alloc allocates a new instance without initialization.
func (ec _EnvironmentClass) Alloc() Environment {
	rv := objc.Send[Environment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EnvironmentClass) New() Environment {
	rv := objc.Send[Environment](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Environment) Init() Environment {
	rv := objc.Send[Environment](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Environment) Autorelease() Environment {
	rv := objc.Send[Environment](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironment creates a new Environment instance.
func NewEnvironment() Environment {
	return getEnvironmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Environment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment
type Environment struct {
	objectivec.Object
}

// EnvironmentFrom constructs a [Environment] from an unsafe.Pointer.
func EnvironmentFrom(ptr unsafe.Pointer) Environment {
	return Environment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Environment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Environment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Environment */

// Environment of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/currentUser
func (ec _EnvironmentClass) CurrentUser() Environment {
	rv := objc.Send[Environment](objc.ID(ec.class), objc.Sel("currentUser"))
	return rv
}/* debug [class_properties_class/property]: currentUser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Environment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/addObserver(_:)
func (e_ Environment) AddObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addObserver:"), observer)
}/* debug [instance_methods/method]: AddObserver */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/removeObserver(_:)
func (e_ Environment) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeObserver:"), observer)
}/* debug [instance_methods/method]: RemoveObserver */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Environment */

// Environment of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/currentUser
func (e_ Environment) CurrentUser() ILAEnvironment {
	rv := objc.Send[Environment](e_.ID, objc.Sel("currentUser"))
	return rv
}/* debug [instance_properties/getter]: currentUser */


// The environment state information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/state-swift.property
func (e_ Environment) State() ILAEnvironmentState {
	rv := objc.Send[EnvironmentState](e_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAEnvironment */



