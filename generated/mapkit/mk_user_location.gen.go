// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An annotation that reflects the user’s location on the map.
//
// You don’t create instances of this class directly. Instead, you retrieve an existing object from the property of the map view that displays in your app.
//
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


// The annotation object that represents the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) UserLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userLocation"))
	return rv
}


// SetUserLocation sets the value of the userLocation property.
// The annotation object that represents the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKUserLocation) SetUserLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserLocation:"), value)
}

// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) IsUpdating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUpdating"))
	return rv
}


// SetIsUpdating sets the value of the isUpdating property.
// A Boolean value that indicates whether the map view is updating the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkuserlocation/isupdating
func (m_ MKUserLocation) SetIsUpdating(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUpdating:"), value)
}

// The heading of the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/heading
func (m_ MKUserLocation) Heading() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("heading"))
	return rv
}

// A Boolean value that indicates whether the map view is updating the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/isUpdating
func (m_ MKUserLocation) Updating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("updating"))
	return rv
}

// The location of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/location
func (m_ MKUserLocation) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("location"))
	return rv
}

// The subtitle to display for the user’s location annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/subtitle
func (m_ MKUserLocation) Subtitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// The subtitle to display for the user’s location annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/subtitle
func (m_ MKUserLocation) SetSubtitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The title to display for the user’s location annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/title
func (m_ MKUserLocation) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title to display for the user’s location annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocation/title
func (m_ MKUserLocation) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}



