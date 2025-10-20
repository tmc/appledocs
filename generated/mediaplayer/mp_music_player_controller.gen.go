// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MusicPlayerController] class.
var (
	MusicPlayerControllerClass     _MusicPlayerControllerClass
	MusicPlayerControllerClassOnce sync.Once
)

func getMusicPlayerControllerClass() _MusicPlayerControllerClass {
	MusicPlayerControllerClassOnce.Do(func() {
		MusicPlayerControllerClass = _MusicPlayerControllerClass{objc.GetClass("MPMusicPlayerController")}
	})
	return MusicPlayerControllerClass
}

type _MusicPlayerControllerClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerController] class.
type IMusicPlayerController interface {
	objectivec.IObject
	AppendQueueDescriptor(descriptor unsafe.Pointer)
	BeginGeneratingPlaybackNotifications()
	EndGeneratingPlaybackNotifications()
	PrepareToPlayWithCompletionHandler(completionHandler unsafe.Pointer)
	PrependQueueDescriptor(descriptor unsafe.Pointer)
	SetQueueWithDescriptor(descriptor unsafe.Pointer)
	SetQueueWithQuery(query unsafe.Pointer)
	SetQueueWithStoreIDs(storeIDs unsafe.Pointer)
	SetQueueWithItemCollection(itemCollection unsafe.Pointer)
	SkipToBeginning()
	SkipToNextItem()
	SkipToPreviousItem()
}

// An object that plays audio media items from the device’s Music app library.
//
// Create an instance of a music player to play media items in your app. There are two types of music players: An plays music locally within your app. It isn’t aware of the Music app’s Now Playing item, nor does it affect the Music app’s state. There are two application music players: and . The application queue player provides greater control over the contents of the queue and is the preferred player. The employs the built-in Music app on your behalf. On instantiation, it takes on the current Music app state, such as the identification of the Now Playing item. If a user switches away from your app while music is playing, that music continues to play. The Music app then has your music player’s most recently-set repeat mode, shuffle mode, playback state, and Now Playing item. Creating a new instance of and not specifying the player type returns a system music player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController
type MusicPlayerController struct {
	objectivec.Object
}

// MusicPlayerControllerFrom constructs a [MusicPlayerController] from an unsafe.Pointer.
//
// An object that plays audio media items from the device’s Music app library.
func MusicPlayerControllerFrom(ptr unsafe.Pointer) MusicPlayerController {
	return MusicPlayerController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerClass) Alloc() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerControllerClass) New() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerController) Init() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerController) Autorelease() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerController creates a new MusicPlayerController instance.
func NewMusicPlayerController() MusicPlayerController {
	return getMusicPlayerControllerClass().New()
}


// Inserts the media items defined by the queue descriptor after the last media item in the current queue.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/append(_:)
func (m_ MusicPlayerController) AppendQueueDescriptor(descriptor unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendQueueDescriptor:"), descriptor)
}

// Starts the generation of playback notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/beginGeneratingPlaybackNotifications()
func (m_ MusicPlayerController) BeginGeneratingPlaybackNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginGeneratingPlaybackNotifications"))
}

// Ends the generation of playback notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/endGeneratingPlaybackNotifications()
func (m_ MusicPlayerController) EndGeneratingPlaybackNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endGeneratingPlaybackNotifications"))
}

// Prepares a music player for playback.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/prepareToPlay(completionHandler:)
func (m_ MusicPlayerController) PrepareToPlayWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("prepareToPlayWithCompletionHandler:"), completionHandler)
}

// Inserts the media items defined by the queue descriptor into the current queue immediately after the currently playing media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/prepend(_:)
func (m_ MusicPlayerController) PrependQueueDescriptor(descriptor unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("prependQueueDescriptor:"), descriptor)
}

// Set the music player’s playback queue using media items that fit the queue descriptor properties.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-1izmj
func (m_ MusicPlayerController) SetQueueWithDescriptor(descriptor unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithDescriptor:"), descriptor)
}

// Sets a music player’s playback queue based on a media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-5rii3
func (m_ MusicPlayerController) SetQueueWithQuery(query unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithQuery:"), query)
}

// Sets a music player’s playback queue using with media items identified by the store identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-8x6xb
func (m_ MusicPlayerController) SetQueueWithStoreIDs(storeIDs unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithStoreIDs:"), storeIDs)
}

// Sets a music player’s playback queue using a media item collection.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-xlwk
func (m_ MusicPlayerController) SetQueueWithItemCollection(itemCollection unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithItemCollection:"), itemCollection)
}

// Restarts playback at the beginning of the currently playing media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToBeginning()
func (m_ MusicPlayerController) SkipToBeginning() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToBeginning"))
}

// Starts playback of the next media item in the playback queue, or if the music player isn’t playing, designates the next media item as the next item to play.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToNextItem()
func (m_ MusicPlayerController) SkipToNextItem() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToNextItem"))
}

// Starts playback of the previous media item in the playback queue, or if the music player isn’t playing, designates the previous media item as the next to play.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToPreviousItem()
func (m_ MusicPlayerController) SkipToPreviousItem() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToPreviousItem"))
}

// The index of the now playing item in the current playback queue.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/indexOfNowPlayingItem
func (m_ MusicPlayerController) IndexOfNowPlayingItem() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("indexOfNowPlayingItem"))
	return rv
}

// The currently-playing media item, or the media item in a queue that you designated to begin playback with.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/nowPlayingItem
func (m_ MusicPlayerController) NowPlayingItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nowPlayingItem"))
	return rv
}


// SetNowPlayingItem sets the value of the nowPlayingItem property.
// The currently-playing media item, or the media item in a queue that you designated to begin playback with.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/nowPlayingItem
func (m_ MusicPlayerController) SetNowPlayingItem(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNowPlayingItem:"), value)
}
// The current playback state of the music player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/playbackState
func (m_ MusicPlayerController) PlaybackState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playbackState"))
	return rv
}

// The current repeat mode of the music player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/repeatMode
func (m_ MusicPlayerController) RepeatMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("repeatMode"))
	return rv
}


// SetRepeatMode sets the value of the repeatMode property.
// The current repeat mode of the music player.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/repeatMode
func (m_ MusicPlayerController) SetRepeatMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRepeatMode:"), value)
}
// The current shuffle mode of the music player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/shuffleMode
func (m_ MusicPlayerController) ShuffleMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("shuffleMode"))
	return rv
}


// SetShuffleMode sets the value of the shuffleMode property.
// The current shuffle mode of the music player.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/shuffleMode
func (m_ MusicPlayerController) SetShuffleMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShuffleMode:"), value)
}
// The audio playback volume for the music player, in the range from (silent) through (maximum volume).
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/volume
func (m_ MusicPlayerController) Volume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The audio playback volume for the music player, in the range from (silent) through (maximum volume).

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/volume
func (m_ MusicPlayerController) SetVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVolume:"), value)
}


