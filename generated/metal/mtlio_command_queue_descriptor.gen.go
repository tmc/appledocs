// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [IOCommandQueueDescriptor] class.
var (
	IOCommandQueueDescriptorClass     _IOCommandQueueDescriptorClass
	IOCommandQueueDescriptorClassOnce sync.Once
)

func getIOCommandQueueDescriptorClass() _IOCommandQueueDescriptorClass {
	IOCommandQueueDescriptorClassOnce.Do(func() {
		IOCommandQueueDescriptorClass = _IOCommandQueueDescriptorClass{objc.GetClass("MTLIOCommandQueueDescriptor")}
	})
	return IOCommandQueueDescriptorClass
}

type _IOCommandQueueDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [IOCommandQueueDescriptor] class.
type IIOCommandQueueDescriptor interface {
	objectivec.IObject
	

	// properties:
	MaxCommandBufferCount() uint
	SetMaxCommandBufferCount(value uint)
	MaxCommandsInFlight() uint
	SetMaxCommandsInFlight(value uint)
	Priority() IOPriority
	SetPriority(value IOPriority)
	ScratchBufferAllocator() unsafe.Pointer
	SetScratchBufferAllocator(value unsafe.Pointer)
	Type() IOCommandQueueType
	SetType(value IOCommandQueueType)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _IOCommandQueueDescriptorClass) Alloc() IOCommandQueueDescriptor {
	rv := objc.Send[IOCommandQueueDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _IOCommandQueueDescriptorClass) New() IOCommandQueueDescriptor {
	rv := objc.Send[IOCommandQueueDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ IOCommandQueueDescriptor) Init() IOCommandQueueDescriptor {
	rv := objc.Send[IOCommandQueueDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ IOCommandQueueDescriptor) Autorelease() IOCommandQueueDescriptor {
	rv := objc.Send[IOCommandQueueDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOCommandQueueDescriptor creates a new IOCommandQueueDescriptor instance.
func NewIOCommandQueueDescriptor() IOCommandQueueDescriptor {
	return getIOCommandQueueDescriptorClass().New()
}





// A configuration template you use to create a new input/output command queue.
//
// Use this descriptor type to configure the settings of each input/output command queue that you create using . To create additional input/output command queues, you can reuse a descriptor instance and optionally reconfigure its properties. Create each input/output queue to meet your apps needs by setting the descriptor’s properties. Select a queue’s relative level of importance with the property. Create a queue that runs multiple input/output command buffers in parallel by setting the property to . Decide how many individual commands a queue can run simultaneously with the property. Choose how many command buffers a queue can have waiting to run with property. Take control of the queue’s scratch memory allocation by implementing and assign an instance of it to the property.


// A configuration template you use to create a new input/output command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor
type IOCommandQueueDescriptor struct {
	objectivec.Object
}

// IOCommandQueueDescriptorFrom constructs a [IOCommandQueueDescriptor] from an unsafe.Pointer.
//
// A configuration template you use to create a new input/output command queue.
func IOCommandQueueDescriptorFrom(ptr unsafe.Pointer) IOCommandQueueDescriptor {
	return IOCommandQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Sets the largest number of outstanding input/output command buffers a queue can have at any point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandBufferCount
func (c_ IOCommandQueueDescriptor) MaxCommandBufferCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCommandBufferCount"))
	return rv
}


// Sets the largest number of outstanding input/output command buffers a queue can have at any point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandBufferCount
func (c_ IOCommandQueueDescriptor) SetMaxCommandBufferCount(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCommandBufferCount:"), value)
}


// Sets the largest number of individual commands that an input/output command queue can run at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandsInFlight
func (c_ IOCommandQueueDescriptor) MaxCommandsInFlight() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCommandsInFlight"))
	return rv
}


// Sets the largest number of individual commands that an input/output command queue can run at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandsInFlight
func (c_ IOCommandQueueDescriptor) SetMaxCommandsInFlight(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCommandsInFlight:"), value)
}


// Configures the priority for a new input/output command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/priority
func (c_ IOCommandQueueDescriptor) Priority() IOPriority {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("priority"))
	return rv
}


// Configures the priority for a new input/output command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/priority
func (c_ IOCommandQueueDescriptor) SetPriority(value IOPriority) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPriority:"), value)
}


// An optional memory allocator that you implement to manage the scratch memory that an input/output command queue requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/scratchBufferAllocator
func (c_ IOCommandQueueDescriptor) ScratchBufferAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scratchBufferAllocator"))
	return rv
}


// An optional memory allocator that you implement to manage the scratch memory that an input/output command queue requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/scratchBufferAllocator
func (c_ IOCommandQueueDescriptor) SetScratchBufferAllocator(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScratchBufferAllocator:"), value)
}


// Configures the queue type for a new input/output command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/type
func (c_ IOCommandQueueDescriptor) Type() IOCommandQueueType {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("type"))
	return rv
}


// Configures the queue type for a new input/output command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/type
func (c_ IOCommandQueueDescriptor) SetType(value IOCommandQueueType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}








