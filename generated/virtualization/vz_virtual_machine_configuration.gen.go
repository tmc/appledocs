// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtualMachineConfiguration */


/* debug [class_header]: Header for VZVirtualMachineConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtualMachineConfiguration */
// An interface definition for the [VZVirtualMachineConfiguration] class.
type IVZVirtualMachineConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZVirtualMachineConfiguration */
	// properties:
	AudioDevices() []VZAudioDeviceConfiguration
	SetAudioDevices(value []VZAudioDeviceConfiguration)
	BootLoader() IVZBootLoader
	SetBootLoader(value IVZBootLoader)
	ConsoleDevices() []VZConsoleDeviceConfiguration
	SetConsoleDevices(value []VZConsoleDeviceConfiguration)
	CPUCount() uint
	SetCPUCount(value uint)
	DirectorySharingDevices() []VZDirectorySharingDeviceConfiguration
	SetDirectorySharingDevices(value []VZDirectorySharingDeviceConfiguration)
	EntropyDevices() []VZEntropyDeviceConfiguration
	SetEntropyDevices(value []VZEntropyDeviceConfiguration)
	GraphicsDevices() []VZGraphicsDeviceConfiguration
	SetGraphicsDevices(value []VZGraphicsDeviceConfiguration)
	Keyboards() []VZKeyboardConfiguration
	SetKeyboards(value []VZKeyboardConfiguration)
	MemoryBalloonDevices() []VZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value []VZMemoryBalloonDeviceConfiguration)
	MemorySize() uint64
	SetMemorySize(value uint64)
	NetworkDevices() []VZNetworkDeviceConfiguration
	SetNetworkDevices(value []VZNetworkDeviceConfiguration)
	Platform() IVZPlatformConfiguration
	SetPlatform(value IVZPlatformConfiguration)
	PointingDevices() []VZPointingDeviceConfiguration
	SetPointingDevices(value []VZPointingDeviceConfiguration)
	SerialPorts() []VZSerialPortConfiguration
	SetSerialPorts(value []VZSerialPortConfiguration)
	SocketDevices() []VZSocketDeviceConfiguration
	SetSocketDevices(value []VZSocketDeviceConfiguration)
	StorageDevices() []VZStorageDeviceConfiguration
	SetStorageDevices(value []VZStorageDeviceConfiguration)
	UsbControllers() []VZUSBControllerConfiguration
	SetUsbControllers(value []VZUSBControllerConfiguration)
	CpuCount() uint
	SetCpuCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtualMachineConfiguration */
	// methods:
	ValidateWithError(error_ objectivec.IObject) bool
	ValidateSaveRestoreSupportWithError(error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtualMachineConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineConfigurationClass) Alloc() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtualMachineConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtualMachineConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtualMachineConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtualMachineConfiguration */

// The maximum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedCPUCount
func (vc _VZVirtualMachineConfigurationClass) MaximumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("maximumAllowedCPUCount"))
	return rv
}/* debug [class_properties_class/property]: maximumAllowedCPUCount */

// The maximum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedMemorySize
func (vc _VZVirtualMachineConfigurationClass) MaximumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(vc.class), objc.Sel("maximumAllowedMemorySize"))
	return rv
}/* debug [class_properties_class/property]: maximumAllowedMemorySize */

// The minimum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedCPUCount
func (vc _VZVirtualMachineConfigurationClass) MinimumAllowedCPUCount() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("minimumAllowedCPUCount"))
	return rv
}/* debug [class_properties_class/property]: minimumAllowedCPUCount */

// The minimum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedMemorySize
func (vc _VZVirtualMachineConfigurationClass) MinimumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(vc.class), objc.Sel("minimumAllowedMemorySize"))
	return rv
}/* debug [class_properties_class/property]: minimumAllowedMemorySize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtualMachineConfiguration */

// Validates the current configuration settings and reports any issues that might prevent the successful initialization of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validate()
func (v_ VZVirtualMachineConfiguration) ValidateWithError(error_ objectivec.IObject) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ValidateWithError */


// Determines whether the framework can save or restore the VM’s current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/validateSaveRestoreSupport()
func (v_ VZVirtualMachineConfiguration) ValidateSaveRestoreSupportWithError(error_ objectivec.IObject) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("validateSaveRestoreSupportWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ValidateSaveRestoreSupportWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtualMachineConfiguration */

// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) AudioDevices() []VZAudioDeviceConfiguration {
	rv := objc.Send[[]VZAudioDeviceConfiguration](v_.ID, objc.Sel("audioDevices"))
	return rv
}/* debug [instance_properties/getter]: audioDevices */


// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/audioDevices
func (v_ VZVirtualMachineConfiguration) SetAudioDevices(value []VZAudioDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: audioDevices */


// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) BootLoader() IVZBootLoader {
	rv := objc.Send[VZBootLoader](v_.ID, objc.Sel("bootLoader"))
	return rv
}/* debug [instance_properties/getter]: bootLoader */


// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/bootLoader
func (v_ VZVirtualMachineConfiguration) SetBootLoader(value IVZBootLoader) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBootLoader:"), value)
}/* debug [instance_properties/setter]: bootLoader */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) ConsoleDevices() []VZConsoleDeviceConfiguration {
	rv := objc.Send[[]VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}/* debug [instance_properties/getter]: consoleDevices */


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/consoleDevices
func (v_ VZVirtualMachineConfiguration) SetConsoleDevices(value []VZConsoleDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: consoleDevices */


// The number of CPUs you make available to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) CPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("CPUCount"))
	return rv
}/* debug [instance_properties/getter]: CPUCount */


// The number of CPUs you make available to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/cpuCount
func (v_ VZVirtualMachineConfiguration) SetCPUCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCPUCount:"), value)
}/* debug [instance_properties/setter]: CPUCount */


// The list of directory sharing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) DirectorySharingDevices() []VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[[]VZDirectorySharingDeviceConfiguration](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}/* debug [instance_properties/getter]: directorySharingDevices */


// The list of directory sharing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/directorySharingDevices
func (v_ VZVirtualMachineConfiguration) SetDirectorySharingDevices(value []VZDirectorySharingDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: directorySharingDevices */


// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) EntropyDevices() []VZEntropyDeviceConfiguration {
	rv := objc.Send[[]VZEntropyDeviceConfiguration](v_.ID, objc.Sel("entropyDevices"))
	return rv
}/* debug [instance_properties/getter]: entropyDevices */


// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/entropyDevices
func (v_ VZVirtualMachineConfiguration) SetEntropyDevices(value []VZEntropyDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: entropyDevices */


// The list of graphics devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) GraphicsDevices() []VZGraphicsDeviceConfiguration {
	rv := objc.Send[[]VZGraphicsDeviceConfiguration](v_.ID, objc.Sel("graphicsDevices"))
	return rv
}/* debug [instance_properties/getter]: graphicsDevices */


// The list of graphics devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/graphicsDevices
func (v_ VZVirtualMachineConfiguration) SetGraphicsDevices(value []VZGraphicsDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: graphicsDevices */


// The list of keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) Keyboards() []VZKeyboardConfiguration {
	rv := objc.Send[[]VZKeyboardConfiguration](v_.ID, objc.Sel("keyboards"))
	return rv
}/* debug [instance_properties/getter]: keyboards */


// The list of keyboards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/keyboards
func (v_ VZVirtualMachineConfiguration) SetKeyboards(value []VZKeyboardConfiguration) {
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
}/* debug [instance_properties/setter]: keyboards */


// The maximum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedCPUCount
func (v_ VZVirtualMachineConfiguration) MaximumAllowedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("maximumAllowedCPUCount"))
	return rv
}/* debug [instance_properties/getter]: maximumAllowedCPUCount */


// The maximum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/maximumAllowedMemorySize
func (v_ VZVirtualMachineConfiguration) MaximumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("maximumAllowedMemorySize"))
	return rv
}/* debug [instance_properties/getter]: maximumAllowedMemorySize */


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) MemoryBalloonDevices() []VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[[]VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}/* debug [instance_properties/getter]: memoryBalloonDevices */


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memoryBalloonDevices
func (v_ VZVirtualMachineConfiguration) SetMemoryBalloonDevices(value []VZMemoryBalloonDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: memoryBalloonDevices */


// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) MemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("memorySize"))
	return rv
}/* debug [instance_properties/getter]: memorySize */


// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/memorySize
func (v_ VZVirtualMachineConfiguration) SetMemorySize(value uint64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemorySize:"), value)
}/* debug [instance_properties/setter]: memorySize */


// The minimum number of CPUs you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedCPUCount
func (v_ VZVirtualMachineConfiguration) MinimumAllowedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("minimumAllowedCPUCount"))
	return rv
}/* debug [instance_properties/getter]: minimumAllowedCPUCount */


// The minimum amount of memory that you may configure for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/minimumAllowedMemorySize
func (v_ VZVirtualMachineConfiguration) MinimumAllowedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("minimumAllowedMemorySize"))
	return rv
}/* debug [instance_properties/getter]: minimumAllowedMemorySize */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) NetworkDevices() []VZNetworkDeviceConfiguration {
	rv := objc.Send[[]VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}/* debug [instance_properties/getter]: networkDevices */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/networkDevices
func (v_ VZVirtualMachineConfiguration) SetNetworkDevices(value []VZNetworkDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: networkDevices */


// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) Platform() IVZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](v_.ID, objc.Sel("platform"))
	return rv
}/* debug [instance_properties/getter]: platform */


// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/platform
func (v_ VZVirtualMachineConfiguration) SetPlatform(value IVZPlatformConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlatform:"), value)
}/* debug [instance_properties/setter]: platform */


// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) PointingDevices() []VZPointingDeviceConfiguration {
	rv := objc.Send[[]VZPointingDeviceConfiguration](v_.ID, objc.Sel("pointingDevices"))
	return rv
}/* debug [instance_properties/getter]: pointingDevices */


// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/pointingDevices
func (v_ VZVirtualMachineConfiguration) SetPointingDevices(value []VZPointingDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: pointingDevices */


// The array of serial ports that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SerialPorts() []VZSerialPortConfiguration {
	rv := objc.Send[[]VZSerialPortConfiguration](v_.ID, objc.Sel("serialPorts"))
	return rv
}/* debug [instance_properties/getter]: serialPorts */


// The array of serial ports that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/serialPorts
func (v_ VZVirtualMachineConfiguration) SetSerialPorts(value []VZSerialPortConfiguration) {
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
}/* debug [instance_properties/setter]: serialPorts */


// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SocketDevices() []VZSocketDeviceConfiguration {
	rv := objc.Send[[]VZSocketDeviceConfiguration](v_.ID, objc.Sel("socketDevices"))
	return rv
}/* debug [instance_properties/getter]: socketDevices */


// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/socketDevices
func (v_ VZVirtualMachineConfiguration) SetSocketDevices(value []VZSocketDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: socketDevices */


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) StorageDevices() []VZStorageDeviceConfiguration {
	rv := objc.Send[[]VZStorageDeviceConfiguration](v_.ID, objc.Sel("storageDevices"))
	return rv
}/* debug [instance_properties/getter]: storageDevices */


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/storageDevices
func (v_ VZVirtualMachineConfiguration) SetStorageDevices(value []VZStorageDeviceConfiguration) {
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
}/* debug [instance_properties/setter]: storageDevices */


// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) UsbControllers() []VZUSBControllerConfiguration {
	rv := objc.Send[[]VZUSBControllerConfiguration](v_.ID, objc.Sel("usbControllers"))
	return rv
}/* debug [instance_properties/getter]: usbControllers */


// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/usbControllers
func (v_ VZVirtualMachineConfiguration) SetUsbControllers(value []VZUSBControllerConfiguration) {
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
}/* debug [instance_properties/setter]: usbControllers */


// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v_ VZVirtualMachineConfiguration) CpuCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("cpuCount"))
	return rv
}/* debug [instance_properties/getter]: cpuCount */


// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v_ VZVirtualMachineConfiguration) SetCpuCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCpuCount:"), value)
}/* debug [instance_properties/setter]: cpuCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtualMachineConfiguration */



