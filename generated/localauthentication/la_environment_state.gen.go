// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EnvironmentState] class.
var (
	EnvironmentStateClass     _EnvironmentStateClass
	EnvironmentStateClassOnce sync.Once
)

func getEnvironmentStateClass() _EnvironmentStateClass {
	EnvironmentStateClassOnce.Do(func() {
		EnvironmentStateClass = _EnvironmentStateClass{objc.GetClass("LAEnvironmentState")}
	})
	return EnvironmentStateClass
}

type _EnvironmentStateClass struct {
	class objc.Class
}

// An interface definition for the [EnvironmentState] class.
type IEnvironmentState interface {
	objectivec.IObject
	// properties:
	AllMechanisms() []EnvironmentMechanism /* primitive/slice/pointer. */
	Biometry() ILAEnvironmentMechanismBiometry
	Companions() []EnvironmentMechanismCompanion /* primitive/slice/pointer. */
	UserPassword() ILAEnvironmentMechanismUserPassword
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class
type EnvironmentState struct {
	objectivec.Object
}

// EnvironmentStateFrom constructs a [EnvironmentState] from an unsafe.Pointer.
func EnvironmentStateFrom(ptr unsafe.Pointer) EnvironmentState {
	return EnvironmentState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EnvironmentStateClass) Alloc() EnvironmentState {
	rv := objc.Send[EnvironmentState](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EnvironmentStateClass) New() EnvironmentState {
	rv := objc.Send[EnvironmentState](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnvironmentState) Init() EnvironmentState {
	rv := objc.Send[EnvironmentState](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnvironmentState) Autorelease() EnvironmentState {
	rv := objc.Send[EnvironmentState](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnvironmentState creates a new EnvironmentState instance.
func NewEnvironmentState() EnvironmentState {
	return getEnvironmentStateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/allMechanisms
func (e_ EnvironmentState) AllMechanisms() []EnvironmentMechanism /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EnvironmentMechanism](e_.ID, objc.Sel("allMechanisms"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/biometry
func (e_ EnvironmentState) Biometry() ILAEnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](e_.ID, objc.Sel("biometry"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/companions
func (e_ EnvironmentState) Companions() []EnvironmentMechanismCompanion /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EnvironmentMechanismCompanion](e_.ID, objc.Sel("companions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/userPassword
func (e_ EnvironmentState) UserPassword() ILAEnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](e_.ID, objc.Sel("userPassword"))
	return rv
}



