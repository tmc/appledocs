// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItemIdentifier */


/* debug [class_header]: Header for MKMapItemIdentifier */
// The class instance for the [MKMapItemIdentifier] class.
var (
	MKMapItemIdentifierClass     _MKMapItemIdentifierClass
	MKMapItemIdentifierClassOnce sync.Once
)

func getMKMapItemIdentifierClass() _MKMapItemIdentifierClass {
	MKMapItemIdentifierClassOnce.Do(func() {
		MKMapItemIdentifierClass = _MKMapItemIdentifierClass{objc.GetClass("MKMapItemIdentifier")}
	})
	return MKMapItemIdentifierClass
}

type _MKMapItemIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItemIdentifier */
// An interface definition for the [MKMapItemIdentifier] class.
type IMKMapItemIdentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapItemIdentifier */
	// properties:
	IdentifierString() objc.IObject /* cross-framework: NSString */
	AlternateIdentifiers() IMKMapItemIdentifier
	SetAlternateIdentifiers(value IMKMapItemIdentifier)
	Identifier() IMKMapItemIdentifier
	SetIdentifier(value IMKMapItemIdentifier)
	IsCurrentLocation() bool
	SetIsCurrentLocation(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PhoneNumber() objc.IObject /* cross-framework: NSString */
	SetPhoneNumber(value objc.IObject /* cross-framework: NSString */)
	Placemark() IMKPlacemark
	SetPlacemark(value IMKPlacemark)
	PointOfInterestCategory() MKPointOfInterestCategory /* typedef */
	SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
	Url() foundation.URL
	SetUrl(value foundation.URL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItemIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItemIdentifier */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemIdentifierClass) Alloc() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapItemIdentifierClass) New() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemIdentifier) Init() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemIdentifier) Autorelease() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemIdentifier creates a new MKMapItemIdentifier instance.
func NewMKMapItemIdentifier() MKMapItemIdentifier {
	return getMKMapItemIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItemIdentifier */
// A unique identifier for a place.


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/Identifier-swift.class
type MKMapItemIdentifier struct {
	objectivec.Object
}

// MKMapItemIdentifierFrom constructs a [MKMapItemIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a place.
func MKMapItemIdentifierFrom(ptr unsafe.Pointer) MKMapItemIdentifier {
	return MKMapItemIdentifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItemIdentifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemIdentifier/initWithIdentifierString:
func NewMKMapItemIdentifierWithIdentifierString(string_ objc.IObject /* cross-framework: NSString */) MKMapItemIdentifier {
	instance := getMKMapItemIdentifierClass().Alloc()
	rv := objc.Send[MKMapItemIdentifier](instance.ID, objc.Sel("initWithIdentifierString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemIdentifierWithIdentifierString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItemIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItemIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItemIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItemIdentifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemIdentifier/identifierString
func (m_ MKMapItemIdentifier) IdentifierString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifierString"))
	return rv
}/* debug [instance_properties/getter]: identifierString */


// A set of alternative identifiers for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/alternateidentifiers
func (m_ MKMapItemIdentifier) AlternateIdentifiers() IMKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("alternateIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: alternateIdentifiers */


// A set of alternative identifiers for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/alternateidentifiers
func (m_ MKMapItemIdentifier) SetAlternateIdentifiers(value IMKMapItemIdentifier) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlternateIdentifiers:"), value)
}/* debug [instance_properties/setter]: alternateIdentifiers */


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItemIdentifier) Identifier() IMKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItemIdentifier) SetIdentifier(value IMKMapItemIdentifier) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItemIdentifier) IsCurrentLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCurrentLocation"))
	return rv
}/* debug [instance_properties/getter]: isCurrentLocation */


// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItemIdentifier) SetIsCurrentLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCurrentLocation:"), value)
}/* debug [instance_properties/setter]: isCurrentLocation */


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItemIdentifier) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItemIdentifier) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItemIdentifier) PhoneNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("phoneNumber"))
	return rv
}/* debug [instance_properties/getter]: phoneNumber */


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItemIdentifier) SetPhoneNumber(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPhoneNumber:"), value)
}/* debug [instance_properties/setter]: phoneNumber */


// The placemark object containing the location information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItemIdentifier) Placemark() IMKPlacemark {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("placemark"))
	return rv
}/* debug [instance_properties/getter]: placemark */


// The placemark object containing the location information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItemIdentifier) SetPlacemark(value IMKPlacemark) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlacemark:"), value)
}/* debug [instance_properties/setter]: placemark */


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItemIdentifier) PointOfInterestCategory() MKPointOfInterestCategory /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestCategory */


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItemIdentifier) SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}/* debug [instance_properties/setter]: pointOfInterestCategory */


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItemIdentifier) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](m_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItemIdentifier) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItemIdentifier) Url() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItemIdentifier) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItemIdentifier */


