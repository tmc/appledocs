// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLCredential

// ExampleNewURLCredentialWithUserPasswordPersistence demonstrates how to create a URLCredential instance using NewURLCredentialWithUserPasswordPersistence.
// Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
func ExampleNewURLCredentialWithUserPasswordPersistence() {
	_ = foundation.NewURLCredentialWithUserPasswordPersistence(
		"user", // user string
		"password", // password string
		foundation.URLCredentialPersistence{}, // persistence URLCredentialPersistence
	)
	// Output:
}
