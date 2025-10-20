// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Geocoder] class.
var (
	geocoderClass     _GeocoderClass
	geocoderClassOnce sync.Once
)

func getGeocoderClass() _GeocoderClass {
	geocoderClassOnce.Do(func() {
		geocoderClass = _GeocoderClass{objc.GetClass("CLGeocoder")}
	})
	return geocoderClass
}

type _GeocoderClass struct {
	class objc.Class
}

// An interface definition for the [Geocoder] class.
type IGeocoder interface {
	objectivec.IObject
	CancelGeocode()
	GeocodeAddressDictionaryCompletionHandler(addressDictionary objc.ID, completionHandler unsafe.Pointer)
	GeocodeAddressStringCompletionHandler(addressString string, completionHandler unsafe.Pointer)
	GeocodeAddressStringInRegionCompletionHandler(addressString string, region unsafe.Pointer, completionHandler unsafe.Pointer)
	GeocodeAddressStringInRegionPreferredLocaleCompletionHandler(addressString string, region unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer)
	GeocodeAddressStringInRegionCenteredAtInRegionRadiusPreferredLocaleCompletionHandler(addressString string, centroid unsafe.Pointer, radius unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer)
	GeocodePostalAddressCompletionHandler(postalAddress unsafe.Pointer, completionHandler unsafe.Pointer)
	GeocodePostalAddressPreferredLocaleCompletionHandler(postalAddress unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer)
	ReverseGeocodeLocationCompletionHandler(location unsafe.Pointer, completionHandler unsafe.Pointer)
	ReverseGeocodeLocationPreferredLocaleCompletionHandler(location unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An interface for converting between geographic coordinates and place names.
//
// The class provides services for converting between a coordinate (specified as a latitude and longitude) and the user-friendly representation of that coordinate. A user-friendly representation of the coordinate typically consists of the street, city, state, and country or region information corresponding to the given location, but it may also contain a relevant point of interest, landmarks, or other identifying information. A geocoder object is a single-shot object that works with a network-based service to look up placemark information for its specified coordinate value. To use a geocoder object, you create it and call one of its forward- or reverse-geocoding methods to begin the request. requests take a latitude and longitude value and find a user-readable address. requests take a user-readable address and find the corresponding latitude and longitude value. Forward-geocoding requests may also return additional information about the specified location, such as a point of interest or building at that location. For both types of request, the results are returned using a object. In the case of forward-geocoding requests, multiple placemark objects may be returned if the provided information yielded multiple possible locations. To make smart decisions about what types of information to return, the geocoder server uses all the information provided to it when processing the request. For example, if the user is moving quickly along a highway, it might return the name of the overall region, and not the name of a small park that the user is passing through.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder
type Geocoder struct {
	objectivec.Object
}

// GeocoderFrom constructs a [Geocoder] from an unsafe.Pointer.
//
// An interface for converting between geographic coordinates and place names.
func GeocoderFrom(ptr unsafe.Pointer) Geocoder {
	return Geocoder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GeocoderClass) Alloc() Geocoder {
	rv := objc.Send[Geocoder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GeocoderClass) New() Geocoder {
	rv := objc.Send[Geocoder](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Geocoder) Init() Geocoder {
	rv := objc.Send[Geocoder](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Geocoder) Autorelease() Geocoder {
	rv := objc.Send[Geocoder](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeocoder creates a new Geocoder instance.
func NewGeocoder() Geocoder {
	return getGeocoderClass().New()
}


// Cancels a pending geocoding request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/cancelGeocode()
func (g_ Geocoder) CancelGeocode() {
	objc.Send[objc.ID](g_.ID, objc.Sel("cancelGeocode"))
}

// Submits a forward-geocoding request using the specified address dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressDictionary(_:completionHandler:)
func (g_ Geocoder) GeocodeAddressDictionaryCompletionHandler(addressDictionary objc.ID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressDictionary:completionHandler:"), addressDictionary, completionHandler)
}

// Submits a forward-geocoding request using the specified string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringCompletionHandler(addressString string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:completionHandler:"), objc.String(addressString), completionHandler)
}

// Submits a forward-geocoding request using the specified string and region information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:in:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionCompletionHandler(addressString string, region unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegion:completionHandler:"), objc.String(addressString), region, completionHandler)
}

// Submits a forward-geocoding requesting using the specified address string and locale information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:in:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionPreferredLocaleCompletionHandler(addressString string, region unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegion:preferredLocale:completionHandler:"), objc.String(addressString), region, locale, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:inRegionCenteredAt:inRegionRadius:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionCenteredAtInRegionRadiusPreferredLocaleCompletionHandler(addressString string, centroid unsafe.Pointer, radius unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegionCenteredAt:inRegionRadius:preferredLocale:completionHandler:"), objc.String(addressString), centroid, radius, locale, completionHandler)
}

// Submits a forward-geocoding requesting using the specified Contacts framework information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodePostalAddress(_:completionHandler:)
func (g_ Geocoder) GeocodePostalAddressCompletionHandler(postalAddress unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodePostalAddress:completionHandler:"), postalAddress, completionHandler)
}

// Submits a forward-geocoding requesting using the specified locale and Contacts framework information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodePostalAddress(_:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodePostalAddressPreferredLocaleCompletionHandler(postalAddress unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodePostalAddress:preferredLocale:completionHandler:"), postalAddress, locale, completionHandler)
}

// Submits a reverse-geocoding request for the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/reverseGeocodeLocation(_:completionHandler:)
func (g_ Geocoder) ReverseGeocodeLocationCompletionHandler(location unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("reverseGeocodeLocation:completionHandler:"), location, completionHandler)
}

// Submits a reverse-geocoding request for the specified location and locale.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/reverseGeocodeLocation(_:preferredLocale:completionHandler:)
func (g_ Geocoder) ReverseGeocodeLocationPreferredLocaleCompletionHandler(location unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("reverseGeocodeLocation:preferredLocale:completionHandler:"), location, locale, completionHandler)
}

// A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/isGeocoding
func (g_ Geocoder) Geocoding() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("geocoding"))
	return rv
}



