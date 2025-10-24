// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/javascriptcore"
	"github.com/tmc/appledocs/generated/objc"
)

// PWebFrameLoadDelegate is the WebFrameLoadDelegate protocol interface.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebFrameLoadDelegate
type PWebFrameLoadDelegate interface {
	// Optional methods
	WebViewDidCancelClientRedirectForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidCancelClientRedirectForFrame() bool
	WebViewDidChangeLocationWithinPageForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidChangeLocationWithinPageForFrame() bool
	WebViewDidClearWindowObjectForFrame(webView IWebView, windowObject IWebScriptObject, frame IWebFrame)
	HasWebViewDidClearWindowObjectForFrame() bool
	WebViewDidCommitLoadForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidCommitLoadForFrame() bool
	WebViewDidCreateJavaScriptContextForFrame(webView IWebView, context javascriptcore.JSContext, frame IWebFrame)
	HasWebViewDidCreateJavaScriptContextForFrame() bool
	WebViewDidFailLoadWithErrorForFrame(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
	HasWebViewDidFailLoadWithErrorForFrame() bool
	WebViewDidFailProvisionalLoadWithErrorForFrame(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
	HasWebViewDidFailProvisionalLoadWithErrorForFrame() bool
	WebViewDidFinishLoadForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidFinishLoadForFrame() bool
	WebViewDidReceiveIconForFrame(sender IWebView, image appkit.Image, frame IWebFrame)
	HasWebViewDidReceiveIconForFrame() bool
	WebViewDidReceiveServerRedirectForProvisionalLoadForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidReceiveServerRedirectForProvisionalLoadForFrame() bool
	WebViewDidReceiveTitleForFrame(sender IWebView, title objc.IObject /* cross-framework: NSString */, frame IWebFrame)
	HasWebViewDidReceiveTitleForFrame() bool
	WebViewDidStartProvisionalLoadForFrame(sender IWebView, frame IWebFrame)
	HasWebViewDidStartProvisionalLoadForFrame() bool
	WebViewWillCloseFrame(sender IWebView, frame IWebFrame)
	HasWebViewWillCloseFrame() bool
	WebViewWillPerformClientRedirectToURLDelayFireDateForFrame(sender IWebView, URL objc.IObject /* cross-framework: NSURL */, seconds float64, date objc.IObject /* cross-framework: NSDate */, frame IWebFrame)
	HasWebViewWillPerformClientRedirectToURLDelayFireDateForFrame() bool
}

// WebFrameLoadDelegate is a delegate implementation builder for the PWebFrameLoadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebFrameLoadDelegate struct {
	_WebViewDidCancelClientRedirectForFrame                     func(sender IWebView, frame IWebFrame)
	_WebViewDidChangeLocationWithinPageForFrame                 func(sender IWebView, frame IWebFrame)
	_WebViewDidClearWindowObjectForFrame                        func(webView IWebView, windowObject IWebScriptObject, frame IWebFrame)
	_WebViewDidCommitLoadForFrame                               func(sender IWebView, frame IWebFrame)
	_WebViewDidCreateJavaScriptContextForFrame                  func(webView IWebView, context javascriptcore.JSContext, frame IWebFrame)
	_WebViewDidFailLoadWithErrorForFrame                        func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
	_WebViewDidFailProvisionalLoadWithErrorForFrame             func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
	_WebViewDidFinishLoadForFrame                               func(sender IWebView, frame IWebFrame)
	_WebViewDidReceiveIconForFrame                              func(sender IWebView, image appkit.Image, frame IWebFrame)
	_WebViewDidReceiveServerRedirectForProvisionalLoadForFrame  func(sender IWebView, frame IWebFrame)
	_WebViewDidReceiveTitleForFrame                             func(sender IWebView, title objc.IObject /* cross-framework: NSString */, frame IWebFrame)
	_WebViewDidStartProvisionalLoadForFrame                     func(sender IWebView, frame IWebFrame)
	_WebViewWillCloseFrame                                      func(sender IWebView, frame IWebFrame)
	_WebViewWillPerformClientRedirectToURLDelayFireDateForFrame func(sender IWebView, URL objc.IObject /* cross-framework: NSURL */, seconds float64, date objc.IObject /* cross-framework: NSDate */, frame IWebFrame)
}

// SetWebViewDidCancelClientRedirectForFrame sets the handler for the WebViewDidCancelClientRedirectForFrame delegate method.
//
// Called when a client redirect is cancelled.
func (d *WebFrameLoadDelegate) SetWebViewDidCancelClientRedirectForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidCancelClientRedirectForFrame = f
}

