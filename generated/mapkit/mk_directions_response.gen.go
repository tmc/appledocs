// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKDirectionsResponse */


/* debug [class_header]: Header for MKDirectionsResponse */
// The class instance for the [MKDirectionsResponse] class.
var (
	MKDirectionsResponseClass     _MKDirectionsResponseClass
	MKDirectionsResponseClassOnce sync.Once
)

func getMKDirectionsResponseClass() _MKDirectionsResponseClass {
	MKDirectionsResponseClassOnce.Do(func() {
		MKDirectionsResponseClass = _MKDirectionsResponseClass{objc.GetClass("MKDirectionsResponse")}
	})
	return MKDirectionsResponseClass
}

type _MKDirectionsResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKDirectionsResponse */
// An interface definition for the [MKDirectionsResponse] class.
type IMKDirectionsResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKDirectionsResponse */
	// properties:
	Destination() IMKMapItem
	Routes() []MKRoute
	Source() IMKMapItem
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKDirectionsResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKDirectionsResponse */
// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsResponseClass) Alloc() MKDirectionsResponse {
	rv := objc.Send[MKDirectionsResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKDirectionsResponseClass) New() MKDirectionsResponse {
	rv := objc.Send[MKDirectionsResponse](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDirectionsResponse) Init() MKDirectionsResponse {
	rv := objc.Send[MKDirectionsResponse](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDirectionsResponse) Autorelease() MKDirectionsResponse {
	rv := objc.Send[MKDirectionsResponse](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDirectionsResponse creates a new MKDirectionsResponse instance.
func NewMKDirectionsResponse() MKDirectionsResponse {
	return getMKDirectionsResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKDirectionsResponse */
// The route information that Apple servers return in response to your request for directions.
//
// You don’t create instances of this class directly. Instead, you initiate a request for directions by calling the method of an object. The completion handler you pass to that method receives an object with the results.


// The route information that Apple servers return in response to your request for directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Response
type MKDirectionsResponse struct {
	objectivec.Object
}

// MKDirectionsResponseFrom constructs a [MKDirectionsResponse] from an unsafe.Pointer.
//
// The route information that Apple servers return in response to your request for directions.
func MKDirectionsResponseFrom(ptr unsafe.Pointer) MKDirectionsResponse {
	return MKDirectionsResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKDirectionsResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKDirectionsResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKDirectionsResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKDirectionsResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKDirectionsResponse */

// The end point of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Response/destination
func (m_ MKDirectionsResponse) Destination() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// An array of route objects representing the directions between the start and end points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Response/routes
func (m_ MKDirectionsResponse) Routes() []MKRoute {
	rv := objc.Send[[]MKRoute](m_.ID, objc.Sel("routes"))
	return rv
}/* debug [instance_properties/getter]: routes */


// The start point of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Response/source
func (m_ MKDirectionsResponse) Source() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKDirectionsResponse */



