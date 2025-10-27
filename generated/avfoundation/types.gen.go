// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation
import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)


// C struct types
// AVCaptionDimension - A structure that defines a caption dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionDimension
type AVCaptionDimension struct {
	Units CaptionUnitsType // The units of the coordinate, such as cells or points.
	Value float64 // The value of the coordinate or length.
}

// AVCaptionPoint - A structure that defines the origin point for a caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionPoint
type AVCaptionPoint struct {
	X CaptionDimension // The point’s x coordinate.
	Y CaptionDimension // The point’s y coordinate.
}

// AVCaptionSize - A structure that defines the height and width of a caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSize
type AVCaptionSize struct {
	Height CaptionDimension // The height of the caption.
	Width CaptionDimension // The width of the caption.
}

// WhiteBalanceChromaticityValues - A structure that defines CIE 1931 xy chromaticity values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceChromaticityValues
type WhiteBalanceChromaticityValues struct {
}

// WhiteBalanceGains - A structure that defines RGB white balance gain values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceGains
type WhiteBalanceGains struct {
}

// WhiteBalanceTemperatureAndTintValues - A structure that defines temperature and tint values correlated to a white-balance color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceTemperatureAndTintValues
type WhiteBalanceTemperatureAndTintValues struct {
}

// AVCaptureTimecode - This structure represents a timecode, adhering to SMPTE standards, which define precise time information and associated timestamps for video or audio synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode
type AVCaptureTimecode struct {
	FrameDuration objectivec.IObject // Frame duration of the timecode. If unknown, the value is  .
	Frames uint8 // Frame component of the timecode, indicating the frame count within the second.
	Hours uint8 // Time component representing the current timecode in hours.
	Minutes uint8 // Time component representing the current timecode in minutes.
	Seconds uint8 // Time component representing the current timecode in seconds.
	SourceType CaptureTimecodeSourceType // Source type of the timecode, indicating the emitter, carriage, or transport mechanism.
	UserBits uint32 // A 32-bit field carrying SMPTE user bits, which are not strictly standardized. User bits are often used for additional metadata such as scene-take information, reel numbers, or dates, but their exact usage is application-dependent.
}

// AVEdgeWidths - A structure that defines edge processing region widths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVEdgeWidths
type AVEdgeWidths struct {
	Bottom float64 // The bottom-edge width.
	Left float64 // The left-edge width.
	Right float64 // The right-edge width.
	Top float64 // The top-edge width.
}

// AVPixelAspectRatio - A structure that defines a pixel aspect ratio for a rendering context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPixelAspectRatio
type AVPixelAspectRatio struct {
	HorizontalSpacing int // The pixel aspect ratio’s horizontal spacing value.
	VerticalSpacing int // The pixel aspect ratio’s vertical spacing value.
}

// AVSampleCursorAudioDependencyInfo - A structure that describes the independent decodability of audio samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorAudioDependencyInfo
type AVSampleCursorAudioDependencyInfo struct {
	AudioSampleIsIndependentlyDecodable bool // A Boolean value indicating whether the sample is independently decodable.
	AudioSamplePacketRefreshCount int // The number of samples, starting at the current sample, that must be fed to the decoder to achieve full decoder refresh.
}

// AVSampleCursorChunkInfo - A value that provides information about a chunk of media samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorChunkInfo
type AVSampleCursorChunkInfo struct {
	ChunkHasUniformFormatDescriptions bool // The samples in the chunk have the same format description.
	ChunkHasUniformSampleDurations bool // The samples in the chunk have the same duration.
	ChunkHasUniformSampleSizes bool // The samples in the chunk occupy the same number of bytes in storage.
	ChunkSampleCount int64 // The count of media samples in the chunk.
}

// AVSampleCursorDependencyInfo - A value for describing dependencies between a media sample and other media samples in the same sample sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorDependencyInfo
type AVSampleCursorDependencyInfo struct {
	SampleDependsOnOthers bool // A Boolean value that determines whether the sample depends on other samples.
	SampleHasDependentSamples bool // A Boolean value that determines whether the sample has dependent samples.
	SampleHasRedundantCoding bool // A Boolean value that determines whether the sample has redundant coding.
	SampleIndicatesWhetherItDependsOnOthers bool // A Boolean value that determines whether the sample indicates that it depends on other samples.
	SampleIndicatesWhetherItHasDependentSamples bool // A Boolean value that determines whether the sample indicates if other samples depend on it.
	SampleIndicatesWhetherItHasRedundantCoding bool // A Boolean value that determines whether the sample indicates that it has redundant coding.
}

// AVSampleCursorStorageRange - A structure that indicates the offset and length of storage for a media sample or its chunk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorStorageRange
type AVSampleCursorStorageRange struct {
	Length int64 // The count of bytes of storage that a media sample or its chunk occupies.
	Offset int64 // The offset of the first byte of storage that a media sample or its chunk occupies.
}

// AVSampleCursorSyncInfo - A structure that describes the attributes of media samples to consider when resynchronizing a decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorSyncInfo
type AVSampleCursorSyncInfo struct {
	SampleIsDroppable bool // A Boolean value that indicates whether a sample is droppable.
	SampleIsFullSync bool // A Boolean value that indicates whether a sample is a full sync sample.
	SampleIsPartialSync bool // A Boolean value that indicates whether a sample is a partial sync sample.
}





