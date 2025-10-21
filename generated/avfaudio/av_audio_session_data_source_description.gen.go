// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionDataSourceDescription] class.
var (
	AudioSessionDataSourceDescriptionClass     _AudioSessionDataSourceDescriptionClass
	AudioSessionDataSourceDescriptionClassOnce sync.Once
)

func getAudioSessionDataSourceDescriptionClass() _AudioSessionDataSourceDescriptionClass {
	AudioSessionDataSourceDescriptionClassOnce.Do(func() {
		AudioSessionDataSourceDescriptionClass = _AudioSessionDataSourceDescriptionClass{objc.GetClass("AVAudioSessionDataSourceDescription")}
	})
	return AudioSessionDataSourceDescriptionClass
}

type _AudioSessionDataSourceDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionDataSourceDescription] class.
type IAudioSessionDataSourceDescription interface {
	objectivec.IObject
	SetPreferredPolarPatternError(pattern unsafe.Pointer, outError unsafe.Pointer) bool
}

// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
//
// You obtain data source descriptions from the shared object or the objects corresponding to its input and output ports. Only built-in microphone ports on certain devices support the location, orientation, and polar pattern properties. If a port doesn’t support these features, the value of its property is . This class is especially useful for differentiating between microphone configurations on devices having more than one built-in microphone. Such devices may also support signal processing features for spatial filtering, or , in which the system makes the device more sensitive to audio signals from a particular direction. See for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription
type AudioSessionDataSourceDescription struct {
	objectivec.Object
}

// AudioSessionDataSourceDescriptionFrom constructs a [AudioSessionDataSourceDescription] from an unsafe.Pointer.
//
// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
func AudioSessionDataSourceDescriptionFrom(ptr unsafe.Pointer) AudioSessionDataSourceDescription {
	return AudioSessionDataSourceDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionDataSourceDescriptionClass) Alloc() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionDataSourceDescriptionClass) New() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionDataSourceDescription) Init() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionDataSourceDescription) Autorelease() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionDataSourceDescription creates a new AudioSessionDataSourceDescription instance.
func NewAudioSessionDataSourceDescription() AudioSessionDataSourceDescription {
	return getAudioSessionDataSourceDescriptionClass().New()
}


// Selects the preferred directivity configuration for the data source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/setPreferredPolarPattern(_:)
func (a_ AudioSessionDataSourceDescription) SetPreferredPolarPatternError(pattern unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredPolarPattern:error:"), pattern, outError)
	return rv
}

// The preferred directivity configuration for the data source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/preferredPolarPattern
func (a_ AudioSessionDataSourceDescription) PreferredPolarPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredPolarPattern"))
	return rv
}

// The data source’s active polar pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/selectedPolarPattern
func (a_ AudioSessionDataSourceDescription) SelectedPolarPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("selectedPolarPattern"))
	return rv
}

// The set of directivity configurations supported by the data source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/supportedPolarPatterns
func (a_ AudioSessionDataSourceDescription) SupportedPolarPatterns() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedPolarPatterns"))
	return rv
}



