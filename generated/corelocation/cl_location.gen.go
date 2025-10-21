// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Location] class.
var (
	LocationClass     _LocationClass
	LocationClassOnce sync.Once
)

func getLocationClass() _LocationClass {
	LocationClassOnce.Do(func() {
		LocationClass = _LocationClass{objc.GetClass("CLLocation")}
	})
	return LocationClass
}

type _LocationClass struct {
	class objc.Class
}

// An interface definition for the [Location] class.
type ILocation interface {
	objectivec.IObject
	Distance()
	GetDistanceFrom()
}

// The latitude, longitude, and course information reported by the system.
//
// A object contains the geographical location and altitude of a device, along with values indicating the accuracy of those measurements and when they were collected. In iOS, a location object also contains course information — that is, the speed and heading in which the device was moving. Typically, you don’t create location objects yourself. After you request location updates from your object, the system uses onboard sensors to gather location data and report that data to your app. Some services also return previously collected location data, which you can use as context to improve your services. You can always retrieve the most recently collected location from the property of your object. You may create location objects yourself when you want to cache custom location data or calculate the distance between two geographical coordinates. Use objects as-is, and don’t subclass them.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation
type Location struct {
	objectivec.Object
}

// LocationFrom constructs a [Location] from an unsafe.Pointer.
//
// The latitude, longitude, and course information reported by the system.
func LocationFrom(ptr unsafe.Pointer) Location {
	return Location{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LocationClass) Alloc() Location {
	rv := objc.Send[Location](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocationClass) New() Location {
	rv := objc.Send[Location](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Location) Init() Location {
	rv := objc.Send[Location](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Location) Autorelease() Location {
	rv := objc.Send[Location](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocation creates a new Location instance.
func NewLocation() Location {
	return getLocationClass().New()
}


// Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/distance(from:)
func (l_ Location) Distance() {
	objc.Send[objc.ID](l_.ID, objc.Sel("distance"))
}

// Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/getDistanceFrom(_:)
func (l_ Location) GetDistanceFrom() {
	objc.Send[objc.ID](l_.ID, objc.Sel("getDistanceFrom"))
}

// The altitude above mean sea level associated with a location, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/altitude
func (l_ Location) Altitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("altitude"))
	return rv
}


// SetAltitude sets the value of the altitude property.
// The altitude above mean sea level associated with a location, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/altitude
func (l_ Location) SetAltitude(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAltitude:"), value)
}

// The geographical coordinate information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/coordinate
func (l_ Location) Coordinate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("coordinate"))
	return rv
}


// SetCoordinate sets the value of the coordinate property.
// The geographical coordinate information.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/coordinate
func (l_ Location) SetCoordinate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCoordinate:"), value)
}

// The direction in which the device is traveling, measured in degrees and relative to due north.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/course
func (l_ Location) Course() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("course"))
	return rv
}


// SetCourse sets the value of the course property.
// The direction in which the device is traveling, measured in degrees and relative to due north.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/course
func (l_ Location) SetCourse(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCourse:"), value)
}

// The accuracy of the course value, measured in degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/courseAccuracy
func (l_ Location) CourseAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("courseAccuracy"))
	return rv
}


// SetCourseAccuracy sets the value of the courseAccuracy property.
// The accuracy of the course value, measured in degrees.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/courseAccuracy
func (l_ Location) SetCourseAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCourseAccuracy:"), value)
}

// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/ellipsoidalAltitude
func (l_ Location) EllipsoidalAltitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("ellipsoidalAltitude"))
	return rv
}


// SetEllipsoidalAltitude sets the value of the ellipsoidalAltitude property.
// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/ellipsoidalAltitude
func (l_ Location) SetEllipsoidalAltitude(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEllipsoidalAltitude:"), value)
}

// The radius of uncertainty for the location, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/horizontalAccuracy
func (l_ Location) HorizontalAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("horizontalAccuracy"))
	return rv
}


// SetHorizontalAccuracy sets the value of the horizontalAccuracy property.
// The radius of uncertainty for the location, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/horizontalAccuracy
func (l_ Location) SetHorizontalAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHorizontalAccuracy:"), value)
}

// Information about the source that provides the location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/sourceInformation
func (l_ Location) SourceInformation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("sourceInformation"))
	return rv
}


// SetSourceInformation sets the value of the sourceInformation property.
// Information about the source that provides the location.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/sourceInformation
func (l_ Location) SetSourceInformation(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSourceInformation:"), value)
}

// The instantaneous speed of the device, measured in meters per second.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speed
func (l_ Location) Speed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("speed"))
	return rv
}


// SetSpeed sets the value of the speed property.
// The instantaneous speed of the device, measured in meters per second.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speed
func (l_ Location) SetSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpeed:"), value)
}

// The accuracy of the speed value, measured in meters per second.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speedAccuracy
func (l_ Location) SpeedAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("speedAccuracy"))
	return rv
}


// SetSpeedAccuracy sets the value of the speedAccuracy property.
// The accuracy of the speed value, measured in meters per second.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speedAccuracy
func (l_ Location) SetSpeedAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpeedAccuracy:"), value)
}

// The time at which this location was determined.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/timestamp
func (l_ Location) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// The time at which this location was determined.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/timestamp
func (l_ Location) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimestamp:"), value)
}

// The validity of the altitude values, and their estimated uncertainty, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/verticalAccuracy
func (l_ Location) VerticalAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("verticalAccuracy"))
	return rv
}


// SetVerticalAccuracy sets the value of the verticalAccuracy property.
// The validity of the altitude values, and their estimated uncertainty, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/verticalAccuracy
func (l_ Location) SetVerticalAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setVerticalAccuracy:"), value)
}



