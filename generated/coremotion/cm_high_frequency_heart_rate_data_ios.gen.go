//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for HighFrequencyHeartRateData


// iOS-only properties

// The confidence level of the heart rate value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/confidence
func (h_ HighFrequencyHeartRateData) Confidence() HighFrequencyHeartRateDataConfidence {
	rv := objc.Send[HighFrequencyHeartRateDataConfidence](h_.ID, objc.Sel("confidence"))
	return rv
}

// The time the heart rate value occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/date
func (h_ HighFrequencyHeartRateData) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("date"))
	return rv
}

// The heart rate value in units of beats per minute (BPM).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/heartRate
func (h_ HighFrequencyHeartRateData) HeartRate() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("heartRate"))
	return rv
}





