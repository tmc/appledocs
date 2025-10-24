// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [OSSystemExtensionInfo] class.
type IOSSystemExtensionInfo interface {
	objectivec.IObject
	// properties:
	BundleIdentifier() objc.IObject /* cross-framework: NSString */
	BundleShortVersion() objc.IObject /* cross-framework: NSString */
	SetBundleShortVersion(value objc.IObject /* cross-framework: NSString */)
	BundleVersion() objc.IObject /* cross-framework: NSString */
	SetBundleVersion(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo
type OSSystemExtensionInfo struct {
	objectivec.Object
}

// OSSystemExtensionInfoFrom constructs a [OSSystemExtensionInfo] from an unsafe.Pointer.
func OSSystemExtensionInfoFrom(ptr unsafe.Pointer) OSSystemExtensionInfo {
	return OSSystemExtensionInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionInfoClass) Alloc() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleIdentifier
func (o_ OSSystemExtensionInfo) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensioninfo/bundleshortversion
func (o_ OSSystemExtensionInfo) BundleShortVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleShortVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensioninfo/bundleshortversion
func (o_ OSSystemExtensionInfo) SetBundleShortVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBundleShortVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensioninfo/bundleversion
func (o_ OSSystemExtensionInfo) BundleVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("bundleVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensioninfo/bundleversion
func (o_ OSSystemExtensionInfo) SetBundleVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBundleVersion:"), value)
}



