// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object that responds to remote control events sent by external accessories and system controls.
//
// Don’t create instances of this class yourself. Instead, use the method to retrieve the shared command center object. The properties of the shared command center object contain objects that respond to the various kinds of remote control events. You configure these objects to respond to the events you’re interested to handle in your app.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/shared()
func (rc _RemoteCommandCenterClass) SharedCommandCenter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("sharedCommandCenter"))
	return rv
}

// The command object for indicating that a user wants to remember a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/bookmarkCommand
func (r_ RemoteCommandCenter) BookmarkCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("bookmarkCommand"))
	return rv
}

// The command object for changing the playback position in a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changePlaybackPositionCommand
func (r_ RemoteCommandCenter) ChangePlaybackPositionCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("changePlaybackPositionCommand"))
	return rv
}

// The command object for changing the playback rate of the current media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changePlaybackRateCommand
func (r_ RemoteCommandCenter) ChangePlaybackRateCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("changePlaybackRateCommand"))
	return rv
}

// The command object for changing the repeat mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeRepeatModeCommand
func (r_ RemoteCommandCenter) ChangeRepeatModeCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("changeRepeatModeCommand"))
	return rv
}

// The command object for changing the shuffle mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/changeShuffleModeCommand
func (r_ RemoteCommandCenter) ChangeShuffleModeCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("changeShuffleModeCommand"))
	return rv
}

// The command object for disabling a language option
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/disableLanguageOptionCommand
func (r_ RemoteCommandCenter) DisableLanguageOptionCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("disableLanguageOptionCommand"))
	return rv
}

// The command object for indicating that a user dislikes what is currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/dislikeCommand
func (r_ RemoteCommandCenter) DislikeCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("dislikeCommand"))
	return rv
}

// The command object for enabling a language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/enableLanguageOptionCommand
func (r_ RemoteCommandCenter) EnableLanguageOptionCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("enableLanguageOptionCommand"))
	return rv
}

// The command object for indicating that a user likes what is currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/likeCommand
func (r_ RemoteCommandCenter) LikeCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("likeCommand"))
	return rv
}

// The command object for selecting the next track.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/nextTrackCommand
func (r_ RemoteCommandCenter) NextTrackCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("nextTrackCommand"))
	return rv
}

// The command object for pausing playback of the current item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/pauseCommand
func (r_ RemoteCommandCenter) PauseCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("pauseCommand"))
	return rv
}

// The command object for starting playback of the current item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/playCommand
func (r_ RemoteCommandCenter) PlayCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("playCommand"))
	return rv
}

// The command object for selecting the previous track.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/previousTrackCommand
func (r_ RemoteCommandCenter) PreviousTrackCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("previousTrackCommand"))
	return rv
}

// The command object for rating a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/ratingCommand
func (r_ RemoteCommandCenter) RatingCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("ratingCommand"))
	return rv
}

// The command object for seeking backward through a single media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekBackwardCommand
func (r_ RemoteCommandCenter) SeekBackwardCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("seekBackwardCommand"))
	return rv
}

// The command object for seeking forward through a single media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/seekForwardCommand
func (r_ RemoteCommandCenter) SeekForwardCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("seekForwardCommand"))
	return rv
}

// The command object for playing a previous point in a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/skipBackwardCommand
func (r_ RemoteCommandCenter) SkipBackwardCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("skipBackwardCommand"))
	return rv
}

// The command object for playing a future point in a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/skipForwardCommand
func (r_ RemoteCommandCenter) SkipForwardCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("skipForwardCommand"))
	return rv
}

// The command object for stopping playback of the current item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/stopCommand
func (r_ RemoteCommandCenter) StopCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("stopCommand"))
	return rv
}

// The command object for toggling between playing and pausing the current item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandCenter/togglePlayPauseCommand
func (r_ RemoteCommandCenter) TogglePlayPauseCommand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("togglePlayPauseCommand"))
	return rv
}



