// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaItem] class.
var (
	MediaItemClass     _MediaItemClass
	MediaItemClassOnce sync.Once
)

func getMediaItemClass() _MediaItemClass {
	MediaItemClassOnce.Do(func() {
		MediaItemClass = _MediaItemClass{objc.GetClass("MPMediaItem")}
	})
	return MediaItemClass
}

type _MediaItemClass struct {
	class objc.Class
}

// An interface definition for the [MediaItem] class.
type IMediaItem interface {
	IMediaEntity
}

// A collection of properties that represents a single item in the media library.
//
// A media item has an overall unique identifier, accessed using the property key, as well as specific identifiers for its metadata. These identifiers persists across application launches. A media item can have a wide range of metadata associated with it. You access this metadata using the method along with the property keys described in this document. You can also access metadata in a batch fashion using the method. Anytime the app accesses more than one property, enumerating over a set of property keys is more efficient than fetching each individual property. defines both of these methods, the abstract superclass of , and described in . You use attributes of media items to build media queries for searching the Media library. , , and describe these attributes. In addition, describes the property, and describes media queries.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem
type MediaItem struct {
	MediaEntity
}

// MediaItemFrom constructs a [MediaItem] from an unsafe.Pointer.
//
// A collection of properties that represents a single item in the media library.
func MediaItemFrom(ptr unsafe.Pointer) MediaItem {
	return MediaItem{
		MediaEntity: MediaEntityFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaItemClass) Alloc() MediaItem {
	rv := objc.Send[MediaItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaItemClass) New() MediaItem {
	rv := objc.Send[MediaItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaItem) Init() MediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaItem) Autorelease() MediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaItem creates a new MediaItem instance.
func NewMediaItem() MediaItem {
	return getMediaItemClass().New()
}


// Obtains the persistent identifier key for a specified grouping type.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/persistentIDProperty(forGroupingType:)
func (mc _MediaItemClass) PersistentIDPropertyForGroupingType(groupingType unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(mc.class), objc.Sel("persistentIDPropertyForGroupingType:"), groupingType)
	return rv
}

// Obtains the title key for a specified grouping type.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/titleProperty(forGroupingType:)
func (mc _MediaItemClass) TitlePropertyForGroupingType(groupingType unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(mc.class), objc.Sel("titlePropertyForGroupingType:"), groupingType)
	return rv
}

// The primary performing artist for an album.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumArtist
func (m_ MediaItem) AlbumArtist() string {
	rv := objc.Send[string](m_.ID, objc.Sel("albumArtist"))
	return rv
}

// The persistent identifier for the primary performing artist for an album.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumArtistPersistentID
func (m_ MediaItem) AlbumArtistPersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("albumArtistPersistentID"))
	return rv
}

// The persistent identifier for an album.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumPersistentID
func (m_ MediaItem) AlbumPersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("albumPersistentID"))
	return rv
}

// The title of an album, such as , rather than the title of an individual song on the album, such as “Crater Dance.”
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTitle
func (m_ MediaItem) AlbumTitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("albumTitle"))
	return rv
}

// The number of tracks for the album that contains the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTrackCount
func (m_ MediaItem) AlbumTrackCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("albumTrackCount"))
	return rv
}

// The track number of the media item, for a media item that’s part of an album.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/albumTrackNumber
func (m_ MediaItem) AlbumTrackNumber() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("albumTrackNumber"))
	return rv
}

// The performing artists for a media item, which may vary from the primary artist for the album that a media item belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artist
func (m_ MediaItem) Artist() string {
	rv := objc.Send[string](m_.ID, objc.Sel("artist"))
	return rv
}

// The persistent identifier for an artist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artistPersistentID
func (m_ MediaItem) ArtistPersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("artistPersistentID"))
	return rv
}

// The artwork image for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/artwork
func (m_ MediaItem) Artwork() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("artwork"))
	return rv
}

// The URL that points to the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/assetURL
func (m_ MediaItem) AssetURL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("assetURL"))
	return rv
}

// The number of musical beats per minute for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/beatsPerMinute
func (m_ MediaItem) BeatsPerMinute() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("beatsPerMinute"))
	return rv
}

// The time of the user’s most recent interaction with the bookmark in the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/bookmarkTime
func (m_ MediaItem) BookmarkTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("bookmarkTime"))
	return rv
}

// Textual information about the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/comments
func (m_ MediaItem) Comments() string {
	rv := objc.Send[string](m_.ID, objc.Sel("comments"))
	return rv
}

// The persistent identifier for a composer.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/composerPersistentID
func (m_ MediaItem) ComposerPersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("composerPersistentID"))
	return rv
}

// The date the user adds the media item to the library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/dateAdded
func (m_ MediaItem) DateAdded() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dateAdded"))
	return rv
}

// The number of discs for the album that contains the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/discCount
func (m_ MediaItem) DiscCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("discCount"))
	return rv
}

// The disc number of the media item, for a media item that’s part of a multidisc album.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/discNumber
func (m_ MediaItem) DiscNumber() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("discNumber"))
	return rv
}

// The music or film genre of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/genre
func (m_ MediaItem) Genre() string {
	rv := objc.Send[string](m_.ID, objc.Sel("genre"))
	return rv
}

// The persistent identifier for a genre.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/genrePersistentID
func (m_ MediaItem) GenrePersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("genrePersistentID"))
	return rv
}

