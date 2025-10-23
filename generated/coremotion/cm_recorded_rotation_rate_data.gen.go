// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RecordedRotationRateData] class.
var (
	RecordedRotationRateDataClass     _RecordedRotationRateDataClass
	RecordedRotationRateDataClassOnce sync.Once
)

func getRecordedRotationRateDataClass() _RecordedRotationRateDataClass {
	RecordedRotationRateDataClassOnce.Do(func() {
		RecordedRotationRateDataClass = _RecordedRotationRateDataClass{objc.GetClass("CMRecordedRotationRateData")}
	})
	return RecordedRotationRateDataClass
}

type _RecordedRotationRateDataClass struct {
	class objc.Class
}

// An interface definition for the [RecordedRotationRateData] class.
type IRecordedRotationRateData interface {
	IRotationRateData
	StartDate() foundation.NSDate
	RotationRate() unsafe.Pointer
	SetRotationRate(value unsafe.Pointer)
}

// A data object that contains a single rotation-rate measurement at a specific time.


// A data object that contains a single rotation-rate measurement at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedRotationRateData
type RecordedRotationRateData struct {
	RotationRateData
}

// RecordedRotationRateDataFrom constructs a [RecordedRotationRateData] from an unsafe.Pointer.
//
// A data object that contains a single rotation-rate measurement at a specific time.
func RecordedRotationRateDataFrom(ptr unsafe.Pointer) RecordedRotationRateData {
	return RecordedRotationRateData{
		RotationRateData: RotationRateDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecordedRotationRateDataClass) Alloc() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecordedRotationRateDataClass) New() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordedRotationRateData) Init() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordedRotationRateData) Autorelease() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordedRotationRateData creates a new RecordedRotationRateData instance.
func NewRecordedRotationRateData() RecordedRotationRateData {
	return getRecordedRotationRateDataClass().New()
}



// The time when the gyroscope measured the rotation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedRotationRateData/startDate
func (r_ RecordedRotationRateData) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](r_.ID, objc.Sel("startDate"))
	return rv
}


// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RecordedRotationRateData) RotationRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rotationRate"))
	return rv
}


// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RecordedRotationRateData) SetRotationRate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRotationRate:"), value)
}



