// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSFSafariExtensionHandling is the SFSafariExtensionHandling protocol interface.
//
// A protocol for implementing event handling in a Safari app extension.
//
// Availability:
//   - macOS 10.12+
//
// See: doc://com.apple.safariservices/documentation/SafariServices/SFSafariExtensionHandling
type PSFSafariExtensionHandling interface {
	// Optional methods
	AdditionalRequestHeadersForURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	HasAdditionalRequestHeadersForURLCompletionHandler() bool
	ContentBlockerWithIdentifierBlockedResourcesWithURLsOnPage(contentBlockerIdentifier objc.IObject /* cross-framework: NSString */, urls []foundation.URL, page ISFSafariPage)
	HasContentBlockerWithIdentifierBlockedResourcesWithURLsOnPage() bool
	ContextMenuItemSelectedWithCommandInPageUserInfo(command objc.IObject /* cross-framework: NSString */, page ISFSafariPage, userInfo foundation.IDictionary)
	HasContextMenuItemSelectedWithCommandInPageUserInfo() bool
	MessageReceivedWithNameFromPageUserInfo(messageName objc.IObject /* cross-framework: NSString */, page ISFSafariPage, userInfo foundation.IDictionary)
	HasMessageReceivedWithNameFromPageUserInfo() bool
	MessageReceivedFromContainingAppWithNameUserInfo(messageName objc.IObject /* cross-framework: NSString */, userInfo foundation.IDictionary)
	HasMessageReceivedFromContainingAppWithNameUserInfo() bool
	PageWillNavigateToURL(page ISFSafariPage, url objc.IObject /* cross-framework: NSURL */)
	HasPageWillNavigateToURL() bool
	PopoverDidCloseInWindow(window ISFSafariWindow)
	HasPopoverDidCloseInWindow() bool
	PopoverViewController() SFSafariExtensionViewController
	HasPopoverViewController() bool
	PopoverWillShowInWindow(window ISFSafariWindow)
	HasPopoverWillShowInWindow() bool
	ToolbarItemClickedInWindow(window ISFSafariWindow)
	HasToolbarItemClickedInWindow() bool
	ValidateContextMenuItemWithCommandInPageUserInfoValidationHandler(command objc.IObject /* cross-framework: NSString */, page ISFSafariPage, userInfo foundation.IDictionary, validationHandler unsafe.Pointer)
	HasValidateContextMenuItemWithCommandInPageUserInfoValidationHandler() bool
	ValidateToolbarItemInWindowValidationHandler(window ISFSafariWindow, validationHandler unsafe.Pointer)
	HasValidateToolbarItemInWindowValidationHandler() bool
}
