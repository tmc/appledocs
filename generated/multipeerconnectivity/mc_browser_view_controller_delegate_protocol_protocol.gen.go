// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMCBrowserViewControllerDelegate is the MCBrowserViewControllerDelegate protocol interface.
//
// The   protocol defines the methods that your delegate object can implement to handle events related to the   class.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.multipeerconnectivity/documentation/MultipeerConnectivity/MCBrowserViewControllerDelegate
type PMCBrowserViewControllerDelegate interface {
	// Required methods
	BrowserViewControllerDidFinish(browserViewController IMCBrowserViewController)/* debug [protocol_interface/required_method]: BrowserViewControllerDidFinish */
	BrowserViewControllerWasCancelled(browserViewController IMCBrowserViewController)/* debug [protocol_interface/required_method]: BrowserViewControllerWasCancelled */
	// Optional methods
	BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo(browserViewController IMCBrowserViewController, peerID IMCPeerID, info foundation.IDictionary) bool
	HasBrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo() bool
}

// MCBrowserViewControllerDelegate is a delegate implementation builder for the PMCBrowserViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MCBrowserViewControllerDelegate struct {
	_BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo func(browserViewController IMCBrowserViewController, peerID IMCPeerID, info foundation.IDictionary) bool
	_BrowserViewControllerDidFinish func(browserViewController IMCBrowserViewController)
	_BrowserViewControllerWasCancelled func(browserViewController IMCBrowserViewController)
}

// SetBrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo sets the handler for the BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo delegate method.
//
// Called when a new peer is discovered to decide whether to show it in the user interface.
func (d *MCBrowserViewControllerDelegate) SetBrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo(f func(browserViewController IMCBrowserViewController, peerID IMCPeerID, info foundation.IDictionary) bool) {
	d._BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo = f
}

// SetBrowserViewControllerDidFinish sets the handler for the BrowserViewControllerDidFinish delegate method.
//
// Called when the browser view controller is dismissed with peers connected in a session.
func (d *MCBrowserViewControllerDelegate) SetBrowserViewControllerDidFinish(f func(browserViewController IMCBrowserViewController)) {
	d._BrowserViewControllerDidFinish = f
}

// SetBrowserViewControllerWasCancelled sets the handler for the BrowserViewControllerWasCancelled delegate method.
//
// Called when the user cancels the browser view controller.
func (d *MCBrowserViewControllerDelegate) SetBrowserViewControllerWasCancelled(f func(browserViewController IMCBrowserViewController)) {
	d._BrowserViewControllerWasCancelled = f
}

// BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo implements the PMCBrowserViewControllerDelegate interface.
func (d *MCBrowserViewControllerDelegate) BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo(browserViewController IMCBrowserViewController, peerID IMCPeerID, info foundation.IDictionary) bool {
	if d._BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo != nil {
		return d._BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo(browserViewController, peerID, info)
	}
	var zero bool
	return zero
}

// HasBrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo returns true if a handler for BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo has been set.
func (d *MCBrowserViewControllerDelegate) HasBrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo() bool {
	return d._BrowserViewControllerShouldPresentNearbyPeerWithDiscoveryInfo != nil
}

// BrowserViewControllerDidFinish implements the PMCBrowserViewControllerDelegate interface.
func (d *MCBrowserViewControllerDelegate) BrowserViewControllerDidFinish(browserViewController IMCBrowserViewController) {
	if d._BrowserViewControllerDidFinish != nil {
		d._BrowserViewControllerDidFinish(browserViewController)
	}
}

// HasBrowserViewControllerDidFinish returns true if a handler for BrowserViewControllerDidFinish has been set.
func (d *MCBrowserViewControllerDelegate) HasBrowserViewControllerDidFinish() bool {
	return d._BrowserViewControllerDidFinish != nil
}

// BrowserViewControllerWasCancelled implements the PMCBrowserViewControllerDelegate interface.
func (d *MCBrowserViewControllerDelegate) BrowserViewControllerWasCancelled(browserViewController IMCBrowserViewController) {
	if d._BrowserViewControllerWasCancelled != nil {
		d._BrowserViewControllerWasCancelled(browserViewController)
	}
}

// HasBrowserViewControllerWasCancelled returns true if a handler for BrowserViewControllerWasCancelled has been set.
func (d *MCBrowserViewControllerDelegate) HasBrowserViewControllerWasCancelled() bool {
	return d._BrowserViewControllerWasCancelled != nil
}
