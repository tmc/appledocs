// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLHandle

// ExampleNewURLHandleWithURLCached demonstrates how to create a URLHandle instance using NewURLHandleWithURLCached.
// Initializes a newly created URL handle with the specified URL.
func ExampleNewURLHandleWithURLCached() {
	_ = foundation.NewURLHandleWithURLCached(
		foundation.URL{}, // anURL URL
		false, // willCache bool
	)
	// Output:
}
