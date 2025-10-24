// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RemoteCommandCenter] class.
var (
	RemoteCommandCenterClass     _RemoteCommandCenterClass
	RemoteCommandCenterClassOnce sync.Once
)

func getRemoteCommandCenterClass() _RemoteCommandCenterClass {
	RemoteCommandCenterClassOnce.Do(func() {
		RemoteCommandCenterClass = _RemoteCommandCenterClass{objc.GetClass("MPRemoteCommandCenter")}
	})
	return RemoteCommandCenterClass
}

type _RemoteCommandCenterClass struct {
	class objc.Class
}

// An interface definition for the [RemoteCommandCenter] class.
type IRemoteCommandCenter interface {
	objectivec.IObject
	// properties:
	BookmarkCommand() IMPFeedbackCommand
	ChangePlaybackPositionCommand() IMPChangePlaybackPositionCommand
	ChangeRepeatModeCommand() objc.IObject /* cross-framework: ChangeRepeatModeCommand */
	ChangeShuffleModeCommand() objc.IObject /* cross-framework: ChangeShuffleModeCommand */
	DisableLanguageOptionCommand() IMPRemoteCommand
	NextTrackCommand() IMPRemoteCommand
	PauseCommand() IMPRemoteCommand
	SeekBackwardCommand() IMPRemoteCommand
	SeekForwardCommand() IMPRemoteCommand
	StopCommand() IMPRemoteCommand
	ChangePlaybackRateCommand() objc.IObject /* cross-framework: ChangePlaybackRateCommand */
	SetChangePlaybackRateCommand(value objc.IObject /* cross-framework: ChangePlaybackRateCommand */)
	DislikeCommand() IMPFeedbackCommand
	SetDislikeCommand(value IMPFeedbackCommand)
	EnableLanguageOptionCommand() IMPRemoteCommand
	SetEnableLanguageOptionCommand(value IMPRemoteCommand)
	LikeCommand() IMPFeedbackCommand
	SetLikeCommand(value IMPFeedbackCommand)
	PlayCommand() IMPRemoteCommand
	SetPlayCommand(value IMPRemoteCommand)
	PreviousTrackCommand() IMPRemoteCommand
	SetPreviousTrackCommand(value IMPRemoteCommand)
	RatingCommand() IMPRatingCommand
	SetRatingCommand(value IMPRatingCommand)
	SkipBackwardCommand() IMPSkipIntervalCommand
	SetSkipBackwardCommand(value IMPSkipIntervalCommand)
	SkipForwardCommand() IMPSkipIntervalCommand
	SetSkipForwardCommand(value IMPSkipIntervalCommand)
	TogglePlayPauseCommand() IMPRemoteCommand
	SetTogglePlayPauseCommand(value IMPRemoteCommand)
	// methods:
}

// An object that responds to remote control events sent by external accessories and system controls.
//
// Don’t create instances of this class yourself. Instead, use the method to retrieve the shared command center object. The properties of the shared command center object contain objects that respond to the various kinds of remote control events. You configure these objects to respond to the events you’re interested to handle in your app.


// An object that responds to remote control events sent by external accessories and system controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter
type RemoteCommandCenter struct {
	objectivec.Object
}

// RemoteCommandCenterFrom constructs a [RemoteCommandCenter] from an unsafe.Pointer.
//
// An object that responds to remote control events sent by external accessories and system controls.
func RemoteCommandCenterFrom(ptr unsafe.Pointer) RemoteCommandCenter {
	return RemoteCommandCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandCenterClass) Alloc() RemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RemoteCommandCenterClass) New() RemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RemoteCommandCenter) Init() RemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RemoteCommandCenter) Autorelease() RemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemoteCommandCenter creates a new RemoteCommandCenter instance.
func NewRemoteCommandCenter() RemoteCommandCenter {
	return getRemoteCommandCenterClass().New()
}



// Returns the shared object you use to access the system’s remote command objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/shared()
func (rc _RemoteCommandCenterClass) SharedCommandCenter() IRemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](objc.ID(rc.class), objc.Sel("sharedCommandCenter"))
	return rv
}


// The command object for indicating that a user wants to remember a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/bookmarkCommand
func (r_ RemoteCommandCenter) BookmarkCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("bookmarkCommand"))
	return rv
}


// The command object for changing the playback position in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changePlaybackPositionCommand
func (r_ RemoteCommandCenter) ChangePlaybackPositionCommand() IMPChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](r_.ID, objc.Sel("changePlaybackPositionCommand"))
	return rv
}


// The command object for changing the repeat mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeRepeatModeCommand
func (r_ RemoteCommandCenter) ChangeRepeatModeCommand() objc.IObject /* cross-framework: ChangeRepeatModeCommand */ {
	rv := objc.Send[ChangeRepeatModeCommand](r_.ID, objc.Sel("changeRepeatModeCommand"))
	return rv
}


// The command object for changing the shuffle mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeShuffleModeCommand
func (r_ RemoteCommandCenter) ChangeShuffleModeCommand() objc.IObject /* cross-framework: ChangeShuffleModeCommand */ {
	rv := objc.Send[ChangeShuffleModeCommand](r_.ID, objc.Sel("changeShuffleModeCommand"))
	return rv
}


