// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCommandBuffer */


/* debug [class_header]: Header for MPSCommandBuffer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CommandBuffer */
// An interface definition for the [CommandBuffer] class.
type ICommandBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CommandBuffer */
	// properties:
	CommandBuffer() CommandBuffer get /* not a class type */
	SetCommandBuffer(value CommandBuffer get /* not a class type */)
	Predicate() IMPSPredicate
	SetPredicate(value IMPSPredicate)
	RootCommandBuffer() CommandBuffer get /* not a class type */
	SetRootCommandBuffer(value CommandBuffer get /* not a class type */)
	HeapProvider() HeapProvider get set /* not a class type */
	SetHeapProvider(value HeapProvider get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CommandBuffer */
	// methods:
	CommitAndContinue()
	PrefetchHeap()
	PrefetchHeapForWorkloadSize(size uintptr /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CommandBuffer */
// Alloc allocates a new instance without initialization.
func (cc _CommandBufferClass) Alloc() CommandBuffer {
	rv := objc.Send[CommandBuffer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CommandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer
type CommandBuffer struct {
	objectivec.Object
}

// CommandBufferFrom constructs a [CommandBuffer] from an unsafe.Pointer.
func CommandBufferFrom(ptr unsafe.Pointer) CommandBuffer {
	return CommandBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CommandBuffer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114031-initwithcommandbuffer
func NewCommandBufferWithCommandBuffer(commandBuffer unsafe.Pointer) CommandBuffer {
	instance := getCommandBufferClass().Alloc()
	rv := objc.Send[CommandBuffer](instance.ID, objc.Sel("initWithCommandBuffer:"), commandBuffer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCommandBufferWithCommandBuffer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CommandBuffer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114029-commandbufferfromcommandqueue
func (cc _CommandBufferClass) CommandBufferFromCommandQueue(commandQueue unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("commandBufferFromCommandQueue:"), commandQueue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CommandBufferFromCommandQueue) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114030-commandbufferwithcommandbuffer
func (cc _CommandBufferClass) CommandBufferWithCommandBuffer(commandBuffer unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("commandBufferWithCommandBuffer:"), commandBuffer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CommandBufferWithCommandBuffer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CommandBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CommandBuffer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3152524-commitandcontinue
func (c_ CommandBuffer) CommitAndContinue() {
	objc.Send[objc.ID](c_.ID, objc.Sel("commitAndContinue"))
}/* debug [instance_methods/method]: CommitAndContinue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229858-prefetchheap
func (c_ CommandBuffer) PrefetchHeap() {
	objc.Send[objc.ID](c_.ID, objc.Sel("prefetchHeap"))
}/* debug [instance_methods/method]: PrefetchHeap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229858-prefetchheapforworkloadsize
func (c_ CommandBuffer) PrefetchHeapForWorkloadSize(size uintptr /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prefetchHeapForWorkloadSize:"), size)
}/* debug [instance_methods/method]: PrefetchHeapForWorkloadSize */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CommandBuffer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114028-commandbuffer
func (c_ CommandBuffer) CommandBuffer() CommandBuffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("commandBuffer"))
	return rv
}/* debug [instance_properties/getter]: commandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114028-commandbuffer
func (c_ CommandBuffer) SetCommandBuffer(value CommandBuffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCommandBuffer:"), value)
}/* debug [instance_properties/setter]: commandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114032-predicate
func (c_ CommandBuffer) Predicate() IMPSPredicate {
	rv := objc.Send[Predicate](c_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3114032-predicate
func (c_ CommandBuffer) SetPredicate(value IMPSPredicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicate:"), value)
}/* debug [instance_properties/setter]: predicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3166772-rootcommandbuffer
func (c_ CommandBuffer) RootCommandBuffer() CommandBuffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("rootCommandBuffer"))
	return rv
}/* debug [instance_properties/getter]: rootCommandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3166772-rootcommandbuffer
func (c_ CommandBuffer) SetRootCommandBuffer(value CommandBuffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootCommandBuffer:"), value)
}/* debug [instance_properties/setter]: rootCommandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229857-heapprovider
func (c_ CommandBuffer) HeapProvider() HeapProvider get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("heapProvider"))
	return rv
}/* debug [instance_properties/getter]: heapProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscommandbuffer/3229857-heapprovider
func (c_ CommandBuffer) SetHeapProvider(value HeapProvider get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeapProvider:"), value)
}/* debug [instance_properties/setter]: heapProvider */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCommandBuffer */


