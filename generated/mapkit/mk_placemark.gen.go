// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
)

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

// An interface definition for the [MKPlacemark] class.
type IMKPlacemark interface {
	IPlacemark
	CountryCode() string
	SetCountryCode(value string)
}

// A user-friendly description of a location on the map.
//
// Placemark data includes information like the country or region, state, city, and street address associated with the specified coordinate. A placemark is a concrete annotation object and conforms to the protocol. Because it’s an annotation, you can add a placemark directly to the map view’s list of annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark
type MKPlacemark struct {
	corelocation.Placemark
}

// MKPlacemarkFrom constructs a [MKPlacemark] from an unsafe.Pointer.
//
// A user-friendly description of a location on the map.
func MKPlacemarkFrom(ptr unsafe.Pointer) MKPlacemark {
	return MKPlacemark{
		Placemark: corelocation.PlacemarkFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKPlacemarkClass) Alloc() MKPlacemark {
	rv := objc.Send[MKPlacemark](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates and returns a placemark object using the specified coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPlacemark/init(coordinate:)
func NewMKPlacemarkWithCoordinate(coordinate unsafe.Pointer) MKPlacemark {
	instance := getMKPlacemarkClass().Alloc()
	rv := objc.Send[MKPlacemark](instance.ID, objc.Sel("initWithCoordinate:"), coordinate)
	rv.Autorelease()
	return rv
}


// The abbreviated country or region name.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkplacemark/countrycode
func (m_ MKPlacemark) CountryCode() string {
	rv := objc.Send[string](m_.ID, objc.Sel("countryCode"))
	return rv
}


// SetCountryCode sets the value of the countryCode property.
// The abbreviated country or region name.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkplacemark/countrycode
func (m_ MKPlacemark) SetCountryCode(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCountryCode:"), objc.String(value))
}


