// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

/* debug [enums.gen.go]: Generating 7 enums for CoreVideo */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CVAttachmentMode (2 cases) */
// AttachmentMode - The propagation modes of a Core Video buffer attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVAttachmentMode
type AttachmentMode uint

const (
	// kCVAttachmentMode_ShouldNotPropagate - Indicates to not propagate the attachment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVAttachmentMode/shouldNotPropagate
	kCVAttachmentMode_ShouldNotPropagate AttachmentMode = 0
	// kCVAttachmentMode_ShouldPropagate - Indicates to copy the attachment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVAttachmentMode/shouldPropagate
	kCVAttachmentMode_ShouldPropagate AttachmentMode = 0
)

/* debug [enums.gen.go]: Processing enum CVPixelBufferLockFlags (1 cases) */
// PixelBufferLockFlags - The flags to pass to 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockFlags
type PixelBufferLockFlags uint

const (
	// kCVPixelBufferLock_ReadOnly - A read-only buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockFlags/readOnly
	kCVPixelBufferLock_ReadOnly PixelBufferLockFlags = 0
)

/* debug [enums.gen.go]: Processing enum CVPixelBufferPoolFlushFlags (1 cases) */
// PixelBufferPoolFlushFlags - The flags to pass to flush the pool.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlushFlags
type PixelBufferPoolFlushFlags uint

const (
	// kCVPixelBufferPoolFlushExcessBuffers - The value to pass to flush all unused buffers regardless of age.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlushFlags/excessBuffers
	kCVPixelBufferPoolFlushExcessBuffers PixelBufferPoolFlushFlags = 0
)

/* debug [enums.gen.go]: Processing enum CVSMPTETimeFlags (2 cases) */
// SMPTETimeFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeFlags
type SMPTETimeFlags uint

const (
	// kCVSMPTETimeRunning - Time is running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeFlags/running
	kCVSMPTETimeRunning SMPTETimeFlags = 0
	// kCVSMPTETimeValid - The full time is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeFlags/valid
	kCVSMPTETimeValid SMPTETimeFlags = 0
)

/* debug [enums.gen.go]: Processing enum CVSMPTETimeType (8 cases) */
// SMPTETimeType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType
type SMPTETimeType uint

const (
	// kCVSMPTETimeType24 - 24 frames per second (standard film).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type24
	kCVSMPTETimeType24 SMPTETimeType = 0
	// kCVSMPTETimeType25 - 25 frames per second (standard PAL).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type25
	kCVSMPTETimeType25 SMPTETimeType = 0
	// kCVSMPTETimeType2997 - 29.97 frames per second (standard NTSC).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type2997
	kCVSMPTETimeType2997 SMPTETimeType = 0
	// kCVSMPTETimeType2997Drop - 29.97 drop frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type2997Drop
	kCVSMPTETimeType2997Drop SMPTETimeType = 0
	// kCVSMPTETimeType30 - 30 frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type30
	kCVSMPTETimeType30 SMPTETimeType = 0
	// kCVSMPTETimeType30Drop - 30 drop frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type30Drop
	kCVSMPTETimeType30Drop SMPTETimeType = 0
	// kCVSMPTETimeType5994 - 59.94 frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type5994
	kCVSMPTETimeType5994 SMPTETimeType = 0
	// kCVSMPTETimeType60 - 60 frames per second.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type60
	kCVSMPTETimeType60 SMPTETimeType = 0
)

/* debug [enums.gen.go]: Processing enum CVTimeFlags (1 cases) */
// TimeFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeFlags
type TimeFlags uint

const (
	// kCVTimeIsIndefinite - The time value is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeFlags/isIndefinite
	kCVTimeIsIndefinite TimeFlags = 0
)

/* debug [enums.gen.go]: Processing enum CVTimeStampFlags (9 cases) */
// TimeStampFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags
type TimeStampFlags uint

const (
	// kCVTimeStampBottomField - The timestamp represents the bottom lines of an interlaced image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/bottomField
	kCVTimeStampBottomField TimeStampFlags = 0
	// kCVTimeStampHostTimeValid - The value in the host time field is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/hostTimeValid
	kCVTimeStampHostTimeValid TimeStampFlags = 0
	// kCVTimeStampIsInterlaced - A convenience constant indicating that the timestamp is for an interlaced image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/isInterlaced
	kCVTimeStampIsInterlaced TimeStampFlags = 0
	// kCVTimeStampRateScalarValid - The value in the rate scalar field is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/rateScalarValid
	kCVTimeStampRateScalarValid TimeStampFlags = 0
	// kCVTimeStampSMPTETimeValid - The value in the SMPTE time field is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/smpteTimeValid
	kCVTimeStampSMPTETimeValid TimeStampFlags = 0
	// kCVTimeStampTopField - The timestamp represents the top lines of an interlaced image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/topField
	kCVTimeStampTopField TimeStampFlags = 0
	// kCVTimeStampVideoHostTimeValid - A convenience constant indicating that both the video time and host time fields are valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/videoHostTimeValid
	kCVTimeStampVideoHostTimeValid TimeStampFlags = 0
	// kCVTimeStampVideoRefreshPeriodValid - The value in the video refresh period field is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/videoRefreshPeriodValid
	kCVTimeStampVideoRefreshPeriodValid TimeStampFlags = 0
	// kCVTimeStampVideoTimeValid - The value in the video time field is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/videoTimeValid
	kCVTimeStampVideoTimeValid TimeStampFlags = 0
)


