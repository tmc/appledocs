//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for AbsoluteAltitudeData


// iOS-only properties

// The estimated uncertainty of the altimeter in meters, based on one standard deviation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/accuracy
func (a_ AbsoluteAltitudeData) Accuracy() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("accuracy"))
	return rv
}

// The absolute altitude of the device relative to sea level, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/altitude
func (a_ AbsoluteAltitudeData) Altitude() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("altitude"))
	return rv
}

// The recommended resolution for the altitude, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/precision
func (a_ AbsoluteAltitudeData) Precision() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("precision"))
	return rv
}





