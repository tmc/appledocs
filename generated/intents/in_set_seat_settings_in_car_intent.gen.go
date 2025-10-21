// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetSeatSettingsInCarIntent] class.
var (
	INSetSeatSettingsInCarIntentClass     _INSetSeatSettingsInCarIntentClass
	INSetSeatSettingsInCarIntentClassOnce sync.Once
)

func getINSetSeatSettingsInCarIntentClass() _INSetSeatSettingsInCarIntentClass {
	INSetSeatSettingsInCarIntentClassOnce.Do(func() {
		INSetSeatSettingsInCarIntentClass = _INSetSeatSettingsInCarIntentClass{objc.GetClass("INSetSeatSettingsInCarIntent")}
	})
	return INSetSeatSettingsInCarIntentClass
}

type _INSetSeatSettingsInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetSeatSettingsInCarIntent] class.
type IINSetSeatSettingsInCarIntent interface {
	IINIntent
}

// A request to change the seat-related settings in a CarPlay-enabled vehicle.
//
// Automotive venders can add support for this intent to an Intents extension that they ship with their automotive apps. When the user asks Siri to change a seat-related setting, Siri creates an object and delivers it to the app’s Intents extension. You use the intent to identify which setting the user wants to change. This class contains properties for multiple types of seat-related settings, but a given instance of this class contains changes for only one setting at a time. When resolving and confirming the parameters of this intent, use the specified properties to modify your vehicle’s seat settings. Assume no changes for other settings. If your vehicle doesn’t support a particular setting, offer a reasonable fallback setting and ask the user for confirmation. The object that handles this intent must adopt the protocol. Use this intent object to identify which setting changed and to create an object indicating the results of changing the seat settings.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetSeatSettingsInCarIntent
type INSetSeatSettingsInCarIntent struct {
	INIntent
}

// INSetSeatSettingsInCarIntentFrom constructs a [INSetSeatSettingsInCarIntent] from an unsafe.Pointer.
//
// A request to change the seat-related settings in a CarPlay-enabled vehicle.
func INSetSeatSettingsInCarIntentFrom(ptr unsafe.Pointer) INSetSeatSettingsInCarIntent {
	return INSetSeatSettingsInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetSeatSettingsInCarIntentClass) Alloc() INSetSeatSettingsInCarIntent {
	rv := objc.Send[INSetSeatSettingsInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetSeatSettingsInCarIntentClass) New() INSetSeatSettingsInCarIntent {
	rv := objc.Send[INSetSeatSettingsInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetSeatSettingsInCarIntent) Init() INSetSeatSettingsInCarIntent {
	rv := objc.Send[INSetSeatSettingsInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetSeatSettingsInCarIntent) Autorelease() INSetSeatSettingsInCarIntent {
	rv := objc.Send[INSetSeatSettingsInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetSeatSettingsInCarIntent creates a new INSetSeatSettingsInCarIntent instance.
func NewINSetSeatSettingsInCarIntent() INSetSeatSettingsInCarIntent {
	return getINSetSeatSettingsInCarIntentClass().New()
}


// A Boolean value indicating whether to enable the seat cooling system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetSeatSettingsInCarIntent/enableCooling-6bcu7
func (i_ INSetSeatSettingsInCarIntent) EnableCooling() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("enableCooling"))
	return rv
}



