// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewDocument

// ExampleNewDocument demonstrates how to create a Document instance.
// Initializes and returns an empty document object.
func ExampleNewDocument() {
	_ = appkit.NewDocument()
	// Output:
}
// ExampleDocument_Close demonstrates using Close on a Document instance.
// Closes all of the document’s windows and removes the document from its document controller.
func ExampleDocument_Close() {
	obj := appkit.NewDocument()
	obj.Close()
	// Output:
	}

// ExampleDocument_DefaultDraftName demonstrates using DefaultDraftName on a Document instance.
// Returns the default draft name for the document subclass.
func ExampleDocument_DefaultDraftName() {
	obj := appkit.NewDocument()
	_ = obj.DefaultDraftName()
	// Output:
	}

// ExampleDocument_InvalidateRestorableState demonstrates using InvalidateRestorableState on a Document instance.
// Marks the document’s interface-related state as dirty.
func ExampleDocument_InvalidateRestorableState() {
	obj := appkit.NewDocument()
	obj.InvalidateRestorableState()
	// Output:
	}

// ExampleDocument_MakeWindowControllers demonstrates using MakeWindowControllers on a Document instance.
// Creates the window controller objects that the document uses to display its content.
func ExampleDocument_MakeWindowControllers() {
	obj := appkit.NewDocument()
	obj.MakeWindowControllers()
	// Output:
	}

// ExampleDocument_PresentedItemDidChange demonstrates using PresentedItemDidChange on a Document instance.
func ExampleDocument_PresentedItemDidChange() {
	obj := appkit.NewDocument()
	obj.PresentedItemDidChange()
	// Output:
	}

// ExampleDocument_ScheduleAutosaving demonstrates using ScheduleAutosaving on a Document instance.
// Schedules periodic autosaving for the purpose of crash protection.
func ExampleDocument_ScheduleAutosaving() {
	obj := appkit.NewDocument()
	obj.ScheduleAutosaving()
	// Output:
	}

// ExampleDocument_ShowWindows demonstrates using ShowWindows on a Document instance.
// Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
func ExampleDocument_ShowWindows() {
	obj := appkit.NewDocument()
	obj.ShowWindows()
	// Output:
	}

// ExampleDocument_UnblockUserInteraction demonstrates using UnblockUserInteraction on a Document instance.
// Unblocks the main thread during asynchronous saving.
func ExampleDocument_UnblockUserInteraction() {
	obj := appkit.NewDocument()
	obj.UnblockUserInteraction()
	// Output:
	}

