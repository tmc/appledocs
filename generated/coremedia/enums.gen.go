// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

// Enum types and constants
// CMPackingType - The type of packing within each video frame, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType
type CMPackingType uint

const (
	// kCMPackingType_None - Each frame contains only a single image, and isn’t frame-packed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType/none
	kCMPackingType_None CMPackingType = 0
	// kCMPackingType_OverUnder - The video contains packed frames that have a left eye image on the top and right eye image on the bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType/overUnder
	kCMPackingType_OverUnder CMPackingType = 0
)

// CMProjectionType - Constants describing the projection surface information in a 3D video buffer or channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType
type CMProjectionType uint

const (
	// kCMProjectionType_Fisheye - Video content displays as a fisheye projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/fisheye
	kCMProjectionType_Fisheye CMProjectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/parametricImmersive
	kCMProjectionType_ParametricImmersive CMProjectionType = 0
)

// CMStereoViewComponents - Constants describing the stereo views contained within a buffer or channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents
type CMStereoViewComponents uint

const (
	// kCMStereoView_None - A constant for video metadata to have no available stereo frames.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/kCMStereoView_None
	kCMStereoView_None CMStereoViewComponents = 0
)

// CMStereoViewInterpretationOptions - Create a set of stereo view interpretation options from a constant.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions
type CMStereoViewInterpretationOptions uint

const (
	// kCMStereoViewInterpretation_AdditionalViews - A flag indicating that the video content contains additional views beyond the left or right eye.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions/additionalViews
	kCMStereoViewInterpretation_AdditionalViews CMStereoViewInterpretationOptions = 0
	// kCMStereoViewInterpretation_Default - The default options for stereo video views.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions/kCMStereoViewInterpretation_Default
	kCMStereoViewInterpretation_Default CMStereoViewInterpretationOptions = 0
)

// CMTagError - Core media tagging errors reported by the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagError
type CMTagError uint

// CMTaggedBufferGroupError - Error codes returned by Core Media when working with tagged buffer groups.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupError
type CMTaggedBufferGroupError uint

const (
	// kCMTaggedBufferGroupError_AllocationFailed - Indicates an internal allocation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupError/kCMTaggedBufferGroupError_AllocationFailed
	kCMTaggedBufferGroupError_AllocationFailed CMTaggedBufferGroupError = 0
	// kCMTaggedBufferGroupError_InternalError - Indicates an error occurred inside of the Core Media framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupError/kCMTaggedBufferGroupError_InternalError
	kCMTaggedBufferGroupError_InternalError CMTaggedBufferGroupError = 0
	// kCMTaggedBufferGroupError_ParamErr - Indicates a parameter to a function was of the wrong type or didn’t meet a necessary condition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupError/kCMTaggedBufferGroupError_ParamErr
	kCMTaggedBufferGroupError_ParamErr CMTaggedBufferGroupError = 0
)


