// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPopUpButtonCell

// ExamplePopUpButtonCell_DismissPopUp demonstrates using DismissPopUp on a PopUpButtonCell instance.
// Dismisses the pop-up button’s menu by ordering its window out.
func ExamplePopUpButtonCell_DismissPopUp() {
	obj := appkit.NewPopUpButtonCell()
	obj.DismissPopUp()
	// Output:
	}

// ExamplePopUpButtonCell_RemoveAllItems demonstrates using RemoveAllItems on a PopUpButtonCell instance.
// Removes all items in the receiver’s item menu.
func ExamplePopUpButtonCell_RemoveAllItems() {
	obj := appkit.NewPopUpButtonCell()
	obj.RemoveAllItems()
	// Output:
	}

// ExamplePopUpButtonCell_SynchronizeTitleAndSelectedItem demonstrates using SynchronizeTitleAndSelectedItem on a PopUpButtonCell instance.
// Synchronizes the pop-up button’s displayed item with the currently selected menu item.
func ExamplePopUpButtonCell_SynchronizeTitleAndSelectedItem() {
	obj := appkit.NewPopUpButtonCell()
	obj.SynchronizeTitleAndSelectedItem()
	// Output:
	}

