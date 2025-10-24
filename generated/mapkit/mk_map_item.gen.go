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

// The class instance for the [MKMapItem] class.
var (
	MKMapItemClass     _MKMapItemClass
	MKMapItemClassOnce sync.Once
)

func getMKMapItemClass() _MKMapItemClass {
	MKMapItemClassOnce.Do(func() {
		MKMapItemClass = _MKMapItemClass{objc.GetClass("MKMapItem")}
	})
	return MKMapItemClass
}

type _MKMapItemClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItem] class.
type IMKMapItem interface {
	objectivec.IObject
	// properties:
	AlternateIdentifiers() unsafe.Pointer
	Address() IMKAddress
	SetAddress(value IMKAddress)
	AddressRepresentations() IMKAddressRepresentations
	SetAddressRepresentations(value IMKAddressRepresentations)
	Identifier() objc.IObject /* cross-framework: MKMapItemIdentifier */
	SetIdentifier(value objc.IObject /* cross-framework: MKMapItemIdentifier */)
	IsCurrentLocation() bool
	SetIsCurrentLocation(value bool)
	Location() objc.IObject /* cross-framework: Location */
	SetLocation(value objc.IObject /* cross-framework: Location */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PhoneNumber() objc.IObject /* cross-framework: NSString */
	SetPhoneNumber(value objc.IObject /* cross-framework: NSString */)
	Placemark() objc.IObject /* cross-framework: MKPlacemark */
	SetPlacemark(value objc.IObject /* cross-framework: MKPlacemark */)
	PointOfInterestCategory() MKPointOfInterestCategory /* typedef */
	SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */)
	TimeZone() objc.IObject /* cross-framework: TimeZone */
	SetTimeZone(value objc.IObject /* cross-framework: TimeZone */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	MKMapItemTypeIdentifier() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A point of interest on the map.
//
// A map item includes a geographic location and any interesting data that might apply to that location, such as the address at that location and the name of a business at that address. You can also create a special object representing the user’s location. Use this class to do the following: Share map-related data with the Maps app. Handle requests for directions that originate from the Maps app. To display information in the Maps app, create an object with the information you want to display and call the method. The Maps app displays that location on the map and shows the information you provide. If you implement a routing app, the Maps app provides two objects representing the start and end points. Use the information in those two objects to plot the route and generate directions.


// A point of interest on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem
type MKMapItem struct {
	objectivec.Object
}

// MKMapItemFrom constructs a [MKMapItem] from an unsafe.Pointer.
//
// A point of interest on the map.
func MKMapItemFrom(ptr unsafe.Pointer) MKMapItem {
	return MKMapItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemClass) Alloc() MKMapItem {
	rv := objc.Send[MKMapItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemClass) New() MKMapItem {
	rv := objc.Send[MKMapItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItem) Init() MKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItem) Autorelease() MKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItem creates a new MKMapItem instance.
func NewMKMapItem() MKMapItem {
	return getMKMapItemClass().New()
}



// A set of alternative identifiers for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/alternateIdentifiers
func (m_ MKMapItem) AlternateIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alternateIdentifiers"))
	return rv
}


// The address object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/address
func (m_ MKMapItem) Address() IMKAddress {
	rv := objc.Send[MKAddress](m_.ID, objc.Sel("address"))
	return rv
}


// The address object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/address
func (m_ MKMapItem) SetAddress(value IMKAddress) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddress:"), value)
}


// The address representations object that contains various address representations useful for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/addressrepresentations
func (m_ MKMapItem) AddressRepresentations() IMKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](m_.ID, objc.Sel("addressRepresentations"))
	return rv
}


// The address representations object that contains various address representations useful for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/addressrepresentations
func (m_ MKMapItem) SetAddressRepresentations(value IMKAddressRepresentations) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressRepresentations:"), value)
}


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItem) Identifier() objc.IObject /* cross-framework: MKMapItemIdentifier */ {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("identifier"))
	return rv
}


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItem) SetIdentifier(value objc.IObject /* cross-framework: MKMapItemIdentifier */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItem) IsCurrentLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCurrentLocation"))
	return rv
}


// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItem) SetIsCurrentLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCurrentLocation:"), value)
}


// The location object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/location
func (m_ MKMapItem) Location() objc.IObject /* cross-framework: Location */ {
	rv := objc.Send[corelocation.Location](m_.ID, objc.Sel("location"))
	return rv
}


// The location object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/location
func (m_ MKMapItem) SetLocation(value objc.IObject /* cross-framework: Location */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItem) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItem) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItem) PhoneNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("phoneNumber"))
	return rv
}


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItem) SetPhoneNumber(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPhoneNumber:"), value)
}


// The placemark object containing the location information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItem) Placemark() objc.IObject /* cross-framework: MKPlacemark */ {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("placemark"))
	return rv
}


// The placemark object containing the location information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItem) SetPlacemark(value objc.IObject /* cross-framework: MKPlacemark */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlacemark:"), value)
}


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItem) PointOfInterestCategory() MKPointOfInterestCategory /* typedef */ {
	rv := objc.Send[MKPointOfInterestCategory](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItem) SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItem) TimeZone() objc.IObject /* cross-framework: TimeZone */ {
	rv := objc.Send[foundation.TimeZone](m_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItem) SetTimeZone(value objc.IObject /* cross-framework: TimeZone */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItem) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItem) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}


// A constant that indicates the type of a serialized map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemtypeidentifier
func (m_ MKMapItem) MKMapItemTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapItemTypeIdentifier"))
	return rv
}



