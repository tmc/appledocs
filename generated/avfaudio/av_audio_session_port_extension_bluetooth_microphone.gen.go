// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioSessionPortExtensionBluetoothMicrophone] class.
type IAudioSessionPortExtensionBluetoothMicrophone interface {
	objectivec.IObject
	FarFieldCapture() IAVAudioSessionCapability
	HighQualityRecording() IAVAudioSessionCapability
}

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

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionPortExtensionBluetoothMicrophoneClass) Alloc() AudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Describes whether this port supports far-field input capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortExtensionBluetoothMicrophone/farFieldCapture
func (a_ AudioSessionPortExtensionBluetoothMicrophone) FarFieldCapture() IAVAudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("farFieldCapture"))
	return rv
}


// Describes whether this port supports Bluetooth high-quality recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortExtensionBluetoothMicrophone/highQualityRecording
func (a_ AudioSessionPortExtensionBluetoothMicrophone) HighQualityRecording() IAVAudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("highQualityRecording"))
	return rv
}



