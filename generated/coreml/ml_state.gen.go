// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLState */


/* debug [class_header]: Header for MLState */
// The class instance for the [State] class.
var (
	StateClass     _StateClass
	StateClassOnce sync.Once
)

func getStateClass() _StateClass {
	StateClassOnce.Do(func() {
		StateClass = _StateClass{objc.GetClass("MLState")}
	})
	return StateClass
}

type _StateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for State */
// An interface definition for the [State] class.
type IState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for State */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for State */
	// methods:
	GetMultiArrayForStateNamedHandler(stateName objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for State */
// Alloc allocates a new instance without initialization.
func (sc _StateClass) Alloc() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StateClass) New() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ State) Init() State {
	rv := objc.Send[State](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ State) Autorelease() State {
	rv := objc.Send[State](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewState creates a new State instance.
func NewState() State {
	return getStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for State */
// Handle to the state buffers.
//
// A stateful model maintains a state from one prediction to another by storing the information in the state buffers. To use such a model, the client must request the model to create state buffers and get object, which is the handle to those buffers. Then, at the prediction time, pass the object in one of the stateful prediction functions. The object is a handle to the state buffers. The client shall not read or write the buffers while a prediction is in-flight. Each stateful prediction that uses the same must be serialized. Otherwise, if two such predictions run concurrently, the behavior is undefined.


// Handle to the state buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLState
type State struct {
	objectivec.Object
}

// StateFrom constructs a [State] from an unsafe.Pointer.
//
// Handle to the state buffers.
func StateFrom(ptr unsafe.Pointer) State {
	return State{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for State *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for State */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for State */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for State */

// Gets a mutable view into a state buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLState/getMultiArrayForStateNamed:handler:
func (s_ State) GetMultiArrayForStateNamedHandler(stateName objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getMultiArrayForStateNamed:handler:"), stateName, handler)
}/* debug [instance_methods/method]: GetMultiArrayForStateNamedHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for State */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLState */



