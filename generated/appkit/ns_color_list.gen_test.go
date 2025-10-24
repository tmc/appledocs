// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColorList

// ExampleColorList_RemoveFile demonstrates using RemoveFile on a ColorList instance.
// Removes the file from which the list was created, if the file is in a standard search path and owned by the user.
func ExampleColorList_RemoveFile() {
	obj := appkit.NewColorList()
	obj.RemoveFile()
	// Output:
}
