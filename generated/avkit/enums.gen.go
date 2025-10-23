// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

// Enum types and constants
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

// AVCaptureEventPhase - Constants that indicate the phase of a system capture event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventPhase
type AVCaptureEventPhase uint

// AVCaptureViewControlsStyle - Constants that describe the capture view’s supported controls styles.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureViewControlsStyle
type AVCaptureViewControlsStyle uint

// AVContentProposalAction - Constant that indicate the action a user takes when dismissing a content proposal.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalAction
type AVContentProposalAction uint

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

// AVKitError - Constants that identify framework error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-swift.struct/Code
type AVKitError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVKitError-c.enum/AVKitErrorRecordingFailed
	AVKitErrorRecordingFailed AVKitError = 0
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
)

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

// AVPlayerViewControlsStyle - Constants that indicate which user interface controls the view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewControlsStyle
type AVPlayerViewControlsStyle uint

// AVPlayerViewTrimResult - Constants that specify an action a user takes when trimming media in a player view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewTrimResult
type AVPlayerViewTrimResult uint

// AVRoutePickerViewButtonState - Constants that describe the available button states.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/ButtonState
type AVRoutePickerViewButtonState uint

// AVRoutePickerViewButtonStyle - Constants that define the button styles a route picker view supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerViewButtonStyle
type AVRoutePickerViewButtonStyle uint

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


