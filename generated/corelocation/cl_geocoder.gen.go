// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Geocoder] class.
var (
	GeocoderClass     _GeocoderClass
	GeocoderClassOnce sync.Once
)

func getGeocoderClass() _GeocoderClass {
	GeocoderClassOnce.Do(func() {
		GeocoderClass = _GeocoderClass{objc.GetClass("CLGeocoder")}
	})
	return GeocoderClass
}

type _GeocoderClass struct {
	class objc.Class
}

// An interface definition for the [Geocoder] class.
type IGeocoder interface {
	objectivec.IObject
	Geocoding() bool
	IsGeocoding() bool
	SetIsGeocoding(value bool)
}

// An interface for converting between geographic coordinates and place names.
//
// The class provides services for converting between a coordinate (specified as a latitude and longitude) and the user-friendly representation of that coordinate. A user-friendly representation of the coordinate typically consists of the street, city, state, and country or region information corresponding to the given location, but it may also contain a relevant point of interest, landmarks, or other identifying information. A geocoder object is a single-shot object that works with a network-based service to look up placemark information for its specified coordinate value. To use a geocoder object, you create it and call one of its forward- or reverse-geocoding methods to begin the request. requests take a latitude and longitude value and find a user-readable address. requests take a user-readable address and find the corresponding latitude and longitude value. Forward-geocoding requests may also return additional information about the specified location, such as a point of interest or building at that location. For both types of request, the results are returned using a object. In the case of forward-geocoding requests, multiple placemark objects may be returned if the provided information yielded multiple possible locations. To make smart decisions about what types of information to return, the geocoder server uses all the information provided to it when processing the request. For example, if the user is moving quickly along a highway, it might return the name of the overall region, and not the name of a small park that the user is passing through.


// An interface for converting between geographic coordinates and place names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder
type Geocoder struct {
	objectivec.Object
}

// GeocoderFrom constructs a [Geocoder] from an unsafe.Pointer.
//
// An interface for converting between geographic coordinates and place names.
func GeocoderFrom(ptr unsafe.Pointer) Geocoder {
	return Geocoder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GeocoderClass) Alloc() Geocoder {
	rv := objc.Send[Geocoder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GeocoderClass) New() Geocoder {
	rv := objc.Send[Geocoder](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Geocoder) Init() Geocoder {
	rv := objc.Send[Geocoder](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Geocoder) Autorelease() Geocoder {
	rv := objc.Send[Geocoder](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeocoder creates a new Geocoder instance.
func NewGeocoder() Geocoder {
	return getGeocoderClass().New()
}



// A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLGeocoder/isGeocoding
func (g_ Geocoder) Geocoding() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("geocoding"))
	return rv
}


// A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/isgeocoding
func (g_ Geocoder) IsGeocoding() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isGeocoding"))
	return rv
}


// A Boolean value indicating whether the receiver is in the middle of geocoding its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocoder/isgeocoding
func (g_ Geocoder) SetIsGeocoding(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsGeocoding:"), value)
}



