//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SensorRecorder


// Retrieves the accelerometer data collected between the specified dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/accelerometerData(from:to:)
func (s_ SensorRecorder) AccelerometerDataFromDateToDate(fromDate objc.IObject /* cross-framework: NSDate */, toDate objc.IObject /* cross-framework: NSDate */) ISensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("accelerometerDataFromDate:toDate:"), fromDate, toDate)
	return rv
}

// Begins recording accelerometer data for the specified period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorRecorder/recordAccelerometer(forDuration:)
func (s_ SensorRecorder) RecordAccelerometerForDuration(duration float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("recordAccelerometerForDuration:"), duration)
}

// iOS-only properties





