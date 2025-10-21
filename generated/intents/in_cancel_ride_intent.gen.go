// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCancelRideIntent] class.
var (
	INCancelRideIntentClass     _INCancelRideIntentClass
	INCancelRideIntentClassOnce sync.Once
)

func getINCancelRideIntentClass() _INCancelRideIntentClass {
	INCancelRideIntentClassOnce.Do(func() {
		INCancelRideIntentClass = _INCancelRideIntentClass{objc.GetClass("INCancelRideIntent")}
	})
	return INCancelRideIntentClass
}

type _INCancelRideIntentClass struct {
	class objc.Class
}

// An interface definition for the [INCancelRideIntent] class.
type IINCancelRideIntent interface {
	IINIntent
}

// An intent requesting the cancellation of a previously booked ride.
//
// When the user cancels a ride that was previously booked through Siri or Maps, SiriKit sends an object to your handler. SiriKit populates this intent object with the ride identifier that you provided when originally booking the ride. Upon receiving this intent, verify the ride information and cancel the ride accordingly. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the status of the ride. SiriKit prefers sending this intent object to your extension over canceling a ride by other means. So when responding to an object, you can continue to set the property of your response’s object to allow cancellation of the ride in your app. SiriKit uses that object only on systems where this intent is unavailable or not supported by your extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCancelRideIntent
type INCancelRideIntent struct {
	INIntent
}

// INCancelRideIntentFrom constructs a [INCancelRideIntent] from an unsafe.Pointer.
//
// An intent requesting the cancellation of a previously booked ride.
func INCancelRideIntentFrom(ptr unsafe.Pointer) INCancelRideIntent {
	return INCancelRideIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCancelRideIntentClass) Alloc() INCancelRideIntent {
	rv := objc.Send[INCancelRideIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCancelRideIntentClass) New() INCancelRideIntent {
	rv := objc.Send[INCancelRideIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCancelRideIntent) Init() INCancelRideIntent {
	rv := objc.Send[INCancelRideIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCancelRideIntent) Autorelease() INCancelRideIntent {
	rv := objc.Send[INCancelRideIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCancelRideIntent creates a new INCancelRideIntent instance.
func NewINCancelRideIntent() INCancelRideIntent {
	return getINCancelRideIntentClass().New()
}




// Initializes the intent object with the specified ride identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCancelRideIntent/init(rideIdentifier:)
func NewINCancelRideIntentWithRideIdentifier(rideIdentifier string) INCancelRideIntent {
	instance := getINCancelRideIntentClass().Alloc()
	rv := objc.Send[INCancelRideIntent](instance.ID, objc.Sel("initWithRideIdentifier:"), objc.String(rideIdentifier))
	rv.Autorelease()
	return rv
}


// The unique identifier that you assigned to the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCancelRideIntent/rideIdentifier
func (i_ INCancelRideIntent) RideIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("rideIdentifier"))
	return rv
}


