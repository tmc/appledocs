//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for RecordedAccelerometerData


// iOS-only properties

// The unique identifier for the accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedAccelerometerData/identifier
func (r_ RecordedAccelerometerData) Identifier() uint64 {
	rv := objc.Send[uint64](r_.ID, objc.Sel("identifier"))
	return rv
}

// The wall clock time when the sensor sample was recorded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedAccelerometerData/startDate
func (r_ RecordedAccelerometerData) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](r_.ID, objc.Sel("startDate"))
	return rv
}





