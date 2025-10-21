// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	PauseWithCompletionHandler(completionHandler unsafe.Pointer)
	RequestStopWithError(error_ unsafe.Pointer) bool
	RestoreMachineStateFromURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer)
	ResumeWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveMachineStateToURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
	StartWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer)
	StopWithCompletionHandler(completionHandler unsafe.Pointer)
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




// Creates the VM and configures it with the specified data.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:)
func NewVZVirtualMachineWithConfiguration(configuration unsafe.Pointer) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[VZVirtualMachine](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}



// Creates and configures the VM with the specified data and dispatch queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/init(configuration:queue:)
func NewVZVirtualMachineWithConfigurationQueue(configuration unsafe.Pointer, queue unsafe.Pointer) VZVirtualMachine {
	instance := getVZVirtualMachineClass().Alloc()
	rv := objc.Send[VZVirtualMachine](instance.ID, objc.Sel("initWithConfiguration:queue:"), configuration, queue)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether the system supports virtualization.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isSupported
func (vc _VZVirtualMachineClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("supported"))
	return rv
}
// Pauses a running VM and notifies the specified completion handler of the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/pause()
func (v_ VZVirtualMachine) PauseWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("pauseWithCompletionHandler:"), completionHandler)
}

// Asks the guest operating system to stop running.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/requestStop()
func (v_ VZVirtualMachine) RequestStopWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("requestStopWithError:"), error_)
	return rv
}

// Restores a VM from a previously saved state.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/restoreMachineStateFrom(url:completionHandler:)
func (v_ VZVirtualMachine) RestoreMachineStateFromURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("restoreMachineStateFromURL:completionHandler:"), saveFileURL, completionHandler)
}

// Resumes a paused VM and notifies the specified completion handler of the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/resume()
func (v_ VZVirtualMachine) ResumeWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resumeWithCompletionHandler:"), completionHandler)
}

// Saves the state of a VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/saveMachineStateTo(url:completionHandler:)
func (v_ VZVirtualMachine) SaveMachineStateToURLCompletionHandler(saveFileURL unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("saveMachineStateToURL:completionHandler:"), saveFileURL, completionHandler)
}

// Starts the VM and notifies the specified completion handler if startup was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start()
func (v_ VZVirtualMachine) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}

// Starts the VM with the options and a completion handler you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/start(options:completionHandler:)
func (v_ VZVirtualMachine) StartWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("startWithOptions:completionHandler:"), options, completionHandler)
}

// Stops a VM that’s in either a running or paused state.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/stop(completionHandler:)
func (v_ VZVirtualMachine) StopWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("stopWithCompletionHandler:"), completionHandler)
}

// A Boolean value that indicates whether you can pause the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canPause
func (v_ VZVirtualMachine) CanPause() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canPause"))
	return rv
}

// A Boolean value that indicates whether you can ask the guest operating system to stop running.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canRequestStop
func (v_ VZVirtualMachine) CanRequestStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canRequestStop"))
	return rv
}

// A Boolean value that indicates whether you can resume the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canResume
func (v_ VZVirtualMachine) CanResume() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canResume"))
	return rv
}

// A Boolean value that indicates whether you can start the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStart
func (v_ VZVirtualMachine) CanStart() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStart"))
	return rv
}

// A Boolean value that indicates whether you can stop the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/canStop
func (v_ VZVirtualMachine) CanStop() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canStop"))
	return rv
}

// The list of configured console devices on the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/consoleDevices
func (v_ VZVirtualMachine) ConsoleDevices() []VZConsoleDevice {
	rv := objc.Send[[]VZConsoleDevice](v_.ID, objc.Sel("consoleDevices"))
	return rv
}

// A custom object you use to determine when the VM stops.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/delegate
func (v_ VZVirtualMachine) Delegate() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A custom object you use to determine when the VM stops.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/delegate
func (v_ VZVirtualMachine) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}

// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/directorySharingDevices
func (v_ VZVirtualMachine) DirectorySharingDevices() []VZDirectorySharingDevice {
	rv := objc.Send[[]VZDirectorySharingDevice](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}

// The list of configured graphics devices on the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/graphicsDevices
func (v_ VZVirtualMachine) GraphicsDevices() []VZGraphicsDevice {
	rv := objc.Send[[]VZGraphicsDevice](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}

// A Boolean value that indicates whether the system supports virtualization.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/isSupported
func (v_ VZVirtualMachine) Supported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("supported"))
	return rv
}

// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/memoryBalloonDevices
func (v_ VZVirtualMachine) MemoryBalloonDevices() []VZMemoryBalloonDevice {
	rv := objc.Send[[]VZMemoryBalloonDevice](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}

// The list of configured network devices on the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/networkDevices
func (v_ VZVirtualMachine) NetworkDevices() []VZNetworkDevice {
	rv := objc.Send[[]VZNetworkDevice](v_.ID, objc.Sel("networkDevices"))
	return rv
}

// The queue associated with this virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/queue
func (v_ VZVirtualMachine) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("queue"))
	return rv
}

// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/socketDevices
func (v_ VZVirtualMachine) SocketDevices() []VZSocketDevice {
	rv := objc.Send[[]VZSocketDevice](v_.ID, objc.Sel("socketDevices"))
	return rv
}

// The current execution state of the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/state-swift.property
func (v_ VZVirtualMachine) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("state"))
	return rv
}

// The list of runtime USB controller objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/usbControllers
func (v_ VZVirtualMachine) UsbControllers() []VZUSBController {
	rv := objc.Send[[]VZUSBController](v_.ID, objc.Sel("usbControllers"))
	return rv
}


