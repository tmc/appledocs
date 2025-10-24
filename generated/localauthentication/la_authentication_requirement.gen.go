// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LAAuthenticationRequirement */


/* debug [class_header]: Header for LAAuthenticationRequirement */
// The class instance for the [AuthenticationRequirement] class.
var (
	AuthenticationRequirementClass     _AuthenticationRequirementClass
	AuthenticationRequirementClassOnce sync.Once
)

func getAuthenticationRequirementClass() _AuthenticationRequirementClass {
	AuthenticationRequirementClassOnce.Do(func() {
		AuthenticationRequirementClass = _AuthenticationRequirementClass{objc.GetClass("LAAuthenticationRequirement")}
	})
	return AuthenticationRequirementClass
}

type _AuthenticationRequirementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthenticationRequirement */
// An interface definition for the [AuthenticationRequirement] class.
type IAuthenticationRequirement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthenticationRequirement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthenticationRequirement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthenticationRequirement */
// Alloc allocates a new instance without initialization.
func (ac _AuthenticationRequirementClass) Alloc() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthenticationRequirementClass) New() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthenticationRequirement) Init() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthenticationRequirement) Autorelease() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthenticationRequirement creates a new AuthenticationRequirement instance.
func NewAuthenticationRequirement() AuthenticationRequirement {
	return getAuthenticationRequirementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthenticationRequirement */
// A set of requirements that protect a right.


// A set of requirements that protect a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement
type AuthenticationRequirement struct {
	objectivec.Object
}

// AuthenticationRequirementFrom constructs a [AuthenticationRequirement] from an unsafe.Pointer.
//
// A set of requirements that protect a right.
func AuthenticationRequirementFrom(ptr unsafe.Pointer) AuthenticationRequirement {
	return AuthenticationRequirement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthenticationRequirement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthenticationRequirement */

// Creates a requirement that requires biometric authentication or a fallback requirement that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry(fallback:)
func (ac _AuthenticationRequirementClass) BiometryRequirementWithFallback(fallback ILABiometryFallbackRequirement) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("biometryRequirementWithFallback:"), fallback)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BiometryRequirementWithFallback) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthenticationRequirement */

// The requirement that requires biometric authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry
func (ac _AuthenticationRequirementClass) BiometryRequirement() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("biometryRequirement"))
	return rv
}/* debug [class_properties_class/property]: biometryRequirement */

// The requirement that requires user authentication with the current set of biometrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometryCurrentSet
func (ac _AuthenticationRequirementClass) BiometryCurrentSetRequirement() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("biometryCurrentSetRequirement"))
	return rv
}/* debug [class_properties_class/property]: biometryCurrentSetRequirement */

// The requirement that requires user authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/default
func (ac _AuthenticationRequirementClass) DefaultRequirement() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("defaultRequirement"))
	return rv
}/* debug [class_properties_class/property]: defaultRequirement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthenticationRequirement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthenticationRequirement */

// The requirement that requires biometric authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry
func (a_ AuthenticationRequirement) BiometryRequirement() ILAAuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](a_.ID, objc.Sel("biometryRequirement"))
	return rv
}/* debug [instance_properties/getter]: biometryRequirement */


// The requirement that requires user authentication with the current set of biometrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometryCurrentSet
func (a_ AuthenticationRequirement) BiometryCurrentSetRequirement() ILAAuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](a_.ID, objc.Sel("biometryCurrentSetRequirement"))
	return rv
}/* debug [instance_properties/getter]: biometryCurrentSetRequirement */


// The requirement that requires user authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/default
func (a_ AuthenticationRequirement) DefaultRequirement() ILAAuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](a_.ID, objc.Sel("defaultRequirement"))
	return rv
}/* debug [instance_properties/getter]: defaultRequirement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAAuthenticationRequirement */



