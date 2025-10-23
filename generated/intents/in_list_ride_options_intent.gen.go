// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
)

// The class instance for the [INListRideOptionsIntent] class.
var (
	INListRideOptionsIntentClass     _INListRideOptionsIntentClass
	INListRideOptionsIntentClassOnce sync.Once
)

func getINListRideOptionsIntentClass() _INListRideOptionsIntentClass {
	INListRideOptionsIntentClassOnce.Do(func() {
		INListRideOptionsIntentClass = _INListRideOptionsIntentClass{objc.GetClass("INListRideOptionsIntent")}
	})
	return INListRideOptionsIntentClass
}

type _INListRideOptionsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INListRideOptionsIntent] class.
type IINListRideOptionsIntent interface {
	IINIntent
	// properties:
	DropOffLocation() corelocation.Placemark
	SetDropOffLocation(value corelocation.Placemark)
	PickupLocation() corelocation.Placemark
	SetPickupLocation(value corelocation.Placemark)
	// methods:
}

// An intent for getting the types of rides available from a ride-booking service.
//
// Maps creates an object when it needs to display the types of vehicles that your service offers. Use the information in this intent to identify the possible vehicles that you can provide to the user right now. The intent contains information about the user’s pickup and drop-off locations, which you can use to determine ride availability, pricing, and estimated pickup times. This intent object represents a request for information and is not a commitment from the user to book any of the specified rides. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should resolve and confirm the request parameters and create an object with the list of options.


// An intent for getting the types of rides available from a ride-booking service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INListRideOptionsIntent
type INListRideOptionsIntent struct {
	INIntent
}

// INListRideOptionsIntentFrom constructs a [INListRideOptionsIntent] from an unsafe.Pointer.
//
// An intent for getting the types of rides available from a ride-booking service.
func INListRideOptionsIntentFrom(ptr unsafe.Pointer) INListRideOptionsIntent {
	return INListRideOptionsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INListRideOptionsIntentClass) Alloc() INListRideOptionsIntent {
	rv := objc.Send[INListRideOptionsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INListRideOptionsIntentClass) New() INListRideOptionsIntent {
	rv := objc.Send[INListRideOptionsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INListRideOptionsIntent) Init() INListRideOptionsIntent {
	rv := objc.Send[INListRideOptionsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INListRideOptionsIntent) Autorelease() INListRideOptionsIntent {
	rv := objc.Send[INListRideOptionsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINListRideOptionsIntent creates a new INListRideOptionsIntent instance.
func NewINListRideOptionsIntent() INListRideOptionsIntent {
	return getINListRideOptionsIntentClass().New()
}



// The user’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlistrideoptionsintent/dropofflocation
func (i_ INListRideOptionsIntent) DropOffLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("dropOffLocation"))
	return rv
}


// The user’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlistrideoptionsintent/dropofflocation
func (i_ INListRideOptionsIntent) SetDropOffLocation(value corelocation.Placemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropOffLocation:"), value)
}


// The user’s starting location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlistrideoptionsintent/pickuplocation
func (i_ INListRideOptionsIntent) PickupLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("pickupLocation"))
	return rv
}


// The user’s starting location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlistrideoptionsintent/pickuplocation
func (i_ INListRideOptionsIntent) SetPickupLocation(value corelocation.Placemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPickupLocation:"), value)
}



