// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
func (o_ OSSystemExtensionProperties) BundleVersion() string {
	rv := objc.Send[string](o_.ID, objc.Sel("bundleVersion"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isAwaitingUserApproval
func (o_ OSSystemExtensionProperties) IsAwaitingUserApproval() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAwaitingUserApproval"))
	return rv
}

// The bundle identifier of the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/bundleidentifier
func (o_ OSSystemExtensionProperties) BundleIdentifier() string {
	rv := objc.Send[string](o_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// SetBundleIdentifier sets the value of the bundleIdentifier property.
// The bundle identifier of the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/bundleidentifier
func (o_ OSSystemExtensionProperties) SetBundleIdentifier(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}

// The bundle short version string of the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/bundleshortversion
func (o_ OSSystemExtensionProperties) BundleShortVersion() string {
	rv := objc.Send[string](o_.ID, objc.Sel("bundleShortVersion"))
	return rv
}


// SetBundleShortVersion sets the value of the bundleShortVersion property.
// The bundle short version string of the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/bundleshortversion
func (o_ OSSystemExtensionProperties) SetBundleShortVersion(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBundleShortVersion:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/isenabled
func (o_ OSSystemExtensionProperties) IsEnabled() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/isenabled
func (o_ OSSystemExtensionProperties) SetIsEnabled(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/isuninstalling
func (o_ OSSystemExtensionProperties) IsUninstalling() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isUninstalling"))
	return rv
}


// SetIsUninstalling sets the value of the isUninstalling property.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/isuninstalling
func (o_ OSSystemExtensionProperties) SetIsUninstalling(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsUninstalling:"), value)
}

// The file URL of the extension bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/url
func (o_ OSSystemExtensionProperties) Url() foundation.URL {
	rv := objc.Send[foundation.URL](o_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The file URL of the extension bundle.

//
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionproperties/url
func (o_ OSSystemExtensionProperties) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUrl:"), value)
}



