// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts_test

import (
	"github.com/tmc/appledocs/generated/accounts"
)

// Suppress unused import errors
var _ = accounts.NewACAccount

// ExampleNewACAccountWithAccountType demonstrates how to create a ACAccount instance using NewACAccountWithAccountType.
// Initializes a new account of the specified type.
func ExampleNewACAccountWithAccountType() {
	_ = accounts.NewACAccountWithAccountType(
		accounts.ACAccountType{}, // type ACAccountType
	)
	// Output:
}
