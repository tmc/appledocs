// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureVideoDataOutput */


/* debug [class_header]: Header for AVCaptureVideoDataOutput */
// The class instance for the [CaptureVideoDataOutput] class.
var (
	CaptureVideoDataOutputClass     _CaptureVideoDataOutputClass
	CaptureVideoDataOutputClassOnce sync.Once
)

func getCaptureVideoDataOutputClass() _CaptureVideoDataOutputClass {
	CaptureVideoDataOutputClassOnce.Do(func() {
		CaptureVideoDataOutputClass = _CaptureVideoDataOutputClass{objc.GetClass("AVCaptureVideoDataOutput")}
	})
	return CaptureVideoDataOutputClass
}

type _CaptureVideoDataOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureVideoDataOutput */
// An interface definition for the [CaptureVideoDataOutput] class.
type ICaptureVideoDataOutput interface {
	ICaptureOutput
	
/* debug [class_interface_properties]: Properties for CaptureVideoDataOutput */
	// properties:
	AlwaysDiscardsLateVideoFrames() bool
	SetAlwaysDiscardsLateVideoFrames(value bool)
	AvailableVideoCVPixelFormatTypes() []foundation.Number
	AvailableVideoCodecTypes() []string
	PreservesDynamicHDRMetadata() bool
	SetPreservesDynamicHDRMetadata(value bool)
	RecommendedMediaTimeScaleForAssetWriter() TimeScale /* not a class type */
	SampleBufferCallbackQueue() objectivec.IObject
	SampleBufferDelegate() unsafe.Pointer
	VideoSettings() foundation.IDictionary
	SetVideoSettings(value foundation.IDictionary)
	AvailableVideoPixelFormatTypes() uint32 /* not a class type */
	SetAvailableVideoPixelFormatTypes(value uint32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureVideoDataOutput */
	// methods:
	AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType /* typedef */) []string
	RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */) []MetadataItem
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */) foundation.IDictionary
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */, outputFileURL objc.IObject /* cross-framework: NSURL */) foundation.IDictionary
	RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType /* typedef */) foundation.IDictionary
	SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureVideoDataOutput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureVideoDataOutputClass) Alloc() CaptureVideoDataOutput {
	rv := objc.Send[CaptureVideoDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureVideoDataOutputClass) New() CaptureVideoDataOutput {
	rv := objc.Send[CaptureVideoDataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureVideoDataOutput) Init() CaptureVideoDataOutput {
	rv := objc.Send[CaptureVideoDataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureVideoDataOutput) Autorelease() CaptureVideoDataOutput {
	rv := objc.Send[CaptureVideoDataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureVideoDataOutput creates a new CaptureVideoDataOutput instance.
func NewCaptureVideoDataOutput() CaptureVideoDataOutput {
	return getCaptureVideoDataOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureVideoDataOutput */
// A capture output that records video and provides access to video frames for processing.
//
// Use this output to process compressed or uncompressed frames from the captured video. You can access the frames with the delegate method. This object supports compressed video data output for macOS only. It can output pixel buffers in several pixel formats. Consider the usability and performance characteristics of these formats and choose the best format for your app.


// A capture output that records video and provides access to video frames for processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput
type CaptureVideoDataOutput struct {
	CaptureOutput
}

// CaptureVideoDataOutputFrom constructs a [CaptureVideoDataOutput] from an unsafe.Pointer.
//
// A capture output that records video and provides access to video frames for processing.
func CaptureVideoDataOutputFrom(ptr unsafe.Pointer) CaptureVideoDataOutput {
	return CaptureVideoDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureVideoDataOutput */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureVideoDataOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureVideoDataOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureVideoDataOutput */

// The video codecs that the output supports for writing video to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypesForAssetWriter(writingTo:)
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType /* typedef */) []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableVideoCodecTypesForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}/* debug [instance_methods/method]: AvailableVideoCodecTypesForAssetWriterWithOutputFileType */


// Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMovieMetadata(forVideoCodecType:assetWriterOutputFileType:)
func (c_ CaptureVideoDataOutput) RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("recommendedMovieMetadataForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}/* debug [instance_methods/method]: RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType */


// Returns a video settings dictionary appropriate for capturing video to a file with the specified codec and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}/* debug [instance_methods/method]: RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType */


// Returns a dictionary of recommended output settings for writing the specified code, file type, and output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:outputFileURL:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType VideoCodecType /* typedef */, outputFileType FileType /* typedef */, outputFileURL objc.IObject /* cross-framework: NSURL */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:outputFileURL:"), videoCodecType, outputFileType, outputFileURL)
	return rv
}/* debug [instance_methods/method]: RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL */


// Specifies the recommended settings for use with an AVAssetWriterInput.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettingsForAssetWriter(writingTo:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType /* typedef */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}/* debug [instance_methods/method]: RecommendedVideoSettingsForAssetWriterWithOutputFileType */


// Sets the sample buffer delegate and the queue for invoking callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}/* debug [instance_methods/method]: SetSampleBufferDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureVideoDataOutput */

// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) AlwaysDiscardsLateVideoFrames() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alwaysDiscardsLateVideoFrames"))
	return rv
}/* debug [instance_properties/getter]: alwaysDiscardsLateVideoFrames */


// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) SetAlwaysDiscardsLateVideoFrames(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlwaysDiscardsLateVideoFrames:"), value)
}/* debug [instance_properties/setter]: alwaysDiscardsLateVideoFrames */


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCVPixelFormatTypes
func (c_ CaptureVideoDataOutput) AvailableVideoCVPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("availableVideoCVPixelFormatTypes"))
	return rv
}/* debug [instance_properties/getter]: availableVideoCVPixelFormatTypes */


// The video codecs that the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypes
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}/* debug [instance_properties/getter]: availableVideoCodecTypes */


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preservesDynamicHDRMetadata
func (c_ CaptureVideoDataOutput) PreservesDynamicHDRMetadata() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesDynamicHDRMetadata"))
	return rv
}/* debug [instance_properties/getter]: preservesDynamicHDRMetadata */


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preservesDynamicHDRMetadata
func (c_ CaptureVideoDataOutput) SetPreservesDynamicHDRMetadata(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesDynamicHDRMetadata:"), value)
}/* debug [instance_properties/setter]: preservesDynamicHDRMetadata */


// Indicates the recommended media timescale for the video track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMediaTimeScaleForAssetWriter
func (c_ CaptureVideoDataOutput) RecommendedMediaTimeScaleForAssetWriter() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](c_.ID, objc.Sel("recommendedMediaTimeScaleForAssetWriter"))
	return rv
}/* debug [instance_properties/getter]: recommendedMediaTimeScaleForAssetWriter */


// The queue on which the system invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferCallbackQueue
func (c_ CaptureVideoDataOutput) SampleBufferCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferCallbackQueue */


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferDelegate
func (c_ CaptureVideoDataOutput) SampleBufferDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBufferDelegate"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferDelegate */


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) VideoSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("videoSettings"))
	return rv
}/* debug [instance_properties/getter]: videoSettings */


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) SetVideoSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoSettings:"), value)
}/* debug [instance_properties/setter]: videoSettings */


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) AvailableVideoPixelFormatTypes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](c_.ID, objc.Sel("availableVideoPixelFormatTypes"))
	return rv
}/* debug [instance_properties/getter]: availableVideoPixelFormatTypes */


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) SetAvailableVideoPixelFormatTypes(value uint32 /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableVideoPixelFormatTypes:"), value)
}/* debug [instance_properties/setter]: availableVideoPixelFormatTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureVideoDataOutput */


