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
	

	// properties:
	AudioSettings() foundation.IDictionary
	OutputFileType() FileType
	SourceAudioFormat() AudioFormatDescriptionRef /* not a class type */
	SetSourceAudioFormat(value AudioFormatDescriptionRef /* not a class type */)
	SourceVideoAverageFrameDuration() objectivec.IObject
	SetSourceVideoAverageFrameDuration(value objectivec.IObject)
	SourceVideoFormat() VideoFormatDescriptionRef /* not a class type */
	SetSourceVideoFormat(value VideoFormatDescriptionRef /* not a class type */)
	SourceVideoMinFrameDuration() objectivec.IObject
	SetSourceVideoMinFrameDuration(value objectivec.IObject)
	VideoSettings() foundation.IDictionary


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OutputSettingsAssistantClass) Alloc() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an output setting assistant with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/init(preset:)
func NewOutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset) OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(getOutputSettingsAssistantClass().class), objc.Sel("outputSettingsAssistantWithPreset:"), presetIdentifier)
	return rv
}







// Returns an array of preset values to use to initialize an output settings assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/availableOutputSettingsPresets()
func (oc _OutputSettingsAssistantClass) AvailableOutputSettingsPresets() []string {
	rv := objc.Send[[]string](objc.ID(oc.class), objc.Sel("availableOutputSettingsPresets"))
	return rv
}


// Creates an output setting assistant with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/init(preset:)
func (oc _OutputSettingsAssistantClass) OutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("outputSettingsAssistantWithPreset:"), presetIdentifier)
	return rv
}

















// An audio settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/audioSettings
func (o_ OutputSettingsAssistant) AudioSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](o_.ID, objc.Sel("audioSettings"))
	return rv
}


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/outputFileType
func (o_ OutputSettingsAssistant) OutputFileType() FileType {
	rv := objc.Send[FileType](o_.ID, objc.Sel("outputFileType"))
	return rv
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceAudioFormat
func (o_ OutputSettingsAssistant) SourceAudioFormat() AudioFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[AudioFormatDescriptionRef](o_.ID, objc.Sel("sourceAudioFormat"))
	return rv
}


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceAudioFormat
func (o_ OutputSettingsAssistant) SetSourceAudioFormat(value AudioFormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceAudioFormat:"), value)
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoAverageFrameDuration
func (o_ OutputSettingsAssistant) SourceVideoAverageFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("sourceVideoAverageFrameDuration"))
	return rv
}


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoAverageFrameDuration
func (o_ OutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoFormat
func (o_ OutputSettingsAssistant) SourceVideoFormat() VideoFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[VideoFormatDescriptionRef](o_.ID, objc.Sel("sourceVideoFormat"))
	return rv
}


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoFormat
func (o_ OutputSettingsAssistant) SetSourceVideoFormat(value VideoFormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoFormat:"), value)
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoMinFrameDuration
func (o_ OutputSettingsAssistant) SourceVideoMinFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("sourceVideoMinFrameDuration"))
	return rv
}


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoMinFrameDuration
func (o_ OutputSettingsAssistant) SetSourceVideoMinFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/videoSettings
func (o_ OutputSettingsAssistant) VideoSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](o_.ID, objc.Sel("videoSettings"))
	return rv
}







