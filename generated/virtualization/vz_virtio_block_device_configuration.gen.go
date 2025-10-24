// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZVirtioBlockDeviceConfiguration] class.
var (
	VZVirtioBlockDeviceConfigurationClass     _VZVirtioBlockDeviceConfigurationClass
	VZVirtioBlockDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioBlockDeviceConfigurationClass() _VZVirtioBlockDeviceConfigurationClass {
	VZVirtioBlockDeviceConfigurationClassOnce.Do(func() {
		VZVirtioBlockDeviceConfigurationClass = _VZVirtioBlockDeviceConfigurationClass{objc.GetClass("VZVirtioBlockDeviceConfiguration")}
	})
	return VZVirtioBlockDeviceConfigurationClass
}

type _VZVirtioBlockDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioBlockDeviceConfiguration] class.
type IVZVirtioBlockDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
	// properties:
	BlockDeviceIdentifier() objc.IObject /* cross-framework: NSString */
	SetBlockDeviceIdentifier(value objc.IObject /* cross-framework: NSString */)
	StorageDevices() IVZStorageDeviceConfiguration
	SetStorageDevices(value IVZStorageDeviceConfiguration)
	// methods:
}

// The configuration object that requests the creation of a virtual storage device in the guest system.
//
// Use a object to create an emulated storage device in your virtual machine. When you add this object to your virtual machine configuration, the virtual machine creates an emulated disk for the guest operating system to use to read and write files. The emulated storage device conforms to the Virtio Block Device specification. When you create a object, specify the attachment object that implements the underlying storage. For example, specify a object to configure the storage device using a disk image in the local file system. Assign your configuration object to the property of your object before creating your virtual machine.


// The configuration object that requests the creation of a virtual storage device in the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration
type VZVirtioBlockDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZVirtioBlockDeviceConfigurationFrom constructs a [VZVirtioBlockDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that requests the creation of a virtual storage device in the guest system.
func VZVirtioBlockDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioBlockDeviceConfiguration {
	return VZVirtioBlockDeviceConfiguration{
		VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioBlockDeviceConfigurationClass) Alloc() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioBlockDeviceConfigurationClass) New() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioBlockDeviceConfiguration) Init() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioBlockDeviceConfiguration) Autorelease() VZVirtioBlockDeviceConfiguration {
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioBlockDeviceConfiguration creates a new VZVirtioBlockDeviceConfiguration instance.
func NewVZVirtioBlockDeviceConfiguration() VZVirtioBlockDeviceConfiguration {
	return getVZVirtioBlockDeviceConfigurationClass().New()
}



// Creates a block device configuration object that uses the specified storage medium.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/init(attachment:)
func NewVZVirtioBlockDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZVirtioBlockDeviceConfiguration {
	instance := getVZVirtioBlockDeviceConfigurationClass().Alloc()
	rv := objc.Send[VZVirtioBlockDeviceConfiguration](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	rv.Autorelease()
	return rv
}



// Checks the validity of a block device identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/validateBlockDeviceIdentifier(_:)
func (vc _VZVirtioBlockDeviceConfigurationClass) ValidateBlockDeviceIdentifierError(blockDeviceIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateBlockDeviceIdentifier:error:"), blockDeviceIdentifier, error_)
	return rv
}


// The string that identifies the VIRTIO block device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/blockDeviceIdentifier
func (v_ VZVirtioBlockDeviceConfiguration) BlockDeviceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("blockDeviceIdentifier"))
	return rv
}


// The string that identifies the VIRTIO block device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/blockDeviceIdentifier
func (v_ VZVirtioBlockDeviceConfiguration) SetBlockDeviceIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBlockDeviceIdentifier:"), value)
}


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/storagedevices
func (v_ VZVirtioBlockDeviceConfiguration) StorageDevices() IVZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](v_.ID, objc.Sel("storageDevices"))
	return rv
}


// The array of storage devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/storagedevices
func (v_ VZVirtioBlockDeviceConfiguration) SetStorageDevices(value IVZStorageDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStorageDevices:"), value)
}


