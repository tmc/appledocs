// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZVirtualMachine] class.
var (
	VZVirtualMachineClass     _VZVirtualMachineClass
	VZVirtualMachineClassOnce sync.Once
)

func getVZVirtualMachineClass() _VZVirtualMachineClass {
	VZVirtualMachineClassOnce.Do(func() {
		VZVirtualMachineClass = _VZVirtualMachineClass{objc.GetClass("VZVirtualMachine")}
	})
	return VZVirtualMachineClass
}

type _VZVirtualMachineClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtualMachine] class.
type IVZVirtualMachine interface {
	objectivec.IObject
	RestoreMachineStateFromURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that manages the overall state and configuration of your VM.
//
// A object emulates a complete hardware machine of the same architecture as the underlying Mac computer. Use the VM to execute a guest operating system and any other apps you install. The VM manages the resources that the guest operating system uses, providing access to some hardware resources while emulating others. Create and configure a object with details about how you want to configure your VM, and use that object to create the object. After creating the VM, call the method (Swift) or the method (Objective-C) to start the VM and boot the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine
type VZVirtualMachine struct {
	objectivec.Object
}

// VZVirtualMachineFrom constructs a [VZVirtualMachine] from an unsafe.Pointer.
//
// An object that manages the overall state and configuration of your VM.
func VZVirtualMachineFrom(ptr unsafe.Pointer) VZVirtualMachine {
	return VZVirtualMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineClass) Alloc() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtualMachineClass) New() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// IsSupported returns a Boolean value that indicates whether the Virtualization framework is available.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isVirtualizationSupported
func (vc _VZVirtualMachineClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("isVirtualizationSupported"))
	return rv
}

// Init initializes the instance.
func (v_ VZVirtualMachine) Init() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtualMachine) Autorelease() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachine creates a new VZVirtualMachine instance.
func NewVZVirtualMachine() VZVirtualMachine {
	return getVZVirtualMachineClass().New()
}

// NewVZVirtualMachineWithConfiguration initializes a virtual machine with a configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:)
func NewVZVirtualMachineWithConfiguration(configuration VZVirtualMachineConfiguration) VZVirtualMachine {
	class := getVZVirtualMachineClass()
	alloc := objc.Send[VZVirtualMachine](objc.ID(class.class), objc.Sel("alloc"))
	inst := objc.Send[VZVirtualMachine](
		alloc.ID,
		objc.Sel("initWithConfiguration:"),
		unsafe.Pointer(configuration.ID),
	)
	return inst
}

// Restores a VM from a previously saved state.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/restoreMachineStateFrom(url:completionHandler:)
func (v_ VZVirtualMachine) RestoreMachineStateFromURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("restoreMachineStateFromURL:completionHandler:"), saveFileURL, completionHandler)
}

// Starts the VM and notifies the specified completion handler if startup was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start()
func (v_ VZVirtualMachine) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}

// A Boolean value that indicates whether you can start the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStart
func (v_ VZVirtualMachine) CanStart() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStart"))
	return rv
}

// The queue associated with this virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/queue
func (v_ VZVirtualMachine) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("queue"))
	return rv
}



