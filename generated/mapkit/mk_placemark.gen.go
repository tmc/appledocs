// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPlacemark */


/* debug [class_header]: Header for MKPlacemark */
// The class instance for the [MKPlacemark] class.
var (
	MKPlacemarkClass     _MKPlacemarkClass
	MKPlacemarkClassOnce sync.Once
)

func getMKPlacemarkClass() _MKPlacemarkClass {
	MKPlacemarkClassOnce.Do(func() {
		MKPlacemarkClass = _MKPlacemarkClass{objc.GetClass("MKPlacemark")}
	})
	return MKPlacemarkClass
}

type _MKPlacemarkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPlacemark */
// An interface definition for the [MKPlacemark] class.
type IMKPlacemark interface {
	IPlacemark
	
/* debug [class_interface_properties]: Properties for MKPlacemark */
	// properties:
	CountryCode() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPlacemark */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPlacemark */
// Alloc allocates a new instance without initialization.
func (mc _MKPlacemarkClass) Alloc() MKPlacemark {
	rv := objc.Send[MKPlacemark](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPlacemarkClass) New() MKPlacemark {
	rv := objc.Send[MKPlacemark](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPlacemark) Init() MKPlacemark {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPlacemark) Autorelease() MKPlacemark {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPlacemark creates a new MKPlacemark instance.
func NewMKPlacemark() MKPlacemark {
	return getMKPlacemarkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPlacemark */
// A user-friendly description of a location on the map.
//
// Placemark data includes information like the country or region, state, city, and street address associated with the specified coordinate. A placemark is a concrete annotation object and conforms to the protocol. Because it’s an annotation, you can add a placemark directly to the map view’s list of annotations.


// A user-friendly description of a location on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark
type MKPlacemark struct {
	Placemark
}

// MKPlacemarkFrom constructs a [MKPlacemark] from an unsafe.Pointer.
//
// A user-friendly description of a location on the map.
func MKPlacemarkFrom(ptr unsafe.Pointer) MKPlacemark {
	return MKPlacemark{
		Placemark: PlacemarkFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPlacemark */

// Creates and returns a placemark object using the specified coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark/init(coordinate:)
func NewMKPlacemarkWithCoordinate(coordinate LocationCoordinate2D /* not a class type */) MKPlacemark {
	instance := getMKPlacemarkClass().Alloc()
	rv := objc.Send[MKPlacemark](instance.ID, objc.Sel("initWithCoordinate:"), coordinate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPlacemarkWithCoordinate */


// Creates and returns a placemark object using the specified coordinate and Address Book dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark/init(coordinate:addressDictionary:)
func NewMKPlacemarkWithCoordinateAddressDictionary(coordinate LocationCoordinate2D /* not a class type */, addressDictionary foundation.IDictionary) MKPlacemark {
	instance := getMKPlacemarkClass().Alloc()
	rv := objc.Send[MKPlacemark](instance.ID, objc.Sel("initWithCoordinate:addressDictionary:"), coordinate, addressDictionary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPlacemarkWithCoordinateAddressDictionary */


// Creates and returns a placemark object with the specified coordinate and postal address from the user’s Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark/init(coordinate:postalAddress:)
func NewMKPlacemarkWithCoordinatePostalAddress(coordinate LocationCoordinate2D /* not a class type */, postalAddress contacts.CNPostalAddress) MKPlacemark {
	instance := getMKPlacemarkClass().Alloc()
	rv := objc.Send[MKPlacemark](instance.ID, objc.Sel("initWithCoordinate:postalAddress:"), coordinate, postalAddress)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPlacemarkWithCoordinatePostalAddress */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPlacemark */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPlacemark */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPlacemark */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPlacemark */

// The abbreviated country or region name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark/countryCode
func (m_ MKPlacemark) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_properties/getter]: countryCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPlacemark */


