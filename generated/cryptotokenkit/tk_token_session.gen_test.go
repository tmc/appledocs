// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit_test

import (
	"github.com/tmc/appledocs/generated/cryptotokenkit"
)

// Suppress unused import errors
var _ = cryptotokenkit.NewTKTokenSession

// ExampleNewTKTokenSessionWithToken demonstrates how to create a TKTokenSession instance using NewTKTokenSessionWithToken.
// Initializes a token session with the specified token.
func ExampleNewTKTokenSessionWithToken() {
	_ = cryptotokenkit.NewTKTokenSessionWithToken(
		cryptotokenkit.TKToken{}, // token TKToken
	)
	// Output:
}
