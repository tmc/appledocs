// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioFormat] class.
type IAudioFormat interface {
	objectivec.IObject
	// properties:
	ChannelCount() objc.IObject /* cross-framework: AudioChannelCount */
	SetChannelCount(value objc.IObject /* cross-framework: AudioChannelCount */)
	ChannelLayout() objc.IObject /* cross-framework: AudioChannelLayout */
	SetChannelLayout(value objc.IObject /* cross-framework: AudioChannelLayout */)
	CommonFormat() AudioCommonFormat /* not a class type */
	SetCommonFormat(value AudioCommonFormat /* not a class type */)
	FormatDescription() AudioFormatDescription /* not a class type */
	SetFormatDescription(value AudioFormatDescription /* not a class type */)
	IsInterleaved() bool
	SetIsInterleaved(value bool)
	IsStandard() bool
	SetIsStandard(value bool)
	MagicCookie() objc.IObject /* cross-framework: Data */
	SetMagicCookie(value objc.IObject /* cross-framework: Data */)
	SampleRate() float64
	SetSampleRate(value float64)
	Settings() objc.IObject /* cross-framework: NSString */
	SetSettings(value objc.IObject /* cross-framework: NSString */)
	StreamDescription() unsafe.Pointer
	SetStreamDescription(value unsafe.Pointer)
	AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _AudioFormatClass) Alloc() AudioFormat {
	rv := objc.Send[AudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channelcount
func (a_ AudioFormat) ChannelCount() objc.IObject /* cross-framework: AudioChannelCount */ {
	rv := objc.Send[AudioChannelCount](a_.ID, objc.Sel("channelCount"))
	return rv
}


// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channelcount
func (a_ AudioFormat) SetChannelCount(value objc.IObject /* cross-framework: AudioChannelCount */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelCount:"), value)
}


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channellayout
func (a_ AudioFormat) ChannelLayout() objc.IObject /* cross-framework: AudioChannelLayout */ {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("channelLayout"))
	return rv
}


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channellayout
func (a_ AudioFormat) SetChannelLayout(value objc.IObject /* cross-framework: AudioChannelLayout */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelLayout:"), value)
}


// The common format identifier instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/commonformat
func (a_ AudioFormat) CommonFormat() AudioCommonFormat /* not a class type */ {
	rv := objc.Send[AudioCommonFormat](a_.ID, objc.Sel("commonFormat"))
	return rv
}


// The common format identifier instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/commonformat
func (a_ AudioFormat) SetCommonFormat(value AudioCommonFormat /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCommonFormat:"), value)
}


// The audio format description to use with Core Media APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/formatdescription
func (a_ AudioFormat) FormatDescription() AudioFormatDescription /* not a class type */ {
	rv := objc.Send[AudioFormatDescription](a_.ID, objc.Sel("formatDescription"))
	return rv
}


// The audio format description to use with Core Media APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/formatdescription
func (a_ AudioFormat) SetFormatDescription(value AudioFormatDescription /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormatDescription:"), value)
}


// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) IsInterleaved() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInterleaved"))
	return rv
}


// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) SetIsInterleaved(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInterleaved:"), value)
}


// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) IsStandard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStandard"))
	return rv
}


// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) SetIsStandard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStandard:"), value)
}


// An object that contains metadata that encoders and decoders require.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/magiccookie
func (a_ AudioFormat) MagicCookie() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("magicCookie"))
	return rv
}


// An object that contains metadata that encoders and decoders require.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/magiccookie
func (a_ AudioFormat) SetMagicCookie(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}


// The audio format sampling rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/samplerate
func (a_ AudioFormat) SampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("sampleRate"))
	return rv
}


// The audio format sampling rate, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/samplerate
func (a_ AudioFormat) SetSampleRate(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRate:"), value)
}


// A dictionary that represents the format as a dictionary using audio setting keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/settings
func (a_ AudioFormat) Settings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("settings"))
	return rv
}


// A dictionary that represents the format as a dictionary using audio setting keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/settings
func (a_ AudioFormat) SetSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSettings:"), value)
}


// The audio format properties of a stream of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/streamdescription
func (a_ AudioFormat) StreamDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("streamDescription"))
	return rv
}


// The audio format properties of a stream of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/streamdescription
func (a_ AudioFormat) SetStreamDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStreamDescription:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioFormat) AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}



