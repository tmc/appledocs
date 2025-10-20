// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZVirtualMachineConfiguration] class.
var (
	VZVirtualMachineConfigurationClass     _VZVirtualMachineConfigurationClass
	VZVirtualMachineConfigurationClassOnce sync.Once
)

func getVZVirtualMachineConfigurationClass() _VZVirtualMachineConfigurationClass {
	VZVirtualMachineConfigurationClassOnce.Do(func() {
		VZVirtualMachineConfigurationClass = _VZVirtualMachineConfigurationClass{objc.GetClass("VZVirtualMachineConfiguration")}
	})
	return VZVirtualMachineConfigurationClass
}

type _VZVirtualMachineConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtualMachineConfiguration] class.
type IVZVirtualMachineConfiguration interface {
	objectivec.IObject
	ValidateWithError(error_ unsafe.Pointer) bool
	ValidateSaveRestoreSupportWithError(error_ unsafe.Pointer) bool
}

// The environment attributes and list of devices to use during the configuration of macOS or Linux VMs.
//
// Use a object to configure the environment for a macOS or Linux VM. This configuration object contains information about the VM environment, including the devices that the VM exposes to the guest operating system. For example, use the configuration object to specify the network interfaces and storage devices that the operating system may access. For more information on the devices that macOS and Linux guests can support, see the Devices section on the framework page. You create and configure objects directly. After validating the configuration object, use it to initialize the object that manages the virtual environment. The smallest valid configuration includes a value for the property; you can also include more devices in the configuration depending on the needs of your app, such as graphics devices, shared directories, and so on. When you finish configuring the object, call the method to determine whether a VM can successfully support your configuration. A configuration object is invalid if your app doesn’t have the entitlement. For more information on using , see and .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration
type VZVirtualMachineConfiguration struct {
	objectivec.Object
}

// VZVirtualMachineConfigurationFrom constructs a [VZVirtualMachineConfiguration] from an unsafe.Pointer.
//
// The environment attributes and list of devices to use during the configuration of macOS or Linux VMs.
func VZVirtualMachineConfigurationFrom(ptr unsafe.Pointer) VZVirtualMachineConfiguration {
	return VZVirtualMachineConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineConfigurationClass) Alloc() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtualMachineConfigurationClass) New() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtualMachineConfiguration) Init() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtualMachineConfiguration) Autorelease() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfiguration creates a new VZVirtualMachineConfiguration instance.
func NewVZVirtualMachineConfiguration() VZVirtualMachineConfiguration {
	return getVZVirtualMachineConfigurationClass().New()
}


// Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validate()
func (v_ VZVirtualMachineConfiguration) ValidateWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateWithError:"), error_)
	return rv
}

// Determines whether the framework can save or restore the VM’s current configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validateSaveRestoreSupport()
func (v_ VZVirtualMachineConfiguration) ValidateSaveRestoreSupportWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateSaveRestoreSupportWithError:"), error_)
	return rv
}

// The list of audio devices.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) AudioDevices() []VZAudioDeviceConfiguration {
	rv := objc.Send[[]VZAudioDeviceConfiguration](v_.ID, objc.Sel("audioDevices"))
	return rv
}


// SetAudioDevices sets the value of the audioDevices property.
// The list of audio devices.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) SetAudioDevices(value []VZAudioDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAudioDevices:"), value)
}
// The guest system to boot when the VM starts.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) BootLoader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("bootLoader"))
	return rv
}


// SetBootLoader sets the value of the bootLoader property.
// The guest system to boot when the VM starts.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) SetBootLoader(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBootLoader:"), value)
}
// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) ConsoleDevices() []VZConsoleDeviceConfiguration {
	rv := objc.Send[[]VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}


// SetConsoleDevices sets the value of the consoleDevices property.
// The array of console devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) SetConsoleDevices(value []VZConsoleDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}
// The number of CPUs you make available to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) CPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("CPUCount"))
	return rv
}


// SetCPUCount sets the value of the CPUCount property.
// The number of CPUs you make available to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) SetCPUCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCPUCount:"), value)
}
// The list of directory sharing devices.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) DirectorySharingDevices() []VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[[]VZDirectorySharingDeviceConfiguration](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}


