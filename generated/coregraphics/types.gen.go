// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics
import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
)


// C struct types
// CGBitmapParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapParameters-1cm7j
type CGBitmapParameters struct {
	AlignedBytesPerRow uintptr
	ByteOrder int32
	BytesPerPixel uintptr
	ColorSpace ColorSpaceRef
	Component Component
	EdrTargetHeadroom float32
	Format ImagePixelFormatInfo
	HasPremultipliedAlpha bool
	Height uintptr
	Layout BitmapLayout
	Width uintptr
}

// CGColorBufferFormat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorBufferFormat
type CGColorBufferFormat struct {
	BitmapInfo BitmapInfo
	BitsPerComponent uintptr
	BitsPerPixel uintptr
	BytesPerRow uintptr
	Version uint32
}

// CGColorDataFormat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorDataFormat
type CGColorDataFormat struct {
	Bitmap_info BitmapInfo
	Bits_per_component uintptr
	Bytes_per_row uintptr
	Colorspace_info TypeRef
	Decode *float64
	Intent ColorRenderingIntent
	Version uint32
}

// CGContentInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContentInfo
type CGContentInfo struct {
	ContentColorModels ColorModel
	DeepestImageComponent Component
	HasTransparency bool
	HasWideGamut bool
	LargestContentHeadroom float32
}

// CGContentToneMappingInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContentToneMappingInfo-c.struct
type CGContentToneMappingInfo struct {
	Method ToneMapping
	Options DictionaryRef
}

// CGDataConsumerCallbacks - A structure that contains pointers to callback functions that manage the copying of data for a data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerCallbacks
type CGDataConsumerCallbacks struct {
	PutBytes DataConsumerPutBytesCallback // A pointer to a function that copies data to the data consumer. For more information, see  .
	ReleaseConsumer DataConsumerReleaseInfoCallback // A pointer to a function that handles clean-up for the data consumer, or  .
}

// CGDataProviderDirectCallbacks - Defines pointers to client-defined callback functions that manage the sending of data for a direct-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderDirectCallbacks
type CGDataProviderDirectCallbacks struct {
	GetBytePointer DataProviderGetBytePointerCallback // A pointer to a function that returns a pointer to the provider’s data. For more information, see  .
	GetBytesAtPosition DataProviderGetBytesAtPositionCallback // A pointer to a function that copies data from the provider.
	ReleaseBytePointer DataProviderReleaseBytePointerCallback // A pointer to a function that Core Graphics calls to release a pointer to the provider’s data. For more information, see  .
	ReleaseInfo DataProviderReleaseInfoCallback // A pointer to a function that handles clean-up for the data provider, or  . For more information, see  .
	Version unsafe.Pointer // The version of this structure. It should be set to 0.
}

// CGDataProviderSequentialCallbacks - Defines a structure containing pointers to client-defined callback functions that manage the sending of data for a sequential-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderSequentialCallbacks
type CGDataProviderSequentialCallbacks struct {
	GetBytes DataProviderGetBytesCallback // A pointer to a function that copies data from the provider. For more information, see  .
	ReleaseInfo DataProviderReleaseInfoCallback // A pointer to a function that handles clean-up for the data provider, or  . For more information, see  .
	Rewind DataProviderRewindCallback // A pointer to a function Core Graphics calls to return the provider to the beginning of the data stream. For more information, see  .
	SkipForward DataProviderSkipForwardCallback // A pointer to a function that Core Graphics calls to advance the stream of data supplied by the provider.
	Version unsafe.Pointer // The version of this structure. It should be set to 0.
}

// CGDeviceColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDeviceColor
type CGDeviceColor struct {
	Blue float32
	Green float32
	Red float32
}

// CGFunctionCallbacks - A structure that contains callbacks needed by a 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunctionCallbacks
type CGFunctionCallbacks struct {
	Evaluate FunctionEvaluateCallback // The callback that evaluates the function.
	ReleaseInfo FunctionReleaseInfoCallback // If non- ,the callback used to release the   parameterpassed to  .
	Version unsafe.Pointer // The structure version number. For this structure,the version should be  .
}

// CGPSConverterCallbacks - A structure for holding the callbacks provided when you create a PostScript converter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverterCallbacks
type CGPSConverterCallbacks struct {
	BeginDocument PSConverterBeginDocumentCallback // The callback called at the beginning of the conversion of the PostScript document, or  .
	BeginPage PSConverterBeginPageCallback // The callback called at the start of the conversion of each page in the PostScript document, or  .
	EndDocument PSConverterEndDocumentCallback // The callback called at the end of conversion of the PostScript document, or  .
	EndPage PSConverterEndPageCallback // The callback called at the end of the conversion of each page in the PostScript document, or  .
	NoteMessage PSConverterMessageCallback // The callback called to pass any messages that might result during the conversion, or  .
	NoteProgress PSConverterProgressCallback // The callback called periodically during the conversion to indicate that conversion is proceeding, or  .
	ReleaseInfo PSConverterReleaseInfoCallback // The callback called when the converter is deallocated, or  .
	Version unsafe.Pointer // The version number of the structure passed in as a parameter to the converter creation functions. The structure defined below is version  .
}

// CGPathElement - A data structure that provides information about a path element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElement
type CGPathElement struct {
	Points corefoundation.CGPoint // An array of one or more points that serve as arguments.
	Type PathElementType // An element type (or operation).
}

// CGPatternCallbacks - A structure that holds a version and two callback functions for drawing a custom pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternCallbacks
type CGPatternCallbacks struct {
	DrawPattern PatternDrawPatternCallback // A pointer to a custom function that draws thepattern. For information about this callback function, see  .
	ReleaseInfo PatternReleaseInfoCallback // An optional pointer to a custom function that’sinvoked when the pattern is released.  .
	Version unsafe.Pointer // The version of the structure passed in as a parameterto the  . Forthis version of the structure, you should set this value to zero.
}

// CGScreenUpdateMoveDelta - The distance, in pixel units, that an onscreen region moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateMoveDelta
type CGScreenUpdateMoveDelta struct {
	DX int32
	DY int32
}





