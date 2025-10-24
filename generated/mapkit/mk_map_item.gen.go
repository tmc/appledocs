// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItem */


/* debug [class_header]: Header for MKMapItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItem */
// An interface definition for the [MKMapItem] class.
type IMKMapItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapItem */
	// properties:
	Address() IMKAddress
	AddressRepresentations() IMKAddressRepresentations
	AlternateIdentifiers() unsafe.Pointer
	Identifier() IMKMapItemIdentifier
	IsCurrentLocation() bool
	Location() corelocation.Location
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PhoneNumber() objc.IObject /* cross-framework: NSString */
	SetPhoneNumber(value objc.IObject /* cross-framework: NSString */)
	Placemark() IMKPlacemark
	PointOfInterestCategory() MKPointOfInterestCategory /* typedef */
	SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
	Url() objc.IObject /* cross-framework: NSURL */
	SetUrl(value objc.IObject /* cross-framework: NSURL */)
	MKMapItemTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItem */
	// methods:
	OpenInMapsWithLaunchOptions(launchOptions foundation.IDictionary) bool
	OpenInMapsWithLaunchOptionsCompletionHandler(launchOptions foundation.IDictionary, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItem */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemClass) Alloc() MKMapItem {
	rv := objc.Send[MKMapItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItem */

// Creates and returns a map item object using the specified location and address objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/init(location:address:)
func NewMKMapItemWithLocationAddress(location corelocation.Location, address IMKAddress) MKMapItem {
	instance := getMKMapItemClass().Alloc()
	rv := objc.Send[MKMapItem](instance.ID, objc.Sel("initWithLocation:address:"), location, address)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemWithLocationAddress */


// Creates and returns a map item object using the specified placemark object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/init(placemark:)
func NewMKMapItemWithPlacemark(placemark IMKPlacemark) MKMapItem {
	instance := getMKMapItemClass().Alloc()
	rv := objc.Send[MKMapItem](instance.ID, objc.Sel("initWithPlacemark:"), placemark)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemWithPlacemark */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItem */

// Creates and returns a singleton map item object representing the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/forCurrentLocation()
func (mc _MKMapItemClass) MapItemForCurrentLocation() MKMapItem {
	rv := objc.Send[MKMapItem](objc.ID(mc.class), objc.Sel("mapItemForCurrentLocation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapItemForCurrentLocation) */


// Opens the Maps app and displays the specified map items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptions(mapItems []MKMapItem, launchOptions foundation.IDictionary) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:"), mapItems, launchOptions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenMapsWithItemsLaunchOptions) */


// Opens the Maps app using the specified map items and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:completionHandler:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptionsCompletionHandler(mapItems []MKMapItem, launchOptions foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:completionHandler:"), mapItems, launchOptions, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenMapsWithItemsLaunchOptionsCompletionHandler) */


// Opens the Maps app from a particular scene using the specified map items and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openMaps(with:launchOptions:from:completionHandler:)
func (mc _MKMapItemClass) OpenMapsWithItemsLaunchOptionsFromSceneCompletionHandler(mapItems []MKMapItem, launchOptions foundation.IDictionary, scene gameplaykit.Scene, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("openMapsWithItems:launchOptions:fromScene:completionHandler:"), mapItems, launchOptions, scene, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenMapsWithItemsLaunchOptionsFromSceneCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItem */

// Opens the Maps app and displays the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openInMaps(launchOptions:)
func (m_ MKMapItem) OpenInMapsWithLaunchOptions(launchOptions foundation.IDictionary) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("openInMapsWithLaunchOptions:"), launchOptions)
	return rv
}/* debug [instance_methods/method]: OpenInMapsWithLaunchOptions */


// Opens the Maps app and displays the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/openInMaps(launchOptions:completionHandler:)
func (m_ MKMapItem) OpenInMapsWithLaunchOptionsCompletionHandler(launchOptions foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openInMapsWithLaunchOptions:completionHandler:"), launchOptions, completion)
}/* debug [instance_methods/method]: OpenInMapsWithLaunchOptionsCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItem */

// The address object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/address
func (m_ MKMapItem) Address() IMKAddress {
	rv := objc.Send[MKAddress](m_.ID, objc.Sel("address"))
	return rv
}/* debug [instance_properties/getter]: address */


// The address representations object that contains various address representations useful for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/addressRepresentations
func (m_ MKMapItem) AddressRepresentations() IMKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](m_.ID, objc.Sel("addressRepresentations"))
	return rv
}/* debug [instance_properties/getter]: addressRepresentations */


// A set of alternative identifiers for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/alternateIdentifiers
func (m_ MKMapItem) AlternateIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alternateIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: alternateIdentifiers */


// A unique identifier for a place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/identifier-swift.property
func (m_ MKMapItem) Identifier() IMKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/isCurrentLocation
func (m_ MKMapItem) IsCurrentLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCurrentLocation"))
	return rv
}/* debug [instance_properties/getter]: isCurrentLocation */


// The location object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/location
func (m_ MKMapItem) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](m_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/name
func (m_ MKMapItem) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The descriptive name associated with the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/name
func (m_ MKMapItem) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/phoneNumber
func (m_ MKMapItem) PhoneNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("phoneNumber"))
	return rv
}/* debug [instance_properties/getter]: phoneNumber */


// The phone number associated with a business at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/phoneNumber
func (m_ MKMapItem) SetPhoneNumber(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPhoneNumber:"), value)
}/* debug [instance_properties/setter]: phoneNumber */


// The placemark object containing the location information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/placemark
func (m_ MKMapItem) Placemark() IMKPlacemark {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("placemark"))
	return rv
}/* debug [instance_properties/getter]: placemark */


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/pointOfInterestCategory
func (m_ MKMapItem) PointOfInterestCategory() MKPointOfInterestCategory /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestCategory */


// The point-of-interest category for the map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/pointOfInterestCategory
func (m_ MKMapItem) SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}/* debug [instance_properties/setter]: pointOfInterestCategory */


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/timeZone
func (m_ MKMapItem) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](m_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone of the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/timeZone
func (m_ MKMapItem) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/url
func (m_ MKMapItem) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL associated with the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/url
func (m_ MKMapItem) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */


// A constant that indicates the type of a serialized map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemtypeidentifier
func (m_ MKMapItem) MKMapItemTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapItemTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: MKMapItemTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItem */


