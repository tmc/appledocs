// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AllowedSubtitleOptionLanguages() []string /* primitive/slice/pointer. */
	SetAllowedSubtitleOptionLanguages(value []string /* primitive/slice/pointer. */)
	AllowsPictureInPicturePlayback() bool /* primitive/slice/pointer. */
	SetAllowsPictureInPicturePlayback(value bool /* primitive/slice/pointer. */)
	AllowsVideoFrameAnalysis() bool /* primitive/slice/pointer. */
	SetAllowsVideoFrameAnalysis(value bool /* primitive/slice/pointer. */)
	AppliesPreferredDisplayCriteriaAutomatically() bool /* primitive/slice/pointer. */
	SetAppliesPreferredDisplayCriteriaAutomatically(value bool /* primitive/slice/pointer. */)
	CanBeginTrimming() bool /* primitive/slice/pointer. */
	CanStartPictureInPictureAutomaticallyFromInline() bool /* primitive/slice/pointer. */
	SetCanStartPictureInPictureAutomaticallyFromInline(value bool /* primitive/slice/pointer. */)
	ContentOverlayView() objc.IObject /* cross-framework: View */
	ContentProposalViewController() IAVContentProposalViewController
	SetContentProposalViewController(value IAVContentProposalViewController)
	ContextualActions() []Action /* primitive/slice/pointer. */
	SetContextualActions(value []Action /* primitive/slice/pointer. */)
	ContextualActionsInfoView() objc.IObject /* cross-framework: View */
	ContextualActionsPreviewImage() objc.IObject /* cross-framework: Image */
	SetContextualActionsPreviewImage(value objc.IObject /* cross-framework: Image */)
	CustomInfoViewController() objc.IObject /* cross-framework: ViewController */
	SetCustomInfoViewController(value objc.IObject /* cross-framework: ViewController */)
	CustomInfoViewControllers() []appkit.objc.IObject /* cross-framework: ViewController */
	SetCustomInfoViewControllers(value []appkit.objc.IObject /* cross-framework: ViewController */)
	CustomOverlayViewController() objc.IObject /* cross-framework: ViewController */
	SetCustomOverlayViewController(value objc.IObject /* cross-framework: ViewController */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	EntersFullScreenWhenPlaybackBegins() bool /* primitive/slice/pointer. */
	SetEntersFullScreenWhenPlaybackBegins(value bool /* primitive/slice/pointer. */)
	ExitsFullScreenWhenPlaybackEnds() bool /* primitive/slice/pointer. */
	SetExitsFullScreenWhenPlaybackEnds(value bool /* primitive/slice/pointer. */)
	GroupExperienceCoordinator() GroupExperienceCoordinator /* not a class type */
	InfoViewActions() []Action /* primitive/slice/pointer. */
	SetInfoViewActions(value []Action /* primitive/slice/pointer. */)
	ReadyForDisplay() bool /* primitive/slice/pointer. */
	SkipBackwardEnabled() bool /* primitive/slice/pointer. */
	SetSkipBackwardEnabled(value bool /* primitive/slice/pointer. */)
	SkipForwardEnabled() bool /* primitive/slice/pointer. */
	SetSkipForwardEnabled(value bool /* primitive/slice/pointer. */)
	PixelBufferAttributes() foundation.IDictionary /* already interface */
	SetPixelBufferAttributes(value foundation.IDictionary /* already interface */)
	PlaybackControlsIncludeInfoViews() bool /* primitive/slice/pointer. */
	SetPlaybackControlsIncludeInfoViews(value bool /* primitive/slice/pointer. */)
	PlaybackControlsIncludeTransportBar() bool /* primitive/slice/pointer. */
	SetPlaybackControlsIncludeTransportBar(value bool /* primitive/slice/pointer. */)
	Player() objc.IObject /* cross-framework: Player */
	SetPlayer(value objc.IObject /* cross-framework: Player */)
	PreferredDisplayDynamicRange() DisplayDynamicRange
	SetPreferredDisplayDynamicRange(value DisplayDynamicRange)
	RequiresFullSubtitles() bool /* primitive/slice/pointer. */
	SetRequiresFullSubtitles(value bool /* primitive/slice/pointer. */)
	RequiresLinearPlayback() bool /* primitive/slice/pointer. */
	SetRequiresLinearPlayback(value bool /* primitive/slice/pointer. */)
	RequiresMonoscopicViewingMode() bool /* primitive/slice/pointer. */
	SetRequiresMonoscopicViewingMode(value bool /* primitive/slice/pointer. */)
	SelectedSpeed() IAVPlaybackSpeed
	ShowsPlaybackControls() bool /* primitive/slice/pointer. */
	SetShowsPlaybackControls(value bool /* primitive/slice/pointer. */)
	ShowsTimecodes() bool /* primitive/slice/pointer. */
	SetShowsTimecodes(value bool /* primitive/slice/pointer. */)
	SkippingBehavior() PlayerViewControllerSkippingBehavior
	SetSkippingBehavior(value PlayerViewControllerSkippingBehavior)
	Speeds() []PlaybackSpeed /* primitive/slice/pointer. */
	SetSpeeds(value []PlaybackSpeed /* primitive/slice/pointer. */)
	ToggleLookupAction() Action /* not a class type */
	TransportBarCustomMenuItems() []MenuElement /* primitive/slice/pointer. */
	SetTransportBarCustomMenuItems(value []MenuElement /* primitive/slice/pointer. */)
	TransportBarIncludesTitleView() bool /* primitive/slice/pointer. */
	SetTransportBarIncludesTitleView(value bool /* primitive/slice/pointer. */)
	UnobscuredContentGuide() objc.IObject /* cross-framework: LayoutGuide */
	UpdatesNowPlayingInfoCenter() bool /* primitive/slice/pointer. */
	SetUpdatesNowPlayingInfoCenter(value bool /* primitive/slice/pointer. */)
	VideoBounds() coregraphics.CGRect
	VideoFrameAnalysisTypes() VideoFrameAnalysisType
	SetVideoFrameAnalysisTypes(value VideoFrameAnalysisType)
	VideoGravity() LayerVideoGravity /* not a class type */
	SetVideoGravity(value LayerVideoGravity /* not a class type */)
	ExperienceController() ExperienceController /* not a class type */
	SetExperienceController(value ExperienceController /* not a class type */)
	IsReadyForDisplay() bool /* primitive/slice/pointer. */
	SetIsReadyForDisplay(value bool /* primitive/slice/pointer. */)
	IsSkipBackwardEnabled() bool /* primitive/slice/pointer. */
	SetIsSkipBackwardEnabled(value bool /* primitive/slice/pointer. */)
	IsSkipForwardEnabled() bool /* primitive/slice/pointer. */
	SetIsSkipForwardEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
	BeginTrimmingWithCompletionHandler(handler unsafe.Pointer)
	SelectSpeed(speed IAVPlaybackSpeed)
}

// A view controller that displays content from a player and presents a native user interface to control playback.
//
// A player view controller makes it simple to add media playback capabilities to your app that match the styling and features of the native system players. Using this object also means that your app automatically adopts the new features and styling of future operating system releases.


// A view controller that displays content from a player and presents a native user interface to control playback.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/mediaCharacteristicsForSupportedCustomMediaSelectionSchemes
func (pc _PlayerViewControllerClass) MediaCharacteristicsForSupportedCustomMediaSelectionSchemes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("mediaCharacteristicsForSupportedCustomMediaSelectionSchemes"))
	return rv
}

// Presents the system trimming interface controls inside the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/beginTrimming(completionHandler:)
func (p_ PlayerViewController) BeginTrimmingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginTrimmingWithCompletionHandler:"), handler)
}


// Selects a specified playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/selectSpeed(_:)
func (p_ PlayerViewController) SelectSpeed(speed IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectSpeed:"), speed)
}


// An array of language codes that restrict the set of subtitle languages available to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowedSubtitleOptionLanguages
func (p_ PlayerViewController) AllowedSubtitleOptionLanguages() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](p_.ID, objc.Sel("allowedSubtitleOptionLanguages"))
	return rv
}


// An array of language codes that restrict the set of subtitle languages available to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowedSubtitleOptionLanguages
func (p_ PlayerViewController) SetAllowedSubtitleOptionLanguages(value []string /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedSubtitleOptionLanguages:"), nsArray)
}


// A Boolean value that indicates whether the player allows Picture in Picture playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsPictureInPicturePlayback
func (p_ PlayerViewController) AllowsPictureInPicturePlayback() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsPictureInPicturePlayback"))
	return rv
}


// A Boolean value that indicates whether the player allows Picture in Picture playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsPictureInPicturePlayback
func (p_ PlayerViewController) SetAllowsPictureInPicturePlayback(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsPictureInPicturePlayback:"), value)
}


// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsVideoFrameAnalysis
func (p_ PlayerViewController) AllowsVideoFrameAnalysis() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsVideoFrameAnalysis"))
	return rv
}


// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/allowsVideoFrameAnalysis
func (p_ PlayerViewController) SetAllowsVideoFrameAnalysis(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsVideoFrameAnalysis:"), value)
}


// A Boolean value that indicates whether the view controller automatically sets the screen’s display criteria to match that of the currently playing asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/appliesPreferredDisplayCriteriaAutomatically
func (p_ PlayerViewController) AppliesPreferredDisplayCriteriaAutomatically() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("appliesPreferredDisplayCriteriaAutomatically"))
	return rv
}


// A Boolean value that indicates whether the view controller automatically sets the screen’s display criteria to match that of the currently playing asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/appliesPreferredDisplayCriteriaAutomatically
func (p_ PlayerViewController) SetAppliesPreferredDisplayCriteriaAutomatically(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAppliesPreferredDisplayCriteriaAutomatically:"), value)
}


// A Boolean value that indicates whether the current media supports trimming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canBeginTrimming
func (p_ PlayerViewController) CanBeginTrimming() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("canBeginTrimming"))
	return rv
}


// A Boolean value that indicates whether Picture in Picture starts automatically when transitioning to the background when the view controller presents its content inline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PlayerViewController) CanStartPictureInPictureAutomaticallyFromInline() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStartPictureInPictureAutomaticallyFromInline"))
	return rv
}


// A Boolean value that indicates whether Picture in Picture starts automatically when transitioning to the background when the view controller presents its content inline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PlayerViewController) SetCanStartPictureInPictureAutomaticallyFromInline(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanStartPictureInPictureAutomaticallyFromInline:"), value)
}


// A view that displays between the video content and the playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentOverlayView
func (p_ PlayerViewController) ContentOverlayView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[View](p_.ID, objc.Sel("contentOverlayView"))
	return rv
}


// The view controller responsible for the presentation of content proposals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentProposalViewController
func (p_ PlayerViewController) ContentProposalViewController() IAVContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](p_.ID, objc.Sel("contentProposalViewController"))
	return rv
}


// The view controller responsible for the presentation of content proposals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contentProposalViewController
func (p_ PlayerViewController) SetContentProposalViewController(value IAVContentProposalViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentProposalViewController:"), value)
}


// An array of action controls to present contextually during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActions
func (p_ PlayerViewController) ContextualActions() []Action /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Action](p_.ID, objc.Sel("contextualActions"))
	return rv
}


// An array of action controls to present contextually during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActions
func (p_ PlayerViewController) SetContextualActions(value []Action /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setContextualActions:"), nsArray)
}


// A view the system shows adjacent to the contextual actions that’s suitable for showing related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsInfoView
func (p_ PlayerViewController) ContextualActionsInfoView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[View](p_.ID, objc.Sel("contextualActionsInfoView"))
	return rv
}


// An image to show alongside the contextual actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsPreviewImage
func (p_ PlayerViewController) ContextualActionsPreviewImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[Image](p_.ID, objc.Sel("contextualActionsPreviewImage"))
	return rv
}


// An image to show alongside the contextual actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/contextualActionsPreviewImage
func (p_ PlayerViewController) SetContextualActionsPreviewImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContextualActionsPreviewImage:"), value)
}


// A view controller that provides client-specific content and controls alongside system-provided information and settings panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewController
func (p_ PlayerViewController) CustomInfoViewController() objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("customInfoViewController"))
	return rv
}


// A view controller that provides client-specific content and controls alongside system-provided information and settings panels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewController
func (p_ PlayerViewController) SetCustomInfoViewController(value objc.IObject /* cross-framework: ViewController */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomInfoViewController:"), value)
}


// An array of view controllers to display as content tabs in the player user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewControllers
func (p_ PlayerViewController) CustomInfoViewControllers() []appkit.objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[[]appkit.ViewController](p_.ID, objc.Sel("customInfoViewControllers"))
	return rv
}


// An array of view controllers to display as content tabs in the player user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customInfoViewControllers
func (p_ PlayerViewController) SetCustomInfoViewControllers(value []appkit.objc.IObject /* cross-framework: ViewController */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomInfoViewControllers:"), nsArray)
}


// A view controller that presents custom content over the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customOverlayViewController
func (p_ PlayerViewController) CustomOverlayViewController() objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("customOverlayViewController"))
	return rv
}


// A view controller that presents custom content over the player view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/customOverlayViewController
func (p_ PlayerViewController) SetCustomOverlayViewController(value objc.IObject /* cross-framework: ViewController */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomOverlayViewController:"), value)
}


// The delegate object for the player view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/delegate
func (p_ PlayerViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the player view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/delegate
func (p_ PlayerViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that determines whether the player automatically displays in full screen when the user taps the play button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/entersFullScreenWhenPlaybackBegins
func (p_ PlayerViewController) EntersFullScreenWhenPlaybackBegins() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("entersFullScreenWhenPlaybackBegins"))
	return rv
}


// A Boolean value that determines whether the player automatically displays in full screen when the user taps the play button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/entersFullScreenWhenPlaybackBegins
func (p_ PlayerViewController) SetEntersFullScreenWhenPlaybackBegins(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEntersFullScreenWhenPlaybackBegins:"), value)
}


// A Boolean value that indicates whether the player exits full-screen mode when playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/exitsFullScreenWhenPlaybackEnds
func (p_ PlayerViewController) ExitsFullScreenWhenPlaybackEnds() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("exitsFullScreenWhenPlaybackEnds"))
	return rv
}


// A Boolean value that indicates whether the player exits full-screen mode when playback ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/exitsFullScreenWhenPlaybackEnds
func (p_ PlayerViewController) SetExitsFullScreenWhenPlaybackEnds(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExitsFullScreenWhenPlaybackEnds:"), value)
}


// The group experience coordinator for this view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/groupExperienceCoordinator
func (p_ PlayerViewController) GroupExperienceCoordinator() GroupExperienceCoordinator /* not a class type */ {
	rv := objc.Send[GroupExperienceCoordinator](p_.ID, objc.Sel("groupExperienceCoordinator"))
	return rv
}


// An array of actions to present in the Info content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/infoViewActions
func (p_ PlayerViewController) InfoViewActions() []Action /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Action](p_.ID, objc.Sel("infoViewActions"))
	return rv
}


// An array of actions to present in the Info content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/infoViewActions
func (p_ PlayerViewController) SetInfoViewActions(value []Action /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setInfoViewActions:"), nsArray)
}


// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isReadyForDisplay
func (p_ PlayerViewController) ReadyForDisplay() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("readyForDisplay"))
	return rv
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipBackwardEnabled
func (p_ PlayerViewController) SkipBackwardEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("skipBackwardEnabled"))
	return rv
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipBackwardEnabled
func (p_ PlayerViewController) SetSkipBackwardEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipBackwardEnabled:"), value)
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipForwardEnabled
func (p_ PlayerViewController) SkipForwardEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("skipForwardEnabled"))
	return rv
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/isSkipForwardEnabled
func (p_ PlayerViewController) SetSkipForwardEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkipForwardEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/mediaCharacteristicsForSupportedCustomMediaSelectionSchemes
func (p_ PlayerViewController) MediaCharacteristicsForSupportedCustomMediaSelectionSchemes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](p_.ID, objc.Sel("mediaCharacteristicsForSupportedCustomMediaSelectionSchemes"))
	return rv
}


// The pixel buffer attributes of the video frames the view controller presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/pixelBufferAttributes
func (p_ PlayerViewController) PixelBufferAttributes() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// The pixel buffer attributes of the video frames the view controller presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/pixelBufferAttributes
func (p_ PlayerViewController) SetPixelBufferAttributes(value foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}


// A Boolean value that indicates whether the player presents video metadata, navigation markers, and playback settings views when the user requests them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeInfoViews
func (p_ PlayerViewController) PlaybackControlsIncludeInfoViews() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackControlsIncludeInfoViews"))
	return rv
}


// A Boolean value that indicates whether the player presents video metadata, navigation markers, and playback settings views when the user requests them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeInfoViews
func (p_ PlayerViewController) SetPlaybackControlsIncludeInfoViews(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackControlsIncludeInfoViews:"), value)
}


// A Boolean value that indicates whether the player shows the transport bar and related controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeTransportBar
func (p_ PlayerViewController) PlaybackControlsIncludeTransportBar() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("playbackControlsIncludeTransportBar"))
	return rv
}


// A Boolean value that indicates whether the player shows the transport bar and related controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/playbackControlsIncludeTransportBar
func (p_ PlayerViewController) SetPlaybackControlsIncludeTransportBar(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackControlsIncludeTransportBar:"), value)
}


// The player object that provides the media content for the view controller to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/player
func (p_ PlayerViewController) Player() objc.IObject /* cross-framework: Player */ {
	rv := objc.Send[Player](p_.ID, objc.Sel("player"))
	return rv
}


// The player object that provides the media content for the view controller to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/player
func (p_ PlayerViewController) SetPlayer(value objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}


// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/preferredDisplayDynamicRange
func (p_ PlayerViewController) PreferredDisplayDynamicRange() DisplayDynamicRange {
	rv := objc.Send[DisplayDynamicRange](p_.ID, objc.Sel("preferredDisplayDynamicRange"))
	return rv
}


// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/preferredDisplayDynamicRange
func (p_ PlayerViewController) SetPreferredDisplayDynamicRange(value DisplayDynamicRange) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredDisplayDynamicRange:"), value)
}


// A Boolean value that indicates whether the user can disable the display of subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresFullSubtitles
func (p_ PlayerViewController) RequiresFullSubtitles() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresFullSubtitles"))
	return rv
}


// A Boolean value that indicates whether the user can disable the display of subtitles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresFullSubtitles
func (p_ PlayerViewController) SetRequiresFullSubtitles(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresFullSubtitles:"), value)
}


// A Boolean value that determines whether the player allows the user to skip media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresLinearPlayback
func (p_ PlayerViewController) RequiresLinearPlayback() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresLinearPlayback"))
	return rv
}


// A Boolean value that determines whether the player allows the user to skip media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresLinearPlayback
func (p_ PlayerViewController) SetRequiresLinearPlayback(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresLinearPlayback:"), value)
}


// A Boolean value that indicates whether to permit playback of 2D video content only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresMonoscopicViewingMode
func (p_ PlayerViewController) RequiresMonoscopicViewingMode() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresMonoscopicViewingMode"))
	return rv
}


// A Boolean value that indicates whether to permit playback of 2D video content only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/requiresMonoscopicViewingMode
func (p_ PlayerViewController) SetRequiresMonoscopicViewingMode(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresMonoscopicViewingMode:"), value)
}


// The currently selected playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/selectedSpeed
func (p_ PlayerViewController) SelectedSpeed() IAVPlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}


// A Boolean value that indicates whether the player view controller shows playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsPlaybackControls
func (p_ PlayerViewController) ShowsPlaybackControls() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsPlaybackControls"))
	return rv
}


// A Boolean value that indicates whether the player view controller shows playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsPlaybackControls
func (p_ PlayerViewController) SetShowsPlaybackControls(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsPlaybackControls:"), value)
}


// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsTimecodes
func (p_ PlayerViewController) ShowsTimecodes() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsTimecodes"))
	return rv
}


// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/showsTimecodes
func (p_ PlayerViewController) SetShowsTimecodes(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsTimecodes:"), value)
}


// The behavior that skipping gestures perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/skippingBehavior
func (p_ PlayerViewController) SkippingBehavior() PlayerViewControllerSkippingBehavior {
	rv := objc.Send[PlayerViewControllerSkippingBehavior](p_.ID, objc.Sel("skippingBehavior"))
	return rv
}


// The behavior that skipping gestures perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/skippingBehavior
func (p_ PlayerViewController) SetSkippingBehavior(value PlayerViewControllerSkippingBehavior) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSkippingBehavior:"), value)
}


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/speeds
func (p_ PlayerViewController) Speeds() []PlaybackSpeed /* primitive/slice/pointer. */ {
	rv := objc.Send[[]PlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/speeds
func (p_ PlayerViewController) SetSpeeds(value []PlaybackSpeed /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpeeds:"), nsArray)
}


// An action that enables the visual lookup interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/toggleLookupAction
func (p_ PlayerViewController) ToggleLookupAction() Action /* not a class type */ {
	rv := objc.Send[Action](p_.ID, objc.Sel("toggleLookupAction"))
	return rv
}


// An array of actions and menus to display with the default player controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarCustomMenuItems
func (p_ PlayerViewController) TransportBarCustomMenuItems() []MenuElement /* primitive/slice/pointer. */ {
	rv := objc.Send[[]MenuElement](p_.ID, objc.Sel("transportBarCustomMenuItems"))
	return rv
}


// An array of actions and menus to display with the default player controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarCustomMenuItems
func (p_ PlayerViewController) SetTransportBarCustomMenuItems(value []MenuElement /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransportBarCustomMenuItems:"), nsArray)
}


// A Boolean value that indicates whether the player user interface shows the title view above the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarIncludesTitleView
func (p_ PlayerViewController) TransportBarIncludesTitleView() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("transportBarIncludesTitleView"))
	return rv
}


// A Boolean value that indicates whether the player user interface shows the title view above the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/transportBarIncludesTitleView
func (p_ PlayerViewController) SetTransportBarIncludesTitleView(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransportBarIncludesTitleView:"), value)
}


// A layout guide that represents an area that fixed-position playback controls don’t obscure when visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/unobscuredContentGuide
func (p_ PlayerViewController) UnobscuredContentGuide() objc.IObject /* cross-framework: LayoutGuide */ {
	rv := objc.Send[LayoutGuide](p_.ID, objc.Sel("unobscuredContentGuide"))
	return rv
}


// A Boolean value that indicates whether the view controller updates Now Playing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/updatesNowPlayingInfoCenter
func (p_ PlayerViewController) UpdatesNowPlayingInfoCenter() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("updatesNowPlayingInfoCenter"))
	return rv
}


// A Boolean value that indicates whether the view controller updates Now Playing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/updatesNowPlayingInfoCenter
func (p_ PlayerViewController) SetUpdatesNowPlayingInfoCenter(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpdatesNowPlayingInfoCenter:"), value)
}


// The size and position of the video image within the bounds of the view controller’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoBounds
func (p_ PlayerViewController) VideoBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("videoBounds"))
	return rv
}


// The types of analysis a player view controller performs on a paused video frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoFrameAnalysisTypes
func (p_ PlayerViewController) VideoFrameAnalysisTypes() VideoFrameAnalysisType {
	rv := objc.Send[VideoFrameAnalysisType](p_.ID, objc.Sel("videoFrameAnalysisTypes"))
	return rv
}


// The types of analysis a player view controller performs on a paused video frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoFrameAnalysisTypes
func (p_ PlayerViewController) SetVideoFrameAnalysisTypes(value VideoFrameAnalysisType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFrameAnalysisTypes:"), value)
}


// A string that specifies how the video displays within the bounds of the view controller’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoGravity
func (p_ PlayerViewController) VideoGravity() LayerVideoGravity /* not a class type */ {
	rv := objc.Send[LayerVideoGravity](p_.ID, objc.Sel("videoGravity"))
	return rv
}


// A string that specifies how the video displays within the bounds of the view controller’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/videoGravity
func (p_ PlayerViewController) SetVideoGravity(value LayerVideoGravity /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoGravity:"), value)
}


// The experience controller for this view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/experiencecontroller
func (p_ PlayerViewController) ExperienceController() ExperienceController /* not a class type */ {
	rv := objc.Send[ExperienceController](p_.ID, objc.Sel("experienceController"))
	return rv
}


// The experience controller for this view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/experiencecontroller
func (p_ PlayerViewController) SetExperienceController(value ExperienceController /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExperienceController:"), value)
}


// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isreadyfordisplay
func (p_ PlayerViewController) IsReadyForDisplay() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isreadyfordisplay
func (p_ PlayerViewController) SetIsReadyForDisplay(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipbackwardenabled
func (p_ PlayerViewController) IsSkipBackwardEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSkipBackwardEnabled"))
	return rv
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipbackwardenabled
func (p_ PlayerViewController) SetIsSkipBackwardEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSkipBackwardEnabled:"), value)
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipforwardenabled
func (p_ PlayerViewController) IsSkipForwardEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSkipForwardEnabled"))
	return rv
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipforwardenabled
func (p_ PlayerViewController) SetIsSkipForwardEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSkipForwardEnabled:"), value)
}



