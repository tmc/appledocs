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


// The end point of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/destination
func (m_ MKETAResponse) Destination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destination"))
	return rv
}


// SetDestination sets the value of the destination property.
// The end point of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/destination
func (m_ MKETAResponse) SetDestination(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestination:"), value)
}

// The expected travel distance, in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/distance
func (m_ MKETAResponse) Distance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("distance"))
	return rv
}


// SetDistance sets the value of the distance property.
// The expected travel distance, in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/distance
func (m_ MKETAResponse) SetDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistance:"), value)
}

// The expected arrival time.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expectedarrivaldate
func (m_ MKETAResponse) ExpectedArrivalDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("expectedArrivalDate"))
	return rv
}


// SetExpectedArrivalDate sets the value of the expectedArrivalDate property.
// The expected arrival time.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expectedarrivaldate
func (m_ MKETAResponse) SetExpectedArrivalDate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedArrivalDate:"), value)
}

// The expected departure time.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expecteddeparturedate
func (m_ MKETAResponse) ExpectedDepartureDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("expectedDepartureDate"))
	return rv
}


// SetExpectedDepartureDate sets the value of the expectedDepartureDate property.
// The expected departure time.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expecteddeparturedate
func (m_ MKETAResponse) SetExpectedDepartureDate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedDepartureDate:"), value)
}

// The expected travel time, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expectedtraveltime
func (m_ MKETAResponse) ExpectedTravelTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("expectedTravelTime"))
	return rv
}


// SetExpectedTravelTime sets the value of the expectedTravelTime property.
// The expected travel time, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/expectedtraveltime
func (m_ MKETAResponse) SetExpectedTravelTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedTravelTime:"), value)
}

// The start point of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/source
func (m_ MKETAResponse) Source() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("source"))
	return rv
}


// SetSource sets the value of the source property.
// The start point of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/source
func (m_ MKETAResponse) SetSource(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSource:"), value)
}

// The type of conveyance to use for determining the travel time.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/transporttype
func (m_ MKETAResponse) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transportType"))
	return rv
}


// SetTransportType sets the value of the transportType property.
// The type of conveyance to use for determining the travel time.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkdirections/etaresponse/transporttype
func (m_ MKETAResponse) SetTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransportType:"), value)
}



