// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AVPlayerView */


/* debug [class_header]: Header for AVPlayerView */
// The class instance for the [PlayerView] class.
var (
	PlayerViewClass     _PlayerViewClass
	PlayerViewClassOnce sync.Once
)

func getPlayerViewClass() _PlayerViewClass {
	PlayerViewClassOnce.Do(func() {
		PlayerViewClass = _PlayerViewClass{objc.GetClass("AVPlayerView")}
	})
	return PlayerViewClass
}

type _PlayerViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerView */
// An interface definition for the [PlayerView] class.
type IPlayerView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for PlayerView */
	// properties:
	ActionPopUpButtonMenu() appkit.Menu
	SetActionPopUpButtonMenu(value appkit.Menu)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	AllowsPictureInPicturePlayback() bool
	SetAllowsPictureInPicturePlayback(value bool)
	AllowsVideoFrameAnalysis() bool
	SetAllowsVideoFrameAnalysis(value bool)
	CanBeginTrimming() bool
	ContentOverlayView() appkit.View
	ControlsStyle() PlayerViewControlsStyle
	SetControlsStyle(value PlayerViewControlsStyle)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ReadyForDisplay() bool
	Magnification() float64
	SetMagnification(value float64)
	PictureInPictureDelegate() unsafe.Pointer
	SetPictureInPictureDelegate(value unsafe.Pointer)
	Player() objc.IObject /* cross-framework: Player */
	SetPlayer(value objc.IObject /* cross-framework: Player */)
	PreferredDisplayDynamicRange() DisplayDynamicRange
	SetPreferredDisplayDynamicRange(value DisplayDynamicRange)
	SelectedSpeed() IAVPlaybackSpeed
	ShowsFrameSteppingButtons() bool
	SetShowsFrameSteppingButtons(value bool)
	ShowsFullScreenToggleButton() bool
	SetShowsFullScreenToggleButton(value bool)
	ShowsSharingServiceButton() bool
	SetShowsSharingServiceButton(value bool)
	ShowsTimecodes() bool
	SetShowsTimecodes(value bool)
	Speeds() []PlaybackSpeed
	SetSpeeds(value []PlaybackSpeed)
	UpdatesNowPlayingInfoCenter() bool
	SetUpdatesNowPlayingInfoCenter(value bool)
	VideoBounds() Rect /* not a class type */
	VideoFrameAnalysisTypes() VideoFrameAnalysisType
	SetVideoFrameAnalysisTypes(value VideoFrameAnalysisType)
	VideoGravity() LayerVideoGravity /* not a class type */
	SetVideoGravity(value LayerVideoGravity /* not a class type */)
	IsReadyForDisplay() bool
	SetIsReadyForDisplay(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerView */
	// methods:
	BeginTrimmingWithCompletionHandler(handler unsafe.Pointer)
	FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle objc.IObject /* cross-framework: NSString */)
	SelectSpeed(speed IAVPlaybackSpeed)
	SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerView */
// Alloc allocates a new instance without initialization.
func (pc _PlayerViewClass) Alloc() PlayerView {
	rv := objc.Send[PlayerView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerViewClass) New() PlayerView {
	rv := objc.Send[PlayerView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerView) Init() PlayerView {
	rv := objc.Send[PlayerView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerView) Autorelease() PlayerView {
	rv := objc.Send[PlayerView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerView creates a new PlayerView instance.
func NewPlayerView() PlayerView {
	return getPlayerViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerView */
// A view that displays content from a player and presents a native user interface to control playback.
//
// The player view supports several controls styles, ranging from no controls to controls matching the look of QuickTime Player. This makes it easy for you to tailor the presentation to best match your use of the player view. Regardless of the selected controls style, the player view always supports the following standard set of keyboard shortcuts to control playback: The Space bar plays and pauses playback. The right and left arrow keys step frame-by-frame through the video. JKL navigation: The J key rewinds. Press it multiple times to cycle through rewind speeds. The K key stops playback. The L key fast-forwards. Press it multiple times to cycle through fast-forward speeds. The player view also makes it simple to add trimming capabilities to your player. Call the view’s method to present a trimming UI that matches the QuickTime Player interface.


// A view that displays content from a player and presents a native user interface to control playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView
type PlayerView struct {
	appkit.View
}

// PlayerViewFrom constructs a [PlayerView] from an unsafe.Pointer.
//
// A view that displays content from a player and presents a native user interface to control playback.
func PlayerViewFrom(ptr unsafe.Pointer) PlayerView {
	return PlayerView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerView */

// Puts the player view into trimming mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/beginTrimming(completionHandler:)
func (p_ PlayerView) BeginTrimmingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginTrimmingWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: BeginTrimmingWithCompletionHandler */


// Displays the chapter number and title in the player view for a brief moment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/flashChapterNumber(_:chapterTitle:)
func (p_ PlayerView) FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("flashChapterNumber:chapterTitle:"), chapterNumber, chapterTitle)
}/* debug [instance_methods/method]: FlashChapterNumberChapterTitle */


// Selects a specified playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/selectSpeed(_:)
func (p_ PlayerView) SelectSpeed(speed IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectSpeed:"), speed)
}/* debug [instance_methods/method]: SelectSpeed */


// Scales the video’s view by a specified factor, and centers the result on a specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/setMagnification(_:centeredAt:)
func (p_ PlayerView) SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}/* debug [instance_methods/method]: SetMagnificationCenteredAtPoint */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerView */

// An action pop-up button menu that the player view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/actionPopUpButtonMenu
func (p_ PlayerView) ActionPopUpButtonMenu() appkit.Menu {
	rv := objc.Send[appkit.Menu](p_.ID, objc.Sel("actionPopUpButtonMenu"))
	return rv
}/* debug [instance_properties/getter]: actionPopUpButtonMenu */


// An action pop-up button menu that the player view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/actionPopUpButtonMenu
func (p_ PlayerView) SetActionPopUpButtonMenu(value appkit.Menu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActionPopUpButtonMenu:"), value)
}/* debug [instance_properties/setter]: actionPopUpButtonMenu */


// A Boolean value that indicates whether the magnify gesture changes the video’s view magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsMagnification
func (p_ PlayerView) AllowsMagnification() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsMagnification"))
	return rv
}/* debug [instance_properties/getter]: allowsMagnification */


// A Boolean value that indicates whether the magnify gesture changes the video’s view magnification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsMagnification
func (p_ PlayerView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsMagnification:"), value)
}/* debug [instance_properties/setter]: allowsMagnification */


// A Boolean value that determines whether the player view allows Picture in Picture playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsPictureInPicturePlayback
func (p_ PlayerView) AllowsPictureInPicturePlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsPictureInPicturePlayback"))
	return rv
}/* debug [instance_properties/getter]: allowsPictureInPicturePlayback */


// A Boolean value that determines whether the player view allows Picture in Picture playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsPictureInPicturePlayback
func (p_ PlayerView) SetAllowsPictureInPicturePlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsPictureInPicturePlayback:"), value)
}/* debug [instance_properties/setter]: allowsPictureInPicturePlayback */


// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsVideoFrameAnalysis
func (p_ PlayerView) AllowsVideoFrameAnalysis() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsVideoFrameAnalysis"))
	return rv
}/* debug [instance_properties/getter]: allowsVideoFrameAnalysis */


// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsVideoFrameAnalysis
func (p_ PlayerView) SetAllowsVideoFrameAnalysis(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsVideoFrameAnalysis:"), value)
}/* debug [instance_properties/setter]: allowsVideoFrameAnalysis */


// A Boolean value that indicates whether the player view can begin trimming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/canBeginTrimming
func (p_ PlayerView) CanBeginTrimming() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canBeginTrimming"))
	return rv
}/* debug [instance_properties/getter]: canBeginTrimming */


// A view that adds additional custom views between the video content and the controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/contentOverlayView
func (p_ PlayerView) ContentOverlayView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("contentOverlayView"))
	return rv
}/* debug [instance_properties/getter]: contentOverlayView */


// The player view’s controls style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/controlsStyle
func (p_ PlayerView) ControlsStyle() PlayerViewControlsStyle {
	rv := objc.Send[PlayerViewControlsStyle](p_.ID, objc.Sel("controlsStyle"))
	return rv
}/* debug [instance_properties/getter]: controlsStyle */


// The player view’s controls style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/controlsStyle
func (p_ PlayerView) SetControlsStyle(value PlayerViewControlsStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlsStyle:"), value)
}/* debug [instance_properties/setter]: controlsStyle */


// The player view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/delegate
func (p_ PlayerView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The player view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/delegate
func (p_ PlayerView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the current player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/isReadyForDisplay
func (p_ PlayerView) ReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readyForDisplay"))
	return rv
}/* debug [instance_properties/getter]: readyForDisplay */


// The factor by which the video’s view is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/magnification
func (p_ PlayerView) Magnification() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("magnification"))
	return rv
}/* debug [instance_properties/getter]: magnification */


// The factor by which the video’s view is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/magnification
func (p_ PlayerView) SetMagnification(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMagnification:"), value)
}/* debug [instance_properties/setter]: magnification */


// The Picture in Picture delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/pictureInPictureDelegate
func (p_ PlayerView) PictureInPictureDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pictureInPictureDelegate"))
	return rv
}/* debug [instance_properties/getter]: pictureInPictureDelegate */


// The Picture in Picture delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/pictureInPictureDelegate
func (p_ PlayerView) SetPictureInPictureDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPictureInPictureDelegate:"), value)
}/* debug [instance_properties/setter]: pictureInPictureDelegate */


// The player instance that provides the media content for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/player
func (p_ PlayerView) Player() objc.IObject /* cross-framework: Player */ {
	rv := objc.Send[appkit.Player](p_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The player instance that provides the media content for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/player
func (p_ PlayerView) SetPlayer(value objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}/* debug [instance_properties/setter]: player */


// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/preferredDisplayDynamicRange
func (p_ PlayerView) PreferredDisplayDynamicRange() DisplayDynamicRange {
	rv := objc.Send[DisplayDynamicRange](p_.ID, objc.Sel("preferredDisplayDynamicRange"))
	return rv
}/* debug [instance_properties/getter]: preferredDisplayDynamicRange */


// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/preferredDisplayDynamicRange
func (p_ PlayerView) SetPreferredDisplayDynamicRange(value DisplayDynamicRange) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredDisplayDynamicRange:"), value)
}/* debug [instance_properties/setter]: preferredDisplayDynamicRange */


// The currently selected playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/selectedSpeed
func (p_ PlayerView) SelectedSpeed() IAVPlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}/* debug [instance_properties/getter]: selectedSpeed */


// A Boolean value that determines whether the player view displays frame stepping buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFrameSteppingButtons
func (p_ PlayerView) ShowsFrameSteppingButtons() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsFrameSteppingButtons"))
	return rv
}/* debug [instance_properties/getter]: showsFrameSteppingButtons */


// A Boolean value that determines whether the player view displays frame stepping buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFrameSteppingButtons
func (p_ PlayerView) SetShowsFrameSteppingButtons(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsFrameSteppingButtons:"), value)
}/* debug [instance_properties/setter]: showsFrameSteppingButtons */


// A Boolean value that determines whether the player view displays a full-screen toggle button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFullScreenToggleButton
func (p_ PlayerView) ShowsFullScreenToggleButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsFullScreenToggleButton"))
	return rv
}/* debug [instance_properties/getter]: showsFullScreenToggleButton */


// A Boolean value that determines whether the player view displays a full-screen toggle button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFullScreenToggleButton
func (p_ PlayerView) SetShowsFullScreenToggleButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsFullScreenToggleButton:"), value)
}/* debug [instance_properties/setter]: showsFullScreenToggleButton */


// A Boolean value that determines whether the player view displays a sharing service button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsSharingServiceButton
func (p_ PlayerView) ShowsSharingServiceButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsSharingServiceButton"))
	return rv
}/* debug [instance_properties/getter]: showsSharingServiceButton */


// A Boolean value that determines whether the player view displays a sharing service button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsSharingServiceButton
func (p_ PlayerView) SetShowsSharingServiceButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsSharingServiceButton:"), value)
}/* debug [instance_properties/setter]: showsSharingServiceButton */


// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsTimecodes
func (p_ PlayerView) ShowsTimecodes() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsTimecodes"))
	return rv
}/* debug [instance_properties/getter]: showsTimecodes */


// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsTimecodes
func (p_ PlayerView) SetShowsTimecodes(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsTimecodes:"), value)
}/* debug [instance_properties/setter]: showsTimecodes */


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/speeds
func (p_ PlayerView) Speeds() []PlaybackSpeed {
	rv := objc.Send[[]PlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}/* debug [instance_properties/getter]: speeds */


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/speeds
func (p_ PlayerView) SetSpeeds(value []PlaybackSpeed) {
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
}/* debug [instance_properties/setter]: speeds */


// A Boolean value that indicates whether the player view controller updates the Now Playing info center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/updatesNowPlayingInfoCenter
func (p_ PlayerView) UpdatesNowPlayingInfoCenter() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("updatesNowPlayingInfoCenter"))
	return rv
}/* debug [instance_properties/getter]: updatesNowPlayingInfoCenter */


// A Boolean value that indicates whether the player view controller updates the Now Playing info center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/updatesNowPlayingInfoCenter
func (p_ PlayerView) SetUpdatesNowPlayingInfoCenter(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpdatesNowPlayingInfoCenter:"), value)
}/* debug [instance_properties/setter]: updatesNowPlayingInfoCenter */


// The current size and position of the video image that displays within the player view’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoBounds
func (p_ PlayerView) VideoBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("videoBounds"))
	return rv
}/* debug [instance_properties/getter]: videoBounds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoFrameAnalysisTypes
func (p_ PlayerView) VideoFrameAnalysisTypes() VideoFrameAnalysisType {
	rv := objc.Send[VideoFrameAnalysisType](p_.ID, objc.Sel("videoFrameAnalysisTypes"))
	return rv
}/* debug [instance_properties/getter]: videoFrameAnalysisTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoFrameAnalysisTypes
func (p_ PlayerView) SetVideoFrameAnalysisTypes(value VideoFrameAnalysisType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFrameAnalysisTypes:"), value)
}/* debug [instance_properties/setter]: videoFrameAnalysisTypes */


// A value that determines how the player view displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoGravity
func (p_ PlayerView) VideoGravity() LayerVideoGravity /* not a class type */ {
	rv := objc.Send[LayerVideoGravity](p_.ID, objc.Sel("videoGravity"))
	return rv
}/* debug [instance_properties/getter]: videoGravity */


// A value that determines how the player view displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoGravity
func (p_ PlayerView) SetVideoGravity(value LayerVideoGravity /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoGravity:"), value)
}/* debug [instance_properties/setter]: videoGravity */


// A Boolean value that indicates whether the current player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/isreadyfordisplay
func (p_ PlayerView) IsReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}/* debug [instance_properties/getter]: isReadyForDisplay */


// A Boolean value that indicates whether the current player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/isreadyfordisplay
func (p_ PlayerView) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}/* debug [instance_properties/setter]: isReadyForDisplay */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerView */



