//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionPortDescription


// Sets the preferred audio data source for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/setPreferredDataSource(_:)
func (a_ AudioSessionPortDescription) SetPreferredDataSourceError(dataSource IAVAudioSessionDataSourceDescription, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredDataSource:error:"), dataSource, outError)
	return rv
}

// iOS-only properties

// An optional port extension that describes capabilities relevant to Bluetooth microphone ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/bluetoothMicrophoneExtension
func (a_ AudioSessionPortDescription) BluetoothMicrophoneExtension() IAVAudioSessionPortExtensionBluetoothMicrophone {
	rv := objc.Send[AudioSessionPortExtensionBluetoothMicrophone](a_.ID, objc.Sel("bluetoothMicrophoneExtension"))
	return rv
}

// An array of channel objects that describe the port’s input or output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/channels
func (a_ AudioSessionPortDescription) Channels() []AudioSessionChannelDescription {
	rv := objc.Send[[]AudioSessionChannelDescription](a_.ID, objc.Sel("channels"))
	return rv
}

// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/dataSources
func (a_ AudioSessionPortDescription) DataSources() []AudioSessionDataSourceDescription {
	rv := objc.Send[[]AudioSessionDataSourceDescription](a_.ID, objc.Sel("dataSources"))
	return rv
}

// A Boolean value that indicates whether the associated hardware port has built-in processing for two-way voice communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/hasHardwareVoiceCallProcessing
func (a_ AudioSessionPortDescription) HasHardwareVoiceCallProcessing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasHardwareVoiceCallProcessing"))
	return rv
}

// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/isSpatialAudioEnabled
func (a_ AudioSessionPortDescription) SpatialAudioEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("spatialAudioEnabled"))
	return rv
}

// A descriptive name for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/portName
func (a_ AudioSessionPortDescription) PortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("portName"))
	return rv
}

// The type of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/portType
func (a_ AudioSessionPortDescription) PortType() AudioSessionPort /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("portType"))
	return rv
}

// The preferred audio data source for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/preferredDataSource
func (a_ AudioSessionPortDescription) PreferredDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("preferredDataSource"))
	return rv
}

// The currently selected audio data source for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/selectedDataSource
func (a_ AudioSessionPortDescription) SelectedDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("selectedDataSource"))
	return rv
}

// A system-assigned unique identifier (UID) for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/uid
func (a_ AudioSessionPortDescription) UID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("UID"))
	return rv
}





