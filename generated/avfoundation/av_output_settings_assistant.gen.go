// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVOutputSettingsAssistant */


/* debug [class_header]: Header for AVOutputSettingsAssistant */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OutputSettingsAssistant */
// An interface definition for the [OutputSettingsAssistant] class.
type IOutputSettingsAssistant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OutputSettingsAssistant */
	// properties:
	AudioSettings() foundation.IDictionary
	OutputFileType() FileType /* typedef */
	SourceAudioFormat() AudioFormatDescriptionRef /* not a class type */
	SetSourceAudioFormat(value AudioFormatDescriptionRef /* not a class type */)
	SourceVideoAverageFrameDuration() objc.IObject /* cross-framework: Time */
	SetSourceVideoAverageFrameDuration(value objc.IObject /* cross-framework: Time */)
	SourceVideoFormat() VideoFormatDescriptionRef /* not a class type */
	SetSourceVideoFormat(value VideoFormatDescriptionRef /* not a class type */)
	SourceVideoMinFrameDuration() objc.IObject /* cross-framework: Time */
	SetSourceVideoMinFrameDuration(value objc.IObject /* cross-framework: Time */)
	VideoSettings() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OutputSettingsAssistant */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OutputSettingsAssistant */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OutputSettingsAssistant */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OutputSettingsAssistant */

// Creates an output setting assistant with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/init(preset:)
func NewOutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset /* typedef */) OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(getOutputSettingsAssistantClass().class), objc.Sel("outputSettingsAssistantWithPreset:"), presetIdentifier)
	return rv
}/* debug [class_init_methods/constructor]: NewOutputSettingsAssistantWithPreset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OutputSettingsAssistant */

// Returns an array of preset values to use to initialize an output settings assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/availableOutputSettingsPresets()
func (oc _OutputSettingsAssistantClass) AvailableOutputSettingsPresets() []string {
	rv := objc.Send[[]string](objc.ID(oc.class), objc.Sel("availableOutputSettingsPresets"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AvailableOutputSettingsPresets) */


// Creates an output setting assistant with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/init(preset:)
func (oc _OutputSettingsAssistantClass) OutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("outputSettingsAssistantWithPreset:"), presetIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OutputSettingsAssistantWithPreset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OutputSettingsAssistant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OutputSettingsAssistant */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OutputSettingsAssistant */

// An audio settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/audioSettings
func (o_ OutputSettingsAssistant) AudioSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](o_.ID, objc.Sel("audioSettings"))
	return rv
}/* debug [instance_properties/getter]: audioSettings */


// A uniform type identifier (UTI) that indicates the type of file to write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/outputFileType
func (o_ OutputSettingsAssistant) OutputFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("outputFileType"))
	return rv
}/* debug [instance_properties/getter]: outputFileType */


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceAudioFormat
func (o_ OutputSettingsAssistant) SourceAudioFormat() AudioFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[AudioFormatDescriptionRef](o_.ID, objc.Sel("sourceAudioFormat"))
	return rv
}/* debug [instance_properties/getter]: sourceAudioFormat */


// The format of the source audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceAudioFormat
func (o_ OutputSettingsAssistant) SetSourceAudioFormat(value AudioFormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceAudioFormat:"), value)
}/* debug [instance_properties/setter]: sourceAudioFormat */


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoAverageFrameDuration
func (o_ OutputSettingsAssistant) SourceVideoAverageFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](o_.ID, objc.Sel("sourceVideoAverageFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: sourceVideoAverageFrameDuration */


// A time value that describes the average frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoAverageFrameDuration
func (o_ OutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}/* debug [instance_properties/setter]: sourceVideoAverageFrameDuration */


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoFormat
func (o_ OutputSettingsAssistant) SourceVideoFormat() VideoFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[VideoFormatDescriptionRef](o_.ID, objc.Sel("sourceVideoFormat"))
	return rv
}/* debug [instance_properties/getter]: sourceVideoFormat */


// The format of the source video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoFormat
func (o_ OutputSettingsAssistant) SetSourceVideoFormat(value VideoFormatDescriptionRef /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoFormat:"), value)
}/* debug [instance_properties/setter]: sourceVideoFormat */


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoMinFrameDuration
func (o_ OutputSettingsAssistant) SourceVideoMinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](o_.ID, objc.Sel("sourceVideoMinFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: sourceVideoMinFrameDuration */


// A time value that describes the minimum frame duration of the video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoMinFrameDuration
func (o_ OutputSettingsAssistant) SetSourceVideoMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}/* debug [instance_properties/setter]: sourceVideoMinFrameDuration */


// A video settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/videoSettings
func (o_ OutputSettingsAssistant) VideoSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](o_.ID, objc.Sel("videoSettings"))
	return rv
}/* debug [instance_properties/getter]: videoSettings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVOutputSettingsAssistant */


