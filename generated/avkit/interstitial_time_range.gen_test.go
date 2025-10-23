// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit_test

import (
	"github.com/tmc/appledocs/generated/avkit"
)

// Suppress unused import errors
var _ = avkit.NewInterstitialTimeRange

// ExampleNewInterstitialTimeRangeWithTimeRange demonstrates how to create a InterstitialTimeRange instance using NewInterstitialTimeRangeWithTimeRange.
// Initializes an interstitial time range object with the specified time range.
func ExampleNewInterstitialTimeRangeWithTimeRange() {
	_ = avkit.NewInterstitialTimeRangeWithTimeRange(
		avkit.TimeRange{}, // timeRange TimeRange
	)
	// Output:
}
