// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AssetReaderTrackOutput] class.
var (
	AssetReaderTrackOutputClass     _AssetReaderTrackOutputClass
	AssetReaderTrackOutputClassOnce sync.Once
)

func getAssetReaderTrackOutputClass() _AssetReaderTrackOutputClass {
	AssetReaderTrackOutputClassOnce.Do(func() {
		AssetReaderTrackOutputClass = _AssetReaderTrackOutputClass{objc.GetClass("AVAssetReaderTrackOutput")}
	})
	return AssetReaderTrackOutputClass
}

type _AssetReaderTrackOutputClass struct {
	class objc.Class
}

// An interface definition for the [AssetReaderTrackOutput] class.
type IAssetReaderTrackOutput interface {
	IAssetReaderOutput
	// properties:
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */)
	OutputSettings() string /* primitive/slice/pointer */
	SetOutputSettings(value string /* primitive/slice/pointer */)
	Track() IAVAssetTrack
	SetTrack(value IAVAssetTrack)
	AVVideoCleanApertureKey() string /* primitive/slice/pointer */
	AVVideoPixelAspectRatioKey() string /* primitive/slice/pointer */
	AVVideoScalingModeKey() string /* primitive/slice/pointer */
	AVFormatIDKey() string /* primitive/slice/pointer */
	AVSampleRateConverterAudioQualityKey() string /* primitive/slice/pointer */
	KAudioFormatLinearPCM() unsafe.Pointer
	SetKAudioFormatLinearPCM(value unsafe.Pointer)
	KCMFormatDescriptionExtension_Depth() String /* not a class type */
	KCVPixelBufferHeightKey() String /* not a class type */
	KCVPixelBufferWidthKey() String /* not a class type */
	KCVPixelFormatType_32ARGB() unsafe.Pointer
	SetKCVPixelFormatType_32ARGB(value unsafe.Pointer)
	KCVPixelFormatType_32BGRA() unsafe.Pointer
	SetKCVPixelFormatType_32BGRA(value unsafe.Pointer)
	KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() unsafe.Pointer
	SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value unsafe.Pointer)
	KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() unsafe.Pointer
	SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value unsafe.Pointer)
	KCVPixelFormatType_422YpCbCr10() unsafe.Pointer
	SetKCVPixelFormatType_422YpCbCr10(value unsafe.Pointer)
	KCVPixelFormatType_422YpCbCr16() unsafe.Pointer
	SetKCVPixelFormatType_422YpCbCr16(value unsafe.Pointer)
	KCVPixelFormatType_422YpCbCr8() unsafe.Pointer
	SetKCVPixelFormatType_422YpCbCr8(value unsafe.Pointer)
	KCVPixelFormatType_4444AYpCbCr16() unsafe.Pointer
	SetKCVPixelFormatType_4444AYpCbCr16(value unsafe.Pointer)
	KCVPixelFormatType_64ARGB() unsafe.Pointer
	SetKCVPixelFormatType_64ARGB(value unsafe.Pointer)
	// methods:
}

// An object that reads media data from a single track of an asset.
//
// Read the media data of an asset track by adding a track output to an asset reader. You can read the media samples in their stored format, or you can convert them to an alternative format. A track output produces uncompressed output. For audio output settings, this means that must be . For video output settings, this means that the dictionary must contain values for uncompressed video output, as defined in . A track output doesn’t support the audio setting key or the following video settings keys: , , and . When constructing video output settings, the choice of pixel format affects the performance and quality of the decompression. For optimal performance when decompressing video, the requested pixel format should be one that the decoder supports natively to avoid unnecessary conversions. Below are some recommendations: For H.264, use or when you know the video is full range. In iOS, use for JPEG output. In macOS, is the preferred pixel format for video and generally provides the best performance when decoding. If you need to work in the RGB domain, use in iOS, and in macOS. ProRes-encoded media can contain up to 12 bits per channel. For ProRes-encoded sources that you wish to preserve more than 8 bits per channel during decompression, use one of the following pixel formats: , , , or . doesn’t support scaling with any of these high-bit-depth pixel formats. If you use the above pixel formats, don’t specify or in the dictionary. Only ProRes encoders support these pixel formats. ProRes 4444-encoded media can contain a mathematically lossless alpha channel. To preserve the alpha channel during decompression, use a pixel format with an alpha component such as or . To test whether your source contains an alpha channel, check that the track’s format description has a key with a value of .


// An object that reads media data from a single track of an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput
type AssetReaderTrackOutput struct {
	AssetReaderOutput
}

// AssetReaderTrackOutputFrom constructs a [AssetReaderTrackOutput] from an unsafe.Pointer.
//
// An object that reads media data from a single track of an asset.
func AssetReaderTrackOutputFrom(ptr unsafe.Pointer) AssetReaderTrackOutput {
	return AssetReaderTrackOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetReaderTrackOutputClass) Alloc() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetReaderTrackOutputClass) New() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderTrackOutput) Init() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderTrackOutput) Autorelease() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderTrackOutput creates a new AssetReaderTrackOutput instance.
func NewAssetReaderTrackOutput() AssetReaderTrackOutput {
	return getAssetReaderTrackOutputClass().New()
}



// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/audiotimepitchalgorithm
func (a_ AssetReaderTrackOutput) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */ {
	rv := objc.Send[AudioTimePitchAlgorithm](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/audiotimepitchalgorithm
func (a_ AssetReaderTrackOutput) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// The output settings for this track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/outputsettings
func (a_ AssetReaderTrackOutput) OutputSettings() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("outputSettings"))
	return rv
}


// The output settings for this track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/outputsettings
func (a_ AssetReaderTrackOutput) SetOutputSettings(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputSettings:"), objc.String(value))
}


// The track from which the output reads sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/track
func (a_ AssetReaderTrackOutput) Track() IAVAssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("track"))
	return rv
}


// The track from which the output reads sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/track
func (a_ AssetReaderTrackOutput) SetTrack(value IAVAssetTrack) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrack:"), value)
}


// A key that defines the region within the video dimension displayed during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocleanaperturekey
func (a_ AssetReaderTrackOutput) AVVideoCleanApertureKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("AVVideoCleanApertureKey"))
	return rv
}


// A key to access the video’s pixel aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideopixelaspectratiokey
func (a_ AssetReaderTrackOutput) AVVideoPixelAspectRatioKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("AVVideoPixelAspectRatioKey"))
	return rv
}


// A key to retrieve the video scaling mode from a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideoscalingmodekey
func (a_ AssetReaderTrackOutput) AVVideoScalingModeKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("AVVideoScalingModeKey"))
	return rv
}


// An integer value that represents the format of the audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVFormatIDKey
func (a_ AssetReaderTrackOutput) AVFormatIDKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("AVFormatIDKey"))
	return rv
}


// An integer value that represents the audio quality for conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAudioQualityKey
func (a_ AssetReaderTrackOutput) AVSampleRateConverterAudioQualityKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](a_.ID, objc.Sel("AVSampleRateConverterAudioQualityKey"))
	return rv
}


// A key that specifies the linear PCM codec, and uses the standard flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
func (a_ AssetReaderTrackOutput) KAudioFormatLinearPCM() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioFormatLinearPCM"))
	return rv
}


// A key that specifies the linear PCM codec, and uses the standard flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
func (a_ AssetReaderTrackOutput) SetKAudioFormatLinearPCM(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioFormatLinearPCM:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Depth
func (a_ AssetReaderTrackOutput) KCMFormatDescriptionExtension_Depth() String /* not a class type */ {
	rv := objc.Send[String](a_.ID, objc.Sel("kCMFormatDescriptionExtension_Depth"))
	return rv
}


// A key to the height of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferHeightKey
func (a_ AssetReaderTrackOutput) KCVPixelBufferHeightKey() String /* not a class type */ {
	rv := objc.Send[String](a_.ID, objc.Sel("kCVPixelBufferHeightKey"))
	return rv
}


// A key to the width of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferWidthKey
func (a_ AssetReaderTrackOutput) KCVPixelBufferWidthKey() String /* not a class type */ {
	rv := objc.Send[String](a_.ID, objc.Sel("kCVPixelBufferWidthKey"))
	return rv
}


// 32-bit ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_32ARGB() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_32ARGB"))
	return rv
}


// 32-bit ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_32ARGB(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_32ARGB:"), value)
}


// 32-bit BGRA.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_32BGRA() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_32BGRA"))
	return rv
}


// 32-bit BGRA.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_32BGRA(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_32BGRA:"), value)
}


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarFullRange"))
	return rv
}


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarFullRange:"), value)
}


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange"))
	return rv
}


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange:"), value)
}


// Component Y’CbCr 10-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr10() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr10"))
	return rv
}


// Component Y’CbCr 10-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr10(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr10:"), value)
}


// Component Y’CbCr 10,12,14,16-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr16() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr16"))
	return rv
}


// Component Y’CbCr 10,12,14,16-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr16(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr16:"), value)
}


// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr8() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr8"))
	return rv
}


// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr8(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr8:"), value)
}


// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_4444AYpCbCr16() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_4444AYpCbCr16"))
	return rv
}


// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_4444AYpCbCr16(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_4444AYpCbCr16:"), value)
}


// 64-bit ARGB, 16-bit big-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_64ARGB() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kCVPixelFormatType_64ARGB"))
	return rv
}


// 64-bit ARGB, 16-bit big-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_64ARGB(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_64ARGB:"), value)
}



