// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
func (i_ INSetSeatSettingsInCarIntent) EnableCooling() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("enableCooling"))
	return rv
}

// The name of the car you applied the settings to.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/carname
func (i_ INSetSeatSettingsInCarIntent) CarName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("carName"))
	return rv
}


// SetCarName sets the value of the carName property.
// The name of the car you applied the settings to.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/carname
func (i_ INSetSeatSettingsInCarIntent) SetCarName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}

// A Boolean value indicating whether to enable the seat heating system.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/enableheating-8auz2
func (i_ INSetSeatSettingsInCarIntent) EnableHeating() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableHeating"))
	return rv
}


// SetEnableHeating sets the value of the enableHeating property.
// A Boolean value indicating whether to enable the seat heating system.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/enableheating-8auz2
func (i_ INSetSeatSettingsInCarIntent) SetEnableHeating(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableHeating:"), value)
}

// A Boolean value indicating whether to enable the seat massage system.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/enablemassage-46ndx
func (i_ INSetSeatSettingsInCarIntent) EnableMassage() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableMassage"))
	return rv
}


// SetEnableMassage sets the value of the enableMassage property.
// A Boolean value indicating whether to enable the seat massage system.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/enablemassage-46ndx
func (i_ INSetSeatSettingsInCarIntent) SetEnableMassage(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableMassage:"), value)
}

// An integer value indicating the desired level for the seat setting.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/level-94975
func (i_ INSetSeatSettingsInCarIntent) Level() int {
	rv := objc.Send[int](i_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
// An integer value indicating the desired level for the seat setting.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/level-94975
func (i_ INSetSeatSettingsInCarIntent) SetLevel(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLevel:"), value)
}

// A relative change to the level value.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/relativelevelsetting
func (i_ INSetSeatSettingsInCarIntent) RelativeLevelSetting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relativeLevelSetting"))
	return rv
}


// SetRelativeLevelSetting sets the value of the relativeLevelSetting property.
// A relative change to the level value.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/relativelevelsetting
func (i_ INSetSeatSettingsInCarIntent) SetRelativeLevelSetting(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelativeLevelSetting:"), value)
}

// The seat position whose settings you want to modify.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/seat
func (i_ INSetSeatSettingsInCarIntent) Seat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("seat"))
	return rv
}


// SetSeat sets the value of the seat property.
// The seat position whose settings you want to modify.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insetseatsettingsincarintent/seat
func (i_ INSetSeatSettingsInCarIntent) SetSeat(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSeat:"), value)
}



