// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetCarLockStatusIntent] class.
var (
	INGetCarLockStatusIntentClass     _INGetCarLockStatusIntentClass
	INGetCarLockStatusIntentClassOnce sync.Once
)

func getINGetCarLockStatusIntentClass() _INGetCarLockStatusIntentClass {
	INGetCarLockStatusIntentClassOnce.Do(func() {
		INGetCarLockStatusIntentClass = _INGetCarLockStatusIntentClass{objc.GetClass("INGetCarLockStatusIntent")}
	})
	return INGetCarLockStatusIntentClass
}

type _INGetCarLockStatusIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetCarLockStatusIntent] class.
type IINGetCarLockStatusIntent interface {
	IINIntent
	CarName() INSpeakableString
	SetCarName(value INSpeakableString)
}

// A request to get the lock status of the user’s car.
//
// When the user asks for the car’s lock status, Siri creates an object . This intent object can contain the name of the user’s car. Use this object to provide information about the car’s current lock status. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetCarLockStatusIntent
type INGetCarLockStatusIntent struct {
	INIntent
}

// INGetCarLockStatusIntentFrom constructs a [INGetCarLockStatusIntent] from an unsafe.Pointer.
//
// A request to get the lock status of the user’s car.
func INGetCarLockStatusIntentFrom(ptr unsafe.Pointer) INGetCarLockStatusIntent {
	return INGetCarLockStatusIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetCarLockStatusIntentClass) Alloc() INGetCarLockStatusIntent {
	rv := objc.Send[INGetCarLockStatusIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetCarLockStatusIntentClass) New() INGetCarLockStatusIntent {
	rv := objc.Send[INGetCarLockStatusIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetCarLockStatusIntent) Init() INGetCarLockStatusIntent {
	rv := objc.Send[INGetCarLockStatusIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetCarLockStatusIntent) Autorelease() INGetCarLockStatusIntent {
	rv := objc.Send[INGetCarLockStatusIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetCarLockStatusIntent creates a new INGetCarLockStatusIntent instance.
func NewINGetCarLockStatusIntent() INGetCarLockStatusIntent {
	return getINGetCarLockStatusIntentClass().New()
}


// A name that identifies the user’s car.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetcarlockstatusintent/carname
func (i_ INGetCarLockStatusIntent) CarName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("carName"))
	return rv
}


// SetCarName sets the value of the carName property.
// A name that identifies the user’s car.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetcarlockstatusintent/carname
func (i_ INGetCarLockStatusIntent) SetCarName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}



