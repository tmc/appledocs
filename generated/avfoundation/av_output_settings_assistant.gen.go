// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AudioSettings() objc.IObject /* cross-framework: NSString */
	SetAudioSettings(value objc.IObject /* cross-framework: NSString */)
	OutputFileType() objc.IObject /* cross-framework: FileType */
	SetOutputFileType(value objc.IObject /* cross-framework: FileType */)
	SourceAudioFormat() AudioFormatDescription /* not a class type */
	SetSourceAudioFormat(value AudioFormatDescription /* not a class type */)
	SourceVideoAverageFrameDuration() Time /* not a class type */
	SetSourceVideoAverageFrameDuration(value Time /* not a class type */)
	SourceVideoFormat() VideoFormatDescription /* not a class type */
	SetSourceVideoFormat(value VideoFormatDescription /* not a class type */)
	SourceVideoMinFrameDuration() Time /* not a class type */
	SetSourceVideoMinFrameDuration(value Time /* not a class type */)
	VideoSettings() objc.IObject /* cross-framework: NSString */
	SetVideoSettings(value objc.IObject /* cross-framework: NSString */)
	// methods:
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
func (o_ OutputSettingsAssistant) AudioSettings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("audioSettings"))
	return rv
}


// An audio settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/audiosettings
func (o_ OutputSettingsAssistant) SetAudioSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAudioSettings:"), value)
}


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/outputfiletype
func (o_ OutputSettingsAssistant) OutputFileType() objc.IObject /* cross-framework: FileType */ {
	rv := objc.Send[FileType](o_.ID, objc.Sel("outputFileType"))
	return rv
}


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/outputfiletype
func (o_ OutputSettingsAssistant) SetOutputFileType(value objc.IObject /* cross-framework: FileType */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOutputFileType:"), value)
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourceaudioformat
func (o_ OutputSettingsAssistant) SourceAudioFormat() AudioFormatDescription /* not a class type */ {
	rv := objc.Send[AudioFormatDescription](o_.ID, objc.Sel("sourceAudioFormat"))
	return rv
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourceaudioformat
func (o_ OutputSettingsAssistant) SetSourceAudioFormat(value AudioFormatDescription /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceAudioFormat:"), value)
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoaverageframeduration
func (o_ OutputSettingsAssistant) SourceVideoAverageFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](o_.ID, objc.Sel("sourceVideoAverageFrameDuration"))
	return rv
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoaverageframeduration
func (o_ OutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoformat
func (o_ OutputSettingsAssistant) SourceVideoFormat() VideoFormatDescription /* not a class type */ {
	rv := objc.Send[VideoFormatDescription](o_.ID, objc.Sel("sourceVideoFormat"))
	return rv
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideoformat
func (o_ OutputSettingsAssistant) SetSourceVideoFormat(value VideoFormatDescription /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoFormat:"), value)
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideominframeduration
func (o_ OutputSettingsAssistant) SourceVideoMinFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](o_.ID, objc.Sel("sourceVideoMinFrameDuration"))
	return rv
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/sourcevideominframeduration
func (o_ OutputSettingsAssistant) SetSourceVideoMinFrameDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/videosettings
func (o_ OutputSettingsAssistant) VideoSettings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("videoSettings"))
	return rv
}


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/videosettings
func (o_ OutputSettingsAssistant) SetVideoSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVideoSettings:"), value)
}



