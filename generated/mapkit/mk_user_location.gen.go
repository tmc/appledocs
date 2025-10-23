// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKUserLocation] class.
type IMKUserLocation interface {
	objectivec.IObject
	// properties:
	Location() corelocation.Location
	UserLocation() IMKUserLocation
	SetUserLocation(value IMKUserLocation)
	Heading() corelocation.Heading
	SetHeading(value corelocation.Heading)
	IsUpdating() bool
	SetIsUpdating(value bool)
	Subtitle() string
	SetSubtitle(value string)
	Title() string
	SetTitle(value string)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MKUserLocationClass) Alloc() MKUserLocation {
	rv := objc.Send[MKUserLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The location of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/location
func (m_ MKUserLocation) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](m_.ID, objc.Sel("location"))
	return rv
}


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) UserLocation() IMKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("userLocation"))
	return rv
}


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) SetUserLocation(value IMKUserLocation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserLocation:"), value)
}


// The heading of the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/heading
func (m_ MKUserLocation) Heading() corelocation.Heading {
	rv := objc.Send[corelocation.Heading](m_.ID, objc.Sel("heading"))
	return rv
}


// The heading of the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/heading
func (m_ MKUserLocation) SetHeading(value corelocation.Heading) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeading:"), value)
}


// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) IsUpdating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdating"))
	return rv
}


// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) SetIsUpdating(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUpdating:"), value)
}


// The subtitle to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/subtitle
func (m_ MKUserLocation) Subtitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subtitle"))
	return rv
}


// The subtitle to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/subtitle
func (m_ MKUserLocation) SetSubtitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}


// The title to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/title
func (m_ MKUserLocation) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// The title to display for the user’s location annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/title
func (m_ MKUserLocation) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}



