// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PContentKeySessionDelegate is the AVContentKeySessionDelegate protocol interface.
//
// A protocol that handles content key requests.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.3+
//   - iPadOS 10.3+
//   - macOS 10.12.4+
//   - tvOS 10.2+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVContentKeySessionDelegate
type PContentKeySessionDelegate interface {
	// Required methods
	ContentKeySessionDidProvideContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest)/* debug [protocol_interface/required_method]: ContentKeySessionDidProvideContentKeyRequest */
	// Optional methods
	ContentKeySessionContentKeyRequestDidFailWithError(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err Error)
	HasContentKeySessionContentKeyRequestDidFailWithError() bool
	ContentKeySessionContentKeyRequestDidSucceed(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	HasContentKeySessionContentKeyRequestDidSucceed() bool
	ContentKeySessionDidProvidePersistableContentKeyRequest(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest)
	HasContentKeySessionDidProvidePersistableContentKeyRequest() bool
	ContentKeySessionDidProvideContentKeyRequestsForInitializationData(session IAVContentKeySession, keyRequests []ContentKeyRequest, initializationData objc.IObject /* cross-framework: NSData */)
	HasContentKeySessionDidProvideContentKeyRequestsForInitializationData() bool
	ContentKeySessionDidProvideRenewingContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	HasContentKeySessionDidProvideRenewingContentKeyRequest() bool
	ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session IAVContentKeySession, persistableContentKey objc.IObject /* cross-framework: NSData */, keyIdentifier objc.IObject)
	HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool
	ContentKeySessionExternalProtectionStatusDidChangeForContentKey(session IAVContentKeySession, contentKey IAVContentKey)
	HasContentKeySessionExternalProtectionStatusDidChangeForContentKey() bool
	ContentKeySessionShouldRetryContentKeyRequestReason(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason ContentKeyRequestRetryReason /* typedef */) bool
	HasContentKeySessionShouldRetryContentKeyRequestReason() bool
	ContentKeySessionContentProtectionSessionIdentifierDidChange(session IAVContentKeySession)
	HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool
	ContentKeySessionDidGenerateExpiredSessionReport(session IAVContentKeySession)
	HasContentKeySessionDidGenerateExpiredSessionReport() bool
}

// ContentKeySessionDelegate is a delegate implementation builder for the PContentKeySessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ContentKeySessionDelegate struct {
	_ContentKeySessionContentKeyRequestDidFailWithError func(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err Error)
	_ContentKeySessionContentKeyRequestDidSucceed func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	_ContentKeySessionDidProvidePersistableContentKeyRequest func(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest)
	_ContentKeySessionDidProvideContentKeyRequestsForInitializationData func(session IAVContentKeySession, keyRequests []ContentKeyRequest, initializationData objc.IObject /* cross-framework: NSData */)
	_ContentKeySessionDidProvideRenewingContentKeyRequest func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	_ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier func(session IAVContentKeySession, persistableContentKey objc.IObject /* cross-framework: NSData */, keyIdentifier objc.IObject)
	_ContentKeySessionExternalProtectionStatusDidChangeForContentKey func(session IAVContentKeySession, contentKey IAVContentKey)
	_ContentKeySessionShouldRetryContentKeyRequestReason func(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason ContentKeyRequestRetryReason /* typedef */) bool
	_ContentKeySessionContentProtectionSessionIdentifierDidChange func(session IAVContentKeySession)
	_ContentKeySessionDidGenerateExpiredSessionReport func(session IAVContentKeySession)
	_ContentKeySessionDidProvideContentKeyRequest func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
}

// SetContentKeySessionContentKeyRequestDidFailWithError sets the handler for the ContentKeySessionContentKeyRequestDidFailWithError delegate method.
//
// Tells the receiver that the content key request failed.
func (d *ContentKeySessionDelegate) SetContentKeySessionContentKeyRequestDidFailWithError(f func(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err Error)) {
	d._ContentKeySessionContentKeyRequestDidFailWithError = f
}

// SetContentKeySessionContentKeyRequestDidSucceed sets the handler for the ContentKeySessionContentKeyRequestDidSucceed delegate method.
//
// Tells the content key session that the response to a content key requeset was successfully processed.
func (d *ContentKeySessionDelegate) SetContentKeySessionContentKeyRequestDidSucceed(f func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)) {
	d._ContentKeySessionContentKeyRequestDidSucceed = f
}

// SetContentKeySessionDidProvidePersistableContentKeyRequest sets the handler for the ContentKeySessionDidProvidePersistableContentKeyRequest delegate method.
//
// Provides the receiver with a new content key request object to process a persistable content key.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidProvidePersistableContentKeyRequest(f func(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest)) {
	d._ContentKeySessionDidProvidePersistableContentKeyRequest = f
}

// SetContentKeySessionDidProvideContentKeyRequestsForInitializationData sets the handler for the ContentKeySessionDidProvideContentKeyRequestsForInitializationData delegate method.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidProvideContentKeyRequestsForInitializationData(f func(session IAVContentKeySession, keyRequests []ContentKeyRequest, initializationData objc.IObject /* cross-framework: NSData */)) {
	d._ContentKeySessionDidProvideContentKeyRequestsForInitializationData = f
}

// SetContentKeySessionDidProvideRenewingContentKeyRequest sets the handler for the ContentKeySessionDidProvideRenewingContentKeyRequest delegate method.
//
// Provides the receiver with a new content key request object for the renewal of an existing content key.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidProvideRenewingContentKeyRequest(f func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)) {
	d._ContentKeySessionDidProvideRenewingContentKeyRequest = f
}

// SetContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier sets the handler for the ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier delegate method.
//
// Provides the receiver with an updated persistable content key for a specific key request.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(f func(session IAVContentKeySession, persistableContentKey objc.IObject /* cross-framework: NSData */, keyIdentifier objc.IObject)) {
	d._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier = f
}

// SetContentKeySessionExternalProtectionStatusDidChangeForContentKey sets the handler for the ContentKeySessionExternalProtectionStatusDidChangeForContentKey delegate method.
//
// Tells the delegate when external protection state has changed.
func (d *ContentKeySessionDelegate) SetContentKeySessionExternalProtectionStatusDidChangeForContentKey(f func(session IAVContentKeySession, contentKey IAVContentKey)) {
	d._ContentKeySessionExternalProtectionStatusDidChangeForContentKey = f
}

// SetContentKeySessionShouldRetryContentKeyRequestReason sets the handler for the ContentKeySessionShouldRetryContentKeyRequestReason delegate method.
//
// Provides the receiver with a content key request object to retry.
func (d *ContentKeySessionDelegate) SetContentKeySessionShouldRetryContentKeyRequestReason(f func(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason ContentKeyRequestRetryReason /* typedef */) bool) {
	d._ContentKeySessionShouldRetryContentKeyRequestReason = f
}

// SetContentKeySessionContentProtectionSessionIdentifierDidChange sets the handler for the ContentKeySessionContentProtectionSessionIdentifierDidChange delegate method.
//
// Tells the receiver the content protection session identifier changed.
func (d *ContentKeySessionDelegate) SetContentKeySessionContentProtectionSessionIdentifierDidChange(f func(session IAVContentKeySession)) {
	d._ContentKeySessionContentProtectionSessionIdentifierDidChange = f
}

// SetContentKeySessionDidGenerateExpiredSessionReport sets the handler for the ContentKeySessionDidGenerateExpiredSessionReport delegate method.
//
// Notifies the sender that an expired session report has been generated.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidGenerateExpiredSessionReport(f func(session IAVContentKeySession)) {
	d._ContentKeySessionDidGenerateExpiredSessionReport = f
}

// SetContentKeySessionDidProvideContentKeyRequest sets the handler for the ContentKeySessionDidProvideContentKeyRequest delegate method.
//
// Provides the receiver with a new content key request object.
func (d *ContentKeySessionDelegate) SetContentKeySessionDidProvideContentKeyRequest(f func(session IAVContentKeySession, keyRequest IAVContentKeyRequest)) {
	d._ContentKeySessionDidProvideContentKeyRequest = f
}

// ContentKeySessionContentKeyRequestDidFailWithError implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionContentKeyRequestDidFailWithError(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err Error) {
	if d._ContentKeySessionContentKeyRequestDidFailWithError != nil {
		d._ContentKeySessionContentKeyRequestDidFailWithError(session, keyRequest, err)
	}
}

// HasContentKeySessionContentKeyRequestDidFailWithError returns true if a handler for ContentKeySessionContentKeyRequestDidFailWithError has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionContentKeyRequestDidFailWithError() bool {
	return d._ContentKeySessionContentKeyRequestDidFailWithError != nil
}

// ContentKeySessionContentKeyRequestDidSucceed implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionContentKeyRequestDidSucceed(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	if d._ContentKeySessionContentKeyRequestDidSucceed != nil {
		d._ContentKeySessionContentKeyRequestDidSucceed(session, keyRequest)
	}
}

// HasContentKeySessionContentKeyRequestDidSucceed returns true if a handler for ContentKeySessionContentKeyRequestDidSucceed has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionContentKeyRequestDidSucceed() bool {
	return d._ContentKeySessionContentKeyRequestDidSucceed != nil
}

// ContentKeySessionDidProvidePersistableContentKeyRequest implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidProvidePersistableContentKeyRequest(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest) {
	if d._ContentKeySessionDidProvidePersistableContentKeyRequest != nil {
		d._ContentKeySessionDidProvidePersistableContentKeyRequest(session, keyRequest)
	}
}

// HasContentKeySessionDidProvidePersistableContentKeyRequest returns true if a handler for ContentKeySessionDidProvidePersistableContentKeyRequest has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidProvidePersistableContentKeyRequest() bool {
	return d._ContentKeySessionDidProvidePersistableContentKeyRequest != nil
}

// ContentKeySessionDidProvideContentKeyRequestsForInitializationData implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidProvideContentKeyRequestsForInitializationData(session IAVContentKeySession, keyRequests []ContentKeyRequest, initializationData objc.IObject /* cross-framework: NSData */) {
	if d._ContentKeySessionDidProvideContentKeyRequestsForInitializationData != nil {
		d._ContentKeySessionDidProvideContentKeyRequestsForInitializationData(session, keyRequests, initializationData)
	}
}

// HasContentKeySessionDidProvideContentKeyRequestsForInitializationData returns true if a handler for ContentKeySessionDidProvideContentKeyRequestsForInitializationData has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidProvideContentKeyRequestsForInitializationData() bool {
	return d._ContentKeySessionDidProvideContentKeyRequestsForInitializationData != nil
}

// ContentKeySessionDidProvideRenewingContentKeyRequest implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidProvideRenewingContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	if d._ContentKeySessionDidProvideRenewingContentKeyRequest != nil {
		d._ContentKeySessionDidProvideRenewingContentKeyRequest(session, keyRequest)
	}
}

// HasContentKeySessionDidProvideRenewingContentKeyRequest returns true if a handler for ContentKeySessionDidProvideRenewingContentKeyRequest has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidProvideRenewingContentKeyRequest() bool {
	return d._ContentKeySessionDidProvideRenewingContentKeyRequest != nil
}

// ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session IAVContentKeySession, persistableContentKey objc.IObject /* cross-framework: NSData */, keyIdentifier objc.IObject) {
	if d._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier != nil {
		d._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session, persistableContentKey, keyIdentifier)
	}
}

// HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier returns true if a handler for ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool {
	return d._ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier != nil
}

// ContentKeySessionExternalProtectionStatusDidChangeForContentKey implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionExternalProtectionStatusDidChangeForContentKey(session IAVContentKeySession, contentKey IAVContentKey) {
	if d._ContentKeySessionExternalProtectionStatusDidChangeForContentKey != nil {
		d._ContentKeySessionExternalProtectionStatusDidChangeForContentKey(session, contentKey)
	}
}

// HasContentKeySessionExternalProtectionStatusDidChangeForContentKey returns true if a handler for ContentKeySessionExternalProtectionStatusDidChangeForContentKey has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionExternalProtectionStatusDidChangeForContentKey() bool {
	return d._ContentKeySessionExternalProtectionStatusDidChangeForContentKey != nil
}

// ContentKeySessionShouldRetryContentKeyRequestReason implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionShouldRetryContentKeyRequestReason(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason ContentKeyRequestRetryReason /* typedef */) bool {
	if d._ContentKeySessionShouldRetryContentKeyRequestReason != nil {
		return d._ContentKeySessionShouldRetryContentKeyRequestReason(session, keyRequest, retryReason)
	}
	var zero bool
	return zero
}

// HasContentKeySessionShouldRetryContentKeyRequestReason returns true if a handler for ContentKeySessionShouldRetryContentKeyRequestReason has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionShouldRetryContentKeyRequestReason() bool {
	return d._ContentKeySessionShouldRetryContentKeyRequestReason != nil
}

// ContentKeySessionContentProtectionSessionIdentifierDidChange implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionContentProtectionSessionIdentifierDidChange(session IAVContentKeySession) {
	if d._ContentKeySessionContentProtectionSessionIdentifierDidChange != nil {
		d._ContentKeySessionContentProtectionSessionIdentifierDidChange(session)
	}
}

// HasContentKeySessionContentProtectionSessionIdentifierDidChange returns true if a handler for ContentKeySessionContentProtectionSessionIdentifierDidChange has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool {
	return d._ContentKeySessionContentProtectionSessionIdentifierDidChange != nil
}

// ContentKeySessionDidGenerateExpiredSessionReport implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidGenerateExpiredSessionReport(session IAVContentKeySession) {
	if d._ContentKeySessionDidGenerateExpiredSessionReport != nil {
		d._ContentKeySessionDidGenerateExpiredSessionReport(session)
	}
}

// HasContentKeySessionDidGenerateExpiredSessionReport returns true if a handler for ContentKeySessionDidGenerateExpiredSessionReport has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidGenerateExpiredSessionReport() bool {
	return d._ContentKeySessionDidGenerateExpiredSessionReport != nil
}

// ContentKeySessionDidProvideContentKeyRequest implements the PContentKeySessionDelegate interface.
func (d *ContentKeySessionDelegate) ContentKeySessionDidProvideContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	if d._ContentKeySessionDidProvideContentKeyRequest != nil {
		d._ContentKeySessionDidProvideContentKeyRequest(session, keyRequest)
	}
}

// HasContentKeySessionDidProvideContentKeyRequest returns true if a handler for ContentKeySessionDidProvideContentKeyRequest has been set.
func (d *ContentKeySessionDelegate) HasContentKeySessionDidProvideContentKeyRequest() bool {
	return d._ContentKeySessionDidProvideContentKeyRequest != nil
}
