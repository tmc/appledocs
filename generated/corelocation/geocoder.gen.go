// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Geocoder] class.
var geocoderClass = _GeocoderClass{objc.GetClass("CLGeocoder")}

type _GeocoderClass struct {
	class objc.Class
}

// An interface for converting between geographic coordinates and place names. [Full Topic]
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

// Cancels a pending geocoding request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/cancelGeocode()
func (g_ Geocoder) CancelGeocode() {
	objc.Send[objc.ID](g_.ID, objc.Sel("cancelGeocode"))
}
// Submits a forward-geocoding request using the specified address dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressDictionary(_:completionHandler:)
func (g_ Geocoder) GeocodeAddressDictionaryCompletionHandler(addressDictionary unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressDictionary:completionHandler:"), addressDictionary, completionHandler)
}
// Submits a forward-geocoding request using the specified string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringCompletionHandler(addressString string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:completionHandler:"), addressString, completionHandler)
}
// Submits a forward-geocoding request using the specified string and region information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:in:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionCompletionHandler(addressString string, region unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegion:completionHandler:"), addressString, region, completionHandler)
}
// Submits a forward-geocoding requesting using the specified address string and locale information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:in:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionPreferredLocaleCompletionHandler(addressString string, region unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegion:preferredLocale:completionHandler:"), addressString, region, locale, completionHandler)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodeAddressString(_:inRegionCenteredAt:inRegionRadius:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodeAddressStringInRegionCenteredAtInRegionRadiusPreferredLocaleCompletionHandler(addressString string, centroid unsafe.Pointer, radius unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodeAddressString:inRegionCenteredAt:inRegionRadius:preferredLocale:completionHandler:"), addressString, centroid, radius, locale, completionHandler)
}
// Submits a forward-geocoding requesting using the specified Contacts framework information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodePostalAddress(_:completionHandler:)
func (g_ Geocoder) GeocodePostalAddressCompletionHandler(postalAddress unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodePostalAddress:completionHandler:"), postalAddress, completionHandler)
}
// Submits a forward-geocoding requesting using the specified locale and Contacts framework information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/geocodePostalAddress(_:preferredLocale:completionHandler:)
func (g_ Geocoder) GeocodePostalAddressPreferredLocaleCompletionHandler(postalAddress unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("geocodePostalAddress:preferredLocale:completionHandler:"), postalAddress, locale, completionHandler)
}
// Submits a reverse-geocoding request for the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/reverseGeocodeLocation(_:completionHandler:)
func (g_ Geocoder) ReverseGeocodeLocationCompletionHandler(location unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("reverseGeocodeLocation:completionHandler:"), location, completionHandler)
}
// Submits a reverse-geocoding request for the specified location and locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/reverseGeocodeLocation(_:preferredLocale:completionHandler:)
func (g_ Geocoder) ReverseGeocodeLocationPreferredLocaleCompletionHandler(location unsafe.Pointer, locale unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("reverseGeocodeLocation:preferredLocale:completionHandler:"), location, locale, completionHandler)
}


