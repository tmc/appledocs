// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationSound] class.
var (
	UNNotificationSoundClass     _UNNotificationSoundClass
	UNNotificationSoundClassOnce sync.Once
)

func getUNNotificationSoundClass() _UNNotificationSoundClass {
	UNNotificationSoundClassOnce.Do(func() {
		UNNotificationSoundClass = _UNNotificationSoundClass{objc.GetClass("UNNotificationSound")}
	})
	return UNNotificationSoundClass
}

type _UNNotificationSoundClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationSound] class.
type IUNNotificationSound interface {
	objectivec.IObject
	// properties:
	Sound() IUNNotificationSound
	SetSound(value IUNNotificationSound)
	// methods:
}

// The sound played upon delivery of a notification.
//
// Create a object when you want the system to play a specific sound when it delivers with your notification. To play the default system sound, create your sound object using the method. If you want to play a custom sound, create a new sound object and specify the name of the audio file that you want to play. For local notifications, assign the sound object to the property of your object. For a remote notification, assign the name of your sound file to the key in the dictionary. You can also use a notification service app extension to add a sound file to a notification shortly before delivery. In your extension, create a object and add it to your notification content in the same way that you’d for a local notification. Audio files must already be on the user’s device before the system can play them. If you use a predefined set of sounds for your notifications, include the audio files in your app’s bundle. For all other sounds, the object looks only in the following locations: The directory of the app’s container directory. The directory of one of the app’s shared group container directories. The main bundle of the current executable.


// The sound played upon delivery of a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound
type UNNotificationSound struct {
	objectivec.Object
}

// UNNotificationSoundFrom constructs a [UNNotificationSound] from an unsafe.Pointer.
//
// The sound played upon delivery of a notification.
func UNNotificationSoundFrom(ptr unsafe.Pointer) UNNotificationSound {
	return UNNotificationSound{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationSoundClass) Alloc() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationSoundClass) New() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationSound) Init() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationSound) Autorelease() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationSound creates a new UNNotificationSound instance.
func NewUNNotificationSound() UNNotificationSound {
	return getUNNotificationSoundClass().New()
}



// Creates a sound object that represents a custom sound file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/init(named:)
func NewUNNotificationSoundNamed(name UNNotificationSoundName /* typedef */) UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(getUNNotificationSoundClass().class), objc.Sel("soundNamed:"), name)
	return rv
}



// Creates a custom sound object for critical alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/criticalSoundNamed(_:)
func (uc _UNNotificationSoundClass) CriticalSoundNamed(name UNNotificationSoundName /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("criticalSoundNamed:"), name)
	return rv
}


// Creates a custom sound object for critical alerts with the volume you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/criticalSoundNamed(_:withAudioVolume:)
func (uc _UNNotificationSoundClass) CriticalSoundNamedWithAudioVolume(name UNNotificationSoundName /* typedef */, volume float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("criticalSoundNamed:withAudioVolume:"), name, volume)
	return rv
}


// Creates a sound object that plays the default critical alert sound at the volume you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultCriticalSound(withAudioVolume:)
func (uc _UNNotificationSoundClass) DefaultCriticalSoundWithAudioVolume(volume float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("defaultCriticalSoundWithAudioVolume:"), volume)
	return rv
}


// Creates a sound object that represents a custom sound file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/init(named:)
func (uc _UNNotificationSoundClass) SoundNamed(name UNNotificationSoundName /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("soundNamed:"), name)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/ringtoneSoundNamed(_:)
func (uc _UNNotificationSoundClass) RingtoneSoundNamed(name UNNotificationSoundName /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("ringtoneSoundNamed:"), name)
	return rv
}


// Returns an object representing the default sound for notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/default
func (uc _UNNotificationSoundClass) DefaultSound() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("defaultSound"))
	return rv
}

// The default sound used for critical alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultCritical
func (uc _UNNotificationSoundClass) DefaultCriticalSound() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("defaultCriticalSound"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultRingtone
func (uc _UNNotificationSoundClass) DefaultRingtoneSound() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("defaultRingtoneSound"))
	return rv
}

// Returns an object representing the default sound for notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/default
func (u_ UNNotificationSound) DefaultSound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("defaultSound"))
	return rv
}


// The default sound used for critical alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultCritical
func (u_ UNNotificationSound) DefaultCriticalSound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("defaultCriticalSound"))
	return rv
}


// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/sound
func (u_ UNNotificationSound) Sound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("sound"))
	return rv
}


// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/sound
func (u_ UNNotificationSound) SetSound(value IUNNotificationSound) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSound:"), value)
}