// The command object for disabling a language option
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/disableLanguageOptionCommand
func (r_ RemoteCommandCenter) DisableLanguageOptionCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("disableLanguageOptionCommand"))
	return rv
}


// The command object for selecting the next track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/nextTrackCommand
func (r_ RemoteCommandCenter) NextTrackCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("nextTrackCommand"))
	return rv
}


// The command object for pausing playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/pauseCommand
func (r_ RemoteCommandCenter) PauseCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("pauseCommand"))
	return rv
}


// The command object for seeking backward through a single media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekBackwardCommand
func (r_ RemoteCommandCenter) SeekBackwardCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("seekBackwardCommand"))
	return rv
}


// The command object for seeking forward through a single media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekForwardCommand
func (r_ RemoteCommandCenter) SeekForwardCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("seekForwardCommand"))
	return rv
}


// The command object for stopping playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/stopCommand
func (r_ RemoteCommandCenter) StopCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("stopCommand"))
	return rv
}


// The command object for changing the playback rate of the current media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/changeplaybackratecommand
func (r_ RemoteCommandCenter) ChangePlaybackRateCommand() objc.IObject /* cross-framework: ChangePlaybackRateCommand */ {
	rv := objc.Send[ChangePlaybackRateCommand](r_.ID, objc.Sel("changePlaybackRateCommand"))
	return rv
}


// The command object for changing the playback rate of the current media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/changeplaybackratecommand
func (r_ RemoteCommandCenter) SetChangePlaybackRateCommand(value objc.IObject /* cross-framework: ChangePlaybackRateCommand */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setChangePlaybackRateCommand:"), value)
}


// The command object for indicating that a user dislikes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/dislikecommand
func (r_ RemoteCommandCenter) DislikeCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("dislikeCommand"))
	return rv
}


// The command object for indicating that a user dislikes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/dislikecommand
func (r_ RemoteCommandCenter) SetDislikeCommand(value IMPFeedbackCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDislikeCommand:"), value)
}


// The command object for enabling a language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/enablelanguageoptioncommand
func (r_ RemoteCommandCenter) EnableLanguageOptionCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("enableLanguageOptionCommand"))
	return rv
}


// The command object for enabling a language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/enablelanguageoptioncommand
func (r_ RemoteCommandCenter) SetEnableLanguageOptionCommand(value IMPRemoteCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEnableLanguageOptionCommand:"), value)
}


// The command object for indicating that a user likes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/likecommand
func (r_ RemoteCommandCenter) LikeCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("likeCommand"))
	return rv
}


// The command object for indicating that a user likes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/likecommand
func (r_ RemoteCommandCenter) SetLikeCommand(value IMPFeedbackCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLikeCommand:"), value)
}


// The command object for starting playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/playcommand
func (r_ RemoteCommandCenter) PlayCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("playCommand"))
	return rv
}


// The command object for starting playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/playcommand
func (r_ RemoteCommandCenter) SetPlayCommand(value IMPRemoteCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPlayCommand:"), value)
}


// The command object for selecting the previous track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/previoustrackcommand
func (r_ RemoteCommandCenter) PreviousTrackCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("previousTrackCommand"))
	return rv
}


// The command object for selecting the previous track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/previoustrackcommand
func (r_ RemoteCommandCenter) SetPreviousTrackCommand(value IMPRemoteCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreviousTrackCommand:"), value)
}


// The command object for rating a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/ratingcommand
func (r_ RemoteCommandCenter) RatingCommand() IMPRatingCommand {
	rv := objc.Send[RatingCommand](r_.ID, objc.Sel("ratingCommand"))
	return rv
}


// The command object for rating a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/ratingcommand
func (r_ RemoteCommandCenter) SetRatingCommand(value IMPRatingCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRatingCommand:"), value)
}


// The command object for playing a previous point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/skipbackwardcommand
func (r_ RemoteCommandCenter) SkipBackwardCommand() IMPSkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](r_.ID, objc.Sel("skipBackwardCommand"))
	return rv
}


// The command object for playing a previous point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/skipbackwardcommand
func (r_ RemoteCommandCenter) SetSkipBackwardCommand(value IMPSkipIntervalCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSkipBackwardCommand:"), value)
}


// The command object for playing a future point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/skipforwardcommand
func (r_ RemoteCommandCenter) SkipForwardCommand() IMPSkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](r_.ID, objc.Sel("skipForwardCommand"))
	return rv
}


// The command object for playing a future point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/skipforwardcommand
func (r_ RemoteCommandCenter) SetSkipForwardCommand(value IMPSkipIntervalCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSkipForwardCommand:"), value)
}


// The command object for toggling between playing and pausing the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/toggleplaypausecommand
func (r_ RemoteCommandCenter) TogglePlayPauseCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("togglePlayPauseCommand"))
	return rv
}


// The command object for toggling between playing and pausing the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandcenter/toggleplaypausecommand
func (r_ RemoteCommandCenter) SetTogglePlayPauseCommand(value IMPRemoteCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTogglePlayPauseCommand:"), value)
}



