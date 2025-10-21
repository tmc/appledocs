// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit_test

import (
	"github.com/tmc/appledocs/generated/cryptotokenkit"
)

// Suppress unused import errors
var _ = cryptotokenkit.NewTKToken

// ExampleNewTKTokenWithTokenDriverInstanceID demonstrates how to create a TKToken instance using NewTKTokenWithTokenDriverInstanceID.
// Initializes a token with the driver you specify.
func ExampleNewTKTokenWithTokenDriverInstanceID() {
	_ = cryptotokenkit.NewTKTokenWithTokenDriverInstanceID(
		cryptotokenkit.TKTokenDriver{}, // tokenDriver TKTokenDriver
		cryptotokenkit.TKTokenInstanceID{}, // instanceID TKTokenInstanceID
	)
	// Output:
}
