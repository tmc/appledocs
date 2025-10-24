// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMCNearbyServiceAdvertiserDelegate is the MCNearbyServiceAdvertiserDelegate protocol interface.
//
// The   protocol describes the methods that the delegate object for an   instance can implement for handling events from the   class.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.multipeerconnectivity/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiserDelegate
type PMCNearbyServiceAdvertiserDelegate interface {
	// Required methods
	AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler(advertiser IMCNearbyServiceAdvertiser, peerID IMCPeerID, context objc.IObject /* cross-framework: NSData */, invitationHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler */
	// Optional methods
	AdvertiserDidNotStartAdvertisingPeer(advertiser IMCNearbyServiceAdvertiser, error_ objc.IObject /* cross-framework: Error */)
	HasAdvertiserDidNotStartAdvertisingPeer() bool
}

// MCNearbyServiceAdvertiserDelegate is a delegate implementation builder for the PMCNearbyServiceAdvertiserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MCNearbyServiceAdvertiserDelegate struct {
	_AdvertiserDidNotStartAdvertisingPeer func(advertiser IMCNearbyServiceAdvertiser, error_ objc.IObject /* cross-framework: Error */)
	_AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler func(advertiser IMCNearbyServiceAdvertiser, peerID IMCPeerID, context objc.IObject /* cross-framework: NSData */, invitationHandler unsafe.Pointer)
}

// SetAdvertiserDidNotStartAdvertisingPeer sets the handler for the AdvertiserDidNotStartAdvertisingPeer delegate method.
//
// Called when advertisement fails.
func (d *MCNearbyServiceAdvertiserDelegate) SetAdvertiserDidNotStartAdvertisingPeer(f func(advertiser IMCNearbyServiceAdvertiser, error_ objc.IObject /* cross-framework: Error */)) {
	d._AdvertiserDidNotStartAdvertisingPeer = f
}

// SetAdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler sets the handler for the AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler delegate method.
//
// Called when an invitation to join a session is received from a nearby peer.
func (d *MCNearbyServiceAdvertiserDelegate) SetAdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler(f func(advertiser IMCNearbyServiceAdvertiser, peerID IMCPeerID, context objc.IObject /* cross-framework: NSData */, invitationHandler unsafe.Pointer)) {
	d._AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler = f
}

// AdvertiserDidNotStartAdvertisingPeer implements the PMCNearbyServiceAdvertiserDelegate interface.
func (d *MCNearbyServiceAdvertiserDelegate) AdvertiserDidNotStartAdvertisingPeer(advertiser IMCNearbyServiceAdvertiser, error_ objc.IObject /* cross-framework: Error */) {
	if d._AdvertiserDidNotStartAdvertisingPeer != nil {
		d._AdvertiserDidNotStartAdvertisingPeer(advertiser, error_)
	}
}

// HasAdvertiserDidNotStartAdvertisingPeer returns true if a handler for AdvertiserDidNotStartAdvertisingPeer has been set.
func (d *MCNearbyServiceAdvertiserDelegate) HasAdvertiserDidNotStartAdvertisingPeer() bool {
	return d._AdvertiserDidNotStartAdvertisingPeer != nil
}

// AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler implements the PMCNearbyServiceAdvertiserDelegate interface.
func (d *MCNearbyServiceAdvertiserDelegate) AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler(advertiser IMCNearbyServiceAdvertiser, peerID IMCPeerID, context objc.IObject /* cross-framework: NSData */, invitationHandler unsafe.Pointer) {
	if d._AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler != nil {
		d._AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler(advertiser, peerID, context, invitationHandler)
	}
}

// HasAdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler returns true if a handler for AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler has been set.
func (d *MCNearbyServiceAdvertiserDelegate) HasAdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler() bool {
	return d._AdvertiserDidReceiveInvitationFromPeerWithContextInvitationHandler != nil
}
