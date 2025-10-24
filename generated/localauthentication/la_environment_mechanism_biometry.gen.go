// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class LAEnvironmentMechanismBiometry */


/* debug [class_header]: Header for LAEnvironmentMechanismBiometry */
// The class instance for the [EnvironmentMechanismBiometry] class.
var (
	EnvironmentMechanismBiometryClass     _EnvironmentMechanismBiometryClass
	EnvironmentMechanismBiometryClassOnce sync.Once
)

func getEnvironmentMechanismBiometryClass() _EnvironmentMechanismBiometryClass {
	EnvironmentMechanismBiometryClassOnce.Do(func() {
		EnvironmentMechanismBiometryClass = _EnvironmentMechanismBiometryClass{objc.GetClass("LAEnvironmentMechanismBiometry")}
	})
	return EnvironmentMechanismBiometryClass
}

type _EnvironmentMechanismBiometryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EnvironmentMechanismBiometry */
// An interface definition for the [EnvironmentMechanismBiometry] class.
type IEnvironmentMechanismBiometry interface {
	IEnvironmentMechanism
	
/* debug [class_interface_properties]: Properties for EnvironmentMechanismBiometry */
	// properties:
	BiometryType() BiometryType
	BuiltInSensorInaccessible() bool
	IsEnrolled() bool
	IsLockedOut() bool
	StateHash() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EnvironmentMechanismBiometry */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EnvironmentMechanismBiometry */
// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismBiometryClass) Alloc() EnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EnvironmentMechanismBiometryClass) New() EnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnvironmentMechanismBiometry) Init() EnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnvironmentMechanismBiometry) Autorelease() EnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironmentMechanismBiometry creates a new EnvironmentMechanismBiometry instance.
func NewEnvironmentMechanismBiometry() EnvironmentMechanismBiometry {
	return getEnvironmentMechanismBiometryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EnvironmentMechanismBiometry */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry
type EnvironmentMechanismBiometry struct {
	EnvironmentMechanism
}

// EnvironmentMechanismBiometryFrom constructs a [EnvironmentMechanismBiometry] from an unsafe.Pointer.
func EnvironmentMechanismBiometryFrom(ptr unsafe.Pointer) EnvironmentMechanismBiometry {
	return EnvironmentMechanismBiometry{
		EnvironmentMechanism: EnvironmentMechanismFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EnvironmentMechanismBiometry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EnvironmentMechanismBiometry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EnvironmentMechanismBiometry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EnvironmentMechanismBiometry */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EnvironmentMechanismBiometry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/biometryType
func (e_ EnvironmentMechanismBiometry) BiometryType() BiometryType {
	rv := objc.Send[BiometryType](e_.ID, objc.Sel("biometryType"))
	return rv
}/* debug [instance_properties/getter]: biometryType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/builtInSensorInaccessible
func (e_ EnvironmentMechanismBiometry) BuiltInSensorInaccessible() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("builtInSensorInaccessible"))
	return rv
}/* debug [instance_properties/getter]: builtInSensorInaccessible */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isEnrolled
func (e_ EnvironmentMechanismBiometry) IsEnrolled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnrolled"))
	return rv
}/* debug [instance_properties/getter]: isEnrolled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isLockedOut
func (e_ EnvironmentMechanismBiometry) IsLockedOut() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isLockedOut"))
	return rv
}/* debug [instance_properties/getter]: isLockedOut */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/stateHash
func (e_ EnvironmentMechanismBiometry) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("stateHash"))
	return rv
}/* debug [instance_properties/getter]: stateHash */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAEnvironmentMechanismBiometry */



