// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPopUpButton

// ExamplePopUpButton_RemoveAllItems demonstrates using RemoveAllItems on a PopUpButton instance.
// Removes all items in the receiver’s item menu.
func ExamplePopUpButton_RemoveAllItems() {
	obj := appkit.NewPopUpButton()
	obj.RemoveAllItems()
	// Output:
	}

// ExamplePopUpButton_SynchronizeTitleAndSelectedItem demonstrates using SynchronizeTitleAndSelectedItem on a PopUpButton instance.
// Ensures that the item being displayed by the receiver agrees with the selected item.
func ExamplePopUpButton_SynchronizeTitleAndSelectedItem() {
	obj := appkit.NewPopUpButton()
	obj.SynchronizeTitleAndSelectedItem()
	// Output:
	}

