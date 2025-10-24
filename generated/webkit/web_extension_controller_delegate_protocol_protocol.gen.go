// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PWebExtensionControllerDelegate is the WKWebExtensionControllerDelegate protocol interface.
//
// A group of methods you use to customize web extension interactions.
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - visionOS 2.4+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKWebExtensionControllerDelegate
type PWebExtensionControllerDelegate interface {
	// Optional methods
	WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler(controller IWKWebExtensionController, port IWKWebExtensionMessagePort, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler() bool
	WebExtensionControllerDidUpdateActionForExtensionContext(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext)
	HasWebExtensionControllerDidUpdateActionForExtensionContext() bool
	WebExtensionControllerFocusedWindowForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) unsafe.Pointer
	HasWebExtensionControllerFocusedWindowForExtensionContext() bool
	WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionTabConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler() bool
	WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionWindowConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler() bool
	WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler() bool
	WebExtensionControllerOpenWindowsForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) []objc.ID
	HasWebExtensionControllerOpenWindowsForExtensionContext() bool
	WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler() bool
	WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, matchPatterns unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler() bool
	WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, permissions unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler() bool
	WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, urls unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasWebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler() bool
	WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler(controller IWKWebExtensionController, message objc.IObject, applicationIdentifier objc.IObject /* cross-framework: NSString */, extensionContext IWKWebExtensionContext, replyHandler unsafe.Pointer)
	HasWebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler() bool
}

// WebExtensionControllerDelegate is a delegate implementation builder for the PWebExtensionControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebExtensionControllerDelegate struct {
	_WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler               func(controller IWKWebExtensionController, port IWKWebExtensionMessagePort, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerDidUpdateActionForExtensionContext                                        func(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext)
	_WebExtensionControllerFocusedWindowForExtensionContext                                          func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) unsafe.Pointer
	_WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler          func(controller IWKWebExtensionController, configuration IWKWebExtensionTabConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler       func(controller IWKWebExtensionController, configuration IWKWebExtensionWindowConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler                       func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerOpenWindowsForExtensionContext                                            func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) []objc.ID
	_WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler                 func(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler func(controller IWKWebExtensionController, matchPatterns unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler             func(controller IWKWebExtensionController, permissions unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler  func(controller IWKWebExtensionController, urls unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)
	_WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler     func(controller IWKWebExtensionController, message objc.IObject, applicationIdentifier objc.IObject /* cross-framework: NSString */, extensionContext IWKWebExtensionContext, replyHandler unsafe.Pointer)
}

// SetWebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context wants to establish a persistent connection to an application.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, port IWKWebExtensionMessagePort, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerDidUpdateActionForExtensionContext sets the handler for the WebExtensionControllerDidUpdateActionForExtensionContext delegate method.
//
// Called when an action’s properties are updated.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerDidUpdateActionForExtensionContext(f func(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext)) {
	d._WebExtensionControllerDidUpdateActionForExtensionContext = f
}

// SetWebExtensionControllerFocusedWindowForExtensionContext sets the handler for the WebExtensionControllerFocusedWindowForExtensionContext delegate method.
//
// Called when an extension context requests the currently focused window.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerFocusedWindowForExtensionContext(f func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) unsafe.Pointer) {
	d._WebExtensionControllerFocusedWindowForExtensionContext = f
}

// SetWebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests a new tab to be opened.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, configuration IWKWebExtensionTabConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests a new window to be opened.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, configuration IWKWebExtensionWindowConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests its options page to be opened.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerOpenWindowsForExtensionContext sets the handler for the WebExtensionControllerOpenWindowsForExtensionContext delegate method.
//
// Called when an extension context requests the list of ordered open windows.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerOpenWindowsForExtensionContext(f func(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) []objc.ID) {
	d._WebExtensionControllerOpenWindowsForExtensionContext = f
}

// SetWebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler delegate method.
//
// Called when a popup is requested to be displayed for a specific action.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests access to a set of match patterns.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, matchPatterns unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests permissions.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, permissions unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler sets the handler for the WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler delegate method.
//
// Called when an extension context requests access to a set of URLs.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler(f func(controller IWKWebExtensionController, urls unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer)) {
	d._WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler = f
}

// SetWebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler sets the handler for the WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler delegate method.
//
// Called when an extension context wants to send a one-time message to an application.
func (d *WebExtensionControllerDelegate) SetWebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler(f func(controller IWKWebExtensionController, message objc.IObject, applicationIdentifier objc.IObject /* cross-framework: NSString */, extensionContext IWKWebExtensionContext, replyHandler unsafe.Pointer)) {
	d._WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler = f
}

// WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler(controller IWKWebExtensionController, port IWKWebExtensionMessagePort, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler(controller, port, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerDidUpdateActionForExtensionContext implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerDidUpdateActionForExtensionContext(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext) {
	if d._WebExtensionControllerDidUpdateActionForExtensionContext != nil {
		d._WebExtensionControllerDidUpdateActionForExtensionContext(controller, action, context)
	}
}

// HasWebExtensionControllerDidUpdateActionForExtensionContext returns true if a handler for WebExtensionControllerDidUpdateActionForExtensionContext has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerDidUpdateActionForExtensionContext() bool {
	return d._WebExtensionControllerDidUpdateActionForExtensionContext != nil
}

// WebExtensionControllerFocusedWindowForExtensionContext implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerFocusedWindowForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) unsafe.Pointer {
	if d._WebExtensionControllerFocusedWindowForExtensionContext != nil {
		return d._WebExtensionControllerFocusedWindowForExtensionContext(controller, extensionContext)
	}
	var zero unsafe.Pointer
	return zero
}

// HasWebExtensionControllerFocusedWindowForExtensionContext returns true if a handler for WebExtensionControllerFocusedWindowForExtensionContext has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerFocusedWindowForExtensionContext() bool {
	return d._WebExtensionControllerFocusedWindowForExtensionContext != nil
}

// WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionTabConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler(controller, configuration, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionWindowConfiguration, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler(controller, configuration, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler(controller, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerOpenWindowsForExtensionContext implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerOpenWindowsForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) []objc.ID {
	if d._WebExtensionControllerOpenWindowsForExtensionContext != nil {
		return d._WebExtensionControllerOpenWindowsForExtensionContext(controller, extensionContext)
	}
	var zero []objc.ID
	return zero
}

// HasWebExtensionControllerOpenWindowsForExtensionContext returns true if a handler for WebExtensionControllerOpenWindowsForExtensionContext has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerOpenWindowsForExtensionContext() bool {
	return d._WebExtensionControllerOpenWindowsForExtensionContext != nil
}

// WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler(controller, action, context, completionHandler)
	}
}

// HasWebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, matchPatterns unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler(controller, matchPatterns, tab, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, permissions unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler(controller, permissions, tab, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler(controller IWKWebExtensionController, urls unsafe.Pointer, tab unsafe.Pointer, extensionContext IWKWebExtensionContext, completionHandler unsafe.Pointer) {
	if d._WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler != nil {
		d._WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler(controller, urls, tab, extensionContext, completionHandler)
	}
}

// HasWebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler returns true if a handler for WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler() bool {
	return d._WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler != nil
}

// WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler implements the PWebExtensionControllerDelegate interface.
func (d *WebExtensionControllerDelegate) WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler(controller IWKWebExtensionController, message objc.IObject, applicationIdentifier objc.IObject /* cross-framework: NSString */, extensionContext IWKWebExtensionContext, replyHandler unsafe.Pointer) {
	if d._WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler != nil {
		d._WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler(controller, message, applicationIdentifier, extensionContext, replyHandler)
	}
}

// HasWebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler returns true if a handler for WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler has been set.
func (d *WebExtensionControllerDelegate) HasWebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler() bool {
	return d._WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler != nil
}
