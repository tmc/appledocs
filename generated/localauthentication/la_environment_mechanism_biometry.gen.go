// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [EnvironmentMechanismBiometry] class.
type IEnvironmentMechanismBiometry interface {
	IEnvironmentMechanism
}

//
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

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismBiometryClass) Alloc() EnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/biometryType
func (e_ EnvironmentMechanismBiometry) BiometryType() BiometryType {
	rv := objc.Send[BiometryType](e_.ID, objc.Sel("biometryType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/builtInSensorInaccessible
func (e_ EnvironmentMechanismBiometry) BuiltInSensorInaccessible() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("builtInSensorInaccessible"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isEnrolled
func (e_ EnvironmentMechanismBiometry) IsEnrolled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnrolled"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/isLockedOut
func (e_ EnvironmentMechanismBiometry) IsLockedOut() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isLockedOut"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismBiometry/stateHash
func (e_ EnvironmentMechanismBiometry) StateHash() foundation.NSData {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("stateHash"))
	return rv
}



