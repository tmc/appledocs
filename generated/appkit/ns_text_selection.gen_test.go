// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextSelection

// ExampleNewTextSelectionWithRangesAffinityGranularity demonstrates how to create a TextSelection instance using NewTextSelectionWithRangesAffinityGranularity.
// Creates a new text selection with the ranges, selection affinity, and granularity you provide.
func ExampleNewTextSelectionWithRangesAffinityGranularity() {
	_ = appkit.NewTextSelectionWithRangesAffinityGranularity(
		[]appkit.TextRange{}, // textRanges []TextRange
		appkit.TextSelectionAffinity{}, // affinity TextSelectionAffinity
		appkit.TextSelectionGranularity{}, // granularity TextSelectionGranularity
	)
	// Output:
}
