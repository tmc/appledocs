// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewWritingToolsCoordinator

// ExampleWritingToolsCoordinator_StopWritingTools demonstrates using StopWritingTools on a WritingToolsCoordinator instance.
// Stops the current Writing Tools operation and dismisses the system UI.
func ExampleWritingToolsCoordinator_StopWritingTools() {
	obj := appkit.NewWritingToolsCoordinator()
	obj.StopWritingTools()
	// Output:
	}

