// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSSearchQuery

// ExampleNewCSSearchQueryWithQueryStringAttributes demonstrates how to create a CSSearchQuery instance using NewCSSearchQueryWithQueryStringAttributes.
// Initializes and returns a query object with the specified query string and item attributes.
func ExampleNewCSSearchQueryWithQueryStringAttributes() {
	_ = corespotlight.NewCSSearchQueryWithQueryStringAttributes(
		"kMDItemFSName == '*.txt'", // queryString string
		[]corespotlight.string{}, // attributes []string
	)
	// Output:
}
// ExampleNewCSSearchQueryWithQueryStringQueryContext demonstrates how to create a CSSearchQuery instance using NewCSSearchQueryWithQueryStringQueryContext.
// Initializes and returns a query object with the specified query string and query context.
func ExampleNewCSSearchQueryWithQueryStringQueryContext() {
	_ = corespotlight.NewCSSearchQueryWithQueryStringQueryContext(
		"kMDItemFSName == '*.txt'", // queryString string
		corespotlight.CSSearchQueryContext{}, // queryContext CSSearchQueryContext
	)
	// Output:
}
