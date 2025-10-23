// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKRoute] class.
type IMKRoute interface {
	objectivec.IObject
	AdvisoryNotices() string
	SetAdvisoryNotices(value string)
	Distance() unsafe.Pointer
	SetDistance(value unsafe.Pointer)
	ExpectedTravelTime() unsafe.Pointer
	SetExpectedTravelTime(value unsafe.Pointer)
	HasHighways() bool
	SetHasHighways(value bool)
	HasTolls() bool
	SetHasTolls(value bool)
	Name() string
	SetName(value string)
	Polyline() MKPolyline
	SetPolyline(value IMKPolyline)
	Steps() MKRouteStep
	SetSteps(value IMKRouteStep)
	TransportType() unsafe.Pointer
	SetTransportType(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (mc _MKRouteClass) Alloc() MKRoute {
	rv := objc.Send[MKRoute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An array of advisory notice strings for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/advisorynotices
func (m_ MKRoute) AdvisoryNotices() string {
	rv := objc.Send[string](m_.ID, objc.Sel("advisoryNotices"))
	return rv
}


// An array of advisory notice strings for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/advisorynotices
func (m_ MKRoute) SetAdvisoryNotices(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdvisoryNotices:"), objc.String(value))
}


// The route distance, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/distance
func (m_ MKRoute) Distance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("distance"))
	return rv
}


// The route distance, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/distance
func (m_ MKRoute) SetDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistance:"), value)
}


// The expected travel time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/expectedtraveltime
func (m_ MKRoute) ExpectedTravelTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("expectedTravelTime"))
	return rv
}


// The expected travel time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/expectedtraveltime
func (m_ MKRoute) SetExpectedTravelTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedTravelTime:"), value)
}


// A Boolean value that indicates whether the route contains highways.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/hashighways
func (m_ MKRoute) HasHighways() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasHighways"))
	return rv
}


// A Boolean value that indicates whether the route contains highways.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/hashighways
func (m_ MKRoute) SetHasHighways(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasHighways:"), value)
}


// A Boolean value that indicates whether the route has tolls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/hastolls
func (m_ MKRoute) HasTolls() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasTolls"))
	return rv
}


// A Boolean value that indicates whether the route has tolls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/hastolls
func (m_ MKRoute) SetHasTolls(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasTolls:"), value)
}


// The assigned name for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/name
func (m_ MKRoute) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The assigned name for the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/name
func (m_ MKRoute) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}


// The detailed route geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/polyline
func (m_ MKRoute) Polyline() MKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("polyline"))
	return rv
}


// The detailed route geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/polyline
func (m_ MKRoute) SetPolyline(value IMKPolyline) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPolyline:"), value)
}


// The array of steps that create the overall route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/steps
func (m_ MKRoute) Steps() MKRouteStep {
	rv := objc.Send[MKRouteStep](m_.ID, objc.Sel("steps"))
	return rv
}


// The array of steps that create the overall route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/steps
func (m_ MKRoute) SetSteps(value IMKRouteStep) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSteps:"), value)
}


// The overall route transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/transporttype
func (m_ MKRoute) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transportType"))
	return rv
}


// The overall route transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkroute/transporttype
func (m_ MKRoute) SetTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransportType:"), value)
}



