// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

/* debug [enums.gen.go]: Generating 12 enums for AVKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum AVAudioSessionRouteSelection (3 cases) */
// AVAudioSessionRouteSelection - Constants that indicate the audio route selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection
type AVAudioSessionRouteSelection int

const (
	// AVAudioSessionRouteSelectionExternal - An external device selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionExternal
	AVAudioSessionRouteSelectionExternal AVAudioSessionRouteSelection = 0
	// AVAudioSessionRouteSelectionLocal - A local device selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionLocal
	AVAudioSessionRouteSelectionLocal AVAudioSessionRouteSelection = 0
	// AVAudioSessionRouteSelectionNone - No route selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionNone
	AVAudioSessionRouteSelectionNone AVAudioSessionRouteSelection = 0
)

/* debug [enums.gen.go]: Processing enum AVKitError (6 cases) */
// AVKitError - Constants that identify framework error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code
type AVKitError uint

const (
	// AVKitErrorContentDisallowedByPasscode - A restriction disallows access to this content, but the user can override the restriction by entering the device passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentDisallowedByPasscode
	AVKitErrorContentDisallowedByPasscode AVKitError = 0
	// AVKitErrorContentDisallowedByProfile - An installed profile restricts access to this content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentDisallowedByProfile
	AVKitErrorContentDisallowedByProfile AVKitError = 0
	// AVKitErrorContentRatingUnknown - The media content rating is missing or unrecognized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentRatingUnknown
	AVKitErrorContentRatingUnknown AVKitError = 0
	// AVKitErrorPictureInPictureStartFailed - The system failed to start Picture in Picture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/pictureInPictureStartFailed
	AVKitErrorPictureInPictureStartFailed AVKitError = 0
	// AVKitErrorUnknown - An unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/unknown
	AVKitErrorUnknown AVKitError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-c.enum/AVKitErrorRecordingFailed
	AVKitErrorRecordingFailed AVKitError = 0
)

/* debug [enums.gen.go]: Processing enum AVRoutePickerViewButtonState (4 cases) */
// AVRoutePickerViewButtonState - Constants that describe the available button states.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState
type AVRoutePickerViewButtonState uint

const (
	// AVRoutePickerViewButtonStateActive - The button state when AirPlay is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/active
	AVRoutePickerViewButtonStateActive AVRoutePickerViewButtonState = 0
	// AVRoutePickerViewButtonStateActiveHighlighted - The highlighted button state when AirPlay is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/activeHighlighted
	AVRoutePickerViewButtonStateActiveHighlighted AVRoutePickerViewButtonState = 0
	// AVRoutePickerViewButtonStateNormal - The normal, or default, button state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/normal
	AVRoutePickerViewButtonStateNormal AVRoutePickerViewButtonState = 0
	// AVRoutePickerViewButtonStateNormalHighlighted - The highlighted button state when a mouse-down event occurs inside the button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/normalHighlighted
	AVRoutePickerViewButtonStateNormalHighlighted AVRoutePickerViewButtonState = 0
)

/* debug [enums.gen.go]: Processing enum AVCaptureEventPhase (3 cases) */
// AVCaptureEventPhase - Constants that indicate the phase of a system capture event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase
type AVCaptureEventPhase uint

const (
	// AVCaptureEventPhaseBegan - A phase that indicates the beginning of a capture event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase/began
	AVCaptureEventPhaseBegan AVCaptureEventPhase = 0
	// AVCaptureEventPhaseCancelled - A phase that indicates the cancellation of a capture event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase/cancelled
	AVCaptureEventPhaseCancelled AVCaptureEventPhase = 0
	// AVCaptureEventPhaseEnded - A phase that indicates the end of a capture event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase/ended
	AVCaptureEventPhaseEnded AVCaptureEventPhase = 0
)

/* debug [enums.gen.go]: Processing enum AVCaptureViewControlsStyle (4 cases) */
// AVCaptureViewControlsStyle - Constants that describe the capture view’s supported controls styles.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle
type AVCaptureViewControlsStyle uint

const (
	// AVCaptureViewControlsStyleDefault - The view’s default controls style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/default
	AVCaptureViewControlsStyleDefault AVCaptureViewControlsStyle = 0
	// AVCaptureViewControlsStyleFloating - The view’s floating controls style, which matches the user interface of QuickTime Player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/floating
	AVCaptureViewControlsStyleFloating AVCaptureViewControlsStyle = 0
	// AVCaptureViewControlsStyleInline - The view’s inline controls style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/inline
	AVCaptureViewControlsStyleInline AVCaptureViewControlsStyle = 0
	// AVCaptureViewControlsStyleInlineDeviceSelection - The view’s inline device selection style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/inlineDeviceSelection
	AVCaptureViewControlsStyleInlineDeviceSelection AVCaptureViewControlsStyle = 0
)

/* debug [enums.gen.go]: Processing enum AVContentProposalAction (3 cases) */
// AVContentProposalAction - Constant that indicate the action a user takes when dismissing a content proposal.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction
type AVContentProposalAction uint

const (
	// AVContentProposalActionAccept - The user accepted the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/accept
	AVContentProposalActionAccept AVContentProposalAction = 0
	// AVContentProposalActionDefer - The user deferred the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/defer
	AVContentProposalActionDefer AVContentProposalAction = 0
	// AVContentProposalActionReject - The user rejected the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/reject
	AVContentProposalActionReject AVContentProposalAction = 0
)

/* debug [enums.gen.go]: Processing enum AVDisplayDynamicRange (4 cases) */
// AVDisplayDynamicRange - Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange
type AVDisplayDynamicRange uint

const (
	// AVDisplayDynamicRangeAutomatic - Defines an automatic dynamic range. Indicates that the dynamic range will be set automatically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange/automatic
	AVDisplayDynamicRangeAutomatic AVDisplayDynamicRange = 0
	// AVDisplayDynamicRangeConstrainedHigh - Defines a constrained high dynamic range. Allows for constrained High Dynamic Range (HDR) video content which is useful for mixing HDR and Standard Dynamic Range (SDR) content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange/constrainedHigh
	AVDisplayDynamicRangeConstrainedHigh AVDisplayDynamicRange = 0
	// AVDisplayDynamicRangeHigh - Defines a high dynamic range. Allows video content to use extended dynamic range if it has dynamic range content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange/high
	AVDisplayDynamicRangeHigh AVDisplayDynamicRange = 0
	// AVDisplayDynamicRangeStandard - Defines a standard dynamic range. Restricts the video content dynamic range to the standard range regardless of the actual range of the video content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange/standard
	AVDisplayDynamicRangeStandard AVDisplayDynamicRange = 0
)

/* debug [enums.gen.go]: Processing enum AVPlayerViewControllerSkippingBehavior (2 cases) */
// AVPlayerViewControllerSkippingBehavior - Constants that represent the player view controller’s skipping behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControllerSkippingBehavior
type AVPlayerViewControllerSkippingBehavior uint

const (
	// AVPlayerViewControllerSkippingBehaviorDefault - The default skipping behavior, which is to skip forward or backward in 10-second intervals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControllerSkippingBehavior/default
	AVPlayerViewControllerSkippingBehaviorDefault AVPlayerViewControllerSkippingBehavior = 0
	// AVPlayerViewControllerSkippingBehaviorSkipItem - Skipping behavior that specifies skipping to the next or previous item in the player’s playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControllerSkippingBehavior/skipItem
	AVPlayerViewControllerSkippingBehaviorSkipItem AVPlayerViewControllerSkippingBehavior = 0
)

/* debug [enums.gen.go]: Processing enum AVPlayerViewControlsStyle (5 cases) */
// AVPlayerViewControlsStyle - Constants that indicate which user interface controls the view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle
type AVPlayerViewControlsStyle uint

