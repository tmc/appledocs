// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan_test

import (
	"github.com/tmc/appledocs/generated/corewlan"
)

// Suppress unused import errors
var _ = corewlan.NewCWWiFiClient

// ExampleNewCWWiFiClient demonstrates how to create a CWWiFiClient instance.
// Initializes a Wi-Fi client object.
func ExampleNewCWWiFiClient() {
	_ = corewlan.NewCWWiFiClient()
	// Output:
}
// ExampleCWWiFiClient_Interface demonstrates using Interface on a CWWiFiClient instance.
// Returns the default Wi-Fi interface.
func ExampleCWWiFiClient_Interface() {
	obj := corewlan.NewCWWiFiClient()
	_ = obj.Interface()
	// Output:
	}

// ExampleCWWiFiClient_InterfaceNames demonstrates using InterfaceNames on a CWWiFiClient instance.
func ExampleCWWiFiClient_InterfaceNames() {
	obj := corewlan.NewCWWiFiClient()
	_ = obj.InterfaceNames()
	// Output:
	}

// ExampleCWWiFiClient_Interfaces demonstrates using Interfaces on a CWWiFiClient instance.
// Returns all available Wi-Fi interfaces.
func ExampleCWWiFiClient_Interfaces() {
	obj := corewlan.NewCWWiFiClient()
	_ = obj.Interfaces()
	// Output:
	}


