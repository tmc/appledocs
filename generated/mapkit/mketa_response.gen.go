// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKETAResponse] class.
var (
	MKETAResponseClass     _MKETAResponseClass
	MKETAResponseClassOnce sync.Once
)

func getMKETAResponseClass() _MKETAResponseClass {
	MKETAResponseClassOnce.Do(func() {
		MKETAResponseClass = _MKETAResponseClass{objc.GetClass("MKETAResponse")}
	})
	return MKETAResponseClass
}

type _MKETAResponseClass struct {
	class objc.Class
}

// An interface definition for the [MKETAResponse] class.
type IMKETAResponse interface {
	objectivec.IObject
}

// The travel-time information that Apple servers return.
//
// You don’t create instances of this class directly. Instead, you initiate a request for the travel time by calling the method of an object. The completion handler you pass to that method receives an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/ETAResponse
type MKETAResponse struct {
	objectivec.Object
}

// MKETAResponseFrom constructs a [MKETAResponse] from an unsafe.Pointer.
//
// The travel-time information that Apple servers return.
func MKETAResponseFrom(ptr unsafe.Pointer) MKETAResponse {
	return MKETAResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKETAResponseClass) Alloc() MKETAResponse {
	rv := objc.Send[MKETAResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKETAResponseClass) New() MKETAResponse {
	rv := objc.Send[MKETAResponse](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKETAResponse) Init() MKETAResponse {
	rv := objc.Send[MKETAResponse](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKETAResponse) Autorelease() MKETAResponse {
	rv := objc.Send[MKETAResponse](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKETAResponse creates a new MKETAResponse instance.
func NewMKETAResponse() MKETAResponse {
	return getMKETAResponseClass().New()
}




