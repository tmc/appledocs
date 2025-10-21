// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewUserDefaults


// ExampleNewUserDefaults demonstrates how to create a UserDefaults instance.
// Creates a user defaults object initialized with the defaults for the app and current user.
func ExampleNewUserDefaults() {
	_ = foundation.NewUserDefaults()
	// Output:
}


// ExampleNewUserDefaultsWithSuiteName demonstrates how to create a UserDefaults instance using NewUserDefaultsWithSuiteName.
// Creates a user defaults object initialized with the defaults for the specified database name.
func ExampleNewUserDefaultsWithSuiteName() {
	_ = foundation.NewUserDefaultsWithSuiteName(
		"suitename", // suitename string
	)
	// Output:
}


