// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCNearbyServiceAdvertiser

// ExampleMCNearbyServiceAdvertiser_StartAdvertisingPeer demonstrates using StartAdvertisingPeer on a MCNearbyServiceAdvertiser instance.
// Begins advertising the service provided by a local peer.
func ExampleMCNearbyServiceAdvertiser_StartAdvertisingPeer() {
	obj := multipeerconnectivity.NewMCNearbyServiceAdvertiser()
	obj.StartAdvertisingPeer()
	// Output:
	}

// ExampleMCNearbyServiceAdvertiser_StopAdvertisingPeer demonstrates using StopAdvertisingPeer on a MCNearbyServiceAdvertiser instance.
// Stops advertising the service provided by a local peer.
func ExampleMCNearbyServiceAdvertiser_StopAdvertisingPeer() {
	obj := multipeerconnectivity.NewMCNearbyServiceAdvertiser()
	obj.StopAdvertisingPeer()
	// Output:
	}

