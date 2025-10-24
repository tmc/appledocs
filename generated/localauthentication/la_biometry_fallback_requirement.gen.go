// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LABiometryFallbackRequirement */


/* debug [class_header]: Header for LABiometryFallbackRequirement */
// The class instance for the [BiometryFallbackRequirement] class.
var (
	BiometryFallbackRequirementClass     _BiometryFallbackRequirementClass
	BiometryFallbackRequirementClassOnce sync.Once
)

func getBiometryFallbackRequirementClass() _BiometryFallbackRequirementClass {
	BiometryFallbackRequirementClassOnce.Do(func() {
		BiometryFallbackRequirementClass = _BiometryFallbackRequirementClass{objc.GetClass("LABiometryFallbackRequirement")}
	})
	return BiometryFallbackRequirementClass
}

type _BiometryFallbackRequirementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BiometryFallbackRequirement */
// An interface definition for the [BiometryFallbackRequirement] class.
type IBiometryFallbackRequirement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BiometryFallbackRequirement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BiometryFallbackRequirement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BiometryFallbackRequirement */
// Alloc allocates a new instance without initialization.
func (bc _BiometryFallbackRequirementClass) Alloc() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BiometryFallbackRequirementClass) New() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BiometryFallbackRequirement) Init() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BiometryFallbackRequirement) Autorelease() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBiometryFallbackRequirement creates a new BiometryFallbackRequirement instance.
func NewBiometryFallbackRequirement() BiometryFallbackRequirement {
	return getBiometryFallbackRequirementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BiometryFallbackRequirement */
// A set of requirements to fall back on if biometrics aren’t present.


// A set of requirements to fall back on if biometrics aren’t present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement
type BiometryFallbackRequirement struct {
	objectivec.Object
}

// BiometryFallbackRequirementFrom constructs a [BiometryFallbackRequirement] from an unsafe.Pointer.
//
// A set of requirements to fall back on if biometrics aren’t present.
func BiometryFallbackRequirementFrom(ptr unsafe.Pointer) BiometryFallbackRequirement {
	return BiometryFallbackRequirement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BiometryFallbackRequirement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BiometryFallbackRequirement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BiometryFallbackRequirement */

// The default biometric fallback requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/default
func (bc _BiometryFallbackRequirementClass) DefaultRequirement() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("defaultRequirement"))
	return rv
}/* debug [class_properties_class/property]: defaultRequirement */

// The fallback requirement that requires entering the device passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/devicePasscode
func (bc _BiometryFallbackRequirementClass) DevicePasscodeRequirement() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("devicePasscodeRequirement"))
	return rv
}/* debug [class_properties_class/property]: devicePasscodeRequirement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BiometryFallbackRequirement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BiometryFallbackRequirement */

// The default biometric fallback requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/default
func (b_ BiometryFallbackRequirement) DefaultRequirement() ILABiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](b_.ID, objc.Sel("defaultRequirement"))
	return rv
}/* debug [instance_properties/getter]: defaultRequirement */


// The fallback requirement that requires entering the device passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/devicePasscode
func (b_ BiometryFallbackRequirement) DevicePasscodeRequirement() ILABiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](b_.ID, objc.Sel("devicePasscodeRequirement"))
	return rv
}/* debug [instance_properties/getter]: devicePasscodeRequirement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LABiometryFallbackRequirement */



