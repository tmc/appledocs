// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [EnvironmentMechanismCompanion] class.
type IEnvironmentMechanismCompanion interface {
	IEnvironmentMechanism
	// properties:
	StateHash() objc.IObject /* cross-framework: NSData */
	Type() CompanionType
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismCompanionClass) Alloc() EnvironmentMechanismCompanion {
	rv := objc.Send[EnvironmentMechanismCompanion](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/stateHash
func (e_ EnvironmentMechanismCompanion) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("stateHash"))
	return rv
}


// Type of the companion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/MechanismCompanion/type
func (e_ EnvironmentMechanismCompanion) Type() CompanionType {
	rv := objc.Send[CompanionType](e_.ID, objc.Sel("type"))
	return rv
}



