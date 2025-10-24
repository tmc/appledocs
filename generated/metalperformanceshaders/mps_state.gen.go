// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

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
		StateClass = _StateClass{objc.GetClass("MPSState")}
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
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type State struct {
	objectivec.Object
}

// StateFrom constructs a [State] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func StateFrom(ptr unsafe.Pointer) State {
	return State{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StateClass) Alloc() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




