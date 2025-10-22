// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Environment] class.
var (
	EnvironmentClass     _EnvironmentClass
	EnvironmentClassOnce sync.Once
)

func getEnvironmentClass() _EnvironmentClass {
	EnvironmentClassOnce.Do(func() {
		EnvironmentClass = _EnvironmentClass{objc.GetClass("LAEnvironment")}
	})
	return EnvironmentClass
}

type _EnvironmentClass struct {
	class objc.Class
}

// An interface definition for the [Environment] class.
type IEnvironment interface {
	objectivec.IObject
	AddObserver(observer objectivec.IObject)
	RemoveObserver(observer objectivec.IObject)
	State() LAEnvironmentState
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment
type Environment struct {
	objectivec.Object
}

// EnvironmentFrom constructs a [Environment] from an unsafe.Pointer.
func EnvironmentFrom(ptr unsafe.Pointer) Environment {
	return Environment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentClass) Alloc() Environment {
	rv := objc.Send[Environment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EnvironmentClass) New() Environment {
	rv := objc.Send[Environment](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Environment) Init() Environment {
	rv := objc.Send[Environment](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Environment) Autorelease() Environment {
	rv := objc.Send[Environment](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironment creates a new Environment instance.
func NewEnvironment() Environment {
	return getEnvironmentClass().New()
}


// Environment of the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/currentUser
func (ec _EnvironmentClass) CurrentUser() Environment {
	rv := objc.Send[LAEnvironment](objc.ID(ec.class), objc.Sel("currentUser"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/addObserver(_:)
func (e_ Environment) AddObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addObserver:"), observer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/removeObserver(_:)
func (e_ Environment) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeObserver:"), observer)
}

// Environment of the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/currentUser
func (e_ Environment) CurrentUser() LAEnvironment {
	rv := objc.Send[LAEnvironment](e_.ID, objc.Sel("currentUser"))
	return rv
}

// The environment state information.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/state-swift.property
func (e_ Environment) State() LAEnvironmentState {
	rv := objc.Send[LAEnvironmentState](e_.ID, objc.Sel("state"))
	return rv
}



