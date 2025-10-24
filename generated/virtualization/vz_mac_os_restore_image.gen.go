// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacOSRestoreImage */

/* debug [class_header]: Header for VZMacOSRestoreImage */
// The class instance for the [VZMacOSRestoreImage] class.
var (
	VZMacOSRestoreImageClass     _VZMacOSRestoreImageClass
	VZMacOSRestoreImageClassOnce sync.Once
)

func getVZMacOSRestoreImageClass() _VZMacOSRestoreImageClass {
	VZMacOSRestoreImageClassOnce.Do(func() {
		VZMacOSRestoreImageClass = _VZMacOSRestoreImageClass{objc.GetClass("VZMacOSRestoreImage")}
	})
	return VZMacOSRestoreImageClass
}

type _VZMacOSRestoreImageClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMacOSRestoreImage */
// An interface definition for the [VZMacOSRestoreImage] class.
type IVZMacOSRestoreImage interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZMacOSRestoreImage */
	// properties:
	BuildVersion() objc.IObject /* cross-framework: NSString */
	Supported() bool
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	OperatingSystemVersion() foundation.OperatingSystemVersion
	URL() objc.IObject /* cross-framework: NSURL */
	IsSupported() bool
	SetIsSupported(value bool)
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMacOSRestoreImage */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMacOSRestoreImage */
// Alloc allocates a new instance without initialization.
func (vc _VZMacOSRestoreImageClass) Alloc() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacOSRestoreImageClass) New() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSRestoreImage) Init() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSRestoreImage) Autorelease() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSRestoreImage creates a new VZMacOSRestoreImage instance.
func NewVZMacOSRestoreImage() VZMacOSRestoreImage {
	return getVZMacOSRestoreImageClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMacOSRestoreImage */
// An object that describes a version of macOS to install on to a virtual machine.
//
// To set up a new VM compatible with the restore image, use to obtain the of the . Then, create a object by loading an installation media file. Initialize a object with this object to install the operating system onto a VM.

// An object that describes a version of macOS to install on to a virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage
type VZMacOSRestoreImage struct {
	objectivec.Object
}

// VZMacOSRestoreImageFrom constructs a [VZMacOSRestoreImage] from an unsafe.Pointer.
//
// An object that describes a version of macOS to install on to a virtual machine.
func VZMacOSRestoreImageFrom(ptr unsafe.Pointer) VZMacOSRestoreImage {
	return VZMacOSRestoreImage{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMacOSRestoreImage */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMacOSRestoreImage */

// Load a restore image from a file on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/image(from:)
func (vc _VZMacOSRestoreImageClass) LoadFileURLCompletionHandler(fileURL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("loadFileURL:completionHandler:"), fileURL, completionHandler)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=LoadFileURLCompletionHandler) */

// Fetches the latest restore image supported by this host from the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/latestSupported
func (vc _VZMacOSRestoreImageClass) FetchLatestSupportedWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("fetchLatestSupportedWithCompletionHandler:"), completionHandler)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=FetchLatestSupportedWithCompletionHandler) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMacOSRestoreImage */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMacOSRestoreImage */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMacOSRestoreImage */

// The build version this restore image contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/buildVersion
func (v_ VZMacOSRestoreImage) BuildVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("buildVersion"))
	return rv
} /* debug [instance_properties/getter]: buildVersion */

// A Boolean value that indicates whether the current host supports this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/isSupported
func (v_ VZMacOSRestoreImage) Supported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("supported"))
	return rv
} /* debug [instance_properties/getter]: supported */

// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/mostFeaturefulSupportedConfiguration
func (v_ VZMacOSRestoreImage) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return rv
} /* debug [instance_properties/getter]: mostFeaturefulSupportedConfiguration */

// The operating system version this restore image contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/operatingSystemVersion
func (v_ VZMacOSRestoreImage) OperatingSystemVersion() foundation.OperatingSystemVersion {
	rv := objc.Send[foundation.OperatingSystemVersion](v_.ID, objc.Sel("operatingSystemVersion"))
	return rv
} /* debug [instance_properties/getter]: operatingSystemVersion */

// The URL of this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/url
func (v_ VZMacOSRestoreImage) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

// A Boolean value that indicates whether the current host supports this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/issupported
func (v_ VZMacOSRestoreImage) IsSupported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isSupported"))
	return rv
} /* debug [instance_properties/getter]: isSupported */

// A Boolean value that indicates whether the current host supports this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/issupported
func (v_ VZMacOSRestoreImage) SetIsSupported(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsSupported:"), value)
} /* debug [instance_properties/setter]: isSupported */

// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacOSRestoreImage) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("hardwareModel"))
	return rv
} /* debug [instance_properties/getter]: hardwareModel */

// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacOSRestoreImage) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHardwareModel:"), value)
} /* debug [instance_properties/setter]: hardwareModel */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMacOSRestoreImage */
