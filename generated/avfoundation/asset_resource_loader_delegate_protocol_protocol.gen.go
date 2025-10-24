// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

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

// AssetResourceLoaderDelegate is a delegate implementation builder for the PAssetResourceLoaderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AssetResourceLoaderDelegate struct {
	_ResourceLoaderDidCancelAuthenticationChallenge func(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
	_ResourceLoaderDidCancelLoadingRequest func(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest)
	_ResourceLoaderShouldWaitForLoadingOfRequestedResource func(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) bool
	_ResourceLoaderShouldWaitForRenewalOfRequestedResource func(resourceLoader IAVAssetResourceLoader, renewalRequest IAVAssetResourceRenewalRequest) bool
	_ResourceLoaderShouldWaitForResponseToAuthenticationChallenge func(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool
}

// SetResourceLoaderDidCancelAuthenticationChallenge sets the handler for the ResourceLoaderDidCancelAuthenticationChallenge delegate method.
//
// Informs the delegate that a prior authentication challenge has been cancelled.
func (d *AssetResourceLoaderDelegate) SetResourceLoaderDidCancelAuthenticationChallenge(f func(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)) {
	d._ResourceLoaderDidCancelAuthenticationChallenge = f
}

// SetResourceLoaderDidCancelLoadingRequest sets the handler for the ResourceLoaderDidCancelLoadingRequest delegate method.
//
// Informs the delegate that a prior loading request has been cancelled.
func (d *AssetResourceLoaderDelegate) SetResourceLoaderDidCancelLoadingRequest(f func(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest)) {
	d._ResourceLoaderDidCancelLoadingRequest = f
}

// SetResourceLoaderShouldWaitForLoadingOfRequestedResource sets the handler for the ResourceLoaderShouldWaitForLoadingOfRequestedResource delegate method.
//
// Asks the delegate if it wants to load the requested resource.
func (d *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForLoadingOfRequestedResource(f func(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) bool) {
	d._ResourceLoaderShouldWaitForLoadingOfRequestedResource = f
}

// SetResourceLoaderShouldWaitForRenewalOfRequestedResource sets the handler for the ResourceLoaderShouldWaitForRenewalOfRequestedResource delegate method.
//
// Tells the delegate when assistance is required of the application to renew a resource.
func (d *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForRenewalOfRequestedResource(f func(resourceLoader IAVAssetResourceLoader, renewalRequest IAVAssetResourceRenewalRequest) bool) {
	d._ResourceLoaderShouldWaitForRenewalOfRequestedResource = f
}

// SetResourceLoaderShouldWaitForResponseToAuthenticationChallenge sets the handler for the ResourceLoaderShouldWaitForResponseToAuthenticationChallenge delegate method.
//
// Tells the delegate that assistance is required of the application to respond to an authentication challenge.
func (d *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForResponseToAuthenticationChallenge(f func(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool) {
	d._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge = f
}

// ResourceLoaderDidCancelAuthenticationChallenge implements the PAssetResourceLoaderDelegate interface.
func (d *AssetResourceLoaderDelegate) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) {
	if d._ResourceLoaderDidCancelAuthenticationChallenge != nil {
		d._ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader, authenticationChallenge)
	}
}

// HasResourceLoaderDidCancelAuthenticationChallenge returns true if a handler for ResourceLoaderDidCancelAuthenticationChallenge has been set.
func (d *AssetResourceLoaderDelegate) HasResourceLoaderDidCancelAuthenticationChallenge() bool {
	return d._ResourceLoaderDidCancelAuthenticationChallenge != nil
}

// ResourceLoaderDidCancelLoadingRequest implements the PAssetResourceLoaderDelegate interface.
func (d *AssetResourceLoaderDelegate) ResourceLoaderDidCancelLoadingRequest(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) {
	if d._ResourceLoaderDidCancelLoadingRequest != nil {
		d._ResourceLoaderDidCancelLoadingRequest(resourceLoader, loadingRequest)
	}
}

// HasResourceLoaderDidCancelLoadingRequest returns true if a handler for ResourceLoaderDidCancelLoadingRequest has been set.
func (d *AssetResourceLoaderDelegate) HasResourceLoaderDidCancelLoadingRequest() bool {
	return d._ResourceLoaderDidCancelLoadingRequest != nil
}

// ResourceLoaderShouldWaitForLoadingOfRequestedResource implements the PAssetResourceLoaderDelegate interface.
func (d *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) bool {
	if d._ResourceLoaderShouldWaitForLoadingOfRequestedResource != nil {
		return d._ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader, loadingRequest)
	}
	var zero bool
	return zero
}

// HasResourceLoaderShouldWaitForLoadingOfRequestedResource returns true if a handler for ResourceLoaderShouldWaitForLoadingOfRequestedResource has been set.
func (d *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForLoadingOfRequestedResource() bool {
	return d._ResourceLoaderShouldWaitForLoadingOfRequestedResource != nil
}

// ResourceLoaderShouldWaitForRenewalOfRequestedResource implements the PAssetResourceLoaderDelegate interface.
func (d *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader IAVAssetResourceLoader, renewalRequest IAVAssetResourceRenewalRequest) bool {
	if d._ResourceLoaderShouldWaitForRenewalOfRequestedResource != nil {
		return d._ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader, renewalRequest)
	}
	var zero bool
	return zero
}

// HasResourceLoaderShouldWaitForRenewalOfRequestedResource returns true if a handler for ResourceLoaderShouldWaitForRenewalOfRequestedResource has been set.
func (d *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForRenewalOfRequestedResource() bool {
	return d._ResourceLoaderShouldWaitForRenewalOfRequestedResource != nil
}

// ResourceLoaderShouldWaitForResponseToAuthenticationChallenge implements the PAssetResourceLoaderDelegate interface.
func (d *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool {
	if d._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge != nil {
		return d._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader, authenticationChallenge)
	}
	var zero bool
	return zero
}

// HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge returns true if a handler for ResourceLoaderShouldWaitForResponseToAuthenticationChallenge has been set.
func (d *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge() bool {
	return d._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge != nil
}
