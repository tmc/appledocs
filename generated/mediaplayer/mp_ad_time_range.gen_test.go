// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewAdTimeRange

// ExampleNewAdTimeRangeWithTimeRange demonstrates how to create a AdTimeRange instance using NewAdTimeRangeWithTimeRange.
// Creates a Media Player time range that indicates where an ad break exists in the current player.
func ExampleNewAdTimeRangeWithTimeRange() {
	_ = mediaplayer.NewAdTimeRangeWithTimeRange(
		mediaplayer.TimeRange /* not a class type */{}, // timeRange TimeRange /* not a class type */
	)
	// Output:
}
