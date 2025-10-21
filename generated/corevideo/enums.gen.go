// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

// Enum types and constants
// CVAttachmentMode - The propagation modes of a Core Video buffer attachment.
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

// CVPixelBufferLockFlags - The flags to pass to 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockFlags
type PixelBufferLockFlags uint

const (
// kCVPixelBufferLock_ReadOnly - A read-only buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockFlags/readOnly
kCVPixelBufferLock_ReadOnly PixelBufferLockFlags = 0
)

// CVPixelBufferPoolFlushFlags - The flags to pass to flush the pool.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlushFlags
type PixelBufferPoolFlushFlags uint

const (
// kCVPixelBufferPoolFlushExcessBuffers - The value to pass to flush all unused buffers regardless of age.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlushFlags/excessBuffers
kCVPixelBufferPoolFlushExcessBuffers PixelBufferPoolFlushFlags = 0
)

// CVSMPTETimeFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeFlags
type SMPTETimeFlags uint

// CVSMPTETimeType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType
type SMPTETimeType uint

const (
// kCVSMPTETimeType24 - 24 frames per second (standard film).
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETimeType/type24
kCVSMPTETimeType24 SMPTETimeType = 0
)

// CVTimeFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeFlags
type TimeFlags uint

// CVTimeStampFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags
type TimeStampFlags uint

const (
// kCVTimeStampHostTimeValid - The value in the host time field is valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/hostTimeValid
kCVTimeStampHostTimeValid TimeStampFlags = 0
// kCVTimeStampVideoRefreshPeriodValid - The value in the video refresh period field is valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/videoRefreshPeriodValid
kCVTimeStampVideoRefreshPeriodValid TimeStampFlags = 0
// kCVTimeStampVideoTimeValid - The value in the video time field is valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStampFlags/videoTimeValid
kCVTimeStampVideoTimeValid TimeStampFlags = 0
)


