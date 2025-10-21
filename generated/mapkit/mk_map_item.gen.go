// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	OpenInMapsWithLaunchOptionsCompletionHandler(launchOptions unsafe.Pointer, completion unsafe.Pointer)
	OpenInMapsWithLaunchOptionsFromSceneCompletionHandler(launchOptions unsafe.Pointer, scene unsafe.Pointer, completion unsafe.Pointer)
}

// A point of interest on the map.
//
// A map item includes a geographic location and any interesting data that might apply to that location, such as the address at that location and the name of a business at that address. You can also create a special object representing the user’s location. Use this class to do the following: Share map-related data with the Maps app. Handle requests for directions that originate from the Maps app. To display information in the Maps app, create an object with the information you want to display and call the method. The Maps app displays that location on the map and shows the information you provide. If you implement a routing app, the Maps app provides two objects representing the start and end points. Use the information in those two objects to plot the route and generate directions.
//
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




// Creates and returns a map item object using the specified location and address objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/init(location:address:)
func NewMKMapItemWithLocationAddress(location unsafe.Pointer, address unsafe.Pointer) MKMapItem {
	instance := getMKMapItemClass().Alloc()
	rv := objc.Send[MKMapItem](instance.ID, objc.Sel("initWithLocation:address:"), location, address)
	rv.Autorelease()
	return rv
}



// Creates and returns a map item object using the specified placemark object.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/init(placemark:)
func NewMKMapItemWithPlacemark(placemark unsafe.Pointer) MKMapItem {
	instance := getMKMapItemClass().Alloc()
	rv := objc.Send[MKMapItem](instance.ID, objc.Sel("initWithPlacemark:"), placemark)
	rv.Autorelease()
	return rv
}


// Opens the Maps app and displays the specified map items.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptions(mapItems unsafe.Pointer, launchOptions unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:"), mapItems, launchOptions)
	return rv
}

// Opens the Maps app using the specified map items and options.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:completionHandler:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptionsCompletionHandler(mapItems unsafe.Pointer, launchOptions unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:completionHandler:"), mapItems, launchOptions, completion)
}

// Opens the Maps app from a particular scene using the specified map items and options.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:from:completionHandler:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptionsFromSceneCompletionHandler(mapItems unsafe.Pointer, launchOptions unsafe.Pointer, scene unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:fromScene:completionHandler:"), mapItems, launchOptions, scene, completion)
}

// Opens the Maps app and displays the map item.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openInMaps(launchOptions:completionHandler:)
func (m_ MKMapItem) OpenInMapsWithLaunchOptionsCompletionHandler(launchOptions unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openInMapsWithLaunchOptions:completionHandler:"), launchOptions, completion)
}

// Opens the Maps app from a particular scene using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openInMaps(launchOptions:from:completionHandler:)
func (m_ MKMapItem) OpenInMapsWithLaunchOptionsFromSceneCompletionHandler(launchOptions unsafe.Pointer, scene unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openInMapsWithLaunchOptions:fromScene:completionHandler:"), launchOptions, scene, completion)
}

// The address representations object that contains various address representations useful for display purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/addressRepresentations
func (m_ MKMapItem) AddressRepresentations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addressRepresentations"))
	return rv
}

// A set of alternative identifiers for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/alternateIdentifiers
func (m_ MKMapItem) AlternateIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alternateIdentifiers"))
	return rv
}

// A unique identifier for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/identifier-swift.property
func (m_ MKMapItem) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/isCurrentLocation
func (m_ MKMapItem) IsCurrentLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCurrentLocation"))
	return rv
}

// The location object.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/location
func (m_ MKMapItem) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("location"))
	return rv
}

// The descriptive name associated with the map item.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/name
func (m_ MKMapItem) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The descriptive name associated with the map item.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/name
func (m_ MKMapItem) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// The phone number associated with a business at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/phoneNumber
func (m_ MKMapItem) PhoneNumber() string {
	rv := objc.Send[string](m_.ID, objc.Sel("phoneNumber"))
	return rv
}


// SetPhoneNumber sets the value of the phoneNumber property.
// The phone number associated with a business at the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/phoneNumber
func (m_ MKMapItem) SetPhoneNumber(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPhoneNumber:"), objc.String(value))
}

// The placemark object containing the location information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/placemark
func (m_ MKMapItem) Placemark() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("placemark"))
	return rv
}

// The point-of-interest category for the map item.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/pointOfInterestCategory
func (m_ MKMapItem) PointOfInterestCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}


// SetPointOfInterestCategory sets the value of the pointOfInterestCategory property.
// The point-of-interest category for the map item.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/pointOfInterestCategory
func (m_ MKMapItem) SetPointOfInterestCategory(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}

// The time zone of the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/timeZone
func (m_ MKMapItem) TimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone of the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/timeZone
func (m_ MKMapItem) SetTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}

// The URL associated with the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/url
func (m_ MKMapItem) Url() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL associated with the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/url
func (m_ MKMapItem) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}

// The address object.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/address
func (m_ MKMapItem) Address() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("address"))
	return rv
}


// SetAddress sets the value of the address property.
// The address object.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/address
func (m_ MKMapItem) SetAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddress:"), value)
}

// A constant that indicates the type of a serialized map item.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemtypeidentifier
func (m_ MKMapItem) MKMapItemTypeIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MKMapItemTypeIdentifier"))
	return rv
}


