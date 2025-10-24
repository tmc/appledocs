// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLStateConstraint */


/* debug [class_header]: Header for MLStateConstraint */
// The class instance for the [StateConstraint] class.
var (
	StateConstraintClass     _StateConstraintClass
	StateConstraintClassOnce sync.Once
)

func getStateConstraintClass() _StateConstraintClass {
	StateConstraintClassOnce.Do(func() {
		StateConstraintClass = _StateConstraintClass{objc.GetClass("MLStateConstraint")}
	})
	return StateConstraintClass
}

type _StateConstraintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StateConstraint */
// An interface definition for the [StateConstraint] class.
type IStateConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StateConstraint */
	// properties:
	BufferShape() []foundation.Number
	DataType() MultiArrayDataType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StateConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StateConstraint */
// Alloc allocates a new instance without initialization.
func (sc _StateConstraintClass) Alloc() StateConstraint {
	rv := objc.Send[StateConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StateConstraintClass) New() StateConstraint {
	rv := objc.Send[StateConstraint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateConstraint) Init() StateConstraint {
	rv := objc.Send[StateConstraint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateConstraint) Autorelease() StateConstraint {
	rv := objc.Send[StateConstraint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateConstraint creates a new StateConstraint instance.
func NewStateConstraint() StateConstraint {
	return getStateConstraintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StateConstraint */
// Constraint of a state feature value.


// Constraint of a state feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint
type StateConstraint struct {
	objectivec.Object
}

// StateConstraintFrom constructs a [StateConstraint] from an unsafe.Pointer.
//
// Constraint of a state feature value.
func StateConstraintFrom(ptr unsafe.Pointer) StateConstraint {
	return StateConstraint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StateConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StateConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StateConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StateConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StateConstraint */

// The shape of the state buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint/bufferShape-6o5vn
func (s_ StateConstraint) BufferShape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](s_.ID, objc.Sel("bufferShape"))
	return rv
}/* debug [instance_properties/getter]: bufferShape */


// The data type of scalars in the state buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint/dataType
func (s_ StateConstraint) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](s_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLStateConstraint */



