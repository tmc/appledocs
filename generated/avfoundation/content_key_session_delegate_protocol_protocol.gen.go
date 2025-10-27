// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	ContentKeySessionDidProvideContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	// Optional methods
	ContentKeySessionContentKeyRequestDidFailWithError(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err foundation.foundation.INSError)
	HasContentKeySessionContentKeyRequestDidFailWithError() bool
	ContentKeySessionContentKeyRequestDidSucceed(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	HasContentKeySessionContentKeyRequestDidSucceed() bool
	ContentKeySessionDidProvidePersistableContentKeyRequest(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest)
	HasContentKeySessionDidProvidePersistableContentKeyRequest() bool
	ContentKeySessionDidProvideContentKeyRequestsForInitializationData(session IAVContentKeySession, keyRequests []ContentKeyRequest, initializationData foundation.foundation.INSData)
	HasContentKeySessionDidProvideContentKeyRequestsForInitializationData() bool
	ContentKeySessionDidProvideRenewingContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
	HasContentKeySessionDidProvideRenewingContentKeyRequest() bool
	ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session IAVContentKeySession, persistableContentKey foundation.foundation.INSData, keyIdentifier objectivec.IObject)
	HasContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier() bool
	ContentKeySessionExternalProtectionStatusDidChangeForContentKey(session IAVContentKeySession, contentKey IAVContentKey)
	HasContentKeySessionExternalProtectionStatusDidChangeForContentKey() bool
	ContentKeySessionShouldRetryContentKeyRequestReason(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason ContentKeyRequestRetryReason) bool
	HasContentKeySessionShouldRetryContentKeyRequestReason() bool
	ContentKeySessionContentProtectionSessionIdentifierDidChange(session IAVContentKeySession)
	HasContentKeySessionContentProtectionSessionIdentifierDidChange() bool
	ContentKeySessionDidGenerateExpiredSessionReport(session IAVContentKeySession)
	HasContentKeySessionDidGenerateExpiredSessionReport() bool
}
