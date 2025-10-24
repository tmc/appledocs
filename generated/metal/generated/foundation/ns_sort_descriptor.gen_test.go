// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSortDescriptor

// ExampleSortDescriptor_AllowEvaluation demonstrates using AllowEvaluation on a SortDescriptor instance.
// Forces a securely decoded sort descriptor to allow evaluation.
func ExampleSortDescriptor_AllowEvaluation() {
	obj := foundation.NewSortDescriptor()
	obj.AllowEvaluation()
	// Output:
	}

