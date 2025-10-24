// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetAudioSourceInCarIntent] class.
var (
	INSetAudioSourceInCarIntentClass     _INSetAudioSourceInCarIntentClass
	INSetAudioSourceInCarIntentClassOnce sync.Once
)

func getINSetAudioSourceInCarIntentClass() _INSetAudioSourceInCarIntentClass {
	INSetAudioSourceInCarIntentClassOnce.Do(func() {
		INSetAudioSourceInCarIntentClass = _INSetAudioSourceInCarIntentClass{objc.GetClass("INSetAudioSourceInCarIntent")}
	})
	return INSetAudioSourceInCarIntentClass
}

type _INSetAudioSourceInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetAudioSourceInCarIntent] class.
type IINSetAudioSourceInCarIntent interface {
	IINIntent
	AudioSource() unsafe.Pointer
	SetAudioSource(value unsafe.Pointer)
	RelativeAudioSourceReference() unsafe.Pointer
	SetRelativeAudioSourceReference(value unsafe.Pointer)
}

// A request to change the source of audio playback in a CarPlay-enabled vehicle.
//
// Automotive venders can add support for this intent to an Intents extension that they ship with their automotive apps. When the user asks Siri to change the audio source of the vehicle, SiriKit creates an object and delivers it to your app’s Intents extension. You use the intent to identify which audio source the user wants to use and to communicate the new audio source information directly to your vehicle’s systems. Users may select audio sources by name or by asking for the next or previous audio source that’s available. When the user asks for an audio source by name, Siri populates the property of this intent object with the specific requested source. When the user asks for the next or previous audio source, Siri places the appropriate value in the property so that you can determine which audio source to select. Only one of these properties contains usable information; Siri sets the other to a constant indicating an unknown status for the value. The object that handles this intent must adopt the protocol. Use this intent object to resolve the audio source details and to create an object indicating the results of changing the audio source.

// A request to change the source of audio playback in a CarPlay-enabled vehicle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetAudioSourceInCarIntent
type INSetAudioSourceInCarIntent struct {
	INIntent
}

// INSetAudioSourceInCarIntentFrom constructs a [INSetAudioSourceInCarIntent] from an unsafe.Pointer.
//
// A request to change the source of audio playback in a CarPlay-enabled vehicle.
func INSetAudioSourceInCarIntentFrom(ptr unsafe.Pointer) INSetAudioSourceInCarIntent {
	return INSetAudioSourceInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetAudioSourceInCarIntentClass) Alloc() INSetAudioSourceInCarIntent {
	rv := objc.Send[INSetAudioSourceInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetAudioSourceInCarIntentClass) New() INSetAudioSourceInCarIntent {
	rv := objc.Send[INSetAudioSourceInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetAudioSourceInCarIntent) Init() INSetAudioSourceInCarIntent {
	rv := objc.Send[INSetAudioSourceInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetAudioSourceInCarIntent) Autorelease() INSetAudioSourceInCarIntent {
	rv := objc.Send[INSetAudioSourceInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetAudioSourceInCarIntent creates a new INSetAudioSourceInCarIntent instance.
func NewINSetAudioSourceInCarIntent() INSetAudioSourceInCarIntent {
	return getINSetAudioSourceInCarIntentClass().New()
}

// The audio source to select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetaudiosourceincarintent/audiosource
func (i_ INSetAudioSourceInCarIntent) AudioSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("audioSource"))
	return rv
}

// The audio source to select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetaudiosourceincarintent/audiosource
func (i_ INSetAudioSourceInCarIntent) SetAudioSource(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAudioSource:"), value)
}

// The relative audio source to select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetaudiosourceincarintent/relativeaudiosourcereference
func (i_ INSetAudioSourceInCarIntent) RelativeAudioSourceReference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relativeAudioSourceReference"))
	return rv
}

// The relative audio source to select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetaudiosourceincarintent/relativeaudiosourcereference
func (i_ INSetAudioSourceInCarIntent) SetRelativeAudioSourceReference(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelativeAudioSourceReference:"), value)
}
