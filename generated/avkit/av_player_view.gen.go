// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [PlayerView] class.
type IPlayerView interface {
	appkit.IView
	BeginTrimmingWithCompletionHandler(handler unsafe.Pointer)
	FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle string)
	SelectSpeed(speed unsafe.Pointer)
	SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint)
}

// A view that displays content from a player and presents a native user interface to control playback.
//
// The player view supports several controls styles, ranging from no controls to controls matching the look of QuickTime Player. This makes it easy for you to tailor the presentation to best match your use of the player view. Regardless of the selected controls style, the player view always supports the following standard set of keyboard shortcuts to control playback: The Space bar plays and pauses playback. The right and left arrow keys step frame-by-frame through the video. JKL navigation: The J key rewinds. Press it multiple times to cycle through rewind speeds. The K key stops playback. The L key fast-forwards. Press it multiple times to cycle through fast-forward speeds. The player view also makes it simple to add trimming capabilities to your player. Call the view’s method to present a trimming UI that matches the QuickTime Player interface.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PlayerViewClass) Alloc() PlayerView {
	rv := objc.Send[PlayerView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Puts the player view into trimming mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/beginTrimming(completionHandler:)
func (p_ PlayerView) BeginTrimmingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginTrimmingWithCompletionHandler:"), handler)
}

// Displays the chapter number and title in the player view for a brief moment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/flashChapterNumber(_:chapterTitle:)
func (p_ PlayerView) FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("flashChapterNumber:chapterTitle:"), chapterNumber, objc.String(chapterTitle))
}

// Selects a specified playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/selectSpeed(_:)
func (p_ PlayerView) SelectSpeed(speed unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectSpeed:"), speed)
}

// Scales the video’s view by a specified factor, and centers the result on a specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/setMagnification(_:centeredAt:)
func (p_ PlayerView) SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// An action pop-up button menu that the player view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/actionPopUpButtonMenu
func (p_ PlayerView) ActionPopUpButtonMenu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("actionPopUpButtonMenu"))
	return rv
}


// SetActionPopUpButtonMenu sets the value of the actionPopUpButtonMenu property.
// An action pop-up button menu that the player view displays.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/actionPopUpButtonMenu
func (p_ PlayerView) SetActionPopUpButtonMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActionPopUpButtonMenu:"), value)
}
// A Boolean value that indicates whether the magnify gesture changes the video’s view magnification.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsMagnification
func (p_ PlayerView) AllowsMagnification() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsMagnification"))
	return rv
}


// SetAllowsMagnification sets the value of the allowsMagnification property.
// A Boolean value that indicates whether the magnify gesture changes the video’s view magnification.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsMagnification
func (p_ PlayerView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsMagnification:"), value)
}
// A Boolean value that determines whether the player view allows Picture in Picture playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsPictureInPicturePlayback
func (p_ PlayerView) AllowsPictureInPicturePlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsPictureInPicturePlayback"))
	return rv
}


// SetAllowsPictureInPicturePlayback sets the value of the allowsPictureInPicturePlayback property.
// A Boolean value that determines whether the player view allows Picture in Picture playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsPictureInPicturePlayback
func (p_ PlayerView) SetAllowsPictureInPicturePlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsPictureInPicturePlayback:"), value)
}
// A Boolean value that indicates whether to perform video frame analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsVideoFrameAnalysis
func (p_ PlayerView) AllowsVideoFrameAnalysis() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsVideoFrameAnalysis"))
	return rv
}


// SetAllowsVideoFrameAnalysis sets the value of the allowsVideoFrameAnalysis property.
// A Boolean value that indicates whether to perform video frame analysis.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/allowsVideoFrameAnalysis
func (p_ PlayerView) SetAllowsVideoFrameAnalysis(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsVideoFrameAnalysis:"), value)
}
// A Boolean value that indicates whether the player view can begin trimming.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/canBeginTrimming
func (p_ PlayerView) CanBeginTrimming() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canBeginTrimming"))
	return rv
}

// A view that adds additional custom views between the video content and the controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/contentOverlayView
func (p_ PlayerView) ContentOverlayView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentOverlayView"))
	return rv
}

// The player view’s controls style.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/controlsStyle
func (p_ PlayerView) ControlsStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("controlsStyle"))
	return rv
}


// SetControlsStyle sets the value of the controlsStyle property.
// The player view’s controls style.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/controlsStyle
func (p_ PlayerView) SetControlsStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlsStyle:"), value)
}
// The player view’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/delegate
func (p_ PlayerView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The player view’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/delegate
func (p_ PlayerView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that indicates whether the current player item’s first video frame is ready for display.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/isReadyForDisplay
func (p_ PlayerView) ReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readyForDisplay"))
	return rv
}

// The factor by which the video’s view is currently scaled.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/magnification
func (p_ PlayerView) Magnification() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("magnification"))
	return rv
}


// SetMagnification sets the value of the magnification property.
// The factor by which the video’s view is currently scaled.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/magnification
func (p_ PlayerView) SetMagnification(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMagnification:"), value)
}
// The Picture in Picture delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/pictureInPictureDelegate
func (p_ PlayerView) PictureInPictureDelegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("pictureInPictureDelegate"))
	return rv
}


// SetPictureInPictureDelegate sets the value of the pictureInPictureDelegate property.
// The Picture in Picture delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/pictureInPictureDelegate
func (p_ PlayerView) SetPictureInPictureDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPictureInPictureDelegate:"), value)
}
// The player instance that provides the media content for the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/player
func (p_ PlayerView) Player() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("player"))
	return rv
}


// SetPlayer sets the value of the player property.
// The player instance that provides the media content for the view.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/player
func (p_ PlayerView) SetPlayer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}
// Describes how High Dynamic Range (HDR) video content renders.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/preferredDisplayDynamicRange
func (p_ PlayerView) PreferredDisplayDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredDisplayDynamicRange"))
	return rv
}


// SetPreferredDisplayDynamicRange sets the value of the preferredDisplayDynamicRange property.
// Describes how High Dynamic Range (HDR) video content renders.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/preferredDisplayDynamicRange
func (p_ PlayerView) SetPreferredDisplayDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredDisplayDynamicRange:"), value)
}
// The currently selected playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/selectedSpeed
func (p_ PlayerView) SelectedSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}

// A Boolean value that determines whether the player view displays frame stepping buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFrameSteppingButtons
func (p_ PlayerView) ShowsFrameSteppingButtons() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsFrameSteppingButtons"))
	return rv
}


// SetShowsFrameSteppingButtons sets the value of the showsFrameSteppingButtons property.
// A Boolean value that determines whether the player view displays frame stepping buttons.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFrameSteppingButtons
func (p_ PlayerView) SetShowsFrameSteppingButtons(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsFrameSteppingButtons:"), value)
}
// A Boolean value that determines whether the player view displays a full-screen toggle button.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFullScreenToggleButton
func (p_ PlayerView) ShowsFullScreenToggleButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsFullScreenToggleButton"))
	return rv
}


// SetShowsFullScreenToggleButton sets the value of the showsFullScreenToggleButton property.
// A Boolean value that determines whether the player view displays a full-screen toggle button.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsFullScreenToggleButton
func (p_ PlayerView) SetShowsFullScreenToggleButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsFullScreenToggleButton:"), value)
}
// A Boolean value that determines whether the player view displays a sharing service button.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsSharingServiceButton
func (p_ PlayerView) ShowsSharingServiceButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsSharingServiceButton"))
	return rv
}


// SetShowsSharingServiceButton sets the value of the showsSharingServiceButton property.
// A Boolean value that determines whether the player view displays a sharing service button.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsSharingServiceButton
func (p_ PlayerView) SetShowsSharingServiceButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsSharingServiceButton:"), value)
}
// A Boolean value that determines whether the player view displays timecodes, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsTimecodes
func (p_ PlayerView) ShowsTimecodes() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsTimecodes"))
	return rv
}


// SetShowsTimecodes sets the value of the showsTimecodes property.
// A Boolean value that determines whether the player view displays timecodes, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/showsTimecodes
func (p_ PlayerView) SetShowsTimecodes(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsTimecodes:"), value)
}
// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/speeds
func (p_ PlayerView) Speeds() []PlaybackSpeed {
	rv := objc.Send[[]PlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}


// SetSpeeds sets the value of the speeds property.
// A list of user-selectable playback speeds to show in the playback speed control.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/speeds
func (p_ PlayerView) SetSpeeds(value []PlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpeeds:"), value)
}
// A Boolean value that indicates whether the player view controller updates the Now Playing info center.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/updatesNowPlayingInfoCenter
func (p_ PlayerView) UpdatesNowPlayingInfoCenter() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("updatesNowPlayingInfoCenter"))
	return rv
}


// SetUpdatesNowPlayingInfoCenter sets the value of the updatesNowPlayingInfoCenter property.
// A Boolean value that indicates whether the player view controller updates the Now Playing info center.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/updatesNowPlayingInfoCenter
func (p_ PlayerView) SetUpdatesNowPlayingInfoCenter(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpdatesNowPlayingInfoCenter:"), value)
}
// The current size and position of the video image that displays within the player view’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoBounds
func (p_ PlayerView) VideoBounds() Rect {
	rv := objc.Send[Rect](p_.ID, objc.Sel("videoBounds"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoFrameAnalysisTypes
func (p_ PlayerView) VideoFrameAnalysisTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoFrameAnalysisTypes"))
	return rv
}


// SetVideoFrameAnalysisTypes sets the value of the videoFrameAnalysisTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoFrameAnalysisTypes
func (p_ PlayerView) SetVideoFrameAnalysisTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFrameAnalysisTypes:"), value)
}
// A value that determines how the player view displays video content within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoGravity
func (p_ PlayerView) VideoGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoGravity"))
	return rv
}


// SetVideoGravity sets the value of the videoGravity property.
// A value that determines how the player view displays video content within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerView/videoGravity
func (p_ PlayerView) SetVideoGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoGravity:"), value)
}


