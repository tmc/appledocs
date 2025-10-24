// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objc"
)

// PUIDelegate is the WKUIDelegate protocol interface.
//
// The methods for presenting native user interface elements on behalf of a webpage.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.webkit/documentation/WebKit/WKUIDelegate
type PUIDelegate interface {
	// Optional methods
	WebViewCommitPreviewingViewController(webView IWKWebView, previewingViewController appkit.ViewController)
	HasWebViewCommitPreviewingViewController() bool
	WebViewContextMenuConfigurationForElementCompletionHandler(webView IWKWebView, elementInfo IWKContextMenuElementInfo, completionHandler func(unsafe.Pointer))
	HasWebViewContextMenuConfigurationForElementCompletionHandler() bool
	WebViewContextMenuDidEndForElement(webView IWKWebView, elementInfo IWKContextMenuElementInfo)
	HasWebViewContextMenuDidEndForElement() bool
	WebViewContextMenuForElementWillCommitWithAnimator(webView IWKWebView, elementInfo IWKContextMenuElementInfo, animator unsafe.Pointer)
	HasWebViewContextMenuForElementWillCommitWithAnimator() bool
	WebViewContextMenuWillPresentForElement(webView IWKWebView, elementInfo IWKContextMenuElementInfo)
	HasWebViewContextMenuWillPresentForElement() bool
	WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView IWKWebView, configuration IWKWebViewConfiguration, navigationAction IWKNavigationAction, windowFeatures IWKWindowFeatures) WebView
	HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool
	WebViewPreviewingViewControllerForElementDefaultActions(webView IWKWebView, elementInfo IWKPreviewElementInfo, previewActions []objc.ID) appkit.ViewController
	HasWebViewPreviewingViewControllerForElementDefaultActions() bool
	WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, decisionHandler func(unsafe.Pointer))
	HasWebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler() bool
	WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, type_ MediaCaptureType, decisionHandler func(unsafe.Pointer))
	HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool
	WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func())
	HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool
	WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))
	HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool
	WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView IWKWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))
	HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool
	WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView IWKWebView, parameters IWKOpenPanelParameters, frame IWKFrameInfo, completionHandler func([]unsafe.Pointer))
	HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler() bool
	WebViewShouldPreviewElement(webView IWKWebView, elementInfo IWKPreviewElementInfo) bool
	HasWebViewShouldPreviewElement() bool
	WebViewShowLockdownModeFirstUseMessageCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))
	HasWebViewShowLockdownModeFirstUseMessageCompletionHandler() bool
	WebViewWillDismissEditMenuWithAnimator(webView IWKWebView, animator unsafe.Pointer)
	HasWebViewWillDismissEditMenuWithAnimator() bool
	WebViewWillPresentEditMenuWithAnimator(webView IWKWebView, animator unsafe.Pointer)
	HasWebViewWillPresentEditMenuWithAnimator() bool
	WebViewDidClose(webView IWKWebView)
	HasWebViewDidClose() bool
}

// UIDelegate is a delegate implementation builder for the PUIDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type UIDelegate struct {
	_WebViewCommitPreviewingViewController                                                      func(webView IWKWebView, previewingViewController appkit.ViewController)
	_WebViewContextMenuConfigurationForElementCompletionHandler                                 func(webView IWKWebView, elementInfo IWKContextMenuElementInfo, completionHandler func(unsafe.Pointer))
	_WebViewContextMenuDidEndForElement                                                         func(webView IWKWebView, elementInfo IWKContextMenuElementInfo)
	_WebViewContextMenuForElementWillCommitWithAnimator                                         func(webView IWKWebView, elementInfo IWKContextMenuElementInfo, animator unsafe.Pointer)
	_WebViewContextMenuWillPresentForElement                                                    func(webView IWKWebView, elementInfo IWKContextMenuElementInfo)
	_WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures                     func(webView IWKWebView, configuration IWKWebViewConfiguration, navigationAction IWKNavigationAction, windowFeatures IWKWindowFeatures) WebView
	_WebViewPreviewingViewControllerForElementDefaultActions                                    func(webView IWKWebView, elementInfo IWKPreviewElementInfo, previewActions []objc.ID) appkit.ViewController
	_WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler func(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, decisionHandler func(unsafe.Pointer))
	_WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler           func(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, type_ MediaCaptureType, decisionHandler func(unsafe.Pointer))
	_WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler                 func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func())
	_WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler               func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))
	_WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler   func(webView IWKWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))
	_WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler                         func(webView IWKWebView, parameters IWKOpenPanelParameters, frame IWKFrameInfo, completionHandler func([]unsafe.Pointer))
	_WebViewShouldPreviewElement                                                                func(webView IWKWebView, elementInfo IWKPreviewElementInfo) bool
	_WebViewShowLockdownModeFirstUseMessageCompletionHandler                                    func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))
	_WebViewWillDismissEditMenuWithAnimator                                                     func(webView IWKWebView, animator unsafe.Pointer)
	_WebViewWillPresentEditMenuWithAnimator                                                     func(webView IWKWebView, animator unsafe.Pointer)
	_WebViewDidClose                                                                            func(webView IWKWebView)
}

// SetWebViewCommitPreviewingViewController sets the handler for the WebViewCommitPreviewingViewController delegate method.
//
// Called when the user performs a pop action on the preview.
func (d *UIDelegate) SetWebViewCommitPreviewingViewController(f func(webView IWKWebView, previewingViewController appkit.ViewController)) {
	d._WebViewCommitPreviewingViewController = f
}

// SetWebViewContextMenuConfigurationForElementCompletionHandler sets the handler for the WebViewContextMenuConfigurationForElementCompletionHandler delegate method.
//
// Tells the delegate that a contextual menu interaction began.
func (d *UIDelegate) SetWebViewContextMenuConfigurationForElementCompletionHandler(f func(webView IWKWebView, elementInfo IWKContextMenuElementInfo, completionHandler func(unsafe.Pointer))) {
	d._WebViewContextMenuConfigurationForElementCompletionHandler = f
}

// SetWebViewContextMenuDidEndForElement sets the handler for the WebViewContextMenuDidEndForElement delegate method.
//
// Tells the delegate that the web view dismissed the contextual menu for the specified element.
func (d *UIDelegate) SetWebViewContextMenuDidEndForElement(f func(webView IWKWebView, elementInfo IWKContextMenuElementInfo)) {
	d._WebViewContextMenuDidEndForElement = f
}

// SetWebViewContextMenuForElementWillCommitWithAnimator sets the handler for the WebViewContextMenuForElementWillCommitWithAnimator delegate method.
//
// Provides the delegate with the animator object that the web view uses to display the contextual menu.
func (d *UIDelegate) SetWebViewContextMenuForElementWillCommitWithAnimator(f func(webView IWKWebView, elementInfo IWKContextMenuElementInfo, animator unsafe.Pointer)) {
	d._WebViewContextMenuForElementWillCommitWithAnimator = f
}

// SetWebViewContextMenuWillPresentForElement sets the handler for the WebViewContextMenuWillPresentForElement delegate method.
//
// Tells the delegate that the web view is about to present the contextual menu for the specified element.
func (d *UIDelegate) SetWebViewContextMenuWillPresentForElement(f func(webView IWKWebView, elementInfo IWKContextMenuElementInfo)) {
	d._WebViewContextMenuWillPresentForElement = f
}

// SetWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures sets the handler for the WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures delegate method.
//
// Creates a new web view.
func (d *UIDelegate) SetWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(f func(webView IWKWebView, configuration IWKWebViewConfiguration, navigationAction IWKNavigationAction, windowFeatures IWKWindowFeatures) WebView) {
	d._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures = f
}

// SetWebViewPreviewingViewControllerForElementDefaultActions sets the handler for the WebViewPreviewingViewControllerForElementDefaultActions delegate method.
//
// Called when the user performs a peek action.
func (d *UIDelegate) SetWebViewPreviewingViewControllerForElementDefaultActions(f func(webView IWKWebView, elementInfo IWKPreviewElementInfo, previewActions []objc.ID) appkit.ViewController) {
	d._WebViewPreviewingViewControllerForElementDefaultActions = f
}

// SetWebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler sets the handler for the WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler delegate method.
//
// Determines whether a web resource, which the security origin object describes, can access the device’s orientation and motion.
func (d *UIDelegate) SetWebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler(f func(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, decisionHandler func(unsafe.Pointer))) {
	d._WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler = f
}

// SetWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler sets the handler for the WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler delegate method.
//
// Determines whether a web resource, which the security origin object describes, can access to the device’s microphone audio and camera video.
func (d *UIDelegate) SetWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(f func(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, type_ MediaCaptureType, decisionHandler func(unsafe.Pointer))) {
	d._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler = f
}

// SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler sets the handler for the WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler delegate method.
//
// Displays a JavaScript alert panel.
func (d *UIDelegate) SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func())) {
	d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler = f
}

// SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler sets the handler for the WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler delegate method.
//
// Displays a JavaScript confirm panel.
func (d *UIDelegate) SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))) {
	d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler = f
}

// SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler sets the handler for the WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler delegate method.
//
// Displays a JavaScript text input panel.
func (d *UIDelegate) SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(f func(webView IWKWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer))) {
	d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler = f
}

// SetWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler sets the handler for the WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler delegate method.
//
// Displays a file upload panel.
func (d *UIDelegate) SetWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(f func(webView IWKWebView, parameters IWKOpenPanelParameters, frame IWKFrameInfo, completionHandler func([]unsafe.Pointer))) {
	d._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler = f
}

// SetWebViewShouldPreviewElement sets the handler for the WebViewShouldPreviewElement delegate method.
//
// Determines whether the given element should show a preview.
func (d *UIDelegate) SetWebViewShouldPreviewElement(f func(webView IWKWebView, elementInfo IWKPreviewElementInfo) bool) {
	d._WebViewShouldPreviewElement = f
}

// SetWebViewShowLockdownModeFirstUseMessageCompletionHandler sets the handler for the WebViewShowLockdownModeFirstUseMessageCompletionHandler delegate method.
//
// Displays a custom Lockdown Mode first use message.
func (d *UIDelegate) SetWebViewShowLockdownModeFirstUseMessageCompletionHandler(f func(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))) {
	d._WebViewShowLockdownModeFirstUseMessageCompletionHandler = f
}

// SetWebViewWillDismissEditMenuWithAnimator sets the handler for the WebViewWillDismissEditMenuWithAnimator delegate method.
//
// Tells the delegate that the web view is about to dismiss an edit menu.
func (d *UIDelegate) SetWebViewWillDismissEditMenuWithAnimator(f func(webView IWKWebView, animator unsafe.Pointer)) {
	d._WebViewWillDismissEditMenuWithAnimator = f
}

// SetWebViewWillPresentEditMenuWithAnimator sets the handler for the WebViewWillPresentEditMenuWithAnimator delegate method.
//
// Tells the delegate that the web view is about to present an edit menu.
func (d *UIDelegate) SetWebViewWillPresentEditMenuWithAnimator(f func(webView IWKWebView, animator unsafe.Pointer)) {
	d._WebViewWillPresentEditMenuWithAnimator = f
}

// SetWebViewDidClose sets the handler for the WebViewDidClose delegate method.
//
// Notifies your app that the DOM window closed successfully.
func (d *UIDelegate) SetWebViewDidClose(f func(webView IWKWebView)) {
	d._WebViewDidClose = f
}

// WebViewCommitPreviewingViewController implements the PUIDelegate interface.
func (d *UIDelegate) WebViewCommitPreviewingViewController(webView IWKWebView, previewingViewController appkit.ViewController) {
	if d._WebViewCommitPreviewingViewController != nil {
		d._WebViewCommitPreviewingViewController(webView, previewingViewController)
	}
}

// HasWebViewCommitPreviewingViewController returns true if a handler for WebViewCommitPreviewingViewController has been set.
func (d *UIDelegate) HasWebViewCommitPreviewingViewController() bool {
	return d._WebViewCommitPreviewingViewController != nil
}

// WebViewContextMenuConfigurationForElementCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewContextMenuConfigurationForElementCompletionHandler(webView IWKWebView, elementInfo IWKContextMenuElementInfo, completionHandler func(unsafe.Pointer)) {
	if d._WebViewContextMenuConfigurationForElementCompletionHandler != nil {
		d._WebViewContextMenuConfigurationForElementCompletionHandler(webView, elementInfo, completionHandler)
	}
}

// HasWebViewContextMenuConfigurationForElementCompletionHandler returns true if a handler for WebViewContextMenuConfigurationForElementCompletionHandler has been set.
func (d *UIDelegate) HasWebViewContextMenuConfigurationForElementCompletionHandler() bool {
	return d._WebViewContextMenuConfigurationForElementCompletionHandler != nil
}

// WebViewContextMenuDidEndForElement implements the PUIDelegate interface.
func (d *UIDelegate) WebViewContextMenuDidEndForElement(webView IWKWebView, elementInfo IWKContextMenuElementInfo) {
	if d._WebViewContextMenuDidEndForElement != nil {
		d._WebViewContextMenuDidEndForElement(webView, elementInfo)
	}
}

// HasWebViewContextMenuDidEndForElement returns true if a handler for WebViewContextMenuDidEndForElement has been set.
func (d *UIDelegate) HasWebViewContextMenuDidEndForElement() bool {
	return d._WebViewContextMenuDidEndForElement != nil
}

// WebViewContextMenuForElementWillCommitWithAnimator implements the PUIDelegate interface.
func (d *UIDelegate) WebViewContextMenuForElementWillCommitWithAnimator(webView IWKWebView, elementInfo IWKContextMenuElementInfo, animator unsafe.Pointer) {
	if d._WebViewContextMenuForElementWillCommitWithAnimator != nil {
		d._WebViewContextMenuForElementWillCommitWithAnimator(webView, elementInfo, animator)
	}
}

// HasWebViewContextMenuForElementWillCommitWithAnimator returns true if a handler for WebViewContextMenuForElementWillCommitWithAnimator has been set.
func (d *UIDelegate) HasWebViewContextMenuForElementWillCommitWithAnimator() bool {
	return d._WebViewContextMenuForElementWillCommitWithAnimator != nil
}

// WebViewContextMenuWillPresentForElement implements the PUIDelegate interface.
func (d *UIDelegate) WebViewContextMenuWillPresentForElement(webView IWKWebView, elementInfo IWKContextMenuElementInfo) {
	if d._WebViewContextMenuWillPresentForElement != nil {
		d._WebViewContextMenuWillPresentForElement(webView, elementInfo)
	}
}

// HasWebViewContextMenuWillPresentForElement returns true if a handler for WebViewContextMenuWillPresentForElement has been set.
func (d *UIDelegate) HasWebViewContextMenuWillPresentForElement() bool {
	return d._WebViewContextMenuWillPresentForElement != nil
}

// WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures implements the PUIDelegate interface.
func (d *UIDelegate) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView IWKWebView, configuration IWKWebViewConfiguration, navigationAction IWKNavigationAction, windowFeatures IWKWindowFeatures) WebView {
	if d._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures != nil {
		return d._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView, configuration, navigationAction, windowFeatures)
	}
	var zero WebView
	return zero
}

// HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures returns true if a handler for WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures has been set.
func (d *UIDelegate) HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool {
	return d._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures != nil
}

// WebViewPreviewingViewControllerForElementDefaultActions implements the PUIDelegate interface.
func (d *UIDelegate) WebViewPreviewingViewControllerForElementDefaultActions(webView IWKWebView, elementInfo IWKPreviewElementInfo, previewActions []objc.ID) appkit.ViewController {
	if d._WebViewPreviewingViewControllerForElementDefaultActions != nil {
		return d._WebViewPreviewingViewControllerForElementDefaultActions(webView, elementInfo, previewActions)
	}
	var zero appkit.ViewController
	return zero
}

// HasWebViewPreviewingViewControllerForElementDefaultActions returns true if a handler for WebViewPreviewingViewControllerForElementDefaultActions has been set.
func (d *UIDelegate) HasWebViewPreviewingViewControllerForElementDefaultActions() bool {
	return d._WebViewPreviewingViewControllerForElementDefaultActions != nil
}

// WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, decisionHandler func(unsafe.Pointer)) {
	if d._WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler != nil {
		d._WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler(webView, origin, frame, decisionHandler)
	}
}

// HasWebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler returns true if a handler for WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler has been set.
func (d *UIDelegate) HasWebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler() bool {
	return d._WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler != nil
}

// WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, type_ MediaCaptureType, decisionHandler func(unsafe.Pointer)) {
	if d._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler != nil {
		d._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView, origin, frame, type_, decisionHandler)
	}
}

// HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler returns true if a handler for WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler has been set.
func (d *UIDelegate) HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool {
	return d._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler != nil
}

// WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func()) {
	if d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler != nil {
		d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
	}
}

// HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler returns true if a handler for WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler has been set.
func (d *UIDelegate) HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return d._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

// WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer)) {
	if d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler != nil {
		d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
	}
}

// HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler returns true if a handler for WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler has been set.
func (d *UIDelegate) HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return d._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

// WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView IWKWebView, prompt objc.IObject /* cross-framework: NSString */, defaultText objc.IObject /* cross-framework: NSString */, frame IWKFrameInfo, completionHandler func(unsafe.Pointer)) {
	if d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler != nil {
		d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView, prompt, defaultText, frame, completionHandler)
	}
}

// HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler returns true if a handler for WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler has been set.
func (d *UIDelegate) HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool {
	return d._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler != nil
}

// WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView IWKWebView, parameters IWKOpenPanelParameters, frame IWKFrameInfo, completionHandler func([]unsafe.Pointer)) {
	if d._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler != nil {
		d._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView, parameters, frame, completionHandler)
	}
}

// HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler returns true if a handler for WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler has been set.
func (d *UIDelegate) HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler() bool {
	return d._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler != nil
}

// WebViewShouldPreviewElement implements the PUIDelegate interface.
func (d *UIDelegate) WebViewShouldPreviewElement(webView IWKWebView, elementInfo IWKPreviewElementInfo) bool {
	if d._WebViewShouldPreviewElement != nil {
		return d._WebViewShouldPreviewElement(webView, elementInfo)
	}
	var zero bool
	return zero
}

// HasWebViewShouldPreviewElement returns true if a handler for WebViewShouldPreviewElement has been set.
func (d *UIDelegate) HasWebViewShouldPreviewElement() bool {
	return d._WebViewShouldPreviewElement != nil
}

// WebViewShowLockdownModeFirstUseMessageCompletionHandler implements the PUIDelegate interface.
func (d *UIDelegate) WebViewShowLockdownModeFirstUseMessageCompletionHandler(webView IWKWebView, message objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer)) {
	if d._WebViewShowLockdownModeFirstUseMessageCompletionHandler != nil {
		d._WebViewShowLockdownModeFirstUseMessageCompletionHandler(webView, message, completionHandler)
	}
}

// HasWebViewShowLockdownModeFirstUseMessageCompletionHandler returns true if a handler for WebViewShowLockdownModeFirstUseMessageCompletionHandler has been set.
func (d *UIDelegate) HasWebViewShowLockdownModeFirstUseMessageCompletionHandler() bool {
	return d._WebViewShowLockdownModeFirstUseMessageCompletionHandler != nil
}

// WebViewWillDismissEditMenuWithAnimator implements the PUIDelegate interface.
func (d *UIDelegate) WebViewWillDismissEditMenuWithAnimator(webView IWKWebView, animator unsafe.Pointer) {
	if d._WebViewWillDismissEditMenuWithAnimator != nil {
		d._WebViewWillDismissEditMenuWithAnimator(webView, animator)
	}
}

// HasWebViewWillDismissEditMenuWithAnimator returns true if a handler for WebViewWillDismissEditMenuWithAnimator has been set.
func (d *UIDelegate) HasWebViewWillDismissEditMenuWithAnimator() bool {
	return d._WebViewWillDismissEditMenuWithAnimator != nil
}

// WebViewWillPresentEditMenuWithAnimator implements the PUIDelegate interface.
func (d *UIDelegate) WebViewWillPresentEditMenuWithAnimator(webView IWKWebView, animator unsafe.Pointer) {
	if d._WebViewWillPresentEditMenuWithAnimator != nil {
		d._WebViewWillPresentEditMenuWithAnimator(webView, animator)
	}
}

// HasWebViewWillPresentEditMenuWithAnimator returns true if a handler for WebViewWillPresentEditMenuWithAnimator has been set.
func (d *UIDelegate) HasWebViewWillPresentEditMenuWithAnimator() bool {
	return d._WebViewWillPresentEditMenuWithAnimator != nil
}

// WebViewDidClose implements the PUIDelegate interface.
func (d *UIDelegate) WebViewDidClose(webView IWKWebView) {
	if d._WebViewDidClose != nil {
		d._WebViewDidClose(webView)
	}
}

// HasWebViewDidClose returns true if a handler for WebViewDidClose has been set.
func (d *UIDelegate) HasWebViewDidClose() bool {
	return d._WebViewDidClose != nil
}
