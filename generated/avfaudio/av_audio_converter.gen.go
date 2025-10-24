// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	ApplicableEncodeBitRates() []foundation.Number
	ApplicableEncodeSampleRates() []foundation.Number
	AudioSyncPacketFrequency() int
	SetAudioSyncPacketFrequency(value int)
	AvailableEncodeBitRates() []foundation.Number
	AvailableEncodeChannelLayoutTags() []foundation.Number
	AvailableEncodeSampleRates() []foundation.Number
	BitRate() int
	SetBitRate(value int)
	BitRateStrategy() objc.IObject /* cross-framework: NSString */
	SetBitRateStrategy(value objc.IObject /* cross-framework: NSString */)
	ChannelMap() []foundation.Number
	SetChannelMap(value []foundation.Number)
	ContentSource() AudioContentSource
	SetContentSource(value AudioContentSource)
	Dither() bool
	SetDither(value bool)
	Downmix() bool
	SetDownmix(value bool)
	DynamicRangeControlConfiguration() AudioDynamicRangeControlConfiguration
	SetDynamicRangeControlConfiguration(value AudioDynamicRangeControlConfiguration)
	InputFormat() IAVAudioFormat
	MagicCookie() objc.IObject /* cross-framework: NSData */
	SetMagicCookie(value objc.IObject /* cross-framework: NSData */)
	MaximumOutputPacketSize() int
	OutputFormat() IAVAudioFormat
	PrimeInfo() objc.IObject /* cross-framework: AVAudioConverterPrimeInfo */
	SetPrimeInfo(value objc.IObject /* cross-framework: AVAudioConverterPrimeInfo */)
	PrimeMethod() AudioConverterPrimeMethod
	SetPrimeMethod(value AudioConverterPrimeMethod)
	SampleRateConverterAlgorithm() objc.IObject /* cross-framework: NSString */
	SetSampleRateConverterAlgorithm(value objc.IObject /* cross-framework: NSString */)
	SampleRateConverterQuality() int
	SetSampleRateConverterQuality(value int)


	

	// methods:
	ConvertToBufferErrorWithInputFromBlock(outputBuffer IAVAudioBuffer, outError objectivec.IObject, inputBlock AudioConverterInputBlock /* not a class type */) AudioConverterOutputStatus
	ConvertToBufferFromBufferError(outputBuffer IAVAudioPCMBuffer, inputBuffer IAVAudioPCMBuffer, outError objectivec.IObject) bool
	Reset()


}





// Alloc allocates a new instance without initialization.
func (ac _AudioConverterClass) Alloc() AudioConverter {
	rv := objc.Send[AudioConverter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an audio converter object from the specified input and output formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/init(from:to:)
func NewAudioConverterFromFormatToFormat(fromFormat IAVAudioFormat, toFormat IAVAudioFormat) AudioConverter {
	instance := getAudioConverterClass().Alloc()
	rv := objc.Send[AudioConverter](instance.ID, objc.Sel("initFromFormat:toFormat:"), fromFormat, toFormat)
	rv.Autorelease()
	return rv
}

















// Performs a conversion between audio formats, if the system supports it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:error:withInputFrom:)
func (a_ AudioConverter) ConvertToBufferErrorWithInputFromBlock(outputBuffer IAVAudioBuffer, outError objectivec.IObject, inputBlock AudioConverterInputBlock /* not a class type */) AudioConverterOutputStatus {
	rv := objc.Send[AudioConverterOutputStatus](a_.ID, objc.Sel("convertToBuffer:error:withInputFromBlock:"), outputBuffer, outError, inputBlock)
	return rv
}


// Performs a basic conversion between audio formats that doesn’t involve converting codecs or sample rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:from:)
func (a_ AudioConverter) ConvertToBufferFromBufferError(outputBuffer IAVAudioPCMBuffer, inputBuffer IAVAudioPCMBuffer, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("convertToBuffer:fromBuffer:error:"), outputBuffer, inputBuffer, outError)
	return rv
}


// Resets the converter so you can convert a new audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/reset()
func (a_ AudioConverter) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}







// An array of bit rates the framework applies during encoding according to the current formats and settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/applicableEncodeBitRates
func (a_ AudioConverter) ApplicableEncodeBitRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("applicableEncodeBitRates"))
	return rv
}


// An array of output sample rates that the converter applies according to the current formats and settings, when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/applicableEncodeSampleRates
func (a_ AudioConverter) ApplicableEncodeSampleRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("applicableEncodeSampleRates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/audioSyncPacketFrequency
func (a_ AudioConverter) AudioSyncPacketFrequency() int {
	rv := objc.Send[int](a_.ID, objc.Sel("audioSyncPacketFrequency"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/audioSyncPacketFrequency
func (a_ AudioConverter) SetAudioSyncPacketFrequency(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioSyncPacketFrequency:"), value)
}


// An array of all bit rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeBitRates
func (a_ AudioConverter) AvailableEncodeBitRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("availableEncodeBitRates"))
	return rv
}


// An array of all output channel layout tags the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeChannelLayoutTags
func (a_ AudioConverter) AvailableEncodeChannelLayoutTags() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("availableEncodeChannelLayoutTags"))
	return rv
}


