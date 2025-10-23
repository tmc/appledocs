// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	Queue() unsafe.Pointer
	CanPause() bool
	SetCanPause(value bool)
	CanRequestStop() bool
	SetCanRequestStop(value bool)
	CanResume() bool
	SetCanResume(value bool)
	CanStart() bool
	SetCanStart(value bool)
	CanStop() bool
	SetCanStop(value bool)
	ConsoleDevices() IVZConsoleDevice
	SetConsoleDevices(value IVZConsoleDevice)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DirectorySharingDevices() VZDirectorySharingDevice
	SetDirectorySharingDevices(value VZDirectorySharingDevice)
	GraphicsDevices() IVZGraphicsDevice
	SetGraphicsDevices(value IVZGraphicsDevice)
	MemoryBalloonDevices() VZMemoryBalloonDevice
	SetMemoryBalloonDevices(value VZMemoryBalloonDevice)
	NetworkDevices() VZNetworkDevice
	SetNetworkDevices(value VZNetworkDevice)
	SocketDevices() VZSocketDevice
	SetSocketDevices(value VZSocketDevice)
	State() coreml.State
	SetState(value coreml.State)
	UsbControllers() VZUSBController
	SetUsbControllers(value VZUSBController)
	RestoreMachineStateFromURLCompletionHandler(saveFileURL foundation.URL, completionHandler unsafe.Pointer)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that manages the overall state and configuration of your VM.
//
// A object emulates a complete hardware machine of the same architecture as the underlying Mac computer. Use the VM to execute a guest operating system and any other apps you install. The VM manages the resources that the guest operating system uses, providing access to some hardware resources while emulating others. Create and configure a object with details about how you want to configure your VM, and use that object to create the object. After creating the VM, call the method (Swift) or the method (Objective-C) to start the VM and boot the guest operating system.


// An object that manages the overall state and configuration of your VM.
//
// [Full Topic]
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



// Restores a VM from a previously saved state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/restoreMachineStateFrom(url:completionHandler:)
func (v_ VZVirtualMachine) RestoreMachineStateFromURLCompletionHandler(saveFileURL foundation.URL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("restoreMachineStateFromURL:completionHandler:"), saveFileURL, completionHandler)
}


// Starts the VM and notifies the specified completion handler if startup was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start()
func (v_ VZVirtualMachine) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}


// The queue associated with this virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/queue
func (v_ VZVirtualMachine) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("queue"))
	return rv
}


// A Boolean value that indicates whether you can pause the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canpause
func (v_ VZVirtualMachine) CanPause() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canPause"))
	return rv
}


// A Boolean value that indicates whether you can pause the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canpause
func (v_ VZVirtualMachine) SetCanPause(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanPause:"), value)
}


// A Boolean value that indicates whether you can ask the guest operating system to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canrequeststop
func (v_ VZVirtualMachine) CanRequestStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canRequestStop"))
	return rv
}


// A Boolean value that indicates whether you can ask the guest operating system to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canrequeststop
func (v_ VZVirtualMachine) SetCanRequestStop(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanRequestStop:"), value)
}


// A Boolean value that indicates whether you can resume the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canresume
func (v_ VZVirtualMachine) CanResume() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canResume"))
	return rv
}


// A Boolean value that indicates whether you can resume the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canresume
func (v_ VZVirtualMachine) SetCanResume(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanResume:"), value)
}


// A Boolean value that indicates whether you can start the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canstart
func (v_ VZVirtualMachine) CanStart() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStart"))
	return rv
}


// A Boolean value that indicates whether you can start the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canstart
func (v_ VZVirtualMachine) SetCanStart(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanStart:"), value)
}


// A Boolean value that indicates whether you can stop the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canstop
func (v_ VZVirtualMachine) CanStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStop"))
	return rv
}


// A Boolean value that indicates whether you can stop the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/canstop
func (v_ VZVirtualMachine) SetCanStop(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanStop:"), value)
}


// The list of configured console devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/consoledevices
func (v_ VZVirtualMachine) ConsoleDevices() IVZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("consoleDevices"))
	return rv
}


// The list of configured console devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/consoledevices
func (v_ VZVirtualMachine) SetConsoleDevices(value IVZConsoleDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}


// A custom object you use to determine when the VM stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/delegate
func (v_ VZVirtualMachine) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
}


// A custom object you use to determine when the VM stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/delegate
func (v_ VZVirtualMachine) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}


// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtualMachine) DirectorySharingDevices() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}


// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtualMachine) SetDirectorySharingDevices(value VZDirectorySharingDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectorySharingDevices:"), value)
}


// The list of configured graphics devices on the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/graphicsdevices
func (v_ VZVirtualMachine) GraphicsDevices() IVZGraphicsDevice {
	rv := objc.Send[VZGraphicsDevice](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}


// The list of configured graphics devices on the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/graphicsdevices
func (v_ VZVirtualMachine) SetGraphicsDevices(value IVZGraphicsDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setGraphicsDevices:"), value)
}


// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/memoryballoondevices
func (v_ VZVirtualMachine) MemoryBalloonDevices() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}


// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/memoryballoondevices
func (v_ VZVirtualMachine) SetMemoryBalloonDevices(value VZMemoryBalloonDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}


// The list of configured network devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/networkdevices
func (v_ VZVirtualMachine) NetworkDevices() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// The list of configured network devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/networkdevices
func (v_ VZVirtualMachine) SetNetworkDevices(value VZNetworkDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}


// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZVirtualMachine) SocketDevices() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("socketDevices"))
	return rv
}


// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZVirtualMachine) SetSocketDevices(value VZSocketDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}


// The current execution state of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/state-swift.property
func (v_ VZVirtualMachine) State() coreml.State {
	rv := objc.Send[coreml.State](v_.ID, objc.Sel("state"))
	return rv
}


// The current execution state of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/state-swift.property
func (v_ VZVirtualMachine) SetState(value coreml.State) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setState:"), value)
}


// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZVirtualMachine) UsbControllers() VZUSBController {
	rv := objc.Send[VZUSBController](v_.ID, objc.Sel("usbControllers"))
	return rv
}


// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZVirtualMachine) SetUsbControllers(value VZUSBController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
}



