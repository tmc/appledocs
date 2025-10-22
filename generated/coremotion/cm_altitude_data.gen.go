// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AltitudeData] class.
var (
	AltitudeDataClass     _AltitudeDataClass
	AltitudeDataClassOnce sync.Once
)

func getAltitudeDataClass() _AltitudeDataClass {
	AltitudeDataClassOnce.Do(func() {
		AltitudeDataClass = _AltitudeDataClass{objc.GetClass("CMAltitudeData")}
	})
	return AltitudeDataClass
}

type _AltitudeDataClass struct {
	class objc.Class
}

// An interface definition for the [AltitudeData] class.
type IAltitudeData interface {
	ILogItem
	Pressure() foundation.Number
	RelativeAltitude() foundation.Number
}

// Data for a recorded change in altitude.
//
// You do not create instances of this class directly. When you want to receive altimeter changes, create an instance of the class and use that object to query for events or to start the delivery of events. The altimeter object creates new instances of this class at appropriate times and delivers them to the handler you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData
type AltitudeData struct {
	LogItem
}

// AltitudeDataFrom constructs a [AltitudeData] from an unsafe.Pointer.
//
// Data for a recorded change in altitude.
func AltitudeDataFrom(ptr unsafe.Pointer) AltitudeData {
	return AltitudeData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AltitudeDataClass) Alloc() AltitudeData {
	rv := objc.Send[AltitudeData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AltitudeDataClass) New() AltitudeData {
	rv := objc.Send[AltitudeData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AltitudeData) Init() AltitudeData {
	rv := objc.Send[AltitudeData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AltitudeData) Autorelease() AltitudeData {
	rv := objc.Send[AltitudeData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAltitudeData creates a new AltitudeData instance.
func NewAltitudeData() AltitudeData {
	return getAltitudeDataClass().New()
}


// The recorded pressure, in kilopascals.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData/pressure
func (a_ AltitudeData) Pressure() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("pressure"))
	return rv
}

// The change in altitude (in meters) since the first reported event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData/relativeAltitude
func (a_ AltitudeData) RelativeAltitude() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("relativeAltitude"))
	return rv
}



