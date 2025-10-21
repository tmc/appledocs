// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZMacOSRestoreImage] class.
type IVZMacOSRestoreImage interface {
	objectivec.IObject
}

// An object that describes a version of macOS to install on to a virtual machine.
//
// To set up a new VM compatible with the restore image, use to obtain the of the . Then, create a object by loading an installation media file. Initialize a object with this object to install the operating system onto a VM.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZMacOSRestoreImageClass) Alloc() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Load a restore image from a file on the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/image(from:)
func (vc _VZMacOSRestoreImageClass) LoadFileURLCompletionHandler(fileURL foundation.URL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("loadFileURL:completionHandler:"), fileURL, completionHandler)
}

// Fetches the latest restore image supported by this host from the network.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/latestSupported
func (vc _VZMacOSRestoreImageClass) FetchLatestSupportedWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("fetchLatestSupportedWithCompletionHandler:"), completionHandler)
}

// The Mac hardware model.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacOSRestoreImage) HardwareModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("hardwareModel"))
	return rv
}


// SetHardwareModel sets the value of the hardwareModel property.
// The Mac hardware model.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (v_ VZMacOSRestoreImage) SetHardwareModel(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHardwareModel:"), value)
}

// A Boolean value that indicates whether the current host supports this restore image.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/issupported
func (v_ VZMacOSRestoreImage) IsSupported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isSupported"))
	return rv
}


// SetIsSupported sets the value of the isSupported property.
// A Boolean value that indicates whether the current host supports this restore image.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/issupported
func (v_ VZMacOSRestoreImage) SetIsSupported(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsSupported:"), value)
}

// The build version this restore image contains.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/buildVersion
func (v_ VZMacOSRestoreImage) BuildVersion() string {
	rv := objc.Send[string](v_.ID, objc.Sel("buildVersion"))
	return rv
}

// A Boolean value that indicates whether the current host supports this restore image.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/isSupported
func (v_ VZMacOSRestoreImage) Supported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("supported"))
	return rv
}

// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/mostFeaturefulSupportedConfiguration
func (v_ VZMacOSRestoreImage) MostFeaturefulSupportedConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return rv
}

// The operating system version this restore image contains.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/operatingSystemVersion
func (v_ VZMacOSRestoreImage) OperatingSystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}

// The URL of this restore image.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/url
func (v_ VZMacOSRestoreImage) URL() foundation.URL {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("URL"))
	return rv
}



