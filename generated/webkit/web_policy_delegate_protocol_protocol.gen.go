// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// PWebPolicyDelegate is the WebPolicyDelegate protocol interface.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebPolicyDelegate
type PWebPolicyDelegate interface {
	// Optional methods
	WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener(webView IWebView, type_ objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)
	HasWebViewDecidePolicyForMIMETypeRequestFrameDecisionListener() bool
	WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)
	HasWebViewDecidePolicyForNavigationActionRequestFrameDecisionListener() bool
	WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frameName objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer)
	HasWebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener() bool
	WebViewUnableToImplementPolicyWithErrorFrame(webView IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
	HasWebViewUnableToImplementPolicyWithErrorFrame() bool
}

// WebPolicyDelegate is a delegate implementation builder for the PWebPolicyDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebPolicyDelegate struct {
	_WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener               func(webView IWebView, type_ objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)
	_WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener       func(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)
	_WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener func(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frameName objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer)
	_WebViewUnableToImplementPolicyWithErrorFrame                             func(webView IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)
}

// SetWebViewDecidePolicyForMIMETypeRequestFrameDecisionListener sets the handler for the WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener delegate method.
//
// Decides whether to display content with a given MIME type.
func (d *WebPolicyDelegate) SetWebViewDecidePolicyForMIMETypeRequestFrameDecisionListener(f func(webView IWebView, type_ objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)) {
	d._WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener = f
}

// SetWebViewDecidePolicyForNavigationActionRequestFrameDecisionListener sets the handler for the WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener delegate method.
//
// Routes a navigation action internally or to an external viewer.
func (d *WebPolicyDelegate) SetWebViewDecidePolicyForNavigationActionRequestFrameDecisionListener(f func(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer)) {
	d._WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener = f
}

// SetWebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener sets the handler for the WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener delegate method.
//
// Decides whether to allow a targeted navigation event, such as opening a link in a new window.
func (d *WebPolicyDelegate) SetWebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener(f func(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frameName objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer)) {
	d._WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener = f
}

// SetWebViewUnableToImplementPolicyWithErrorFrame sets the handler for the WebViewUnableToImplementPolicyWithErrorFrame delegate method.
//
// Handles or drops events that were rejected by a policy maker.
func (d *WebPolicyDelegate) SetWebViewUnableToImplementPolicyWithErrorFrame(f func(webView IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame)) {
	d._WebViewUnableToImplementPolicyWithErrorFrame = f
}

// WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener implements the PWebPolicyDelegate interface.
func (d *WebPolicyDelegate) WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener(webView IWebView, type_ objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer) {
	if d._WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener != nil {
		d._WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener(webView, type_, request, frame, listener)
	}
}

// HasWebViewDecidePolicyForMIMETypeRequestFrameDecisionListener returns true if a handler for WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener has been set.
func (d *WebPolicyDelegate) HasWebViewDecidePolicyForMIMETypeRequestFrameDecisionListener() bool {
	return d._WebViewDecidePolicyForMIMETypeRequestFrameDecisionListener != nil
}

// WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener implements the PWebPolicyDelegate interface.
func (d *WebPolicyDelegate) WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frame IWebFrame, listener unsafe.Pointer) {
	if d._WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener != nil {
		d._WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener(webView, actionInformation, request, frame, listener)
	}
}

// HasWebViewDecidePolicyForNavigationActionRequestFrameDecisionListener returns true if a handler for WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener has been set.
func (d *WebPolicyDelegate) HasWebViewDecidePolicyForNavigationActionRequestFrameDecisionListener() bool {
	return d._WebViewDecidePolicyForNavigationActionRequestFrameDecisionListener != nil
}

// WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener implements the PWebPolicyDelegate interface.
func (d *WebPolicyDelegate) WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener(webView IWebView, actionInformation objc.IObject /* cross-framework: NSDictionary */, request foundation.URLRequest, frameName objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer) {
	if d._WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener != nil {
		d._WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener(webView, actionInformation, request, frameName, listener)
	}
}

// HasWebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener returns true if a handler for WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener has been set.
func (d *WebPolicyDelegate) HasWebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener() bool {
	return d._WebViewDecidePolicyForNewWindowActionRequestNewFrameNameDecisionListener != nil
}

// WebViewUnableToImplementPolicyWithErrorFrame implements the PWebPolicyDelegate interface.
func (d *WebPolicyDelegate) WebViewUnableToImplementPolicyWithErrorFrame(webView IWebView, error_ objc.IObject /* cross-framework: Error */, frame IWebFrame) {
	if d._WebViewUnableToImplementPolicyWithErrorFrame != nil {
		d._WebViewUnableToImplementPolicyWithErrorFrame(webView, error_, frame)
	}
}

// HasWebViewUnableToImplementPolicyWithErrorFrame returns true if a handler for WebViewUnableToImplementPolicyWithErrorFrame has been set.
func (d *WebPolicyDelegate) HasWebViewUnableToImplementPolicyWithErrorFrame() bool {
	return d._WebViewUnableToImplementPolicyWithErrorFrame != nil
}
