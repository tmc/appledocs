// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation


import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
	corevideo "github.com/tmc/appledocs/generated/corevideo"
)


// AVFoundation Functions (10 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AVCaptionDimensionMake func(float64, CaptionUnitsType) CaptionDimension
	_AVCaptionPointMake func(CaptionDimension, CaptionDimension) CaptionPoint
	_AVCaptionSizeMake func(CaptionDimension, CaptionDimension) CaptionSize
	_CMTagCollectionCreateWithVideoOutputPreset func(AllocatorRef, TagCollectionVideoOutputPreset, unsafe.Pointer) unsafe.Pointer
	_AVCaptureReactionSystemImageNameForType func(CaptureReactionType) unsafe.Pointer
	_AVCaptureTimecodeAdvancedByFrames func(CaptureTimecode, int64) CaptureTimecode
	_AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp func(CaptureTimecode, corevideo.Time) SampleBufferRef
	_AVCaptureTimecodeCreateMetadataSampleBufferForDuration func(CaptureTimecode, corevideo.Time) SampleBufferRef
	_AVMakeRectWithAspectRatioInsideRect func(corefoundation.CGSize, corefoundation.CGRect) corefoundation.CGRect
	_AVSampleBufferAttachContentKey func(SampleBufferRef, unsafe.Pointer, unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AVCaptionDimensionMake, lib, "AVCaptionDimensionMake")
	tryRegister(&_AVCaptionPointMake, lib, "AVCaptionPointMake")
	tryRegister(&_AVCaptionSizeMake, lib, "AVCaptionSizeMake")
	tryRegister(&_CMTagCollectionCreateWithVideoOutputPreset, lib, "CMTagCollectionCreateWithVideoOutputPreset")
	tryRegister(&_AVCaptureReactionSystemImageNameForType, lib, "AVCaptureReactionSystemImageNameForType")
	tryRegister(&_AVCaptureTimecodeAdvancedByFrames, lib, "AVCaptureTimecodeAdvancedByFrames")
	tryRegister(&_AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp, lib, "AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp")
	tryRegister(&_AVCaptureTimecodeCreateMetadataSampleBufferForDuration, lib, "AVCaptureTimecodeCreateMetadataSampleBufferForDuration")
	tryRegister(&_AVMakeRectWithAspectRatioInsideRect, lib, "AVMakeRectWithAspectRatioInsideRect")
	tryRegister(&_AVSampleBufferAttachContentKey, lib, "AVSampleBufferAttachContentKey")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Creates a caption dimension with a value and unit type.
//
// Added in macOS 12.0.
// Creates a caption dimension with a value and unit type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionDimensionMake
func AVCaptionDimensionMake(value float64, units CaptionUnitsType) CaptionDimension {
	return _AVCaptionDimensionMake(value, units)
}

// Creates a caption point with the specified x and y positions.
//
// Added in macOS 12.0.
// Creates a caption point with the specified x and y positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionPointMake
func AVCaptionPointMake(x CaptionDimension, y CaptionDimension) CaptionPoint {
	return _AVCaptionPointMake(x, y)
}

// Creates a caption size with the specified width and height.
//
// Added in macOS 12.0.
// Creates a caption size with the specified width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSizeMake
func AVCaptionSizeMake(width CaptionDimension, height CaptionDimension) CaptionSize {
	return _AVCaptionSizeMake(width, height)
}

// Creates a collection with the required tags to describe the specified video output requirements.
//
// Added in macOS 14.2.
// Creates a collection with the required tags to describe the specified video output requirements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionCreateWithVideoOutputPreset
func CMTagCollectionCreateWithVideoOutputPreset(allocator AllocatorRef, preset TagCollectionVideoOutputPreset, newCollectionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionCreateWithVideoOutputPreset(allocator, preset, newCollectionOut)
}

// Returns the name of a system image that displays the recommended iconography for a specified reaction type.
//
// Added in macOS 14.0.
// Returns the name of a system image that displays the recommended iconography for a specified reaction type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionType/systemImageName
func AVCaptureReactionSystemImageNameForType(reactionType CaptureReactionType) unsafe.Pointer {
	return _AVCaptureReactionSystemImageNameForType(reactionType)
}

// Generates a new timecode by adding a specified number of frames to the given timecode, handling overflow for seconds, minutes, and hours.
//
// Added in macOS 26.0.
// Generates a new timecode by adding a specified number of frames to the given timecode, handling overflow for seconds, minutes, and hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/advanced(_:by:)
func AVCaptureTimecodeAdvancedByFrames(timecode CaptureTimecode, framesToAdd int64) CaptureTimecode {
	return _AVCaptureTimecodeAdvancedByFrames(timecode, framesToAdd)
}

// Creates a sample buffer containing Timecode Media Description metadata for integration with a video track.
//
// Added in macOS 26.0.
// Creates a sample buffer containing Timecode Media Description metadata for integration with a video track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:associatedWithPresentationTimeStamp:)
func AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode CaptureTimecode, presentationTimeStamp corevideo.Time) SampleBufferRef {
	return _AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode, presentationTimeStamp)
}

// Creates a sample buffer containing Timecode Media Description metadata for a specified duration.
//
// Added in macOS 26.0.
// Creates a sample buffer containing Timecode Media Description metadata for a specified duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:forDuration:)
func AVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode CaptureTimecode, duration corevideo.Time) SampleBufferRef {
	return _AVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode, duration)
}

// Returns a scaled rectangle that maintains the specified aspect ratio within a bounding rectangle.
//
// Added in macOS 10.7.
// Returns a scaled rectangle that maintains the specified aspect ratio within a bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMakeRect(aspectRatio:insideRect:)
func AVMakeRectWithAspectRatioInsideRect(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) corefoundation.CGRect {
	return _AVMakeRectWithAspectRatioInsideRect(aspectRatio, boundingRect)
}

// Attaches a content key to a sample buffer for the purpose of content decryption.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.10.
// Attaches a content key to a sample buffer for the purpose of content decryption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAttachContentKey(_:_:_:)
func AVSampleBufferAttachContentKey(sbuf SampleBufferRef, contentKey unsafe.Pointer, outError unsafe.Pointer) bool {
	return _AVSampleBufferAttachContentKey(sbuf, contentKey, outError)
}




