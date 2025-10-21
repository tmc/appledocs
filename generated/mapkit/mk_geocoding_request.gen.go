// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKGeocodingRequest] class.
var (
	MKGeocodingRequestClass     _MKGeocodingRequestClass
	MKGeocodingRequestClassOnce sync.Once
)

func getMKGeocodingRequestClass() _MKGeocodingRequestClass {
	MKGeocodingRequestClassOnce.Do(func() {
		MKGeocodingRequestClass = _MKGeocodingRequestClass{objc.GetClass("MKGeocodingRequest")}
	})
	return MKGeocodingRequestClass
}

type _MKGeocodingRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKGeocodingRequest] class.
type IMKGeocodingRequest interface {
	objectivec.IObject
}

// A class that looks up a geographic coordinate using the provided string.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeocodingRequest
type MKGeocodingRequest struct {
	objectivec.Object
}

// MKGeocodingRequestFrom constructs a [MKGeocodingRequest] from an unsafe.Pointer.
//
// A class that looks up a geographic coordinate using the provided string.
func MKGeocodingRequestFrom(ptr unsafe.Pointer) MKGeocodingRequest {
	return MKGeocodingRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKGeocodingRequestClass) Alloc() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKGeocodingRequestClass) New() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeocodingRequest) Init() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeocodingRequest) Autorelease() MKGeocodingRequest {
	rv := objc.Send[MKGeocodingRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeocodingRequest creates a new MKGeocodingRequest instance.
func NewMKGeocodingRequest() MKGeocodingRequest {
	return getMKGeocodingRequestClass().New()
}


// The string used to initialize the geocoder.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/addressstring
func (m_ MKGeocodingRequest) AddressString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("addressString"))
	return rv
}


// SetAddressString sets the value of the addressString property.
// The string used to initialize the geocoder.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/addressstring
func (m_ MKGeocodingRequest) SetAddressString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressString:"), objc.String(value))
}

// A Boolean value that indicates whether the current geocoding request is in a cancelled state.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/iscancelled
func (m_ MKGeocodingRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}


// SetIsCancelled sets the value of the isCancelled property.
// A Boolean value that indicates whether the current geocoding request is in a cancelled state.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/iscancelled
func (m_ MKGeocodingRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}

// A Boolean value that indicates whether the current geocoding request is in a loading state.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/isloading
func (m_ MKGeocodingRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}


// SetIsLoading sets the value of the isLoading property.
// A Boolean value that indicates whether the current geocoding request is in a loading state.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/isloading
func (m_ MKGeocodingRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}

// A value that indicates the default locale the geocoder should use when processing requests.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/preferredlocale
func (m_ MKGeocodingRequest) PreferredLocale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredLocale"))
	return rv
}


// SetPreferredLocale sets the value of the preferredLocale property.
// A value that indicates the default locale the geocoder should use when processing requests.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/preferredlocale
func (m_ MKGeocodingRequest) SetPreferredLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredLocale:"), value)
}

// The geographic region for the framework to use as the bounds for the request; defaults to a region that covers the whole world.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/region
func (m_ MKGeocodingRequest) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// SetRegion sets the value of the region property.
// The geographic region for the framework to use as the bounds for the request; defaults to a region that covers the whole world.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeocodingrequest/region
func (m_ MKGeocodingRequest) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}



