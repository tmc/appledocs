// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReaderTrackOutput */


/* debug [class_header]: Header for AVAssetReaderTrackOutput */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReaderTrackOutput */
// An interface definition for the [AssetReaderTrackOutput] class.
type IAssetReaderTrackOutput interface {
	IAssetReaderOutput
	
/* debug [class_interface_properties]: Properties for AssetReaderTrackOutput */
	// properties:
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */)
	OutputSettings() foundation.IDictionary
	Track() IAVAssetTrack
	AVVideoCleanApertureKey() objc.IObject /* cross-framework: NSString */
	AVVideoPixelAspectRatioKey() objc.IObject /* cross-framework: NSString */
	AVVideoScalingModeKey() objc.IObject /* cross-framework: NSString */
	AVFormatIDKey() objc.IObject /* cross-framework: NSString */
	AVSampleRateConverterAudioQualityKey() objc.IObject /* cross-framework: NSString */
	KAudioFormatLinearPCM() objectivec.IObject
	SetKAudioFormatLinearPCM(value objectivec.IObject)
	KCMFormatDescriptionExtension_Depth() foundation.String
	KCVPixelBufferHeightKey() foundation.String
	KCVPixelBufferWidthKey() foundation.String
	KCVPixelFormatType_32ARGB() uint32 /* not a class type */
	SetKCVPixelFormatType_32ARGB(value uint32 /* not a class type */)
	KCVPixelFormatType_32BGRA() uint32 /* not a class type */
	SetKCVPixelFormatType_32BGRA(value uint32 /* not a class type */)
	KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() uint32 /* not a class type */
	SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value uint32 /* not a class type */)
	KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() uint32 /* not a class type */
	SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value uint32 /* not a class type */)
	KCVPixelFormatType_422YpCbCr10() uint32 /* not a class type */
	SetKCVPixelFormatType_422YpCbCr10(value uint32 /* not a class type */)
	KCVPixelFormatType_422YpCbCr16() uint32 /* not a class type */
	SetKCVPixelFormatType_422YpCbCr16(value uint32 /* not a class type */)
	KCVPixelFormatType_422YpCbCr8() uint32 /* not a class type */
	SetKCVPixelFormatType_422YpCbCr8(value uint32 /* not a class type */)
	KCVPixelFormatType_4444AYpCbCr16() uint32 /* not a class type */
	SetKCVPixelFormatType_4444AYpCbCr16(value uint32 /* not a class type */)
	KCVPixelFormatType_64ARGB() uint32 /* not a class type */
	SetKCVPixelFormatType_64ARGB(value uint32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReaderTrackOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReaderTrackOutput */
// Alloc allocates a new instance without initialization.
func (ac _AssetReaderTrackOutputClass) Alloc() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReaderTrackOutput */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReaderTrackOutput */

// Creates an object that reads media data from an asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/init(track:outputSettings:)
func NewAssetReaderTrackOutputWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.IDictionary) AssetReaderTrackOutput {
	instance := getAssetReaderTrackOutputClass().Alloc()
	rv := objc.Send[AssetReaderTrackOutput](instance.ID, objc.Sel("initWithTrack:outputSettings:"), track, outputSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetReaderTrackOutputWithTrackOutputSettings */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReaderTrackOutput */

// Returns a new object that reads media data from an asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/assetReaderTrackOutputWithTrack:outputSettings:
func (ac _AssetReaderTrackOutputClass) AssetReaderTrackOutputWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderTrackOutputWithTrack:outputSettings:"), track, outputSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetReaderTrackOutputWithTrackOutputSettings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReaderTrackOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReaderTrackOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReaderTrackOutput */

// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/audioTimePitchAlgorithm
func (a_ AssetReaderTrackOutput) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/audioTimePitchAlgorithm
func (a_ AssetReaderTrackOutput) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}/* debug [instance_properties/setter]: audioTimePitchAlgorithm */


// The output settings for this track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/outputSettings
func (a_ AssetReaderTrackOutput) OutputSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("outputSettings"))
	return rv
}/* debug [instance_properties/getter]: outputSettings */


// The track from which the output reads sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/track
func (a_ AssetReaderTrackOutput) Track() IAVAssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("track"))
	return rv
}/* debug [instance_properties/getter]: track */


// A key that defines the region within the video dimension displayed during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocleanaperturekey
func (a_ AssetReaderTrackOutput) AVVideoCleanApertureKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVVideoCleanApertureKey"))
	return rv
}/* debug [instance_properties/getter]: AVVideoCleanApertureKey */


// A key to access the video’s pixel aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideopixelaspectratiokey
func (a_ AssetReaderTrackOutput) AVVideoPixelAspectRatioKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVVideoPixelAspectRatioKey"))
	return rv
}/* debug [instance_properties/getter]: AVVideoPixelAspectRatioKey */


// A key to retrieve the video scaling mode from a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideoscalingmodekey
func (a_ AssetReaderTrackOutput) AVVideoScalingModeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVVideoScalingModeKey"))
	return rv
}/* debug [instance_properties/getter]: AVVideoScalingModeKey */


// An integer value that represents the format of the audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVFormatIDKey
func (a_ AssetReaderTrackOutput) AVFormatIDKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVFormatIDKey"))
	return rv
}/* debug [instance_properties/getter]: AVFormatIDKey */


// An integer value that represents the audio quality for conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAudioQualityKey
func (a_ AssetReaderTrackOutput) AVSampleRateConverterAudioQualityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVSampleRateConverterAudioQualityKey"))
	return rv
}/* debug [instance_properties/getter]: AVSampleRateConverterAudioQualityKey */


// A key that specifies the linear PCM codec, and uses the standard flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
func (a_ AssetReaderTrackOutput) KAudioFormatLinearPCM() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioFormatLinearPCM"))
	return rv
}/* debug [instance_properties/getter]: kAudioFormatLinearPCM */


// A key that specifies the linear PCM codec, and uses the standard flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
func (a_ AssetReaderTrackOutput) SetKAudioFormatLinearPCM(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioFormatLinearPCM:"), value)
}/* debug [instance_properties/setter]: kAudioFormatLinearPCM */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Depth
func (a_ AssetReaderTrackOutput) KCMFormatDescriptionExtension_Depth() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCMFormatDescriptionExtension_Depth"))
	return rv
}/* debug [instance_properties/getter]: kCMFormatDescriptionExtension_Depth */


// A key to the height of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferHeightKey
func (a_ AssetReaderTrackOutput) KCVPixelBufferHeightKey() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCVPixelBufferHeightKey"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelBufferHeightKey */


// A key to the width of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferWidthKey
func (a_ AssetReaderTrackOutput) KCVPixelBufferWidthKey() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCVPixelBufferWidthKey"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelBufferWidthKey */


// 32-bit ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_32ARGB() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_32ARGB"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_32ARGB */


// 32-bit ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_32ARGB(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_32ARGB:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_32ARGB */


// 32-bit BGRA.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_32BGRA() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_32BGRA"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_32BGRA */


// 32-bit BGRA.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_32BGRA(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_32BGRA:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_32BGRA */


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarFullRange"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_420YpCbCr8BiPlanarFullRange */


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarFullRange:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_420YpCbCr8BiPlanarFullRange */


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange */


// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]). `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange */


// Component Y’CbCr 10-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr10() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr10"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_422YpCbCr10 */


// Component Y’CbCr 10-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr10(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr10:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_422YpCbCr10 */


// Component Y’CbCr 10,12,14,16-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr16() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr16"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_422YpCbCr16 */


// Component Y’CbCr 10,12,14,16-bit 4:2:2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr16(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr16:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_422YpCbCr16 */


// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr8() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_422YpCbCr8"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_422YpCbCr8 */


// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr8(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr8:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_422YpCbCr8 */


// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_4444AYpCbCr16() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_4444AYpCbCr16"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_4444AYpCbCr16 */


// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_4444AYpCbCr16(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_4444AYpCbCr16:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_4444AYpCbCr16 */


// 64-bit ARGB, 16-bit big-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
func (a_ AssetReaderTrackOutput) KCVPixelFormatType_64ARGB() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("kCVPixelFormatType_64ARGB"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_64ARGB */


// 64-bit ARGB, 16-bit big-endian samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
func (a_ AssetReaderTrackOutput) SetKCVPixelFormatType_64ARGB(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKCVPixelFormatType_64ARGB:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_64ARGB */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReaderTrackOutput */


