// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewButton


// ExampleNewButtonWithTitleTargetAction demonstrates how to create a Button instance using NewButtonWithTitleTargetAction.
// Creates a standard push button with the title you specify.
func ExampleNewButtonWithTitleTargetAction() {
	_ = appkit.NewButtonWithTitleTargetAction(
		"title", // title string
		0, // target objc.ID
		0, // action objc.SEL
	)
	// Output:
}

// ExampleNewCheckboxWithTitleTargetAction demonstrates how to create a Button instance using NewCheckboxWithTitleTargetAction.
// Creates a standard checkbox with the title you specify.
func ExampleNewCheckboxWithTitleTargetAction() {
	_ = appkit.NewCheckboxWithTitleTargetAction(
		"title", // title string
		0, // target objc.ID
		0, // action objc.SEL
	)
	// Output:
}



