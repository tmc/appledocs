// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSessionPortExtensionBluetoothMicrophone */


/* debug [class_header]: Header for AVAudioSessionPortExtensionBluetoothMicrophone */
// The class instance for the [AudioSessionPortExtensionBluetoothMicrophone] class.
var (
	AudioSessionPortExtensionBluetoothMicrophoneClass     _AudioSessionPortExtensionBluetoothMicrophoneClass
	AudioSessionPortExtensionBluetoothMicrophoneClassOnce sync.Once
)

func getAudioSessionPortExtensionBluetoothMicrophoneClass() _AudioSessionPortExtensionBluetoothMicrophoneClass {
	AudioSessionPortExtensionBluetoothMicrophoneClassOnce.Do(func() {
		AudioSessionPortExtensionBluetoothMicrophoneClass = _AudioSessionPortExtensionBluetoothMicrophoneClass{objc.GetClass("AVAudioSessionPortExtensionBluetoothMicrophone")}
	})
	return AudioSessionPortExtensionBluetoothMicrophoneClass
}

type _AudioSessionPortExtensionBluetoothMicrophoneClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSessionPortExtensionBluetoothMicrophone */
// An interface definition for the [AudioSessionPortExtensionBluetoothMicrophone] class.
type IAudioSessionPortExtensionBluetoothMicrophone interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSessionPortExtensionBluetoothMicrophone */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSessionPortExtensionBluetoothMicrophone */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSessionPortExtensionBluetoothMicrophone */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionPortExtensionBluetoothMicrophoneClass) Alloc() AudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionPortExtensionBluetoothMicrophoneClass) New() AudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionPortExtensionBluetoothMicrophone) Init() AudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionPortExtensionBluetoothMicrophone) Autorelease() AudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionPortExtensionBluetoothMicrophone creates a new AudioSessionPortExtensionBluetoothMicrophone instance.
func NewAudioSessionPortExtensionBluetoothMicrophone() AudioSessionPortExtensionBluetoothMicrophone {
	return getAudioSessionPortExtensionBluetoothMicrophoneClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSessionPortExtensionBluetoothMicrophone */
// An object that describes capabilities of Bluetooth microphone ports.


// An object that describes capabilities of Bluetooth microphone ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortExtensionBluetoothMicrophone
type AudioSessionPortExtensionBluetoothMicrophone struct {
	objectivec.Object
}

// AudioSessionPortExtensionBluetoothMicrophoneFrom constructs a [AudioSessionPortExtensionBluetoothMicrophone] from an unsafe.Pointer.
//
// An object that describes capabilities of Bluetooth microphone ports.
func AudioSessionPortExtensionBluetoothMicrophoneFrom(ptr unsafe.Pointer) AudioSessionPortExtensionBluetoothMicrophone {
	return AudioSessionPortExtensionBluetoothMicrophone{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSessionPortExtensionBluetoothMicrophone *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSessionPortExtensionBluetoothMicrophone */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSessionPortExtensionBluetoothMicrophone */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSessionPortExtensionBluetoothMicrophone */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSessionPortExtensionBluetoothMicrophone */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSessionPortExtensionBluetoothMicrophone */


