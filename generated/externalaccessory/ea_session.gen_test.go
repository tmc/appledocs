// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory_test

import (
	"github.com/tmc/appledocs/generated/externalaccessory"
)

// Suppress unused import errors
var _ = externalaccessory.NewEASession

// ExampleNewEASessionWithAccessoryForProtocol demonstrates how to create a EASession instance using NewEASessionWithAccessoryForProtocol.
// Initializes the session for the specified accessory and protocol.
func ExampleNewEASessionWithAccessoryForProtocol() {
	_ = externalaccessory.NewEASessionWithAccessoryForProtocol(
		externalaccessory.EAAccessory{}, // accessory EAAccessory
		"protocolString", // protocolString string
	)
	// Output:
}
