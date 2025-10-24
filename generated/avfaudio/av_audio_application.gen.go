// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioApplication */


/* debug [class_header]: Header for AVAudioApplication */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioApplication */
// An interface definition for the [AudioApplication] class.
type IAudioApplication interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioApplication */
	// properties:
	InputMuted() bool
	RecordPermission() AudioApplicationRecordPermission
	IsInputMuted() bool
	SetIsInputMuted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioApplication */
	// methods:
	SetInputMutedError(muted bool, outError objectivec.IObject) bool
	SetInputMuteStateChangeHandlerError(inputMuteHandler unsafe.Pointer, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioApplication */
// Alloc allocates a new instance without initialization.
func (ac _AudioApplicationClass) Alloc() AudioApplication {
	rv := objc.Send[AudioApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioApplication */
// An object that manages one or more audio sessions that belong to an app.
//
// Access the shared audio application instance to control app-level audio operations, such as requesting microphone permission and controlling audio input muting.


// An object that manages one or more audio sessions that belong to an app.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioApplication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioApplication */

// Requests the app’s permission to add audio to calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestMicrophoneInjectionPermission(completionHandler:)
func (ac _AudioApplicationClass) RequestMicrophoneInjectionPermissionWithCompletionHandler(response unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestMicrophoneInjectionPermissionWithCompletionHandler:"), response)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestMicrophoneInjectionPermissionWithCompletionHandler) */


// Determines whether the app has permission to record audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestRecordPermission(completionHandler:)
func (ac _AudioApplicationClass) RequestRecordPermissionWithCompletionHandler(response unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestRecordPermissionWithCompletionHandler:"), response)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestRecordPermissionWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioApplication */

// Accesses the shared audio application instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/shared
func (ac _AudioApplicationClass) SharedInstance() AudioApplication {
	rv := objc.Send[AudioApplication](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioApplication */

// Sets a Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuted(_:)
func (a_ AudioApplication) SetInputMutedError(muted bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputMuted:error:"), muted, outError)
	return rv
}/* debug [instance_methods/method]: SetInputMutedError */


// Sets a callback to handle changes to application-level audio muting states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuteStateChangeHandler(_:)
func (a_ AudioApplication) SetInputMuteStateChangeHandlerError(inputMuteHandler unsafe.Pointer, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setInputMuteStateChangeHandler:error:"), inputMuteHandler, outError)
	return rv
}/* debug [instance_methods/method]: SetInputMuteStateChangeHandlerError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioApplication */

// A Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/isInputMuted
func (a_ AudioApplication) InputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("inputMuted"))
	return rv
}/* debug [instance_properties/getter]: inputMuted */


// The app’s permission to record audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.property
func (a_ AudioApplication) RecordPermission() AudioApplicationRecordPermission {
	rv := objc.Send[AudioApplicationRecordPermission](a_.ID, objc.Sel("recordPermission"))
	return rv
}/* debug [instance_properties/getter]: recordPermission */


// Accesses the shared audio application instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/shared
func (a_ AudioApplication) SharedInstance() IAVAudioApplication {
	rv := objc.Send[AudioApplication](a_.ID, objc.Sel("sharedInstance"))
	return rv
}/* debug [instance_properties/getter]: sharedInstance */


// A Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioapplication/isinputmuted
func (a_ AudioApplication) IsInputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInputMuted"))
	return rv
}/* debug [instance_properties/getter]: isInputMuted */


// A Boolean value that indicates whether the app’s audio input is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioapplication/isinputmuted
func (a_ AudioApplication) SetIsInputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInputMuted:"), value)
}/* debug [instance_properties/setter]: isInputMuted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioApplication */


