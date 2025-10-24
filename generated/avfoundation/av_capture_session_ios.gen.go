//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSession


// iOS-only properties

// A Boolean value that indicates whether the capture session automatically changes settings in the app’s shared audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresApplicationAudioSession
func (c_ CaptureSession) AutomaticallyConfiguresApplicationAudioSession() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresApplicationAudioSession"))
	return rv
}
func (c_ CaptureSession) SetAutomaticallyConfiguresApplicationAudioSession(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyConfiguresApplicationAudioSession:"), value)
}

// A Boolean value that specifies whether the session should automatically use wide-gamut color where available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresCaptureDeviceForWideColor
func (c_ CaptureSession) AutomaticallyConfiguresCaptureDeviceForWideColor() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresCaptureDeviceForWideColor"))
	return rv
}
func (c_ CaptureSession) SetAutomaticallyConfiguresCaptureDeviceForWideColor(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyConfiguresCaptureDeviceForWideColor:"), value)
}

// A Boolean value that indicates whether the capture session configures the app’s audio session for bluetooth high-quality recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionForBluetoothHighQualityRecording
func (c_ CaptureSession) ConfiguresApplicationAudioSessionForBluetoothHighQualityRecording() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("configuresApplicationAudioSessionForBluetoothHighQualityRecording"))
	return rv
}
func (c_ CaptureSession) SetConfiguresApplicationAudioSessionForBluetoothHighQualityRecording(value bool) {
	c_.ID.Send(objc.RegisterName("setConfiguresApplicationAudioSessionForBluetoothHighQualityRecording:"), value)
}

// A Boolean value that Indicates whether the capture session configures the app’s audio session to mix with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionToMixWithOthers
func (c_ CaptureSession) ConfiguresApplicationAudioSessionToMixWithOthers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("configuresApplicationAudioSessionToMixWithOthers"))
	return rv
}
func (c_ CaptureSession) SetConfiguresApplicationAudioSessionToMixWithOthers(value bool) {
	c_.ID.Send(objc.RegisterName("setConfiguresApplicationAudioSessionToMixWithOthers:"), value)
}

// A value that indicates the percentage of the session’s available hardware budget in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/hardwareCost
func (c_ CaptureSession) HardwareCost() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("hardwareCost"))
	return rv
}

// A Boolean value that indicates whether the capture session is in an interrupted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isInterrupted
func (c_ CaptureSession) Interrupted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("interrupted"))
	return rv
}

// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessEnabled
func (c_ CaptureSession) MultitaskingCameraAccessEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multitaskingCameraAccessEnabled"))
	return rv
}
func (c_ CaptureSession) SetMultitaskingCameraAccessEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setMultitaskingCameraAccessEnabled:"), value)
}

// A Boolean value that indicates whether the capture session supports using the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessSupported
func (c_ CaptureSession) MultitaskingCameraAccessSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multitaskingCameraAccessSupported"))
	return rv
}

// A Boolean value that indicates whether the capture session uses the app’s shared audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/usesApplicationAudioSession
func (c_ CaptureSession) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}
func (c_ CaptureSession) SetUsesApplicationAudioSession(value bool) {
	c_.ID.Send(objc.RegisterName("setUsesApplicationAudioSession:"), value)
}





