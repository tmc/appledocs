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

/* debug [class.gen.go]: Generating class MKUserLocation */


/* debug [class_header]: Header for MKUserLocation */
// The class instance for the [MKUserLocation] class.
var (
	MKUserLocationClass     _MKUserLocationClass
	MKUserLocationClassOnce sync.Once
)

func getMKUserLocationClass() _MKUserLocationClass {
	MKUserLocationClassOnce.Do(func() {
		MKUserLocationClass = _MKUserLocationClass{objc.GetClass("MKUserLocation")}
	})
	return MKUserLocationClass
}

type _MKUserLocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKUserLocation */
// An interface definition for the [MKUserLocation] class.
type IMKUserLocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKUserLocation */
	// properties:
	Heading() corelocation.Heading
	Updating() bool
	Location() corelocation.Location
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UserLocation() IMKUserLocation
	SetUserLocation(value IMKUserLocation)
	IsUpdating() bool
	SetIsUpdating(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKUserLocation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKUserLocation */
// Alloc allocates a new instance without initialization.
func (mc _MKUserLocationClass) Alloc() MKUserLocation {
	rv := objc.Send[MKUserLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKUserLocationClass) New() MKUserLocation {
	rv := objc.Send[MKUserLocation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserLocation) Init() MKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserLocation) Autorelease() MKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserLocation creates a new MKUserLocation instance.
func NewMKUserLocation() MKUserLocation {
	return getMKUserLocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKUserLocation */
// An annotation that reflects the user’s location on the map.
//
// You don’t create instances of this class directly. Instead, you retrieve an existing object from the property of the map view that displays in your app.


// An annotation that reflects the user’s location on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation
type MKUserLocation struct {
	objectivec.Object
}

// MKUserLocationFrom constructs a [MKUserLocation] from an unsafe.Pointer.
//
// An annotation that reflects the user’s location on the map.
func MKUserLocationFrom(ptr unsafe.Pointer) MKUserLocation {
	return MKUserLocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKUserLocation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKUserLocation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKUserLocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKUserLocation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKUserLocation */

// The heading of the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/heading
func (m_ MKUserLocation) Heading() corelocation.Heading {
	rv := objc.Send[corelocation.Heading](m_.ID, objc.Sel("heading"))
	return rv
}/* debug [instance_properties/getter]: heading */


// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/isUpdating
func (m_ MKUserLocation) Updating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("updating"))
	return rv
}/* debug [instance_properties/getter]: updating */


// The location of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/location
func (m_ MKUserLocation) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](m_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The subtitle to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/subtitle
func (m_ MKUserLocation) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The subtitle to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/subtitle
func (m_ MKUserLocation) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// The title to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/title
func (m_ MKUserLocation) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/title
func (m_ MKUserLocation) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) UserLocation() IMKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("userLocation"))
	return rv
}/* debug [instance_properties/getter]: userLocation */


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) SetUserLocation(value IMKUserLocation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserLocation:"), value)
}/* debug [instance_properties/setter]: userLocation */


// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) IsUpdating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdating"))
	return rv
}/* debug [instance_properties/getter]: isUpdating */


// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) SetIsUpdating(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUpdating:"), value)
}/* debug [instance_properties/setter]: isUpdating */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKUserLocation */



