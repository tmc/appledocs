// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLDownloadDelegate is the NSURLDownloadDelegate protocol interface.
//
// A protocol that URL download delegates implement to interact with a URL download request.
//
// Availability:
//   - macOS 10.2+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSURLDownloadDelegate
type PURLDownloadDelegate interface {
	// Optional methods
	DownloadCanAuthenticateAgainstProtectionSpace(connection IURLDownload, protectionSpace IURLProtectionSpace) bool
	HasDownloadCanAuthenticateAgainstProtectionSpace() bool
	DownloadDecideDestinationWithSuggestedFilename(download IURLDownload, filename IString)
	HasDownloadDecideDestinationWithSuggestedFilename() bool
	DownloadDidCancelAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge)
	HasDownloadDidCancelAuthenticationChallenge() bool
	DownloadDidCreateDestination(download IURLDownload, path IString)
	HasDownloadDidCreateDestination() bool
	DownloadDidFailWithError(download IURLDownload, error_ IError)
	HasDownloadDidFailWithError() bool
	DownloadDidReceiveAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge)
	HasDownloadDidReceiveAuthenticationChallenge() bool
	DownloadDidReceiveResponse(download IURLDownload, response IURLResponse)
	HasDownloadDidReceiveResponse() bool
	DownloadDidReceiveDataOfLength(download IURLDownload, length uint)
	HasDownloadDidReceiveDataOfLength() bool
	DownloadShouldDecodeSourceDataOfMIMEType(download IURLDownload, encodingType IString) bool
	HasDownloadShouldDecodeSourceDataOfMIMEType() bool
	DownloadWillResumeWithResponseFromByte(download IURLDownload, response IURLResponse, startingByte objectivec.IObject)
	HasDownloadWillResumeWithResponseFromByte() bool
	DownloadWillSendRequestRedirectResponse(download IURLDownload, request IURLRequest, redirectResponse IURLResponse) IURLRequest
	HasDownloadWillSendRequestRedirectResponse() bool
	DownloadDidBegin(download IURLDownload)
	HasDownloadDidBegin() bool
	DownloadDidFinish(download IURLDownload)
	HasDownloadDidFinish() bool
	DownloadShouldUseCredentialStorage(download IURLDownload) bool
	HasDownloadShouldUseCredentialStorage() bool
}

// URLDownloadDelegate is a delegate implementation builder for the PURLDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLDownloadDelegate struct {
	_DownloadCanAuthenticateAgainstProtectionSpace func(connection IURLDownload, protectionSpace IURLProtectionSpace) bool
	_DownloadDecideDestinationWithSuggestedFilename func(download IURLDownload, filename IString)
	_DownloadDidCancelAuthenticationChallenge func(download IURLDownload, challenge IURLAuthenticationChallenge)
	_DownloadDidCreateDestination func(download IURLDownload, path IString)
	_DownloadDidFailWithError func(download IURLDownload, error_ IError)
	_DownloadDidReceiveAuthenticationChallenge func(download IURLDownload, challenge IURLAuthenticationChallenge)
	_DownloadDidReceiveResponse func(download IURLDownload, response IURLResponse)
	_DownloadDidReceiveDataOfLength func(download IURLDownload, length uint)
	_DownloadShouldDecodeSourceDataOfMIMEType func(download IURLDownload, encodingType IString) bool
	_DownloadWillResumeWithResponseFromByte func(download IURLDownload, response IURLResponse, startingByte objectivec.IObject)
	_DownloadWillSendRequestRedirectResponse func(download IURLDownload, request IURLRequest, redirectResponse IURLResponse) IURLRequest
	_DownloadDidBegin func(download IURLDownload)
	_DownloadDidFinish func(download IURLDownload)
	_DownloadShouldUseCredentialStorage func(download IURLDownload) bool
}

// SetDownloadCanAuthenticateAgainstProtectionSpace sets the handler for the DownloadCanAuthenticateAgainstProtectionSpace delegate method.
//
// Sent to determine whether the delegate is able to respond to a protection space’s form of authentication.
func (d *URLDownloadDelegate) SetDownloadCanAuthenticateAgainstProtectionSpace(f func(connection IURLDownload, protectionSpace IURLProtectionSpace) bool) {
	d._DownloadCanAuthenticateAgainstProtectionSpace = f
}

// SetDownloadDecideDestinationWithSuggestedFilename sets the handler for the DownloadDecideDestinationWithSuggestedFilename delegate method.
//
// The delegate receives this message when   has determined a suggested filename for the downloaded file.
func (d *URLDownloadDelegate) SetDownloadDecideDestinationWithSuggestedFilename(f func(download IURLDownload, filename IString)) {
	d._DownloadDecideDestinationWithSuggestedFilename = f
}

// SetDownloadDidCancelAuthenticationChallenge sets the handler for the DownloadDidCancelAuthenticationChallenge delegate method.
//
// Sent if an authentication challenge is canceled due to the protocol implementation encountering an error.
func (d *URLDownloadDelegate) SetDownloadDidCancelAuthenticationChallenge(f func(download IURLDownload, challenge IURLAuthenticationChallenge)) {
	d._DownloadDidCancelAuthenticationChallenge = f
}

// SetDownloadDidCreateDestination sets the handler for the DownloadDidCreateDestination delegate method.
//
// Sent when the destination file is created.
func (d *URLDownloadDelegate) SetDownloadDidCreateDestination(f func(download IURLDownload, path IString)) {
	d._DownloadDidCreateDestination = f
}

// SetDownloadDidFailWithError sets the handler for the DownloadDidFailWithError delegate method.
//
// Sent if the download fails or if an I/O error occurs when the file is written to disk.
func (d *URLDownloadDelegate) SetDownloadDidFailWithError(f func(download IURLDownload, error_ IError)) {
	d._DownloadDidFailWithError = f
}

// SetDownloadDidReceiveAuthenticationChallenge sets the handler for the DownloadDidReceiveAuthenticationChallenge delegate method.
//
// Sent when the URL download must authenticate a challenge in order to download the request.
func (d *URLDownloadDelegate) SetDownloadDidReceiveAuthenticationChallenge(f func(download IURLDownload, challenge IURLAuthenticationChallenge)) {
	d._DownloadDidReceiveAuthenticationChallenge = f
}

// SetDownloadDidReceiveResponse sets the handler for the DownloadDidReceiveResponse delegate method.
//
// Sent when a download object has received sufficient load data to construct the NSURLResponse object for the download.
func (d *URLDownloadDelegate) SetDownloadDidReceiveResponse(f func(download IURLDownload, response IURLResponse)) {
	d._DownloadDidReceiveResponse = f
}

// SetDownloadDidReceiveDataOfLength sets the handler for the DownloadDidReceiveDataOfLength delegate method.
//
// Sent as a download object receives data incrementally.
func (d *URLDownloadDelegate) SetDownloadDidReceiveDataOfLength(f func(download IURLDownload, length uint)) {
	d._DownloadDidReceiveDataOfLength = f
}

// SetDownloadShouldDecodeSourceDataOfMIMEType sets the handler for the DownloadShouldDecodeSourceDataOfMIMEType delegate method.
//
// Sent when a download object determines that the downloaded file is encoded to inquire whether the file should be automatically decoded.
func (d *URLDownloadDelegate) SetDownloadShouldDecodeSourceDataOfMIMEType(f func(download IURLDownload, encodingType IString) bool) {
	d._DownloadShouldDecodeSourceDataOfMIMEType = f
}

// SetDownloadWillResumeWithResponseFromByte sets the handler for the DownloadWillResumeWithResponseFromByte delegate method.
//
// Sent when a download object has received a response from the server after attempting to resume a download.
func (d *URLDownloadDelegate) SetDownloadWillResumeWithResponseFromByte(f func(download IURLDownload, response IURLResponse, startingByte objectivec.IObject)) {
	d._DownloadWillResumeWithResponseFromByte = f
}

// SetDownloadWillSendRequestRedirectResponse sets the handler for the DownloadWillSendRequestRedirectResponse delegate method.
//
// Sent when the download object determines that it must change URLs in order to continue loading a request.
func (d *URLDownloadDelegate) SetDownloadWillSendRequestRedirectResponse(f func(download IURLDownload, request IURLRequest, redirectResponse IURLResponse) IURLRequest) {
	d._DownloadWillSendRequestRedirectResponse = f
}

// SetDownloadDidBegin sets the handler for the DownloadDidBegin delegate method.
//
// Sent immediately after a download object begins a download.
func (d *URLDownloadDelegate) SetDownloadDidBegin(f func(download IURLDownload)) {
	d._DownloadDidBegin = f
}

// SetDownloadDidFinish sets the handler for the DownloadDidFinish delegate method.
//
// Sent when a download object has completed downloading successfully and has written its results to disk.
func (d *URLDownloadDelegate) SetDownloadDidFinish(f func(download IURLDownload)) {
	d._DownloadDidFinish = f
}

// SetDownloadShouldUseCredentialStorage sets the handler for the DownloadShouldUseCredentialStorage delegate method.
//
// Sent to determine whether the URL loader should consult the credential storage to authenticate the download.
func (d *URLDownloadDelegate) SetDownloadShouldUseCredentialStorage(f func(download IURLDownload) bool) {
	d._DownloadShouldUseCredentialStorage = f
}

// DownloadCanAuthenticateAgainstProtectionSpace implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadCanAuthenticateAgainstProtectionSpace(connection IURLDownload, protectionSpace IURLProtectionSpace) bool {
	if d._DownloadCanAuthenticateAgainstProtectionSpace != nil {
		return d._DownloadCanAuthenticateAgainstProtectionSpace(connection, protectionSpace)
	}
	var zero bool
	return zero
}

// HasDownloadCanAuthenticateAgainstProtectionSpace returns true if a handler for DownloadCanAuthenticateAgainstProtectionSpace has been set.
func (d *URLDownloadDelegate) HasDownloadCanAuthenticateAgainstProtectionSpace() bool {
	return d._DownloadCanAuthenticateAgainstProtectionSpace != nil
}

// DownloadDecideDestinationWithSuggestedFilename implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDecideDestinationWithSuggestedFilename(download IURLDownload, filename IString) {
	if d._DownloadDecideDestinationWithSuggestedFilename != nil {
		d._DownloadDecideDestinationWithSuggestedFilename(download, filename)
	}
}

// HasDownloadDecideDestinationWithSuggestedFilename returns true if a handler for DownloadDecideDestinationWithSuggestedFilename has been set.
func (d *URLDownloadDelegate) HasDownloadDecideDestinationWithSuggestedFilename() bool {
	return d._DownloadDecideDestinationWithSuggestedFilename != nil
}

// DownloadDidCancelAuthenticationChallenge implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidCancelAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge) {
	if d._DownloadDidCancelAuthenticationChallenge != nil {
		d._DownloadDidCancelAuthenticationChallenge(download, challenge)
	}
}

// HasDownloadDidCancelAuthenticationChallenge returns true if a handler for DownloadDidCancelAuthenticationChallenge has been set.
func (d *URLDownloadDelegate) HasDownloadDidCancelAuthenticationChallenge() bool {
	return d._DownloadDidCancelAuthenticationChallenge != nil
}

// DownloadDidCreateDestination implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidCreateDestination(download IURLDownload, path IString) {
	if d._DownloadDidCreateDestination != nil {
		d._DownloadDidCreateDestination(download, path)
	}
}

// HasDownloadDidCreateDestination returns true if a handler for DownloadDidCreateDestination has been set.
func (d *URLDownloadDelegate) HasDownloadDidCreateDestination() bool {
	return d._DownloadDidCreateDestination != nil
}

// DownloadDidFailWithError implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidFailWithError(download IURLDownload, error_ IError) {
	if d._DownloadDidFailWithError != nil {
		d._DownloadDidFailWithError(download, error_)
	}
}

// HasDownloadDidFailWithError returns true if a handler for DownloadDidFailWithError has been set.
func (d *URLDownloadDelegate) HasDownloadDidFailWithError() bool {
	return d._DownloadDidFailWithError != nil
}

// DownloadDidReceiveAuthenticationChallenge implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidReceiveAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge) {
	if d._DownloadDidReceiveAuthenticationChallenge != nil {
		d._DownloadDidReceiveAuthenticationChallenge(download, challenge)
	}
}

// HasDownloadDidReceiveAuthenticationChallenge returns true if a handler for DownloadDidReceiveAuthenticationChallenge has been set.
func (d *URLDownloadDelegate) HasDownloadDidReceiveAuthenticationChallenge() bool {
	return d._DownloadDidReceiveAuthenticationChallenge != nil
}

// DownloadDidReceiveResponse implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidReceiveResponse(download IURLDownload, response IURLResponse) {
	if d._DownloadDidReceiveResponse != nil {
		d._DownloadDidReceiveResponse(download, response)
	}
}

// HasDownloadDidReceiveResponse returns true if a handler for DownloadDidReceiveResponse has been set.
func (d *URLDownloadDelegate) HasDownloadDidReceiveResponse() bool {
	return d._DownloadDidReceiveResponse != nil
}

// DownloadDidReceiveDataOfLength implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidReceiveDataOfLength(download IURLDownload, length uint) {
	if d._DownloadDidReceiveDataOfLength != nil {
		d._DownloadDidReceiveDataOfLength(download, length)
	}
}

// HasDownloadDidReceiveDataOfLength returns true if a handler for DownloadDidReceiveDataOfLength has been set.
func (d *URLDownloadDelegate) HasDownloadDidReceiveDataOfLength() bool {
	return d._DownloadDidReceiveDataOfLength != nil
}

// DownloadShouldDecodeSourceDataOfMIMEType implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadShouldDecodeSourceDataOfMIMEType(download IURLDownload, encodingType IString) bool {
	if d._DownloadShouldDecodeSourceDataOfMIMEType != nil {
		return d._DownloadShouldDecodeSourceDataOfMIMEType(download, encodingType)
	}
	var zero bool
	return zero
}

// HasDownloadShouldDecodeSourceDataOfMIMEType returns true if a handler for DownloadShouldDecodeSourceDataOfMIMEType has been set.
func (d *URLDownloadDelegate) HasDownloadShouldDecodeSourceDataOfMIMEType() bool {
	return d._DownloadShouldDecodeSourceDataOfMIMEType != nil
}

// DownloadWillResumeWithResponseFromByte implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadWillResumeWithResponseFromByte(download IURLDownload, response IURLResponse, startingByte objectivec.IObject) {
	if d._DownloadWillResumeWithResponseFromByte != nil {
		d._DownloadWillResumeWithResponseFromByte(download, response, startingByte)
	}
}

// HasDownloadWillResumeWithResponseFromByte returns true if a handler for DownloadWillResumeWithResponseFromByte has been set.
func (d *URLDownloadDelegate) HasDownloadWillResumeWithResponseFromByte() bool {
	return d._DownloadWillResumeWithResponseFromByte != nil
}

// DownloadWillSendRequestRedirectResponse implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadWillSendRequestRedirectResponse(download IURLDownload, request IURLRequest, redirectResponse IURLResponse) IURLRequest {
	if d._DownloadWillSendRequestRedirectResponse != nil {
		return d._DownloadWillSendRequestRedirectResponse(download, request, redirectResponse)
	}
	var zero IURLRequest
	return zero
}

// HasDownloadWillSendRequestRedirectResponse returns true if a handler for DownloadWillSendRequestRedirectResponse has been set.
func (d *URLDownloadDelegate) HasDownloadWillSendRequestRedirectResponse() bool {
	return d._DownloadWillSendRequestRedirectResponse != nil
}

// DownloadDidBegin implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidBegin(download IURLDownload) {
	if d._DownloadDidBegin != nil {
		d._DownloadDidBegin(download)
	}
}

// HasDownloadDidBegin returns true if a handler for DownloadDidBegin has been set.
func (d *URLDownloadDelegate) HasDownloadDidBegin() bool {
	return d._DownloadDidBegin != nil
}

// DownloadDidFinish implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadDidFinish(download IURLDownload) {
	if d._DownloadDidFinish != nil {
		d._DownloadDidFinish(download)
	}
}

// HasDownloadDidFinish returns true if a handler for DownloadDidFinish has been set.
func (d *URLDownloadDelegate) HasDownloadDidFinish() bool {
	return d._DownloadDidFinish != nil
}

// DownloadShouldUseCredentialStorage implements the PURLDownloadDelegate interface.
func (d *URLDownloadDelegate) DownloadShouldUseCredentialStorage(download IURLDownload) bool {
	if d._DownloadShouldUseCredentialStorage != nil {
		return d._DownloadShouldUseCredentialStorage(download)
	}
	var zero bool
	return zero
}

// HasDownloadShouldUseCredentialStorage returns true if a handler for DownloadShouldUseCredentialStorage has been set.
func (d *URLDownloadDelegate) HasDownloadShouldUseCredentialStorage() bool {
	return d._DownloadShouldUseCredentialStorage != nil
}

// URLDownloadDelegateObject wraps an existing Objective-C object that conforms to the PURLDownloadDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLDownloadDelegateObject struct {
	objectivec.Object
}

// NewURLDownloadDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLDownloadDelegate protocol.
func NewURLDownloadDelegateObject(obj objectivec.Object) *URLDownloadDelegateObject {
	return &URLDownloadDelegateObject{obj}
}

// Make sure URLDownloadDelegateObject implements PURLDownloadDelegate.
var _ PURLDownloadDelegate = (*URLDownloadDelegateObject)(nil)

// DownloadCanAuthenticateAgainstProtectionSpace implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadCanAuthenticateAgainstProtectionSpace(connection IURLDownload, protectionSpace IURLProtectionSpace) bool {
	return objc.Send[bool](o.ID, objc.Sel("download:canAuthenticateAgainstProtectionSpace:"), connection, protectionSpace)
}

// HasDownloadCanAuthenticateAgainstProtectionSpace returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadCanAuthenticateAgainstProtectionSpace() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDecideDestinationWithSuggestedFilename implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDecideDestinationWithSuggestedFilename(download IURLDownload, filename IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:decideDestinationWithSuggestedFilename:"), download, filename)
}

// HasDownloadDecideDestinationWithSuggestedFilename returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDecideDestinationWithSuggestedFilename() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidCancelAuthenticationChallenge implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidCancelAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didCancelAuthenticationChallenge:"), download, challenge)
}

// HasDownloadDidCancelAuthenticationChallenge returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidCancelAuthenticationChallenge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidCreateDestination implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidCreateDestination(download IURLDownload, path IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didCreateDestination:"), download, path)
}

// HasDownloadDidCreateDestination returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidCreateDestination() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidFailWithError implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidFailWithError(download IURLDownload, error_ IError) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didFailWithError:"), download, error_)
}

// HasDownloadDidFailWithError returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidFailWithError() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidReceiveAuthenticationChallenge implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidReceiveAuthenticationChallenge(download IURLDownload, challenge IURLAuthenticationChallenge) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didReceiveAuthenticationChallenge:"), download, challenge)
}

// HasDownloadDidReceiveAuthenticationChallenge returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidReceiveAuthenticationChallenge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidReceiveResponse implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidReceiveResponse(download IURLDownload, response IURLResponse) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didReceiveResponse:"), download, response)
}

// HasDownloadDidReceiveResponse returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidReceiveResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidReceiveDataOfLength implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidReceiveDataOfLength(download IURLDownload, length uint) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:didReceiveDataOfLength:"), download, length)
}

// HasDownloadDidReceiveDataOfLength returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidReceiveDataOfLength() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadShouldDecodeSourceDataOfMIMEType implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadShouldDecodeSourceDataOfMIMEType(download IURLDownload, encodingType IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("download:shouldDecodeSourceDataOfMIMEType:"), download, encodingType)
}

// HasDownloadShouldDecodeSourceDataOfMIMEType returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadShouldDecodeSourceDataOfMIMEType() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadWillResumeWithResponseFromByte implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadWillResumeWithResponseFromByte(download IURLDownload, response IURLResponse, startingByte objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("download:willResumeWithResponse:fromByte:"), download, response, startingByte)
}

// HasDownloadWillResumeWithResponseFromByte returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadWillResumeWithResponseFromByte() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadWillSendRequestRedirectResponse implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadWillSendRequestRedirectResponse(download IURLDownload, request IURLRequest, redirectResponse IURLResponse) IURLRequest {
	return objc.Send[IURLRequest](o.ID, objc.Sel("download:willSendRequest:redirectResponse:"), download, request, redirectResponse)
}

// HasDownloadWillSendRequestRedirectResponse returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadWillSendRequestRedirectResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidBegin implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidBegin(download IURLDownload) {
	objc.Send[objc.ID](o.ID, objc.Sel("downloadDidBegin:"), download)
}

// HasDownloadDidBegin returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidBegin() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadDidFinish implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadDidFinish(download IURLDownload) {
	objc.Send[objc.ID](o.ID, objc.Sel("downloadDidFinish:"), download)
}

// HasDownloadDidFinish returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadDidFinish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// DownloadShouldUseCredentialStorage implements the PURLDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLDownloadDelegateObject) DownloadShouldUseCredentialStorage(download IURLDownload) bool {
	return objc.Send[bool](o.ID, objc.Sel("downloadShouldUseCredentialStorage:"), download)
}

// HasDownloadShouldUseCredentialStorage returns true; this is a placeholder for optional method checks.
func (o *URLDownloadDelegateObject) HasDownloadShouldUseCredentialStorage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
