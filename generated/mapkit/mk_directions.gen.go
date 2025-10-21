// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKDirections] class.
var (
	MKDirectionsClass     _MKDirectionsClass
	MKDirectionsClassOnce sync.Once
)

func getMKDirectionsClass() _MKDirectionsClass {
	MKDirectionsClassOnce.Do(func() {
		MKDirectionsClass = _MKDirectionsClass{objc.GetClass("MKDirections")}
	})
	return MKDirectionsClass
}

type _MKDirectionsClass struct {
	class objc.Class
}

// An interface definition for the [MKDirections] class.
type IMKDirections interface {
	objectivec.IObject
	CalculateETAWithCompletionHandler(completionHandler unsafe.Pointer)
}

// A utility object that computes directions and travel-time information based on the route information you provide.
//
// You use an object to ask the Apple servers to provide walking or driving directions for a route, which you specify using an object. After making a request, MapKit delivers the results asynchronously to the completion handler that you provide. You can also get the estimated travel time for the route. Each object handles a single request for directions, although you can cancel and restart that request as needed. You can create multiple instances of this class and process different route requests at the same time, but make requests only when you plan to present the corresponding route information to the user. Apps may receive an error if the device makes too many requests in too short a time period.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections
type MKDirections struct {
	objectivec.Object
}

// MKDirectionsFrom constructs a [MKDirections] from an unsafe.Pointer.
//
// A utility object that computes directions and travel-time information based on the route information you provide.
func MKDirectionsFrom(ptr unsafe.Pointer) MKDirections {
	return MKDirections{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKDirectionsClass) Alloc() MKDirections {
	rv := objc.Send[MKDirections](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKDirectionsClass) New() MKDirections {
	rv := objc.Send[MKDirections](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDirections) Init() MKDirections {
	rv := objc.Send[MKDirections](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDirections) Autorelease() MKDirections {
	rv := objc.Send[MKDirections](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDirections creates a new MKDirections instance.
func NewMKDirections() MKDirections {
	return getMKDirectionsClass().New()
}


// Begins calculating the requested travel-time information asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/calculateETA(completionHandler:)
func (m_ MKDirections) CalculateETAWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("calculateETAWithCompletionHandler:"), completionHandler)
}

// A Boolean value that indicates whether a request is in process.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/iscalculating
func (m_ MKDirections) IsCalculating() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCalculating"))
	return rv
}


// SetIsCalculating sets the value of the isCalculating property.
// A Boolean value that indicates whether a request is in process.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/iscalculating
func (m_ MKDirections) SetIsCalculating(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCalculating:"), value)
}



