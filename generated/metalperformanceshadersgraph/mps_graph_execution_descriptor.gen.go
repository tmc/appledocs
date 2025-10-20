// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphExecutionDescriptor] class.
var (
	GraphExecutionDescriptorClass     _GraphExecutionDescriptorClass
	GraphExecutionDescriptorClassOnce sync.Once
)

func getGraphExecutionDescriptorClass() _GraphExecutionDescriptorClass {
	GraphExecutionDescriptorClassOnce.Do(func() {
		GraphExecutionDescriptorClass = _GraphExecutionDescriptorClass{objc.GetClass("MPSGraphExecutionDescriptor")}
	})
	return GraphExecutionDescriptorClass
}

type _GraphExecutionDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphExecutionDescriptor] class.
type IGraphExecutionDescriptor interface {
	IGraphObject
	SignalEventAtExecutionEventValue(event objc.ID, executionStage unsafe.Pointer, value unsafe.Pointer)
	WaitForEventValue(event objc.ID, value unsafe.Pointer)
}

// A class that consists of all the levers to synchronize and schedule graph execution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor
type GraphExecutionDescriptor struct {
	GraphObject
}

// GraphExecutionDescriptorFrom constructs a [GraphExecutionDescriptor] from an unsafe.Pointer.
//
// A class that consists of all the levers to synchronize and schedule graph execution.
func GraphExecutionDescriptorFrom(ptr unsafe.Pointer) GraphExecutionDescriptor {
	return GraphExecutionDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphExecutionDescriptorClass) Alloc() GraphExecutionDescriptor {
	rv := objc.Send[GraphExecutionDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphExecutionDescriptorClass) New() GraphExecutionDescriptor {
	rv := objc.Send[GraphExecutionDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutionDescriptor) Init() GraphExecutionDescriptor {
	rv := objc.Send[GraphExecutionDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutionDescriptor) Autorelease() GraphExecutionDescriptor {
	rv := objc.Send[GraphExecutionDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutionDescriptor creates a new GraphExecutionDescriptor instance.
func NewGraphExecutionDescriptor() GraphExecutionDescriptor {
	return getGraphExecutionDescriptorClass().New()
}


// Executable signals these shared events at execution stage and immediately proceeds.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g_ GraphExecutionDescriptor) SignalEventAtExecutionEventValue(event objc.ID, executionStage unsafe.Pointer, value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}

// Executable waits on these shared events before scheduling execution on the HW, this does not include encoding which can still continue.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/wait(for:value:)
func (g_ GraphExecutionDescriptor) WaitForEventValue(event objc.ID, value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("waitForEvent:value:"), event, value)
}



