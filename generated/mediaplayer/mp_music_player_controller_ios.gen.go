//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MusicPlayerController


// Inserts the media items defined by the queue descriptor after the last media item in the current queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/append(_:)
func (m_ MusicPlayerController) AppendQueueDescriptor(descriptor IMPMusicPlayerQueueDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendQueueDescriptor:"), descriptor)
}

// Starts the generation of playback notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/beginGeneratingPlaybackNotifications()
func (m_ MusicPlayerController) BeginGeneratingPlaybackNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginGeneratingPlaybackNotifications"))
}

// Ends the generation of playback notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/endGeneratingPlaybackNotifications()
func (m_ MusicPlayerController) EndGeneratingPlaybackNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endGeneratingPlaybackNotifications"))
}

// Prepares a music player for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/prepareToPlay(completionHandler:)
func (m_ MusicPlayerController) PrepareToPlayWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("prepareToPlayWithCompletionHandler:"), completionHandler)
}

// Inserts the media items defined by the queue descriptor into the current queue immediately after the currently playing media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/prepend(_:)
func (m_ MusicPlayerController) PrependQueueDescriptor(descriptor IMPMusicPlayerQueueDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("prependQueueDescriptor:"), descriptor)
}

// Set the music player’s playback queue using media items that fit the queue descriptor properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-1izmj
func (m_ MusicPlayerController) SetQueueWithDescriptor(descriptor IMPMusicPlayerQueueDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithDescriptor:"), descriptor)
}

// Sets a music player’s playback queue based on a media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-5rii3
func (m_ MusicPlayerController) SetQueueWithQuery(query IMPMediaQuery) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithQuery:"), query)
}

// Sets a music player’s playback queue using with media items identified by the store identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-8x6xb
func (m_ MusicPlayerController) SetQueueWithStoreIDs(storeIDs []string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithStoreIDs:"), storeIDs)
}

// Sets a music player’s playback queue using a media item collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/setQueue(with:)-xlwk
func (m_ MusicPlayerController) SetQueueWithItemCollection(itemCollection IMPMediaItemCollection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueueWithItemCollection:"), itemCollection)
}

// Restarts playback at the beginning of the currently playing media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToBeginning()
func (m_ MusicPlayerController) SkipToBeginning() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToBeginning"))
}

// Starts playback of the next media item in the playback queue, or if the music player isn’t playing, designates the next media item as the next item to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToNextItem()
func (m_ MusicPlayerController) SkipToNextItem() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToNextItem"))
}

// Starts playback of the previous media item in the playback queue, or if the music player isn’t playing, designates the previous media item as the next to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/skipToPreviousItem()
func (m_ MusicPlayerController) SkipToPreviousItem() {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipToPreviousItem"))
}

// iOS-only properties

// The index of the now playing item in the current playback queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/indexOfNowPlayingItem
func (m_ MusicPlayerController) IndexOfNowPlayingItem() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("indexOfNowPlayingItem"))
	return rv
}

// The currently-playing media item, or the media item in a queue that you designated to begin playback with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/nowPlayingItem
func (m_ MusicPlayerController) NowPlayingItem() IMPMediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("nowPlayingItem"))
	return rv
}
func (m_ MusicPlayerController) SetNowPlayingItem(value IMPMediaItem) {
	m_.ID.Send(objc.RegisterName("setNowPlayingItem:"), value)
}

// The current playback state of the music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/playbackState
func (m_ MusicPlayerController) PlaybackState() MusicPlaybackState {
	rv := objc.Send[MusicPlaybackState](m_.ID, objc.Sel("playbackState"))
	return rv
}

// The current repeat mode of the music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/repeatMode
func (m_ MusicPlayerController) RepeatMode() MusicRepeatMode {
	rv := objc.Send[MusicRepeatMode](m_.ID, objc.Sel("repeatMode"))
	return rv
}
func (m_ MusicPlayerController) SetRepeatMode(value MusicRepeatMode) {
	m_.ID.Send(objc.RegisterName("setRepeatMode:"), value)
}

// The current shuffle mode of the music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/shuffleMode
func (m_ MusicPlayerController) ShuffleMode() MusicShuffleMode {
	rv := objc.Send[MusicShuffleMode](m_.ID, objc.Sel("shuffleMode"))
	return rv
}
func (m_ MusicPlayerController) SetShuffleMode(value MusicShuffleMode) {
	m_.ID.Send(objc.RegisterName("setShuffleMode:"), value)
}

// The audio playback volume for the music player, in the range from (silent) through (maximum volume).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/volume
func (m_ MusicPlayerController) Volume() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("volume"))
	return rv
}
func (m_ MusicPlayerController) SetVolume(value float32) {
	m_.ID.Send(objc.RegisterName("setVolume:"), value)
}





