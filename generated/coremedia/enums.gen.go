// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

// Enum types and constants
// CMPackingType - The type of packing within each video frame, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType
type PackingType uint

const (
	// kCMPackingType_None - Each frame contains only a single image, and isn’t frame-packed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType/none
	kCMPackingType_None PackingType = 0
	// kCMPackingType_OverUnder - The video contains packed frames that have a left eye image on the top and right eye image on the bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMPackingType/overUnder
	kCMPackingType_OverUnder PackingType = 0
)

// CMProjectionType - Constants describing the projection surface information in a 3D video buffer or channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType
type ProjectionType uint

const (
	// kCMProjectionType_HalfEquirectangular - Video content displays as a 180 degree equirectangular projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/halfEquirectangular
	kCMProjectionType_HalfEquirectangular ProjectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/parametricImmersive
	kCMProjectionType_ParametricImmersive ProjectionType = 0
	// kCMProjectionType_Rectangular - Video content displays on a flat, rectangular 2D surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMProjectionType/rectangular
	kCMProjectionType_Rectangular ProjectionType = 0
)

// CMStereoViewComponents - Constants describing the stereo views contained within a buffer or channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents
type StereoViewComponents uint

const (
	// kCMStereoView_None - A constant for video metadata to have no available stereo frames.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/kCMStereoView_None
	kCMStereoView_None StereoViewComponents = 0
)

// CMStereoViewInterpretationOptions - Create a set of stereo view interpretation options from a constant.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions
type StereoViewInterpretationOptions uint

const (
	// kCMStereoViewInterpretation_AdditionalViews - A flag indicating that the video content contains additional views beyond the left or right eye.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions/additionalViews
	kCMStereoViewInterpretation_AdditionalViews StereoViewInterpretationOptions = 0
	// kCMStereoViewInterpretation_Default - The default options for stereo video views.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewInterpretationOptions/kCMStereoViewInterpretation_Default
	kCMStereoViewInterpretation_Default StereoViewInterpretationOptions = 0
)

// CMTagError - Core media tagging errors reported by the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTagError
type TagError uint

// CMTimeFlags - A structure that defines the flags for a time value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags
type TimeFlags uint

const (
	// kCMTimeFlags_HasBeenRounded - A flag that indicates a previous time calculation rounded the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/hasBeenRounded
	kCMTimeFlags_HasBeenRounded TimeFlags = 0
	// kCMTimeFlags_ImpliedValueFlagsMask - A flag that indicates the time is positive or negative infinity, or indefinite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/impliedValueFlagsMask
	kCMTimeFlags_ImpliedValueFlagsMask TimeFlags = 0
	// kCMTimeFlags_Indefinite - A flag that indicates the time is indefinite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/indefinite
	kCMTimeFlags_Indefinite TimeFlags = 0
	// kCMTimeFlags_NegativeInfinity - A flag that indicates the time is negative infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/negativeInfinity
	kCMTimeFlags_NegativeInfinity TimeFlags = 0
	// kCMTimeFlags_PositiveInfinity - A flag that indicates the time is positive infinity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/positiveInfinity
	kCMTimeFlags_PositiveInfinity TimeFlags = 0
	// kCMTimeFlags_Valid - A flag that indicates a time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeFlags/valid
	kCMTimeFlags_Valid TimeFlags = 0
)

// CMTimeRoundingMethod - An enumeration of rounding methods to use when performing time calculations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/CMTimeRoundingMethod
type TimeRoundingMethod uint


