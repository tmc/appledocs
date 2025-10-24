//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionPortExtensionBluetoothMicrophone


// iOS-only properties

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





