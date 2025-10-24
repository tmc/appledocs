// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMCAdvertiserAssistantDelegate is the MCAdvertiserAssistantDelegate protocol interface.
//
// The   protocol describes the methods that the delegate object for an   instance can implement to handle advertising-related events.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.multipeerconnectivity/documentation/MultipeerConnectivity/MCAdvertiserAssistantDelegate
type PMCAdvertiserAssistantDelegate interface {
	// Optional methods
	AdvertiserAssistantDidDismissInvitation(advertiserAssistant IMCAdvertiserAssistant)
	HasAdvertiserAssistantDidDismissInvitation() bool
	AdvertiserAssistantWillPresentInvitation(advertiserAssistant IMCAdvertiserAssistant)
	HasAdvertiserAssistantWillPresentInvitation() bool
}

// MCAdvertiserAssistantDelegate is a delegate implementation builder for the PMCAdvertiserAssistantDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MCAdvertiserAssistantDelegate struct {
	_AdvertiserAssistantDidDismissInvitation func(advertiserAssistant IMCAdvertiserAssistant)
	_AdvertiserAssistantWillPresentInvitation func(advertiserAssistant IMCAdvertiserAssistant)
}

// SetAdvertiserAssistantDidDismissInvitation sets the handler for the AdvertiserAssistantDidDismissInvitation delegate method.
//
// Indicates that the advertiser assistant finished showing the invitation to the user.
func (d *MCAdvertiserAssistantDelegate) SetAdvertiserAssistantDidDismissInvitation(f func(advertiserAssistant IMCAdvertiserAssistant)) {
	d._AdvertiserAssistantDidDismissInvitation = f
}

// SetAdvertiserAssistantWillPresentInvitation sets the handler for the AdvertiserAssistantWillPresentInvitation delegate method.
//
// Indicates that the advertiser assistant is about to present an invitation to the user.
func (d *MCAdvertiserAssistantDelegate) SetAdvertiserAssistantWillPresentInvitation(f func(advertiserAssistant IMCAdvertiserAssistant)) {
	d._AdvertiserAssistantWillPresentInvitation = f
}

// AdvertiserAssistantDidDismissInvitation implements the PMCAdvertiserAssistantDelegate interface.
func (d *MCAdvertiserAssistantDelegate) AdvertiserAssistantDidDismissInvitation(advertiserAssistant IMCAdvertiserAssistant) {
	if d._AdvertiserAssistantDidDismissInvitation != nil {
		d._AdvertiserAssistantDidDismissInvitation(advertiserAssistant)
	}
}

// HasAdvertiserAssistantDidDismissInvitation returns true if a handler for AdvertiserAssistantDidDismissInvitation has been set.
func (d *MCAdvertiserAssistantDelegate) HasAdvertiserAssistantDidDismissInvitation() bool {
	return d._AdvertiserAssistantDidDismissInvitation != nil
}

// AdvertiserAssistantWillPresentInvitation implements the PMCAdvertiserAssistantDelegate interface.
func (d *MCAdvertiserAssistantDelegate) AdvertiserAssistantWillPresentInvitation(advertiserAssistant IMCAdvertiserAssistant) {
	if d._AdvertiserAssistantWillPresentInvitation != nil {
		d._AdvertiserAssistantWillPresentInvitation(advertiserAssistant)
	}
}

// HasAdvertiserAssistantWillPresentInvitation returns true if a handler for AdvertiserAssistantWillPresentInvitation has been set.
func (d *MCAdvertiserAssistantDelegate) HasAdvertiserAssistantWillPresentInvitation() bool {
	return d._AdvertiserAssistantWillPresentInvitation != nil
}
