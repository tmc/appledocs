// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewLogicalTest

// ExampleNewLogicalTestAndTestWithTests demonstrates how to create a LogicalTest instance using NewLogicalTestAndTestWithTests.
// Returns an   object initialized to perform an   operation with the   objects in a given array.
func ExampleNewLogicalTestAndTestWithTests() {
	_ = foundation.NewLogicalTestAndTestWithTests(
		[]foundation.SpecifierTest{}, // subTests []SpecifierTest
	)
	// Output:
}
// ExampleNewLogicalTestOrTestWithTests demonstrates how to create a LogicalTest instance using NewLogicalTestOrTestWithTests.
// Returns an   object initialized to perform an   operation with the   objects in a given array.
func ExampleNewLogicalTestOrTestWithTests() {
	_ = foundation.NewLogicalTestOrTestWithTests(
		[]foundation.SpecifierTest{}, // subTests []SpecifierTest
	)
	// Output:
}
