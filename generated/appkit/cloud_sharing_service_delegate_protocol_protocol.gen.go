// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCloudSharingServiceDelegate is the NSCloudSharingServiceDelegate protocol interface.
//
// A set of methods for responding to the life cycle events of the cloud-sharing service.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCloudSharingServiceDelegate
type PCloudSharingServiceDelegate interface {
	// Optional methods
	OptionsForSharingServiceShareProvider(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	HasOptionsForSharingServiceShareProvider() bool
	SharingServiceDidStopSharing(sharingService ISharingService, share objc.IObject)
	HasSharingServiceDidStopSharing() bool
	SharingServiceDidCompleteForItemsError(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)
	HasSharingServiceDidCompleteForItemsError() bool
	SharingServiceDidSaveShare(sharingService ISharingService, share objc.IObject)
	HasSharingServiceDidSaveShare() bool
}

// CloudSharingServiceDelegate is a delegate implementation builder for the PCloudSharingServiceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CloudSharingServiceDelegate struct {
	_OptionsForSharingServiceShareProvider func(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	_SharingServiceDidStopSharing func(sharingService ISharingService, share objc.IObject)
	_SharingServiceDidCompleteForItemsError func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)
	_SharingServiceDidSaveShare func(sharingService ISharingService, share objc.IObject)
}

// SetOptionsForSharingServiceShareProvider sets the handler for the OptionsForSharingServiceShareProvider delegate method.
//
// Asks the delegate for the participant options for the cloud-sharing service.
func (d *CloudSharingServiceDelegate) SetOptionsForSharingServiceShareProvider(f func(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions) {
	d._OptionsForSharingServiceShareProvider = f
}

// SetSharingServiceDidStopSharing sets the handler for the SharingServiceDidStopSharing delegate method.
//
// Tells the delegate when the user stops sharing the CloudKit share.
func (d *CloudSharingServiceDelegate) SetSharingServiceDidStopSharing(f func(sharingService ISharingService, share objc.IObject)) {
	d._SharingServiceDidStopSharing = f
}

// SetSharingServiceDidCompleteForItemsError sets the handler for the SharingServiceDidCompleteForItemsError delegate method.
//
// Tells the delegate when the cloud-sharing service completes.
func (d *CloudSharingServiceDelegate) SetSharingServiceDidCompleteForItemsError(f func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)) {
	d._SharingServiceDidCompleteForItemsError = f
}

// SetSharingServiceDidSaveShare sets the handler for the SharingServiceDidSaveShare delegate method.
//
// Tells the delegate when the cloud-sharing service saves the CloudKit share.
func (d *CloudSharingServiceDelegate) SetSharingServiceDidSaveShare(f func(sharingService ISharingService, share objc.IObject)) {
	d._SharingServiceDidSaveShare = f
}

// OptionsForSharingServiceShareProvider implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) OptionsForSharingServiceShareProvider(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions {
	if d._OptionsForSharingServiceShareProvider != nil {
		return d._OptionsForSharingServiceShareProvider(cloudKitSharingService, provider)
	}
	var zero CloudKitSharingServiceOptions
	return zero
}

// HasOptionsForSharingServiceShareProvider returns true if a handler for OptionsForSharingServiceShareProvider has been set.
func (d *CloudSharingServiceDelegate) HasOptionsForSharingServiceShareProvider() bool {
	return d._OptionsForSharingServiceShareProvider != nil
}

// SharingServiceDidStopSharing implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) SharingServiceDidStopSharing(sharingService ISharingService, share objc.IObject) {
	if d._SharingServiceDidStopSharing != nil {
		d._SharingServiceDidStopSharing(sharingService, share)
	}
}

// HasSharingServiceDidStopSharing returns true if a handler for SharingServiceDidStopSharing has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidStopSharing() bool {
	return d._SharingServiceDidStopSharing != nil
}

// SharingServiceDidCompleteForItemsError implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) SharingServiceDidCompleteForItemsError(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */) {
	if d._SharingServiceDidCompleteForItemsError != nil {
		d._SharingServiceDidCompleteForItemsError(sharingService, items, error_)
	}
}

// HasSharingServiceDidCompleteForItemsError returns true if a handler for SharingServiceDidCompleteForItemsError has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidCompleteForItemsError() bool {
	return d._SharingServiceDidCompleteForItemsError != nil
}

// SharingServiceDidSaveShare implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) SharingServiceDidSaveShare(sharingService ISharingService, share objc.IObject) {
	if d._SharingServiceDidSaveShare != nil {
		d._SharingServiceDidSaveShare(sharingService, share)
	}
}

// HasSharingServiceDidSaveShare returns true if a handler for SharingServiceDidSaveShare has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidSaveShare() bool {
	return d._SharingServiceDidSaveShare != nil
}
