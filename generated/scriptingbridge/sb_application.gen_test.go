// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge_test

import (
	"github.com/tmc/appledocs/generated/scriptingbridge"
)

// Suppress unused import errors
var _ = scriptingbridge.NewSBApplication

// ExampleNewSBApplicationWithBundleIdentifier demonstrates how to create a SBApplication instance using NewSBApplicationWithBundleIdentifier.
// Returns an instance of an   subclass that represents the   target application identified by the given bundle identifier.
func ExampleNewSBApplicationWithBundleIdentifier() {
	_ = scriptingbridge.NewSBApplicationWithBundleIdentifier(
		"ident", // ident string
	)
	// Output:
}
