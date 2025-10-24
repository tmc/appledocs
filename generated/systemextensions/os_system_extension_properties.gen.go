// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSSystemExtensionProperties */

/* debug [class_header]: Header for OSSystemExtensionProperties */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OSSystemExtensionProperties */
// An interface definition for the [OSSystemExtensionProperties] class.
type IOSSystemExtensionProperties interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for OSSystemExtensionProperties */
	// properties:
	BundleIdentifier() objc.IObject   /* cross-framework: NSString */
	BundleShortVersion() objc.IObject /* cross-framework: NSString */
	BundleVersion() objc.IObject      /* cross-framework: NSString */
	IsAwaitingUserApproval() bool
	IsEnabled() bool
	IsUninstalling() bool
	URL() objc.IObject /* cross-framework: NSURL */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OSSystemExtensionProperties */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OSSystemExtensionProperties */
// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionPropertiesClass) Alloc() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OSSystemExtensionProperties */
// Properties that identify a specific version of a system extension.

// Properties that identify a specific version of a system extension.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OSSystemExtensionProperties */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OSSystemExtensionProperties */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OSSystemExtensionProperties */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OSSystemExtensionProperties */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OSSystemExtensionProperties */

// The bundle identifier of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleIdentifier
func (o_ OSSystemExtensionProperties) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleIdentifier"))
	return rv
} /* debug [instance_properties/getter]: bundleIdentifier */

// The bundle short version string of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleShortVersion
func (o_ OSSystemExtensionProperties) BundleShortVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleShortVersion"))
	return rv
} /* debug [instance_properties/getter]: bundleShortVersion */

// The bundle version of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleVersion
func (o_ OSSystemExtensionProperties) BundleVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleVersion"))
	return rv
} /* debug [instance_properties/getter]: bundleVersion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isAwaitingUserApproval
func (o_ OSSystemExtensionProperties) IsAwaitingUserApproval() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAwaitingUserApproval"))
	return rv
} /* debug [instance_properties/getter]: isAwaitingUserApproval */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isEnabled
func (o_ OSSystemExtensionProperties) IsEnabled() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEnabled"))
	return rv
} /* debug [instance_properties/getter]: isEnabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isUninstalling
func (o_ OSSystemExtensionProperties) IsUninstalling() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isUninstalling"))
	return rv
} /* debug [instance_properties/getter]: isUninstalling */

// The file URL of the extension bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/url
func (o_ OSSystemExtensionProperties) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](o_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class OSSystemExtensionProperties */
