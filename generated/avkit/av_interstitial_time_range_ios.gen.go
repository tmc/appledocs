//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for InterstitialTimeRange


// iOS-only properties

// The time range identified as interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange/timeRange
func (i_ InterstitialTimeRange) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](i_.ID, objc.Sel("timeRange"))
	return rv
}




