// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [WebExtension] class.
type IWebExtension interface {
	objectivec.IObject
}

// An object that encapsulates a web extension’s resources that the manifest file defines.
//
// This class reads and parses the file along with the supporting resources like icons and localizations.
//
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

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionClass) Alloc() WebExtension {
	rv := objc.Send[WebExtension](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The set of websites that the extension requires access to for injected content and for receiving messages from websites.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/allrequestedmatchpatterns
func (w_ WebExtension) AllRequestedMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("allRequestedMatchPatterns"))
	return rv
}


// SetAllRequestedMatchPatterns sets the value of the allRequestedMatchPatterns property.
// The set of websites that the extension requires access to for injected content and for receiving messages from websites.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/allrequestedmatchpatterns
func (w_ WebExtension) SetAllRequestedMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllRequestedMatchPatterns:"), value)
}

// The default locale for the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/defaultlocale
func (w_ WebExtension) DefaultLocale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("defaultLocale"))
	return rv
}


// SetDefaultLocale sets the value of the defaultLocale property.
// The default locale for the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/defaultlocale
func (w_ WebExtension) SetDefaultLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultLocale:"), value)
}

// The default localized extension action label.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayactionlabel
func (w_ WebExtension) DisplayActionLabel() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayActionLabel"))
	return rv
}


// SetDisplayActionLabel sets the value of the displayActionLabel property.
// The default localized extension action label.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayactionlabel
func (w_ WebExtension) SetDisplayActionLabel(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayActionLabel:"), objc.String(value))
}

// The localized extension description.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displaydescription
func (w_ WebExtension) DisplayDescription() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayDescription"))
	return rv
}


// SetDisplayDescription sets the value of the displayDescription property.
// The localized extension description.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displaydescription
func (w_ WebExtension) SetDisplayDescription(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayDescription:"), objc.String(value))
}

// The localized extension name.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayname
func (w_ WebExtension) DisplayName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayName"))
	return rv
}


// SetDisplayName sets the value of the displayName property.
// The localized extension name.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayname
func (w_ WebExtension) SetDisplayName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}

// The localized extension short name.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayshortname
func (w_ WebExtension) DisplayShortName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayShortName"))
	return rv
}


// SetDisplayShortName sets the value of the displayShortName property.
// The localized extension short name.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayshortname
func (w_ WebExtension) SetDisplayShortName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayShortName:"), objc.String(value))
}

// The localized extension display version.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayversion
func (w_ WebExtension) DisplayVersion() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayVersion"))
	return rv
}


// SetDisplayVersion sets the value of the displayVersion property.
// The localized extension display version.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayversion
func (w_ WebExtension) SetDisplayVersion(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayVersion:"), objc.String(value))
}

// An array of all errors that occurred during the processing of the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/errors
func (w_ WebExtension) Errors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("errors"))
	return rv
}


// SetErrors sets the value of the errors property.
// An array of all errors that occurred during the processing of the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/errors
func (w_ WebExtension) SetErrors(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setErrors:"), value)
}

// A Boolean value indicating whether the extension has background content that can run when needed.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasbackgroundcontent
func (w_ WebExtension) HasBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasBackgroundContent"))
	return rv
}


// SetHasBackgroundContent sets the value of the hasBackgroundContent property.
// A Boolean value indicating whether the extension has background content that can run when needed.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasbackgroundcontent
func (w_ WebExtension) SetHasBackgroundContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasBackgroundContent:"), value)
}

// A Boolean value indicating whether the extension includes commands that users can invoke.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascommands
func (w_ WebExtension) HasCommands() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasCommands"))
	return rv
}


// SetHasCommands sets the value of the hasCommands property.
// A Boolean value indicating whether the extension includes commands that users can invoke.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascommands
func (w_ WebExtension) SetHasCommands(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasCommands:"), value)
}

// A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascontentmodificationrules
func (w_ WebExtension) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
}


// SetHasContentModificationRules sets the value of the hasContentModificationRules property.
// A Boolean value indicating whether the extension includes rules used for content modification or blocking.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascontentmodificationrules
func (w_ WebExtension) SetHasContentModificationRules(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasContentModificationRules:"), value)
}

// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasinjectedcontent
func (w_ WebExtension) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
}


// SetHasInjectedContent sets the value of the hasInjectedContent property.
// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasinjectedcontent
func (w_ WebExtension) SetHasInjectedContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasInjectedContent:"), value)
}

// A Boolean value indicating whether the extension has an options page.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoptionspage
func (w_ WebExtension) HasOptionsPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOptionsPage"))
	return rv
}


// SetHasOptionsPage sets the value of the hasOptionsPage property.
// A Boolean value indicating whether the extension has an options page.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoptionspage
func (w_ WebExtension) SetHasOptionsPage(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasOptionsPage:"), value)
}

// A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoverridenewtabpage
func (w_ WebExtension) HasOverrideNewTabPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOverrideNewTabPage"))
	return rv
}


// SetHasOverrideNewTabPage sets the value of the hasOverrideNewTabPage property.
// A Boolean value indicating whether the extension provides an alternative to the default new tab page.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoverridenewtabpage
func (w_ WebExtension) SetHasOverrideNewTabPage(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasOverrideNewTabPage:"), value)
}

// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/haspersistentbackgroundcontent
func (w_ WebExtension) HasPersistentBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasPersistentBackgroundContent"))
	return rv
}


// SetHasPersistentBackgroundContent sets the value of the hasPersistentBackgroundContent property.
// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/haspersistentbackgroundcontent
func (w_ WebExtension) SetHasPersistentBackgroundContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasPersistentBackgroundContent:"), value)
}

// The parsed manifest as a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifest
func (w_ WebExtension) Manifest() string {
	rv := objc.Send[string](w_.ID, objc.Sel("manifest"))
	return rv
}


// SetManifest sets the value of the manifest property.
// The parsed manifest as a dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifest
func (w_ WebExtension) SetManifest(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setManifest:"), objc.String(value))
}

// The parsed manifest version, or
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifestversion
func (w_ WebExtension) ManifestVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("manifestVersion"))
	return rv
}


// SetManifestVersion sets the value of the manifestVersion property.
// The parsed manifest version, or

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifestversion
func (w_ WebExtension) SetManifestVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setManifestVersion:"), value)
}

// The set of websites that the extension may need access to for optional functionality.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissionmatchpatterns
func (w_ WebExtension) OptionalPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("optionalPermissionMatchPatterns"))
	return rv
}


// SetOptionalPermissionMatchPatterns sets the value of the optionalPermissionMatchPatterns property.
// The set of websites that the extension may need access to for optional functionality.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissionmatchpatterns
func (w_ WebExtension) SetOptionalPermissionMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionalPermissionMatchPatterns:"), value)
}

// The set of permissions that the extension may need for optional functionality.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissions
func (w_ WebExtension) OptionalPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("optionalPermissions"))
	return rv
}


// SetOptionalPermissions sets the value of the optionalPermissions property.
// The set of permissions that the extension may need for optional functionality.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissions
func (w_ WebExtension) SetOptionalPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionalPermissions:"), value)
}

// The set of websites that the extension requires access to for its base functionality.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissionmatchpatterns
func (w_ WebExtension) RequestedPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("requestedPermissionMatchPatterns"))
	return rv
}


// SetRequestedPermissionMatchPatterns sets the value of the requestedPermissionMatchPatterns property.
// The set of websites that the extension requires access to for its base functionality.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissionmatchpatterns
func (w_ WebExtension) SetRequestedPermissionMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequestedPermissionMatchPatterns:"), value)
}

// The set of permissions that the extension requires for its base functionality.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissions
func (w_ WebExtension) RequestedPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("requestedPermissions"))
	return rv
}


// SetRequestedPermissions sets the value of the requestedPermissions property.
// The set of permissions that the extension requires for its base functionality.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissions
func (w_ WebExtension) SetRequestedPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequestedPermissions:"), value)
}

// The extension version.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/version
func (w_ WebExtension) Version() string {
	rv := objc.Send[string](w_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The extension version.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/version
func (w_ WebExtension) SetVersion(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setVersion:"), objc.String(value))
}



