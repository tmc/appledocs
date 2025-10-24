// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// PWebResourceLoadDelegate is the WebResourceLoadDelegate protocol interface.
//
// Web view resource load delegates implement this protocol to be notified on the progress of loading individual resources. Note that there can be hundreds of resources, such as images and other media, per page. So, if you just want to get page loading status see the WebFrameLoadDelegate protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebResourceLoadDelegate
type PWebResourceLoadDelegate interface {
	// Optional methods
	WebViewIdentifierForInitialRequestFromDataSource(sender IWebView, request foundation.URLRequest, dataSource IWebDataSource) objc.ID
	HasWebViewIdentifierForInitialRequestFromDataSource() bool
	WebViewPlugInFailedWithErrorDataSource(sender IWebView, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)
	HasWebViewPlugInFailedWithErrorDataSource() bool
	WebViewResourceDidCancelAuthenticationChallengeFromDataSource(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)
	HasWebViewResourceDidCancelAuthenticationChallengeFromDataSource() bool
	WebViewResourceDidFailLoadingWithErrorFromDataSource(sender IWebView, identifier objc.IObject, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)
	HasWebViewResourceDidFailLoadingWithErrorFromDataSource() bool
	WebViewResourceDidFinishLoadingFromDataSource(sender IWebView, identifier objc.IObject, dataSource IWebDataSource)
	HasWebViewResourceDidFinishLoadingFromDataSource() bool
	WebViewResourceDidReceiveResponseFromDataSource(sender IWebView, identifier objc.IObject, response foundation.URLResponse, dataSource IWebDataSource)
	HasWebViewResourceDidReceiveResponseFromDataSource() bool
	WebViewResourceDidReceiveAuthenticationChallengeFromDataSource(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)
	HasWebViewResourceDidReceiveAuthenticationChallengeFromDataSource() bool
	WebViewResourceDidReceiveContentLengthFromDataSource(sender IWebView, identifier objc.IObject, length int, dataSource IWebDataSource)
	HasWebViewResourceDidReceiveContentLengthFromDataSource() bool
	WebViewResourceWillSendRequestRedirectResponseFromDataSource(sender IWebView, identifier objc.IObject, request foundation.URLRequest, redirectResponse foundation.URLResponse, dataSource IWebDataSource) foundation.URLRequest
	HasWebViewResourceWillSendRequestRedirectResponseFromDataSource() bool
}

// WebResourceLoadDelegate is a delegate implementation builder for the PWebResourceLoadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebResourceLoadDelegate struct {
	_WebViewIdentifierForInitialRequestFromDataSource               func(sender IWebView, request foundation.URLRequest, dataSource IWebDataSource) objc.ID
	_WebViewPlugInFailedWithErrorDataSource                         func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)
	_WebViewResourceDidCancelAuthenticationChallengeFromDataSource  func(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)
	_WebViewResourceDidFailLoadingWithErrorFromDataSource           func(sender IWebView, identifier objc.IObject, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)
	_WebViewResourceDidFinishLoadingFromDataSource                  func(sender IWebView, identifier objc.IObject, dataSource IWebDataSource)
	_WebViewResourceDidReceiveResponseFromDataSource                func(sender IWebView, identifier objc.IObject, response foundation.URLResponse, dataSource IWebDataSource)
	_WebViewResourceDidReceiveAuthenticationChallengeFromDataSource func(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)
	_WebViewResourceDidReceiveContentLengthFromDataSource           func(sender IWebView, identifier objc.IObject, length int, dataSource IWebDataSource)
	_WebViewResourceWillSendRequestRedirectResponseFromDataSource   func(sender IWebView, identifier objc.IObject, request foundation.URLRequest, redirectResponse foundation.URLResponse, dataSource IWebDataSource) foundation.URLRequest
}

// SetWebViewIdentifierForInitialRequestFromDataSource sets the handler for the WebViewIdentifierForInitialRequestFromDataSource delegate method.
//
// Returns an identifier object used to track the progress of loading a single resource.
func (d *WebResourceLoadDelegate) SetWebViewIdentifierForInitialRequestFromDataSource(f func(sender IWebView, request foundation.URLRequest, dataSource IWebDataSource) objc.ID) {
	d._WebViewIdentifierForInitialRequestFromDataSource = f
}

// SetWebViewPlugInFailedWithErrorDataSource sets the handler for the WebViewPlugInFailedWithErrorDataSource delegate method.
//
// Invoked when a plug-in fails to load.
func (d *WebResourceLoadDelegate) SetWebViewPlugInFailedWithErrorDataSource(f func(sender IWebView, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)) {
	d._WebViewPlugInFailedWithErrorDataSource = f
}

// SetWebViewResourceDidCancelAuthenticationChallengeFromDataSource sets the handler for the WebViewResourceDidCancelAuthenticationChallengeFromDataSource delegate method.
//
// Invoked when an authentication challenge for a resource was canceled.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidCancelAuthenticationChallengeFromDataSource(f func(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)) {
	d._WebViewResourceDidCancelAuthenticationChallengeFromDataSource = f
}

// SetWebViewResourceDidFailLoadingWithErrorFromDataSource sets the handler for the WebViewResourceDidFailLoadingWithErrorFromDataSource delegate method.
//
// Invoked when a resource failed to load.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidFailLoadingWithErrorFromDataSource(f func(sender IWebView, identifier objc.IObject, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)) {
	d._WebViewResourceDidFailLoadingWithErrorFromDataSource = f
}

// SetWebViewResourceDidFinishLoadingFromDataSource sets the handler for the WebViewResourceDidFinishLoadingFromDataSource delegate method.
//
// Invoked when all of the data for a given resource is loaded.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidFinishLoadingFromDataSource(f func(sender IWebView, identifier objc.IObject, dataSource IWebDataSource)) {
	d._WebViewResourceDidFinishLoadingFromDataSource = f
}

// SetWebViewResourceDidReceiveResponseFromDataSource sets the handler for the WebViewResourceDidReceiveResponseFromDataSource delegate method.
//
// Invoked after a resource has been loaded.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidReceiveResponseFromDataSource(f func(sender IWebView, identifier objc.IObject, response foundation.URLResponse, dataSource IWebDataSource)) {
	d._WebViewResourceDidReceiveResponseFromDataSource = f
}

// SetWebViewResourceDidReceiveAuthenticationChallengeFromDataSource sets the handler for the WebViewResourceDidReceiveAuthenticationChallengeFromDataSource delegate method.
//
// Invoked when an authentication challenge has been received for a resource.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidReceiveAuthenticationChallengeFromDataSource(f func(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource)) {
	d._WebViewResourceDidReceiveAuthenticationChallengeFromDataSource = f
}

// SetWebViewResourceDidReceiveContentLengthFromDataSource sets the handler for the WebViewResourceDidReceiveContentLengthFromDataSource delegate method.
//
// Invoked when some of the data for a given resource has arrived.
func (d *WebResourceLoadDelegate) SetWebViewResourceDidReceiveContentLengthFromDataSource(f func(sender IWebView, identifier objc.IObject, length int, dataSource IWebDataSource)) {
	d._WebViewResourceDidReceiveContentLengthFromDataSource = f
}

// SetWebViewResourceWillSendRequestRedirectResponseFromDataSource sets the handler for the WebViewResourceWillSendRequestRedirectResponseFromDataSource delegate method.
//
// Invoked before a request is initiated for a resource and returns a possibly modified request.
func (d *WebResourceLoadDelegate) SetWebViewResourceWillSendRequestRedirectResponseFromDataSource(f func(sender IWebView, identifier objc.IObject, request foundation.URLRequest, redirectResponse foundation.URLResponse, dataSource IWebDataSource) foundation.URLRequest) {
	d._WebViewResourceWillSendRequestRedirectResponseFromDataSource = f
}

// WebViewIdentifierForInitialRequestFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewIdentifierForInitialRequestFromDataSource(sender IWebView, request foundation.URLRequest, dataSource IWebDataSource) objc.ID {
	if d._WebViewIdentifierForInitialRequestFromDataSource != nil {
		return d._WebViewIdentifierForInitialRequestFromDataSource(sender, request, dataSource)
	}
	var zero objc.ID
	return zero
}

// HasWebViewIdentifierForInitialRequestFromDataSource returns true if a handler for WebViewIdentifierForInitialRequestFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewIdentifierForInitialRequestFromDataSource() bool {
	return d._WebViewIdentifierForInitialRequestFromDataSource != nil
}

// WebViewPlugInFailedWithErrorDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewPlugInFailedWithErrorDataSource(sender IWebView, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource) {
	if d._WebViewPlugInFailedWithErrorDataSource != nil {
		d._WebViewPlugInFailedWithErrorDataSource(sender, error_, dataSource)
	}
}

// HasWebViewPlugInFailedWithErrorDataSource returns true if a handler for WebViewPlugInFailedWithErrorDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewPlugInFailedWithErrorDataSource() bool {
	return d._WebViewPlugInFailedWithErrorDataSource != nil
}

// WebViewResourceDidCancelAuthenticationChallengeFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidCancelAuthenticationChallengeFromDataSource(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource) {
	if d._WebViewResourceDidCancelAuthenticationChallengeFromDataSource != nil {
		d._WebViewResourceDidCancelAuthenticationChallengeFromDataSource(sender, identifier, challenge, dataSource)
	}
}

// HasWebViewResourceDidCancelAuthenticationChallengeFromDataSource returns true if a handler for WebViewResourceDidCancelAuthenticationChallengeFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidCancelAuthenticationChallengeFromDataSource() bool {
	return d._WebViewResourceDidCancelAuthenticationChallengeFromDataSource != nil
}

// WebViewResourceDidFailLoadingWithErrorFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidFailLoadingWithErrorFromDataSource(sender IWebView, identifier objc.IObject, error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource) {
	if d._WebViewResourceDidFailLoadingWithErrorFromDataSource != nil {
		d._WebViewResourceDidFailLoadingWithErrorFromDataSource(sender, identifier, error_, dataSource)
	}
}

// HasWebViewResourceDidFailLoadingWithErrorFromDataSource returns true if a handler for WebViewResourceDidFailLoadingWithErrorFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidFailLoadingWithErrorFromDataSource() bool {
	return d._WebViewResourceDidFailLoadingWithErrorFromDataSource != nil
}

// WebViewResourceDidFinishLoadingFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidFinishLoadingFromDataSource(sender IWebView, identifier objc.IObject, dataSource IWebDataSource) {
	if d._WebViewResourceDidFinishLoadingFromDataSource != nil {
		d._WebViewResourceDidFinishLoadingFromDataSource(sender, identifier, dataSource)
	}
}

// HasWebViewResourceDidFinishLoadingFromDataSource returns true if a handler for WebViewResourceDidFinishLoadingFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidFinishLoadingFromDataSource() bool {
	return d._WebViewResourceDidFinishLoadingFromDataSource != nil
}

// WebViewResourceDidReceiveResponseFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidReceiveResponseFromDataSource(sender IWebView, identifier objc.IObject, response foundation.URLResponse, dataSource IWebDataSource) {
	if d._WebViewResourceDidReceiveResponseFromDataSource != nil {
		d._WebViewResourceDidReceiveResponseFromDataSource(sender, identifier, response, dataSource)
	}
}

// HasWebViewResourceDidReceiveResponseFromDataSource returns true if a handler for WebViewResourceDidReceiveResponseFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidReceiveResponseFromDataSource() bool {
	return d._WebViewResourceDidReceiveResponseFromDataSource != nil
}

// WebViewResourceDidReceiveAuthenticationChallengeFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidReceiveAuthenticationChallengeFromDataSource(sender IWebView, identifier objc.IObject, challenge foundation.URLAuthenticationChallenge, dataSource IWebDataSource) {
	if d._WebViewResourceDidReceiveAuthenticationChallengeFromDataSource != nil {
		d._WebViewResourceDidReceiveAuthenticationChallengeFromDataSource(sender, identifier, challenge, dataSource)
	}
}

// HasWebViewResourceDidReceiveAuthenticationChallengeFromDataSource returns true if a handler for WebViewResourceDidReceiveAuthenticationChallengeFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidReceiveAuthenticationChallengeFromDataSource() bool {
	return d._WebViewResourceDidReceiveAuthenticationChallengeFromDataSource != nil
}

// WebViewResourceDidReceiveContentLengthFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceDidReceiveContentLengthFromDataSource(sender IWebView, identifier objc.IObject, length int, dataSource IWebDataSource) {
	if d._WebViewResourceDidReceiveContentLengthFromDataSource != nil {
		d._WebViewResourceDidReceiveContentLengthFromDataSource(sender, identifier, length, dataSource)
	}
}

// HasWebViewResourceDidReceiveContentLengthFromDataSource returns true if a handler for WebViewResourceDidReceiveContentLengthFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceDidReceiveContentLengthFromDataSource() bool {
	return d._WebViewResourceDidReceiveContentLengthFromDataSource != nil
}

// WebViewResourceWillSendRequestRedirectResponseFromDataSource implements the PWebResourceLoadDelegate interface.
func (d *WebResourceLoadDelegate) WebViewResourceWillSendRequestRedirectResponseFromDataSource(sender IWebView, identifier objc.IObject, request foundation.URLRequest, redirectResponse foundation.URLResponse, dataSource IWebDataSource) foundation.URLRequest {
	if d._WebViewResourceWillSendRequestRedirectResponseFromDataSource != nil {
		return d._WebViewResourceWillSendRequestRedirectResponseFromDataSource(sender, identifier, request, redirectResponse, dataSource)
	}
	var zero foundation.URLRequest
	return zero
}

// HasWebViewResourceWillSendRequestRedirectResponseFromDataSource returns true if a handler for WebViewResourceWillSendRequestRedirectResponseFromDataSource has been set.
func (d *WebResourceLoadDelegate) HasWebViewResourceWillSendRequestRedirectResponseFromDataSource() bool {
	return d._WebViewResourceWillSendRequestRedirectResponseFromDataSource != nil
}
