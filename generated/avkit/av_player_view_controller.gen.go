// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PlayerViewController] class.
var (
	PlayerViewControllerClass     _PlayerViewControllerClass
	PlayerViewControllerClassOnce sync.Once
)

func getPlayerViewControllerClass() _PlayerViewControllerClass {
	PlayerViewControllerClassOnce.Do(func() {
		PlayerViewControllerClass = _PlayerViewControllerClass{objc.GetClass("AVPlayerViewController")}
	})
	return PlayerViewControllerClass
}

type _PlayerViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [PlayerViewController] class.
type IPlayerViewController interface {
	appkit.IViewController
	BeginTrimmingWithCompletionHandler(handler unsafe.Pointer)
	SelectSpeed(speed unsafe.Pointer)
}

// A view controller that displays content from a player and presents a native user interface to control playback.
//
// A player view controller makes it simple to add media playback capabilities to your app that match the styling and features of the native system players. Using this object also means that your app automatically adopts the new features and styling of future operating system releases.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
type PlayerViewController struct {
	appkit.ViewController
}

// PlayerViewControllerFrom constructs a [PlayerViewController] from an unsafe.Pointer.
//
// A view controller that displays content from a player and presents a native user interface to control playback.
func PlayerViewControllerFrom(ptr unsafe.Pointer) PlayerViewController {
	return PlayerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerViewControllerClass) Alloc() PlayerViewController {
	rv := objc.Send[PlayerViewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerViewControllerClass) New() PlayerViewController {
	rv := objc.Send[PlayerViewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerViewController) Init() PlayerViewController {
	rv := objc.Send[PlayerViewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerViewController) Autorelease() PlayerViewController {
	rv := objc.Send[PlayerViewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerViewController creates a new PlayerViewController instance.
func NewPlayerViewController() PlayerViewController {
	return getPlayerViewControllerClass().New()
}


// Presents the system trimming interface controls inside the player view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/beginTrimming(completionHandler:)
func (p_ PlayerViewController) BeginTrimmingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginTrimmingWithCompletionHandler:"), handler)
}

// Selects a specified playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/selectSpeed(_:)
func (p_ PlayerViewController) SelectSpeed(speed unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectSpeed:"), speed)
}

// An array of language codes that restrict the set of subtitle languages available to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowedSubtitleOptionLanguages
func (p_ PlayerViewController) AllowedSubtitleOptionLanguages() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("allowedSubtitleOptionLanguages"))
	return rv
}


// SetAllowedSubtitleOptionLanguages sets the value of the allowedSubtitleOptionLanguages property.
// An array of language codes that restrict the set of subtitle languages available to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowedSubtitleOptionLanguages
func (p_ PlayerViewController) SetAllowedSubtitleOptionLanguages(value []string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedSubtitleOptionLanguages:"), value)
}
// A Boolean value that indicates whether the player allows Picture in Picture playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsPictureInPicturePlayback
func (p_ PlayerViewController) AllowsPictureInPicturePlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsPictureInPicturePlayback"))
	return rv
}


// SetAllowsPictureInPicturePlayback sets the value of the allowsPictureInPicturePlayback property.
// A Boolean value that indicates whether the player allows Picture in Picture playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsPictureInPicturePlayback
func (p_ PlayerViewController) SetAllowsPictureInPicturePlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsPictureInPicturePlayback:"), value)
}
// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsVideoFrameAnalysis
func (p_ PlayerViewController) AllowsVideoFrameAnalysis() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsVideoFrameAnalysis"))
	return rv
}


// SetAllowsVideoFrameAnalysis sets the value of the allowsVideoFrameAnalysis property.
// A Boolean value that indicates whether to perform video frame analysis.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsVideoFrameAnalysis
func (p_ PlayerViewController) SetAllowsVideoFrameAnalysis(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsVideoFrameAnalysis:"), value)
}
// A Boolean value that indicates whether the view controller automatically sets the screen’s display criteria to match that of the currently playing asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/appliesPreferredDisplayCriteriaAutomatically
func (p_ PlayerViewController) AppliesPreferredDisplayCriteriaAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("appliesPreferredDisplayCriteriaAutomatically"))
	return rv
}


// SetAppliesPreferredDisplayCriteriaAutomatically sets the value of the appliesPreferredDisplayCriteriaAutomatically property.
// A Boolean value that indicates whether the view controller automatically sets the screen’s display criteria to match that of the currently playing asset.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/appliesPreferredDisplayCriteriaAutomatically
func (p_ PlayerViewController) SetAppliesPreferredDisplayCriteriaAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppliesPreferredDisplayCriteriaAutomatically:"), value)
}
// A Boolean value that indicates whether the current media supports trimming.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canBeginTrimming
func (p_ PlayerViewController) CanBeginTrimming() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canBeginTrimming"))
	return rv
}

// A Boolean value that indicates whether Picture in Picture starts automatically when transitioning to the background when the view controller presents its content inline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PlayerViewController) CanStartPictureInPictureAutomaticallyFromInline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStartPictureInPictureAutomaticallyFromInline"))
	return rv
}


// SetCanStartPictureInPictureAutomaticallyFromInline sets the value of the canStartPictureInPictureAutomaticallyFromInline property.
// A Boolean value that indicates whether Picture in Picture starts automatically when transitioning to the background when the view controller presents its content inline.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PlayerViewController) SetCanStartPictureInPictureAutomaticallyFromInline(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanStartPictureInPictureAutomaticallyFromInline:"), value)
}
// A view that displays between the video content and the playback controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentOverlayView
func (p_ PlayerViewController) ContentOverlayView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentOverlayView"))
	return rv
}

// The view controller responsible for the presentation of content proposals.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentProposalViewController
func (p_ PlayerViewController) ContentProposalViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentProposalViewController"))
	return rv
}


// SetContentProposalViewController sets the value of the contentProposalViewController property.
// The view controller responsible for the presentation of content proposals.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentProposalViewController
func (p_ PlayerViewController) SetContentProposalViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentProposalViewController:"), value)
}
// An array of action controls to present contextually during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActions
func (p_ PlayerViewController) ContextualActions() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("contextualActions"))
	return rv
}


// SetContextualActions sets the value of the contextualActions property.
// An array of action controls to present contextually during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActions
func (p_ PlayerViewController) SetContextualActions(value []unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContextualActions:"), value)
}
// A view the system shows adjacent to the contextual actions that’s suitable for showing related information.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsInfoView
func (p_ PlayerViewController) ContextualActionsInfoView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contextualActionsInfoView"))
	return rv
}

// An image to show alongside the contextual actions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsPreviewImage
func (p_ PlayerViewController) ContextualActionsPreviewImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contextualActionsPreviewImage"))
	return rv
}


// SetContextualActionsPreviewImage sets the value of the contextualActionsPreviewImage property.
// An image to show alongside the contextual actions.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsPreviewImage
func (p_ PlayerViewController) SetContextualActionsPreviewImage(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContextualActionsPreviewImage:"), value)
}
// A view controller that provides client-specific content and controls alongside system-provided information and settings panels.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewController
func (p_ PlayerViewController) CustomInfoViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("customInfoViewController"))
	return rv
}


// SetCustomInfoViewController sets the value of the customInfoViewController property.
// A view controller that provides client-specific content and controls alongside system-provided information and settings panels.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewController
func (p_ PlayerViewController) SetCustomInfoViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomInfoViewController:"), value)
}
// An array of view controllers to display as content tabs in the player user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewControllers
func (p_ PlayerViewController) CustomInfoViewControllers() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("customInfoViewControllers"))
	return rv
}


// SetCustomInfoViewControllers sets the value of the customInfoViewControllers property.
// An array of view controllers to display as content tabs in the player user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewControllers
func (p_ PlayerViewController) SetCustomInfoViewControllers(value []unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomInfoViewControllers:"), value)
}
// A view controller that presents custom content over the player view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customOverlayViewController
func (p_ PlayerViewController) CustomOverlayViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("customOverlayViewController"))
	return rv
}


// SetCustomOverlayViewController sets the value of the customOverlayViewController property.
// A view controller that presents custom content over the player view.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customOverlayViewController
func (p_ PlayerViewController) SetCustomOverlayViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomOverlayViewController:"), value)
}
// The delegate object for the player view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/delegate
func (p_ PlayerViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the player view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/delegate
func (p_ PlayerViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that determines whether the player automatically displays in full screen when the user taps the play button.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/entersFullScreenWhenPlaybackBegins
func (p_ PlayerViewController) EntersFullScreenWhenPlaybackBegins() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("entersFullScreenWhenPlaybackBegins"))
	return rv
}


// SetEntersFullScreenWhenPlaybackBegins sets the value of the entersFullScreenWhenPlaybackBegins property.
// A Boolean value that determines whether the player automatically displays in full screen when the user taps the play button.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/entersFullScreenWhenPlaybackBegins
func (p_ PlayerViewController) SetEntersFullScreenWhenPlaybackBegins(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEntersFullScreenWhenPlaybackBegins:"), value)
}
// A Boolean value that indicates whether the player exits full-screen mode when playback ends.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/exitsFullScreenWhenPlaybackEnds
func (p_ PlayerViewController) ExitsFullScreenWhenPlaybackEnds() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("exitsFullScreenWhenPlaybackEnds"))
	return rv
}


// SetExitsFullScreenWhenPlaybackEnds sets the value of the exitsFullScreenWhenPlaybackEnds property.
// A Boolean value that indicates whether the player exits full-screen mode when playback ends.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/exitsFullScreenWhenPlaybackEnds
func (p_ PlayerViewController) SetExitsFullScreenWhenPlaybackEnds(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExitsFullScreenWhenPlaybackEnds:"), value)
}
// The group experience coordinator for this view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/groupExperienceCoordinator
func (p_ PlayerViewController) GroupExperienceCoordinator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("groupExperienceCoordinator"))
	return rv
}

// An array of actions to present in the Info content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/infoViewActions
func (p_ PlayerViewController) InfoViewActions() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("infoViewActions"))
	return rv
}


// SetInfoViewActions sets the value of the infoViewActions property.
// An array of actions to present in the Info content view.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/infoViewActions
func (p_ PlayerViewController) SetInfoViewActions(value []unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInfoViewActions:"), value)
}
// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isReadyForDisplay
func (p_ PlayerViewController) ReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readyForDisplay"))
	return rv
}

// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipBackwardEnabled
func (p_ PlayerViewController) SkipBackwardEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("skipBackwardEnabled"))
	return rv
}


// SetSkipBackwardEnabled sets the value of the skipBackwardEnabled property.
// A Boolean value that indicates whether backward-skipping is available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipBackwardEnabled
func (p_ PlayerViewController) SetSkipBackwardEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipBackwardEnabled:"), value)
}
// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipForwardEnabled
func (p_ PlayerViewController) SkipForwardEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("skipForwardEnabled"))
	return rv
}


// SetSkipForwardEnabled sets the value of the skipForwardEnabled property.
// A Boolean value that indicates whether forward-skipping is available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipForwardEnabled
func (p_ PlayerViewController) SetSkipForwardEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipForwardEnabled:"), value)
}
// The pixel buffer attributes of the video frames the view controller presents.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/pixelBufferAttributes
func (p_ PlayerViewController) PixelBufferAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// SetPixelBufferAttributes sets the value of the pixelBufferAttributes property.
// The pixel buffer attributes of the video frames the view controller presents.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/pixelBufferAttributes
func (p_ PlayerViewController) SetPixelBufferAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}
// A Boolean value that indicates whether the player presents video metadata, navigation markers, and playback settings views when the user requests them.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeInfoViews
func (p_ PlayerViewController) PlaybackControlsIncludeInfoViews() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackControlsIncludeInfoViews"))
	return rv
}


// SetPlaybackControlsIncludeInfoViews sets the value of the playbackControlsIncludeInfoViews property.
// A Boolean value that indicates whether the player presents video metadata, navigation markers, and playback settings views when the user requests them.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeInfoViews
func (p_ PlayerViewController) SetPlaybackControlsIncludeInfoViews(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackControlsIncludeInfoViews:"), value)
}
// A Boolean value that indicates whether the player shows the transport bar and related controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeTransportBar
func (p_ PlayerViewController) PlaybackControlsIncludeTransportBar() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackControlsIncludeTransportBar"))
	return rv
}


// SetPlaybackControlsIncludeTransportBar sets the value of the playbackControlsIncludeTransportBar property.
// A Boolean value that indicates whether the player shows the transport bar and related controls.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeTransportBar
func (p_ PlayerViewController) SetPlaybackControlsIncludeTransportBar(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackControlsIncludeTransportBar:"), value)
}
// The player object that provides the media content for the view controller to display.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/player
func (p_ PlayerViewController) Player() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("player"))
	return rv
}


// SetPlayer sets the value of the player property.
// The player object that provides the media content for the view controller to display.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/player
func (p_ PlayerViewController) SetPlayer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}
// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/preferredDisplayDynamicRange
func (p_ PlayerViewController) PreferredDisplayDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredDisplayDynamicRange"))
	return rv
}


// SetPreferredDisplayDynamicRange sets the value of the preferredDisplayDynamicRange property.
// Describes how High Dynamic Range (HDR) video content renders.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/preferredDisplayDynamicRange
func (p_ PlayerViewController) SetPreferredDisplayDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredDisplayDynamicRange:"), value)
}
// A Boolean value that indicates whether the user can disable the display of subtitles.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresFullSubtitles
func (p_ PlayerViewController) RequiresFullSubtitles() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresFullSubtitles"))
	return rv
}


// SetRequiresFullSubtitles sets the value of the requiresFullSubtitles property.
// A Boolean value that indicates whether the user can disable the display of subtitles.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresFullSubtitles
func (p_ PlayerViewController) SetRequiresFullSubtitles(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresFullSubtitles:"), value)
}
// A Boolean value that determines whether the player allows the user to skip media content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresLinearPlayback
func (p_ PlayerViewController) RequiresLinearPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresLinearPlayback"))
	return rv
}


// SetRequiresLinearPlayback sets the value of the requiresLinearPlayback property.
// A Boolean value that determines whether the player allows the user to skip media content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresLinearPlayback
func (p_ PlayerViewController) SetRequiresLinearPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresLinearPlayback:"), value)
}
// A Boolean value that indicates whether to permit playback of 2D video content only.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresMonoscopicViewingMode
func (p_ PlayerViewController) RequiresMonoscopicViewingMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresMonoscopicViewingMode"))
	return rv
}


// SetRequiresMonoscopicViewingMode sets the value of the requiresMonoscopicViewingMode property.
// A Boolean value that indicates whether to permit playback of 2D video content only.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresMonoscopicViewingMode
func (p_ PlayerViewController) SetRequiresMonoscopicViewingMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresMonoscopicViewingMode:"), value)
}
// The currently selected playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/selectedSpeed
func (p_ PlayerViewController) SelectedSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}

// A Boolean value that indicates whether the player view controller shows playback controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsPlaybackControls
func (p_ PlayerViewController) ShowsPlaybackControls() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsPlaybackControls"))
	return rv
}


// SetShowsPlaybackControls sets the value of the showsPlaybackControls property.
// A Boolean value that indicates whether the player view controller shows playback controls.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsPlaybackControls
func (p_ PlayerViewController) SetShowsPlaybackControls(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsPlaybackControls:"), value)
}
// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsTimecodes
func (p_ PlayerViewController) ShowsTimecodes() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsTimecodes"))
	return rv
}


