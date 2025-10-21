// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXTransaction

// ExampleNewCXTransactionWithAction demonstrates how to create a CXTransaction instance using NewCXTransactionWithAction.
// Initializes a new transaction with the specified action.
func ExampleNewCXTransactionWithAction() {
	_ = callkit.NewCXTransactionWithAction(
		callkit.CXAction{}, // action CXAction
	)
	// Output:
}
// ExampleNewCXTransactionWithActions demonstrates how to create a CXTransaction instance using NewCXTransactionWithActions.
// Initializes a new transaction with the specified actions.
func ExampleNewCXTransactionWithActions() {
	_ = callkit.NewCXTransactionWithActions(
		[]callkit.CXAction{}, // actions []CXAction
	)
	// Output:
}

