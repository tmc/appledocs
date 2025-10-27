// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	SharingServiceDidStopSharing(sharingService ISharingService, share objectivec.IObject)
	HasSharingServiceDidStopSharing() bool
	SharingServiceDidCompleteForItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)
	HasSharingServiceDidCompleteForItemsError() bool
	SharingServiceDidSaveShare(sharingService ISharingService, share objectivec.IObject)
	HasSharingServiceDidSaveShare() bool
}

// CloudSharingServiceDelegate is a delegate implementation builder for the PCloudSharingServiceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CloudSharingServiceDelegate struct {
	_OptionsForSharingServiceShareProvider func(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	_SharingServiceDidStopSharing func(sharingService ISharingService, share objectivec.IObject)
	_SharingServiceDidCompleteForItemsError func(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)
	_SharingServiceDidSaveShare func(sharingService ISharingService, share objectivec.IObject)
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
func (d *CloudSharingServiceDelegate) SetSharingServiceDidStopSharing(f func(sharingService ISharingService, share objectivec.IObject)) {
	d._SharingServiceDidStopSharing = f
}

// SetSharingServiceDidCompleteForItemsError sets the handler for the SharingServiceDidCompleteForItemsError delegate method.
//
// Tells the delegate when the cloud-sharing service completes.
func (d *CloudSharingServiceDelegate) SetSharingServiceDidCompleteForItemsError(f func(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)) {
	d._SharingServiceDidCompleteForItemsError = f
}

// SetSharingServiceDidSaveShare sets the handler for the SharingServiceDidSaveShare delegate method.
//
// Tells the delegate when the cloud-sharing service saves the CloudKit share.
func (d *CloudSharingServiceDelegate) SetSharingServiceDidSaveShare(f func(sharingService ISharingService, share objectivec.IObject)) {
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
func (d *CloudSharingServiceDelegate) SharingServiceDidStopSharing(sharingService ISharingService, share objectivec.IObject) {
	if d._SharingServiceDidStopSharing != nil {
		d._SharingServiceDidStopSharing(sharingService, share)
	}
}

// HasSharingServiceDidStopSharing returns true if a handler for SharingServiceDidStopSharing has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidStopSharing() bool {
	return d._SharingServiceDidStopSharing != nil
}

// SharingServiceDidCompleteForItemsError implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) SharingServiceDidCompleteForItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError) {
	if d._SharingServiceDidCompleteForItemsError != nil {
		d._SharingServiceDidCompleteForItemsError(sharingService, items, error_)
	}
}

// HasSharingServiceDidCompleteForItemsError returns true if a handler for SharingServiceDidCompleteForItemsError has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidCompleteForItemsError() bool {
	return d._SharingServiceDidCompleteForItemsError != nil
}

// SharingServiceDidSaveShare implements the PCloudSharingServiceDelegate interface.
func (d *CloudSharingServiceDelegate) SharingServiceDidSaveShare(sharingService ISharingService, share objectivec.IObject) {
	if d._SharingServiceDidSaveShare != nil {
		d._SharingServiceDidSaveShare(sharingService, share)
	}
}

// HasSharingServiceDidSaveShare returns true if a handler for SharingServiceDidSaveShare has been set.
func (d *CloudSharingServiceDelegate) HasSharingServiceDidSaveShare() bool {
	return d._SharingServiceDidSaveShare != nil
}

// CloudSharingServiceDelegateObject wraps an existing Objective-C object that conforms to the PCloudSharingServiceDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type CloudSharingServiceDelegateObject struct {
	objectivec.Object
}

// NewCloudSharingServiceDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSCloudSharingServiceDelegate protocol.
func NewCloudSharingServiceDelegateObject(obj objectivec.Object) *CloudSharingServiceDelegateObject {
	return &CloudSharingServiceDelegateObject{obj}
}

// Make sure CloudSharingServiceDelegateObject implements PCloudSharingServiceDelegate.
var _ PCloudSharingServiceDelegate = (*CloudSharingServiceDelegateObject)(nil)

// OptionsForSharingServiceShareProvider implements the PCloudSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CloudSharingServiceDelegateObject) OptionsForSharingServiceShareProvider(cloudKitSharingService ISharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions {
	return objc.Send[CloudKitSharingServiceOptions](o.ID, objc.Sel("optionsForSharingService:shareProvider:"), cloudKitSharingService, provider)
}

// HasOptionsForSharingServiceShareProvider returns true; this is a placeholder for optional method checks.
func (o *CloudSharingServiceDelegateObject) HasOptionsForSharingServiceShareProvider() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceDidStopSharing implements the PCloudSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CloudSharingServiceDelegateObject) SharingServiceDidStopSharing(sharingService ISharingService, share objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:didStopSharing:"), sharingService, share)
}

// HasSharingServiceDidStopSharing returns true; this is a placeholder for optional method checks.
func (o *CloudSharingServiceDelegateObject) HasSharingServiceDidStopSharing() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceDidCompleteForItemsError implements the PCloudSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CloudSharingServiceDelegateObject) SharingServiceDidCompleteForItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:didCompleteForItems:error:"), sharingService, items, error_)
}

// HasSharingServiceDidCompleteForItemsError returns true; this is a placeholder for optional method checks.
func (o *CloudSharingServiceDelegateObject) HasSharingServiceDidCompleteForItemsError() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceDidSaveShare implements the PCloudSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CloudSharingServiceDelegateObject) SharingServiceDidSaveShare(sharingService ISharingService, share objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:didSaveShare:"), sharingService, share)
}

// HasSharingServiceDidSaveShare returns true; this is a placeholder for optional method checks.
func (o *CloudSharingServiceDelegateObject) HasSharingServiceDidSaveShare() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
