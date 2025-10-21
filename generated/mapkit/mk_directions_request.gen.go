// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKDirectionsRequest] class.
type IMKDirectionsRequest interface {
	objectivec.IObject
}

// The start and end points of a route, along with the planned mode of transportation.
//
// You use an object when requesting or providing directions. If your app provides directions, use this class to decode the URL that the Maps app sends to you. If you need to request directions from Apple, pass an instance of this class to an object. For example, an app that provides subway directions might request walking directions to and from relevant subway stations. Prior to iOS 14, for apps that provide directions, you receive direction-related URLs in your app delegate’s method. Upon receiving a URL, call the method of this class to determine whether the URL relates to routing directions. If it does, create an instance of this class using the provided URL and extract the map items associated with the start and end points.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsRequestClass) Alloc() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The arrival date for the trip.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/arrivaldate
func (m_ MKDirectionsRequest) ArrivalDate() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("arrivalDate"))
	return rv
}


// SetArrivalDate sets the value of the arrivalDate property.
// The arrival date for the trip.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/arrivaldate
func (m_ MKDirectionsRequest) SetArrivalDate(value foundation.IDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArrivalDate:"), value)
}

// The departure date for the trip.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/departuredate
func (m_ MKDirectionsRequest) DepartureDate() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("departureDate"))
	return rv
}


// SetDepartureDate sets the value of the departureDate property.
// The departure date for the trip.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/departuredate
func (m_ MKDirectionsRequest) SetDepartureDate(value foundation.IDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepartureDate:"), value)
}

// The end point for routing directions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/destination
func (m_ MKDirectionsRequest) Destination() MKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("destination"))
	return rv
}


// SetDestination sets the value of the destination property.
// The end point for routing directions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/destination
func (m_ MKDirectionsRequest) SetDestination(value IMKMapItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestination:"), value)
}

// The value that indicates whether the framework uses or avoids highways when providing directions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/highwaypreference
func (m_ MKDirectionsRequest) HighwayPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("highwayPreference"))
	return rv
}


// SetHighwayPreference sets the value of the highwayPreference property.
// The value that indicates whether the framework uses or avoids highways when providing directions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/highwaypreference
func (m_ MKDirectionsRequest) SetHighwayPreference(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHighwayPreference:"), value)
}

// A Boolean value that indicates whether your app requests multiple routes when they’re available.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/requestsalternateroutes
func (m_ MKDirectionsRequest) RequestsAlternateRoutes() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requestsAlternateRoutes"))
	return rv
}


// SetRequestsAlternateRoutes sets the value of the requestsAlternateRoutes property.
// A Boolean value that indicates whether your app requests multiple routes when they’re available.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/requestsalternateroutes
func (m_ MKDirectionsRequest) SetRequestsAlternateRoutes(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestsAlternateRoutes:"), value)
}

// The starting point for routing directions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/source
func (m_ MKDirectionsRequest) Source() MKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("source"))
	return rv
}


// SetSource sets the value of the source property.
// The starting point for routing directions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/source
func (m_ MKDirectionsRequest) SetSource(value IMKMapItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSource:"), value)
}

// The value that indicates whether the framework avoids routes that have tolls when providing directions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/tollpreference
func (m_ MKDirectionsRequest) TollPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tollPreference"))
	return rv
}


// SetTollPreference sets the value of the tollPreference property.
// The value that indicates whether the framework avoids routes that have tolls when providing directions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/tollpreference
func (m_ MKDirectionsRequest) SetTollPreference(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTollPreference:"), value)
}

// The type of conveyance that the directions apply to.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/transporttype
func (m_ MKDirectionsRequest) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transportType"))
	return rv
}


// SetTransportType sets the value of the transportType property.
// The type of conveyance that the directions apply to.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/request/transporttype
func (m_ MKDirectionsRequest) SetTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransportType:"), value)
}

// The virtual camera to use for viewing the map.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionscamerakey
func (m_ MKDirectionsRequest) MKLaunchOptionsCameraKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsCameraKey"))
	return rv
}

// Cycling directions between the specified start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodecycling
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeCycling() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeCycling"))
	return rv
}

// Directions that match the user’s preferred transportation type.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodedefault
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeDefault() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeDefault"))
	return rv
}

// Driving directions between the specified start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodedriving
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeDriving() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeDriving"))
	return rv
}

// The mode of transportation.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodekey
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeKey"))
	return rv
}

// Public transit directions between the specified start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodetransit
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeTransit() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeTransit"))
	return rv
}

// Walking directions between the specified start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsdirectionsmodewalking
func (m_ MKDirectionsRequest) MKLaunchOptionsDirectionsModeWalking() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsDirectionsModeWalking"))
	return rv
}

// The coordinate value on which to center the map.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmapcenterkey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapCenterKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsMapCenterKey"))
	return rv
}

// The amount of the map to display.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmapspankey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapSpanKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsMapSpanKey"))
	return rv
}

// The type of map (standard, satellite, or hybrid) to display.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsmaptypekey
func (m_ MKDirectionsRequest) MKLaunchOptionsMapTypeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsMapTypeKey"))
	return rv
}

// A Boolean value that indicates whether to display traffic information.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklaunchoptionsshowstraffickey
func (m_ MKDirectionsRequest) MKLaunchOptionsShowsTrafficKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("MKLaunchOptionsShowsTrafficKey"))
	return rv
}



