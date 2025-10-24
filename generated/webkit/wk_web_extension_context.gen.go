// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionContext */

/* debug [class_header]: Header for WKWebExtensionContext */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebExtensionContext */
// An interface definition for the [WebExtensionContext] class.
type IWebExtensionContext interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebExtensionContext */
	// properties:
	BaseURL() objc.IObject /* cross-framework: NSURL */
	SetBaseURL(value objc.IObject /* cross-framework: NSURL */)
	Commands() []WebExtensionCommand
	CurrentPermissionMatchPatterns() unsafe.Pointer
	CurrentPermissions() unsafe.Pointer
	DeniedPermissionMatchPatterns() foundation.IDictionary
	SetDeniedPermissionMatchPatterns(value foundation.IDictionary)
	DeniedPermissions() foundation.IDictionary
	SetDeniedPermissions(value foundation.IDictionary)
	Errors() []objc.IObject /* cross-framework: Error */
	FocusedWindow() unsafe.Pointer
	GrantedPermissionMatchPatterns() foundation.IDictionary
	SetGrantedPermissionMatchPatterns(value foundation.IDictionary)
	GrantedPermissions() foundation.IDictionary
	SetGrantedPermissions(value foundation.IDictionary)
	HasAccessToAllHosts() bool
	HasAccessToAllURLs() bool
	HasAccessToPrivateData() bool
	SetHasAccessToPrivateData(value bool)
	HasContentModificationRules() bool
	HasInjectedContent() bool
	HasRequestedOptionalAccessToAllHosts() bool
	SetHasRequestedOptionalAccessToAllHosts(value bool)
	InspectionName() objc.IObject /* cross-framework: NSString */
	SetInspectionName(value objc.IObject /* cross-framework: NSString */)
	Inspectable() bool
	SetInspectable(value bool)
	Loaded() bool
	OpenTabs() unsafe.Pointer
	OpenWindows() []objc.ID
	OptionsPageURL() objc.IObject        /* cross-framework: NSURL */
	OverrideNewTabPageURL() objc.IObject /* cross-framework: NSURL */
	UniqueIdentifier() objc.IObject      /* cross-framework: NSString */
	SetUniqueIdentifier(value objc.IObject /* cross-framework: NSString */)
	UnsupportedAPIs() unsafe.Pointer
	SetUnsupportedAPIs(value unsafe.Pointer)
	WebExtension() IWKWebExtension
	WebExtensionController() IWKWebExtensionController
	WebViewConfiguration() IWKWebViewConfiguration
	IsInspectable() bool
	SetIsInspectable(value bool)
	IsLoaded() bool
	SetIsLoaded(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebExtensionContext */
	// methods:
	ActionForTab(tab unsafe.Pointer) IWebExtensionAction
	ClearUserGestureInTab(tab unsafe.Pointer)
	CommandForEvent(event appkit.Event) IWebExtensionCommand
	DidChangeTabPropertiesForTab(properties WebExtensionTabChangedProperties, changedTab unsafe.Pointer)
	DidCloseWindow(closedWindow unsafe.Pointer)
	DidDeselectTabs(deselectedTabs []objc.ID)
	DidFocusWindow(focusedWindow unsafe.Pointer)
	DidOpenTab(newTab unsafe.Pointer)
	DidOpenWindow(newWindow unsafe.Pointer)
	DidReplaceTabWithTab(oldTab unsafe.Pointer, newTab unsafe.Pointer)
	DidSelectTabs(selectedTabs []objc.ID)
	HasAccessToURL(url objc.IObject /* cross-framework: NSURL */) bool
	HasAccessToURLInTab(url objc.IObject /* cross-framework: NSURL */, tab unsafe.Pointer) bool
	HasActiveUserGestureInTab(tab unsafe.Pointer) bool
	HasInjectedContentForURL(url objc.IObject /* cross-framework: NSURL */) bool
	HasPermission(permission WebExtensionPermission /* typedef */) bool
	HasPermissionInTab(permission WebExtensionPermission /* typedef */, tab unsafe.Pointer) bool
	LoadBackgroundContentWithCompletionHandler(completionHandler unsafe.Pointer)
	MenuItemsForTab(tab unsafe.Pointer) []appkit.MenuItem
	PerformActionForTab(tab unsafe.Pointer)
	PerformCommand(command IWKWebExtensionCommand)
	PerformCommandForEvent(event appkit.Event) bool
	PermissionStatusForPermission(permission WebExtensionPermission /* typedef */) WebExtensionContextPermissionStatus
	PermissionStatusForMatchPattern(pattern IWKWebExtensionMatchPattern) WebExtensionContextPermissionStatus
	PermissionStatusForURL(url objc.IObject /* cross-framework: NSURL */) WebExtensionContextPermissionStatus
	PermissionStatusForPermissionInTab(permission WebExtensionPermission /* typedef */, tab unsafe.Pointer) WebExtensionContextPermissionStatus
	PermissionStatusForURLInTab(url objc.IObject /* cross-framework: NSURL */, tab unsafe.Pointer) WebExtensionContextPermissionStatus
	PermissionStatusForMatchPatternInTab(pattern IWKWebExtensionMatchPattern, tab unsafe.Pointer) WebExtensionContextPermissionStatus
	SetPermissionStatusForPermission(status WebExtensionContextPermissionStatus, permission WebExtensionPermission /* typedef */)
	SetPermissionStatusForURL(status WebExtensionContextPermissionStatus, url objc.IObject /* cross-framework: NSURL */)
	SetPermissionStatusForMatchPattern(status WebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern)
	SetPermissionStatusForURLExpirationDate(status WebExtensionContextPermissionStatus, url objc.IObject /* cross-framework: NSURL */, expirationDate objc.IObject /* cross-framework: NSDate */)
	SetPermissionStatusForPermissionExpirationDate(status WebExtensionContextPermissionStatus, permission WebExtensionPermission /* typedef */, expirationDate objc.IObject /* cross-framework: NSDate */)
	SetPermissionStatusForMatchPatternExpirationDate(status WebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern, expirationDate objc.IObject /* cross-framework: NSDate */)
	UserGesturePerformedInTab(tab unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebExtensionContext */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionContextClass) Alloc() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebExtensionContext */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebExtensionContext */

// Returns a web extension context initialized with a specified extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/init(for:)
func NewWebExtensionContextForExtension(extension IWKWebExtension) WebExtensionContext {
	instance := getWebExtensionContextClass().Alloc()
	rv := objc.Send[WebExtensionContext](instance.ID, objc.Sel("initForExtension:"), extension)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewWebExtensionContextForExtension */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebExtensionContext */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebExtensionContext */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebExtensionContext */

// Retrieves the extension action for a given tab, or the default action if is passed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/action(for:)
func (w_ WebExtensionContext) ActionForTab(tab unsafe.Pointer) IWebExtensionAction {
	rv := objc.Send[WebExtensionAction](w_.ID, objc.Sel("actionForTab:"), tab)
	return rv
} /* debug [instance_methods/method]: ActionForTab */

// Called by the app to clear a user gesture in a specific tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/clearUserGesture(in:)
func (w_ WebExtensionContext) ClearUserGestureInTab(tab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("clearUserGestureInTab:"), tab)
} /* debug [instance_methods/method]: ClearUserGestureInTab */

// Retrieves the command associated with the given event without performing it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/command(for:)
func (w_ WebExtensionContext) CommandForEvent(event appkit.Event) IWebExtensionCommand {
	rv := objc.Send[WebExtensionCommand](w_.ID, objc.Sel("commandForEvent:"), event)
	return rv
} /* debug [instance_methods/method]: CommandForEvent */

// Called by the app when the properties of a tab are changed to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didChangeTabProperties(_:for:)
func (w_ WebExtensionContext) DidChangeTabPropertiesForTab(properties WebExtensionTabChangedProperties, changedTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didChangeTabProperties:forTab:"), properties, changedTab)
} /* debug [instance_methods/method]: DidChangeTabPropertiesForTab */

// Called by the app when a window is closed to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didCloseWindow(_:)
func (w_ WebExtensionContext) DidCloseWindow(closedWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didCloseWindow:"), closedWindow)
} /* debug [instance_methods/method]: DidCloseWindow */

// Called by the app when tabs are deselected to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didDeselectTabs(_:)
func (w_ WebExtensionContext) DidDeselectTabs(deselectedTabs []objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didDeselectTabs:"), deselectedTabs)
} /* debug [instance_methods/method]: DidDeselectTabs */

// Called by the app when a window gains focus to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didFocusWindow(_:)
func (w_ WebExtensionContext) DidFocusWindow(focusedWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didFocusWindow:"), focusedWindow)
} /* debug [instance_methods/method]: DidFocusWindow */

// Called by the app when a new tab is opened to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didOpenTab(_:)
func (w_ WebExtensionContext) DidOpenTab(newTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didOpenTab:"), newTab)
} /* debug [instance_methods/method]: DidOpenTab */

// Called by the app when a new window is opened to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didOpenWindow(_:)
func (w_ WebExtensionContext) DidOpenWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didOpenWindow:"), newWindow)
} /* debug [instance_methods/method]: DidOpenWindow */

