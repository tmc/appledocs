// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetClimateSettingsInCarIntent] class.
var (
	INSetClimateSettingsInCarIntentClass     _INSetClimateSettingsInCarIntentClass
	INSetClimateSettingsInCarIntentClassOnce sync.Once
)

func getINSetClimateSettingsInCarIntentClass() _INSetClimateSettingsInCarIntentClass {
	INSetClimateSettingsInCarIntentClassOnce.Do(func() {
		INSetClimateSettingsInCarIntentClass = _INSetClimateSettingsInCarIntentClass{objc.GetClass("INSetClimateSettingsInCarIntent")}
	})
	return INSetClimateSettingsInCarIntentClass
}

type _INSetClimateSettingsInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetClimateSettingsInCarIntent] class.
type IINSetClimateSettingsInCarIntent interface {
	IINIntent
}

// A request to change the climate settings in a CarPlay-enabled vehicle.
//
// Automotive vendors can add support for this intent to an Intents extension that they ship with their automotive apps. When the user asks Siri to change a setting related to the vehicle’s climate control, Siri creates an object and delivers it to the app’s Intents extension. You use the intent to identify which setting the user wants to change. This class contains properties for multiple types of climate settings, but a given instance of this class contains changes for only one system at a time. When resolving and confirming the parameters of this intent, use the specified properties to modify your vehicle’s settings. Assume no changes for other settings. If your vehicle doesn’t support a particular setting, offer a reasonable fallback setting and ask the user for confirmation. The object that handles this intent must adopt the protocol. Use this intent object to identify which setting changed and to create an object indicating the results of changing the climate settings.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetClimateSettingsInCarIntent
type INSetClimateSettingsInCarIntent struct {
	INIntent
}

// INSetClimateSettingsInCarIntentFrom constructs a [INSetClimateSettingsInCarIntent] from an unsafe.Pointer.
//
// A request to change the climate settings in a CarPlay-enabled vehicle.
func INSetClimateSettingsInCarIntentFrom(ptr unsafe.Pointer) INSetClimateSettingsInCarIntent {
	return INSetClimateSettingsInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetClimateSettingsInCarIntentClass) Alloc() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetClimateSettingsInCarIntentClass) New() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetClimateSettingsInCarIntent) Init() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetClimateSettingsInCarIntent) Autorelease() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetClimateSettingsInCarIntent creates a new INSetClimateSettingsInCarIntent instance.
func NewINSetClimateSettingsInCarIntent() INSetClimateSettingsInCarIntent {
	return getINSetClimateSettingsInCarIntentClass().New()
}