// SetWebViewDidChangeLocationWithinPageForFrame sets the handler for the WebViewDidChangeLocationWithinPageForFrame delegate method.
//
// Called when the scroll position within a frame changes.
func (d *WebFrameLoadDelegate) SetWebViewDidChangeLocationWithinPageForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidChangeLocationWithinPageForFrame = f
}

// SetWebViewDidClearWindowObjectForFrame sets the handler for the WebViewDidClearWindowObjectForFrame delegate method.
//
// Called when the JavaScript window object in a frame is ready for loading.
func (d *WebFrameLoadDelegate) SetWebViewDidClearWindowObjectForFrame(f func(webView IWebView, windowObject IWebScriptObject, frame IWebFrame)) {
	d._WebViewDidClearWindowObjectForFrame = f
}

// SetWebViewDidCommitLoadForFrame sets the handler for the WebViewDidCommitLoadForFrame delegate method.
//
// Called when content starts arriving for a page load.
func (d *WebFrameLoadDelegate) SetWebViewDidCommitLoadForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidCommitLoadForFrame = f
}

// SetWebViewDidCreateJavaScriptContextForFrame sets the handler for the WebViewDidCreateJavaScriptContextForFrame delegate method.
//
// Notifies the delegate that a new JavaScript context has been created.
func (d *WebFrameLoadDelegate) SetWebViewDidCreateJavaScriptContextForFrame(f func(webView IWebView, context javascriptcore.JSContext, frame IWebFrame)) {
	d._WebViewDidCreateJavaScriptContextForFrame = f
}

// SetWebViewDidFailLoadWithErrorForFrame sets the handler for the WebViewDidFailLoadWithErrorForFrame delegate method.
//
// Called when an error occurs loading a committed data source.
func (d *WebFrameLoadDelegate) SetWebViewDidFailLoadWithErrorForFrame(f func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)) {
	d._WebViewDidFailLoadWithErrorForFrame = f
}

// SetWebViewDidFailProvisionalLoadWithErrorForFrame sets the handler for the WebViewDidFailProvisionalLoadWithErrorForFrame delegate method.
//
// Called if an error occurs when starting to load data for a page.
func (d *WebFrameLoadDelegate) SetWebViewDidFailProvisionalLoadWithErrorForFrame(f func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)) {
	d._WebViewDidFailProvisionalLoadWithErrorForFrame = f
}

// SetWebViewDidFinishLoadForFrame sets the handler for the WebViewDidFinishLoadForFrame delegate method.
//
// Called when a page load completes.
func (d *WebFrameLoadDelegate) SetWebViewDidFinishLoadForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidFinishLoadForFrame = f
}

// SetWebViewDidReceiveIconForFrame sets the handler for the WebViewDidReceiveIconForFrame delegate method.
//
// Called when a page icon changes.
func (d *WebFrameLoadDelegate) SetWebViewDidReceiveIconForFrame(f func(sender IWebView, image appkit.Image, frame IWebFrame)) {
	d._WebViewDidReceiveIconForFrame = f
}

// SetWebViewDidReceiveServerRedirectForProvisionalLoadForFrame sets the handler for the WebViewDidReceiveServerRedirectForProvisionalLoadForFrame delegate method.
//
// Called when a provisional data source for a frame receives a server redirect.
func (d *WebFrameLoadDelegate) SetWebViewDidReceiveServerRedirectForProvisionalLoadForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidReceiveServerRedirectForProvisionalLoadForFrame = f
}

// SetWebViewDidReceiveTitleForFrame sets the handler for the WebViewDidReceiveTitleForFrame delegate method.
//
// Called when the page title of a frame loads or changes.
func (d *WebFrameLoadDelegate) SetWebViewDidReceiveTitleForFrame(f func(sender IWebView, title objc.IObject /* cross-framework: NSString */, frame IWebFrame)) {
	d._WebViewDidReceiveTitleForFrame = f
}

// SetWebViewDidStartProvisionalLoadForFrame sets the handler for the WebViewDidStartProvisionalLoadForFrame delegate method.
//
// Called when a page load is in progress in a given frame.
func (d *WebFrameLoadDelegate) SetWebViewDidStartProvisionalLoadForFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewDidStartProvisionalLoadForFrame = f
}

// SetWebViewWillCloseFrame sets the handler for the WebViewWillCloseFrame delegate method.
//
// Called when a frame will be closed.
func (d *WebFrameLoadDelegate) SetWebViewWillCloseFrame(f func(sender IWebView, frame IWebFrame)) {
	d._WebViewWillCloseFrame = f
}

// SetWebViewWillPerformClientRedirectToURLDelayFireDateForFrame sets the handler for the WebViewWillPerformClientRedirectToURLDelayFireDateForFrame delegate method.
//
// Called when a frame receives a client redirect and before it is fired.
func (d *WebFrameLoadDelegate) SetWebViewWillPerformClientRedirectToURLDelayFireDateForFrame(f func(sender IWebView, URL objc.IObject /* cross-framework: NSURL */, seconds float64, date objc.IObject /* cross-framework: NSDate */, frame IWebFrame)) {
	d._WebViewWillPerformClientRedirectToURLDelayFireDateForFrame = f
}

// WebViewDidCancelClientRedirectForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidCancelClientRedirectForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidCancelClientRedirectForFrame != nil {
		d._WebViewDidCancelClientRedirectForFrame(sender, frame)
	}
}

// HasWebViewDidCancelClientRedirectForFrame returns true if a handler for WebViewDidCancelClientRedirectForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidCancelClientRedirectForFrame() bool {
	return d._WebViewDidCancelClientRedirectForFrame != nil
}

// WebViewDidChangeLocationWithinPageForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidChangeLocationWithinPageForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidChangeLocationWithinPageForFrame != nil {
		d._WebViewDidChangeLocationWithinPageForFrame(sender, frame)
	}
}

// HasWebViewDidChangeLocationWithinPageForFrame returns true if a handler for WebViewDidChangeLocationWithinPageForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidChangeLocationWithinPageForFrame() bool {
	return d._WebViewDidChangeLocationWithinPageForFrame != nil
}

// WebViewDidClearWindowObjectForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidClearWindowObjectForFrame(webView IWebView, windowObject IWebScriptObject, frame IWebFrame) {
	if d._WebViewDidClearWindowObjectForFrame != nil {
		d._WebViewDidClearWindowObjectForFrame(webView, windowObject, frame)
	}
}

// HasWebViewDidClearWindowObjectForFrame returns true if a handler for WebViewDidClearWindowObjectForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidClearWindowObjectForFrame() bool {
	return d._WebViewDidClearWindowObjectForFrame != nil
}

// WebViewDidCommitLoadForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidCommitLoadForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidCommitLoadForFrame != nil {
		d._WebViewDidCommitLoadForFrame(sender, frame)
	}
}

// HasWebViewDidCommitLoadForFrame returns true if a handler for WebViewDidCommitLoadForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidCommitLoadForFrame() bool {
	return d._WebViewDidCommitLoadForFrame != nil
}

// WebViewDidCreateJavaScriptContextForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidCreateJavaScriptContextForFrame(webView IWebView, context javascriptcore.JSContext, frame IWebFrame) {
	if d._WebViewDidCreateJavaScriptContextForFrame != nil {
		d._WebViewDidCreateJavaScriptContextForFrame(webView, context, frame)
	}
}

// HasWebViewDidCreateJavaScriptContextForFrame returns true if a handler for WebViewDidCreateJavaScriptContextForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidCreateJavaScriptContextForFrame() bool {
	return d._WebViewDidCreateJavaScriptContextForFrame != nil
}

// WebViewDidFailLoadWithErrorForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidFailLoadWithErrorForFrame(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame) {
	if d._WebViewDidFailLoadWithErrorForFrame != nil {
		d._WebViewDidFailLoadWithErrorForFrame(sender, error_, frame)
	}
}

// HasWebViewDidFailLoadWithErrorForFrame returns true if a handler for WebViewDidFailLoadWithErrorForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidFailLoadWithErrorForFrame() bool {
	return d._WebViewDidFailLoadWithErrorForFrame != nil
}

// WebViewDidFailProvisionalLoadWithErrorForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidFailProvisionalLoadWithErrorForFrame(sender IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame) {
	if d._WebViewDidFailProvisionalLoadWithErrorForFrame != nil {
		d._WebViewDidFailProvisionalLoadWithErrorForFrame(sender, error_, frame)
	}
}

// HasWebViewDidFailProvisionalLoadWithErrorForFrame returns true if a handler for WebViewDidFailProvisionalLoadWithErrorForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidFailProvisionalLoadWithErrorForFrame() bool {
	return d._WebViewDidFailProvisionalLoadWithErrorForFrame != nil
}

// WebViewDidFinishLoadForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidFinishLoadForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidFinishLoadForFrame != nil {
		d._WebViewDidFinishLoadForFrame(sender, frame)
	}
}

// HasWebViewDidFinishLoadForFrame returns true if a handler for WebViewDidFinishLoadForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidFinishLoadForFrame() bool {
	return d._WebViewDidFinishLoadForFrame != nil
}

// WebViewDidReceiveIconForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidReceiveIconForFrame(sender IWebView, image appkit.Image, frame IWebFrame) {
	if d._WebViewDidReceiveIconForFrame != nil {
		d._WebViewDidReceiveIconForFrame(sender, image, frame)
	}
}

// HasWebViewDidReceiveIconForFrame returns true if a handler for WebViewDidReceiveIconForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidReceiveIconForFrame() bool {
	return d._WebViewDidReceiveIconForFrame != nil
}

// WebViewDidReceiveServerRedirectForProvisionalLoadForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidReceiveServerRedirectForProvisionalLoadForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidReceiveServerRedirectForProvisionalLoadForFrame != nil {
		d._WebViewDidReceiveServerRedirectForProvisionalLoadForFrame(sender, frame)
	}
}

// HasWebViewDidReceiveServerRedirectForProvisionalLoadForFrame returns true if a handler for WebViewDidReceiveServerRedirectForProvisionalLoadForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidReceiveServerRedirectForProvisionalLoadForFrame() bool {
	return d._WebViewDidReceiveServerRedirectForProvisionalLoadForFrame != nil
}

// WebViewDidReceiveTitleForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidReceiveTitleForFrame(sender IWebView, title objc.IObject /* cross-framework: NSString */, frame IWebFrame) {
	if d._WebViewDidReceiveTitleForFrame != nil {
		d._WebViewDidReceiveTitleForFrame(sender, title, frame)
	}
}

// HasWebViewDidReceiveTitleForFrame returns true if a handler for WebViewDidReceiveTitleForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidReceiveTitleForFrame() bool {
	return d._WebViewDidReceiveTitleForFrame != nil
}

// WebViewDidStartProvisionalLoadForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewDidStartProvisionalLoadForFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewDidStartProvisionalLoadForFrame != nil {
		d._WebViewDidStartProvisionalLoadForFrame(sender, frame)
	}
}

// HasWebViewDidStartProvisionalLoadForFrame returns true if a handler for WebViewDidStartProvisionalLoadForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewDidStartProvisionalLoadForFrame() bool {
	return d._WebViewDidStartProvisionalLoadForFrame != nil
}

// WebViewWillCloseFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewWillCloseFrame(sender IWebView, frame IWebFrame) {
	if d._WebViewWillCloseFrame != nil {
		d._WebViewWillCloseFrame(sender, frame)
	}
}

// HasWebViewWillCloseFrame returns true if a handler for WebViewWillCloseFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewWillCloseFrame() bool {
	return d._WebViewWillCloseFrame != nil
}

// WebViewWillPerformClientRedirectToURLDelayFireDateForFrame implements the PWebFrameLoadDelegate interface.
func (d *WebFrameLoadDelegate) WebViewWillPerformClientRedirectToURLDelayFireDateForFrame(sender IWebView, URL objc.IObject /* cross-framework: NSURL */, seconds float64, date objc.IObject /* cross-framework: NSDate */, frame IWebFrame) {
	if d._WebViewWillPerformClientRedirectToURLDelayFireDateForFrame != nil {
		d._WebViewWillPerformClientRedirectToURLDelayFireDateForFrame(sender, URL, seconds, date, frame)
	}
}

// HasWebViewWillPerformClientRedirectToURLDelayFireDateForFrame returns true if a handler for WebViewWillPerformClientRedirectToURLDelayFireDateForFrame has been set.
func (d *WebFrameLoadDelegate) HasWebViewWillPerformClientRedirectToURLDelayFireDateForFrame() bool {
	return d._WebViewWillPerformClientRedirectToURLDelayFireDateForFrame != nil
}
