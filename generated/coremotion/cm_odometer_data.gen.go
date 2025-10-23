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
	// properties:
	DeltaAltitude() LocationDistance /* not a class type */
	DeltaDistance() LocationDistance /* not a class type */
	DeltaDistanceAccuracy() LocationAccuracy /* not a class type */
	EndDate() foundation.objc.IObject /* cross-framework: NSDate */
	GpsDate() foundation.objc.IObject /* cross-framework: NSDate */
	MaxAbsSlope() foundation.objc.IObject /* cross-framework: Number */
	OriginDevice() OdometerOriginDevice
	Slope() foundation.objc.IObject /* cross-framework: Number */
	Speed() LocationSpeed /* not a class type */
	SpeedAccuracy() LocationSpeedAccuracy /* not a class type */
	StartDate() foundation.objc.IObject /* cross-framework: NSDate */
	VerticalAccuracy() LocationAccuracy /* not a class type */
	// methods:
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
func (o_ OdometerData) DeltaAltitude() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](o_.ID, objc.Sel("deltaAltitude"))
	return rv
}


// The change in distance that the user travels since the last location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistance
func (o_ OdometerData) DeltaDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](o_.ID, objc.Sel("deltaDistance"))
	return rv
}


// The accuracy of the change in distance value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistanceAccuracy
func (o_ OdometerData) DeltaDistanceAccuracy() LocationAccuracy /* not a class type */ {
	rv := objc.Send[LocationAccuracy](o_.ID, objc.Sel("deltaDistanceAccuracy"))
	return rv
}


// The time that the device stops recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/endDate
func (o_ OdometerData) EndDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("endDate"))
	return rv
}


// The time of the GPS measurement associated with the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/gpsDate
func (o_ OdometerData) GpsDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("gpsDate"))
	return rv
}


// The maximum absolute slope at the location toward all directions, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/maxAbsSlope-96ulr
func (o_ OdometerData) MaxAbsSlope() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](o_.ID, objc.Sel("maxAbsSlope"))
	return rv
}


// The device that measures the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/originDevice
func (o_ OdometerData) OriginDevice() OdometerOriginDevice {
	rv := objc.Send[OdometerOriginDevice](o_.ID, objc.Sel("originDevice"))
	return rv
}


// The slope at the location toward the direction of travel, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/slope-96hlt
func (o_ OdometerData) Slope() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](o_.ID, objc.Sel("slope"))
	return rv
}


// The instantaneous velocity of the device, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speed
func (o_ OdometerData) Speed() LocationSpeed /* not a class type */ {
	rv := objc.Send[LocationSpeed](o_.ID, objc.Sel("speed"))
	return rv
}


// The accuracy of the speed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speedAccuracy
func (o_ OdometerData) SpeedAccuracy() LocationSpeedAccuracy /* not a class type */ {
	rv := objc.Send[LocationSpeedAccuracy](o_.ID, objc.Sel("speedAccuracy"))
	return rv
}


// The time that the device starts recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/startDate
func (o_ OdometerData) StartDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("startDate"))
	return rv
}


// The validity of the altitude values and their estimated uncertainty, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/verticalAccuracy
func (o_ OdometerData) VerticalAccuracy() LocationAccuracy /* not a class type */ {
	rv := objc.Send[LocationAccuracy](o_.ID, objc.Sel("verticalAccuracy"))
	return rv
}



