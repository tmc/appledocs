// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtualMachine */

/* debug [class_header]: Header for VZVirtualMachine */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtualMachine */
// An interface definition for the [VZVirtualMachine] class.
type IVZVirtualMachine interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZVirtualMachine */
	// properties:
	CanPause() bool
	CanRequestStop() bool
	CanResume() bool
	CanStart() bool
	CanStop() bool
	ConsoleDevices() []VZConsoleDevice
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DirectorySharingDevices() []VZDirectorySharingDevice
	GraphicsDevices() []VZGraphicsDevice
	MemoryBalloonDevices() []VZMemoryBalloonDevice
	NetworkDevices() []VZNetworkDevice
	Queue() unsafe.Pointer
	SocketDevices() []VZSocketDevice
	State() VZVirtualMachineState
	UsbControllers() []VZUSBController
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtualMachine */
	// methods:
	PauseWithCompletionHandler(completionHandler unsafe.Pointer)
	RequestStopWithError(error_ unsafe.Pointer) bool
	RestoreMachineStateFromURLCompletionHandler(saveFileURL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	ResumeWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveMachineStateToURLCompletionHandler(saveFileURL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
	StartWithOptionsCompletionHandler(options IVZVirtualMachineStartOptions, completionHandler unsafe.Pointer)
	StopWithCompletionHandler(completionHandler unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtualMachine */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineClass) Alloc() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtualMachine */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtualMachine */

// Creates the VM and configures it with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:)
func NewVZVirtualMachineWithConfiguration(configuration IVZVirtualMachineConfiguration) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[VZVirtualMachine](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZVirtualMachineWithConfiguration */

// Creates and configures the VM with the specified data and dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:queue:)
func NewVZVirtualMachineWithConfigurationQueue(configuration IVZVirtualMachineConfiguration, queue unsafe.Pointer) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[VZVirtualMachine](instance.ID, objc.Sel("initWithConfiguration:queue:"), configuration, queue)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZVirtualMachineWithConfigurationQueue */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtualMachine */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtualMachine */

// A Boolean value that indicates whether the system supports virtualization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isSupported
func (vc _VZVirtualMachineClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("supported"))
	return rv
} /* debug [class_properties_class/property]: supported */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtualMachine */

// Pauses a running VM and notifies the specified completion handler of the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/pause()
func (v_ VZVirtualMachine) PauseWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("pauseWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: PauseWithCompletionHandler */

// Asks the guest operating system to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/requestStop()
func (v_ VZVirtualMachine) RequestStopWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("requestStopWithError:"), error_)
	return rv
} /* debug [instance_methods/method]: RequestStopWithError */

// Restores a VM from a previously saved state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/restoreMachineStateFrom(url:completionHandler:)
func (v_ VZVirtualMachine) RestoreMachineStateFromURLCompletionHandler(saveFileURL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("restoreMachineStateFromURL:completionHandler:"), saveFileURL, completionHandler)
} /* debug [instance_methods/method]: RestoreMachineStateFromURLCompletionHandler */

// Resumes a paused VM and notifies the specified completion handler of the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/resume()
func (v_ VZVirtualMachine) ResumeWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resumeWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: ResumeWithCompletionHandler */

// Saves the state of a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/saveMachineStateTo(url:completionHandler:)
func (v_ VZVirtualMachine) SaveMachineStateToURLCompletionHandler(saveFileURL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("saveMachineStateToURL:completionHandler:"), saveFileURL, completionHandler)
} /* debug [instance_methods/method]: SaveMachineStateToURLCompletionHandler */

// Starts the VM and notifies the specified completion handler if startup was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start()
func (v_ VZVirtualMachine) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: StartWithCompletionHandler */

// Starts the VM with the options and a completion handler you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start(options:completionHandler:)
func (v_ VZVirtualMachine) StartWithOptionsCompletionHandler(options IVZVirtualMachineStartOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithOptions:completionHandler:"), options, completionHandler)
} /* debug [instance_methods/method]: StartWithOptionsCompletionHandler */

// Stops a VM that’s in either a running or paused state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/stop(completionHandler:)
func (v_ VZVirtualMachine) StopWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("stopWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: StopWithCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtualMachine */

// A Boolean value that indicates whether you can pause the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canPause
func (v_ VZVirtualMachine) CanPause() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canPause"))
	return rv
} /* debug [instance_properties/getter]: canPause */

// A Boolean value that indicates whether you can ask the guest operating system to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canRequestStop
func (v_ VZVirtualMachine) CanRequestStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canRequestStop"))
	return rv
} /* debug [instance_properties/getter]: canRequestStop */

// A Boolean value that indicates whether you can resume the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canResume
func (v_ VZVirtualMachine) CanResume() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canResume"))
	return rv
} /* debug [instance_properties/getter]: canResume */

// A Boolean value that indicates whether you can start the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStart
func (v_ VZVirtualMachine) CanStart() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStart"))
	return rv
} /* debug [instance_properties/getter]: canStart */

// A Boolean value that indicates whether you can stop the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStop
func (v_ VZVirtualMachine) CanStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStop"))
	return rv
} /* debug [instance_properties/getter]: canStop */

// The list of configured console devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/consoleDevices
func (v_ VZVirtualMachine) ConsoleDevices() []VZConsoleDevice {
	rv := objc.Send[[]VZConsoleDevice](v_.ID, objc.Sel("consoleDevices"))
	return rv
} /* debug [instance_properties/getter]: consoleDevices */

// A custom object you use to determine when the VM stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/delegate
func (v_ VZVirtualMachine) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
} /* debug [instance_properties/getter]: delegate */

// A custom object you use to determine when the VM stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/delegate
func (v_ VZVirtualMachine) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
} /* debug [instance_properties/setter]: delegate */

// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/directorySharingDevices
func (v_ VZVirtualMachine) DirectorySharingDevices() []VZDirectorySharingDevice {
	rv := objc.Send[[]VZDirectorySharingDevice](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
} /* debug [instance_properties/getter]: directorySharingDevices */

// The list of configured graphics devices on the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/graphicsDevices
func (v_ VZVirtualMachine) GraphicsDevices() []VZGraphicsDevice {
	rv := objc.Send[[]VZGraphicsDevice](v_.ID, objc.Sel("graphicsDevices"))
	return rv
} /* debug [instance_properties/getter]: graphicsDevices */

// A Boolean value that indicates whether the system supports virtualization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isSupported
func (v_ VZVirtualMachine) Supported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("supported"))
	return rv
} /* debug [instance_properties/getter]: supported */

// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/memoryBalloonDevices
func (v_ VZVirtualMachine) MemoryBalloonDevices() []VZMemoryBalloonDevice {
	rv := objc.Send[[]VZMemoryBalloonDevice](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
} /* debug [instance_properties/getter]: memoryBalloonDevices */

// The list of configured network devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/networkDevices
func (v_ VZVirtualMachine) NetworkDevices() []VZNetworkDevice {
	rv := objc.Send[[]VZNetworkDevice](v_.ID, objc.Sel("networkDevices"))
	return rv
} /* debug [instance_properties/getter]: networkDevices */

// The queue associated with this virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/queue
func (v_ VZVirtualMachine) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("queue"))
	return rv
} /* debug [instance_properties/getter]: queue */

// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/socketDevices
func (v_ VZVirtualMachine) SocketDevices() []VZSocketDevice {
	rv := objc.Send[[]VZSocketDevice](v_.ID, objc.Sel("socketDevices"))
	return rv
} /* debug [instance_properties/getter]: socketDevices */

// The current execution state of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/state-swift.property
func (v_ VZVirtualMachine) State() VZVirtualMachineState {
	rv := objc.Send[VZVirtualMachineState](v_.ID, objc.Sel("state"))
	return rv
} /* debug [instance_properties/getter]: state */

// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/usbControllers
func (v_ VZVirtualMachine) UsbControllers() []VZUSBController {
	rv := objc.Send[[]VZUSBController](v_.ID, objc.Sel("usbControllers"))
	return rv
} /* debug [instance_properties/getter]: usbControllers */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtualMachine */
