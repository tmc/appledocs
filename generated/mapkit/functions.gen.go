// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

/* debug [functions.gen.go]: Generating 18 functions for MapKit */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MapKit Functions (18 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MKCoordinateRegionForMapRect func(MKMapRect) MKCoordinateRegion
	_MKCoordinateRegionMakeWithDistance func(LocationCoordinate2D, LocationDistance, LocationDistance) MKCoordinateRegion
	_MKCoordinateForMapPoint func(MKMapPoint) LocationCoordinate2D
	_MKMetersBetweenMapPoints func(MKMapPoint, MKMapPoint) LocationDistance
	_MKMapPointForCoordinate func(LocationCoordinate2D) MKMapPoint
	_MKMapPointsPerMeterAtLatitude func(LocationDegrees) float64
	_MKMapRectContainsRect func(MKMapRect, MKMapRect) bool
	_MKMapRectContainsPoint func(MKMapRect, MKMapPoint) bool
	_MKMapRectInset func(MKMapRect, float64, float64) MKMapRect
	_MKMapRectIntersection func(MKMapRect, MKMapRect) MKMapRect
	_MKMapRectIntersectsRect func(MKMapRect, MKMapRect) bool
	_MKMapRectOffset func(MKMapRect, float64, float64) MKMapRect
	_MKMapRectRemainder func(MKMapRect) MKMapRect
	_MKMapRectSpans180thMeridian func(MKMapRect) bool
	_MKMapRectUnion func(MKMapRect, MKMapRect) MKMapRect
	_MKMapRectDivide func(MKMapRect, unsafe.Pointer, unsafe.Pointer, float64, RectEdge)
	_MKMetersPerMapPointAtLatitude func(LocationDegrees) LocationDistance
	_MKRoadWidthAtZoomScale func(MKZoomScale) float64
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MKCoordinateRegionForMapRect, lib, "MKCoordinateRegionForMapRect")
	tryRegister(&_MKCoordinateRegionMakeWithDistance, lib, "MKCoordinateRegionMakeWithDistance")
	tryRegister(&_MKCoordinateForMapPoint, lib, "MKCoordinateForMapPoint")
	tryRegister(&_MKMetersBetweenMapPoints, lib, "MKMetersBetweenMapPoints")
	tryRegister(&_MKMapPointForCoordinate, lib, "MKMapPointForCoordinate")
	tryRegister(&_MKMapPointsPerMeterAtLatitude, lib, "MKMapPointsPerMeterAtLatitude")
	tryRegister(&_MKMapRectContainsRect, lib, "MKMapRectContainsRect")
	tryRegister(&_MKMapRectContainsPoint, lib, "MKMapRectContainsPoint")
	tryRegister(&_MKMapRectInset, lib, "MKMapRectInset")
	tryRegister(&_MKMapRectIntersection, lib, "MKMapRectIntersection")
	tryRegister(&_MKMapRectIntersectsRect, lib, "MKMapRectIntersectsRect")
	tryRegister(&_MKMapRectOffset, lib, "MKMapRectOffset")
	tryRegister(&_MKMapRectRemainder, lib, "MKMapRectRemainder")
	tryRegister(&_MKMapRectSpans180thMeridian, lib, "MKMapRectSpans180thMeridian")
	tryRegister(&_MKMapRectUnion, lib, "MKMapRectUnion")
	tryRegister(&_MKMapRectDivide, lib, "MKMapRectDivide")
	tryRegister(&_MKMetersPerMapPointAtLatitude, lib, "MKMetersPerMapPointAtLatitude")
	tryRegister(&_MKRoadWidthAtZoomScale, lib, "MKRoadWidthAtZoomScale")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the region that corresponds to the specified map rectangle.
//
// Added in macOS 10.9.
// Returns the region that corresponds to the specified map rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCoordinateRegion/init(_:)
func MKCoordinateRegionForMapRect(rect MKMapRect) MKCoordinateRegion {
	return _MKCoordinateRegionForMapRect(rect)
}/* debug [functions.gen.go/function]: MKCoordinateRegionForMapRect */

// Creates a new coordinate region from the specified coordinate and distance values.
//
// Added in macOS .
// Creates a new coordinate region from the specified coordinate and distance values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCoordinateRegion/init(center:latitudinalMeters:longitudinalMeters:)
func MKCoordinateRegionMakeWithDistance(centerCoordinate LocationCoordinate2D, latitudinalMeters LocationDistance, longitudinalMeters LocationDistance) MKCoordinateRegion {
	return _MKCoordinateRegionMakeWithDistance(centerCoordinate, latitudinalMeters, longitudinalMeters)
}/* debug [functions.gen.go/function]: MKCoordinateRegionMakeWithDistance */

// A 2D coordinate that corresponds to the latitude and longitude of the specified map point.
//
// Added in macOS 10.9.
// A 2D coordinate that corresponds to the latitude and longitude of the specified map point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapPoint/coordinate
func MKCoordinateForMapPoint(mapPoint MKMapPoint) LocationCoordinate2D {
	return _MKCoordinateForMapPoint(mapPoint)
}/* debug [functions.gen.go/function]: MKCoordinateForMapPoint */

// Returns the number of meters between two map points.
//
// Added in macOS 10.9.
// Returns the number of meters between two map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapPoint/distance(to:)
func MKMetersBetweenMapPoints(a MKMapPoint, b MKMapPoint) LocationDistance {
	return _MKMetersBetweenMapPoints(a, b)
}/* debug [functions.gen.go/function]: MKMetersBetweenMapPoints */

// Creates the map point data structure that corresponds to the specified coordinate.
//
// Added in macOS 10.9.
// Creates the map point data structure that corresponds to the specified coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapPoint/init(_:)
func MKMapPointForCoordinate(coordinate LocationCoordinate2D) MKMapPoint {
	return _MKMapPointForCoordinate(coordinate)
}/* debug [functions.gen.go/function]: MKMapPointForCoordinate */

// Returns the number of map points that represent one meter at the specified latitude.
//
// Added in macOS 10.9.
// Returns the number of map points that represent one meter at the specified latitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapPointsPerMeterAtLatitude(_:)
func MKMapPointsPerMeterAtLatitude(latitude LocationDegrees) float64 {
	return _MKMapPointsPerMeterAtLatitude(latitude)
}/* debug [functions.gen.go/function]: MKMapPointsPerMeterAtLatitude */

// Returns a Boolean value that indicates whether one rectangle contains another.
//
// Added in macOS 10.9.
// Returns a Boolean value that indicates whether one rectangle contains another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/contains(_:)-1z5oa
func MKMapRectContainsRect(rect1 MKMapRect, rect2 MKMapRect) bool {
	return _MKMapRectContainsRect(rect1, rect2)
}/* debug [functions.gen.go/function]: MKMapRectContainsRect */

// Returns a Boolean value that indicates whether the specified map point lies within the rectangle.
//
// Added in macOS 10.9.
// Returns a Boolean value that indicates whether the specified map point lies within the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/contains(_:)-79tjt
func MKMapRectContainsPoint(rect MKMapRect, point MKMapPoint) bool {
	return _MKMapRectContainsPoint(rect, point)
}/* debug [functions.gen.go/function]: MKMapRectContainsPoint */

// Returns the specified rectangle with an inset by the specified amounts.
//
// Added in macOS 10.9.
// Returns the specified rectangle with an inset by the specified amounts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/insetBy(dx:dy:)
func MKMapRectInset(rect MKMapRect, dx float64, dy float64) MKMapRect {
	return _MKMapRectInset(rect, dx, dy)
}/* debug [functions.gen.go/function]: MKMapRectInset */

// Returns the rectangle that represents the intersection of two rectangles.
//
// Added in macOS 10.9.
// Returns the rectangle that represents the intersection of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/intersection(_:)
func MKMapRectIntersection(rect1 MKMapRect, rect2 MKMapRect) MKMapRect {
	return _MKMapRectIntersection(rect1, rect2)
}/* debug [functions.gen.go/function]: MKMapRectIntersection */

// Returns a Boolean value that indicates whether two rectangles intersect each other.
//
// Added in macOS 10.9.
// Returns a Boolean value that indicates whether two rectangles intersect each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/intersects(_:)
func MKMapRectIntersectsRect(rect1 MKMapRect, rect2 MKMapRect) bool {
	return _MKMapRectIntersectsRect(rect1, rect2)
}/* debug [functions.gen.go/function]: MKMapRectIntersectsRect */

// Returns a rectangle with an origin point that shifts by the specified amount.
//
// Added in macOS 10.9.
// Returns a rectangle with an origin point that shifts by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/offsetBy(dx:dy:)
func MKMapRectOffset(rect MKMapRect, dx float64, dy float64) MKMapRect {
	return _MKMapRectOffset(rect, dx, dy)
}/* debug [functions.gen.go/function]: MKMapRectOffset */

// A rectangle that represents the normalized portion of the specified rectangle that lies outside the world map boundaries.
//
// Added in macOS 10.9.
// A rectangle that represents the normalized portion of the specified rectangle that lies outside the world map boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/remainder
func MKMapRectRemainder(rect MKMapRect) MKMapRect {
	return _MKMapRectRemainder(rect)
}/* debug [functions.gen.go/function]: MKMapRectRemainder */

// A Boolean value that indicates whether the specified map rectangle crosses the 180th meridian.
//
// Added in macOS 10.9.
// A Boolean value that indicates whether the specified map rectangle crosses the 180th meridian.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/spans180thMeridian
func MKMapRectSpans180thMeridian(rect MKMapRect) bool {
	return _MKMapRectSpans180thMeridian(rect)
}/* debug [functions.gen.go/function]: MKMapRectSpans180thMeridian */

// Returns a rectangle that represents the union of two rectangles.
//
// Added in macOS 10.9.
// Returns a rectangle that represents the union of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect/union(_:)
func MKMapRectUnion(rect1 MKMapRect, rect2 MKMapRect) MKMapRect {
	return _MKMapRectUnion(rect1, rect2)
}/* debug [functions.gen.go/function]: MKMapRectUnion */

// Divides the specified rectangle into two smaller rectangles.
//
// Added in macOS 10.9.
// Divides the specified rectangle into two smaller rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRectDivide(_:_:_:_:_:)
func MKMapRectDivide(rect MKMapRect, slice unsafe.Pointer, remainder unsafe.Pointer, amount float64, edge RectEdge) {
	_MKMapRectDivide(rect, slice, remainder, amount, edge)
}/* debug [functions.gen.go/function]: MKMapRectDivide */

// Returns the distance that one map point spans at the specified latitude.
//
// Added in macOS 10.9.
// Returns the distance that one map point spans at the specified latitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMetersPerMapPointAtLatitude(_:)
func MKMetersPerMapPointAtLatitude(latitude LocationDegrees) LocationDistance {
	return _MKMetersPerMapPointAtLatitude(latitude)
}/* debug [functions.gen.go/function]: MKMetersPerMapPointAtLatitude */

// Returns the width (in screen points) of roads on a map at the specified zoom level.
//
// Added in macOS 10.9.
// Returns the width (in screen points) of roads on a map at the specified zoom level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoadWidthAtZoomScale(_:)
func MKRoadWidthAtZoomScale(zoomScale MKZoomScale) float64 {
	return _MKRoadWidthAtZoomScale(zoomScale)
}/* debug [functions.gen.go/function]: MKRoadWidthAtZoomScale */




