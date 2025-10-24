// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLCommandQueueDescriptor */


/* debug [class_header]: Header for MTLCommandQueueDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CommandQueueDescriptor */
// An interface definition for the [CommandQueueDescriptor] class.
type ICommandQueueDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CommandQueueDescriptor */
	// properties:
	LogState() unsafe.Pointer
	SetLogState(value unsafe.Pointer)
	MaxCommandBufferCount() uint
	SetMaxCommandBufferCount(value uint)
	MTLCommandBufferErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CommandQueueDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CommandQueueDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CommandQueueDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CommandQueueDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CommandQueueDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CommandQueueDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CommandQueueDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CommandQueueDescriptor */

// The shader logging configuration that the command queue uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/logState
func (c_ CommandQueueDescriptor) LogState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("logState"))
	return rv
}/* debug [instance_properties/getter]: logState */


// The shader logging configuration that the command queue uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/logState
func (c_ CommandQueueDescriptor) SetLogState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLogState:"), value)
}/* debug [instance_properties/setter]: logState */


// An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/maxCommandBufferCount
func (c_ CommandQueueDescriptor) MaxCommandBufferCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCommandBufferCount"))
	return rv
}/* debug [instance_properties/getter]: maxCommandBufferCount */


// An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/maxCommandBufferCount
func (c_ CommandQueueDescriptor) SetMaxCommandBufferCount(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCommandBufferCount:"), value)
}/* debug [instance_properties/setter]: maxCommandBufferCount */


// The domain for Metal command buffer errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffererrordomain
func (c_ CommandQueueDescriptor) MTLCommandBufferErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("MTLCommandBufferErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTLCommandBufferErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLCommandQueueDescriptor */



