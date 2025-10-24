// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AllRequestedMatchPatterns() WebExtensionMatchPattern /* not a class type */
	SetAllRequestedMatchPatterns(value WebExtensionMatchPattern /* not a class type */)
	DefaultLocale() objc.IObject /* cross-framework: Locale */
	SetDefaultLocale(value objc.IObject /* cross-framework: Locale */)
	DisplayActionLabel() objc.IObject /* cross-framework: NSString */
	SetDisplayActionLabel(value objc.IObject /* cross-framework: NSString */)
	DisplayDescription() objc.IObject /* cross-framework: NSString */
	SetDisplayDescription(value objc.IObject /* cross-framework: NSString */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	DisplayShortName() objc.IObject /* cross-framework: NSString */
	SetDisplayShortName(value objc.IObject /* cross-framework: NSString */)
	DisplayVersion() objc.IObject /* cross-framework: NSString */
	SetDisplayVersion(value objc.IObject /* cross-framework: NSString */)
	Errors() objc.IObject /* cross-framework: Error */
	SetErrors(value objc.IObject /* cross-framework: Error */)
	HasBackgroundContent() bool
	SetHasBackgroundContent(value bool)
	HasCommands() bool
	SetHasCommands(value bool)
	HasContentModificationRules() bool
	SetHasContentModificationRules(value bool)
	HasInjectedContent() bool
	SetHasInjectedContent(value bool)
	HasOptionsPage() bool
	SetHasOptionsPage(value bool)
	HasOverrideNewTabPage() bool
	SetHasOverrideNewTabPage(value bool)
	HasPersistentBackgroundContent() bool
	SetHasPersistentBackgroundContent(value bool)
	Manifest() objc.IObject /* cross-framework: NSString */
	SetManifest(value objc.IObject /* cross-framework: NSString */)
	ManifestVersion() float64
	SetManifestVersion(value float64)
	OptionalPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */
	SetOptionalPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */)
	OptionalPermissions() unsafe.Pointer
	SetOptionalPermissions(value unsafe.Pointer)
	RequestedPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */
	SetRequestedPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */)
	RequestedPermissions() unsafe.Pointer
	SetRequestedPermissions(value unsafe.Pointer)
	Version() objc.IObject /* cross-framework: NSString */
	SetVersion(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/allrequestedmatchpatterns
func (w_ WebExtension) AllRequestedMatchPatterns() WebExtensionMatchPattern /* not a class type */ {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("allRequestedMatchPatterns"))
	return rv
}


// The set of websites that the extension requires access to for injected content and for receiving messages from websites.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/allrequestedmatchpatterns
func (w_ WebExtension) SetAllRequestedMatchPatterns(value WebExtensionMatchPattern /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllRequestedMatchPatterns:"), value)
}


// The default locale for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/defaultlocale
func (w_ WebExtension) DefaultLocale() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](w_.ID, objc.Sel("defaultLocale"))
	return rv
}


// The default locale for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/defaultlocale
func (w_ WebExtension) SetDefaultLocale(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultLocale:"), value)
}


// The default localized extension action label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayactionlabel
func (w_ WebExtension) DisplayActionLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayActionLabel"))
	return rv
}


// The default localized extension action label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayactionlabel
func (w_ WebExtension) SetDisplayActionLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayActionLabel:"), value)
}


// The localized extension description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displaydescription
func (w_ WebExtension) DisplayDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayDescription"))
	return rv
}


// The localized extension description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displaydescription
func (w_ WebExtension) SetDisplayDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayDescription:"), value)
}


// The localized extension name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayname
func (w_ WebExtension) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayName"))
	return rv
}


// The localized extension name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayname
func (w_ WebExtension) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayName:"), value)
}


// The localized extension short name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayshortname
func (w_ WebExtension) DisplayShortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayShortName"))
	return rv
}


// The localized extension short name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayshortname
func (w_ WebExtension) SetDisplayShortName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayShortName:"), value)
}


// The localized extension display version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayversion
func (w_ WebExtension) DisplayVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayVersion"))
	return rv
}


// The localized extension display version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/displayversion
func (w_ WebExtension) SetDisplayVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayVersion:"), value)
}


// An array of all errors that occurred during the processing of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/errors
func (w_ WebExtension) Errors() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](w_.ID, objc.Sel("errors"))
	return rv
}


// An array of all errors that occurred during the processing of the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/errors
func (w_ WebExtension) SetErrors(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setErrors:"), value)
}


// A Boolean value indicating whether the extension has background content that can run when needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasbackgroundcontent
func (w_ WebExtension) HasBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasBackgroundContent"))
	return rv
}


// A Boolean value indicating whether the extension has background content that can run when needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasbackgroundcontent
func (w_ WebExtension) SetHasBackgroundContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasBackgroundContent:"), value)
}


// A Boolean value indicating whether the extension includes commands that users can invoke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascommands
func (w_ WebExtension) HasCommands() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasCommands"))
	return rv
}


// A Boolean value indicating whether the extension includes commands that users can invoke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascommands
func (w_ WebExtension) SetHasCommands(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasCommands:"), value)
}


// A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascontentmodificationrules
func (w_ WebExtension) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
}


// A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hascontentmodificationrules
func (w_ WebExtension) SetHasContentModificationRules(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasContentModificationRules:"), value)
}


// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasinjectedcontent
func (w_ WebExtension) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
}


// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasinjectedcontent
func (w_ WebExtension) SetHasInjectedContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasInjectedContent:"), value)
}


// A Boolean value indicating whether the extension has an options page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoptionspage
func (w_ WebExtension) HasOptionsPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOptionsPage"))
	return rv
}


// A Boolean value indicating whether the extension has an options page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoptionspage
func (w_ WebExtension) SetHasOptionsPage(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasOptionsPage:"), value)
}


// A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoverridenewtabpage
func (w_ WebExtension) HasOverrideNewTabPage() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasOverrideNewTabPage"))
	return rv
}


// A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/hasoverridenewtabpage
func (w_ WebExtension) SetHasOverrideNewTabPage(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasOverrideNewTabPage:"), value)
}


// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/haspersistentbackgroundcontent
func (w_ WebExtension) HasPersistentBackgroundContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasPersistentBackgroundContent"))
	return rv
}


// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/haspersistentbackgroundcontent
func (w_ WebExtension) SetHasPersistentBackgroundContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasPersistentBackgroundContent:"), value)
}


// The parsed manifest as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifest
func (w_ WebExtension) Manifest() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("manifest"))
	return rv
}


// The parsed manifest as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifest
func (w_ WebExtension) SetManifest(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setManifest:"), value)
}


// The parsed manifest version, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifestversion
func (w_ WebExtension) ManifestVersion() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("manifestVersion"))
	return rv
}


// The parsed manifest version, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/manifestversion
func (w_ WebExtension) SetManifestVersion(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setManifestVersion:"), value)
}


// The set of websites that the extension may need access to for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissionmatchpatterns
func (w_ WebExtension) OptionalPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */ {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("optionalPermissionMatchPatterns"))
	return rv
}


// The set of websites that the extension may need access to for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissionmatchpatterns
func (w_ WebExtension) SetOptionalPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionalPermissionMatchPatterns:"), value)
}


// The set of permissions that the extension may need for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissions
func (w_ WebExtension) OptionalPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("optionalPermissions"))
	return rv
}


// The set of permissions that the extension may need for optional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/optionalpermissions
func (w_ WebExtension) SetOptionalPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionalPermissions:"), value)
}


// The set of websites that the extension requires access to for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissionmatchpatterns
func (w_ WebExtension) RequestedPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */ {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("requestedPermissionMatchPatterns"))
	return rv
}


// The set of websites that the extension requires access to for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissionmatchpatterns
func (w_ WebExtension) SetRequestedPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequestedPermissionMatchPatterns:"), value)
}


// The set of permissions that the extension requires for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissions
func (w_ WebExtension) RequestedPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("requestedPermissions"))
	return rv
}


// The set of permissions that the extension requires for its base functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/requestedpermissions
func (w_ WebExtension) SetRequestedPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequestedPermissions:"), value)
}


// The extension version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/version
func (w_ WebExtension) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("version"))
	return rv
}


// The extension version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/version
func (w_ WebExtension) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setVersion:"), value)
}



