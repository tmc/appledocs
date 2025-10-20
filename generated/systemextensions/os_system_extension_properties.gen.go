// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSSystemExtensionProperties] class.
var (
	OSSystemExtensionPropertiesClass     _OSSystemExtensionPropertiesClass
	OSSystemExtensionPropertiesClassOnce sync.Once
)

func getOSSystemExtensionPropertiesClass() _OSSystemExtensionPropertiesClass {
	OSSystemExtensionPropertiesClassOnce.Do(func() {
		OSSystemExtensionPropertiesClass = _OSSystemExtensionPropertiesClass{objc.GetClass("OSSystemExtensionProperties")}
	})
	return OSSystemExtensionPropertiesClass
}

type _OSSystemExtensionPropertiesClass struct {
	class objc.Class
}

// An interface definition for the [OSSystemExtensionProperties] class.
type IOSSystemExtensionProperties interface {
	objectivec.IObject
}

// Properties that identify a specific version of a system extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties
type OSSystemExtensionProperties struct {
	objectivec.Object
}

// OSSystemExtensionPropertiesFrom constructs a [OSSystemExtensionProperties] from an unsafe.Pointer.
//
// Properties that identify a specific version of a system extension.
func OSSystemExtensionPropertiesFrom(ptr unsafe.Pointer) OSSystemExtensionProperties {
	return OSSystemExtensionProperties{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionPropertiesClass) Alloc() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OSSystemExtensionPropertiesClass) New() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionProperties) Init() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionProperties) Autorelease() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionProperties creates a new OSSystemExtensionProperties instance.
func NewOSSystemExtensionProperties() OSSystemExtensionProperties {
	return getOSSystemExtensionPropertiesClass().New()
}


// The bundle version of the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleVersion
func (o_ OSSystemExtensionProperties) BundleVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("bundleVersion"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isAwaitingUserApproval
func (o_ OSSystemExtensionProperties) IsAwaitingUserApproval() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAwaitingUserApproval"))
	return rv
}



