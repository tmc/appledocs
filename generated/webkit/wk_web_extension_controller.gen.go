// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionController */


/* debug [class_header]: Header for WKWebExtensionController */
// The class instance for the [WebExtensionController] class.
var (
	WebExtensionControllerClass     _WebExtensionControllerClass
	WebExtensionControllerClassOnce sync.Once
)

func getWebExtensionControllerClass() _WebExtensionControllerClass {
	WebExtensionControllerClassOnce.Do(func() {
		WebExtensionControllerClass = _WebExtensionControllerClass{objc.GetClass("WKWebExtensionController")}
	})
	return WebExtensionControllerClass
}

type _WebExtensionControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionController */
// An interface definition for the [WebExtensionController] class.
type IWebExtensionController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionController */
	// properties:
	Configuration() IWKWebExtensionControllerConfiguration
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ExtensionContexts() unsafe.Pointer
	Extensions() unsafe.Pointer
	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionController */
	// methods:
	DidActivateTabPreviousActiveTab(activatedTab unsafe.Pointer, previousTab unsafe.Pointer)
	DidChangeTabPropertiesForTab(properties WebExtensionTabChangedProperties, changedTab unsafe.Pointer)
	DidCloseTabWindowIsClosing(closedTab unsafe.Pointer, windowIsClosing bool)
	DidCloseWindow(closedWindow unsafe.Pointer)
	DidDeselectTabs(deselectedTabs []objc.ID)
	DidFocusWindow(focusedWindow unsafe.Pointer)
	DidMoveTabFromIndexInWindow(movedTab unsafe.Pointer, index uint, oldWindow unsafe.Pointer)
	DidOpenTab(newTab unsafe.Pointer)
	DidOpenWindow(newWindow unsafe.Pointer)
	DidReplaceTabWithTab(oldTab unsafe.Pointer, newTab unsafe.Pointer)
	DidSelectTabs(selectedTabs []objc.ID)
	ExtensionContextForURL(URL objc.IObject /* cross-framework: NSURL */) IWebExtensionContext
	ExtensionContextForExtension(extension IWKWebExtension) IWebExtensionContext
	FetchDataRecordOfTypesForExtensionContextCompletionHandler(dataTypes unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	FetchDataRecordsOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler unsafe.Pointer)
	LoadExtensionContextError(extensionContext IWKWebExtensionContext, error_ objectivec.IObject) bool
	RemoveDataOfTypesFromDataRecordsCompletionHandler(dataTypes unsafe.Pointer, dataRecords []WebExtensionDataRecord, completionHandler unsafe.Pointer)
	UnloadExtensionContextError(extensionContext IWKWebExtensionContext, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionController */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionControllerClass) Alloc() WebExtensionController {
	rv := objc.Send[WebExtensionController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionControllerClass) New() WebExtensionController {
	rv := objc.Send[WebExtensionController](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionController) Init() WebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionController) Autorelease() WebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionController creates a new WebExtensionController instance.
func NewWebExtensionController() WebExtensionController {
	return getWebExtensionControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionController */
// An object that manages a set of loaded extension contexts.
//
// You can have one or more extension controller instances, allowing different parts of the app to use different sets of extensions. You can associate a controller with using the property on .


// An object that manages a set of loaded extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController
type WebExtensionController struct {
	objectivec.Object
}

// WebExtensionControllerFrom constructs a [WebExtensionController] from an unsafe.Pointer.
//
// An object that manages a set of loaded extension contexts.
func WebExtensionControllerFrom(ptr unsafe.Pointer) WebExtensionController {
	return WebExtensionController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionController */

// Returns a web extension controller initialized with the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/init(configuration:)
func NewWebExtensionControllerWithConfiguration(configuration IWKWebExtensionControllerConfiguration) WebExtensionController {
	instance := getWebExtensionControllerClass().Alloc()
	rv := objc.Send[WebExtensionController](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebExtensionControllerWithConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionController */

// Returns a set of all available extension data types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/allExtensionDataTypes
func (wc _WebExtensionControllerClass) AllExtensionDataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("allExtensionDataTypes"))
	return rv
}/* debug [class_properties_class/property]: allExtensionDataTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionController */

// Should be called by the app when a tab is activated to notify all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didActivateTab:previousActiveTab:
func (w_ WebExtensionController) DidActivateTabPreviousActiveTab(activatedTab unsafe.Pointer, previousTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didActivateTab:previousActiveTab:"), activatedTab, previousTab)
}/* debug [instance_methods/method]: DidActivateTabPreviousActiveTab */


// Should be called by the app when the properties of a tab are changed to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didChangeTabProperties(_:for:)
func (w_ WebExtensionController) DidChangeTabPropertiesForTab(properties WebExtensionTabChangedProperties, changedTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didChangeTabProperties:forTab:"), properties, changedTab)
}/* debug [instance_methods/method]: DidChangeTabPropertiesForTab */


// Should be called by the app when a tab is closed to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didCloseTab:windowIsClosing:
func (w_ WebExtensionController) DidCloseTabWindowIsClosing(closedTab unsafe.Pointer, windowIsClosing bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didCloseTab:windowIsClosing:"), closedTab, windowIsClosing)
}/* debug [instance_methods/method]: DidCloseTabWindowIsClosing */


// Should be called by the app when a window is closed to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didCloseWindow(_:)
func (w_ WebExtensionController) DidCloseWindow(closedWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didCloseWindow:"), closedWindow)
}/* debug [instance_methods/method]: DidCloseWindow */


// Should be called by the app when tabs are deselected to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didDeselectTabs(_:)
func (w_ WebExtensionController) DidDeselectTabs(deselectedTabs []objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didDeselectTabs:"), deselectedTabs)
}/* debug [instance_methods/method]: DidDeselectTabs */


// Should be called by the app when a window gains focus to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didFocusWindow(_:)
func (w_ WebExtensionController) DidFocusWindow(focusedWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didFocusWindow:"), focusedWindow)
}/* debug [instance_methods/method]: DidFocusWindow */


// Should be called by the app when a tab is moved to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didMoveTab:fromIndex:inWindow:
func (w_ WebExtensionController) DidMoveTabFromIndexInWindow(movedTab unsafe.Pointer, index uint, oldWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didMoveTab:fromIndex:inWindow:"), movedTab, index, oldWindow)
}/* debug [instance_methods/method]: DidMoveTabFromIndexInWindow */


// Should be called by the app when a new tab is opened to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didOpenTab(_:)
func (w_ WebExtensionController) DidOpenTab(newTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didOpenTab:"), newTab)
}/* debug [instance_methods/method]: DidOpenTab */


// Should be called by the app when a new window is opened to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didOpenWindow(_:)
func (w_ WebExtensionController) DidOpenWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didOpenWindow:"), newWindow)
}/* debug [instance_methods/method]: DidOpenWindow */


// Should be called by the app when a tab is replaced by another tab to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didReplaceTab(_:with:)
func (w_ WebExtensionController) DidReplaceTabWithTab(oldTab unsafe.Pointer, newTab unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didReplaceTab:withTab:"), oldTab, newTab)
}/* debug [instance_methods/method]: DidReplaceTabWithTab */


// Should be called by the app when tabs are selected to fire appropriate events with all loaded web extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didSelectTabs(_:)
func (w_ WebExtensionController) DidSelectTabs(selectedTabs []objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("didSelectTabs:"), selectedTabs)
}/* debug [instance_methods/method]: DidSelectTabs */


// Returns a loaded extension context matching the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContext(for:)-2kr4
func (w_ WebExtensionController) ExtensionContextForURL(URL objc.IObject /* cross-framework: NSURL */) IWebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("extensionContextForURL:"), URL)
	return rv
}/* debug [instance_methods/method]: ExtensionContextForURL */


