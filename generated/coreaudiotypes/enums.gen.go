// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes

/* debug [enums.gen.go]: Generating 8 enums for CoreAudioTypes */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum AVAudioSessionErrorCode (15 cases) */
// AVAudioSessionErrorCode - Codes that describe error conditions that may occur when performing audio session operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode
type AVAudioSessionErrorCode uint

const (
	// AVAudioSessionErrorCodeBadParam - An error code that indicates an attempt to set a property to an illegal value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/badParam
	AVAudioSessionErrorCodeBadParam AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeCannotInterruptOthers - An error code that indictates an attempt to make a nonmixable audio session active while the app was in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotInterruptOthers
	AVAudioSessionErrorCodeCannotInterruptOthers AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeCannotStartPlaying - An error code that indicates an attempt to start audio playback when it wasn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartPlaying
	AVAudioSessionErrorCodeCannotStartPlaying AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeCannotStartRecording - An error code that indicates an attempt to start audio recording, but the operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartRecording
	AVAudioSessionErrorCodeCannotStartRecording AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeExpiredSession - An error code that indicates that an operation failed because the system deallocated the associated session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/expiredSession
	AVAudioSessionErrorCodeExpiredSession AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeIncompatibleCategory - An error code that indicates an attempt to perform an operation that the current audio session category doesn’t support.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/incompatibleCategory
	AVAudioSessionErrorCodeIncompatibleCategory AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeInsufficientPriority - An error code that indicates the app isn’t allowed to set the audio category because it’s in use by another app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/insufficientPriority
	AVAudioSessionErrorCodeInsufficientPriority AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/isBusy
	AVAudioSessionErrorCodeIsBusy AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeMediaServicesFailed - An error code that indictates an attempt to use the audio session during or after a Media Services failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/mediaServicesFailed
	AVAudioSessionErrorCodeMediaServicesFailed AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeMissingEntitlement - An error code that indicates an attempt to perform an operation for which the app doesn’t have the required entitlements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/missingEntitlement
	AVAudioSessionErrorCodeMissingEntitlement AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeNone - An error code that indicates the operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/none
	AVAudioSessionErrorCodeNone AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeResourceNotAvailable - An error code that indicates that an operation failed because the device doesn’t have sufficient hardware resources to complete the action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/resourceNotAvailable
	AVAudioSessionErrorCodeResourceNotAvailable AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeSessionNotActive - An error code that indicates the operation failed because the session isn’t active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/sessionNotActive
	AVAudioSessionErrorCodeSessionNotActive AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeSiriIsRecording - An error code that indicates an attempt to perform an operation that isn’t allowed while Siri is recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/siriIsRecording
	AVAudioSessionErrorCodeSiriIsRecording AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeUnspecified - An error code that indicates an unspecified error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/unspecified
	AVAudioSessionErrorCodeUnspecified AVAudioSessionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum AudioTimeStampFlags (7 cases) */
// AudioTimeStampFlags - A structure that represents flags for a timestamp.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags
type AudioTimeStampFlags uint

const (
	// kAudioTimeStampHostTimeValid - A flag that indicates that the host time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/hostTimeValid
	kAudioTimeStampHostTimeValid AudioTimeStampFlags = 0
	// kAudioTimeStampNothingValid - A flag that indicates no fields are valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/kAudioTimeStampNothingValid
	kAudioTimeStampNothingValid AudioTimeStampFlags = 0
	// kAudioTimeStampRateScalarValid - A flag that indicates that the rate scalar is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/rateScalarValid
	kAudioTimeStampRateScalarValid AudioTimeStampFlags = 0
	// kAudioTimeStampSampleHostTimeValid - A flag that indicates that the sample frame time and the host time are valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/sampleHostTimeValid
	kAudioTimeStampSampleHostTimeValid AudioTimeStampFlags = 0
	// kAudioTimeStampSampleTimeValid - A flag that indicates that the sample frame time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/sampleTimeValid
	kAudioTimeStampSampleTimeValid AudioTimeStampFlags = 0
	// kAudioTimeStampSMPTETimeValid - A flag that indicates that the SMPTE time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/smpteTimeValid
	kAudioTimeStampSMPTETimeValid AudioTimeStampFlags = 0
	// kAudioTimeStampWordClockTimeValid - A flag that indicates that the word clock time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/wordClockTimeValid
	kAudioTimeStampWordClockTimeValid AudioTimeStampFlags = 0
)

/* debug [enums.gen.go]: Processing enum MPEG4ObjectID (9 cases) */
// MPEG4ObjectID - Constants that define the type of MPEG-4 audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID
type MPEG4ObjectID uint

const (
	// kMPEG4Object_AAC_LC - A constant that specifies lossless coding, which provides compression with no loss of quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_LC
	kMPEG4Object_AAC_LC MPEG4ObjectID = 0
	// kMPEG4Object_AAC_LTP - A constant that specifies long-term prediction, which reduces redundancy in a coded signal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_LTP
	kMPEG4Object_AAC_LTP MPEG4ObjectID = 0
	// kMPEG4Object_AAC_Main - A constant that specifies advanced audio coding, which is the basic MPEG-4 technology.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/aac_Main
	kMPEG4Object_AAC_Main MPEG4ObjectID = 0
	// kMPEG4Object_AAC_SBR - A constant that specifies spectral band replication, which reconstructs high-frequency content from lower frequencies and side information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_SBR
	kMPEG4Object_AAC_SBR MPEG4ObjectID = 0
	// kMPEG4Object_AAC_Scalable - A constant that specifies scalable lossless coding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/aac_Scalable
	kMPEG4Object_AAC_Scalable MPEG4ObjectID = 0
	// kMPEG4Object_AAC_SSR - A constant that specifies scalable sampling rate, which provides different sampling frequencies for different targets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_SSR
	kMPEG4Object_AAC_SSR MPEG4ObjectID = 0
	// kMPEG4Object_CELP - A constant that specifies code-excited linear prediction, which is a narrow-band/wide-band speech codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/CELP
	kMPEG4Object_CELP MPEG4ObjectID = 0
	// kMPEG4Object_HVXC - A constant that specifies harmonic vector excitation coding, which is a very-low bit-rate parametric speech codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/HVXC
	kMPEG4Object_HVXC MPEG4ObjectID = 0
	// kMPEG4Object_TwinVQ - A constant that specifies transform-domain weighted interleaved vector quantization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/twinVQ
	kMPEG4Object_TwinVQ MPEG4ObjectID = 0
)

/* debug [enums.gen.go]: Processing enum SMPTETimeFlags (3 cases) */
// SMPTETimeFlags - A structure that defines SMPTE time flags.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags
type SMPTETimeFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/kSMPTETimeUnknown
	kSMPTETimeUnknown SMPTETimeFlags = 0
	// kSMPTETimeRunning - Time is running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/running
	kSMPTETimeRunning SMPTETimeFlags = 0
	// kSMPTETimeValid - The full time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/valid
	kSMPTETimeValid SMPTETimeFlags = 0
)

/* debug [enums.gen.go]: Processing enum SMPTETimeType (12 cases) */
// SMPTETimeType - Constants that define SMPTE time types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType
type SMPTETimeType uint

const (
	// kSMPTETimeType2398 - 23.98 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2398
	kSMPTETimeType2398 SMPTETimeType = 0
	// kSMPTETimeType24 - 24 video frames per second—standard for 16mm and 35mm film.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type24
	kSMPTETimeType24 SMPTETimeType = 0
	// kSMPTETimeType25 - 25 video frames per second—standard for PAL and SECAM video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type25
	kSMPTETimeType25 SMPTETimeType = 0
	// kSMPTETimeType2997 - 29.97 video frames per second—standard for NTSC video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2997
	kSMPTETimeType2997 SMPTETimeType = 0
	// kSMPTETimeType2997Drop - 29.97 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2997Drop
	kSMPTETimeType2997Drop SMPTETimeType = 0
	// kSMPTETimeType30 - 30 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type30
	kSMPTETimeType30 SMPTETimeType = 0
	// kSMPTETimeType30Drop - 30 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type30Drop
	kSMPTETimeType30Drop SMPTETimeType = 0
	// kSMPTETimeType50 - 50 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type50
	kSMPTETimeType50 SMPTETimeType = 0
	// kSMPTETimeType5994 - 59.94 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type5994
	kSMPTETimeType5994 SMPTETimeType = 0
	// kSMPTETimeType5994Drop - 59.94 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type5994Drop
	kSMPTETimeType5994Drop SMPTETimeType = 0
	// kSMPTETimeType60 - 60 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type60
	kSMPTETimeType60 SMPTETimeType = 0
	// kSMPTETimeType60Drop - 60 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type60Drop
	kSMPTETimeType60Drop SMPTETimeType = 0
)

/* debug [enums.gen.go]: Processing enum AudioChannelBitmap (27 cases) */
// AudioChannelBitmap - The supported channel bitmaps to use when defining channel layouts.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap
type AudioChannelBitmap uint

const (
	// kAudioChannelBit_Center - The center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_Center
	kAudioChannelBit_Center AudioChannelBitmap = 0
	// kAudioChannelBit_CenterSurround - The center surround channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_CenterSurround
	kAudioChannelBit_CenterSurround AudioChannelBitmap = 0
	// kAudioChannelBit_CenterTopFront - The top-front center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_CenterTopFront
	kAudioChannelBit_CenterTopFront AudioChannelBitmap = 0
	// kAudioChannelBit_CenterTopMiddle - The top-middle center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_CenterTopMiddle
	kAudioChannelBit_CenterTopMiddle AudioChannelBitmap = 0
	// kAudioChannelBit_CenterTopRear - The top-right center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_CenterTopRear
	kAudioChannelBit_CenterTopRear AudioChannelBitmap = 0
	// kAudioChannelBit_Left - The left channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_Left
	kAudioChannelBit_Left AudioChannelBitmap = 0
	// kAudioChannelBit_LeftCenter - The left center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftCenter
	kAudioChannelBit_LeftCenter AudioChannelBitmap = 0
	// kAudioChannelBit_LeftSurround - The left surround channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftSurround
	kAudioChannelBit_LeftSurround AudioChannelBitmap = 0
	// kAudioChannelBit_LeftSurroundDirect - The left surround direct channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftSurroundDirect
	kAudioChannelBit_LeftSurroundDirect AudioChannelBitmap = 0
	// kAudioChannelBit_LeftTopFront - The left-top front channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftTopFront
	kAudioChannelBit_LeftTopFront AudioChannelBitmap = 0
	// kAudioChannelBit_LeftTopMiddle - The left-top middle channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftTopMiddle
	kAudioChannelBit_LeftTopMiddle AudioChannelBitmap = 0
	// kAudioChannelBit_LeftTopRear - The left-top rear channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LeftTopRear
	kAudioChannelBit_LeftTopRear AudioChannelBitmap = 0
	// kAudioChannelBit_LFEScreen - The Low Frequency Effects (LFE) screen channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LFEScreen
	kAudioChannelBit_LFEScreen AudioChannelBitmap = 0
	// kAudioChannelBit_Right - The right channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_Right
	kAudioChannelBit_Right AudioChannelBitmap = 0
	// kAudioChannelBit_RightCenter - The right center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightCenter
	kAudioChannelBit_RightCenter AudioChannelBitmap = 0
	// kAudioChannelBit_RightSurround - The rIght surround channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightSurround
	kAudioChannelBit_RightSurround AudioChannelBitmap = 0
	// kAudioChannelBit_RightSurroundDirect - The right surround direct channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightSurroundDirect
	kAudioChannelBit_RightSurroundDirect AudioChannelBitmap = 0
	// kAudioChannelBit_RightTopFront - The top-front front channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightTopFront
	kAudioChannelBit_RightTopFront AudioChannelBitmap = 0
	// kAudioChannelBit_RightTopMiddle - The top-middle right channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightTopMiddle
	kAudioChannelBit_RightTopMiddle AudioChannelBitmap = 0
	// kAudioChannelBit_RightTopRear - The top-rear right channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_RightTopRear
	kAudioChannelBit_RightTopRear AudioChannelBitmap = 0
	// kAudioChannelBit_TopBackCenter - The top-back center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_TopBackCenter
	kAudioChannelBit_TopBackCenter AudioChannelBitmap = 0
	// kAudioChannelBit_TopBackLeft - The top-back left channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_TopBackLeft
	kAudioChannelBit_TopBackLeft AudioChannelBitmap = 0
	// kAudioChannelBit_TopBackRight - The top-back right channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_TopBackRight
	kAudioChannelBit_TopBackRight AudioChannelBitmap = 0
	// kAudioChannelBit_TopCenterSurround - The top center surround channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_TopCenterSurround
	kAudioChannelBit_TopCenterSurround AudioChannelBitmap = 0
	// kAudioChannelBit_VerticalHeightCenter - The vertical height center channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_VerticalHeightCenter
	kAudioChannelBit_VerticalHeightCenter AudioChannelBitmap = 0
	// kAudioChannelBit_VerticalHeightLeft - The vertical height left channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_VerticalHeightLeft
	kAudioChannelBit_VerticalHeightLeft AudioChannelBitmap = 0
	// kAudioChannelBit_VerticalHeightRight - The vertical height right channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_VerticalHeightRight
	kAudioChannelBit_VerticalHeightRight AudioChannelBitmap = 0
)

/* debug [enums.gen.go]: Processing enum AudioChannelCoordinateIndex (6 cases) */
// AudioChannelCoordinateIndex - Indexes the fields of the
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex
type AudioChannelCoordinateIndex uint

const (
	// kAudioChannelCoordinates_Azimuth - For spherical coordinates,   is front center, positive is right, negative is left, and measurements are in degrees.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_Azimuth
	kAudioChannelCoordinates_Azimuth AudioChannelCoordinateIndex = 0
	// kAudioChannelCoordinates_BackFront - For rectangular coordinates, negative is back and positive is front. The units are specified by the   field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_BackFront
	kAudioChannelCoordinates_BackFront AudioChannelCoordinateIndex = 0
	// kAudioChannelCoordinates_Distance - For spherical coordinates, distance is radially from the center. The units are specified by the   field of the   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_Distance
	kAudioChannelCoordinates_Distance AudioChannelCoordinateIndex = 0
	// kAudioChannelCoordinates_DownUp - For rectangular coordinates, negative is below ground level,   is ground level, and positive is above ground level. The units are specified by the   field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_DownUp
	kAudioChannelCoordinates_DownUp AudioChannelCoordinateIndex = 0
	// kAudioChannelCoordinates_Elevation - For spherical coordinates,   is zenith,   is horizontal,   is nadir, and measurements are in degrees.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_Elevation
	kAudioChannelCoordinates_Elevation AudioChannelCoordinateIndex = 0
	// kAudioChannelCoordinates_LeftRight - For rectangular coordinates, negative is left and positive is right. The units are specified by the   field of the   structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelCoordinateIndex/coordinates_LeftRight
	kAudioChannelCoordinates_LeftRight AudioChannelCoordinateIndex = 0
)

/* debug [enums.gen.go]: Processing enum AudioChannelFlags (4 cases) */
// AudioChannelFlags - Constants that define the audio channel flags of an audio channel description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags
type AudioChannelFlags uint

const (
	// kAudioChannelFlags_AllOff - All flags are clear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags/kAudioChannelFlags_AllOff
	kAudioChannelFlags_AllOff AudioChannelFlags = 0
	// kAudioChannelFlags_Meters - A flag that indicates that unit values are in meters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags/meters
	kAudioChannelFlags_Meters AudioChannelFlags = 0
	// kAudioChannelFlags_RectangularCoordinates - A flag that indicates the channel uses the speaker position’s cartesian coordinates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags/rectangularCoordinates
	kAudioChannelFlags_RectangularCoordinates AudioChannelFlags = 0
	// kAudioChannelFlags_SphericalCoordinates - A flag that indicates the channel uses the speaker position’s spherical coordinates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelFlags/sphericalCoordinates
	kAudioChannelFlags_SphericalCoordinates AudioChannelFlags = 0
)