// A Boolean value that indicates whether the media item has a protected asset.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/hasProtectedAsset
func (m_ MediaItem) ProtectedAsset() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("protectedAsset"))
	return rv
}

// A Boolean value that indicates whether the media item is an iCloud Music Library item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isCloudItem
func (m_ MediaItem) CloudItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cloudItem"))
	return rv
}

// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isCompilation
func (m_ MediaItem) Compilation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("compilation"))
	return rv
}

// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isExplicitItem
func (m_ MediaItem) ExplicitItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("explicitItem"))
	return rv
}

// A Boolean value that indicates whether the media item is a preorder.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/isPreorder
func (m_ MediaItem) Preorder() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("preorder"))
	return rv
}

// The most recent play date of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/lastPlayedDate
func (m_ MediaItem) LastPlayedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lastPlayedDate"))
	return rv
}

// The lyrics for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/lyrics
func (m_ MediaItem) Lyrics() string {
	rv := objc.Send[string](m_.ID, objc.Sel("lyrics"))
	return rv
}

// The media type of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/mediaType
func (m_ MediaItem) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaType"))
	return rv
}

// The persistent identifier for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/persistentID
func (m_ MediaItem) PersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("persistentID"))
	return rv
}

// The number of times the user plays the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playCount
func (m_ MediaItem) PlayCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("playCount"))
	return rv
}

// The playback duration of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playbackDuration
func (m_ MediaItem) PlaybackDuration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("playbackDuration"))
	return rv
}

// The ID of a media item from the Apple Music catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/playbackStoreID
func (m_ MediaItem) PlaybackStoreID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("playbackStoreID"))
	return rv
}

// The persistent identifier for an audio podcast.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/podcastPersistentID
func (m_ MediaItem) PodcastPersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("podcastPersistentID"))
	return rv
}

// The title of a podcast, such as , rather than the title of an individual episode of a podcast, such as “Episode 12: Another Cold Day at the Pole.”
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/podcastTitle
func (m_ MediaItem) PodcastTitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("podcastTitle"))
	return rv
}

// The user-specified rating of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/rating
func (m_ MediaItem) Rating() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rating"))
	return rv
}

// The date of the media item’s first public release.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/releaseDate
func (m_ MediaItem) ReleaseDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("releaseDate"))
	return rv
}

// The number of times the user skips playing the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/skipCount
func (m_ MediaItem) SkipCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("skipCount"))
	return rv
}

// The title or name of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/title
func (m_ MediaItem) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}

// Grouping information for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/userGrouping
func (m_ MediaItem) UserGrouping() string {
	rv := objc.Send[string](m_.ID, objc.Sel("userGrouping"))
	return rv
}

// The persistent identifier for a media entity.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaentitypropertypersistentid
func (m_ MediaItem) MPMediaEntityPropertyPersistentID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MPMediaEntityPropertyPersistentID"))
	return rv
}

// The musical composer for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/composer
func (m_ MediaItem) Composer() string {
	rv := objc.Send[string](m_.ID, objc.Sel("composer"))
	return rv
}


// SetComposer sets the value of the composer property.
// The musical composer for the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/composer
func (m_ MediaItem) SetComposer(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComposer:"), objc.String(value))
}

// A Boolean value that indicates whether the media item has a protected asset.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/hasprotectedasset
func (m_ MediaItem) HasProtectedAsset() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasProtectedAsset"))
	return rv
}


// SetHasProtectedAsset sets the value of the hasProtectedAsset property.
// A Boolean value that indicates whether the media item has a protected asset.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/hasprotectedasset
func (m_ MediaItem) SetHasProtectedAsset(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasProtectedAsset:"), value)
}

// A Boolean value that indicates whether the media item is an iCloud Music Library item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isclouditem
func (m_ MediaItem) IsCloudItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCloudItem"))
	return rv
}


// SetIsCloudItem sets the value of the isCloudItem property.
// A Boolean value that indicates whether the media item is an iCloud Music Library item.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isclouditem
func (m_ MediaItem) SetIsCloudItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCloudItem:"), value)
}

// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/iscompilation
func (m_ MediaItem) IsCompilation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCompilation"))
	return rv
}


// SetIsCompilation sets the value of the isCompilation property.
// A Boolean value that indicates whether the media item is part of a compilation.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/iscompilation
func (m_ MediaItem) SetIsCompilation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCompilation:"), value)
}

// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isexplicititem
func (m_ MediaItem) IsExplicitItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isExplicitItem"))
	return rv
}


// SetIsExplicitItem sets the value of the isExplicitItem property.
// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isexplicititem
func (m_ MediaItem) SetIsExplicitItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsExplicitItem:"), value)
}

// A Boolean value that indicates whether the media item is a preorder.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/ispreorder
func (m_ MediaItem) IsPreorder() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPreorder"))
	return rv
}


// SetIsPreorder sets the value of the isPreorder property.
// A Boolean value that indicates whether the media item is a preorder.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/ispreorder
func (m_ MediaItem) SetIsPreorder(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPreorder:"), value)
}

// The key for the persistent identifier for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertypersistentid
func (m_ MediaItem) MPMediaItemPropertyPersistentID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MPMediaItemPropertyPersistentID"))
	return rv
}



