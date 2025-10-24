// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPRemoteCommandCenter */


/* debug [class_header]: Header for MPRemoteCommandCenter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RemoteCommandCenter */
// An interface definition for the [RemoteCommandCenter] class.
type IRemoteCommandCenter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RemoteCommandCenter */
	// properties:
	BookmarkCommand() IMPFeedbackCommand
	ChangePlaybackPositionCommand() IMPChangePlaybackPositionCommand
	ChangePlaybackRateCommand() IMPChangePlaybackRateCommand
	ChangeRepeatModeCommand() IMPChangeRepeatModeCommand
	ChangeShuffleModeCommand() IMPChangeShuffleModeCommand
	DisableLanguageOptionCommand() IMPRemoteCommand
	DislikeCommand() IMPFeedbackCommand
	EnableLanguageOptionCommand() IMPRemoteCommand
	LikeCommand() IMPFeedbackCommand
	NextTrackCommand() IMPRemoteCommand
	PauseCommand() IMPRemoteCommand
	PlayCommand() IMPRemoteCommand
	PreviousTrackCommand() IMPRemoteCommand
	RatingCommand() IMPRatingCommand
	SeekBackwardCommand() IMPRemoteCommand
	SeekForwardCommand() IMPRemoteCommand
	SkipBackwardCommand() IMPSkipIntervalCommand
	SkipForwardCommand() IMPSkipIntervalCommand
	StopCommand() IMPRemoteCommand
	TogglePlayPauseCommand() IMPRemoteCommand
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RemoteCommandCenter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RemoteCommandCenter */
// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandCenterClass) Alloc() RemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RemoteCommandCenter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RemoteCommandCenter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RemoteCommandCenter */

// Returns the shared object you use to access the system’s remote command objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/shared()
func (rc _RemoteCommandCenterClass) SharedCommandCenter() IRemoteCommandCenter {
	rv := objc.Send[RemoteCommandCenter](objc.ID(rc.class), objc.Sel("sharedCommandCenter"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCommandCenter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RemoteCommandCenter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RemoteCommandCenter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RemoteCommandCenter */

// The command object for indicating that a user wants to remember a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/bookmarkCommand
func (r_ RemoteCommandCenter) BookmarkCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("bookmarkCommand"))
	return rv
}/* debug [instance_properties/getter]: bookmarkCommand */


// The command object for changing the playback position in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changePlaybackPositionCommand
func (r_ RemoteCommandCenter) ChangePlaybackPositionCommand() IMPChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](r_.ID, objc.Sel("changePlaybackPositionCommand"))
	return rv
}/* debug [instance_properties/getter]: changePlaybackPositionCommand */


// The command object for changing the playback rate of the current media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changePlaybackRateCommand
func (r_ RemoteCommandCenter) ChangePlaybackRateCommand() IMPChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](r_.ID, objc.Sel("changePlaybackRateCommand"))
	return rv
}/* debug [instance_properties/getter]: changePlaybackRateCommand */


// The command object for changing the repeat mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeRepeatModeCommand
func (r_ RemoteCommandCenter) ChangeRepeatModeCommand() IMPChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](r_.ID, objc.Sel("changeRepeatModeCommand"))
	return rv
}/* debug [instance_properties/getter]: changeRepeatModeCommand */


// The command object for changing the shuffle mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeShuffleModeCommand
func (r_ RemoteCommandCenter) ChangeShuffleModeCommand() IMPChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](r_.ID, objc.Sel("changeShuffleModeCommand"))
	return rv
}/* debug [instance_properties/getter]: changeShuffleModeCommand */


// The command object for disabling a language option
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/disableLanguageOptionCommand
func (r_ RemoteCommandCenter) DisableLanguageOptionCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("disableLanguageOptionCommand"))
	return rv
}/* debug [instance_properties/getter]: disableLanguageOptionCommand */


// The command object for indicating that a user dislikes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/dislikeCommand
func (r_ RemoteCommandCenter) DislikeCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("dislikeCommand"))
	return rv
}/* debug [instance_properties/getter]: dislikeCommand */


// The command object for enabling a language option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/enableLanguageOptionCommand
func (r_ RemoteCommandCenter) EnableLanguageOptionCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("enableLanguageOptionCommand"))
	return rv
}/* debug [instance_properties/getter]: enableLanguageOptionCommand */


// The command object for indicating that a user likes what is currently playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/likeCommand
func (r_ RemoteCommandCenter) LikeCommand() IMPFeedbackCommand {
	rv := objc.Send[FeedbackCommand](r_.ID, objc.Sel("likeCommand"))
	return rv
}/* debug [instance_properties/getter]: likeCommand */


// The command object for selecting the next track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/nextTrackCommand
func (r_ RemoteCommandCenter) NextTrackCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("nextTrackCommand"))
	return rv
}/* debug [instance_properties/getter]: nextTrackCommand */


// The command object for pausing playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/pauseCommand
func (r_ RemoteCommandCenter) PauseCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("pauseCommand"))
	return rv
}/* debug [instance_properties/getter]: pauseCommand */


// The command object for starting playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/playCommand
func (r_ RemoteCommandCenter) PlayCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("playCommand"))
	return rv
}/* debug [instance_properties/getter]: playCommand */


// The command object for selecting the previous track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/previousTrackCommand
func (r_ RemoteCommandCenter) PreviousTrackCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("previousTrackCommand"))
	return rv
}/* debug [instance_properties/getter]: previousTrackCommand */


// The command object for rating a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/ratingCommand
func (r_ RemoteCommandCenter) RatingCommand() IMPRatingCommand {
	rv := objc.Send[RatingCommand](r_.ID, objc.Sel("ratingCommand"))
	return rv
}/* debug [instance_properties/getter]: ratingCommand */


// The command object for seeking backward through a single media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekBackwardCommand
func (r_ RemoteCommandCenter) SeekBackwardCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("seekBackwardCommand"))
	return rv
}/* debug [instance_properties/getter]: seekBackwardCommand */


// The command object for seeking forward through a single media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekForwardCommand
func (r_ RemoteCommandCenter) SeekForwardCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("seekForwardCommand"))
	return rv
}/* debug [instance_properties/getter]: seekForwardCommand */


// The command object for playing a previous point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/skipBackwardCommand
func (r_ RemoteCommandCenter) SkipBackwardCommand() IMPSkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](r_.ID, objc.Sel("skipBackwardCommand"))
	return rv
}/* debug [instance_properties/getter]: skipBackwardCommand */


// The command object for playing a future point in a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/skipForwardCommand
func (r_ RemoteCommandCenter) SkipForwardCommand() IMPSkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](r_.ID, objc.Sel("skipForwardCommand"))
	return rv
}/* debug [instance_properties/getter]: skipForwardCommand */


// The command object for stopping playback of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/stopCommand
func (r_ RemoteCommandCenter) StopCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("stopCommand"))
	return rv
}/* debug [instance_properties/getter]: stopCommand */


// The command object for toggling between playing and pausing the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/togglePlayPauseCommand
func (r_ RemoteCommandCenter) TogglePlayPauseCommand() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("togglePlayPauseCommand"))
	return rv
}/* debug [instance_properties/getter]: togglePlayPauseCommand */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPRemoteCommandCenter */



