// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSharingServicePickerDelegate is the NSSharingServicePickerDelegate protocol interface.
//
// An interface for managing content for the macOS share sheet.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePickerDelegate
type PSharingServicePickerDelegate interface {
	// Optional methods
	SharingServicePickerDelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) unsafe.Pointer
	HasSharingServicePickerDelegateForSharingService() bool
	SharingServicePickerDidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService)
	HasSharingServicePickerDidChooseSharingService() bool
	SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker ISharingServicePicker, items foundation.foundation.INSArray, proposedServices []SharingService) []SharingService
	HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool
	SharingServicePickerCollaborationModeRestrictions(sharingServicePicker ISharingServicePicker) []SharingCollaborationModeRestriction
	HasSharingServicePickerCollaborationModeRestrictions() bool
}

// SharingServicePickerDelegate is a delegate implementation builder for the PSharingServicePickerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SharingServicePickerDelegate struct {
	_SharingServicePickerDelegateForSharingService func(sharingServicePicker ISharingServicePicker, sharingService ISharingService) unsafe.Pointer
	_SharingServicePickerDidChooseSharingService func(sharingServicePicker ISharingServicePicker, service ISharingService)
	_SharingServicePickerSharingServicesForItemsProposedSharingServices func(sharingServicePicker ISharingServicePicker, items foundation.foundation.INSArray, proposedServices []SharingService) []SharingService
	_SharingServicePickerCollaborationModeRestrictions func(sharingServicePicker ISharingServicePicker) []SharingCollaborationModeRestriction
}

// SetSharingServicePickerDelegateForSharingService sets the handler for the SharingServicePickerDelegateForSharingService delegate method.
//
// Asks your delegate to provide an object that the selected sharing service can use as its delegate.
func (d *SharingServicePickerDelegate) SetSharingServicePickerDelegateForSharingService(f func(sharingServicePicker ISharingServicePicker, sharingService ISharingService) unsafe.Pointer) {
	d._SharingServicePickerDelegateForSharingService = f
}

// SetSharingServicePickerDidChooseSharingService sets the handler for the SharingServicePickerDidChooseSharingService delegate method.
//
// Tells the delegate that the person selected a sharing service for the current item.
func (d *SharingServicePickerDelegate) SetSharingServicePickerDidChooseSharingService(f func(sharingServicePicker ISharingServicePicker, service ISharingService)) {
	d._SharingServicePickerDidChooseSharingService = f
}

// SetSharingServicePickerSharingServicesForItemsProposedSharingServices sets the handler for the SharingServicePickerSharingServicesForItemsProposedSharingServices delegate method.
//
// Asks the delegate to specify which services to make available from the sharing service picker.
func (d *SharingServicePickerDelegate) SetSharingServicePickerSharingServicesForItemsProposedSharingServices(f func(sharingServicePicker ISharingServicePicker, items foundation.foundation.INSArray, proposedServices []SharingService) []SharingService) {
	d._SharingServicePickerSharingServicesForItemsProposedSharingServices = f
}

// SetSharingServicePickerCollaborationModeRestrictions sets the handler for the SharingServicePickerCollaborationModeRestrictions delegate method.
//
// Used to specify the case where the share picker should not support some modes of sharing even if they are supported by the items being shared.   Disabling all possible modes at the same time is not supported behavior.
func (d *SharingServicePickerDelegate) SetSharingServicePickerCollaborationModeRestrictions(f func(sharingServicePicker ISharingServicePicker) []SharingCollaborationModeRestriction) {
	d._SharingServicePickerCollaborationModeRestrictions = f
}

// SharingServicePickerDelegateForSharingService implements the PSharingServicePickerDelegate interface.
func (d *SharingServicePickerDelegate) SharingServicePickerDelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) unsafe.Pointer {
	if d._SharingServicePickerDelegateForSharingService != nil {
		return d._SharingServicePickerDelegateForSharingService(sharingServicePicker, sharingService)
	}
	var zero unsafe.Pointer
	return zero
}

// HasSharingServicePickerDelegateForSharingService returns true if a handler for SharingServicePickerDelegateForSharingService has been set.
func (d *SharingServicePickerDelegate) HasSharingServicePickerDelegateForSharingService() bool {
	return d._SharingServicePickerDelegateForSharingService != nil
}

// SharingServicePickerDidChooseSharingService implements the PSharingServicePickerDelegate interface.
func (d *SharingServicePickerDelegate) SharingServicePickerDidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	if d._SharingServicePickerDidChooseSharingService != nil {
		d._SharingServicePickerDidChooseSharingService(sharingServicePicker, service)
	}
}

// HasSharingServicePickerDidChooseSharingService returns true if a handler for SharingServicePickerDidChooseSharingService has been set.
func (d *SharingServicePickerDelegate) HasSharingServicePickerDidChooseSharingService() bool {
	return d._SharingServicePickerDidChooseSharingService != nil
}

// SharingServicePickerSharingServicesForItemsProposedSharingServices implements the PSharingServicePickerDelegate interface.
func (d *SharingServicePickerDelegate) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker ISharingServicePicker, items foundation.foundation.INSArray, proposedServices []SharingService) []SharingService {
	if d._SharingServicePickerSharingServicesForItemsProposedSharingServices != nil {
		return d._SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker, items, proposedServices)
	}
	var zero []SharingService
	return zero
}

// HasSharingServicePickerSharingServicesForItemsProposedSharingServices returns true if a handler for SharingServicePickerSharingServicesForItemsProposedSharingServices has been set.
func (d *SharingServicePickerDelegate) HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return d._SharingServicePickerSharingServicesForItemsProposedSharingServices != nil
}

// SharingServicePickerCollaborationModeRestrictions implements the PSharingServicePickerDelegate interface.
func (d *SharingServicePickerDelegate) SharingServicePickerCollaborationModeRestrictions(sharingServicePicker ISharingServicePicker) []SharingCollaborationModeRestriction {
	if d._SharingServicePickerCollaborationModeRestrictions != nil {
		return d._SharingServicePickerCollaborationModeRestrictions(sharingServicePicker)
	}
	var zero []SharingCollaborationModeRestriction
	return zero
}

// HasSharingServicePickerCollaborationModeRestrictions returns true if a handler for SharingServicePickerCollaborationModeRestrictions has been set.
func (d *SharingServicePickerDelegate) HasSharingServicePickerCollaborationModeRestrictions() bool {
	return d._SharingServicePickerCollaborationModeRestrictions != nil
}

// SharingServicePickerDelegateObject wraps an existing Objective-C object that conforms to the PSharingServicePickerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SharingServicePickerDelegateObject struct {
	objectivec.Object
}

// NewSharingServicePickerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSharingServicePickerDelegate protocol.
func NewSharingServicePickerDelegateObject(obj objectivec.Object) *SharingServicePickerDelegateObject {
	return &SharingServicePickerDelegateObject{obj}
}

// Make sure SharingServicePickerDelegateObject implements PSharingServicePickerDelegate.
var _ PSharingServicePickerDelegate = (*SharingServicePickerDelegateObject)(nil)

// SharingServicePickerDelegateForSharingService implements the PSharingServicePickerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServicePickerDelegateObject) SharingServicePickerDelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) unsafe.Pointer {
	return objc.Send[unsafe.Pointer](o.ID, objc.Sel("sharingServicePicker:delegateForSharingService:"), sharingServicePicker, sharingService)
}

// HasSharingServicePickerDelegateForSharingService returns true; this is a placeholder for optional method checks.
func (o *SharingServicePickerDelegateObject) HasSharingServicePickerDelegateForSharingService() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServicePickerDidChooseSharingService implements the PSharingServicePickerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServicePickerDelegateObject) SharingServicePickerDidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingServicePicker:didChooseSharingService:"), sharingServicePicker, service)
}

// HasSharingServicePickerDidChooseSharingService returns true; this is a placeholder for optional method checks.
func (o *SharingServicePickerDelegateObject) HasSharingServicePickerDidChooseSharingService() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServicePickerSharingServicesForItemsProposedSharingServices implements the PSharingServicePickerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServicePickerDelegateObject) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker ISharingServicePicker, items foundation.foundation.INSArray, proposedServices []SharingService) []SharingService {
	return objc.Send[[]SharingService](o.ID, objc.Sel("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), sharingServicePicker, items, proposedServices)
}

// HasSharingServicePickerSharingServicesForItemsProposedSharingServices returns true; this is a placeholder for optional method checks.
func (o *SharingServicePickerDelegateObject) HasSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServicePickerCollaborationModeRestrictions implements the PSharingServicePickerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServicePickerDelegateObject) SharingServicePickerCollaborationModeRestrictions(sharingServicePicker ISharingServicePicker) []SharingCollaborationModeRestriction {
	return objc.Send[[]SharingCollaborationModeRestriction](o.ID, objc.Sel("sharingServicePickerCollaborationModeRestrictions:"), sharingServicePicker)
}

// HasSharingServicePickerCollaborationModeRestrictions returns true; this is a placeholder for optional method checks.
func (o *SharingServicePickerDelegateObject) HasSharingServicePickerCollaborationModeRestrictions() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
