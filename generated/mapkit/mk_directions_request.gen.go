// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKDirectionsRequest */


/* debug [class_header]: Header for MKDirectionsRequest */
// The class instance for the [MKDirectionsRequest] class.
var (
	MKDirectionsRequestClass     _MKDirectionsRequestClass
	MKDirectionsRequestClassOnce sync.Once
)

func getMKDirectionsRequestClass() _MKDirectionsRequestClass {
	MKDirectionsRequestClassOnce.Do(func() {
		MKDirectionsRequestClass = _MKDirectionsRequestClass{objc.GetClass("MKDirectionsRequest")}
	})
	return MKDirectionsRequestClass
}

type _MKDirectionsRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKDirectionsRequest */
// An interface definition for the [MKDirectionsRequest] class.
type IMKDirectionsRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKDirectionsRequest */
	// properties:
	ArrivalDate() objc.IObject /* cross-framework: NSDate */
	SetArrivalDate(value objc.IObject /* cross-framework: NSDate */)
	DepartureDate() objc.IObject /* cross-framework: NSDate */
	SetDepartureDate(value objc.IObject /* cross-framework: NSDate */)
	Destination() IMKMapItem
	SetDestination(value IMKMapItem)
	HighwayPreference() MKDirectionsRoutePreference
	SetHighwayPreference(value MKDirectionsRoutePreference)
	RequestsAlternateRoutes() bool
	SetRequestsAlternateRoutes(value bool)
	Source() IMKMapItem
	SetSource(value IMKMapItem)
	TollPreference() MKDirectionsRoutePreference
	SetTollPreference(value MKDirectionsRoutePreference)
	TransportType() MKDirectionsTransportType
	SetTransportType(value MKDirectionsTransportType)
	MKLaunchOptionsCameraKey() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeCycling() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeDefault() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeDriving() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeKey() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeTransit() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsDirectionsModeWalking() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsMapCenterKey() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsMapSpanKey() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsMapTypeKey() objc.IObject /* cross-framework: NSString */
	MKLaunchOptionsShowsTrafficKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKDirectionsRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKDirectionsRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsRequestClass) Alloc() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKDirectionsRequestClass) New() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDirectionsRequest) Init() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDirectionsRequest) Autorelease() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDirectionsRequest creates a new MKDirectionsRequest instance.
func NewMKDirectionsRequest() MKDirectionsRequest {
	return getMKDirectionsRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKDirectionsRequest */
// The start and end points of a route, along with the planned mode of transportation.
//
// You use an object when requesting or providing directions. If your app provides directions, use this class to decode the URL that the Maps app sends to you. If you need to request directions from Apple, pass an instance of this class to an object. For example, an app that provides subway directions might request walking directions to and from relevant subway stations. Prior to iOS 14, for apps that provide directions, you receive direction-related URLs in your app delegate’s method. Upon receiving a URL, call the method of this class to determine whether the URL relates to routing directions. If it does, create an instance of this class using the provided URL and extract the map items associated with the start and end points.


// The start and end points of a route, along with the planned mode of transportation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request
type MKDirectionsRequest struct {
	objectivec.Object
}

// MKDirectionsRequestFrom constructs a [MKDirectionsRequest] from an unsafe.Pointer.
//
// The start and end points of a route, along with the planned mode of transportation.
func MKDirectionsRequestFrom(ptr unsafe.Pointer) MKDirectionsRequest {
	return MKDirectionsRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKDirectionsRequest */

// Creates and returns a directions request object using the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/init(contentsOfURL:)
func NewMKDirectionsRequestWithContentsOfURL(url objc.IObject /* cross-framework: NSURL */) MKDirectionsRequest {
	instance := getMKDirectionsRequestClass().Alloc()
	rv := objc.Send[MKDirectionsRequest](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKDirectionsRequestWithContentsOfURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKDirectionsRequest */

// Returns a Boolean value that indicates whether the specified URL contains a directions request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/isDirectionsRequest(_:)
func (mc _MKDirectionsRequestClass) IsDirectionsRequestURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isDirectionsRequestURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsDirectionsRequestURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKDirectionsRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKDirectionsRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKDirectionsRequest */

// The arrival date for the trip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/arrivalDate
func (m_ MKDirectionsRequest) ArrivalDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("arrivalDate"))
	return rv
}/* debug [instance_properties/getter]: arrivalDate */


// The arrival date for the trip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/arrivalDate
func (m_ MKDirectionsRequest) SetArrivalDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArrivalDate:"), value)
}/* debug [instance_properties/setter]: arrivalDate */


// The departure date for the trip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/departureDate
func (m_ MKDirectionsRequest) DepartureDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("departureDate"))
	return rv
}/* debug [instance_properties/getter]: departureDate */


// The departure date for the trip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/departureDate
func (m_ MKDirectionsRequest) SetDepartureDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepartureDate:"), value)
}/* debug [instance_properties/setter]: departureDate */


// The end point for routing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/destination
func (m_ MKDirectionsRequest) Destination() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// The end point for routing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/destination
func (m_ MKDirectionsRequest) SetDestination(value IMKMapItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */


// The value that indicates whether the framework uses or avoids highways when providing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/highwayPreference
func (m_ MKDirectionsRequest) HighwayPreference() MKDirectionsRoutePreference {
	rv := objc.Send[MKDirectionsRoutePreference](m_.ID, objc.Sel("highwayPreference"))
	return rv
}/* debug [instance_properties/getter]: highwayPreference */


// The value that indicates whether the framework uses or avoids highways when providing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/highwayPreference
func (m_ MKDirectionsRequest) SetHighwayPreference(value MKDirectionsRoutePreference) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHighwayPreference:"), value)
}/* debug [instance_properties/setter]: highwayPreference */


// A Boolean value that indicates whether your app requests multiple routes when they’re available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/requestsAlternateRoutes
func (m_ MKDirectionsRequest) RequestsAlternateRoutes() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requestsAlternateRoutes"))
	return rv
}/* debug [instance_properties/getter]: requestsAlternateRoutes */


// A Boolean value that indicates whether your app requests multiple routes when they’re available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/requestsAlternateRoutes
func (m_ MKDirectionsRequest) SetRequestsAlternateRoutes(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestsAlternateRoutes:"), value)
}/* debug [instance_properties/setter]: requestsAlternateRoutes */


// The starting point for routing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/source
func (m_ MKDirectionsRequest) Source() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// The starting point for routing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/source
func (m_ MKDirectionsRequest) SetSource(value IMKMapItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSource:"), value)
}/* debug [instance_properties/setter]: source */


// The value that indicates whether the framework avoids routes that have tolls when providing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/tollPreference
func (m_ MKDirectionsRequest) TollPreference() MKDirectionsRoutePreference {
	rv := objc.Send[MKDirectionsRoutePreference](m_.ID, objc.Sel("tollPreference"))
	return rv
}/* debug [instance_properties/getter]: tollPreference */


// The value that indicates whether the framework avoids routes that have tolls when providing directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/tollPreference
func (m_ MKDirectionsRequest) SetTollPreference(value MKDirectionsRoutePreference) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTollPreference:"), value)
}/* debug [instance_properties/setter]: tollPreference */


// The type of conveyance that the directions apply to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/transportType
func (m_ MKDirectionsRequest) TransportType() MKDirectionsTransportType {
	rv := objc.Send[MKDirectionsTransportType](m_.ID, objc.Sel("transportType"))
	return rv
}/* debug [instance_properties/getter]: transportType */


// The type of conveyance that the directions apply to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request/transportType
func (m_ MKDirectionsRequest) SetTransportType(value MKDirectionsTransportType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransportType:"), value)
}/* debug [instance_properties/setter]: transportType */


// The virtual camera to use for viewing the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionscamerakey
func (m_ MKDirectionsRequest) MKLaunchOptionsCameraKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsCameraKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsCameraKey */


// Cycling directions between the specified start and end points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodecycling
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeCycling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeCycling"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeCycling */


// Directions that match the user’s preferred transportation type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodedefault
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeDefault() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeDefault"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeDefault */


// Driving directions between the specified start and end points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodedriving
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeDriving() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeDriving"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeDriving */


// The mode of transportation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodekey
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeKey */


// Public transit directions between the specified start and end points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodetransit
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeTransit() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeTransit"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeTransit */


// Walking directions between the specified start and end points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodewalking
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeWalking() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeWalking"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsDirectionsModeWalking */


// The coordinate value on which to center the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmapcenterkey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapCenterKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsMapCenterKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsMapCenterKey */


// The amount of the map to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmapspankey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapSpanKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsMapSpanKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsMapSpanKey */


// The type of map (standard, satellite, or hybrid) to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmaptypekey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsMapTypeKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsMapTypeKey */


// A Boolean value that indicates whether to display traffic information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsshowstraffickey
func (m_ MKDirectionsRequest) MKLaunchOptionsShowsTrafficKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKLaunchOptionsShowsTrafficKey"))
	return rv
}/* debug [instance_properties/getter]: MKLaunchOptionsShowsTrafficKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKDirectionsRequest */


