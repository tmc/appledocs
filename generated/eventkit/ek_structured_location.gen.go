// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/mapkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKStructuredLocation */


/* debug [class_header]: Header for EKStructuredLocation */
// The class instance for the [EKStructuredLocation] class.
var (
	EKStructuredLocationClass     _EKStructuredLocationClass
	EKStructuredLocationClassOnce sync.Once
)

func getEKStructuredLocationClass() _EKStructuredLocationClass {
	EKStructuredLocationClassOnce.Do(func() {
		EKStructuredLocationClass = _EKStructuredLocationClass{objc.GetClass("EKStructuredLocation")}
	})
	return EKStructuredLocationClass
}

type _EKStructuredLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKStructuredLocation */
// An interface definition for the [EKStructuredLocation] class.
type IEKStructuredLocation interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKStructuredLocation */
	// properties:
	GeoLocation() corelocation.Location
	SetGeoLocation(value corelocation.Location)
	Radius() float64
	SetRadius(value float64)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	StructuredLocation() IEKStructuredLocation
	SetStructuredLocation(value IEKStructuredLocation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKStructuredLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKStructuredLocation */
// Alloc allocates a new instance without initialization.
func (ec _EKStructuredLocationClass) Alloc() EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKStructuredLocationClass) New() EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKStructuredLocation) Init() EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKStructuredLocation) Autorelease() EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKStructuredLocation creates a new EKStructuredLocation instance.
func NewEKStructuredLocation() EKStructuredLocation {
	return getEKStructuredLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKStructuredLocation */
// class that specifies a geofence to activate the alarm of a calendar item.
//
// Use to create a new structured location, then set it to the property of an object.


// class that specifies a geofence to activate the alarm of a calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation
type EKStructuredLocation struct {
	EKObject
}

// EKStructuredLocationFrom constructs a [EKStructuredLocation] from an unsafe.Pointer.
//
// class that specifies a geofence to activate the alarm of a calendar item.
func EKStructuredLocationFrom(ptr unsafe.Pointer) EKStructuredLocation {
	return EKStructuredLocation{
		EKObject: EKObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKStructuredLocation */

// Creates a new structured location with the specified map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(mapItem:)
func NewEKStructuredLocationWithMapItem(mapItem mapkit.MKMapItem) EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(getEKStructuredLocationClass().class), objc.Sel("locationWithMapItem:"), mapItem)
	return rv
}/* debug [class_init_methods/constructor]: NewEKStructuredLocationWithMapItem */


// Creates a new structured location with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(title:)
func NewEKStructuredLocationWithTitle(title objc.IObject /* cross-framework: NSString */) EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(getEKStructuredLocationClass().class), objc.Sel("locationWithTitle:"), title)
	return rv
}/* debug [class_init_methods/constructor]: NewEKStructuredLocationWithTitle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKStructuredLocation */

// Creates a new structured location with the specified map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(mapItem:)
func (ec _EKStructuredLocationClass) LocationWithMapItem(mapItem mapkit.MKMapItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("locationWithMapItem:"), mapItem)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocationWithMapItem) */


// Creates a new structured location with the specified title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(title:)
func (ec _EKStructuredLocationClass) LocationWithTitle(title objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("locationWithTitle:"), title)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocationWithTitle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKStructuredLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKStructuredLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKStructuredLocation */

// The core location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/geoLocation
func (e_ EKStructuredLocation) GeoLocation() corelocation.Location {
	rv := objc.Send[corelocation.Location](e_.ID, objc.Sel("geoLocation"))
	return rv
}/* debug [instance_properties/getter]: geoLocation */


// The core location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/geoLocation
func (e_ EKStructuredLocation) SetGeoLocation(value corelocation.Location) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGeoLocation:"), value)
}/* debug [instance_properties/setter]: geoLocation */


// A minimum distance from the core location that would trigger the alarm or reminder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/radius
func (e_ EKStructuredLocation) Radius() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// A minimum distance from the core location that would trigger the alarm or reminder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/radius
func (e_ EKStructuredLocation) SetRadius(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRadius:"), value)
}/* debug [instance_properties/setter]: radius */


// The title of the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/title
func (e_ EKStructuredLocation) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/title
func (e_ EKStructuredLocation) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekalarm/structuredlocation
func (e_ EKStructuredLocation) StructuredLocation() IEKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("structuredLocation"))
	return rv
}/* debug [instance_properties/getter]: structuredLocation */


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekalarm/structuredlocation
func (e_ EKStructuredLocation) SetStructuredLocation(value IEKStructuredLocation) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}/* debug [instance_properties/setter]: structuredLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKStructuredLocation */


