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

// The class instance for the [WebExtensionContext] class.
var (
	WebExtensionContextClass     _WebExtensionContextClass
	WebExtensionContextClassOnce sync.Once
)

func getWebExtensionContextClass() _WebExtensionContextClass {
	WebExtensionContextClassOnce.Do(func() {
		WebExtensionContextClass = _WebExtensionContextClass{objc.GetClass("WKWebExtensionContext")}
	})
	return WebExtensionContextClass
}

type _WebExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [WebExtensionContext] class.
type IWebExtensionContext interface {
	objectivec.IObject
	// properties:
	BaseURL() objc.IObject /* cross-framework: URL */
	SetBaseURL(value objc.IObject /* cross-framework: URL */)
	Commands() WebExtensionCommand /* not a class type */
	SetCommands(value WebExtensionCommand /* not a class type */)
	CurrentPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */
	SetCurrentPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */)
	CurrentPermissions() unsafe.Pointer
	SetCurrentPermissions(value unsafe.Pointer)
	DeniedPermissionMatchPatterns() objc.IObject /* cross-framework: Date */
	SetDeniedPermissionMatchPatterns(value objc.IObject /* cross-framework: Date */)
	DeniedPermissions() objc.IObject /* cross-framework: Date */
	SetDeniedPermissions(value objc.IObject /* cross-framework: Date */)
	Errors() objc.IObject /* cross-framework: Error */
	SetErrors(value objc.IObject /* cross-framework: Error */)
	FocusedWindow() WebExtensionWindow /* not a class type */
	SetFocusedWindow(value WebExtensionWindow /* not a class type */)
	GrantedPermissionMatchPatterns() objc.IObject /* cross-framework: Date */
	SetGrantedPermissionMatchPatterns(value objc.IObject /* cross-framework: Date */)
	GrantedPermissions() objc.IObject /* cross-framework: Date */
	SetGrantedPermissions(value objc.IObject /* cross-framework: Date */)
	HasAccessToAllHosts() bool
	SetHasAccessToAllHosts(value bool)
	HasAccessToAllURLs() bool
	SetHasAccessToAllURLs(value bool)
	HasAccessToPrivateData() bool
	SetHasAccessToPrivateData(value bool)
	HasContentModificationRules() bool
	SetHasContentModificationRules(value bool)
	HasInjectedContent() bool
	SetHasInjectedContent(value bool)
	HasRequestedOptionalAccessToAllHosts() bool
	SetHasRequestedOptionalAccessToAllHosts(value bool)
	InspectionName() objc.IObject /* cross-framework: NSString */
	SetInspectionName(value objc.IObject /* cross-framework: NSString */)
	IsInspectable() bool
	SetIsInspectable(value bool)
	IsLoaded() bool
	SetIsLoaded(value bool)
	OpenTabs() unsafe.Pointer
	SetOpenTabs(value unsafe.Pointer)
	OpenWindows() WebExtensionWindow /* not a class type */
	SetOpenWindows(value WebExtensionWindow /* not a class type */)
	OptionsPageURL() objc.IObject /* cross-framework: URL */
	SetOptionsPageURL(value objc.IObject /* cross-framework: URL */)
	OverrideNewTabPageURL() objc.IObject /* cross-framework: URL */
	SetOverrideNewTabPageURL(value objc.IObject /* cross-framework: URL */)
	UniqueIdentifier() objc.IObject /* cross-framework: NSString */
	SetUniqueIdentifier(value objc.IObject /* cross-framework: NSString */)
	UnsupportedAPIs() objc.IObject /* cross-framework: NSString */
	SetUnsupportedAPIs(value objc.IObject /* cross-framework: NSString */)
	WebExtension() IWKWebExtension
	SetWebExtension(value IWKWebExtension)
	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
	WebViewConfiguration() IWKWebViewConfiguration
	SetWebViewConfiguration(value IWKWebViewConfiguration)
	// methods:
}

// An object that represents the runtime environment for a web extension.
//
// This class provides methods for managing the extension’s permissions, allowing it to inject content, run background logic, show popovers, and display other web-based UI to the user.


// An object that represents the runtime environment for a web extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext
type WebExtensionContext struct {
	objectivec.Object
}

// WebExtensionContextFrom constructs a [WebExtensionContext] from an unsafe.Pointer.
//
// An object that represents the runtime environment for a web extension.
func WebExtensionContextFrom(ptr unsafe.Pointer) WebExtensionContext {
	return WebExtensionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionContextClass) Alloc() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebExtensionContextClass) New() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionContext) Init() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionContext) Autorelease() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionContext creates a new WebExtensionContext instance.
func NewWebExtensionContext() WebExtensionContext {
	return getWebExtensionContextClass().New()
}



// The base URL the context uses for loading extension resources or injecting content into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/baseurl
func (w_ WebExtensionContext) BaseURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("baseURL"))
	return rv
}


// The base URL the context uses for loading extension resources or injecting content into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/baseurl
func (w_ WebExtensionContext) SetBaseURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBaseURL:"), value)
}


// The commands associated with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/commands
func (w_ WebExtensionContext) Commands() WebExtensionCommand /* not a class type */ {
	rv := objc.Send[WebExtensionCommand](w_.ID, objc.Sel("commands"))
	return rv
}


// The commands associated with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/commands
func (w_ WebExtensionContext) SetCommands(value WebExtensionCommand /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCommands:"), value)
}


// The currently granted permission match patterns that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissionmatchpatterns
func (w_ WebExtensionContext) CurrentPermissionMatchPatterns() WebExtensionMatchPattern /* not a class type */ {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("currentPermissionMatchPatterns"))
	return rv
}


// The currently granted permission match patterns that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissionmatchpatterns
func (w_ WebExtensionContext) SetCurrentPermissionMatchPatterns(value WebExtensionMatchPattern /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCurrentPermissionMatchPatterns:"), value)
}


// The currently granted permissions that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissions
func (w_ WebExtensionContext) CurrentPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentPermissions"))
	return rv
}


// The currently granted permissions that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissions
func (w_ WebExtensionContext) SetCurrentPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCurrentPermissions:"), value)
}


// The currently denied permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissionmatchpatterns
func (w_ WebExtensionContext) DeniedPermissionMatchPatterns() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](w_.ID, objc.Sel("deniedPermissionMatchPatterns"))
	return rv
}


// The currently denied permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissionmatchpatterns
func (w_ WebExtensionContext) SetDeniedPermissionMatchPatterns(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissionMatchPatterns:"), value)
}


// The currently denied permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissions
func (w_ WebExtensionContext) DeniedPermissions() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](w_.ID, objc.Sel("deniedPermissions"))
	return rv
}


// The currently denied permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissions
func (w_ WebExtensionContext) SetDeniedPermissions(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissions:"), value)
}


// All errors that occurred in the extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/errors
func (w_ WebExtensionContext) Errors() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](w_.ID, objc.Sel("errors"))
	return rv
}


// All errors that occurred in the extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/errors
func (w_ WebExtensionContext) SetErrors(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setErrors:"), value)
}


// The window that currently has focus for this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/focusedwindow
func (w_ WebExtensionContext) FocusedWindow() WebExtensionWindow /* not a class type */ {
	rv := objc.Send[WebExtensionWindow](w_.ID, objc.Sel("focusedWindow"))
	return rv
}


// The window that currently has focus for this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/focusedwindow
func (w_ WebExtensionContext) SetFocusedWindow(value WebExtensionWindow /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFocusedWindow:"), value)
}


// The currently granted permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissionmatchpatterns
func (w_ WebExtensionContext) GrantedPermissionMatchPatterns() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](w_.ID, objc.Sel("grantedPermissionMatchPatterns"))
	return rv
}


// The currently granted permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissionmatchpatterns
func (w_ WebExtensionContext) SetGrantedPermissionMatchPatterns(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissionMatchPatterns:"), value)
}


// The currently granted permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissions
func (w_ WebExtensionContext) GrantedPermissions() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](w_.ID, objc.Sel("grantedPermissions"))
	return rv
}


// The currently granted permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissions
func (w_ WebExtensionContext) SetGrantedPermissions(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissions:"), value)
}


// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallhosts
func (w_ WebExtensionContext) HasAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllHosts"))
	return rv
}


// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallhosts
func (w_ WebExtensionContext) SetHasAccessToAllHosts(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToAllHosts:"), value)
}


// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallurls
func (w_ WebExtensionContext) HasAccessToAllURLs() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllURLs"))
	return rv
}


// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallurls
func (w_ WebExtensionContext) SetHasAccessToAllURLs(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToAllURLs:"), value)
}


// A Boolean value indicating if the extension has access to private data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoprivatedata
func (w_ WebExtensionContext) HasAccessToPrivateData() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToPrivateData"))
	return rv
}


// A Boolean value indicating if the extension has access to private data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoprivatedata
func (w_ WebExtensionContext) SetHasAccessToPrivateData(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToPrivateData:"), value)
}


// A boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hascontentmodificationrules
func (w_ WebExtensionContext) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
}


// A boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hascontentmodificationrules
func (w_ WebExtensionContext) SetHasContentModificationRules(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasContentModificationRules:"), value)
}


// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasinjectedcontent
func (w_ WebExtensionContext) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
}


// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasinjectedcontent
func (w_ WebExtensionContext) SetHasInjectedContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasInjectedContent:"), value)
}


// A Boolean value indicating if the extension has requested optional access to all hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasrequestedoptionalaccesstoallhosts
func (w_ WebExtensionContext) HasRequestedOptionalAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasRequestedOptionalAccessToAllHosts"))
	return rv
}


// A Boolean value indicating if the extension has requested optional access to all hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasrequestedoptionalaccesstoallhosts
func (w_ WebExtensionContext) SetHasRequestedOptionalAccessToAllHosts(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasRequestedOptionalAccessToAllHosts:"), value)
}


// The name shown when inspecting the background web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/inspectionname
func (w_ WebExtensionContext) InspectionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("inspectionName"))
	return rv
}


// The name shown when inspecting the background web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/inspectionname
func (w_ WebExtensionContext) SetInspectionName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectionName:"), value)
}


// Determines whether Web Inspector can inspect the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
}


// Determines whether Web Inspector can inspect the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
}


// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) IsLoaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoaded"))
	return rv
}


// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) SetIsLoaded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoaded:"), value)
}


// A set of open tabs in all open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/opentabs
func (w_ WebExtensionContext) OpenTabs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("openTabs"))
	return rv
}


// A set of open tabs in all open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/opentabs
func (w_ WebExtensionContext) SetOpenTabs(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpenTabs:"), value)
}


// The open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/openwindows
func (w_ WebExtensionContext) OpenWindows() WebExtensionWindow /* not a class type */ {
	rv := objc.Send[WebExtensionWindow](w_.ID, objc.Sel("openWindows"))
	return rv
}


// The open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/openwindows
func (w_ WebExtensionContext) SetOpenWindows(value WebExtensionWindow /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpenWindows:"), value)
}


// The URL of the extension’s options page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/optionspageurl
func (w_ WebExtensionContext) OptionsPageURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("optionsPageURL"))
	return rv
}


// The URL of the extension’s options page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/optionspageurl
func (w_ WebExtensionContext) SetOptionsPageURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionsPageURL:"), value)
}


// The URL to use as an alternative to the default new tab page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/overridenewtabpageurl
func (w_ WebExtensionContext) OverrideNewTabPageURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("overrideNewTabPageURL"))
	return rv
}


// The URL to use as an alternative to the default new tab page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/overridenewtabpageurl
func (w_ WebExtensionContext) SetOverrideNewTabPageURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOverrideNewTabPageURL:"), value)
}


// A unique identifier used to distinguish the extension from other extensions and target it for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/uniqueidentifier
func (w_ WebExtensionContext) UniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// A unique identifier used to distinguish the extension from other extensions and target it for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/uniqueidentifier
func (w_ WebExtensionContext) SetUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUniqueIdentifier:"), value)
}


// Specifies unsupported APIs for this extension, making them
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/unsupportedapis
func (w_ WebExtensionContext) UnsupportedAPIs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("unsupportedAPIs"))
	return rv
}


// Specifies unsupported APIs for this extension, making them
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/unsupportedapis
func (w_ WebExtensionContext) SetUnsupportedAPIs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnsupportedAPIs:"), value)
}


// The extension this context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextension
func (w_ WebExtensionContext) WebExtension() IWKWebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("webExtension"))
	return rv
}


// The extension this context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextension
func (w_ WebExtensionContext) SetWebExtension(value IWKWebExtension) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtension:"), value)
}


// The extension controller this context is loaded in, otherwise
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextensioncontroller
func (w_ WebExtensionContext) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("webExtensionController"))
	return rv
}


// The extension controller this context is loaded in, otherwise
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextensioncontroller
func (w_ WebExtensionContext) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}


// The web view configuration to use for web views that load pages from this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webviewconfiguration
func (w_ WebExtensionContext) WebViewConfiguration() IWKWebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("webViewConfiguration"))
	return rv
}


// The web view configuration to use for web views that load pages from this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webviewconfiguration
func (w_ WebExtensionContext) SetWebViewConfiguration(value IWKWebViewConfiguration) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebViewConfiguration:"), value)
}



