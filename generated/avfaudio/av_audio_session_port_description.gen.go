// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionPortDescription] class.
var (
	AudioSessionPortDescriptionClass     _AudioSessionPortDescriptionClass
	AudioSessionPortDescriptionClassOnce sync.Once
)

func getAudioSessionPortDescriptionClass() _AudioSessionPortDescriptionClass {
	AudioSessionPortDescriptionClassOnce.Do(func() {
		AudioSessionPortDescriptionClass = _AudioSessionPortDescriptionClass{objc.GetClass("AVAudioSessionPortDescription")}
	})
	return AudioSessionPortDescriptionClass
}

type _AudioSessionPortDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionPortDescription] class.
type IAudioSessionPortDescription interface {
	objectivec.IObject
}

// Information about the capabilities of the port and the hardware channels it supports.
//
// A port description object describes a single input or output port associated with an audio route. Examples of audio ports include a device’s built-in speaker, a microphone on a wired headset, and a Bluetooth device supporting the Advanced Audio Distribution Profile (A2DP). You can query the audio session’s property to get information about the active set of input and output ports. To change the current audio routing, call the method. For example, on a device with a wired headset attached, the audio session’s array may contain two port descriptions: one for the headset microphone and one for the device’s built-in microphone. You can use the audio session’s method to select the headset or built-in microphone for audio input.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription
type AudioSessionPortDescription struct {
	objectivec.Object
}

// AudioSessionPortDescriptionFrom constructs a [AudioSessionPortDescription] from an unsafe.Pointer.
//
// Information about the capabilities of the port and the hardware channels it supports.
func AudioSessionPortDescriptionFrom(ptr unsafe.Pointer) AudioSessionPortDescription {
	return AudioSessionPortDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionPortDescriptionClass) Alloc() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionPortDescriptionClass) New() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionPortDescription) Init() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionPortDescription) Autorelease() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionPortDescription creates a new AudioSessionPortDescription instance.
func NewAudioSessionPortDescription() AudioSessionPortDescription {
	return getAudioSessionPortDescriptionClass().New()
}


// The preferred audio data source for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/preferredDataSource
func (a_ AudioSessionPortDescription) PreferredDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredDataSource"))
	return rv
}

// An array of input ports available for audio routing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/availableinputs
func (a_ AudioSessionPortDescription) AvailableInputs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableInputs"))
	return rv
}


// SetAvailableInputs sets the value of the availableInputs property.
// An array of input ports available for audio routing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/availableinputs
func (a_ AudioSessionPortDescription) SetAvailableInputs(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableInputs:"), value)
}

// A description of the current audio route’s input and output ports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionPortDescription) CurrentRoute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentRoute"))
	return rv
}


// SetCurrentRoute sets the value of the currentRoute property.
// A description of the current audio route’s input and output ports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/currentroute
func (a_ AudioSessionPortDescription) SetCurrentRoute(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentRoute:"), value)
}

// An optional port extension that describes capabilities relevant to Bluetooth microphone ports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/bluetoothmicrophoneextension
func (a_ AudioSessionPortDescription) BluetoothMicrophoneExtension() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("bluetoothMicrophoneExtension"))
	return rv
}


// SetBluetoothMicrophoneExtension sets the value of the bluetoothMicrophoneExtension property.
// An optional port extension that describes capabilities relevant to Bluetooth microphone ports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/bluetoothmicrophoneextension
func (a_ AudioSessionPortDescription) SetBluetoothMicrophoneExtension(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBluetoothMicrophoneExtension:"), value)
}

// An array of channel objects that describe the port’s input or output channels.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/channels
func (a_ AudioSessionPortDescription) Channels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("channels"))
	return rv
}


// SetChannels sets the value of the channels property.
// An array of channel objects that describe the port’s input or output channels.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/channels
func (a_ AudioSessionPortDescription) SetChannels(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannels:"), value)
}

// The available data sources for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionPortDescription) DataSources() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("dataSources"))
	return rv
}


// SetDataSources sets the value of the dataSources property.
// The available data sources for the port.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionPortDescription) SetDataSources(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSources:"), value)
}

// A Boolean value that indicates whether the associated hardware port has built-in processing for two-way voice communication.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/hashardwarevoicecallprocessing
func (a_ AudioSessionPortDescription) HasHardwareVoiceCallProcessing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasHardwareVoiceCallProcessing"))
	return rv
}


// SetHasHardwareVoiceCallProcessing sets the value of the hasHardwareVoiceCallProcessing property.
// A Boolean value that indicates whether the associated hardware port has built-in processing for two-way voice communication.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/hashardwarevoicecallprocessing
func (a_ AudioSessionPortDescription) SetHasHardwareVoiceCallProcessing(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasHardwareVoiceCallProcessing:"), value)
}

// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionPortDescription) IsSpatialAudioEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSpatialAudioEnabled"))
	return rv
}


// SetIsSpatialAudioEnabled sets the value of the isSpatialAudioEnabled property.
// A Boolean value that indicates whether the port supports spatial audio playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionPortDescription) SetIsSpatialAudioEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSpatialAudioEnabled:"), value)
}

// A descriptive name for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/portname
func (a_ AudioSessionPortDescription) PortName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("portName"))
	return rv
}


// SetPortName sets the value of the portName property.
// A descriptive name for the port.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/portname
func (a_ AudioSessionPortDescription) SetPortName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPortName:"), objc.String(value))
}

// The type of the port.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/porttype
func (a_ AudioSessionPortDescription) PortType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("portType"))
	return rv
}


// SetPortType sets the value of the portType property.
// The type of the port.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/porttype
func (a_ AudioSessionPortDescription) SetPortType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPortType:"), value)
}

// The currently selected audio data source for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/selecteddatasource
func (a_ AudioSessionPortDescription) SelectedDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("selectedDataSource"))
	return rv
}


// SetSelectedDataSource sets the value of the selectedDataSource property.
// The currently selected audio data source for the port.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/selecteddatasource
func (a_ AudioSessionPortDescription) SetSelectedDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedDataSource:"), value)
}

// A system-assigned unique identifier (UID) for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/uid
func (a_ AudioSessionPortDescription) Uid() string {
	rv := objc.Send[string](a_.ID, objc.Sel("uid"))
	return rv
}


// SetUid sets the value of the uid property.
// A system-assigned unique identifier (UID) for the port.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/uid
func (a_ AudioSessionPortDescription) SetUid(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUid:"), objc.String(value))
}



