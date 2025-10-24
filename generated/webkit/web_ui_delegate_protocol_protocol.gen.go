// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
)

// PWebUIDelegate is the WebUIDelegate protocol interface.
//
// Web view user interface delegates implement this protocol to control the opening of new windows, augment the behavior of default menu items displayed when the user clicks elements, and perform other user interface–related tasks. These methods can be invoked as a result of handling JavaScript or other plug-in content. Delegates that display more than one web view per window, for example, need to implement some of these methods to handle that case. The default implementation assumes one window per web view, so non-conventional user interfaces might implement a user interface delegate.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebUIDelegate
type PWebUIDelegate interface {
	// Required methods
	WebViewRunJavaScriptAlertPanelWithMessage(sender IWebView, message objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: WebViewRunJavaScriptAlertPanelWithMessage */
	WebViewRunJavaScriptConfirmPanelWithMessage(sender IWebView, message objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: WebViewRunJavaScriptConfirmPanelWithMessage */
	WebViewRunJavaScriptTextInputPanelWithPromptDefaultText(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */) foundation.String/* debug [protocol_interface/required_method]: WebViewRunJavaScriptTextInputPanelWithPromptDefaultText */
	WebViewSetContentRect(sender IWebView, frame Rect /* not a class type */)/* debug [protocol_interface/required_method]: WebViewSetContentRect */
	WebViewContentRect(sender IWebView) Rect/* debug [protocol_interface/required_method]: WebViewContentRect */
	// Optional methods
	WebViewContextMenuItemsForElementDefaultMenuItems(sender IWebView, element objc.IObject /* cross-framework: NSDictionary */, defaultMenuItems objc.IObject /* cross-framework: NSArray */) foundation.Array
	HasWebViewContextMenuItemsForElementDefaultMenuItems() bool
	WebViewCreateWebViewModalDialogWithRequest(sender IWebView, request foundation.URLRequest) WebView
	HasWebViewCreateWebViewModalDialogWithRequest() bool
	WebViewCreateWebViewWithRequest(sender IWebView, request foundation.URLRequest) WebView
	HasWebViewCreateWebViewWithRequest() bool
	WebViewDragDestinationActionMaskForDraggingInfo(webView IWebView, draggingInfo unsafe.Pointer) uint
	HasWebViewDragDestinationActionMaskForDraggingInfo() bool
	WebViewDragSourceActionMaskForPoint(webView IWebView, point vision.Point) uint
	HasWebViewDragSourceActionMaskForPoint() bool
	WebViewDrawFooterInRect(sender IWebView, rect Rect /* not a class type */)
	HasWebViewDrawFooterInRect() bool
	WebViewDrawHeaderInRect(sender IWebView, rect Rect /* not a class type */)
	HasWebViewDrawHeaderInRect() bool
	WebViewMakeFirstResponder(sender IWebView, responder appkit.Responder)
	HasWebViewMakeFirstResponder() bool
	WebViewMouseDidMoveOverElementModifierFlags(sender IWebView, elementInformation objc.IObject /* cross-framework: NSDictionary */, modifierFlags uint)
	HasWebViewMouseDidMoveOverElementModifierFlags() bool
	WebViewPrintFrameView(sender IWebView, frameView IWebFrameView)
	HasWebViewPrintFrameView() bool
	WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool
	HasWebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame() bool
	WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame)
	HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame() bool
	WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool
	HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame() bool
	WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWebFrame) foundation.String
	HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame() bool
	WebViewRunOpenPanelForFileButtonWithResultListener(sender IWebView, resultListener unsafe.Pointer)
	HasWebViewRunOpenPanelForFileButtonWithResultListener() bool
	WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles(sender IWebView, resultListener unsafe.Pointer, allowMultipleFiles bool)
	HasWebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles() bool
	WebViewSetFrame(sender IWebView, frame Rect /* not a class type */)
	HasWebViewSetFrame() bool
	WebViewSetResizable(sender IWebView, resizable bool)
	HasWebViewSetResizable() bool
	WebViewSetStatusBarVisible(sender IWebView, visible bool)
	HasWebViewSetStatusBarVisible() bool
	WebViewSetStatusText(sender IWebView, text objc.IObject /* cross-framework: NSString */)
	HasWebViewSetStatusText() bool
	WebViewSetToolbarsVisible(sender IWebView, visible bool)
	HasWebViewSetToolbarsVisible() bool
	WebViewShouldPerformActionFromSender(webView IWebView, action objc.SEL, sender objc.IObject) bool
	HasWebViewShouldPerformActionFromSender() bool
	WebViewValidateUserInterfaceItemDefaultValidation(webView IWebView, item unsafe.Pointer, defaultValidation bool) bool
	HasWebViewValidateUserInterfaceItemDefaultValidation() bool
	WebViewWillPerformDragDestinationActionForDraggingInfo(webView IWebView, action WebDragDestinationAction, draggingInfo unsafe.Pointer)
	HasWebViewWillPerformDragDestinationActionForDraggingInfo() bool
	WebViewWillPerformDragSourceActionFromPointWithPasteboard(webView IWebView, action WebDragSourceAction, point vision.Point, pasteboard appkit.Pasteboard)
	HasWebViewWillPerformDragSourceActionFromPointWithPasteboard() bool
	WebViewAreToolbarsVisible(sender IWebView) bool
	HasWebViewAreToolbarsVisible() bool
	WebViewClose(sender IWebView)
	HasWebViewClose() bool
	WebViewFirstResponder(sender IWebView) appkit.Responder
	HasWebViewFirstResponder() bool
	WebViewFocus(sender IWebView)
	HasWebViewFocus() bool
	WebViewFooterHeight(sender IWebView) float32
	HasWebViewFooterHeight() bool
	WebViewFrame(sender IWebView) Rect
	HasWebViewFrame() bool
	WebViewHeaderHeight(sender IWebView) float32
	HasWebViewHeaderHeight() bool
	WebViewIsResizable(sender IWebView) bool
	HasWebViewIsResizable() bool
	WebViewIsStatusBarVisible(sender IWebView) bool
	HasWebViewIsStatusBarVisible() bool
	WebViewRunModal(sender IWebView)
	HasWebViewRunModal() bool
	WebViewShow(sender IWebView)
	HasWebViewShow() bool
	WebViewStatusText(sender IWebView) foundation.String
	HasWebViewStatusText() bool
	WebViewUnfocus(sender IWebView)
	HasWebViewUnfocus() bool
}

// WebUIDelegate is a delegate implementation builder for the PWebUIDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebUIDelegate struct {
	_WebViewContextMenuItemsForElementDefaultMenuItems func(sender IWebView, element objc.IObject /* cross-framework: NSDictionary */, defaultMenuItems objc.IObject /* cross-framework: NSArray */) foundation.Array
	_WebViewCreateWebViewModalDialogWithRequest func(sender IWebView, request foundation.URLRequest) WebView
	_WebViewCreateWebViewWithRequest func(sender IWebView, request foundation.URLRequest) WebView
	_WebViewDragDestinationActionMaskForDraggingInfo func(webView IWebView, draggingInfo unsafe.Pointer) uint
	_WebViewDragSourceActionMaskForPoint func(webView IWebView, point vision.Point) uint
	_WebViewDrawFooterInRect func(sender IWebView, rect Rect /* not a class type */)
	_WebViewDrawHeaderInRect func(sender IWebView, rect Rect /* not a class type */)
	_WebViewMakeFirstResponder func(sender IWebView, responder appkit.Responder)
	_WebViewMouseDidMoveOverElementModifierFlags func(sender IWebView, elementInformation objc.IObject /* cross-framework: NSDictionary */, modifierFlags uint)
	_WebViewPrintFrameView func(sender IWebView, frameView IWebFrameView)
	_WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool
	_WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame)
	_WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool
	_WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame func(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWebFrame) foundation.String
	_WebViewRunOpenPanelForFileButtonWithResultListener func(sender IWebView, resultListener unsafe.Pointer)
	_WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles func(sender IWebView, resultListener unsafe.Pointer, allowMultipleFiles bool)
	_WebViewSetFrame func(sender IWebView, frame Rect /* not a class type */)
	_WebViewSetResizable func(sender IWebView, resizable bool)
	_WebViewSetStatusBarVisible func(sender IWebView, visible bool)
	_WebViewSetStatusText func(sender IWebView, text objc.IObject /* cross-framework: NSString */)
	_WebViewSetToolbarsVisible func(sender IWebView, visible bool)
	_WebViewShouldPerformActionFromSender func(webView IWebView, action objc.SEL, sender objc.IObject) bool
	_WebViewValidateUserInterfaceItemDefaultValidation func(webView IWebView, item unsafe.Pointer, defaultValidation bool) bool
	_WebViewWillPerformDragDestinationActionForDraggingInfo func(webView IWebView, action WebDragDestinationAction, draggingInfo unsafe.Pointer)
	_WebViewWillPerformDragSourceActionFromPointWithPasteboard func(webView IWebView, action WebDragSourceAction, point vision.Point, pasteboard appkit.Pasteboard)
	_WebViewAreToolbarsVisible func(sender IWebView) bool
	_WebViewClose func(sender IWebView)
	_WebViewFirstResponder func(sender IWebView) appkit.Responder
	_WebViewFocus func(sender IWebView)
	_WebViewFooterHeight func(sender IWebView) float32
	_WebViewFrame func(sender IWebView) Rect
	_WebViewHeaderHeight func(sender IWebView) float32
	_WebViewIsResizable func(sender IWebView) bool
	_WebViewIsStatusBarVisible func(sender IWebView) bool
	_WebViewRunModal func(sender IWebView)
	_WebViewShow func(sender IWebView)
	_WebViewStatusText func(sender IWebView) foundation.String
	_WebViewUnfocus func(sender IWebView)
	_WebViewRunJavaScriptAlertPanelWithMessage func(sender IWebView, message objc.IObject /* cross-framework: NSString */)
	_WebViewRunJavaScriptConfirmPanelWithMessage func(sender IWebView, message objc.IObject /* cross-framework: NSString */) bool
	_WebViewRunJavaScriptTextInputPanelWithPromptDefaultText func(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */) foundation.String
	_WebViewSetContentRect func(sender IWebView, frame Rect /* not a class type */)
	_WebViewContentRect func(sender IWebView) Rect
}

// SetWebViewContextMenuItemsForElementDefaultMenuItems sets the handler for the WebViewContextMenuItemsForElementDefaultMenuItems delegate method.
//
// Returns menu items to display in an element’s contextual menu.
func (d *WebUIDelegate) SetWebViewContextMenuItemsForElementDefaultMenuItems(f func(sender IWebView, element objc.IObject /* cross-framework: NSDictionary */, defaultMenuItems objc.IObject /* cross-framework: NSArray */) foundation.Array) {
	d._WebViewContextMenuItemsForElementDefaultMenuItems = f
}

// SetWebViewCreateWebViewModalDialogWithRequest sets the handler for the WebViewCreateWebViewModalDialogWithRequest delegate method.
//
// Creates a modal window containing a web view that loads the specified request.
func (d *WebUIDelegate) SetWebViewCreateWebViewModalDialogWithRequest(f func(sender IWebView, request foundation.URLRequest) WebView) {
	d._WebViewCreateWebViewModalDialogWithRequest = f
}

// SetWebViewCreateWebViewWithRequest sets the handler for the WebViewCreateWebViewWithRequest delegate method.
//
// Creates a window containing a web view to load the specified request.
func (d *WebUIDelegate) SetWebViewCreateWebViewWithRequest(f func(sender IWebView, request foundation.URLRequest) WebView) {
	d._WebViewCreateWebViewWithRequest = f
}

// SetWebViewDragDestinationActionMaskForDraggingInfo sets the handler for the WebViewDragDestinationActionMaskForDraggingInfo delegate method.
//
// Returns a mask indicating which drag operations are allowed by the sender.
func (d *WebUIDelegate) SetWebViewDragDestinationActionMaskForDraggingInfo(f func(webView IWebView, draggingInfo unsafe.Pointer) uint) {
	d._WebViewDragDestinationActionMaskForDraggingInfo = f
}

// SetWebViewDragSourceActionMaskForPoint sets the handler for the WebViewDragSourceActionMaskForPoint delegate method.
//
// Returns a mask indicating which drag-source actions are allowed for a drag that begins at the specified location.
func (d *WebUIDelegate) SetWebViewDragSourceActionMaskForPoint(f func(webView IWebView, point vision.Point) uint) {
	d._WebViewDragSourceActionMaskForPoint = f
}

// SetWebViewDrawFooterInRect sets the handler for the WebViewDrawFooterInRect delegate method.
//
// Draws the web view’s footer in the specified rectangle.
func (d *WebUIDelegate) SetWebViewDrawFooterInRect(f func(sender IWebView, rect Rect /* not a class type */)) {
	d._WebViewDrawFooterInRect = f
}

// SetWebViewDrawHeaderInRect sets the handler for the WebViewDrawHeaderInRect delegate method.
//
// Draws the web view’s header in the specified rectangle.
func (d *WebUIDelegate) SetWebViewDrawHeaderInRect(f func(sender IWebView, rect Rect /* not a class type */)) {
	d._WebViewDrawHeaderInRect = f
}

// SetWebViewMakeFirstResponder sets the handler for the WebViewMakeFirstResponder delegate method.
//
// Sets the first responder of a web view’s window to the specified view.
func (d *WebUIDelegate) SetWebViewMakeFirstResponder(f func(sender IWebView, responder appkit.Responder)) {
	d._WebViewMakeFirstResponder = f
}

// SetWebViewMouseDidMoveOverElementModifierFlags sets the handler for the WebViewMouseDidMoveOverElementModifierFlags delegate method.
//
// Updates information about the element the user is mousing over.
func (d *WebUIDelegate) SetWebViewMouseDidMoveOverElementModifierFlags(f func(sender IWebView, elementInformation objc.IObject /* cross-framework: NSDictionary */, modifierFlags uint)) {
	d._WebViewMouseDidMoveOverElementModifierFlags = f
}

// SetWebViewPrintFrameView sets the handler for the WebViewPrintFrameView delegate method.
//
// Prints the contents of a web frame view.
func (d *WebUIDelegate) SetWebViewPrintFrameView(f func(sender IWebView, frameView IWebFrameView)) {
	d._WebViewPrintFrameView = f
}

// SetWebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame sets the handler for the WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame delegate method.
//
// Displays a confirmation panel containing the specified message before a window closes.
func (d *WebUIDelegate) SetWebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame(f func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool) {
	d._WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame = f
}

// SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame sets the handler for the WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame delegate method.
//
// Displays a JavaScript alert panel containing the specified message.
func (d *WebUIDelegate) SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame(f func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame)) {
	d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame = f
}

// SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame sets the handler for the WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame delegate method.
//
// Displays a JavaScript confirmation panel with the specified message.
func (d *WebUIDelegate) SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame(f func(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool) {
	d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame = f
}

// SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame sets the handler for the WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame delegate method.
//
// Displays a JavaScript text input panel and returns the entered text.
func (d *WebUIDelegate) SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame(f func(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWebFrame) foundation.String) {
	d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame = f
}

// SetWebViewRunOpenPanelForFileButtonWithResultListener sets the handler for the WebViewRunOpenPanelForFileButtonWithResultListener delegate method.
//
// Displays an open panel for a file input control.
func (d *WebUIDelegate) SetWebViewRunOpenPanelForFileButtonWithResultListener(f func(sender IWebView, resultListener unsafe.Pointer)) {
	d._WebViewRunOpenPanelForFileButtonWithResultListener = f
}

// SetWebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles sets the handler for the WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles delegate method.
//
// Displays an open panel for a file input control.
func (d *WebUIDelegate) SetWebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles(f func(sender IWebView, resultListener unsafe.Pointer, allowMultipleFiles bool)) {
	d._WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles = f
}

// SetWebViewSetFrame sets the handler for the WebViewSetFrame delegate method.
//
// Sets the frame rectangle of a web view’s window to the specified frame size.
func (d *WebUIDelegate) SetWebViewSetFrame(f func(sender IWebView, frame Rect /* not a class type */)) {
	d._WebViewSetFrame = f
}

// SetWebViewSetResizable sets the handler for the WebViewSetResizable delegate method.
//
// Sets whether a web view’s window can be resized.
func (d *WebUIDelegate) SetWebViewSetResizable(f func(sender IWebView, resizable bool)) {
	d._WebViewSetResizable = f
}

// SetWebViewSetStatusBarVisible sets the handler for the WebViewSetStatusBarVisible delegate method.
//
// Sets the visibility of the status bar in a web view’s window.
func (d *WebUIDelegate) SetWebViewSetStatusBarVisible(f func(sender IWebView, visible bool)) {
	d._WebViewSetStatusBarVisible = f
}

// SetWebViewSetStatusText sets the handler for the WebViewSetStatusText delegate method.
//
// Sets the status message displayed by a web view’s window, if any, to the specified text.
func (d *WebUIDelegate) SetWebViewSetStatusText(f func(sender IWebView, text objc.IObject /* cross-framework: NSString */)) {
	d._WebViewSetStatusText = f
}

// SetWebViewSetToolbarsVisible sets the handler for the WebViewSetToolbarsVisible delegate method.
//
// Sets whether a web view’s toolbars should be visible.
func (d *WebUIDelegate) SetWebViewSetToolbarsVisible(f func(sender IWebView, visible bool)) {
	d._WebViewSetToolbarsVisible = f
}

// SetWebViewShouldPerformActionFromSender sets the handler for the WebViewShouldPerformActionFromSender delegate method.
//
// Returns a Boolean value that indicates whether the action sent by the specified object should be performed.
func (d *WebUIDelegate) SetWebViewShouldPerformActionFromSender(f func(webView IWebView, action objc.SEL, sender objc.IObject) bool) {
	d._WebViewShouldPerformActionFromSender = f
}

// SetWebViewValidateUserInterfaceItemDefaultValidation sets the handler for the WebViewValidateUserInterfaceItemDefaultValidation delegate method.
//
// Returns a Boolean value that indicates whether the specified user interface item is valid.
func (d *WebUIDelegate) SetWebViewValidateUserInterfaceItemDefaultValidation(f func(webView IWebView, item unsafe.Pointer, defaultValidation bool) bool) {
	d._WebViewValidateUserInterfaceItemDefaultValidation = f
}

// SetWebViewWillPerformDragDestinationActionForDraggingInfo sets the handler for the WebViewWillPerformDragDestinationActionForDraggingInfo delegate method.
//
// Tells the receiver that the sending web view will perform the specified drag-destination action.
func (d *WebUIDelegate) SetWebViewWillPerformDragDestinationActionForDraggingInfo(f func(webView IWebView, action WebDragDestinationAction, draggingInfo unsafe.Pointer)) {
	d._WebViewWillPerformDragDestinationActionForDraggingInfo = f
}

// SetWebViewWillPerformDragSourceActionFromPointWithPasteboard sets the handler for the WebViewWillPerformDragSourceActionFromPointWithPasteboard delegate method.
//
// Tells the receiver that the sending web view will perform the specified drag-source action.
func (d *WebUIDelegate) SetWebViewWillPerformDragSourceActionFromPointWithPasteboard(f func(webView IWebView, action WebDragSourceAction, point vision.Point, pasteboard appkit.Pasteboard)) {
	d._WebViewWillPerformDragSourceActionFromPointWithPasteboard = f
}

// SetWebViewAreToolbarsVisible sets the handler for the WebViewAreToolbarsVisible delegate method.
//
// Returns a Boolean value indicating whether any toolbars are visible in a web view’s window.
func (d *WebUIDelegate) SetWebViewAreToolbarsVisible(f func(sender IWebView) bool) {
	d._WebViewAreToolbarsVisible = f
}

// SetWebViewClose sets the handler for the WebViewClose delegate method.
//
// Closes a web view in a window.
func (d *WebUIDelegate) SetWebViewClose(f func(sender IWebView)) {
	d._WebViewClose = f
}

// SetWebViewFirstResponder sets the handler for the WebViewFirstResponder delegate method.
//
// Returns the first responder of the web view’s window.
func (d *WebUIDelegate) SetWebViewFirstResponder(f func(sender IWebView) appkit.Responder) {
	d._WebViewFirstResponder = f
}

// SetWebViewFocus sets the handler for the WebViewFocus delegate method.
//
// Brings a web view’s window to the front and makes it the active window.
func (d *WebUIDelegate) SetWebViewFocus(f func(sender IWebView)) {
	d._WebViewFocus = f
}

// SetWebViewFooterHeight sets the handler for the WebViewFooterHeight delegate method.
//
// Returns the height of the web view’s printed page footer.
func (d *WebUIDelegate) SetWebViewFooterHeight(f func(sender IWebView) float32) {
	d._WebViewFooterHeight = f
}

// SetWebViewFrame sets the handler for the WebViewFrame delegate method.
//
// Returns the frame rectangle of a web view’s window.
func (d *WebUIDelegate) SetWebViewFrame(f func(sender IWebView) Rect) {
	d._WebViewFrame = f
}

// SetWebViewHeaderHeight sets the handler for the WebViewHeaderHeight delegate method.
//
// Returns the height of the web view’s printed page header.
func (d *WebUIDelegate) SetWebViewHeaderHeight(f func(sender IWebView) float32) {
	d._WebViewHeaderHeight = f
}

// SetWebViewIsResizable sets the handler for the WebViewIsResizable delegate method.
//
// Returns a Boolean value indicating whether a web view’s window can be resized.
func (d *WebUIDelegate) SetWebViewIsResizable(f func(sender IWebView) bool) {
	d._WebViewIsResizable = f
}

// SetWebViewIsStatusBarVisible sets the handler for the WebViewIsStatusBarVisible delegate method.
//
// Returns a Boolean value indicating whether the status bar in a web view’s window is visible.
func (d *WebUIDelegate) SetWebViewIsStatusBarVisible(f func(sender IWebView) bool) {
	d._WebViewIsStatusBarVisible = f
}

// SetWebViewRunModal sets the handler for the WebViewRunModal delegate method.
//
// Displays a web view in a modal window.
func (d *WebUIDelegate) SetWebViewRunModal(f func(sender IWebView)) {
	d._WebViewRunModal = f
}

// SetWebViewShow sets the handler for the WebViewShow delegate method.
//
// Displays a web view’s window and moves it to the front.
func (d *WebUIDelegate) SetWebViewShow(f func(sender IWebView)) {
	d._WebViewShow = f
}

// SetWebViewStatusText sets the handler for the WebViewStatusText delegate method.
//
// Returns the current status message from a web view’s window.
func (d *WebUIDelegate) SetWebViewStatusText(f func(sender IWebView) foundation.String) {
	d._WebViewStatusText = f
}

// SetWebViewUnfocus sets the handler for the WebViewUnfocus delegate method.
//
// Relinquishes focus on a web view’s window.
func (d *WebUIDelegate) SetWebViewUnfocus(f func(sender IWebView)) {
	d._WebViewUnfocus = f
}

// SetWebViewRunJavaScriptAlertPanelWithMessage sets the handler for the WebViewRunJavaScriptAlertPanelWithMessage delegate method.
//
// Displays a JavaScript alert panel.
func (d *WebUIDelegate) SetWebViewRunJavaScriptAlertPanelWithMessage(f func(sender IWebView, message objc.IObject /* cross-framework: NSString */)) {
	d._WebViewRunJavaScriptAlertPanelWithMessage = f
}

// SetWebViewRunJavaScriptConfirmPanelWithMessage sets the handler for the WebViewRunJavaScriptConfirmPanelWithMessage delegate method.
//
// Displays a JavaScript confirm panel.
func (d *WebUIDelegate) SetWebViewRunJavaScriptConfirmPanelWithMessage(f func(sender IWebView, message objc.IObject /* cross-framework: NSString */) bool) {
	d._WebViewRunJavaScriptConfirmPanelWithMessage = f
}

// SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultText sets the handler for the WebViewRunJavaScriptTextInputPanelWithPromptDefaultText delegate method.
//
// Displays a JavaScript text input panel and returns the entered text.
func (d *WebUIDelegate) SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultText(f func(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */) foundation.String) {
	d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultText = f
}

// SetWebViewSetContentRect sets the handler for the WebViewSetContentRect delegate method.
//
// Sets the window’s content view frame to the specified content rectangle.
func (d *WebUIDelegate) SetWebViewSetContentRect(f func(sender IWebView, frame Rect /* not a class type */)) {
	d._WebViewSetContentRect = f
}

// SetWebViewContentRect sets the handler for the WebViewContentRect delegate method.
//
// Returns a web view window’s content rectangle.
func (d *WebUIDelegate) SetWebViewContentRect(f func(sender IWebView) Rect) {
	d._WebViewContentRect = f
}

// WebViewContextMenuItemsForElementDefaultMenuItems implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewContextMenuItemsForElementDefaultMenuItems(sender IWebView, element objc.IObject /* cross-framework: NSDictionary */, defaultMenuItems objc.IObject /* cross-framework: NSArray */) foundation.Array {
	if d._WebViewContextMenuItemsForElementDefaultMenuItems != nil {
		return d._WebViewContextMenuItemsForElementDefaultMenuItems(sender, element, defaultMenuItems)
	}
	var zero foundation.Array
	return zero
}

// HasWebViewContextMenuItemsForElementDefaultMenuItems returns true if a handler for WebViewContextMenuItemsForElementDefaultMenuItems has been set.
func (d *WebUIDelegate) HasWebViewContextMenuItemsForElementDefaultMenuItems() bool {
	return d._WebViewContextMenuItemsForElementDefaultMenuItems != nil
}

// WebViewCreateWebViewModalDialogWithRequest implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewCreateWebViewModalDialogWithRequest(sender IWebView, request foundation.URLRequest) WebView {
	if d._WebViewCreateWebViewModalDialogWithRequest != nil {
		return d._WebViewCreateWebViewModalDialogWithRequest(sender, request)
	}
	var zero WebView
	return zero
}

// HasWebViewCreateWebViewModalDialogWithRequest returns true if a handler for WebViewCreateWebViewModalDialogWithRequest has been set.
func (d *WebUIDelegate) HasWebViewCreateWebViewModalDialogWithRequest() bool {
	return d._WebViewCreateWebViewModalDialogWithRequest != nil
}

// WebViewCreateWebViewWithRequest implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewCreateWebViewWithRequest(sender IWebView, request foundation.URLRequest) WebView {
	if d._WebViewCreateWebViewWithRequest != nil {
		return d._WebViewCreateWebViewWithRequest(sender, request)
	}
	var zero WebView
	return zero
}

// HasWebViewCreateWebViewWithRequest returns true if a handler for WebViewCreateWebViewWithRequest has been set.
func (d *WebUIDelegate) HasWebViewCreateWebViewWithRequest() bool {
	return d._WebViewCreateWebViewWithRequest != nil
}

// WebViewDragDestinationActionMaskForDraggingInfo implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewDragDestinationActionMaskForDraggingInfo(webView IWebView, draggingInfo unsafe.Pointer) uint {
	if d._WebViewDragDestinationActionMaskForDraggingInfo != nil {
		return d._WebViewDragDestinationActionMaskForDraggingInfo(webView, draggingInfo)
	}
	var zero uint
	return zero
}

// HasWebViewDragDestinationActionMaskForDraggingInfo returns true if a handler for WebViewDragDestinationActionMaskForDraggingInfo has been set.
func (d *WebUIDelegate) HasWebViewDragDestinationActionMaskForDraggingInfo() bool {
	return d._WebViewDragDestinationActionMaskForDraggingInfo != nil
}

// WebViewDragSourceActionMaskForPoint implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewDragSourceActionMaskForPoint(webView IWebView, point vision.Point) uint {
	if d._WebViewDragSourceActionMaskForPoint != nil {
		return d._WebViewDragSourceActionMaskForPoint(webView, point)
	}
	var zero uint
	return zero
}

// HasWebViewDragSourceActionMaskForPoint returns true if a handler for WebViewDragSourceActionMaskForPoint has been set.
func (d *WebUIDelegate) HasWebViewDragSourceActionMaskForPoint() bool {
	return d._WebViewDragSourceActionMaskForPoint != nil
}

// WebViewDrawFooterInRect implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewDrawFooterInRect(sender IWebView, rect Rect /* not a class type */) {
	if d._WebViewDrawFooterInRect != nil {
		d._WebViewDrawFooterInRect(sender, rect)
	}
}

// HasWebViewDrawFooterInRect returns true if a handler for WebViewDrawFooterInRect has been set.
func (d *WebUIDelegate) HasWebViewDrawFooterInRect() bool {
	return d._WebViewDrawFooterInRect != nil
}

// WebViewDrawHeaderInRect implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewDrawHeaderInRect(sender IWebView, rect Rect /* not a class type */) {
	if d._WebViewDrawHeaderInRect != nil {
		d._WebViewDrawHeaderInRect(sender, rect)
	}
}

// HasWebViewDrawHeaderInRect returns true if a handler for WebViewDrawHeaderInRect has been set.
func (d *WebUIDelegate) HasWebViewDrawHeaderInRect() bool {
	return d._WebViewDrawHeaderInRect != nil
}

// WebViewMakeFirstResponder implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewMakeFirstResponder(sender IWebView, responder appkit.Responder) {
	if d._WebViewMakeFirstResponder != nil {
		d._WebViewMakeFirstResponder(sender, responder)
	}
}

// HasWebViewMakeFirstResponder returns true if a handler for WebViewMakeFirstResponder has been set.
func (d *WebUIDelegate) HasWebViewMakeFirstResponder() bool {
	return d._WebViewMakeFirstResponder != nil
}

// WebViewMouseDidMoveOverElementModifierFlags implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewMouseDidMoveOverElementModifierFlags(sender IWebView, elementInformation objc.IObject /* cross-framework: NSDictionary */, modifierFlags uint) {
	if d._WebViewMouseDidMoveOverElementModifierFlags != nil {
		d._WebViewMouseDidMoveOverElementModifierFlags(sender, elementInformation, modifierFlags)
	}
}

// HasWebViewMouseDidMoveOverElementModifierFlags returns true if a handler for WebViewMouseDidMoveOverElementModifierFlags has been set.
func (d *WebUIDelegate) HasWebViewMouseDidMoveOverElementModifierFlags() bool {
	return d._WebViewMouseDidMoveOverElementModifierFlags != nil
}

// WebViewPrintFrameView implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewPrintFrameView(sender IWebView, frameView IWebFrameView) {
	if d._WebViewPrintFrameView != nil {
		d._WebViewPrintFrameView(sender, frameView)
	}
}

// HasWebViewPrintFrameView returns true if a handler for WebViewPrintFrameView has been set.
func (d *WebUIDelegate) HasWebViewPrintFrameView() bool {
	return d._WebViewPrintFrameView != nil
}

// WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool {
	if d._WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame != nil {
		return d._WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame(sender, message, frame)
	}
	var zero bool
	return zero
}

// HasWebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame returns true if a handler for WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame has been set.
func (d *WebUIDelegate) HasWebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame() bool {
	return d._WebViewRunBeforeUnloadConfirmPanelWithMessageInitiatedByFrame != nil
}

// WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) {
	if d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame != nil {
		d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame(sender, message, frame)
	}
}

// HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame returns true if a handler for WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame() bool {
	return d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrame != nil
}

// WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame(sender IWebView, message objc.IObject /* cross-framework: NSString */, frame IWebFrame) bool {
	if d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame != nil {
		return d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame(sender, message, frame)
	}
	var zero bool
	return zero
}

// HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame returns true if a handler for WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame() bool {
	return d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrame != nil
}

// WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWebFrame) foundation.String {
	if d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame != nil {
		return d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame(sender, prompt, defaultText, frame)
	}
	var zero foundation.String
	return zero
}

// HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame returns true if a handler for WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame() bool {
	return d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrame != nil
}

// WebViewRunOpenPanelForFileButtonWithResultListener implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunOpenPanelForFileButtonWithResultListener(sender IWebView, resultListener unsafe.Pointer) {
	if d._WebViewRunOpenPanelForFileButtonWithResultListener != nil {
		d._WebViewRunOpenPanelForFileButtonWithResultListener(sender, resultListener)
	}
}

// HasWebViewRunOpenPanelForFileButtonWithResultListener returns true if a handler for WebViewRunOpenPanelForFileButtonWithResultListener has been set.
func (d *WebUIDelegate) HasWebViewRunOpenPanelForFileButtonWithResultListener() bool {
	return d._WebViewRunOpenPanelForFileButtonWithResultListener != nil
}

// WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles(sender IWebView, resultListener unsafe.Pointer, allowMultipleFiles bool) {
	if d._WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles != nil {
		d._WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles(sender, resultListener, allowMultipleFiles)
	}
}

// HasWebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles returns true if a handler for WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles has been set.
func (d *WebUIDelegate) HasWebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles() bool {
	return d._WebViewRunOpenPanelForFileButtonWithResultListenerAllowMultipleFiles != nil
}

// WebViewSetFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetFrame(sender IWebView, frame Rect /* not a class type */) {
	if d._WebViewSetFrame != nil {
		d._WebViewSetFrame(sender, frame)
	}
}

// HasWebViewSetFrame returns true if a handler for WebViewSetFrame has been set.
func (d *WebUIDelegate) HasWebViewSetFrame() bool {
	return d._WebViewSetFrame != nil
}

// WebViewSetResizable implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetResizable(sender IWebView, resizable bool) {
	if d._WebViewSetResizable != nil {
		d._WebViewSetResizable(sender, resizable)
	}
}

// HasWebViewSetResizable returns true if a handler for WebViewSetResizable has been set.
func (d *WebUIDelegate) HasWebViewSetResizable() bool {
	return d._WebViewSetResizable != nil
}

// WebViewSetStatusBarVisible implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetStatusBarVisible(sender IWebView, visible bool) {
	if d._WebViewSetStatusBarVisible != nil {
		d._WebViewSetStatusBarVisible(sender, visible)
	}
}

// HasWebViewSetStatusBarVisible returns true if a handler for WebViewSetStatusBarVisible has been set.
func (d *WebUIDelegate) HasWebViewSetStatusBarVisible() bool {
	return d._WebViewSetStatusBarVisible != nil
}

// WebViewSetStatusText implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetStatusText(sender IWebView, text objc.IObject /* cross-framework: NSString */) {
	if d._WebViewSetStatusText != nil {
		d._WebViewSetStatusText(sender, text)
	}
}

// HasWebViewSetStatusText returns true if a handler for WebViewSetStatusText has been set.
func (d *WebUIDelegate) HasWebViewSetStatusText() bool {
	return d._WebViewSetStatusText != nil
}

// WebViewSetToolbarsVisible implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetToolbarsVisible(sender IWebView, visible bool) {
	if d._WebViewSetToolbarsVisible != nil {
		d._WebViewSetToolbarsVisible(sender, visible)
	}
}

// HasWebViewSetToolbarsVisible returns true if a handler for WebViewSetToolbarsVisible has been set.
func (d *WebUIDelegate) HasWebViewSetToolbarsVisible() bool {
	return d._WebViewSetToolbarsVisible != nil
}

// WebViewShouldPerformActionFromSender implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewShouldPerformActionFromSender(webView IWebView, action objc.SEL, sender objc.IObject) bool {
	if d._WebViewShouldPerformActionFromSender != nil {
		return d._WebViewShouldPerformActionFromSender(webView, action, sender)
	}
	var zero bool
	return zero
}

// HasWebViewShouldPerformActionFromSender returns true if a handler for WebViewShouldPerformActionFromSender has been set.
func (d *WebUIDelegate) HasWebViewShouldPerformActionFromSender() bool {
	return d._WebViewShouldPerformActionFromSender != nil
}

// WebViewValidateUserInterfaceItemDefaultValidation implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewValidateUserInterfaceItemDefaultValidation(webView IWebView, item unsafe.Pointer, defaultValidation bool) bool {
	if d._WebViewValidateUserInterfaceItemDefaultValidation != nil {
		return d._WebViewValidateUserInterfaceItemDefaultValidation(webView, item, defaultValidation)
	}
	var zero bool
	return zero
}

// HasWebViewValidateUserInterfaceItemDefaultValidation returns true if a handler for WebViewValidateUserInterfaceItemDefaultValidation has been set.
func (d *WebUIDelegate) HasWebViewValidateUserInterfaceItemDefaultValidation() bool {
	return d._WebViewValidateUserInterfaceItemDefaultValidation != nil
}

// WebViewWillPerformDragDestinationActionForDraggingInfo implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewWillPerformDragDestinationActionForDraggingInfo(webView IWebView, action WebDragDestinationAction, draggingInfo unsafe.Pointer) {
	if d._WebViewWillPerformDragDestinationActionForDraggingInfo != nil {
		d._WebViewWillPerformDragDestinationActionForDraggingInfo(webView, action, draggingInfo)
	}
}

// HasWebViewWillPerformDragDestinationActionForDraggingInfo returns true if a handler for WebViewWillPerformDragDestinationActionForDraggingInfo has been set.
func (d *WebUIDelegate) HasWebViewWillPerformDragDestinationActionForDraggingInfo() bool {
	return d._WebViewWillPerformDragDestinationActionForDraggingInfo != nil
}

// WebViewWillPerformDragSourceActionFromPointWithPasteboard implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewWillPerformDragSourceActionFromPointWithPasteboard(webView IWebView, action WebDragSourceAction, point vision.Point, pasteboard appkit.Pasteboard) {
	if d._WebViewWillPerformDragSourceActionFromPointWithPasteboard != nil {
		d._WebViewWillPerformDragSourceActionFromPointWithPasteboard(webView, action, point, pasteboard)
	}
}

// HasWebViewWillPerformDragSourceActionFromPointWithPasteboard returns true if a handler for WebViewWillPerformDragSourceActionFromPointWithPasteboard has been set.
func (d *WebUIDelegate) HasWebViewWillPerformDragSourceActionFromPointWithPasteboard() bool {
	return d._WebViewWillPerformDragSourceActionFromPointWithPasteboard != nil
}

// WebViewAreToolbarsVisible implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewAreToolbarsVisible(sender IWebView) bool {
	if d._WebViewAreToolbarsVisible != nil {
		return d._WebViewAreToolbarsVisible(sender)
	}
	var zero bool
	return zero
}

// HasWebViewAreToolbarsVisible returns true if a handler for WebViewAreToolbarsVisible has been set.
func (d *WebUIDelegate) HasWebViewAreToolbarsVisible() bool {
	return d._WebViewAreToolbarsVisible != nil
}

// WebViewClose implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewClose(sender IWebView) {
	if d._WebViewClose != nil {
		d._WebViewClose(sender)
	}
}

// HasWebViewClose returns true if a handler for WebViewClose has been set.
func (d *WebUIDelegate) HasWebViewClose() bool {
	return d._WebViewClose != nil
}

// WebViewFirstResponder implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewFirstResponder(sender IWebView) appkit.Responder {
	if d._WebViewFirstResponder != nil {
		return d._WebViewFirstResponder(sender)
	}
	var zero appkit.Responder
	return zero
}

// HasWebViewFirstResponder returns true if a handler for WebViewFirstResponder has been set.
func (d *WebUIDelegate) HasWebViewFirstResponder() bool {
	return d._WebViewFirstResponder != nil
}

// WebViewFocus implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewFocus(sender IWebView) {
	if d._WebViewFocus != nil {
		d._WebViewFocus(sender)
	}
}

// HasWebViewFocus returns true if a handler for WebViewFocus has been set.
func (d *WebUIDelegate) HasWebViewFocus() bool {
	return d._WebViewFocus != nil
}

// WebViewFooterHeight implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewFooterHeight(sender IWebView) float32 {
	if d._WebViewFooterHeight != nil {
		return d._WebViewFooterHeight(sender)
	}
	var zero float32
	return zero
}

// HasWebViewFooterHeight returns true if a handler for WebViewFooterHeight has been set.
func (d *WebUIDelegate) HasWebViewFooterHeight() bool {
	return d._WebViewFooterHeight != nil
}

// WebViewFrame implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewFrame(sender IWebView) Rect {
	if d._WebViewFrame != nil {
		return d._WebViewFrame(sender)
	}
	var zero Rect
	return zero
}

// HasWebViewFrame returns true if a handler for WebViewFrame has been set.
func (d *WebUIDelegate) HasWebViewFrame() bool {
	return d._WebViewFrame != nil
}

// WebViewHeaderHeight implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewHeaderHeight(sender IWebView) float32 {
	if d._WebViewHeaderHeight != nil {
		return d._WebViewHeaderHeight(sender)
	}
	var zero float32
	return zero
}

// HasWebViewHeaderHeight returns true if a handler for WebViewHeaderHeight has been set.
func (d *WebUIDelegate) HasWebViewHeaderHeight() bool {
	return d._WebViewHeaderHeight != nil
}

// WebViewIsResizable implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewIsResizable(sender IWebView) bool {
	if d._WebViewIsResizable != nil {
		return d._WebViewIsResizable(sender)
	}
	var zero bool
	return zero
}

// HasWebViewIsResizable returns true if a handler for WebViewIsResizable has been set.
func (d *WebUIDelegate) HasWebViewIsResizable() bool {
	return d._WebViewIsResizable != nil
}

// WebViewIsStatusBarVisible implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewIsStatusBarVisible(sender IWebView) bool {
	if d._WebViewIsStatusBarVisible != nil {
		return d._WebViewIsStatusBarVisible(sender)
	}
	var zero bool
	return zero
}

// HasWebViewIsStatusBarVisible returns true if a handler for WebViewIsStatusBarVisible has been set.
func (d *WebUIDelegate) HasWebViewIsStatusBarVisible() bool {
	return d._WebViewIsStatusBarVisible != nil
}

// WebViewRunModal implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunModal(sender IWebView) {
	if d._WebViewRunModal != nil {
		d._WebViewRunModal(sender)
	}
}

// HasWebViewRunModal returns true if a handler for WebViewRunModal has been set.
func (d *WebUIDelegate) HasWebViewRunModal() bool {
	return d._WebViewRunModal != nil
}

// WebViewShow implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewShow(sender IWebView) {
	if d._WebViewShow != nil {
		d._WebViewShow(sender)
	}
}

// HasWebViewShow returns true if a handler for WebViewShow has been set.
func (d *WebUIDelegate) HasWebViewShow() bool {
	return d._WebViewShow != nil
}

// WebViewStatusText implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewStatusText(sender IWebView) foundation.String {
	if d._WebViewStatusText != nil {
		return d._WebViewStatusText(sender)
	}
	var zero foundation.String
	return zero
}

// HasWebViewStatusText returns true if a handler for WebViewStatusText has been set.
func (d *WebUIDelegate) HasWebViewStatusText() bool {
	return d._WebViewStatusText != nil
}

// WebViewUnfocus implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewUnfocus(sender IWebView) {
	if d._WebViewUnfocus != nil {
		d._WebViewUnfocus(sender)
	}
}

// HasWebViewUnfocus returns true if a handler for WebViewUnfocus has been set.
func (d *WebUIDelegate) HasWebViewUnfocus() bool {
	return d._WebViewUnfocus != nil
}

// WebViewRunJavaScriptAlertPanelWithMessage implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptAlertPanelWithMessage(sender IWebView, message objc.IObject /* cross-framework: NSString */) {
	if d._WebViewRunJavaScriptAlertPanelWithMessage != nil {
		d._WebViewRunJavaScriptAlertPanelWithMessage(sender, message)
	}
}

// HasWebViewRunJavaScriptAlertPanelWithMessage returns true if a handler for WebViewRunJavaScriptAlertPanelWithMessage has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptAlertPanelWithMessage() bool {
	return d._WebViewRunJavaScriptAlertPanelWithMessage != nil
}

// WebViewRunJavaScriptConfirmPanelWithMessage implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptConfirmPanelWithMessage(sender IWebView, message objc.IObject /* cross-framework: NSString */) bool {
	if d._WebViewRunJavaScriptConfirmPanelWithMessage != nil {
		return d._WebViewRunJavaScriptConfirmPanelWithMessage(sender, message)
	}
	var zero bool
	return zero
}

// HasWebViewRunJavaScriptConfirmPanelWithMessage returns true if a handler for WebViewRunJavaScriptConfirmPanelWithMessage has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptConfirmPanelWithMessage() bool {
	return d._WebViewRunJavaScriptConfirmPanelWithMessage != nil
}

// WebViewRunJavaScriptTextInputPanelWithPromptDefaultText implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewRunJavaScriptTextInputPanelWithPromptDefaultText(sender IWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */) foundation.String {
	if d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultText != nil {
		return d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultText(sender, prompt, defaultText)
	}
	var zero foundation.String
	return zero
}

// HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultText returns true if a handler for WebViewRunJavaScriptTextInputPanelWithPromptDefaultText has been set.
func (d *WebUIDelegate) HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultText() bool {
	return d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultText != nil
}

// WebViewSetContentRect implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewSetContentRect(sender IWebView, frame Rect /* not a class type */) {
	if d._WebViewSetContentRect != nil {
		d._WebViewSetContentRect(sender, frame)
	}
}

// HasWebViewSetContentRect returns true if a handler for WebViewSetContentRect has been set.
func (d *WebUIDelegate) HasWebViewSetContentRect() bool {
	return d._WebViewSetContentRect != nil
}

// WebViewContentRect implements the PWebUIDelegate interface.
func (d *WebUIDelegate) WebViewContentRect(sender IWebView) Rect {
	if d._WebViewContentRect != nil {
		return d._WebViewContentRect(sender)
	}
	var zero Rect
	return zero
}

// HasWebViewContentRect returns true if a handler for WebViewContentRect has been set.
func (d *WebUIDelegate) HasWebViewContentRect() bool {
	return d._WebViewContentRect != nil
}
