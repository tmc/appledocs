// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CommandQueueDescriptor] class.
var (
	CommandQueueDescriptorClass     _CommandQueueDescriptorClass
	CommandQueueDescriptorClassOnce sync.Once
)

func getCommandQueueDescriptorClass() _CommandQueueDescriptorClass {
	CommandQueueDescriptorClassOnce.Do(func() {
		CommandQueueDescriptorClass = _CommandQueueDescriptorClass{objc.GetClass("MTLCommandQueueDescriptor")}
	})
	return CommandQueueDescriptorClass
}

type _CommandQueueDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CommandQueueDescriptor] class.
type ICommandQueueDescriptor interface {
	objectivec.IObject
	

	// properties:
	LogState() unsafe.Pointer
	SetLogState(value unsafe.Pointer)
	MaxCommandBufferCount() uint
	SetMaxCommandBufferCount(value uint)
	MTLCommandBufferErrorDomain() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CommandQueueDescriptorClass) Alloc() CommandQueueDescriptor {
	rv := objc.Send[CommandQueueDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CommandQueueDescriptorClass) New() CommandQueueDescriptor {
	rv := objc.Send[CommandQueueDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CommandQueueDescriptor) Init() CommandQueueDescriptor {
	rv := objc.Send[CommandQueueDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CommandQueueDescriptor) Autorelease() CommandQueueDescriptor {
	rv := objc.Send[CommandQueueDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCommandQueueDescriptor creates a new CommandQueueDescriptor instance.
func NewCommandQueueDescriptor() CommandQueueDescriptor {
	return getCommandQueueDescriptorClass().New()
}





// A configuration that customizes the behavior for a new command queue.


// A configuration that customizes the behavior for a new command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor
type CommandQueueDescriptor struct {
	objectivec.Object
}

// CommandQueueDescriptorFrom constructs a [CommandQueueDescriptor] from an unsafe.Pointer.
//
// A configuration that customizes the behavior for a new command queue.
func CommandQueueDescriptorFrom(ptr unsafe.Pointer) CommandQueueDescriptor {
	return CommandQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// The shader logging configuration that the command queue uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/logState
func (c_ CommandQueueDescriptor) LogState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("logState"))
	return rv
}


// The shader logging configuration that the command queue uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/logState
func (c_ CommandQueueDescriptor) SetLogState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLogState:"), value)
}


// An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/maxCommandBufferCount
func (c_ CommandQueueDescriptor) MaxCommandBufferCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCommandBufferCount"))
	return rv
}


// An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/maxCommandBufferCount
func (c_ CommandQueueDescriptor) SetMaxCommandBufferCount(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCommandBufferCount:"), value)
}


// The domain for Metal command buffer errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffererrordomain
func (c_ CommandQueueDescriptor) MTLCommandBufferErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("MTLCommandBufferErrorDomain"))
	return rv
}








