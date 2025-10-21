// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HighFrequencyHeartRateData] class.
var (
	HighFrequencyHeartRateDataClass     _HighFrequencyHeartRateDataClass
	HighFrequencyHeartRateDataClassOnce sync.Once
)

func getHighFrequencyHeartRateDataClass() _HighFrequencyHeartRateDataClass {
	HighFrequencyHeartRateDataClassOnce.Do(func() {
		HighFrequencyHeartRateDataClass = _HighFrequencyHeartRateDataClass{objc.GetClass("CMHighFrequencyHeartRateData")}
	})
	return HighFrequencyHeartRateDataClass
}

type _HighFrequencyHeartRateDataClass struct {
	class objc.Class
}

// An interface definition for the [HighFrequencyHeartRateData] class.
type IHighFrequencyHeartRateData interface {
	ILogItem
}

// A class that represents heart rate data collected at 1 Hz.
//
// Use the property to get the data, and the property for the accuracy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData
type HighFrequencyHeartRateData struct {
	LogItem
}

// HighFrequencyHeartRateDataFrom constructs a [HighFrequencyHeartRateData] from an unsafe.Pointer.
//
// A class that represents heart rate data collected at 1 Hz.
func HighFrequencyHeartRateDataFrom(ptr unsafe.Pointer) HighFrequencyHeartRateData {
	return HighFrequencyHeartRateData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HighFrequencyHeartRateDataClass) Alloc() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HighFrequencyHeartRateDataClass) New() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HighFrequencyHeartRateData) Init() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HighFrequencyHeartRateData) Autorelease() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHighFrequencyHeartRateData creates a new HighFrequencyHeartRateData instance.
func NewHighFrequencyHeartRateData() HighFrequencyHeartRateData {
	return getHighFrequencyHeartRateDataClass().New()
}


// The confidence level of the heart rate value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/confidence
func (h_ HighFrequencyHeartRateData) Confidence() HighFrequencyHeartRateDataConfidence {
	rv := objc.Send[HighFrequencyHeartRateDataConfidence](h_.ID, objc.Sel("confidence"))
	return rv
}

// The time the heart rate value occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/date
func (h_ HighFrequencyHeartRateData) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("date"))
	return rv
}

// The heart rate value in units of beats per minute (BPM).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData/heartRate
func (h_ HighFrequencyHeartRateData) HeartRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("heartRate"))
	return rv
}



