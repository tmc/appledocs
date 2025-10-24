// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetRestaurantGuestIntent] class.
var (
	INGetRestaurantGuestIntentClass     _INGetRestaurantGuestIntentClass
	INGetRestaurantGuestIntentClassOnce sync.Once
)

func getINGetRestaurantGuestIntentClass() _INGetRestaurantGuestIntentClass {
	INGetRestaurantGuestIntentClassOnce.Do(func() {
		INGetRestaurantGuestIntentClass = _INGetRestaurantGuestIntentClass{objc.GetClass("INGetRestaurantGuestIntent")}
	})
	return INGetRestaurantGuestIntentClass
}

type _INGetRestaurantGuestIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetRestaurantGuestIntent] class.
type IINGetRestaurantGuestIntent interface {
	IINIntent
}

// A request for information about the guest who is making reservations.
//
// An object is a request for information about the person whose name should appear on reservations. Maps sends this intent to your Intents extension when it wants information about the person making reservations. Your response contains the identity of the person making the reservation. The response can also contain information about whether your service allows the user to modify the identity of the guest. For example, you can specify whether your service allows one user to book reservations on behalf of another user. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should create an object with information about the guest and your app’s preferences for modifying that guest’s identity.

// A request for information about the guest who is making reservations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetRestaurantGuestIntent
type INGetRestaurantGuestIntent struct {
	INIntent
}

// INGetRestaurantGuestIntentFrom constructs a [INGetRestaurantGuestIntent] from an unsafe.Pointer.
//
// A request for information about the guest who is making reservations.
func INGetRestaurantGuestIntentFrom(ptr unsafe.Pointer) INGetRestaurantGuestIntent {
	return INGetRestaurantGuestIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetRestaurantGuestIntentClass) Alloc() INGetRestaurantGuestIntent {
	rv := objc.Send[INGetRestaurantGuestIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetRestaurantGuestIntentClass) New() INGetRestaurantGuestIntent {
	rv := objc.Send[INGetRestaurantGuestIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetRestaurantGuestIntent) Init() INGetRestaurantGuestIntent {
	rv := objc.Send[INGetRestaurantGuestIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetRestaurantGuestIntent) Autorelease() INGetRestaurantGuestIntent {
	rv := objc.Send[INGetRestaurantGuestIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetRestaurantGuestIntent creates a new INGetRestaurantGuestIntent instance.
func NewINGetRestaurantGuestIntent() INGetRestaurantGuestIntent {
	return getINGetRestaurantGuestIntentClass().New()
}
