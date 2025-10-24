// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/func(unsafe"
)

// PNavigationDelegate is the WKNavigationDelegate protocol interface.
//
// Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.webkit/documentation/WebKit/WKNavigationDelegate
type PNavigationDelegate interface {
	// Optional methods
	WebViewAuthenticationChallengeShouldAllowDeprecatedTLS(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, decisionHandler func(unsafe.Pointer))
	HasWebViewAuthenticationChallengeShouldAllowDeprecatedTLS() bool
	WebViewDecidePolicyForNavigationResponseDecisionHandler(webView IWKWebView, navigationResponse IWKNavigationResponse, decisionHandler func(unsafe.Pointer))
	HasWebViewDecidePolicyForNavigationResponseDecisionHandler() bool
	WebViewDecidePolicyForNavigationActionDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, decisionHandler func(unsafe.Pointer))
	HasWebViewDecidePolicyForNavigationActionDecisionHandler() bool
	WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, preferences IWKWebpagePreferences, decisionHandler func(unsafe.Pointer, unsafe.Pointer))
	HasWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler() bool
	WebViewDidCommitNavigation(webView IWKWebView, navigation IWKNavigation)
	HasWebViewDidCommitNavigation() bool
	WebViewDidFailNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)
	HasWebViewDidFailNavigationWithError() bool
	WebViewDidFailProvisionalNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)
	HasWebViewDidFailProvisionalNavigationWithError() bool
	WebViewDidFinishNavigation(webView IWKWebView, navigation IWKNavigation)
	HasWebViewDidFinishNavigation() bool
	WebViewDidReceiveAuthenticationChallengeCompletionHandler(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	HasWebViewDidReceiveAuthenticationChallengeCompletionHandler() bool
	WebViewDidReceiveServerRedirectForProvisionalNavigation(webView IWKWebView, navigation IWKNavigation)
	HasWebViewDidReceiveServerRedirectForProvisionalNavigation() bool
	WebViewDidStartProvisionalNavigation(webView IWKWebView, navigation IWKNavigation)
	HasWebViewDidStartProvisionalNavigation() bool
	WebViewNavigationActionDidBecomeDownload(webView IWKWebView, navigationAction IWKNavigationAction, download IWKDownload)
	HasWebViewNavigationActionDidBecomeDownload() bool
	WebViewNavigationResponseDidBecomeDownload(webView IWKWebView, navigationResponse IWKNavigationResponse, download IWKDownload)
	HasWebViewNavigationResponseDidBecomeDownload() bool
	WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler(webView IWKWebView, backForwardListItem IWKBackForwardListItem, willUseInstantBack bool, completionHandler unsafe.Pointer)
	HasWebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler() bool
	WebViewWebContentProcessDidTerminate(webView IWKWebView)
	HasWebViewWebContentProcessDidTerminate() bool
}

// NavigationDelegate is a delegate implementation builder for the PNavigationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NavigationDelegate struct {
	_WebViewAuthenticationChallengeShouldAllowDeprecatedTLS func(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, decisionHandler func(unsafe.Pointer))
	_WebViewDecidePolicyForNavigationResponseDecisionHandler func(webView IWKWebView, navigationResponse IWKNavigationResponse, decisionHandler func(unsafe.Pointer))
	_WebViewDecidePolicyForNavigationActionDecisionHandler func(webView IWKWebView, navigationAction IWKNavigationAction, decisionHandler func(unsafe.Pointer))
	_WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler func(webView IWKWebView, navigationAction IWKNavigationAction, preferences IWKWebpagePreferences, decisionHandler func(unsafe.Pointer, unsafe.Pointer))
	_WebViewDidCommitNavigation func(webView IWKWebView, navigation IWKNavigation)
	_WebViewDidFailNavigationWithError func(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)
	_WebViewDidFailProvisionalNavigationWithError func(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)
	_WebViewDidFinishNavigation func(webView IWKWebView, navigation IWKNavigation)
	_WebViewDidReceiveAuthenticationChallengeCompletionHandler func(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	_WebViewDidReceiveServerRedirectForProvisionalNavigation func(webView IWKWebView, navigation IWKNavigation)
	_WebViewDidStartProvisionalNavigation func(webView IWKWebView, navigation IWKNavigation)
	_WebViewNavigationActionDidBecomeDownload func(webView IWKWebView, navigationAction IWKNavigationAction, download IWKDownload)
	_WebViewNavigationResponseDidBecomeDownload func(webView IWKWebView, navigationResponse IWKNavigationResponse, download IWKDownload)
	_WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler func(webView IWKWebView, backForwardListItem IWKBackForwardListItem, willUseInstantBack bool, completionHandler unsafe.Pointer)
	_WebViewWebContentProcessDidTerminate func(webView IWKWebView)
}

// SetWebViewAuthenticationChallengeShouldAllowDeprecatedTLS sets the handler for the WebViewAuthenticationChallengeShouldAllowDeprecatedTLS delegate method.
//
// Asks the delegate whether to continue with a connection that uses a deprecated version of TLS.
func (d *NavigationDelegate) SetWebViewAuthenticationChallengeShouldAllowDeprecatedTLS(f func(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, decisionHandler func(unsafe.Pointer))) {
	d._WebViewAuthenticationChallengeShouldAllowDeprecatedTLS = f
}

// SetWebViewDecidePolicyForNavigationResponseDecisionHandler sets the handler for the WebViewDecidePolicyForNavigationResponseDecisionHandler delegate method.
//
// Asks the delegate for permission to navigate to new content after the response to the navigation request is known.
func (d *NavigationDelegate) SetWebViewDecidePolicyForNavigationResponseDecisionHandler(f func(webView IWKWebView, navigationResponse IWKNavigationResponse, decisionHandler func(unsafe.Pointer))) {
	d._WebViewDecidePolicyForNavigationResponseDecisionHandler = f
}

// SetWebViewDecidePolicyForNavigationActionDecisionHandler sets the handler for the WebViewDecidePolicyForNavigationActionDecisionHandler delegate method.
//
// Asks the delegate for permission to navigate to new content based on the specified action information.
func (d *NavigationDelegate) SetWebViewDecidePolicyForNavigationActionDecisionHandler(f func(webView IWKWebView, navigationAction IWKNavigationAction, decisionHandler func(unsafe.Pointer))) {
	d._WebViewDecidePolicyForNavigationActionDecisionHandler = f
}

// SetWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler sets the handler for the WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler delegate method.
//
// Asks the delegate for permission to navigate to new content based on the specified preferences and action information.
func (d *NavigationDelegate) SetWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(f func(webView IWKWebView, navigationAction IWKNavigationAction, preferences IWKWebpagePreferences, decisionHandler func(unsafe.Pointer, unsafe.Pointer))) {
	d._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler = f
}

// SetWebViewDidCommitNavigation sets the handler for the WebViewDidCommitNavigation delegate method.
//
// Tells the delegate that the web view has started to receive content for the main frame.
func (d *NavigationDelegate) SetWebViewDidCommitNavigation(f func(webView IWKWebView, navigation IWKNavigation)) {
	d._WebViewDidCommitNavigation = f
}

// SetWebViewDidFailNavigationWithError sets the handler for the WebViewDidFailNavigationWithError delegate method.
//
// Tells the delegate that an error occurred during navigation.
func (d *NavigationDelegate) SetWebViewDidFailNavigationWithError(f func(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)) {
	d._WebViewDidFailNavigationWithError = f
}

// SetWebViewDidFailProvisionalNavigationWithError sets the handler for the WebViewDidFailProvisionalNavigationWithError delegate method.
//
// Tells the delegate that an error occurred during the early navigation process.
func (d *NavigationDelegate) SetWebViewDidFailProvisionalNavigationWithError(f func(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */)) {
	d._WebViewDidFailProvisionalNavigationWithError = f
}

// SetWebViewDidFinishNavigation sets the handler for the WebViewDidFinishNavigation delegate method.
//
// Tells the delegate that navigation is complete.
func (d *NavigationDelegate) SetWebViewDidFinishNavigation(f func(webView IWKWebView, navigation IWKNavigation)) {
	d._WebViewDidFinishNavigation = f
}

// SetWebViewDidReceiveAuthenticationChallengeCompletionHandler sets the handler for the WebViewDidReceiveAuthenticationChallengeCompletionHandler delegate method.
//
// Asks the delegate to respond to an authentication challenge.
func (d *NavigationDelegate) SetWebViewDidReceiveAuthenticationChallengeCompletionHandler(f func(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))) {
	d._WebViewDidReceiveAuthenticationChallengeCompletionHandler = f
}

// SetWebViewDidReceiveServerRedirectForProvisionalNavigation sets the handler for the WebViewDidReceiveServerRedirectForProvisionalNavigation delegate method.
//
// Tells the delegate that the web view received a server redirect for a request.
func (d *NavigationDelegate) SetWebViewDidReceiveServerRedirectForProvisionalNavigation(f func(webView IWKWebView, navigation IWKNavigation)) {
	d._WebViewDidReceiveServerRedirectForProvisionalNavigation = f
}

// SetWebViewDidStartProvisionalNavigation sets the handler for the WebViewDidStartProvisionalNavigation delegate method.
//
// Tells the delegate that navigation from the main frame has started.
func (d *NavigationDelegate) SetWebViewDidStartProvisionalNavigation(f func(webView IWKWebView, navigation IWKNavigation)) {
	d._WebViewDidStartProvisionalNavigation = f
}

// SetWebViewNavigationActionDidBecomeDownload sets the handler for the WebViewNavigationActionDidBecomeDownload delegate method.
//
// Tells the delegate that a navigation action became a download.
func (d *NavigationDelegate) SetWebViewNavigationActionDidBecomeDownload(f func(webView IWKWebView, navigationAction IWKNavigationAction, download IWKDownload)) {
	d._WebViewNavigationActionDidBecomeDownload = f
}

// SetWebViewNavigationResponseDidBecomeDownload sets the handler for the WebViewNavigationResponseDidBecomeDownload delegate method.
//
// Tells the delegate that a navigation response became a download.
func (d *NavigationDelegate) SetWebViewNavigationResponseDidBecomeDownload(f func(webView IWKWebView, navigationResponse IWKNavigationResponse, download IWKDownload)) {
	d._WebViewNavigationResponseDidBecomeDownload = f
}

// SetWebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler sets the handler for the WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler delegate method.
func (d *NavigationDelegate) SetWebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler(f func(webView IWKWebView, backForwardListItem IWKBackForwardListItem, willUseInstantBack bool, completionHandler unsafe.Pointer)) {
	d._WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler = f
}

// SetWebViewWebContentProcessDidTerminate sets the handler for the WebViewWebContentProcessDidTerminate delegate method.
//
// Tells the delegate that the web view’s content process was terminated.
func (d *NavigationDelegate) SetWebViewWebContentProcessDidTerminate(f func(webView IWKWebView)) {
	d._WebViewWebContentProcessDidTerminate = f
}

// WebViewAuthenticationChallengeShouldAllowDeprecatedTLS implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewAuthenticationChallengeShouldAllowDeprecatedTLS(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, decisionHandler func(unsafe.Pointer)) {
	if d._WebViewAuthenticationChallengeShouldAllowDeprecatedTLS != nil {
		d._WebViewAuthenticationChallengeShouldAllowDeprecatedTLS(webView, challenge, decisionHandler)
	}
}

// HasWebViewAuthenticationChallengeShouldAllowDeprecatedTLS returns true if a handler for WebViewAuthenticationChallengeShouldAllowDeprecatedTLS has been set.
func (d *NavigationDelegate) HasWebViewAuthenticationChallengeShouldAllowDeprecatedTLS() bool {
	return d._WebViewAuthenticationChallengeShouldAllowDeprecatedTLS != nil
}

// WebViewDecidePolicyForNavigationResponseDecisionHandler implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDecidePolicyForNavigationResponseDecisionHandler(webView IWKWebView, navigationResponse IWKNavigationResponse, decisionHandler func(unsafe.Pointer)) {
	if d._WebViewDecidePolicyForNavigationResponseDecisionHandler != nil {
		d._WebViewDecidePolicyForNavigationResponseDecisionHandler(webView, navigationResponse, decisionHandler)
	}
}

// HasWebViewDecidePolicyForNavigationResponseDecisionHandler returns true if a handler for WebViewDecidePolicyForNavigationResponseDecisionHandler has been set.
func (d *NavigationDelegate) HasWebViewDecidePolicyForNavigationResponseDecisionHandler() bool {
	return d._WebViewDecidePolicyForNavigationResponseDecisionHandler != nil
}

// WebViewDecidePolicyForNavigationActionDecisionHandler implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDecidePolicyForNavigationActionDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, decisionHandler func(unsafe.Pointer)) {
	if d._WebViewDecidePolicyForNavigationActionDecisionHandler != nil {
		d._WebViewDecidePolicyForNavigationActionDecisionHandler(webView, navigationAction, decisionHandler)
	}
}

// HasWebViewDecidePolicyForNavigationActionDecisionHandler returns true if a handler for WebViewDecidePolicyForNavigationActionDecisionHandler has been set.
func (d *NavigationDelegate) HasWebViewDecidePolicyForNavigationActionDecisionHandler() bool {
	return d._WebViewDecidePolicyForNavigationActionDecisionHandler != nil
}

// WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, preferences IWKWebpagePreferences, decisionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	if d._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler != nil {
		d._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView, navigationAction, preferences, decisionHandler)
	}
}

// HasWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler returns true if a handler for WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler has been set.
func (d *NavigationDelegate) HasWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler() bool {
	return d._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler != nil
}

// WebViewDidCommitNavigation implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidCommitNavigation(webView IWKWebView, navigation IWKNavigation) {
	if d._WebViewDidCommitNavigation != nil {
		d._WebViewDidCommitNavigation(webView, navigation)
	}
}

// HasWebViewDidCommitNavigation returns true if a handler for WebViewDidCommitNavigation has been set.
func (d *NavigationDelegate) HasWebViewDidCommitNavigation() bool {
	return d._WebViewDidCommitNavigation != nil
}

// WebViewDidFailNavigationWithError implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidFailNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */) {
	if d._WebViewDidFailNavigationWithError != nil {
		d._WebViewDidFailNavigationWithError(webView, navigation, error_)
	}
}

// HasWebViewDidFailNavigationWithError returns true if a handler for WebViewDidFailNavigationWithError has been set.
func (d *NavigationDelegate) HasWebViewDidFailNavigationWithError() bool {
	return d._WebViewDidFailNavigationWithError != nil
}

// WebViewDidFailProvisionalNavigationWithError implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidFailProvisionalNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ objc.IObject /* cross-framework: Error */) {
	if d._WebViewDidFailProvisionalNavigationWithError != nil {
		d._WebViewDidFailProvisionalNavigationWithError(webView, navigation, error_)
	}
}

// HasWebViewDidFailProvisionalNavigationWithError returns true if a handler for WebViewDidFailProvisionalNavigationWithError has been set.
func (d *NavigationDelegate) HasWebViewDidFailProvisionalNavigationWithError() bool {
	return d._WebViewDidFailProvisionalNavigationWithError != nil
}

// WebViewDidFinishNavigation implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidFinishNavigation(webView IWKWebView, navigation IWKNavigation) {
	if d._WebViewDidFinishNavigation != nil {
		d._WebViewDidFinishNavigation(webView, navigation)
	}
}

// HasWebViewDidFinishNavigation returns true if a handler for WebViewDidFinishNavigation has been set.
func (d *NavigationDelegate) HasWebViewDidFinishNavigation() bool {
	return d._WebViewDidFinishNavigation != nil
}

// WebViewDidReceiveAuthenticationChallengeCompletionHandler implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidReceiveAuthenticationChallengeCompletionHandler(webView IWKWebView, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	if d._WebViewDidReceiveAuthenticationChallengeCompletionHandler != nil {
		d._WebViewDidReceiveAuthenticationChallengeCompletionHandler(webView, challenge, completionHandler)
	}
}

// HasWebViewDidReceiveAuthenticationChallengeCompletionHandler returns true if a handler for WebViewDidReceiveAuthenticationChallengeCompletionHandler has been set.
func (d *NavigationDelegate) HasWebViewDidReceiveAuthenticationChallengeCompletionHandler() bool {
	return d._WebViewDidReceiveAuthenticationChallengeCompletionHandler != nil
}

// WebViewDidReceiveServerRedirectForProvisionalNavigation implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidReceiveServerRedirectForProvisionalNavigation(webView IWKWebView, navigation IWKNavigation) {
	if d._WebViewDidReceiveServerRedirectForProvisionalNavigation != nil {
		d._WebViewDidReceiveServerRedirectForProvisionalNavigation(webView, navigation)
	}
}

// HasWebViewDidReceiveServerRedirectForProvisionalNavigation returns true if a handler for WebViewDidReceiveServerRedirectForProvisionalNavigation has been set.
func (d *NavigationDelegate) HasWebViewDidReceiveServerRedirectForProvisionalNavigation() bool {
	return d._WebViewDidReceiveServerRedirectForProvisionalNavigation != nil
}

// WebViewDidStartProvisionalNavigation implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewDidStartProvisionalNavigation(webView IWKWebView, navigation IWKNavigation) {
	if d._WebViewDidStartProvisionalNavigation != nil {
		d._WebViewDidStartProvisionalNavigation(webView, navigation)
	}
}

// HasWebViewDidStartProvisionalNavigation returns true if a handler for WebViewDidStartProvisionalNavigation has been set.
func (d *NavigationDelegate) HasWebViewDidStartProvisionalNavigation() bool {
	return d._WebViewDidStartProvisionalNavigation != nil
}

// WebViewNavigationActionDidBecomeDownload implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewNavigationActionDidBecomeDownload(webView IWKWebView, navigationAction IWKNavigationAction, download IWKDownload) {
	if d._WebViewNavigationActionDidBecomeDownload != nil {
		d._WebViewNavigationActionDidBecomeDownload(webView, navigationAction, download)
	}
}

// HasWebViewNavigationActionDidBecomeDownload returns true if a handler for WebViewNavigationActionDidBecomeDownload has been set.
func (d *NavigationDelegate) HasWebViewNavigationActionDidBecomeDownload() bool {
	return d._WebViewNavigationActionDidBecomeDownload != nil
}

// WebViewNavigationResponseDidBecomeDownload implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewNavigationResponseDidBecomeDownload(webView IWKWebView, navigationResponse IWKNavigationResponse, download IWKDownload) {
	if d._WebViewNavigationResponseDidBecomeDownload != nil {
		d._WebViewNavigationResponseDidBecomeDownload(webView, navigationResponse, download)
	}
}

// HasWebViewNavigationResponseDidBecomeDownload returns true if a handler for WebViewNavigationResponseDidBecomeDownload has been set.
func (d *NavigationDelegate) HasWebViewNavigationResponseDidBecomeDownload() bool {
	return d._WebViewNavigationResponseDidBecomeDownload != nil
}

// WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler(webView IWKWebView, backForwardListItem IWKBackForwardListItem, willUseInstantBack bool, completionHandler unsafe.Pointer) {
	if d._WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler != nil {
		d._WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler(webView, backForwardListItem, willUseInstantBack, completionHandler)
	}
}

// HasWebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler returns true if a handler for WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler has been set.
func (d *NavigationDelegate) HasWebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler() bool {
	return d._WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler != nil
}

// WebViewWebContentProcessDidTerminate implements the PNavigationDelegate interface.
func (d *NavigationDelegate) WebViewWebContentProcessDidTerminate(webView IWKWebView) {
	if d._WebViewWebContentProcessDidTerminate != nil {
		d._WebViewWebContentProcessDidTerminate(webView)
	}
}

// HasWebViewWebContentProcessDidTerminate returns true if a handler for WebViewWebContentProcessDidTerminate has been set.
func (d *NavigationDelegate) HasWebViewWebContentProcessDidTerminate() bool {
	return d._WebViewWebContentProcessDidTerminate != nil
}
