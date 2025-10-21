// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZStorageDeviceConfiguration] class.
var (
	VZStorageDeviceConfigurationClass     _VZStorageDeviceConfigurationClass
	VZStorageDeviceConfigurationClassOnce sync.Once
)

func getVZStorageDeviceConfigurationClass() _VZStorageDeviceConfigurationClass {
	VZStorageDeviceConfigurationClassOnce.Do(func() {
		VZStorageDeviceConfigurationClass = _VZStorageDeviceConfigurationClass{objc.GetClass("VZStorageDeviceConfiguration")}
	})
	return VZStorageDeviceConfigurationClass
}

type _VZStorageDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZStorageDeviceConfiguration] class.
type IVZStorageDeviceConfiguration interface {
	objectivec.IObject
}

// The common configuration traits for storage device requests.
//
// Don’t create a object directly. Instead, instantiate one of its subclasses, such as . Use the property of this class to access the device’s underlying storage.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration
type VZStorageDeviceConfiguration struct {
	objectivec.Object
}

// VZStorageDeviceConfigurationFrom constructs a [VZStorageDeviceConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for storage device requests.
func VZStorageDeviceConfigurationFrom(ptr unsafe.Pointer) VZStorageDeviceConfiguration {
	return VZStorageDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceConfigurationClass) Alloc() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZStorageDeviceConfigurationClass) New() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZStorageDeviceConfiguration) Init() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZStorageDeviceConfiguration) Autorelease() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceConfiguration creates a new VZStorageDeviceConfiguration instance.
func NewVZStorageDeviceConfiguration() VZStorageDeviceConfiguration {
	return getVZStorageDeviceConfigurationClass().New()
}


// The attachment object that provides the underlying storage for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzstoragedeviceconfiguration/attachment
func (v_ VZStorageDeviceConfiguration) Attachment() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The attachment object that provides the underlying storage for the device.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzstoragedeviceconfiguration/attachment
func (v_ VZStorageDeviceConfiguration) SetAttachment(value IVZStorageDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}



