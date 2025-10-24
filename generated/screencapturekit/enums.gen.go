// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

/* debug [enums.gen.go]: Generating 12 enums for ScreenCaptureKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SCContentSharingPickerMode (5 cases) */
// SCContentSharingPickerMode - Available modes for selecting streaming content from a picker presented by the operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode
type SCContentSharingPickerMode uint

const (
	// SCContentSharingPickerModeMultipleApplications - The mode allowing the selection of multiple applications through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/multipleApplications
	SCContentSharingPickerModeMultipleApplications SCContentSharingPickerMode = 0
	// SCContentSharingPickerModeMultipleWindows - The mode allowing the selection of multiple windows through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/multipleWindows
	SCContentSharingPickerModeMultipleWindows SCContentSharingPickerMode = 0
	// SCContentSharingPickerModeSingleApplication - The mode allowing the selection of a single application through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleApplication
	SCContentSharingPickerModeSingleApplication SCContentSharingPickerMode = 0
	// SCContentSharingPickerModeSingleDisplay - The mode allowing the selection of a single display through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleDisplay
	SCContentSharingPickerModeSingleDisplay SCContentSharingPickerMode = 0
	// SCContentSharingPickerModeSingleWindow - The mode allowing the selection of a single window through the presented picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode/singleWindow
	SCContentSharingPickerModeSingleWindow SCContentSharingPickerMode = 0
)

/* debug [enums.gen.go]: Processing enum SCFrameStatus (6 cases) */
// SCFrameStatus - Status values for a frame from a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus
type SCFrameStatus uint

const (
	// SCFrameStatusBlank - A status that indicates the system didn’t generate a new frame because the display is blank.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/blank
	SCFrameStatusBlank SCFrameStatus = 0
	// SCFrameStatusComplete - A status that indicates the system successfully generated a new frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/complete
	SCFrameStatusComplete SCFrameStatus = 0
	// SCFrameStatusIdle - A status that indicates the system didn’t generate a new frame because the display didn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/idle
	SCFrameStatusIdle SCFrameStatus = 0
	// SCFrameStatusStarted - A status that indicates the frame is the first one sent after the stream starts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/started
	SCFrameStatusStarted SCFrameStatus = 0
	// SCFrameStatusStopped - A status that indicates the frame is in a stopped state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/stopped
	SCFrameStatusStopped SCFrameStatus = 0
	// SCFrameStatusSuspended - A status that indicates the system didn’t generate a new frame because you suspended updates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus/suspended
	SCFrameStatusSuspended SCFrameStatus = 0
)

/* debug [enums.gen.go]: Processing enum SCScreenshotDisplayIntent (2 cases) */
// SCScreenshotDisplayIntent - A value that specifies the type of display a screenshot rendering optimizes for.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum
type SCScreenshotDisplayIntent uint

const (
	// SCScreenshotDisplayIntentCanonical - Specifies that the screenshot renders with canonical display attributes optimizing output for presentation on a high dynamic range display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/canonical
	SCScreenshotDisplayIntentCanonical SCScreenshotDisplayIntent = 0
	// SCScreenshotDisplayIntentLocal - Specifies that the screenshot renders with local display attributes optimizing output for presentation on the capture display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum/local
	SCScreenshotDisplayIntentLocal SCScreenshotDisplayIntent = 0
)

/* debug [enums.gen.go]: Processing enum SCScreenshotDynamicRange (3 cases) */
// SCScreenshotDynamicRange - Specifies the type of images returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum
type SCScreenshotDynamicRange uint

const (
	// SCScreenshotDynamicRangeSDRAndHDR - Returns both standard dynamic range and high dynamic range image versions to the client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/bothSDRAndHDR
	SCScreenshotDynamicRangeSDRAndHDR SCScreenshotDynamicRange = 0
	// SCScreenshotDynamicRangeHDR - Returns a high dynamic range image to the client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/hdr
	SCScreenshotDynamicRangeHDR SCScreenshotDynamicRange = 0
	// SCScreenshotDynamicRangeSDR - Returns a standard dynamic range image to the client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum/sdr
	SCScreenshotDynamicRangeSDR SCScreenshotDynamicRange = 0
)

/* debug [enums.gen.go]: Processing enum SCShareableContentStyle (4 cases) */
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
	// SCShareableContentStyleNone - The stream isn’t currently presenting any content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/none
	SCShareableContentStyleNone SCShareableContentStyle = 0
	// SCShareableContentStyleWindow - The stream is currently presenting one or more windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle/window
	SCShareableContentStyleWindow SCShareableContentStyle = 0
)

/* debug [enums.gen.go]: Processing enum SCStreamConfigurationPreset (5 cases) */
// SCStreamConfigurationPreset enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset
type SCStreamConfigurationPreset uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset/captureHDRRecordingPreservedSDRHDR10
	SCStreamConfigurationPresetCaptureHDRRecordingPreservedSDRHDR10 SCStreamConfigurationPreset = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset/captureHDRScreenshotCanonicalDisplay
	SCStreamConfigurationPresetCaptureHDRScreenshotCanonicalDisplay SCStreamConfigurationPreset = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset/captureHDRScreenshotLocalDisplay
	SCStreamConfigurationPresetCaptureHDRScreenshotLocalDisplay SCStreamConfigurationPreset = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset/captureHDRStreamCanonicalDisplay
	SCStreamConfigurationPresetCaptureHDRStreamCanonicalDisplay SCStreamConfigurationPreset = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset/captureHDRStreamLocalDisplay
	SCStreamConfigurationPresetCaptureHDRStreamLocalDisplay SCStreamConfigurationPreset = 0
)

/* debug [enums.gen.go]: Processing enum SCStreamErrorCode (21 cases) */
// SCStreamErrorCode - Codes for user cancellation events and errors that can occur in ScreenCaptureKit.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code
type SCStreamErrorCode uint

const (
	// SCStreamErrorAttemptToConfigState - An error message that indicates a stream couldn’t update its configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/attemptToConfigState
	SCStreamErrorAttemptToConfigState SCStreamErrorCode = 0
	// SCStreamErrorAttemptToStartStreamState - An error message that indicates a stream is already running or doesn’t exist when trying to start a stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/attemptToStartStreamState
	SCStreamErrorAttemptToStartStreamState SCStreamErrorCode = 0
	// SCStreamErrorAttemptToStopStreamState - An error message that indicates a stream is already stopped or doesn’t exist when trying to stop a stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/attemptToStopStreamState
	SCStreamErrorAttemptToStopStreamState SCStreamErrorCode = 0
	// SCStreamErrorAttemptToUpdateFilterState - An error message that indicates a stream couldn’t update its content filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/attemptToUpdateFilterState
	SCStreamErrorAttemptToUpdateFilterState SCStreamErrorCode = 0
	// SCStreamErrorFailedApplicationConnectionInterrupted - An error message that indicates there was an interruption in a connection to an app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedApplicationConnectionInterrupted
	SCStreamErrorFailedApplicationConnectionInterrupted SCStreamErrorCode = 0
	// SCStreamErrorFailedApplicationConnectionInvalid - An error message that indicates the stream lost its connection to an app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedApplicationConnectionInvalid
	SCStreamErrorFailedApplicationConnectionInvalid SCStreamErrorCode = 0
	// SCStreamErrorFailedNoMatchingApplicationContext - An error message that indicates there isn’t a matching app context for streaming.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedNoMatchingApplicationContext
	SCStreamErrorFailedNoMatchingApplicationContext SCStreamErrorCode = 0
	// SCStreamErrorFailedToStart - An error message that indicates a stream failed to start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedToStart
	SCStreamErrorFailedToStart SCStreamErrorCode = 0
	// SCStreamErrorFailedToStartAudioCapture - An error message that indicates an audio stream failed to start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedToStartAudioCapture
	SCStreamErrorFailedToStartAudioCapture SCStreamErrorCode = 0
	// SCStreamErrorFailedToStartMicrophoneCapture - An error message that indicates microphone capture failed to start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedToStartMicrophoneCapture
	SCStreamErrorFailedToStartMicrophoneCapture SCStreamErrorCode = 0
	// SCStreamErrorFailedToStopAudioCapture - An error message that indicates an audio stream failed to stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/failedToStopAudioCapture
	SCStreamErrorFailedToStopAudioCapture SCStreamErrorCode = 0
	// SCStreamErrorInternalError - An error message that indicates a stream can’t start due to a failure in ScreenCaptureKit’s internals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/internalError
	SCStreamErrorInternalError SCStreamErrorCode = 0
	// SCStreamErrorInvalidParameter - An error message that indicates an operation failed because of an invalid parameter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/invalidParameter
	SCStreamErrorInvalidParameter SCStreamErrorCode = 0
	// SCStreamErrorMissingEntitlements - An error message that indicates missing entitlements in your app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/missingEntitlements
	SCStreamErrorMissingEntitlements SCStreamErrorCode = 0
	// SCStreamErrorNoCaptureSource - An error message that indicates a stream doesn’t have a source to capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/noCaptureSource
	SCStreamErrorNoCaptureSource SCStreamErrorCode = 0
	// SCStreamErrorNoDisplayList - An error message that indicates a stream doesn’t have displays available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/noDisplayList
	SCStreamErrorNoDisplayList SCStreamErrorCode = 0
	// SCStreamErrorNoWindowList - An error message that indicates a stream doesn’t have windows available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/noWindowList
	SCStreamErrorNoWindowList SCStreamErrorCode = 0
	// SCStreamErrorRemovingStream - An error message that indicates a stream wasn’t removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/removingStream
	SCStreamErrorRemovingStream SCStreamErrorCode = 0
	// SCStreamErrorSystemStoppedStream - An error message that indicates the system stopped the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/systemStoppedStream
	SCStreamErrorSystemStoppedStream SCStreamErrorCode = 0
	// SCStreamErrorUserDeclined - An error message that indicates the user didn’t grant Screen Recording permission to your app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/userDeclined
	SCStreamErrorUserDeclined SCStreamErrorCode = 0
	// SCStreamErrorUserStopped - An error message that indicates the user stopped the stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code/userStopped
	SCStreamErrorUserStopped SCStreamErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SCStreamOutputType (3 cases) */
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
	// SCStreamOutputTypeScreen - An output type that represents a screen capture sample buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType/screen
	SCStreamOutputTypeScreen SCStreamOutputType = 0
)

/* debug [enums.gen.go]: Processing enum SCCaptureDynamicRange (3 cases) */
// SCCaptureDynamicRange - Specifies whether the captured screen output is standard or high dynamic range.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureDynamicRange
type SCCaptureDynamicRange uint

const (
	// SCCaptureDynamicRangeHDRCanonicalDisplay - Specifies that the system captures the screen in high dynamic range with attributes of the canonical display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureDynamicRange/hdrCanonicalDisplay
	SCCaptureDynamicRangeHDRCanonicalDisplay SCCaptureDynamicRange = 0
	// SCCaptureDynamicRangeHDRLocalDisplay - Specifies that the system captures the screen in high dynamic range with attributes of the local display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureDynamicRange/hdrLocalDisplay
	SCCaptureDynamicRangeHDRLocalDisplay SCCaptureDynamicRange = 0
	// SCCaptureDynamicRangeSDR - Specifies that the system captures the screen in standard dynamic range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureDynamicRange/SDR
	SCCaptureDynamicRangeSDR SCCaptureDynamicRange = 0
)

/* debug [enums.gen.go]: Processing enum SCCaptureResolutionType (3 cases) */
// SCCaptureResolutionType - Available resolutions for content capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureResolutionType
type SCCaptureResolutionType uint

const (
	// SCCaptureResolutionAutomatic - Allow ScreenCaptureKit to automatically select the quality of content depending on factors such as network connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureResolutionType/automatic
	SCCaptureResolutionAutomatic SCCaptureResolutionType = 0
	// SCCaptureResolutionBest - Capture streaming content at the best available resolution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureResolutionType/best
	SCCaptureResolutionBest SCCaptureResolutionType = 0
	// SCCaptureResolutionNominal - Capture streaming content with a one point to one pixel conversion factor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureResolutionType/nominal
	SCCaptureResolutionNominal SCCaptureResolutionType = 0
)

/* debug [enums.gen.go]: Processing enum SCPresenterOverlayAlertSetting (3 cases) */
// SCPresenterOverlayAlertSetting - Configures how to present streaming notifications to a streamer of Presenter Overlay.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCPresenterOverlayAlertSetting
type SCPresenterOverlayAlertSetting uint

const (
	// SCPresenterOverlayAlertSettingAlways - Always display an alert when using Presenter Overlay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCPresenterOverlayAlertSetting/always
	SCPresenterOverlayAlertSettingAlways SCPresenterOverlayAlertSetting = 0
	// SCPresenterOverlayAlertSettingNever - Never display an alert when using Presenter Overlay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCPresenterOverlayAlertSetting/never
	SCPresenterOverlayAlertSettingNever SCPresenterOverlayAlertSetting = 0
	// SCPresenterOverlayAlertSettingSystem - Displays an alert when using Presenter Overlay based on the System Settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCPresenterOverlayAlertSetting/system
	SCPresenterOverlayAlertSettingSystem SCPresenterOverlayAlertSetting = 0
)

/* debug [enums.gen.go]: Processing enum SCStreamType (2 cases) */
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


