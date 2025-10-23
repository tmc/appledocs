// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioFileSystemDeviceConfiguration] class.
var (
	VZVirtioFileSystemDeviceConfigurationClass     _VZVirtioFileSystemDeviceConfigurationClass
	VZVirtioFileSystemDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioFileSystemDeviceConfigurationClass() _VZVirtioFileSystemDeviceConfigurationClass {
	VZVirtioFileSystemDeviceConfigurationClassOnce.Do(func() {
		VZVirtioFileSystemDeviceConfigurationClass = _VZVirtioFileSystemDeviceConfigurationClass{objc.GetClass("VZVirtioFileSystemDeviceConfiguration")}
	})
	return VZVirtioFileSystemDeviceConfigurationClass
}

type _VZVirtioFileSystemDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioFileSystemDeviceConfiguration] class.
type IVZVirtioFileSystemDeviceConfiguration interface {
	IVZDirectorySharingDeviceConfiguration
	Share() VZDirectoryShare
	SetShare(value VZDirectoryShare)
	Tag() string
	SetTag(value string)
}

// An object that represents the configuration of a Virtio file system device.
//
// Use to create a Virtio file system device which allows the host to expose directories to a guest using a label. The example below shows the creation of a that shares a single directory that the user can manually mount after creating a mount point in the guest VM: A can also share multiple directories. The example below demonstrates sharing the and directories from the user’s home directory to the guest VM:


// An object that represents the configuration of a Virtio file system device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration
type VZVirtioFileSystemDeviceConfiguration struct {
	VZDirectorySharingDeviceConfiguration
}

// VZVirtioFileSystemDeviceConfigurationFrom constructs a [VZVirtioFileSystemDeviceConfiguration] from an unsafe.Pointer.
//
// An object that represents the configuration of a Virtio file system device.
func VZVirtioFileSystemDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioFileSystemDeviceConfiguration {
	return VZVirtioFileSystemDeviceConfiguration{
		VZDirectorySharingDeviceConfiguration: VZDirectorySharingDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioFileSystemDeviceConfigurationClass) Alloc() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioFileSystemDeviceConfigurationClass) New() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioFileSystemDeviceConfiguration) Init() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioFileSystemDeviceConfiguration) Autorelease() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFileSystemDeviceConfiguration creates a new VZVirtioFileSystemDeviceConfiguration instance.
func NewVZVirtioFileSystemDeviceConfiguration() VZVirtioFileSystemDeviceConfiguration {
	return getVZVirtioFileSystemDeviceConfigurationClass().New()
}



// Checks to see whether a Virtio tag is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration/validateTag(_:)
func (vc _VZVirtioFileSystemDeviceConfigurationClass) ValidateTagError(tag string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateTag:error:"), objc.String(tag), error_)
	return rv
}


// A value that defines how the host exposes resources to the guest virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiofilesystemdeviceconfiguration/share
func (v_ VZVirtioFileSystemDeviceConfiguration) Share() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("share"))
	return rv
}


// A value that defines how the host exposes resources to the guest virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiofilesystemdeviceconfiguration/share
func (v_ VZVirtioFileSystemDeviceConfiguration) SetShare(value VZDirectoryShare) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShare:"), value)
}


// A label that identifies this device in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiofilesystemdeviceconfiguration/tag
func (v_ VZVirtioFileSystemDeviceConfiguration) Tag() string {
	rv := objc.Send[string](v_.ID, objc.Sel("tag"))
	return rv
}


// A label that identifies this device in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiofilesystemdeviceconfiguration/tag
func (v_ VZVirtioFileSystemDeviceConfiguration) SetTag(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTag:"), objc.String(value))
}



