// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LAEnvironmentState */


/* debug [class_header]: Header for LAEnvironmentState */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EnvironmentState */
// An interface definition for the [EnvironmentState] class.
type IEnvironmentState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EnvironmentState */
	// properties:
	AllMechanisms() []EnvironmentMechanism
	Biometry() ILAEnvironmentMechanismBiometry
	Companions() []EnvironmentMechanismCompanion
	UserPassword() ILAEnvironmentMechanismUserPassword
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EnvironmentState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EnvironmentState */
// Alloc allocates a new instance without initialization.
func (ec _EnvironmentStateClass) Alloc() EnvironmentState {
	rv := objc.Send[EnvironmentState](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EnvironmentState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class
type EnvironmentState struct {
	objectivec.Object
}

// EnvironmentStateFrom constructs a [EnvironmentState] from an unsafe.Pointer.
func EnvironmentStateFrom(ptr unsafe.Pointer) EnvironmentState {
	return EnvironmentState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EnvironmentState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EnvironmentState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EnvironmentState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EnvironmentState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EnvironmentState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/allMechanisms
func (e_ EnvironmentState) AllMechanisms() []EnvironmentMechanism {
	rv := objc.Send[[]EnvironmentMechanism](e_.ID, objc.Sel("allMechanisms"))
	return rv
}/* debug [instance_properties/getter]: allMechanisms */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/biometry
func (e_ EnvironmentState) Biometry() ILAEnvironmentMechanismBiometry {
	rv := objc.Send[EnvironmentMechanismBiometry](e_.ID, objc.Sel("biometry"))
	return rv
}/* debug [instance_properties/getter]: biometry */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/companions
func (e_ EnvironmentState) Companions() []EnvironmentMechanismCompanion {
	rv := objc.Send[[]EnvironmentMechanismCompanion](e_.ID, objc.Sel("companions"))
	return rv
}/* debug [instance_properties/getter]: companions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAEnvironment/State-swift.class/userPassword
func (e_ EnvironmentState) UserPassword() ILAEnvironmentMechanismUserPassword {
	rv := objc.Send[EnvironmentMechanismUserPassword](e_.ID, objc.Sel("userPassword"))
	return rv
}/* debug [instance_properties/getter]: userPassword */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LAEnvironmentState */



