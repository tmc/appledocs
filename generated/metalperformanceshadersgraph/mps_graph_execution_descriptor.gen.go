// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	CompilationDescriptor() IMPSGraphCompilationDescriptor
	SetCompilationDescriptor(value IMPSGraphCompilationDescriptor)
	CompletionHandler() GraphCompletionHandler /* not a class type */
	SetCompletionHandler(value GraphCompletionHandler /* not a class type */)
	ScheduledHandler() GraphScheduledHandler /* not a class type */
	SetScheduledHandler(value GraphScheduledHandler /* not a class type */)
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
	// methods:
	SignalEventAtExecutionEventValue(event objectivec.IObject, executionStage GraphExecutionStage, value uint64)
}

// A class that consists of all the levers to synchronize and schedule graph execution.


// A class that consists of all the levers to synchronize and schedule graph execution.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g_ GraphExecutionDescriptor) SignalEventAtExecutionEventValue(event objectivec.IObject, executionStage GraphExecutionStage, value uint64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}


// The compilation descriptor for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/compilationdescriptor
func (g_ GraphExecutionDescriptor) CompilationDescriptor() IMPSGraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](g_.ID, objc.Sel("compilationDescriptor"))
	return rv
}


// The compilation descriptor for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/compilationdescriptor
func (g_ GraphExecutionDescriptor) SetCompilationDescriptor(value IMPSGraphCompilationDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompilationDescriptor:"), value)
}


// The handler that graph calls at the completion of the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/completionhandler
func (g_ GraphExecutionDescriptor) CompletionHandler() GraphCompletionHandler /* not a class type */ {
	rv := objc.Send[GraphCompletionHandler](g_.ID, objc.Sel("completionHandler"))
	return rv
}


// The handler that graph calls at the completion of the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/completionhandler
func (g_ GraphExecutionDescriptor) SetCompletionHandler(value GraphCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompletionHandler:"), value)
}


// The handler that graph calls when it schedules the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/scheduledhandler
func (g_ GraphExecutionDescriptor) ScheduledHandler() GraphScheduledHandler /* not a class type */ {
	rv := objc.Send[GraphScheduledHandler](g_.ID, objc.Sel("scheduledHandler"))
	return rv
}


// The handler that graph calls when it schedules the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/scheduledhandler
func (g_ GraphExecutionDescriptor) SetScheduledHandler(value GraphScheduledHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScheduledHandler:"), value)
}


// The flag that blocks the execution call until the entire execution is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/waituntilcompleted
func (g_ GraphExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitUntilCompleted"))
	return rv
}


// The flag that blocks the execution call until the entire execution is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/waituntilcompleted
func (g_ GraphExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitUntilCompleted:"), value)
}



