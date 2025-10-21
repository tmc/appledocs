// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKReverseGeocodingRequest] class.
var (
	MKReverseGeocodingRequestClass     _MKReverseGeocodingRequestClass
	MKReverseGeocodingRequestClassOnce sync.Once
)

func getMKReverseGeocodingRequestClass() _MKReverseGeocodingRequestClass {
	MKReverseGeocodingRequestClassOnce.Do(func() {
		MKReverseGeocodingRequestClass = _MKReverseGeocodingRequestClass{objc.GetClass("MKReverseGeocodingRequest")}
	})
	return MKReverseGeocodingRequestClass
}

type _MKReverseGeocodingRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKReverseGeocodingRequest] class.
type IMKReverseGeocodingRequest interface {
	objectivec.IObject
}

// A class that looks up address strings for the provided geographic coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocodingRequest
type MKReverseGeocodingRequest struct {
	objectivec.Object
}

// MKReverseGeocodingRequestFrom constructs a [MKReverseGeocodingRequest] from an unsafe.Pointer.
//
// A class that looks up address strings for the provided geographic coordinates.
func MKReverseGeocodingRequestFrom(ptr unsafe.Pointer) MKReverseGeocodingRequest {
	return MKReverseGeocodingRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKReverseGeocodingRequestClass) Alloc() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKReverseGeocodingRequestClass) New() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKReverseGeocodingRequest) Init() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKReverseGeocodingRequest) Autorelease() MKReverseGeocodingRequest {
	rv := objc.Send[MKReverseGeocodingRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKReverseGeocodingRequest creates a new MKReverseGeocodingRequest instance.
func NewMKReverseGeocodingRequest() MKReverseGeocodingRequest {
	return getMKReverseGeocodingRequestClass().New()
}


// A Boolean value that indicates whether the current reverse geocoding request is in a loading state.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/isloading
func (m_ MKReverseGeocodingRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}


// SetIsLoading sets the value of the isLoading property.
// A Boolean value that indicates whether the current reverse geocoding request is in a loading state.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/isloading
func (m_ MKReverseGeocodingRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}

// A value that indicates the preferred locale for the addresses the request returns, or
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/preferredlocale
func (m_ MKReverseGeocodingRequest) PreferredLocale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredLocale"))
	return rv
}


// SetPreferredLocale sets the value of the preferredLocale property.
// A value that indicates the preferred locale for the addresses the request returns, or

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/preferredlocale
func (m_ MKReverseGeocodingRequest) SetPreferredLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredLocale:"), value)
}

// The location provided to the initializer.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/location
func (m_ MKReverseGeocodingRequest) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The location provided to the initializer.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/location
func (m_ MKReverseGeocodingRequest) SetLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}

// A Boolean value that indicates whether the current reverse geocoding request is in a cancelled state.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/iscancelled
func (m_ MKReverseGeocodingRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}


// SetIsCancelled sets the value of the isCancelled property.
// A Boolean value that indicates whether the current reverse geocoding request is in a cancelled state.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkreversegeocodingrequest/iscancelled
func (m_ MKReverseGeocodingRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}



