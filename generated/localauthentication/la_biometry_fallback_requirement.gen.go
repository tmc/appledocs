// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BiometryFallbackRequirement] class.
type IBiometryFallbackRequirement interface {
	objectivec.IObject
}

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

// Alloc allocates a new instance without initialization.
func (bc _BiometryFallbackRequirementClass) Alloc() BiometryFallbackRequirement {
	rv := objc.Send[BiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The default biometric fallback requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/default
func (bc _BiometryFallbackRequirementClass) DefaultRequirement() BiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("defaultRequirement"))
	return rv
}

// The fallback requirement that requires entering the device passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/devicePasscode
func (bc _BiometryFallbackRequirementClass) DevicePasscodeRequirement() BiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](objc.ID(bc.class), objc.Sel("devicePasscodeRequirement"))
	return rv
}

// The default biometric fallback requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/default
func (b_ BiometryFallbackRequirement) DefaultRequirement() LABiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](b_.ID, objc.Sel("defaultRequirement"))
	return rv
}


// The fallback requirement that requires entering the device passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryFallbackRequirement/devicePasscode
func (b_ BiometryFallbackRequirement) DevicePasscodeRequirement() LABiometryFallbackRequirement {
	rv := objc.Send[LABiometryFallbackRequirement](b_.ID, objc.Sel("devicePasscodeRequirement"))
	return rv
}



