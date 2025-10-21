// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [EKStructuredLocation] class.
type IEKStructuredLocation interface {
	IEKObject
}

// class that specifies a geofence to activate the alarm of a calendar item.
//
// Use to create a new structured location, then set it to the property of an object.
//
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

// Alloc allocates a new instance without initialization.
func (ec _EKStructuredLocationClass) Alloc() EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new structured location with the specified map item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(mapItem:)
func NewEKStructuredLocationWithMapItem(mapItem unsafe.Pointer) EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(getEKStructuredLocationClass().class), objc.Sel("locationWithMapItem:"), mapItem)
	return rv
}



// Creates a new structured location with the specified title.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(title:)
func NewEKStructuredLocationWithTitle(title string) EKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](objc.ID(getEKStructuredLocationClass().class), objc.Sel("locationWithTitle:"), objc.String(title))
	return rv
}


// Creates a new structured location with the specified map item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(mapItem:)
func (ec _EKStructuredLocationClass) LocationWithMapItem(mapItem unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("locationWithMapItem:"), mapItem)
	return rv
}

// Creates a new structured location with the specified title.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/init(title:)
func (ec _EKStructuredLocationClass) LocationWithTitle(title string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("locationWithTitle:"), objc.String(title))
	return rv
}

// The core location.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/geoLocation
func (e_ EKStructuredLocation) GeoLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("geoLocation"))
	return rv
}


// SetGeoLocation sets the value of the geoLocation property.
// The core location.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/geoLocation
func (e_ EKStructuredLocation) SetGeoLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGeoLocation:"), value)
}

// A minimum distance from the core location that would trigger the alarm or reminder.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/radius
func (e_ EKStructuredLocation) Radius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("radius"))
	return rv
}


// SetRadius sets the value of the radius property.
// A minimum distance from the core location that would trigger the alarm or reminder.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/radius
func (e_ EKStructuredLocation) SetRadius(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRadius:"), value)
}

// The title of the location.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/title
func (e_ EKStructuredLocation) Title() string {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the location.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKStructuredLocation/title
func (e_ EKStructuredLocation) SetTitle(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), objc.String(value))
}


