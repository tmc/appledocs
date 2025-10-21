// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes

// Enum types and constants
// AVAudioSessionErrorCode - Codes that describe error conditions that may occur when performing audio session operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode
type AudioSessionErrorCode uint

const (
	// AudioSessionErrorCodeBadParam - An error code that indicates an attempt to set a property to an illegal value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/badParam
	AudioSessionErrorCodeBadParam AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotInterruptOthers - An error code that indictates an attempt to make a nonmixable audio session active while the app was in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotInterruptOthers
	AudioSessionErrorCodeCannotInterruptOthers AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotStartPlaying - An error code that indicates an attempt to start audio playback when it wasn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartPlaying
	AudioSessionErrorCodeCannotStartPlaying AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotStartRecording - An error code that indicates an attempt to start audio recording, but the operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartRecording
	AudioSessionErrorCodeCannotStartRecording AudioSessionErrorCode = 0
	// AudioSessionErrorCodeExpiredSession - An error code that indicates that an operation failed because the system deallocated the associated session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/expiredSession
	AudioSessionErrorCodeExpiredSession AudioSessionErrorCode = 0
	// AudioSessionErrorCodeInsufficientPriority - An error code that indicates the app isn’t allowed to set the audio category because it’s in use by another app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/insufficientPriority
	AudioSessionErrorCodeInsufficientPriority AudioSessionErrorCode = 0
	// AudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/isBusy
	AudioSessionErrorCodeIsBusy AudioSessionErrorCode = 0
	// AudioSessionErrorCodeMediaServicesFailed - An error code that indictates an attempt to use the audio session during or after a Media Services failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/mediaServicesFailed
	AudioSessionErrorCodeMediaServicesFailed AudioSessionErrorCode = 0
	// AudioSessionErrorCodeMissingEntitlement - An error code that indicates an attempt to perform an operation for which the app doesn’t have the required entitlements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/missingEntitlement
	AudioSessionErrorCodeMissingEntitlement AudioSessionErrorCode = 0
	// AudioSessionErrorCodeNone - An error code that indicates the operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/none
	AudioSessionErrorCodeNone AudioSessionErrorCode = 0
	// AudioSessionErrorCodeSessionNotActive - An error code that indicates the operation failed because the session isn’t active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/sessionNotActive
	AudioSessionErrorCodeSessionNotActive AudioSessionErrorCode = 0
	// AudioSessionErrorCodeSiriIsRecording - An error code that indicates an attempt to perform an operation that isn’t allowed while Siri is recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/siriIsRecording
	AudioSessionErrorCodeSiriIsRecording AudioSessionErrorCode = 0
	// AudioSessionErrorCodeUnspecified - An error code that indicates an unspecified error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/unspecified
	AudioSessionErrorCodeUnspecified AudioSessionErrorCode = 0
	// AudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/avaudiosessionerrorcode/avaudiosessionerrorcodeisbusy
	AudioSessionErrorCodeIsBusy AudioSessionErrorCode = 0
)

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
	// kAudioChannelBit_LFEScreen - The Low Frequency Effects (LFE) screen channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelBitmap/bit_LFEScreen
	kAudioChannelBit_LFEScreen AudioChannelBitmap = 0
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

// MPEG4ObjectID - Constants that define the type of MPEG-4 audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID
type EG4ObjectID uint

const (
	// kMPEG4Object_AAC_LTP - A constant that specifies long-term prediction, which reduces redundancy in a coded signal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_LTP
	kMPEG4Object_AAC_LTP EG4ObjectID = 0
	// kMPEG4Object_AAC_SBR - A constant that specifies spectral band replication, which reconstructs high-frequency content from lower frequencies and side information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/AAC_SBR
	kMPEG4Object_AAC_SBR EG4ObjectID = 0
	// kMPEG4Object_AAC_Scalable - A constant that specifies scalable lossless coding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/aac_Scalable
	kMPEG4Object_AAC_Scalable EG4ObjectID = 0
)

// SMPTETimeFlags - A structure that defines SMPTE time flags.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags
type SMPTETimeFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/kSMPTETimeUnknown
	kSMPTETimeUnknown SMPTETimeFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/running
	kSMPTETimeRunning SMPTETimeFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeFlags/valid
	kSMPTETimeValid SMPTETimeFlags = 0
)

// SMPTETimeType - Constants that define SMPTE time types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType
type SMPTETimeType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2997
	kSMPTETimeType2997 SMPTETimeType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2997Drop
	kSMPTETimeType2997Drop SMPTETimeType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type30
	kSMPTETimeType30 SMPTETimeType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type30Drop
	kSMPTETimeType30Drop SMPTETimeType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type5994
	kSMPTETimeType5994 SMPTETimeType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type60Drop
	kSMPTETimeType60Drop SMPTETimeType = 0
)

// AVAudioSessionErrorCode - Codes that describe error conditions that may occur when performing audio session operations.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/errorcode
type AudioSessionErrorCode uint

const (
	// AudioSessionErrorCodeBadParam - An error code that indicates an attempt to set a property to an illegal value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/badParam
	AudioSessionErrorCodeBadParam AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotInterruptOthers - An error code that indictates an attempt to make a nonmixable audio session active while the app was in the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotInterruptOthers
	AudioSessionErrorCodeCannotInterruptOthers AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotStartPlaying - An error code that indicates an attempt to start audio playback when it wasn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartPlaying
	AudioSessionErrorCodeCannotStartPlaying AudioSessionErrorCode = 0
	// AudioSessionErrorCodeCannotStartRecording - An error code that indicates an attempt to start audio recording, but the operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/cannotStartRecording
	AudioSessionErrorCodeCannotStartRecording AudioSessionErrorCode = 0
	// AudioSessionErrorCodeExpiredSession - An error code that indicates that an operation failed because the system deallocated the associated session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/expiredSession
	AudioSessionErrorCodeExpiredSession AudioSessionErrorCode = 0
	// AudioSessionErrorCodeInsufficientPriority - An error code that indicates the app isn’t allowed to set the audio category because it’s in use by another app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/insufficientPriority
	AudioSessionErrorCodeInsufficientPriority AudioSessionErrorCode = 0
	// AudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/isBusy
	AudioSessionErrorCodeIsBusy AudioSessionErrorCode = 0
	// AudioSessionErrorCodeMediaServicesFailed - An error code that indictates an attempt to use the audio session during or after a Media Services failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/mediaServicesFailed
	AudioSessionErrorCodeMediaServicesFailed AudioSessionErrorCode = 0
	// AudioSessionErrorCodeMissingEntitlement - An error code that indicates an attempt to perform an operation for which the app doesn’t have the required entitlements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/missingEntitlement
	AudioSessionErrorCodeMissingEntitlement AudioSessionErrorCode = 0
	// AudioSessionErrorCodeNone - An error code that indicates the operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/none
	AudioSessionErrorCodeNone AudioSessionErrorCode = 0
	// AudioSessionErrorCodeSessionNotActive - An error code that indicates the operation failed because the session isn’t active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/sessionNotActive
	AudioSessionErrorCodeSessionNotActive AudioSessionErrorCode = 0
	// AudioSessionErrorCodeSiriIsRecording - An error code that indicates an attempt to perform an operation that isn’t allowed while Siri is recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/siriIsRecording
	AudioSessionErrorCodeSiriIsRecording AudioSessionErrorCode = 0
	// AudioSessionErrorCodeUnspecified - An error code that indicates an unspecified error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/unspecified
	AudioSessionErrorCodeUnspecified AudioSessionErrorCode = 0
	// AudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/coreaudiotypes/avaudiosessionerrorcode/avaudiosessionerrorcodeisbusy
	AudioSessionErrorCodeIsBusy AudioSessionErrorCode = 0
)