// SetShowsTimecodes sets the value of the showsTimecodes property.
// A Boolean value that determines whether the player view displays timecodes, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsTimecodes
func (p_ PlayerViewController) SetShowsTimecodes(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsTimecodes:"), value)
}
// The behavior that skipping gestures perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/skippingBehavior
func (p_ PlayerViewController) SkippingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("skippingBehavior"))
	return rv
}


// SetSkippingBehavior sets the value of the skippingBehavior property.
// The behavior that skipping gestures perform.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/skippingBehavior
func (p_ PlayerViewController) SetSkippingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkippingBehavior:"), value)
}
// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/speeds
func (p_ PlayerViewController) Speeds() []PlaybackSpeed {
	rv := objc.Send[[]PlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}


// SetSpeeds sets the value of the speeds property.
// A list of user-selectable playback speeds to show in the playback speed control.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/speeds
func (p_ PlayerViewController) SetSpeeds(value []PlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpeeds:"), value)
}
// An action that enables the visual lookup interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/toggleLookupAction
func (p_ PlayerViewController) ToggleLookupAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("toggleLookupAction"))
	return rv
}

// An array of actions and menus to display with the default player controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarCustomMenuItems
func (p_ PlayerViewController) TransportBarCustomMenuItems() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("transportBarCustomMenuItems"))
	return rv
}


// SetTransportBarCustomMenuItems sets the value of the transportBarCustomMenuItems property.
// An array of actions and menus to display with the default player controls.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarCustomMenuItems
func (p_ PlayerViewController) SetTransportBarCustomMenuItems(value []unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransportBarCustomMenuItems:"), value)
}
// A Boolean value that indicates whether the player user interface shows the title view above the scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarIncludesTitleView
func (p_ PlayerViewController) TransportBarIncludesTitleView() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("transportBarIncludesTitleView"))
	return rv
}


// SetTransportBarIncludesTitleView sets the value of the transportBarIncludesTitleView property.
// A Boolean value that indicates whether the player user interface shows the title view above the scrubber.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarIncludesTitleView
func (p_ PlayerViewController) SetTransportBarIncludesTitleView(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransportBarIncludesTitleView:"), value)
}
// A layout guide that represents an area that fixed-position playback controls don’t obscure when visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/unobscuredContentGuide
func (p_ PlayerViewController) UnobscuredContentGuide() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("unobscuredContentGuide"))
	return rv
}

// A Boolean value that indicates whether the view controller updates Now Playing information.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/updatesNowPlayingInfoCenter
func (p_ PlayerViewController) UpdatesNowPlayingInfoCenter() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("updatesNowPlayingInfoCenter"))
	return rv
}


// SetUpdatesNowPlayingInfoCenter sets the value of the updatesNowPlayingInfoCenter property.
// A Boolean value that indicates whether the view controller updates Now Playing information.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/updatesNowPlayingInfoCenter
func (p_ PlayerViewController) SetUpdatesNowPlayingInfoCenter(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpdatesNowPlayingInfoCenter:"), value)
}
// The size and position of the video image within the bounds of the view controller’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoBounds
func (p_ PlayerViewController) VideoBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("videoBounds"))
	return rv
}

// The types of analysis a player view controller performs on a paused video frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoFrameAnalysisTypes
func (p_ PlayerViewController) VideoFrameAnalysisTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoFrameAnalysisTypes"))
	return rv
}


// SetVideoFrameAnalysisTypes sets the value of the videoFrameAnalysisTypes property.
// The types of analysis a player view controller performs on a paused video frame.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoFrameAnalysisTypes
func (p_ PlayerViewController) SetVideoFrameAnalysisTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFrameAnalysisTypes:"), value)
}
// A string that specifies how the video displays within the bounds of the view controller’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoGravity
func (p_ PlayerViewController) VideoGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoGravity"))
	return rv
}


// SetVideoGravity sets the value of the videoGravity property.
// A string that specifies how the video displays within the bounds of the view controller’s view.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoGravity
func (p_ PlayerViewController) SetVideoGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoGravity:"), value)
}


