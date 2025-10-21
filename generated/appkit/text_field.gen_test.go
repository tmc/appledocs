// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextField

// ExampleNewTextFieldLabelWithAttributedString demonstrates how to create a TextField instance using NewTextFieldLabelWithAttributedString.
// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
func ExampleNewTextFieldLabelWithAttributedString() {
	_ = appkit.NewTextFieldLabelWithAttributedString(
		appkit.AttributedString{}, // attributedStringValue AttributedString
	)
	// Output:
}
