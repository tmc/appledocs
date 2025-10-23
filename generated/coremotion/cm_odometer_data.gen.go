// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OdometerData] class.
var (
	OdometerDataClass     _OdometerDataClass
	OdometerDataClassOnce sync.Once
)

func getOdometerDataClass() _OdometerDataClass {
	OdometerDataClassOnce.Do(func() {
		OdometerDataClass = _OdometerDataClass{objc.GetClass("CMOdometerData")}
	})
	return OdometerDataClass
}

type _OdometerDataClass struct {
	class objc.Class
}

// An interface definition for the [OdometerData] class.
type IOdometerData interface {
	objectivec.IObject
	DeltaAltitude() unsafe.Pointer
	DeltaDistance() unsafe.Pointer
	DeltaDistanceAccuracy() unsafe.Pointer
	EndDate() foundation.NSDate
	GpsDate() foundation.NSDate
	MaxAbsSlope() foundation.Number
	OriginDevice() CMOdometerOriginDevice
	Slope() foundation.Number
	Speed() unsafe.Pointer
	SpeedAccuracy() unsafe.Pointer
	StartDate() foundation.NSDate
	VerticalAccuracy() unsafe.Pointer
}

// A class that represents odometer data for workouts.
//
// To get the measurements, use the and properties. To compute distances, use the and properties.


// A class that represents odometer data for workouts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData
type OdometerData struct {
	objectivec.Object
}

// OdometerDataFrom constructs a [OdometerData] from an unsafe.Pointer.
//
// A class that represents odometer data for workouts.
func OdometerDataFrom(ptr unsafe.Pointer) OdometerData {
	return OdometerData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OdometerDataClass) Alloc() OdometerData {
	rv := objc.Send[OdometerData](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OdometerDataClass) New() OdometerData {
	rv := objc.Send[OdometerData](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OdometerData) Init() OdometerData {
	rv := objc.Send[OdometerData](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OdometerData) Autorelease() OdometerData {
	rv := objc.Send[OdometerData](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOdometerData creates a new OdometerData instance.
func NewOdometerData() OdometerData {
	return getOdometerDataClass().New()
}



// The change in altitude above mean sea level associated with the location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaAltitude
func (o_ OdometerData) DeltaAltitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("deltaAltitude"))
	return rv
}


// The change in distance that the user travels since the last location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistance
func (o_ OdometerData) DeltaDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("deltaDistance"))
	return rv
}


// The accuracy of the change in distance value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistanceAccuracy
func (o_ OdometerData) DeltaDistanceAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("deltaDistanceAccuracy"))
	return rv
}


// The time that the device stops recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/endDate
func (o_ OdometerData) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("endDate"))
	return rv
}


// The time of the GPS measurement associated with the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/gpsDate
func (o_ OdometerData) GpsDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("gpsDate"))
	return rv
}


// The maximum absolute slope at the location toward all directions, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/maxAbsSlope-96ulr
func (o_ OdometerData) MaxAbsSlope() foundation.Number {
	rv := objc.Send[foundation.Number](o_.ID, objc.Sel("maxAbsSlope"))
	return rv
}


// The device that measures the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/originDevice
func (o_ OdometerData) OriginDevice() CMOdometerOriginDevice {
	rv := objc.Send[CMOdometerOriginDevice](o_.ID, objc.Sel("originDevice"))
	return rv
}


// The slope at the location toward the direction of travel, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/slope-96hlt
func (o_ OdometerData) Slope() foundation.Number {
	rv := objc.Send[foundation.Number](o_.ID, objc.Sel("slope"))
	return rv
}


// The instantaneous velocity of the device, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speed
func (o_ OdometerData) Speed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("speed"))
	return rv
}


// The accuracy of the speed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speedAccuracy
func (o_ OdometerData) SpeedAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("speedAccuracy"))
	return rv
}


// The time that the device starts recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/startDate
func (o_ OdometerData) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("startDate"))
	return rv
}


// The validity of the altitude values and their estimated uncertainty, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/verticalAccuracy
func (o_ OdometerData) VerticalAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("verticalAccuracy"))
	return rv
}



