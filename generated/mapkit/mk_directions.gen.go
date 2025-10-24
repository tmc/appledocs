// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKDirections */


/* debug [class_header]: Header for MKDirections */
// The class instance for the [MKDirections] class.
var (
	MKDirectionsClass     _MKDirectionsClass
	MKDirectionsClassOnce sync.Once
)

func getMKDirectionsClass() _MKDirectionsClass {
	MKDirectionsClassOnce.Do(func() {
		MKDirectionsClass = _MKDirectionsClass{objc.GetClass("MKDirections")}
	})
	return MKDirectionsClass
}

type _MKDirectionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKDirections */
// An interface definition for the [MKDirections] class.
type IMKDirections interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKDirections */
	// properties:
	Calculating() bool
	IsCalculating() bool
	SetIsCalculating(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKDirections */
	// methods:
	CalculateDirectionsWithCompletionHandler(completionHandler objectivec.IObject)
	CalculateETAWithCompletionHandler(completionHandler objectivec.IObject)
	Cancel()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKDirections */
// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsClass) Alloc() MKDirections {
	rv := objc.Send[MKDirections](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKDirectionsClass) New() MKDirections {
	rv := objc.Send[MKDirections](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDirections) Init() MKDirections {
	rv := objc.Send[MKDirections](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDirections) Autorelease() MKDirections {
	rv := objc.Send[MKDirections](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDirections creates a new MKDirections instance.
func NewMKDirections() MKDirections {
	return getMKDirectionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKDirections */
// A utility object that computes directions and travel-time information based on the route information you provide.
//
// You use an object to ask the Apple servers to provide walking or driving directions for a route, which you specify using an object. After making a request, MapKit delivers the results asynchronously to the completion handler that you provide. You can also get the estimated travel time for the route. Each object handles a single request for directions, although you can cancel and restart that request as needed. You can create multiple instances of this class and process different route requests at the same time, but make requests only when you plan to present the corresponding route information to the user. Apps may receive an error if the device makes too many requests in too short a time period.


// A utility object that computes directions and travel-time information based on the route information you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections
type MKDirections struct {
	objectivec.Object
}

// MKDirectionsFrom constructs a [MKDirections] from an unsafe.Pointer.
//
// A utility object that computes directions and travel-time information based on the route information you provide.
func MKDirectionsFrom(ptr unsafe.Pointer) MKDirections {
	return MKDirections{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKDirections */

// Creates and returns a directions object using the specified request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/init(request:)
func NewMKDirectionsWithRequest(request IMKDirectionsRequest) MKDirections {
	instance := getMKDirectionsClass().Alloc()
	rv := objc.Send[MKDirections](instance.ID, objc.Sel("initWithRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKDirectionsWithRequest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKDirections */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKDirections */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKDirections */

// Begins calculating the requested route information asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/calculate(completionHandler:)
func (m_ MKDirections) CalculateDirectionsWithCompletionHandler(completionHandler objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("calculateDirectionsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: CalculateDirectionsWithCompletionHandler */


// Begins calculating the requested travel-time information asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/calculateETA(completionHandler:)
func (m_ MKDirections) CalculateETAWithCompletionHandler(completionHandler objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("calculateETAWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: CalculateETAWithCompletionHandler */


// Cancels a pending request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/cancel()
func (m_ MKDirections) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKDirections */

// A Boolean value that indicates whether a request is in process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/isCalculating
func (m_ MKDirections) Calculating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("calculating"))
	return rv
}/* debug [instance_properties/getter]: calculating */


// A Boolean value that indicates whether a request is in process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/iscalculating
func (m_ MKDirections) IsCalculating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCalculating"))
	return rv
}/* debug [instance_properties/getter]: isCalculating */


// A Boolean value that indicates whether a request is in process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/iscalculating
func (m_ MKDirections) SetIsCalculating(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCalculating:"), value)
}/* debug [instance_properties/setter]: isCalculating */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKDirections */


