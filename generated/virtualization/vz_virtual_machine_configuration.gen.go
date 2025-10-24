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
	// properties:
	AudioDevices() []IVZAudioDeviceConfiguration
	SetAudioDevices(value []IVZAudioDeviceConfiguration)
	BootLoader() IVZBootLoader
	SetBootLoader(value IVZBootLoader)
	ConsoleDevices() []IVZConsoleDeviceConfiguration
	SetConsoleDevices(value []IVZConsoleDeviceConfiguration)
	CPUCount() uint
	SetCPUCount(value uint)
	DirectorySharingDevices() []IVZDirectorySharingDeviceConfiguration
	SetDirectorySharingDevices(value []IVZDirectorySharingDeviceConfiguration)
	EntropyDevices() []IVZEntropyDeviceConfiguration
	SetEntropyDevices(value []IVZEntropyDeviceConfiguration)
	GraphicsDevices() []IVZGraphicsDeviceConfiguration
	SetGraphicsDevices(value []IVZGraphicsDeviceConfiguration)
	Keyboards() []IVZKeyboardConfiguration
	SetKeyboards(value []IVZKeyboardConfiguration)
	MemoryBalloonDevices() []IVZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value []IVZMemoryBalloonDeviceConfiguration)
	MemorySize() uint64
	SetMemorySize(value uint64)
	NetworkDevices() []IVZNetworkDeviceConfiguration
	SetNetworkDevices(value []IVZNetworkDeviceConfiguration)
	Platform() IVZPlatformConfiguration
	SetPlatform(value IVZPlatformConfiguration)
	PointingDevices() []IVZPointingDeviceConfiguration
	SetPointingDevices(value []IVZPointingDeviceConfiguration)
	SerialPorts() []IVZSerialPortConfiguration
	SetSerialPorts(value []IVZSerialPortConfiguration)
	SocketDevices() []IVZSocketDeviceConfiguration
	SetSocketDevices(value []IVZSocketDeviceConfiguration)
	StorageDevices() []IVZStorageDeviceConfiguration
	SetStorageDevices(value []IVZStorageDeviceConfiguration)
	UsbControllers() []IVZUSBControllerConfiguration
	SetUsbControllers(value []IVZUSBControllerConfiguration)
	CpuCount() uint
	SetCpuCount(value uint)
	// methods:
	ValidateWithError(error_ unsafe.Pointer) bool
	ValidateSaveRestoreSupportWithError(error_ unsafe.Pointer) bool
}

// The environment attributes and list of devices to use during the configuration of macOS or Linux VMs.
//
// Use a object to configure the environment for a macOS or Linux VM. This configuration object contains information about the VM environment, including the devices that the VM exposes to the guest operating system. For example, use the configuration object to specify the network interfaces and storage devices that the operating system may access. For more information on the devices that macOS and Linux guests can support, see the Devices section on the framework page. You create and configure objects directly. After validating the configuration object, use it to initialize the object that manages the virtual environment. The smallest valid configuration includes a value for the property; you can also include more devices in the configuration depending on the needs of your app, such as graphics devices, shared directories, and so on. When you finish configuring the object, call the method to determine whether a VM can successfully support your configuration. A configuration object is invalid if your app doesn’t have the entitlement. For more information on using , see and .


// The environment attributes and list of devices to use during the configuration of macOS or Linux VMs.
//
// [Full Topic]
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



// The maximum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedCPUCount
func (vc _VZVirtualMachineConfigurationClass) MaximumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("maximumAllowedCPUCount"))
	return rv
}

// The maximum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedMemorySize
func (vc _VZVirtualMachineConfigurationClass) MaximumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(vc.class), objc.Sel("maximumAllowedMemorySize"))
	return rv
}

// The minimum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedCPUCount
func (vc _VZVirtualMachineConfigurationClass) MinimumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("minimumAllowedCPUCount"))
	return rv
}

// The minimum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedMemorySize
func (vc _VZVirtualMachineConfigurationClass) MinimumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(vc.class), objc.Sel("minimumAllowedMemorySize"))
	return rv
}

// Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validate()
func (v_ VZVirtualMachineConfiguration) ValidateWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateWithError:"), error_)
	return rv
}


// Determines whether the framework can save or restore the VM’s current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validateSaveRestoreSupport()
func (v_ VZVirtualMachineConfiguration) ValidateSaveRestoreSupportWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateSaveRestoreSupportWithError:"), error_)
	return rv
}


// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) AudioDevices() []IVZAudioDeviceConfiguration {
	rv := objc.Send[[]VZAudioDeviceConfiguration](v_.ID, objc.Sel("audioDevices"))
	return rv
}


// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) SetAudioDevices(value []IVZAudioDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setAudioDevices:"), nsArray)
}


// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) BootLoader() IVZBootLoader {
	rv := objc.Send[VZBootLoader](v_.ID, objc.Sel("bootLoader"))
	return rv
}


// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) SetBootLoader(value IVZBootLoader) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBootLoader:"), value)
}


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) ConsoleDevices() []IVZConsoleDeviceConfiguration {
	rv := objc.Send[[]VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) SetConsoleDevices(value []IVZConsoleDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), nsArray)
}


// The number of CPUs you make available to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) CPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("CPUCount"))
	return rv
}


// The number of CPUs you make available to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) SetCPUCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCPUCount:"), value)
}


// The list of directory sharing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) DirectorySharingDevices() []IVZDirectorySharingDeviceConfiguration {
	rv := objc.Send[[]VZDirectorySharingDeviceConfiguration](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}


// The list of directory sharing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) SetDirectorySharingDevices(value []IVZDirectorySharingDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectorySharingDevices:"), nsArray)
}


// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) EntropyDevices() []IVZEntropyDeviceConfiguration {
	rv := objc.Send[[]VZEntropyDeviceConfiguration](v_.ID, objc.Sel("entropyDevices"))
	return rv
}


// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) SetEntropyDevices(value []IVZEntropyDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setEntropyDevices:"), nsArray)
}


// The list of graphics devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) GraphicsDevices() []IVZGraphicsDeviceConfiguration {
	rv := objc.Send[[]VZGraphicsDeviceConfiguration](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}


// The list of graphics devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) SetGraphicsDevices(value []IVZGraphicsDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setGraphicsDevices:"), nsArray)
}


// The list of keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) Keyboards() []IVZKeyboardConfiguration {
	rv := objc.Send[[]VZKeyboardConfiguration](v_.ID, objc.Sel("keyboards"))
	return rv
}


// The list of keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) SetKeyboards(value []IVZKeyboardConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setKeyboards:"), nsArray)
}


// The maximum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedCPUCount
func (v_ VZVirtualMachineConfiguration) MaximumAllowedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("maximumAllowedCPUCount"))
	return rv
}


// The maximum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedMemorySize
func (v_ VZVirtualMachineConfiguration) MaximumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("maximumAllowedMemorySize"))
	return rv
}


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) MemoryBalloonDevices() []IVZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[[]VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) SetMemoryBalloonDevices(value []IVZMemoryBalloonDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), nsArray)
}


// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) MemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("memorySize"))
	return rv
}


// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) SetMemorySize(value uint64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemorySize:"), value)
}


// The minimum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedCPUCount
func (v_ VZVirtualMachineConfiguration) MinimumAllowedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("minimumAllowedCPUCount"))
	return rv
}


// The minimum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedMemorySize
func (v_ VZVirtualMachineConfiguration) MinimumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("minimumAllowedMemorySize"))
	return rv
}


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) NetworkDevices() []IVZNetworkDeviceConfiguration {
	rv := objc.Send[[]VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) SetNetworkDevices(value []IVZNetworkDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), nsArray)
}


// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) Platform() IVZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](v_.ID, objc.Sel("platform"))
	return rv
}


// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) SetPlatform(value IVZPlatformConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlatform:"), value)
}


// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) PointingDevices() []IVZPointingDeviceConfiguration {
	rv := objc.Send[[]VZPointingDeviceConfiguration](v_.ID, objc.Sel("pointingDevices"))
	return rv
}


// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) SetPointingDevices(value []IVZPointingDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointingDevices:"), nsArray)
}


// The array of serial ports that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SerialPorts() []IVZSerialPortConfiguration {
	rv := objc.Send[[]VZSerialPortConfiguration](v_.ID, objc.Sel("serialPorts"))
	return rv
}


// The array of serial ports that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SetSerialPorts(value []IVZSerialPortConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setSerialPorts:"), nsArray)
}


// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SocketDevices() []IVZSocketDeviceConfiguration {
	rv := objc.Send[[]VZSocketDeviceConfiguration](v_.ID, objc.Sel("socketDevices"))
	return rv
}


// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SetSocketDevices(value []IVZSocketDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), nsArray)
}


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) StorageDevices() []IVZStorageDeviceConfiguration {
	rv := objc.Send[[]VZStorageDeviceConfiguration](v_.ID, objc.Sel("storageDevices"))
	return rv
}


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) SetStorageDevices(value []IVZStorageDeviceConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setStorageDevices:"), nsArray)
}


// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) UsbControllers() []IVZUSBControllerConfiguration {
	rv := objc.Send[[]VZUSBControllerConfiguration](v_.ID, objc.Sel("usbControllers"))
	return rv
}


// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) SetUsbControllers(value []IVZUSBControllerConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), nsArray)
}


// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v_ VZVirtualMachineConfiguration) CpuCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("cpuCount"))
	return rv
}


// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v_ VZVirtualMachineConfiguration) SetCpuCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCpuCount:"), value)
}