// An array of all output sample rates the codec provides when encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/availableEncodeSampleRates
func (a_ AudioConverter) AvailableEncodeSampleRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("availableEncodeSampleRates"))
	return rv
}


// The bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRate
func (a_ AudioConverter) BitRate() int {
	rv := objc.Send[int](a_.ID, objc.Sel("bitRate"))
	return rv
}


// The bit rate, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRate
func (a_ AudioConverter) SetBitRate(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRate:"), value)
}


// A key value constant the framework uses during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRateStrategy
func (a_ AudioConverter) BitRateStrategy() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("bitRateStrategy"))
	return rv
}


// A key value constant the framework uses during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/bitRateStrategy
func (a_ AudioConverter) SetBitRateStrategy(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBitRateStrategy:"), value)
}


// An array of integers that indicates which input to derive each output from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/channelMap
func (a_ AudioConverter) ChannelMap() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("channelMap"))
	return rv
}


// An array of integers that indicates which input to derive each output from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/channelMap
func (a_ AudioConverter) SetChannelMap(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelMap:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/contentSource
func (a_ AudioConverter) ContentSource() AudioContentSource {
	rv := objc.Send[AudioContentSource](a_.ID, objc.Sel("contentSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/contentSource
func (a_ AudioConverter) SetContentSource(value AudioContentSource) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentSource:"), value)
}


// A Boolean value that indicates whether dither is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dither
func (a_ AudioConverter) Dither() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("dither"))
	return rv
}


// A Boolean value that indicates whether dither is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dither
func (a_ AudioConverter) SetDither(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDither:"), value)
}


// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/downmix
func (a_ AudioConverter) Downmix() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downmix"))
	return rv
}


// A Boolean value that indicates whether the framework mixes the channels instead of remapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/downmix
func (a_ AudioConverter) SetDownmix(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDownmix:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dynamicRangeControlConfiguration
func (a_ AudioConverter) DynamicRangeControlConfiguration() AudioDynamicRangeControlConfiguration {
	rv := objc.Send[AudioDynamicRangeControlConfiguration](a_.ID, objc.Sel("dynamicRangeControlConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/dynamicRangeControlConfiguration
func (a_ AudioConverter) SetDynamicRangeControlConfiguration(value AudioDynamicRangeControlConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDynamicRangeControlConfiguration:"), value)
}


// The format of the input audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/inputFormat
func (a_ AudioConverter) InputFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("inputFormat"))
	return rv
}


// An object that contains metadata for encoders and decoders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/magicCookie
func (a_ AudioConverter) MagicCookie() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("magicCookie"))
	return rv
}


// An object that contains metadata for encoders and decoders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/magicCookie
func (a_ AudioConverter) SetMagicCookie(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMagicCookie:"), value)
}


// The maximum size of an output packet, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/maximumOutputPacketSize
func (a_ AudioConverter) MaximumOutputPacketSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumOutputPacketSize"))
	return rv
}


// The format of the output audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/outputFormat
func (a_ AudioConverter) OutputFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("outputFormat"))
	return rv
}


// The number of priming frames the converter uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeInfo
func (a_ AudioConverter) PrimeInfo() objc.IObject /* cross-framework: AVAudioConverterPrimeInfo */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("primeInfo"))
	return rv
}


// The number of priming frames the converter uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeInfo
func (a_ AudioConverter) SetPrimeInfo(value objc.IObject /* cross-framework: AVAudioConverterPrimeInfo */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeInfo:"), value)
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeMethod
func (a_ AudioConverter) PrimeMethod() AudioConverterPrimeMethod {
	rv := objc.Send[AudioConverterPrimeMethod](a_.ID, objc.Sel("primeMethod"))
	return rv
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/primeMethod
func (a_ AudioConverter) SetPrimeMethod(value AudioConverterPrimeMethod) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimeMethod:"), value)
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterAlgorithm
func (a_ AudioConverter) SampleRateConverterAlgorithm() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("sampleRateConverterAlgorithm"))
	return rv
}


// The priming method the sample rate converter or decoder uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterAlgorithm
func (a_ AudioConverter) SetSampleRateConverterAlgorithm(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterAlgorithm:"), value)
}


// A sample rate converter algorithm key value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterQuality
func (a_ AudioConverter) SampleRateConverterQuality() int {
	rv := objc.Send[int](a_.ID, objc.Sel("sampleRateConverterQuality"))
	return rv
}


// A sample rate converter algorithm key value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/sampleRateConverterQuality
func (a_ AudioConverter) SetSampleRateConverterQuality(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRateConverterQuality:"), value)
}







