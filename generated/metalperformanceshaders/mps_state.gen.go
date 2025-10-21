// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An opaque data container for large storage in MPS CNN filters.
//
// Some MPS CNN kernels produce additional information beyond an . These may be pooling indices where the result came from, convolution weights, or other information not contained in the usual result from a . An object typically contains one or more expensive objects such as textures or buffers to store this information. It provides a base class with interfaces for managing this storage. Child classes may add additional functionality specific to their contents. Some objects are temporary. Temporary state objects, for example, and , are for very short lived storage, perhaps just a few lines of code within the scope of a single . They are very efficient for storage, as several temporary objects can share the same memory over the course of a command buffer. This can improve both memory usage and time spent in the kernel wiring down memory and such. You may find that some large CNN tasks can not be computed without them, as nontemporary storage would simply take up too much memory. In exchange, the lifetime of the underlying storage in temporary objects needs to be carefully managed. ARC often waits until the end of scope to release objects. Temporary storage often needs to be released sooner than that. Consequently the lifetime of the data in the underlying Metal resources is managed by a property. Each time a reads a temporary object the is automatically decremented. When it reaches 0, the underlying storage is recycled for use by other MPS temporary objects, and the data is becomes undefined. If you need to consume the data multiple times, you should set the to a larger number to prevent the data from becoming undefined. You may set the to 0 yourself to return the storage to MPS, if for any reason, you realize that the object will no longer be used. The contents of a temporary object are only valid from creation to the time the reaches 0. The data is only valid for the on which it was created. Nontemporary objects are valid on any on the same device until they are released.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState
type State struct {
	objectivec.Object
}

// StateFrom constructs a [State] from an unsafe.Pointer.
//
// An opaque data container for large storage in MPS CNN filters.
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




