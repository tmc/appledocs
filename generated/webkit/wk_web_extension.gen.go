// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtension */

/* debug [class_header]: Header for WKWebExtension */
// The class instance for the [WebExtension] class.
var (
	WebExtensionClass     _WebExtensionClass
	WebExtensionClassOnce sync.Once
)

func getWebExtensionClass() _WebExtensionClass {
	WebExtensionClassOnce.Do(func() {
		WebExtensionClass = _WebExtensionClass{objc.GetClass("WKWebExtension")}
	})
	return WebExtensionClass
}

type _WebExtensionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebExtension */
// An interface definition for the [WebExtension] class.
type IWebExtension interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebExtension */
	// properties:
	AllRequestedMatchPatterns() unsafe.Pointer
	DefaultLocale() foundation.Locale
	DisplayActionLabel() objc.IObject /* cross-framework: NSString */
	DisplayDescription() objc.IObject /* cross-framework: NSString */
	DisplayName() objc.IObject        /* cross-framework: NSString */
	DisplayShortName() objc.IObject   /* cross-framework: NSString */
	DisplayVersion() objc.IObject     /* cross-framework: NSString */
	Errors() []objc.IObject           /* cross-framework: Error */
	HasBackgroundContent() bool
	HasCommands() bool
	HasContentModificationRules() bool
	HasInjectedContent() bool
	HasOptionsPage() bool
	HasOverrideNewTabPage() bool
	HasPersistentBackgroundContent() bool
	Manifest() foundation.IDictionary
	ManifestVersion() float64
	OptionalPermissionMatchPatterns() unsafe.Pointer
	OptionalPermissions() unsafe.Pointer
	RequestedPermissionMatchPatterns() unsafe.Pointer
	RequestedPermissions() unsafe.Pointer
	Version() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebExtension */
	// methods:
	ActionIconForSize(size corefoundation.CGSize) appkit.Image
	IconForSize(size corefoundation.CGSize) appkit.Image
	SupportsManifestVersion(manifestVersion float64) bool
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebExtension */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionClass) Alloc() WebExtension {
	rv := objc.Send[WebExtension](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionClass) New() WebExtension {
	rv := objc.Send[WebExtension](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtension) Init() WebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtension) Autorelease() WebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtension creates a new WebExtension instance.
func NewWebExtension() WebExtension {
	return getWebExtensionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebExtension */
// An object that encapsulates a web extension’s resources that the manifest file defines.
//
// This class reads and parses the file along with the supporting resources like icons and localizations.

// An object that encapsulates a web extension’s resources that the manifest file defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension
type WebExtension struct {
	objectivec.Object
}

// WebExtensionFrom constructs a [WebExtension] from an unsafe.Pointer.
//
// An object that encapsulates a web extension’s resources that the manifest file defines.
func WebExtensionFrom(ptr unsafe.Pointer) WebExtension {
	return WebExtension{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebExtension */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebExtension */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebExtension */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebExtension */

// Returns the default action icon for the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/actionIcon(for:)
func (w_ WebExtension) ActionIconForSize(size corefoundation.CGSize) appkit.Image {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("actionIconForSize:"), size)
	return rv
} /* debug [instance_methods/method]: ActionIconForSize */

// Returns the extension’s icon image for the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/icon(for:)
func (w_ WebExtension) IconForSize(size corefoundation.CGSize) appkit.Image {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("iconForSize:"), size)
	return rv
} /* debug [instance_methods/method]: IconForSize */

// Checks if a manifest version is supported by the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/supportsManifestVersion(_:)
func (w_ WebExtension) SupportsManifestVersion(manifestVersion float64) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("supportsManifestVersion:"), manifestVersion)
	return rv
} /* debug [instance_methods/method]: SupportsManifestVersion */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebExtension */

// The set of websites that the extension requires access to for injected content and for receiving messages from websites.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/allRequestedMatchPatterns
func (w_ WebExtension) AllRequestedMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("allRequestedMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: allRequestedMatchPatterns */

// The default locale for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/defaultLocale
func (w_ WebExtension) DefaultLocale() foundation.Locale {
	rv := objc.Send[foundation.Locale](w_.ID, objc.Sel("defaultLocale"))
	return rv
} /* debug [instance_properties/getter]: defaultLocale */

// The default localized extension action label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayActionLabel
func (w_ WebExtension) DisplayActionLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayActionLabel"))
	return rv
} /* debug [instance_properties/getter]: displayActionLabel */

// The localized extension description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayDescription
func (w_ WebExtension) DisplayDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayDescription"))
	return rv
} /* debug [instance_properties/getter]: displayDescription */

// The localized extension name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayName
func (w_ WebExtension) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayName"))
	return rv
} /* debug [instance_properties/getter]: displayName */

// The localized extension short name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayShortName
func (w_ WebExtension) DisplayShortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayShortName"))
	return rv
} /* debug [instance_properties/getter]: displayShortName */

// The localized extension display version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayVersion
func (w_ WebExtension) DisplayVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayVersion"))
	return rv
} /* debug [instance_properties/getter]: displayVersion */

// An array of all errors that occurred during the processing of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/errors
func (w_ WebExtension) Errors() []objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[[]coretelephony.Error](w_.ID, objc.Sel("errors"))
	return rv
} /* debug [instance_properties/getter]: errors */

// A Boolean value indicating whether the extension has background content that can run when needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasBackgroundContent
func (w_ WebExtension) HasBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasBackgroundContent"))
	return rv
} /* debug [instance_properties/getter]: hasBackgroundContent */

// A Boolean value indicating whether the extension includes commands that users can invoke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasCommands
func (w_ WebExtension) HasCommands() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasCommands"))
	return rv
} /* debug [instance_properties/getter]: hasCommands */

// A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasContentModificationRules
func (w_ WebExtension) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
} /* debug [instance_properties/getter]: hasContentModificationRules */

// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasInjectedContent
func (w_ WebExtension) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
} /* debug [instance_properties/getter]: hasInjectedContent */

// A Boolean value indicating whether the extension has an options page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasOptionsPage
func (w_ WebExtension) HasOptionsPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOptionsPage"))
	return rv
} /* debug [instance_properties/getter]: hasOptionsPage */

// A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasOverrideNewTabPage
func (w_ WebExtension) HasOverrideNewTabPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOverrideNewTabPage"))
	return rv
} /* debug [instance_properties/getter]: hasOverrideNewTabPage */

// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasPersistentBackgroundContent
func (w_ WebExtension) HasPersistentBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasPersistentBackgroundContent"))
	return rv
} /* debug [instance_properties/getter]: hasPersistentBackgroundContent */

// The parsed manifest as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/manifest
func (w_ WebExtension) Manifest() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("manifest"))
	return rv
} /* debug [instance_properties/getter]: manifest */

// The parsed manifest version, or if there is no version specified in the manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/manifestVersion
func (w_ WebExtension) ManifestVersion() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("manifestVersion"))
	return rv
} /* debug [instance_properties/getter]: manifestVersion */

// The set of websites that the extension may need access to for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/optionalPermissionMatchPatterns
func (w_ WebExtension) OptionalPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("optionalPermissionMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: optionalPermissionMatchPatterns */

// The set of permissions that the extension may need for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/optionalPermissions
func (w_ WebExtension) OptionalPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("optionalPermissions"))
	return rv
} /* debug [instance_properties/getter]: optionalPermissions */

// The set of websites that the extension requires access to for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/requestedPermissionMatchPatterns
func (w_ WebExtension) RequestedPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("requestedPermissionMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: requestedPermissionMatchPatterns */

// The set of permissions that the extension requires for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/requestedPermissions
func (w_ WebExtension) RequestedPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("requestedPermissions"))
	return rv
} /* debug [instance_properties/getter]: requestedPermissions */

// The extension version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/version
func (w_ WebExtension) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("version"))
	return rv
} /* debug [instance_properties/getter]: version */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKWebExtension */
