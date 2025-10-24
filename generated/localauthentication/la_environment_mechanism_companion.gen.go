// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class LAEnvironmentMechanismCompanion */


/* debug [class_header]: Header for LAEnvironmentMechanismCompanion */
// The class instance for the [EnvironmentMechanismCompanion] class.
var (
	EnvironmentMechanismCompanionClass     _EnvironmentMechanismCompanionClass
	EnvironmentMechanismCompanionClassOnce sync.Once
)

func getEnvironmentMechanismCompanionClass() _EnvironmentMechanismCompanionClass {
	EnvironmentMechanismCompanionClassOnce.Do(func() {
		EnvironmentMechanismCompanionClass = _EnvironmentMechanismCompanionClass{objc.GetClass("LAEnvironmentMechanismCompanion")}
	})
	return EnvironmentMechanismCompanionClass
}

type _EnvironmentMechanismCompanionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EnvironmentMechanismCompanion */
// An interface definition for the [EnvironmentMechanismCompanion] class.
type IEnvironmentMechanismCompanion interface {
	IEnvironmentMechanism
	
/* debug [class_interface_properties]: Properties for EnvironmentMechanismCompanion */
	// properties:
	StateHash() objc.IObject /* cross-framework: NSData */
	Type() CompanionType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EnvironmentMechanismCompanion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EnvironmentMechanismCompanion */
// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismCompanionClass) Alloc() EnvironmentMechanismCompanion {
	rv := objc.Send[EnvironmentMechanismCompanion](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EnvironmentMechanismCompanionClass) New() EnvironmentMechanismCompanion {
	rv := objc.Send[EnvironmentMechanismCompanion](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnvironmentMechanismCompanion) Init() EnvironmentMechanismCompanion {
	rv := objc.Send[EnvironmentMechanismCompanion](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnvironmentMechanismCompanion) Autorelease() EnvironmentMechanismCompanion {
	rv := objc.Send[EnvironmentMechanismCompanion](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironmentMechanismCompanion creates a new EnvironmentMechanismCompanion instance.
func NewEnvironmentMechanismCompanion() EnvironmentMechanismCompanion {
	return getEnvironmentMechanismCompanionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EnvironmentMechanismCompanion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion
type EnvironmentMechanismCompanion struct {
	EnvironmentMechanism
}

// EnvironmentMechanismCompanionFrom constructs a [EnvironmentMechanismCompanion] from an unsafe.Pointer.
func EnvironmentMechanismCompanionFrom(ptr unsafe.Pointer) EnvironmentMechanismCompanion {
	return EnvironmentMechanismCompanion{
		EnvironmentMechanism: EnvironmentMechanismFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EnvironmentMechanismCompanion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EnvironmentMechanismCompanion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EnvironmentMechanismCompanion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EnvironmentMechanismCompanion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EnvironmentMechanismCompanion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/stateHash
func (e_ EnvironmentMechanismCompanion) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("stateHash"))
	return rv
}/* debug [instance_properties/getter]: stateHash */


// Type of the companion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/type
func (e_ EnvironmentMechanismCompanion) Type() CompanionType {
	rv := objc.Send[CompanionType](e_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAEnvironmentMechanismCompanion */



