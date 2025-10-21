// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

// Enum types and constants
// SCContentSharingPickerMode - Available modes for selecting streaming content from a picker presented by the operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode
type ContentSharingPickerMode uint

const (
	// ContentSharingPickerModeSingleDisplay - The mode allowing the selection of a single display through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleDisplay
	ContentSharingPickerModeSingleDisplay ContentSharingPickerMode = 0
	// ContentSharingPickerModeSingleWindow - The mode allowing the selection of a single window through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleWindow
	ContentSharingPickerModeSingleWindow ContentSharingPickerMode = 0
)

// SCFrameStatus - Status values for a frame from a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus
type FrameStatus uint

const (
	// FrameStatusIdle - A status that indicates the system didn’t generate a new frame because the display didn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/idle
	FrameStatusIdle FrameStatus = 0
)

// SCScreenshotDisplayIntent enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum
type ScreenshotDisplayIntent uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/canonical
	ScreenshotDisplayIntentCanonical ScreenshotDisplayIntent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/local
	ScreenshotDisplayIntentLocal ScreenshotDisplayIntent = 0
)

// SCScreenshotDynamicRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum
type ScreenshotDynamicRange uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/hdr
	ScreenshotDynamicRangeHDR ScreenshotDynamicRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/sdr
	ScreenshotDynamicRangeSDR ScreenshotDynamicRange = 0
)

// SCShareableContentStyle - The style of content presented in a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle
type ShareableContentStyle uint

const (
	// ShareableContentStyleApplication - The stream is currently presenting one or more applications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/application
	ShareableContentStyleApplication ShareableContentStyle = 0
	// ShareableContentStyleWindow - The stream is currently presenting one or more windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/window
	ShareableContentStyleWindow ShareableContentStyle = 0
)

// SCStreamErrorCode - Codes for user cancellation events and errors that can occur in ScreenCaptureKit.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code
type StreamErrorCode uint

const (
	// StreamErrorInternalError - An error message that indicates a stream can’t start due to a failure in ScreenCaptureKit’s internals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/internalError
	StreamErrorInternalError StreamErrorCode = 0
	// StreamErrorRemovingStream - An error message that indicates a stream wasn’t removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/removingStream
	StreamErrorRemovingStream StreamErrorCode = 0
)

// SCStreamOutputType - Constants that represent output types for a stream frame.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType
type StreamOutputType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType/microphone
	StreamOutputTypeMicrophone StreamOutputType = 0
	// StreamOutputTypeScreen - An output type that represents a screen capture sample buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType/screen
	StreamOutputTypeScreen StreamOutputType = 0
)


