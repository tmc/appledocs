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
	// kCMPackingType_SideBySide - The video contains packed frames that have a left eye image on the left and right eye image on the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType/sideBySide
	kCMPackingType_SideBySide CMPackingType = 0
)


// CMProjectionType - Constants describing the projection surface information in a 3D video buffer or channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType
type CMProjectionType uint

const (
	// kCMProjectionType_Equirectangular - Video content displays as a 360 degree equirectangular projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/equirectangular
	kCMProjectionType_Equirectangular CMProjectionType = 0
	// kCMProjectionType_Fisheye - Video content displays as a fisheye projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/fisheye
	kCMProjectionType_Fisheye CMProjectionType = 0
	// kCMProjectionType_HalfEquirectangular - Video content displays as a 180 degree equirectangular projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/halfEquirectangular
	kCMProjectionType_HalfEquirectangular CMProjectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/parametricImmersive
	kCMProjectionType_ParametricImmersive CMProjectionType = 0
	// kCMProjectionType_Rectangular - Video content displays on a flat, rectangular 2D surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/rectangular
	kCMProjectionType_Rectangular CMProjectionType = 0
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
	// kCMStereoView_LeftEye - The stereo video track includes a left eye layer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/leftEye
	kCMStereoView_LeftEye CMStereoViewComponents = 0
	// kCMStereoView_RightEye - The stereo video track includes a right eye layer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/rightEye
	kCMStereoView_RightEye CMStereoViewComponents = 0
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
	// kCMStereoViewInterpretation_StereoOrderReversed - Changes the default ordering of eye data, switching it from left-to-right to right-to-left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions/stereoOrderReversed
	kCMStereoViewInterpretation_StereoOrderReversed CMStereoViewInterpretationOptions = 0
)


// CMTagCategory - A 64-bit representation of a tag’s category.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory
type CMTagCategory uint

const (
	// kCMTagCategory_ChannelID - A category used for tagging a channel ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_ChannelID
	kCMTagCategory_ChannelID CMTagCategory = 0
	// kCMTagCategory_MediaSubType - A category used for tagging media subtype metadata.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_MediaSubType
	kCMTagCategory_MediaSubType CMTagCategory = 0
	// kCMTagCategory_MediaType - A category used for tagging media type metadata.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_MediaType
	kCMTagCategory_MediaType CMTagCategory = 0
	// kCMTagCategory_PackingType - A category used for tagging frame-packing information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_PackingType
	kCMTagCategory_PackingType CMTagCategory = 0
	// kCMTagCategory_PixelFormat - A category used for tagging pixel format information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_PixelFormat
	kCMTagCategory_PixelFormat CMTagCategory = 0
	// kCMTagCategory_ProjectionType - A category used for tagging projection surface information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_ProjectionType
	kCMTagCategory_ProjectionType CMTagCategory = 0
	// kCMTagCategory_StereoView - A category used for tagging eye information for 3D video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_StereoView
	kCMTagCategory_StereoView CMTagCategory = 0
	// kCMTagCategory_StereoViewInterpretation - A category used for tagging how to interpret stereo view metadata.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_StereoViewInterpretation
	kCMTagCategory_StereoViewInterpretation CMTagCategory = 0
	// kCMTagCategory_TrackID - A category used for tagging a track ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_TrackID
	kCMTagCategory_TrackID CMTagCategory = 0
	// kCMTagCategory_Undefined - An unknown or undefined tag category.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_Undefined
	kCMTagCategory_Undefined CMTagCategory = 0
	// kCMTagCategory_VideoLayerID - A category used for tagging a video layer ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCategory/kCMTagCategory_VideoLayerID
	kCMTagCategory_VideoLayerID CMTagCategory = 0
)


// CMTagCollectionError - Error codes returned by Core Media when working with tag collections.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError
type CMTagCollectionError uint

const (
	// kCMTagCollectionError_AllocationFailed - Indicates an internal allocation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_AllocationFailed
	kCMTagCollectionError_AllocationFailed CMTagCollectionError = 0
	// kCMTagCollectionError_ExhaustedBufferSize - Indicates that a buffer was smaller than the number of requested tags.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_ExhaustedBufferSize
	kCMTagCollectionError_ExhaustedBufferSize CMTagCollectionError = 0
	// kCMTagCollectionError_InternalError - Indicates an error occurred inside of the Core Media framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_InternalError
	kCMTagCollectionError_InternalError CMTagCollectionError = 0
	// kCMTagCollectionError_InvalidTag - Indicates that the collection contains an invalid tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_InvalidTag
	kCMTagCollectionError_InvalidTag CMTagCollectionError = 0
	// kCMTagCollectionError_InvalidTagCollectionData - Indicates that a Core Foundation data instance failed to initialize a new tag collection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_InvalidTagCollectionData
	kCMTagCollectionError_InvalidTagCollectionData CMTagCollectionError = 0
	// kCMTagCollectionError_InvalidTagCollectionDataVersion - Indicates that a Core Foundation data instance failed to initialize a new tag collection due to a versioning problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_InvalidTagCollectionDataVersion
	kCMTagCollectionError_InvalidTagCollectionDataVersion CMTagCollectionError = 0
	// kCMTagCollectionError_InvalidTagCollectionDictionary - Indicates that a Core Foundation dictionary instance failed to initialize a new tag collection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_InvalidTagCollectionDictionary
	kCMTagCollectionError_InvalidTagCollectionDictionary CMTagCollectionError = 0
	// kCMTagCollectionError_NotYetImplemented - Indicates a function lacks a necessary backing implementation in Core Media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_NotYetImplemented
	kCMTagCollectionError_NotYetImplemented CMTagCollectionError = 0
	// kCMTagCollectionError_ParamErr - Indicates a parameter to a function was of the wrong type or didn’t meet a necessary condition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_ParamErr
	kCMTagCollectionError_ParamErr CMTagCollectionError = 0
	// kCMTagCollectionError_TagNotFound - Indicates that there was no match in a collection for a tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionError/kCMTagCollectionError_TagNotFound
	kCMTagCollectionError_TagNotFound CMTagCollectionError = 0
)


// CMTagDataType - The data type of a tag’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType
type CMTagDataType uint

const (
	// kCMTagDataType_Flags - The tag value is a 64-bit wide bitflag field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType/kCMTagDataType_Flags
	kCMTagDataType_Flags CMTagDataType = 0
	// kCMTagDataType_Float64 - The tag value is a 64-bit floating point number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType/kCMTagDataType_Float64
	kCMTagDataType_Float64 CMTagDataType = 0
	// kCMTagDataType_Invalid - The tag value isn’t associated with any known data type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType/kCMTagDataType_Invalid
	kCMTagDataType_Invalid CMTagDataType = 0
	// kCMTagDataType_OSType - The tag value is a 64-bit identifier used by the operating system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType/kCMTagDataType_OSType
	kCMTagDataType_OSType CMTagDataType = 0
	// kCMTagDataType_SInt64 - The tag value is a signed 64-bit integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagDataType/kCMTagDataType_SInt64
	kCMTagDataType_SInt64 CMTagDataType = 0
)


// CMTagError - Core media tagging errors reported by the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagError
type CMTagError uint

const (
	// kCMTagError_AllocationFailed - An error where the system can’t allocate enough memory for the tag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagError/kCMTagError_AllocationFailed
	kCMTagError_AllocationFailed CMTagError = 0
	// kCMTagError_ParamErr - An error where input or output parameters didn’t match the requirements of Core Media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagError/kCMTagError_ParamErr
	kCMTagError_ParamErr CMTagError = 0
)


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


// CMTimeFlags - A structure that defines the flags for a time value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags
type CMTimeFlags uint

const (
	// kCMTimeFlags_HasBeenRounded - A flag that indicates a previous time calculation rounded the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/hasBeenRounded
	kCMTimeFlags_HasBeenRounded CMTimeFlags = 0
	// kCMTimeFlags_ImpliedValueFlagsMask - A flag that indicates the time is positive or negative infinity, or indefinite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/impliedValueFlagsMask
	kCMTimeFlags_ImpliedValueFlagsMask CMTimeFlags = 0
	// kCMTimeFlags_Indefinite - A flag that indicates the time is indefinite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/indefinite
	kCMTimeFlags_Indefinite CMTimeFlags = 0
	// kCMTimeFlags_NegativeInfinity - A flag that indicates the time is negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/negativeInfinity
	kCMTimeFlags_NegativeInfinity CMTimeFlags = 0
	// kCMTimeFlags_PositiveInfinity - A flag that indicates the time is positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/positiveInfinity
	kCMTimeFlags_PositiveInfinity CMTimeFlags = 0
	// kCMTimeFlags_Valid - A flag that indicates a time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/valid
	kCMTimeFlags_Valid CMTimeFlags = 0
)


// CMTimeRoundingMethod - An enumeration of rounding methods to use when performing time calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod
type CMTimeRoundingMethod uint

const (
	// kCMTimeRoundingMethod_Default - The default rounding method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/default
	kCMTimeRoundingMethod_Default CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_QuickTime - Rounds using the QuickTime method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/quickTime
	kCMTimeRoundingMethod_QuickTime CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_RoundAwayFromZero - Rounds away from zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/roundAwayFromZero
	kCMTimeRoundingMethod_RoundAwayFromZero CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_RoundHalfAwayFromZero - Rounds half away from zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/roundHalfAwayFromZero
	kCMTimeRoundingMethod_RoundHalfAwayFromZero CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_RoundTowardNegativeInfinity - Rounds toward negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/roundTowardNegativeInfinity
	kCMTimeRoundingMethod_RoundTowardNegativeInfinity CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_RoundTowardPositiveInfinity - Rounds toward positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/roundTowardPositiveInfinity
	kCMTimeRoundingMethod_RoundTowardPositiveInfinity CMTimeRoundingMethod = 0
	// kCMTimeRoundingMethod_RoundTowardZero - Rounds toward zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod/roundTowardZero
	kCMTimeRoundingMethod_RoundTowardZero CMTimeRoundingMethod = 0
)


