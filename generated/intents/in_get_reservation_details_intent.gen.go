// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetReservationDetailsIntent] class.
var (
	INGetReservationDetailsIntentClass     _INGetReservationDetailsIntentClass
	INGetReservationDetailsIntentClassOnce sync.Once
)

func getINGetReservationDetailsIntentClass() _INGetReservationDetailsIntentClass {
	INGetReservationDetailsIntentClassOnce.Do(func() {
		INGetReservationDetailsIntentClass = _INGetReservationDetailsIntentClass{objc.GetClass("INGetReservationDetailsIntent")}
	})
	return INGetReservationDetailsIntentClass
}

type _INGetReservationDetailsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetReservationDetailsIntent] class.
type IINGetReservationDetailsIntent interface {
	IINIntent
}

// A request for details about one or more reservations.
//
// Create an object when the user makes a request to see the details about one or more reservations in your app. The app creates a new object containing the and object and donates it to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetReservationDetailsIntent
type INGetReservationDetailsIntent struct {
	INIntent
}

// INGetReservationDetailsIntentFrom constructs a [INGetReservationDetailsIntent] from an unsafe.Pointer.
//
// A request for details about one or more reservations.
func INGetReservationDetailsIntentFrom(ptr unsafe.Pointer) INGetReservationDetailsIntent {
	return INGetReservationDetailsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetReservationDetailsIntentClass) Alloc() INGetReservationDetailsIntent {
	rv := objc.Send[INGetReservationDetailsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetReservationDetailsIntentClass) New() INGetReservationDetailsIntent {
	rv := objc.Send[INGetReservationDetailsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetReservationDetailsIntent) Init() INGetReservationDetailsIntent {
	rv := objc.Send[INGetReservationDetailsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetReservationDetailsIntent) Autorelease() INGetReservationDetailsIntent {
	rv := objc.Send[INGetReservationDetailsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetReservationDetailsIntent creates a new INGetReservationDetailsIntent instance.
func NewINGetReservationDetailsIntent() INGetReservationDetailsIntent {
	return getINGetReservationDetailsIntentClass().New()
}


// A unique identifier for the array containing the reservation objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetReservationDetailsIntent/reservationContainerReference
func (i_ INGetReservationDetailsIntent) ReservationContainerReference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reservationContainerReference"))
	return rv
}

// An array of unique identifiers for previously created reservations.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetReservationDetailsIntent/reservationItemReferences
func (i_ INGetReservationDetailsIntent) ReservationItemReferences() []INSpeakableString {
	rv := objc.Send[[]INSpeakableString](i_.ID, objc.Sel("reservationItemReferences"))
	return rv
}



