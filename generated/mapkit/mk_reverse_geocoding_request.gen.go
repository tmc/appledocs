// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKReverseGeocodingRequest */


/* debug [class_header]: Header for MKReverseGeocodingRequest */
// The class instance for the [MKReverseGeocodingRequest] class.
var (
	MKReverseGeocodingRequestClass     _MKReverseGeocodingRequestClass
	MKReverseGeocodingRequestClassOnce sync.Once
)

func getMKReverseGeocodingRequestClass() _MKReverseGeocodingRequestClass {
	MKReverseGeocodingRequestClassOnce.Do(func() {
		MKReverseGeocodingRequestClass = _MKReverseGeocodingRequestClass{objc.GetClass("MKReverseGeocodingRequest")}
	})
	return MKReverseGeocodingRequestClass
}

type _MKReverseGeocodingRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKReverseGeocodingRequest */
// An interface definition for the [MKReverseGeocodingRequest] class.
type IMKReverseGeocodingRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKReverseGeocodingRequest */
	// properties:
	Cancelled() bool
	Loading() bool
	Location() corelocation.Location
	PreferredLocale() foundation.Locale
	SetPreferredLocale(value foundation.Locale)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKReverseGeocodingRequest */
	// methods:
	Cancel()
	GetMapItemsWithCompletionHandler(completionHandler func([]unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKReverseGeocodingRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKReverseGeocodingRequestClass) Alloc() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKReverseGeocodingRequestClass) New() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKReverseGeocodingRequest) Init() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKReverseGeocodingRequest) Autorelease() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKReverseGeocodingRequest creates a new MKReverseGeocodingRequest instance.
func NewMKReverseGeocodingRequest() MKReverseGeocodingRequest {
	return getMKReverseGeocodingRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKReverseGeocodingRequest */
// A class that looks up address strings for the provided geographic coordinates.
//
// Use this class to look up an address by a coordinate you provide. This example shows how to use a modifier on a SwiftUI view to reverse geocodes an array of coordinates to the corresponding addresses that MapKit returns in an array of objects.


// A class that looks up address strings for the provided geographic coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest
type MKReverseGeocodingRequest struct {
	objectivec.Object
}

// MKReverseGeocodingRequestFrom constructs a [MKReverseGeocodingRequest] from an unsafe.Pointer.
//
// A class that looks up address strings for the provided geographic coordinates.
func MKReverseGeocodingRequestFrom(ptr unsafe.Pointer) MKReverseGeocodingRequest {
	return MKReverseGeocodingRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKReverseGeocodingRequest */

// Initializes a new reverse geocoder request object with the provided location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/init(location:)
func NewMKReverseGeocodingRequestWithLocation(location corelocation.Location) MKReverseGeocodingRequest {
	instance := getMKReverseGeocodingRequestClass().Alloc()
	rv := objc.Send[MKReverseGeocodingRequest](instance.ID, objc.Sel("initWithLocation:"), location)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKReverseGeocodingRequestWithLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKReverseGeocodingRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKReverseGeocodingRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKReverseGeocodingRequest */

// A method you call to cancel a reverse geocoding request that’s in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/cancel()
func (m_ MKReverseGeocodingRequest) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Returns the map items relevant to the reverse geocoded location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/getMapItems(completionHandler:)
func (m_ MKReverseGeocodingRequest) GetMapItemsWithCompletionHandler(completionHandler func([]unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMapItemsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetMapItemsWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKReverseGeocodingRequest */

// A Boolean value that indicates whether the current reverse geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/isCancelled
func (m_ MKReverseGeocodingRequest) Cancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates whether the current reverse geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/isLoading
func (m_ MKReverseGeocodingRequest) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// The location provided to the initializer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/location
func (m_ MKReverseGeocodingRequest) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](m_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// A value that indicates the preferred locale for the addresses the request returns, or if the framework should use the device locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/preferredLocale
func (m_ MKReverseGeocodingRequest) PreferredLocale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("preferredLocale"))
	return rv
}/* debug [instance_properties/getter]: preferredLocale */


// A value that indicates the preferred locale for the addresses the request returns, or if the framework should use the device locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest/preferredLocale
func (m_ MKReverseGeocodingRequest) SetPreferredLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredLocale:"), value)
}/* debug [instance_properties/setter]: preferredLocale */


// A Boolean value that indicates whether the current reverse geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/iscancelled
func (m_ MKReverseGeocodingRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates whether the current reverse geocoding request is in a cancelled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/iscancelled
func (m_ MKReverseGeocodingRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates whether the current reverse geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/isloading
func (m_ MKReverseGeocodingRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether the current reverse geocoding request is in a loading state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/isloading
func (m_ MKReverseGeocodingRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKReverseGeocodingRequest */


