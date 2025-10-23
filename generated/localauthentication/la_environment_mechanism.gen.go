// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EnvironmentMechanism] class.
var (
	EnvironmentMechanismClass     _EnvironmentMechanismClass
	EnvironmentMechanismClassOnce sync.Once
)

func getEnvironmentMechanismClass() _EnvironmentMechanismClass {
	EnvironmentMechanismClassOnce.Do(func() {
		EnvironmentMechanismClass = _EnvironmentMechanismClass{objc.GetClass("LAEnvironmentMechanism")}
	})
	return EnvironmentMechanismClass
}

type _EnvironmentMechanismClass struct {
	class objc.Class
}

// An interface definition for the [EnvironmentMechanism] class.
type IEnvironmentMechanism interface {
	objectivec.IObject
	// properties:
	IconSystemName() string
	IsUsable() bool
	LocalizedName() string
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism
type EnvironmentMechanism struct {
	objectivec.Object
}

// EnvironmentMechanismFrom constructs a [EnvironmentMechanism] from an unsafe.Pointer.
func EnvironmentMechanismFrom(ptr unsafe.Pointer) EnvironmentMechanism {
	return EnvironmentMechanism{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentMechanismClass) Alloc() EnvironmentMechanism {
	rv := objc.Send[EnvironmentMechanism](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EnvironmentMechanismClass) New() EnvironmentMechanism {
	rv := objc.Send[EnvironmentMechanism](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnvironmentMechanism) Init() EnvironmentMechanism {
	rv := objc.Send[EnvironmentMechanism](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnvironmentMechanism) Autorelease() EnvironmentMechanism {
	rv := objc.Send[EnvironmentMechanism](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironmentMechanism creates a new EnvironmentMechanism instance.
func NewEnvironmentMechanism() EnvironmentMechanism {
	return getEnvironmentMechanismClass().New()
}



// Name of the SF Symbol representing this authentication mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/iconSystemName
func (e_ EnvironmentMechanism) IconSystemName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("iconSystemName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/isUsable
func (e_ EnvironmentMechanism) IsUsable() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isUsable"))
	return rv
}


// The localized name of the authentication mechanism, e.g. “Touch ID”, “Face ID” etc.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/Mechanism/localizedName
func (e_ EnvironmentMechanism) LocalizedName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("localizedName"))
	return rv
}



