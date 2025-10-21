// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioFileSystemDevice] class.
var (
	VZVirtioFileSystemDeviceClass     _VZVirtioFileSystemDeviceClass
	VZVirtioFileSystemDeviceClassOnce sync.Once
)

func getVZVirtioFileSystemDeviceClass() _VZVirtioFileSystemDeviceClass {
	VZVirtioFileSystemDeviceClassOnce.Do(func() {
		VZVirtioFileSystemDeviceClass = _VZVirtioFileSystemDeviceClass{objc.GetClass("VZVirtioFileSystemDevice")}
	})
	return VZVirtioFileSystemDeviceClass
}

type _VZVirtioFileSystemDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioFileSystemDevice] class.
type IVZVirtioFileSystemDevice interface {
	IVZDirectorySharingDevice
}

// An object the defines a VIRTIO file system device.
//
// This device exposes host resources to the guest as a file system mount. The directory share defines which resources the host exposes to the guest. Create this device by instantiating a in a . The file system device is available in the . property. The guest can use the label to mount and access the host resources. With , the framework enforces several permissions policies for shared directories: The framework reads and writes files using the user ID (UID) of the effective user, which is the UID of the current user, rather than the UID of the system process. The framework doesn’t allow reading or overwriting of files with permissions where the file is inaccessible to the current user. The framework ignores requests from guest operating systems to change the UID or group ID (GID) of files on the host.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice
type VZVirtioFileSystemDevice struct {
	VZDirectorySharingDevice
}

// VZVirtioFileSystemDeviceFrom constructs a [VZVirtioFileSystemDevice] from an unsafe.Pointer.
//
// An object the defines a VIRTIO file system device.
func VZVirtioFileSystemDeviceFrom(ptr unsafe.Pointer) VZVirtioFileSystemDevice {
	return VZVirtioFileSystemDevice{
		VZDirectorySharingDevice: VZDirectorySharingDeviceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioFileSystemDeviceClass) Alloc() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioFileSystemDeviceClass) New() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioFileSystemDevice) Init() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioFileSystemDevice) Autorelease() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFileSystemDevice creates a new VZVirtioFileSystemDevice instance.
func NewVZVirtioFileSystemDevice() VZVirtioFileSystemDevice {
	return getVZVirtioFileSystemDeviceClass().New()
}


// A value that defines the directory share the host exposes to the guest VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/share
func (v_ VZVirtioFileSystemDevice) Share() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("share"))
	return rv
}


// SetShare sets the value of the share property.
// A value that defines the directory share the host exposes to the guest VM.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/share
func (v_ VZVirtioFileSystemDevice) SetShare(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShare:"), value)
}

// A string that identifies the device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/tag
func (v_ VZVirtioFileSystemDevice) Tag() string {
	rv := objc.Send[string](v_.ID, objc.Sel("tag"))
	return rv
}

// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtioFileSystemDevice) DirectorySharingDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
}


// SetDirectorySharingDevices sets the value of the directorySharingDevices property.
// The list of configured directory-sharing devices on the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtioFileSystemDevice) SetDirectorySharingDevices(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectorySharingDevices:"), value)
}



