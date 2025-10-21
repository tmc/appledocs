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
}

// A class that consists of all the levers to synchronize and schedule executable execution.
//
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


// A notification that appears when graph-executable execution is finished.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/completionHandler
func (g_ GraphExecutableExecutionDescriptor) CompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("completionHandler"))
	return rv
}


// SetCompletionHandler sets the value of the completionHandler property.
// A notification that appears when graph-executable execution is finished.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/completionHandler
func (g_ GraphExecutableExecutionDescriptor) SetCompletionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompletionHandler:"), value)
}

// A notification that appears when graph-executable execution is scheduled.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/scheduledHandler
func (g_ GraphExecutableExecutionDescriptor) ScheduledHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("scheduledHandler"))
	return rv
}


// SetScheduledHandler sets the value of the scheduledHandler property.
// A notification that appears when graph-executable execution is scheduled.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/scheduledHandler
func (g_ GraphExecutableExecutionDescriptor) SetScheduledHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScheduledHandler:"), value)
}

// Flag for the graph executable to wait till the execution has completed.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("waitUntilCompleted"))
	return rv
}


// SetWaitUntilCompleted sets the value of the waitUntilCompleted property.
// Flag for the graph executable to wait till the execution has completed.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g_ GraphExecutableExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWaitUntilCompleted:"), value)
}



