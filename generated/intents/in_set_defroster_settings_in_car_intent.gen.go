// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetDefrosterSettingsInCarIntent] class.
var (
	INSetDefrosterSettingsInCarIntentClass     _INSetDefrosterSettingsInCarIntentClass
	INSetDefrosterSettingsInCarIntentClassOnce sync.Once
)

func getINSetDefrosterSettingsInCarIntentClass() _INSetDefrosterSettingsInCarIntentClass {
	INSetDefrosterSettingsInCarIntentClassOnce.Do(func() {
		INSetDefrosterSettingsInCarIntentClass = _INSetDefrosterSettingsInCarIntentClass{objc.GetClass("INSetDefrosterSettingsInCarIntent")}
	})
	return INSetDefrosterSettingsInCarIntentClass
}

type _INSetDefrosterSettingsInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetDefrosterSettingsInCarIntent] class.
type IINSetDefrosterSettingsInCarIntent interface {
	IINIntent
	Enable() foundation.Number
	CarName() INSpeakableString
	SetCarName(value INSpeakableString)
	Defroster() unsafe.Pointer
	SetDefroster(value unsafe.Pointer)
}

// A request to change the defroster settings in a CarPlay-enabled vehicle.
//
// Automotive venders can add support for this intent to an Intents extension that they ship with their automotive apps. When the user asks Siri to change the defroster settings for a vehicle, SiriKit creates an object and delivers it to the app’s Intents extension. You use the intent to identify which defroster the user wants to enable or disable and to communicate the changes directly to your vehicle’s systems. The object that handles this intent must adopt the protocol. Use this intent object to resolve the defroster details and to create an object indicating whether you were able to make the change successfully.

// A request to change the defroster settings in a CarPlay-enabled vehicle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetDefrosterSettingsInCarIntent
type INSetDefrosterSettingsInCarIntent struct {
	INIntent
}

// INSetDefrosterSettingsInCarIntentFrom constructs a [INSetDefrosterSettingsInCarIntent] from an unsafe.Pointer.
//
// A request to change the defroster settings in a CarPlay-enabled vehicle.
func INSetDefrosterSettingsInCarIntentFrom(ptr unsafe.Pointer) INSetDefrosterSettingsInCarIntent {
	return INSetDefrosterSettingsInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetDefrosterSettingsInCarIntentClass) Alloc() INSetDefrosterSettingsInCarIntent {
	rv := objc.Send[INSetDefrosterSettingsInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetDefrosterSettingsInCarIntentClass) New() INSetDefrosterSettingsInCarIntent {
	rv := objc.Send[INSetDefrosterSettingsInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetDefrosterSettingsInCarIntent) Init() INSetDefrosterSettingsInCarIntent {
	rv := objc.Send[INSetDefrosterSettingsInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetDefrosterSettingsInCarIntent) Autorelease() INSetDefrosterSettingsInCarIntent {
	rv := objc.Send[INSetDefrosterSettingsInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetDefrosterSettingsInCarIntent creates a new INSetDefrosterSettingsInCarIntent instance.
func NewINSetDefrosterSettingsInCarIntent() INSetDefrosterSettingsInCarIntent {
	return getINSetDefrosterSettingsInCarIntentClass().New()
}

// A Boolean indicating whether to enable or disable the defroster.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetDefrosterSettingsInCarIntent/enable-8tf0i
func (i_ INSetDefrosterSettingsInCarIntent) Enable() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("enable"))
	return rv
}

// The name of the car you applied the settings to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetdefrostersettingsincarintent/carname
func (i_ INSetDefrosterSettingsInCarIntent) CarName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("carName"))
	return rv
}

// The name of the car you applied the settings to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetdefrostersettingsincarintent/carname
func (i_ INSetDefrosterSettingsInCarIntent) SetCarName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}

// The defroster to enable or disable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetdefrostersettingsincarintent/defroster
func (i_ INSetDefrosterSettingsInCarIntent) Defroster() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("defroster"))
	return rv
}

// The defroster to enable or disable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetdefrostersettingsincarintent/defroster
func (i_ INSetDefrosterSettingsInCarIntent) SetDefroster(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefroster:"), value)
}
