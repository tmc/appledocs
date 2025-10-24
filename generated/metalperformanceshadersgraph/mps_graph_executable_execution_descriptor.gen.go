// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphExecutableExecutionDescriptor */


/* debug [class_header]: Header for MPSGraphExecutableExecutionDescriptor */
// The class instance for the [GraphExecutableExecutionDescriptor] class.
var (
	GraphExecutableExecutionDescriptorClass     _GraphExecutableExecutionDescriptorClass
	GraphExecutableExecutionDescriptorClassOnce sync.Once
)

func getGraphExecutableExecutionDescriptorClass() _GraphExecutableExecutionDescriptorClass {
	GraphExecutableExecutionDescriptorClassOnce.Do(func() {
		GraphExecutableExecutionDescriptorClass = _GraphExecutableExecutionDescriptorClass{objc.GetClass("MPSGraphExecutableExecutionDescriptor")}
	})
	return GraphExecutableExecutionDescriptorClass
}

type _GraphExecutableExecutionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphExecutableExecutionDescriptor */
// An interface definition for the [GraphExecutableExecutionDescriptor] class.
type IGraphExecutableExecutionDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphExecutableExecutionDescriptor */
	// properties:
	CompletionHandler() GraphExecutableCompletionHandler /* not a class type */
	SetCompletionHandler(value GraphExecutableCompletionHandler /* not a class type */)
	ScheduledHandler() GraphExecutableScheduledHandler /* not a class type */
	SetScheduledHandler(value GraphExecutableScheduledHandler /* not a class type */)
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphExecutableExecutionDescriptor */
	// methods:
	SignalEventAtExecutionEventValue(event unsafe.Pointer, executionStage GraphExecutionStage, value uint64)
	WaitForEventValue(event unsafe.Pointer, value uint64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphExecutableExecutionDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableExecutionDescriptorClass) Alloc() GraphExecutableExecutionDescriptor {
	rv := objc.Send[GraphExecutableExecutionDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphExecutableExecutionDescriptorClass) New() GraphExecutableExecutionDescriptor {
	rv := objc.Send[GraphExecutableExecutionDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutableExecutionDescriptor) Init() GraphExecutableExecutionDescriptor {
	rv := objc.Send[GraphExecutableExecutionDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutableExecutionDescriptor) Autorelease() GraphExecutableExecutionDescriptor {
	rv := objc.Send[GraphExecutableExecutionDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutableExecutionDescriptor creates a new GraphExecutableExecutionDescriptor instance.
func NewGraphExecutableExecutionDescriptor() GraphExecutableExecutionDescriptor {
	return getGraphExecutableExecutionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphExecutableExecutionDescriptor */
// A class that consists of all the levers to synchronize and schedule executable execution.


// A class that consists of all the levers to synchronize and schedule executable execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor
type GraphExecutableExecutionDescriptor struct {
	GraphObject
}

// GraphExecutableExecutionDescriptorFrom constructs a [GraphExecutableExecutionDescriptor] from an unsafe.Pointer.
//
// A class that consists of all the levers to synchronize and schedule executable execution.
func GraphExecutableExecutionDescriptorFrom(ptr unsafe.Pointer) GraphExecutableExecutionDescriptor {
	return GraphExecutableExecutionDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphExecutableExecutionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphExecutableExecutionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphExecutableExecutionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphExecutableExecutionDescriptor */

// Signals these shared events at execution stage and immediately proceeds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g_ GraphExecutableExecutionDescriptor) SignalEventAtExecutionEventValue(event unsafe.Pointer, executionStage GraphExecutionStage, value uint64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}/* debug [instance_methods/method]: SignalEventAtExecutionEventValue */


// Waits on these shared events before scheduling execution on the HW.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/wait(for:value:)
func (g_ GraphExecutableExecutionDescriptor) WaitForEventValue(event unsafe.Pointer, value uint64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("waitForEvent:value:"), event, value)
}/* debug [instance_methods/method]: WaitForEventValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphExecutableExecutionDescriptor */

// A notification that appears when graph-executable execution is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/completionHandler
func (g_ GraphExecutableExecutionDescriptor) CompletionHandler() GraphExecutableCompletionHandler /* not a class type */ {
	rv := objc.Send[GraphExecutableCompletionHandler](g_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// A notification that appears when graph-executable execution is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/completionHandler
func (g_ GraphExecutableExecutionDescriptor) SetCompletionHandler(value GraphExecutableCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompletionHandler:"), value)
}/* debug [instance_properties/setter]: completionHandler */


// A notification that appears when graph-executable execution is scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/scheduledHandler
func (g_ GraphExecutableExecutionDescriptor) ScheduledHandler() GraphExecutableScheduledHandler /* not a class type */ {
	rv := objc.Send[GraphExecutableScheduledHandler](g_.ID, objc.Sel("scheduledHandler"))
	return rv
}/* debug [instance_properties/getter]: scheduledHandler */


// A notification that appears when graph-executable execution is scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/scheduledHandler
func (g_ GraphExecutableExecutionDescriptor) SetScheduledHandler(value GraphExecutableScheduledHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScheduledHandler:"), value)
}/* debug [instance_properties/setter]: scheduledHandler */


// Flag for the graph executable to wait till the execution has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitUntilCompleted"))
	return rv
}/* debug [instance_properties/getter]: waitUntilCompleted */


// Flag for the graph executable to wait till the execution has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitUntilCompleted:"), value)
}/* debug [instance_properties/setter]: waitUntilCompleted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphExecutableExecutionDescriptor */



