// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptureVideoDataOutput] class.
type ICaptureVideoDataOutput interface {
	ICaptureOutput
	

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


	

	// methods:
	AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType) []string
	RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) []MetadataItem
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) foundation.IDictionary
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType VideoCodecType, outputFileType FileType, outputFileURL foundation.foundation.INSURL) foundation.IDictionary
	RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType) foundation.IDictionary
	SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject)


}





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





















// The video codecs that the output supports for writing video to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypesForAssetWriter(writingTo:)
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType) []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableVideoCodecTypesForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}


// Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMovieMetadata(forVideoCodecType:assetWriterOutputFileType:)
func (c_ CaptureVideoDataOutput) RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("recommendedMovieMetadataForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}


// Returns a video settings dictionary appropriate for capturing video to a file with the specified codec and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}


// Returns a dictionary of recommended output settings for writing the specified code, file type, and output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:outputFileURL:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType VideoCodecType, outputFileType FileType, outputFileURL foundation.foundation.INSURL) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:outputFileURL:"), videoCodecType, outputFileType, outputFileURL)
	return rv
}


// Specifies the recommended settings for use with an AVAssetWriterInput.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettingsForAssetWriter(writingTo:)
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recommendedVideoSettingsForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}


// Sets the sample buffer delegate and the queue for invoking callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}







// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) AlwaysDiscardsLateVideoFrames() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alwaysDiscardsLateVideoFrames"))
	return rv
}


// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) SetAlwaysDiscardsLateVideoFrames(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlwaysDiscardsLateVideoFrames:"), value)
}


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCVPixelFormatTypes
func (c_ CaptureVideoDataOutput) AvailableVideoCVPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("availableVideoCVPixelFormatTypes"))
	return rv
}


// The video codecs that the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypes
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preservesDynamicHDRMetadata
func (c_ CaptureVideoDataOutput) PreservesDynamicHDRMetadata() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesDynamicHDRMetadata"))
	return rv
}


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preservesDynamicHDRMetadata
func (c_ CaptureVideoDataOutput) SetPreservesDynamicHDRMetadata(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesDynamicHDRMetadata:"), value)
}


// Indicates the recommended media timescale for the video track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMediaTimeScaleForAssetWriter
func (c_ CaptureVideoDataOutput) RecommendedMediaTimeScaleForAssetWriter() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](c_.ID, objc.Sel("recommendedMediaTimeScaleForAssetWriter"))
	return rv
}


// The queue on which the system invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferCallbackQueue
func (c_ CaptureVideoDataOutput) SampleBufferCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferDelegate
func (c_ CaptureVideoDataOutput) SampleBufferDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBufferDelegate"))
	return rv
}


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) VideoSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("videoSettings"))
	return rv
}


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) SetVideoSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoSettings:"), value)
}


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) AvailableVideoPixelFormatTypes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](c_.ID, objc.Sel("availableVideoPixelFormatTypes"))
	return rv
}


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) SetAvailableVideoPixelFormatTypes(value uint32 /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableVideoPixelFormatTypes:"), value)
}







