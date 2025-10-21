// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




