// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest_test

import (
	"github.com/tmc/appledocs/generated/storekittest"
)

// Suppress unused import errors
var _ = storekittest.NewTestSession

// ExampleTestSession_AllTransactions demonstrates using AllTransactions on a TestSession instance.
// Gets a list of all transactions in the test environment.
func ExampleTestSession_AllTransactions() {
	obj := storekittest.NewTestSession()
	_ = obj.AllTransactions()
	// Output:
}

// ExampleTestSession_ClearTransactions demonstrates using ClearTransactions on a TestSession instance.
// Removes all transactions from the test environment.
func ExampleTestSession_ClearTransactions() {
	obj := storekittest.NewTestSession()
	obj.ClearTransactions()
	// Output:
}

// ExampleTestSession_ResetToDefaultState demonstrates using ResetToDefaultState on a TestSession instance.
// Removes all property overrides and resets all test session settings to their default state.
func ExampleTestSession_ResetToDefaultState() {
	obj := storekittest.NewTestSession()
	obj.ResetToDefaultState()
	// Output:
}
