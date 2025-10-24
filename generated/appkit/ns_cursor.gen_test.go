// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCursor

// ExampleCursor_Pop demonstrates using Pop on a Cursor instance.
// Sends a   message to the receiver’s class.
func ExampleCursor_Pop() {
	obj := appkit.NewCursor()
	obj.Pop()
	// Output:
	}

// ExampleCursor_Push demonstrates using Push on a Cursor instance.
// Puts the receiver on top of the cursor stack and makes it the current cursor.
func ExampleCursor_Push() {
	obj := appkit.NewCursor()
	obj.Push()
	// Output:
	}

// ExampleCursor_Set demonstrates using Set on a Cursor instance.
// Makes the receiver the current cursor.
func ExampleCursor_Set() {
	obj := appkit.NewCursor()
	obj.Set()
	// Output:
	}

