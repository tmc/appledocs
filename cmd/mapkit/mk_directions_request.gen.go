// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKDirectionsRequest] class.
var (
	MKDirectionsRequestClass     _MKDirectionsRequestClass
	MKDirectionsRequestClassOnce sync.Once
)

func getMKDirectionsRequestClass() _MKDirectionsRequestClass {
	MKDirectionsRequestClassOnce.Do(func() {
		MKDirectionsRequestClass = _MKDirectionsRequestClass{objc.GetClass("MKDirectionsRequest")}
	})
	return MKDirectionsRequestClass
}

type _MKDirectionsRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKDirectionsRequest] class.
type IMKDirectionsRequest interface {
	objectivec.IObject
}

// The start and end points of a route, along with the planned mode of transportation.
//
// You use an object when requesting or providing directions. If your app provides directions, use this class to decode the URL that the Maps app sends to you. If you need to request directions from Apple, pass an instance of this class to an object. For example, an app that provides subway directions might request walking directions to and from relevant subway stations. Prior to iOS 14, for apps that provide directions, you receive direction-related URLs in your app delegate’s method. Upon receiving a URL, call the method of this class to determine whether the URL relates to routing directions. If it does, create an instance of this class using the provided URL and extract the map items associated with the start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/Request
type MKDirectionsRequest struct {
	objectivec.Object
}

// MKDirectionsRequestFrom constructs a [MKDirectionsRequest] from an unsafe.Pointer.
//
// The start and end points of a route, along with the planned mode of transportation.
func MKDirectionsRequestFrom(ptr unsafe.Pointer) MKDirectionsRequest {
	return MKDirectionsRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsRequestClass) Alloc() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKDirectionsRequestClass) New() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDirectionsRequest) Init() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDirectionsRequest) Autorelease() MKDirectionsRequest {
	rv := objc.Send[MKDirectionsRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDirectionsRequest creates a new MKDirectionsRequest instance.
func NewMKDirectionsRequest() MKDirectionsRequest {
	return getMKDirectionsRequestClass().New()
}




