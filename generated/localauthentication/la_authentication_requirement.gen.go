// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AuthenticationRequirement] class.
type IAuthenticationRequirement interface {
	objectivec.IObject
}

// A set of requirements that protect a right.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AuthenticationRequirementClass) Alloc() AuthenticationRequirement {
	rv := objc.Send[AuthenticationRequirement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a requirement that requires biometric authentication or a fallback requirement that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry(fallback:)
func (ac _AuthenticationRequirementClass) BiometryRequirementWithFallback(fallback ILABiometryFallbackRequirement) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("biometryRequirementWithFallback:"), fallback)
	return rv
}

// The requirement that requires biometric authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry
func (ac _AuthenticationRequirementClass) BiometryRequirement() AuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](objc.ID(ac.class), objc.Sel("biometryRequirement"))
	return rv
}
// The requirement that requires user authentication with the current set of biometrics.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometryCurrentSet
func (ac _AuthenticationRequirementClass) BiometryCurrentSetRequirement() AuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](objc.ID(ac.class), objc.Sel("biometryCurrentSetRequirement"))
	return rv
}
// The requirement that requires user authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/default
func (ac _AuthenticationRequirementClass) DefaultRequirement() AuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](objc.ID(ac.class), objc.Sel("defaultRequirement"))
	return rv
}
// The requirement that requires biometric authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometry
func (a_ AuthenticationRequirement) BiometryRequirement() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](a_.ID, objc.Sel("biometryRequirement"))
	return rv
}

// The requirement that requires user authentication with the current set of biometrics.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/biometryCurrentSet
func (a_ AuthenticationRequirement) BiometryCurrentSetRequirement() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](a_.ID, objc.Sel("biometryCurrentSetRequirement"))
	return rv
}

// The requirement that requires user authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAuthenticationRequirement/default
func (a_ AuthenticationRequirement) DefaultRequirement() LAAuthenticationRequirement {
	rv := objc.Send[LAAuthenticationRequirement](a_.ID, objc.Sel("defaultRequirement"))
	return rv
}



