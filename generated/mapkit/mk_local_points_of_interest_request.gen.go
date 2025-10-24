// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalPointsOfInterestRequest */


/* debug [class_header]: Header for MKLocalPointsOfInterestRequest */
// The class instance for the [MKLocalPointsOfInterestRequest] class.
var (
	MKLocalPointsOfInterestRequestClass     _MKLocalPointsOfInterestRequestClass
	MKLocalPointsOfInterestRequestClassOnce sync.Once
)

func getMKLocalPointsOfInterestRequestClass() _MKLocalPointsOfInterestRequestClass {
	MKLocalPointsOfInterestRequestClassOnce.Do(func() {
		MKLocalPointsOfInterestRequestClass = _MKLocalPointsOfInterestRequestClass{objc.GetClass("MKLocalPointsOfInterestRequest")}
	})
	return MKLocalPointsOfInterestRequestClass
}

type _MKLocalPointsOfInterestRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalPointsOfInterestRequest */
// An interface definition for the [MKLocalPointsOfInterestRequest] class.
type IMKLocalPointsOfInterestRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalPointsOfInterestRequest */
	// properties:
	Coordinate() LocationCoordinate2D /* not a class type */
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	Radius() LocationDistance /* not a class type */
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalPointsOfInterestRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalPointsOfInterestRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalPointsOfInterestRequestClass) Alloc() MKLocalPointsOfInterestRequest {
	rv := objc.Send[MKLocalPointsOfInterestRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalPointsOfInterestRequestClass) New() MKLocalPointsOfInterestRequest {
	rv := objc.Send[MKLocalPointsOfInterestRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalPointsOfInterestRequest) Init() MKLocalPointsOfInterestRequest {
	rv := objc.Send[MKLocalPointsOfInterestRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalPointsOfInterestRequest) Autorelease() MKLocalPointsOfInterestRequest {
	rv := objc.Send[MKLocalPointsOfInterestRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalPointsOfInterestRequest creates a new MKLocalPointsOfInterestRequest instance.
func NewMKLocalPointsOfInterestRequest() MKLocalPointsOfInterestRequest {
	return getMKLocalPointsOfInterestRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalPointsOfInterestRequest */
// A structured request to use when searching for points of interest.
//
// You create an to fetch points of interest within a rectangular bounding box or circular area. To leverage the phone’s viewport to request points of interest, create a request with a rectangular bounding box using an . The request fetches points of interest within the rectangular region. To retrieve points of interest nearby or “around the user,” create a request with a circular area defined by and a in meters. The fetch returns points of interest up to the maximum distance defined by . You may optionally specifying an describing categories to include or exclude. The default behavior of the fetch returns all points of interest.


// A structured request to use when searching for points of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest
type MKLocalPointsOfInterestRequest struct {
	objectivec.Object
}

// MKLocalPointsOfInterestRequestFrom constructs a [MKLocalPointsOfInterestRequest] from an unsafe.Pointer.
//
// A structured request to use when searching for points of interest.
func MKLocalPointsOfInterestRequestFrom(ptr unsafe.Pointer) MKLocalPointsOfInterestRequest {
	return MKLocalPointsOfInterestRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalPointsOfInterestRequest */

// Creates a points of interest search request centered on the provided coordinate with the provided radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/init(center:radius:)
func NewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius(coordinate LocationCoordinate2D /* not a class type */, radius LocationDistance /* not a class type */) MKLocalPointsOfInterestRequest {
	instance := getMKLocalPointsOfInterestRequestClass().Alloc()
	rv := objc.Send[MKLocalPointsOfInterestRequest](instance.ID, objc.Sel("initWithCenterCoordinate:radius:"), coordinate, radius)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius */


// Creates a points of interest search request based on existing region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/init(coordinateRegion:)
func NewMKLocalPointsOfInterestRequestWithCoordinateRegion(region objc.IObject /* cross-framework: MKCoordinateRegion */) MKLocalPointsOfInterestRequest {
	instance := getMKLocalPointsOfInterestRequestClass().Alloc()
	rv := objc.Send[MKLocalPointsOfInterestRequest](instance.ID, objc.Sel("initWithCoordinateRegion:"), region)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalPointsOfInterestRequestWithCoordinateRegion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalPointsOfInterestRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalPointsOfInterestRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalPointsOfInterestRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalPointsOfInterestRequest */

// The center of the point of request as latitude and longitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/coordinate
func (m_ MKLocalPointsOfInterestRequest) Coordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("coordinate"))
	return rv
}/* debug [instance_properties/getter]: coordinate */


// A filter that lists points of interest categories to include or exclude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/pointOfInterestFilter
func (m_ MKLocalPointsOfInterestRequest) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// A filter that lists points of interest categories to include or exclude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/pointOfInterestFilter
func (m_ MKLocalPointsOfInterestRequest) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The distance provided in meters or the longest distance derived from the center point to the region’s bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/radius
func (m_ MKLocalPointsOfInterestRequest) Radius() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// The region of the bounding box of the request provided or the derived bounding box of the circle created by the radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalPointsOfInterestRequest/region
func (m_ MKLocalPointsOfInterestRequest) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalPointsOfInterestRequest */


