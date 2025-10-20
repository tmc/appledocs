// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	objectivec.IObject
}

// The configuration object that requests the creation of a virtual storage device in the guest system.
//
// Use a object to create an emulated storage device in your virtual machine. When you add this object to your virtual machine configuration, the virtual machine creates an emulated disk for the guest operating system to use to read and write files. The emulated storage device conforms to the Virtio Block Device specification. When you create a object, specify the attachment object that implements the underlying storage. For example, specify a object to configure the storage device using a disk image in the local file system. Assign your configuration object to the property of your object before creating your virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration
type VZVirtioBlockDeviceConfiguration struct {
	objectivec.Object
}

// VZVirtioBlockDeviceConfigurationFrom constructs a [VZVirtioBlockDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that requests the creation of a virtual storage device in the guest system.
func VZVirtioBlockDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioBlockDeviceConfiguration {
	return VZVirtioBlockDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
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

// The storage device attachment for this block device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/attachment
func (v_ VZVirtioBlockDeviceConfiguration) Attachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("attachment"))
	return rv
}

// SetAttachment sets the value of the attachment property.
// The storage device attachment for this block device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/attachment
func (v_ VZVirtioBlockDeviceConfiguration) SetAttachment(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}

// Checks the validity of a block device identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioBlockDeviceConfiguration/validateBlockDeviceIdentifier(_:)
func (vc _VZVirtioBlockDeviceConfigurationClass) ValidateBlockDeviceIdentifierError(blockDeviceIdentifier string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateBlockDeviceIdentifier:error:"), objc.String(blockDeviceIdentifier), error_)
	return rv
}



