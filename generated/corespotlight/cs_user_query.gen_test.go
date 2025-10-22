// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSUserQuery

// ExampleNewCSUserQueryWithUserQueryStringUserQueryContext demonstrates how to create a CSUserQuery instance using NewCSUserQueryWithUserQueryStringUserQueryContext.
// Creates a new user query that searches for the specified term.
func ExampleNewCSUserQueryWithUserQueryStringUserQueryContext() {
	_ = corespotlight.NewCSUserQueryWithUserQueryStringUserQueryContext(
		"kMDItemFSName == '*.txt'", // userQueryString string
		corespotlight.CSUserQueryContext{}, // userQueryContext CSUserQueryContext
	)
	// Output:
}