// Called by the app when a tab is replaced by another tab to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didReplaceTab(_:with:)
func (w_ WebExtensionContext) DidReplaceTabWithTab(oldTab unsafe.Pointer, newTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didReplaceTab:withTab:"), oldTab, newTab)
} /* debug [instance_methods/method]: DidReplaceTabWithTab */

// Called by the app when tabs are selected to fire appropriate events with only this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didSelectTabs(_:)
func (w_ WebExtensionContext) DidSelectTabs(selectedTabs []objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didSelectTabs:"), selectedTabs)
} /* debug [instance_methods/method]: DidSelectTabs */

// Checks the specified URL against the currently granted permission match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccess(to:)
func (w_ WebExtensionContext) HasAccessToURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToURL:"), url)
	return rv
} /* debug [instance_methods/method]: HasAccessToURL */

// Checks the specified URL against the currently granted permission match patterns in a specific tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccess(to:in:)
func (w_ WebExtensionContext) HasAccessToURLInTab(url objc.IObject /* cross-framework: NSURL */, tab unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToURL:inTab:"), url, tab)
	return rv
} /* debug [instance_methods/method]: HasAccessToURLInTab */

// Indicates if a user gesture is currently active in the specified tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasActiveUserGesture(in:)
func (w_ WebExtensionContext) HasActiveUserGestureInTab(tab unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasActiveUserGestureInTab:"), tab)
	return rv
} /* debug [instance_methods/method]: HasActiveUserGestureInTab */

// Checks if the extension has script or stylesheet content that can be injected into the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasInjectedContent(for:)
func (w_ WebExtensionContext) HasInjectedContentForURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContentForURL:"), url)
	return rv
} /* debug [instance_methods/method]: HasInjectedContentForURL */

// Checks the specified permission against the currently granted permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasPermission(_:)
func (w_ WebExtensionContext) HasPermission(permission WebExtensionPermission /* typedef */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasPermission:"), permission)
	return rv
} /* debug [instance_methods/method]: HasPermission */

// Checks the specified permission against the currently granted permissions in a specific tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasPermission(_:in:)
func (w_ WebExtensionContext) HasPermissionInTab(permission WebExtensionPermission /* typedef */, tab unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasPermission:inTab:"), permission, tab)
	return rv
} /* debug [instance_methods/method]: HasPermissionInTab */

// Loads the background content if needed for the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/loadBackgroundContent(completionHandler:)
func (w_ WebExtensionContext) LoadBackgroundContentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("loadBackgroundContentWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: LoadBackgroundContentWithCompletionHandler */

// Retrieves the menu items for a given tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/menuItems(for:)
func (w_ WebExtensionContext) MenuItemsForTab(tab unsafe.Pointer) []appkit.MenuItem {
	rv := objc.Send[[]appkit.MenuItem](w_.ID, objc.Sel("menuItemsForTab:"), tab)
	return rv
} /* debug [instance_methods/method]: MenuItemsForTab */

// Performs the extension action associated with the specified tab or performs the default action if is passed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performAction(for:)
func (w_ WebExtensionContext) PerformActionForTab(tab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performActionForTab:"), tab)
} /* debug [instance_methods/method]: PerformActionForTab */

// Performs the specified command, triggering events specific to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(_:)
func (w_ WebExtensionContext) PerformCommand(command IWKWebExtensionCommand) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performCommand:"), command)
} /* debug [instance_methods/method]: PerformCommand */

// Performs the command associated with the given event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(for:)-8btj0
func (w_ WebExtensionContext) PerformCommandForEvent(event appkit.Event) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("performCommandForEvent:"), event)
	return rv
} /* debug [instance_methods/method]: PerformCommandForEvent */

// Checks the specified permission against the currently denied, granted, and requested permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-3qq2w
func (w_ WebExtensionContext) PermissionStatusForPermission(permission WebExtensionPermission /* typedef */) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForPermission:"), permission)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForPermission */

// Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-7mu8
func (w_ WebExtensionContext) PermissionStatusForMatchPattern(pattern IWKWebExtensionMatchPattern) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForMatchPattern:"), pattern)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForMatchPattern */

// Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-7ojrb
func (w_ WebExtensionContext) PermissionStatusForURL(url objc.IObject /* cross-framework: NSURL */) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForURL:"), url)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForURL */

// Checks the specified permission against the currently denied, granted, and requested permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-4h82n
func (w_ WebExtensionContext) PermissionStatusForPermissionInTab(permission WebExtensionPermission /* typedef */, tab unsafe.Pointer) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForPermission:inTab:"), permission, tab)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForPermissionInTab */

// Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-96xaf
func (w_ WebExtensionContext) PermissionStatusForURLInTab(url objc.IObject /* cross-framework: NSURL */, tab unsafe.Pointer) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForURL:inTab:"), url, tab)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForURLInTab */

// Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-nqhm
func (w_ WebExtensionContext) PermissionStatusForMatchPatternInTab(pattern IWKWebExtensionMatchPattern, tab unsafe.Pointer) WebExtensionContextPermissionStatus {
	rv := objc.Send[WebExtensionContextPermissionStatus](w_.ID, objc.Sel("permissionStatusForMatchPattern:inTab:"), pattern, tab)
	return rv
} /* debug [instance_methods/method]: PermissionStatusForMatchPatternInTab */

// Sets the status of a permission with a distant future expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-4u95f
func (w_ WebExtensionContext) SetPermissionStatusForPermission(status WebExtensionContextPermissionStatus, permission WebExtensionPermission /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forPermission:"), status, permission)
} /* debug [instance_methods/method]: SetPermissionStatusForPermission */

// Sets the permission status of a URL with a distant future expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-5xahd
func (w_ WebExtensionContext) SetPermissionStatusForURL(status WebExtensionContextPermissionStatus, url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forURL:"), status, url)
} /* debug [instance_methods/method]: SetPermissionStatusForURL */

// Sets the status of a match pattern with a distant future expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-6auqv
func (w_ WebExtensionContext) SetPermissionStatusForMatchPattern(status WebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forMatchPattern:"), status, pattern)
} /* debug [instance_methods/method]: SetPermissionStatusForMatchPattern */

// Sets the permission status of a URL with a distant future expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-5q9id
func (w_ WebExtensionContext) SetPermissionStatusForURLExpirationDate(status WebExtensionContextPermissionStatus, url objc.IObject /* cross-framework: NSURL */, expirationDate objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forURL:expirationDate:"), status, url, expirationDate)
} /* debug [instance_methods/method]: SetPermissionStatusForURLExpirationDate */

// Sets the status of a permission with a specific expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-692ui
func (w_ WebExtensionContext) SetPermissionStatusForPermissionExpirationDate(status WebExtensionContextPermissionStatus, permission WebExtensionPermission /* typedef */, expirationDate objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forPermission:expirationDate:"), status, permission, expirationDate)
} /* debug [instance_methods/method]: SetPermissionStatusForPermissionExpirationDate */

// Sets the status of a match pattern with a specific expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-7038f
func (w_ WebExtensionContext) SetPermissionStatusForMatchPatternExpirationDate(status WebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern, expirationDate objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPermissionStatus:forMatchPattern:expirationDate:"), status, pattern, expirationDate)
} /* debug [instance_methods/method]: SetPermissionStatusForMatchPatternExpirationDate */

// Should be called by the app when a user gesture is performed in a specific tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/userGesturePerformed(in:)
func (w_ WebExtensionContext) UserGesturePerformedInTab(tab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("userGesturePerformedInTab:"), tab)
} /* debug [instance_methods/method]: UserGesturePerformedInTab */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebExtensionContext */

// The base URL the context uses for loading extension resources or injecting content into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/baseURL
func (w_ WebExtensionContext) BaseURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("baseURL"))
	return rv
} /* debug [instance_properties/getter]: baseURL */

// The base URL the context uses for loading extension resources or injecting content into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/baseURL
func (w_ WebExtensionContext) SetBaseURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBaseURL:"), value)
} /* debug [instance_properties/setter]: baseURL */

// The commands associated with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/commands
func (w_ WebExtensionContext) Commands() []WebExtensionCommand {
	rv := objc.Send[[]WebExtensionCommand](w_.ID, objc.Sel("commands"))
	return rv
} /* debug [instance_properties/getter]: commands */

// The currently granted permission match patterns that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/currentPermissionMatchPatterns
func (w_ WebExtensionContext) CurrentPermissionMatchPatterns() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentPermissionMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: currentPermissionMatchPatterns */

// The currently granted permissions that have not expired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/currentPermissions
func (w_ WebExtensionContext) CurrentPermissions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentPermissions"))
	return rv
} /* debug [instance_properties/getter]: currentPermissions */

// The currently denied permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissionMatchPatterns
func (w_ WebExtensionContext) DeniedPermissionMatchPatterns() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("deniedPermissionMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: deniedPermissionMatchPatterns */

// The currently denied permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissionMatchPatterns
func (w_ WebExtensionContext) SetDeniedPermissionMatchPatterns(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissionMatchPatterns:"), value)
} /* debug [instance_properties/setter]: deniedPermissionMatchPatterns */

// The currently denied permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissions
func (w_ WebExtensionContext) DeniedPermissions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("deniedPermissions"))
	return rv
} /* debug [instance_properties/getter]: deniedPermissions */

// The currently denied permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissions
func (w_ WebExtensionContext) SetDeniedPermissions(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDeniedPermissions:"), value)
} /* debug [instance_properties/setter]: deniedPermissions */

// All errors that occurred in the extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/errors
func (w_ WebExtensionContext) Errors() []objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[[]coretelephony.Error](w_.ID, objc.Sel("errors"))
	return rv
} /* debug [instance_properties/getter]: errors */

// The window that currently has focus for this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/focusedWindow
func (w_ WebExtensionContext) FocusedWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("focusedWindow"))
	return rv
} /* debug [instance_properties/getter]: focusedWindow */

// The currently granted permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissionMatchPatterns
func (w_ WebExtensionContext) GrantedPermissionMatchPatterns() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("grantedPermissionMatchPatterns"))
	return rv
} /* debug [instance_properties/getter]: grantedPermissionMatchPatterns */

// The currently granted permission match patterns and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissionMatchPatterns
func (w_ WebExtensionContext) SetGrantedPermissionMatchPatterns(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissionMatchPatterns:"), value)
} /* debug [instance_properties/setter]: grantedPermissionMatchPatterns */

// The currently granted permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissions
func (w_ WebExtensionContext) GrantedPermissions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("grantedPermissions"))
	return rv
} /* debug [instance_properties/getter]: grantedPermissions */

// The currently granted permissions and their expiration dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissions
func (w_ WebExtensionContext) SetGrantedPermissions(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setGrantedPermissions:"), value)
} /* debug [instance_properties/setter]: grantedPermissions */

// A Boolean value indicating if the currently granted permission match patterns set contains the pattern or any host patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToAllHosts
func (w_ WebExtensionContext) HasAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllHosts"))
	return rv
} /* debug [instance_properties/getter]: hasAccessToAllHosts */

// A Boolean value indicating if the currently granted permission match patterns set contains the pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToAllURLs
func (w_ WebExtensionContext) HasAccessToAllURLs() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToAllURLs"))
	return rv
} /* debug [instance_properties/getter]: hasAccessToAllURLs */

// A Boolean value indicating if the extension has access to private data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToPrivateData
func (w_ WebExtensionContext) HasAccessToPrivateData() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasAccessToPrivateData"))
	return rv
} /* debug [instance_properties/getter]: hasAccessToPrivateData */

// A Boolean value indicating if the extension has access to private data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToPrivateData
func (w_ WebExtensionContext) SetHasAccessToPrivateData(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasAccessToPrivateData:"), value)
} /* debug [instance_properties/setter]: hasAccessToPrivateData */

// A boolean value indicating whether the extension includes rules used for content modification or blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasContentModificationRules
func (w_ WebExtensionContext) HasContentModificationRules() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasContentModificationRules"))
	return rv
} /* debug [instance_properties/getter]: hasContentModificationRules */

// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasInjectedContent
func (w_ WebExtensionContext) HasInjectedContent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasInjectedContent"))
	return rv
} /* debug [instance_properties/getter]: hasInjectedContent */

// A Boolean value indicating if the extension has requested optional access to all hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasRequestedOptionalAccessToAllHosts
func (w_ WebExtensionContext) HasRequestedOptionalAccessToAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasRequestedOptionalAccessToAllHosts"))
	return rv
} /* debug [instance_properties/getter]: hasRequestedOptionalAccessToAllHosts */

// A Boolean value indicating if the extension has requested optional access to all hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasRequestedOptionalAccessToAllHosts
func (w_ WebExtensionContext) SetHasRequestedOptionalAccessToAllHosts(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasRequestedOptionalAccessToAllHosts:"), value)
} /* debug [instance_properties/setter]: hasRequestedOptionalAccessToAllHosts */

// The name shown when inspecting the background web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/inspectionName
func (w_ WebExtensionContext) InspectionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("inspectionName"))
	return rv
} /* debug [instance_properties/getter]: inspectionName */

// The name shown when inspecting the background web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/inspectionName
func (w_ WebExtensionContext) SetInspectionName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectionName:"), value)
} /* debug [instance_properties/setter]: inspectionName */

// Determines whether Web Inspector can inspect the instances for this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/isInspectable
func (w_ WebExtensionContext) Inspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("inspectable"))
	return rv
} /* debug [instance_properties/getter]: inspectable */

// Determines whether Web Inspector can inspect the instances for this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/isInspectable
func (w_ WebExtensionContext) SetInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectable:"), value)
} /* debug [instance_properties/setter]: inspectable */

// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/isLoaded
func (w_ WebExtensionContext) Loaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loaded"))
	return rv
} /* debug [instance_properties/getter]: loaded */

// A set of open tabs in all open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/openTabs
func (w_ WebExtensionContext) OpenTabs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("openTabs"))
	return rv
} /* debug [instance_properties/getter]: openTabs */

// The open windows that are exposed to this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/openWindows
func (w_ WebExtensionContext) OpenWindows() []objc.ID {
	rv := objc.Send[[]objc.ID](w_.ID, objc.Sel("openWindows"))
	return rv
} /* debug [instance_properties/getter]: openWindows */

// The URL of the extension’s options page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/optionsPageURL
func (w_ WebExtensionContext) OptionsPageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("optionsPageURL"))
	return rv
} /* debug [instance_properties/getter]: optionsPageURL */

// The URL to use as an alternative to the default new tab page, if the extension has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/overrideNewTabPageURL
func (w_ WebExtensionContext) OverrideNewTabPageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("overrideNewTabPageURL"))
	return rv
} /* debug [instance_properties/getter]: overrideNewTabPageURL */

// A unique identifier used to distinguish the extension from other extensions and target it for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/uniqueIdentifier
func (w_ WebExtensionContext) UniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("uniqueIdentifier"))
	return rv
} /* debug [instance_properties/getter]: uniqueIdentifier */

// A unique identifier used to distinguish the extension from other extensions and target it for messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/uniqueIdentifier
func (w_ WebExtensionContext) SetUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUniqueIdentifier:"), value)
} /* debug [instance_properties/setter]: uniqueIdentifier */

// Specifies unsupported APIs for this extension, making them in JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/unsupportedAPIs
func (w_ WebExtensionContext) UnsupportedAPIs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("unsupportedAPIs"))
	return rv
} /* debug [instance_properties/getter]: unsupportedAPIs */

// Specifies unsupported APIs for this extension, making them in JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/unsupportedAPIs
func (w_ WebExtensionContext) SetUnsupportedAPIs(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUnsupportedAPIs:"), value)
} /* debug [instance_properties/setter]: unsupportedAPIs */

// The extension this context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webExtension
func (w_ WebExtensionContext) WebExtension() IWKWebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("webExtension"))
	return rv
} /* debug [instance_properties/getter]: webExtension */

// The extension controller this context is loaded in, otherwise if it isn’t loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webExtensionController
func (w_ WebExtensionContext) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("webExtensionController"))
	return rv
} /* debug [instance_properties/getter]: webExtensionController */

// The web view configuration to use for web views that load pages from this extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webViewConfiguration
func (w_ WebExtensionContext) WebViewConfiguration() IWKWebViewConfiguration {
	rv := objc.Send[WebViewConfiguration](w_.ID, objc.Sel("webViewConfiguration"))
	return rv
} /* debug [instance_properties/getter]: webViewConfiguration */

// Determines whether Web Inspector can inspect the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) IsInspectable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isInspectable"))
	return rv
} /* debug [instance_properties/getter]: isInspectable */

// Determines whether Web Inspector can inspect the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isinspectable
func (w_ WebExtensionContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsInspectable:"), value)
} /* debug [instance_properties/setter]: isInspectable */

// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) IsLoaded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoaded"))
	return rv
} /* debug [instance_properties/getter]: isLoaded */

// A Boolean value indicating if this context is loaded in an extension controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontext/isloaded
func (w_ WebExtensionContext) SetIsLoaded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoaded:"), value)
} /* debug [instance_properties/setter]: isLoaded */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKWebExtensionContext */
