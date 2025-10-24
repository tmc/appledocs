//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AdTimeRange


// iOS-only properties

// A Media Player time range that indicates where an ad break exists in the current player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPAdTimeRange/timeRange
func (a_ AdTimeRange) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}
func (a_ AdTimeRange) SetTimeRange(value TimeRange /* not a class type */) {
	a_.ID.Send(objc.RegisterName("setTimeRange:"), value)
}




