// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextSelectionNavigation

// ExampleTextSelectionNavigation_FlushLayoutCache demonstrates using FlushLayoutCache on a TextSelectionNavigation instance.
// Flushes cached layout information.
func ExampleTextSelectionNavigation_FlushLayoutCache() {
	obj := appkit.NewTextSelectionNavigation()
	obj.FlushLayoutCache()
	// Output:
	}

