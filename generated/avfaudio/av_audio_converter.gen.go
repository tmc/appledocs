// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioConverter] class.
var (
	AudioConverterClass     _AudioConverterClass
	AudioConverterClassOnce sync.Once
)

func getAudioConverterClass() _AudioConverterClass {
	AudioConverterClassOnce.Do(func() {
		AudioConverterClass = _AudioConverterClass{objc.GetClass("AVAudioConverter")}
	})
	return AudioConverterClass
}

type _AudioConverterClass struct {
	class objc.Class
}

// An interface definition for the [AudioConverter] class.
type IAudioConverter interface {
	objectivec.IObject
	ApplicableEncodeBitRates() foundation.Number
	SetApplicableEncodeBitRates(value foundation.Number)
	ApplicableEncodeSampleRates() foundation.Number
	SetApplicableEncodeSampleRates(value foundation.Number)
	AudioSyncPacketFrequency() int
	SetAudioSyncPacketFrequency(value int)
	AvailableEncodeBitRates() foundation.Number
	SetAvailableEncodeBitRates(value foundation.Number)
	AvailableEncodeChannelLayoutTags() foundation.Number
	SetAvailableEncodeChannelLayoutTags(value foundation.Number)
	AvailableEncodeSampleRates() foundation.Number
	SetAvailableEncodeSampleRates(value foundation.Number)
	BitRate() int
	SetBitRate(value int)
	BitRateStrategy() string
	SetBitRateStrategy(value string)
	ChannelMap() foundation.Number
	SetChannelMap(value foundation.Number)
	ContentSource() AVAudioContentSource
	SetContentSource(value AVAudioContentSource)
	Dither() bool
	SetDither(value bool)
	Downmix() bool
	SetDownmix(value bool)
	DynamicRangeControlConfiguration() AVAudioDynamicRangeControlConfiguration
	SetDynamicRangeControlConfiguration(value AVAudioDynamicRangeControlConfiguration)
	InputFormat() IAVAudioFormat
	SetInputFormat(value IAVAudioFormat)
	MagicCookie() foundation.Data
	SetMagicCookie(value foundation.Data)
	MaximumOutputPacketSize() int
	SetMaximumOutputPacketSize(value int)
	OutputFormat() IAVAudioFormat
	SetOutputFormat(value IAVAudioFormat)
	PrimeInfo() unsafe.Pointer
	SetPrimeInfo(value unsafe.Pointer)
	PrimeMethod() unsafe.Pointer
	SetPrimeMethod(value unsafe.Pointer)
	SampleRateConverterAlgorithm() string
	SetSampleRateConverterAlgorithm(value string)
	SampleRateConverterQuality() int
	SetSampleRateConverterQuality(value int)
}

// An object that converts streams of audio between formats.
//
// The audio converter class transforms audio between file formats and audio encodings. Supported transformations include: PCM float, integer, or bit depth conversions PCM sample rate conversion PCM interleaving and deinterleaving Encoding PCM to compressed formats Decoding compressed formats to PCM A single audio converter instance may perform more than one of the above transformations.


// An object that converts streams of audio between formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter
type AudioConverter struct {
	objectivec.Object
}

// AudioConverterFrom constructs a [AudioConverter] from an unsafe.Pointer.
//
// An object that converts streams of audio between formats.
func AudioConverterFrom(ptr unsafe.Pointer) AudioConverter {
	return AudioConverter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioConverterClass) Alloc() AudioConverter {
	rv := objc.Send[AudioConverter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioConverterClass) New() AudioConverter {
	rv := objc.Send[AudioConverter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioConverter) Init() AudioConverter {
	rv := objc.Send[AudioConverter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioConverter) Autorelease() AudioConverter {
	rv := objc.Send[AudioConverter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioConverter creates a new AudioConverter instance.
func NewAudioConverter() AudioConverter {
	return getAudioConverterClass().New()
}



// An array of bit rates the framework applies during encoding according to the current formats and settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodebitrates
func (a_ AudioConverter) ApplicableEncodeBitRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("applicableEncodeBitRates"))
	return rv
}


// An array of bit rates the framework applies during encoding according to the current formats and settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodebitrates
func (a_ AudioConverter) SetApplicableEncodeBitRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicableEncodeBitRates:"), value)
}


// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodesamplerates
func (a_ AudioConverter) ApplicableEncodeSampleRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("applicableEncodeSampleRates"))
	return rv
}


// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodesamplerates
func (a_ AudioConverter) SetApplicableEncodeSampleRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicableEncodeSampleRates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/audiosyncpacketfrequency
func (a_ AudioConverter) AudioSyncPacketFrequency() int {
	rv := objc.Send[int](a_.ID, objc.Sel("audioSyncPacketFrequency"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/audiosyncpacketfrequency
func (a_ AudioConverter) SetAudioSyncPacketFrequency(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioSyncPacketFrequency:"), value)
}


// An array of all bit rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodebitrates
func (a_ AudioConverter) AvailableEncodeBitRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeBitRates"))
	return rv
}


// An array of all bit rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodebitrates
func (a_ AudioConverter) SetAvailableEncodeBitRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeBitRates:"), value)
}


// An array of all output channel layout tags the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodechannellayouttags
func (a_ AudioConverter) AvailableEncodeChannelLayoutTags() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeChannelLayoutTags"))
	return rv
}


