// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INRequestRideIntent] class.
var (
	INRequestRideIntentClass     _INRequestRideIntentClass
	INRequestRideIntentClassOnce sync.Once
)

func getINRequestRideIntentClass() _INRequestRideIntentClass {
	INRequestRideIntentClassOnce.Do(func() {
		INRequestRideIntentClass = _INRequestRideIntentClass{objc.GetClass("INRequestRideIntent")}
	})
	return INRequestRideIntentClass
}

type _INRequestRideIntentClass struct {
	class objc.Class
}

// An interface definition for the [INRequestRideIntent] class.
type IINRequestRideIntent interface {
	IINIntent
	PickupLocation() corelocation.Placemark
	DropOffLocation() corelocation.Placemark
	SetDropOffLocation(value corelocation.IPlacemark)
	PartySize() int
	SetPartySize(value int)
	PaymentMethod() unsafe.Pointer
	SetPaymentMethod(value unsafe.Pointer)
	RideOptionName() INSpeakableString
	SetRideOptionName(value INSpeakableString)
	ScheduledPickupTime() INDateComponentsRange
	SetScheduledPickupTime(value INDateComponentsRange)
}

// A request to book the specified ride from your service.
//
// SiriKit creates an object when the user asks to book a ride using your app. A ride request intent contains user-supplied information about the ride, such as its starting point and the number of people. Use this intent object to identify possible ride options for the user to choose from and to book the ride after the user confirms it. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with information about whether your app successfully booked the ride. When implementing ride-booking support, provide a GeoJSON file with the regions for which you are able to provide rides and upload that file as your app’s Routing App Coverage File in App Store Connect. When it needs to suggest apps capable of providing a ride, Maps uses your coverage information to determine whether it should suggest your app. If you do not provide a coverage file and your app is not installed on the user’s device, Maps does not suggest your app. For information about how to create and upload a Routing App Coverage File, see .

// A request to book the specified ride from your service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRequestRideIntent
type INRequestRideIntent struct {
	INIntent
}

// INRequestRideIntentFrom constructs a [INRequestRideIntent] from an unsafe.Pointer.
//
// A request to book the specified ride from your service.
func INRequestRideIntentFrom(ptr unsafe.Pointer) INRequestRideIntent {
	return INRequestRideIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INRequestRideIntentClass) Alloc() INRequestRideIntent {
	rv := objc.Send[INRequestRideIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRequestRideIntentClass) New() INRequestRideIntent {
	rv := objc.Send[INRequestRideIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRequestRideIntent) Init() INRequestRideIntent {
	rv := objc.Send[INRequestRideIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRequestRideIntent) Autorelease() INRequestRideIntent {
	rv := objc.Send[INRequestRideIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRequestRideIntent creates a new INRequestRideIntent instance.
func NewINRequestRideIntent() INRequestRideIntent {
	return getINRequestRideIntentClass().New()
}

// The user’s starting location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRequestRideIntent/pickupLocation
func (i_ INRequestRideIntent) PickupLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("pickupLocation"))
	return rv
}

// The user’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/dropofflocation
func (i_ INRequestRideIntent) DropOffLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("dropOffLocation"))
	return rv
}

// The user’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/dropofflocation
func (i_ INRequestRideIntent) SetDropOffLocation(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropOffLocation:"), value)
}

// The number of passengers that the ride must accommodate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/partysize-39k7z
func (i_ INRequestRideIntent) PartySize() int {
	rv := objc.Send[int](i_.ID, objc.Sel("partySize"))
	return rv
}

// The number of passengers that the ride must accommodate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/partysize-39k7z
func (i_ INRequestRideIntent) SetPartySize(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPartySize:"), value)
}

// The user’s requested payment method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/paymentmethod
func (i_ INRequestRideIntent) PaymentMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("paymentMethod"))
	return rv
}

// The user’s requested payment method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/paymentmethod
func (i_ INRequestRideIntent) SetPaymentMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPaymentMethod:"), value)
}

// The name of the ride option selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/rideoptionname
func (i_ INRequestRideIntent) RideOptionName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("rideOptionName"))
	return rv
}

// The name of the ride option selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/rideoptionname
func (i_ INRequestRideIntent) SetRideOptionName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRideOptionName:"), value)
}

// The time at which to pick up the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/scheduledpickuptime
func (i_ INRequestRideIntent) ScheduledPickupTime() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("scheduledPickupTime"))
	return rv
}

// The time at which to pick up the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrequestrideintent/scheduledpickuptime
func (i_ INRequestRideIntent) SetScheduledPickupTime(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScheduledPickupTime:"), value)
}
