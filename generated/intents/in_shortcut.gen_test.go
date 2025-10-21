// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINShortcut

// ExampleNewINShortcutWithIntent demonstrates how to create a INShortcut instance using NewINShortcutWithIntent.
// Creates a shortcut with the specified intent.
func ExampleNewINShortcutWithIntent() {
	_ = intents.NewINShortcutWithIntent(
		intents.INIntent{}, // intent INIntent
	)
	// Output:
}
