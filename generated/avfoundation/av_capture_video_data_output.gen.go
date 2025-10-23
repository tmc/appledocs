// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AlwaysDiscardsLateVideoFrames() bool /* primitive/slice/pointer */
	SetAlwaysDiscardsLateVideoFrames(value bool /* primitive/slice/pointer */)
	MinFrameDuration() Time /* not a class type */
	SetMinFrameDuration(value Time /* not a class type */)
	SampleBufferCallbackQueue() unsafe.Pointer
	VideoSettings() foundation.IDictionary /* already interface */
	SetVideoSettings(value foundation.IDictionary /* already interface */)
	AutomaticallyConfiguresOutputBufferDimensions() bool /* primitive/slice/pointer */
	SetAutomaticallyConfiguresOutputBufferDimensions(value bool /* primitive/slice/pointer */)
	AvailableVideoCodecTypes() VideoCodecType /* not a class type */
	SetAvailableVideoCodecTypes(value VideoCodecType /* not a class type */)
	AvailableVideoPixelFormatTypes() unsafe.Pointer
	SetAvailableVideoPixelFormatTypes(value unsafe.Pointer)
	DeliversPreviewSizedOutputBuffers() bool /* primitive/slice/pointer */
	SetDeliversPreviewSizedOutputBuffers(value bool /* primitive/slice/pointer */)
	PreparesCellularRadioForNetworkConnection() bool /* primitive/slice/pointer */
	SetPreparesCellularRadioForNetworkConnection(value bool /* primitive/slice/pointer */)
	PreservesDynamicHDRMetadata() bool /* primitive/slice/pointer */
	SetPreservesDynamicHDRMetadata(value bool /* primitive/slice/pointer */)
	RecommendedMediaTimeScaleForAssetWriter() TimeScale /* not a class type */
	SetRecommendedMediaTimeScaleForAssetWriter(value TimeScale /* not a class type */)
	SampleBufferDelegate() CaptureVideoDataOutputSampleBufferDelegate /* not a class type */
	SetSampleBufferDelegate(value CaptureVideoDataOutputSampleBufferDelegate /* not a class type */)
	// methods:
	RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* not a class type */, outputFileType FileType /* not a class type */) []MetadataItem /* primitive/slice/pointer */
	SetSampleBufferDelegateQueue(sampleBufferDelegate objectivec.IObject, sampleBufferCallbackQueue unsafe.Pointer)
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureVideoDataOutputClass) Alloc() CaptureVideoDataOutput {
	rv := objc.Send[CaptureVideoDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMovieMetadata(forVideoCodecType:assetWriterOutputFileType:)
func (c_ CaptureVideoDataOutput) RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType /* not a class type */, outputFileType FileType /* not a class type */) []MetadataItem /* primitive/slice/pointer */ {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("recommendedMovieMetadataForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}


// Sets the sample buffer delegate and the queue for invoking callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate objectivec.IObject, sampleBufferCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}


// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) AlwaysDiscardsLateVideoFrames() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("alwaysDiscardsLateVideoFrames"))
	return rv
}


// Indicates whether to drop video frames if they arrive late.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c_ CaptureVideoDataOutput) SetAlwaysDiscardsLateVideoFrames(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlwaysDiscardsLateVideoFrames:"), value)
}


// The minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/minFrameDuration
func (c_ CaptureVideoDataOutput) MinFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/minFrameDuration
func (c_ CaptureVideoDataOutput) SetMinFrameDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// The queue on which the system invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferCallbackQueue
func (c_ CaptureVideoDataOutput) SampleBufferCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) VideoSettings() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("videoSettings"))
	return rv
}


// A dictionary that contains the compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c_ CaptureVideoDataOutput) SetVideoSettings(value foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoSettings:"), value)
}


// A Boolean value that indicates whether the output automatically configures the size of output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/automaticallyconfiguresoutputbufferdimensions
func (c_ CaptureVideoDataOutput) AutomaticallyConfiguresOutputBufferDimensions() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresOutputBufferDimensions"))
	return rv
}


// A Boolean value that indicates whether the output automatically configures the size of output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/automaticallyconfiguresoutputbufferdimensions
func (c_ CaptureVideoDataOutput) SetAutomaticallyConfiguresOutputBufferDimensions(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyConfiguresOutputBufferDimensions:"), value)
}


// The video codecs that the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideocodectypes
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypes() VideoCodecType /* not a class type */ {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}


// The video codecs that the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideocodectypes
func (c_ CaptureVideoDataOutput) SetAvailableVideoCodecTypes(value VideoCodecType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableVideoCodecTypes:"), value)
}


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) AvailableVideoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableVideoPixelFormatTypes"))
	return rv
}


// The video pixel formats the output supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c_ CaptureVideoDataOutput) SetAvailableVideoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableVideoPixelFormatTypes:"), value)
}


// A Boolean value that indicates whether the output is configured to deliver preview-sized buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/deliverspreviewsizedoutputbuffers
func (c_ CaptureVideoDataOutput) DeliversPreviewSizedOutputBuffers() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("deliversPreviewSizedOutputBuffers"))
	return rv
}


// A Boolean value that indicates whether the output is configured to deliver preview-sized buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/deliverspreviewsizedoutputbuffers
func (c_ CaptureVideoDataOutput) SetDeliversPreviewSizedOutputBuffers(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliversPreviewSizedOutputBuffers:"), value)
}


// Indicates whether the receiver should prepare the cellular radio for imminent network activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/preparescellularradiofornetworkconnection
func (c_ CaptureVideoDataOutput) PreparesCellularRadioForNetworkConnection() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("preparesCellularRadioForNetworkConnection"))
	return rv
}


// Indicates whether the receiver should prepare the cellular radio for imminent network activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/preparescellularradiofornetworkconnection
func (c_ CaptureVideoDataOutput) SetPreparesCellularRadioForNetworkConnection(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreparesCellularRadioForNetworkConnection:"), value)
}


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/preservesdynamichdrmetadata
func (c_ CaptureVideoDataOutput) PreservesDynamicHDRMetadata() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesDynamicHDRMetadata"))
	return rv
}


// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/preservesdynamichdrmetadata
func (c_ CaptureVideoDataOutput) SetPreservesDynamicHDRMetadata(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesDynamicHDRMetadata:"), value)
}


// Indicates the recommended media timescale for the video track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/recommendedmediatimescaleforassetwriter
func (c_ CaptureVideoDataOutput) RecommendedMediaTimeScaleForAssetWriter() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](c_.ID, objc.Sel("recommendedMediaTimeScaleForAssetWriter"))
	return rv
}


// Indicates the recommended media timescale for the video track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/recommendedmediatimescaleforassetwriter
func (c_ CaptureVideoDataOutput) SetRecommendedMediaTimeScaleForAssetWriter(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecommendedMediaTimeScaleForAssetWriter:"), value)
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/samplebufferdelegate
func (c_ CaptureVideoDataOutput) SampleBufferDelegate() CaptureVideoDataOutputSampleBufferDelegate /* not a class type */ {
	rv := objc.Send[CaptureVideoDataOutputSampleBufferDelegate](c_.ID, objc.Sel("sampleBufferDelegate"))
	return rv
}


// The capture object’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/samplebufferdelegate
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegate(value CaptureVideoDataOutputSampleBufferDelegate /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferDelegate:"), value)
}



