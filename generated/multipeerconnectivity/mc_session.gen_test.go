// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCSession

// ExampleNewMCSessionWithPeer demonstrates how to create a MCSession instance using NewMCSessionWithPeer.
// Creates a Multipeer Connectivity session.
func ExampleNewMCSessionWithPeer() {
	_ = multipeerconnectivity.NewMCSessionWithPeer(
		multipeerconnectivity.MCPeerID{}, // myPeerID MCPeerID
	)
	// Output:
}

