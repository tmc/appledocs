// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetCarPowerLevelStatusIntent] class.
var (
	INGetCarPowerLevelStatusIntentClass     _INGetCarPowerLevelStatusIntentClass
	INGetCarPowerLevelStatusIntentClassOnce sync.Once
)

func getINGetCarPowerLevelStatusIntentClass() _INGetCarPowerLevelStatusIntentClass {
	INGetCarPowerLevelStatusIntentClassOnce.Do(func() {
		INGetCarPowerLevelStatusIntentClass = _INGetCarPowerLevelStatusIntentClass{objc.GetClass("INGetCarPowerLevelStatusIntent")}
	})
	return INGetCarPowerLevelStatusIntentClass
}

type _INGetCarPowerLevelStatusIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetCarPowerLevelStatusIntent] class.
type IINGetCarPowerLevelStatusIntent interface {
	IINIntent
}

// A request for the current power level of the user’s car.
//
// When asked for the car’s power level, Siri creates an object. This intent object can contain the name of the user’s car. Use this object to provide information about the car’s current power level. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetCarPowerLevelStatusIntent
type INGetCarPowerLevelStatusIntent struct {
	INIntent
}

// INGetCarPowerLevelStatusIntentFrom constructs a [INGetCarPowerLevelStatusIntent] from an unsafe.Pointer.
//
// A request for the current power level of the user’s car.
func INGetCarPowerLevelStatusIntentFrom(ptr unsafe.Pointer) INGetCarPowerLevelStatusIntent {
	return INGetCarPowerLevelStatusIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetCarPowerLevelStatusIntentClass) Alloc() INGetCarPowerLevelStatusIntent {
	rv := objc.Send[INGetCarPowerLevelStatusIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetCarPowerLevelStatusIntentClass) New() INGetCarPowerLevelStatusIntent {
	rv := objc.Send[INGetCarPowerLevelStatusIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetCarPowerLevelStatusIntent) Init() INGetCarPowerLevelStatusIntent {
	rv := objc.Send[INGetCarPowerLevelStatusIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetCarPowerLevelStatusIntent) Autorelease() INGetCarPowerLevelStatusIntent {
	rv := objc.Send[INGetCarPowerLevelStatusIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetCarPowerLevelStatusIntent creates a new INGetCarPowerLevelStatusIntent instance.
func NewINGetCarPowerLevelStatusIntent() INGetCarPowerLevelStatusIntent {
	return getINGetCarPowerLevelStatusIntentClass().New()
}


// A name that identifies the user’s car.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetcarpowerlevelstatusintent/carname
func (i_ INGetCarPowerLevelStatusIntent) CarName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("carName"))
	return rv
}


// SetCarName sets the value of the carName property.
// A name that identifies the user’s car.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetcarpowerlevelstatusintent/carname
func (i_ INGetCarPowerLevelStatusIntent) SetCarName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}



