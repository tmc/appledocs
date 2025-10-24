// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioFormat */


/* debug [class_header]: Header for AVAudioFormat */
// The class instance for the [AudioFormat] class.
var (
	AudioFormatClass     _AudioFormatClass
	AudioFormatClassOnce sync.Once
)

func getAudioFormatClass() _AudioFormatClass {
	AudioFormatClassOnce.Do(func() {
		AudioFormatClass = _AudioFormatClass{objc.GetClass("AVAudioFormat")}
	})
	return AudioFormatClass
}

type _AudioFormatClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioFormat */
// An interface definition for the [AudioFormat] class.
type IAudioFormat interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioFormat */
	// properties:
	ChannelCount() AudioChannelCount /* typedef */
	ChannelLayout() IAVAudioChannelLayout
	CommonFormat() AudioCommonFormat
	FormatDescription() AudioFormatDescriptionRef /* not a class type */
	Interleaved() bool
	Standard() bool
	MagicCookie() objc.IObject /* cross-framework: NSData */
	SetMagicCookie(value objc.IObject /* cross-framework: NSData */)
	SampleRate() float64
	Settings() foundation.IDictionary
	StreamDescription() objc.IObject
	IsInterleaved() bool
	SetIsInterleaved(value bool)
	IsStandard() bool
	SetIsStandard(value bool)
	AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioFormat */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioFormat */
// Alloc allocates a new instance without initialization.
func (ac _AudioFormatClass) Alloc() AudioFormat {
	rv := objc.Send[AudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioFormatClass) New() AudioFormat {
	rv := objc.Send[AudioFormat](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioFormat) Init() AudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioFormat) Autorelease() AudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioFormat creates a new AudioFormat instance.
func NewAudioFormat() AudioFormat {
	return getAudioFormatClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioFormat */
// An object that describes the representation of an audio format.
//
// The class wraps Core Audio’s , and includes convenience initializers and accessors for common formats, including Core Audio’s standard deinterleaved 32-bit floating point format. Instances of this class are immutable.


// An object that describes the representation of an audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
type AudioFormat struct {
	objectivec.Object
}

// AudioFormatFrom constructs a [AudioFormat] from an unsafe.Pointer.
//
// An object that describes the representation of an audio format.
func AudioFormatFrom(ptr unsafe.Pointer) AudioFormat {
	return AudioFormat{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioFormat */

// Creates an audio format instance as a deinterleaved float with the specified sample rate and channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channelLayout:)
func NewAudioFormatStandardFormatWithSampleRateChannelLayout(sampleRate float64, layout IAVAudioChannelLayout) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initStandardFormatWithSampleRate:channelLayout:"), sampleRate, layout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatStandardFormatWithSampleRateChannelLayout */


// Creates an audio format instance with the specified sample rate and channel count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channels:)
func NewAudioFormatStandardFormatWithSampleRateChannels(sampleRate float64, channels AudioChannelCount /* typedef */) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initStandardFormatWithSampleRate:channels:"), sampleRate, channels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatStandardFormatWithSampleRateChannels */


// Creates an audio format instance from a Core Media audio format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(cmAudioFormatDescription:)
func NewAudioFormatWithCMAudioFormatDescription(formatDescription AudioFormatDescriptionRef /* not a class type */) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithCMAudioFormatDescription:"), formatDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithCMAudioFormatDescription */


// Creates an audio format instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:channels:interleaved:)
func NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved(format AudioCommonFormat, sampleRate float64, channels AudioChannelCount /* typedef */, interleaved bool) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithCommonFormat:sampleRate:channels:interleaved:"), format, sampleRate, channels, interleaved)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved */


// Creates an audio format instance with the specified audio format, sample rate, interleaved state, and channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:interleaved:channelLayout:)
func NewAudioFormatWithCommonFormatSampleRateInterleavedChannelLayout(format AudioCommonFormat, sampleRate float64, interleaved bool, layout IAVAudioChannelLayout) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithCommonFormat:sampleRate:interleaved:channelLayout:"), format, sampleRate, interleaved, layout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithCommonFormatSampleRateInterleavedChannelLayout */


// Creates an audio format instance using the specified settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(settings:)
func NewAudioFormatWithSettings(settings foundation.IDictionary) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithSettings:"), settings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithSettings */


// Creates an audio format instance from a stream description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:)
func NewAudioFormatWithStreamDescription(asbd objc.IObject) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithStreamDescription:"), asbd)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithStreamDescription */


// Creates an audio format instance from a stream description and channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:channelLayout:)
func NewAudioFormatWithStreamDescriptionChannelLayout(asbd objc.IObject, layout IAVAudioChannelLayout) AudioFormat {
	instance := getAudioFormatClass().Alloc()
	rv := objc.Send[AudioFormat](instance.ID, objc.Sel("initWithStreamDescription:channelLayout:"), asbd, layout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioFormatWithStreamDescriptionChannelLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioFormat */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioFormat */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioFormat */

// Indicates whether the audio format instance and a specified object are functionally equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/isEqual(_:)
func (a_ AudioFormat) IsEqual(object objc.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqual:"), object)
	return rv
}/* debug [instance_methods/method]: IsEqual */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioFormat */

// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/channelCount
func (a_ AudioFormat) ChannelCount() AudioChannelCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("channelCount"))
	return rv
}/* debug [instance_properties/getter]: channelCount */


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/channelLayout
func (a_ AudioFormat) ChannelLayout() IAVAudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("channelLayout"))
	return rv
}/* debug [instance_properties/getter]: channelLayout */


// The common format identifier instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/commonFormat
func (a_ AudioFormat) CommonFormat() AudioCommonFormat {
	rv := objc.Send[AudioCommonFormat](a_.ID, objc.Sel("commonFormat"))
	return rv
}/* debug [instance_properties/getter]: commonFormat */


// The audio format description to use with Core Media APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/formatDescription
func (a_ AudioFormat) FormatDescription() AudioFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[AudioFormatDescriptionRef](a_.ID, objc.Sel("formatDescription"))
	return rv
}/* debug [instance_properties/getter]: formatDescription */


// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/isInterleaved
func (a_ AudioFormat) Interleaved() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("interleaved"))
	return rv
}/* debug [instance_properties/getter]: interleaved */


// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/isStandard
func (a_ AudioFormat) Standard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("standard"))
	return rv
}/* debug [instance_properties/getter]: standard */


// An object that contains metadata that encoders and decoders require.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/magicCookie
func (a_ AudioFormat) MagicCookie() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("magicCookie"))
	return rv
}/* debug [instance_properties/getter]: magicCookie */


// An object that contains metadata that encoders and decoders require.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/magicCookie
func (a_ AudioFormat) SetMagicCookie(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}/* debug [instance_properties/setter]: magicCookie */


// The audio format sampling rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/sampleRate
func (a_ AudioFormat) SampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("sampleRate"))
	return rv
}/* debug [instance_properties/getter]: sampleRate */


// A dictionary that represents the format as a dictionary using audio setting keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/settings
func (a_ AudioFormat) Settings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("settings"))
	return rv
}/* debug [instance_properties/getter]: settings */


// The audio format properties of a stream of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/streamDescription
func (a_ AudioFormat) StreamDescription() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("streamDescription"))
	return rv
}/* debug [instance_properties/getter]: streamDescription */


// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) IsInterleaved() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInterleaved"))
	return rv
}/* debug [instance_properties/getter]: isInterleaved */


// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) SetIsInterleaved(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInterleaved:"), value)
}/* debug [instance_properties/setter]: isInterleaved */


// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) IsStandard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStandard"))
	return rv
}/* debug [instance_properties/getter]: isStandard */


// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) SetIsStandard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStandard:"), value)
}/* debug [instance_properties/setter]: isStandard */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioFormat) AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}/* debug [instance_properties/getter]: AVChannelLayoutKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioFormat */


