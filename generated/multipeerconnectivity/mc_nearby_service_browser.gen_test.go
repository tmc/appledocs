// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCNearbyServiceBrowser

// ExampleMCNearbyServiceBrowser_StartBrowsingForPeers demonstrates using StartBrowsingForPeers on a MCNearbyServiceBrowser instance.
// Starts browsing for peers.
func ExampleMCNearbyServiceBrowser_StartBrowsingForPeers() {
	obj := multipeerconnectivity.NewMCNearbyServiceBrowser()
	obj.StartBrowsingForPeers()
	// Output:
	}

// ExampleMCNearbyServiceBrowser_StopBrowsingForPeers demonstrates using StopBrowsingForPeers on a MCNearbyServiceBrowser instance.
// Stops browsing for peers.
func ExampleMCNearbyServiceBrowser_StopBrowsingForPeers() {
	obj := multipeerconnectivity.NewMCNearbyServiceBrowser()
	obj.StopBrowsingForPeers()
	// Output:
	}

