//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for RecordedPressureData


// iOS-only properties

// A value that uniquely identifies this measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData/identifier
func (r_ RecordedPressureData) Identifier() uint64 {
	rv := objc.Send[uint64](r_.ID, objc.Sel("identifier"))
	return rv
}

// The time and date when the system recorded the measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData/startDate
func (r_ RecordedPressureData) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](r_.ID, objc.Sel("startDate"))
	return rv
}





