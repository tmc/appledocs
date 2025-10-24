// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCSession

// ExampleMCSession_Disconnect demonstrates using Disconnect on a MCSession instance.
// Disconnects the local peer from the session.
func ExampleMCSession_Disconnect() {
	obj := multipeerconnectivity.NewMCSession()
	obj.Disconnect()
	// Output:
	}


