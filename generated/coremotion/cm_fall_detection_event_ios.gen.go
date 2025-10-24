//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FallDetectionEvent


// iOS-only properties

// The event’s time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/date
func (f_ FallDetectionEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](f_.ID, objc.Sel("date"))
	return rv
}

// The event’s resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/resolution
func (f_ FallDetectionEvent) Resolution() FallDetectionEventUserResolution {
	rv := objc.Send[FallDetectionEventUserResolution](f_.ID, objc.Sel("resolution"))
	return rv
}





