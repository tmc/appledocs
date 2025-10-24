// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSSystemExtensionInfo */

/* debug [class_header]: Header for OSSystemExtensionInfo */
// The class instance for the [OSSystemExtensionInfo] class.
var (
	OSSystemExtensionInfoClass     _OSSystemExtensionInfoClass
	OSSystemExtensionInfoClassOnce sync.Once
)

func getOSSystemExtensionInfoClass() _OSSystemExtensionInfoClass {
	OSSystemExtensionInfoClassOnce.Do(func() {
		OSSystemExtensionInfoClass = _OSSystemExtensionInfoClass{objc.GetClass("OSSystemExtensionInfo")}
	})
	return OSSystemExtensionInfoClass
}

type _OSSystemExtensionInfoClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OSSystemExtensionInfo */
// An interface definition for the [OSSystemExtensionInfo] class.
type IOSSystemExtensionInfo interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for OSSystemExtensionInfo */
	// properties:
	BundleIdentifier() objc.IObject   /* cross-framework: NSString */
	BundleShortVersion() objc.IObject /* cross-framework: NSString */
	BundleVersion() objc.IObject      /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OSSystemExtensionInfo */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OSSystemExtensionInfo */
// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionInfoClass) Alloc() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSSystemExtensionInfoClass) New() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionInfo) Init() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionInfo) Autorelease() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionInfo creates a new OSSystemExtensionInfo instance.
func NewOSSystemExtensionInfo() OSSystemExtensionInfo {
	return getOSSystemExtensionInfoClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OSSystemExtensionInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo
type OSSystemExtensionInfo struct {
	objectivec.Object
}

// OSSystemExtensionInfoFrom constructs a [OSSystemExtensionInfo] from an unsafe.Pointer.
func OSSystemExtensionInfoFrom(ptr unsafe.Pointer) OSSystemExtensionInfo {
	return OSSystemExtensionInfo{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OSSystemExtensionInfo */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OSSystemExtensionInfo */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OSSystemExtensionInfo */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OSSystemExtensionInfo */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OSSystemExtensionInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleIdentifier
func (o_ OSSystemExtensionInfo) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleIdentifier"))
	return rv
} /* debug [instance_properties/getter]: bundleIdentifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleShortVersion
func (o_ OSSystemExtensionInfo) BundleShortVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleShortVersion"))
	return rv
} /* debug [instance_properties/getter]: bundleShortVersion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleVersion
func (o_ OSSystemExtensionInfo) BundleVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleVersion"))
	return rv
} /* debug [instance_properties/getter]: bundleVersion */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class OSSystemExtensionInfo */
