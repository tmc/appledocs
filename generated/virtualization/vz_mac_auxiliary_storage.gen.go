// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacAuxiliaryStorage */


/* debug [class_header]: Header for VZMacAuxiliaryStorage */
// The class instance for the [VZMacAuxiliaryStorage] class.
var (
	VZMacAuxiliaryStorageClass     _VZMacAuxiliaryStorageClass
	VZMacAuxiliaryStorageClassOnce sync.Once
)

func getVZMacAuxiliaryStorageClass() _VZMacAuxiliaryStorageClass {
	VZMacAuxiliaryStorageClassOnce.Do(func() {
		VZMacAuxiliaryStorageClass = _VZMacAuxiliaryStorageClass{objc.GetClass("VZMacAuxiliaryStorage")}
	})
	return VZMacAuxiliaryStorageClass
}

type _VZMacAuxiliaryStorageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacAuxiliaryStorage */
// An interface definition for the [VZMacAuxiliaryStorage] class.
type IVZMacAuxiliaryStorage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZMacAuxiliaryStorage */
	// properties:
	URL() objc.IObject /* cross-framework: NSURL */
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements)
	AuxiliaryStorage() IVZMacAuxiliaryStorage
	SetAuxiliaryStorage(value IVZMacAuxiliaryStorage)
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacAuxiliaryStorage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacAuxiliaryStorage */
// Alloc allocates a new instance without initialization.
func (vc _VZMacAuxiliaryStorageClass) Alloc() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacAuxiliaryStorageClass) New() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacAuxiliaryStorage) Init() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacAuxiliaryStorage) Autorelease() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacAuxiliaryStorage creates a new VZMacAuxiliaryStorage instance.
func NewVZMacAuxiliaryStorage() VZMacAuxiliaryStorage {
	return getVZMacAuxiliaryStorageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacAuxiliaryStorage */
// An object that contains information the boot loader needs for booting macOS as a guest operating system.
//
// The Mac auxiliary storage contains data used by the boot loader and the guest operating system. It’s necessary to boot a macOS guest OS. When creating a new VM, use to create a default initialized auxiliary storage. The hardware model you use when creating the new auxiliary storage depends on the restore image that you’ll use for installation. From the restore image, use to get a supported configuration. A configuration has a associated with it. After initializing the new auxiliary storage, set it on . . The hardware model in . must be identical to the one used to create the empty auxiliary storage., otherwise the behavior isn’t defined. When installing macOS, the lays out data on the auxiliary storage. After installation, the macOS guest uses the auxiliary storage for every subsequent boot. When moving or performing a backup of a VM, you must move or copy the file containing the auxiliary storage along with the main disk image. To boot a VM created with , use to set up the auxiliary storage from the existing file used during installation. When using an existing file, the hardware model of the . must match the hardware model used when creating the original file.


// An object that contains information the boot loader needs for booting macOS as a guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
type VZMacAuxiliaryStorage struct {
	objectivec.Object
}

// VZMacAuxiliaryStorageFrom constructs a [VZMacAuxiliaryStorage] from an unsafe.Pointer.
//
// An object that contains information the boot loader needs for booting macOS as a guest operating system.
func VZMacAuxiliaryStorageFrom(ptr unsafe.Pointer) VZMacAuxiliaryStorage {
	return VZMacAuxiliaryStorage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacAuxiliaryStorage */

// Creates an initialized Mac auxiliary storage instance that describes a specific hardware model at a URL you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(creatingStorageAt:hardwareModel:options:)
func NewVZMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError(URL objc.IObject /* cross-framework: NSURL */, hardwareModel IVZMacHardwareModel, options VZMacAuxiliaryStorageInitializationOptions, error_ objectivec.IObject) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initCreatingStorageAtURL:hardwareModel:options:error:"), URL, hardwareModel, options, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError */


// Initializes an auxiliary storage object with data from the location at the URL you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOfURL:)
func NewVZMacAuxiliaryStorageWithContentsOfURL(URL objc.IObject /* cross-framework: NSURL */) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initWithContentsOfURL:"), URL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMacAuxiliaryStorageWithContentsOfURL */


// Initializes an auxiliary storage object with data from the location at the URL you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(url:)
func NewVZMacAuxiliaryStorageWithURL(URL objc.IObject /* cross-framework: NSURL */) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMacAuxiliaryStorageWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacAuxiliaryStorage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacAuxiliaryStorage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacAuxiliaryStorage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacAuxiliaryStorage */

// The URL of the auxiliary storage on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/url
func (v_ VZMacAuxiliaryStorage) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacAuxiliaryStorage) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return rv
}/* debug [instance_properties/getter]: mostFeaturefulSupportedConfiguration */


// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacAuxiliaryStorage) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
}/* debug [instance_properties/setter]: mostFeaturefulSupportedConfiguration */


// The Mac auxiliary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/auxiliarystorage
func (v_ VZMacAuxiliaryStorage) AuxiliaryStorage() IVZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("auxiliaryStorage"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryStorage */


// The Mac auxiliary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/auxiliarystorage
func (v_ VZMacAuxiliaryStorage) SetAuxiliaryStorage(value IVZMacAuxiliaryStorage) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAuxiliaryStorage:"), value)
}/* debug [instance_properties/setter]: auxiliaryStorage */


// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacAuxiliaryStorage) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("hardwareModel"))
	return rv
}/* debug [instance_properties/getter]: hardwareModel */


// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacAuxiliaryStorage) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHardwareModel:"), value)
}/* debug [instance_properties/setter]: hardwareModel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacAuxiliaryStorage */


