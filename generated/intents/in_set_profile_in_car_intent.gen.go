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
	CarName() INSpeakableString
	SetCarName(value INSpeakableString)
	DefaultProfile() int
	SetDefaultProfile(value int)
	IsDefaultProfile() bool
	SetIsDefaultProfile(value bool)
	ProfileLabel() string
	SetProfileLabel(value string)
	ProfileName() string
	SetProfileName(value string)
	ProfileNumber() int
	SetProfileNumber(value int)
}

// A request to change the user’s vehicle environment settings to the ones from the specified profile.
//
// Automotive vendors whose cars support the saving of seat and other environment settings can add support for this intent to an Intents extension that they ship with their automotive apps. When users want to restore settings from a profile, SiriKit creates an object and delivers it to the app’s Intents extension. You use the intent object to get the name or index of the profile whose settings you use to configure the vehicle. You’re responsible for determining which settings to save and restore with user profiles. You can restore seat-related settings, climate control settings, defroster settings, radio settings, other settings in your vehicle, or any combination of those settings. Siri handles only the name or index of the profile and doesn’t ask you to provide a list of the settings that you restored. Users can restore settings regardless for whether the profile was originally created through Siri or through your vehicle’s built-in interface. The object that handles this intent must adopt the protocol. Use this intent object to resolve the profile information and to create an object indicating the results of restoring the profile.

// A request to change the user’s vehicle environment settings to the ones from the specified profile.
//
// [Full Topic]
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

// The name of the car associated with the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/carname
func (i_ INSetProfileInCarIntent) CarName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("carName"))
	return rv
}

// The name of the car associated with the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/carname
func (i_ INSetProfileInCarIntent) SetCarName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/defaultprofile-19jwc
func (i_ INSetProfileInCarIntent) DefaultProfile() int {
	rv := objc.Send[int](i_.ID, objc.Sel("defaultProfile"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/defaultprofile-19jwc
func (i_ INSetProfileInCarIntent) SetDefaultProfile(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultProfile:"), value)
}

// A Boolean value indicating whether to make the profile the default profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/isdefaultprofile
func (i_ INSetProfileInCarIntent) IsDefaultProfile() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isDefaultProfile"))
	return rv
}

// A Boolean value indicating whether to make the profile the default profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/isdefaultprofile
func (i_ INSetProfileInCarIntent) SetIsDefaultProfile(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsDefaultProfile:"), value)
}

// The name assigned to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilelabel
func (i_ INSetProfileInCarIntent) ProfileLabel() string {
	rv := objc.Send[string](i_.ID, objc.Sel("profileLabel"))
	return rv
}

// The name assigned to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilelabel
func (i_ INSetProfileInCarIntent) SetProfileLabel(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileLabel:"), objc.String(value))
}

// The name assigned to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilename
func (i_ INSetProfileInCarIntent) ProfileName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("profileName"))
	return rv
}

// The name assigned to the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilename
func (i_ INSetProfileInCarIntent) SetProfileName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileName:"), objc.String(value))
}

// The profile index from which to restore the settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilenumber-37vj8
func (i_ INSetProfileInCarIntent) ProfileNumber() int {
	rv := objc.Send[int](i_.ID, objc.Sel("profileNumber"))
	return rv
}

// The profile index from which to restore the settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetprofileincarintent/profilenumber-37vj8
func (i_ INSetProfileInCarIntent) SetProfileNumber(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProfileNumber:"), value)
}
