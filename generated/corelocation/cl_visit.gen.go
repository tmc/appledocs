// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Visit] class.
var (
	VisitClass     _VisitClass
	VisitClassOnce sync.Once
)

func getVisitClass() _VisitClass {
	VisitClassOnce.Do(func() {
		VisitClass = _VisitClass{objc.GetClass("CLVisit")}
	})
	return VisitClass
}

type _VisitClass struct {
	class objc.Class
}





// An interface definition for the [Visit] class.
type IVisit interface {
	objectivec.IObject
	

	// properties:
	ArrivalDate() foundation.foundation.INSDate
	Coordinate() CLLocationCoordinate2D
	DepartureDate() foundation.foundation.INSDate
	HorizontalAccuracy() LocationAccuracy /* not a class type */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VisitClass) Alloc() Visit {
	rv := objc.Send[Visit](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VisitClass) New() Visit {
	rv := objc.Send[Visit](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Visit) Init() Visit {
	rv := objc.Send[Visit](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Visit) Autorelease() Visit {
	rv := objc.Send[Visit](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisit creates a new Visit instance.
func NewVisit() Visit {
	return getVisitClass().New()
}





// Information about the user’s location during a specific period of time.
//
// A object encapsulates information about places that the user has been. Visit objects are created by the system and delivered by the object to its delegate after you start the delivery of events. The visit includes the location where the visit occurred and information about the arrival and departure times as relevant. You do not create visit objects directly, nor should you subclass . Visit objects contain as much information about the visit as possible but may not always include both the arrival and departure times. For example, when the user arrives at a location, the system may send an event with only an arrival time. When the user departs a location, the event can contain both the arrival time (if your app was monitoring visits prior to the user’s arrival) and the departure time.


// Information about the user’s location during a specific period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit
type Visit struct {
	objectivec.Object
}

// VisitFrom constructs a [Visit] from an unsafe.Pointer.
//
// Information about the user’s location during a specific period of time.
func VisitFrom(ptr unsafe.Pointer) Visit {
	return Visit{objectivec.Object{objc.ID(ptr)}}
}

























// The approximate time at which the user arrived at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit/arrivalDate
func (v_ Visit) ArrivalDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](v_.ID, objc.Sel("arrivalDate"))
	return rv
}


// The geographical coordinate information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit/coordinate
func (v_ Visit) Coordinate() CLLocationCoordinate2D {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("coordinate"))
	return rv
}


// The approximate time at which the user left the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit/departureDate
func (v_ Visit) DepartureDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](v_.ID, objc.Sel("departureDate"))
	return rv
}


// The horizontal accuracy (in meters) of the specified coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit/horizontalAccuracy
func (v_ Visit) HorizontalAccuracy() LocationAccuracy /* not a class type */ {
	rv := objc.Send[LocationAccuracy](v_.ID, objc.Sel("horizontalAccuracy"))
	return rv
}








