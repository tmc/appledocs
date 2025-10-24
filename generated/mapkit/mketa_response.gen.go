// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKETAResponse */


/* debug [class_header]: Header for MKETAResponse */
// The class instance for the [MKETAResponse] class.
var (
	MKETAResponseClass     _MKETAResponseClass
	MKETAResponseClassOnce sync.Once
)

func getMKETAResponseClass() _MKETAResponseClass {
	MKETAResponseClassOnce.Do(func() {
		MKETAResponseClass = _MKETAResponseClass{objc.GetClass("MKETAResponse")}
	})
	return MKETAResponseClass
}

type _MKETAResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKETAResponse */
// An interface definition for the [MKETAResponse] class.
type IMKETAResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKETAResponse */
	// properties:
	Destination() IMKMapItem
	Distance() LocationDistance /* not a class type */
	ExpectedArrivalDate() objc.IObject /* cross-framework: NSDate */
	ExpectedDepartureDate() objc.IObject /* cross-framework: NSDate */
	ExpectedTravelTime() float64
	Source() IMKMapItem
	TransportType() MKDirectionsTransportType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKETAResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKETAResponse */
// Alloc allocates a new instance without initialization.
func (mc _MKETAResponseClass) Alloc() MKETAResponse {
	rv := objc.Send[MKETAResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKETAResponseClass) New() MKETAResponse {
	rv := objc.Send[MKETAResponse](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKETAResponse) Init() MKETAResponse {
	rv := objc.Send[MKETAResponse](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKETAResponse) Autorelease() MKETAResponse {
	rv := objc.Send[MKETAResponse](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKETAResponse creates a new MKETAResponse instance.
func NewMKETAResponse() MKETAResponse {
	return getMKETAResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKETAResponse */
// The travel-time information that Apple servers return.
//
// You don’t create instances of this class directly. Instead, you initiate a request for the travel time by calling the method of an object. The completion handler you pass to that method receives an object with the results.


// The travel-time information that Apple servers return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse
type MKETAResponse struct {
	objectivec.Object
}

// MKETAResponseFrom constructs a [MKETAResponse] from an unsafe.Pointer.
//
// The travel-time information that Apple servers return.
func MKETAResponseFrom(ptr unsafe.Pointer) MKETAResponse {
	return MKETAResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKETAResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKETAResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKETAResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKETAResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKETAResponse */

// The end point of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/destination
func (m_ MKETAResponse) Destination() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// The expected travel distance, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/distance
func (m_ MKETAResponse) Distance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("distance"))
	return rv
}/* debug [instance_properties/getter]: distance */


// The expected arrival time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/expectedArrivalDate
func (m_ MKETAResponse) ExpectedArrivalDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("expectedArrivalDate"))
	return rv
}/* debug [instance_properties/getter]: expectedArrivalDate */


// The expected departure time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/expectedDepartureDate
func (m_ MKETAResponse) ExpectedDepartureDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("expectedDepartureDate"))
	return rv
}/* debug [instance_properties/getter]: expectedDepartureDate */


// The expected travel time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/expectedTravelTime
func (m_ MKETAResponse) ExpectedTravelTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("expectedTravelTime"))
	return rv
}/* debug [instance_properties/getter]: expectedTravelTime */


// The start point of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/source
func (m_ MKETAResponse) Source() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// The type of conveyance to use for determining the travel time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse/transportType
func (m_ MKETAResponse) TransportType() MKDirectionsTransportType {
	rv := objc.Send[MKDirectionsTransportType](m_.ID, objc.Sel("transportType"))
	return rv
}/* debug [instance_properties/getter]: transportType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKETAResponse */



