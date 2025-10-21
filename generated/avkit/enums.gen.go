// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

// Enum types and constants
// AVAudioSessionRouteSelection - Constants that indicate the audio route selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection
type AudioSessionRouteSelection uint

const (
	// AudioSessionRouteSelectionExternal - An external device selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionExternal
	AudioSessionRouteSelectionExternal AudioSessionRouteSelection = 0
	// AudioSessionRouteSelectionLocal - A local device selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionLocal
	AudioSessionRouteSelectionLocal AudioSessionRouteSelection = 0
	// AudioSessionRouteSelectionNone - No route selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVAudioSessionRouteSelection/AVAudioSessionRouteSelectionNone
	AudioSessionRouteSelectionNone AudioSessionRouteSelection = 0
)

// AVCaptureEventPhase - Constants that indicate the phase of a system capture event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase
type CaptureEventPhase uint

// AVCaptureViewControlsStyle - Constants that describe the capture view’s supported controls styles.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle
type CaptureViewControlsStyle uint

const (
	// CaptureViewControlsStyleDefault - The view’s default controls style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/default
	CaptureViewControlsStyleDefault CaptureViewControlsStyle = 0
	// CaptureViewControlsStyleFloating - The view’s floating controls style, which matches the user interface of QuickTime Player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/floating
	CaptureViewControlsStyleFloating CaptureViewControlsStyle = 0
	// CaptureViewControlsStyleInline - The view’s inline controls style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/inline
	CaptureViewControlsStyleInline CaptureViewControlsStyle = 0
	// CaptureViewControlsStyleInlineDeviceSelection - The view’s inline device selection style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle/inlineDeviceSelection
	CaptureViewControlsStyleInlineDeviceSelection CaptureViewControlsStyle = 0
)

// AVContentProposalAction - Constant that indicate the action a user takes when dismissing a content proposal.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction
type ContentProposalAction uint

const (
	// ContentProposalActionAccept - The user accepted the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/accept
	ContentProposalActionAccept ContentProposalAction = 0
	// ContentProposalActionDefer - The user deferred the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/defer
	ContentProposalActionDefer ContentProposalAction = 0
	// ContentProposalActionReject - The user rejected the content proposal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction/reject
	ContentProposalActionReject ContentProposalAction = 0
)

// AVDisplayDynamicRange - Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayDynamicRange
type DisplayDynamicRange uint

// AVKitError - Constants that identify framework error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code
type KitError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-c.enum/AVKitErrorRecordingFailed
	KitErrorRecordingFailed KitError = 0
	// KitErrorContentDisallowedByPasscode - A restriction disallows access to this content, but the user can override the restriction by entering the device passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentDisallowedByPasscode
	KitErrorContentDisallowedByPasscode KitError = 0
	// KitErrorContentDisallowedByProfile - An installed profile restricts access to this content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentDisallowedByProfile
	KitErrorContentDisallowedByProfile KitError = 0
	// KitErrorContentRatingUnknown - The media content rating is missing or unrecognized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/contentRatingUnknown
	KitErrorContentRatingUnknown KitError = 0
	// KitErrorPictureInPictureStartFailed - The system failed to start Picture in Picture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/pictureInPictureStartFailed
	KitErrorPictureInPictureStartFailed KitError = 0
	// KitErrorUnknown - An unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code/unknown
	KitErrorUnknown KitError = 0
)

// AVPlayerViewControllerSkippingBehavior - Constants that represent the player view controller’s skipping behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControllerSkippingBehavior
type PlayerViewControllerSkippingBehavior uint

// AVPlayerViewControlsStyle - Constants that indicate which user interface controls the view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle
type PlayerViewControlsStyle uint

// AVPlayerViewTrimResult - Constants that specify an action a user takes when trimming media in a player view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewTrimResult
type PlayerViewTrimResult uint

// AVRoutePickerViewButtonState - Constants that describe the available button states.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState
type RoutePickerViewButtonState uint

const (
	// RoutePickerViewButtonStateActive - The button state when AirPlay is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/active
	RoutePickerViewButtonStateActive RoutePickerViewButtonState = 0
	// RoutePickerViewButtonStateActiveHighlighted - The highlighted button state when AirPlay is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/activeHighlighted
	RoutePickerViewButtonStateActiveHighlighted RoutePickerViewButtonState = 0
	// RoutePickerViewButtonStateNormal - The normal, or default, button state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/normal
	RoutePickerViewButtonStateNormal RoutePickerViewButtonState = 0
	// RoutePickerViewButtonStateNormalHighlighted - The highlighted button state when a mouse-down event occurs inside the button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState/normalHighlighted
	RoutePickerViewButtonStateNormalHighlighted RoutePickerViewButtonState = 0
)

// AVRoutePickerViewButtonStyle - Constants that define the button styles a route picker view supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle
type RoutePickerViewButtonStyle uint

const (
	// RoutePickerViewButtonStyleCustom - A custom button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/custom
	RoutePickerViewButtonStyleCustom RoutePickerViewButtonStyle = 0
	// RoutePickerViewButtonStylePlain - A plain button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/plain
	RoutePickerViewButtonStylePlain RoutePickerViewButtonStyle = 0
	// RoutePickerViewButtonStyleSystem - A system-defined button style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle/system
	RoutePickerViewButtonStyleSystem RoutePickerViewButtonStyle = 0
)

// AVVideoFrameAnalysisType - Constants that define the types of analysis a player view controller may perform on a paused video frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVVideoFrameAnalysisType
type VideoFrameAnalysisType uint


