// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PWebExtensionWindow is the WKWebExtensionWindow protocol interface.
//
// A protocol with methods that represent a window to web extensions.
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - visionOS 2.4+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKWebExtensionWindow
type PWebExtensionWindow interface {
	// Optional methods
	ActiveTabForWebExtensionContext(context IWKWebExtensionContext) unsafe.Pointer
	HasActiveTabForWebExtensionContext() bool
	CloseForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasCloseForWebExtensionContextCompletionHandler() bool
	FocusForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasFocusForWebExtensionContextCompletionHandler() bool
	FrameForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGRect
	HasFrameForWebExtensionContext() bool
	IsPrivateForWebExtensionContext(context IWKWebExtensionContext) bool
	HasIsPrivateForWebExtensionContext() bool
	ScreenFrameForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGRect
	HasScreenFrameForWebExtensionContext() bool
	SetFrameForWebExtensionContextCompletionHandler(frame corefoundation.CGRect, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetFrameForWebExtensionContextCompletionHandler() bool
	SetWindowStateForWebExtensionContextCompletionHandler(state WebExtensionWindowState, context IWKWebExtensionContext, completionHandler unsafe.Pointer)
	HasSetWindowStateForWebExtensionContextCompletionHandler() bool
	TabsForWebExtensionContext(context IWKWebExtensionContext) []objc.ID
	HasTabsForWebExtensionContext() bool
	WindowStateForWebExtensionContext(context IWKWebExtensionContext) WebExtensionWindowState
	HasWindowStateForWebExtensionContext() bool
	WindowTypeForWebExtensionContext(context IWKWebExtensionContext) WebExtensionWindowType
	HasWindowTypeForWebExtensionContext() bool
}
