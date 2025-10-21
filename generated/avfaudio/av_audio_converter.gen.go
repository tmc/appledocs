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
	ConvertToBufferErrorWithInputFromBlock(outputBuffer unsafe.Pointer, outError unsafe.Pointer, inputBlock unsafe.Pointer) unsafe.Pointer
}

// An object that converts streams of audio between formats.
//
// The audio converter class transforms audio between file formats and audio encodings. Supported transformations include: PCM float, integer, or bit depth conversions PCM sample rate conversion PCM interleaving and deinterleaving Encoding PCM to compressed formats Decoding compressed formats to PCM A single audio converter instance may perform more than one of the above transformations.
//
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


// Performs a conversion between audio formats, if the system supports it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:error:withInputFrom:)
func (a_ AudioConverter) ConvertToBufferErrorWithInputFromBlock(outputBuffer unsafe.Pointer, outError unsafe.Pointer, inputBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("convertToBuffer:error:withInputFromBlock:"), outputBuffer, outError, inputBlock)
	return rv
}

// An array of bit rates the framework applies during encoding according to the current formats and settings.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodebitrates
func (a_ AudioConverter) ApplicableEncodeBitRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("applicableEncodeBitRates"))
	return rv
}


// SetApplicableEncodeBitRates sets the value of the applicableEncodeBitRates property.
// An array of bit rates the framework applies during encoding according to the current formats and settings.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodebitrates
func (a_ AudioConverter) SetApplicableEncodeBitRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicableEncodeBitRates:"), value)
}

// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodesamplerates
func (a_ AudioConverter) ApplicableEncodeSampleRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("applicableEncodeSampleRates"))
	return rv
}


// SetApplicableEncodeSampleRates sets the value of the applicableEncodeSampleRates property.
// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/applicableencodesamplerates
func (a_ AudioConverter) SetApplicableEncodeSampleRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicableEncodeSampleRates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/audiosyncpacketfrequency
func (a_ AudioConverter) AudioSyncPacketFrequency() int {
	rv := objc.Send[int](a_.ID, objc.Sel("audioSyncPacketFrequency"))
	return rv
}


// SetAudioSyncPacketFrequency sets the value of the audioSyncPacketFrequency property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/audiosyncpacketfrequency
func (a_ AudioConverter) SetAudioSyncPacketFrequency(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioSyncPacketFrequency:"), value)
}

// An array of all bit rates the codec provides when encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodebitrates
func (a_ AudioConverter) AvailableEncodeBitRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeBitRates"))
	return rv
}


// SetAvailableEncodeBitRates sets the value of the availableEncodeBitRates property.
// An array of all bit rates the codec provides when encoding.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodebitrates
func (a_ AudioConverter) SetAvailableEncodeBitRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeBitRates:"), value)
}

// An array of all output channel layout tags the codec provides when encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodechannellayouttags
func (a_ AudioConverter) AvailableEncodeChannelLayoutTags() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeChannelLayoutTags"))
	return rv
}


// SetAvailableEncodeChannelLayoutTags sets the value of the availableEncodeChannelLayoutTags property.
// An array of all output channel layout tags the codec provides when encoding.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodechannellayouttags
func (a_ AudioConverter) SetAvailableEncodeChannelLayoutTags(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeChannelLayoutTags:"), value)
}

// An array of all output sample rates the codec provides when encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodesamplerates
func (a_ AudioConverter) AvailableEncodeSampleRates() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("availableEncodeSampleRates"))
	return rv
}


// SetAvailableEncodeSampleRates sets the value of the availableEncodeSampleRates property.
// An array of all output sample rates the codec provides when encoding.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/availableencodesamplerates
func (a_ AudioConverter) SetAvailableEncodeSampleRates(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableEncodeSampleRates:"), value)
}

// The bit rate, in bits per second.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitrate
func (a_ AudioConverter) BitRate() int {
	rv := objc.Send[int](a_.ID, objc.Sel("bitRate"))
	return rv
}


// SetBitRate sets the value of the bitRate property.
// The bit rate, in bits per second.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitrate
func (a_ AudioConverter) SetBitRate(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRate:"), value)
}

// A key value constant the framework uses during encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitratestrategy
func (a_ AudioConverter) BitRateStrategy() string {
	rv := objc.Send[string](a_.ID, objc.Sel("bitRateStrategy"))
	return rv
}


// SetBitRateStrategy sets the value of the bitRateStrategy property.
// A key value constant the framework uses during encoding.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/bitratestrategy
func (a_ AudioConverter) SetBitRateStrategy(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRateStrategy:"), objc.String(value))
}

// An array of integers that indicates which input to derive each output from.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/channelmap
func (a_ AudioConverter) ChannelMap() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("channelMap"))
	return rv
}


// SetChannelMap sets the value of the channelMap property.
// An array of integers that indicates which input to derive each output from.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/channelmap
func (a_ AudioConverter) SetChannelMap(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelMap:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/contentsource
func (a_ AudioConverter) ContentSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("contentSource"))
	return rv
}


// SetContentSource sets the value of the contentSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/contentsource
func (a_ AudioConverter) SetContentSource(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentSource:"), value)
}

// A Boolean value that indicates whether dither is on.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dither
func (a_ AudioConverter) Dither() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("dither"))
	return rv
}


// SetDither sets the value of the dither property.
// A Boolean value that indicates whether dither is on.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dither
func (a_ AudioConverter) SetDither(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDither:"), value)
}

// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/downmix
func (a_ AudioConverter) Downmix() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downmix"))
	return rv
}


// SetDownmix sets the value of the downmix property.
// A Boolean value that indicates whether the framework mixes the channels instead of remapping.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/downmix
func (a_ AudioConverter) SetDownmix(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDownmix:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dynamicrangecontrolconfiguration
func (a_ AudioConverter) DynamicRangeControlConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("dynamicRangeControlConfiguration"))
	return rv
}


// SetDynamicRangeControlConfiguration sets the value of the dynamicRangeControlConfiguration property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/dynamicrangecontrolconfiguration
func (a_ AudioConverter) SetDynamicRangeControlConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDynamicRangeControlConfiguration:"), value)
}

// The format of the input audio stream.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/inputformat
func (a_ AudioConverter) InputFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputFormat"))
	return rv
}


// SetInputFormat sets the value of the inputFormat property.
// The format of the input audio stream.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/inputformat
func (a_ AudioConverter) SetInputFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputFormat:"), value)
}

// An object that contains metadata for encoders and decoders.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/magiccookie
func (a_ AudioConverter) MagicCookie() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("magicCookie"))
	return rv
}


// SetMagicCookie sets the value of the magicCookie property.
// An object that contains metadata for encoders and decoders.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/magiccookie
func (a_ AudioConverter) SetMagicCookie(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}

// The maximum size of an output packet, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/maximumoutputpacketsize
func (a_ AudioConverter) MaximumOutputPacketSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumOutputPacketSize"))
	return rv
}


// SetMaximumOutputPacketSize sets the value of the maximumOutputPacketSize property.
// The maximum size of an output packet, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/maximumoutputpacketsize
func (a_ AudioConverter) SetMaximumOutputPacketSize(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumOutputPacketSize:"), value)
}

// The format of the output audio stream.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/outputformat
func (a_ AudioConverter) OutputFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputFormat"))
	return rv
}


// SetOutputFormat sets the value of the outputFormat property.
// The format of the output audio stream.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/outputformat
func (a_ AudioConverter) SetOutputFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFormat:"), value)
}

// The number of priming frames the converter uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primeinfo
func (a_ AudioConverter) PrimeInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primeInfo"))
	return rv
}


// SetPrimeInfo sets the value of the primeInfo property.
// The number of priming frames the converter uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primeinfo
func (a_ AudioConverter) SetPrimeInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeInfo:"), value)
}

// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primemethod
func (a_ AudioConverter) PrimeMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primeMethod"))
	return rv
}


// SetPrimeMethod sets the value of the primeMethod property.
// The priming method the sample rate converter or decoder uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/primemethod
func (a_ AudioConverter) SetPrimeMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeMethod:"), value)
}

// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverteralgorithm
func (a_ AudioConverter) SampleRateConverterAlgorithm() string {
	rv := objc.Send[string](a_.ID, objc.Sel("sampleRateConverterAlgorithm"))
	return rv
}


// SetSampleRateConverterAlgorithm sets the value of the sampleRateConverterAlgorithm property.
// The priming method the sample rate converter or decoder uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverteralgorithm
func (a_ AudioConverter) SetSampleRateConverterAlgorithm(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterAlgorithm:"), objc.String(value))
}

// A sample rate converter algorithm key value.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverterquality
func (a_ AudioConverter) SampleRateConverterQuality() int {
	rv := objc.Send[int](a_.ID, objc.Sel("sampleRateConverterQuality"))
	return rv
}


// SetSampleRateConverterQuality sets the value of the sampleRateConverterQuality property.
// A sample rate converter algorithm key value.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioconverter/samplerateconverterquality
func (a_ AudioConverter) SetSampleRateConverterQuality(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterQuality:"), value)
}



