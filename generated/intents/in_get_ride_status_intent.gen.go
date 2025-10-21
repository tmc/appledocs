// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetRideStatusIntent] class.
var (
	INGetRideStatusIntentClass     _INGetRideStatusIntentClass
	INGetRideStatusIntentClassOnce sync.Once
)

func getINGetRideStatusIntentClass() _INGetRideStatusIntentClass {
	INGetRideStatusIntentClassOnce.Do(func() {
		INGetRideStatusIntentClass = _INGetRideStatusIntentClass{objc.GetClass("INGetRideStatusIntent")}
	})
	return INGetRideStatusIntentClass
}

type _INGetRideStatusIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetRideStatusIntent] class.
type IINGetRideStatusIntent interface {
	IINIntent
}

// A request for the current status of a previously booked ride.
//
// When there’s a request for the status of a booked ride, SiriKit sends an object to your handler object. Upon receiving this intent, fetch the ride status and return it in your response object. The intent object has no additional parameters. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the status of the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetRideStatusIntent
type INGetRideStatusIntent struct {
	INIntent
}

// INGetRideStatusIntentFrom constructs a [INGetRideStatusIntent] from an unsafe.Pointer.
//
// A request for the current status of a previously booked ride.
func INGetRideStatusIntentFrom(ptr unsafe.Pointer) INGetRideStatusIntent {
	return INGetRideStatusIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetRideStatusIntentClass) Alloc() INGetRideStatusIntent {
	rv := objc.Send[INGetRideStatusIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetRideStatusIntentClass) New() INGetRideStatusIntent {
	rv := objc.Send[INGetRideStatusIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetRideStatusIntent) Init() INGetRideStatusIntent {
	rv := objc.Send[INGetRideStatusIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetRideStatusIntent) Autorelease() INGetRideStatusIntent {
	rv := objc.Send[INGetRideStatusIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetRideStatusIntent creates a new INGetRideStatusIntent instance.
func NewINGetRideStatusIntent() INGetRideStatusIntent {
	return getINGetRideStatusIntentClass().New()
}




