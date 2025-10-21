// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts_test

import (
	"github.com/tmc/appledocs/generated/accounts"
)

// Suppress unused import errors
var _ = accounts.NewACAccountCredential



// ExampleNewACAccountCredentialWithOAuthTokenTokenSecret demonstrates how to create a ACAccountCredential instance using NewACAccountCredentialWithOAuthTokenTokenSecret.
// Initializes an account credential using OAuth.
func ExampleNewACAccountCredentialWithOAuthTokenTokenSecret() {
	_ = accounts.NewACAccountCredentialWithOAuthTokenTokenSecret(
		"token", // token string
		"secret", // secret string
	)
	// Output:
}


