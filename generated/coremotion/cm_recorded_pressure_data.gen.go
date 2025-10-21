// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecordedPressureData] class.
var (
	RecordedPressureDataClass     _RecordedPressureDataClass
	RecordedPressureDataClassOnce sync.Once
)

func getRecordedPressureDataClass() _RecordedPressureDataClass {
	RecordedPressureDataClassOnce.Do(func() {
		RecordedPressureDataClass = _RecordedPressureDataClass{objc.GetClass("CMRecordedPressureData")}
	})
	return RecordedPressureDataClass
}

type _RecordedPressureDataClass struct {
	class objc.Class
}

// An interface definition for the [RecordedPressureData] class.
type IRecordedPressureData interface {
	IAmbientPressureData
}

// A recorded measurement of pressure data.
//
// Use SensorKit’s sensor to read ambient pressure data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData
type RecordedPressureData struct {
	AmbientPressureData
}

// RecordedPressureDataFrom constructs a [RecordedPressureData] from an unsafe.Pointer.
//
// A recorded measurement of pressure data.
func RecordedPressureDataFrom(ptr unsafe.Pointer) RecordedPressureData {
	return RecordedPressureData{
		AmbientPressureData: AmbientPressureDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecordedPressureDataClass) Alloc() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecordedPressureDataClass) New() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordedPressureData) Init() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordedPressureData) Autorelease() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordedPressureData creates a new RecordedPressureData instance.
func NewRecordedPressureData() RecordedPressureData {
	return getRecordedPressureDataClass().New()
}


// A value that uniquely identifies this measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData/identifier
func (r_ RecordedPressureData) Identifier() uint64 {
	rv := objc.Send[uint64](r_.ID, objc.Sel("identifier"))
	return rv
}

// The time and date when the system recorded the measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData/startDate
func (r_ RecordedPressureData) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("startDate"))
	return rv
}



