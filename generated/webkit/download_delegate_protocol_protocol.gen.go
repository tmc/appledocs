// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/func(unsafe"
)

// PDownloadDelegate is the WKDownloadDelegate protocol interface.
//
// A protocol you implement to track download progress and handle redirects, authentication challenges, and failures.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.webkit/documentation/WebKit/WKDownloadDelegate
type PDownloadDelegate interface {
	// Required methods
	DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download IWKDownload, response foundation.URLResponse, suggestedFilename objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))/* debug [protocol_interface/required_method]: DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler */
	// Optional methods
	DownloadDecidePlaceholderPolicy(download IWKDownload, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	HasDownloadDecidePlaceholderPolicy() bool
	DownloadDidFailWithErrorResumeData(download IWKDownload, error_ objc.IObject /* cross-framework: Error */, resumeData objc.IObject /* cross-framework: NSData */)
	HasDownloadDidFailWithErrorResumeData() bool
	DownloadDidReceiveAuthenticationChallengeCompletionHandler(download IWKDownload, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	HasDownloadDidReceiveAuthenticationChallengeCompletionHandler() bool
	DownloadDidReceiveFinalURL(download IWKDownload, url objc.IObject /* cross-framework: NSURL */)
	HasDownloadDidReceiveFinalURL() bool
	DownloadDidReceivePlaceholderURLCompletionHandler(download IWKDownload, url objc.IObject /* cross-framework: NSURL */, completionHandler func())
	HasDownloadDidReceivePlaceholderURLCompletionHandler() bool
	DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download IWKDownload, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(unsafe.Pointer))
	HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler() bool
	DownloadDidFinish(download IWKDownload)
	HasDownloadDidFinish() bool
}

// DownloadDelegate is a delegate implementation builder for the PDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DownloadDelegate struct {
	_DownloadDecidePlaceholderPolicy func(download IWKDownload, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	_DownloadDidFailWithErrorResumeData func(download IWKDownload, error_ objc.IObject /* cross-framework: Error */, resumeData objc.IObject /* cross-framework: NSData */)
	_DownloadDidReceiveAuthenticationChallengeCompletionHandler func(download IWKDownload, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))
	_DownloadDidReceiveFinalURL func(download IWKDownload, url objc.IObject /* cross-framework: NSURL */)
	_DownloadDidReceivePlaceholderURLCompletionHandler func(download IWKDownload, url objc.IObject /* cross-framework: NSURL */, completionHandler func())
	_DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler func(download IWKDownload, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(unsafe.Pointer))
	_DownloadDidFinish func(download IWKDownload)
	_DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler func(download IWKDownload, response foundation.URLResponse, suggestedFilename objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))
}

// SetDownloadDecidePlaceholderPolicy sets the handler for the DownloadDecidePlaceholderPolicy delegate method.
func (d *DownloadDelegate) SetDownloadDecidePlaceholderPolicy(f func(download IWKDownload, completionHandler func(unsafe.Pointer, unsafe.Pointer))) {
	d._DownloadDecidePlaceholderPolicy = f
}

// SetDownloadDidFailWithErrorResumeData sets the handler for the DownloadDidFailWithErrorResumeData delegate method.
//
// Tells the delegate that the download failed, with error information and data you can use to restart the download.
func (d *DownloadDelegate) SetDownloadDidFailWithErrorResumeData(f func(download IWKDownload, error_ objc.IObject /* cross-framework: Error */, resumeData objc.IObject /* cross-framework: NSData */)) {
	d._DownloadDidFailWithErrorResumeData = f
}

// SetDownloadDidReceiveAuthenticationChallengeCompletionHandler sets the handler for the DownloadDidReceiveAuthenticationChallengeCompletionHandler delegate method.
//
// Asks the delegate to respond to an authentication challenge.
func (d *DownloadDelegate) SetDownloadDidReceiveAuthenticationChallengeCompletionHandler(f func(download IWKDownload, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer))) {
	d._DownloadDidReceiveAuthenticationChallengeCompletionHandler = f
}

// SetDownloadDidReceiveFinalURL sets the handler for the DownloadDidReceiveFinalURL delegate method.
func (d *DownloadDelegate) SetDownloadDidReceiveFinalURL(f func(download IWKDownload, url objc.IObject /* cross-framework: NSURL */)) {
	d._DownloadDidReceiveFinalURL = f
}

// SetDownloadDidReceivePlaceholderURLCompletionHandler sets the handler for the DownloadDidReceivePlaceholderURLCompletionHandler delegate method.
func (d *DownloadDelegate) SetDownloadDidReceivePlaceholderURLCompletionHandler(f func(download IWKDownload, url objc.IObject /* cross-framework: NSURL */, completionHandler func())) {
	d._DownloadDidReceivePlaceholderURLCompletionHandler = f
}

// SetDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler sets the handler for the DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler delegate method.
//
// Asks the delegate to respond to the download’s redirect response.
func (d *DownloadDelegate) SetDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(f func(download IWKDownload, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(unsafe.Pointer))) {
	d._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler = f
}

// SetDownloadDidFinish sets the handler for the DownloadDidFinish delegate method.
//
// Tells the delegate that the download finished.
func (d *DownloadDelegate) SetDownloadDidFinish(f func(download IWKDownload)) {
	d._DownloadDidFinish = f
}

// SetDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler sets the handler for the DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler delegate method.
//
// Asks the delegate to provide a file destination where the system should write the download data.
func (d *DownloadDelegate) SetDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(f func(download IWKDownload, response foundation.URLResponse, suggestedFilename objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer))) {
	d._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler = f
}

// DownloadDecidePlaceholderPolicy implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDecidePlaceholderPolicy(download IWKDownload, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	if d._DownloadDecidePlaceholderPolicy != nil {
		d._DownloadDecidePlaceholderPolicy(download, completionHandler)
	}
}

// HasDownloadDecidePlaceholderPolicy returns true if a handler for DownloadDecidePlaceholderPolicy has been set.
func (d *DownloadDelegate) HasDownloadDecidePlaceholderPolicy() bool {
	return d._DownloadDecidePlaceholderPolicy != nil
}

// DownloadDidFailWithErrorResumeData implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDidFailWithErrorResumeData(download IWKDownload, error_ objc.IObject /* cross-framework: Error */, resumeData objc.IObject /* cross-framework: NSData */) {
	if d._DownloadDidFailWithErrorResumeData != nil {
		d._DownloadDidFailWithErrorResumeData(download, error_, resumeData)
	}
}

// HasDownloadDidFailWithErrorResumeData returns true if a handler for DownloadDidFailWithErrorResumeData has been set.
func (d *DownloadDelegate) HasDownloadDidFailWithErrorResumeData() bool {
	return d._DownloadDidFailWithErrorResumeData != nil
}

// DownloadDidReceiveAuthenticationChallengeCompletionHandler implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDidReceiveAuthenticationChallengeCompletionHandler(download IWKDownload, challenge foundation.URLAuthenticationChallenge, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	if d._DownloadDidReceiveAuthenticationChallengeCompletionHandler != nil {
		d._DownloadDidReceiveAuthenticationChallengeCompletionHandler(download, challenge, completionHandler)
	}
}

// HasDownloadDidReceiveAuthenticationChallengeCompletionHandler returns true if a handler for DownloadDidReceiveAuthenticationChallengeCompletionHandler has been set.
func (d *DownloadDelegate) HasDownloadDidReceiveAuthenticationChallengeCompletionHandler() bool {
	return d._DownloadDidReceiveAuthenticationChallengeCompletionHandler != nil
}

// DownloadDidReceiveFinalURL implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDidReceiveFinalURL(download IWKDownload, url objc.IObject /* cross-framework: NSURL */) {
	if d._DownloadDidReceiveFinalURL != nil {
		d._DownloadDidReceiveFinalURL(download, url)
	}
}

// HasDownloadDidReceiveFinalURL returns true if a handler for DownloadDidReceiveFinalURL has been set.
func (d *DownloadDelegate) HasDownloadDidReceiveFinalURL() bool {
	return d._DownloadDidReceiveFinalURL != nil
}

// DownloadDidReceivePlaceholderURLCompletionHandler implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDidReceivePlaceholderURLCompletionHandler(download IWKDownload, url objc.IObject /* cross-framework: NSURL */, completionHandler func()) {
	if d._DownloadDidReceivePlaceholderURLCompletionHandler != nil {
		d._DownloadDidReceivePlaceholderURLCompletionHandler(download, url, completionHandler)
	}
}

// HasDownloadDidReceivePlaceholderURLCompletionHandler returns true if a handler for DownloadDidReceivePlaceholderURLCompletionHandler has been set.
func (d *DownloadDelegate) HasDownloadDidReceivePlaceholderURLCompletionHandler() bool {
	return d._DownloadDidReceivePlaceholderURLCompletionHandler != nil
}

// DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download IWKDownload, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(unsafe.Pointer)) {
	if d._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler != nil {
		d._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download, response, request, decisionHandler)
	}
}

// HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler returns true if a handler for DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler has been set.
func (d *DownloadDelegate) HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler() bool {
	return d._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler != nil
}

// DownloadDidFinish implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDidFinish(download IWKDownload) {
	if d._DownloadDidFinish != nil {
		d._DownloadDidFinish(download)
	}
}

// HasDownloadDidFinish returns true if a handler for DownloadDidFinish has been set.
func (d *DownloadDelegate) HasDownloadDidFinish() bool {
	return d._DownloadDidFinish != nil
}

// DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler implements the PDownloadDelegate interface.
func (d *DownloadDelegate) DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download IWKDownload, response foundation.URLResponse, suggestedFilename objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer)) {
	if d._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler != nil {
		d._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download, response, suggestedFilename, completionHandler)
	}
}

// HasDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler returns true if a handler for DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler has been set.
func (d *DownloadDelegate) HasDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler() bool {
	return d._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler != nil
}
