// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZStorageDeviceConfiguration */

/* debug [class_header]: Header for VZStorageDeviceConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZStorageDeviceConfiguration */
// An interface definition for the [VZStorageDeviceConfiguration] class.
type IVZStorageDeviceConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZStorageDeviceConfiguration */
	// properties:
	Attachment() IVZStorageDeviceAttachment
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZStorageDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZStorageDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceConfigurationClass) Alloc() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZStorageDeviceConfiguration */
// The common configuration traits for storage device requests.
//
// Don’t create a object directly. Instead, instantiate one of its subclasses, such as . Use the property of this class to access the device’s underlying storage.

// The common configuration traits for storage device requests.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZStorageDeviceConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZStorageDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZStorageDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZStorageDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZStorageDeviceConfiguration */

// The attachment object that provides the underlying storage for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/attachment
func (v_ VZStorageDeviceConfiguration) Attachment() IVZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
} /* debug [instance_properties/getter]: attachment */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZStorageDeviceConfiguration */
