// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GraphExecutableExecutionDescriptor] class.
type IGraphExecutableExecutionDescriptor interface {
	IGraphObject
	// properties:
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
	CompletionHandler() GraphExecutableCompletionHandler /* not a class type */
	SetCompletionHandler(value GraphExecutableCompletionHandler /* not a class type */)
	ScheduledHandler() GraphExecutableScheduledHandler /* not a class type */
	SetScheduledHandler(value GraphExecutableScheduledHandler /* not a class type */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableExecutionDescriptorClass) Alloc() GraphExecutableExecutionDescriptor {
	rv := objc.Send[GraphExecutableExecutionDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Flag for the graph executable to wait till the execution has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitUntilCompleted"))
	return rv
}


// Flag for the graph executable to wait till the execution has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitUntilCompleted:"), value)
}


// A notification that appears when graph-executable execution is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/completionhandler
func (g_ GraphExecutableExecutionDescriptor) CompletionHandler() GraphExecutableCompletionHandler /* not a class type */ {
	rv := objc.Send[GraphExecutableCompletionHandler](g_.ID, objc.Sel("completionHandler"))
	return rv
}


// A notification that appears when graph-executable execution is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/completionhandler
func (g_ GraphExecutableExecutionDescriptor) SetCompletionHandler(value GraphExecutableCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompletionHandler:"), value)
}


// A notification that appears when graph-executable execution is scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/scheduledhandler
func (g_ GraphExecutableExecutionDescriptor) ScheduledHandler() GraphExecutableScheduledHandler /* not a class type */ {
	rv := objc.Send[GraphExecutableScheduledHandler](g_.ID, objc.Sel("scheduledHandler"))
	return rv
}


// A notification that appears when graph-executable execution is scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/scheduledhandler
func (g_ GraphExecutableExecutionDescriptor) SetScheduledHandler(value GraphExecutableScheduledHandler /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScheduledHandler:"), value)
}



