// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewMergePolicy

// ExampleNewMergePolicyWithMergeType demonstrates how to create a MergePolicy instance using NewMergePolicyWithMergeType.
// Returns a merge policy initialized with a given policy type.
func ExampleNewMergePolicyWithMergeType() {
	_ = coredata.NewMergePolicyWithMergeType(
		coredata.MergePolicyType{}, // ty MergePolicyType
	)
	// Output:
}
