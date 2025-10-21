// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioApplication] class.
var (
	AudioApplicationClass     _AudioApplicationClass
	AudioApplicationClassOnce sync.Once
)

func getAudioApplicationClass() _AudioApplicationClass {
	AudioApplicationClassOnce.Do(func() {
		AudioApplicationClass = _AudioApplicationClass{objc.GetClass("AVAudioApplication")}
	})
	return AudioApplicationClass
}

type _AudioApplicationClass struct {
	class objc.Class
}

// An interface definition for the [AudioApplication] class.
type IAudioApplication interface {
	objectivec.IObject
	SetInputMuteStateChangeHandlerError(inputMuteHandler unsafe.Pointer, outError unsafe.Pointer) bool
	SetInputMutedError(muted bool, outError unsafe.Pointer) bool
}

// An object that manages one or more audio sessions that belong to an app.
//
// Access the shared audio application instance to control app-level audio operations, such as requesting microphone permission and controlling audio input muting.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication
type AudioApplication struct {
	objectivec.Object
}

// AudioApplicationFrom constructs a [AudioApplication] from an unsafe.Pointer.
//
// An object that manages one or more audio sessions that belong to an app.
func AudioApplicationFrom(ptr unsafe.Pointer) AudioApplication {
	return AudioApplication{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioApplicationClass) Alloc() AudioApplication {
	rv := objc.Send[AudioApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioApplicationClass) New() AudioApplication {
	rv := objc.Send[AudioApplication](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioApplication) Init() AudioApplication {
	rv := objc.Send[AudioApplication](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioApplication) Autorelease() AudioApplication {
	rv := objc.Send[AudioApplication](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioApplication creates a new AudioApplication instance.
func NewAudioApplication() AudioApplication {
	return getAudioApplicationClass().New()
}


// Requests the app’s permission to add audio to calls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestMicrophoneInjectionPermission(completionHandler:)
func (ac _AudioApplicationClass) RequestMicrophoneInjectionPermissionWithCompletionHandler(response unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestMicrophoneInjectionPermissionWithCompletionHandler:"), response)
}

// Determines whether the app has permission to record audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestRecordPermission(completionHandler:)
func (ac _AudioApplicationClass) RequestRecordPermissionWithCompletionHandler(response unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestRecordPermissionWithCompletionHandler:"), response)
}

// Accesses the shared audio application instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/shared
func (ac _AudioApplicationClass) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}
// Sets a callback to handle changes to application-level audio muting states.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuteStateChangeHandler(_:)
func (a_ AudioApplication) SetInputMuteStateChangeHandlerError(inputMuteHandler unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputMuteStateChangeHandler:error:"), inputMuteHandler, outError)
	return rv
}

// Sets a Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuted(_:)
func (a_ AudioApplication) SetInputMutedError(muted bool, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputMuted:error:"), muted, outError)
	return rv
}

// A Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioapplication/isinputmuted
func (a_ AudioApplication) IsInputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInputMuted"))
	return rv
}


// SetIsInputMuted sets the value of the isInputMuted property.
// A Boolean value that indicates whether the app’s audio input is in a muted state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioapplication/isinputmuted
func (a_ AudioApplication) SetIsInputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInputMuted:"), value)
}

// A Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/isInputMuted
func (a_ AudioApplication) InputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputMuted"))
	return rv
}

// A value that indicates an app’s permission to add audio to calls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/microphoneInjectionPermission-swift.property
func (a_ AudioApplication) MicrophoneInjectionPermission() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("microphoneInjectionPermission"))
	return rv
}

// The app’s permission to record audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.property
func (a_ AudioApplication) RecordPermission() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("recordPermission"))
	return rv
}

// Accesses the shared audio application instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/shared
func (a_ AudioApplication) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sharedInstance"))
	return rv
}



