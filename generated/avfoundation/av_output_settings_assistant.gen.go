// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OutputSettingsAssistant] class.
var (
	OutputSettingsAssistantClass     _OutputSettingsAssistantClass
	OutputSettingsAssistantClassOnce sync.Once
)

func getOutputSettingsAssistantClass() _OutputSettingsAssistantClass {
	OutputSettingsAssistantClassOnce.Do(func() {
		OutputSettingsAssistantClass = _OutputSettingsAssistantClass{objc.GetClass("AVOutputSettingsAssistant")}
	})
	return OutputSettingsAssistantClass
}

type _OutputSettingsAssistantClass struct {
	class objc.Class
}

// An interface definition for the [OutputSettingsAssistant] class.
type IOutputSettingsAssistant interface {
	objectivec.IObject
	AudioSettings() string
	SetAudioSettings(value string)
	OutputFileType() unsafe.Pointer
	SetOutputFileType(value unsafe.Pointer)
	SourceAudioFormat() unsafe.Pointer
	SetSourceAudioFormat(value unsafe.Pointer)
	SourceVideoAverageFrameDuration() unsafe.Pointer
	SetSourceVideoAverageFrameDuration(value unsafe.Pointer)
	SourceVideoFormat() unsafe.Pointer
	SetSourceVideoFormat(value unsafe.Pointer)
	SourceVideoMinFrameDuration() unsafe.Pointer
	SetSourceVideoMinFrameDuration(value unsafe.Pointer)
	VideoSettings() string
	SetVideoSettings(value string)
}

// An object that builds audio and video output settings dictionaries.
//
// Use an output settings assistant to create the audio and video settings that you use to configure instances of and . You create an assistant with a specific preset configuration, such as or . You can accept the settings dictionaries as is to generate a file that conforms to the criteria that the preset implies. You may also use the dictionaries it generates as a base configuration that you can customize as you require. Providing the assistant additional details about your source media helps it generate more complete results. For example, setting a value for its property ensures that the assistant generates settings that don’t scale up video frames from a smaller size.


// An object that builds audio and video output settings dictionaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant
type OutputSettingsAssistant struct {
	objectivec.Object
}

// OutputSettingsAssistantFrom constructs a [OutputSettingsAssistant] from an unsafe.Pointer.
//
// An object that builds audio and video output settings dictionaries.
func OutputSettingsAssistantFrom(ptr unsafe.Pointer) OutputSettingsAssistant {
	return OutputSettingsAssistant{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OutputSettingsAssistantClass) Alloc() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OutputSettingsAssistantClass) New() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OutputSettingsAssistant) Init() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OutputSettingsAssistant) Autorelease() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutputSettingsAssistant creates a new OutputSettingsAssistant instance.
func NewOutputSettingsAssistant() OutputSettingsAssistant {
	return getOutputSettingsAssistantClass().New()
}



// An audio settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/audiosettings
func (o_ OutputSettingsAssistant) AudioSettings() string {
	rv := objc.Send[string](o_.ID, objc.Sel("audioSettings"))
	return rv
}


// An audio settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/audiosettings
func (o_ OutputSettingsAssistant) SetAudioSettings(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAudioSettings:"), objc.String(value))
}


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/outputfiletype
func (o_ OutputSettingsAssistant) OutputFileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("outputFileType"))
	return rv
}


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/outputfiletype
func (o_ OutputSettingsAssistant) SetOutputFileType(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOutputFileType:"), value)
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourceaudioformat
func (o_ OutputSettingsAssistant) SourceAudioFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sourceAudioFormat"))
	return rv
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourceaudioformat
func (o_ OutputSettingsAssistant) SetSourceAudioFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceAudioFormat:"), value)
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoaverageframeduration
func (o_ OutputSettingsAssistant) SourceVideoAverageFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sourceVideoAverageFrameDuration"))
	return rv
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoaverageframeduration
func (o_ OutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoformat
func (o_ OutputSettingsAssistant) SourceVideoFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sourceVideoFormat"))
	return rv
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoformat
func (o_ OutputSettingsAssistant) SetSourceVideoFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoFormat:"), value)
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideominframeduration
func (o_ OutputSettingsAssistant) SourceVideoMinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sourceVideoMinFrameDuration"))
	return rv
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideominframeduration
func (o_ OutputSettingsAssistant) SetSourceVideoMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/videosettings
func (o_ OutputSettingsAssistant) VideoSettings() string {
	rv := objc.Send[string](o_.ID, objc.Sel("videoSettings"))
	return rv
}


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/videosettings
func (o_ OutputSettingsAssistant) SetVideoSettings(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVideoSettings:"), objc.String(value))
}