// SetDirectorySharingDevices sets the value of the directorySharingDevices property.
// The list of directory sharing devices.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) SetDirectorySharingDevices(value []VZDirectorySharingDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectorySharingDevices:"), value)
}
// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) EntropyDevices() []VZEntropyDeviceConfiguration {
	rv := objc.Send[[]VZEntropyDeviceConfiguration](v_.ID, objc.Sel("entropyDevices"))
	return rv
}


// SetEntropyDevices sets the value of the entropyDevices property.
// The array of randomization devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) SetEntropyDevices(value []VZEntropyDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEntropyDevices:"), value)
}
// The list of graphics devices.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) GraphicsDevices() []VZGraphicsDeviceConfiguration {
	rv := objc.Send[[]VZGraphicsDeviceConfiguration](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}


// SetGraphicsDevices sets the value of the graphicsDevices property.
// The list of graphics devices.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) SetGraphicsDevices(value []VZGraphicsDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setGraphicsDevices:"), value)
}
// The list of keyboards.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) Keyboards() []VZKeyboardConfiguration {
	rv := objc.Send[[]VZKeyboardConfiguration](v_.ID, objc.Sel("keyboards"))
	return rv
}


// SetKeyboards sets the value of the keyboards property.
// The list of keyboards.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) SetKeyboards(value []VZKeyboardConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKeyboards:"), value)
}
// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) MemoryBalloonDevices() []VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[[]VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}


// SetMemoryBalloonDevices sets the value of the memoryBalloonDevices property.
// An array that you configure with a memory balloon device, used to update the memory in the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) SetMemoryBalloonDevices(value []VZMemoryBalloonDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}
// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) MemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("memorySize"))
	return rv
}


// SetMemorySize sets the value of the memorySize property.
// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) SetMemorySize(value uint64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemorySize:"), value)
}
// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) NetworkDevices() []VZNetworkDeviceConfiguration {
	rv := objc.Send[[]VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// SetNetworkDevices sets the value of the networkDevices property.
// The array of network devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) SetNetworkDevices(value []VZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}
// The hardware platform to use.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) Platform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("platform"))
	return rv
}


// SetPlatform sets the value of the platform property.
// The hardware platform to use.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) SetPlatform(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlatform:"), value)
}
// The list of pointing devices.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) PointingDevices() []VZPointingDeviceConfiguration {
	rv := objc.Send[[]VZPointingDeviceConfiguration](v_.ID, objc.Sel("pointingDevices"))
	return rv
}


// SetPointingDevices sets the value of the pointingDevices property.
// The list of pointing devices.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) SetPointingDevices(value []VZPointingDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointingDevices:"), value)
}
// The array of serial ports that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SerialPorts() []VZSerialPortConfiguration {
	rv := objc.Send[[]VZSerialPortConfiguration](v_.ID, objc.Sel("serialPorts"))
	return rv
}


// SetSerialPorts sets the value of the serialPorts property.
// The array of serial ports that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SetSerialPorts(value []VZSerialPortConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSerialPorts:"), value)
}
// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SocketDevices() []VZSocketDeviceConfiguration {
	rv := objc.Send[[]VZSocketDeviceConfiguration](v_.ID, objc.Sel("socketDevices"))
	return rv
}


// SetSocketDevices sets the value of the socketDevices property.
// The socket device that you use to implement port-based communication with the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SetSocketDevices(value []VZSocketDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}
// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) StorageDevices() []VZStorageDeviceConfiguration {
	rv := objc.Send[[]VZStorageDeviceConfiguration](v_.ID, objc.Sel("storageDevices"))
	return rv
}


// SetStorageDevices sets the value of the storageDevices property.
// The array of storage devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) SetStorageDevices(value []VZStorageDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStorageDevices:"), value)
}
// The list of configured USB controllers for the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) UsbControllers() []VZUSBControllerConfiguration {
	rv := objc.Send[[]VZUSBControllerConfiguration](v_.ID, objc.Sel("usbControllers"))
	return rv
}


// SetUsbControllers sets the value of the usbControllers property.
// The list of configured USB controllers for the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) SetUsbControllers(value []VZUSBControllerConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
}
// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v_ VZVirtualMachineConfiguration) CpuCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("cpuCount"))
	return rv
}


// SetCpuCount sets the value of the cpuCount property.
// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount.
func (v_ VZVirtualMachineConfiguration) SetCpuCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCpuCount:"), value)
}


