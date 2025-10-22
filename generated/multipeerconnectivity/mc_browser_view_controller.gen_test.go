// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCBrowserViewController

// ExampleNewMCBrowserViewControllerWithBrowserSession demonstrates how to create a MCBrowserViewController instance using NewMCBrowserViewControllerWithBrowserSession.
// Initializes a browser view controller with the provided browser and session.
func ExampleNewMCBrowserViewControllerWithBrowserSession() {
	_ = multipeerconnectivity.NewMCBrowserViewControllerWithBrowserSession(
		multipeerconnectivity.MCNearbyServiceBrowser{}, // browser MCNearbyServiceBrowser
		multipeerconnectivity.MCSession{}, // session MCSession
	)
	// Output:
}
// ExampleNewMCBrowserViewControllerWithServiceTypeSession demonstrates how to create a MCBrowserViewController instance using NewMCBrowserViewControllerWithServiceTypeSession.
// Initializes a browser view controller using the provided service type and session.
func ExampleNewMCBrowserViewControllerWithServiceTypeSession() {
	_ = multipeerconnectivity.NewMCBrowserViewControllerWithServiceTypeSession(
		"serviceType", // serviceType string
		multipeerconnectivity.MCSession{}, // session MCSession
	)
	// Output:
}
