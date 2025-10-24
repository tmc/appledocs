// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [EnvironmentMechanismUserPassword] class.
type IEnvironmentMechanismUserPassword interface {
	IEnvironmentMechanism
	// properties:
	IsSet() bool
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismUserPasswordClass) Alloc() EnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismUserPassword/isSet
func (e_ EnvironmentMechanismUserPassword) IsSet() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isSet"))
	return rv
}



