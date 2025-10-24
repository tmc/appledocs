// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMOdometerData */


/* debug [class_header]: Header for CMOdometerData */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OdometerData */
// An interface definition for the [OdometerData] class.
type IOdometerData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OdometerData */
	// properties:
	DeltaAltitude() LocationDistance /* not a class type */
	DeltaDistance() LocationDistance /* not a class type */
	DeltaDistanceAccuracy() LocationAccuracy /* not a class type */
	EndDate() objc.IObject /* cross-framework: NSDate */
	GpsDate() objc.IObject /* cross-framework: NSDate */
	MaxAbsSlope() objc.IObject /* cross-framework: NSNumber */
	OriginDevice() OdometerOriginDevice
	Slope() objc.IObject /* cross-framework: NSNumber */
	Speed() LocationSpeed /* not a class type */
	SpeedAccuracy() LocationSpeedAccuracy /* not a class type */
	StartDate() objc.IObject /* cross-framework: NSDate */
	VerticalAccuracy() LocationAccuracy /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OdometerData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OdometerData */
// Alloc allocates a new instance without initialization.
func (oc _OdometerDataClass) Alloc() OdometerData {
	rv := objc.Send[OdometerData](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OdometerData */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OdometerData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OdometerData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OdometerData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OdometerData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OdometerData */

// The change in altitude above mean sea level associated with the location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaAltitude
func (o_ OdometerData) DeltaAltitude() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](o_.ID, objc.Sel("deltaAltitude"))
	return rv
}/* debug [instance_properties/getter]: deltaAltitude */


// The change in distance that the user travels since the last location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistance
func (o_ OdometerData) DeltaDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](o_.ID, objc.Sel("deltaDistance"))
	return rv
}/* debug [instance_properties/getter]: deltaDistance */


// The accuracy of the change in distance value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/deltaDistanceAccuracy
func (o_ OdometerData) DeltaDistanceAccuracy() LocationAccuracy /* not a class type */ {
	rv := objc.Send[LocationAccuracy](o_.ID, objc.Sel("deltaDistanceAccuracy"))
	return rv
}/* debug [instance_properties/getter]: deltaDistanceAccuracy */


// The time that the device stops recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/endDate
func (o_ OdometerData) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The time of the GPS measurement associated with the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/gpsDate
func (o_ OdometerData) GpsDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("gpsDate"))
	return rv
}/* debug [instance_properties/getter]: gpsDate */


// The maximum absolute slope at the location toward all directions, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/maxAbsSlope-96ulr
func (o_ OdometerData) MaxAbsSlope() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](o_.ID, objc.Sel("maxAbsSlope"))
	return rv
}/* debug [instance_properties/getter]: maxAbsSlope */


// The device that measures the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/originDevice
func (o_ OdometerData) OriginDevice() OdometerOriginDevice {
	rv := objc.Send[OdometerOriginDevice](o_.ID, objc.Sel("originDevice"))
	return rv
}/* debug [instance_properties/getter]: originDevice */


// The slope at the location toward the direction of travel, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/slope-96hlt
func (o_ OdometerData) Slope() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](o_.ID, objc.Sel("slope"))
	return rv
}/* debug [instance_properties/getter]: slope */


// The instantaneous velocity of the device, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speed
func (o_ OdometerData) Speed() LocationSpeed /* not a class type */ {
	rv := objc.Send[LocationSpeed](o_.ID, objc.Sel("speed"))
	return rv
}/* debug [instance_properties/getter]: speed */


// The accuracy of the speed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/speedAccuracy
func (o_ OdometerData) SpeedAccuracy() LocationSpeedAccuracy /* not a class type */ {
	rv := objc.Send[LocationSpeedAccuracy](o_.ID, objc.Sel("speedAccuracy"))
	return rv
}/* debug [instance_properties/getter]: speedAccuracy */


// The time that the device starts recording the odometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/startDate
func (o_ OdometerData) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](o_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The validity of the altitude values and their estimated uncertainty, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerData/verticalAccuracy
func (o_ OdometerData) VerticalAccuracy() LocationAccuracy /* not a class type */ {
	rv := objc.Send[LocationAccuracy](o_.ID, objc.Sel("verticalAccuracy"))
	return rv
}/* debug [instance_properties/getter]: verticalAccuracy */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMOdometerData */



