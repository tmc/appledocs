//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MediaItem


// iOS-only properties

// The primary performing artist for an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumArtist
func (m_ MediaItem) AlbumArtist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("albumArtist"))
	return rv
}

// The persistent identifier for the primary performing artist for an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumArtistPersistentID
func (m_ MediaItem) AlbumArtistPersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("albumArtistPersistentID"))
	return rv
}

// The persistent identifier for an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumPersistentID
func (m_ MediaItem) AlbumPersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("albumPersistentID"))
	return rv
}

// The title of an album, such as , rather than the title of an individual song on the album, such as “Crater Dance.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTitle
func (m_ MediaItem) AlbumTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("albumTitle"))
	return rv
}

// The number of tracks for the album that contains the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTrackCount
func (m_ MediaItem) AlbumTrackCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("albumTrackCount"))
	return rv
}

// The track number of the media item, for a media item that’s part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTrackNumber
func (m_ MediaItem) AlbumTrackNumber() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("albumTrackNumber"))
	return rv
}

// The performing artists for a media item, which may vary from the primary artist for the album that a media item belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artist
func (m_ MediaItem) Artist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("artist"))
	return rv
}

// The persistent identifier for an artist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artistPersistentID
func (m_ MediaItem) ArtistPersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("artistPersistentID"))
	return rv
}

// The artwork image for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artwork
func (m_ MediaItem) Artwork() IMPMediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](m_.ID, objc.Sel("artwork"))
	return rv
}

// The URL that points to the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/assetURL
func (m_ MediaItem) AssetURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("assetURL"))
	return rv
}

// The number of musical beats per minute for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/beatsPerMinute
func (m_ MediaItem) BeatsPerMinute() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("beatsPerMinute"))
	return rv
}

// The time of the user’s most recent interaction with the bookmark in the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/bookmarkTime
func (m_ MediaItem) BookmarkTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("bookmarkTime"))
	return rv
}

// Textual information about the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/comments
func (m_ MediaItem) Comments() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("comments"))
	return rv
}

// The persistent identifier for a composer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/composerPersistentID
func (m_ MediaItem) ComposerPersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("composerPersistentID"))
	return rv
}

// The date the user adds the media item to the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/dateAdded
func (m_ MediaItem) DateAdded() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("dateAdded"))
	return rv
}

// The number of discs for the album that contains the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/discCount
func (m_ MediaItem) DiscCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("discCount"))
	return rv
}

// The disc number of the media item, for a media item that’s part of a multidisc album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/discNumber
func (m_ MediaItem) DiscNumber() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("discNumber"))
	return rv
}

// The music or film genre of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/genre
func (m_ MediaItem) Genre() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("genre"))
	return rv
}

// The persistent identifier for a genre.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/genrePersistentID
func (m_ MediaItem) GenrePersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("genrePersistentID"))
	return rv
}

// A Boolean value that indicates whether the media item has a protected asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/hasProtectedAsset
func (m_ MediaItem) ProtectedAsset() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("protectedAsset"))
	return rv
}

// A Boolean value that indicates whether the media item is an iCloud Music Library item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isCloudItem
func (m_ MediaItem) CloudItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cloudItem"))
	return rv
}

// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isCompilation
func (m_ MediaItem) Compilation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("compilation"))
	return rv
}

// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isExplicitItem
func (m_ MediaItem) ExplicitItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("explicitItem"))
	return rv
}

// A Boolean value that indicates whether the media item is a preorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isPreorder
func (m_ MediaItem) Preorder() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("preorder"))
	return rv
}

// The most recent play date of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/lastPlayedDate
func (m_ MediaItem) LastPlayedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("lastPlayedDate"))
	return rv
}

// The lyrics for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/lyrics
func (m_ MediaItem) Lyrics() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("lyrics"))
	return rv
}

// The media type of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/mediaType
func (m_ MediaItem) MediaType() MediaType {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}

// The persistent identifier for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/persistentID
func (m_ MediaItem) PersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("persistentID"))
	return rv
}

// The playback duration of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playbackDuration
func (m_ MediaItem) PlaybackDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("playbackDuration"))
	return rv
}

// The ID of a media item from the Apple Music catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playbackStoreID
func (m_ MediaItem) PlaybackStoreID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("playbackStoreID"))
	return rv
}

// The number of times the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playCount
func (m_ MediaItem) PlayCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("playCount"))
	return rv
}

// The persistent identifier for an audio podcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/podcastPersistentID
func (m_ MediaItem) PodcastPersistentID() MediaEntityPersistentID /* typedef */ {
	rv := objc.Send[uint64](m_.ID, objc.Sel("podcastPersistentID"))
	return rv
}

// The title of a podcast, such as , rather than the title of an individual episode of a podcast, such as “Episode 12: Another Cold Day at the Pole.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/podcastTitle
func (m_ MediaItem) PodcastTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("podcastTitle"))
	return rv
}

// The user-specified rating of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/rating
func (m_ MediaItem) Rating() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rating"))
	return rv
}

// The date of the media item’s first public release.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/releaseDate
func (m_ MediaItem) ReleaseDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("releaseDate"))
	return rv
}

// The number of times the user skips playing the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/skipCount
func (m_ MediaItem) SkipCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("skipCount"))
	return rv
}

// The title or name of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/title
func (m_ MediaItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}

// Grouping information for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/userGrouping
func (m_ MediaItem) UserGrouping() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userGrouping"))
	return rv
}





