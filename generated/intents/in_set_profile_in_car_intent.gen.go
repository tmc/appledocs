// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetProfileInCarIntent] class.
var (
	INSetProfileInCarIntentClass     _INSetProfileInCarIntentClass
	INSetProfileInCarIntentClassOnce sync.Once
)

func getINSetProfileInCarIntentClass() _INSetProfileInCarIntentClass {
	INSetProfileInCarIntentClassOnce.Do(func() {
		INSetProfileInCarIntentClass = _INSetProfileInCarIntentClass{objc.GetClass("INSetProfileInCarIntent")}
	})
	return INSetProfileInCarIntentClass
}

type _INSetProfileInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetProfileInCarIntent] class.
type IINSetProfileInCarIntent interface {
	IINIntent
}

// A request to change the user’s vehicle environment settings to the ones from the specified profile.
//
// Automotive vendors whose cars support the saving of seat and other environment settings can add support for this intent to an Intents extension that they ship with their automotive apps. When users want to restore settings from a profile, SiriKit creates an object and delivers it to the app’s Intents extension. You use the intent object to get the name or index of the profile whose settings you use to configure the vehicle. You’re responsible for determining which settings to save and restore with user profiles. You can restore seat-related settings, climate control settings, defroster settings, radio settings, other settings in your vehicle, or any combination of those settings. Siri handles only the name or index of the profile and doesn’t ask you to provide a list of the settings that you restored. Users can restore settings regardless for whether the profile was originally created through Siri or through your vehicle’s built-in interface. The object that handles this intent must adopt the protocol. Use this intent object to resolve the profile information and to create an object indicating the results of restoring the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetProfileInCarIntent
type INSetProfileInCarIntent struct {
	INIntent
}

// INSetProfileInCarIntentFrom constructs a [INSetProfileInCarIntent] from an unsafe.Pointer.
//
// A request to change the user’s vehicle environment settings to the ones from the specified profile.
func INSetProfileInCarIntentFrom(ptr unsafe.Pointer) INSetProfileInCarIntent {
	return INSetProfileInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetProfileInCarIntentClass) Alloc() INSetProfileInCarIntent {
	rv := objc.Send[INSetProfileInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetProfileInCarIntentClass) New() INSetProfileInCarIntent {
	rv := objc.Send[INSetProfileInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetProfileInCarIntent) Init() INSetProfileInCarIntent {
	rv := objc.Send[INSetProfileInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetProfileInCarIntent) Autorelease() INSetProfileInCarIntent {
	rv := objc.Send[INSetProfileInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetProfileInCarIntent creates a new INSetProfileInCarIntent instance.
func NewINSetProfileInCarIntent() INSetProfileInCarIntent {
	return getINSetProfileInCarIntentClass().New()
}




