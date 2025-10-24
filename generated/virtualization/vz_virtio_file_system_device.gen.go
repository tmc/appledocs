// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioFileSystemDevice */

/* debug [class_header]: Header for VZVirtioFileSystemDevice */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioFileSystemDevice */
// An interface definition for the [VZVirtioFileSystemDevice] class.
type IVZVirtioFileSystemDevice interface {
	IVZDirectorySharingDevice

	/* debug [class_interface_properties]: Properties for VZVirtioFileSystemDevice */
	// properties:
	Share() IVZDirectoryShare
	SetShare(value IVZDirectoryShare)
	Tag() objc.IObject /* cross-framework: NSString */
	DirectorySharingDevices() IVZDirectorySharingDevice
	SetDirectorySharingDevices(value IVZDirectorySharingDevice)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioFileSystemDevice */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioFileSystemDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioFileSystemDeviceClass) Alloc() VZVirtioFileSystemDevice {
	rv := objc.Send[VZVirtioFileSystemDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioFileSystemDevice */
// An object the defines a VIRTIO file system device.
//
// This device exposes host resources to the guest as a file system mount. The directory share defines which resources the host exposes to the guest. Create this device by instantiating a in a . The file system device is available in the . property. The guest can use the label to mount and access the host resources. With , the framework enforces several permissions policies for shared directories: The framework reads and writes files using the user ID (UID) of the effective user, which is the UID of the current user, rather than the UID of the system process. The framework doesn’t allow reading or overwriting of files with permissions where the file is inaccessible to the current user. The framework ignores requests from guest operating systems to change the UID or group ID (GID) of files on the host.

// An object the defines a VIRTIO file system device.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioFileSystemDevice */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioFileSystemDevice */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioFileSystemDevice */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioFileSystemDevice */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioFileSystemDevice */

// A value that defines the directory share the host exposes to the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/share
func (v_ VZVirtioFileSystemDevice) Share() IVZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("share"))
	return rv
} /* debug [instance_properties/getter]: share */

// A value that defines the directory share the host exposes to the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/share
func (v_ VZVirtioFileSystemDevice) SetShare(value IVZDirectoryShare) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShare:"), value)
} /* debug [instance_properties/setter]: share */

// A string that identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDevice/tag
func (v_ VZVirtioFileSystemDevice) Tag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("tag"))
	return rv
} /* debug [instance_properties/getter]: tag */

// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtioFileSystemDevice) DirectorySharingDevices() IVZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("directorySharingDevices"))
	return rv
} /* debug [instance_properties/getter]: directorySharingDevices */

// The list of configured directory-sharing devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/directorysharingdevices
func (v_ VZVirtioFileSystemDevice) SetDirectorySharingDevices(value IVZDirectorySharingDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectorySharingDevices:"), value)
} /* debug [instance_properties/setter]: directorySharingDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioFileSystemDevice */
