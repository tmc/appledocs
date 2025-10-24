// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMCNearbyServiceBrowserDelegate is the MCNearbyServiceBrowserDelegate protocol interface.
//
// The   protocol defines methods that a   object’s delegate can implement to handle browser-related events.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.multipeerconnectivity/documentation/MultipeerConnectivity/MCNearbyServiceBrowserDelegate
type PMCNearbyServiceBrowserDelegate interface {
	// Required methods
	BrowserFoundPeerWithDiscoveryInfo(browser IMCNearbyServiceBrowser, peerID IMCPeerID, info foundation.IDictionary)/* debug [protocol_interface/required_method]: BrowserFoundPeerWithDiscoveryInfo */
	BrowserLostPeer(browser IMCNearbyServiceBrowser, peerID IMCPeerID)/* debug [protocol_interface/required_method]: BrowserLostPeer */
	// Optional methods
	BrowserDidNotStartBrowsingForPeers(browser IMCNearbyServiceBrowser, error_ objc.IObject /* cross-framework: Error */)
	HasBrowserDidNotStartBrowsingForPeers() bool
}

// MCNearbyServiceBrowserDelegate is a delegate implementation builder for the PMCNearbyServiceBrowserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MCNearbyServiceBrowserDelegate struct {
	_BrowserDidNotStartBrowsingForPeers func(browser IMCNearbyServiceBrowser, error_ objc.IObject /* cross-framework: Error */)
	_BrowserFoundPeerWithDiscoveryInfo func(browser IMCNearbyServiceBrowser, peerID IMCPeerID, info foundation.IDictionary)
	_BrowserLostPeer func(browser IMCNearbyServiceBrowser, peerID IMCPeerID)
}

// SetBrowserDidNotStartBrowsingForPeers sets the handler for the BrowserDidNotStartBrowsingForPeers delegate method.
//
// Called when a browser failed to start browsing for peers.
func (d *MCNearbyServiceBrowserDelegate) SetBrowserDidNotStartBrowsingForPeers(f func(browser IMCNearbyServiceBrowser, error_ objc.IObject /* cross-framework: Error */)) {
	d._BrowserDidNotStartBrowsingForPeers = f
}

// SetBrowserFoundPeerWithDiscoveryInfo sets the handler for the BrowserFoundPeerWithDiscoveryInfo delegate method.
//
// Called when a nearby peer is found.
func (d *MCNearbyServiceBrowserDelegate) SetBrowserFoundPeerWithDiscoveryInfo(f func(browser IMCNearbyServiceBrowser, peerID IMCPeerID, info foundation.IDictionary)) {
	d._BrowserFoundPeerWithDiscoveryInfo = f
}

// SetBrowserLostPeer sets the handler for the BrowserLostPeer delegate method.
//
// Called when a nearby peer is lost.
func (d *MCNearbyServiceBrowserDelegate) SetBrowserLostPeer(f func(browser IMCNearbyServiceBrowser, peerID IMCPeerID)) {
	d._BrowserLostPeer = f
}

// BrowserDidNotStartBrowsingForPeers implements the PMCNearbyServiceBrowserDelegate interface.
func (d *MCNearbyServiceBrowserDelegate) BrowserDidNotStartBrowsingForPeers(browser IMCNearbyServiceBrowser, error_ objc.IObject /* cross-framework: Error */) {
	if d._BrowserDidNotStartBrowsingForPeers != nil {
		d._BrowserDidNotStartBrowsingForPeers(browser, error_)
	}
}

// HasBrowserDidNotStartBrowsingForPeers returns true if a handler for BrowserDidNotStartBrowsingForPeers has been set.
func (d *MCNearbyServiceBrowserDelegate) HasBrowserDidNotStartBrowsingForPeers() bool {
	return d._BrowserDidNotStartBrowsingForPeers != nil
}

// BrowserFoundPeerWithDiscoveryInfo implements the PMCNearbyServiceBrowserDelegate interface.
func (d *MCNearbyServiceBrowserDelegate) BrowserFoundPeerWithDiscoveryInfo(browser IMCNearbyServiceBrowser, peerID IMCPeerID, info foundation.IDictionary) {
	if d._BrowserFoundPeerWithDiscoveryInfo != nil {
		d._BrowserFoundPeerWithDiscoveryInfo(browser, peerID, info)
	}
}

// HasBrowserFoundPeerWithDiscoveryInfo returns true if a handler for BrowserFoundPeerWithDiscoveryInfo has been set.
func (d *MCNearbyServiceBrowserDelegate) HasBrowserFoundPeerWithDiscoveryInfo() bool {
	return d._BrowserFoundPeerWithDiscoveryInfo != nil
}

// BrowserLostPeer implements the PMCNearbyServiceBrowserDelegate interface.
func (d *MCNearbyServiceBrowserDelegate) BrowserLostPeer(browser IMCNearbyServiceBrowser, peerID IMCPeerID) {
	if d._BrowserLostPeer != nil {
		d._BrowserLostPeer(browser, peerID)
	}
}

// HasBrowserLostPeer returns true if a handler for BrowserLostPeer has been set.
func (d *MCNearbyServiceBrowserDelegate) HasBrowserLostPeer() bool {
	return d._BrowserLostPeer != nil
}
