// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCNearbyServiceBrowser

// ExampleNewMCNearbyServiceBrowserWithPeerServiceType demonstrates how to create a MCNearbyServiceBrowser instance using NewMCNearbyServiceBrowserWithPeerServiceType.
// Initializes the nearby service browser object.
func ExampleNewMCNearbyServiceBrowserWithPeerServiceType() {
	_ = multipeerconnectivity.NewMCNearbyServiceBrowserWithPeerServiceType(
		multipeerconnectivity.MCPeerID{}, // myPeerID MCPeerID
		"serviceType", // serviceType string
	)
	// Output:
}
