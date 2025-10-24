// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

// Enum types and constants
// SCContentSharingPickerMode - Available modes for selecting streaming content from a picker presented by the operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode
type SCContentSharingPickerMode uint

const (
	// SCContentSharingPickerModeSingleApplication - The mode allowing the selection of a single application through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleApplication
	SCContentSharingPickerModeSingleApplication SCContentSharingPickerMode = 0
)

// SCFrameStatus - Status values for a frame from a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus
type SCFrameStatus uint

const (
	// SCFrameStatusComplete - A status that indicates the system successfully generated a new frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/complete
	SCFrameStatusComplete SCFrameStatus = 0
	// SCFrameStatusIdle - A status that indicates the system didn’t generate a new frame because the display didn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/idle
	SCFrameStatusIdle SCFrameStatus = 0
	// SCFrameStatusStopped - A status that indicates the frame is in a stopped state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/stopped
	SCFrameStatusStopped SCFrameStatus = 0
)

// SCScreenshotDisplayIntent - A value that specifies the type of display a screenshot rendering optimizes for.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum
type SCScreenshotDisplayIntent uint

const (
	// SCScreenshotDisplayIntentCanonical - Specifies that the screenshot renders with canonical display attributes optimizing output for presentation on a high dynamic range display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/canonical
	SCScreenshotDisplayIntentCanonical SCScreenshotDisplayIntent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/local
	SCScreenshotDisplayIntentLocal SCScreenshotDisplayIntent = 0
)

// SCScreenshotDynamicRange - Specifies the type of images returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum
type SCScreenshotDynamicRange uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/sdr
	SCScreenshotDynamicRangeSDR SCScreenshotDynamicRange = 0
)

// SCShareableContentStyle - The style of content presented in a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle
type SCShareableContentStyle uint

const (
	// SCShareableContentStyleApplication - The stream is currently presenting one or more applications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/application
	SCShareableContentStyleApplication SCShareableContentStyle = 0
	// SCShareableContentStyleDisplay - The stream is currently presenting a complete display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/display
	SCShareableContentStyleDisplay SCShareableContentStyle = 0
)

// SCStreamErrorCode - Codes for user cancellation events and errors that can occur in ScreenCaptureKit.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code
type SCStreamErrorCode uint

const (
	// SCStreamErrorFailedToStart - An error message that indicates a stream failed to start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedToStart
	SCStreamErrorFailedToStart SCStreamErrorCode = 0
	// SCStreamErrorRemovingStream - An error message that indicates a stream wasn’t removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/removingStream
	SCStreamErrorRemovingStream SCStreamErrorCode = 0
	// SCStreamErrorUserDeclined - An error message that indicates the user didn’t grant Screen Recording permission to your app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/userDeclined
	SCStreamErrorUserDeclined SCStreamErrorCode = 0
)

// SCStreamOutputType - Constants that represent output types for a stream frame.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType
type SCStreamOutputType uint

const (
	// SCStreamOutputTypeAudio - An output type that represents an audio capture sample buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType/audio
	SCStreamOutputTypeAudio SCStreamOutputType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType/microphone
	SCStreamOutputTypeMicrophone SCStreamOutputType = 0
)

// SCStreamType - The display type of the presented stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamType
type SCStreamType uint

const (
	// SCStreamTypeDisplay - The stream is currently on a complete display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamType/display
	SCStreamTypeDisplay SCStreamType = 0
	// SCStreamTypeWindow - The stream is currently presented as a window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamType/window
	SCStreamTypeWindow SCStreamType = 0
)


