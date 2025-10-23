// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PedometerData] class.
var (
	PedometerDataClass     _PedometerDataClass
	PedometerDataClassOnce sync.Once
)

func getPedometerDataClass() _PedometerDataClass {
	PedometerDataClassOnce.Do(func() {
		PedometerDataClass = _PedometerDataClass{objc.GetClass("CMPedometerData")}
	})
	return PedometerDataClass
}

type _PedometerDataClass struct {
	class objc.Class
}

// An interface definition for the [PedometerData] class.
type IPedometerData interface {
	objectivec.IObject
	AverageActivePace() foundation.Number
	CurrentCadence() foundation.Number
	CurrentPace() foundation.Number
	Distance() foundation.Number
	EndDate() foundation.NSDate
	FloorsAscended() foundation.Number
	FloorsDescended() foundation.Number
	NumberOfSteps() foundation.Number
	StartDate() foundation.NSDate
}

// Information about the distance traveled by a user on foot.
//
// You do not create instances of this class yourself. Instead, you use a object to request pedometer data from the system. The data for each request is packaged into an instance of this class and delivered to the handlers you registered with the pedometer object.


// Information about the distance traveled by a user on foot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData
type PedometerData struct {
	objectivec.Object
}

// PedometerDataFrom constructs a [PedometerData] from an unsafe.Pointer.
//
// Information about the distance traveled by a user on foot.
func PedometerDataFrom(ptr unsafe.Pointer) PedometerData {
	return PedometerData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PedometerDataClass) Alloc() PedometerData {
	rv := objc.Send[PedometerData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PedometerDataClass) New() PedometerData {
	rv := objc.Send[PedometerData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PedometerData) Init() PedometerData {
	rv := objc.Send[PedometerData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PedometerData) Autorelease() PedometerData {
	rv := objc.Send[PedometerData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPedometerData creates a new PedometerData instance.
func NewPedometerData() PedometerData {
	return getPedometerDataClass().New()
}



// The average pace of the user, measured in seconds per meter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/averageActivePace
func (p_ PedometerData) AverageActivePace() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("averageActivePace"))
	return rv
}


// The rate at which steps are taken, measured in steps per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/currentCadence
func (p_ PedometerData) CurrentCadence() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("currentCadence"))
	return rv
}


// The current pace of the user, measured in seconds per meter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/currentPace
func (p_ PedometerData) CurrentPace() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("currentPace"))
	return rv
}


// The estimated distance (in meters) traveled by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/distance
func (p_ PedometerData) Distance() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("distance"))
	return rv
}


// The end time for the pedometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/endDate
func (p_ PedometerData) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("endDate"))
	return rv
}


// The approximate number of floors ascended by walking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/floorsAscended
func (p_ PedometerData) FloorsAscended() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("floorsAscended"))
	return rv
}


// The approximate number of floors descended by walking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/floorsDescended
func (p_ PedometerData) FloorsDescended() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("floorsDescended"))
	return rv
}


// The number of steps taken by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/numberOfSteps
func (p_ PedometerData) NumberOfSteps() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("numberOfSteps"))
	return rv
}


// The start time for the pedometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/startDate
func (p_ PedometerData) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}



