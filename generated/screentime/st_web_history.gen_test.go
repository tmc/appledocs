// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime_test

import (
	"github.com/tmc/appledocs/generated/screentime"
)

// Suppress unused import errors
var _ = screentime.NewSTWebHistory

// ExampleNewSTWebHistoryWithProfileIdentifier demonstrates how to create a STWebHistory instance using NewSTWebHistoryWithProfileIdentifier.
// Creates a web history instance to delete web-usage data associated to the   profile identifier you specify.
func ExampleNewSTWebHistoryWithProfileIdentifier() {
	_ = screentime.NewSTWebHistoryWithProfileIdentifier(
		screentime.STWebHistoryProfileIdentifier /* typedef */{}, // profileIdentifier STWebHistoryProfileIdentifier /* typedef */
	)
	// Output:
}