// Returns a loaded extension context for the specified extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContext(for:)-6ecpm
func (w_ WebExtensionController) ExtensionContextForExtension(extension IWKWebExtension) IWebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("extensionContextForExtension:"), extension)
	return rv
}/* debug [instance_methods/method]: ExtensionContextForExtension */


// Fetches a data record containing the given extension data types for a specific known web extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/fetchDataRecord(ofTypes:for:completionHandler:)
func (w_ WebExtensionController) FetchDataRecordOfTypesForExtensionContextCompletionHandler(dataTypes unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataRecordOfTypes:forExtensionContext:completionHandler:"), dataTypes, extensionContext, completionHandler)
}/* debug [instance_methods/method]: FetchDataRecordOfTypesForExtensionContextCompletionHandler */


// Fetches data records containing the given extension data types for all known extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/fetchDataRecords(ofTypes:completionHandler:)
func (w_ WebExtensionController) FetchDataRecordsOfTypesCompletionHandler(dataTypes unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("fetchDataRecordsOfTypes:completionHandler:"), dataTypes, completionHandler)
}/* debug [instance_methods/method]: FetchDataRecordsOfTypesCompletionHandler */


// Loads the specified extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/load(_:)
func (w_ WebExtensionController) LoadExtensionContextError(extensionContext IWKWebExtensionContext, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loadExtensionContext:error:"), extensionContext, error_)
	return rv
}/* debug [instance_methods/method]: LoadExtensionContextError */


// Removes extension data of the given types for the given data records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/removeData(ofTypes:from:completionHandler:)
func (w_ WebExtensionController) RemoveDataOfTypesFromDataRecordsCompletionHandler(dataTypes unsafe.Pointer, dataRecords []WebExtensionDataRecord, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeDataOfTypes:fromDataRecords:completionHandler:"), dataTypes, dataRecords, completionHandler)
}/* debug [instance_methods/method]: RemoveDataOfTypesFromDataRecordsCompletionHandler */


// Unloads the specified extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/unload(_:)
func (w_ WebExtensionController) UnloadExtensionContextError(extensionContext IWKWebExtensionContext, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unloadExtensionContext:error:"), extensionContext, error_)
	return rv
}/* debug [instance_methods/method]: UnloadExtensionContextError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionController */

// Returns a set of all available extension data types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/allExtensionDataTypes
func (w_ WebExtensionController) AllExtensionDataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("allExtensionDataTypes"))
	return rv
}/* debug [instance_properties/getter]: allExtensionDataTypes */


// A copy of the configuration with which the web extension controller was initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/configuration-swift.property
func (w_ WebExtensionController) Configuration() IWKWebExtensionControllerConfiguration {
	rv := objc.Send[WebExtensionControllerConfiguration](w_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The extension controller delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/delegate
func (w_ WebExtensionController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The extension controller delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/delegate
func (w_ WebExtensionController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A set of all the currently loaded extension contexts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContexts
func (w_ WebExtensionController) ExtensionContexts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("extensionContexts"))
	return rv
}/* debug [instance_properties/getter]: extensionContexts */


// A set of all the currently loaded extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensions
func (w_ WebExtensionController) Extensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("extensions"))
	return rv
}/* debug [instance_properties/getter]: extensions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebExtensionController) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("webExtensionController"))
	return rv
}/* debug [instance_properties/getter]: webExtensionController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebExtensionController) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}/* debug [instance_properties/setter]: webExtensionController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionController */


