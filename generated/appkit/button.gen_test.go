// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewButton


// ExampleNewButtonCheckboxWithTitleTargetAction demonstrates how to create a Button instance using NewButtonCheckboxWithTitleTargetAction.
// Creates a standard checkbox with the title you specify.
func ExampleNewButtonCheckboxWithTitleTargetAction() {
	_ = appkit.NewButtonCheckboxWithTitleTargetAction(
		"title", // title string
		0, // target objc.ID
		0, // action objc.SEL
	)
	// Output:
}


// ExampleNewButtonRadioButtonWithTitleTargetAction demonstrates how to create a Button instance using NewButtonRadioButtonWithTitleTargetAction.
// Creates a standard radio button with the title you specify.
func ExampleNewButtonRadioButtonWithTitleTargetAction() {
	_ = appkit.NewButtonRadioButtonWithTitleTargetAction(
		"title", // title string
		0, // target objc.ID
		0, // action objc.SEL
	)
	// Output:
}


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


