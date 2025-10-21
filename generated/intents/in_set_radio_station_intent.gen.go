// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetRadioStationIntent] class.
var (
	INSetRadioStationIntentClass     _INSetRadioStationIntentClass
	INSetRadioStationIntentClassOnce sync.Once
)

func getINSetRadioStationIntentClass() _INSetRadioStationIntentClass {
	INSetRadioStationIntentClassOnce.Do(func() {
		INSetRadioStationIntentClass = _INSetRadioStationIntentClass{objc.GetClass("INSetRadioStationIntent")}
	})
	return INSetRadioStationIntentClass
}

type _INSetRadioStationIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetRadioStationIntent] class.
type IINSetRadioStationIntent interface {
	IINIntent
}

// A request to change the current radio station.
//
// When the user asks Siri to change the current radio station, SiriKit creates an object and delivers it to the app’s Intents extension. You use the intent to identify which radio station the user wants. Automotive vendors can use this intent to change the settings on a vehicle’s built-in entertainment system. This properties of this class support identifying a radio station in several different ways, but a given instance of this class contains doesn’t populate all of those properties. When resolving and confirming the parameters of this intent, use the properties that are available to change the station. The object that handles this intent must adopt the protocol. Use this intent object to identify the selected station and to create an object indicating the results of changing the station.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetRadioStationIntent
type INSetRadioStationIntent struct {
	INIntent
}

// INSetRadioStationIntentFrom constructs a [INSetRadioStationIntent] from an unsafe.Pointer.
//
// A request to change the current radio station.
func INSetRadioStationIntentFrom(ptr unsafe.Pointer) INSetRadioStationIntent {
	return INSetRadioStationIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetRadioStationIntentClass) Alloc() INSetRadioStationIntent {
	rv := objc.Send[INSetRadioStationIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetRadioStationIntentClass) New() INSetRadioStationIntent {
	rv := objc.Send[INSetRadioStationIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetRadioStationIntent) Init() INSetRadioStationIntent {
	rv := objc.Send[INSetRadioStationIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetRadioStationIntent) Autorelease() INSetRadioStationIntent {
	rv := objc.Send[INSetRadioStationIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetRadioStationIntent creates a new INSetRadioStationIntent instance.
func NewINSetRadioStationIntent() INSetRadioStationIntent {
	return getINSetRadioStationIntentClass().New()
}




