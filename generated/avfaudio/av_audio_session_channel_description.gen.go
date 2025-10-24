// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioSessionChannelDescription] class.
var (
	AudioSessionChannelDescriptionClass     _AudioSessionChannelDescriptionClass
	AudioSessionChannelDescriptionClassOnce sync.Once
)

func getAudioSessionChannelDescriptionClass() _AudioSessionChannelDescriptionClass {
	AudioSessionChannelDescriptionClassOnce.Do(func() {
		AudioSessionChannelDescriptionClass = _AudioSessionChannelDescriptionClass{objc.GetClass("AVAudioSessionChannelDescription")}
	})
	return AudioSessionChannelDescriptionClass
}

type _AudioSessionChannelDescriptionClass struct {
	class objc.Class
}





// An interface definition for the [AudioSessionChannelDescription] class.
type IAudioSessionChannelDescription interface {
	objectivec.IObject
	

	// properties:
	Channels() IAVAudioSessionChannelDescription
	SetChannels(value IAVAudioSessionChannelDescription)
	HasHardwareVoiceCallProcessing() bool
	SetHasHardwareVoiceCallProcessing(value bool)
	IsSpatialAudioEnabled() bool
	SetIsSpatialAudioEnabled(value bool)
	PortName() objc.IObject /* cross-framework: NSString */
	SetPortName(value objc.IObject /* cross-framework: NSString */)
	PortType() foundation.Port
	SetPortType(value foundation.Port)
	Uid() objc.IObject /* cross-framework: NSString */
	SetUid(value objc.IObject /* cross-framework: NSString */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioSessionChannelDescriptionClass) Alloc() AudioSessionChannelDescription {
	rv := objc.Send[AudioSessionChannelDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionChannelDescriptionClass) New() AudioSessionChannelDescription {
	rv := objc.Send[AudioSessionChannelDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionChannelDescription) Init() AudioSessionChannelDescription {
	rv := objc.Send[AudioSessionChannelDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionChannelDescription) Autorelease() AudioSessionChannelDescription {
	rv := objc.Send[AudioSessionChannelDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionChannelDescription creates a new AudioSessionChannelDescription instance.
func NewAudioSessionChannelDescription() AudioSessionChannelDescription {
	return getAudioSessionChannelDescriptionClass().New()
}





// A class that describes a hardware channel on the current device.


// A class that describes a hardware channel on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionChannelDescription
type AudioSessionChannelDescription struct {
	objectivec.Object
}

// AudioSessionChannelDescriptionFrom constructs a [AudioSessionChannelDescription] from an unsafe.Pointer.
//
// A class that describes a hardware channel on the current device.
func AudioSessionChannelDescriptionFrom(ptr unsafe.Pointer) AudioSessionChannelDescription {
	return AudioSessionChannelDescription{objectivec.Object{objc.ID(ptr)}}
}

























// An array of channel objects that describe the port’s input or output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/channels
func (a_ AudioSessionChannelDescription) Channels() IAVAudioSessionChannelDescription {
	rv := objc.Send[AudioSessionChannelDescription](a_.ID, objc.Sel("channels"))
	return rv
}


// An array of channel objects that describe the port’s input or output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/channels
func (a_ AudioSessionChannelDescription) SetChannels(value IAVAudioSessionChannelDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannels:"), value)
}


// A Boolean value that indicates whether the associated hardware port has built-in processing for two-way voice communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/hashardwarevoicecallprocessing
func (a_ AudioSessionChannelDescription) HasHardwareVoiceCallProcessing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasHardwareVoiceCallProcessing"))
	return rv
}


// A Boolean value that indicates whether the associated hardware port has built-in processing for two-way voice communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/hashardwarevoicecallprocessing
func (a_ AudioSessionChannelDescription) SetHasHardwareVoiceCallProcessing(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHasHardwareVoiceCallProcessing:"), value)
}


// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionChannelDescription) IsSpatialAudioEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSpatialAudioEnabled"))
	return rv
}


// A Boolean value that indicates whether the port supports spatial audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/isspatialaudioenabled
func (a_ AudioSessionChannelDescription) SetIsSpatialAudioEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSpatialAudioEnabled:"), value)
}


// A descriptive name for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/portname
func (a_ AudioSessionChannelDescription) PortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("portName"))
	return rv
}


// A descriptive name for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/portname
func (a_ AudioSessionChannelDescription) SetPortName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPortName:"), value)
}


// The type of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/porttype
func (a_ AudioSessionChannelDescription) PortType() foundation.Port {
	rv := objc.Send[foundation.Port](a_.ID, objc.Sel("portType"))
	return rv
}


// The type of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/porttype
func (a_ AudioSessionChannelDescription) SetPortType(value foundation.Port) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPortType:"), value)
}


// A system-assigned unique identifier (UID) for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/uid
func (a_ AudioSessionChannelDescription) Uid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("uid"))
	return rv
}


// A system-assigned unique identifier (UID) for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/uid
func (a_ AudioSessionChannelDescription) SetUid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUid:"), value)
}







