// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit_test

import (
	"github.com/tmc/appledocs/generated/cryptotokenkit"
)

// Suppress unused import errors
var _ = cryptotokenkit.NewTKTokenKeychainItem

// ExampleNewTKTokenKeychainItemWithObjectID demonstrates how to create a TKTokenKeychainItem instance using NewTKTokenKeychainItemWithObjectID.
// Initializes a token keychain item with the specified object ID.
func ExampleNewTKTokenKeychainItemWithObjectID() {
	_ = cryptotokenkit.NewTKTokenKeychainItemWithObjectID(
		cryptotokenkit.TKTokenObjectID /* typedef */{}, // objectID TKTokenObjectID /* typedef */
	)
	// Output:
}
