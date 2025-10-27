// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [State] class.
type IState interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	GetMultiArrayForStateNamedHandler(stateName foundation.foundation.INSString, handler unsafe.Pointer)


}





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




















// Gets a mutable view into a state buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLState/getMultiArrayForStateNamed:handler:
func (s_ State) GetMultiArrayForStateNamedHandler(stateName foundation.foundation.INSString, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getMultiArrayForStateNamed:handler:"), stateName, handler)
}













