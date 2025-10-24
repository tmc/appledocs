// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphExecutionDescriptor */


/* debug [class_header]: Header for MPSGraphExecutionDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphExecutionDescriptor */
// An interface definition for the [GraphExecutionDescriptor] class.
type IGraphExecutionDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphExecutionDescriptor */
	// properties:
	CompilationDescriptor() IMPSGraphCompilationDescriptor
	SetCompilationDescriptor(value IMPSGraphCompilationDescriptor)
	CompletionHandler() GraphCompletionHandler /* not a class type */
	SetCompletionHandler(value GraphCompletionHandler /* not a class type */)
	ScheduledHandler() GraphScheduledHandler /* not a class type */
	SetScheduledHandler(value GraphScheduledHandler /* not a class type */)
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphExecutionDescriptor */
	// methods:
	SignalEventAtExecutionEventValue(event unsafe.Pointer, executionStage GraphExecutionStage, value uint64)
	WaitForEventValue(event unsafe.Pointer, value uint64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphExecutionDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphExecutionDescriptorClass) Alloc() GraphExecutionDescriptor {
	rv := objc.Send[GraphExecutionDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphExecutionDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphExecutionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphExecutionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphExecutionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphExecutionDescriptor */

// Executable signals these shared events at execution stage and immediately proceeds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g_ GraphExecutionDescriptor) SignalEventAtExecutionEventValue(event unsafe.Pointer, executionStage GraphExecutionStage, value uint64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}/* debug [instance_methods/method]: SignalEventAtExecutionEventValue */


// Executable waits on these shared events before scheduling execution on the HW, this does not include encoding which can still continue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/wait(for:value:)
func (g_ GraphExecutionDescriptor) WaitForEventValue(event unsafe.Pointer, value uint64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("waitForEvent:value:"), event, value)
}/* debug [instance_methods/method]: WaitForEventValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphExecutionDescriptor */

// The compilation descriptor for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/compilationDescriptor
func (g_ GraphExecutionDescriptor) CompilationDescriptor() IMPSGraphCompilationDescriptor {
	rv := objc.Send[GraphCompilationDescriptor](g_.ID, objc.Sel("compilationDescriptor"))
	return rv
}/* debug [instance_properties/getter]: compilationDescriptor */


// The compilation descriptor for the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/compilationDescriptor
func (g_ GraphExecutionDescriptor) SetCompilationDescriptor(value IMPSGraphCompilationDescriptor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompilationDescriptor:"), value)
}/* debug [instance_properties/setter]: compilationDescriptor */


// The handler that graph calls at the completion of the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/completionHandler
func (g_ GraphExecutionDescriptor) CompletionHandler() GraphCompletionHandler /* not a class type */ {
	rv := objc.Send[GraphCompletionHandler](g_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// The handler that graph calls at the completion of the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/completionHandler
func (g_ GraphExecutionDescriptor) SetCompletionHandler(value GraphCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompletionHandler:"), value)
}/* debug [instance_properties/setter]: completionHandler */


// The handler that graph calls when it schedules the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/scheduledHandler
func (g_ GraphExecutionDescriptor) ScheduledHandler() GraphScheduledHandler /* not a class type */ {
	rv := objc.Send[GraphScheduledHandler](g_.ID, objc.Sel("scheduledHandler"))
	return rv
}/* debug [instance_properties/getter]: scheduledHandler */


// The handler that graph calls when it schedules the execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/scheduledHandler
func (g_ GraphExecutionDescriptor) SetScheduledHandler(value GraphScheduledHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScheduledHandler:"), value)
}/* debug [instance_properties/setter]: scheduledHandler */


// The flag that blocks the execution call until the entire execution is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitUntilCompleted"))
	return rv
}/* debug [instance_properties/getter]: waitUntilCompleted */


// The flag that blocks the execution call until the entire execution is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitUntilCompleted:"), value)
}/* debug [instance_properties/setter]: waitUntilCompleted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphExecutionDescriptor */



