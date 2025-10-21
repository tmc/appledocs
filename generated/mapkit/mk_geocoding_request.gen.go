// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




