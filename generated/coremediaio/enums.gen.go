// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio


// Enum types and constants

// CMIOExtensionStreamClockType - Constants that indicate the clock type of a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/ClockType-swift.enum
type CMIOExtensionStreamClockType uint

const (
	// CMIOExtensionStreamClockTypeCustom - Indicates that the stream’s clock is specific to the device hosting the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/ClockType-swift.enum/custom
	CMIOExtensionStreamClockTypeCustom CMIOExtensionStreamClockType = 0
	// CMIOExtensionStreamClockTypeHostTime - Indicates that the stream uses the host time clock.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/ClockType-swift.enum/hostTime
	CMIOExtensionStreamClockTypeHostTime CMIOExtensionStreamClockType = 0
	// CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID - Indicates that the stream uses the clock of the linked Core Audio device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/ClockType-swift.enum/linkedCoreAudioDeviceUID
	CMIOExtensionStreamClockTypeLinkedCoreAudioDeviceUID CMIOExtensionStreamClockType = 0
)


// CMIOExtensionStreamDirection - Constants that define the data-flow direction of the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/Direction-swift.enum
type CMIOExtensionStreamDirection uint

const (
	// CMIOExtensionStreamDirectionSink - A stream that consumes sample buffers for playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/Direction-swift.enum/sink
	CMIOExtensionStreamDirectionSink CMIOExtensionStreamDirection = 0
	// CMIOExtensionStreamDirectionSource - A stream that provides sample buffers for capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/Direction-swift.enum/source
	CMIOExtensionStreamDirectionSource CMIOExtensionStreamDirection = 0
)


// CMIOExtensionStreamDiscontinuityFlags - Constants that specify the types of discontinuities that can occur in a media stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/DiscontinuityFlags
type CMIOExtensionStreamDiscontinuityFlags uint

const (
	// CMIOExtensionStreamDiscontinuityFlagSampleDropped - A flag that indicates a discontinuity in the stream due to a dropped frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/DiscontinuityFlags/sampleDropped
	CMIOExtensionStreamDiscontinuityFlagSampleDropped CMIOExtensionStreamDiscontinuityFlags = 0
	// CMIOExtensionStreamDiscontinuityFlagTime - A flag that indicates a time discontinuity in the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/DiscontinuityFlags/time
	CMIOExtensionStreamDiscontinuityFlagTime CMIOExtensionStreamDiscontinuityFlags = 0
	// CMIOExtensionStreamDiscontinuityFlagUnknown - A flag that indicates a stream discontinuity due to an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStream/DiscontinuityFlags/unknown
	CMIOExtensionStreamDiscontinuityFlagUnknown CMIOExtensionStreamDiscontinuityFlags = 0
	// CMIOExtensionStreamDiscontinuityFlagNone - A flag that indicates there’s no discontinuity in the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStreamDiscontinuityFlags/CMIOExtensionStreamDiscontinuityFlagNone
	CMIOExtensionStreamDiscontinuityFlagNone CMIOExtensionStreamDiscontinuityFlags = 0
)


