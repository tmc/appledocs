// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AbsoluteAltitudeData] class.
var (
	AbsoluteAltitudeDataClass     _AbsoluteAltitudeDataClass
	AbsoluteAltitudeDataClassOnce sync.Once
)

func getAbsoluteAltitudeDataClass() _AbsoluteAltitudeDataClass {
	AbsoluteAltitudeDataClassOnce.Do(func() {
		AbsoluteAltitudeDataClass = _AbsoluteAltitudeDataClass{objc.GetClass("CMAbsoluteAltitudeData")}
	})
	return AbsoluteAltitudeDataClass
}

type _AbsoluteAltitudeDataClass struct {
	class objc.Class
}

// An interface definition for the [AbsoluteAltitudeData] class.
type IAbsoluteAltitudeData interface {
	ILogItem
	Accuracy() float64
	Altitude() float64
	Precision() float64
}

// Data that records a change in absolute altitude.
//
// Absolute altitude is only available on iPhone 12 and later and Apple Watch 6 or SE and later.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData
type AbsoluteAltitudeData struct {
	LogItem
}

// AbsoluteAltitudeDataFrom constructs a [AbsoluteAltitudeData] from an unsafe.Pointer.
//
// Data that records a change in absolute altitude.
func AbsoluteAltitudeDataFrom(ptr unsafe.Pointer) AbsoluteAltitudeData {
	return AbsoluteAltitudeData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AbsoluteAltitudeDataClass) Alloc() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AbsoluteAltitudeDataClass) New() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AbsoluteAltitudeData) Init() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AbsoluteAltitudeData) Autorelease() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAbsoluteAltitudeData creates a new AbsoluteAltitudeData instance.
func NewAbsoluteAltitudeData() AbsoluteAltitudeData {
	return getAbsoluteAltitudeDataClass().New()
}


// The estimated uncertainty of the altimeter in meters, based on one standard deviation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/accuracy
func (a_ AbsoluteAltitudeData) Accuracy() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("accuracy"))
	return rv
}

// The absolute altitude of the device relative to sea level, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/altitude
func (a_ AbsoluteAltitudeData) Altitude() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("altitude"))
	return rv
}

// The recommended resolution for the altitude, in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData/precision
func (a_ AbsoluteAltitudeData) Precision() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("precision"))
	return rv
}



