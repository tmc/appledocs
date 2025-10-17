// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)


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

// ExampleNewButtonWithTitleImageTargetAction demonstrates how to create a Button instance using NewButtonWithTitleImageTargetAction.
// Creates a standard push button with a title and image.
func ExampleNewButtonWithTitleImageTargetAction() {
	_ = appkit.NewButtonWithTitleImageTargetAction(
		"title", // title string
		nil, // image unsafe.Pointer
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


