// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that describes the representation of an audio format.
//
// The class wraps Core Audio’s , and includes convenience initializers and accessors for common formats, including Core Audio’s standard deinterleaved 32-bit floating point format. Instances of this class are immutable.
//
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


// A dictionary that represents the format as a dictionary using audio setting keys.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/settings
func (a_ AudioFormat) Settings() string {
	rv := objc.Send[string](a_.ID, objc.Sel("settings"))
	return rv
}


// SetSettings sets the value of the settings property.
// A dictionary that represents the format as a dictionary using audio setting keys.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/settings
func (a_ AudioFormat) SetSettings(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSettings:"), objc.String(value))
}

// A Boolean value that indicates whether the samples mix into one stream.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) IsInterleaved() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInterleaved"))
	return rv
}


// SetIsInterleaved sets the value of the isInterleaved property.
// A Boolean value that indicates whether the samples mix into one stream.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isinterleaved
func (a_ AudioFormat) SetIsInterleaved(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsInterleaved:"), value)
}

// The audio format properties of a stream of audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/streamdescription
func (a_ AudioFormat) StreamDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("streamDescription"))
	return rv
}


// SetStreamDescription sets the value of the streamDescription property.
// The audio format properties of a stream of audio data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/streamdescription
func (a_ AudioFormat) SetStreamDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStreamDescription:"), value)
}

// The audio format sampling rate, in hertz.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/samplerate
func (a_ AudioFormat) SampleRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleRate"))
	return rv
}


// SetSampleRate sets the value of the sampleRate property.
// The audio format sampling rate, in hertz.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/samplerate
func (a_ AudioFormat) SetSampleRate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRate:"), value)
}

// The underlying audio channel layout.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channellayout
func (a_ AudioFormat) ChannelLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("channelLayout"))
	return rv
}


// SetChannelLayout sets the value of the channelLayout property.
// The underlying audio channel layout.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channellayout
func (a_ AudioFormat) SetChannelLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelLayout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioFormat) AVChannelLayoutKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}

// The common format identifier instance.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/commonformat
func (a_ AudioFormat) CommonFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("commonFormat"))
	return rv
}


// SetCommonFormat sets the value of the commonFormat property.
// The common format identifier instance.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/commonformat
func (a_ AudioFormat) SetCommonFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCommonFormat:"), value)
}

// An object that contains metadata that encoders and decoders require.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/magiccookie
func (a_ AudioFormat) MagicCookie() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("magicCookie"))
	return rv
}


// SetMagicCookie sets the value of the magicCookie property.
// An object that contains metadata that encoders and decoders require.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/magiccookie
func (a_ AudioFormat) SetMagicCookie(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}

// The number of channels of audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channelcount
func (a_ AudioFormat) ChannelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("channelCount"))
	return rv
}


// SetChannelCount sets the value of the channelCount property.
// The number of channels of audio data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/channelcount
func (a_ AudioFormat) SetChannelCount(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelCount:"), value)
}

// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) IsStandard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStandard"))
	return rv
}


// SetIsStandard sets the value of the isStandard property.
// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioformat/isstandard
func (a_ AudioFormat) SetIsStandard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStandard:"), value)
}

// The audio format description to use with Core Media APIs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/formatDescription
func (a_ AudioFormat) FormatDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("formatDescription"))
	return rv
}



