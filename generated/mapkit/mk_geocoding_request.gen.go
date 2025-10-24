// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKGeocodingRequest */


/* debug [class_header]: Header for MKGeocodingRequest */
// The class instance for the [MKGeocodingRequest] class.
var (
	MKGeocodingRequestClass     _MKGeocodingRequestClass
	MKGeocodingRequestClassOnce sync.Once
)

func getMKGeocodingRequestClass() _MKGeocodingRequestClass {
	MKGeocodingRequestClassOnce.Do(func() {
		MKGeocodingRequestClass = _MKGeocodingRequestClass{objc.GetClass("MKGeocodingRequest")}
	})
	return MKGeocodingRequestClass
}

type _MKGeocodingRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKGeocodingRequest */
// An interface definition for the [MKGeocodingRequest] class.
type IMKGeocodingRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKGeocodingRequest */
	// properties:
	AddressString() objc.IObject /* cross-framework: NSString */
	Cancelled() bool
	Loading() bool
	PreferredLocale() foundation.Locale
	SetPreferredLocale(value foundation.Locale)
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKGeocodingRequest */
	// methods:
	Cancel()
	GetMapItemsWithCompletionHandler(completionHandler func([]unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKGeocodingRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKGeocodingRequestClass) Alloc() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKGeocodingRequestClass) New() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeocodingRequest) Init() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeocodingRequest) Autorelease() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeocodingRequest creates a new MKGeocodingRequest instance.
func NewMKGeocodingRequest() MKGeocodingRequest {
	return getMKGeocodingRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKGeocodingRequest */
// A class that looks up a geographic coordinate using the provided string.
//
// Use this class to look up the coordinate for an address string you provide, for example if you want to display the location in a map. This example shows how to use a modifier on a SwiftUI view to geocode an array of street addresses to the corresponding coordinates that MapKit returns in an array of objects.


// A class that looks up a geographic coordinate using the provided string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest
type MKGeocodingRequest struct {
	objectivec.Object
}

// MKGeocodingRequestFrom constructs a [MKGeocodingRequest] from an unsafe.Pointer.
//
// A class that looks up a geographic coordinate using the provided string.
func MKGeocodingRequestFrom(ptr unsafe.Pointer) MKGeocodingRequest {
	return MKGeocodingRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKGeocodingRequest */

// Initializes a new geocoder request object with the provided address string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/init(addressString:)
func NewMKGeocodingRequestWithAddressString(addressString objc.IObject /* cross-framework: NSString */) MKGeocodingRequest {
	instance := getMKGeocodingRequestClass().Alloc()
	rv := objc.Send[MKGeocodingRequest](instance.ID, objc.Sel("initWithAddressString:"), addressString)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKGeocodingRequestWithAddressString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKGeocodingRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKGeocodingRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKGeocodingRequest */

// A function you call to cancel a geocoding request that’s in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/cancel()
func (m_ MKGeocodingRequest) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Returns the map items relevant to the geocoded location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/getMapItems(completionHandler:)
func (m_ MKGeocodingRequest) GetMapItemsWithCompletionHandler(completionHandler func([]unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMapItemsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetMapItemsWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKGeocodingRequest */

// The string used to initialize the geocoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/addressString
func (m_ MKGeocodingRequest) AddressString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("addressString"))
	return rv
}/* debug [instance_properties/getter]: addressString */


// A Boolean value that indicates whether the current geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/isCancelled
func (m_ MKGeocodingRequest) Cancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates whether the current geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/isLoading
func (m_ MKGeocodingRequest) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// A value that indicates the default locale the geocoder should use when processing requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/preferredLocale
func (m_ MKGeocodingRequest) PreferredLocale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("preferredLocale"))
	return rv
}/* debug [instance_properties/getter]: preferredLocale */


// A value that indicates the default locale the geocoder should use when processing requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/preferredLocale
func (m_ MKGeocodingRequest) SetPreferredLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredLocale:"), value)
}/* debug [instance_properties/setter]: preferredLocale */


// The geographic region for the framework to use as the bounds for the request; defaults to a region that covers the whole world.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/region
func (m_ MKGeocodingRequest) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The geographic region for the framework to use as the bounds for the request; defaults to a region that covers the whole world.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest/region
func (m_ MKGeocodingRequest) SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */


// A Boolean value that indicates whether the current geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/iscancelled
func (m_ MKGeocodingRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates whether the current geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/iscancelled
func (m_ MKGeocodingRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates whether the current geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/isloading
func (m_ MKGeocodingRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether the current geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/isloading
func (m_ MKGeocodingRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKGeocodingRequest */


