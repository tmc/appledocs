// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INMediaSearch] class.
var (
	INMediaSearchClass     _INMediaSearchClass
	INMediaSearchClassOnce sync.Once
)

func getINMediaSearchClass() _INMediaSearchClass {
	INMediaSearchClassOnce.Do(func() {
		INMediaSearchClass = _INMediaSearchClass{objc.GetClass("INMediaSearch")}
	})
	return INMediaSearchClass
}

type _INMediaSearchClass struct {
	class objc.Class
}

// An interface definition for the [INMediaSearch] class.
type IINMediaSearch interface {
	objectivec.IObject
	ActivityNames() string
	SetActivityNames(value string)
	AlbumName() string
	SetAlbumName(value string)
	ArtistName() string
	SetArtistName(value string)
	GenreNames() string
	SetGenreNames(value string)
	MediaIdentifier() string
	SetMediaIdentifier(value string)
	MediaName() string
	SetMediaName(value string)
	MediaType() unsafe.Pointer
	SetMediaType(value unsafe.Pointer)
	MoodNames() string
	SetMoodNames(value string)
	Reference() unsafe.Pointer
	SetReference(value unsafe.Pointer)
	ReleaseDate() INDateComponentsRange
	SetReleaseDate(value INDateComponentsRange)
	SortOrder() unsafe.Pointer
	SetSortOrder(value unsafe.Pointer)
}

// An object that describes a media type to search for, such as a station name, song name, or album name.

// An object that describes a media type to search for, such as a station name, song name, or album name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMediaSearch
type INMediaSearch struct {
	objectivec.Object
}

// INMediaSearchFrom constructs a [INMediaSearch] from an unsafe.Pointer.
//
// An object that describes a media type to search for, such as a station name, song name, or album name.
func INMediaSearchFrom(ptr unsafe.Pointer) INMediaSearch {
	return INMediaSearch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INMediaSearchClass) Alloc() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INMediaSearchClass) New() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INMediaSearch) Init() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INMediaSearch) Autorelease() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINMediaSearch creates a new INMediaSearch instance.
func NewINMediaSearch() INMediaSearch {
	return getINMediaSearchClass().New()
}

// The activity names to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/activitynames
func (i_ INMediaSearch) ActivityNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("activityNames"))
	return rv
}

// The activity names to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/activitynames
func (i_ INMediaSearch) SetActivityNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setActivityNames:"), objc.String(value))
}

// The name of the album to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/albumname
func (i_ INMediaSearch) AlbumName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("albumName"))
	return rv
}

// The name of the album to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/albumname
func (i_ INMediaSearch) SetAlbumName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlbumName:"), objc.String(value))
}

// The name of the artist to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/artistname
func (i_ INMediaSearch) ArtistName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("artistName"))
	return rv
}

// The name of the artist to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/artistname
func (i_ INMediaSearch) SetArtistName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArtistName:"), objc.String(value))
}

// The media genres to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/genrenames
func (i_ INMediaSearch) GenreNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("genreNames"))
	return rv
}

// The media genres to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/genrenames
func (i_ INMediaSearch) SetGenreNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGenreNames:"), objc.String(value))
}

// The unique media identifier to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediaidentifier
func (i_ INMediaSearch) MediaIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("mediaIdentifier"))
	return rv
}

// The unique media identifier to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediaidentifier
func (i_ INMediaSearch) SetMediaIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaIdentifier:"), objc.String(value))
}

// The name of the media to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/medianame
func (i_ INMediaSearch) MediaName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("mediaName"))
	return rv
}

// The name of the media to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/medianame
func (i_ INMediaSearch) SetMediaName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaName:"), objc.String(value))
}

// The type of media to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediatype
func (i_ INMediaSearch) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaType"))
	return rv
}

// The type of media to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediatype
func (i_ INMediaSearch) SetMediaType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaType:"), value)
}

// The moods to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/moodnames
func (i_ INMediaSearch) MoodNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("moodNames"))
	return rv
}

// The moods to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/moodnames
func (i_ INMediaSearch) SetMoodNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMoodNames:"), objc.String(value))
}

// A reference for the media item to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/reference
func (i_ INMediaSearch) Reference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reference"))
	return rv
}

// A reference for the media item to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/reference
func (i_ INMediaSearch) SetReference(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReference:"), value)
}

// The release date to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/releasedate
func (i_ INMediaSearch) ReleaseDate() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("releaseDate"))
	return rv
}

// The release date to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/releasedate
func (i_ INMediaSearch) SetReleaseDate(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReleaseDate:"), value)
}

// The sort order for the found media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/sortorder
func (i_ INMediaSearch) SortOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("sortOrder"))
	return rv
}

// The sort order for the found media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/sortorder
func (i_ INMediaSearch) SetSortOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSortOrder:"), value)
}
