// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAssetResourceLoaderDelegate is the AVAssetResourceLoaderDelegate protocol interface.
//
// Methods you can implement to handle resource-loading requests coming from a URL asset.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVAssetResourceLoaderDelegate
type PAssetResourceLoaderDelegate interface {
	// Optional methods
	ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
	HasResourceLoaderDidCancelAuthenticationChallenge() bool
	ResourceLoaderDidCancelLoadingRequest(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest)
	HasResourceLoaderDidCancelLoadingRequest() bool
	ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) bool
	HasResourceLoaderShouldWaitForLoadingOfRequestedResource() bool
	ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader IAVAssetResourceLoader, renewalRequest IAVAssetResourceRenewalRequest) bool
	HasResourceLoaderShouldWaitForRenewalOfRequestedResource() bool
	ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool
	HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge() bool
}