// An array of all output channel layout tags the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodechannellayouttags
func (a_ AudioConverter) SetAvailableEncodeChannelLayoutTags(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeChannelLayoutTags:"), value)
}


// An array of all output sample rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodesamplerates
func (a_ AudioConverter) AvailableEncodeSampleRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeSampleRates"))
	return rv
}


// An array of all output sample rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodesamplerates
func (a_ AudioConverter) SetAvailableEncodeSampleRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeSampleRates:"), value)
}


// The bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitrate
func (a_ AudioConverter) BitRate() int {
	rv := objc.Send[int](a_.ID, objc.Sel("bitRate"))
	return rv
}


// The bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitrate
func (a_ AudioConverter) SetBitRate(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRate:"), value)
}


// A key value constant the framework uses during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitratestrategy
func (a_ AudioConverter) BitRateStrategy() string {
	rv := objc.Send[string](a_.ID, objc.Sel("bitRateStrategy"))
	return rv
}


// A key value constant the framework uses during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitratestrategy
func (a_ AudioConverter) SetBitRateStrategy(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRateStrategy:"), objc.String(value))
}


// An array of integers that indicates which input to derive each output from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/channelmap
func (a_ AudioConverter) ChannelMap() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("channelMap"))
	return rv
}


// An array of integers that indicates which input to derive each output from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/channelmap
func (a_ AudioConverter) SetChannelMap(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelMap:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/contentsource
func (a_ AudioConverter) ContentSource() AVAudioContentSource {
	rv := objc.Send[AVAudioContentSource](a_.ID, objc.Sel("contentSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/contentsource
func (a_ AudioConverter) SetContentSource(value AVAudioContentSource) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentSource:"), value)
}


// A Boolean value that indicates whether dither is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dither
func (a_ AudioConverter) Dither() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("dither"))
	return rv
}


// A Boolean value that indicates whether dither is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dither
func (a_ AudioConverter) SetDither(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDither:"), value)
}


// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/downmix
func (a_ AudioConverter) Downmix() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downmix"))
	return rv
}


// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/downmix
func (a_ AudioConverter) SetDownmix(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDownmix:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dynamicrangecontrolconfiguration
func (a_ AudioConverter) DynamicRangeControlConfiguration() AVAudioDynamicRangeControlConfiguration {
	rv := objc.Send[AVAudioDynamicRangeControlConfiguration](a_.ID, objc.Sel("dynamicRangeControlConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dynamicrangecontrolconfiguration
func (a_ AudioConverter) SetDynamicRangeControlConfiguration(value AVAudioDynamicRangeControlConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDynamicRangeControlConfiguration:"), value)
}


// The format of the input audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/inputformat
func (a_ AudioConverter) InputFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("inputFormat"))
	return rv
}


// The format of the input audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/inputformat
func (a_ AudioConverter) SetInputFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputFormat:"), value)
}


// An object that contains metadata for encoders and decoders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/magiccookie
func (a_ AudioConverter) MagicCookie() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("magicCookie"))
	return rv
}


// An object that contains metadata for encoders and decoders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/magiccookie
func (a_ AudioConverter) SetMagicCookie(value foundation.Data) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}


// The maximum size of an output packet, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/maximumoutputpacketsize
func (a_ AudioConverter) MaximumOutputPacketSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumOutputPacketSize"))
	return rv
}


// The maximum size of an output packet, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/maximumoutputpacketsize
func (a_ AudioConverter) SetMaximumOutputPacketSize(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumOutputPacketSize:"), value)
}


// The format of the output audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/outputformat
func (a_ AudioConverter) OutputFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("outputFormat"))
	return rv
}


// The format of the output audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/outputformat
func (a_ AudioConverter) SetOutputFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFormat:"), value)
}


// The number of priming frames the converter uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primeinfo
func (a_ AudioConverter) PrimeInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primeInfo"))
	return rv
}


// The number of priming frames the converter uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primeinfo
func (a_ AudioConverter) SetPrimeInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeInfo:"), value)
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primemethod
func (a_ AudioConverter) PrimeMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primeMethod"))
	return rv
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primemethod
func (a_ AudioConverter) SetPrimeMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeMethod:"), value)
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverteralgorithm
func (a_ AudioConverter) SampleRateConverterAlgorithm() string {
	rv := objc.Send[string](a_.ID, objc.Sel("sampleRateConverterAlgorithm"))
	return rv
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverteralgorithm
func (a_ AudioConverter) SetSampleRateConverterAlgorithm(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterAlgorithm:"), objc.String(value))
}


// A sample rate converter algorithm key value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverterquality
func (a_ AudioConverter) SampleRateConverterQuality() int {
	rv := objc.Send[int](a_.ID, objc.Sel("sampleRateConverterQuality"))
	return rv
}


// A sample rate converter algorithm key value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverterquality
func (a_ AudioConverter) SetSampleRateConverterQuality(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterQuality:"), value)
}



