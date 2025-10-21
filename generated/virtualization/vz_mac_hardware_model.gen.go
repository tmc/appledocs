// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMacHardwareModel] class.
var (
	VZMacHardwareModelClass     _VZMacHardwareModelClass
	VZMacHardwareModelClassOnce sync.Once
)

func getVZMacHardwareModelClass() _VZMacHardwareModelClass {
	VZMacHardwareModelClassOnce.Do(func() {
		VZMacHardwareModelClass = _VZMacHardwareModelClass{objc.GetClass("VZMacHardwareModel")}
	})
	return VZMacHardwareModelClass
}

type _VZMacHardwareModelClass struct {
	class objc.Class
}

// An interface definition for the [VZMacHardwareModel] class.
type IVZMacHardwareModel interface {
	objectivec.IObject
}

// A specification for the hardware elements and configurations present in a particular Mac hardware model.
//
// The Mac hardware model abstracts a set of virtualized hardware elements and configurations. A version of macOS may only run on certain hardware models. Additionally, the host may also only provide certain hardware models based on the version of macOS and the underlying hardware. The property allows you to discover if the current host supports a particular hardware model. Choosing the hardware model starts from a restore image with . A restore image describes its supported configuration requirements through its property. A configuration requirements object has a corresponding hardware model that you can use to configure a VM that meets the requirements. After obtaining the hardware model, use the platform configuration’s to configure the Mac platform object and use to create its auxiliary storage. After creating the VM, use to install macOS on it. If you serialize the VM on disk, preserve the hardware model used for installation for subsequent boots. The property provides a unique binary representation that you serialize to the file system. You can recreate the hardware model from the serialized binary representation with .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel
type VZMacHardwareModel struct {
	objectivec.Object
}

// VZMacHardwareModelFrom constructs a [VZMacHardwareModel] from an unsafe.Pointer.
//
// A specification for the hardware elements and configurations present in a particular Mac hardware model.
func VZMacHardwareModelFrom(ptr unsafe.Pointer) VZMacHardwareModel {
	return VZMacHardwareModel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacHardwareModelClass) Alloc() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacHardwareModelClass) New() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacHardwareModel) Init() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacHardwareModel) Autorelease() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModel creates a new VZMacHardwareModel instance.
func NewVZMacHardwareModel() VZMacHardwareModel {
	return getVZMacHardwareModelClass().New()
}




// Creates an instance of the hardware model described by the specified data representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/init(dataRepresentation:)
func NewVZMacHardwareModelWithDataRepresentation(dataRepresentation foundation.IData) VZMacHardwareModel {
	instance := getVZMacHardwareModelClass().Alloc()
	rv := objc.Send[VZMacHardwareModel](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	rv.Autorelease()
	return rv
}


// Returns the opaque data representation of the hardware model.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/dataRepresentation
func (v_ VZMacHardwareModel) DataRepresentation() foundation.NSData {
	rv := objc.Send[foundation.NSData](v_.ID, objc.Sel("dataRepresentation"))
	return rv
}

// A Boolean value that indicates whether the host supports this hardware model.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/isSupported
func (v_ VZMacHardwareModel) Supported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("supported"))
	return rv
}

// A Boolean value that indicates whether the host supports this hardware model.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmachardwaremodel/issupported
func (v_ VZMacHardwareModel) IsSupported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isSupported"))
	return rv
}


// SetIsSupported sets the value of the isSupported property.
// A Boolean value that indicates whether the host supports this hardware model.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmachardwaremodel/issupported
func (v_ VZMacHardwareModel) SetIsSupported(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsSupported:"), value)
}

// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacHardwareModel) MostFeaturefulSupportedConfiguration() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return rv
}


// SetMostFeaturefulSupportedConfiguration sets the value of the mostFeaturefulSupportedConfiguration property.
// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacHardwareModel) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
}

// The Mac hardware model.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacHardwareModel) HardwareModel() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("hardwareModel"))
	return rv
}


// SetHardwareModel sets the value of the hardwareModel property.
// The Mac hardware model.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacHardwareModel) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHardwareModel:"), value)
}