const (
	// AVPlayerViewControlsStyleDefault - The view’s default controls style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle/default
	AVPlayerViewControlsStyleDefault AVPlayerViewControlsStyle = 0
	// AVPlayerViewControlsStyleFloating - The view displays playback controls in a floating window over the video content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle/floating
	AVPlayerViewControlsStyleFloating AVPlayerViewControlsStyle = 0
	// AVPlayerViewControlsStyleInline - The view displays playback controls in a bar along the view’s bottom edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle/inline
	AVPlayerViewControlsStyleInline AVPlayerViewControlsStyle = 0
	// AVPlayerViewControlsStyleMinimal - The view presents basic controls to play and pause playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle/minimal
	AVPlayerViewControlsStyleMinimal AVPlayerViewControlsStyle = 0
	// AVPlayerViewControlsStyleNone - The view displays no playback controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle/none
	AVPlayerViewControlsStyleNone AVPlayerViewControlsStyle = 0
)

/* debug [enums.gen.go]: Processing enum AVPlayerViewTrimResult (2 cases) */
// AVPlayerViewTrimResult - Constants that specify an action a user takes when trimming media in a player view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewTrimResult
type AVPlayerViewTrimResult uint

const (
	// AVPlayerViewTrimCancelButton - The user clicked the Cancel button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewTrimResult/cancelButton
	AVPlayerViewTrimCancelButton AVPlayerViewTrimResult = 0
	// AVPlayerViewTrimOKButton - The user clicked the Trim button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewTrimResult/okButton
	AVPlayerViewTrimOKButton AVPlayerViewTrimResult = 0
)

/* debug [enums.gen.go]: Processing enum AVRoutePickerViewButtonStyle (3 cases) */
// AVRoutePickerViewButtonStyle - Constants that define the button styles a route picker view supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle
type AVRoutePickerViewButtonStyle uint

const (
	// AVRoutePickerViewButtonStyleCustom - A custom button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/custom
	AVRoutePickerViewButtonStyleCustom AVRoutePickerViewButtonStyle = 0
	// AVRoutePickerViewButtonStylePlain - A plain button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/plain
	AVRoutePickerViewButtonStylePlain AVRoutePickerViewButtonStyle = 0
	// AVRoutePickerViewButtonStyleSystem - A system-defined button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/system
	AVRoutePickerViewButtonStyleSystem AVRoutePickerViewButtonStyle = 0
)

/* debug [enums.gen.go]: Processing enum AVVideoFrameAnalysisType (6 cases) */
// AVVideoFrameAnalysisType - Constants that define the types of analysis a player view controller may perform on a paused video frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType
type AVVideoFrameAnalysisType uint

const (
	// AVVideoFrameAnalysisTypeNone - A type that performs no analysis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/AVVideoFrameAnalysisTypeNone
	AVVideoFrameAnalysisTypeNone AVVideoFrameAnalysisType = 0
	// AVVideoFrameAnalysisTypeDefault - The default types of analysis to perform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/default
	AVVideoFrameAnalysisTypeDefault AVVideoFrameAnalysisType = 0
	// AVVideoFrameAnalysisTypeMachineReadableCode - A type that recognizes machine-readable codes, such as QR codes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/machineReadableCode
	AVVideoFrameAnalysisTypeMachineReadableCode AVVideoFrameAnalysisType = 0
	// AVVideoFrameAnalysisTypeSubject - A type that finds a subject that a user can copy out of frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/subject
	AVVideoFrameAnalysisTypeSubject AVVideoFrameAnalysisType = 0
	// AVVideoFrameAnalysisTypeText - A type that finds text in a paused video frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/text
	AVVideoFrameAnalysisTypeText AVVideoFrameAnalysisType = 0
	// AVVideoFrameAnalysisTypeVisualSearch - A type that identifies objects, landmarks, art, and so on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType/visualSearch
	AVVideoFrameAnalysisTypeVisualSearch AVVideoFrameAnalysisType = 0
)


