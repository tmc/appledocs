// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes

// Enum types and constants
// AVAudioSessionErrorCode - Codes that describe error conditions that may occur when performing audio session operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode
type AVAudioSessionErrorCode uint

const (
	// AVAudioSessionErrorCodeBadParam - An error code that indicates an attempt to set a property to an illegal value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/badParam
	AVAudioSessionErrorCodeBadParam AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeExpiredSession - An error code that indicates that an operation failed because the system deallocated the associated session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/expiredSession
	AVAudioSessionErrorCodeExpiredSession AVAudioSessionErrorCode = 0
	// AVAudioSessionErrorCodeIsBusy - An error code that indicates an attempt to deactivate the audio session while it’s still playing or recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AVAudioSession/ErrorCode/isBusy
	AVAudioSessionErrorCodeIsBusy AVAudioSessionErrorCode = 0
)

// AudioTimeStampFlags - A structure that represents flags for a timestamp.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags
type AudioTimeStampFlags uint

const (
	// kAudioTimeStampSampleHostTimeValid - A flag that indicates that the sample frame time and the host time are valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStampFlags/sampleHostTimeValid
	kAudioTimeStampSampleHostTimeValid AudioTimeStampFlags = 0
)

// MPEG4ObjectID - Constants that define the type of MPEG-4 audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID
type MPEG4ObjectID uint

const (
	// kMPEG4Object_AAC_Scalable - A constant that specifies scalable lossless coding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/MPEG4ObjectID/aac_Scalable
	kMPEG4Object_AAC_Scalable MPEG4ObjectID = 0
)

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
)

// SMPTETimeType - Constants that define SMPTE time types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType
type SMPTETimeType uint

const (
	// kSMPTETimeType2997Drop - 29.97 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type2997Drop
	kSMPTETimeType2997Drop SMPTETimeType = 0
	// kSMPTETimeType30 - 30 video frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type30
	kSMPTETimeType30 SMPTETimeType = 0
	// kSMPTETimeType60Drop - 60 video frames per second, with video-frame numbers adjusted to ensure that the timecode matches elapsed clock time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETimeType/type60Drop
	kSMPTETimeType60Drop SMPTETimeType = 0
)


