// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that represents the runtime environment for a web extension.
//
// This class provides methods for managing the extension’s permissions, allowing it to inject content, run background logic, show popovers, and display other web-based UI to the user.
//
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
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/baseurl
func (w_ WebExtensionContext) BaseURL() foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("baseURL"))
	return rv
}


// SetBaseURL sets the value of the baseURL property.
// The base URL the context uses for loading extension resources or injecting content into webpages.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/baseurl
func (w_ WebExtensionContext) SetBaseURL(value foundation.URL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBaseURL:"), value)
}

// The commands associated with the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/commands
func (w_ WebExtensionContext) Commands() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("commands"))
	return rv
}


// SetCommands sets the value of the commands property.
// The commands associated with the extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/commands
func (w_ WebExtensionContext) SetCommands(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCommands:"), value)
}

// The currently granted permission match patterns that have not expired.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissionmatchpatterns
func (w_ WebExtensionContext) CurrentPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentPermissionMatchPatterns"))
	return rv
}


// SetCurrentPermissionMatchPatterns sets the value of the currentPermissionMatchPatterns property.
// The currently granted permission match patterns that have not expired.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissionmatchpatterns
func (w_ WebExtensionContext) SetCurrentPermissionMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCurrentPermissionMatchPatterns:"), value)
}

// The currently granted permissions that have not expired.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissions
func (w_ WebExtensionContext) CurrentPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentPermissions"))
	return rv
}


// SetCurrentPermissions sets the value of the currentPermissions property.
// The currently granted permissions that have not expired.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/currentpermissions
func (w_ WebExtensionContext) SetCurrentPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCurrentPermissions:"), value)
}

// The currently denied permission match patterns and their expiration dates.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissionmatchpatterns
func (w_ WebExtensionContext) DeniedPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("deniedPermissionMatchPatterns"))
	return rv
}


// SetDeniedPermissionMatchPatterns sets the value of the deniedPermissionMatchPatterns property.
// The currently denied permission match patterns and their expiration dates.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissionmatchpatterns
func (w_ WebExtensionContext) SetDeniedPermissionMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissionMatchPatterns:"), value)
}

// The currently denied permissions and their expiration dates.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissions
func (w_ WebExtensionContext) DeniedPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("deniedPermissions"))
	return rv
}


// SetDeniedPermissions sets the value of the deniedPermissions property.
// The currently denied permissions and their expiration dates.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/deniedpermissions
func (w_ WebExtensionContext) SetDeniedPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissions:"), value)
}

// All errors that occurred in the extension context.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/errors
func (w_ WebExtensionContext) Errors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("errors"))
	return rv
}


// SetErrors sets the value of the errors property.
// All errors that occurred in the extension context.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/errors
func (w_ WebExtensionContext) SetErrors(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setErrors:"), value)
}

// The window that currently has focus for this extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/focusedwindow
func (w_ WebExtensionContext) FocusedWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("focusedWindow"))
	return rv
}


// SetFocusedWindow sets the value of the focusedWindow property.
// The window that currently has focus for this extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/focusedwindow
func (w_ WebExtensionContext) SetFocusedWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFocusedWindow:"), value)
}

// The currently granted permission match patterns and their expiration dates.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissionmatchpatterns
func (w_ WebExtensionContext) GrantedPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("grantedPermissionMatchPatterns"))
	return rv
}


// SetGrantedPermissionMatchPatterns sets the value of the grantedPermissionMatchPatterns property.
// The currently granted permission match patterns and their expiration dates.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissionmatchpatterns
func (w_ WebExtensionContext) SetGrantedPermissionMatchPatterns(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissionMatchPatterns:"), value)
}

// The currently granted permissions and their expiration dates.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissions
func (w_ WebExtensionContext) GrantedPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("grantedPermissions"))
	return rv
}


// SetGrantedPermissions sets the value of the grantedPermissions property.
// The currently granted permissions and their expiration dates.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/grantedpermissions
func (w_ WebExtensionContext) SetGrantedPermissions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissions:"), value)
}

// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallhosts
func (w_ WebExtensionContext) HasAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllHosts"))
	return rv
}


// SetHasAccessToAllHosts sets the value of the hasAccessToAllHosts property.
// A Boolean value indicating if the currently granted permission match patterns set contains the

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallhosts
func (w_ WebExtensionContext) SetHasAccessToAllHosts(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToAllHosts:"), value)
}

// A Boolean value indicating if the currently granted permission match patterns set contains the
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallurls
func (w_ WebExtensionContext) HasAccessToAllURLs() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllURLs"))
	return rv
}


// SetHasAccessToAllURLs sets the value of the hasAccessToAllURLs property.
// A Boolean value indicating if the currently granted permission match patterns set contains the

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoallurls
func (w_ WebExtensionContext) SetHasAccessToAllURLs(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToAllURLs:"), value)
}

// A Boolean value indicating if the extension has access to private data.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoprivatedata
func (w_ WebExtensionContext) HasAccessToPrivateData() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToPrivateData"))
	return rv
}


// SetHasAccessToPrivateData sets the value of the hasAccessToPrivateData property.
// A Boolean value indicating if the extension has access to private data.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasaccesstoprivatedata
func (w_ WebExtensionContext) SetHasAccessToPrivateData(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToPrivateData:"), value)
}

// A boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hascontentmodificationrules
func (w_ WebExtensionContext) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
}


// SetHasContentModificationRules sets the value of the hasContentModificationRules property.
// A boolean value indicating whether the extension includes rules used for content modification or blocking.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hascontentmodificationrules
func (w_ WebExtensionContext) SetHasContentModificationRules(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasContentModificationRules:"), value)
}

// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasinjectedcontent
func (w_ WebExtensionContext) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
}


// SetHasInjectedContent sets the value of the hasInjectedContent property.
// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasinjectedcontent
func (w_ WebExtensionContext) SetHasInjectedContent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasInjectedContent:"), value)
}

// A Boolean value indicating if the extension has requested optional access to all hosts.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasrequestedoptionalaccesstoallhosts
func (w_ WebExtensionContext) HasRequestedOptionalAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasRequestedOptionalAccessToAllHosts"))
	return rv
}


// SetHasRequestedOptionalAccessToAllHosts sets the value of the hasRequestedOptionalAccessToAllHosts property.
// A Boolean value indicating if the extension has requested optional access to all hosts.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/hasrequestedoptionalaccesstoallhosts
func (w_ WebExtensionContext) SetHasRequestedOptionalAccessToAllHosts(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasRequestedOptionalAccessToAllHosts:"), value)
}

// The name shown when inspecting the background web view.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/inspectionname
func (w_ WebExtensionContext) InspectionName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("inspectionName"))
	return rv
}


// SetInspectionName sets the value of the inspectionName property.
// The name shown when inspecting the background web view.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/inspectionname
func (w_ WebExtensionContext) SetInspectionName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectionName:"), objc.String(value))
}

// Determines whether Web Inspector can inspect the
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
}


// SetIsInspectable sets the value of the isInspectable property.
// Determines whether Web Inspector can inspect the

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
}

// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) IsLoaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoaded"))
	return rv
}


// SetIsLoaded sets the value of the isLoaded property.
// A Boolean value indicating if this context is loaded in an extension controller.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) SetIsLoaded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoaded:"), value)
}

// A set of open tabs in all open windows that are exposed to this extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/opentabs
func (w_ WebExtensionContext) OpenTabs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("openTabs"))
	return rv
}


// SetOpenTabs sets the value of the openTabs property.
// A set of open tabs in all open windows that are exposed to this extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/opentabs
func (w_ WebExtensionContext) SetOpenTabs(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpenTabs:"), value)
}

// The open windows that are exposed to this extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/openwindows
func (w_ WebExtensionContext) OpenWindows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("openWindows"))
	return rv
}


// SetOpenWindows sets the value of the openWindows property.
// The open windows that are exposed to this extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/openwindows
func (w_ WebExtensionContext) SetOpenWindows(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpenWindows:"), value)
}

// The URL of the extension’s options page, if the extension has one.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/optionspageurl
func (w_ WebExtensionContext) OptionsPageURL() foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("optionsPageURL"))
	return rv
}


// SetOptionsPageURL sets the value of the optionsPageURL property.
// The URL of the extension’s options page, if the extension has one.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/optionspageurl
func (w_ WebExtensionContext) SetOptionsPageURL(value foundation.URL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOptionsPageURL:"), value)
}

// The URL to use as an alternative to the default new tab page, if the extension has one.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/overridenewtabpageurl
func (w_ WebExtensionContext) OverrideNewTabPageURL() foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("overrideNewTabPageURL"))
	return rv
}


// SetOverrideNewTabPageURL sets the value of the overrideNewTabPageURL property.
// The URL to use as an alternative to the default new tab page, if the extension has one.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/overridenewtabpageurl
func (w_ WebExtensionContext) SetOverrideNewTabPageURL(value foundation.URL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOverrideNewTabPageURL:"), value)
}

// A unique identifier used to distinguish the extension from other extensions and target it for messages.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/uniqueidentifier
func (w_ WebExtensionContext) UniqueIdentifier() string {
	rv := objc.Send[string](w_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// SetUniqueIdentifier sets the value of the uniqueIdentifier property.
// A unique identifier used to distinguish the extension from other extensions and target it for messages.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/uniqueidentifier
func (w_ WebExtensionContext) SetUniqueIdentifier(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUniqueIdentifier:"), objc.String(value))
}

// Specifies unsupported APIs for this extension, making them
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/unsupportedapis
func (w_ WebExtensionContext) UnsupportedAPIs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("unsupportedAPIs"))
	return rv
}


// SetUnsupportedAPIs sets the value of the unsupportedAPIs property.
// Specifies unsupported APIs for this extension, making them

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/unsupportedapis
func (w_ WebExtensionContext) SetUnsupportedAPIs(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnsupportedAPIs:"), value)
}

// The extension this context represents.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextension
func (w_ WebExtensionContext) WebExtension() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("webExtension"))
	return rv
}


// SetWebExtension sets the value of the webExtension property.
// The extension this context represents.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextension
func (w_ WebExtensionContext) SetWebExtension(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtension:"), value)
}

// The extension controller this context is loaded in, otherwise
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextensioncontroller
func (w_ WebExtensionContext) WebExtensionController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("webExtensionController"))
	return rv
}


// SetWebExtensionController sets the value of the webExtensionController property.
// The extension controller this context is loaded in, otherwise

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webextensioncontroller
func (w_ WebExtensionContext) SetWebExtensionController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}

// The web view configuration to use for web views that load pages from this extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webviewconfiguration
func (w_ WebExtensionContext) WebViewConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("webViewConfiguration"))
	return rv
}


// SetWebViewConfiguration sets the value of the webViewConfiguration property.
// The web view configuration to use for web views that load pages from this extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/webviewconfiguration
func (w_ WebExtensionContext) SetWebViewConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebViewConfiguration:"), value)
}



