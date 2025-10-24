// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKRoute */


/* debug [class_header]: Header for MKRoute */
// The class instance for the [MKRoute] class.
var (
	MKRouteClass     _MKRouteClass
	MKRouteClassOnce sync.Once
)

func getMKRouteClass() _MKRouteClass {
	MKRouteClassOnce.Do(func() {
		MKRouteClass = _MKRouteClass{objc.GetClass("MKRoute")}
	})
	return MKRouteClass
}

type _MKRouteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKRoute */
// An interface definition for the [MKRoute] class.
type IMKRoute interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKRoute */
	// properties:
	AdvisoryNotices() []string
	Distance() LocationDistance /* not a class type */
	ExpectedTravelTime() float64
	HasHighways() bool
	HasTolls() bool
	Name() objc.IObject /* cross-framework: NSString */
	Polyline() IMKPolyline
	Steps() []MKRouteStep
	TransportType() MKDirectionsTransportType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKRoute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKRoute */
// Alloc allocates a new instance without initialization.
func (mc _MKRouteClass) Alloc() MKRoute {
	rv := objc.Send[MKRoute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKRouteClass) New() MKRoute {
	rv := objc.Send[MKRoute](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKRoute) Init() MKRoute {
	rv := objc.Send[MKRoute](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKRoute) Autorelease() MKRoute {
	rv := objc.Send[MKRoute](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKRoute creates a new MKRoute instance.
func NewMKRoute() MKRoute {
	return getMKRouteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKRoute */
// A single route between a requested start and end point.
//
// An object defines the geometry for the route — that is, it contains line segments associated with specific map coordinates. A route object may also include other information, such as the name of the route, its distance, and the expected travel time. You don’t create instances of this class directly. When you use an object to request directions from Apple, the returned object contains the possible routes.


// A single route between a requested start and end point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute
type MKRoute struct {
	objectivec.Object
}

// MKRouteFrom constructs a [MKRoute] from an unsafe.Pointer.
//
// A single route between a requested start and end point.
func MKRouteFrom(ptr unsafe.Pointer) MKRoute {
	return MKRoute{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKRoute *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKRoute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKRoute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKRoute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKRoute */

// An array of advisory notice strings for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/advisoryNotices
func (m_ MKRoute) AdvisoryNotices() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("advisoryNotices"))
	return rv
}/* debug [instance_properties/getter]: advisoryNotices */


// The route distance, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/distance
func (m_ MKRoute) Distance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("distance"))
	return rv
}/* debug [instance_properties/getter]: distance */


// The expected travel time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/expectedTravelTime
func (m_ MKRoute) ExpectedTravelTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("expectedTravelTime"))
	return rv
}/* debug [instance_properties/getter]: expectedTravelTime */


// A Boolean value that indicates whether the route contains highways.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/hasHighways
func (m_ MKRoute) HasHighways() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasHighways"))
	return rv
}/* debug [instance_properties/getter]: hasHighways */


// A Boolean value that indicates whether the route has tolls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/hasTolls
func (m_ MKRoute) HasTolls() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasTolls"))
	return rv
}/* debug [instance_properties/getter]: hasTolls */


// The assigned name for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/name
func (m_ MKRoute) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The detailed route geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/polyline
func (m_ MKRoute) Polyline() IMKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("polyline"))
	return rv
}/* debug [instance_properties/getter]: polyline */


// The array of steps that create the overall route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/steps
func (m_ MKRoute) Steps() []MKRouteStep {
	rv := objc.Send[[]MKRouteStep](m_.ID, objc.Sel("steps"))
	return rv
}/* debug [instance_properties/getter]: steps */


// The overall route transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/transportType
func (m_ MKRoute) TransportType() MKDirectionsTransportType {
	rv := objc.Send[MKDirectionsTransportType](m_.ID, objc.Sel("transportType"))
	return rv
}/* debug [instance_properties/getter]: transportType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKRoute */



