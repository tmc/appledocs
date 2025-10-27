// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CommandBufferDescriptor] class.
var (
	CommandBufferDescriptorClass     _CommandBufferDescriptorClass
	CommandBufferDescriptorClassOnce sync.Once
)

func getCommandBufferDescriptorClass() _CommandBufferDescriptorClass {
	CommandBufferDescriptorClassOnce.Do(func() {
		CommandBufferDescriptorClass = _CommandBufferDescriptorClass{objc.GetClass("MTLCommandBufferDescriptor")}
	})
	return CommandBufferDescriptorClass
}

type _CommandBufferDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CommandBufferDescriptor] class.
type ICommandBufferDescriptor interface {
	objectivec.IObject
	

	// properties:
	ErrorOptions() CommandBufferErrorOption
	SetErrorOptions(value CommandBufferErrorOption)
	LogState() unsafe.Pointer
	SetLogState(value unsafe.Pointer)
	RetainedReferences() bool
	SetRetainedReferences(value bool)
	MTLCommandBufferErrorDomain() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CommandBufferDescriptorClass) Alloc() CommandBufferDescriptor {
	rv := objc.Send[CommandBufferDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CommandBufferDescriptorClass) New() CommandBufferDescriptor {
	rv := objc.Send[CommandBufferDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CommandBufferDescriptor) Init() CommandBufferDescriptor {
	rv := objc.Send[CommandBufferDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CommandBufferDescriptor) Autorelease() CommandBufferDescriptor {
	rv := objc.Send[CommandBufferDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCommandBufferDescriptor creates a new CommandBufferDescriptor instance.
func NewCommandBufferDescriptor() CommandBufferDescriptor {
	return getCommandBufferDescriptorClass().New()
}





// A configuration that customizes the behavior for a new command buffer.
//
// Create a command buffer with a custom configuration by creating an instance and passing it to an instance’s method. You can configure whether the command buffer retains references to resources that its commands refer to with the property. The command buffer can save extra error information, which is useful during development, by setting its property to .


// A configuration that customizes the behavior for a new command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor
type CommandBufferDescriptor struct {
	objectivec.Object
}

// CommandBufferDescriptorFrom constructs a [CommandBufferDescriptor] from an unsafe.Pointer.
//
// A configuration that customizes the behavior for a new command buffer.
func CommandBufferDescriptorFrom(ptr unsafe.Pointer) CommandBufferDescriptor {
	return CommandBufferDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/errorOptions
func (c_ CommandBufferDescriptor) ErrorOptions() CommandBufferErrorOption {
	rv := objc.Send[CommandBufferErrorOption](c_.ID, objc.Sel("errorOptions"))
	return rv
}


// The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/errorOptions
func (c_ CommandBufferDescriptor) SetErrorOptions(value CommandBufferErrorOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setErrorOptions:"), value)
}


// The shader logging configuration that the command buffer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/logState
func (c_ CommandBufferDescriptor) LogState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("logState"))
	return rv
}


// The shader logging configuration that the command buffer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/logState
func (c_ CommandBufferDescriptor) SetLogState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLogState:"), value)
}


// A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/retainedReferences
func (c_ CommandBufferDescriptor) RetainedReferences() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("retainedReferences"))
	return rv
}


// A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/retainedReferences
func (c_ CommandBufferDescriptor) SetRetainedReferences(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRetainedReferences:"), value)
}


// The domain for Metal command buffer errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffererrordomain
func (c_ CommandBufferDescriptor) MTLCommandBufferErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("MTLCommandBufferErrorDomain"))
	return rv
}








