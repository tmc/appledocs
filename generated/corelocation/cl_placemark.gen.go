// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Placemark] class.
var (
	PlacemarkClass     _PlacemarkClass
	PlacemarkClassOnce sync.Once
)

func getPlacemarkClass() _PlacemarkClass {
	PlacemarkClassOnce.Do(func() {
		PlacemarkClass = _PlacemarkClass{objc.GetClass("CLPlacemark")}
	})
	return PlacemarkClass
}

type _PlacemarkClass struct {
	class objc.Class
}





// An interface definition for the [Placemark] class.
type IPlacemark interface {
	objectivec.IObject
	

	// properties:
	AddressDictionary() foundation.foundation.INSDictionary
	AdministrativeArea() foundation.foundation.INSString
	AreasOfInterest() []string
	Country() foundation.foundation.INSString
	InlandWater() foundation.foundation.INSString
	ISOcountryCode() foundation.foundation.INSString
	Locality() foundation.foundation.INSString
	Location() ICLLocation
	Name() foundation.foundation.INSString
	Ocean() foundation.foundation.INSString
	PostalAddress() objectivec.IObject
	PostalCode() foundation.foundation.INSString
	Region() ICLRegion
	SubAdministrativeArea() foundation.foundation.INSString
	SubLocality() foundation.foundation.INSString
	SubThoroughfare() foundation.foundation.INSString
	Thoroughfare() foundation.foundation.INSString
	TimeZone() foundation.TimeZone


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlacemarkClass) Alloc() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlacemarkClass) New() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Placemark) Init() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Placemark) Autorelease() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlacemark creates a new Placemark instance.
func NewPlacemark() Placemark {
	return getPlacemarkClass().New()
}





// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information.
//
// A object stores placemark data for a given latitude and longitude. Placemark data includes information such as the country or region, state, city, and street address associated with the specified coordinate. It can also include points of interest and geographically related data. When you reverse geocode a geographic coordinate using a object, you receive a object containing the descriptive information for that location. You can also create object and fill it with address information yourself, which you might do when you want to determine the geographic coordinate associated with the location.


// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark
type Placemark struct {
	objectivec.Object
}

// PlacemarkFrom constructs a [Placemark] from an unsafe.Pointer.
//
// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information.
func PlacemarkFrom(ptr unsafe.Pointer) Placemark {
	return Placemark{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(location:name:postalAddress:)
func NewPlacemarkWithLocationNamePostalAddress(location ICLLocation, name foundation.foundation.INSString, postalAddress objectivec.IObject) Placemark {
	rv := objc.Send[Placemark](objc.ID(getPlacemarkClass().class), objc.Sel("placemarkWithLocation:name:postalAddress:"), location, name, postalAddress)
	return rv
}


// Initializes and returns a placemark object from another placemark object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(placemark:)
func NewPlacemarkWithPlacemark(placemark ICLPlacemark) Placemark {
	instance := getPlacemarkClass().Alloc()
	rv := objc.Send[Placemark](instance.ID, objc.Sel("initWithPlacemark:"), placemark)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(location:name:postalAddress:)
func (pc _PlacemarkClass) PlacemarkWithLocationNamePostalAddress(location ICLLocation, name foundation.foundation.INSString, postalAddress objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("placemarkWithLocation:name:postalAddress:"), location, name, postalAddress)
	return rv
}

















// A dictionary containing the Address Book keys and values for the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/addressDictionary
func (p_ Placemark) AddressDictionary() foundation.foundation.INSDictionary {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("addressDictionary"))
	return rv
}


// The state or province associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/administrativeArea
func (p_ Placemark) AdministrativeArea() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("administrativeArea"))
	return rv
}


// The relevant areas of interest associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/areasOfInterest
func (p_ Placemark) AreasOfInterest() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("areasOfInterest"))
	return rv
}


// The name of the country or region associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/country
func (p_ Placemark) Country() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("country"))
	return rv
}


// The name of the inland water body associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/inlandWater
func (p_ Placemark) InlandWater() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("inlandWater"))
	return rv
}


// The abbreviated country or region name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/isoCountryCode
func (p_ Placemark) ISOcountryCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("ISOcountryCode"))
	return rv
}


// The city associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/locality
func (p_ Placemark) Locality() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("locality"))
	return rv
}


// The location object containing latitude and longitude information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/location
func (p_ Placemark) Location() ICLLocation {
	rv := objc.Send[Location](p_.ID, objc.Sel("location"))
	return rv
}


// The name of the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/name
func (p_ Placemark) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("name"))
	return rv
}


// The name of the ocean associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/ocean
func (p_ Placemark) Ocean() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("ocean"))
	return rv
}


// The postal address associated with the location, formatted for use with the Contacts framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/postalAddress
func (p_ Placemark) PostalAddress() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("postalAddress"))
	return rv
}


// The postal code associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/postalCode
func (p_ Placemark) PostalCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("postalCode"))
	return rv
}


// The geographic region associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/region
func (p_ Placemark) Region() ICLRegion {
	rv := objc.Send[Region](p_.ID, objc.Sel("region"))
	return rv
}


// Additional administrative area information for the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/subAdministrativeArea
func (p_ Placemark) SubAdministrativeArea() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}


// Additional city-level information for the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/subLocality
func (p_ Placemark) SubLocality() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subLocality"))
	return rv
}


// Additional street-level information for the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/subThoroughfare
func (p_ Placemark) SubThoroughfare() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// The street address associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/thoroughfare
func (p_ Placemark) Thoroughfare() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("thoroughfare"))
	return rv
}


// The time zone associated with the placemark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/timeZone
func (p_ Placemark) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](p_.ID, objc.Sel("timeZone"))
	return rv
}







