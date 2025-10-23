// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSaveProfileInCarIntent] class.
var (
	INSaveProfileInCarIntentClass     _INSaveProfileInCarIntentClass
	INSaveProfileInCarIntentClassOnce sync.Once
)

func getINSaveProfileInCarIntentClass() _INSaveProfileInCarIntentClass {
	INSaveProfileInCarIntentClassOnce.Do(func() {
		INSaveProfileInCarIntentClass = _INSaveProfileInCarIntentClass{objc.GetClass("INSaveProfileInCarIntent")}
	})
	return INSaveProfileInCarIntentClass
}

type _INSaveProfileInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSaveProfileInCarIntent] class.
type IINSaveProfileInCarIntent interface {
	IINIntent
	ProfileLabel() string
	SetProfileLabel(value string)
	ProfileName() string
	SetProfileName(value string)
	ProfileNumber() int
	SetProfileNumber(value int)
}

// A request to save the user’s vehicle environment settings in a CarPlay-enabled vehicle.
//
// Automotive vendors whose cars support the saving of seat and other environment settings can add support for this intent to an Intents extension that they ship with their automotive apps. When users engage Siri to save the current environment settings, SiriKit creates an object and delivers it to the app’s Intents extension. You use the intent object to get the name or index of the profile to use when saving the settings. You’re responsible for determining which settings to save and restore with user profiles. You can save seat-related settings, climate control settings, defroster settings, radio settings, other settings in your vehicle, or any combination of those settings. Siri handles only the name or index of the profile and doesn’t ask you to provide a list of the settings that you saved. The object that handles this intent must adopt the protocol. Use this intent object to resolve the audio source details and to create an object indicating the results of changing the audio source.


// A request to save the user’s vehicle environment settings in a CarPlay-enabled vehicle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSaveProfileInCarIntent
type INSaveProfileInCarIntent struct {
	INIntent
}

// INSaveProfileInCarIntentFrom constructs a [INSaveProfileInCarIntent] from an unsafe.Pointer.
//
// A request to save the user’s vehicle environment settings in a CarPlay-enabled vehicle.
func INSaveProfileInCarIntentFrom(ptr unsafe.Pointer) INSaveProfileInCarIntent {
	return INSaveProfileInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSaveProfileInCarIntentClass) Alloc() INSaveProfileInCarIntent {
	rv := objc.Send[INSaveProfileInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSaveProfileInCarIntentClass) New() INSaveProfileInCarIntent {
	rv := objc.Send[INSaveProfileInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSaveProfileInCarIntent) Init() INSaveProfileInCarIntent {
	rv := objc.Send[INSaveProfileInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSaveProfileInCarIntent) Autorelease() INSaveProfileInCarIntent {
	rv := objc.Send[INSaveProfileInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSaveProfileInCarIntent creates a new INSaveProfileInCarIntent instance.
func NewINSaveProfileInCarIntent() INSaveProfileInCarIntent {
	return getINSaveProfileInCarIntentClass().New()
}



// The name to assign to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilelabel
func (i_ INSaveProfileInCarIntent) ProfileLabel() string {
	rv := objc.Send[string](i_.ID, objc.Sel("profileLabel"))
	return rv
}


// The name to assign to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilelabel
func (i_ INSaveProfileInCarIntent) SetProfileLabel(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileLabel:"), objc.String(value))
}


// The name to assign to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilename
func (i_ INSaveProfileInCarIntent) ProfileName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("profileName"))
	return rv
}


// The name to assign to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilename
func (i_ INSaveProfileInCarIntent) SetProfileName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileName:"), objc.String(value))
}


// The profile index in which to save the settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilenumber-2q84c
func (i_ INSaveProfileInCarIntent) ProfileNumber() int {
	rv := objc.Send[int](i_.ID, objc.Sel("profileNumber"))
	return rv
}


// The profile index in which to save the settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insaveprofileincarintent/profilenumber-2q84c
func (i_ INSaveProfileInCarIntent) SetProfileNumber(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileNumber:"), value)
}



