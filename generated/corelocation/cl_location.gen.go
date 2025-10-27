// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	Altitude() LocationDistance get /* not a class type */
	SetAltitude(value LocationDistance get /* not a class type */)
	Coordinate() LocationCoordinate2D get /* not a class type */
	SetCoordinate(value LocationCoordinate2D get /* not a class type */)
	Course() LocationDirection get /* not a class type */
	SetCourse(value LocationDirection get /* not a class type */)
	CourseAccuracy() LocationDirectionAccuracy get /* not a class type */
	SetCourseAccuracy(value LocationDirectionAccuracy get /* not a class type */)
	EllipsoidalAltitude() LocationDistance get /* not a class type */
	SetEllipsoidalAltitude(value LocationDistance get /* not a class type */)
	HorizontalAccuracy() LocationAccuracy get /* not a class type */
	SetHorizontalAccuracy(value LocationAccuracy get /* not a class type */)
	SourceInformation() ICLLocationSourceInformation
	SetSourceInformation(value ICLLocationSourceInformation)
	Speed() LocationSpeed get /* not a class type */
	SetSpeed(value LocationSpeed get /* not a class type */)
	SpeedAccuracy() LocationSpeedAccuracy get /* not a class type */
	SetSpeedAccuracy(value LocationSpeedAccuracy get /* not a class type */)
	Timestamp() objectivec.IObject
	SetTimestamp(value objectivec.IObject)
	VerticalAccuracy() LocationAccuracy get /* not a class type */
	SetVerticalAccuracy(value LocationAccuracy get /* not a class type */)


	

	// methods:
	Distance()


}





// Alloc allocates a new instance without initialization.
func (lc _LocationClass) Alloc() Location {
	rv := objc.Send[Location](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The latitude, longitude, and course information reported by the system.
//
// A object contains the geographical location and altitude of a device, along with values indicating the accuracy of those measurements and when they were collected. In iOS, a location object also contains course information — that is, the speed and heading in which the device was moving. Typically, you don’t create location objects yourself. After you request location updates from your object, the system uses onboard sensors to gather location data and report that data to your app. Some services also return previously collected location data, which you can use as context to improve your services. You can always retrieve the most recently collected location from the property of your object. You may create location objects yourself when you want to cache custom location data or calculate the distance between two geographical coordinates. Use objects as-is, and don’t subclass them.


// The latitude, longitude, and course information reported by the system.
//
// [Full Topic]
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




















// Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/distance(from:)
func (l_ Location) Distance() {
	objc.Send[objc.ID](l_.ID, objc.Sel("distance"))
}







// The altitude above mean sea level associated with a location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/altitude
func (l_ Location) Altitude() LocationDistance get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude above mean sea level associated with a location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/altitude
func (l_ Location) SetAltitude(value LocationDistance get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAltitude:"), value)
}


// The geographical coordinate information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/coordinate
func (l_ Location) Coordinate() LocationCoordinate2D get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("coordinate"))
	return rv
}


// The geographical coordinate information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/coordinate
func (l_ Location) SetCoordinate(value LocationCoordinate2D get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCoordinate:"), value)
}


// The direction in which the device is traveling, measured in degrees and relative to due north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/course
func (l_ Location) Course() LocationDirection get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("course"))
	return rv
}


// The direction in which the device is traveling, measured in degrees and relative to due north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/course
func (l_ Location) SetCourse(value LocationDirection get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCourse:"), value)
}


// The accuracy of the course value, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/courseAccuracy
func (l_ Location) CourseAccuracy() LocationDirectionAccuracy get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("courseAccuracy"))
	return rv
}


// The accuracy of the course value, measured in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/courseAccuracy
func (l_ Location) SetCourseAccuracy(value LocationDirectionAccuracy get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCourseAccuracy:"), value)
}


// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/ellipsoidalAltitude
func (l_ Location) EllipsoidalAltitude() LocationDistance get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("ellipsoidalAltitude"))
	return rv
}


// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/ellipsoidalAltitude
func (l_ Location) SetEllipsoidalAltitude(value LocationDistance get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEllipsoidalAltitude:"), value)
}


// The radius of uncertainty for the location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/horizontalAccuracy
func (l_ Location) HorizontalAccuracy() LocationAccuracy get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("horizontalAccuracy"))
	return rv
}


// The radius of uncertainty for the location, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/horizontalAccuracy
func (l_ Location) SetHorizontalAccuracy(value LocationAccuracy get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHorizontalAccuracy:"), value)
}


// Information about the source that provides the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/sourceInformation
func (l_ Location) SourceInformation() ICLLocationSourceInformation {
	rv := objc.Send[LocationSourceInformation](l_.ID, objc.Sel("sourceInformation"))
	return rv
}


// Information about the source that provides the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/sourceInformation
func (l_ Location) SetSourceInformation(value ICLLocationSourceInformation) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSourceInformation:"), value)
}


// The instantaneous speed of the device, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speed
func (l_ Location) Speed() LocationSpeed get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("speed"))
	return rv
}


// The instantaneous speed of the device, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speed
func (l_ Location) SetSpeed(value LocationSpeed get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpeed:"), value)
}


// The accuracy of the speed value, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speedAccuracy
func (l_ Location) SpeedAccuracy() LocationSpeedAccuracy get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("speedAccuracy"))
	return rv
}


// The accuracy of the speed value, measured in meters per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/speedAccuracy
func (l_ Location) SetSpeedAccuracy(value LocationSpeedAccuracy get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpeedAccuracy:"), value)
}


// The time at which this location was determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/timestamp
func (l_ Location) Timestamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("timestamp"))
	return rv
}


// The time at which this location was determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/timestamp
func (l_ Location) SetTimestamp(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimestamp:"), value)
}


// The validity of the altitude values, and their estimated uncertainty, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/verticalAccuracy
func (l_ Location) VerticalAccuracy() LocationAccuracy get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("verticalAccuracy"))
	return rv
}


// The validity of the altitude values, and their estimated uncertainty, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/verticalAccuracy
func (l_ Location) SetVerticalAccuracy(value LocationAccuracy get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setVerticalAccuracy:"), value)
}








