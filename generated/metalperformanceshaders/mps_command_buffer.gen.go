// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CommandBuffer] class.
var (
	CommandBufferClass     _CommandBufferClass
	CommandBufferClassOnce sync.Once
)

func getCommandBufferClass() _CommandBufferClass {
	CommandBufferClassOnce.Do(func() {
		CommandBufferClass = _CommandBufferClass{objc.GetClass("MPSCommandBuffer")}
	})
	return CommandBufferClass
}

type _CommandBufferClass struct {
	class objc.Class
}

// An interface definition for the [CommandBuffer] class.
type ICommandBuffer interface {
	objectivec.IObject
	PrefetchHeapForWorkloadSize(size unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer
type CommandBuffer struct {
	objectivec.Object
}

// CommandBufferFrom constructs a [CommandBuffer] from an unsafe.Pointer.
func CommandBufferFrom(ptr unsafe.Pointer) CommandBuffer {
	return CommandBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CommandBufferClass) Alloc() CommandBuffer {
	rv := objc.Send[CommandBuffer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CommandBufferClass) New() CommandBuffer {
	rv := objc.Send[CommandBuffer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CommandBuffer) Init() CommandBuffer {
	rv := objc.Send[CommandBuffer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CommandBuffer) Autorelease() CommandBuffer {
	rv := objc.Send[CommandBuffer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCommandBuffer creates a new CommandBuffer instance.
func NewCommandBuffer() CommandBuffer {
	return getCommandBufferClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/prefetchHeap(forWorkloadSize:)
func (c_ CommandBuffer) PrefetchHeapForWorkloadSize(size unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prefetchHeapForWorkloadSize:"), size)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/commandBuffer
func (c_ CommandBuffer) CommandBuffer() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("commandBuffer"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/heapProvider
func (c_ CommandBuffer) HeapProvider() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("heapProvider"))
	return rv
}


// SetHeapProvider sets the value of the heapProvider property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/heapProvider
func (c_ CommandBuffer) SetHeapProvider(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeapProvider:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/predicate
func (c_ CommandBuffer) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/predicate
func (c_ CommandBuffer) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicate:"), value)
}


